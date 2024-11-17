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

int gfx_windows_init(HMODULE* inst);

void gfx_windows_init_callback();

int gfx_windows_new_window(
        uint64_t wid,
        LPCWSTR title,
        int width,
        int height,
        HWND* res
);

void gfx_windows_resize_callback(uint64_t id, int64_t width, int64_t height);

void gfx_windows_close_requested_callback(uint64_t wid);

void gfx_windows_window_closed_callback(uint64_t wid);

void gfx_windows_close_window(HWND w);

void gfx_windows_exit();

#endif //GFX_WINDOWS_H
