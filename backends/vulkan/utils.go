package vulkan

import "github.com/csnewman/go-gfx/ffi"

func convertStringArray(arena *ffi.Arena, in []string) ffi.Ref[ffi.CString] {
	ref := ffi.RefAlloc[ffi.CString](arena, len(in))

	for i, v := range in {
		ref.Offset(i).Set(ffi.CStringAlloc(arena, v))
	}

	return ref
}
