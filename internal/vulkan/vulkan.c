#include "vulkan.h"
#include <vulkan/vulkan.h>
#include <vulkan/vulkan_metal.h>

uint32_t gfx_vk_version(int a, int b, int c) {
    return VK_MAKE_VERSION(a, b, c);
}
