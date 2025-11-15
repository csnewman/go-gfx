package ffi

import "unsafe"

/*
#include <stdlib.h>
*/
import "C"

var GlobalAllocator Allocator = CAllocator{}

type CAllocator struct{}

func (a CAllocator) Allocate(size int) unsafe.Pointer {
	return C.calloc(C.size_t(size), C.size_t(1))
}

func (a CAllocator) Free(ptr unsafe.Pointer) {
	C.free(ptr)
}

type Arena struct {
}

func NewArena() *Arena {
	return &Arena{}
}

func (a *Arena) Checkpoint() func() {
	return func() {}
}

func (a *Arena) Close() {}

func (a *Arena) Allocate(size int) unsafe.Pointer {
	// TODO: replace
	//return C.malloc(C.size_t(size))
	return C.calloc(C.size_t(size), C.size_t(1))
}
