#include "vulkan.h"

uint32_t gfx_vk_version(int a, int b, int c) {
    return VK_MAKE_VERSION(a, b, c);
}

#undef GFX_FUNC
#define GFX_FUNC(RES, NAME, ...) PFN_ ## NAME ptr_ ## NAME;

#include "vulkan_funcs.h"

#undef GFX_FUNC
#define GFX_FUNC(RES, NAME, ...) RES NAME ( MAP_LIST(GFX_PARAMS_PAIR, __VA_ARGS__) ) {  return ptr_ ## NAME (  MAP_LIST(GFX_PARAM_NAMES_PAIR, __VA_ARGS__)   ); }

#include "vulkan_funcs.h"

#undef GFX_FUNC
#define GFX_FUNC(RES, NAME, ...) if (ptr_ ## NAME == NULL) ptr_ ## NAME = (PFN_ ## NAME) ptr_vkGetInstanceProcAddr(context, # NAME);

PFN_vkGetInstanceProcAddr ptr_vkGetInstanceProcAddr;

void gfx_vkInit(void* loader) {
	ptr_vkGetInstanceProcAddr = (PFN_vkGetInstanceProcAddr) loader;

	void* context = NULL;

    #include "vulkan_funcs.h"
}

PFN_vkVoidFunction gfx_vkGetInstanceProcAddr(VkInstance instance, const char* pName) {
    return ptr_vkGetInstanceProcAddr(instance, pName);
}

void gfx_vkInitInstance(VkInstance context) {
    #include "vulkan_funcs.h"
}
