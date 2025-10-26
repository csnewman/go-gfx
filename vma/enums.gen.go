package vma

import "fmt"

// #include "vma.h"
import "C"

// DefragmentationMoveOperation wraps the bitmask VmaDefragmentationMoveOperation.
//
//	Operation performed on single defragmentation move. See structure #VmaDefragmentationMove.
type DefragmentationMoveOperation int32

const (
	// DEFRAGMENTATION_MOVE_OPERATION_COPY wraps VMA_DEFRAGMENTATION_MOVE_OPERATION_COPY.
	//
	//	Buffer/image has been recreated at `dstTmpAllocation`, data has been copied, old buffer/image has been destroyed. `srcAllocation` should be changed to point to the new place. This is the default value set by vmaBeginDefragmentationPass().
	DEFRAGMENTATION_MOVE_OPERATION_COPY DefragmentationMoveOperation = C.VMA_DEFRAGMENTATION_MOVE_OPERATION_COPY
	// DEFRAGMENTATION_MOVE_OPERATION_IGNORE wraps VMA_DEFRAGMENTATION_MOVE_OPERATION_IGNORE.
	//
	//	Set this value if you cannot move the allocation. New place reserved at `dstTmpAllocation` will be freed. `srcAllocation` will remain unchanged.
	DEFRAGMENTATION_MOVE_OPERATION_IGNORE DefragmentationMoveOperation = C.VMA_DEFRAGMENTATION_MOVE_OPERATION_IGNORE
	// DEFRAGMENTATION_MOVE_OPERATION_DESTROY wraps VMA_DEFRAGMENTATION_MOVE_OPERATION_DESTROY.
	//
	//	Set this value if you decide to abandon the allocation and you destroyed the buffer/image. New place reserved at `dstTmpAllocation` will be freed, along with `srcAllocation`, which will be destroyed.
	DEFRAGMENTATION_MOVE_OPERATION_DESTROY DefragmentationMoveOperation = C.VMA_DEFRAGMENTATION_MOVE_OPERATION_DESTROY
)

func (e DefragmentationMoveOperation) String() string {
	return fmt.Sprintf("VmaDefragmentationMoveOperation(%b)", e)
}

// MemoryUsage wraps the bitmask VmaMemoryUsage.
//
//	\brief Intended usage of the allocated memory.
type MemoryUsage int32

const (
	// MEMORY_USAGE_UNKNOWN wraps VMA_MEMORY_USAGE_UNKNOWN.
	/*
	   No intended memory usage specified.
	   Use other members of VmaAllocationCreateInfo to specify your requirements.
	*/
	MEMORY_USAGE_UNKNOWN MemoryUsage = C.VMA_MEMORY_USAGE_UNKNOWN
	// MEMORY_USAGE_GPU_ONLY wraps VMA_MEMORY_USAGE_GPU_ONLY.
	/*
	   \deprecated Obsolete, preserved for backward compatibility.
	   Prefers `VK_MEMORY_PROPERTY_DEVICE_LOCAL_BIT`.
	*/
	MEMORY_USAGE_GPU_ONLY MemoryUsage = C.VMA_MEMORY_USAGE_GPU_ONLY
	// MEMORY_USAGE_CPU_ONLY wraps VMA_MEMORY_USAGE_CPU_ONLY.
	/*
	   \deprecated Obsolete, preserved for backward compatibility.
	   Guarantees `VK_MEMORY_PROPERTY_HOST_VISIBLE_BIT` and `VK_MEMORY_PROPERTY_HOST_COHERENT_BIT`.
	*/
	MEMORY_USAGE_CPU_ONLY MemoryUsage = C.VMA_MEMORY_USAGE_CPU_ONLY
	// MEMORY_USAGE_CPU_TO_GPU wraps VMA_MEMORY_USAGE_CPU_TO_GPU.
	/*
	   \deprecated Obsolete, preserved for backward compatibility.
	   Guarantees `VK_MEMORY_PROPERTY_HOST_VISIBLE_BIT`, prefers `VK_MEMORY_PROPERTY_DEVICE_LOCAL_BIT`.
	*/
	MEMORY_USAGE_CPU_TO_GPU MemoryUsage = C.VMA_MEMORY_USAGE_CPU_TO_GPU
	// MEMORY_USAGE_GPU_TO_CPU wraps VMA_MEMORY_USAGE_GPU_TO_CPU.
	/*
	   \deprecated Obsolete, preserved for backward compatibility.
	   Guarantees `VK_MEMORY_PROPERTY_HOST_VISIBLE_BIT`, prefers `VK_MEMORY_PROPERTY_HOST_CACHED_BIT`.
	*/
	MEMORY_USAGE_GPU_TO_CPU MemoryUsage = C.VMA_MEMORY_USAGE_GPU_TO_CPU
	// MEMORY_USAGE_CPU_COPY wraps VMA_MEMORY_USAGE_CPU_COPY.
	/*
	   \deprecated Obsolete, preserved for backward compatibility.
	   Prefers not `VK_MEMORY_PROPERTY_DEVICE_LOCAL_BIT`.
	*/
	MEMORY_USAGE_CPU_COPY MemoryUsage = C.VMA_MEMORY_USAGE_CPU_COPY
	// MEMORY_USAGE_GPU_LAZILY_ALLOCATED wraps VMA_MEMORY_USAGE_GPU_LAZILY_ALLOCATED.
	/*
	   Lazily allocated GPU memory having `VK_MEMORY_PROPERTY_LAZILY_ALLOCATED_BIT`.
	   Exists mostly on mobile platforms. Using it on desktop PC or other GPUs with no such memory type present will fail the allocation.

	   Usage: Memory for transient attachment images (color attachments, depth attachments etc.), created with `VK_IMAGE_USAGE_TRANSIENT_ATTACHMENT_BIT`.

	   Allocations with this usage are always created as dedicated - it implies #VMA_ALLOCATION_CREATE_DEDICATED_MEMORY_BIT.
	*/
	MEMORY_USAGE_GPU_LAZILY_ALLOCATED MemoryUsage = C.VMA_MEMORY_USAGE_GPU_LAZILY_ALLOCATED
	// MEMORY_USAGE_AUTO wraps VMA_MEMORY_USAGE_AUTO.
	/*
	   Selects best memory type automatically.
	   This flag is recommended for most common use cases.

	   When using this flag, if you want to map the allocation (using vmaMapMemory() or #VMA_ALLOCATION_CREATE_MAPPED_BIT),
	   you must pass one of the flags: #VMA_ALLOCATION_CREATE_HOST_ACCESS_SEQUENTIAL_WRITE_BIT or #VMA_ALLOCATION_CREATE_HOST_ACCESS_RANDOM_BIT
	   in VmaAllocationCreateInfo::flags.

	   It can be used only with functions that let the library know `VkBufferCreateInfo` or `VkImageCreateInfo`, e.g.
	   vmaCreateBuffer(), vmaCreateImage(), vmaFindMemoryTypeIndexForBufferInfo(), vmaFindMemoryTypeIndexForImageInfo()
	   and not with generic memory allocation functions.
	*/
	MEMORY_USAGE_AUTO MemoryUsage = C.VMA_MEMORY_USAGE_AUTO
	// MEMORY_USAGE_AUTO_PREFER_DEVICE wraps VMA_MEMORY_USAGE_AUTO_PREFER_DEVICE.
	/*
	   Selects best memory type automatically with preference for GPU (device) memory.

	   When using this flag, if you want to map the allocation (using vmaMapMemory() or #VMA_ALLOCATION_CREATE_MAPPED_BIT),
	   you must pass one of the flags: #VMA_ALLOCATION_CREATE_HOST_ACCESS_SEQUENTIAL_WRITE_BIT or #VMA_ALLOCATION_CREATE_HOST_ACCESS_RANDOM_BIT
	   in VmaAllocationCreateInfo::flags.

	   It can be used only with functions that let the library know `VkBufferCreateInfo` or `VkImageCreateInfo`, e.g.
	   vmaCreateBuffer(), vmaCreateImage(), vmaFindMemoryTypeIndexForBufferInfo(), vmaFindMemoryTypeIndexForImageInfo()
	   and not with generic memory allocation functions.
	*/
	MEMORY_USAGE_AUTO_PREFER_DEVICE MemoryUsage = C.VMA_MEMORY_USAGE_AUTO_PREFER_DEVICE
	// MEMORY_USAGE_AUTO_PREFER_HOST wraps VMA_MEMORY_USAGE_AUTO_PREFER_HOST.
	/*
	   Selects best memory type automatically with preference for CPU (host) memory.

	   When using this flag, if you want to map the allocation (using vmaMapMemory() or #VMA_ALLOCATION_CREATE_MAPPED_BIT),
	   you must pass one of the flags: #VMA_ALLOCATION_CREATE_HOST_ACCESS_SEQUENTIAL_WRITE_BIT or #VMA_ALLOCATION_CREATE_HOST_ACCESS_RANDOM_BIT
	   in VmaAllocationCreateInfo::flags.

	   It can be used only with functions that let the library know `VkBufferCreateInfo` or `VkImageCreateInfo`, e.g.
	   vmaCreateBuffer(), vmaCreateImage(), vmaFindMemoryTypeIndexForBufferInfo(), vmaFindMemoryTypeIndexForImageInfo()
	   and not with generic memory allocation functions.
	*/
	MEMORY_USAGE_AUTO_PREFER_HOST MemoryUsage = C.VMA_MEMORY_USAGE_AUTO_PREFER_HOST
	// MEMORY_USAGE_MAX_ENUM wraps VMA_MEMORY_USAGE_MAX_ENUM.
	/*
	   Selects best memory type automatically with preference for CPU (host) memory.

	   When using this flag, if you want to map the allocation (using vmaMapMemory() or #VMA_ALLOCATION_CREATE_MAPPED_BIT),
	   you must pass one of the flags: #VMA_ALLOCATION_CREATE_HOST_ACCESS_SEQUENTIAL_WRITE_BIT or #VMA_ALLOCATION_CREATE_HOST_ACCESS_RANDOM_BIT
	   in VmaAllocationCreateInfo::flags.

	   It can be used only with functions that let the library know `VkBufferCreateInfo` or `VkImageCreateInfo`, e.g.
	   vmaCreateBuffer(), vmaCreateImage(), vmaFindMemoryTypeIndexForBufferInfo(), vmaFindMemoryTypeIndexForImageInfo()
	   and not with generic memory allocation functions.
	*/
	MEMORY_USAGE_MAX_ENUM MemoryUsage = C.VMA_MEMORY_USAGE_MAX_ENUM
)

func (e MemoryUsage) String() string {
	return fmt.Sprintf("VmaMemoryUsage(%b)", e)
}
