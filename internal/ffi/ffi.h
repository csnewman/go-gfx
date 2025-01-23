#ifndef GFX_FFI_H
#define GFX_FFI_H

void* gfx_ffi_open(const char* path);

extern void* gfx_ffi_get(void* lib, const char* name);

#endif //GFX_FFI_H
