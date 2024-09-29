package appkit

import (
	"log"
	"sync"
	"sync/atomic"
)

/*
#include "appkit.h"
*/
import "C"

var (
	windowCounter atomic.Uint32
	windows       sync.Map
)

type Window uint32

func NewWindow(width int, height int) error {
	var res C.id

	id := Window(windowCounter.Add(1))

	r := C.gfx_ak_new_window(C.uint32_t(id), C.int(width), C.int(height), &res)

	log.Println(id, res)

	windows.Store(id, res)

	switch r {
	case C.GFX_SUCCESS:
		return nil

	default:
		panic("unexpected response")
	}
}

//export gfx_ak_close_requested_callback
func gfx_ak_close_requested_callback(id uint32) {
	log.Println("close requested", id)

}

//export gfx_ak_window_closed_callback
func gfx_ak_window_closed_callback(id uint32) {
	log.Println("closed", id)
}
