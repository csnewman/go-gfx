#include <stdint.h>

#if defined(__OBJC__)

#import <objc/objc.h>

#else
typedef void *id;
#endif

#define GFX_NOT_MAIN_THREAD (-10)
#define GFX_SUCCESS 1

int gfx_ak_run();

void gfx_ak_stop();

void gfx_ak_init_callback();

int gfx_ak_new_window(uint32_t wid, int width, int height, id *res);

void gfx_ak_close_requested_callback(uint32_t wid);

void gfx_ak_window_closed_callback(uint32_t wid);

void gfx_ak_close_window(id w);

void gfx_ak_free_context(id w);
