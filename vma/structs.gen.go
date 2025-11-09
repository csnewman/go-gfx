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

// AllocationCreateInfo wraps struct VmaAllocationCreateInfo.
/*
  \brief Parameters of new #VmaAllocation.

  To be used with functions like vmaCreateBuffer(), vmaCreateImage(), and many others.
*/
type AllocationCreateInfo uintptr

// AllocationCreateInfoNil is a null pointer.
var AllocationCreateInfoNil AllocationCreateInfo

// AllocationCreateInfoSizeOf is the byte size of VmaAllocationCreateInfo.
const AllocationCreateInfoSizeOf = int(C.sizeof_VmaAllocationCreateInfo)

// AllocationCreateInfoAlloc allocates a continuous block of AllocationCreateInfo.
func AllocationCreateInfoAlloc(alloc ffi.Allocator, count int) AllocationCreateInfo {
	ptr := alloc.Allocate(AllocationCreateInfoSizeOf * count)
	return AllocationCreateInfo(ptr)
}

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p AllocationCreateInfo) Offset(offset int) AllocationCreateInfo {
	return p + AllocationCreateInfo(offset*AllocationCreateInfoSizeOf)
}

// GetFlags returns the value in flags.
func (p AllocationCreateInfo) GetFlags() AllocationCreateFlags {
	ptr := (*C.VmaAllocationCreateInfo)(unsafe.Pointer(p))
	return AllocationCreateFlags(ptr.flags)
}

// SetFlags sets the value in flags.
func (p AllocationCreateInfo) SetFlags(value AllocationCreateFlags) {
	ptr := (*C.VmaAllocationCreateInfo)(unsafe.Pointer(p))
	ptr.flags = (C.VmaAllocationCreateFlags)(value)
}

// GetMemoryTypeBits returns the value in memoryTypeBits.
func (p AllocationCreateInfo) GetMemoryTypeBits() uint32 {
	ptr := (*C.VmaAllocationCreateInfo)(unsafe.Pointer(p))
	return uint32(ptr.memoryTypeBits)
}

// SetMemoryTypeBits sets the value in memoryTypeBits.
func (p AllocationCreateInfo) SetMemoryTypeBits(value uint32) {
	ptr := (*C.VmaAllocationCreateInfo)(unsafe.Pointer(p))
	ptr.memoryTypeBits = (C.uint32_t)(value)
}

// GetPUserData returns the value in pUserData.
func (p AllocationCreateInfo) GetPUserData() uintptr {
	ptr := (*C.VmaAllocationCreateInfo)(unsafe.Pointer(p))
	return uintptr(unsafe.Pointer(ptr.pUserData))
}

// SetPUserData sets the value in pUserData.
func (p AllocationCreateInfo) SetPUserData(value uintptr) {
	ptr := (*C.VmaAllocationCreateInfo)(unsafe.Pointer(p))
	ptr.pUserData = unsafe.Pointer(value)
}

// GetPool returns the value in pool.
func (p AllocationCreateInfo) GetPool() Pool {
	ptr := (*C.VmaAllocationCreateInfo)(unsafe.Pointer(p))
	return Pool(unsafe.Pointer(ptr.pool))
}

// SetPool sets the value in pool.
func (p AllocationCreateInfo) SetPool(value Pool) {
	ptr := (*C.VmaAllocationCreateInfo)(unsafe.Pointer(p))
	ptr.pool = (C.VmaPool)(unsafe.Pointer(value))
}

// GetPreferredFlags returns the value in preferredFlags.
func (p AllocationCreateInfo) GetPreferredFlags() vk.MemoryPropertyFlags {
	ptr := (*C.VmaAllocationCreateInfo)(unsafe.Pointer(p))
	return vk.MemoryPropertyFlags(ptr.preferredFlags)
}

// SetPreferredFlags sets the value in preferredFlags.
func (p AllocationCreateInfo) SetPreferredFlags(value vk.MemoryPropertyFlags) {
	ptr := (*C.VmaAllocationCreateInfo)(unsafe.Pointer(p))
	ptr.preferredFlags = (C.VkMemoryPropertyFlags)(value)
}

// GetPriority returns the value in priority.
func (p AllocationCreateInfo) GetPriority() float32 {
	ptr := (*C.VmaAllocationCreateInfo)(unsafe.Pointer(p))
	return float32(ptr.priority)
}

// SetPriority sets the value in priority.
func (p AllocationCreateInfo) SetPriority(value float32) {
	ptr := (*C.VmaAllocationCreateInfo)(unsafe.Pointer(p))
	ptr.priority = (C.float)(value)
}

// GetRequiredFlags returns the value in requiredFlags.
func (p AllocationCreateInfo) GetRequiredFlags() vk.MemoryPropertyFlags {
	ptr := (*C.VmaAllocationCreateInfo)(unsafe.Pointer(p))
	return vk.MemoryPropertyFlags(ptr.requiredFlags)
}

// SetRequiredFlags sets the value in requiredFlags.
func (p AllocationCreateInfo) SetRequiredFlags(value vk.MemoryPropertyFlags) {
	ptr := (*C.VmaAllocationCreateInfo)(unsafe.Pointer(p))
	ptr.requiredFlags = (C.VkMemoryPropertyFlags)(value)
}

// GetUsage returns the value in usage.
func (p AllocationCreateInfo) GetUsage() MemoryUsage {
	ptr := (*C.VmaAllocationCreateInfo)(unsafe.Pointer(p))
	return MemoryUsage(ptr.usage)
}

// SetUsage sets the value in usage.
func (p AllocationCreateInfo) SetUsage(value MemoryUsage) {
	ptr := (*C.VmaAllocationCreateInfo)(unsafe.Pointer(p))
	ptr.usage = (C.VmaMemoryUsage)(value)
}

// AllocationInfo wraps struct VmaAllocationInfo.
/*
  Parameters of #VmaAllocation objects, that can be retrieved using function vmaGetAllocationInfo().

  There is also an extended version of this structure that carries additional parameters: #VmaAllocationInfo2.
*/
type AllocationInfo uintptr

// AllocationInfoNil is a null pointer.
var AllocationInfoNil AllocationInfo

// AllocationInfoSizeOf is the byte size of VmaAllocationInfo.
const AllocationInfoSizeOf = int(C.sizeof_VmaAllocationInfo)

// AllocationInfoAlloc allocates a continuous block of AllocationInfo.
func AllocationInfoAlloc(alloc ffi.Allocator, count int) AllocationInfo {
	ptr := alloc.Allocate(AllocationInfoSizeOf * count)
	return AllocationInfo(ptr)
}

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p AllocationInfo) Offset(offset int) AllocationInfo {
	return p + AllocationInfo(offset*AllocationInfoSizeOf)
}

// GetDeviceMemory returns the value in deviceMemory.
func (p AllocationInfo) GetDeviceMemory() vk.DeviceMemory {
	ptr := (*C.VmaAllocationInfo)(unsafe.Pointer(p))
	return vk.DeviceMemory(ptr.deviceMemory)
}

// SetDeviceMemory sets the value in deviceMemory.
func (p AllocationInfo) SetDeviceMemory(value vk.DeviceMemory) {
	ptr := (*C.VmaAllocationInfo)(unsafe.Pointer(p))
	ptr.deviceMemory = (C.VkDeviceMemory)(value)
}

// GetMemoryType returns the value in memoryType.
func (p AllocationInfo) GetMemoryType() uint32 {
	ptr := (*C.VmaAllocationInfo)(unsafe.Pointer(p))
	return uint32(ptr.memoryType)
}

// SetMemoryType sets the value in memoryType.
func (p AllocationInfo) SetMemoryType(value uint32) {
	ptr := (*C.VmaAllocationInfo)(unsafe.Pointer(p))
	ptr.memoryType = (C.uint32_t)(value)
}

// GetOffset returns the value in offset.
func (p AllocationInfo) GetOffset() vk.DeviceSize {
	ptr := (*C.VmaAllocationInfo)(unsafe.Pointer(p))
	return vk.DeviceSize(ptr.offset)
}

// SetOffset sets the value in offset.
func (p AllocationInfo) SetOffset(value vk.DeviceSize) {
	ptr := (*C.VmaAllocationInfo)(unsafe.Pointer(p))
	ptr.offset = (C.VkDeviceSize)(value)
}

// GetPMappedData returns the value in pMappedData.
func (p AllocationInfo) GetPMappedData() uintptr {
	ptr := (*C.VmaAllocationInfo)(unsafe.Pointer(p))
	return uintptr(unsafe.Pointer(ptr.pMappedData))
}

// SetPMappedData sets the value in pMappedData.
func (p AllocationInfo) SetPMappedData(value uintptr) {
	ptr := (*C.VmaAllocationInfo)(unsafe.Pointer(p))
	ptr.pMappedData = unsafe.Pointer(value)
}

// GetPName returns the value in pName.
func (p AllocationInfo) GetPName() ffi.CString {
	ptr := (*C.VmaAllocationInfo)(unsafe.Pointer(p))
	return ffi.CStringFromPtr((unsafe.Pointer)(ptr.pName))
}

// SetPName sets the value in pName.
func (p AllocationInfo) SetPName(value ffi.CString) {
	ptr := (*C.VmaAllocationInfo)(unsafe.Pointer(p))
	ptr.pName = (*C.char)(value.Raw())
}

// GetPUserData returns the value in pUserData.
func (p AllocationInfo) GetPUserData() uintptr {
	ptr := (*C.VmaAllocationInfo)(unsafe.Pointer(p))
	return uintptr(unsafe.Pointer(ptr.pUserData))
}

// SetPUserData sets the value in pUserData.
func (p AllocationInfo) SetPUserData(value uintptr) {
	ptr := (*C.VmaAllocationInfo)(unsafe.Pointer(p))
	ptr.pUserData = unsafe.Pointer(value)
}

// GetSize returns the value in size.
func (p AllocationInfo) GetSize() vk.DeviceSize {
	ptr := (*C.VmaAllocationInfo)(unsafe.Pointer(p))
	return vk.DeviceSize(ptr.size)
}

// SetSize sets the value in size.
func (p AllocationInfo) SetSize(value vk.DeviceSize) {
	ptr := (*C.VmaAllocationInfo)(unsafe.Pointer(p))
	ptr.size = (C.VkDeviceSize)(value)
}

// AllocationInfo2 wraps struct VmaAllocationInfo2.
//
//	Extended parameters of a #VmaAllocation object that can be retrieved using function vmaGetAllocationInfo2().
type AllocationInfo2 uintptr

// AllocationInfo2Nil is a null pointer.
var AllocationInfo2Nil AllocationInfo2

// AllocationInfo2SizeOf is the byte size of VmaAllocationInfo2.
const AllocationInfo2SizeOf = int(C.sizeof_VmaAllocationInfo2)

// AllocationInfo2Alloc allocates a continuous block of AllocationInfo2.
func AllocationInfo2Alloc(alloc ffi.Allocator, count int) AllocationInfo2 {
	ptr := alloc.Allocate(AllocationInfo2SizeOf * count)
	return AllocationInfo2(ptr)
}

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p AllocationInfo2) Offset(offset int) AllocationInfo2 {
	return p + AllocationInfo2(offset*AllocationInfo2SizeOf)
}

// RefAllocationInfo returns pointer to the allocationInfo field.
func (p AllocationInfo2) RefAllocationInfo() AllocationInfo {
	return AllocationInfo(p + AllocationInfo2(C.offsetof_VmaAllocationInfo2_allocationInfo))
}

// GetBlockSize returns the value in blockSize.
func (p AllocationInfo2) GetBlockSize() vk.DeviceSize {
	ptr := (*C.VmaAllocationInfo2)(unsafe.Pointer(p))
	return vk.DeviceSize(ptr.blockSize)
}

// SetBlockSize sets the value in blockSize.
func (p AllocationInfo2) SetBlockSize(value vk.DeviceSize) {
	ptr := (*C.VmaAllocationInfo2)(unsafe.Pointer(p))
	ptr.blockSize = (C.VkDeviceSize)(value)
}

// GetDedicatedMemory returns the value in dedicatedMemory.
func (p AllocationInfo2) GetDedicatedMemory() bool {
	ptr := (*C.VmaAllocationInfo2)(unsafe.Pointer(p))
	return ptr.dedicatedMemory != 0
}

// SetDedicatedMemory sets the value in dedicatedMemory.
func (p AllocationInfo2) SetDedicatedMemory(value bool) {
	ptr := (*C.VmaAllocationInfo2)(unsafe.Pointer(p))
	if value {
		ptr.dedicatedMemory = C.VkBool32(1)
	} else {
		ptr.dedicatedMemory = C.VkBool32(0)
	}
}

// AllocatorCreateInfo wraps struct VmaAllocatorCreateInfo.
//
//	Description of a Allocator to be created.
type AllocatorCreateInfo uintptr

// AllocatorCreateInfoNil is a null pointer.
var AllocatorCreateInfoNil AllocatorCreateInfo

// AllocatorCreateInfoSizeOf is the byte size of VmaAllocatorCreateInfo.
const AllocatorCreateInfoSizeOf = int(C.sizeof_VmaAllocatorCreateInfo)

// AllocatorCreateInfoAlloc allocates a continuous block of AllocatorCreateInfo.
func AllocatorCreateInfoAlloc(alloc ffi.Allocator, count int) AllocatorCreateInfo {
	ptr := alloc.Allocate(AllocatorCreateInfoSizeOf * count)
	return AllocatorCreateInfo(ptr)
}

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p AllocatorCreateInfo) Offset(offset int) AllocatorCreateInfo {
	return p + AllocatorCreateInfo(offset*AllocatorCreateInfoSizeOf)
}

// GetDevice returns the value in device.
func (p AllocatorCreateInfo) GetDevice() vk.Device {
	ptr := (*C.VmaAllocatorCreateInfo)(unsafe.Pointer(p))
	return vk.Device(unsafe.Pointer(ptr.device))
}

// SetDevice sets the value in device.
func (p AllocatorCreateInfo) SetDevice(value vk.Device) {
	ptr := (*C.VmaAllocatorCreateInfo)(unsafe.Pointer(p))
	ptr.device = (C.VkDevice)(unsafe.Pointer(value))
}

// GetFlags returns the value in flags.
func (p AllocatorCreateInfo) GetFlags() AllocatorCreateFlags {
	ptr := (*C.VmaAllocatorCreateInfo)(unsafe.Pointer(p))
	return AllocatorCreateFlags(ptr.flags)
}

// SetFlags sets the value in flags.
func (p AllocatorCreateInfo) SetFlags(value AllocatorCreateFlags) {
	ptr := (*C.VmaAllocatorCreateInfo)(unsafe.Pointer(p))
	ptr.flags = (C.VmaAllocatorCreateFlags)(value)
}

// GetInstance returns the value in instance.
func (p AllocatorCreateInfo) GetInstance() vk.Instance {
	ptr := (*C.VmaAllocatorCreateInfo)(unsafe.Pointer(p))
	return vk.Instance(unsafe.Pointer(ptr.instance))
}

// SetInstance sets the value in instance.
func (p AllocatorCreateInfo) SetInstance(value vk.Instance) {
	ptr := (*C.VmaAllocatorCreateInfo)(unsafe.Pointer(p))
	ptr.instance = (C.VkInstance)(unsafe.Pointer(value))
}

// GetPAllocationCallbacks returns the value in pAllocationCallbacks.
func (p AllocatorCreateInfo) GetPAllocationCallbacks() vk.AllocationCallbacks {
	ptr := (*C.VmaAllocatorCreateInfo)(unsafe.Pointer(p))
	return vk.AllocationCallbacks(unsafe.Pointer(ptr.pAllocationCallbacks))
}

// SetPAllocationCallbacks sets the value in pAllocationCallbacks.
func (p AllocatorCreateInfo) SetPAllocationCallbacks(value vk.AllocationCallbacks) {
	ptr := (*C.VmaAllocatorCreateInfo)(unsafe.Pointer(p))
	ptr.pAllocationCallbacks = (*C.VkAllocationCallbacks)(unsafe.Pointer(value))
}

// GetPDeviceMemoryCallbacks returns the value in pDeviceMemoryCallbacks.
func (p AllocatorCreateInfo) GetPDeviceMemoryCallbacks() DeviceMemoryCallbacks {
	ptr := (*C.VmaAllocatorCreateInfo)(unsafe.Pointer(p))
	return DeviceMemoryCallbacks(unsafe.Pointer(ptr.pDeviceMemoryCallbacks))
}

// SetPDeviceMemoryCallbacks sets the value in pDeviceMemoryCallbacks.
func (p AllocatorCreateInfo) SetPDeviceMemoryCallbacks(value DeviceMemoryCallbacks) {
	ptr := (*C.VmaAllocatorCreateInfo)(unsafe.Pointer(p))
	ptr.pDeviceMemoryCallbacks = (*C.VmaDeviceMemoryCallbacks)(unsafe.Pointer(value))
}

// GetPHeapSizeLimit returns the value in pHeapSizeLimit.
func (p AllocatorCreateInfo) GetPHeapSizeLimit() ffi.Ref[vk.DeviceSize] {
	ptr := (*C.VmaAllocatorCreateInfo)(unsafe.Pointer(p))
	return ffi.RefFromPtr[vk.DeviceSize](unsafe.Pointer(ptr.pHeapSizeLimit))
}

// SetPHeapSizeLimit sets the value in pHeapSizeLimit.
func (p AllocatorCreateInfo) SetPHeapSizeLimit(value ffi.Ref[vk.DeviceSize]) {
	ptr := (*C.VmaAllocatorCreateInfo)(unsafe.Pointer(p))
	ptr.pHeapSizeLimit = (*C.VkDeviceSize)(value.Raw())
}

// GetPTypeExternalMemoryHandleTypes returns the value in pTypeExternalMemoryHandleTypes.
func (p AllocatorCreateInfo) GetPTypeExternalMemoryHandleTypes() ffi.Ref[vk.ExternalMemoryHandleTypeFlags] {
	ptr := (*C.VmaAllocatorCreateInfo)(unsafe.Pointer(p))
	return ffi.RefFromPtr[vk.ExternalMemoryHandleTypeFlags](unsafe.Pointer(ptr.pTypeExternalMemoryHandleTypes))
}

// SetPTypeExternalMemoryHandleTypes sets the value in pTypeExternalMemoryHandleTypes.
func (p AllocatorCreateInfo) SetPTypeExternalMemoryHandleTypes(value ffi.Ref[vk.ExternalMemoryHandleTypeFlags]) {
	ptr := (*C.VmaAllocatorCreateInfo)(unsafe.Pointer(p))
	ptr.pTypeExternalMemoryHandleTypes = (*C.VkExternalMemoryHandleTypeFlagsKHR)(value.Raw())
}

// GetPVulkanFunctions returns the value in pVulkanFunctions.
func (p AllocatorCreateInfo) GetPVulkanFunctions() VulkanFunctions {
	ptr := (*C.VmaAllocatorCreateInfo)(unsafe.Pointer(p))
	return VulkanFunctions(unsafe.Pointer(ptr.pVulkanFunctions))
}

// SetPVulkanFunctions sets the value in pVulkanFunctions.
func (p AllocatorCreateInfo) SetPVulkanFunctions(value VulkanFunctions) {
	ptr := (*C.VmaAllocatorCreateInfo)(unsafe.Pointer(p))
	ptr.pVulkanFunctions = (*C.VmaVulkanFunctions)(unsafe.Pointer(value))
}

// GetPhysicalDevice returns the value in physicalDevice.
func (p AllocatorCreateInfo) GetPhysicalDevice() vk.PhysicalDevice {
	ptr := (*C.VmaAllocatorCreateInfo)(unsafe.Pointer(p))
	return vk.PhysicalDevice(unsafe.Pointer(ptr.physicalDevice))
}

// SetPhysicalDevice sets the value in physicalDevice.
func (p AllocatorCreateInfo) SetPhysicalDevice(value vk.PhysicalDevice) {
	ptr := (*C.VmaAllocatorCreateInfo)(unsafe.Pointer(p))
	ptr.physicalDevice = (C.VkPhysicalDevice)(unsafe.Pointer(value))
}

// GetPreferredLargeHeapBlockSize returns the value in preferredLargeHeapBlockSize.
func (p AllocatorCreateInfo) GetPreferredLargeHeapBlockSize() vk.DeviceSize {
	ptr := (*C.VmaAllocatorCreateInfo)(unsafe.Pointer(p))
	return vk.DeviceSize(ptr.preferredLargeHeapBlockSize)
}

// SetPreferredLargeHeapBlockSize sets the value in preferredLargeHeapBlockSize.
func (p AllocatorCreateInfo) SetPreferredLargeHeapBlockSize(value vk.DeviceSize) {
	ptr := (*C.VmaAllocatorCreateInfo)(unsafe.Pointer(p))
	ptr.preferredLargeHeapBlockSize = (C.VkDeviceSize)(value)
}

// GetVulkanApiVersion returns the value in vulkanApiVersion.
func (p AllocatorCreateInfo) GetVulkanApiVersion() uint32 {
	ptr := (*C.VmaAllocatorCreateInfo)(unsafe.Pointer(p))
	return uint32(ptr.vulkanApiVersion)
}

// SetVulkanApiVersion sets the value in vulkanApiVersion.
func (p AllocatorCreateInfo) SetVulkanApiVersion(value uint32) {
	ptr := (*C.VmaAllocatorCreateInfo)(unsafe.Pointer(p))
	ptr.vulkanApiVersion = (C.uint32_t)(value)
}

// AllocatorInfo wraps struct VmaAllocatorInfo.
//
//	Information about existing #VmaAllocator object.
type AllocatorInfo uintptr

// AllocatorInfoNil is a null pointer.
var AllocatorInfoNil AllocatorInfo

// AllocatorInfoSizeOf is the byte size of VmaAllocatorInfo.
const AllocatorInfoSizeOf = int(C.sizeof_VmaAllocatorInfo)

// AllocatorInfoAlloc allocates a continuous block of AllocatorInfo.
func AllocatorInfoAlloc(alloc ffi.Allocator, count int) AllocatorInfo {
	ptr := alloc.Allocate(AllocatorInfoSizeOf * count)
	return AllocatorInfo(ptr)
}

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p AllocatorInfo) Offset(offset int) AllocatorInfo {
	return p + AllocatorInfo(offset*AllocatorInfoSizeOf)
}

// GetDevice returns the value in device.
func (p AllocatorInfo) GetDevice() vk.Device {
	ptr := (*C.VmaAllocatorInfo)(unsafe.Pointer(p))
	return vk.Device(unsafe.Pointer(ptr.device))
}

// SetDevice sets the value in device.
func (p AllocatorInfo) SetDevice(value vk.Device) {
	ptr := (*C.VmaAllocatorInfo)(unsafe.Pointer(p))
	ptr.device = (C.VkDevice)(unsafe.Pointer(value))
}

// GetInstance returns the value in instance.
func (p AllocatorInfo) GetInstance() vk.Instance {
	ptr := (*C.VmaAllocatorInfo)(unsafe.Pointer(p))
	return vk.Instance(unsafe.Pointer(ptr.instance))
}

// SetInstance sets the value in instance.
func (p AllocatorInfo) SetInstance(value vk.Instance) {
	ptr := (*C.VmaAllocatorInfo)(unsafe.Pointer(p))
	ptr.instance = (C.VkInstance)(unsafe.Pointer(value))
}

// GetPhysicalDevice returns the value in physicalDevice.
func (p AllocatorInfo) GetPhysicalDevice() vk.PhysicalDevice {
	ptr := (*C.VmaAllocatorInfo)(unsafe.Pointer(p))
	return vk.PhysicalDevice(unsafe.Pointer(ptr.physicalDevice))
}

// SetPhysicalDevice sets the value in physicalDevice.
func (p AllocatorInfo) SetPhysicalDevice(value vk.PhysicalDevice) {
	ptr := (*C.VmaAllocatorInfo)(unsafe.Pointer(p))
	ptr.physicalDevice = (C.VkPhysicalDevice)(unsafe.Pointer(value))
}

// Budget wraps struct VmaBudget.
/*
  \brief Statistics of current memory usage and available budget for a specific memory heap.

  These are fast to calculate.
  See function vmaGetHeapBudgets().
*/
type Budget uintptr

// BudgetNil is a null pointer.
var BudgetNil Budget

// BudgetSizeOf is the byte size of VmaBudget.
const BudgetSizeOf = int(C.sizeof_VmaBudget)

// BudgetAlloc allocates a continuous block of Budget.
func BudgetAlloc(alloc ffi.Allocator, count int) Budget {
	ptr := alloc.Allocate(BudgetSizeOf * count)
	return Budget(ptr)
}

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p Budget) Offset(offset int) Budget {
	return p + Budget(offset*BudgetSizeOf)
}

// GetBudget returns the value in budget.
func (p Budget) GetBudget() vk.DeviceSize {
	ptr := (*C.VmaBudget)(unsafe.Pointer(p))
	return vk.DeviceSize(ptr.budget)
}

// SetBudget sets the value in budget.
func (p Budget) SetBudget(value vk.DeviceSize) {
	ptr := (*C.VmaBudget)(unsafe.Pointer(p))
	ptr.budget = (C.VkDeviceSize)(value)
}

// RefStatistics returns pointer to the statistics field.
func (p Budget) RefStatistics() Statistics {
	return Statistics(p + Budget(C.offsetof_VmaBudget_statistics))
}

// GetUsage returns the value in usage.
func (p Budget) GetUsage() vk.DeviceSize {
	ptr := (*C.VmaBudget)(unsafe.Pointer(p))
	return vk.DeviceSize(ptr.usage)
}

// SetUsage sets the value in usage.
func (p Budget) SetUsage(value vk.DeviceSize) {
	ptr := (*C.VmaBudget)(unsafe.Pointer(p))
	ptr.usage = (C.VkDeviceSize)(value)
}

// DefragmentationInfo wraps struct VmaDefragmentationInfo.
/*
  \brief Parameters for defragmentation.

  To be used with function vmaBeginDefragmentation().
*/
type DefragmentationInfo uintptr

// DefragmentationInfoNil is a null pointer.
var DefragmentationInfoNil DefragmentationInfo

// DefragmentationInfoSizeOf is the byte size of VmaDefragmentationInfo.
const DefragmentationInfoSizeOf = int(C.sizeof_VmaDefragmentationInfo)

// DefragmentationInfoAlloc allocates a continuous block of DefragmentationInfo.
func DefragmentationInfoAlloc(alloc ffi.Allocator, count int) DefragmentationInfo {
	ptr := alloc.Allocate(DefragmentationInfoSizeOf * count)
	return DefragmentationInfo(ptr)
}

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p DefragmentationInfo) Offset(offset int) DefragmentationInfo {
	return p + DefragmentationInfo(offset*DefragmentationInfoSizeOf)
}

// GetFlags returns the value in flags.
func (p DefragmentationInfo) GetFlags() DefragmentationFlags {
	ptr := (*C.VmaDefragmentationInfo)(unsafe.Pointer(p))
	return DefragmentationFlags(ptr.flags)
}

// SetFlags sets the value in flags.
func (p DefragmentationInfo) SetFlags(value DefragmentationFlags) {
	ptr := (*C.VmaDefragmentationInfo)(unsafe.Pointer(p))
	ptr.flags = (C.VmaDefragmentationFlags)(value)
}

// GetMaxAllocationsPerPass returns the value in maxAllocationsPerPass.
func (p DefragmentationInfo) GetMaxAllocationsPerPass() uint32 {
	ptr := (*C.VmaDefragmentationInfo)(unsafe.Pointer(p))
	return uint32(ptr.maxAllocationsPerPass)
}

// SetMaxAllocationsPerPass sets the value in maxAllocationsPerPass.
func (p DefragmentationInfo) SetMaxAllocationsPerPass(value uint32) {
	ptr := (*C.VmaDefragmentationInfo)(unsafe.Pointer(p))
	ptr.maxAllocationsPerPass = (C.uint32_t)(value)
}

// GetMaxBytesPerPass returns the value in maxBytesPerPass.
func (p DefragmentationInfo) GetMaxBytesPerPass() vk.DeviceSize {
	ptr := (*C.VmaDefragmentationInfo)(unsafe.Pointer(p))
	return vk.DeviceSize(ptr.maxBytesPerPass)
}

// SetMaxBytesPerPass sets the value in maxBytesPerPass.
func (p DefragmentationInfo) SetMaxBytesPerPass(value vk.DeviceSize) {
	ptr := (*C.VmaDefragmentationInfo)(unsafe.Pointer(p))
	ptr.maxBytesPerPass = (C.VkDeviceSize)(value)
}

// GetPBreakCallbackUserData returns the value in pBreakCallbackUserData.
func (p DefragmentationInfo) GetPBreakCallbackUserData() uintptr {
	ptr := (*C.VmaDefragmentationInfo)(unsafe.Pointer(p))
	return uintptr(unsafe.Pointer(ptr.pBreakCallbackUserData))
}

// SetPBreakCallbackUserData sets the value in pBreakCallbackUserData.
func (p DefragmentationInfo) SetPBreakCallbackUserData(value uintptr) {
	ptr := (*C.VmaDefragmentationInfo)(unsafe.Pointer(p))
	ptr.pBreakCallbackUserData = unsafe.Pointer(value)
}

// DefragmentationInfo.pfnBreakCallback is unsupported: unknown type PFN_vmaCheckDefragmentationBreakFunction.

// GetPool returns the value in pool.
func (p DefragmentationInfo) GetPool() Pool {
	ptr := (*C.VmaDefragmentationInfo)(unsafe.Pointer(p))
	return Pool(unsafe.Pointer(ptr.pool))
}

// SetPool sets the value in pool.
func (p DefragmentationInfo) SetPool(value Pool) {
	ptr := (*C.VmaDefragmentationInfo)(unsafe.Pointer(p))
	ptr.pool = (C.VmaPool)(unsafe.Pointer(value))
}

// DefragmentationMove wraps struct VmaDefragmentationMove.
//
//	Single move of an allocation to be done for defragmentation.
type DefragmentationMove uintptr

// DefragmentationMoveNil is a null pointer.
var DefragmentationMoveNil DefragmentationMove

// DefragmentationMoveSizeOf is the byte size of VmaDefragmentationMove.
const DefragmentationMoveSizeOf = int(C.sizeof_VmaDefragmentationMove)

// DefragmentationMoveAlloc allocates a continuous block of DefragmentationMove.
func DefragmentationMoveAlloc(alloc ffi.Allocator, count int) DefragmentationMove {
	ptr := alloc.Allocate(DefragmentationMoveSizeOf * count)
	return DefragmentationMove(ptr)
}

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p DefragmentationMove) Offset(offset int) DefragmentationMove {
	return p + DefragmentationMove(offset*DefragmentationMoveSizeOf)
}

// GetDstTmpAllocation returns the value in dstTmpAllocation.
func (p DefragmentationMove) GetDstTmpAllocation() Allocation {
	ptr := (*C.VmaDefragmentationMove)(unsafe.Pointer(p))
	return Allocation(unsafe.Pointer(ptr.dstTmpAllocation))
}

// SetDstTmpAllocation sets the value in dstTmpAllocation.
func (p DefragmentationMove) SetDstTmpAllocation(value Allocation) {
	ptr := (*C.VmaDefragmentationMove)(unsafe.Pointer(p))
	ptr.dstTmpAllocation = (C.VmaAllocation)(unsafe.Pointer(value))
}

// GetOperation returns the value in operation.
func (p DefragmentationMove) GetOperation() DefragmentationMoveOperation {
	ptr := (*C.VmaDefragmentationMove)(unsafe.Pointer(p))
	return DefragmentationMoveOperation(ptr.operation)
}

// SetOperation sets the value in operation.
func (p DefragmentationMove) SetOperation(value DefragmentationMoveOperation) {
	ptr := (*C.VmaDefragmentationMove)(unsafe.Pointer(p))
	ptr.operation = (C.VmaDefragmentationMoveOperation)(value)
}

// GetSrcAllocation returns the value in srcAllocation.
func (p DefragmentationMove) GetSrcAllocation() Allocation {
	ptr := (*C.VmaDefragmentationMove)(unsafe.Pointer(p))
	return Allocation(unsafe.Pointer(ptr.srcAllocation))
}

// SetSrcAllocation sets the value in srcAllocation.
func (p DefragmentationMove) SetSrcAllocation(value Allocation) {
	ptr := (*C.VmaDefragmentationMove)(unsafe.Pointer(p))
	ptr.srcAllocation = (C.VmaAllocation)(unsafe.Pointer(value))
}

// DefragmentationPassMoveInfo wraps struct VmaDefragmentationPassMoveInfo.
/*
  \brief Parameters for incremental defragmentation steps.

  To be used with function vmaBeginDefragmentationPass().
*/
type DefragmentationPassMoveInfo uintptr

// DefragmentationPassMoveInfoNil is a null pointer.
var DefragmentationPassMoveInfoNil DefragmentationPassMoveInfo

// DefragmentationPassMoveInfoSizeOf is the byte size of VmaDefragmentationPassMoveInfo.
const DefragmentationPassMoveInfoSizeOf = int(C.sizeof_VmaDefragmentationPassMoveInfo)

// DefragmentationPassMoveInfoAlloc allocates a continuous block of DefragmentationPassMoveInfo.
func DefragmentationPassMoveInfoAlloc(alloc ffi.Allocator, count int) DefragmentationPassMoveInfo {
	ptr := alloc.Allocate(DefragmentationPassMoveInfoSizeOf * count)
	return DefragmentationPassMoveInfo(ptr)
}

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p DefragmentationPassMoveInfo) Offset(offset int) DefragmentationPassMoveInfo {
	return p + DefragmentationPassMoveInfo(offset*DefragmentationPassMoveInfoSizeOf)
}

// GetMoveCount returns the value in moveCount.
func (p DefragmentationPassMoveInfo) GetMoveCount() uint32 {
	ptr := (*C.VmaDefragmentationPassMoveInfo)(unsafe.Pointer(p))
	return uint32(ptr.moveCount)
}

// SetMoveCount sets the value in moveCount.
func (p DefragmentationPassMoveInfo) SetMoveCount(value uint32) {
	ptr := (*C.VmaDefragmentationPassMoveInfo)(unsafe.Pointer(p))
	ptr.moveCount = (C.uint32_t)(value)
}

// GetPMoves returns the value in pMoves.
func (p DefragmentationPassMoveInfo) GetPMoves() DefragmentationMove {
	ptr := (*C.VmaDefragmentationPassMoveInfo)(unsafe.Pointer(p))
	return DefragmentationMove(unsafe.Pointer(ptr.pMoves))
}

// SetPMoves sets the value in pMoves.
func (p DefragmentationPassMoveInfo) SetPMoves(value DefragmentationMove) {
	ptr := (*C.VmaDefragmentationPassMoveInfo)(unsafe.Pointer(p))
	ptr.pMoves = (*C.VmaDefragmentationMove)(unsafe.Pointer(value))
}

// DefragmentationStats wraps struct VmaDefragmentationStats.
//
//	Statistics returned for defragmentation process in function vmaEndDefragmentation().
type DefragmentationStats uintptr

// DefragmentationStatsNil is a null pointer.
var DefragmentationStatsNil DefragmentationStats

// DefragmentationStatsSizeOf is the byte size of VmaDefragmentationStats.
const DefragmentationStatsSizeOf = int(C.sizeof_VmaDefragmentationStats)

// DefragmentationStatsAlloc allocates a continuous block of DefragmentationStats.
func DefragmentationStatsAlloc(alloc ffi.Allocator, count int) DefragmentationStats {
	ptr := alloc.Allocate(DefragmentationStatsSizeOf * count)
	return DefragmentationStats(ptr)
}

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p DefragmentationStats) Offset(offset int) DefragmentationStats {
	return p + DefragmentationStats(offset*DefragmentationStatsSizeOf)
}

// GetAllocationsMoved returns the value in allocationsMoved.
func (p DefragmentationStats) GetAllocationsMoved() uint32 {
	ptr := (*C.VmaDefragmentationStats)(unsafe.Pointer(p))
	return uint32(ptr.allocationsMoved)
}

// SetAllocationsMoved sets the value in allocationsMoved.
func (p DefragmentationStats) SetAllocationsMoved(value uint32) {
	ptr := (*C.VmaDefragmentationStats)(unsafe.Pointer(p))
	ptr.allocationsMoved = (C.uint32_t)(value)
}

// GetBytesFreed returns the value in bytesFreed.
func (p DefragmentationStats) GetBytesFreed() vk.DeviceSize {
	ptr := (*C.VmaDefragmentationStats)(unsafe.Pointer(p))
	return vk.DeviceSize(ptr.bytesFreed)
}

// SetBytesFreed sets the value in bytesFreed.
func (p DefragmentationStats) SetBytesFreed(value vk.DeviceSize) {
	ptr := (*C.VmaDefragmentationStats)(unsafe.Pointer(p))
	ptr.bytesFreed = (C.VkDeviceSize)(value)
}

// GetBytesMoved returns the value in bytesMoved.
func (p DefragmentationStats) GetBytesMoved() vk.DeviceSize {
	ptr := (*C.VmaDefragmentationStats)(unsafe.Pointer(p))
	return vk.DeviceSize(ptr.bytesMoved)
}

// SetBytesMoved sets the value in bytesMoved.
func (p DefragmentationStats) SetBytesMoved(value vk.DeviceSize) {
	ptr := (*C.VmaDefragmentationStats)(unsafe.Pointer(p))
	ptr.bytesMoved = (C.VkDeviceSize)(value)
}

// GetDeviceMemoryBlocksFreed returns the value in deviceMemoryBlocksFreed.
func (p DefragmentationStats) GetDeviceMemoryBlocksFreed() uint32 {
	ptr := (*C.VmaDefragmentationStats)(unsafe.Pointer(p))
	return uint32(ptr.deviceMemoryBlocksFreed)
}

// SetDeviceMemoryBlocksFreed sets the value in deviceMemoryBlocksFreed.
func (p DefragmentationStats) SetDeviceMemoryBlocksFreed(value uint32) {
	ptr := (*C.VmaDefragmentationStats)(unsafe.Pointer(p))
	ptr.deviceMemoryBlocksFreed = (C.uint32_t)(value)
}

// DetailedStatistics wraps struct VmaDetailedStatistics.
/*
  \brief More detailed statistics than #VmaStatistics.

  These are slower to calculate. Use for debugging purposes.
  See functions: vmaCalculateStatistics(), vmaCalculatePoolStatistics().

  Previous version of the statistics API provided averages, but they have been removed
  because they can be easily calculated as:

  \code
  VkDeviceSize allocationSizeAvg = detailedStats.statistics.allocationBytes / detailedStats.statistics.allocationCount;
  VkDeviceSize unusedBytes = detailedStats.statistics.blockBytes - detailedStats.statistics.allocationBytes;
  VkDeviceSize unusedRangeSizeAvg = unusedBytes / detailedStats.unusedRangeCount;
  \endcode
*/
type DetailedStatistics uintptr

// DetailedStatisticsNil is a null pointer.
var DetailedStatisticsNil DetailedStatistics

// DetailedStatisticsSizeOf is the byte size of VmaDetailedStatistics.
const DetailedStatisticsSizeOf = int(C.sizeof_VmaDetailedStatistics)

// DetailedStatisticsAlloc allocates a continuous block of DetailedStatistics.
func DetailedStatisticsAlloc(alloc ffi.Allocator, count int) DetailedStatistics {
	ptr := alloc.Allocate(DetailedStatisticsSizeOf * count)
	return DetailedStatistics(ptr)
}

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p DetailedStatistics) Offset(offset int) DetailedStatistics {
	return p + DetailedStatistics(offset*DetailedStatisticsSizeOf)
}

// GetAllocationSizeMax returns the value in allocationSizeMax.
func (p DetailedStatistics) GetAllocationSizeMax() vk.DeviceSize {
	ptr := (*C.VmaDetailedStatistics)(unsafe.Pointer(p))
	return vk.DeviceSize(ptr.allocationSizeMax)
}

// SetAllocationSizeMax sets the value in allocationSizeMax.
func (p DetailedStatistics) SetAllocationSizeMax(value vk.DeviceSize) {
	ptr := (*C.VmaDetailedStatistics)(unsafe.Pointer(p))
	ptr.allocationSizeMax = (C.VkDeviceSize)(value)
}

// GetAllocationSizeMin returns the value in allocationSizeMin.
func (p DetailedStatistics) GetAllocationSizeMin() vk.DeviceSize {
	ptr := (*C.VmaDetailedStatistics)(unsafe.Pointer(p))
	return vk.DeviceSize(ptr.allocationSizeMin)
}

// SetAllocationSizeMin sets the value in allocationSizeMin.
func (p DetailedStatistics) SetAllocationSizeMin(value vk.DeviceSize) {
	ptr := (*C.VmaDetailedStatistics)(unsafe.Pointer(p))
	ptr.allocationSizeMin = (C.VkDeviceSize)(value)
}

// RefStatistics returns pointer to the statistics field.
func (p DetailedStatistics) RefStatistics() Statistics {
	return Statistics(p + DetailedStatistics(C.offsetof_VmaDetailedStatistics_statistics))
}

// GetUnusedRangeCount returns the value in unusedRangeCount.
func (p DetailedStatistics) GetUnusedRangeCount() uint32 {
	ptr := (*C.VmaDetailedStatistics)(unsafe.Pointer(p))
	return uint32(ptr.unusedRangeCount)
}

// SetUnusedRangeCount sets the value in unusedRangeCount.
func (p DetailedStatistics) SetUnusedRangeCount(value uint32) {
	ptr := (*C.VmaDetailedStatistics)(unsafe.Pointer(p))
	ptr.unusedRangeCount = (C.uint32_t)(value)
}

// GetUnusedRangeSizeMax returns the value in unusedRangeSizeMax.
func (p DetailedStatistics) GetUnusedRangeSizeMax() vk.DeviceSize {
	ptr := (*C.VmaDetailedStatistics)(unsafe.Pointer(p))
	return vk.DeviceSize(ptr.unusedRangeSizeMax)
}

// SetUnusedRangeSizeMax sets the value in unusedRangeSizeMax.
func (p DetailedStatistics) SetUnusedRangeSizeMax(value vk.DeviceSize) {
	ptr := (*C.VmaDetailedStatistics)(unsafe.Pointer(p))
	ptr.unusedRangeSizeMax = (C.VkDeviceSize)(value)
}

// GetUnusedRangeSizeMin returns the value in unusedRangeSizeMin.
func (p DetailedStatistics) GetUnusedRangeSizeMin() vk.DeviceSize {
	ptr := (*C.VmaDetailedStatistics)(unsafe.Pointer(p))
	return vk.DeviceSize(ptr.unusedRangeSizeMin)
}

// SetUnusedRangeSizeMin sets the value in unusedRangeSizeMin.
func (p DetailedStatistics) SetUnusedRangeSizeMin(value vk.DeviceSize) {
	ptr := (*C.VmaDetailedStatistics)(unsafe.Pointer(p))
	ptr.unusedRangeSizeMin = (C.VkDeviceSize)(value)
}

// DeviceMemoryCallbacks wraps struct VmaDeviceMemoryCallbacks.
/*
  \brief Set of callbacks that the library will call for `vkAllocateMemory` and `vkFreeMemory`.

  Provided for informative purpose, e.g. to gather statistics about number of
  allocations or total amount of memory allocated in Vulkan.

  Used in VmaAllocatorCreateInfo::pDeviceMemoryCallbacks.
*/
type DeviceMemoryCallbacks uintptr

// DeviceMemoryCallbacksNil is a null pointer.
var DeviceMemoryCallbacksNil DeviceMemoryCallbacks

// DeviceMemoryCallbacksSizeOf is the byte size of VmaDeviceMemoryCallbacks.
const DeviceMemoryCallbacksSizeOf = int(C.sizeof_VmaDeviceMemoryCallbacks)

// DeviceMemoryCallbacksAlloc allocates a continuous block of DeviceMemoryCallbacks.
func DeviceMemoryCallbacksAlloc(alloc ffi.Allocator, count int) DeviceMemoryCallbacks {
	ptr := alloc.Allocate(DeviceMemoryCallbacksSizeOf * count)
	return DeviceMemoryCallbacks(ptr)
}

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p DeviceMemoryCallbacks) Offset(offset int) DeviceMemoryCallbacks {
	return p + DeviceMemoryCallbacks(offset*DeviceMemoryCallbacksSizeOf)
}

// GetPUserData returns the value in pUserData.
func (p DeviceMemoryCallbacks) GetPUserData() uintptr {
	ptr := (*C.VmaDeviceMemoryCallbacks)(unsafe.Pointer(p))
	return uintptr(unsafe.Pointer(ptr.pUserData))
}

// SetPUserData sets the value in pUserData.
func (p DeviceMemoryCallbacks) SetPUserData(value uintptr) {
	ptr := (*C.VmaDeviceMemoryCallbacks)(unsafe.Pointer(p))
	ptr.pUserData = unsafe.Pointer(value)
}

// DeviceMemoryCallbacks.pfnAllocate is unsupported: unknown type PFN_vmaAllocateDeviceMemoryFunction.

// DeviceMemoryCallbacks.pfnFree is unsupported: unknown type PFN_vmaFreeDeviceMemoryFunction.

// PoolCreateInfo wraps struct VmaPoolCreateInfo.
//
//	Describes parameter of created #VmaPool.
type PoolCreateInfo uintptr

// PoolCreateInfoNil is a null pointer.
var PoolCreateInfoNil PoolCreateInfo

// PoolCreateInfoSizeOf is the byte size of VmaPoolCreateInfo.
const PoolCreateInfoSizeOf = int(C.sizeof_VmaPoolCreateInfo)

// PoolCreateInfoAlloc allocates a continuous block of PoolCreateInfo.
func PoolCreateInfoAlloc(alloc ffi.Allocator, count int) PoolCreateInfo {
	ptr := alloc.Allocate(PoolCreateInfoSizeOf * count)
	return PoolCreateInfo(ptr)
}

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p PoolCreateInfo) Offset(offset int) PoolCreateInfo {
	return p + PoolCreateInfo(offset*PoolCreateInfoSizeOf)
}

// GetBlockSize returns the value in blockSize.
func (p PoolCreateInfo) GetBlockSize() vk.DeviceSize {
	ptr := (*C.VmaPoolCreateInfo)(unsafe.Pointer(p))
	return vk.DeviceSize(ptr.blockSize)
}

// SetBlockSize sets the value in blockSize.
func (p PoolCreateInfo) SetBlockSize(value vk.DeviceSize) {
	ptr := (*C.VmaPoolCreateInfo)(unsafe.Pointer(p))
	ptr.blockSize = (C.VkDeviceSize)(value)
}

// GetFlags returns the value in flags.
func (p PoolCreateInfo) GetFlags() PoolCreateFlags {
	ptr := (*C.VmaPoolCreateInfo)(unsafe.Pointer(p))
	return PoolCreateFlags(ptr.flags)
}

// SetFlags sets the value in flags.
func (p PoolCreateInfo) SetFlags(value PoolCreateFlags) {
	ptr := (*C.VmaPoolCreateInfo)(unsafe.Pointer(p))
	ptr.flags = (C.VmaPoolCreateFlags)(value)
}

// GetMaxBlockCount returns the value in maxBlockCount.
func (p PoolCreateInfo) GetMaxBlockCount() uintptr {
	ptr := (*C.VmaPoolCreateInfo)(unsafe.Pointer(p))
	return uintptr(ptr.maxBlockCount)
}

// SetMaxBlockCount sets the value in maxBlockCount.
func (p PoolCreateInfo) SetMaxBlockCount(value uintptr) {
	ptr := (*C.VmaPoolCreateInfo)(unsafe.Pointer(p))
	ptr.maxBlockCount = (C.size_t)(value)
}

// GetMemoryTypeIndex returns the value in memoryTypeIndex.
func (p PoolCreateInfo) GetMemoryTypeIndex() uint32 {
	ptr := (*C.VmaPoolCreateInfo)(unsafe.Pointer(p))
	return uint32(ptr.memoryTypeIndex)
}

// SetMemoryTypeIndex sets the value in memoryTypeIndex.
func (p PoolCreateInfo) SetMemoryTypeIndex(value uint32) {
	ptr := (*C.VmaPoolCreateInfo)(unsafe.Pointer(p))
	ptr.memoryTypeIndex = (C.uint32_t)(value)
}

// GetMinAllocationAlignment returns the value in minAllocationAlignment.
func (p PoolCreateInfo) GetMinAllocationAlignment() vk.DeviceSize {
	ptr := (*C.VmaPoolCreateInfo)(unsafe.Pointer(p))
	return vk.DeviceSize(ptr.minAllocationAlignment)
}

// SetMinAllocationAlignment sets the value in minAllocationAlignment.
func (p PoolCreateInfo) SetMinAllocationAlignment(value vk.DeviceSize) {
	ptr := (*C.VmaPoolCreateInfo)(unsafe.Pointer(p))
	ptr.minAllocationAlignment = (C.VkDeviceSize)(value)
}

// GetMinBlockCount returns the value in minBlockCount.
func (p PoolCreateInfo) GetMinBlockCount() uintptr {
	ptr := (*C.VmaPoolCreateInfo)(unsafe.Pointer(p))
	return uintptr(ptr.minBlockCount)
}

// SetMinBlockCount sets the value in minBlockCount.
func (p PoolCreateInfo) SetMinBlockCount(value uintptr) {
	ptr := (*C.VmaPoolCreateInfo)(unsafe.Pointer(p))
	ptr.minBlockCount = (C.size_t)(value)
}

// GetPMemoryAllocateNext returns the value in pMemoryAllocateNext.
func (p PoolCreateInfo) GetPMemoryAllocateNext() uintptr {
	ptr := (*C.VmaPoolCreateInfo)(unsafe.Pointer(p))
	return uintptr(unsafe.Pointer(ptr.pMemoryAllocateNext))
}

// SetPMemoryAllocateNext sets the value in pMemoryAllocateNext.
func (p PoolCreateInfo) SetPMemoryAllocateNext(value uintptr) {
	ptr := (*C.VmaPoolCreateInfo)(unsafe.Pointer(p))
	ptr.pMemoryAllocateNext = unsafe.Pointer(value)
}

// GetPriority returns the value in priority.
func (p PoolCreateInfo) GetPriority() float32 {
	ptr := (*C.VmaPoolCreateInfo)(unsafe.Pointer(p))
	return float32(ptr.priority)
}

// SetPriority sets the value in priority.
func (p PoolCreateInfo) SetPriority(value float32) {
	ptr := (*C.VmaPoolCreateInfo)(unsafe.Pointer(p))
	ptr.priority = (C.float)(value)
}

// Statistics wraps struct VmaStatistics.
/*
  \brief Calculated statistics of memory usage e.g. in a specific memory type, heap, custom pool, or total.

  These are fast to calculate.
  See functions: vmaGetHeapBudgets(), vmaGetPoolStatistics().
*/
type Statistics uintptr

// StatisticsNil is a null pointer.
var StatisticsNil Statistics

// StatisticsSizeOf is the byte size of VmaStatistics.
const StatisticsSizeOf = int(C.sizeof_VmaStatistics)

// StatisticsAlloc allocates a continuous block of Statistics.
func StatisticsAlloc(alloc ffi.Allocator, count int) Statistics {
	ptr := alloc.Allocate(StatisticsSizeOf * count)
	return Statistics(ptr)
}

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p Statistics) Offset(offset int) Statistics {
	return p + Statistics(offset*StatisticsSizeOf)
}

// GetAllocationBytes returns the value in allocationBytes.
func (p Statistics) GetAllocationBytes() vk.DeviceSize {
	ptr := (*C.VmaStatistics)(unsafe.Pointer(p))
	return vk.DeviceSize(ptr.allocationBytes)
}

// SetAllocationBytes sets the value in allocationBytes.
func (p Statistics) SetAllocationBytes(value vk.DeviceSize) {
	ptr := (*C.VmaStatistics)(unsafe.Pointer(p))
	ptr.allocationBytes = (C.VkDeviceSize)(value)
}

// GetAllocationCount returns the value in allocationCount.
func (p Statistics) GetAllocationCount() uint32 {
	ptr := (*C.VmaStatistics)(unsafe.Pointer(p))
	return uint32(ptr.allocationCount)
}

// SetAllocationCount sets the value in allocationCount.
func (p Statistics) SetAllocationCount(value uint32) {
	ptr := (*C.VmaStatistics)(unsafe.Pointer(p))
	ptr.allocationCount = (C.uint32_t)(value)
}

// GetBlockBytes returns the value in blockBytes.
func (p Statistics) GetBlockBytes() vk.DeviceSize {
	ptr := (*C.VmaStatistics)(unsafe.Pointer(p))
	return vk.DeviceSize(ptr.blockBytes)
}

// SetBlockBytes sets the value in blockBytes.
func (p Statistics) SetBlockBytes(value vk.DeviceSize) {
	ptr := (*C.VmaStatistics)(unsafe.Pointer(p))
	ptr.blockBytes = (C.VkDeviceSize)(value)
}

// GetBlockCount returns the value in blockCount.
func (p Statistics) GetBlockCount() uint32 {
	ptr := (*C.VmaStatistics)(unsafe.Pointer(p))
	return uint32(ptr.blockCount)
}

// SetBlockCount sets the value in blockCount.
func (p Statistics) SetBlockCount(value uint32) {
	ptr := (*C.VmaStatistics)(unsafe.Pointer(p))
	ptr.blockCount = (C.uint32_t)(value)
}

// TotalStatistics wraps struct VmaTotalStatistics.
/*
  \brief  General statistics from current state of the Allocator -
  total memory usage across all memory heaps and types.

  These are slower to calculate. Use for debugging purposes.
  See function vmaCalculateStatistics().
*/
type TotalStatistics uintptr

// TotalStatisticsNil is a null pointer.
var TotalStatisticsNil TotalStatistics

// TotalStatisticsSizeOf is the byte size of VmaTotalStatistics.
const TotalStatisticsSizeOf = int(C.sizeof_VmaTotalStatistics)

// TotalStatisticsAlloc allocates a continuous block of TotalStatistics.
func TotalStatisticsAlloc(alloc ffi.Allocator, count int) TotalStatistics {
	ptr := alloc.Allocate(TotalStatisticsSizeOf * count)
	return TotalStatistics(ptr)
}

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p TotalStatistics) Offset(offset int) TotalStatistics {
	return p + TotalStatistics(offset*TotalStatisticsSizeOf)
}

// TotalStatistics.memoryHeap is unsupported: category array.

// TotalStatistics.memoryType is unsupported: category array.

// RefTotal returns pointer to the total field.
func (p TotalStatistics) RefTotal() DetailedStatistics {
	return DetailedStatistics(p + TotalStatistics(C.offsetof_VmaTotalStatistics_total))
}

// VirtualAllocationCreateInfo wraps struct VmaVirtualAllocationCreateInfo.
//
//	Parameters of created virtual allocation to be passed to vmaVirtualAllocate().
type VirtualAllocationCreateInfo uintptr

// VirtualAllocationCreateInfoNil is a null pointer.
var VirtualAllocationCreateInfoNil VirtualAllocationCreateInfo

// VirtualAllocationCreateInfoSizeOf is the byte size of VmaVirtualAllocationCreateInfo.
const VirtualAllocationCreateInfoSizeOf = int(C.sizeof_VmaVirtualAllocationCreateInfo)

// VirtualAllocationCreateInfoAlloc allocates a continuous block of VirtualAllocationCreateInfo.
func VirtualAllocationCreateInfoAlloc(alloc ffi.Allocator, count int) VirtualAllocationCreateInfo {
	ptr := alloc.Allocate(VirtualAllocationCreateInfoSizeOf * count)
	return VirtualAllocationCreateInfo(ptr)
}

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p VirtualAllocationCreateInfo) Offset(offset int) VirtualAllocationCreateInfo {
	return p + VirtualAllocationCreateInfo(offset*VirtualAllocationCreateInfoSizeOf)
}

// GetAlignment returns the value in alignment.
func (p VirtualAllocationCreateInfo) GetAlignment() vk.DeviceSize {
	ptr := (*C.VmaVirtualAllocationCreateInfo)(unsafe.Pointer(p))
	return vk.DeviceSize(ptr.alignment)
}

// SetAlignment sets the value in alignment.
func (p VirtualAllocationCreateInfo) SetAlignment(value vk.DeviceSize) {
	ptr := (*C.VmaVirtualAllocationCreateInfo)(unsafe.Pointer(p))
	ptr.alignment = (C.VkDeviceSize)(value)
}

// GetFlags returns the value in flags.
func (p VirtualAllocationCreateInfo) GetFlags() VirtualAllocationCreateFlags {
	ptr := (*C.VmaVirtualAllocationCreateInfo)(unsafe.Pointer(p))
	return VirtualAllocationCreateFlags(ptr.flags)
}

// SetFlags sets the value in flags.
func (p VirtualAllocationCreateInfo) SetFlags(value VirtualAllocationCreateFlags) {
	ptr := (*C.VmaVirtualAllocationCreateInfo)(unsafe.Pointer(p))
	ptr.flags = (C.VmaVirtualAllocationCreateFlags)(value)
}

// GetPUserData returns the value in pUserData.
func (p VirtualAllocationCreateInfo) GetPUserData() uintptr {
	ptr := (*C.VmaVirtualAllocationCreateInfo)(unsafe.Pointer(p))
	return uintptr(unsafe.Pointer(ptr.pUserData))
}

// SetPUserData sets the value in pUserData.
func (p VirtualAllocationCreateInfo) SetPUserData(value uintptr) {
	ptr := (*C.VmaVirtualAllocationCreateInfo)(unsafe.Pointer(p))
	ptr.pUserData = unsafe.Pointer(value)
}

// GetSize returns the value in size.
func (p VirtualAllocationCreateInfo) GetSize() vk.DeviceSize {
	ptr := (*C.VmaVirtualAllocationCreateInfo)(unsafe.Pointer(p))
	return vk.DeviceSize(ptr.size)
}

// SetSize sets the value in size.
func (p VirtualAllocationCreateInfo) SetSize(value vk.DeviceSize) {
	ptr := (*C.VmaVirtualAllocationCreateInfo)(unsafe.Pointer(p))
	ptr.size = (C.VkDeviceSize)(value)
}

// VirtualAllocationInfo wraps struct VmaVirtualAllocationInfo.
//
//	Parameters of an existing virtual allocation, returned by vmaGetVirtualAllocationInfo().
type VirtualAllocationInfo uintptr

// VirtualAllocationInfoNil is a null pointer.
var VirtualAllocationInfoNil VirtualAllocationInfo

// VirtualAllocationInfoSizeOf is the byte size of VmaVirtualAllocationInfo.
const VirtualAllocationInfoSizeOf = int(C.sizeof_VmaVirtualAllocationInfo)

// VirtualAllocationInfoAlloc allocates a continuous block of VirtualAllocationInfo.
func VirtualAllocationInfoAlloc(alloc ffi.Allocator, count int) VirtualAllocationInfo {
	ptr := alloc.Allocate(VirtualAllocationInfoSizeOf * count)
	return VirtualAllocationInfo(ptr)
}

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p VirtualAllocationInfo) Offset(offset int) VirtualAllocationInfo {
	return p + VirtualAllocationInfo(offset*VirtualAllocationInfoSizeOf)
}

// GetOffset returns the value in offset.
func (p VirtualAllocationInfo) GetOffset() vk.DeviceSize {
	ptr := (*C.VmaVirtualAllocationInfo)(unsafe.Pointer(p))
	return vk.DeviceSize(ptr.offset)
}

// SetOffset sets the value in offset.
func (p VirtualAllocationInfo) SetOffset(value vk.DeviceSize) {
	ptr := (*C.VmaVirtualAllocationInfo)(unsafe.Pointer(p))
	ptr.offset = (C.VkDeviceSize)(value)
}

// GetPUserData returns the value in pUserData.
func (p VirtualAllocationInfo) GetPUserData() uintptr {
	ptr := (*C.VmaVirtualAllocationInfo)(unsafe.Pointer(p))
	return uintptr(unsafe.Pointer(ptr.pUserData))
}

// SetPUserData sets the value in pUserData.
func (p VirtualAllocationInfo) SetPUserData(value uintptr) {
	ptr := (*C.VmaVirtualAllocationInfo)(unsafe.Pointer(p))
	ptr.pUserData = unsafe.Pointer(value)
}

// GetSize returns the value in size.
func (p VirtualAllocationInfo) GetSize() vk.DeviceSize {
	ptr := (*C.VmaVirtualAllocationInfo)(unsafe.Pointer(p))
	return vk.DeviceSize(ptr.size)
}

// SetSize sets the value in size.
func (p VirtualAllocationInfo) SetSize(value vk.DeviceSize) {
	ptr := (*C.VmaVirtualAllocationInfo)(unsafe.Pointer(p))
	ptr.size = (C.VkDeviceSize)(value)
}

// VirtualBlockCreateInfo wraps struct VmaVirtualBlockCreateInfo.
//
//	Parameters of created #VmaVirtualBlock object to be passed to vmaCreateVirtualBlock().
type VirtualBlockCreateInfo uintptr

// VirtualBlockCreateInfoNil is a null pointer.
var VirtualBlockCreateInfoNil VirtualBlockCreateInfo

// VirtualBlockCreateInfoSizeOf is the byte size of VmaVirtualBlockCreateInfo.
const VirtualBlockCreateInfoSizeOf = int(C.sizeof_VmaVirtualBlockCreateInfo)

// VirtualBlockCreateInfoAlloc allocates a continuous block of VirtualBlockCreateInfo.
func VirtualBlockCreateInfoAlloc(alloc ffi.Allocator, count int) VirtualBlockCreateInfo {
	ptr := alloc.Allocate(VirtualBlockCreateInfoSizeOf * count)
	return VirtualBlockCreateInfo(ptr)
}

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p VirtualBlockCreateInfo) Offset(offset int) VirtualBlockCreateInfo {
	return p + VirtualBlockCreateInfo(offset*VirtualBlockCreateInfoSizeOf)
}

// GetFlags returns the value in flags.
func (p VirtualBlockCreateInfo) GetFlags() VirtualBlockCreateFlags {
	ptr := (*C.VmaVirtualBlockCreateInfo)(unsafe.Pointer(p))
	return VirtualBlockCreateFlags(ptr.flags)
}

// SetFlags sets the value in flags.
func (p VirtualBlockCreateInfo) SetFlags(value VirtualBlockCreateFlags) {
	ptr := (*C.VmaVirtualBlockCreateInfo)(unsafe.Pointer(p))
	ptr.flags = (C.VmaVirtualBlockCreateFlags)(value)
}

// GetPAllocationCallbacks returns the value in pAllocationCallbacks.
func (p VirtualBlockCreateInfo) GetPAllocationCallbacks() vk.AllocationCallbacks {
	ptr := (*C.VmaVirtualBlockCreateInfo)(unsafe.Pointer(p))
	return vk.AllocationCallbacks(unsafe.Pointer(ptr.pAllocationCallbacks))
}

// SetPAllocationCallbacks sets the value in pAllocationCallbacks.
func (p VirtualBlockCreateInfo) SetPAllocationCallbacks(value vk.AllocationCallbacks) {
	ptr := (*C.VmaVirtualBlockCreateInfo)(unsafe.Pointer(p))
	ptr.pAllocationCallbacks = (*C.VkAllocationCallbacks)(unsafe.Pointer(value))
}

// GetSize returns the value in size.
func (p VirtualBlockCreateInfo) GetSize() vk.DeviceSize {
	ptr := (*C.VmaVirtualBlockCreateInfo)(unsafe.Pointer(p))
	return vk.DeviceSize(ptr.size)
}

// SetSize sets the value in size.
func (p VirtualBlockCreateInfo) SetSize(value vk.DeviceSize) {
	ptr := (*C.VmaVirtualBlockCreateInfo)(unsafe.Pointer(p))
	ptr.size = (C.VkDeviceSize)(value)
}

// VulkanFunctions wraps struct VmaVulkanFunctions.
/*
  \brief Pointers to some Vulkan functions - a subset used by the library.

  Used in VmaAllocatorCreateInfo::pVulkanFunctions.
*/
type VulkanFunctions uintptr

// VulkanFunctionsNil is a null pointer.
var VulkanFunctionsNil VulkanFunctions

// VulkanFunctionsSizeOf is the byte size of VmaVulkanFunctions.
const VulkanFunctionsSizeOf = int(C.sizeof_VmaVulkanFunctions)

// VulkanFunctionsAlloc allocates a continuous block of VulkanFunctions.
func VulkanFunctionsAlloc(alloc ffi.Allocator, count int) VulkanFunctions {
	ptr := alloc.Allocate(VulkanFunctionsSizeOf * count)
	return VulkanFunctions(ptr)
}

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p VulkanFunctions) Offset(offset int) VulkanFunctions {
	return p + VulkanFunctions(offset*VulkanFunctionsSizeOf)
}

// VulkanFunctions.vkAllocateMemory is unsupported: unknown type PFN_vkAllocateMemory.

// VulkanFunctions.vkBindBufferMemory is unsupported: unknown type PFN_vkBindBufferMemory.

// VulkanFunctions.vkBindBufferMemory2KHR is unsupported: unknown type PFN_vkBindBufferMemory2KHR.

// VulkanFunctions.vkBindImageMemory is unsupported: unknown type PFN_vkBindImageMemory.

// VulkanFunctions.vkBindImageMemory2KHR is unsupported: unknown type PFN_vkBindImageMemory2KHR.

// VulkanFunctions.vkCmdCopyBuffer is unsupported: unknown type PFN_vkCmdCopyBuffer.

// VulkanFunctions.vkCreateBuffer is unsupported: unknown type PFN_vkCreateBuffer.

// VulkanFunctions.vkCreateImage is unsupported: unknown type PFN_vkCreateImage.

// VulkanFunctions.vkDestroyBuffer is unsupported: unknown type PFN_vkDestroyBuffer.

// VulkanFunctions.vkDestroyImage is unsupported: unknown type PFN_vkDestroyImage.

// VulkanFunctions.vkFlushMappedMemoryRanges is unsupported: unknown type PFN_vkFlushMappedMemoryRanges.

// VulkanFunctions.vkFreeMemory is unsupported: unknown type PFN_vkFreeMemory.

// VulkanFunctions.vkGetBufferMemoryRequirements is unsupported: unknown type PFN_vkGetBufferMemoryRequirements.

// VulkanFunctions.vkGetBufferMemoryRequirements2KHR is unsupported: unknown type PFN_vkGetBufferMemoryRequirements2KHR.

// VulkanFunctions.vkGetDeviceBufferMemoryRequirements is unsupported: unknown type PFN_vkGetDeviceBufferMemoryRequirementsKHR.

// VulkanFunctions.vkGetDeviceImageMemoryRequirements is unsupported: unknown type PFN_vkGetDeviceImageMemoryRequirementsKHR.

// VulkanFunctions.vkGetDeviceProcAddr is unsupported: unknown type PFN_vkGetDeviceProcAddr.

// VulkanFunctions.vkGetImageMemoryRequirements is unsupported: unknown type PFN_vkGetImageMemoryRequirements.

// VulkanFunctions.vkGetImageMemoryRequirements2KHR is unsupported: unknown type PFN_vkGetImageMemoryRequirements2KHR.

// VulkanFunctions.vkGetInstanceProcAddr is unsupported: unknown type PFN_vkGetInstanceProcAddr.

// GetVkGetMemoryWin32HandleKHR returns the value in vkGetMemoryWin32HandleKHR.
func (p VulkanFunctions) GetVkGetMemoryWin32HandleKHR() uintptr {
	ptr := (*C.VmaVulkanFunctions)(unsafe.Pointer(p))
	return uintptr(unsafe.Pointer(ptr.vkGetMemoryWin32HandleKHR))
}

// SetVkGetMemoryWin32HandleKHR sets the value in vkGetMemoryWin32HandleKHR.
func (p VulkanFunctions) SetVkGetMemoryWin32HandleKHR(value uintptr) {
	ptr := (*C.VmaVulkanFunctions)(unsafe.Pointer(p))
	ptr.vkGetMemoryWin32HandleKHR = unsafe.Pointer(value)
}

// VulkanFunctions.vkGetPhysicalDeviceMemoryProperties is unsupported: unknown type PFN_vkGetPhysicalDeviceMemoryProperties.

// VulkanFunctions.vkGetPhysicalDeviceMemoryProperties2KHR is unsupported: unknown type PFN_vkGetPhysicalDeviceMemoryProperties2KHR.

// VulkanFunctions.vkGetPhysicalDeviceProperties is unsupported: unknown type PFN_vkGetPhysicalDeviceProperties.

// VulkanFunctions.vkInvalidateMappedMemoryRanges is unsupported: unknown type PFN_vkInvalidateMappedMemoryRanges.

// VulkanFunctions.vkMapMemory is unsupported: unknown type PFN_vkMapMemory.

// VulkanFunctions.vkUnmapMemory is unsupported: unknown type PFN_vkUnmapMemory.
