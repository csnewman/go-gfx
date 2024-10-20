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

id gfx_mtl_get_drawable_texture(id <CAMetalDrawable> drawable) {
    @autoreleasepool {
        [drawable retain];
        return [drawable texture];
    }
}

void gfx_mtl_present_texture(id <MTLCommandQueue> queue, id <CAMetalDrawable> draw) {
    @autoreleasepool {
        id <MTLCommandBuffer> buffer = [queue commandBuffer];
        [buffer presentDrawable:draw];
        [buffer commit];
        [draw release];
    }
}

void gfx_mtl_discard_surface_texture(id <CAMetalDrawable> draw) {
    @autoreleasepool {
         [draw release];
    }
}

void gfx_mtl_create_command_buf(id <MTLCommandQueue> queue, id *res) {
    @autoreleasepool {
        id <MTLCommandBuffer> buf = [queue commandBuffer];

        *res = [buf retain];
    }
}

void gfx_mtl_cbuf_submit(id <MTLCommandBuffer> buf) {
    @autoreleasepool {
        [buf commit];
        [buf release];
    }
}

void gfx_mtl_begin_rpass(
        id <MTLCommandBuffer> buf,
        const struct ColorAttachment *colors,
        uint64_t colors_len,
        id *res
) {
    @autoreleasepool {
        MTLRenderPassDescriptor *des = [[MTLRenderPassDescriptor new] autorelease];

        for (int i = 0; i < colors_len; ++i) {
            const struct ColorAttachment attachment = colors[i];

            des.colorAttachments[i].texture = attachment.view;

            if (attachment.load) {
                des.colorAttachments[i].loadAction = MTLLoadActionLoad;
            } else {
                des.colorAttachments[i].loadAction = MTLLoadActionClear;
                des.colorAttachments[i].clearColor = MTLClearColorMake(
                        attachment.r, attachment.g, attachment.b, attachment.a
                );
            }

            if (attachment.store) {
                des.colorAttachments[i].storeAction = MTLStoreActionStore;
            } else {
                des.colorAttachments[i].storeAction = MTLStoreActionDontCare;
            }
        }

        id <MTLRenderCommandEncoder> enc = [buf renderCommandEncoderWithDescriptor:des];
        *res = [enc retain];
    }
}

void gfx_mtl_end_rpass(id <MTLRenderCommandEncoder> enc) {
    @autoreleasepool {
        [enc endEncoding];
        [enc release];
    }
}
