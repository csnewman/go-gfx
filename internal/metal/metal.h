#include <stdint.h>
#include <stdbool.h>

#if defined(__OBJC__)

#import <objc/objc.h>

#else
typedef void *id;
#endif

#define GFX_SUCCESS 1

int gfx_mtl_open(id *res, id *res_queue);

int gfx_mtl_configure_surface(id device, id layer);

void gfx_mtl_acquire_texture(id layer, id *res_draw, id *res_tex);

void gfx_mtl_present_texture(id queue, id draw);

void gfx_mtl_discard_surface_texture(id draw);

void gfx_mtl_create_command_buf(id queue, id *res);

typedef struct ColorAttachment {
    id view;
    bool load;
    bool store;
    double r;
    double g;
    double b;
    double a;
} ColorAttachment;

void gfx_mtl_begin_rpass(id buf, const struct ColorAttachment *colors, uint64_t colors_len, id *res);

void gfx_mtl_end_rpass(id enc);

void gfx_mtl_cbuf_submit(id buffer);
