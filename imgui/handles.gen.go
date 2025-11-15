package imgui

// #include "imgui_wrapper.h"
import "C"

// DrawIdx wraps the handle ImDrawIdx.
type DrawIdx uint16

// DrawIdxNil is a null pointer.
var DrawIdxNil DrawIdx

// FontAtlasRectId wraps the handle ImFontAtlasRectId.
type FontAtlasRectId uint32

// FontAtlasRectIdNil is a null pointer.
var FontAtlasRectIdNil FontAtlasRectId
