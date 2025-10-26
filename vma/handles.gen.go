package vma

// #include "vma.h"
import "C"

// Allocation wraps VmaAllocation.
type Allocation uintptr

// AllocationNil is a null pointer.
var AllocationNil Allocation

// Allocator wraps VmaAllocator.
type Allocator uintptr

// AllocatorNil is a null pointer.
var AllocatorNil Allocator

// DefragmentationContext wraps VmaDefragmentationContext.
type DefragmentationContext uintptr

// DefragmentationContextNil is a null pointer.
var DefragmentationContextNil DefragmentationContext

// Pool wraps VmaPool.
type Pool uintptr

// PoolNil is a null pointer.
var PoolNil Pool

// VirtualAllocation wraps VmaVirtualAllocation.
type VirtualAllocation uint64

// VirtualAllocationNil is a null pointer.
var VirtualAllocationNil VirtualAllocation

// VirtualBlock wraps VmaVirtualBlock.
type VirtualBlock uintptr

// VirtualBlockNil is a null pointer.
var VirtualBlockNil VirtualBlock
