#ifndef VULKAN_H_
#define VULKAN_H_

#include <stdlib.h>
#include <stdint.h>

#define VK_NO_PROTOTYPES 1
#define VK_USE_64_BIT_PTR_DEFINES 0

#if defined(_WIN32)
#include <windows.h>
#else
typedef void *HANDLE;
#define DECLARE_HANDLE(name) struct name##__ { int unused; }; typedef struct name##__ *name
typedef unsigned short WCHAR;
typedef WCHAR *LPCWSTR;
DECLARE_HANDLE(HWND);
DECLARE_HANDLE(HINSTANCE);
DECLARE_HANDLE(HMONITOR);
typedef void SECURITY_ATTRIBUTES;
typedef uint32_t DWORD;
#endif

#include "vulkan/vk_platform.h"
#include "vulkan/vulkan_core.h"
#include "vulkan/vulkan_metal.h"
#include "vulkan/vulkan_win32.h"

#endif //VULKAN_H_
