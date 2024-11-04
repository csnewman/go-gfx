package vulkan

/*
#include <stdlib.h>
#include <vulkan/vulkan.h>
*/
import "C"

type Texture struct {
	img C.VkImage
}

type TextureView struct {
	view C.VkImageView
}
