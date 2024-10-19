#import <Metal/Metal.h>
#import <QuartzCore/CAMetalLayer.h>
#include "metal.h"

int gfx_mtl_open(id *res, id *res_queue) {
    @autoreleasepool {
        id <MTLDevice> device = MTLCreateSystemDefaultDevice();

        *res = device;
        *res_queue = [device newCommandQueue];

        return GFX_SUCCESS;
    }
}

int gfx_mtl_configure_surface(id <MTLDevice> device, CAMetalLayer *layer) {
    @autoreleasepool {
        [layer setDevice:device];
        // TODO: pixelFormat
        [layer setFramebufferOnly:YES];
        // TODO: colorspace?
        // TODO: wantsExtendedDynamicRangeContent?
        // TODO: expose
        [layer setDisplaySyncEnabled:YES];

        return GFX_SUCCESS;
    }
}

void gfx_mtl_acquire_texture(CAMetalLayer *layer, id *res) {
    @autoreleasepool {
        id <CAMetalDrawable> drawable = [layer nextDrawable];

        if (drawable != nil) {
            *res = [drawable retain];
        }
    }
}

void gfx_mtl_present_texture(id <MTLCommandQueue> queue, id <CAMetalDrawable> texture) {
    @autoreleasepool {
        id <MTLCommandBuffer> buffer = [queue commandBuffer];
        [buffer presentDrawable:texture];
        [buffer commit];
        [texture release];
    }
}

void gfx_mtl_discard_surface_texture(id <CAMetalDrawable> texture) {
    @autoreleasepool {
        [texture release];
    }
}
