package vma

// #include "vma.h"
import "C"

// Allocation wraps the handle VmaAllocation.
type Allocation uintptr

// AllocationNil is a null pointer.
var AllocationNil Allocation

// Allocator wraps the handle VmaAllocator.
type Allocator uintptr

// AllocatorNil is a null pointer.
var AllocatorNil Allocator

// DefragmentationContext wraps the handle VmaDefragmentationContext.
type DefragmentationContext uintptr

// DefragmentationContextNil is a null pointer.
var DefragmentationContextNil DefragmentationContext

// Pool wraps the handle VmaPool.
type Pool uintptr

// PoolNil is a null pointer.
var PoolNil Pool

// VirtualAllocation wraps the handle VmaVirtualAllocation.
type VirtualAllocation uint64

// VirtualAllocationNil is a null pointer.
var VirtualAllocationNil VirtualAllocation

// VirtualBlock wraps the handle VmaVirtualBlock.
type VirtualBlock uintptr

// VirtualBlockNil is a null pointer.
var VirtualBlockNil VirtualBlock
