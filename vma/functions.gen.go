package vma

import (
	ffi "github.com/csnewman/go-gfx/ffi"
	vk "github.com/csnewman/go-gfx/vk"
	"unsafe"
)

// #include "vma.h"
import "C"

// AllocateDedicatedMemory wraps vmaAllocateDedicatedMemory.
/*
  \brief General purpose allocation of a dedicated memory.

  This function is similar vmaAllocateMemory(), but
  it always allocates dedicated memory - flag #VMA_ALLOCATION_CREATE_DEDICATED_MEMORY_BIT is implied.
  It offers additional parameter `pMemoryAllocateNext`,
  which can be used to attach `pNext` chain to the `VkMemoryAllocateInfo` structure.
  It can be useful for importing external memory. For more information, see \ref other_api_interop.
*/
func AllocateDedicatedMemory(allocator Allocator, pVkMemoryRequirements vk.MemoryRequirements, pCreateInfo AllocationCreateInfo, pMemoryAllocateNext uintptr, pAllocation ffi.Ref[Allocation], pAllocationInfo AllocationInfo) vk.Result {
	ret := C.vmaAllocateDedicatedMemory(C.VmaAllocator(unsafe.Pointer(allocator)), (*C.VkMemoryRequirements)(unsafe.Pointer(pVkMemoryRequirements)), (*C.VmaAllocationCreateInfo)(unsafe.Pointer(pCreateInfo)), unsafe.Pointer(pMemoryAllocateNext), (*C.VmaAllocation)(pAllocation.Raw()), (*C.VmaAllocationInfo)(unsafe.Pointer(pAllocationInfo)))

	return vk.Result(ret)
}

// AllocateMemory wraps vmaAllocateMemory.
/*
  \brief General purpose memory allocation.

  \param allocator The main allocator object.
  \param pVkMemoryRequirements Requirements for the allocated memory.
  \param pCreateInfo Allocation creation parameters.
  \param[out] pAllocation Handle to allocated memory.
  \param[out] pAllocationInfo Optional, can be null. Information about allocated memory. It can be also fetched later using vmaGetAllocationInfo().

  The function creates a #VmaAllocation object without creating a buffer or an image together with it.

  - It is recommended to use vmaAllocateMemoryForBuffer(), vmaAllocateMemoryForImage(),
  vmaCreateBuffer(), vmaCreateImage() instead whenever possible.
  - You can also create a buffer or an image later in an existing allocation using
  vmaCreateAliasingBuffer2(), vmaCreateAliasingImage2().
  - You can also create a buffer or an image on your own and bind it to an existing allocation
  using vmaBindBufferMemory2(), vmaBindImageMemory2().

  You must free the returned allocation object using vmaFreeMemory() or vmaFreeMemoryPages().

  There is also extended version of this function: vmaAllocateDedicatedMemory()
  that offers additional parameter `pMemoryAllocateNext`.
*/
func AllocateMemory(allocator Allocator, pVkMemoryRequirements vk.MemoryRequirements, pCreateInfo AllocationCreateInfo, pAllocation ffi.Ref[Allocation], pAllocationInfo AllocationInfo) vk.Result {
	ret := C.vmaAllocateMemory(C.VmaAllocator(unsafe.Pointer(allocator)), (*C.VkMemoryRequirements)(unsafe.Pointer(pVkMemoryRequirements)), (*C.VmaAllocationCreateInfo)(unsafe.Pointer(pCreateInfo)), (*C.VmaAllocation)(pAllocation.Raw()), (*C.VmaAllocationInfo)(unsafe.Pointer(pAllocationInfo)))

	return vk.Result(ret)
}

// AllocateMemoryForBuffer wraps vmaAllocateMemoryForBuffer.
/*
  \brief Allocates memory suitable for given `VkBuffer`.

  \param allocator
  \param buffer
  \param pCreateInfo
  \param[out] pAllocation Handle to allocated memory.
  \param[out] pAllocationInfo Optional. Information about allocated memory. It can be later fetched using function vmaGetAllocationInfo().

  It only creates #VmaAllocation. To bind the memory to the buffer, use vmaBindBufferMemory().

  This is a special-purpose function. In most cases you should use vmaCreateBuffer().

  You must free the allocation using vmaFreeMemory() when no longer needed.
*/
func AllocateMemoryForBuffer(allocator Allocator, buffer vk.Buffer, pCreateInfo AllocationCreateInfo, pAllocation ffi.Ref[Allocation], pAllocationInfo AllocationInfo) vk.Result {
	ret := C.vmaAllocateMemoryForBuffer(C.VmaAllocator(unsafe.Pointer(allocator)), C.VkBuffer(buffer), (*C.VmaAllocationCreateInfo)(unsafe.Pointer(pCreateInfo)), (*C.VmaAllocation)(pAllocation.Raw()), (*C.VmaAllocationInfo)(unsafe.Pointer(pAllocationInfo)))

	return vk.Result(ret)
}

// AllocateMemoryForImage wraps vmaAllocateMemoryForImage.
/*
  \brief Allocates memory suitable for given `VkImage`.

  \param allocator
  \param image
  \param pCreateInfo
  \param[out] pAllocation Handle to allocated memory.
  \param[out] pAllocationInfo Optional. Information about allocated memory. It can be later fetched using function vmaGetAllocationInfo().

  It only creates #VmaAllocation. To bind the memory to the buffer, use vmaBindImageMemory().

  This is a special-purpose function. In most cases you should use vmaCreateImage().

  You must free the allocation using vmaFreeMemory() when no longer needed.
*/
func AllocateMemoryForImage(allocator Allocator, image vk.Image, pCreateInfo AllocationCreateInfo, pAllocation ffi.Ref[Allocation], pAllocationInfo AllocationInfo) vk.Result {
	ret := C.vmaAllocateMemoryForImage(C.VmaAllocator(unsafe.Pointer(allocator)), C.VkImage(image), (*C.VmaAllocationCreateInfo)(unsafe.Pointer(pCreateInfo)), (*C.VmaAllocation)(pAllocation.Raw()), (*C.VmaAllocationInfo)(unsafe.Pointer(pAllocationInfo)))

	return vk.Result(ret)
}

// AllocateMemoryPages wraps vmaAllocateMemoryPages.
/*
  \brief General purpose memory allocation for multiple allocation objects at once.

  \param allocator Allocator object.
  \param pVkMemoryRequirements Memory requirements for each allocation.
  \param pCreateInfo Creation parameters for each allocation.
  \param allocationCount Number of allocations to make.
  \param[out] pAllocations Pointer to array that will be filled with handles to created allocations.
  \param[out] pAllocationInfo Optional. Pointer to array that will be filled with parameters of created allocations.

  You should free the memory using vmaFreeMemory() or vmaFreeMemoryPages().

  Word "pages" is just a suggestion to use this function to allocate pieces of memory needed for sparse binding.
  It is just a general purpose allocation function able to make multiple allocations at once.
  It may be internally optimized to be more efficient than calling vmaAllocateMemory() `allocationCount` times.

  All allocations are made using same parameters. All of them are created out of the same memory pool and type.
  If any allocation fails, all allocations already made within this function call are also freed, so that when
  returned result is not `VK_SUCCESS`, `pAllocation` array is always entirely filled with `VK_NULL_HANDLE`.
*/
func AllocateMemoryPages(allocator Allocator, pVkMemoryRequirements vk.MemoryRequirements, pCreateInfo AllocationCreateInfo, allocationCount uintptr, pAllocations ffi.Ref[Allocation], pAllocationInfo AllocationInfo) vk.Result {
	ret := C.vmaAllocateMemoryPages(C.VmaAllocator(unsafe.Pointer(allocator)), (*C.VkMemoryRequirements)(unsafe.Pointer(pVkMemoryRequirements)), (*C.VmaAllocationCreateInfo)(unsafe.Pointer(pCreateInfo)), C.size_t(allocationCount), (*C.VmaAllocation)(pAllocations.Raw()), (*C.VmaAllocationInfo)(unsafe.Pointer(pAllocationInfo)))

	return vk.Result(ret)
}

// BeginDefragmentation wraps vmaBeginDefragmentation.
/*
  \brief Begins defragmentation process.

  \param allocator Allocator object.
  \param pInfo Structure filled with parameters of defragmentation.
  \param[out] pContext Context object that must be passed to vmaEndDefragmentation() to finish defragmentation.
  \returns
  - `VK_SUCCESS` if defragmentation can begin.
  - `VK_ERROR_FEATURE_NOT_PRESENT` if defragmentation is not supported.

  For more information about defragmentation, see documentation chapter:
  [Defragmentation](@ref defragmentation).
*/
func BeginDefragmentation(allocator Allocator, pInfo DefragmentationInfo, pContext ffi.Ref[DefragmentationContext]) vk.Result {
	ret := C.vmaBeginDefragmentation(C.VmaAllocator(unsafe.Pointer(allocator)), (*C.VmaDefragmentationInfo)(unsafe.Pointer(pInfo)), (*C.VmaDefragmentationContext)(pContext.Raw()))

	return vk.Result(ret)
}

// BeginDefragmentationPass wraps vmaBeginDefragmentationPass.
/*
  \brief Starts single defragmentation pass.

  \param allocator Allocator object.
  \param context Context object that has been created by vmaBeginDefragmentation().
  \param[out] pPassInfo Computed information for current pass.
  \returns
  - `VK_SUCCESS` if no more moves are possible. Then you can omit call to vmaEndDefragmentationPass() and simply end whole defragmentation.
  - `VK_INCOMPLETE` if there are pending moves returned in `pPassInfo`. You need to perform them, call vmaEndDefragmentationPass(),
  and then preferably try another pass with vmaBeginDefragmentationPass().
*/
func BeginDefragmentationPass(allocator Allocator, context DefragmentationContext, pPassInfo DefragmentationPassMoveInfo) vk.Result {
	ret := C.vmaBeginDefragmentationPass(C.VmaAllocator(unsafe.Pointer(allocator)), C.VmaDefragmentationContext(unsafe.Pointer(context)), (*C.VmaDefragmentationPassMoveInfo)(unsafe.Pointer(pPassInfo)))

	return vk.Result(ret)
}

// BindBufferMemory wraps vmaBindBufferMemory.
/*
  \brief Binds buffer to allocation.

  Binds specified buffer to region of memory represented by specified allocation.
  Gets `VkDeviceMemory` handle and offset from the allocation.
  If you want to create a buffer, allocate memory for it and bind them together separately,
  you should use this function for binding instead of standard `vkBindBufferMemory()`,
  because it ensures proper synchronization so that when a `VkDeviceMemory` object is used by multiple
  allocations, calls to `vkBind*Memory()` or `vkMapMemory()` won't happen from multiple threads simultaneously
  (which is illegal in Vulkan).

  It is recommended to use function vmaCreateBuffer() instead of this one.
*/
func BindBufferMemory(allocator Allocator, allocation Allocation, buffer vk.Buffer) vk.Result {
	ret := C.vmaBindBufferMemory(C.VmaAllocator(unsafe.Pointer(allocator)), C.VmaAllocation(unsafe.Pointer(allocation)), C.VkBuffer(buffer))

	return vk.Result(ret)
}

// BindBufferMemory2 wraps vmaBindBufferMemory2.
/*
  \brief Binds buffer to allocation with additional parameters.

  \param allocator
  \param allocation
  \param allocationLocalOffset Additional offset to be added while binding, relative to the beginning of the `allocation`. Normally it should be 0.
  \param buffer
  \param pNext A chain of structures to be attached to `VkBindBufferMemoryInfoKHR` structure used internally. Normally it should be null.

  This function is similar to vmaBindBufferMemory(), but it provides additional parameters.

  If `pNext` is not null, #VmaAllocator object must have been created with #VMA_ALLOCATOR_CREATE_KHR_BIND_MEMORY2_BIT flag
  or with VmaAllocatorCreateInfo::vulkanApiVersion `>= VK_API_VERSION_1_1`. Otherwise the call fails.
*/
func BindBufferMemory2(allocator Allocator, allocation Allocation, allocationLocalOffset vk.DeviceSize, buffer vk.Buffer, pNext uintptr) vk.Result {
	ret := C.vmaBindBufferMemory2(C.VmaAllocator(unsafe.Pointer(allocator)), C.VmaAllocation(unsafe.Pointer(allocation)), C.VkDeviceSize(allocationLocalOffset), C.VkBuffer(buffer), unsafe.Pointer(pNext))

	return vk.Result(ret)
}

// BindImageMemory wraps vmaBindImageMemory.
/*
  \brief Binds image to allocation.

  Binds specified image to region of memory represented by specified allocation.
  Gets `VkDeviceMemory` handle and offset from the allocation.
  If you want to create an image, allocate memory for it and bind them together separately,
  you should use this function for binding instead of standard `vkBindImageMemory()`,
  because it ensures proper synchronization so that when a `VkDeviceMemory` object is used by multiple
  allocations, calls to `vkBind*Memory()` or `vkMapMemory()` won't happen from multiple threads simultaneously
  (which is illegal in Vulkan).

  It is recommended to use function vmaCreateImage() instead of this one.
*/
func BindImageMemory(allocator Allocator, allocation Allocation, image vk.Image) vk.Result {
	ret := C.vmaBindImageMemory(C.VmaAllocator(unsafe.Pointer(allocator)), C.VmaAllocation(unsafe.Pointer(allocation)), C.VkImage(image))

	return vk.Result(ret)
}

// BindImageMemory2 wraps vmaBindImageMemory2.
/*
  \brief Binds image to allocation with additional parameters.

  \param allocator
  \param allocation
  \param allocationLocalOffset Additional offset to be added while binding, relative to the beginning of the `allocation`. Normally it should be 0.
  \param image
  \param pNext A chain of structures to be attached to `VkBindImageMemoryInfoKHR` structure used internally. Normally it should be null.

  This function is similar to vmaBindImageMemory(), but it provides additional parameters.

  If `pNext` is not null, #VmaAllocator object must have been created with #VMA_ALLOCATOR_CREATE_KHR_BIND_MEMORY2_BIT flag
  or with VmaAllocatorCreateInfo::vulkanApiVersion `>= VK_API_VERSION_1_1`. Otherwise the call fails.
*/
func BindImageMemory2(allocator Allocator, allocation Allocation, allocationLocalOffset vk.DeviceSize, image vk.Image, pNext uintptr) vk.Result {
	ret := C.vmaBindImageMemory2(C.VmaAllocator(unsafe.Pointer(allocator)), C.VmaAllocation(unsafe.Pointer(allocation)), C.VkDeviceSize(allocationLocalOffset), C.VkImage(image), unsafe.Pointer(pNext))

	return vk.Result(ret)
}

// BuildStatsString wraps vmaBuildStatsString.
/*
  \brief Builds and returns statistics as a null-terminated string in JSON format.
  \param allocator
  \param[out] ppStatsString Must be freed using vmaFreeStatsString() function.
  \param detailedMap
*/
func BuildStatsString(allocator Allocator, ppStatsString ffi.Ref[ffi.CString], detailedMap bool) {
	tmp_detailedMap := 0

	if detailedMap {
		tmp_detailedMap = 1
	}

	C.vmaBuildStatsString(C.VmaAllocator(unsafe.Pointer(allocator)), (**C.char)(ppStatsString.Raw()), C.VkBool32(tmp_detailedMap))
}

// BuildVirtualBlockStatsString wraps vmaBuildVirtualBlockStatsString.
/*
  \brief Builds and returns a null-terminated string in JSON format with information about given #VmaVirtualBlock.
  \param virtualBlock Virtual block.
  \param[out] ppStatsString Returned string.
  \param detailedMap Pass `VK_FALSE` to only obtain statistics as returned by vmaCalculateVirtualBlockStatistics(). Pass `VK_TRUE` to also obtain full list of allocations and free spaces.

  Returned string must be freed using vmaFreeVirtualBlockStatsString().
*/
func BuildVirtualBlockStatsString(virtualBlock VirtualBlock, ppStatsString ffi.Ref[ffi.CString], detailedMap bool) {
	tmp_detailedMap := 0

	if detailedMap {
		tmp_detailedMap = 1
	}

	C.vmaBuildVirtualBlockStatsString(C.VmaVirtualBlock(unsafe.Pointer(virtualBlock)), (**C.char)(ppStatsString.Raw()), C.VkBool32(tmp_detailedMap))
}

// CalculatePoolStatistics wraps vmaCalculatePoolStatistics.
/*
  \brief Retrieves detailed statistics of existing #VmaPool object.

  \param allocator Allocator object.
  \param pool Pool object.
  \param[out] pPoolStats Statistics of specified pool.
*/
func CalculatePoolStatistics(allocator Allocator, pool Pool, pPoolStats DetailedStatistics) {
	C.vmaCalculatePoolStatistics(C.VmaAllocator(unsafe.Pointer(allocator)), C.VmaPool(unsafe.Pointer(pool)), (*C.VmaDetailedStatistics)(unsafe.Pointer(pPoolStats)))
}

// CalculateStatistics wraps vmaCalculateStatistics.
/*
  \brief Retrieves statistics from current state of the Allocator.

  This function is called "calculate" not "get" because it has to traverse all
  internal data structures, so it may be quite slow. Use it for debugging purposes.
  For faster but more brief statistics suitable to be called every frame or every allocation,
  use vmaGetHeapBudgets().

  Note that when using allocator from multiple threads, returned information may immediately
  become outdated.
*/
func CalculateStatistics(allocator Allocator, pStats TotalStatistics) {
	C.vmaCalculateStatistics(C.VmaAllocator(unsafe.Pointer(allocator)), (*C.VmaTotalStatistics)(unsafe.Pointer(pStats)))
}

// CalculateVirtualBlockStatistics wraps vmaCalculateVirtualBlockStatistics.
/*
  \brief Calculates and returns detailed statistics about virtual allocations and memory usage in given #VmaVirtualBlock.

  This function is slow to call. Use for debugging purposes.
  For less detailed statistics, see vmaGetVirtualBlockStatistics().
*/
func CalculateVirtualBlockStatistics(virtualBlock VirtualBlock, pStats DetailedStatistics) {
	C.vmaCalculateVirtualBlockStatistics(C.VmaVirtualBlock(unsafe.Pointer(virtualBlock)), (*C.VmaDetailedStatistics)(unsafe.Pointer(pStats)))
}

// CheckCorruption wraps vmaCheckCorruption.
/*
  \brief Checks magic number in margins around all allocations in given memory types (in both default and custom pools) in search for corruptions.

  \param allocator
  \param memoryTypeBits Bit mask, where each bit set means that a memory type with that index should be checked.

  Corruption detection is enabled only when `VMA_DEBUG_DETECT_CORRUPTION` macro is defined to nonzero,
  `VMA_DEBUG_MARGIN` is defined to nonzero and only for memory types that are
  `HOST_VISIBLE` and `HOST_COHERENT`. For more information, see [Corruption detection](@ref debugging_memory_usage_corruption_detection).

  Possible return values:

  - `VK_ERROR_FEATURE_NOT_PRESENT` - corruption detection is not enabled for any of specified memory types.
  - `VK_SUCCESS` - corruption detection has been performed and succeeded.
  - `VK_ERROR_UNKNOWN` - corruption detection has been performed and found memory corruptions around one of the allocations.
  `VMA_ASSERT` is also fired in that case.
  - Other value: Error returned by Vulkan, e.g. memory mapping failure.
*/
func CheckCorruption(allocator Allocator, memoryTypeBits uint32) vk.Result {
	ret := C.vmaCheckCorruption(C.VmaAllocator(unsafe.Pointer(allocator)), C.uint32_t(memoryTypeBits))

	return vk.Result(ret)
}

// CheckPoolCorruption wraps vmaCheckPoolCorruption.
/*
  \brief Checks magic number in margins around all allocations in given memory pool in search for corruptions.

  Corruption detection is enabled only when `VMA_DEBUG_DETECT_CORRUPTION` macro is defined to nonzero,
  `VMA_DEBUG_MARGIN` is defined to nonzero and the pool is created in memory type that is
  `HOST_VISIBLE` and `HOST_COHERENT`. For more information, see [Corruption detection](@ref debugging_memory_usage_corruption_detection).

  Possible return values:

  - `VK_ERROR_FEATURE_NOT_PRESENT` - corruption detection is not enabled for specified pool.
  - `VK_SUCCESS` - corruption detection has been performed and succeeded.
  - `VK_ERROR_UNKNOWN` - corruption detection has been performed and found memory corruptions around one of the allocations.
  `VMA_ASSERT` is also fired in that case.
  - Other value: Error returned by Vulkan, e.g. memory mapping failure.
*/
func CheckPoolCorruption(allocator Allocator, pool Pool) vk.Result {
	ret := C.vmaCheckPoolCorruption(C.VmaAllocator(unsafe.Pointer(allocator)), C.VmaPool(unsafe.Pointer(pool)))

	return vk.Result(ret)
}

// ClearVirtualBlock wraps vmaClearVirtualBlock.
/*
  \brief Frees all virtual allocations inside given #VmaVirtualBlock.

  You must either call this function or free each virtual allocation individually with vmaVirtualFree()
  before destroying a virtual block. Otherwise, an assert is called.

  If you keep pointer to some additional metadata associated with your virtual allocation in its `pUserData`,
  don't forget to free it as well.
*/
func ClearVirtualBlock(virtualBlock VirtualBlock) {
	C.vmaClearVirtualBlock(C.VmaVirtualBlock(unsafe.Pointer(virtualBlock)))
}

// CopyAllocationToMemory wraps vmaCopyAllocationToMemory.
/*
  \brief Invalidates memory in the host caches if needed, maps the allocation temporarily if needed, and copies data from it to a specified host pointer.

  \param allocator
  \param srcAllocation   Handle to the allocation that becomes source of the copy.
  \param srcAllocationLocalOffset  Offset within `srcAllocation` where to read copied data, in bytes.
  \param pDstHostPointer Pointer to the host memory that become destination of the copy.
  \param size            Number of bytes to copy.

  This is a convenience function that allows to copy data from an allocation to a host pointer easily.
  Same behavior can be achieved by calling vmaInvalidateAllocation(), vmaMapMemory(), `memcpy()`, vmaUnmapMemory().

  This function should be called only for allocations created in a memory type that has `VK_MEMORY_PROPERTY_HOST_VISIBLE_BIT`
  and `VK_MEMORY_PROPERTY_HOST_CACHED_BIT` flag.
  It can be ensured e.g. by using #VMA_MEMORY_USAGE_AUTO and #VMA_ALLOCATION_CREATE_HOST_ACCESS_RANDOM_BIT.
  Otherwise, the function may fail and generate a Validation Layers error.
  It may also work very slowly when reading from an uncached memory.

  `srcAllocationLocalOffset` is relative to the contents of given `srcAllocation`.
  If you mean whole allocation, you should pass 0.
  Do not pass allocation's offset within device memory block as this parameter!
*/
func CopyAllocationToMemory(allocator Allocator, srcAllocation Allocation, srcAllocationLocalOffset vk.DeviceSize, pDstHostPointer uintptr, size vk.DeviceSize) vk.Result {
	ret := C.vmaCopyAllocationToMemory(C.VmaAllocator(unsafe.Pointer(allocator)), C.VmaAllocation(unsafe.Pointer(srcAllocation)), C.VkDeviceSize(srcAllocationLocalOffset), unsafe.Pointer(pDstHostPointer), C.VkDeviceSize(size))

	return vk.Result(ret)
}

// CopyMemoryToAllocation wraps vmaCopyMemoryToAllocation.
/*
  \brief Maps the allocation temporarily if needed, copies data from specified host pointer to it, and flushes the memory from the host caches if needed.

  \param allocator
  \param pSrcHostPointer Pointer to the host data that become source of the copy.
  \param dstAllocation   Handle to the allocation that becomes destination of the copy.
  \param dstAllocationLocalOffset  Offset within `dstAllocation` where to write copied data, in bytes.
  \param size            Number of bytes to copy.

  This is a convenience function that allows to copy data from a host pointer to an allocation easily.
  Same behavior can be achieved by calling vmaMapMemory(), `memcpy()`, vmaUnmapMemory(), vmaFlushAllocation().

  This function can be called only for allocations created in a memory type that has `VK_MEMORY_PROPERTY_HOST_VISIBLE_BIT` flag.
  It can be ensured e.g. by using #VMA_MEMORY_USAGE_AUTO and #VMA_ALLOCATION_CREATE_HOST_ACCESS_SEQUENTIAL_WRITE_BIT or
  #VMA_ALLOCATION_CREATE_HOST_ACCESS_RANDOM_BIT.
  Otherwise, the function will fail and generate a Validation Layers error.

  `dstAllocationLocalOffset` is relative to the contents of given `dstAllocation`.
  If you mean whole allocation, you should pass 0.
  Do not pass allocation's offset within device memory block this parameter!
*/
func CopyMemoryToAllocation(allocator Allocator, pSrcHostPointer uintptr, dstAllocation Allocation, dstAllocationLocalOffset vk.DeviceSize, size vk.DeviceSize) vk.Result {
	ret := C.vmaCopyMemoryToAllocation(C.VmaAllocator(unsafe.Pointer(allocator)), unsafe.Pointer(pSrcHostPointer), C.VmaAllocation(unsafe.Pointer(dstAllocation)), C.VkDeviceSize(dstAllocationLocalOffset), C.VkDeviceSize(size))

	return vk.Result(ret)
}

// CreateAliasingBuffer wraps vmaCreateAliasingBuffer.
/*
  \brief Creates a new `VkBuffer`, binds already created memory for it.

  \param allocator
  \param allocation Allocation that provides memory to be used for binding new buffer to it.
  \param pBufferCreateInfo
  \param[out] pBuffer Buffer that was created.

  This function automatically:

  -# Creates buffer.
  -# Binds the buffer with the supplied memory.

  If any of these operations fail, buffer is not created,
  returned value is negative error code and `*pBuffer` is null.

  If the function succeeded, you must destroy the buffer when you
  no longer need it using `vkDestroyBuffer()`. If you want to also destroy the corresponding
  allocation you can use convenience function vmaDestroyBuffer().

  \note There is a new version of this function augmented with parameter `allocationLocalOffset` - see vmaCreateAliasingBuffer2().
*/
func CreateAliasingBuffer(allocator Allocator, allocation Allocation, pBufferCreateInfo vk.BufferCreateInfo, pBuffer ffi.Ref[vk.Buffer]) vk.Result {
	ret := C.vmaCreateAliasingBuffer(C.VmaAllocator(unsafe.Pointer(allocator)), C.VmaAllocation(unsafe.Pointer(allocation)), (*C.VkBufferCreateInfo)(unsafe.Pointer(pBufferCreateInfo)), (*C.VkBuffer)(pBuffer.Raw()))

	return vk.Result(ret)
}

// CreateAliasingBuffer2 wraps vmaCreateAliasingBuffer2.
/*
  \brief Creates a new `VkBuffer`, binds already created memory for it.

  \param allocator
  \param allocation Allocation that provides memory to be used for binding new buffer to it.
  \param allocationLocalOffset Additional offset to be added while binding, relative to the beginning of the allocation. Normally it should be 0.
  \param pBufferCreateInfo
  \param[out] pBuffer Buffer that was created.

  This function automatically:

  -# Creates buffer.
  -# Binds the buffer with the supplied memory.

  If any of these operations fail, buffer is not created,
  returned value is negative error code and `*pBuffer` is null.

  If the function succeeded, you must destroy the buffer when you
  no longer need it using `vkDestroyBuffer()`. If you want to also destroy the corresponding
  allocation you can use convenience function vmaDestroyBuffer().

  \note This is a new version of the function augmented with parameter `allocationLocalOffset`.
*/
func CreateAliasingBuffer2(allocator Allocator, allocation Allocation, allocationLocalOffset vk.DeviceSize, pBufferCreateInfo vk.BufferCreateInfo, pBuffer ffi.Ref[vk.Buffer]) vk.Result {
	ret := C.vmaCreateAliasingBuffer2(C.VmaAllocator(unsafe.Pointer(allocator)), C.VmaAllocation(unsafe.Pointer(allocation)), C.VkDeviceSize(allocationLocalOffset), (*C.VkBufferCreateInfo)(unsafe.Pointer(pBufferCreateInfo)), (*C.VkBuffer)(pBuffer.Raw()))

	return vk.Result(ret)
}

// CreateAliasingImage wraps vmaCreateAliasingImage.
//
//	Function similar to vmaCreateAliasingBuffer() but for images.
func CreateAliasingImage(allocator Allocator, allocation Allocation, pImageCreateInfo vk.ImageCreateInfo, pImage ffi.Ref[vk.Image]) vk.Result {
	ret := C.vmaCreateAliasingImage(C.VmaAllocator(unsafe.Pointer(allocator)), C.VmaAllocation(unsafe.Pointer(allocation)), (*C.VkImageCreateInfo)(unsafe.Pointer(pImageCreateInfo)), (*C.VkImage)(pImage.Raw()))

	return vk.Result(ret)
}

// CreateAliasingImage2 wraps vmaCreateAliasingImage2.
//
//	Function similar to vmaCreateAliasingBuffer2() but for images.
func CreateAliasingImage2(allocator Allocator, allocation Allocation, allocationLocalOffset vk.DeviceSize, pImageCreateInfo vk.ImageCreateInfo, pImage ffi.Ref[vk.Image]) vk.Result {
	ret := C.vmaCreateAliasingImage2(C.VmaAllocator(unsafe.Pointer(allocator)), C.VmaAllocation(unsafe.Pointer(allocation)), C.VkDeviceSize(allocationLocalOffset), (*C.VkImageCreateInfo)(unsafe.Pointer(pImageCreateInfo)), (*C.VkImage)(pImage.Raw()))

	return vk.Result(ret)
}

// CreateAllocator wraps vmaCreateAllocator.
//
//	Creates #VmaAllocator object.
func CreateAllocator(pCreateInfo AllocatorCreateInfo, pAllocator ffi.Ref[Allocator]) vk.Result {
	ret := C.vmaCreateAllocator((*C.VmaAllocatorCreateInfo)(unsafe.Pointer(pCreateInfo)), (*C.VmaAllocator)(pAllocator.Raw()))

	return vk.Result(ret)
}

// CreateBuffer wraps vmaCreateBuffer.
/*
  \brief Creates a new `VkBuffer`, allocates and binds memory for it.

  \param allocator The main allocator object.
  \param pBufferCreateInfo Buffer creation parameters.
  \param pAllocationCreateInfo Allocation creation parameters.
  \param[out] pBuffer Buffer that was created.
  \param[out] pAllocation Allocation that was created.
  \param[out] pAllocationInfo Optional, can be null. Information about allocated memory.
  It can be also fetched later using vmaGetAllocationInfo().

  This function automatically:

  -# Creates buffer.
  -# Allocates appropriate memory for it.
  -# Binds the buffer with the memory.

  If any of these operations fail, buffer and allocation are not created,
  returned value is negative error code, `*pBuffer` and `*pAllocation` are returned as null.

  If the function succeeded, you must destroy both buffer and allocation when you
  no longer need them using either convenience function vmaDestroyBuffer() or
  separately, using `vkDestroyBuffer()` and vmaFreeMemory().

  If VK_KHR_dedicated_allocation extenion or Vulkan version >= 1.1 is used,
  the function queries the driver whether
  it requires or prefers the new buffer to have dedicated allocation. If yes,
  and if dedicated allocation is possible
  (#VMA_ALLOCATION_CREATE_NEVER_ALLOCATE_BIT is not used), it creates dedicated
  allocation for this buffer, just like when using
  #VMA_ALLOCATION_CREATE_DEDICATED_MEMORY_BIT.

  \note This function creates a new `VkBuffer`. Sub-allocation of parts of one large buffer,
  although recommended as a good practice, is out of scope of this library and could be implemented
  by the user as a higher-level logic on top of VMA.

  There are also extended versions of this function available:
  - With additional parameter `minAlignment` - see vmaCreateBufferWithAlignment().
  - With additional parameter `pMemoryAllocateNext` - see vmaCreateDedicatedBuffer().
*/
func CreateBuffer(allocator Allocator, pBufferCreateInfo vk.BufferCreateInfo, pAllocationCreateInfo AllocationCreateInfo, pBuffer ffi.Ref[vk.Buffer], pAllocation ffi.Ref[Allocation], pAllocationInfo AllocationInfo) vk.Result {
	ret := C.vmaCreateBuffer(C.VmaAllocator(unsafe.Pointer(allocator)), (*C.VkBufferCreateInfo)(unsafe.Pointer(pBufferCreateInfo)), (*C.VmaAllocationCreateInfo)(unsafe.Pointer(pAllocationCreateInfo)), (*C.VkBuffer)(pBuffer.Raw()), (*C.VmaAllocation)(pAllocation.Raw()), (*C.VmaAllocationInfo)(unsafe.Pointer(pAllocationInfo)))

	return vk.Result(ret)
}

// CreateBufferWithAlignment wraps vmaCreateBufferWithAlignment.
/*
  \brief Creates a buffer with additional minimum alignment.

  Similar to vmaCreateBuffer() but provides additional parameter `minAlignment` which allows to specify custom,
  minimum alignment to be used when placing the buffer inside a larger memory block, which may be needed e.g.
  for interop with OpenGL.
*/
func CreateBufferWithAlignment(allocator Allocator, pBufferCreateInfo vk.BufferCreateInfo, pAllocationCreateInfo AllocationCreateInfo, minAlignment vk.DeviceSize, pBuffer ffi.Ref[vk.Buffer], pAllocation ffi.Ref[Allocation], pAllocationInfo AllocationInfo) vk.Result {
	ret := C.vmaCreateBufferWithAlignment(C.VmaAllocator(unsafe.Pointer(allocator)), (*C.VkBufferCreateInfo)(unsafe.Pointer(pBufferCreateInfo)), (*C.VmaAllocationCreateInfo)(unsafe.Pointer(pAllocationCreateInfo)), C.VkDeviceSize(minAlignment), (*C.VkBuffer)(pBuffer.Raw()), (*C.VmaAllocation)(pAllocation.Raw()), (*C.VmaAllocationInfo)(unsafe.Pointer(pAllocationInfo)))

	return vk.Result(ret)
}

// CreateDedicatedBuffer wraps vmaCreateDedicatedBuffer.
/*
  \brief Creates a dedicated buffer while offering extra parameter `pMemoryAllocateNext`.

  This function is similar vmaCreateBuffer(), but
  it always allocates dedicated memory for the buffer - flag #VMA_ALLOCATION_CREATE_DEDICATED_MEMORY_BIT is implied.
  It offers additional parameter `pMemoryAllocateNext`,
  which can be used to attach `pNext` chain to the `VkMemoryAllocateInfo` structure.
  It can be useful for importing external memory. For more information, see \ref other_api_interop.
*/
func CreateDedicatedBuffer(allocator Allocator, pBufferCreateInfo vk.BufferCreateInfo, pAllocationCreateInfo AllocationCreateInfo, pMemoryAllocateNext uintptr, pBuffer ffi.Ref[vk.Buffer], pAllocation ffi.Ref[Allocation], pAllocationInfo AllocationInfo) vk.Result {
	ret := C.vmaCreateDedicatedBuffer(C.VmaAllocator(unsafe.Pointer(allocator)), (*C.VkBufferCreateInfo)(unsafe.Pointer(pBufferCreateInfo)), (*C.VmaAllocationCreateInfo)(unsafe.Pointer(pAllocationCreateInfo)), unsafe.Pointer(pMemoryAllocateNext), (*C.VkBuffer)(pBuffer.Raw()), (*C.VmaAllocation)(pAllocation.Raw()), (*C.VmaAllocationInfo)(unsafe.Pointer(pAllocationInfo)))

	return vk.Result(ret)
}

// CreateDedicatedImage wraps vmaCreateDedicatedImage.
/*
  \brief Function similar to vmaCreateDedicatedBuffer() but for images.

  This function is similar vmaCreateImage(), but
  it always allocates dedicated memory for the image - flag #VMA_ALLOCATION_CREATE_DEDICATED_MEMORY_BIT is implied.
  It offers additional parameter `pMemoryAllocateNext`,
  which can be used to attach `pNext` chain to the `VkMemoryAllocateInfo` structure.
  It can be useful for importing external memory. For more information, see \ref other_api_interop.
*/
func CreateDedicatedImage(allocator Allocator, pImageCreateInfo vk.ImageCreateInfo, pAllocationCreateInfo AllocationCreateInfo, pMemoryAllocateNext uintptr, pImage ffi.Ref[vk.Image], pAllocation ffi.Ref[Allocation], pAllocationInfo AllocationInfo) vk.Result {
	ret := C.vmaCreateDedicatedImage(C.VmaAllocator(unsafe.Pointer(allocator)), (*C.VkImageCreateInfo)(unsafe.Pointer(pImageCreateInfo)), (*C.VmaAllocationCreateInfo)(unsafe.Pointer(pAllocationCreateInfo)), unsafe.Pointer(pMemoryAllocateNext), (*C.VkImage)(pImage.Raw()), (*C.VmaAllocation)(pAllocation.Raw()), (*C.VmaAllocationInfo)(unsafe.Pointer(pAllocationInfo)))

	return vk.Result(ret)
}

// CreateImage wraps vmaCreateImage.
/*
  \brief Function similar to vmaCreateBuffer() but for images.

  There is also an extended version of this function available: vmaCreateDedicatedImage()
  which offers additional parameter `pMemoryAllocateNext`.
*/
func CreateImage(allocator Allocator, pImageCreateInfo vk.ImageCreateInfo, pAllocationCreateInfo AllocationCreateInfo, pImage ffi.Ref[vk.Image], pAllocation ffi.Ref[Allocation], pAllocationInfo AllocationInfo) vk.Result {
	ret := C.vmaCreateImage(C.VmaAllocator(unsafe.Pointer(allocator)), (*C.VkImageCreateInfo)(unsafe.Pointer(pImageCreateInfo)), (*C.VmaAllocationCreateInfo)(unsafe.Pointer(pAllocationCreateInfo)), (*C.VkImage)(pImage.Raw()), (*C.VmaAllocation)(pAllocation.Raw()), (*C.VmaAllocationInfo)(unsafe.Pointer(pAllocationInfo)))

	return vk.Result(ret)
}

// CreatePool wraps vmaCreatePool.
/*
  \brief Allocates Vulkan device memory and creates #VmaPool object.

  \param allocator Allocator object.
  \param pCreateInfo Parameters of pool to create.
  \param[out] pPool Handle to created pool.
*/
func CreatePool(allocator Allocator, pCreateInfo PoolCreateInfo, pPool ffi.Ref[Pool]) vk.Result {
	ret := C.vmaCreatePool(C.VmaAllocator(unsafe.Pointer(allocator)), (*C.VmaPoolCreateInfo)(unsafe.Pointer(pCreateInfo)), (*C.VmaPool)(pPool.Raw()))

	return vk.Result(ret)
}

// CreateVirtualBlock wraps vmaCreateVirtualBlock.
/*
  \brief Creates new #VmaVirtualBlock object.

  \param pCreateInfo Parameters for creation.
  \param[out] pVirtualBlock Returned virtual block object or `VMA_NULL` if creation failed.
*/
func CreateVirtualBlock(pCreateInfo VirtualBlockCreateInfo, pVirtualBlock ffi.Ref[VirtualBlock]) vk.Result {
	ret := C.vmaCreateVirtualBlock((*C.VmaVirtualBlockCreateInfo)(unsafe.Pointer(pCreateInfo)), (*C.VmaVirtualBlock)(pVirtualBlock.Raw()))

	return vk.Result(ret)
}

// DestroyAllocator wraps vmaDestroyAllocator.
//
//	Destroys allocator object.
func DestroyAllocator(allocator Allocator) {
	C.vmaDestroyAllocator(C.VmaAllocator(unsafe.Pointer(allocator)))
}

// DestroyBuffer wraps vmaDestroyBuffer.
/*
  \brief Destroys Vulkan buffer and frees allocated memory.

  This is just a convenience function equivalent to:

  \code
  vkDestroyBuffer(device, buffer, allocationCallbacks);
  vmaFreeMemory(allocator, allocation);
  \endcode

  It is safe to pass null as buffer and/or allocation.
*/
func DestroyBuffer(allocator Allocator, buffer vk.Buffer, allocation Allocation) {
	C.vmaDestroyBuffer(C.VmaAllocator(unsafe.Pointer(allocator)), C.VkBuffer(buffer), C.VmaAllocation(unsafe.Pointer(allocation)))
}

// DestroyImage wraps vmaDestroyImage.
/*
  \brief Destroys Vulkan image and frees allocated memory.

  This is just a convenience function equivalent to:

  \code
  vkDestroyImage(device, image, allocationCallbacks);
  vmaFreeMemory(allocator, allocation);
  \endcode

  It is safe to pass null as image and/or allocation.
*/
func DestroyImage(allocator Allocator, image vk.Image, allocation Allocation) {
	C.vmaDestroyImage(C.VmaAllocator(unsafe.Pointer(allocator)), C.VkImage(image), C.VmaAllocation(unsafe.Pointer(allocation)))
}

// DestroyPool wraps vmaDestroyPool.
//
//	\brief Destroys #VmaPool object and frees Vulkan device memory.
func DestroyPool(allocator Allocator, pool Pool) {
	C.vmaDestroyPool(C.VmaAllocator(unsafe.Pointer(allocator)), C.VmaPool(unsafe.Pointer(pool)))
}

// DestroyVirtualBlock wraps vmaDestroyVirtualBlock.
/*
  \brief Destroys #VmaVirtualBlock object.

  Please note that you should consciously handle virtual allocations that could remain unfreed in the block.
  You should either free them individually using vmaVirtualFree() or call vmaClearVirtualBlock()
  if you are sure this is what you want. If you do neither, an assert is called.

  If you keep pointers to some additional metadata associated with your virtual allocations in their `pUserData`,
  don't forget to free them.
*/
func DestroyVirtualBlock(virtualBlock VirtualBlock) {
	C.vmaDestroyVirtualBlock(C.VmaVirtualBlock(unsafe.Pointer(virtualBlock)))
}

// EndDefragmentation wraps vmaEndDefragmentation.
/*
  \brief Ends defragmentation process.

  \param allocator Allocator object.
  \param context Context object that has been created by vmaBeginDefragmentation().
  \param[out] pStats Optional stats for the defragmentation. Can be null.

  Use this function to finish defragmentation started by vmaBeginDefragmentation().
*/
func EndDefragmentation(allocator Allocator, context DefragmentationContext, pStats DefragmentationStats) {
	C.vmaEndDefragmentation(C.VmaAllocator(unsafe.Pointer(allocator)), C.VmaDefragmentationContext(unsafe.Pointer(context)), (*C.VmaDefragmentationStats)(unsafe.Pointer(pStats)))
}

// EndDefragmentationPass wraps vmaEndDefragmentationPass.
/*
  \brief Ends single defragmentation pass.

  \param allocator Allocator object.
  \param context Context object that has been created by vmaBeginDefragmentation().
  \param pPassInfo Computed information for current pass filled by vmaBeginDefragmentationPass() and possibly modified by you.

  Returns `VK_SUCCESS` if no more moves are possible or `VK_INCOMPLETE` if more defragmentations are possible.

  Ends incremental defragmentation pass and commits all defragmentation moves from `pPassInfo`.
  After this call:

  - Allocations at `pPassInfo[i].srcAllocation` that had `pPassInfo[i].operation ==` #VMA_DEFRAGMENTATION_MOVE_OPERATION_COPY
  (which is the default) will be pointing to the new destination place.
  - Allocation at `pPassInfo[i].srcAllocation` that had `pPassInfo[i].operation ==` #VMA_DEFRAGMENTATION_MOVE_OPERATION_DESTROY
  will be freed.

  If no more moves are possible you can end whole defragmentation.
*/
func EndDefragmentationPass(allocator Allocator, context DefragmentationContext, pPassInfo DefragmentationPassMoveInfo) vk.Result {
	ret := C.vmaEndDefragmentationPass(C.VmaAllocator(unsafe.Pointer(allocator)), C.VmaDefragmentationContext(unsafe.Pointer(context)), (*C.VmaDefragmentationPassMoveInfo)(unsafe.Pointer(pPassInfo)))

	return vk.Result(ret)
}

// FindMemoryTypeIndex wraps vmaFindMemoryTypeIndex.
/*
  \brief Helps to find `memoryTypeIndex`, given `memoryTypeBits` and #VmaAllocationCreateInfo.

  This algorithm tries to find a memory type that:

  - Is allowed by `memoryTypeBits`.
  - Contains all the flags from `pAllocationCreateInfo->requiredFlags`.
  - Matches intended usage.
  - Has as many flags from `pAllocationCreateInfo->preferredFlags` as possible.

  \return Returns `VK_ERROR_FEATURE_NOT_PRESENT` if not found. Receiving such result
  from this function or any other allocating function probably means that your
  device doesn't support any memory type with requested features for the specific
  type of resource you want to use it for. Please check parameters of your
  resource, like image layout (`OPTIMAL` versus `LINEAR`) or mip level count.
*/
func FindMemoryTypeIndex(allocator Allocator, memoryTypeBits uint32, pAllocationCreateInfo AllocationCreateInfo, pMemoryTypeIndex ffi.Ref[uint32]) vk.Result {
	ret := C.vmaFindMemoryTypeIndex(C.VmaAllocator(unsafe.Pointer(allocator)), C.uint32_t(memoryTypeBits), (*C.VmaAllocationCreateInfo)(unsafe.Pointer(pAllocationCreateInfo)), (*C.uint32_t)(pMemoryTypeIndex.Raw()))

	return vk.Result(ret)
}

// FindMemoryTypeIndexForBufferInfo wraps vmaFindMemoryTypeIndexForBufferInfo.
/*
  \brief Helps to find `memoryTypeIndex`, given `VkBufferCreateInfo` and #VmaAllocationCreateInfo.

  It can be useful e.g. to determine value to be used as VmaPoolCreateInfo::memoryTypeIndex.
  It may need to internally create a temporary, dummy buffer that never has memory bound.
*/
func FindMemoryTypeIndexForBufferInfo(allocator Allocator, pBufferCreateInfo vk.BufferCreateInfo, pAllocationCreateInfo AllocationCreateInfo, pMemoryTypeIndex ffi.Ref[uint32]) vk.Result {
	ret := C.vmaFindMemoryTypeIndexForBufferInfo(C.VmaAllocator(unsafe.Pointer(allocator)), (*C.VkBufferCreateInfo)(unsafe.Pointer(pBufferCreateInfo)), (*C.VmaAllocationCreateInfo)(unsafe.Pointer(pAllocationCreateInfo)), (*C.uint32_t)(pMemoryTypeIndex.Raw()))

	return vk.Result(ret)
}

// FindMemoryTypeIndexForImageInfo wraps vmaFindMemoryTypeIndexForImageInfo.
/*
  \brief Helps to find `memoryTypeIndex`, given `VkImageCreateInfo` and #VmaAllocationCreateInfo.

  It can be useful e.g. to determine value to be used as VmaPoolCreateInfo::memoryTypeIndex.
  It may need to internally create a temporary, dummy image that never has memory bound.
*/
func FindMemoryTypeIndexForImageInfo(allocator Allocator, pImageCreateInfo vk.ImageCreateInfo, pAllocationCreateInfo AllocationCreateInfo, pMemoryTypeIndex ffi.Ref[uint32]) vk.Result {
	ret := C.vmaFindMemoryTypeIndexForImageInfo(C.VmaAllocator(unsafe.Pointer(allocator)), (*C.VkImageCreateInfo)(unsafe.Pointer(pImageCreateInfo)), (*C.VmaAllocationCreateInfo)(unsafe.Pointer(pAllocationCreateInfo)), (*C.uint32_t)(pMemoryTypeIndex.Raw()))

	return vk.Result(ret)
}

// FlushAllocation wraps vmaFlushAllocation.
/*
  \brief Flushes memory of given allocation.

  Calls `vkFlushMappedMemoryRanges()` for memory associated with given range of given allocation.
  It needs to be called after writing to a mapped memory for memory types that are not `HOST_COHERENT`.
  Unmap operation doesn't do that automatically.

  - `offset` must be relative to the beginning of allocation.
  - `size` can be `VK_WHOLE_SIZE`. It means all memory from `offset` the the end of given allocation.
  - `offset` and `size` don't have to be aligned.
  They are internally rounded down/up to multiply of `nonCoherentAtomSize`.
  - If `size` is 0, this call is ignored.
  - If memory type that the `allocation` belongs to is not `HOST_VISIBLE` or it is `HOST_COHERENT`,
  this call is ignored.

  Warning! `offset` and `size` are relative to the contents of given `allocation`.
  If you mean whole allocation, you can pass 0 and `VK_WHOLE_SIZE`, respectively.
  Do not pass allocation's offset as `offset`!!!

  This function returns the `VkResult` from `vkFlushMappedMemoryRanges` if it is
  called, otherwise `VK_SUCCESS`.
*/
func FlushAllocation(allocator Allocator, allocation Allocation, offset vk.DeviceSize, size vk.DeviceSize) vk.Result {
	ret := C.vmaFlushAllocation(C.VmaAllocator(unsafe.Pointer(allocator)), C.VmaAllocation(unsafe.Pointer(allocation)), C.VkDeviceSize(offset), C.VkDeviceSize(size))

	return vk.Result(ret)
}

// FlushAllocations wraps vmaFlushAllocations.
/*
  \brief Flushes memory of given set of allocations.

  Calls `vkFlushMappedMemoryRanges()` for memory associated with given ranges of given allocations.
  For more information, see documentation of vmaFlushAllocation().

  \param allocator
  \param allocationCount
  \param allocations
  \param offsets If not null, it must point to an array of offsets of regions to flush, relative to the beginning of respective allocations. Null means all offsets are zero.
  \param sizes If not null, it must point to an array of sizes of regions to flush in respective allocations. Null means `VK_WHOLE_SIZE` for all allocations.

  This function returns the `VkResult` from `vkFlushMappedMemoryRanges` if it is
  called, otherwise `VK_SUCCESS`.
*/
func FlushAllocations(allocator Allocator, allocationCount uint32, allocations ffi.Ref[Allocation], offsets ffi.Ref[vk.DeviceSize], sizes ffi.Ref[vk.DeviceSize]) vk.Result {
	ret := C.vmaFlushAllocations(C.VmaAllocator(unsafe.Pointer(allocator)), C.uint32_t(allocationCount), (*C.VmaAllocation)(allocations.Raw()), (*C.VkDeviceSize)(offsets.Raw()), (*C.VkDeviceSize)(sizes.Raw()))

	return vk.Result(ret)
}

// FreeMemory wraps vmaFreeMemory.
/*
  \brief Frees memory previously allocated using vmaAllocateMemory(), vmaAllocateMemoryForBuffer(), or vmaAllocateMemoryForImage().

  Passing `VK_NULL_HANDLE` as `allocation` is valid. Such function call is just skipped.
*/
func FreeMemory(allocator Allocator, allocation Allocation) {
	C.vmaFreeMemory(C.VmaAllocator(unsafe.Pointer(allocator)), C.VmaAllocation(unsafe.Pointer(allocation)))
}

// FreeMemoryPages wraps vmaFreeMemoryPages.
/*
  \brief Frees memory and destroys multiple allocations.

  Word "pages" is just a suggestion to use this function to free pieces of memory used for sparse binding.
  It is just a general purpose function to free memory and destroy allocations made using e.g. vmaAllocateMemory(),
  vmaAllocateMemoryPages() and other functions.
  It may be internally optimized to be more efficient than calling vmaFreeMemory() `allocationCount` times.

  Allocations in `pAllocations` array can come from any memory pools and types.
  Passing `VK_NULL_HANDLE` as elements of `pAllocations` array is valid. Such entries are just skipped.
*/
func FreeMemoryPages(allocator Allocator, allocationCount uintptr, pAllocations ffi.Ref[Allocation]) {
	C.vmaFreeMemoryPages(C.VmaAllocator(unsafe.Pointer(allocator)), C.size_t(allocationCount), (*C.VmaAllocation)(pAllocations.Raw()))
}

// FreeStatsString wraps vmaFreeStatsString.
func FreeStatsString(allocator Allocator, pStatsString ffi.CString) {
	C.vmaFreeStatsString(C.VmaAllocator(unsafe.Pointer(allocator)), (*C.char)(pStatsString.Raw()))
}

// FreeVirtualBlockStatsString wraps vmaFreeVirtualBlockStatsString.
//
//	Frees a string returned by vmaBuildVirtualBlockStatsString().
func FreeVirtualBlockStatsString(virtualBlock VirtualBlock, pStatsString ffi.CString) {
	C.vmaFreeVirtualBlockStatsString(C.VmaVirtualBlock(unsafe.Pointer(virtualBlock)), (*C.char)(pStatsString.Raw()))
}

// GetAllocationInfo wraps vmaGetAllocationInfo.
/*
  \brief Returns current information about specified allocation.

  Current parameters of given allocation are returned in `pAllocationInfo`.

  Although this function doesn't lock any mutex, so it should be quite efficient,
  you should avoid calling it too often.
  You can retrieve same VmaAllocationInfo structure while creating your resource, from function
  vmaCreateBuffer(), vmaCreateImage(). You can remember it if you are sure parameters don't change
  (e.g. due to defragmentation).

  There is also a new function vmaGetAllocationInfo2() that offers extended information
  about the allocation, returned using new structure #VmaAllocationInfo2.
*/
func GetAllocationInfo(allocator Allocator, allocation Allocation, pAllocationInfo AllocationInfo) {
	C.vmaGetAllocationInfo(C.VmaAllocator(unsafe.Pointer(allocator)), C.VmaAllocation(unsafe.Pointer(allocation)), (*C.VmaAllocationInfo)(unsafe.Pointer(pAllocationInfo)))
}

// GetAllocationInfo2 wraps vmaGetAllocationInfo2.
/*
  \brief Returns extended information about specified allocation.

  Current parameters of given allocation are returned in `pAllocationInfo`.
  Extended parameters in structure #VmaAllocationInfo2 include memory block size
  and a flag telling whether the allocation has dedicated memory.
  It can be useful e.g. for interop with OpenGL.
*/
func GetAllocationInfo2(allocator Allocator, allocation Allocation, pAllocationInfo AllocationInfo2) {
	C.vmaGetAllocationInfo2(C.VmaAllocator(unsafe.Pointer(allocator)), C.VmaAllocation(unsafe.Pointer(allocation)), (*C.VmaAllocationInfo2)(unsafe.Pointer(pAllocationInfo)))
}

// GetAllocationMemoryProperties wraps vmaGetAllocationMemoryProperties.
/*
  \brief Given an allocation, returns Property Flags of its memory type.

  This is just a convenience function. Same information can be obtained using
  vmaGetAllocationInfo() + vmaGetMemoryProperties().
*/
func GetAllocationMemoryProperties(allocator Allocator, allocation Allocation, pFlags ffi.Ref[vk.MemoryPropertyFlags]) {
	C.vmaGetAllocationMemoryProperties(C.VmaAllocator(unsafe.Pointer(allocator)), C.VmaAllocation(unsafe.Pointer(allocation)), (*C.VkMemoryPropertyFlags)(pFlags.Raw()))
}

// GetAllocatorInfo wraps vmaGetAllocatorInfo.
/*
  \brief Returns information about existing #VmaAllocator object - handle to Vulkan device etc.

  It might be useful if you want to keep just the #VmaAllocator handle and fetch other required handles to
  `VkPhysicalDevice`, `VkDevice` etc. every time using this function.
*/
func GetAllocatorInfo(allocator Allocator, pAllocatorInfo AllocatorInfo) {
	C.vmaGetAllocatorInfo(C.VmaAllocator(unsafe.Pointer(allocator)), (*C.VmaAllocatorInfo)(unsafe.Pointer(pAllocatorInfo)))
}

// GetHeapBudgets wraps vmaGetHeapBudgets.
/*
  \brief Retrieves information about current memory usage and budget for all memory heaps.

  \param allocator
  \param[out] pBudgets Must point to array with number of elements at least equal to number of memory heaps in physical device used.

  This function is called "get" not "calculate" because it is very fast, suitable to be called
  every frame or every allocation. For more detailed statistics use vmaCalculateStatistics().

  Note that when using allocator from multiple threads, returned information may immediately
  become outdated.
*/
func GetHeapBudgets(allocator Allocator, pBudgets Budget) {
	C.vmaGetHeapBudgets(C.VmaAllocator(unsafe.Pointer(allocator)), (*C.VmaBudget)(unsafe.Pointer(pBudgets)))
}

// vmaGetMemoryProperties.ppPhysicalDeviceMemoryProperties is unsupported: category pointer2 type VkPhysicalDeviceMemoryProperties.

// GetMemoryTypeProperties wraps vmaGetMemoryTypeProperties.
/*
  \brief Given Memory Type Index, returns Property Flags of this memory type.

  This is just a convenience function. Same information can be obtained using
  vmaGetMemoryProperties().
*/
func GetMemoryTypeProperties(allocator Allocator, memoryTypeIndex uint32, pFlags ffi.Ref[vk.MemoryPropertyFlags]) {
	C.vmaGetMemoryTypeProperties(C.VmaAllocator(unsafe.Pointer(allocator)), C.uint32_t(memoryTypeIndex), (*C.VkMemoryPropertyFlags)(pFlags.Raw()))
}

// vmaGetPhysicalDeviceProperties.ppPhysicalDeviceProperties is unsupported: category pointer2 type VkPhysicalDeviceProperties.

// GetPoolName wraps vmaGetPoolName.
/*
  \brief Retrieves name of a custom pool.

  After the call `ppName` is either null or points to an internally-owned null-terminated string
  containing name of the pool that was previously set. The pointer becomes invalid when the pool is
  destroyed or its name is changed using vmaSetPoolName().
*/
func GetPoolName(allocator Allocator, pool Pool, ppName ffi.Ref[ffi.CString]) {
	C.vmaGetPoolName(C.VmaAllocator(unsafe.Pointer(allocator)), C.VmaPool(unsafe.Pointer(pool)), (**C.char)(ppName.Raw()))
}

// GetPoolStatistics wraps vmaGetPoolStatistics.
/*
  \brief Retrieves statistics of existing #VmaPool object.

  \param allocator Allocator object.
  \param pool Pool object.
  \param[out] pPoolStats Statistics of specified pool.

  Note that when using the pool from multiple threads, returned information may immediately
  become outdated.
*/
func GetPoolStatistics(allocator Allocator, pool Pool, pPoolStats Statistics) {
	C.vmaGetPoolStatistics(C.VmaAllocator(unsafe.Pointer(allocator)), C.VmaPool(unsafe.Pointer(pool)), (*C.VmaStatistics)(unsafe.Pointer(pPoolStats)))
}

// GetVirtualAllocationInfo wraps vmaGetVirtualAllocationInfo.
//
//	\brief Returns information about a specific virtual allocation within a virtual block, like its size and `pUserData` pointer.
func GetVirtualAllocationInfo(virtualBlock VirtualBlock, allocation VirtualAllocation, pVirtualAllocInfo VirtualAllocationInfo) {
	C.vmaGetVirtualAllocationInfo(C.VmaVirtualBlock(unsafe.Pointer(virtualBlock)), C.VmaVirtualAllocation(allocation), (*C.VmaVirtualAllocationInfo)(unsafe.Pointer(pVirtualAllocInfo)))
}

// GetVirtualBlockStatistics wraps vmaGetVirtualBlockStatistics.
/*
  \brief Calculates and returns statistics about virtual allocations and memory usage in given #VmaVirtualBlock.

  This function is fast to call. For more detailed statistics, see vmaCalculateVirtualBlockStatistics().
*/
func GetVirtualBlockStatistics(virtualBlock VirtualBlock, pStats Statistics) {
	C.vmaGetVirtualBlockStatistics(C.VmaVirtualBlock(unsafe.Pointer(virtualBlock)), (*C.VmaStatistics)(unsafe.Pointer(pStats)))
}

// InvalidateAllocation wraps vmaInvalidateAllocation.
/*
  \brief Invalidates memory of given allocation.

  Calls `vkInvalidateMappedMemoryRanges()` for memory associated with given range of given allocation.
  It needs to be called before reading from a mapped memory for memory types that are not `HOST_COHERENT`.
  Map operation doesn't do that automatically.

  - `offset` must be relative to the beginning of allocation.
  - `size` can be `VK_WHOLE_SIZE`. It means all memory from `offset` the the end of given allocation.
  - `offset` and `size` don't have to be aligned.
  They are internally rounded down/up to multiply of `nonCoherentAtomSize`.
  - If `size` is 0, this call is ignored.
  - If memory type that the `allocation` belongs to is not `HOST_VISIBLE` or it is `HOST_COHERENT`,
  this call is ignored.

  Warning! `offset` and `size` are relative to the contents of given `allocation`.
  If you mean whole allocation, you can pass 0 and `VK_WHOLE_SIZE`, respectively.
  Do not pass allocation's offset as `offset`!!!

  This function returns the `VkResult` from `vkInvalidateMappedMemoryRanges` if
  it is called, otherwise `VK_SUCCESS`.
*/
func InvalidateAllocation(allocator Allocator, allocation Allocation, offset vk.DeviceSize, size vk.DeviceSize) vk.Result {
	ret := C.vmaInvalidateAllocation(C.VmaAllocator(unsafe.Pointer(allocator)), C.VmaAllocation(unsafe.Pointer(allocation)), C.VkDeviceSize(offset), C.VkDeviceSize(size))

	return vk.Result(ret)
}

// InvalidateAllocations wraps vmaInvalidateAllocations.
/*
  \brief Invalidates memory of given set of allocations.

  Calls `vkInvalidateMappedMemoryRanges()` for memory associated with given ranges of given allocations.
  For more information, see documentation of vmaInvalidateAllocation().

  \param allocator
  \param allocationCount
  \param allocations
  \param offsets If not null, it must point to an array of offsets of regions to flush, relative to the beginning of respective allocations. Null means all offsets are zero.
  \param sizes If not null, it must point to an array of sizes of regions to flush in respective allocations. Null means `VK_WHOLE_SIZE` for all allocations.

  This function returns the `VkResult` from `vkInvalidateMappedMemoryRanges` if it is
  called, otherwise `VK_SUCCESS`.
*/
func InvalidateAllocations(allocator Allocator, allocationCount uint32, allocations ffi.Ref[Allocation], offsets ffi.Ref[vk.DeviceSize], sizes ffi.Ref[vk.DeviceSize]) vk.Result {
	ret := C.vmaInvalidateAllocations(C.VmaAllocator(unsafe.Pointer(allocator)), C.uint32_t(allocationCount), (*C.VmaAllocation)(allocations.Raw()), (*C.VkDeviceSize)(offsets.Raw()), (*C.VkDeviceSize)(sizes.Raw()))

	return vk.Result(ret)
}

// IsVirtualBlockEmpty wraps vmaIsVirtualBlockEmpty.
//
//	\brief Returns true of the #VmaVirtualBlock is empty - contains 0 virtual allocations and has all its space available for new allocations.
func IsVirtualBlockEmpty(virtualBlock VirtualBlock) bool {
	ret := C.vmaIsVirtualBlockEmpty(C.VmaVirtualBlock(unsafe.Pointer(virtualBlock)))

	if ret > 0 {
		return true
	} else {
		return false
	}
}

// MapMemory wraps vmaMapMemory.
/*
  \brief Maps memory represented by given allocation and returns pointer to it.

  Maps memory represented by given allocation to make it accessible to CPU code.
  When succeeded, `*ppData` contains pointer to first byte of this memory.

  \warning
  If the allocation is part of a bigger `VkDeviceMemory` block, returned pointer is
  correctly offsetted to the beginning of region assigned to this particular allocation.
  Unlike the result of `vkMapMemory`, it points to the allocation, not to the beginning of the whole block.
  You should not add VmaAllocationInfo::offset to it!

  Mapping is internally reference-counted and synchronized, so despite raw Vulkan
  function `vkMapMemory()` cannot be used to map same block of `VkDeviceMemory`
  multiple times simultaneously, it is safe to call this function on allocations
  assigned to the same memory block. Actual Vulkan memory will be mapped on first
  mapping and unmapped on last unmapping.

  If the function succeeded, you must call vmaUnmapMemory() to unmap the
  allocation when mapping is no longer needed or before freeing the allocation, at
  the latest.

  It also safe to call this function multiple times on the same allocation. You
  must call vmaUnmapMemory() same number of times as you called vmaMapMemory().

  It is also safe to call this function on allocation created with
  #VMA_ALLOCATION_CREATE_MAPPED_BIT flag. Its memory stays mapped all the time.
  You must still call vmaUnmapMemory() same number of times as you called
  vmaMapMemory(). You must not call vmaUnmapMemory() additional time to free the
  "0-th" mapping made automatically due to #VMA_ALLOCATION_CREATE_MAPPED_BIT flag.

  This function fails when used on allocation made in memory type that is not
  `HOST_VISIBLE`.

  This function doesn't automatically flush or invalidate caches.
  If the allocation is made from a memory types that is not `HOST_COHERENT`,
  you also need to use vmaInvalidateAllocation() / vmaFlushAllocation(), as required by Vulkan specification.
*/
func MapMemory(allocator Allocator, allocation Allocation, ppData ffi.Ref[uintptr]) vk.Result {
	ret := C.vmaMapMemory(C.VmaAllocator(unsafe.Pointer(allocator)), C.VmaAllocation(unsafe.Pointer(allocation)), (*unsafe.Pointer)(ppData.Raw()))

	return vk.Result(ret)
}

// SetAllocationName wraps vmaSetAllocationName.
/*
  \brief Sets pName in given allocation to new value.

  `pName` must be either null, or pointer to a null-terminated string. The function
  makes local copy of the string and sets it as allocation's `pName`. String
  passed as pName doesn't need to be valid for whole lifetime of the allocation -
  you can free it after this call. String previously pointed by allocation's
  `pName` is freed from memory.
*/
func SetAllocationName(allocator Allocator, allocation Allocation, pName ffi.CString) {
	C.vmaSetAllocationName(C.VmaAllocator(unsafe.Pointer(allocator)), C.VmaAllocation(unsafe.Pointer(allocation)), (*C.char)(pName.Raw()))
}

// SetAllocationUserData wraps vmaSetAllocationUserData.
/*
  \brief Sets pUserData in given allocation to new value.

  The value of pointer `pUserData` is copied to allocation's `pUserData`.
  It is opaque, so you can use it however you want - e.g.
  as a pointer, ordinal number or some handle to you own data.
*/
func SetAllocationUserData(allocator Allocator, allocation Allocation, pUserData uintptr) {
	C.vmaSetAllocationUserData(C.VmaAllocator(unsafe.Pointer(allocator)), C.VmaAllocation(unsafe.Pointer(allocation)), unsafe.Pointer(pUserData))
}

// SetCurrentFrameIndex wraps vmaSetCurrentFrameIndex.
//
//	\brief Sets index of the current frame.
func SetCurrentFrameIndex(allocator Allocator, frameIndex uint32) {
	C.vmaSetCurrentFrameIndex(C.VmaAllocator(unsafe.Pointer(allocator)), C.uint32_t(frameIndex))
}

// SetPoolName wraps vmaSetPoolName.
/*
  \brief Sets name of a custom pool.

  `pName` can be either null or pointer to a null-terminated string with new name for the pool.
  Function makes internal copy of the string, so it can be changed or freed immediately after this call.
*/
func SetPoolName(allocator Allocator, pool Pool, pName ffi.CString) {
	C.vmaSetPoolName(C.VmaAllocator(unsafe.Pointer(allocator)), C.VmaPool(unsafe.Pointer(pool)), (*C.char)(pName.Raw()))
}

// SetVirtualAllocationUserData wraps vmaSetVirtualAllocationUserData.
//
//	\brief Changes custom pointer associated with given virtual allocation.
func SetVirtualAllocationUserData(virtualBlock VirtualBlock, allocation VirtualAllocation, pUserData uintptr) {
	C.vmaSetVirtualAllocationUserData(C.VmaVirtualBlock(unsafe.Pointer(virtualBlock)), C.VmaVirtualAllocation(allocation), unsafe.Pointer(pUserData))
}

// UnmapMemory wraps vmaUnmapMemory.
/*
  \brief Unmaps memory represented by given allocation, mapped previously using vmaMapMemory().

  For details, see description of vmaMapMemory().

  This function doesn't automatically flush or invalidate caches.
  If the allocation is made from a memory types that is not `HOST_COHERENT`,
  you also need to use vmaInvalidateAllocation() / vmaFlushAllocation(), as required by Vulkan specification.
*/
func UnmapMemory(allocator Allocator, allocation Allocation) {
	C.vmaUnmapMemory(C.VmaAllocator(unsafe.Pointer(allocator)), C.VmaAllocation(unsafe.Pointer(allocation)))
}

// VirtualAllocate wraps vmaVirtualAllocate.
/*
  \brief Allocates new virtual allocation inside given #VmaVirtualBlock.

  If the allocation fails due to not enough free space available, `VK_ERROR_OUT_OF_DEVICE_MEMORY` is returned
  (despite the function doesn't ever allocate actual GPU memory).
  `pAllocation` is then set to `VK_NULL_HANDLE` and `pOffset`, if not null, it set to `UINT64_MAX`.

  \param virtualBlock Virtual block
  \param pCreateInfo Parameters for the allocation
  \param[out] pAllocation Returned handle of the new allocation
  \param[out] pOffset Returned offset of the new allocation. Optional, can be null.
*/
func VirtualAllocate(virtualBlock VirtualBlock, pCreateInfo VirtualAllocationCreateInfo, pAllocation ffi.Ref[VirtualAllocation], pOffset ffi.Ref[vk.DeviceSize]) vk.Result {
	ret := C.vmaVirtualAllocate(C.VmaVirtualBlock(unsafe.Pointer(virtualBlock)), (*C.VmaVirtualAllocationCreateInfo)(unsafe.Pointer(pCreateInfo)), (*C.VmaVirtualAllocation)(pAllocation.Raw()), (*C.VkDeviceSize)(pOffset.Raw()))

	return vk.Result(ret)
}

// VirtualFree wraps vmaVirtualFree.
/*
  \brief Frees virtual allocation inside given #VmaVirtualBlock.

  It is correct to call this function with `allocation == VK_NULL_HANDLE` - it does nothing.
*/
func VirtualFree(virtualBlock VirtualBlock, allocation VirtualAllocation) {
	C.vmaVirtualFree(C.VmaVirtualBlock(unsafe.Pointer(virtualBlock)), C.VmaVirtualAllocation(allocation))
}
