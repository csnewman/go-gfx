#include "sdl_wrapper.h"

//file(GLOB SOURCE_FILES
//  ${SDL2_SOURCE_DIR}/src/*.c
#include "../thirdparty/SDL/src/SDL.c"
#include "../thirdparty/SDL/src/SDL_assert.c"
#include "../thirdparty/SDL/src/SDL_dataqueue.c"
#include "../thirdparty/SDL/src/SDL_error.c"
#include "../thirdparty/SDL/src/SDL_guid.c"
#include "../thirdparty/SDL/src/SDL_hints.c"
#include "../thirdparty/SDL/src/SDL_list.c"
#include "../thirdparty/SDL/src/SDL_log.c"
#include "../thirdparty/SDL/src/SDL_utils.c"

//  ${SDL2_SOURCE_DIR}/src/atomic/*.c
#include "../thirdparty/SDL/src/atomic/SDL_atomic.c"
#include "../thirdparty/SDL/src/atomic/SDL_spinlock.c"

//  ${SDL2_SOURCE_DIR}/src/audio/*.c
//  ${SDL2_SOURCE_DIR}/src/cpuinfo/*.c
#include "../thirdparty/SDL/src/cpuinfo/SDL_cpuinfo.c"

//  ${SDL2_SOURCE_DIR}/src/dynapi/*.c
//  ${SDL2_SOURCE_DIR}/src/events/*.c
#include "../thirdparty/SDL/src/events/SDL_events.c"
#include "../thirdparty/SDL/src/events/SDL_mouse.c"
#include "../thirdparty/SDL/src/events/SDL_touch.c"
#include "../thirdparty/SDL/src/events/SDL_gesture.c"
#include "../thirdparty/SDL/src/events/SDL_keyboard.c"
#include "../thirdparty/SDL/src/events/SDL_windowevents.c"
#include "../thirdparty/SDL/src/events/SDL_displayevents.c"
#include "../thirdparty/SDL/src/events/SDL_quit.c"
#include "../thirdparty/SDL/src/events/SDL_dropevents.c"
#include "../thirdparty/SDL/src/events/SDL_clipboardevents.c"

//  ${SDL2_SOURCE_DIR}/src/file/*.c
#include "../thirdparty/SDL/src/file/SDL_rwops.c"

//  ${SDL2_SOURCE_DIR}/src/joystick/*.c
//  ${SDL2_SOURCE_DIR}/src/haptic/*.c
//  ${SDL2_SOURCE_DIR}/src/hidapi/*.c
//  ${SDL2_SOURCE_DIR}/src/libm/*.c
//  ${SDL2_SOURCE_DIR}/src/locale/*.c
//  ${SDL2_SOURCE_DIR}/src/misc/*.c
//  ${SDL2_SOURCE_DIR}/src/power/*.c
//  ${SDL2_SOURCE_DIR}/src/render/*.c
#include "../thirdparty/SDL/src/render/SDL_render.c"

//  ${SDL2_SOURCE_DIR}/src/render/*/*.c
//  ${SDL2_SOURCE_DIR}/src/sensor/*.c
//  ${SDL2_SOURCE_DIR}/src/stdlib/*.c
#include "../thirdparty/SDL/src/stdlib/SDL_crc16.c"
#include "../thirdparty/SDL/src/stdlib/SDL_crc32.c"
#include "../thirdparty/SDL/src/stdlib/SDL_getenv.c"
#include "../thirdparty/SDL/src/stdlib/SDL_iconv.c"
#include "../thirdparty/SDL/src/stdlib/SDL_malloc.c"
#include "../thirdparty/SDL/src/stdlib/SDL_mslibc.c"
//#include "../thirdparty/SDL/src/stdlib/SDL_qsort.c"
#include "../thirdparty/SDL/src/stdlib/SDL_stdlib.c"
#include "../thirdparty/SDL/src/stdlib/SDL_string.c"
#include "../thirdparty/SDL/src/stdlib/SDL_strtokr.c"

//  ${SDL2_SOURCE_DIR}/src/thread/*.c
#include "../thirdparty/SDL/src/thread/SDL_thread.c"

#if defined(__WIN32__)
#include "../thirdparty/SDL/src/thread/generic/SDL_syscond.c"
//#include "../thirdparty/SDL/src/thread/windows/SDL_syscond_cv.c"
#include "../thirdparty/SDL/src/thread/windows/SDL_sysmutex.c"
//#include "../thirdparty/SDL/src/thread/windows/SDL_syssem.c"
#include "../thirdparty/SDL/src/thread/windows/SDL_systhread.c"
#include "../thirdparty/SDL/src/thread/windows/SDL_systls.c"
#else
//#include "../thirdparty/SDL/src/thread/pthread/SDL_syscond.c"
#include "../thirdparty/SDL/src/thread/pthread/SDL_sysmutex.c"
//#include "../thirdparty/SDL/src/thread/pthread/SDL_syssem.c"
#include "../thirdparty/SDL/src/thread/pthread/SDL_systhread.c"
#include "../thirdparty/SDL/src/thread/pthread/SDL_systls.c"
#endif

//  ${SDL2_SOURCE_DIR}/src/timer/*.c
#include "../thirdparty/SDL/src/timer/SDL_timer.c"

#if defined(__WIN32__)
#include "../thirdparty/SDL/src/timer/windows/SDL_systimer.c"
#else
#include "../thirdparty/SDL/src/timer/unix/SDL_systimer.c"
#endif

//  ${SDL2_SOURCE_DIR}/src/video/*.c
#include "../thirdparty/SDL/src/video/SDL_video.c"
#include "../thirdparty/SDL/src/video/SDL_stretch.c"
#include "../thirdparty/SDL/src/video/SDL_surface.c"
#include "../thirdparty/SDL/src/video/SDL_pixels.c"
#include "../thirdparty/SDL/src/video/SDL_rect.c"
#include "../thirdparty/SDL/src/video/SDL_fillrect.c"
#include "../thirdparty/SDL/src/video/SDL_blit.c"
#include "../thirdparty/SDL/src/video/SDL_blit_copy.c"
#include "../thirdparty/SDL/src/video/SDL_blit_slow.c"
#include "../thirdparty/SDL/src/video/SDL_vulkan_utils.c"
#include "../thirdparty/SDL/src/video/SDL_shape.c"

#if defined(__WIN32__)
#include "../thirdparty/SDL/src/loadso/windows/SDL_sysloadso.c"
#else
#include "../thirdparty/SDL/src/loadso/dlopen/SDL_sysloadso.c"
#endif

//  ${SDL2_SOURCE_DIR}/src/video/yuv2rgb/*.c)
