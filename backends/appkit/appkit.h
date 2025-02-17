#include <stdint.h>

#if defined(__OBJC__)

#import <objc/objc.h>

#else
typedef void *id;
#endif

#define GFX_NOT_MAIN_THREAD (-10)
#define GFX_SUCCESS 1

int gfx_ak_run();

void gfx_ak_process_events();

void gfx_ak_stop();

void gfx_ak_init_callback();

int gfx_ak_new_window(
        uint64_t wid,
        const void *title,
        int title_len,
        int width,
        int height,
        id *res,
        id *res_wh
);

void gfx_ak_close_requested_callback(uint64_t wid);

void gfx_ak_window_closed_callback(uint64_t wid);

void gfx_ak_draw_callback(uint64_t wid, id drawable);

void gfx_ak_resize_callback(uint64_t wid, double width, double height, double scale);

void gfx_ak_size(id w, double *width, double *height, double *scale);

void gfx_ak_close_window(id w);

void gfx_ak_free_context(id w);
