package gfx

import (
	"unsafe"
)

/*
#include "vulkan.h"
*/
import "C"

type Buffer struct {
	buffer     C.VkBuffer
	allocation C.VmaAllocation
}

func (g *Graphics) CreateBuffer(data []byte) *Buffer {
	var createInfo C.VkBufferCreateInfo
	createInfo.sType = C.VK_STRUCTURE_TYPE_BUFFER_CREATE_INFO
	createInfo.size = C.VkDeviceSize(len(data))
	createInfo.usage = C.VK_BUFFER_USAGE_VERTEX_BUFFER_BIT
	createInfo.sharingMode = C.VK_SHARING_MODE_EXCLUSIVE

	var allocInfo C.VmaAllocationCreateInfo
	allocInfo.usage = C.VMA_MEMORY_USAGE_CPU_TO_GPU

	var buffer C.VkBuffer
	var allocation C.VmaAllocation

	if err := mapError(C.vmaCreateBuffer(
		g.memoryAllocator,
		&createInfo,
		&allocInfo,
		&buffer,
		&allocation,
		nil,
	)); err != nil {
		// TODO: handle
		panic(err)
	}

	var ptr unsafe.Pointer

	if err := mapError(C.vmaMapMemory(g.memoryAllocator, allocation, &ptr)); err != nil {
		panic(err)
	}

	dst := unsafe.Slice((*byte)(ptr), len(data))

	copy(dst, data)

	C.vmaUnmapMemory(g.memoryAllocator, allocation)

	return &Buffer{
		buffer:     buffer,
		allocation: allocation,
	}
}
