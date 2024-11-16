#ifndef GFX_WINDOWS_H
#define GFX_WINDOWS_H

#undef WINVER
#undef _WIN32_WINNT

// Target Windows 10 upwards
#define WINVER 0x0A00
#define _WIN32_WINNT 0x0A00

#include <windows.h>
#include <stdint.h>

#define GFX_CLASS_ERROR (-1002)
#define GFX_CALL_ERROR (-1001)
#define GFX_MODULE_ERROR (-1000)
#define GFX_SEE_ERROR (-10)
#define GFX_SUCCESS 1

int gfx_windows_init();

void gfx_windows_init_callback();

int gfx_windows_new_window(
        uint64_t wid,
        LPCWSTR title,
        int width,
        int height
);

#endif //GFX_WINDOWS_H
