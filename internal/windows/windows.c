#include "helper.h"

#define GFX_WEVENT_INIT (WM_APP+1)

HMODULE gfx_win_module;

int gfx_windows_init() {
    if (!GetModuleHandleExW(
            GET_MODULE_HANDLE_EX_FLAG_FROM_ADDRESS | GET_MODULE_HANDLE_EX_FLAG_UNCHANGED_REFCOUNT,
            NULL,
            &gfx_win_module
    )) {
        return GFX_MODULE_ERROR;
    }

    SetProcessDpiAwarenessContext(DPI_AWARENESS_CONTEXT_PER_MONITOR_AWARE_V2);

    DWORD main = GetCurrentThreadId();
    PostThreadMessage(main, GFX_WEVENT_INIT, 0, 0);

    MSG msg;

    while (GetMessage(&msg, NULL, 0, 0)) {
        TranslateMessage(&msg);

        if (msg.message == GFX_WEVENT_INIT) {
            gfx_windows_init_callback();
        }

        DispatchMessage(&msg);
    }

    return GFX_SUCCESS;
}
