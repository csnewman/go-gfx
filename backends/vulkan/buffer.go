package vulkan

import (
	"errors"
	"unsafe"

	"github.com/csnewman/go-gfx/ffi"
	"github.com/csnewman/go-gfx/gfx"
	"github.com/csnewman/go-gfx/vk"
	"github.com/csnewman/go-gfx/vma"
)

var ErrNotMapped = errors.New("not mapped")

type Buffer struct {
	graphics   *Graphics
	buffer     vk.Buffer
	allocation vma.Allocation
	mappedPtr  unsafe.Pointer
	mapped     int
	size       int
}

func (g *Graphics) CreateBuffer(des gfx.BufferDescriptor) (gfx.Buffer, error) {
	arena := ffi.NewArena()
	defer arena.Close()

	createInfo := vk.BufferCreateInfoAlloc(arena, 1)
	createInfo.SetSType(vk.STRUCTURE_TYPE_BUFFER_CREATE_INFO)
	createInfo.SetSize(vk.DeviceSize(des.Size))

	// For now enable most common usage types. Supposedly does not affect performance on most hardware.
	createInfo.SetUsage(vk.BUFFER_USAGE_TRANSFER_SRC_BIT | vk.BUFFER_USAGE_TRANSFER_DST_BIT |
		vk.BUFFER_USAGE_UNIFORM_BUFFER_BIT | vk.BUFFER_USAGE_STORAGE_BUFFER_BIT |
		vk.BUFFER_USAGE_INDEX_BUFFER_BIT | vk.BUFFER_USAGE_VERTEX_BUFFER_BIT |
		vk.BUFFER_USAGE_SHADER_DEVICE_ADDRESS_BIT)
	createInfo.SetSharingMode(vk.SHARING_MODE_EXCLUSIVE)

	allocCreateInfo := vma.AllocationCreateInfoAlloc(arena, 1)
	allocCreateInfo.SetUsage(vma.MEMORY_USAGE_AUTO)

	var flags vma.AllocationCreateFlags

	if des.Usage&gfx.BufferUsageHostRandomAccess != 0 {
		flags |= vma.ALLOCATION_CREATE_HOST_ACCESS_RANDOM_BIT
	}

	if des.Usage&gfx.BufferUsageHostUpload != 0 {
		flags |= vma.ALLOCATION_CREATE_HOST_ACCESS_SEQUENTIAL_WRITE_BIT
	}

	if des.Usage&gfx.BufferUsagePersistentMap != 0 {
		flags |= vma.ALLOCATION_CREATE_MAPPED_BIT
	}

	allocCreateInfo.SetFlags(flags)

	bufferRef := ffi.RefAlloc[vk.Buffer](arena, 1)
	allocationRef := ffi.RefAlloc[vma.Allocation](arena, 1)
	allocationInfo := vma.AllocationInfoAlloc(arena, 1)

	if err := mapError(vma.CreateBuffer(
		g.memoryAllocator,
		createInfo,
		allocCreateInfo,
		bufferRef,
		allocationRef,
		allocationInfo,
	)); err != nil {
		return nil, err
	}

	b := &Buffer{
		graphics:   g,
		buffer:     bufferRef.Get(),
		allocation: allocationRef.Get(),
		mappedPtr:  unsafe.Pointer(allocationInfo.GetPMappedData()),
		size:       des.Size,
	}

	if b.mappedPtr != nil {
		b.mapped = -1
	}

	return b, nil
}

func (b *Buffer) Close() {
	vma.DestroyBuffer(b.graphics.memoryAllocator, b.buffer, b.allocation)
}

func (b *Buffer) Size() int {
	return b.size
}

func (b *Buffer) MappedPtr() (unsafe.Pointer, bool) {
	return b.mappedPtr, true
}

func (b *Buffer) Map() (unsafe.Pointer, error) {
	if b.mapped == -1 {
		return b.mappedPtr, nil
	}

	arena := ffi.NewArena()
	defer arena.Close()

	ptrRef := ffi.RefAlloc[uintptr](arena, 1)

	if err := mapError(vma.MapMemory(b.graphics.memoryAllocator, b.allocation, ptrRef)); err != nil {
		return nil, err
	}

	b.mappedPtr = unsafe.Pointer(ptrRef.Get())

	b.mapped++

	return b.mappedPtr, nil
}

func (b *Buffer) Flush() error {
	if err := mapError(vma.FlushAllocation(b.graphics.memoryAllocator, b.allocation, 0, vk.DeviceSize(vk.WHOLE_SIZE))); err != nil {
		return err
	}

	return nil
}

func (b *Buffer) Unmap() {
	if b.mapped <= 0 {
		return
	}

	vma.UnmapMemory(b.graphics.memoryAllocator, b.allocation)

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
