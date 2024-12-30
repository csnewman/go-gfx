package gfx

import (
	"errors"
	"unsafe"
)

/*
#include "vulkan.h"
*/
import "C"

var ErrNotMapped = errors.New("not mapped")

type Buffer struct {
	graphics   *Graphics
	buffer     C.VkBuffer
	allocation C.VmaAllocation
	mappedPtr  unsafe.Pointer
	mapped     int
	Size       int
}

type BufferUsage int

const (
	BufferUsageHostRandomAccess BufferUsage = 1 << iota
	BufferUsageHostUpload
	BufferUsagePersistentMap
)

type BufferDescriptor struct {
	Size  int
	Usage BufferUsage
}

func (g *Graphics) CreateBuffer(des BufferDescriptor) (*Buffer, error) {
	var createInfo C.VkBufferCreateInfo
	createInfo.sType = C.VK_STRUCTURE_TYPE_BUFFER_CREATE_INFO
	createInfo.size = C.VkDeviceSize(des.Size)

	// For now enable most common usage types. Supposedly does not affect performance on most hardware.
	createInfo.usage = C.VK_BUFFER_USAGE_TRANSFER_SRC_BIT | C.VK_BUFFER_USAGE_TRANSFER_DST_BIT |
		C.VK_BUFFER_USAGE_UNIFORM_BUFFER_BIT | C.VK_BUFFER_USAGE_STORAGE_BUFFER_BIT |
		C.VK_BUFFER_USAGE_INDEX_BUFFER_BIT | C.VK_BUFFER_USAGE_VERTEX_BUFFER_BIT |
		C.VK_BUFFER_USAGE_SHADER_DEVICE_ADDRESS_BIT
	createInfo.sharingMode = C.VK_SHARING_MODE_EXCLUSIVE

	var allocCreateInfo C.VmaAllocationCreateInfo
	allocCreateInfo.usage = C.VMA_MEMORY_USAGE_AUTO

	if des.Usage&BufferUsageHostRandomAccess != 0 {
		allocCreateInfo.flags |= C.VMA_ALLOCATION_CREATE_HOST_ACCESS_RANDOM_BIT
	}

	if des.Usage&BufferUsageHostUpload != 0 {
		allocCreateInfo.flags |= C.VMA_ALLOCATION_CREATE_HOST_ACCESS_SEQUENTIAL_WRITE_BIT
	}

	if des.Usage&BufferUsagePersistentMap != 0 {
		allocCreateInfo.flags |= C.VMA_ALLOCATION_CREATE_MAPPED_BIT
	}

	var (
		buffer     C.VkBuffer
		allocation C.VmaAllocation
		allocInfo  C.VmaAllocationInfo
	)

	if err := mapError(C.vmaCreateBuffer(
		g.memoryAllocator,
		&createInfo,
		&allocCreateInfo,
		&buffer,
		&allocation,
		&allocInfo,
	)); err != nil {
		return nil, err
	}

	b := &Buffer{
		graphics:   g,
		buffer:     buffer,
		allocation: allocation,
		mappedPtr:  unsafe.Pointer(allocInfo.pMappedData),
		Size:       des.Size,
	}

	if b.mappedPtr != nil {
		b.mapped = -1
	}

	return b, nil
}

func (b *Buffer) Close() {
	C.vmaDestroyBuffer(b.graphics.memoryAllocator, b.buffer, b.allocation)
}

func (b *Buffer) MappedPtr() (unsafe.Pointer, bool) {
	return b.mappedPtr, true
}

func (b *Buffer) Map() (unsafe.Pointer, error) {
	if b.mapped == -1 {
		return b.mappedPtr, nil
	}

	if err := mapError(C.vmaMapMemory(b.graphics.memoryAllocator, b.allocation, &b.mappedPtr)); err != nil {
		return nil, err
	}

	b.mapped++

	return b.mappedPtr, nil
}

func (b *Buffer) Flush() error {
	if err := mapError(C.vmaFlushAllocation(b.graphics.memoryAllocator, b.allocation, 0, C.VK_WHOLE_SIZE)); err != nil {
		return err
	}

	return nil
}

func (b *Buffer) Unmap() {
	if b.mapped <= 0 {
		return
	}

	C.vmaUnmapMemory(b.graphics.memoryAllocator, b.allocation)

	b.mapped--
}

func (b *Buffer) CopyFrom(data []byte) error {
	if b.mapped == 0 {
		return ErrNotMapped
	}

	dst := unsafe.Slice((*byte)(b.mappedPtr), len(data))

	copy(dst, data)

	return nil
}
