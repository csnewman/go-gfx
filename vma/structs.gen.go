package vma

import (
	ffi "github.com/csnewman/go-gfx/ffi"
	vk "github.com/csnewman/go-gfx/vk"
	"unsafe"
)

// #include "vma.h"
// static const int offsetof_VmaAllocationInfo2_allocationInfo = offsetof(VmaAllocationInfo2, allocationInfo);
// static const int offsetof_VmaBudget_statistics = offsetof(VmaBudget, statistics);
// static const int offsetof_VmaDetailedStatistics_statistics = offsetof(VmaDetailedStatistics, statistics);
// static const int offsetof_VmaTotalStatistics_total = offsetof(VmaTotalStatistics, total);
import "C"

// AllocationCreateInfo wraps VmaAllocationCreateInfo.
type AllocationCreateInfo struct {
	ptr *C.VmaAllocationCreateInfo
}

// AllocationCreateInfoNil is a null pointer.
var AllocationCreateInfoNil AllocationCreateInfo

// AllocationCreateInfoSizeOf is the byte size of VmaAllocationCreateInfo.
const AllocationCreateInfoSizeOf = int(C.sizeof_VmaAllocationCreateInfo)

// AllocationCreateInfoFromPtr converts a raw pointer to a AllocationCreateInfo.
func AllocationCreateInfoFromPtr(ptr unsafe.Pointer) AllocationCreateInfo {
	return AllocationCreateInfo{ptr: (*C.VmaAllocationCreateInfo)(ptr)}
}

// AllocationCreateInfoAlloc allocates a continuous block of VmaAllocationCreateInfo.
func AllocationCreateInfoAlloc(alloc ffi.Allocator, count int) AllocationCreateInfo {
	ptr := alloc.Allocate(AllocationCreateInfoSizeOf * count)
	return AllocationCreateInfo{ptr: (*C.VmaAllocationCreateInfo)(ptr)}
}

// Raw returns a raw pointer to the struct.
func (p AllocationCreateInfo) Raw() unsafe.Pointer {
	return unsafe.Pointer(p.ptr)
}

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p AllocationCreateInfo) Offset(offset int) AllocationCreateInfo {
	ptr := unsafe.Add(unsafe.Pointer(p.ptr), offset*AllocationCreateInfoSizeOf)
	return AllocationCreateInfo{ptr: (*C.VmaAllocationCreateInfo)(ptr)}
}

// GetFlags returns the value in flags.
func (p AllocationCreateInfo) GetFlags() AllocationCreateFlags {
	return AllocationCreateFlags(p.ptr.flags)
}

// SetFlags sets the value in flags.
func (p AllocationCreateInfo) SetFlags(value AllocationCreateFlags) {
	p.ptr.flags = (C.VmaAllocationCreateFlags)(value)
}

// GetUsage returns the value in usage.
func (p AllocationCreateInfo) GetUsage() MemoryUsage {
	return MemoryUsage(p.ptr.usage)
}

// SetUsage sets the value in usage.
func (p AllocationCreateInfo) SetUsage(value MemoryUsage) {
	p.ptr.usage = (C.VmaMemoryUsage)(value)
}

// GetRequiredFlags returns the value in requiredFlags.
func (p AllocationCreateInfo) GetRequiredFlags() vk.MemoryPropertyFlags {
	return vk.MemoryPropertyFlags(p.ptr.requiredFlags)
}

// SetRequiredFlags sets the value in requiredFlags.
func (p AllocationCreateInfo) SetRequiredFlags(value vk.MemoryPropertyFlags) {
	p.ptr.requiredFlags = (C.VkMemoryPropertyFlags)(value)
}

// GetPreferredFlags returns the value in preferredFlags.
func (p AllocationCreateInfo) GetPreferredFlags() vk.MemoryPropertyFlags {
	return vk.MemoryPropertyFlags(p.ptr.preferredFlags)
}

// SetPreferredFlags sets the value in preferredFlags.
func (p AllocationCreateInfo) SetPreferredFlags(value vk.MemoryPropertyFlags) {
	p.ptr.preferredFlags = (C.VkMemoryPropertyFlags)(value)
}

// GetMemoryTypeBits returns the value in memoryTypeBits.
func (p AllocationCreateInfo) GetMemoryTypeBits() uint32 {
	return uint32(p.ptr.memoryTypeBits)
}

// SetMemoryTypeBits sets the value in memoryTypeBits.
func (p AllocationCreateInfo) SetMemoryTypeBits(value uint32) {
	p.ptr.memoryTypeBits = (C.uint32_t)(value)
}

// GetPool returns the value in pool.
func (p AllocationCreateInfo) GetPool() Pool {
	return Pool(unsafe.Pointer(p.ptr.pool))
}

// SetPool sets the value in pool.
func (p AllocationCreateInfo) SetPool(value Pool) {
	p.ptr.pool = (C.VmaPool)(unsafe.Pointer(value))
}

// GetPUserData returns the value in pUserData.
func (p AllocationCreateInfo) GetPUserData() unsafe.Pointer {
	return unsafe.Pointer(p.ptr.pUserData)
}

// SetPUserData sets the value in pUserData.
func (p AllocationCreateInfo) SetPUserData(value unsafe.Pointer) {
	p.ptr.pUserData = value
}

// GetPriority returns the value in priority.
func (p AllocationCreateInfo) GetPriority() float32 {
	return float32(p.ptr.priority)
}

// SetPriority sets the value in priority.
func (p AllocationCreateInfo) SetPriority(value float32) {
	p.ptr.priority = (C.float)(value)
}

// AllocationInfo wraps VmaAllocationInfo.
type AllocationInfo struct {
	ptr *C.VmaAllocationInfo
}

// AllocationInfoNil is a null pointer.
var AllocationInfoNil AllocationInfo

// AllocationInfoSizeOf is the byte size of VmaAllocationInfo.
const AllocationInfoSizeOf = int(C.sizeof_VmaAllocationInfo)

// AllocationInfoFromPtr converts a raw pointer to a AllocationInfo.
func AllocationInfoFromPtr(ptr unsafe.Pointer) AllocationInfo {
	return AllocationInfo{ptr: (*C.VmaAllocationInfo)(ptr)}
}

// AllocationInfoAlloc allocates a continuous block of VmaAllocationInfo.
func AllocationInfoAlloc(alloc ffi.Allocator, count int) AllocationInfo {
	ptr := alloc.Allocate(AllocationInfoSizeOf * count)
	return AllocationInfo{ptr: (*C.VmaAllocationInfo)(ptr)}
}

// Raw returns a raw pointer to the struct.
func (p AllocationInfo) Raw() unsafe.Pointer {
	return unsafe.Pointer(p.ptr)
}

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p AllocationInfo) Offset(offset int) AllocationInfo {
	ptr := unsafe.Add(unsafe.Pointer(p.ptr), offset*AllocationInfoSizeOf)
	return AllocationInfo{ptr: (*C.VmaAllocationInfo)(ptr)}
}

// GetMemoryType returns the value in memoryType.
func (p AllocationInfo) GetMemoryType() uint32 {
	return uint32(p.ptr.memoryType)
}

// SetMemoryType sets the value in memoryType.
func (p AllocationInfo) SetMemoryType(value uint32) {
	p.ptr.memoryType = (C.uint32_t)(value)
}

// GetDeviceMemory returns the value in deviceMemory.
func (p AllocationInfo) GetDeviceMemory() vk.DeviceMemory {
	return vk.DeviceMemory(p.ptr.deviceMemory)
}

// SetDeviceMemory sets the value in deviceMemory.
func (p AllocationInfo) SetDeviceMemory(value vk.DeviceMemory) {
	p.ptr.deviceMemory = (C.VkDeviceMemory)(value)
}

// GetOffset returns the value in offset.
func (p AllocationInfo) GetOffset() vk.DeviceSize {
	return vk.DeviceSize(p.ptr.offset)
}

// SetOffset sets the value in offset.
func (p AllocationInfo) SetOffset(value vk.DeviceSize) {
	p.ptr.offset = (C.VkDeviceSize)(value)
}

// GetSize returns the value in size.
func (p AllocationInfo) GetSize() vk.DeviceSize {
	return vk.DeviceSize(p.ptr.size)
}

// SetSize sets the value in size.
func (p AllocationInfo) SetSize(value vk.DeviceSize) {
	p.ptr.size = (C.VkDeviceSize)(value)
}

// GetPMappedData returns the value in pMappedData.
func (p AllocationInfo) GetPMappedData() unsafe.Pointer {
	return unsafe.Pointer(p.ptr.pMappedData)
}

// SetPMappedData sets the value in pMappedData.
func (p AllocationInfo) SetPMappedData(value unsafe.Pointer) {
	p.ptr.pMappedData = value
}

// GetPUserData returns the value in pUserData.
func (p AllocationInfo) GetPUserData() unsafe.Pointer {
	return unsafe.Pointer(p.ptr.pUserData)
}

// SetPUserData sets the value in pUserData.
func (p AllocationInfo) SetPUserData(value unsafe.Pointer) {
	p.ptr.pUserData = value
}

// GetPName returns the value in pName.
func (p AllocationInfo) GetPName() ffi.CString {
	return ffi.CStringFromPtr((unsafe.Pointer)(p.ptr.pName))
}

// SetPName sets the value in pName.
func (p AllocationInfo) SetPName(value ffi.CString) {
	p.ptr.pName = (*C.char)(value.Raw())
}

// AllocationInfo2 wraps VmaAllocationInfo2.
type AllocationInfo2 struct {
	ptr *C.VmaAllocationInfo2
}

// AllocationInfo2Nil is a null pointer.
var AllocationInfo2Nil AllocationInfo2

// AllocationInfo2SizeOf is the byte size of VmaAllocationInfo2.
const AllocationInfo2SizeOf = int(C.sizeof_VmaAllocationInfo2)

// AllocationInfo2FromPtr converts a raw pointer to a AllocationInfo2.
func AllocationInfo2FromPtr(ptr unsafe.Pointer) AllocationInfo2 {
	return AllocationInfo2{ptr: (*C.VmaAllocationInfo2)(ptr)}
}

// AllocationInfo2Alloc allocates a continuous block of VmaAllocationInfo2.
func AllocationInfo2Alloc(alloc ffi.Allocator, count int) AllocationInfo2 {
	ptr := alloc.Allocate(AllocationInfo2SizeOf * count)
	return AllocationInfo2{ptr: (*C.VmaAllocationInfo2)(ptr)}
}

// Raw returns a raw pointer to the struct.
func (p AllocationInfo2) Raw() unsafe.Pointer {
	return unsafe.Pointer(p.ptr)
}

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p AllocationInfo2) Offset(offset int) AllocationInfo2 {
	ptr := unsafe.Add(unsafe.Pointer(p.ptr), offset*AllocationInfo2SizeOf)
	return AllocationInfo2{ptr: (*C.VmaAllocationInfo2)(ptr)}
}

// RefAllocationInfo returns pointer to the allocationInfo field.
func (p AllocationInfo2) RefAllocationInfo() AllocationInfo {
	return AllocationInfo{ptr: (*C.VmaAllocationInfo)(unsafe.Add(unsafe.Pointer(p.ptr), uintptr(C.offsetof_VmaAllocationInfo2_allocationInfo)))}
}

// GetBlockSize returns the value in blockSize.
func (p AllocationInfo2) GetBlockSize() vk.DeviceSize {
	return vk.DeviceSize(p.ptr.blockSize)
}

// SetBlockSize sets the value in blockSize.
func (p AllocationInfo2) SetBlockSize(value vk.DeviceSize) {
	p.ptr.blockSize = (C.VkDeviceSize)(value)
}

// GetDedicatedMemory returns the value in dedicatedMemory.
func (p AllocationInfo2) GetDedicatedMemory() bool {
	return p.ptr.dedicatedMemory != 0
}

// SetDedicatedMemory sets the value in dedicatedMemory.
func (p AllocationInfo2) SetDedicatedMemory(value bool) {
	if value {
		p.ptr.dedicatedMemory = C.VkBool32(1)
	} else {
		p.ptr.dedicatedMemory = C.VkBool32(0)
	}
}

// AllocatorCreateInfo wraps VmaAllocatorCreateInfo.
type AllocatorCreateInfo struct {
	ptr *C.VmaAllocatorCreateInfo
}

// AllocatorCreateInfoNil is a null pointer.
var AllocatorCreateInfoNil AllocatorCreateInfo

// AllocatorCreateInfoSizeOf is the byte size of VmaAllocatorCreateInfo.
const AllocatorCreateInfoSizeOf = int(C.sizeof_VmaAllocatorCreateInfo)

// AllocatorCreateInfoFromPtr converts a raw pointer to a AllocatorCreateInfo.
func AllocatorCreateInfoFromPtr(ptr unsafe.Pointer) AllocatorCreateInfo {
	return AllocatorCreateInfo{ptr: (*C.VmaAllocatorCreateInfo)(ptr)}
}

// AllocatorCreateInfoAlloc allocates a continuous block of VmaAllocatorCreateInfo.
func AllocatorCreateInfoAlloc(alloc ffi.Allocator, count int) AllocatorCreateInfo {
	ptr := alloc.Allocate(AllocatorCreateInfoSizeOf * count)
	return AllocatorCreateInfo{ptr: (*C.VmaAllocatorCreateInfo)(ptr)}
}

// Raw returns a raw pointer to the struct.
func (p AllocatorCreateInfo) Raw() unsafe.Pointer {
	return unsafe.Pointer(p.ptr)
}

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p AllocatorCreateInfo) Offset(offset int) AllocatorCreateInfo {
	ptr := unsafe.Add(unsafe.Pointer(p.ptr), offset*AllocatorCreateInfoSizeOf)
	return AllocatorCreateInfo{ptr: (*C.VmaAllocatorCreateInfo)(ptr)}
}

// GetFlags returns the value in flags.
func (p AllocatorCreateInfo) GetFlags() AllocatorCreateFlags {
	return AllocatorCreateFlags(p.ptr.flags)
}

// SetFlags sets the value in flags.
func (p AllocatorCreateInfo) SetFlags(value AllocatorCreateFlags) {
	p.ptr.flags = (C.VmaAllocatorCreateFlags)(value)
}

// GetPhysicalDevice returns the value in physicalDevice.
func (p AllocatorCreateInfo) GetPhysicalDevice() vk.PhysicalDevice {
	return vk.PhysicalDevice(unsafe.Pointer(p.ptr.physicalDevice))
}

// SetPhysicalDevice sets the value in physicalDevice.
func (p AllocatorCreateInfo) SetPhysicalDevice(value vk.PhysicalDevice) {
	p.ptr.physicalDevice = (C.VkPhysicalDevice)(unsafe.Pointer(value))
}

// GetDevice returns the value in device.
func (p AllocatorCreateInfo) GetDevice() vk.Device {
	return vk.Device(unsafe.Pointer(p.ptr.device))
}

// SetDevice sets the value in device.
func (p AllocatorCreateInfo) SetDevice(value vk.Device) {
	p.ptr.device = (C.VkDevice)(unsafe.Pointer(value))
}

// GetPreferredLargeHeapBlockSize returns the value in preferredLargeHeapBlockSize.
func (p AllocatorCreateInfo) GetPreferredLargeHeapBlockSize() vk.DeviceSize {
	return vk.DeviceSize(p.ptr.preferredLargeHeapBlockSize)
}

// SetPreferredLargeHeapBlockSize sets the value in preferredLargeHeapBlockSize.
func (p AllocatorCreateInfo) SetPreferredLargeHeapBlockSize(value vk.DeviceSize) {
	p.ptr.preferredLargeHeapBlockSize = (C.VkDeviceSize)(value)
}

// AllocatorCreateInfo.pAllocationCallbacks is unsupported: pointer category vk struct.

// AllocatorCreateInfo.pDeviceMemoryCallbacks is unsupported: pointer category unknown type VmaDeviceMemoryCallbacks.

// GetPHeapSizeLimit returns the value in pHeapSizeLimit.
func (p AllocatorCreateInfo) GetPHeapSizeLimit() ffi.Ref[vk.DeviceSize] {
	return ffi.RefFromPtr[vk.DeviceSize](unsafe.Pointer(p.ptr.pHeapSizeLimit))
}

// SetPHeapSizeLimit sets the value in pHeapSizeLimit.
func (p AllocatorCreateInfo) SetPHeapSizeLimit(value ffi.Ref[vk.DeviceSize]) {
	p.ptr.pHeapSizeLimit = (*C.VkDeviceSize)(value.Raw())
}

// AllocatorCreateInfo.pVulkanFunctions is unsupported: pointer category unknown type VmaVulkanFunctions.

// GetInstance returns the value in instance.
func (p AllocatorCreateInfo) GetInstance() vk.Instance {
	return vk.Instance(unsafe.Pointer(p.ptr.instance))
}

// SetInstance sets the value in instance.
func (p AllocatorCreateInfo) SetInstance(value vk.Instance) {
	p.ptr.instance = (C.VkInstance)(unsafe.Pointer(value))
}

// GetVulkanApiVersion returns the value in vulkanApiVersion.
func (p AllocatorCreateInfo) GetVulkanApiVersion() uint32 {
	return uint32(p.ptr.vulkanApiVersion)
}

// SetVulkanApiVersion sets the value in vulkanApiVersion.
func (p AllocatorCreateInfo) SetVulkanApiVersion(value uint32) {
	p.ptr.vulkanApiVersion = (C.uint32_t)(value)
}

// AllocatorCreateInfo.pTypeExternalMemoryHandleTypes is unsupported: pointer category unknown type VkExternalMemoryHandleTypeFlagsKHR.

// AllocatorInfo wraps VmaAllocatorInfo.
type AllocatorInfo struct {
	ptr *C.VmaAllocatorInfo
}

// AllocatorInfoNil is a null pointer.
var AllocatorInfoNil AllocatorInfo

// AllocatorInfoSizeOf is the byte size of VmaAllocatorInfo.
const AllocatorInfoSizeOf = int(C.sizeof_VmaAllocatorInfo)

// AllocatorInfoFromPtr converts a raw pointer to a AllocatorInfo.
func AllocatorInfoFromPtr(ptr unsafe.Pointer) AllocatorInfo {
	return AllocatorInfo{ptr: (*C.VmaAllocatorInfo)(ptr)}
}

// AllocatorInfoAlloc allocates a continuous block of VmaAllocatorInfo.
func AllocatorInfoAlloc(alloc ffi.Allocator, count int) AllocatorInfo {
	ptr := alloc.Allocate(AllocatorInfoSizeOf * count)
	return AllocatorInfo{ptr: (*C.VmaAllocatorInfo)(ptr)}
}

// Raw returns a raw pointer to the struct.
func (p AllocatorInfo) Raw() unsafe.Pointer {
	return unsafe.Pointer(p.ptr)
}

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p AllocatorInfo) Offset(offset int) AllocatorInfo {
	ptr := unsafe.Add(unsafe.Pointer(p.ptr), offset*AllocatorInfoSizeOf)
	return AllocatorInfo{ptr: (*C.VmaAllocatorInfo)(ptr)}
}

// GetInstance returns the value in instance.
func (p AllocatorInfo) GetInstance() vk.Instance {
	return vk.Instance(unsafe.Pointer(p.ptr.instance))
}

// SetInstance sets the value in instance.
func (p AllocatorInfo) SetInstance(value vk.Instance) {
	p.ptr.instance = (C.VkInstance)(unsafe.Pointer(value))
}

// GetPhysicalDevice returns the value in physicalDevice.
func (p AllocatorInfo) GetPhysicalDevice() vk.PhysicalDevice {
	return vk.PhysicalDevice(unsafe.Pointer(p.ptr.physicalDevice))
}

// SetPhysicalDevice sets the value in physicalDevice.
func (p AllocatorInfo) SetPhysicalDevice(value vk.PhysicalDevice) {
	p.ptr.physicalDevice = (C.VkPhysicalDevice)(unsafe.Pointer(value))
}

// GetDevice returns the value in device.
func (p AllocatorInfo) GetDevice() vk.Device {
	return vk.Device(unsafe.Pointer(p.ptr.device))
}

// SetDevice sets the value in device.
func (p AllocatorInfo) SetDevice(value vk.Device) {
	p.ptr.device = (C.VkDevice)(unsafe.Pointer(value))
}

// Budget wraps VmaBudget.
type Budget struct {
	ptr *C.VmaBudget
}

// BudgetNil is a null pointer.
var BudgetNil Budget

// BudgetSizeOf is the byte size of VmaBudget.
const BudgetSizeOf = int(C.sizeof_VmaBudget)

// BudgetFromPtr converts a raw pointer to a Budget.
func BudgetFromPtr(ptr unsafe.Pointer) Budget {
	return Budget{ptr: (*C.VmaBudget)(ptr)}
}

// BudgetAlloc allocates a continuous block of VmaBudget.
func BudgetAlloc(alloc ffi.Allocator, count int) Budget {
	ptr := alloc.Allocate(BudgetSizeOf * count)
	return Budget{ptr: (*C.VmaBudget)(ptr)}
}

// Raw returns a raw pointer to the struct.
func (p Budget) Raw() unsafe.Pointer {
	return unsafe.Pointer(p.ptr)
}

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p Budget) Offset(offset int) Budget {
	ptr := unsafe.Add(unsafe.Pointer(p.ptr), offset*BudgetSizeOf)
	return Budget{ptr: (*C.VmaBudget)(ptr)}
}

// RefStatistics returns pointer to the statistics field.
func (p Budget) RefStatistics() Statistics {
	return Statistics{ptr: (*C.VmaStatistics)(unsafe.Add(unsafe.Pointer(p.ptr), uintptr(C.offsetof_VmaBudget_statistics)))}
}

// GetUsage returns the value in usage.
func (p Budget) GetUsage() vk.DeviceSize {
	return vk.DeviceSize(p.ptr.usage)
}

// SetUsage sets the value in usage.
func (p Budget) SetUsage(value vk.DeviceSize) {
	p.ptr.usage = (C.VkDeviceSize)(value)
}

// GetBudget returns the value in budget.
func (p Budget) GetBudget() vk.DeviceSize {
	return vk.DeviceSize(p.ptr.budget)
}

// SetBudget sets the value in budget.
func (p Budget) SetBudget(value vk.DeviceSize) {
	p.ptr.budget = (C.VkDeviceSize)(value)
}

// DefragmentationInfo wraps VmaDefragmentationInfo.
type DefragmentationInfo struct {
	ptr *C.VmaDefragmentationInfo
}

// DefragmentationInfoNil is a null pointer.
var DefragmentationInfoNil DefragmentationInfo

// DefragmentationInfoSizeOf is the byte size of VmaDefragmentationInfo.
const DefragmentationInfoSizeOf = int(C.sizeof_VmaDefragmentationInfo)

// DefragmentationInfoFromPtr converts a raw pointer to a DefragmentationInfo.
func DefragmentationInfoFromPtr(ptr unsafe.Pointer) DefragmentationInfo {
	return DefragmentationInfo{ptr: (*C.VmaDefragmentationInfo)(ptr)}
}

// DefragmentationInfoAlloc allocates a continuous block of VmaDefragmentationInfo.
func DefragmentationInfoAlloc(alloc ffi.Allocator, count int) DefragmentationInfo {
	ptr := alloc.Allocate(DefragmentationInfoSizeOf * count)
	return DefragmentationInfo{ptr: (*C.VmaDefragmentationInfo)(ptr)}
}

// Raw returns a raw pointer to the struct.
func (p DefragmentationInfo) Raw() unsafe.Pointer {
	return unsafe.Pointer(p.ptr)
}

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p DefragmentationInfo) Offset(offset int) DefragmentationInfo {
	ptr := unsafe.Add(unsafe.Pointer(p.ptr), offset*DefragmentationInfoSizeOf)
	return DefragmentationInfo{ptr: (*C.VmaDefragmentationInfo)(ptr)}
}

// GetFlags returns the value in flags.
func (p DefragmentationInfo) GetFlags() DefragmentationFlags {
	return DefragmentationFlags(p.ptr.flags)
}

// SetFlags sets the value in flags.
func (p DefragmentationInfo) SetFlags(value DefragmentationFlags) {
	p.ptr.flags = (C.VmaDefragmentationFlags)(value)
}

// GetPool returns the value in pool.
func (p DefragmentationInfo) GetPool() Pool {
	return Pool(unsafe.Pointer(p.ptr.pool))
}

// SetPool sets the value in pool.
func (p DefragmentationInfo) SetPool(value Pool) {
	p.ptr.pool = (C.VmaPool)(unsafe.Pointer(value))
}

// GetMaxBytesPerPass returns the value in maxBytesPerPass.
func (p DefragmentationInfo) GetMaxBytesPerPass() vk.DeviceSize {
	return vk.DeviceSize(p.ptr.maxBytesPerPass)
}

// SetMaxBytesPerPass sets the value in maxBytesPerPass.
func (p DefragmentationInfo) SetMaxBytesPerPass(value vk.DeviceSize) {
	p.ptr.maxBytesPerPass = (C.VkDeviceSize)(value)
}

// GetMaxAllocationsPerPass returns the value in maxAllocationsPerPass.
func (p DefragmentationInfo) GetMaxAllocationsPerPass() uint32 {
	return uint32(p.ptr.maxAllocationsPerPass)
}

// SetMaxAllocationsPerPass sets the value in maxAllocationsPerPass.
func (p DefragmentationInfo) SetMaxAllocationsPerPass(value uint32) {
	p.ptr.maxAllocationsPerPass = (C.uint32_t)(value)
}

// DefragmentationInfo.pfnBreakCallback is unsupported: direct category unknown type PFN_vmaCheckDefragmentationBreakFunction.

// GetPBreakCallbackUserData returns the value in pBreakCallbackUserData.
func (p DefragmentationInfo) GetPBreakCallbackUserData() unsafe.Pointer {
	return unsafe.Pointer(p.ptr.pBreakCallbackUserData)
}

// SetPBreakCallbackUserData sets the value in pBreakCallbackUserData.
func (p DefragmentationInfo) SetPBreakCallbackUserData(value unsafe.Pointer) {
	p.ptr.pBreakCallbackUserData = value
}

// DefragmentationMove wraps VmaDefragmentationMove.
type DefragmentationMove struct {
	ptr *C.VmaDefragmentationMove
}

// DefragmentationMoveNil is a null pointer.
var DefragmentationMoveNil DefragmentationMove

// DefragmentationMoveSizeOf is the byte size of VmaDefragmentationMove.
const DefragmentationMoveSizeOf = int(C.sizeof_VmaDefragmentationMove)

// DefragmentationMoveFromPtr converts a raw pointer to a DefragmentationMove.
func DefragmentationMoveFromPtr(ptr unsafe.Pointer) DefragmentationMove {
	return DefragmentationMove{ptr: (*C.VmaDefragmentationMove)(ptr)}
}

// DefragmentationMoveAlloc allocates a continuous block of VmaDefragmentationMove.
func DefragmentationMoveAlloc(alloc ffi.Allocator, count int) DefragmentationMove {
	ptr := alloc.Allocate(DefragmentationMoveSizeOf * count)
	return DefragmentationMove{ptr: (*C.VmaDefragmentationMove)(ptr)}
}

// Raw returns a raw pointer to the struct.
func (p DefragmentationMove) Raw() unsafe.Pointer {
	return unsafe.Pointer(p.ptr)
}

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p DefragmentationMove) Offset(offset int) DefragmentationMove {
	ptr := unsafe.Add(unsafe.Pointer(p.ptr), offset*DefragmentationMoveSizeOf)
	return DefragmentationMove{ptr: (*C.VmaDefragmentationMove)(ptr)}
}

// GetOperation returns the value in operation.
func (p DefragmentationMove) GetOperation() DefragmentationMoveOperation {
	return DefragmentationMoveOperation(p.ptr.operation)
}

// SetOperation sets the value in operation.
func (p DefragmentationMove) SetOperation(value DefragmentationMoveOperation) {
	p.ptr.operation = (C.VmaDefragmentationMoveOperation)(value)
}

// GetSrcAllocation returns the value in srcAllocation.
func (p DefragmentationMove) GetSrcAllocation() Allocation {
	return Allocation(unsafe.Pointer(p.ptr.srcAllocation))
}

// SetSrcAllocation sets the value in srcAllocation.
func (p DefragmentationMove) SetSrcAllocation(value Allocation) {
	p.ptr.srcAllocation = (C.VmaAllocation)(unsafe.Pointer(value))
}

// GetDstTmpAllocation returns the value in dstTmpAllocation.
func (p DefragmentationMove) GetDstTmpAllocation() Allocation {
	return Allocation(unsafe.Pointer(p.ptr.dstTmpAllocation))
}

// SetDstTmpAllocation sets the value in dstTmpAllocation.
func (p DefragmentationMove) SetDstTmpAllocation(value Allocation) {
	p.ptr.dstTmpAllocation = (C.VmaAllocation)(unsafe.Pointer(value))
}

// DefragmentationPassMoveInfo wraps VmaDefragmentationPassMoveInfo.
type DefragmentationPassMoveInfo struct {
	ptr *C.VmaDefragmentationPassMoveInfo
}

// DefragmentationPassMoveInfoNil is a null pointer.
var DefragmentationPassMoveInfoNil DefragmentationPassMoveInfo

// DefragmentationPassMoveInfoSizeOf is the byte size of VmaDefragmentationPassMoveInfo.
const DefragmentationPassMoveInfoSizeOf = int(C.sizeof_VmaDefragmentationPassMoveInfo)

// DefragmentationPassMoveInfoFromPtr converts a raw pointer to a DefragmentationPassMoveInfo.
func DefragmentationPassMoveInfoFromPtr(ptr unsafe.Pointer) DefragmentationPassMoveInfo {
	return DefragmentationPassMoveInfo{ptr: (*C.VmaDefragmentationPassMoveInfo)(ptr)}
}

// DefragmentationPassMoveInfoAlloc allocates a continuous block of VmaDefragmentationPassMoveInfo.
func DefragmentationPassMoveInfoAlloc(alloc ffi.Allocator, count int) DefragmentationPassMoveInfo {
	ptr := alloc.Allocate(DefragmentationPassMoveInfoSizeOf * count)
	return DefragmentationPassMoveInfo{ptr: (*C.VmaDefragmentationPassMoveInfo)(ptr)}
}

// Raw returns a raw pointer to the struct.
func (p DefragmentationPassMoveInfo) Raw() unsafe.Pointer {
	return unsafe.Pointer(p.ptr)
}

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p DefragmentationPassMoveInfo) Offset(offset int) DefragmentationPassMoveInfo {
	ptr := unsafe.Add(unsafe.Pointer(p.ptr), offset*DefragmentationPassMoveInfoSizeOf)
	return DefragmentationPassMoveInfo{ptr: (*C.VmaDefragmentationPassMoveInfo)(ptr)}
}

// GetMoveCount returns the value in moveCount.
func (p DefragmentationPassMoveInfo) GetMoveCount() uint32 {
	return uint32(p.ptr.moveCount)
}

// SetMoveCount sets the value in moveCount.
func (p DefragmentationPassMoveInfo) SetMoveCount(value uint32) {
	p.ptr.moveCount = (C.uint32_t)(value)
}

// DefragmentationPassMoveInfo.pMoves is unsupported: pointer category unknown type VmaDefragmentationMove.

// DefragmentationStats wraps VmaDefragmentationStats.
type DefragmentationStats struct {
	ptr *C.VmaDefragmentationStats
}

// DefragmentationStatsNil is a null pointer.
var DefragmentationStatsNil DefragmentationStats

// DefragmentationStatsSizeOf is the byte size of VmaDefragmentationStats.
const DefragmentationStatsSizeOf = int(C.sizeof_VmaDefragmentationStats)

// DefragmentationStatsFromPtr converts a raw pointer to a DefragmentationStats.
func DefragmentationStatsFromPtr(ptr unsafe.Pointer) DefragmentationStats {
	return DefragmentationStats{ptr: (*C.VmaDefragmentationStats)(ptr)}
}

// DefragmentationStatsAlloc allocates a continuous block of VmaDefragmentationStats.
func DefragmentationStatsAlloc(alloc ffi.Allocator, count int) DefragmentationStats {
	ptr := alloc.Allocate(DefragmentationStatsSizeOf * count)
	return DefragmentationStats{ptr: (*C.VmaDefragmentationStats)(ptr)}
}

// Raw returns a raw pointer to the struct.
func (p DefragmentationStats) Raw() unsafe.Pointer {
	return unsafe.Pointer(p.ptr)
}

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p DefragmentationStats) Offset(offset int) DefragmentationStats {
	ptr := unsafe.Add(unsafe.Pointer(p.ptr), offset*DefragmentationStatsSizeOf)
	return DefragmentationStats{ptr: (*C.VmaDefragmentationStats)(ptr)}
}

// GetBytesMoved returns the value in bytesMoved.
func (p DefragmentationStats) GetBytesMoved() vk.DeviceSize {
	return vk.DeviceSize(p.ptr.bytesMoved)
}

// SetBytesMoved sets the value in bytesMoved.
func (p DefragmentationStats) SetBytesMoved(value vk.DeviceSize) {
	p.ptr.bytesMoved = (C.VkDeviceSize)(value)
}

// GetBytesFreed returns the value in bytesFreed.
func (p DefragmentationStats) GetBytesFreed() vk.DeviceSize {
	return vk.DeviceSize(p.ptr.bytesFreed)
}

// SetBytesFreed sets the value in bytesFreed.
func (p DefragmentationStats) SetBytesFreed(value vk.DeviceSize) {
	p.ptr.bytesFreed = (C.VkDeviceSize)(value)
}

// GetAllocationsMoved returns the value in allocationsMoved.
func (p DefragmentationStats) GetAllocationsMoved() uint32 {
	return uint32(p.ptr.allocationsMoved)
}

// SetAllocationsMoved sets the value in allocationsMoved.
func (p DefragmentationStats) SetAllocationsMoved(value uint32) {
	p.ptr.allocationsMoved = (C.uint32_t)(value)
}

// GetDeviceMemoryBlocksFreed returns the value in deviceMemoryBlocksFreed.
func (p DefragmentationStats) GetDeviceMemoryBlocksFreed() uint32 {
	return uint32(p.ptr.deviceMemoryBlocksFreed)
}

// SetDeviceMemoryBlocksFreed sets the value in deviceMemoryBlocksFreed.
func (p DefragmentationStats) SetDeviceMemoryBlocksFreed(value uint32) {
	p.ptr.deviceMemoryBlocksFreed = (C.uint32_t)(value)
}

// DetailedStatistics wraps VmaDetailedStatistics.
type DetailedStatistics struct {
	ptr *C.VmaDetailedStatistics
}

// DetailedStatisticsNil is a null pointer.
var DetailedStatisticsNil DetailedStatistics

// DetailedStatisticsSizeOf is the byte size of VmaDetailedStatistics.
const DetailedStatisticsSizeOf = int(C.sizeof_VmaDetailedStatistics)

// DetailedStatisticsFromPtr converts a raw pointer to a DetailedStatistics.
func DetailedStatisticsFromPtr(ptr unsafe.Pointer) DetailedStatistics {
	return DetailedStatistics{ptr: (*C.VmaDetailedStatistics)(ptr)}
}

// DetailedStatisticsAlloc allocates a continuous block of VmaDetailedStatistics.
func DetailedStatisticsAlloc(alloc ffi.Allocator, count int) DetailedStatistics {
	ptr := alloc.Allocate(DetailedStatisticsSizeOf * count)
	return DetailedStatistics{ptr: (*C.VmaDetailedStatistics)(ptr)}
}

// Raw returns a raw pointer to the struct.
func (p DetailedStatistics) Raw() unsafe.Pointer {
	return unsafe.Pointer(p.ptr)
}

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p DetailedStatistics) Offset(offset int) DetailedStatistics {
	ptr := unsafe.Add(unsafe.Pointer(p.ptr), offset*DetailedStatisticsSizeOf)
	return DetailedStatistics{ptr: (*C.VmaDetailedStatistics)(ptr)}
}

// RefStatistics returns pointer to the statistics field.
func (p DetailedStatistics) RefStatistics() Statistics {
	return Statistics{ptr: (*C.VmaStatistics)(unsafe.Add(unsafe.Pointer(p.ptr), uintptr(C.offsetof_VmaDetailedStatistics_statistics)))}
}

// GetUnusedRangeCount returns the value in unusedRangeCount.
func (p DetailedStatistics) GetUnusedRangeCount() uint32 {
	return uint32(p.ptr.unusedRangeCount)
}

// SetUnusedRangeCount sets the value in unusedRangeCount.
func (p DetailedStatistics) SetUnusedRangeCount(value uint32) {
	p.ptr.unusedRangeCount = (C.uint32_t)(value)
}

// GetAllocationSizeMin returns the value in allocationSizeMin.
func (p DetailedStatistics) GetAllocationSizeMin() vk.DeviceSize {
	return vk.DeviceSize(p.ptr.allocationSizeMin)
}

// SetAllocationSizeMin sets the value in allocationSizeMin.
func (p DetailedStatistics) SetAllocationSizeMin(value vk.DeviceSize) {
	p.ptr.allocationSizeMin = (C.VkDeviceSize)(value)
}

// GetAllocationSizeMax returns the value in allocationSizeMax.
func (p DetailedStatistics) GetAllocationSizeMax() vk.DeviceSize {
	return vk.DeviceSize(p.ptr.allocationSizeMax)
}

// SetAllocationSizeMax sets the value in allocationSizeMax.
func (p DetailedStatistics) SetAllocationSizeMax(value vk.DeviceSize) {
	p.ptr.allocationSizeMax = (C.VkDeviceSize)(value)
}

// GetUnusedRangeSizeMin returns the value in unusedRangeSizeMin.
func (p DetailedStatistics) GetUnusedRangeSizeMin() vk.DeviceSize {
	return vk.DeviceSize(p.ptr.unusedRangeSizeMin)
}

// SetUnusedRangeSizeMin sets the value in unusedRangeSizeMin.
func (p DetailedStatistics) SetUnusedRangeSizeMin(value vk.DeviceSize) {
	p.ptr.unusedRangeSizeMin = (C.VkDeviceSize)(value)
}

// GetUnusedRangeSizeMax returns the value in unusedRangeSizeMax.
func (p DetailedStatistics) GetUnusedRangeSizeMax() vk.DeviceSize {
	return vk.DeviceSize(p.ptr.unusedRangeSizeMax)
}

// SetUnusedRangeSizeMax sets the value in unusedRangeSizeMax.
func (p DetailedStatistics) SetUnusedRangeSizeMax(value vk.DeviceSize) {
	p.ptr.unusedRangeSizeMax = (C.VkDeviceSize)(value)
}

// DeviceMemoryCallbacks wraps VmaDeviceMemoryCallbacks.
type DeviceMemoryCallbacks struct {
	ptr *C.VmaDeviceMemoryCallbacks
}

// DeviceMemoryCallbacksNil is a null pointer.
var DeviceMemoryCallbacksNil DeviceMemoryCallbacks

// DeviceMemoryCallbacksSizeOf is the byte size of VmaDeviceMemoryCallbacks.
const DeviceMemoryCallbacksSizeOf = int(C.sizeof_VmaDeviceMemoryCallbacks)

// DeviceMemoryCallbacksFromPtr converts a raw pointer to a DeviceMemoryCallbacks.
func DeviceMemoryCallbacksFromPtr(ptr unsafe.Pointer) DeviceMemoryCallbacks {
	return DeviceMemoryCallbacks{ptr: (*C.VmaDeviceMemoryCallbacks)(ptr)}
}

// DeviceMemoryCallbacksAlloc allocates a continuous block of VmaDeviceMemoryCallbacks.
func DeviceMemoryCallbacksAlloc(alloc ffi.Allocator, count int) DeviceMemoryCallbacks {
	ptr := alloc.Allocate(DeviceMemoryCallbacksSizeOf * count)
	return DeviceMemoryCallbacks{ptr: (*C.VmaDeviceMemoryCallbacks)(ptr)}
}

// Raw returns a raw pointer to the struct.
func (p DeviceMemoryCallbacks) Raw() unsafe.Pointer {
	return unsafe.Pointer(p.ptr)
}

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p DeviceMemoryCallbacks) Offset(offset int) DeviceMemoryCallbacks {
	ptr := unsafe.Add(unsafe.Pointer(p.ptr), offset*DeviceMemoryCallbacksSizeOf)
	return DeviceMemoryCallbacks{ptr: (*C.VmaDeviceMemoryCallbacks)(ptr)}
}

// DeviceMemoryCallbacks.pfnAllocate is unsupported: direct category unknown type PFN_vmaAllocateDeviceMemoryFunction.

// DeviceMemoryCallbacks.pfnFree is unsupported: direct category unknown type PFN_vmaFreeDeviceMemoryFunction.

// GetPUserData returns the value in pUserData.
func (p DeviceMemoryCallbacks) GetPUserData() unsafe.Pointer {
	return unsafe.Pointer(p.ptr.pUserData)
}

// SetPUserData sets the value in pUserData.
func (p DeviceMemoryCallbacks) SetPUserData(value unsafe.Pointer) {
	p.ptr.pUserData = value
}

// PoolCreateInfo wraps VmaPoolCreateInfo.
type PoolCreateInfo struct {
	ptr *C.VmaPoolCreateInfo
}

// PoolCreateInfoNil is a null pointer.
var PoolCreateInfoNil PoolCreateInfo

// PoolCreateInfoSizeOf is the byte size of VmaPoolCreateInfo.
const PoolCreateInfoSizeOf = int(C.sizeof_VmaPoolCreateInfo)

// PoolCreateInfoFromPtr converts a raw pointer to a PoolCreateInfo.
func PoolCreateInfoFromPtr(ptr unsafe.Pointer) PoolCreateInfo {
	return PoolCreateInfo{ptr: (*C.VmaPoolCreateInfo)(ptr)}
}

// PoolCreateInfoAlloc allocates a continuous block of VmaPoolCreateInfo.
func PoolCreateInfoAlloc(alloc ffi.Allocator, count int) PoolCreateInfo {
	ptr := alloc.Allocate(PoolCreateInfoSizeOf * count)
	return PoolCreateInfo{ptr: (*C.VmaPoolCreateInfo)(ptr)}
}

// Raw returns a raw pointer to the struct.
func (p PoolCreateInfo) Raw() unsafe.Pointer {
	return unsafe.Pointer(p.ptr)
}

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p PoolCreateInfo) Offset(offset int) PoolCreateInfo {
	ptr := unsafe.Add(unsafe.Pointer(p.ptr), offset*PoolCreateInfoSizeOf)
	return PoolCreateInfo{ptr: (*C.VmaPoolCreateInfo)(ptr)}
}

// GetMemoryTypeIndex returns the value in memoryTypeIndex.
func (p PoolCreateInfo) GetMemoryTypeIndex() uint32 {
	return uint32(p.ptr.memoryTypeIndex)
}

// SetMemoryTypeIndex sets the value in memoryTypeIndex.
func (p PoolCreateInfo) SetMemoryTypeIndex(value uint32) {
	p.ptr.memoryTypeIndex = (C.uint32_t)(value)
}

// GetFlags returns the value in flags.
func (p PoolCreateInfo) GetFlags() PoolCreateFlags {
	return PoolCreateFlags(p.ptr.flags)
}

// SetFlags sets the value in flags.
func (p PoolCreateInfo) SetFlags(value PoolCreateFlags) {
	p.ptr.flags = (C.VmaPoolCreateFlags)(value)
}

// GetBlockSize returns the value in blockSize.
func (p PoolCreateInfo) GetBlockSize() vk.DeviceSize {
	return vk.DeviceSize(p.ptr.blockSize)
}

// SetBlockSize sets the value in blockSize.
func (p PoolCreateInfo) SetBlockSize(value vk.DeviceSize) {
	p.ptr.blockSize = (C.VkDeviceSize)(value)
}

// GetMinBlockCount returns the value in minBlockCount.
func (p PoolCreateInfo) GetMinBlockCount() uintptr {
	return uintptr(p.ptr.minBlockCount)
}

// SetMinBlockCount sets the value in minBlockCount.
func (p PoolCreateInfo) SetMinBlockCount(value uintptr) {
	p.ptr.minBlockCount = (C.size_t)(value)
}

// GetMaxBlockCount returns the value in maxBlockCount.
func (p PoolCreateInfo) GetMaxBlockCount() uintptr {
	return uintptr(p.ptr.maxBlockCount)
}

// SetMaxBlockCount sets the value in maxBlockCount.
func (p PoolCreateInfo) SetMaxBlockCount(value uintptr) {
	p.ptr.maxBlockCount = (C.size_t)(value)
}

// GetPriority returns the value in priority.
func (p PoolCreateInfo) GetPriority() float32 {
	return float32(p.ptr.priority)
}

// SetPriority sets the value in priority.
func (p PoolCreateInfo) SetPriority(value float32) {
	p.ptr.priority = (C.float)(value)
}

// GetMinAllocationAlignment returns the value in minAllocationAlignment.
func (p PoolCreateInfo) GetMinAllocationAlignment() vk.DeviceSize {
	return vk.DeviceSize(p.ptr.minAllocationAlignment)
}

// SetMinAllocationAlignment sets the value in minAllocationAlignment.
func (p PoolCreateInfo) SetMinAllocationAlignment(value vk.DeviceSize) {
	p.ptr.minAllocationAlignment = (C.VkDeviceSize)(value)
}

// GetPMemoryAllocateNext returns the value in pMemoryAllocateNext.
func (p PoolCreateInfo) GetPMemoryAllocateNext() unsafe.Pointer {
	return unsafe.Pointer(p.ptr.pMemoryAllocateNext)
}

// SetPMemoryAllocateNext sets the value in pMemoryAllocateNext.
func (p PoolCreateInfo) SetPMemoryAllocateNext(value unsafe.Pointer) {
	p.ptr.pMemoryAllocateNext = value
}

// Statistics wraps VmaStatistics.
type Statistics struct {
	ptr *C.VmaStatistics
}

// StatisticsNil is a null pointer.
var StatisticsNil Statistics

// StatisticsSizeOf is the byte size of VmaStatistics.
const StatisticsSizeOf = int(C.sizeof_VmaStatistics)

// StatisticsFromPtr converts a raw pointer to a Statistics.
func StatisticsFromPtr(ptr unsafe.Pointer) Statistics {
	return Statistics{ptr: (*C.VmaStatistics)(ptr)}
}

// StatisticsAlloc allocates a continuous block of VmaStatistics.
func StatisticsAlloc(alloc ffi.Allocator, count int) Statistics {
	ptr := alloc.Allocate(StatisticsSizeOf * count)
	return Statistics{ptr: (*C.VmaStatistics)(ptr)}
}

// Raw returns a raw pointer to the struct.
func (p Statistics) Raw() unsafe.Pointer {
	return unsafe.Pointer(p.ptr)
}

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p Statistics) Offset(offset int) Statistics {
	ptr := unsafe.Add(unsafe.Pointer(p.ptr), offset*StatisticsSizeOf)
	return Statistics{ptr: (*C.VmaStatistics)(ptr)}
}

// GetBlockCount returns the value in blockCount.
func (p Statistics) GetBlockCount() uint32 {
	return uint32(p.ptr.blockCount)
}

// SetBlockCount sets the value in blockCount.
func (p Statistics) SetBlockCount(value uint32) {
	p.ptr.blockCount = (C.uint32_t)(value)
}

// GetAllocationCount returns the value in allocationCount.
func (p Statistics) GetAllocationCount() uint32 {
	return uint32(p.ptr.allocationCount)
}

// SetAllocationCount sets the value in allocationCount.
func (p Statistics) SetAllocationCount(value uint32) {
	p.ptr.allocationCount = (C.uint32_t)(value)
}

// GetBlockBytes returns the value in blockBytes.
func (p Statistics) GetBlockBytes() vk.DeviceSize {
	return vk.DeviceSize(p.ptr.blockBytes)
}

// SetBlockBytes sets the value in blockBytes.
func (p Statistics) SetBlockBytes(value vk.DeviceSize) {
	p.ptr.blockBytes = (C.VkDeviceSize)(value)
}

// GetAllocationBytes returns the value in allocationBytes.
func (p Statistics) GetAllocationBytes() vk.DeviceSize {
	return vk.DeviceSize(p.ptr.allocationBytes)
}

// SetAllocationBytes sets the value in allocationBytes.
func (p Statistics) SetAllocationBytes(value vk.DeviceSize) {
	p.ptr.allocationBytes = (C.VkDeviceSize)(value)
}

// TotalStatistics wraps VmaTotalStatistics.
type TotalStatistics struct {
	ptr *C.VmaTotalStatistics
}

// TotalStatisticsNil is a null pointer.
var TotalStatisticsNil TotalStatistics

// TotalStatisticsSizeOf is the byte size of VmaTotalStatistics.
const TotalStatisticsSizeOf = int(C.sizeof_VmaTotalStatistics)

// TotalStatisticsFromPtr converts a raw pointer to a TotalStatistics.
func TotalStatisticsFromPtr(ptr unsafe.Pointer) TotalStatistics {
	return TotalStatistics{ptr: (*C.VmaTotalStatistics)(ptr)}
}

// TotalStatisticsAlloc allocates a continuous block of VmaTotalStatistics.
func TotalStatisticsAlloc(alloc ffi.Allocator, count int) TotalStatistics {
	ptr := alloc.Allocate(TotalStatisticsSizeOf * count)
	return TotalStatistics{ptr: (*C.VmaTotalStatistics)(ptr)}
}

// Raw returns a raw pointer to the struct.
func (p TotalStatistics) Raw() unsafe.Pointer {
	return unsafe.Pointer(p.ptr)
}

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p TotalStatistics) Offset(offset int) TotalStatistics {
	ptr := unsafe.Add(unsafe.Pointer(p.ptr), offset*TotalStatisticsSizeOf)
	return TotalStatistics{ptr: (*C.VmaTotalStatistics)(ptr)}
}

// TotalStatistics.memoryType is unsupported: category array.

// TotalStatistics.memoryHeap is unsupported: category array.

// RefTotal returns pointer to the total field.
func (p TotalStatistics) RefTotal() DetailedStatistics {
	return DetailedStatistics{ptr: (*C.VmaDetailedStatistics)(unsafe.Add(unsafe.Pointer(p.ptr), uintptr(C.offsetof_VmaTotalStatistics_total)))}
}

// VirtualAllocationCreateInfo wraps VmaVirtualAllocationCreateInfo.
type VirtualAllocationCreateInfo struct {
	ptr *C.VmaVirtualAllocationCreateInfo
}

// VirtualAllocationCreateInfoNil is a null pointer.
var VirtualAllocationCreateInfoNil VirtualAllocationCreateInfo

// VirtualAllocationCreateInfoSizeOf is the byte size of VmaVirtualAllocationCreateInfo.
const VirtualAllocationCreateInfoSizeOf = int(C.sizeof_VmaVirtualAllocationCreateInfo)

// VirtualAllocationCreateInfoFromPtr converts a raw pointer to a VirtualAllocationCreateInfo.
func VirtualAllocationCreateInfoFromPtr(ptr unsafe.Pointer) VirtualAllocationCreateInfo {
	return VirtualAllocationCreateInfo{ptr: (*C.VmaVirtualAllocationCreateInfo)(ptr)}
}

// VirtualAllocationCreateInfoAlloc allocates a continuous block of VmaVirtualAllocationCreateInfo.
func VirtualAllocationCreateInfoAlloc(alloc ffi.Allocator, count int) VirtualAllocationCreateInfo {
	ptr := alloc.Allocate(VirtualAllocationCreateInfoSizeOf * count)
	return VirtualAllocationCreateInfo{ptr: (*C.VmaVirtualAllocationCreateInfo)(ptr)}
}

// Raw returns a raw pointer to the struct.
func (p VirtualAllocationCreateInfo) Raw() unsafe.Pointer {
	return unsafe.Pointer(p.ptr)
}

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p VirtualAllocationCreateInfo) Offset(offset int) VirtualAllocationCreateInfo {
	ptr := unsafe.Add(unsafe.Pointer(p.ptr), offset*VirtualAllocationCreateInfoSizeOf)
	return VirtualAllocationCreateInfo{ptr: (*C.VmaVirtualAllocationCreateInfo)(ptr)}
}

// GetSize returns the value in size.
func (p VirtualAllocationCreateInfo) GetSize() vk.DeviceSize {
	return vk.DeviceSize(p.ptr.size)
}

// SetSize sets the value in size.
func (p VirtualAllocationCreateInfo) SetSize(value vk.DeviceSize) {
	p.ptr.size = (C.VkDeviceSize)(value)
}

// GetAlignment returns the value in alignment.
func (p VirtualAllocationCreateInfo) GetAlignment() vk.DeviceSize {
	return vk.DeviceSize(p.ptr.alignment)
}

// SetAlignment sets the value in alignment.
func (p VirtualAllocationCreateInfo) SetAlignment(value vk.DeviceSize) {
	p.ptr.alignment = (C.VkDeviceSize)(value)
}

// GetFlags returns the value in flags.
func (p VirtualAllocationCreateInfo) GetFlags() VirtualAllocationCreateFlags {
	return VirtualAllocationCreateFlags(p.ptr.flags)
}

// SetFlags sets the value in flags.
func (p VirtualAllocationCreateInfo) SetFlags(value VirtualAllocationCreateFlags) {
	p.ptr.flags = (C.VmaVirtualAllocationCreateFlags)(value)
}

// GetPUserData returns the value in pUserData.
func (p VirtualAllocationCreateInfo) GetPUserData() unsafe.Pointer {
	return unsafe.Pointer(p.ptr.pUserData)
}

// SetPUserData sets the value in pUserData.
func (p VirtualAllocationCreateInfo) SetPUserData(value unsafe.Pointer) {
	p.ptr.pUserData = value
}

// VirtualAllocationInfo wraps VmaVirtualAllocationInfo.
type VirtualAllocationInfo struct {
	ptr *C.VmaVirtualAllocationInfo
}

// VirtualAllocationInfoNil is a null pointer.
var VirtualAllocationInfoNil VirtualAllocationInfo

// VirtualAllocationInfoSizeOf is the byte size of VmaVirtualAllocationInfo.
const VirtualAllocationInfoSizeOf = int(C.sizeof_VmaVirtualAllocationInfo)

// VirtualAllocationInfoFromPtr converts a raw pointer to a VirtualAllocationInfo.
func VirtualAllocationInfoFromPtr(ptr unsafe.Pointer) VirtualAllocationInfo {
	return VirtualAllocationInfo{ptr: (*C.VmaVirtualAllocationInfo)(ptr)}
}

// VirtualAllocationInfoAlloc allocates a continuous block of VmaVirtualAllocationInfo.
func VirtualAllocationInfoAlloc(alloc ffi.Allocator, count int) VirtualAllocationInfo {
	ptr := alloc.Allocate(VirtualAllocationInfoSizeOf * count)
	return VirtualAllocationInfo{ptr: (*C.VmaVirtualAllocationInfo)(ptr)}
}

// Raw returns a raw pointer to the struct.
func (p VirtualAllocationInfo) Raw() unsafe.Pointer {
	return unsafe.Pointer(p.ptr)
}

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p VirtualAllocationInfo) Offset(offset int) VirtualAllocationInfo {
	ptr := unsafe.Add(unsafe.Pointer(p.ptr), offset*VirtualAllocationInfoSizeOf)
	return VirtualAllocationInfo{ptr: (*C.VmaVirtualAllocationInfo)(ptr)}
}

// GetOffset returns the value in offset.
func (p VirtualAllocationInfo) GetOffset() vk.DeviceSize {
	return vk.DeviceSize(p.ptr.offset)
}

// SetOffset sets the value in offset.
func (p VirtualAllocationInfo) SetOffset(value vk.DeviceSize) {
	p.ptr.offset = (C.VkDeviceSize)(value)
}

// GetSize returns the value in size.
func (p VirtualAllocationInfo) GetSize() vk.DeviceSize {
	return vk.DeviceSize(p.ptr.size)
}

// SetSize sets the value in size.
func (p VirtualAllocationInfo) SetSize(value vk.DeviceSize) {
	p.ptr.size = (C.VkDeviceSize)(value)
}

// GetPUserData returns the value in pUserData.
func (p VirtualAllocationInfo) GetPUserData() unsafe.Pointer {
	return unsafe.Pointer(p.ptr.pUserData)
}

// SetPUserData sets the value in pUserData.
func (p VirtualAllocationInfo) SetPUserData(value unsafe.Pointer) {
	p.ptr.pUserData = value
}

// VirtualBlockCreateInfo wraps VmaVirtualBlockCreateInfo.
type VirtualBlockCreateInfo struct {
	ptr *C.VmaVirtualBlockCreateInfo
}

// VirtualBlockCreateInfoNil is a null pointer.
var VirtualBlockCreateInfoNil VirtualBlockCreateInfo

// VirtualBlockCreateInfoSizeOf is the byte size of VmaVirtualBlockCreateInfo.
const VirtualBlockCreateInfoSizeOf = int(C.sizeof_VmaVirtualBlockCreateInfo)

// VirtualBlockCreateInfoFromPtr converts a raw pointer to a VirtualBlockCreateInfo.
func VirtualBlockCreateInfoFromPtr(ptr unsafe.Pointer) VirtualBlockCreateInfo {
	return VirtualBlockCreateInfo{ptr: (*C.VmaVirtualBlockCreateInfo)(ptr)}
}

// VirtualBlockCreateInfoAlloc allocates a continuous block of VmaVirtualBlockCreateInfo.
func VirtualBlockCreateInfoAlloc(alloc ffi.Allocator, count int) VirtualBlockCreateInfo {
	ptr := alloc.Allocate(VirtualBlockCreateInfoSizeOf * count)
	return VirtualBlockCreateInfo{ptr: (*C.VmaVirtualBlockCreateInfo)(ptr)}
}

// Raw returns a raw pointer to the struct.
func (p VirtualBlockCreateInfo) Raw() unsafe.Pointer {
	return unsafe.Pointer(p.ptr)
}

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p VirtualBlockCreateInfo) Offset(offset int) VirtualBlockCreateInfo {
	ptr := unsafe.Add(unsafe.Pointer(p.ptr), offset*VirtualBlockCreateInfoSizeOf)
	return VirtualBlockCreateInfo{ptr: (*C.VmaVirtualBlockCreateInfo)(ptr)}
}

// GetSize returns the value in size.
func (p VirtualBlockCreateInfo) GetSize() vk.DeviceSize {
	return vk.DeviceSize(p.ptr.size)
}

// SetSize sets the value in size.
func (p VirtualBlockCreateInfo) SetSize(value vk.DeviceSize) {
	p.ptr.size = (C.VkDeviceSize)(value)
}

// GetFlags returns the value in flags.
func (p VirtualBlockCreateInfo) GetFlags() VirtualBlockCreateFlags {
	return VirtualBlockCreateFlags(p.ptr.flags)
}

// SetFlags sets the value in flags.
func (p VirtualBlockCreateInfo) SetFlags(value VirtualBlockCreateFlags) {
	p.ptr.flags = (C.VmaVirtualBlockCreateFlags)(value)
}

// VirtualBlockCreateInfo.pAllocationCallbacks is unsupported: pointer category vk struct.

// VulkanFunctions wraps VmaVulkanFunctions.
type VulkanFunctions struct {
	ptr *C.VmaVulkanFunctions
}

// VulkanFunctionsNil is a null pointer.
var VulkanFunctionsNil VulkanFunctions

// VulkanFunctionsSizeOf is the byte size of VmaVulkanFunctions.
const VulkanFunctionsSizeOf = int(C.sizeof_VmaVulkanFunctions)

// VulkanFunctionsFromPtr converts a raw pointer to a VulkanFunctions.
func VulkanFunctionsFromPtr(ptr unsafe.Pointer) VulkanFunctions {
	return VulkanFunctions{ptr: (*C.VmaVulkanFunctions)(ptr)}
}

// VulkanFunctionsAlloc allocates a continuous block of VmaVulkanFunctions.
func VulkanFunctionsAlloc(alloc ffi.Allocator, count int) VulkanFunctions {
	ptr := alloc.Allocate(VulkanFunctionsSizeOf * count)
	return VulkanFunctions{ptr: (*C.VmaVulkanFunctions)(ptr)}
}

// Raw returns a raw pointer to the struct.
func (p VulkanFunctions) Raw() unsafe.Pointer {
	return unsafe.Pointer(p.ptr)
}

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p VulkanFunctions) Offset(offset int) VulkanFunctions {
	ptr := unsafe.Add(unsafe.Pointer(p.ptr), offset*VulkanFunctionsSizeOf)
	return VulkanFunctions{ptr: (*C.VmaVulkanFunctions)(ptr)}
}

// VulkanFunctions.vkGetInstanceProcAddr is unsupported: direct category unknown type PFN_vkGetInstanceProcAddr.

// VulkanFunctions.vkGetDeviceProcAddr is unsupported: direct category unknown type PFN_vkGetDeviceProcAddr.

// VulkanFunctions.vkGetPhysicalDeviceProperties is unsupported: direct category unknown type PFN_vkGetPhysicalDeviceProperties.

// VulkanFunctions.vkGetPhysicalDeviceMemoryProperties is unsupported: direct category unknown type PFN_vkGetPhysicalDeviceMemoryProperties.

// VulkanFunctions.vkAllocateMemory is unsupported: direct category unknown type PFN_vkAllocateMemory.

// VulkanFunctions.vkFreeMemory is unsupported: direct category unknown type PFN_vkFreeMemory.

// VulkanFunctions.vkMapMemory is unsupported: direct category unknown type PFN_vkMapMemory.

// VulkanFunctions.vkUnmapMemory is unsupported: direct category unknown type PFN_vkUnmapMemory.

// VulkanFunctions.vkFlushMappedMemoryRanges is unsupported: direct category unknown type PFN_vkFlushMappedMemoryRanges.

// VulkanFunctions.vkInvalidateMappedMemoryRanges is unsupported: direct category unknown type PFN_vkInvalidateMappedMemoryRanges.

// VulkanFunctions.vkBindBufferMemory is unsupported: direct category unknown type PFN_vkBindBufferMemory.

// VulkanFunctions.vkBindImageMemory is unsupported: direct category unknown type PFN_vkBindImageMemory.

// VulkanFunctions.vkGetBufferMemoryRequirements is unsupported: direct category unknown type PFN_vkGetBufferMemoryRequirements.

// VulkanFunctions.vkGetImageMemoryRequirements is unsupported: direct category unknown type PFN_vkGetImageMemoryRequirements.

// VulkanFunctions.vkCreateBuffer is unsupported: direct category unknown type PFN_vkCreateBuffer.

// VulkanFunctions.vkDestroyBuffer is unsupported: direct category unknown type PFN_vkDestroyBuffer.

// VulkanFunctions.vkCreateImage is unsupported: direct category unknown type PFN_vkCreateImage.

// VulkanFunctions.vkDestroyImage is unsupported: direct category unknown type PFN_vkDestroyImage.

// VulkanFunctions.vkCmdCopyBuffer is unsupported: direct category unknown type PFN_vkCmdCopyBuffer.

// VulkanFunctions.vkGetBufferMemoryRequirements2KHR is unsupported: direct category unknown type PFN_vkGetBufferMemoryRequirements2KHR.

// VulkanFunctions.vkGetImageMemoryRequirements2KHR is unsupported: direct category unknown type PFN_vkGetImageMemoryRequirements2KHR.

// VulkanFunctions.vkBindBufferMemory2KHR is unsupported: direct category unknown type PFN_vkBindBufferMemory2KHR.

// VulkanFunctions.vkBindImageMemory2KHR is unsupported: direct category unknown type PFN_vkBindImageMemory2KHR.

// VulkanFunctions.vkGetPhysicalDeviceMemoryProperties2KHR is unsupported: direct category unknown type PFN_vkGetPhysicalDeviceMemoryProperties2KHR.

// VulkanFunctions.vkGetDeviceBufferMemoryRequirements is unsupported: direct category unknown type PFN_vkGetDeviceBufferMemoryRequirementsKHR.

// VulkanFunctions.vkGetDeviceImageMemoryRequirements is unsupported: direct category unknown type PFN_vkGetDeviceImageMemoryRequirementsKHR.

// GetVkGetMemoryWin32HandleKHR returns the value in vkGetMemoryWin32HandleKHR.
func (p VulkanFunctions) GetVkGetMemoryWin32HandleKHR() unsafe.Pointer {
	return unsafe.Pointer(p.ptr.vkGetMemoryWin32HandleKHR)
}

// SetVkGetMemoryWin32HandleKHR sets the value in vkGetMemoryWin32HandleKHR.
func (p VulkanFunctions) SetVkGetMemoryWin32HandleKHR(value unsafe.Pointer) {
	p.ptr.vkGetMemoryWin32HandleKHR = value
}
