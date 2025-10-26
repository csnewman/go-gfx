package vma

import (
	ffi "github.com/csnewman/go-gfx/ffi"
	vk "github.com/csnewman/go-gfx/vk"
	"unsafe"
)

// #include "vma.h"
import "C"

// AllocateDedicatedMemory wraps vmaAllocateDedicatedMemory.
func AllocateDedicatedMemory(allocator Allocator, pVkMemoryRequirements vk.MemoryRequirements, pCreateInfo AllocationCreateInfo, pMemoryAllocateNext unsafe.Pointer, pAllocation ffi.Ref[Allocation], pAllocationInfo AllocationInfo) vk.Result {
	ret := C.vmaAllocateDedicatedMemory(C.VmaAllocator(unsafe.Pointer(allocator)), (*C.VkMemoryRequirements)(pVkMemoryRequirements.Raw()), pCreateInfo.ptr, pMemoryAllocateNext, (*C.VmaAllocation)(pAllocation.Raw()), pAllocationInfo.ptr)

	return vk.Result(ret)
}

// AllocateMemory wraps vmaAllocateMemory.
func AllocateMemory(allocator Allocator, pVkMemoryRequirements vk.MemoryRequirements, pCreateInfo AllocationCreateInfo, pAllocation ffi.Ref[Allocation], pAllocationInfo AllocationInfo) vk.Result {
	ret := C.vmaAllocateMemory(C.VmaAllocator(unsafe.Pointer(allocator)), (*C.VkMemoryRequirements)(pVkMemoryRequirements.Raw()), pCreateInfo.ptr, (*C.VmaAllocation)(pAllocation.Raw()), pAllocationInfo.ptr)

	return vk.Result(ret)
}

// AllocateMemoryForBuffer wraps vmaAllocateMemoryForBuffer.
func AllocateMemoryForBuffer(allocator Allocator, buffer vk.Buffer, pCreateInfo AllocationCreateInfo, pAllocation ffi.Ref[Allocation], pAllocationInfo AllocationInfo) vk.Result {
	ret := C.vmaAllocateMemoryForBuffer(C.VmaAllocator(unsafe.Pointer(allocator)), C.VkBuffer(buffer), pCreateInfo.ptr, (*C.VmaAllocation)(pAllocation.Raw()), pAllocationInfo.ptr)

	return vk.Result(ret)
}

// AllocateMemoryForImage wraps vmaAllocateMemoryForImage.
func AllocateMemoryForImage(allocator Allocator, image vk.Image, pCreateInfo AllocationCreateInfo, pAllocation ffi.Ref[Allocation], pAllocationInfo AllocationInfo) vk.Result {
	ret := C.vmaAllocateMemoryForImage(C.VmaAllocator(unsafe.Pointer(allocator)), C.VkImage(image), pCreateInfo.ptr, (*C.VmaAllocation)(pAllocation.Raw()), pAllocationInfo.ptr)

	return vk.Result(ret)
}

// AllocateMemoryPages wraps vmaAllocateMemoryPages.
func AllocateMemoryPages(allocator Allocator, pVkMemoryRequirements vk.MemoryRequirements, pCreateInfo AllocationCreateInfo, allocationCount uintptr, pAllocations ffi.Ref[Allocation], pAllocationInfo AllocationInfo) vk.Result {
	ret := C.vmaAllocateMemoryPages(C.VmaAllocator(unsafe.Pointer(allocator)), (*C.VkMemoryRequirements)(pVkMemoryRequirements.Raw()), pCreateInfo.ptr, C.size_t(allocationCount), (*C.VmaAllocation)(pAllocations.Raw()), pAllocationInfo.ptr)

	return vk.Result(ret)
}

// BeginDefragmentation wraps vmaBeginDefragmentation.
func BeginDefragmentation(allocator Allocator, pInfo DefragmentationInfo, pContext ffi.Ref[DefragmentationContext]) vk.Result {
	ret := C.vmaBeginDefragmentation(C.VmaAllocator(unsafe.Pointer(allocator)), pInfo.ptr, (*C.VmaDefragmentationContext)(pContext.Raw()))

	return vk.Result(ret)
}

// BeginDefragmentationPass wraps vmaBeginDefragmentationPass.
func BeginDefragmentationPass(allocator Allocator, context DefragmentationContext, pPassInfo DefragmentationPassMoveInfo) vk.Result {
	ret := C.vmaBeginDefragmentationPass(C.VmaAllocator(unsafe.Pointer(allocator)), C.VmaDefragmentationContext(unsafe.Pointer(context)), pPassInfo.ptr)

	return vk.Result(ret)
}

// BindBufferMemory wraps vmaBindBufferMemory.
func BindBufferMemory(allocator Allocator, allocation Allocation, buffer vk.Buffer) vk.Result {
	ret := C.vmaBindBufferMemory(C.VmaAllocator(unsafe.Pointer(allocator)), C.VmaAllocation(unsafe.Pointer(allocation)), C.VkBuffer(buffer))

	return vk.Result(ret)
}

// BindBufferMemory2 wraps vmaBindBufferMemory2.
func BindBufferMemory2(allocator Allocator, allocation Allocation, allocationLocalOffset vk.DeviceSize, buffer vk.Buffer, pNext unsafe.Pointer) vk.Result {
	ret := C.vmaBindBufferMemory2(C.VmaAllocator(unsafe.Pointer(allocator)), C.VmaAllocation(unsafe.Pointer(allocation)), C.VkDeviceSize(allocationLocalOffset), C.VkBuffer(buffer), pNext)

	return vk.Result(ret)
}

// BindImageMemory wraps vmaBindImageMemory.
func BindImageMemory(allocator Allocator, allocation Allocation, image vk.Image) vk.Result {
	ret := C.vmaBindImageMemory(C.VmaAllocator(unsafe.Pointer(allocator)), C.VmaAllocation(unsafe.Pointer(allocation)), C.VkImage(image))

	return vk.Result(ret)
}

// BindImageMemory2 wraps vmaBindImageMemory2.
func BindImageMemory2(allocator Allocator, allocation Allocation, allocationLocalOffset vk.DeviceSize, image vk.Image, pNext unsafe.Pointer) vk.Result {
	ret := C.vmaBindImageMemory2(C.VmaAllocator(unsafe.Pointer(allocator)), C.VmaAllocation(unsafe.Pointer(allocation)), C.VkDeviceSize(allocationLocalOffset), C.VkImage(image), pNext)

	return vk.Result(ret)
}

// vmaBuildStatsString.ppStatsString is unsupported: param: unknown pointer2 type char.

// vmaBuildVirtualBlockStatsString.ppStatsString is unsupported: param: unknown pointer2 type char.

// CalculatePoolStatistics wraps vmaCalculatePoolStatistics.
func CalculatePoolStatistics(allocator Allocator, pool Pool, pPoolStats DetailedStatistics) {
	C.vmaCalculatePoolStatistics(C.VmaAllocator(unsafe.Pointer(allocator)), C.VmaPool(unsafe.Pointer(pool)), pPoolStats.ptr)
}

// CalculateStatistics wraps vmaCalculateStatistics.
func CalculateStatistics(allocator Allocator, pStats TotalStatistics) {
	C.vmaCalculateStatistics(C.VmaAllocator(unsafe.Pointer(allocator)), pStats.ptr)
}

// CalculateVirtualBlockStatistics wraps vmaCalculateVirtualBlockStatistics.
func CalculateVirtualBlockStatistics(virtualBlock VirtualBlock, pStats DetailedStatistics) {
	C.vmaCalculateVirtualBlockStatistics(C.VmaVirtualBlock(unsafe.Pointer(virtualBlock)), pStats.ptr)
}

// CheckCorruption wraps vmaCheckCorruption.
func CheckCorruption(allocator Allocator, memoryTypeBits uint32) vk.Result {
	ret := C.vmaCheckCorruption(C.VmaAllocator(unsafe.Pointer(allocator)), C.uint32_t(memoryTypeBits))

	return vk.Result(ret)
}

// CheckPoolCorruption wraps vmaCheckPoolCorruption.
func CheckPoolCorruption(allocator Allocator, pool Pool) vk.Result {
	ret := C.vmaCheckPoolCorruption(C.VmaAllocator(unsafe.Pointer(allocator)), C.VmaPool(unsafe.Pointer(pool)))

	return vk.Result(ret)
}

// ClearVirtualBlock wraps vmaClearVirtualBlock.
func ClearVirtualBlock(virtualBlock VirtualBlock) {
	C.vmaClearVirtualBlock(C.VmaVirtualBlock(unsafe.Pointer(virtualBlock)))
}

// CopyAllocationToMemory wraps vmaCopyAllocationToMemory.
func CopyAllocationToMemory(allocator Allocator, srcAllocation Allocation, srcAllocationLocalOffset vk.DeviceSize, pDstHostPointer unsafe.Pointer, size vk.DeviceSize) vk.Result {
	ret := C.vmaCopyAllocationToMemory(C.VmaAllocator(unsafe.Pointer(allocator)), C.VmaAllocation(unsafe.Pointer(srcAllocation)), C.VkDeviceSize(srcAllocationLocalOffset), pDstHostPointer, C.VkDeviceSize(size))

	return vk.Result(ret)
}

// CopyMemoryToAllocation wraps vmaCopyMemoryToAllocation.
func CopyMemoryToAllocation(allocator Allocator, pSrcHostPointer unsafe.Pointer, dstAllocation Allocation, dstAllocationLocalOffset vk.DeviceSize, size vk.DeviceSize) vk.Result {
	ret := C.vmaCopyMemoryToAllocation(C.VmaAllocator(unsafe.Pointer(allocator)), pSrcHostPointer, C.VmaAllocation(unsafe.Pointer(dstAllocation)), C.VkDeviceSize(dstAllocationLocalOffset), C.VkDeviceSize(size))

	return vk.Result(ret)
}

// CreateAliasingBuffer wraps vmaCreateAliasingBuffer.
func CreateAliasingBuffer(allocator Allocator, allocation Allocation, pBufferCreateInfo vk.BufferCreateInfo, pBuffer ffi.Ref[vk.Buffer]) vk.Result {
	ret := C.vmaCreateAliasingBuffer(C.VmaAllocator(unsafe.Pointer(allocator)), C.VmaAllocation(unsafe.Pointer(allocation)), (*C.VkBufferCreateInfo)(pBufferCreateInfo.Raw()), (*C.VkBuffer)(pBuffer.Raw()))

	return vk.Result(ret)
}

// CreateAliasingBuffer2 wraps vmaCreateAliasingBuffer2.
func CreateAliasingBuffer2(allocator Allocator, allocation Allocation, allocationLocalOffset vk.DeviceSize, pBufferCreateInfo vk.BufferCreateInfo, pBuffer ffi.Ref[vk.Buffer]) vk.Result {
	ret := C.vmaCreateAliasingBuffer2(C.VmaAllocator(unsafe.Pointer(allocator)), C.VmaAllocation(unsafe.Pointer(allocation)), C.VkDeviceSize(allocationLocalOffset), (*C.VkBufferCreateInfo)(pBufferCreateInfo.Raw()), (*C.VkBuffer)(pBuffer.Raw()))

	return vk.Result(ret)
}

// CreateAliasingImage wraps vmaCreateAliasingImage.
func CreateAliasingImage(allocator Allocator, allocation Allocation, pImageCreateInfo vk.ImageCreateInfo, pImage ffi.Ref[vk.Image]) vk.Result {
	ret := C.vmaCreateAliasingImage(C.VmaAllocator(unsafe.Pointer(allocator)), C.VmaAllocation(unsafe.Pointer(allocation)), (*C.VkImageCreateInfo)(pImageCreateInfo.Raw()), (*C.VkImage)(pImage.Raw()))

	return vk.Result(ret)
}

// CreateAliasingImage2 wraps vmaCreateAliasingImage2.
func CreateAliasingImage2(allocator Allocator, allocation Allocation, allocationLocalOffset vk.DeviceSize, pImageCreateInfo vk.ImageCreateInfo, pImage ffi.Ref[vk.Image]) vk.Result {
	ret := C.vmaCreateAliasingImage2(C.VmaAllocator(unsafe.Pointer(allocator)), C.VmaAllocation(unsafe.Pointer(allocation)), C.VkDeviceSize(allocationLocalOffset), (*C.VkImageCreateInfo)(pImageCreateInfo.Raw()), (*C.VkImage)(pImage.Raw()))

	return vk.Result(ret)
}

// CreateAllocator wraps vmaCreateAllocator.
func CreateAllocator(pCreateInfo AllocatorCreateInfo, pAllocator ffi.Ref[Allocator]) vk.Result {
	ret := C.vmaCreateAllocator(pCreateInfo.ptr, (*C.VmaAllocator)(pAllocator.Raw()))

	return vk.Result(ret)
}

// CreateBuffer wraps vmaCreateBuffer.
func CreateBuffer(allocator Allocator, pBufferCreateInfo vk.BufferCreateInfo, pAllocationCreateInfo AllocationCreateInfo, pBuffer ffi.Ref[vk.Buffer], pAllocation ffi.Ref[Allocation], pAllocationInfo AllocationInfo) vk.Result {
	ret := C.vmaCreateBuffer(C.VmaAllocator(unsafe.Pointer(allocator)), (*C.VkBufferCreateInfo)(pBufferCreateInfo.Raw()), pAllocationCreateInfo.ptr, (*C.VkBuffer)(pBuffer.Raw()), (*C.VmaAllocation)(pAllocation.Raw()), pAllocationInfo.ptr)

	return vk.Result(ret)
}

// CreateBufferWithAlignment wraps vmaCreateBufferWithAlignment.
func CreateBufferWithAlignment(allocator Allocator, pBufferCreateInfo vk.BufferCreateInfo, pAllocationCreateInfo AllocationCreateInfo, minAlignment vk.DeviceSize, pBuffer ffi.Ref[vk.Buffer], pAllocation ffi.Ref[Allocation], pAllocationInfo AllocationInfo) vk.Result {
	ret := C.vmaCreateBufferWithAlignment(C.VmaAllocator(unsafe.Pointer(allocator)), (*C.VkBufferCreateInfo)(pBufferCreateInfo.Raw()), pAllocationCreateInfo.ptr, C.VkDeviceSize(minAlignment), (*C.VkBuffer)(pBuffer.Raw()), (*C.VmaAllocation)(pAllocation.Raw()), pAllocationInfo.ptr)

	return vk.Result(ret)
}

// CreateDedicatedBuffer wraps vmaCreateDedicatedBuffer.
func CreateDedicatedBuffer(allocator Allocator, pBufferCreateInfo vk.BufferCreateInfo, pAllocationCreateInfo AllocationCreateInfo, pMemoryAllocateNext unsafe.Pointer, pBuffer ffi.Ref[vk.Buffer], pAllocation ffi.Ref[Allocation], pAllocationInfo AllocationInfo) vk.Result {
	ret := C.vmaCreateDedicatedBuffer(C.VmaAllocator(unsafe.Pointer(allocator)), (*C.VkBufferCreateInfo)(pBufferCreateInfo.Raw()), pAllocationCreateInfo.ptr, pMemoryAllocateNext, (*C.VkBuffer)(pBuffer.Raw()), (*C.VmaAllocation)(pAllocation.Raw()), pAllocationInfo.ptr)

	return vk.Result(ret)
}

// CreateDedicatedImage wraps vmaCreateDedicatedImage.
func CreateDedicatedImage(allocator Allocator, pImageCreateInfo vk.ImageCreateInfo, pAllocationCreateInfo AllocationCreateInfo, pMemoryAllocateNext unsafe.Pointer, pImage ffi.Ref[vk.Image], pAllocation ffi.Ref[Allocation], pAllocationInfo AllocationInfo) vk.Result {
	ret := C.vmaCreateDedicatedImage(C.VmaAllocator(unsafe.Pointer(allocator)), (*C.VkImageCreateInfo)(pImageCreateInfo.Raw()), pAllocationCreateInfo.ptr, pMemoryAllocateNext, (*C.VkImage)(pImage.Raw()), (*C.VmaAllocation)(pAllocation.Raw()), pAllocationInfo.ptr)

	return vk.Result(ret)
}

// CreateImage wraps vmaCreateImage.
func CreateImage(allocator Allocator, pImageCreateInfo vk.ImageCreateInfo, pAllocationCreateInfo AllocationCreateInfo, pImage ffi.Ref[vk.Image], pAllocation ffi.Ref[Allocation], pAllocationInfo AllocationInfo) vk.Result {
	ret := C.vmaCreateImage(C.VmaAllocator(unsafe.Pointer(allocator)), (*C.VkImageCreateInfo)(pImageCreateInfo.Raw()), pAllocationCreateInfo.ptr, (*C.VkImage)(pImage.Raw()), (*C.VmaAllocation)(pAllocation.Raw()), pAllocationInfo.ptr)

	return vk.Result(ret)
}

// CreatePool wraps vmaCreatePool.
func CreatePool(allocator Allocator, pCreateInfo PoolCreateInfo, pPool ffi.Ref[Pool]) vk.Result {
	ret := C.vmaCreatePool(C.VmaAllocator(unsafe.Pointer(allocator)), pCreateInfo.ptr, (*C.VmaPool)(pPool.Raw()))

	return vk.Result(ret)
}

// CreateVirtualBlock wraps vmaCreateVirtualBlock.
func CreateVirtualBlock(pCreateInfo VirtualBlockCreateInfo, pVirtualBlock ffi.Ref[VirtualBlock]) vk.Result {
	ret := C.vmaCreateVirtualBlock(pCreateInfo.ptr, (*C.VmaVirtualBlock)(pVirtualBlock.Raw()))

	return vk.Result(ret)
}

// DestroyAllocator wraps vmaDestroyAllocator.
func DestroyAllocator(allocator Allocator) {
	C.vmaDestroyAllocator(C.VmaAllocator(unsafe.Pointer(allocator)))
}

// DestroyBuffer wraps vmaDestroyBuffer.
func DestroyBuffer(allocator Allocator, buffer vk.Buffer, allocation Allocation) {
	C.vmaDestroyBuffer(C.VmaAllocator(unsafe.Pointer(allocator)), C.VkBuffer(buffer), C.VmaAllocation(unsafe.Pointer(allocation)))
}

// DestroyImage wraps vmaDestroyImage.
func DestroyImage(allocator Allocator, image vk.Image, allocation Allocation) {
	C.vmaDestroyImage(C.VmaAllocator(unsafe.Pointer(allocator)), C.VkImage(image), C.VmaAllocation(unsafe.Pointer(allocation)))
}

// DestroyPool wraps vmaDestroyPool.
func DestroyPool(allocator Allocator, pool Pool) {
	C.vmaDestroyPool(C.VmaAllocator(unsafe.Pointer(allocator)), C.VmaPool(unsafe.Pointer(pool)))
}

// DestroyVirtualBlock wraps vmaDestroyVirtualBlock.
func DestroyVirtualBlock(virtualBlock VirtualBlock) {
	C.vmaDestroyVirtualBlock(C.VmaVirtualBlock(unsafe.Pointer(virtualBlock)))
}

// EndDefragmentation wraps vmaEndDefragmentation.
func EndDefragmentation(allocator Allocator, context DefragmentationContext, pStats DefragmentationStats) {
	C.vmaEndDefragmentation(C.VmaAllocator(unsafe.Pointer(allocator)), C.VmaDefragmentationContext(unsafe.Pointer(context)), pStats.ptr)
}

// EndDefragmentationPass wraps vmaEndDefragmentationPass.
func EndDefragmentationPass(allocator Allocator, context DefragmentationContext, pPassInfo DefragmentationPassMoveInfo) vk.Result {
	ret := C.vmaEndDefragmentationPass(C.VmaAllocator(unsafe.Pointer(allocator)), C.VmaDefragmentationContext(unsafe.Pointer(context)), pPassInfo.ptr)

	return vk.Result(ret)
}

// vmaFindMemoryTypeIndex.pMemoryTypeIndex is unsupported: param: unknown pointer type uint32_t.

// vmaFindMemoryTypeIndexForBufferInfo.pMemoryTypeIndex is unsupported: param: unknown pointer type uint32_t.

// vmaFindMemoryTypeIndexForImageInfo.pMemoryTypeIndex is unsupported: param: unknown pointer type uint32_t.

// FlushAllocation wraps vmaFlushAllocation.
func FlushAllocation(allocator Allocator, allocation Allocation, offset vk.DeviceSize, size vk.DeviceSize) vk.Result {
	ret := C.vmaFlushAllocation(C.VmaAllocator(unsafe.Pointer(allocator)), C.VmaAllocation(unsafe.Pointer(allocation)), C.VkDeviceSize(offset), C.VkDeviceSize(size))

	return vk.Result(ret)
}

// vmaFlushAllocations.offsets is unsupported: param: unknown pointer vk category direct.

// FreeMemory wraps vmaFreeMemory.
func FreeMemory(allocator Allocator, allocation Allocation) {
	C.vmaFreeMemory(C.VmaAllocator(unsafe.Pointer(allocator)), C.VmaAllocation(unsafe.Pointer(allocation)))
}

// FreeMemoryPages wraps vmaFreeMemoryPages.
func FreeMemoryPages(allocator Allocator, allocationCount uintptr, pAllocations ffi.Ref[Allocation]) {
	C.vmaFreeMemoryPages(C.VmaAllocator(unsafe.Pointer(allocator)), C.size_t(allocationCount), (*C.VmaAllocation)(pAllocations.Raw()))
}

// FreeStatsString wraps vmaFreeStatsString.
func FreeStatsString(allocator Allocator, pStatsString ffi.CString) {
	C.vmaFreeStatsString(C.VmaAllocator(unsafe.Pointer(allocator)), (*C.char)(pStatsString.Raw()))
}

// FreeVirtualBlockStatsString wraps vmaFreeVirtualBlockStatsString.
func FreeVirtualBlockStatsString(virtualBlock VirtualBlock, pStatsString ffi.CString) {
	C.vmaFreeVirtualBlockStatsString(C.VmaVirtualBlock(unsafe.Pointer(virtualBlock)), (*C.char)(pStatsString.Raw()))
}

// GetAllocationInfo wraps vmaGetAllocationInfo.
func GetAllocationInfo(allocator Allocator, allocation Allocation, pAllocationInfo AllocationInfo) {
	C.vmaGetAllocationInfo(C.VmaAllocator(unsafe.Pointer(allocator)), C.VmaAllocation(unsafe.Pointer(allocation)), pAllocationInfo.ptr)
}

// GetAllocationInfo2 wraps vmaGetAllocationInfo2.
func GetAllocationInfo2(allocator Allocator, allocation Allocation, pAllocationInfo AllocationInfo2) {
	C.vmaGetAllocationInfo2(C.VmaAllocator(unsafe.Pointer(allocator)), C.VmaAllocation(unsafe.Pointer(allocation)), pAllocationInfo.ptr)
}

// GetAllocationMemoryProperties wraps vmaGetAllocationMemoryProperties.
func GetAllocationMemoryProperties(allocator Allocator, allocation Allocation, pFlags ffi.Ref[vk.MemoryPropertyFlags]) {
	C.vmaGetAllocationMemoryProperties(C.VmaAllocator(unsafe.Pointer(allocator)), C.VmaAllocation(unsafe.Pointer(allocation)), (*C.VkMemoryPropertyFlags)(pFlags.Raw()))
}

// GetAllocatorInfo wraps vmaGetAllocatorInfo.
func GetAllocatorInfo(allocator Allocator, pAllocatorInfo AllocatorInfo) {
	C.vmaGetAllocatorInfo(C.VmaAllocator(unsafe.Pointer(allocator)), pAllocatorInfo.ptr)
}

// GetHeapBudgets wraps vmaGetHeapBudgets.
func GetHeapBudgets(allocator Allocator, pBudgets Budget) {
	C.vmaGetHeapBudgets(C.VmaAllocator(unsafe.Pointer(allocator)), pBudgets.ptr)
}

// vmaGetMemoryProperties.ppPhysicalDeviceMemoryProperties is unsupported: param: unknown pointer2 type VkPhysicalDeviceMemoryProperties.

// GetMemoryTypeProperties wraps vmaGetMemoryTypeProperties.
func GetMemoryTypeProperties(allocator Allocator, memoryTypeIndex uint32, pFlags ffi.Ref[vk.MemoryPropertyFlags]) {
	C.vmaGetMemoryTypeProperties(C.VmaAllocator(unsafe.Pointer(allocator)), C.uint32_t(memoryTypeIndex), (*C.VkMemoryPropertyFlags)(pFlags.Raw()))
}

// vmaGetPhysicalDeviceProperties.ppPhysicalDeviceProperties is unsupported: param: unknown pointer2 type VkPhysicalDeviceProperties.

// vmaGetPoolName.ppName is unsupported: param: unknown pointer2 type char.

// GetPoolStatistics wraps vmaGetPoolStatistics.
func GetPoolStatistics(allocator Allocator, pool Pool, pPoolStats Statistics) {
	C.vmaGetPoolStatistics(C.VmaAllocator(unsafe.Pointer(allocator)), C.VmaPool(unsafe.Pointer(pool)), pPoolStats.ptr)
}

// GetVirtualAllocationInfo wraps vmaGetVirtualAllocationInfo.
func GetVirtualAllocationInfo(virtualBlock VirtualBlock, allocation VirtualAllocation, pVirtualAllocInfo VirtualAllocationInfo) {
	C.vmaGetVirtualAllocationInfo(C.VmaVirtualBlock(unsafe.Pointer(virtualBlock)), C.VmaVirtualAllocation(allocation), pVirtualAllocInfo.ptr)
}

// GetVirtualBlockStatistics wraps vmaGetVirtualBlockStatistics.
func GetVirtualBlockStatistics(virtualBlock VirtualBlock, pStats Statistics) {
	C.vmaGetVirtualBlockStatistics(C.VmaVirtualBlock(unsafe.Pointer(virtualBlock)), pStats.ptr)
}

// InvalidateAllocation wraps vmaInvalidateAllocation.
func InvalidateAllocation(allocator Allocator, allocation Allocation, offset vk.DeviceSize, size vk.DeviceSize) vk.Result {
	ret := C.vmaInvalidateAllocation(C.VmaAllocator(unsafe.Pointer(allocator)), C.VmaAllocation(unsafe.Pointer(allocation)), C.VkDeviceSize(offset), C.VkDeviceSize(size))

	return vk.Result(ret)
}

// vmaInvalidateAllocations.offsets is unsupported: param: unknown pointer vk category direct.

// IsVirtualBlockEmpty wraps vmaIsVirtualBlockEmpty.
func IsVirtualBlockEmpty(virtualBlock VirtualBlock) bool {
	ret := C.vmaIsVirtualBlockEmpty(C.VmaVirtualBlock(unsafe.Pointer(virtualBlock)))

	if ret > 0 {
		return true
	} else {
		return false
	}
}

// MapMemory wraps vmaMapMemory.
func MapMemory(allocator Allocator, allocation Allocation, ppData ffi.Ref[unsafe.Pointer]) vk.Result {
	ret := C.vmaMapMemory(C.VmaAllocator(unsafe.Pointer(allocator)), C.VmaAllocation(unsafe.Pointer(allocation)), (*unsafe.Pointer)(ppData.Raw()))

	return vk.Result(ret)
}

// SetAllocationName wraps vmaSetAllocationName.
func SetAllocationName(allocator Allocator, allocation Allocation, pName ffi.CString) {
	C.vmaSetAllocationName(C.VmaAllocator(unsafe.Pointer(allocator)), C.VmaAllocation(unsafe.Pointer(allocation)), (*C.char)(pName.Raw()))
}

// SetAllocationUserData wraps vmaSetAllocationUserData.
func SetAllocationUserData(allocator Allocator, allocation Allocation, pUserData unsafe.Pointer) {
	C.vmaSetAllocationUserData(C.VmaAllocator(unsafe.Pointer(allocator)), C.VmaAllocation(unsafe.Pointer(allocation)), pUserData)
}

// SetCurrentFrameIndex wraps vmaSetCurrentFrameIndex.
func SetCurrentFrameIndex(allocator Allocator, frameIndex uint32) {
	C.vmaSetCurrentFrameIndex(C.VmaAllocator(unsafe.Pointer(allocator)), C.uint32_t(frameIndex))
}

// SetPoolName wraps vmaSetPoolName.
func SetPoolName(allocator Allocator, pool Pool, pName ffi.CString) {
	C.vmaSetPoolName(C.VmaAllocator(unsafe.Pointer(allocator)), C.VmaPool(unsafe.Pointer(pool)), (*C.char)(pName.Raw()))
}

// SetVirtualAllocationUserData wraps vmaSetVirtualAllocationUserData.
func SetVirtualAllocationUserData(virtualBlock VirtualBlock, allocation VirtualAllocation, pUserData unsafe.Pointer) {
	C.vmaSetVirtualAllocationUserData(C.VmaVirtualBlock(unsafe.Pointer(virtualBlock)), C.VmaVirtualAllocation(allocation), pUserData)
}

// UnmapMemory wraps vmaUnmapMemory.
func UnmapMemory(allocator Allocator, allocation Allocation) {
	C.vmaUnmapMemory(C.VmaAllocator(unsafe.Pointer(allocator)), C.VmaAllocation(unsafe.Pointer(allocation)))
}

// vmaVirtualAllocate.pAllocation is unsupported: param: unknown pointer vma category handle-non-dispatchable.

// VirtualFree wraps vmaVirtualFree.
func VirtualFree(virtualBlock VirtualBlock, allocation VirtualAllocation) {
	C.vmaVirtualFree(C.VmaVirtualBlock(unsafe.Pointer(virtualBlock)), C.VmaVirtualAllocation(allocation))
}
