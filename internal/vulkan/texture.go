package vulkan

/*
#include "vulkan.h"
*/
import "C"

type Texture struct {
	img C.VkImage
}

type TextureView struct {
	view C.VkImageView
}
