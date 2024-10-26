package vulkan

/*
#include <vulkan/vulkan.h>
#include "vulkan.h"
*/
import "C"
import (
	"log/slog"
	"unsafe"
)

//export gfx_vk_debug_callback
func gfx_vk_debug_callback(
	messageSeverity C.VkDebugUtilsMessageSeverityFlagBitsEXT,
	messageTypes C.VkDebugUtilsMessageTypeFlagsEXT,
	pCallbackData *C.VkDebugUtilsMessengerCallbackDataEXT,
	pUserData unsafe.Pointer,
) C.VkBool32 {
	if messageSeverity&C.VK_DEBUG_UTILS_MESSAGE_SEVERITY_ERROR_BIT_EXT != 0 {
		slog.Error("vulkan", "id", pCallbackData.messageIdNumber, "name", C.GoString(pCallbackData.pMessageIdName), "msg", C.GoString(pCallbackData.pMessage))
	} else if messageSeverity&C.VK_DEBUG_UTILS_MESSAGE_SEVERITY_WARNING_BIT_EXT != 0 {
		slog.Warn("vulkan", "id", pCallbackData.messageIdNumber, "name", C.GoString(pCallbackData.pMessageIdName), "msg", C.GoString(pCallbackData.pMessage))
	} else if messageSeverity&C.VK_DEBUG_UTILS_MESSAGE_SEVERITY_INFO_BIT_EXT != 0 {
		slog.Info("vulkan", "id", pCallbackData.messageIdNumber, "name", C.GoString(pCallbackData.pMessageIdName), "msg", C.GoString(pCallbackData.pMessage))
	}

	return 0
}
