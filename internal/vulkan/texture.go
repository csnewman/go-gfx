package vulkan

/*
#include <stdlib.h>
#include <vulkan/vulkan.h>
*/
import "C"

type TextureView struct {
	view C.VkImageView
}
