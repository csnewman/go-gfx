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

func CStringAlloc(alloc Allocator, value string) CString {
	ptr := alloc.Allocate(len(value) + 1)

	dst := unsafe.Slice((*byte)(ptr), len(value)+1)

	copy(dst, []byte(value))
	dst[len(value)] = 0

	return CString{ptr: ptr}
}

// Raw returns a raw pointer to the string.
func (p CString) Raw() unsafe.Pointer {
	return p.ptr
}

type Primitive interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr | ~float32 | ~float64 | CString
}

type Ref[T Primitive] struct {
	ptr unsafe.Pointer
}

func RefFromPtr[T Primitive](ptr unsafe.Pointer) Ref[T] {
	return Ref[T]{ptr: ptr}
}

func RefAlloc[T Primitive](alloc Allocator, count int) Ref[T] {
	var empty T

	size := unsafe.Sizeof(empty)

	ptr := alloc.Allocate(int(size) * count)

	return Ref[T]{
		ptr: ptr,
	}
}

func RefNil[T Primitive]() Ref[T] {
	return Ref[T]{}
}

func (r Ref[T]) Set(value T) {
	*(*T)(r.ptr) = value
}

func (r Ref[T]) Get() T {
	return *(*T)(r.ptr)
}

func (r Ref[T]) Offset(offset int) Ref[T] {
	var empty T

	size := unsafe.Sizeof(empty)
	ptr := unsafe.Add(r.ptr, offset*int(size))

	return Ref[T]{
		ptr: ptr,
	}
}

func (r Ref[T]) Raw() unsafe.Pointer {
	return r.ptr
}
