#ifndef VULKAN_H_
#define VULKAN_H_

#include <stdlib.h>
#include <stdint.h>

#define VK_NO_PROTOTYPES 1
#define VMA_NULLABLE
#define VMA_NOT_NULL
#define VMA_STATIC_VULKAN_FUNCTIONS 1
#define VMA_DYNAMIC_VULKAN_FUNCTIONS 0
#define VMA_EXTERNAL_MEMORY_WIN32 0

//#if defined(_WIN32)
//#include <windows.h>
//#else
//typedef void *HANDLE;
//#define DECLARE_HANDLE(name) struct name##__ { int unused; }; typedef struct name##__ *name
//typedef unsigned short WCHAR;
//typedef WCHAR *LPCWSTR;
//DECLARE_HANDLE(HWND);
//DECLARE_HANDLE(HINSTANCE);
//DECLARE_HANDLE(HMONITOR);
//typedef void SECURITY_ATTRIBUTES;
//typedef uint32_t DWORD;
//#endif

#include "vulkan/vk_platform.h"
#include "vulkan/vulkan_core.h"
//#include "vulkan/vulkan_metal.h"
//#include "vulkan/vulkan_win32.h"

#include "macromap.h"

#ifdef __cplusplus
extern "C" {
#endif

uint32_t gfx_vk_version(int a, int b, int c);

void gfx_vkInit(void* loader);
void gfx_vkInitInstance(VkInstance context);

#define GFX_PARAMS(type, var) type var
#define GFX_PARAMS_PAIR(pair) GFX_PARAMS pair
#define GFX_PARAM_NAMES(type, var) var
#define GFX_PARAM_NAMES_PAIR(pair) GFX_PARAM_NAMES pair

#define GFX_FUNC(RES, NAME, ...) RES NAME ( MAP_LIST(GFX_PARAMS_PAIR, __VA_ARGS__) );

#include "vulkan_funcs.h"

PFN_vkVoidFunction gfx_vkGetInstanceProcAddr(VkInstance instance, const char* pName);

#define vkGetInstanceProcAddr gfx_vkGetInstanceProcAddr

#ifdef __cplusplus
}
#endif

#include "vk_mem_alloc.h"

#endif //VULKAN_H_
