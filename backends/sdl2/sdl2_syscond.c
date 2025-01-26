#include "sdl_wrapper.h"

#if defined(__WIN32__)
#include "../../thirdparty/SDL/src/thread/windows/SDL_syscond_cv.c"
#else
#include "../../thirdparty/SDL/src/thread/pthread/SDL_syscond.c"
#endif
