#include "ffi.h"
#include <dlfcn.h>

void* gfx_ffi_open(const char* path) {
    return dlopen(path, RTLD_LAZY | RTLD_LOCAL);
}

void* gfx_ffi_get(void* lib, const char* name) {
    return dlsym(lib, name);
}
