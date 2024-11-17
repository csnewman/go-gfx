package spvc

import (
	"log"
	"unsafe"
)

import "C"

//export gfx_spvc_error_callback
func gfx_spvc_error_callback(userData unsafe.Pointer, ptr *C.char) {
	val := C.GoString(ptr)

	log.Println("error callback", val)
}
