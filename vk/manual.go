package vk

import (
	"unsafe"
)

/*
#include "vulkan.h"

VkClearValue gfx_vk_clear_color(float r, float g, float b, float a) {
	VkClearValue col;
	col.color.float32[0] = r;
	col.color.float32[1] = g;
	col.color.float32[2] = b;
	col.color.float32[3] = a;
	return col;
}

VkBool32 gfx_vk_debug_callback(
	VkDebugUtilsMessageSeverityFlagBitsEXT           messageSeverity,
    VkDebugUtilsMessageTypeFlagsEXT                  messageTypes,
    const VkDebugUtilsMessengerCallbackDataEXT*      pCallbackData,
    void*                                            pUserData
);
*/
import "C"

// SampleMask wraps vkSampleMask.
type SampleMask uint32

// DeviceAddress wraps vkDeviceAddress.
type DeviceAddress uint64

// DeviceSize wraps vkDeviceSize.
type DeviceSize uint64

// MakeAPIVersion produces a vulkan version identifier.
func MakeAPIVersion(variant int, major int, minor int, patch int) uint32 {
	return ((uint32(variant)) << 29) | ((uint32(major)) << 22) | ((uint32(minor)) << 12) | uint32(patch)
}

// APIVersion_1_4 represents vulkan 1.4.
const APIVersion_1_4 = ((uint32(1)) << 22) | ((uint32(4)) << 12)

const APIVersion_1_3 = ((uint32(1)) << 22) | ((uint32(3)) << 12)

func (p ExtensionProperties) GetExtensionName() [256]byte {
	ptr := (*C.VkExtensionProperties)(unsafe.Pointer(p))

	return *(*[256]byte)(unsafe.Pointer(&ptr.extensionName))
}

func (p ExtensionProperties) GetExtensionNameString() string {
	ptr := (*C.VkExtensionProperties)(unsafe.Pointer(p))

	return C.GoString(&ptr.extensionName[0])
}

func (p PhysicalDeviceProperties) GetDeviceNameString() string {
	ptr := (*C.VkPhysicalDeviceProperties)(unsafe.Pointer(p))

	return C.GoString(&ptr.deviceName[0])
}

func (p DebugUtilsMessengerCreateInfoEXT) TempSetCallback() {
	ptr := (*C.VkDebugUtilsMessengerCreateInfoEXT)(unsafe.Pointer(p))

	ptr.pfnUserCallback = C.PFN_vkDebugUtilsMessengerCallbackEXT(C.gfx_vk_debug_callback)
}

func (p RenderingAttachmentInfo) TempSetClearValue(r float32, g float32, b float32, a float32) {
	ptr := (*C.VkRenderingAttachmentInfo)(unsafe.Pointer(p))

	ptr.clearValue = C.gfx_vk_clear_color(
		C.float(r),
		C.float(g),
		C.float(b),
		C.float(a),
	)
}
