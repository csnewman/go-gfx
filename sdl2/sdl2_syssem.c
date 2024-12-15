#include "sdl_wrapper.h"

#if defined(__WIN32__)
#include "../thirdparty/SDL/src/thread/windows/SDL_syssem.c"
#else
#include "../thirdparty/SDL/src/thread/pthread/SDL_syssem.c"
#endif
