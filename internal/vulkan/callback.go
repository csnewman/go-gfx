package vulkan

/*
#include <vulkan/vulkan.h>
#include "vulkan.h"
*/
import "C"
import (
	"log"
	"unsafe"
)

//export gfx_vk_debug_callback
func gfx_vk_debug_callback(
	messageSeverity C.VkDebugUtilsMessageSeverityFlagBitsEXT,
	messageTypes C.VkDebugUtilsMessageTypeFlagsEXT,
	pCallbackData *C.VkDebugUtilsMessengerCallbackDataEXT,
	pUserData unsafe.Pointer,
) C.VkBool32 {
	log.Println("gfx_vk_debug_callback")

	return 0
}
