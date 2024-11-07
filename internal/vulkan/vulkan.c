#include "vulkan.h"
#include "volk.c"

uint32_t gfx_vk_version(int a, int b, int c) {
    return VK_MAKE_VERSION(a, b, c);
}

#undef GFX_FUNC
#define GFX_FUNC(RES, NAME, ...) RES gfx_ ## NAME ( MAP_LIST(GFX_PARAMS_PAIR, __VA_ARGS__) ) {  return NAME (  MAP_LIST(GFX_PARAM_NAMES_PAIR, __VA_ARGS__)   ); }

#include "vulkan_funcs.h"
