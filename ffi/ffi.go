package ffi

import "unsafe"

/*
#include <stdlib.h>
*/
import "C"

type Allocator interface {
	Allocate(size int) unsafe.Pointer
}

type Arena struct {
}

func (a *Arena) Checkpoint() func() {
	return func() {}
}

func (a *Arena) Allocate(size int) unsafe.Pointer {
	// TODO: replace
	return C.malloc(C.size_t(size))
}

type CString struct {
	ptr unsafe.Pointer
}

func CStringFromPtr(ptr unsafe.Pointer) CString {
	return CString{ptr: ptr}
}

func CStringAlloc[T Allocator](alloc T, value string) CString {
	ptr := alloc.Allocate(len(value) + 1)

	dst := unsafe.Slice((*byte)(ptr), len(value))

	copy(dst, []byte(value))
	dst[len(value)] = 0

	return CString{ptr: ptr}
}

// Raw returns a raw pointer to the string.
func (p CString) Raw() unsafe.Pointer {
	return p.ptr
}
