package gfx

import (
	"unsafe"
)

type VulkanPlatform interface {
	VKGetInstanceProcAddr() unsafe.Pointer

	RequiredVKExtensions() []string
}
