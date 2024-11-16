#include "helper.h"

#define GFX_WEVENT_INIT (WM_APP+1)

HMODULE gfx_win_module;
ATOM gfx_win_class;

LRESULT gfx_window_proc(HWND hwnd, UINT uMsg, WPARAM wParam, LPARAM lParam);

HMODULE GetCurrentModuleHandle() {
    HMODULE ImageBase;
    if (GetModuleHandleExW(
            GET_MODULE_HANDLE_EX_FLAG_FROM_ADDRESS | GET_MODULE_HANDLE_EX_FLAG_UNCHANGED_REFCOUNT,
            (LPCWSTR) &GetCurrentModuleHandle,
            &ImageBase
    )) {
        return ImageBase;
    }
    return 0;
}

int gfx_windows_init() {
    gfx_win_module = GetCurrentModuleHandle();
    if (!gfx_win_module) {
        return GFX_MODULE_ERROR;
    }

    SetProcessDpiAwarenessContext(DPI_AWARENESS_CONTEXT_PER_MONITOR_AWARE_V2);

    WNDCLASSEXW wc = {sizeof(wc)};
    wc.style = CS_OWNDC;
    wc.lpfnWndProc = (WNDPROC) gfx_window_proc;
    wc.hInstance = gfx_win_module;
    wc.hCursor = LoadCursorW(NULL, (LPTSTR) IDC_ARROW);
    wc.lpszClassName = L"GFXWindowClass";

    gfx_win_class = RegisterClassExW(&wc);
    if (!gfx_win_class) {
        return GFX_CLASS_ERROR;
    }

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

int gfx_windows_new_window(
        uint64_t wid,
        LPCWSTR title,
        int width,
        int height
) {
    HWND hwnd = CreateWindowExW(
            WS_EX_OVERLAPPEDWINDOW | WS_EX_APPWINDOW,
            MAKEINTATOM(gfx_win_class),
            title,
            WS_CLIPSIBLINGS | WS_CLIPCHILDREN | WS_TILEDWINDOW,
            CW_USEDEFAULT, CW_USEDEFAULT, width, height,
            NULL, NULL,
            gfx_win_module,
            NULL
    );
    if (hwnd == NULL) {
        return GFX_CALL_ERROR;
    }

    ShowWindow(hwnd, SW_NORMAL);

    return GFX_SUCCESS;
}

LRESULT gfx_window_proc(HWND hwnd, UINT uMsg, WPARAM wParam, LPARAM lParam) {
    return DefWindowProcW(hwnd, uMsg, wParam, lParam);
}
