#ifndef GO_GFX_VULKAN_H
#define GO_GFX_VULKAN_H

#include <stdint.h>

uint32_t gfx_vk_version(int a, int b, int c);

#define GFX_VK_EXT_FUNC2(ID, ...) PFN_ ## ID func = (PFN_ ## ID)(vkGetInstanceProcAddr(instance, #ID)); func(__VA_ARGS__)


#define GFX_VK_EXT_FUNC(ID, ...) ((PFN_ ## ID)(vkGetInstanceProcAddr(instance, #ID)))(__VA_ARGS__)

#endif //GO_GFX_VULKAN_H
