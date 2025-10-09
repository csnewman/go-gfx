package vk

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
