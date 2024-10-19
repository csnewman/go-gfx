#include <stdint.h>

#if defined(__OBJC__)

#import <objc/objc.h>

#else
typedef void *id;
#endif

#define GFX_SUCCESS 1

int gfx_mtl_open(id *res, id *res_queue);

int gfx_mtl_configure_surface(id device, id layer);

void gfx_mtl_acquire_texture(id layer, id *res);

void gfx_mtl_present_texture(id queue, id text);

void gfx_mtl_discard_surface_texture(id text);
