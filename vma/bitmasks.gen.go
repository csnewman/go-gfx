package vma

import "fmt"

// #include "vma.h"
import "C"

// AllocationCreateFlags wraps the bitmask VmaAllocationCreateFlags.
//
//	Flags to be passed as VmaAllocationCreateInfo::flags.
type AllocationCreateFlags int32

const (
	// ALLOCATION_CREATE_CAN_ALIAS_BIT wraps VMA_ALLOCATION_CREATE_CAN_ALIAS_BIT.
	/*
	   \brief Set this flag if the allocated memory will have aliasing resources.

	   Usage of this flag prevents supplying `VkMemoryDedicatedAllocateInfoKHR` when #VMA_ALLOCATION_CREATE_DEDICATED_MEMORY_BIT is specified.
	   Otherwise created dedicated memory will not be suitable for aliasing resources, resulting in Vulkan Validation Layer errors.
	*/
	ALLOCATION_CREATE_CAN_ALIAS_BIT AllocationCreateFlags = C.VMA_ALLOCATION_CREATE_CAN_ALIAS_BIT
	// ALLOCATION_CREATE_DEDICATED_MEMORY_BIT wraps VMA_ALLOCATION_CREATE_DEDICATED_MEMORY_BIT.
	/*
	   \brief Set this flag if the allocation should have its own memory block.

	   Use it for special, big resources, like fullscreen images used as attachments.

	   If you use this flag while creating a buffer or an image, `VkMemoryDedicatedAllocateInfo`
	   structure is applied if possible.
	*/
	ALLOCATION_CREATE_DEDICATED_MEMORY_BIT AllocationCreateFlags = C.VMA_ALLOCATION_CREATE_DEDICATED_MEMORY_BIT
	// ALLOCATION_CREATE_DONT_BIND_BIT wraps VMA_ALLOCATION_CREATE_DONT_BIND_BIT.
	/*
	   Create both buffer/image and allocation, but don't bind them together.
	   It is useful when you want to bind yourself to do some more advanced binding, e.g. using some extensions.
	   The flag is meaningful only with functions that bind by default: vmaCreateBuffer(), vmaCreateImage().
	   Otherwise it is ignored.

	   If you want to make sure the new buffer/image is not tied to the new memory allocation
	   through `VkMemoryDedicatedAllocateInfoKHR` structure in case the allocation ends up in its own memory block,
	   use also flag #VMA_ALLOCATION_CREATE_CAN_ALIAS_BIT.
	*/
	ALLOCATION_CREATE_DONT_BIND_BIT AllocationCreateFlags = C.VMA_ALLOCATION_CREATE_DONT_BIND_BIT
	// ALLOCATION_CREATE_FLAG_BITS_MAX_ENUM wraps VMA_ALLOCATION_CREATE_FLAG_BITS_MAX_ENUM.
	//
	//	A bit mask to extract only `STRATEGY` bits from entire set of flags.
	ALLOCATION_CREATE_FLAG_BITS_MAX_ENUM AllocationCreateFlags = C.VMA_ALLOCATION_CREATE_FLAG_BITS_MAX_ENUM
	// ALLOCATION_CREATE_HOST_ACCESS_ALLOW_TRANSFER_INSTEAD_BIT wraps VMA_ALLOCATION_CREATE_HOST_ACCESS_ALLOW_TRANSFER_INSTEAD_BIT.
	/*
	   Together with #VMA_ALLOCATION_CREATE_HOST_ACCESS_SEQUENTIAL_WRITE_BIT or #VMA_ALLOCATION_CREATE_HOST_ACCESS_RANDOM_BIT,
	   it says that despite request for host access, a not-`HOST_VISIBLE` memory type can be selected
	   if it may improve performance.

	   By using this flag, you declare that you will check if the allocation ended up in a `HOST_VISIBLE` memory type
	   (e.g. using vmaGetAllocationMemoryProperties()) and if not, you will create some "staging" buffer and
	   issue an explicit transfer to write/read your data.
	   To prepare for this possibility, don't forget to add appropriate flags like
	   `VK_BUFFER_USAGE_TRANSFER_DST_BIT`, `VK_BUFFER_USAGE_TRANSFER_SRC_BIT` to the parameters of created buffer or image.
	*/
	ALLOCATION_CREATE_HOST_ACCESS_ALLOW_TRANSFER_INSTEAD_BIT AllocationCreateFlags = C.VMA_ALLOCATION_CREATE_HOST_ACCESS_ALLOW_TRANSFER_INSTEAD_BIT
	// ALLOCATION_CREATE_HOST_ACCESS_RANDOM_BIT wraps VMA_ALLOCATION_CREATE_HOST_ACCESS_RANDOM_BIT.
	/*
	   Requests possibility to map the allocation (using vmaMapMemory() or #VMA_ALLOCATION_CREATE_MAPPED_BIT).

	   - If you use #VMA_MEMORY_USAGE_AUTO or other `VMA_MEMORY_USAGE_AUTO*` value,
	   you must use this flag to be able to map the allocation. Otherwise, mapping is incorrect.
	   - If you use other value of #VmaMemoryUsage, this flag is ignored and mapping is always possible in memory types that are `HOST_VISIBLE`.
	   This includes allocations created in \ref custom_memory_pools.

	   Declares that mapped memory can be read, written, and accessed in random order,
	   so a `HOST_CACHED` memory type is preferred.
	*/
	ALLOCATION_CREATE_HOST_ACCESS_RANDOM_BIT AllocationCreateFlags = C.VMA_ALLOCATION_CREATE_HOST_ACCESS_RANDOM_BIT
	// ALLOCATION_CREATE_HOST_ACCESS_SEQUENTIAL_WRITE_BIT wraps VMA_ALLOCATION_CREATE_HOST_ACCESS_SEQUENTIAL_WRITE_BIT.
	/*
	   Requests possibility to map the allocation (using vmaMapMemory() or #VMA_ALLOCATION_CREATE_MAPPED_BIT).

	   - If you use #VMA_MEMORY_USAGE_AUTO or other `VMA_MEMORY_USAGE_AUTO*` value,
	   you must use this flag to be able to map the allocation. Otherwise, mapping is incorrect.
	   - If you use other value of #VmaMemoryUsage, this flag is ignored and mapping is always possible in memory types that are `HOST_VISIBLE`.
	   This includes allocations created in \ref custom_memory_pools.

	   Declares that mapped memory will only be written sequentially, e.g. using `memcpy()` or a loop writing number-by-number,
	   never read or accessed randomly, so a memory type can be selected that is uncached and write-combined.

	   \warning Violating this declaration may work correctly, but will likely be very slow.
	   Watch out for implicit reads introduced by doing e.g. `pMappedData[i] += x;`
	   Better prepare your data in a local variable and `memcpy()` it to the mapped pointer all at once.
	*/
	ALLOCATION_CREATE_HOST_ACCESS_SEQUENTIAL_WRITE_BIT AllocationCreateFlags = C.VMA_ALLOCATION_CREATE_HOST_ACCESS_SEQUENTIAL_WRITE_BIT
	// ALLOCATION_CREATE_MAPPED_BIT wraps VMA_ALLOCATION_CREATE_MAPPED_BIT.
	/*
	   \brief Set this flag to use a memory that will be persistently mapped and retrieve pointer to it.

	   Pointer to mapped memory will be returned through VmaAllocationInfo::pMappedData.

	   It is valid to use this flag for allocation made from memory type that is not
	   `HOST_VISIBLE`. This flag is then ignored and memory is not mapped. This is
	   useful if you need an allocation that is efficient to use on GPU
	   (`DEVICE_LOCAL`) and still want to map it directly if possible on platforms that
	   support it (e.g. Intel GPU).
	*/
	ALLOCATION_CREATE_MAPPED_BIT AllocationCreateFlags = C.VMA_ALLOCATION_CREATE_MAPPED_BIT
	// ALLOCATION_CREATE_NEVER_ALLOCATE_BIT wraps VMA_ALLOCATION_CREATE_NEVER_ALLOCATE_BIT.
	/*
	   \brief Set this flag to only try to allocate from existing `VkDeviceMemory` blocks and never create new such block.

	   If new allocation cannot be placed in any of the existing blocks, allocation
	   fails with `VK_ERROR_OUT_OF_DEVICE_MEMORY` error.

	   You should not use #VMA_ALLOCATION_CREATE_DEDICATED_MEMORY_BIT and
	   #VMA_ALLOCATION_CREATE_NEVER_ALLOCATE_BIT at the same time. It makes no sense.
	*/
	ALLOCATION_CREATE_NEVER_ALLOCATE_BIT AllocationCreateFlags = C.VMA_ALLOCATION_CREATE_NEVER_ALLOCATE_BIT
	// ALLOCATION_CREATE_STRATEGY_BEST_FIT_BIT wraps VMA_ALLOCATION_CREATE_STRATEGY_BEST_FIT_BIT.
	//
	//	Alias to #VMA_ALLOCATION_CREATE_STRATEGY_MIN_MEMORY_BIT.
	ALLOCATION_CREATE_STRATEGY_BEST_FIT_BIT AllocationCreateFlags = C.VMA_ALLOCATION_CREATE_STRATEGY_BEST_FIT_BIT
	// ALLOCATION_CREATE_STRATEGY_FIRST_FIT_BIT wraps VMA_ALLOCATION_CREATE_STRATEGY_FIRST_FIT_BIT.
	//
	//	Alias to #VMA_ALLOCATION_CREATE_STRATEGY_MIN_TIME_BIT.
	ALLOCATION_CREATE_STRATEGY_FIRST_FIT_BIT AllocationCreateFlags = C.VMA_ALLOCATION_CREATE_STRATEGY_FIRST_FIT_BIT
	// ALLOCATION_CREATE_STRATEGY_MASK wraps VMA_ALLOCATION_CREATE_STRATEGY_MASK.
	//
	//	A bit mask to extract only `STRATEGY` bits from entire set of flags.
	ALLOCATION_CREATE_STRATEGY_MASK AllocationCreateFlags = C.VMA_ALLOCATION_CREATE_STRATEGY_MASK
	// ALLOCATION_CREATE_STRATEGY_MIN_MEMORY_BIT wraps VMA_ALLOCATION_CREATE_STRATEGY_MIN_MEMORY_BIT.
	/*
	   Allocation strategy that chooses smallest possible free range for the allocation
	   to minimize memory usage and fragmentation, possibly at the expense of allocation time.
	*/
	ALLOCATION_CREATE_STRATEGY_MIN_MEMORY_BIT AllocationCreateFlags = C.VMA_ALLOCATION_CREATE_STRATEGY_MIN_MEMORY_BIT
	// ALLOCATION_CREATE_STRATEGY_MIN_OFFSET_BIT wraps VMA_ALLOCATION_CREATE_STRATEGY_MIN_OFFSET_BIT.
	/*
	   Allocation strategy that chooses always the lowest offset in available space.
	   This is not the most efficient strategy but achieves highly packed data.
	   Used internally by defragmentation, not recommended in typical usage.
	*/
	ALLOCATION_CREATE_STRATEGY_MIN_OFFSET_BIT AllocationCreateFlags = C.VMA_ALLOCATION_CREATE_STRATEGY_MIN_OFFSET_BIT
	// ALLOCATION_CREATE_STRATEGY_MIN_TIME_BIT wraps VMA_ALLOCATION_CREATE_STRATEGY_MIN_TIME_BIT.
	/*
	   Allocation strategy that chooses first suitable free range for the allocation -
	   not necessarily in terms of the smallest offset but the one that is easiest and fastest to find
	   to minimize allocation time, possibly at the expense of allocation quality.
	*/
	ALLOCATION_CREATE_STRATEGY_MIN_TIME_BIT AllocationCreateFlags = C.VMA_ALLOCATION_CREATE_STRATEGY_MIN_TIME_BIT
	// ALLOCATION_CREATE_UPPER_ADDRESS_BIT wraps VMA_ALLOCATION_CREATE_UPPER_ADDRESS_BIT.
	/*
	   Allocation will be created from upper stack in a double stack pool.

	   This flag is only allowed for custom pools created with #VMA_POOL_CREATE_LINEAR_ALGORITHM_BIT flag.
	*/
	ALLOCATION_CREATE_UPPER_ADDRESS_BIT AllocationCreateFlags = C.VMA_ALLOCATION_CREATE_UPPER_ADDRESS_BIT
	// ALLOCATION_CREATE_USER_DATA_COPY_STRING_BIT wraps VMA_ALLOCATION_CREATE_USER_DATA_COPY_STRING_BIT.
	/*
	   \deprecated Preserved for backward compatibility. Consider using vmaSetAllocationName() instead.

	   Set this flag to treat VmaAllocationCreateInfo::pUserData as pointer to a
	   null-terminated string. Instead of copying pointer value, a local copy of the
	   string is made and stored in allocation's `pName`. The string is automatically
	   freed together with the allocation. It is also used in vmaBuildStatsString().
	*/
	ALLOCATION_CREATE_USER_DATA_COPY_STRING_BIT AllocationCreateFlags = C.VMA_ALLOCATION_CREATE_USER_DATA_COPY_STRING_BIT
	// ALLOCATION_CREATE_WITHIN_BUDGET_BIT wraps VMA_ALLOCATION_CREATE_WITHIN_BUDGET_BIT.
	/*
	   Create allocation only if additional device memory required for it, if any, won't exceed
	   memory budget. Otherwise return `VK_ERROR_OUT_OF_DEVICE_MEMORY`.
	*/
	ALLOCATION_CREATE_WITHIN_BUDGET_BIT AllocationCreateFlags = C.VMA_ALLOCATION_CREATE_WITHIN_BUDGET_BIT
)

func (e AllocationCreateFlags) String() string {
	return fmt.Sprintf("VmaAllocationCreateFlags(%b)", e)
}

// AllocatorCreateFlags wraps the bitmask VmaAllocatorCreateFlags.
//
//	Flags for created #VmaAllocator.
type AllocatorCreateFlags int32

const (
	// ALLOCATOR_CREATE_AMD_DEVICE_COHERENT_MEMORY_BIT wraps VMA_ALLOCATOR_CREATE_AMD_DEVICE_COHERENT_MEMORY_BIT.
	/*
	   Enables usage of VK_AMD_device_coherent_memory extension.

	   You may set this flag only if you:

	   - found out that this device extension is supported and enabled it while creating Vulkan device passed as VmaAllocatorCreateInfo::device,
	   - checked that `VkPhysicalDeviceCoherentMemoryFeaturesAMD::deviceCoherentMemory` is true and set it while creating the Vulkan device,
	   - want it to be used internally by this library.

	   The extension and accompanying device feature provide access to memory types with
	   `VK_MEMORY_PROPERTY_DEVICE_COHERENT_BIT_AMD` and `VK_MEMORY_PROPERTY_DEVICE_UNCACHED_BIT_AMD` flags.
	   They are useful mostly for writing breadcrumb markers - a common method for debugging GPU crash/hang/TDR.

	   When the extension is not enabled, such memory types are still enumerated, but their usage is illegal.
	   To protect from this error, if you don't create the allocator with this flag, it will refuse to allocate any memory or create a custom pool in such memory type,
	   returning `VK_ERROR_FEATURE_NOT_PRESENT`.
	*/
	ALLOCATOR_CREATE_AMD_DEVICE_COHERENT_MEMORY_BIT AllocatorCreateFlags = C.VMA_ALLOCATOR_CREATE_AMD_DEVICE_COHERENT_MEMORY_BIT
	// ALLOCATOR_CREATE_BUFFER_DEVICE_ADDRESS_BIT wraps VMA_ALLOCATOR_CREATE_BUFFER_DEVICE_ADDRESS_BIT.
	/*
	   Enables usage of "buffer device address" feature, which allows you to use function
	   `vkGetBufferDeviceAddress*` to get raw GPU pointer to a buffer and pass it for usage inside a shader.

	   You may set this flag only if you:

	   1. (For Vulkan version < 1.2) Found as available and enabled device extension
	   VK_KHR_buffer_device_address.
	   This extension is promoted to core Vulkan 1.2.
	   2. Found as available and enabled device feature `VkPhysicalDeviceBufferDeviceAddressFeatures::bufferDeviceAddress`.

	   When this flag is set, you can create buffers with `VK_BUFFER_USAGE_SHADER_DEVICE_ADDRESS_BIT` using VMA.
	   The library automatically adds `VK_MEMORY_ALLOCATE_DEVICE_ADDRESS_BIT` to
	   allocated memory blocks wherever it might be needed.

	   For more information, see documentation chapter \ref enabling_buffer_device_address.
	*/
	ALLOCATOR_CREATE_BUFFER_DEVICE_ADDRESS_BIT AllocatorCreateFlags = C.VMA_ALLOCATOR_CREATE_BUFFER_DEVICE_ADDRESS_BIT
	// ALLOCATOR_CREATE_EXTERNALLY_SYNCHRONIZED_BIT wraps VMA_ALLOCATOR_CREATE_EXTERNALLY_SYNCHRONIZED_BIT.
	/*
	   \brief Allocator and all objects created from it will not be synchronized internally, so you must guarantee they are used from only one thread at a time or synchronized externally by you.

	   Using this flag may increase performance because internal mutexes are not used.
	*/
	ALLOCATOR_CREATE_EXTERNALLY_SYNCHRONIZED_BIT AllocatorCreateFlags = C.VMA_ALLOCATOR_CREATE_EXTERNALLY_SYNCHRONIZED_BIT
	// ALLOCATOR_CREATE_EXT_MEMORY_BUDGET_BIT wraps VMA_ALLOCATOR_CREATE_EXT_MEMORY_BUDGET_BIT.
	/*
	   Enables usage of VK_EXT_memory_budget extension.

	   You may set this flag only if you found out that this device extension is supported,
	   you enabled it while creating Vulkan device passed as VmaAllocatorCreateInfo::device,
	   and you want it to be used internally by this library, along with another instance extension
	   VK_KHR_get_physical_device_properties2, which is required by it (or Vulkan 1.1, where this extension is promoted).

	   The extension provides query for current memory usage and budget, which will probably
	   be more accurate than an estimation used by the library otherwise.
	*/
	ALLOCATOR_CREATE_EXT_MEMORY_BUDGET_BIT AllocatorCreateFlags = C.VMA_ALLOCATOR_CREATE_EXT_MEMORY_BUDGET_BIT
	// ALLOCATOR_CREATE_EXT_MEMORY_PRIORITY_BIT wraps VMA_ALLOCATOR_CREATE_EXT_MEMORY_PRIORITY_BIT.
	/*
	   Enables usage of VK_EXT_memory_priority extension in the library.

	   You may set this flag only if you found available and enabled this device extension,
	   along with `VkPhysicalDeviceMemoryPriorityFeaturesEXT::memoryPriority == VK_TRUE`,
	   while creating Vulkan device passed as VmaAllocatorCreateInfo::device.

	   When this flag is used, VmaAllocationCreateInfo::priority and VmaPoolCreateInfo::priority
	   are used to set priorities of allocated Vulkan memory. Without it, these variables are ignored.

	   A priority must be a floating-point value between 0 and 1, indicating the priority of the allocation relative to other memory allocations.
	   Larger values are higher priority. The granularity of the priorities is implementation-dependent.
	   It is automatically passed to every call to `vkAllocateMemory` done by the library using structure `VkMemoryPriorityAllocateInfoEXT`.
	   The value to be used for default priority is 0.5.
	   For more details, see the documentation of the VK_EXT_memory_priority extension.
	*/
	ALLOCATOR_CREATE_EXT_MEMORY_PRIORITY_BIT AllocatorCreateFlags = C.VMA_ALLOCATOR_CREATE_EXT_MEMORY_PRIORITY_BIT
	// ALLOCATOR_CREATE_FLAG_BITS_MAX_ENUM wraps VMA_ALLOCATOR_CREATE_FLAG_BITS_MAX_ENUM.
	/*
	   Enables usage of VK_KHR_external_memory_win32 extension in the library.

	   You should set this flag if you found available and enabled this device extension,
	   while creating Vulkan device passed as VmaAllocatorCreateInfo::device.
	   For more information, see \ref other_api_interop.
	*/
	ALLOCATOR_CREATE_FLAG_BITS_MAX_ENUM AllocatorCreateFlags = C.VMA_ALLOCATOR_CREATE_FLAG_BITS_MAX_ENUM
	// ALLOCATOR_CREATE_KHR_BIND_MEMORY2_BIT wraps VMA_ALLOCATOR_CREATE_KHR_BIND_MEMORY2_BIT.
	/*
	   Enables usage of VK_KHR_bind_memory2 extension.

	   The flag works only if VmaAllocatorCreateInfo::vulkanApiVersion `== VK_API_VERSION_1_0`.
	   When it is `VK_API_VERSION_1_1`, the flag is ignored because the extension has been promoted to Vulkan 1.1.

	   You may set this flag only if you found out that this device extension is supported,
	   you enabled it while creating Vulkan device passed as VmaAllocatorCreateInfo::device,
	   and you want it to be used internally by this library.

	   The extension provides functions `vkBindBufferMemory2KHR` and `vkBindImageMemory2KHR`,
	   which allow to pass a chain of `pNext` structures while binding.
	   This flag is required if you use `pNext` parameter in vmaBindBufferMemory2() or vmaBindImageMemory2().
	*/
	ALLOCATOR_CREATE_KHR_BIND_MEMORY2_BIT AllocatorCreateFlags = C.VMA_ALLOCATOR_CREATE_KHR_BIND_MEMORY2_BIT
	// ALLOCATOR_CREATE_KHR_DEDICATED_ALLOCATION_BIT wraps VMA_ALLOCATOR_CREATE_KHR_DEDICATED_ALLOCATION_BIT.
	/*
	   \brief Enables usage of VK_KHR_dedicated_allocation extension.

	   The flag works only if VmaAllocatorCreateInfo::vulkanApiVersion `== VK_API_VERSION_1_0`.
	   When it is `VK_API_VERSION_1_1`, the flag is ignored because the extension has been promoted to Vulkan 1.1.

	   Using this extension will automatically allocate dedicated blocks of memory for
	   some buffers and images instead of suballocating place for them out of bigger
	   memory blocks (as if you explicitly used #VMA_ALLOCATION_CREATE_DEDICATED_MEMORY_BIT
	   flag) when it is recommended by the driver. It may improve performance on some
	   GPUs.

	   You may set this flag only if you found out that following device extensions are
	   supported, you enabled them while creating Vulkan device passed as
	   VmaAllocatorCreateInfo::device, and you want them to be used internally by this
	   library:

	   - VK_KHR_get_memory_requirements2 (device extension)
	   - VK_KHR_dedicated_allocation (device extension)

	   When this flag is set, you can experience following warnings reported by Vulkan
	   validation layer. You can ignore them.

	   > vkBindBufferMemory(): Binding memory to buffer 0x2d but vkGetBufferMemoryRequirements() has not been called on that buffer.
	*/
	ALLOCATOR_CREATE_KHR_DEDICATED_ALLOCATION_BIT AllocatorCreateFlags = C.VMA_ALLOCATOR_CREATE_KHR_DEDICATED_ALLOCATION_BIT
	// ALLOCATOR_CREATE_KHR_EXTERNAL_MEMORY_WIN32_BIT wraps VMA_ALLOCATOR_CREATE_KHR_EXTERNAL_MEMORY_WIN32_BIT.
	/*
	   Enables usage of VK_KHR_external_memory_win32 extension in the library.

	   You should set this flag if you found available and enabled this device extension,
	   while creating Vulkan device passed as VmaAllocatorCreateInfo::device.
	   For more information, see \ref other_api_interop.
	*/
	ALLOCATOR_CREATE_KHR_EXTERNAL_MEMORY_WIN32_BIT AllocatorCreateFlags = C.VMA_ALLOCATOR_CREATE_KHR_EXTERNAL_MEMORY_WIN32_BIT
	// ALLOCATOR_CREATE_KHR_MAINTENANCE4_BIT wraps VMA_ALLOCATOR_CREATE_KHR_MAINTENANCE4_BIT.
	/*
	   Enables usage of VK_KHR_maintenance4 extension in the library.

	   You may set this flag only if you found available and enabled this device extension,
	   while creating Vulkan device passed as VmaAllocatorCreateInfo::device.
	*/
	ALLOCATOR_CREATE_KHR_MAINTENANCE4_BIT AllocatorCreateFlags = C.VMA_ALLOCATOR_CREATE_KHR_MAINTENANCE4_BIT
	// ALLOCATOR_CREATE_KHR_MAINTENANCE5_BIT wraps VMA_ALLOCATOR_CREATE_KHR_MAINTENANCE5_BIT.
	/*
	   Enables usage of VK_KHR_maintenance5 extension in the library.

	   You should set this flag if you found available and enabled this device extension,
	   while creating Vulkan device passed as VmaAllocatorCreateInfo::device.
	*/
	ALLOCATOR_CREATE_KHR_MAINTENANCE5_BIT AllocatorCreateFlags = C.VMA_ALLOCATOR_CREATE_KHR_MAINTENANCE5_BIT
)

func (e AllocatorCreateFlags) String() string {
	return fmt.Sprintf("VmaAllocatorCreateFlags(%b)", e)
}

// DefragmentationFlags wraps the bitmask VmaDefragmentationFlags.
//
//	Flags to be passed as VmaDefragmentationInfo::flags.
type DefragmentationFlags int32

const (
	// DEFRAGMENTATION_FLAG_ALGORITHM_BALANCED_BIT wraps VMA_DEFRAGMENTATION_FLAG_ALGORITHM_BALANCED_BIT.
	/*
	   \brief Default defragmentation algorithm, applied also when no `ALGORITHM` flag is specified.
	   Offers a balance between defragmentation quality and the amount of allocations and bytes that need to be moved.
	*/
	DEFRAGMENTATION_FLAG_ALGORITHM_BALANCED_BIT DefragmentationFlags = C.VMA_DEFRAGMENTATION_FLAG_ALGORITHM_BALANCED_BIT
	// DEFRAGMENTATION_FLAG_ALGORITHM_EXTENSIVE_BIT wraps VMA_DEFRAGMENTATION_FLAG_ALGORITHM_EXTENSIVE_BIT.
	/*
	   \brief Use the most roboust algorithm at the cost of time to compute and number of copies to make.
	   Only available when bufferImageGranularity is greater than 1, since it aims to reduce
	   alignment issues between different types of resources.
	   Otherwise falls back to same behavior as #VMA_DEFRAGMENTATION_FLAG_ALGORITHM_FULL_BIT.
	*/
	DEFRAGMENTATION_FLAG_ALGORITHM_EXTENSIVE_BIT DefragmentationFlags = C.VMA_DEFRAGMENTATION_FLAG_ALGORITHM_EXTENSIVE_BIT
	// DEFRAGMENTATION_FLAG_ALGORITHM_FAST_BIT wraps VMA_DEFRAGMENTATION_FLAG_ALGORITHM_FAST_BIT.
	/*
	   \brief Use simple but fast algorithm for defragmentation.
	   May not achieve best results but will require least time to compute and least allocations to copy.
	*/
	DEFRAGMENTATION_FLAG_ALGORITHM_FAST_BIT DefragmentationFlags = C.VMA_DEFRAGMENTATION_FLAG_ALGORITHM_FAST_BIT
	// DEFRAGMENTATION_FLAG_ALGORITHM_FULL_BIT wraps VMA_DEFRAGMENTATION_FLAG_ALGORITHM_FULL_BIT.
	/*
	   \brief Perform full defragmentation of memory.
	   Can result in notably more time to compute and allocations to copy, but will achieve best memory packing.
	*/
	DEFRAGMENTATION_FLAG_ALGORITHM_FULL_BIT DefragmentationFlags = C.VMA_DEFRAGMENTATION_FLAG_ALGORITHM_FULL_BIT
	// DEFRAGMENTATION_FLAG_ALGORITHM_MASK wraps VMA_DEFRAGMENTATION_FLAG_ALGORITHM_MASK.
	//
	//	A bit mask to extract only `ALGORITHM` bits from entire set of flags.
	DEFRAGMENTATION_FLAG_ALGORITHM_MASK DefragmentationFlags = C.VMA_DEFRAGMENTATION_FLAG_ALGORITHM_MASK
	// DEFRAGMENTATION_FLAG_BITS_MAX_ENUM wraps VMA_DEFRAGMENTATION_FLAG_BITS_MAX_ENUM.
	//
	//	A bit mask to extract only `ALGORITHM` bits from entire set of flags.
	DEFRAGMENTATION_FLAG_BITS_MAX_ENUM DefragmentationFlags = C.VMA_DEFRAGMENTATION_FLAG_BITS_MAX_ENUM
)

func (e DefragmentationFlags) String() string {
	return fmt.Sprintf("VmaDefragmentationFlags(%b)", e)
}

// PoolCreateFlags wraps the bitmask VmaPoolCreateFlags.
//
//	Flags to be passed as VmaPoolCreateInfo::flags.
type PoolCreateFlags int32

const (
	// POOL_CREATE_ALGORITHM_MASK wraps VMA_POOL_CREATE_ALGORITHM_MASK.
	//
	//	Bit mask to extract only `ALGORITHM` bits from entire set of flags.
	POOL_CREATE_ALGORITHM_MASK PoolCreateFlags = C.VMA_POOL_CREATE_ALGORITHM_MASK
	// POOL_CREATE_FLAG_BITS_MAX_ENUM wraps VMA_POOL_CREATE_FLAG_BITS_MAX_ENUM.
	//
	//	Bit mask to extract only `ALGORITHM` bits from entire set of flags.
	POOL_CREATE_FLAG_BITS_MAX_ENUM PoolCreateFlags = C.VMA_POOL_CREATE_FLAG_BITS_MAX_ENUM
	// POOL_CREATE_IGNORE_BUFFER_IMAGE_GRANULARITY_BIT wraps VMA_POOL_CREATE_IGNORE_BUFFER_IMAGE_GRANULARITY_BIT.
	/*
	   \brief Use this flag if you always allocate only buffers and linear images or only optimal images out of this pool and so Buffer-Image Granularity can be ignored.

	   This is an optional optimization flag.

	   If you always allocate using vmaCreateBuffer(), vmaCreateImage(),
	   vmaAllocateMemoryForBuffer(), then you don't need to use it because allocator
	   knows exact type of your allocations so it can handle Buffer-Image Granularity
	   in the optimal way.

	   If you also allocate using vmaAllocateMemoryForImage() or vmaAllocateMemory(),
	   exact type of such allocations is not known, so allocator must be conservative
	   in handling Buffer-Image Granularity, which can lead to suboptimal allocation
	   (wasted memory). In that case, if you can make sure you always allocate only
	   buffers and linear images or only optimal images out of this pool, use this flag
	   to make allocator disregard Buffer-Image Granularity and so make allocations
	   faster and more optimal.
	*/
	POOL_CREATE_IGNORE_BUFFER_IMAGE_GRANULARITY_BIT PoolCreateFlags = C.VMA_POOL_CREATE_IGNORE_BUFFER_IMAGE_GRANULARITY_BIT
	// POOL_CREATE_LINEAR_ALGORITHM_BIT wraps VMA_POOL_CREATE_LINEAR_ALGORITHM_BIT.
	/*
	   \brief Enables alternative, linear allocation algorithm in this pool.

	   Specify this flag to enable linear allocation algorithm, which always creates
	   new allocations after last one and doesn't reuse space from allocations freed in
	   between. It trades memory consumption for simplified algorithm and data
	   structure, which has better performance and uses less memory for metadata.

	   By using this flag, you can achieve behavior of free-at-once, stack,
	   ring buffer, and double stack.
	   For details, see documentation chapter \ref linear_algorithm.
	*/
	POOL_CREATE_LINEAR_ALGORITHM_BIT PoolCreateFlags = C.VMA_POOL_CREATE_LINEAR_ALGORITHM_BIT
)

func (e PoolCreateFlags) String() string {
	return fmt.Sprintf("VmaPoolCreateFlags(%b)", e)
}

// VirtualAllocationCreateFlags wraps the bitmask VmaVirtualAllocationCreateFlags.
//
//	Flags to be passed as VmaVirtualAllocationCreateInfo::flags.
type VirtualAllocationCreateFlags int32

const (
	// VIRTUAL_ALLOCATION_CREATE_FLAG_BITS_MAX_ENUM wraps VMA_VIRTUAL_ALLOCATION_CREATE_FLAG_BITS_MAX_ENUM.
	/*
	   \brief A bit mask to extract only `STRATEGY` bits from entire set of flags.

	   These strategy flags are binary compatible with equivalent flags in #VmaAllocationCreateFlagBits.
	*/
	VIRTUAL_ALLOCATION_CREATE_FLAG_BITS_MAX_ENUM VirtualAllocationCreateFlags = C.VMA_VIRTUAL_ALLOCATION_CREATE_FLAG_BITS_MAX_ENUM
	// VIRTUAL_ALLOCATION_CREATE_STRATEGY_MASK wraps VMA_VIRTUAL_ALLOCATION_CREATE_STRATEGY_MASK.
	/*
	   \brief A bit mask to extract only `STRATEGY` bits from entire set of flags.

	   These strategy flags are binary compatible with equivalent flags in #VmaAllocationCreateFlagBits.
	*/
	VIRTUAL_ALLOCATION_CREATE_STRATEGY_MASK VirtualAllocationCreateFlags = C.VMA_VIRTUAL_ALLOCATION_CREATE_STRATEGY_MASK
	// VIRTUAL_ALLOCATION_CREATE_STRATEGY_MIN_MEMORY_BIT wraps VMA_VIRTUAL_ALLOCATION_CREATE_STRATEGY_MIN_MEMORY_BIT.
	//
	//	\brief Allocation strategy that tries to minimize memory usage.
	VIRTUAL_ALLOCATION_CREATE_STRATEGY_MIN_MEMORY_BIT VirtualAllocationCreateFlags = C.VMA_VIRTUAL_ALLOCATION_CREATE_STRATEGY_MIN_MEMORY_BIT
	// VIRTUAL_ALLOCATION_CREATE_STRATEGY_MIN_OFFSET_BIT wraps VMA_VIRTUAL_ALLOCATION_CREATE_STRATEGY_MIN_OFFSET_BIT.
	/*
	   Allocation strategy that chooses always the lowest offset in available space.
	   This is not the most efficient strategy but achieves highly packed data.
	*/
	VIRTUAL_ALLOCATION_CREATE_STRATEGY_MIN_OFFSET_BIT VirtualAllocationCreateFlags = C.VMA_VIRTUAL_ALLOCATION_CREATE_STRATEGY_MIN_OFFSET_BIT
	// VIRTUAL_ALLOCATION_CREATE_STRATEGY_MIN_TIME_BIT wraps VMA_VIRTUAL_ALLOCATION_CREATE_STRATEGY_MIN_TIME_BIT.
	//
	//	\brief Allocation strategy that tries to minimize allocation time.
	VIRTUAL_ALLOCATION_CREATE_STRATEGY_MIN_TIME_BIT VirtualAllocationCreateFlags = C.VMA_VIRTUAL_ALLOCATION_CREATE_STRATEGY_MIN_TIME_BIT
	// VIRTUAL_ALLOCATION_CREATE_UPPER_ADDRESS_BIT wraps VMA_VIRTUAL_ALLOCATION_CREATE_UPPER_ADDRESS_BIT.
	/*
	   \brief Allocation will be created from upper stack in a double stack pool.

	   This flag is only allowed for virtual blocks created with #VMA_VIRTUAL_BLOCK_CREATE_LINEAR_ALGORITHM_BIT flag.
	*/
	VIRTUAL_ALLOCATION_CREATE_UPPER_ADDRESS_BIT VirtualAllocationCreateFlags = C.VMA_VIRTUAL_ALLOCATION_CREATE_UPPER_ADDRESS_BIT
)

func (e VirtualAllocationCreateFlags) String() string {
	return fmt.Sprintf("VmaVirtualAllocationCreateFlags(%b)", e)
}

// VirtualBlockCreateFlags wraps the bitmask VmaVirtualBlockCreateFlags.
//
//	Flags to be passed as VmaVirtualBlockCreateInfo::flags.
type VirtualBlockCreateFlags int32

const (
	// VIRTUAL_BLOCK_CREATE_ALGORITHM_MASK wraps VMA_VIRTUAL_BLOCK_CREATE_ALGORITHM_MASK.
	//
	//	\brief Bit mask to extract only `ALGORITHM` bits from entire set of flags.
	VIRTUAL_BLOCK_CREATE_ALGORITHM_MASK VirtualBlockCreateFlags = C.VMA_VIRTUAL_BLOCK_CREATE_ALGORITHM_MASK
	// VIRTUAL_BLOCK_CREATE_FLAG_BITS_MAX_ENUM wraps VMA_VIRTUAL_BLOCK_CREATE_FLAG_BITS_MAX_ENUM.
	//
	//	\brief Bit mask to extract only `ALGORITHM` bits from entire set of flags.
	VIRTUAL_BLOCK_CREATE_FLAG_BITS_MAX_ENUM VirtualBlockCreateFlags = C.VMA_VIRTUAL_BLOCK_CREATE_FLAG_BITS_MAX_ENUM
	// VIRTUAL_BLOCK_CREATE_LINEAR_ALGORITHM_BIT wraps VMA_VIRTUAL_BLOCK_CREATE_LINEAR_ALGORITHM_BIT.
	/*
	   \brief Enables alternative, linear allocation algorithm in this virtual block.

	   Specify this flag to enable linear allocation algorithm, which always creates
	   new allocations after last one and doesn't reuse space from allocations freed in
	   between. It trades memory consumption for simplified algorithm and data
	   structure, which has better performance and uses less memory for metadata.

	   By using this flag, you can achieve behavior of free-at-once, stack,
	   ring buffer, and double stack.
	   For details, see documentation chapter \ref linear_algorithm.
	*/
	VIRTUAL_BLOCK_CREATE_LINEAR_ALGORITHM_BIT VirtualBlockCreateFlags = C.VMA_VIRTUAL_BLOCK_CREATE_LINEAR_ALGORITHM_BIT
)

func (e VirtualBlockCreateFlags) String() string {
	return fmt.Sprintf("VmaVirtualBlockCreateFlags(%b)", e)
}
