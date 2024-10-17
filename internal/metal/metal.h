#include <stdint.h>

#if defined(__OBJC__)

#import <objc/objc.h>

#else
typedef void *id;
#endif

#define GFX_SUCCESS 1

int gfx_mtl_open(id *res);
