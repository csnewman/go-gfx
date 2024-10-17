#import <Metal/MTLDevice.h>
#include "metal.h"

int gfx_mtl_open(id *res) {
    @autoreleasepool {
        id<MTLDevice> device = MTLCreateSystemDefaultDevice();

        *res = device;

        return GFX_SUCCESS;
    }
}
