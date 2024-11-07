#ifndef VULKAN_H_
#define VULKAN_H_

#include <stdlib.h>
#include <stdint.h>

#define VK_NO_PROTOTYPES 1
#define VMA_NULLABLE
#define VMA_NOT_NULL
#define VMA_STATIC_VULKAN_FUNCTIONS 1
#define VMA_DYNAMIC_VULKAN_FUNCTIONS 0

#include "vulkan/vk_platform.h"
#include "vulkan/vulkan_core.h"
#include "vulkan/vulkan_metal.h"
#include "volk.h"
#include "vk_mem_alloc.h"
#include "map.h"

uint32_t gfx_vk_version(int a, int b, int c);

#define GFX_PARAMS(type,var) type var
#define GFX_PARAMS_PAIR(pair) GFX_PARAMS pair
#define GFX_PARAM_NAMES(type,var) var
#define GFX_PARAM_NAMES_PAIR(pair) GFX_PARAM_NAMES pair

#define GFX_FUNC(RES, NAME, ...) RES (gfx_ ## NAME) ( MAP_LIST(GFX_PARAMS_PAIR, __VA_ARGS__) );

#include "vulkan_funcs.h"

#endif //VULKAN_H_
