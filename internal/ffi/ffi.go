package ffi

/*
#include <stdlib.h>
#include "ffi.h"
*/
import "C"

import (
	"log"
	"unsafe"
)

type Library struct {
	ptr unsafe.Pointer
}

func Open(libs []string) *Library {
	for _, lib := range libs {
		log.Println("trying", lib)

		str := C.CString(lib)
		defer C.free(unsafe.Pointer(str))

		ptr := C.gfx_ffi_open(str)
		if ptr == nil {
			continue
		}

		return &Library{
			ptr: ptr,
		}
	}

	return nil
}

func (l *Library) Handle() unsafe.Pointer {
	return l.ptr
}

func (l *Library) Symbol(name string) unsafe.Pointer {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))

	return C.gfx_ffi_get(l.ptr, cName)
}
