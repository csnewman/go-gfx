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

int gfx_mtl_configure_surface(id <MTLDevice> device, CAMetalLayer *layer, int pixelFormat) {
    @autoreleasepool {
        [layer setDevice:device];
        [layer setPixelFormat:pixelFormat];
        [layer setFramebufferOnly:YES];
        // TODO: colorspace?
        // TODO: wantsExtendedDynamicRangeContent?
        // TODO: expose
        [layer setDisplaySyncEnabled:YES];

        return GFX_SUCCESS;
    }
}

void gfx_mtl_acquire_surface(CAMetalLayer *layer, id* res_draw, id* res_text) {
    @autoreleasepool {
        id<CAMetalDrawable> drawable = [[layer nextDrawable] retain];
        *res_draw = drawable;
        *res_text = [drawable texture];
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

int gfx_mtl_create_shader(id <MTLDevice> device, const void *src, int src_len, id *res_lib, char **res_err) {
    @autoreleasepool {
        NSError *error = nil;
        NSString *libSrc = [[[NSString alloc] initWithBytes:src length:src_len encoding:NSUTF8StringEncoding] autorelease];

        id <MTLLibrary> lib = [device newLibraryWithSource:libSrc options:nil error:&error];

        if (error != nil) {
            *res_err = strdup([error.localizedDescription UTF8String]);
            return GFX_SEE_ERROR;
        }

        *res_lib = lib;

        return GFX_SUCCESS;
    }
}

void gfx_mtl_get_shader_function(id <MTLLibrary> lib, const void *name, int name_len, id *res) {
    @autoreleasepool {
        NSString *fnName = [[[NSString alloc] initWithBytes:name length:name_len encoding:NSUTF8StringEncoding] autorelease];
        *res = [lib newFunctionWithName:fnName];
    }
}

void gfx_mtl_buffer_from_bytes(id <MTLDevice> device, const void *data, int data_len, id *res) {
    // TODO: no copy + options
    *res = [device newBufferWithBytes:data length:data_len options:MTLResourceStorageModeShared];
}

int gfx_mtl_create_render_pipeline(
        id <MTLDevice> device,
        id <MTLFunction> vertFunc,
        id <MTLFunction> fragFunc,
        const struct PipelineColorAttachment *colors,
        uint64_t colors_len,
        id *res_lib,
        char **res_err
) {
    @autoreleasepool {
        MTLRenderPipelineDescriptor *des = [[MTLRenderPipelineDescriptor new] autorelease];
        // TODO: label

        [des setVertexFunction:vertFunc];
        [des setFragmentFunction:fragFunc];


        for (int i = 0; i < colors_len; ++i) {
            const struct PipelineColorAttachment attachment = colors[i];

            des.colorAttachments[i].pixelFormat = attachment.format;
            // TODO: rest of options
        }

        NSError *error = nil;
        id <MTLRenderPipelineState> state = [device newRenderPipelineStateWithDescriptor:des error:&error];

        if (error != nil) {
            *res_err = strdup([error.localizedDescription UTF8String]);
            return GFX_SEE_ERROR;
        }

        *res_lib = state;

        return GFX_SUCCESS;
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

void gfx_mtl_set_render_pipeline(id <MTLRenderCommandEncoder> enc, id <MTLRenderPipelineState> pipeline) {
    @autoreleasepool {
        [enc setRenderPipelineState:pipeline];
    }
}

void gfx_mtl_set_vertex_buffer(id <MTLRenderCommandEncoder> enc, id <MTLBuffer> buffer) {
    @autoreleasepool {
        [enc setVertexBuffer:buffer offset:0 atIndex:0];
    }
}

void gfx_mtl_draw(id <MTLRenderCommandEncoder> enc, int start, int count) {
    @autoreleasepool {
        [enc drawPrimitives:MTLPrimitiveTypeTriangle vertexStart:start vertexCount:count];
    }
}

void gfx_mtl_end_rpass(id <MTLRenderCommandEncoder> enc) {
    @autoreleasepool {
        [enc endEncoding];
        [enc release];
    }
}
