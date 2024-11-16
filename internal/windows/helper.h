#ifndef GFX_WINDOWS_H
#define GFX_WINDOWS_H

#undef WINVER
#undef _WIN32_WINNT

// Target Windows 10 upwards
#define WINVER 0x0A00
#define _WIN32_WINNT 0x0A00

#include <windows.h>

#define GFX_MODULE_ERROR (-1000)
#define GFX_SEE_ERROR (-10)
#define GFX_SUCCESS 1

int gfx_windows_init();

void gfx_windows_init_callback();

#endif //GFX_WINDOWS_H
