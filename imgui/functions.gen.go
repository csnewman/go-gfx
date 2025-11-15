package imgui

import (
	ffi "github.com/csnewman/go-gfx/ffi"
	"unsafe"
)

// #include "imgui_wrapper.h"
import "C"

// ImColor_HSV wraps ImColor_HSV.
func ImColor_HSV(pOut Color, h float32, s float32, v float32, a float32) {
	C.ImColor_HSV((*C.ImColor)(unsafe.Pointer(pOut)), C.float(h), C.float(s), C.float(v), C.float(a))
}

// ColorNew_Float wraps ImColor_ImColor_Float.
func ColorNew_Float(r float32, g float32, b float32, a float32) Color {
	ret := C.ImColor_ImColor_Float(C.float(r), C.float(g), C.float(b), C.float(a))

	return Color(unsafe.Pointer(ret))
}

// ColorNew_Int wraps ImColor_ImColor_Int.
func ColorNew_Int(r int32, g int32, b int32, a int32) Color {
	ret := C.ImColor_ImColor_Int(C.int32_t(r), C.int32_t(g), C.int32_t(b), C.int32_t(a))

	return Color(unsafe.Pointer(ret))
}

// ColorNew_Nil wraps ImColor_ImColor_Nil.
func ColorNew_Nil() Color {
	ret := C.ImColor_ImColor_Nil()

	return Color(unsafe.Pointer(ret))
}

// ColorNew_U32 wraps ImColor_ImColor_U32.
func ColorNew_U32(rgba uint32) Color {
	ret := C.ImColor_ImColor_U32(C.uint32_t(rgba))

	return Color(unsafe.Pointer(ret))
}

// ColorNew_Vec4 wraps ImColor_ImColor_Vec4.
func ColorNew_Vec4(col Vec4) Color {
	ret := C.ImColor_ImColor_Vec4(*(*C.ImVec4)(unsafe.Pointer(col)))

	return Color(unsafe.Pointer(ret))
}

// SetHSV wraps ImColor_SetHSV.
func (p Color) SetHSV(h float32, s float32, v float32, a float32) {
	C.ImColor_SetHSV((*C.ImColor)(unsafe.Pointer(p)), C.float(h), C.float(s), C.float(v), C.float(a))
}

// Destroy wraps ImColor_destroy.
func (p Color) Destroy() {
	C.ImColor_destroy((*C.ImColor)(unsafe.Pointer(p)))
}

// ImDrawCmd_GetTexID is unsupported: unknown return category unsupported.

// DrawCmdNew wraps ImDrawCmd_ImDrawCmd.
func DrawCmdNew() DrawCmd {
	ret := C.ImDrawCmd_ImDrawCmd()

	return DrawCmd(unsafe.Pointer(ret))
}

// Destroy wraps ImDrawCmd_destroy.
func (p DrawCmd) Destroy() {
	C.ImDrawCmd_destroy((*C.ImDrawCmd)(unsafe.Pointer(p)))
}

// AddDrawList wraps ImDrawData_AddDrawList.
func (p DrawData) AddDrawList(draw_list DrawList) {
	C.ImDrawData_AddDrawList((*C.ImDrawData)(unsafe.Pointer(p)), (*C.ImDrawList)(unsafe.Pointer(draw_list)))
}

// Clear wraps ImDrawData_Clear.
func (p DrawData) Clear() {
	C.ImDrawData_Clear((*C.ImDrawData)(unsafe.Pointer(p)))
}

// DeIndexAllBuffers wraps ImDrawData_DeIndexAllBuffers.
func (p DrawData) DeIndexAllBuffers() {
	C.ImDrawData_DeIndexAllBuffers((*C.ImDrawData)(unsafe.Pointer(p)))
}

// DrawDataNew wraps ImDrawData_ImDrawData.
func DrawDataNew() DrawData {
	ret := C.ImDrawData_ImDrawData()

	return DrawData(unsafe.Pointer(ret))
}

// ScaleClipRects wraps ImDrawData_ScaleClipRects.
func (p DrawData) ScaleClipRects(fb_scale Vec2) {
	C.ImDrawData_ScaleClipRects((*C.ImDrawData)(unsafe.Pointer(p)), *(*C.ImVec2)(unsafe.Pointer(fb_scale)))
}

// Destroy wraps ImDrawData_destroy.
func (p DrawData) Destroy() {
	C.ImDrawData_destroy((*C.ImDrawData)(unsafe.Pointer(p)))
}

// Clear wraps ImDrawListSplitter_Clear.
func (p DrawListSplitter) Clear() {
	C.ImDrawListSplitter_Clear((*C.ImDrawListSplitter)(unsafe.Pointer(p)))
}

// ClearFreeMemory wraps ImDrawListSplitter_ClearFreeMemory.
func (p DrawListSplitter) ClearFreeMemory() {
	C.ImDrawListSplitter_ClearFreeMemory((*C.ImDrawListSplitter)(unsafe.Pointer(p)))
}

// DrawListSplitterNew wraps ImDrawListSplitter_ImDrawListSplitter.
func DrawListSplitterNew() DrawListSplitter {
	ret := C.ImDrawListSplitter_ImDrawListSplitter()

	return DrawListSplitter(unsafe.Pointer(ret))
}

// Merge wraps ImDrawListSplitter_Merge.
func (p DrawListSplitter) Merge(draw_list DrawList) {
	C.ImDrawListSplitter_Merge((*C.ImDrawListSplitter)(unsafe.Pointer(p)), (*C.ImDrawList)(unsafe.Pointer(draw_list)))
}

// SetCurrentChannel wraps ImDrawListSplitter_SetCurrentChannel.
func (p DrawListSplitter) SetCurrentChannel(draw_list DrawList, channel_idx int32) {
	C.ImDrawListSplitter_SetCurrentChannel((*C.ImDrawListSplitter)(unsafe.Pointer(p)), (*C.ImDrawList)(unsafe.Pointer(draw_list)), C.int32_t(channel_idx))
}

// Split wraps ImDrawListSplitter_Split.
func (p DrawListSplitter) Split(draw_list DrawList, count int32) {
	C.ImDrawListSplitter_Split((*C.ImDrawListSplitter)(unsafe.Pointer(p)), (*C.ImDrawList)(unsafe.Pointer(draw_list)), C.int32_t(count))
}

// Destroy wraps ImDrawListSplitter_destroy.
func (p DrawListSplitter) Destroy() {
	C.ImDrawListSplitter_destroy((*C.ImDrawListSplitter)(unsafe.Pointer(p)))
}

// AddBezierCubic wraps ImDrawList_AddBezierCubic.
func (p DrawList) AddBezierCubic(p1 Vec2, p2 Vec2, p3 Vec2, p4 Vec2, col uint32, thickness float32, num_segments int32) {
	C.ImDrawList_AddBezierCubic((*C.ImDrawList)(unsafe.Pointer(p)), *(*C.ImVec2)(unsafe.Pointer(p1)), *(*C.ImVec2)(unsafe.Pointer(p2)), *(*C.ImVec2)(unsafe.Pointer(p3)), *(*C.ImVec2)(unsafe.Pointer(p4)), C.uint32_t(col), C.float(thickness), C.int32_t(num_segments))
}

// AddBezierQuadratic wraps ImDrawList_AddBezierQuadratic.
func (p DrawList) AddBezierQuadratic(p1 Vec2, p2 Vec2, p3 Vec2, col uint32, thickness float32, num_segments int32) {
	C.ImDrawList_AddBezierQuadratic((*C.ImDrawList)(unsafe.Pointer(p)), *(*C.ImVec2)(unsafe.Pointer(p1)), *(*C.ImVec2)(unsafe.Pointer(p2)), *(*C.ImVec2)(unsafe.Pointer(p3)), C.uint32_t(col), C.float(thickness), C.int32_t(num_segments))
}

// ImDrawList_AddCallback.callback is unsupported: category unsupported.

// AddCircle wraps ImDrawList_AddCircle.
func (p DrawList) AddCircle(center Vec2, radius float32, col uint32, num_segments int32, thickness float32) {
	C.ImDrawList_AddCircle((*C.ImDrawList)(unsafe.Pointer(p)), *(*C.ImVec2)(unsafe.Pointer(center)), C.float(radius), C.uint32_t(col), C.int32_t(num_segments), C.float(thickness))
}

// AddCircleFilled wraps ImDrawList_AddCircleFilled.
func (p DrawList) AddCircleFilled(center Vec2, radius float32, col uint32, num_segments int32) {
	C.ImDrawList_AddCircleFilled((*C.ImDrawList)(unsafe.Pointer(p)), *(*C.ImVec2)(unsafe.Pointer(center)), C.float(radius), C.uint32_t(col), C.int32_t(num_segments))
}

// AddConcavePolyFilled wraps ImDrawList_AddConcavePolyFilled.
func (p DrawList) AddConcavePolyFilled(points Vec2, num_points int32, col uint32) {
	C.ImDrawList_AddConcavePolyFilled((*C.ImDrawList)(unsafe.Pointer(p)), (*C.ImVec2)(unsafe.Pointer(points)), C.int32_t(num_points), C.uint32_t(col))
}

// AddConvexPolyFilled wraps ImDrawList_AddConvexPolyFilled.
func (p DrawList) AddConvexPolyFilled(points Vec2, num_points int32, col uint32) {
	C.ImDrawList_AddConvexPolyFilled((*C.ImDrawList)(unsafe.Pointer(p)), (*C.ImVec2)(unsafe.Pointer(points)), C.int32_t(num_points), C.uint32_t(col))
}

// AddDrawCmd wraps ImDrawList_AddDrawCmd.
func (p DrawList) AddDrawCmd() {
	C.ImDrawList_AddDrawCmd((*C.ImDrawList)(unsafe.Pointer(p)))
}

// AddEllipse wraps ImDrawList_AddEllipse.
func (p DrawList) AddEllipse(center Vec2, radius Vec2, col uint32, rot float32, num_segments int32, thickness float32) {
	C.ImDrawList_AddEllipse((*C.ImDrawList)(unsafe.Pointer(p)), *(*C.ImVec2)(unsafe.Pointer(center)), *(*C.ImVec2)(unsafe.Pointer(radius)), C.uint32_t(col), C.float(rot), C.int32_t(num_segments), C.float(thickness))
}

// AddEllipseFilled wraps ImDrawList_AddEllipseFilled.
func (p DrawList) AddEllipseFilled(center Vec2, radius Vec2, col uint32, rot float32, num_segments int32) {
	C.ImDrawList_AddEllipseFilled((*C.ImDrawList)(unsafe.Pointer(p)), *(*C.ImVec2)(unsafe.Pointer(center)), *(*C.ImVec2)(unsafe.Pointer(radius)), C.uint32_t(col), C.float(rot), C.int32_t(num_segments))
}

// AddImage wraps ImDrawList_AddImage.
func (p DrawList) AddImage(tex_ref TextureRef, p_min Vec2, p_max Vec2, uv_min Vec2, uv_max Vec2, col uint32) {
	C.ImDrawList_AddImage((*C.ImDrawList)(unsafe.Pointer(p)), *(*C.ImTextureRef)(unsafe.Pointer(tex_ref)), *(*C.ImVec2)(unsafe.Pointer(p_min)), *(*C.ImVec2)(unsafe.Pointer(p_max)), *(*C.ImVec2)(unsafe.Pointer(uv_min)), *(*C.ImVec2)(unsafe.Pointer(uv_max)), C.uint32_t(col))
}

// AddImageQuad wraps ImDrawList_AddImageQuad.
func (p DrawList) AddImageQuad(tex_ref TextureRef, p1 Vec2, p2 Vec2, p3 Vec2, p4 Vec2, uv1 Vec2, uv2 Vec2, uv3 Vec2, uv4 Vec2, col uint32) {
	C.ImDrawList_AddImageQuad((*C.ImDrawList)(unsafe.Pointer(p)), *(*C.ImTextureRef)(unsafe.Pointer(tex_ref)), *(*C.ImVec2)(unsafe.Pointer(p1)), *(*C.ImVec2)(unsafe.Pointer(p2)), *(*C.ImVec2)(unsafe.Pointer(p3)), *(*C.ImVec2)(unsafe.Pointer(p4)), *(*C.ImVec2)(unsafe.Pointer(uv1)), *(*C.ImVec2)(unsafe.Pointer(uv2)), *(*C.ImVec2)(unsafe.Pointer(uv3)), *(*C.ImVec2)(unsafe.Pointer(uv4)), C.uint32_t(col))
}

// AddImageRounded wraps ImDrawList_AddImageRounded.
func (p DrawList) AddImageRounded(tex_ref TextureRef, p_min Vec2, p_max Vec2, uv_min Vec2, uv_max Vec2, col uint32, rounding float32, flags DrawFlags) {
	C.ImDrawList_AddImageRounded((*C.ImDrawList)(unsafe.Pointer(p)), *(*C.ImTextureRef)(unsafe.Pointer(tex_ref)), *(*C.ImVec2)(unsafe.Pointer(p_min)), *(*C.ImVec2)(unsafe.Pointer(p_max)), *(*C.ImVec2)(unsafe.Pointer(uv_min)), *(*C.ImVec2)(unsafe.Pointer(uv_max)), C.uint32_t(col), C.float(rounding), C.ImDrawFlags(flags))
}

// AddLine wraps ImDrawList_AddLine.
func (p DrawList) AddLine(p1 Vec2, p2 Vec2, col uint32, thickness float32) {
	C.ImDrawList_AddLine((*C.ImDrawList)(unsafe.Pointer(p)), *(*C.ImVec2)(unsafe.Pointer(p1)), *(*C.ImVec2)(unsafe.Pointer(p2)), C.uint32_t(col), C.float(thickness))
}

// AddNgon wraps ImDrawList_AddNgon.
func (p DrawList) AddNgon(center Vec2, radius float32, col uint32, num_segments int32, thickness float32) {
	C.ImDrawList_AddNgon((*C.ImDrawList)(unsafe.Pointer(p)), *(*C.ImVec2)(unsafe.Pointer(center)), C.float(radius), C.uint32_t(col), C.int32_t(num_segments), C.float(thickness))
}

// AddNgonFilled wraps ImDrawList_AddNgonFilled.
func (p DrawList) AddNgonFilled(center Vec2, radius float32, col uint32, num_segments int32) {
	C.ImDrawList_AddNgonFilled((*C.ImDrawList)(unsafe.Pointer(p)), *(*C.ImVec2)(unsafe.Pointer(center)), C.float(radius), C.uint32_t(col), C.int32_t(num_segments))
}

// AddPolyline wraps ImDrawList_AddPolyline.
func (p DrawList) AddPolyline(points Vec2, num_points int32, col uint32, flags DrawFlags, thickness float32) {
	C.ImDrawList_AddPolyline((*C.ImDrawList)(unsafe.Pointer(p)), (*C.ImVec2)(unsafe.Pointer(points)), C.int32_t(num_points), C.uint32_t(col), C.ImDrawFlags(flags), C.float(thickness))
}

// AddQuad wraps ImDrawList_AddQuad.
func (p DrawList) AddQuad(p1 Vec2, p2 Vec2, p3 Vec2, p4 Vec2, col uint32, thickness float32) {
	C.ImDrawList_AddQuad((*C.ImDrawList)(unsafe.Pointer(p)), *(*C.ImVec2)(unsafe.Pointer(p1)), *(*C.ImVec2)(unsafe.Pointer(p2)), *(*C.ImVec2)(unsafe.Pointer(p3)), *(*C.ImVec2)(unsafe.Pointer(p4)), C.uint32_t(col), C.float(thickness))
}

// AddQuadFilled wraps ImDrawList_AddQuadFilled.
func (p DrawList) AddQuadFilled(p1 Vec2, p2 Vec2, p3 Vec2, p4 Vec2, col uint32) {
	C.ImDrawList_AddQuadFilled((*C.ImDrawList)(unsafe.Pointer(p)), *(*C.ImVec2)(unsafe.Pointer(p1)), *(*C.ImVec2)(unsafe.Pointer(p2)), *(*C.ImVec2)(unsafe.Pointer(p3)), *(*C.ImVec2)(unsafe.Pointer(p4)), C.uint32_t(col))
}

// AddRect wraps ImDrawList_AddRect.
func (p DrawList) AddRect(p_min Vec2, p_max Vec2, col uint32, rounding float32, flags DrawFlags, thickness float32) {
	C.ImDrawList_AddRect((*C.ImDrawList)(unsafe.Pointer(p)), *(*C.ImVec2)(unsafe.Pointer(p_min)), *(*C.ImVec2)(unsafe.Pointer(p_max)), C.uint32_t(col), C.float(rounding), C.ImDrawFlags(flags), C.float(thickness))
}

// AddRectFilled wraps ImDrawList_AddRectFilled.
func (p DrawList) AddRectFilled(p_min Vec2, p_max Vec2, col uint32, rounding float32, flags DrawFlags) {
	C.ImDrawList_AddRectFilled((*C.ImDrawList)(unsafe.Pointer(p)), *(*C.ImVec2)(unsafe.Pointer(p_min)), *(*C.ImVec2)(unsafe.Pointer(p_max)), C.uint32_t(col), C.float(rounding), C.ImDrawFlags(flags))
}

// AddRectFilledMultiColor wraps ImDrawList_AddRectFilledMultiColor.
func (p DrawList) AddRectFilledMultiColor(p_min Vec2, p_max Vec2, col_upr_left uint32, col_upr_right uint32, col_bot_right uint32, col_bot_left uint32) {
	C.ImDrawList_AddRectFilledMultiColor((*C.ImDrawList)(unsafe.Pointer(p)), *(*C.ImVec2)(unsafe.Pointer(p_min)), *(*C.ImVec2)(unsafe.Pointer(p_max)), C.uint32_t(col_upr_left), C.uint32_t(col_upr_right), C.uint32_t(col_bot_right), C.uint32_t(col_bot_left))
}

// AddText_FontPtr wraps ImDrawList_AddText_FontPtr.
func (p DrawList) AddText_FontPtr(font Font, font_size float32, pos Vec2, col uint32, text_begin ffi.CString, text_end ffi.CString, wrap_width float32, cpu_fine_clip_rect Vec4) {
	C.ImDrawList_AddText_FontPtr((*C.ImDrawList)(unsafe.Pointer(p)), (*C.ImFont)(unsafe.Pointer(font)), C.float(font_size), *(*C.ImVec2)(unsafe.Pointer(pos)), C.uint32_t(col), (*C.char)(text_begin.Raw()), (*C.char)(text_end.Raw()), C.float(wrap_width), (*C.ImVec4)(unsafe.Pointer(cpu_fine_clip_rect)))
}

// AddText_Vec2 wraps ImDrawList_AddText_Vec2.
func (p DrawList) AddText_Vec2(pos Vec2, col uint32, text_begin ffi.CString, text_end ffi.CString) {
	C.ImDrawList_AddText_Vec2((*C.ImDrawList)(unsafe.Pointer(p)), *(*C.ImVec2)(unsafe.Pointer(pos)), C.uint32_t(col), (*C.char)(text_begin.Raw()), (*C.char)(text_end.Raw()))
}

// AddTriangle wraps ImDrawList_AddTriangle.
func (p DrawList) AddTriangle(p1 Vec2, p2 Vec2, p3 Vec2, col uint32, thickness float32) {
	C.ImDrawList_AddTriangle((*C.ImDrawList)(unsafe.Pointer(p)), *(*C.ImVec2)(unsafe.Pointer(p1)), *(*C.ImVec2)(unsafe.Pointer(p2)), *(*C.ImVec2)(unsafe.Pointer(p3)), C.uint32_t(col), C.float(thickness))
}

// AddTriangleFilled wraps ImDrawList_AddTriangleFilled.
func (p DrawList) AddTriangleFilled(p1 Vec2, p2 Vec2, p3 Vec2, col uint32) {
	C.ImDrawList_AddTriangleFilled((*C.ImDrawList)(unsafe.Pointer(p)), *(*C.ImVec2)(unsafe.Pointer(p1)), *(*C.ImVec2)(unsafe.Pointer(p2)), *(*C.ImVec2)(unsafe.Pointer(p3)), C.uint32_t(col))
}

// ChannelsMerge wraps ImDrawList_ChannelsMerge.
func (p DrawList) ChannelsMerge() {
	C.ImDrawList_ChannelsMerge((*C.ImDrawList)(unsafe.Pointer(p)))
}

// ChannelsSetCurrent wraps ImDrawList_ChannelsSetCurrent.
func (p DrawList) ChannelsSetCurrent(n int32) {
	C.ImDrawList_ChannelsSetCurrent((*C.ImDrawList)(unsafe.Pointer(p)), C.int32_t(n))
}

// ChannelsSplit wraps ImDrawList_ChannelsSplit.
func (p DrawList) ChannelsSplit(count int32) {
	C.ImDrawList_ChannelsSplit((*C.ImDrawList)(unsafe.Pointer(p)), C.int32_t(count))
}

// CloneOutput wraps ImDrawList_CloneOutput.
func (p DrawList) CloneOutput() DrawList {
	ret := C.ImDrawList_CloneOutput((*C.ImDrawList)(unsafe.Pointer(p)))

	return DrawList(unsafe.Pointer(ret))
}

// GetClipRectMax wraps ImDrawList_GetClipRectMax.
func (p DrawList) GetClipRectMax(pOut Vec2) {
	C.ImDrawList_GetClipRectMax((*C.ImVec2)(unsafe.Pointer(pOut)), (*C.ImDrawList)(unsafe.Pointer(p)))
}

// GetClipRectMin wraps ImDrawList_GetClipRectMin.
func (p DrawList) GetClipRectMin(pOut Vec2) {
	C.ImDrawList_GetClipRectMin((*C.ImVec2)(unsafe.Pointer(pOut)), (*C.ImDrawList)(unsafe.Pointer(p)))
}

// DrawListNew wraps ImDrawList_ImDrawList.
func DrawListNew(shared_data DrawListSharedData) DrawList {
	ret := C.ImDrawList_ImDrawList((*C.ImDrawListSharedData)(unsafe.Pointer(shared_data)))

	return DrawList(unsafe.Pointer(ret))
}

// PathArcTo wraps ImDrawList_PathArcTo.
func (p DrawList) PathArcTo(center Vec2, radius float32, a_min float32, a_max float32, num_segments int32) {
	C.ImDrawList_PathArcTo((*C.ImDrawList)(unsafe.Pointer(p)), *(*C.ImVec2)(unsafe.Pointer(center)), C.float(radius), C.float(a_min), C.float(a_max), C.int32_t(num_segments))
}

// PathArcToFast wraps ImDrawList_PathArcToFast.
func (p DrawList) PathArcToFast(center Vec2, radius float32, a_min_of_12 int32, a_max_of_12 int32) {
	C.ImDrawList_PathArcToFast((*C.ImDrawList)(unsafe.Pointer(p)), *(*C.ImVec2)(unsafe.Pointer(center)), C.float(radius), C.int32_t(a_min_of_12), C.int32_t(a_max_of_12))
}

// PathBezierCubicCurveTo wraps ImDrawList_PathBezierCubicCurveTo.
func (p DrawList) PathBezierCubicCurveTo(p2 Vec2, p3 Vec2, p4 Vec2, num_segments int32) {
	C.ImDrawList_PathBezierCubicCurveTo((*C.ImDrawList)(unsafe.Pointer(p)), *(*C.ImVec2)(unsafe.Pointer(p2)), *(*C.ImVec2)(unsafe.Pointer(p3)), *(*C.ImVec2)(unsafe.Pointer(p4)), C.int32_t(num_segments))
}

// PathBezierQuadraticCurveTo wraps ImDrawList_PathBezierQuadraticCurveTo.
func (p DrawList) PathBezierQuadraticCurveTo(p2 Vec2, p3 Vec2, num_segments int32) {
	C.ImDrawList_PathBezierQuadraticCurveTo((*C.ImDrawList)(unsafe.Pointer(p)), *(*C.ImVec2)(unsafe.Pointer(p2)), *(*C.ImVec2)(unsafe.Pointer(p3)), C.int32_t(num_segments))
}

// PathClear wraps ImDrawList_PathClear.
func (p DrawList) PathClear() {
	C.ImDrawList_PathClear((*C.ImDrawList)(unsafe.Pointer(p)))
}

// PathEllipticalArcTo wraps ImDrawList_PathEllipticalArcTo.
func (p DrawList) PathEllipticalArcTo(center Vec2, radius Vec2, rot float32, a_min float32, a_max float32, num_segments int32) {
	C.ImDrawList_PathEllipticalArcTo((*C.ImDrawList)(unsafe.Pointer(p)), *(*C.ImVec2)(unsafe.Pointer(center)), *(*C.ImVec2)(unsafe.Pointer(radius)), C.float(rot), C.float(a_min), C.float(a_max), C.int32_t(num_segments))
}

// PathFillConcave wraps ImDrawList_PathFillConcave.
func (p DrawList) PathFillConcave(col uint32) {
	C.ImDrawList_PathFillConcave((*C.ImDrawList)(unsafe.Pointer(p)), C.uint32_t(col))
}

// PathFillConvex wraps ImDrawList_PathFillConvex.
func (p DrawList) PathFillConvex(col uint32) {
	C.ImDrawList_PathFillConvex((*C.ImDrawList)(unsafe.Pointer(p)), C.uint32_t(col))
}

// PathLineTo wraps ImDrawList_PathLineTo.
func (p DrawList) PathLineTo(pos Vec2) {
	C.ImDrawList_PathLineTo((*C.ImDrawList)(unsafe.Pointer(p)), *(*C.ImVec2)(unsafe.Pointer(pos)))
}

// PathLineToMergeDuplicate wraps ImDrawList_PathLineToMergeDuplicate.
func (p DrawList) PathLineToMergeDuplicate(pos Vec2) {
	C.ImDrawList_PathLineToMergeDuplicate((*C.ImDrawList)(unsafe.Pointer(p)), *(*C.ImVec2)(unsafe.Pointer(pos)))
}

// PathRect wraps ImDrawList_PathRect.
func (p DrawList) PathRect(rect_min Vec2, rect_max Vec2, rounding float32, flags DrawFlags) {
	C.ImDrawList_PathRect((*C.ImDrawList)(unsafe.Pointer(p)), *(*C.ImVec2)(unsafe.Pointer(rect_min)), *(*C.ImVec2)(unsafe.Pointer(rect_max)), C.float(rounding), C.ImDrawFlags(flags))
}

// PathStroke wraps ImDrawList_PathStroke.
func (p DrawList) PathStroke(col uint32, flags DrawFlags, thickness float32) {
	C.ImDrawList_PathStroke((*C.ImDrawList)(unsafe.Pointer(p)), C.uint32_t(col), C.ImDrawFlags(flags), C.float(thickness))
}

// PopClipRect wraps ImDrawList_PopClipRect.
func (p DrawList) PopClipRect() {
	C.ImDrawList_PopClipRect((*C.ImDrawList)(unsafe.Pointer(p)))
}

// PopTexture wraps ImDrawList_PopTexture.
func (p DrawList) PopTexture() {
	C.ImDrawList_PopTexture((*C.ImDrawList)(unsafe.Pointer(p)))
}

// PrimQuadUV wraps ImDrawList_PrimQuadUV.
func (p DrawList) PrimQuadUV(a Vec2, b Vec2, c Vec2, d Vec2, uv_a Vec2, uv_b Vec2, uv_c Vec2, uv_d Vec2, col uint32) {
	C.ImDrawList_PrimQuadUV((*C.ImDrawList)(unsafe.Pointer(p)), *(*C.ImVec2)(unsafe.Pointer(a)), *(*C.ImVec2)(unsafe.Pointer(b)), *(*C.ImVec2)(unsafe.Pointer(c)), *(*C.ImVec2)(unsafe.Pointer(d)), *(*C.ImVec2)(unsafe.Pointer(uv_a)), *(*C.ImVec2)(unsafe.Pointer(uv_b)), *(*C.ImVec2)(unsafe.Pointer(uv_c)), *(*C.ImVec2)(unsafe.Pointer(uv_d)), C.uint32_t(col))
}

// PrimRect wraps ImDrawList_PrimRect.
func (p DrawList) PrimRect(a Vec2, b Vec2, col uint32) {
	C.ImDrawList_PrimRect((*C.ImDrawList)(unsafe.Pointer(p)), *(*C.ImVec2)(unsafe.Pointer(a)), *(*C.ImVec2)(unsafe.Pointer(b)), C.uint32_t(col))
}

// PrimRectUV wraps ImDrawList_PrimRectUV.
func (p DrawList) PrimRectUV(a Vec2, b Vec2, uv_a Vec2, uv_b Vec2, col uint32) {
	C.ImDrawList_PrimRectUV((*C.ImDrawList)(unsafe.Pointer(p)), *(*C.ImVec2)(unsafe.Pointer(a)), *(*C.ImVec2)(unsafe.Pointer(b)), *(*C.ImVec2)(unsafe.Pointer(uv_a)), *(*C.ImVec2)(unsafe.Pointer(uv_b)), C.uint32_t(col))
}

// PrimReserve wraps ImDrawList_PrimReserve.
func (p DrawList) PrimReserve(idx_count int32, vtx_count int32) {
	C.ImDrawList_PrimReserve((*C.ImDrawList)(unsafe.Pointer(p)), C.int32_t(idx_count), C.int32_t(vtx_count))
}

// PrimUnreserve wraps ImDrawList_PrimUnreserve.
func (p DrawList) PrimUnreserve(idx_count int32, vtx_count int32) {
	C.ImDrawList_PrimUnreserve((*C.ImDrawList)(unsafe.Pointer(p)), C.int32_t(idx_count), C.int32_t(vtx_count))
}

// PrimVtx wraps ImDrawList_PrimVtx.
func (p DrawList) PrimVtx(pos Vec2, uv Vec2, col uint32) {
	C.ImDrawList_PrimVtx((*C.ImDrawList)(unsafe.Pointer(p)), *(*C.ImVec2)(unsafe.Pointer(pos)), *(*C.ImVec2)(unsafe.Pointer(uv)), C.uint32_t(col))
}

// PrimWriteIdx wraps ImDrawList_PrimWriteIdx.
func (p DrawList) PrimWriteIdx(idx DrawIdx) {
	C.ImDrawList_PrimWriteIdx((*C.ImDrawList)(unsafe.Pointer(p)), C.ImDrawIdx(idx))
}

// PrimWriteVtx wraps ImDrawList_PrimWriteVtx.
func (p DrawList) PrimWriteVtx(pos Vec2, uv Vec2, col uint32) {
	C.ImDrawList_PrimWriteVtx((*C.ImDrawList)(unsafe.Pointer(p)), *(*C.ImVec2)(unsafe.Pointer(pos)), *(*C.ImVec2)(unsafe.Pointer(uv)), C.uint32_t(col))
}

// PushClipRect wraps ImDrawList_PushClipRect.
func (p DrawList) PushClipRect(clip_rect_min Vec2, clip_rect_max Vec2, intersect_with_current_clip_rect bool) {
	C.ImDrawList_PushClipRect((*C.ImDrawList)(unsafe.Pointer(p)), *(*C.ImVec2)(unsafe.Pointer(clip_rect_min)), *(*C.ImVec2)(unsafe.Pointer(clip_rect_max)), C.bool(intersect_with_current_clip_rect))
}

// PushClipRectFullScreen wraps ImDrawList_PushClipRectFullScreen.
func (p DrawList) PushClipRectFullScreen() {
	C.ImDrawList_PushClipRectFullScreen((*C.ImDrawList)(unsafe.Pointer(p)))
}

// PushTexture wraps ImDrawList_PushTexture.
func (p DrawList) PushTexture(tex_ref TextureRef) {
	C.ImDrawList_PushTexture((*C.ImDrawList)(unsafe.Pointer(p)), *(*C.ImTextureRef)(unsafe.Pointer(tex_ref)))
}

// _CalcCircleAutoSegmentCount wraps ImDrawList__CalcCircleAutoSegmentCount.
func (p DrawList) _CalcCircleAutoSegmentCount(radius float32) int32 {
	ret := C.ImDrawList__CalcCircleAutoSegmentCount((*C.ImDrawList)(unsafe.Pointer(p)), C.float(radius))

	return int32(ret)
}

// _ClearFreeMemory wraps ImDrawList__ClearFreeMemory.
func (p DrawList) _ClearFreeMemory() {
	C.ImDrawList__ClearFreeMemory((*C.ImDrawList)(unsafe.Pointer(p)))
}

// _OnChangedClipRect wraps ImDrawList__OnChangedClipRect.
func (p DrawList) _OnChangedClipRect() {
	C.ImDrawList__OnChangedClipRect((*C.ImDrawList)(unsafe.Pointer(p)))
}

// _OnChangedTexture wraps ImDrawList__OnChangedTexture.
func (p DrawList) _OnChangedTexture() {
	C.ImDrawList__OnChangedTexture((*C.ImDrawList)(unsafe.Pointer(p)))
}

// _OnChangedVtxOffset wraps ImDrawList__OnChangedVtxOffset.
func (p DrawList) _OnChangedVtxOffset() {
	C.ImDrawList__OnChangedVtxOffset((*C.ImDrawList)(unsafe.Pointer(p)))
}

// _PathArcToFastEx wraps ImDrawList__PathArcToFastEx.
func (p DrawList) _PathArcToFastEx(center Vec2, radius float32, a_min_sample int32, a_max_sample int32, a_step int32) {
	C.ImDrawList__PathArcToFastEx((*C.ImDrawList)(unsafe.Pointer(p)), *(*C.ImVec2)(unsafe.Pointer(center)), C.float(radius), C.int32_t(a_min_sample), C.int32_t(a_max_sample), C.int32_t(a_step))
}

// _PathArcToN wraps ImDrawList__PathArcToN.
func (p DrawList) _PathArcToN(center Vec2, radius float32, a_min float32, a_max float32, num_segments int32) {
	C.ImDrawList__PathArcToN((*C.ImDrawList)(unsafe.Pointer(p)), *(*C.ImVec2)(unsafe.Pointer(center)), C.float(radius), C.float(a_min), C.float(a_max), C.int32_t(num_segments))
}

// _PopUnusedDrawCmd wraps ImDrawList__PopUnusedDrawCmd.
func (p DrawList) _PopUnusedDrawCmd() {
	C.ImDrawList__PopUnusedDrawCmd((*C.ImDrawList)(unsafe.Pointer(p)))
}

// _ResetForNewFrame wraps ImDrawList__ResetForNewFrame.
func (p DrawList) _ResetForNewFrame() {
	C.ImDrawList__ResetForNewFrame((*C.ImDrawList)(unsafe.Pointer(p)))
}

// _SetDrawListSharedData wraps ImDrawList__SetDrawListSharedData.
func (p DrawList) _SetDrawListSharedData(data DrawListSharedData) {
	C.ImDrawList__SetDrawListSharedData((*C.ImDrawList)(unsafe.Pointer(p)), (*C.ImDrawListSharedData)(unsafe.Pointer(data)))
}

// _SetTexture wraps ImDrawList__SetTexture.
func (p DrawList) _SetTexture(tex_ref TextureRef) {
	C.ImDrawList__SetTexture((*C.ImDrawList)(unsafe.Pointer(p)), *(*C.ImTextureRef)(unsafe.Pointer(tex_ref)))
}

// _TryMergeDrawCmds wraps ImDrawList__TryMergeDrawCmds.
func (p DrawList) _TryMergeDrawCmds() {
	C.ImDrawList__TryMergeDrawCmds((*C.ImDrawList)(unsafe.Pointer(p)))
}

// Destroy wraps ImDrawList_destroy.
func (p DrawList) Destroy() {
	C.ImDrawList_destroy((*C.ImDrawList)(unsafe.Pointer(p)))
}

// FontAtlasRectNew wraps ImFontAtlasRect_ImFontAtlasRect.
func FontAtlasRectNew() FontAtlasRect {
	ret := C.ImFontAtlasRect_ImFontAtlasRect()

	return FontAtlasRect(unsafe.Pointer(ret))
}

// Destroy wraps ImFontAtlasRect_destroy.
func (p FontAtlasRect) Destroy() {
	C.ImFontAtlasRect_destroy((*C.ImFontAtlasRect)(unsafe.Pointer(p)))
}

// AddCustomRect wraps ImFontAtlas_AddCustomRect.
func (p FontAtlas) AddCustomRect(width int32, height int32, out_r FontAtlasRect) FontAtlasRectId {
	ret := C.ImFontAtlas_AddCustomRect((*C.ImFontAtlas)(unsafe.Pointer(p)), C.int32_t(width), C.int32_t(height), (*C.ImFontAtlasRect)(unsafe.Pointer(out_r)))

	return FontAtlasRectId(ret)
}

// AddFont wraps ImFontAtlas_AddFont.
func (p FontAtlas) AddFont(font_cfg FontConfig) Font {
	ret := C.ImFontAtlas_AddFont((*C.ImFontAtlas)(unsafe.Pointer(p)), (*C.ImFontConfig)(unsafe.Pointer(font_cfg)))

	return Font(unsafe.Pointer(ret))
}

// AddFontDefault wraps ImFontAtlas_AddFontDefault.
func (p FontAtlas) AddFontDefault(font_cfg FontConfig) Font {
	ret := C.ImFontAtlas_AddFontDefault((*C.ImFontAtlas)(unsafe.Pointer(p)), (*C.ImFontConfig)(unsafe.Pointer(font_cfg)))

	return Font(unsafe.Pointer(ret))
}

// ImFontAtlas_AddFontFromFileTTF.glyph_ranges is unsupported: category unsupported.

// ImFontAtlas_AddFontFromMemoryCompressedBase85TTF.glyph_ranges is unsupported: category unsupported.

// ImFontAtlas_AddFontFromMemoryCompressedTTF.glyph_ranges is unsupported: category unsupported.

// ImFontAtlas_AddFontFromMemoryTTF.glyph_ranges is unsupported: category unsupported.

// Clear wraps ImFontAtlas_Clear.
func (p FontAtlas) Clear() {
	C.ImFontAtlas_Clear((*C.ImFontAtlas)(unsafe.Pointer(p)))
}

// ClearFonts wraps ImFontAtlas_ClearFonts.
func (p FontAtlas) ClearFonts() {
	C.ImFontAtlas_ClearFonts((*C.ImFontAtlas)(unsafe.Pointer(p)))
}

// ClearInputData wraps ImFontAtlas_ClearInputData.
func (p FontAtlas) ClearInputData() {
	C.ImFontAtlas_ClearInputData((*C.ImFontAtlas)(unsafe.Pointer(p)))
}

// ClearTexData wraps ImFontAtlas_ClearTexData.
func (p FontAtlas) ClearTexData() {
	C.ImFontAtlas_ClearTexData((*C.ImFontAtlas)(unsafe.Pointer(p)))
}

// CompactCache wraps ImFontAtlas_CompactCache.
func (p FontAtlas) CompactCache() {
	C.ImFontAtlas_CompactCache((*C.ImFontAtlas)(unsafe.Pointer(p)))
}

// GetCustomRect wraps ImFontAtlas_GetCustomRect.
func (p FontAtlas) GetCustomRect(id FontAtlasRectId, out_r FontAtlasRect) bool {
	ret := C.ImFontAtlas_GetCustomRect((*C.ImFontAtlas)(unsafe.Pointer(p)), C.ImFontAtlasRectId(id), (*C.ImFontAtlasRect)(unsafe.Pointer(out_r)))

	return bool(ret)
}

// ImFontAtlas_GetGlyphRangesDefault is unsupported: unknown return category unsupported.

// FontAtlasNew wraps ImFontAtlas_ImFontAtlas.
func FontAtlasNew() FontAtlas {
	ret := C.ImFontAtlas_ImFontAtlas()

	return FontAtlas(unsafe.Pointer(ret))
}

// RemoveCustomRect wraps ImFontAtlas_RemoveCustomRect.
func (p FontAtlas) RemoveCustomRect(id FontAtlasRectId) {
	C.ImFontAtlas_RemoveCustomRect((*C.ImFontAtlas)(unsafe.Pointer(p)), C.ImFontAtlasRectId(id))
}

// RemoveFont wraps ImFontAtlas_RemoveFont.
func (p FontAtlas) RemoveFont(font Font) {
	C.ImFontAtlas_RemoveFont((*C.ImFontAtlas)(unsafe.Pointer(p)), (*C.ImFont)(unsafe.Pointer(font)))
}

// SetFontLoader wraps ImFontAtlas_SetFontLoader.
func (p FontAtlas) SetFontLoader(font_loader FontLoader) {
	C.ImFontAtlas_SetFontLoader((*C.ImFontAtlas)(unsafe.Pointer(p)), (*C.ImFontLoader)(unsafe.Pointer(font_loader)))
}

// Destroy wraps ImFontAtlas_destroy.
func (p FontAtlas) Destroy() {
	C.ImFontAtlas_destroy((*C.ImFontAtlas)(unsafe.Pointer(p)))
}

// ClearOutputData wraps ImFontBaked_ClearOutputData.
func (p FontBaked) ClearOutputData() {
	C.ImFontBaked_ClearOutputData((*C.ImFontBaked)(unsafe.Pointer(p)))
}

// ImFontBaked_FindGlyph.c is unsupported: category unsupported.

// ImFontBaked_FindGlyphNoFallback.c is unsupported: category unsupported.

// ImFontBaked_GetCharAdvance.c is unsupported: category unsupported.

// FontBakedNew wraps ImFontBaked_ImFontBaked.
func FontBakedNew() FontBaked {
	ret := C.ImFontBaked_ImFontBaked()

	return FontBaked(unsafe.Pointer(ret))
}

// ImFontBaked_IsGlyphLoaded.c is unsupported: category unsupported.

// Destroy wraps ImFontBaked_destroy.
func (p FontBaked) Destroy() {
	C.ImFontBaked_destroy((*C.ImFontBaked)(unsafe.Pointer(p)))
}

// FontConfigNew wraps ImFontConfig_ImFontConfig.
func FontConfigNew() FontConfig {
	ret := C.ImFontConfig_ImFontConfig()

	return FontConfig(unsafe.Pointer(ret))
}

// Destroy wraps ImFontConfig_destroy.
func (p FontConfig) Destroy() {
	C.ImFontConfig_destroy((*C.ImFontConfig)(unsafe.Pointer(p)))
}

// ImFontGlyphRangesBuilder_AddChar.c is unsupported: category unsupported.

// ImFontGlyphRangesBuilder_AddRanges.ranges is unsupported: category unsupported.

// AddText wraps ImFontGlyphRangesBuilder_AddText.
func (p FontGlyphRangesBuilder) AddText(text ffi.CString, text_end ffi.CString) {
	C.ImFontGlyphRangesBuilder_AddText((*C.ImFontGlyphRangesBuilder)(unsafe.Pointer(p)), (*C.char)(text.Raw()), (*C.char)(text_end.Raw()))
}

// BuildRanges wraps ImFontGlyphRangesBuilder_BuildRanges.
func (p FontGlyphRangesBuilder) BuildRanges(out_ranges Vector_ImWchar) {
	C.ImFontGlyphRangesBuilder_BuildRanges((*C.ImFontGlyphRangesBuilder)(unsafe.Pointer(p)), (*C.ImVector_ImWchar)(unsafe.Pointer(out_ranges)))
}

// Clear wraps ImFontGlyphRangesBuilder_Clear.
func (p FontGlyphRangesBuilder) Clear() {
	C.ImFontGlyphRangesBuilder_Clear((*C.ImFontGlyphRangesBuilder)(unsafe.Pointer(p)))
}

// ImFontGlyphRangesBuilder_GetBit.n is unsupported: category unsupported.

// FontGlyphRangesBuilderNew wraps ImFontGlyphRangesBuilder_ImFontGlyphRangesBuilder.
func FontGlyphRangesBuilderNew() FontGlyphRangesBuilder {
	ret := C.ImFontGlyphRangesBuilder_ImFontGlyphRangesBuilder()

	return FontGlyphRangesBuilder(unsafe.Pointer(ret))
}

// ImFontGlyphRangesBuilder_SetBit.n is unsupported: category unsupported.

// Destroy wraps ImFontGlyphRangesBuilder_destroy.
func (p FontGlyphRangesBuilder) Destroy() {
	C.ImFontGlyphRangesBuilder_destroy((*C.ImFontGlyphRangesBuilder)(unsafe.Pointer(p)))
}

// FontGlyphNew wraps ImFontGlyph_ImFontGlyph.
func FontGlyphNew() FontGlyph {
	ret := C.ImFontGlyph_ImFontGlyph()

	return FontGlyph(unsafe.Pointer(ret))
}

// Destroy wraps ImFontGlyph_destroy.
func (p FontGlyph) Destroy() {
	C.ImFontGlyph_destroy((*C.ImFontGlyph)(unsafe.Pointer(p)))
}

// ImFont_AddRemapChar.from_codepoint is unsupported: category unsupported.

// CalcTextSizeA wraps ImFont_CalcTextSizeA.
func (p Font) CalcTextSizeA(pOut Vec2, size float32, max_width float32, wrap_width float32, text_begin ffi.CString, text_end ffi.CString, out_remaining ffi.Ref[ffi.CString]) {
	C.ImFont_CalcTextSizeA((*C.ImVec2)(unsafe.Pointer(pOut)), (*C.ImFont)(unsafe.Pointer(p)), C.float(size), C.float(max_width), C.float(wrap_width), (*C.char)(text_begin.Raw()), (*C.char)(text_end.Raw()), (**C.char)(out_remaining.Raw()))
}

// ImFont_CalcWordWrapPosition is unsupported: return pointer type -> ?? char.

// ClearOutputData wraps ImFont_ClearOutputData.
func (p Font) ClearOutputData() {
	C.ImFont_ClearOutputData((*C.ImFont)(unsafe.Pointer(p)))
}

// ImFont_GetDebugName is unsupported: return pointer type -> ?? char.

// GetFontBaked wraps ImFont_GetFontBaked.
func (p Font) GetFontBaked(font_size float32, density float32) FontBaked {
	ret := C.ImFont_GetFontBaked((*C.ImFont)(unsafe.Pointer(p)), C.float(font_size), C.float(density))

	return FontBaked(unsafe.Pointer(ret))
}

// FontNew wraps ImFont_ImFont.
func FontNew() Font {
	ret := C.ImFont_ImFont()

	return Font(unsafe.Pointer(ret))
}

// ImFont_IsGlyphInFont.c is unsupported: category unsupported.

// IsGlyphRangeUnused wraps ImFont_IsGlyphRangeUnused.
func (p Font) IsGlyphRangeUnused(c_begin uint32, c_last uint32) bool {
	ret := C.ImFont_IsGlyphRangeUnused((*C.ImFont)(unsafe.Pointer(p)), C.uint32_t(c_begin), C.uint32_t(c_last))

	return bool(ret)
}

// IsLoaded wraps ImFont_IsLoaded.
func (p Font) IsLoaded() bool {
	ret := C.ImFont_IsLoaded((*C.ImFont)(unsafe.Pointer(p)))

	return bool(ret)
}

// ImFont_RenderChar.c is unsupported: category unsupported.

// RenderText wraps ImFont_RenderText.
func (p Font) RenderText(draw_list DrawList, size float32, pos Vec2, col uint32, clip_rect Vec4, text_begin ffi.CString, text_end ffi.CString, wrap_width float32, flags DrawTextFlags) {
	C.ImFont_RenderText((*C.ImFont)(unsafe.Pointer(p)), (*C.ImDrawList)(unsafe.Pointer(draw_list)), C.float(size), *(*C.ImVec2)(unsafe.Pointer(pos)), C.uint32_t(col), *(*C.ImVec4)(unsafe.Pointer(clip_rect)), (*C.char)(text_begin.Raw()), (*C.char)(text_end.Raw()), C.float(wrap_width), C.ImDrawTextFlags(flags))
}

// Destroy wraps ImFont_destroy.
func (p Font) Destroy() {
	C.ImFont_destroy((*C.ImFont)(unsafe.Pointer(p)))
}

// AddFocusEvent wraps ImGuiIO_AddFocusEvent.
func (p GuiIO) AddFocusEvent(focused bool) {
	C.ImGuiIO_AddFocusEvent((*C.ImGuiIO)(unsafe.Pointer(p)), C.bool(focused))
}

// AddInputCharacter wraps ImGuiIO_AddInputCharacter.
func (p GuiIO) AddInputCharacter(c uint32) {
	C.ImGuiIO_AddInputCharacter((*C.ImGuiIO)(unsafe.Pointer(p)), C.uint32_t(c))
}

// ImGuiIO_AddInputCharacterUTF16.c is unsupported: category unsupported.

// AddInputCharactersUTF8 wraps ImGuiIO_AddInputCharactersUTF8.
func (p GuiIO) AddInputCharactersUTF8(str ffi.CString) {
	C.ImGuiIO_AddInputCharactersUTF8((*C.ImGuiIO)(unsafe.Pointer(p)), (*C.char)(str.Raw()))
}

// AddKeyAnalogEvent wraps ImGuiIO_AddKeyAnalogEvent.
func (p GuiIO) AddKeyAnalogEvent(key GuiKey, down bool, v float32) {
	C.ImGuiIO_AddKeyAnalogEvent((*C.ImGuiIO)(unsafe.Pointer(p)), C.ImGuiKey(key), C.bool(down), C.float(v))
}

// AddKeyEvent wraps ImGuiIO_AddKeyEvent.
func (p GuiIO) AddKeyEvent(key GuiKey, down bool) {
	C.ImGuiIO_AddKeyEvent((*C.ImGuiIO)(unsafe.Pointer(p)), C.ImGuiKey(key), C.bool(down))
}

// AddMouseButtonEvent wraps ImGuiIO_AddMouseButtonEvent.
func (p GuiIO) AddMouseButtonEvent(button int32, down bool) {
	C.ImGuiIO_AddMouseButtonEvent((*C.ImGuiIO)(unsafe.Pointer(p)), C.int32_t(button), C.bool(down))
}

// AddMousePosEvent wraps ImGuiIO_AddMousePosEvent.
func (p GuiIO) AddMousePosEvent(x float32, y float32) {
	C.ImGuiIO_AddMousePosEvent((*C.ImGuiIO)(unsafe.Pointer(p)), C.float(x), C.float(y))
}

// AddMouseSourceEvent wraps ImGuiIO_AddMouseSourceEvent.
func (p GuiIO) AddMouseSourceEvent(source GuiMouseSource) {
	C.ImGuiIO_AddMouseSourceEvent((*C.ImGuiIO)(unsafe.Pointer(p)), C.ImGuiMouseSource(source))
}

// ImGuiIO_AddMouseViewportEvent.id is unsupported: category unsupported.

// AddMouseWheelEvent wraps ImGuiIO_AddMouseWheelEvent.
func (p GuiIO) AddMouseWheelEvent(wheel_x float32, wheel_y float32) {
	C.ImGuiIO_AddMouseWheelEvent((*C.ImGuiIO)(unsafe.Pointer(p)), C.float(wheel_x), C.float(wheel_y))
}

// ClearEventsQueue wraps ImGuiIO_ClearEventsQueue.
func (p GuiIO) ClearEventsQueue() {
	C.ImGuiIO_ClearEventsQueue((*C.ImGuiIO)(unsafe.Pointer(p)))
}

// ClearInputKeys wraps ImGuiIO_ClearInputKeys.
func (p GuiIO) ClearInputKeys() {
	C.ImGuiIO_ClearInputKeys((*C.ImGuiIO)(unsafe.Pointer(p)))
}

// ClearInputMouse wraps ImGuiIO_ClearInputMouse.
func (p GuiIO) ClearInputMouse() {
	C.ImGuiIO_ClearInputMouse((*C.ImGuiIO)(unsafe.Pointer(p)))
}

// GuiIONew wraps ImGuiIO_ImGuiIO.
func GuiIONew() GuiIO {
	ret := C.ImGuiIO_ImGuiIO()

	return GuiIO(unsafe.Pointer(ret))
}

// SetAppAcceptingEvents wraps ImGuiIO_SetAppAcceptingEvents.
func (p GuiIO) SetAppAcceptingEvents(accepting_events bool) {
	C.ImGuiIO_SetAppAcceptingEvents((*C.ImGuiIO)(unsafe.Pointer(p)), C.bool(accepting_events))
}

// SetKeyEventNativeData wraps ImGuiIO_SetKeyEventNativeData.
func (p GuiIO) SetKeyEventNativeData(key GuiKey, native_keycode int32, native_scancode int32, native_legacy_index int32) {
	C.ImGuiIO_SetKeyEventNativeData((*C.ImGuiIO)(unsafe.Pointer(p)), C.ImGuiKey(key), C.int32_t(native_keycode), C.int32_t(native_scancode), C.int32_t(native_legacy_index))
}

// Destroy wraps ImGuiIO_destroy.
func (p GuiIO) Destroy() {
	C.ImGuiIO_destroy((*C.ImGuiIO)(unsafe.Pointer(p)))
}

// ClearSelection wraps ImGuiInputTextCallbackData_ClearSelection.
func (p GuiInputTextCallbackData) ClearSelection() {
	C.ImGuiInputTextCallbackData_ClearSelection((*C.ImGuiInputTextCallbackData)(unsafe.Pointer(p)))
}

// DeleteChars wraps ImGuiInputTextCallbackData_DeleteChars.
func (p GuiInputTextCallbackData) DeleteChars(pos int32, bytes_count int32) {
	C.ImGuiInputTextCallbackData_DeleteChars((*C.ImGuiInputTextCallbackData)(unsafe.Pointer(p)), C.int32_t(pos), C.int32_t(bytes_count))
}

// HasSelection wraps ImGuiInputTextCallbackData_HasSelection.
func (p GuiInputTextCallbackData) HasSelection() bool {
	ret := C.ImGuiInputTextCallbackData_HasSelection((*C.ImGuiInputTextCallbackData)(unsafe.Pointer(p)))

	return bool(ret)
}

// GuiInputTextCallbackDataNew wraps ImGuiInputTextCallbackData_ImGuiInputTextCallbackData.
func GuiInputTextCallbackDataNew() GuiInputTextCallbackData {
	ret := C.ImGuiInputTextCallbackData_ImGuiInputTextCallbackData()

	return GuiInputTextCallbackData(unsafe.Pointer(ret))
}

// InsertChars wraps ImGuiInputTextCallbackData_InsertChars.
func (p GuiInputTextCallbackData) InsertChars(pos int32, text ffi.CString, text_end ffi.CString) {
	C.ImGuiInputTextCallbackData_InsertChars((*C.ImGuiInputTextCallbackData)(unsafe.Pointer(p)), C.int32_t(pos), (*C.char)(text.Raw()), (*C.char)(text_end.Raw()))
}

// SelectAll wraps ImGuiInputTextCallbackData_SelectAll.
func (p GuiInputTextCallbackData) SelectAll() {
	C.ImGuiInputTextCallbackData_SelectAll((*C.ImGuiInputTextCallbackData)(unsafe.Pointer(p)))
}

// Destroy wraps ImGuiInputTextCallbackData_destroy.
func (p GuiInputTextCallbackData) Destroy() {
	C.ImGuiInputTextCallbackData_destroy((*C.ImGuiInputTextCallbackData)(unsafe.Pointer(p)))
}

// Begin wraps ImGuiListClipper_Begin.
func (p GuiListClipper) Begin(items_count int32, items_height float32) {
	C.ImGuiListClipper_Begin((*C.ImGuiListClipper)(unsafe.Pointer(p)), C.int32_t(items_count), C.float(items_height))
}

// End wraps ImGuiListClipper_End.
func (p GuiListClipper) End() {
	C.ImGuiListClipper_End((*C.ImGuiListClipper)(unsafe.Pointer(p)))
}

// GuiListClipperNew wraps ImGuiListClipper_ImGuiListClipper.
func GuiListClipperNew() GuiListClipper {
	ret := C.ImGuiListClipper_ImGuiListClipper()

	return GuiListClipper(unsafe.Pointer(ret))
}

// IncludeItemByIndex wraps ImGuiListClipper_IncludeItemByIndex.
func (p GuiListClipper) IncludeItemByIndex(item_index int32) {
	C.ImGuiListClipper_IncludeItemByIndex((*C.ImGuiListClipper)(unsafe.Pointer(p)), C.int32_t(item_index))
}

// IncludeItemsByIndex wraps ImGuiListClipper_IncludeItemsByIndex.
func (p GuiListClipper) IncludeItemsByIndex(item_begin int32, item_end int32) {
	C.ImGuiListClipper_IncludeItemsByIndex((*C.ImGuiListClipper)(unsafe.Pointer(p)), C.int32_t(item_begin), C.int32_t(item_end))
}

// SeekCursorForItem wraps ImGuiListClipper_SeekCursorForItem.
func (p GuiListClipper) SeekCursorForItem(item_index int32) {
	C.ImGuiListClipper_SeekCursorForItem((*C.ImGuiListClipper)(unsafe.Pointer(p)), C.int32_t(item_index))
}

// Step wraps ImGuiListClipper_Step.
func (p GuiListClipper) Step() bool {
	ret := C.ImGuiListClipper_Step((*C.ImGuiListClipper)(unsafe.Pointer(p)))

	return bool(ret)
}

// Destroy wraps ImGuiListClipper_destroy.
func (p GuiListClipper) Destroy() {
	C.ImGuiListClipper_destroy((*C.ImGuiListClipper)(unsafe.Pointer(p)))
}

// GuiOnceUponAFrameNew wraps ImGuiOnceUponAFrame_ImGuiOnceUponAFrame.
func GuiOnceUponAFrameNew() GuiOnceUponAFrame {
	ret := C.ImGuiOnceUponAFrame_ImGuiOnceUponAFrame()

	return GuiOnceUponAFrame(unsafe.Pointer(ret))
}

// Destroy wraps ImGuiOnceUponAFrame_destroy.
func (p GuiOnceUponAFrame) Destroy() {
	C.ImGuiOnceUponAFrame_destroy((*C.ImGuiOnceUponAFrame)(unsafe.Pointer(p)))
}

// Clear wraps ImGuiPayload_Clear.
func (p GuiPayload) Clear() {
	C.ImGuiPayload_Clear((*C.ImGuiPayload)(unsafe.Pointer(p)))
}

// GuiPayloadNew wraps ImGuiPayload_ImGuiPayload.
func GuiPayloadNew() GuiPayload {
	ret := C.ImGuiPayload_ImGuiPayload()

	return GuiPayload(unsafe.Pointer(ret))
}

// IsDataType wraps ImGuiPayload_IsDataType.
func (p GuiPayload) IsDataType(type_ ffi.CString) bool {
	ret := C.ImGuiPayload_IsDataType((*C.ImGuiPayload)(unsafe.Pointer(p)), (*C.char)(type_.Raw()))

	return bool(ret)
}

// IsDelivery wraps ImGuiPayload_IsDelivery.
func (p GuiPayload) IsDelivery() bool {
	ret := C.ImGuiPayload_IsDelivery((*C.ImGuiPayload)(unsafe.Pointer(p)))

	return bool(ret)
}

// IsPreview wraps ImGuiPayload_IsPreview.
func (p GuiPayload) IsPreview() bool {
	ret := C.ImGuiPayload_IsPreview((*C.ImGuiPayload)(unsafe.Pointer(p)))

	return bool(ret)
}

// Destroy wraps ImGuiPayload_destroy.
func (p GuiPayload) Destroy() {
	C.ImGuiPayload_destroy((*C.ImGuiPayload)(unsafe.Pointer(p)))
}

// ClearPlatformHandlers wraps ImGuiPlatformIO_ClearPlatformHandlers.
func (p GuiPlatformIO) ClearPlatformHandlers() {
	C.ImGuiPlatformIO_ClearPlatformHandlers((*C.ImGuiPlatformIO)(unsafe.Pointer(p)))
}

// ClearRendererHandlers wraps ImGuiPlatformIO_ClearRendererHandlers.
func (p GuiPlatformIO) ClearRendererHandlers() {
	C.ImGuiPlatformIO_ClearRendererHandlers((*C.ImGuiPlatformIO)(unsafe.Pointer(p)))
}

// GuiPlatformIONew wraps ImGuiPlatformIO_ImGuiPlatformIO.
func GuiPlatformIONew() GuiPlatformIO {
	ret := C.ImGuiPlatformIO_ImGuiPlatformIO()

	return GuiPlatformIO(unsafe.Pointer(ret))
}

// Destroy wraps ImGuiPlatformIO_destroy.
func (p GuiPlatformIO) Destroy() {
	C.ImGuiPlatformIO_destroy((*C.ImGuiPlatformIO)(unsafe.Pointer(p)))
}

// GuiPlatformImeDataNew wraps ImGuiPlatformImeData_ImGuiPlatformImeData.
func GuiPlatformImeDataNew() GuiPlatformImeData {
	ret := C.ImGuiPlatformImeData_ImGuiPlatformImeData()

	return GuiPlatformImeData(unsafe.Pointer(ret))
}

// Destroy wraps ImGuiPlatformImeData_destroy.
func (p GuiPlatformImeData) Destroy() {
	C.ImGuiPlatformImeData_destroy((*C.ImGuiPlatformImeData)(unsafe.Pointer(p)))
}

// GuiPlatformMonitorNew wraps ImGuiPlatformMonitor_ImGuiPlatformMonitor.
func GuiPlatformMonitorNew() GuiPlatformMonitor {
	ret := C.ImGuiPlatformMonitor_ImGuiPlatformMonitor()

	return GuiPlatformMonitor(unsafe.Pointer(ret))
}

// Destroy wraps ImGuiPlatformMonitor_destroy.
func (p GuiPlatformMonitor) Destroy() {
	C.ImGuiPlatformMonitor_destroy((*C.ImGuiPlatformMonitor)(unsafe.Pointer(p)))
}

// ApplyRequests wraps ImGuiSelectionBasicStorage_ApplyRequests.
func (p GuiSelectionBasicStorage) ApplyRequests(ms_io GuiMultiSelectIO) {
	C.ImGuiSelectionBasicStorage_ApplyRequests((*C.ImGuiSelectionBasicStorage)(unsafe.Pointer(p)), (*C.ImGuiMultiSelectIO)(unsafe.Pointer(ms_io)))
}

// Clear wraps ImGuiSelectionBasicStorage_Clear.
func (p GuiSelectionBasicStorage) Clear() {
	C.ImGuiSelectionBasicStorage_Clear((*C.ImGuiSelectionBasicStorage)(unsafe.Pointer(p)))
}

// ImGuiSelectionBasicStorage_Contains.id is unsupported: category unsupported.

// ImGuiSelectionBasicStorage_GetNextSelectedItem.out_id is unsupported: category unsupported.

// ImGuiSelectionBasicStorage_GetStorageIdFromIndex is unsupported: unknown return category unsupported.

// GuiSelectionBasicStorageNew wraps ImGuiSelectionBasicStorage_ImGuiSelectionBasicStorage.
func GuiSelectionBasicStorageNew() GuiSelectionBasicStorage {
	ret := C.ImGuiSelectionBasicStorage_ImGuiSelectionBasicStorage()

	return GuiSelectionBasicStorage(unsafe.Pointer(ret))
}

// ImGuiSelectionBasicStorage_SetItemSelected.id is unsupported: category unsupported.

// Swap wraps ImGuiSelectionBasicStorage_Swap.
func (p GuiSelectionBasicStorage) Swap(r GuiSelectionBasicStorage) {
	C.ImGuiSelectionBasicStorage_Swap((*C.ImGuiSelectionBasicStorage)(unsafe.Pointer(p)), (*C.ImGuiSelectionBasicStorage)(unsafe.Pointer(r)))
}

// Destroy wraps ImGuiSelectionBasicStorage_destroy.
func (p GuiSelectionBasicStorage) Destroy() {
	C.ImGuiSelectionBasicStorage_destroy((*C.ImGuiSelectionBasicStorage)(unsafe.Pointer(p)))
}

// ApplyRequests wraps ImGuiSelectionExternalStorage_ApplyRequests.
func (p GuiSelectionExternalStorage) ApplyRequests(ms_io GuiMultiSelectIO) {
	C.ImGuiSelectionExternalStorage_ApplyRequests((*C.ImGuiSelectionExternalStorage)(unsafe.Pointer(p)), (*C.ImGuiMultiSelectIO)(unsafe.Pointer(ms_io)))
}

// GuiSelectionExternalStorageNew wraps ImGuiSelectionExternalStorage_ImGuiSelectionExternalStorage.
func GuiSelectionExternalStorageNew() GuiSelectionExternalStorage {
	ret := C.ImGuiSelectionExternalStorage_ImGuiSelectionExternalStorage()

	return GuiSelectionExternalStorage(unsafe.Pointer(ret))
}

// Destroy wraps ImGuiSelectionExternalStorage_destroy.
func (p GuiSelectionExternalStorage) Destroy() {
	C.ImGuiSelectionExternalStorage_destroy((*C.ImGuiSelectionExternalStorage)(unsafe.Pointer(p)))
}

// ImGuiStoragePair_ImGuiStoragePair_Float._key is unsupported: category unsupported.

// ImGuiStoragePair_ImGuiStoragePair_Int._key is unsupported: category unsupported.

// ImGuiStoragePair_ImGuiStoragePair_Ptr._key is unsupported: category unsupported.

// Destroy wraps ImGuiStoragePair_destroy.
func (p GuiStoragePair) Destroy() {
	C.ImGuiStoragePair_destroy((*C.ImGuiStoragePair)(unsafe.Pointer(p)))
}

// BuildSortByKey wraps ImGuiStorage_BuildSortByKey.
func (p GuiStorage) BuildSortByKey() {
	C.ImGuiStorage_BuildSortByKey((*C.ImGuiStorage)(unsafe.Pointer(p)))
}

// Clear wraps ImGuiStorage_Clear.
func (p GuiStorage) Clear() {
	C.ImGuiStorage_Clear((*C.ImGuiStorage)(unsafe.Pointer(p)))
}

// ImGuiStorage_GetBool.key is unsupported: category unsupported.

// ImGuiStorage_GetBoolRef is unsupported: return pointer type -> ?? bool.

// ImGuiStorage_GetFloat.key is unsupported: category unsupported.

// ImGuiStorage_GetFloatRef is unsupported: return pointer type -> ?? float.

// ImGuiStorage_GetInt.key is unsupported: category unsupported.

// ImGuiStorage_GetIntRef is unsupported: return pointer type -> ?? int32_t.

// ImGuiStorage_GetVoidPtr is unsupported: return pointer type -> ?? void.

// ImGuiStorage_GetVoidPtrRef is unsupported: unknown return category pointer2.

// SetAllInt wraps ImGuiStorage_SetAllInt.
func (p GuiStorage) SetAllInt(val int32) {
	C.ImGuiStorage_SetAllInt((*C.ImGuiStorage)(unsafe.Pointer(p)), C.int32_t(val))
}

// ImGuiStorage_SetBool.key is unsupported: category unsupported.

// ImGuiStorage_SetFloat.key is unsupported: category unsupported.

// ImGuiStorage_SetInt.key is unsupported: category unsupported.

// ImGuiStorage_SetVoidPtr.key is unsupported: category unsupported.

// GuiStyleNew wraps ImGuiStyle_ImGuiStyle.
func GuiStyleNew() GuiStyle {
	ret := C.ImGuiStyle_ImGuiStyle()

	return GuiStyle(unsafe.Pointer(ret))
}

// ScaleAllSizes wraps ImGuiStyle_ScaleAllSizes.
func (p GuiStyle) ScaleAllSizes(scale_factor float32) {
	C.ImGuiStyle_ScaleAllSizes((*C.ImGuiStyle)(unsafe.Pointer(p)), C.float(scale_factor))
}

// Destroy wraps ImGuiStyle_destroy.
func (p GuiStyle) Destroy() {
	C.ImGuiStyle_destroy((*C.ImGuiStyle)(unsafe.Pointer(p)))
}

// GuiTableColumnSortSpecsNew wraps ImGuiTableColumnSortSpecs_ImGuiTableColumnSortSpecs.
func GuiTableColumnSortSpecsNew() GuiTableColumnSortSpecs {
	ret := C.ImGuiTableColumnSortSpecs_ImGuiTableColumnSortSpecs()

	return GuiTableColumnSortSpecs(unsafe.Pointer(ret))
}

// Destroy wraps ImGuiTableColumnSortSpecs_destroy.
func (p GuiTableColumnSortSpecs) Destroy() {
	C.ImGuiTableColumnSortSpecs_destroy((*C.ImGuiTableColumnSortSpecs)(unsafe.Pointer(p)))
}

// GuiTableSortSpecsNew wraps ImGuiTableSortSpecs_ImGuiTableSortSpecs.
func GuiTableSortSpecsNew() GuiTableSortSpecs {
	ret := C.ImGuiTableSortSpecs_ImGuiTableSortSpecs()

	return GuiTableSortSpecs(unsafe.Pointer(ret))
}

// Destroy wraps ImGuiTableSortSpecs_destroy.
func (p GuiTableSortSpecs) Destroy() {
	C.ImGuiTableSortSpecs_destroy((*C.ImGuiTableSortSpecs)(unsafe.Pointer(p)))
}

// GuiTextBufferNew wraps ImGuiTextBuffer_ImGuiTextBuffer.
func GuiTextBufferNew() GuiTextBuffer {
	ret := C.ImGuiTextBuffer_ImGuiTextBuffer()

	return GuiTextBuffer(unsafe.Pointer(ret))
}

// Append wraps ImGuiTextBuffer_append.
func (p GuiTextBuffer) Append(str ffi.CString, str_end ffi.CString) {
	C.ImGuiTextBuffer_append((*C.ImGuiTextBuffer)(unsafe.Pointer(p)), (*C.char)(str.Raw()), (*C.char)(str_end.Raw()))
}

// ImGuiTextBuffer_appendf.fmt is unsupported: category unsupported.

// ImGuiTextBuffer_appendfv.args is unsupported: category unsupported.

// ImGuiTextBuffer_begin is unsupported: return pointer type -> ?? char.

// ImGuiTextBuffer_c_str is unsupported: return pointer type -> ?? char.

// Clear wraps ImGuiTextBuffer_clear.
func (p GuiTextBuffer) Clear() {
	C.ImGuiTextBuffer_clear((*C.ImGuiTextBuffer)(unsafe.Pointer(p)))
}

// Destroy wraps ImGuiTextBuffer_destroy.
func (p GuiTextBuffer) Destroy() {
	C.ImGuiTextBuffer_destroy((*C.ImGuiTextBuffer)(unsafe.Pointer(p)))
}

// Empty wraps ImGuiTextBuffer_empty.
func (p GuiTextBuffer) Empty() bool {
	ret := C.ImGuiTextBuffer_empty((*C.ImGuiTextBuffer)(unsafe.Pointer(p)))

	return bool(ret)
}

// ImGuiTextBuffer_end is unsupported: return pointer type -> ?? char.

// Reserve wraps ImGuiTextBuffer_reserve.
func (p GuiTextBuffer) Reserve(capacity int32) {
	C.ImGuiTextBuffer_reserve((*C.ImGuiTextBuffer)(unsafe.Pointer(p)), C.int32_t(capacity))
}

// Resize wraps ImGuiTextBuffer_resize.
func (p GuiTextBuffer) Resize(size int32) {
	C.ImGuiTextBuffer_resize((*C.ImGuiTextBuffer)(unsafe.Pointer(p)), C.int32_t(size))
}

// Size wraps ImGuiTextBuffer_size.
func (p GuiTextBuffer) Size() int32 {
	ret := C.ImGuiTextBuffer_size((*C.ImGuiTextBuffer)(unsafe.Pointer(p)))

	return int32(ret)
}

// Build wraps ImGuiTextFilter_Build.
func (p GuiTextFilter) Build() {
	C.ImGuiTextFilter_Build((*C.ImGuiTextFilter)(unsafe.Pointer(p)))
}

// Clear wraps ImGuiTextFilter_Clear.
func (p GuiTextFilter) Clear() {
	C.ImGuiTextFilter_Clear((*C.ImGuiTextFilter)(unsafe.Pointer(p)))
}

// Draw wraps ImGuiTextFilter_Draw.
func (p GuiTextFilter) Draw(label ffi.CString, width float32) bool {
	ret := C.ImGuiTextFilter_Draw((*C.ImGuiTextFilter)(unsafe.Pointer(p)), (*C.char)(label.Raw()), C.float(width))

	return bool(ret)
}

// GuiTextFilterNew wraps ImGuiTextFilter_ImGuiTextFilter.
func GuiTextFilterNew(default_filter ffi.CString) GuiTextFilter {
	ret := C.ImGuiTextFilter_ImGuiTextFilter((*C.char)(default_filter.Raw()))

	return GuiTextFilter(unsafe.Pointer(ret))
}

// IsActive wraps ImGuiTextFilter_IsActive.
func (p GuiTextFilter) IsActive() bool {
	ret := C.ImGuiTextFilter_IsActive((*C.ImGuiTextFilter)(unsafe.Pointer(p)))

	return bool(ret)
}

// PassFilter wraps ImGuiTextFilter_PassFilter.
func (p GuiTextFilter) PassFilter(text ffi.CString, text_end ffi.CString) bool {
	ret := C.ImGuiTextFilter_PassFilter((*C.ImGuiTextFilter)(unsafe.Pointer(p)), (*C.char)(text.Raw()), (*C.char)(text_end.Raw()))

	return bool(ret)
}

// Destroy wraps ImGuiTextFilter_destroy.
func (p GuiTextFilter) Destroy() {
	C.ImGuiTextFilter_destroy((*C.ImGuiTextFilter)(unsafe.Pointer(p)))
}

// GuiTextRangeNew_Nil wraps ImGuiTextRange_ImGuiTextRange_Nil.
func GuiTextRangeNew_Nil() GuiTextRange {
	ret := C.ImGuiTextRange_ImGuiTextRange_Nil()

	return GuiTextRange(unsafe.Pointer(ret))
}

// GuiTextRangeNew_Str wraps ImGuiTextRange_ImGuiTextRange_Str.
func GuiTextRangeNew_Str(_b ffi.CString, _e ffi.CString) GuiTextRange {
	ret := C.ImGuiTextRange_ImGuiTextRange_Str((*C.char)(_b.Raw()), (*C.char)(_e.Raw()))

	return GuiTextRange(unsafe.Pointer(ret))
}

// Destroy wraps ImGuiTextRange_destroy.
func (p GuiTextRange) Destroy() {
	C.ImGuiTextRange_destroy((*C.ImGuiTextRange)(unsafe.Pointer(p)))
}

// Empty wraps ImGuiTextRange_empty.
func (p GuiTextRange) Empty() bool {
	ret := C.ImGuiTextRange_empty((*C.ImGuiTextRange)(unsafe.Pointer(p)))

	return bool(ret)
}

// ImGuiTextRange_split.separator is unsupported: unknown member direct type char.

// GetCenter wraps ImGuiViewport_GetCenter.
func (p GuiViewport) GetCenter(pOut Vec2) {
	C.ImGuiViewport_GetCenter((*C.ImVec2)(unsafe.Pointer(pOut)), (*C.ImGuiViewport)(unsafe.Pointer(p)))
}

// GetWorkCenter wraps ImGuiViewport_GetWorkCenter.
func (p GuiViewport) GetWorkCenter(pOut Vec2) {
	C.ImGuiViewport_GetWorkCenter((*C.ImVec2)(unsafe.Pointer(pOut)), (*C.ImGuiViewport)(unsafe.Pointer(p)))
}

// GuiViewportNew wraps ImGuiViewport_ImGuiViewport.
func GuiViewportNew() GuiViewport {
	ret := C.ImGuiViewport_ImGuiViewport()

	return GuiViewport(unsafe.Pointer(ret))
}

// Destroy wraps ImGuiViewport_destroy.
func (p GuiViewport) Destroy() {
	C.ImGuiViewport_destroy((*C.ImGuiViewport)(unsafe.Pointer(p)))
}

// GuiWindowClassNew wraps ImGuiWindowClass_ImGuiWindowClass.
func GuiWindowClassNew() GuiWindowClass {
	ret := C.ImGuiWindowClass_ImGuiWindowClass()

	return GuiWindowClass(unsafe.Pointer(ret))
}

// Destroy wraps ImGuiWindowClass_destroy.
func (p GuiWindowClass) Destroy() {
	C.ImGuiWindowClass_destroy((*C.ImGuiWindowClass)(unsafe.Pointer(p)))
}

// Create wraps ImTextureData_Create.
func (p TextureData) Create(format TextureFormat, w int32, h int32) {
	C.ImTextureData_Create((*C.ImTextureData)(unsafe.Pointer(p)), C.ImTextureFormat(format), C.int32_t(w), C.int32_t(h))
}

// DestroyPixels wraps ImTextureData_DestroyPixels.
func (p TextureData) DestroyPixels() {
	C.ImTextureData_DestroyPixels((*C.ImTextureData)(unsafe.Pointer(p)))
}

// GetPitch wraps ImTextureData_GetPitch.
func (p TextureData) GetPitch() int32 {
	ret := C.ImTextureData_GetPitch((*C.ImTextureData)(unsafe.Pointer(p)))

	return int32(ret)
}

// ImTextureData_GetPixels is unsupported: return pointer type -> ?? void.

// ImTextureData_GetPixelsAt is unsupported: return pointer type -> ?? void.

// GetSizeInBytes wraps ImTextureData_GetSizeInBytes.
func (p TextureData) GetSizeInBytes() int32 {
	ret := C.ImTextureData_GetSizeInBytes((*C.ImTextureData)(unsafe.Pointer(p)))

	return int32(ret)
}

// ImTextureData_GetTexID is unsupported: unknown return category unsupported.

// GetTexRef wraps ImTextureData_GetTexRef.
func (p TextureData) GetTexRef(pOut TextureRef) {
	C.ImTextureData_GetTexRef((*C.ImTextureRef)(unsafe.Pointer(pOut)), (*C.ImTextureData)(unsafe.Pointer(p)))
}

// TextureDataNew wraps ImTextureData_ImTextureData.
func TextureDataNew() TextureData {
	ret := C.ImTextureData_ImTextureData()

	return TextureData(unsafe.Pointer(ret))
}

// SetStatus wraps ImTextureData_SetStatus.
func (p TextureData) SetStatus(status TextureStatus) {
	C.ImTextureData_SetStatus((*C.ImTextureData)(unsafe.Pointer(p)), C.ImTextureStatus(status))
}

// ImTextureData_SetTexID.tex_id is unsupported: category unsupported.

// Destroy wraps ImTextureData_destroy.
func (p TextureData) Destroy() {
	C.ImTextureData_destroy((*C.ImTextureData)(unsafe.Pointer(p)))
}

// ImTextureRef_GetTexID is unsupported: unknown return category unsupported.

// TextureRefNew_Nil wraps ImTextureRef_ImTextureRef_Nil.
func TextureRefNew_Nil() TextureRef {
	ret := C.ImTextureRef_ImTextureRef_Nil()

	return TextureRef(unsafe.Pointer(ret))
}

// ImTextureRef_ImTextureRef_TextureID.tex_id is unsupported: category unsupported.

// Destroy wraps ImTextureRef_destroy.
func (p TextureRef) Destroy() {
	C.ImTextureRef_destroy((*C.ImTextureRef)(unsafe.Pointer(p)))
}

// Vec2New_Float wraps ImVec2_ImVec2_Float.
func Vec2New_Float(_x float32, _y float32) Vec2 {
	ret := C.ImVec2_ImVec2_Float(C.float(_x), C.float(_y))

	return Vec2(unsafe.Pointer(ret))
}

// Vec2New_Nil wraps ImVec2_ImVec2_Nil.
func Vec2New_Nil() Vec2 {
	ret := C.ImVec2_ImVec2_Nil()

	return Vec2(unsafe.Pointer(ret))
}

// Destroy wraps ImVec2_destroy.
func (p Vec2) Destroy() {
	C.ImVec2_destroy((*C.ImVec2)(unsafe.Pointer(p)))
}

// Vec4New_Float wraps ImVec4_ImVec4_Float.
func Vec4New_Float(_x float32, _y float32, _z float32, _w float32) Vec4 {
	ret := C.ImVec4_ImVec4_Float(C.float(_x), C.float(_y), C.float(_z), C.float(_w))

	return Vec4(unsafe.Pointer(ret))
}

// Vec4New_Nil wraps ImVec4_ImVec4_Nil.
func Vec4New_Nil() Vec4 {
	ret := C.ImVec4_ImVec4_Nil()

	return Vec4(unsafe.Pointer(ret))
}

// Destroy wraps ImVec4_destroy.
func (p Vec4) Destroy() {
	C.ImVec4_destroy((*C.ImVec4)(unsafe.Pointer(p)))
}

// AcceptDragDropPayload wraps igAcceptDragDropPayload.
func AcceptDragDropPayload(type_ ffi.CString, flags GuiDragDropFlags) GuiPayload {
	ret := C.igAcceptDragDropPayload((*C.char)(type_.Raw()), C.ImGuiDragDropFlags(flags))

	return GuiPayload(unsafe.Pointer(ret))
}

// AlignTextToFramePadding wraps igAlignTextToFramePadding.
func AlignTextToFramePadding() {
	C.igAlignTextToFramePadding()
}

// ArrowButton wraps igArrowButton.
func ArrowButton(str_id ffi.CString, dir GuiDir) bool {
	ret := C.igArrowButton((*C.char)(str_id.Raw()), C.ImGuiDir(dir))

	return bool(ret)
}

// Begin wraps igBegin.
func Begin(name ffi.CString, p_open ffi.Ref[bool], flags GuiWindowFlags) bool {
	ret := C.igBegin((*C.char)(name.Raw()), (*C.bool)(p_open.Raw()), C.ImGuiWindowFlags(flags))

	return bool(ret)
}

// igBeginChild_ID.id is unsupported: category unsupported.

// BeginChild_Str wraps igBeginChild_Str.
func BeginChild_Str(str_id ffi.CString, size Vec2, child_flags GuiChildFlags, window_flags GuiWindowFlags) bool {
	ret := C.igBeginChild_Str((*C.char)(str_id.Raw()), *(*C.ImVec2)(unsafe.Pointer(size)), C.ImGuiChildFlags(child_flags), C.ImGuiWindowFlags(window_flags))

	return bool(ret)
}

// BeginCombo wraps igBeginCombo.
func BeginCombo(label ffi.CString, preview_value ffi.CString, flags GuiComboFlags) bool {
	ret := C.igBeginCombo((*C.char)(label.Raw()), (*C.char)(preview_value.Raw()), C.ImGuiComboFlags(flags))

	return bool(ret)
}

// BeginDisabled wraps igBeginDisabled.
func BeginDisabled(disabled bool) {
	C.igBeginDisabled(C.bool(disabled))
}

// BeginDragDropSource wraps igBeginDragDropSource.
func BeginDragDropSource(flags GuiDragDropFlags) bool {
	ret := C.igBeginDragDropSource(C.ImGuiDragDropFlags(flags))

	return bool(ret)
}

// BeginDragDropTarget wraps igBeginDragDropTarget.
func BeginDragDropTarget() bool {
	ret := C.igBeginDragDropTarget()

	return bool(ret)
}

// BeginGroup wraps igBeginGroup.
func BeginGroup() {
	C.igBeginGroup()
}

// BeginItemTooltip wraps igBeginItemTooltip.
func BeginItemTooltip() bool {
	ret := C.igBeginItemTooltip()

	return bool(ret)
}

// BeginListBox wraps igBeginListBox.
func BeginListBox(label ffi.CString, size Vec2) bool {
	ret := C.igBeginListBox((*C.char)(label.Raw()), *(*C.ImVec2)(unsafe.Pointer(size)))

	return bool(ret)
}

// BeginMainMenuBar wraps igBeginMainMenuBar.
func BeginMainMenuBar() bool {
	ret := C.igBeginMainMenuBar()

	return bool(ret)
}

// BeginMenu wraps igBeginMenu.
func BeginMenu(label ffi.CString, enabled bool) bool {
	ret := C.igBeginMenu((*C.char)(label.Raw()), C.bool(enabled))

	return bool(ret)
}

// BeginMenuBar wraps igBeginMenuBar.
func BeginMenuBar() bool {
	ret := C.igBeginMenuBar()

	return bool(ret)
}

// BeginMultiSelect wraps igBeginMultiSelect.
func BeginMultiSelect(flags GuiMultiSelectFlags, selection_size int32, items_count int32) GuiMultiSelectIO {
	ret := C.igBeginMultiSelect(C.ImGuiMultiSelectFlags(flags), C.int32_t(selection_size), C.int32_t(items_count))

	return GuiMultiSelectIO(unsafe.Pointer(ret))
}

// BeginPopup wraps igBeginPopup.
func BeginPopup(str_id ffi.CString, flags GuiWindowFlags) bool {
	ret := C.igBeginPopup((*C.char)(str_id.Raw()), C.ImGuiWindowFlags(flags))

	return bool(ret)
}

// BeginPopupContextItem wraps igBeginPopupContextItem.
func BeginPopupContextItem(str_id ffi.CString, popup_flags GuiPopupFlags) bool {
	ret := C.igBeginPopupContextItem((*C.char)(str_id.Raw()), C.ImGuiPopupFlags(popup_flags))

	return bool(ret)
}

// BeginPopupContextVoid wraps igBeginPopupContextVoid.
func BeginPopupContextVoid(str_id ffi.CString, popup_flags GuiPopupFlags) bool {
	ret := C.igBeginPopupContextVoid((*C.char)(str_id.Raw()), C.ImGuiPopupFlags(popup_flags))

	return bool(ret)
}

// BeginPopupContextWindow wraps igBeginPopupContextWindow.
func BeginPopupContextWindow(str_id ffi.CString, popup_flags GuiPopupFlags) bool {
	ret := C.igBeginPopupContextWindow((*C.char)(str_id.Raw()), C.ImGuiPopupFlags(popup_flags))

	return bool(ret)
}

// BeginPopupModal wraps igBeginPopupModal.
func BeginPopupModal(name ffi.CString, p_open ffi.Ref[bool], flags GuiWindowFlags) bool {
	ret := C.igBeginPopupModal((*C.char)(name.Raw()), (*C.bool)(p_open.Raw()), C.ImGuiWindowFlags(flags))

	return bool(ret)
}

// BeginTabBar wraps igBeginTabBar.
func BeginTabBar(str_id ffi.CString, flags GuiTabBarFlags) bool {
	ret := C.igBeginTabBar((*C.char)(str_id.Raw()), C.ImGuiTabBarFlags(flags))

	return bool(ret)
}

// BeginTabItem wraps igBeginTabItem.
func BeginTabItem(label ffi.CString, p_open ffi.Ref[bool], flags GuiTabItemFlags) bool {
	ret := C.igBeginTabItem((*C.char)(label.Raw()), (*C.bool)(p_open.Raw()), C.ImGuiTabItemFlags(flags))

	return bool(ret)
}

// BeginTable wraps igBeginTable.
func BeginTable(str_id ffi.CString, columns int32, flags GuiTableFlags, outer_size Vec2, inner_width float32) bool {
	ret := C.igBeginTable((*C.char)(str_id.Raw()), C.int32_t(columns), C.ImGuiTableFlags(flags), *(*C.ImVec2)(unsafe.Pointer(outer_size)), C.float(inner_width))

	return bool(ret)
}

// BeginTooltip wraps igBeginTooltip.
func BeginTooltip() bool {
	ret := C.igBeginTooltip()

	return bool(ret)
}

// Bullet wraps igBullet.
func Bullet() {
	C.igBullet()
}

// igBulletText.... is unsupported: category unsupported.

// igBulletTextV.args is unsupported: category unsupported.

// Button wraps igButton.
func Button(label ffi.CString, size Vec2) bool {
	ret := C.igButton((*C.char)(label.Raw()), *(*C.ImVec2)(unsafe.Pointer(size)))

	return bool(ret)
}

// CalcItemWidth wraps igCalcItemWidth.
func CalcItemWidth() float32 {
	ret := C.igCalcItemWidth()

	return float32(ret)
}

// CalcTextSize wraps igCalcTextSize.
func CalcTextSize(pOut Vec2, text ffi.CString, text_end ffi.CString, hide_text_after_double_hash bool, wrap_width float32) {
	C.igCalcTextSize((*C.ImVec2)(unsafe.Pointer(pOut)), (*C.char)(text.Raw()), (*C.char)(text_end.Raw()), C.bool(hide_text_after_double_hash), C.float(wrap_width))
}

// Checkbox wraps igCheckbox.
func Checkbox(label ffi.CString, v ffi.Ref[bool]) bool {
	ret := C.igCheckbox((*C.char)(label.Raw()), (*C.bool)(v.Raw()))

	return bool(ret)
}

// CheckboxFlags_IntPtr wraps igCheckboxFlags_IntPtr.
func CheckboxFlags_IntPtr(label ffi.CString, flags ffi.Ref[int32], flags_value int32) bool {
	ret := C.igCheckboxFlags_IntPtr((*C.char)(label.Raw()), (*C.int32_t)(flags.Raw()), C.int32_t(flags_value))

	return bool(ret)
}

// CheckboxFlags_UintPtr wraps igCheckboxFlags_UintPtr.
func CheckboxFlags_UintPtr(label ffi.CString, flags ffi.Ref[uint32], flags_value uint32) bool {
	ret := C.igCheckboxFlags_UintPtr((*C.char)(label.Raw()), (*C.uint32_t)(flags.Raw()), C.uint32_t(flags_value))

	return bool(ret)
}

// CloseCurrentPopup wraps igCloseCurrentPopup.
func CloseCurrentPopup() {
	C.igCloseCurrentPopup()
}

// CollapsingHeader_BoolPtr wraps igCollapsingHeader_BoolPtr.
func CollapsingHeader_BoolPtr(label ffi.CString, p_visible ffi.Ref[bool], flags GuiTreeNodeFlags) bool {
	ret := C.igCollapsingHeader_BoolPtr((*C.char)(label.Raw()), (*C.bool)(p_visible.Raw()), C.ImGuiTreeNodeFlags(flags))

	return bool(ret)
}

// CollapsingHeader_TreeNodeFlags wraps igCollapsingHeader_TreeNodeFlags.
func CollapsingHeader_TreeNodeFlags(label ffi.CString, flags GuiTreeNodeFlags) bool {
	ret := C.igCollapsingHeader_TreeNodeFlags((*C.char)(label.Raw()), C.ImGuiTreeNodeFlags(flags))

	return bool(ret)
}

// ColorButton wraps igColorButton.
func ColorButton(desc_id ffi.CString, col Vec4, flags GuiColorEditFlags, size Vec2) bool {
	ret := C.igColorButton((*C.char)(desc_id.Raw()), *(*C.ImVec4)(unsafe.Pointer(col)), C.ImGuiColorEditFlags(flags), *(*C.ImVec2)(unsafe.Pointer(size)))

	return bool(ret)
}

// ColorConvertFloat4ToU32 wraps igColorConvertFloat4ToU32.
func ColorConvertFloat4ToU32(in Vec4) uint32 {
	ret := C.igColorConvertFloat4ToU32(*(*C.ImVec4)(unsafe.Pointer(in)))

	return uint32(ret)
}

// ColorConvertHSVtoRGB wraps igColorConvertHSVtoRGB.
func ColorConvertHSVtoRGB(h float32, s float32, v float32, out_r ffi.Ref[float32], out_g ffi.Ref[float32], out_b ffi.Ref[float32]) {
	C.igColorConvertHSVtoRGB(C.float(h), C.float(s), C.float(v), (*C.float)(out_r.Raw()), (*C.float)(out_g.Raw()), (*C.float)(out_b.Raw()))
}

// ColorConvertRGBtoHSV wraps igColorConvertRGBtoHSV.
func ColorConvertRGBtoHSV(r float32, g float32, b float32, out_h ffi.Ref[float32], out_s ffi.Ref[float32], out_v ffi.Ref[float32]) {
	C.igColorConvertRGBtoHSV(C.float(r), C.float(g), C.float(b), (*C.float)(out_h.Raw()), (*C.float)(out_s.Raw()), (*C.float)(out_v.Raw()))
}

// ColorConvertU32ToFloat4 wraps igColorConvertU32ToFloat4.
func ColorConvertU32ToFloat4(pOut Vec4, in uint32) {
	C.igColorConvertU32ToFloat4((*C.ImVec4)(unsafe.Pointer(pOut)), C.uint32_t(in))
}

// igColorEdit3.col is unsupported: category unsupported.

// igColorEdit4.col is unsupported: category unsupported.

// igColorPicker3.col is unsupported: category unsupported.

// igColorPicker4.col is unsupported: category unsupported.

// Columns wraps igColumns.
func Columns(count int32, id ffi.CString, borders bool) {
	C.igColumns(C.int32_t(count), (*C.char)(id.Raw()), C.bool(borders))
}

// igCombo_FnStrPtr.getter is unsupported: category unsupported.

// Combo_Str wraps igCombo_Str.
func Combo_Str(label ffi.CString, current_item ffi.Ref[int32], items_separated_by_zeros ffi.CString, popup_max_height_in_items int32) bool {
	ret := C.igCombo_Str((*C.char)(label.Raw()), (*C.int32_t)(current_item.Raw()), (*C.char)(items_separated_by_zeros.Raw()), C.int32_t(popup_max_height_in_items))

	return bool(ret)
}

// igCombo_Str_arr.items is unsupported: category unsupported.

// CreateContext wraps igCreateContext.
func CreateContext(shared_font_atlas FontAtlas) GuiContext {
	ret := C.igCreateContext((*C.ImFontAtlas)(unsafe.Pointer(shared_font_atlas)))

	return GuiContext(unsafe.Pointer(ret))
}

// igDebugCheckVersionAndDataLayout.sz_io is unsupported: category unsupported.

// DebugFlashStyleColor wraps igDebugFlashStyleColor.
func DebugFlashStyleColor(idx GuiCol) {
	C.igDebugFlashStyleColor(C.ImGuiCol(idx))
}

// igDebugLog.... is unsupported: category unsupported.

// igDebugLogV.args is unsupported: category unsupported.

// DebugStartItemPicker wraps igDebugStartItemPicker.
func DebugStartItemPicker() {
	C.igDebugStartItemPicker()
}

// DebugTextEncoding wraps igDebugTextEncoding.
func DebugTextEncoding(text ffi.CString) {
	C.igDebugTextEncoding((*C.char)(text.Raw()))
}

// DestroyContext wraps igDestroyContext.
func DestroyContext(ctx GuiContext) {
	C.igDestroyContext((*C.ImGuiContext)(unsafe.Pointer(ctx)))
}

// DestroyPlatformWindows wraps igDestroyPlatformWindows.
func DestroyPlatformWindows() {
	C.igDestroyPlatformWindows()
}

// igDockSpace is unsupported: unknown return category unsupported.

// igDockSpaceOverViewport is unsupported: unknown return category unsupported.

// DragFloat wraps igDragFloat.
func DragFloat(label ffi.CString, v ffi.Ref[float32], v_speed float32, v_min float32, v_max float32, format ffi.CString, flags GuiSliderFlags) bool {
	ret := C.igDragFloat((*C.char)(label.Raw()), (*C.float)(v.Raw()), C.float(v_speed), C.float(v_min), C.float(v_max), (*C.char)(format.Raw()), C.ImGuiSliderFlags(flags))

	return bool(ret)
}

// igDragFloat2.v is unsupported: category unsupported.

// igDragFloat3.v is unsupported: category unsupported.

// igDragFloat4.v is unsupported: category unsupported.

// DragFloatRange2 wraps igDragFloatRange2.
func DragFloatRange2(label ffi.CString, v_current_min ffi.Ref[float32], v_current_max ffi.Ref[float32], v_speed float32, v_min float32, v_max float32, format ffi.CString, format_max ffi.CString, flags GuiSliderFlags) bool {
	ret := C.igDragFloatRange2((*C.char)(label.Raw()), (*C.float)(v_current_min.Raw()), (*C.float)(v_current_max.Raw()), C.float(v_speed), C.float(v_min), C.float(v_max), (*C.char)(format.Raw()), (*C.char)(format_max.Raw()), C.ImGuiSliderFlags(flags))

	return bool(ret)
}

// DragInt wraps igDragInt.
func DragInt(label ffi.CString, v ffi.Ref[int32], v_speed float32, v_min int32, v_max int32, format ffi.CString, flags GuiSliderFlags) bool {
	ret := C.igDragInt((*C.char)(label.Raw()), (*C.int32_t)(v.Raw()), C.float(v_speed), C.int32_t(v_min), C.int32_t(v_max), (*C.char)(format.Raw()), C.ImGuiSliderFlags(flags))

	return bool(ret)
}

// igDragInt2.v is unsupported: category unsupported.

// igDragInt3.v is unsupported: category unsupported.

// igDragInt4.v is unsupported: category unsupported.

// DragIntRange2 wraps igDragIntRange2.
func DragIntRange2(label ffi.CString, v_current_min ffi.Ref[int32], v_current_max ffi.Ref[int32], v_speed float32, v_min int32, v_max int32, format ffi.CString, format_max ffi.CString, flags GuiSliderFlags) bool {
	ret := C.igDragIntRange2((*C.char)(label.Raw()), (*C.int32_t)(v_current_min.Raw()), (*C.int32_t)(v_current_max.Raw()), C.float(v_speed), C.int32_t(v_min), C.int32_t(v_max), (*C.char)(format.Raw()), (*C.char)(format_max.Raw()), C.ImGuiSliderFlags(flags))

	return bool(ret)
}

// DragScalar wraps igDragScalar.
func DragScalar(label ffi.CString, data_type GuiDataType, p_data uintptr, v_speed float32, p_min uintptr, p_max uintptr, format ffi.CString, flags GuiSliderFlags) bool {
	ret := C.igDragScalar((*C.char)(label.Raw()), C.ImGuiDataType(data_type), unsafe.Pointer(p_data), C.float(v_speed), unsafe.Pointer(p_min), unsafe.Pointer(p_max), (*C.char)(format.Raw()), C.ImGuiSliderFlags(flags))

	return bool(ret)
}

// DragScalarN wraps igDragScalarN.
func DragScalarN(label ffi.CString, data_type GuiDataType, p_data uintptr, components int32, v_speed float32, p_min uintptr, p_max uintptr, format ffi.CString, flags GuiSliderFlags) bool {
	ret := C.igDragScalarN((*C.char)(label.Raw()), C.ImGuiDataType(data_type), unsafe.Pointer(p_data), C.int32_t(components), C.float(v_speed), unsafe.Pointer(p_min), unsafe.Pointer(p_max), (*C.char)(format.Raw()), C.ImGuiSliderFlags(flags))

	return bool(ret)
}

// Dummy wraps igDummy.
func Dummy(size Vec2) {
	C.igDummy(*(*C.ImVec2)(unsafe.Pointer(size)))
}

// End wraps igEnd.
func End() {
	C.igEnd()
}

// EndChild wraps igEndChild.
func EndChild() {
	C.igEndChild()
}

// EndCombo wraps igEndCombo.
func EndCombo() {
	C.igEndCombo()
}

// EndDisabled wraps igEndDisabled.
func EndDisabled() {
	C.igEndDisabled()
}

// EndDragDropSource wraps igEndDragDropSource.
func EndDragDropSource() {
	C.igEndDragDropSource()
}

// EndDragDropTarget wraps igEndDragDropTarget.
func EndDragDropTarget() {
	C.igEndDragDropTarget()
}

// EndFrame wraps igEndFrame.
func EndFrame() {
	C.igEndFrame()
}

// EndGroup wraps igEndGroup.
func EndGroup() {
	C.igEndGroup()
}

// EndListBox wraps igEndListBox.
func EndListBox() {
	C.igEndListBox()
}

// EndMainMenuBar wraps igEndMainMenuBar.
func EndMainMenuBar() {
	C.igEndMainMenuBar()
}

// EndMenu wraps igEndMenu.
func EndMenu() {
	C.igEndMenu()
}

// EndMenuBar wraps igEndMenuBar.
func EndMenuBar() {
	C.igEndMenuBar()
}

// EndMultiSelect wraps igEndMultiSelect.
func EndMultiSelect() GuiMultiSelectIO {
	ret := C.igEndMultiSelect()

	return GuiMultiSelectIO(unsafe.Pointer(ret))
}

// EndPopup wraps igEndPopup.
func EndPopup() {
	C.igEndPopup()
}

// EndTabBar wraps igEndTabBar.
func EndTabBar() {
	C.igEndTabBar()
}

// EndTabItem wraps igEndTabItem.
func EndTabItem() {
	C.igEndTabItem()
}

// EndTable wraps igEndTable.
func EndTable() {
	C.igEndTable()
}

// EndTooltip wraps igEndTooltip.
func EndTooltip() {
	C.igEndTooltip()
}

// igFindViewportByID.id is unsupported: category unsupported.

// FindViewportByPlatformHandle wraps igFindViewportByPlatformHandle.
func FindViewportByPlatformHandle(platform_handle uintptr) GuiViewport {
	ret := C.igFindViewportByPlatformHandle(unsafe.Pointer(platform_handle))

	return GuiViewport(unsafe.Pointer(ret))
}

// igGetAllocatorFunctions.p_alloc_func is unsupported: category unsupported.

// GetBackgroundDrawList wraps igGetBackgroundDrawList.
func GetBackgroundDrawList(viewport GuiViewport) DrawList {
	ret := C.igGetBackgroundDrawList((*C.ImGuiViewport)(unsafe.Pointer(viewport)))

	return DrawList(unsafe.Pointer(ret))
}

// igGetClipboardText is unsupported: return pointer type -> ?? char.

// GetColorU32_Col wraps igGetColorU32_Col.
func GetColorU32_Col(idx GuiCol, alpha_mul float32) uint32 {
	ret := C.igGetColorU32_Col(C.ImGuiCol(idx), C.float(alpha_mul))

	return uint32(ret)
}

// GetColorU32_U32 wraps igGetColorU32_U32.
func GetColorU32_U32(col uint32, alpha_mul float32) uint32 {
	ret := C.igGetColorU32_U32(C.uint32_t(col), C.float(alpha_mul))

	return uint32(ret)
}

// GetColorU32_Vec4 wraps igGetColorU32_Vec4.
func GetColorU32_Vec4(col Vec4) uint32 {
	ret := C.igGetColorU32_Vec4(*(*C.ImVec4)(unsafe.Pointer(col)))

	return uint32(ret)
}

// GetColumnIndex wraps igGetColumnIndex.
func GetColumnIndex() int32 {
	ret := C.igGetColumnIndex()

	return int32(ret)
}

// GetColumnOffset wraps igGetColumnOffset.
func GetColumnOffset(column_index int32) float32 {
	ret := C.igGetColumnOffset(C.int32_t(column_index))

	return float32(ret)
}

// GetColumnWidth wraps igGetColumnWidth.
func GetColumnWidth(column_index int32) float32 {
	ret := C.igGetColumnWidth(C.int32_t(column_index))

	return float32(ret)
}

// GetColumnsCount wraps igGetColumnsCount.
func GetColumnsCount() int32 {
	ret := C.igGetColumnsCount()

	return int32(ret)
}

// GetContentRegionAvail wraps igGetContentRegionAvail.
func GetContentRegionAvail(pOut Vec2) {
	C.igGetContentRegionAvail((*C.ImVec2)(unsafe.Pointer(pOut)))
}

// GetCurrentContext wraps igGetCurrentContext.
func GetCurrentContext() GuiContext {
	ret := C.igGetCurrentContext()

	return GuiContext(unsafe.Pointer(ret))
}

// GetCursorPos wraps igGetCursorPos.
func GetCursorPos(pOut Vec2) {
	C.igGetCursorPos((*C.ImVec2)(unsafe.Pointer(pOut)))
}

// GetCursorPosX wraps igGetCursorPosX.
func GetCursorPosX() float32 {
	ret := C.igGetCursorPosX()

	return float32(ret)
}

// GetCursorPosY wraps igGetCursorPosY.
func GetCursorPosY() float32 {
	ret := C.igGetCursorPosY()

	return float32(ret)
}

// GetCursorScreenPos wraps igGetCursorScreenPos.
func GetCursorScreenPos(pOut Vec2) {
	C.igGetCursorScreenPos((*C.ImVec2)(unsafe.Pointer(pOut)))
}

// GetCursorStartPos wraps igGetCursorStartPos.
func GetCursorStartPos(pOut Vec2) {
	C.igGetCursorStartPos((*C.ImVec2)(unsafe.Pointer(pOut)))
}

// GetDragDropPayload wraps igGetDragDropPayload.
func GetDragDropPayload() GuiPayload {
	ret := C.igGetDragDropPayload()

	return GuiPayload(unsafe.Pointer(ret))
}

// GetDrawData wraps igGetDrawData.
func GetDrawData() DrawData {
	ret := C.igGetDrawData()

	return DrawData(unsafe.Pointer(ret))
}

// GetDrawListSharedData wraps igGetDrawListSharedData.
func GetDrawListSharedData() DrawListSharedData {
	ret := C.igGetDrawListSharedData()

	return DrawListSharedData(unsafe.Pointer(ret))
}

// GetFont wraps igGetFont.
func GetFont() Font {
	ret := C.igGetFont()

	return Font(unsafe.Pointer(ret))
}

// GetFontBaked wraps igGetFontBaked.
func GetFontBaked() FontBaked {
	ret := C.igGetFontBaked()

	return FontBaked(unsafe.Pointer(ret))
}

// GetFontSize wraps igGetFontSize.
func GetFontSize() float32 {
	ret := C.igGetFontSize()

	return float32(ret)
}

// GetFontTexUvWhitePixel wraps igGetFontTexUvWhitePixel.
func GetFontTexUvWhitePixel(pOut Vec2) {
	C.igGetFontTexUvWhitePixel((*C.ImVec2)(unsafe.Pointer(pOut)))
}

// GetForegroundDrawList wraps igGetForegroundDrawList_ViewportPtr.
func GetForegroundDrawList(viewport GuiViewport) DrawList {
	ret := C.igGetForegroundDrawList_ViewportPtr((*C.ImGuiViewport)(unsafe.Pointer(viewport)))

	return DrawList(unsafe.Pointer(ret))
}

// GetFrameCount wraps igGetFrameCount.
func GetFrameCount() int32 {
	ret := C.igGetFrameCount()

	return int32(ret)
}

// GetFrameHeight wraps igGetFrameHeight.
func GetFrameHeight() float32 {
	ret := C.igGetFrameHeight()

	return float32(ret)
}

// GetFrameHeightWithSpacing wraps igGetFrameHeightWithSpacing.
func GetFrameHeightWithSpacing() float32 {
	ret := C.igGetFrameHeightWithSpacing()

	return float32(ret)
}

// igGetID_Int is unsupported: unknown return category unsupported.

// igGetID_Ptr is unsupported: unknown return category unsupported.

// igGetID_Str is unsupported: unknown return category unsupported.

// igGetID_StrStr is unsupported: unknown return category unsupported.

// GetIO wraps igGetIO_Nil.
func GetIO() GuiIO {
	ret := C.igGetIO_Nil()

	return GuiIO(unsafe.Pointer(ret))
}

// igGetItemID is unsupported: unknown return category unsupported.

// GetItemRectMax wraps igGetItemRectMax.
func GetItemRectMax(pOut Vec2) {
	C.igGetItemRectMax((*C.ImVec2)(unsafe.Pointer(pOut)))
}

// GetItemRectMin wraps igGetItemRectMin.
func GetItemRectMin(pOut Vec2) {
	C.igGetItemRectMin((*C.ImVec2)(unsafe.Pointer(pOut)))
}

// GetItemRectSize wraps igGetItemRectSize.
func GetItemRectSize(pOut Vec2) {
	C.igGetItemRectSize((*C.ImVec2)(unsafe.Pointer(pOut)))
}

// igGetKeyName is unsupported: return pointer type -> ?? char.

// GetKeyPressedAmount wraps igGetKeyPressedAmount.
func GetKeyPressedAmount(key GuiKey, repeat_delay float32, rate float32) int32 {
	ret := C.igGetKeyPressedAmount(C.ImGuiKey(key), C.float(repeat_delay), C.float(rate))

	return int32(ret)
}

// GetMainViewport wraps igGetMainViewport.
func GetMainViewport() GuiViewport {
	ret := C.igGetMainViewport()

	return GuiViewport(unsafe.Pointer(ret))
}

// GetMouseClickedCount wraps igGetMouseClickedCount.
func GetMouseClickedCount(button GuiMouseButton) int32 {
	ret := C.igGetMouseClickedCount(C.ImGuiMouseButton(button))

	return int32(ret)
}

// GetMouseCursor wraps igGetMouseCursor.
func GetMouseCursor() GuiMouseCursor {
	ret := C.igGetMouseCursor()

	return GuiMouseCursor(ret)
}

// GetMouseDragDelta wraps igGetMouseDragDelta.
func GetMouseDragDelta(pOut Vec2, button GuiMouseButton, lock_threshold float32) {
	C.igGetMouseDragDelta((*C.ImVec2)(unsafe.Pointer(pOut)), C.ImGuiMouseButton(button), C.float(lock_threshold))
}

// GetMousePos wraps igGetMousePos.
func GetMousePos(pOut Vec2) {
	C.igGetMousePos((*C.ImVec2)(unsafe.Pointer(pOut)))
}

// GetMousePosOnOpeningCurrentPopup wraps igGetMousePosOnOpeningCurrentPopup.
func GetMousePosOnOpeningCurrentPopup(pOut Vec2) {
	C.igGetMousePosOnOpeningCurrentPopup((*C.ImVec2)(unsafe.Pointer(pOut)))
}

// GetPlatformIO wraps igGetPlatformIO_Nil.
func GetPlatformIO() GuiPlatformIO {
	ret := C.igGetPlatformIO_Nil()

	return GuiPlatformIO(unsafe.Pointer(ret))
}

// GetScrollMaxX wraps igGetScrollMaxX.
func GetScrollMaxX() float32 {
	ret := C.igGetScrollMaxX()

	return float32(ret)
}

// GetScrollMaxY wraps igGetScrollMaxY.
func GetScrollMaxY() float32 {
	ret := C.igGetScrollMaxY()

	return float32(ret)
}

// GetScrollX wraps igGetScrollX.
func GetScrollX() float32 {
	ret := C.igGetScrollX()

	return float32(ret)
}

// GetScrollY wraps igGetScrollY.
func GetScrollY() float32 {
	ret := C.igGetScrollY()

	return float32(ret)
}

// GetStateStorage wraps igGetStateStorage.
func GetStateStorage() GuiStorage {
	ret := C.igGetStateStorage()

	return GuiStorage(unsafe.Pointer(ret))
}

// GetStyle wraps igGetStyle.
func GetStyle() GuiStyle {
	ret := C.igGetStyle()

	return GuiStyle(unsafe.Pointer(ret))
}

// igGetStyleColorName is unsupported: return pointer type -> ?? char.

// GetStyleColorVec4 wraps igGetStyleColorVec4.
func GetStyleColorVec4(idx GuiCol) Vec4 {
	ret := C.igGetStyleColorVec4(C.ImGuiCol(idx))

	return Vec4(unsafe.Pointer(ret))
}

// GetTextLineHeight wraps igGetTextLineHeight.
func GetTextLineHeight() float32 {
	ret := C.igGetTextLineHeight()

	return float32(ret)
}

// GetTextLineHeightWithSpacing wraps igGetTextLineHeightWithSpacing.
func GetTextLineHeightWithSpacing() float32 {
	ret := C.igGetTextLineHeightWithSpacing()

	return float32(ret)
}

// igGetTime is unsupported: unknown return category unsupported.

// GetTreeNodeToLabelSpacing wraps igGetTreeNodeToLabelSpacing.
func GetTreeNodeToLabelSpacing() float32 {
	ret := C.igGetTreeNodeToLabelSpacing()

	return float32(ret)
}

// igGetVersion is unsupported: return pointer type -> ?? char.

// igGetWindowDockID is unsupported: unknown return category unsupported.

// GetWindowDpiScale wraps igGetWindowDpiScale.
func GetWindowDpiScale() float32 {
	ret := C.igGetWindowDpiScale()

	return float32(ret)
}

// GetWindowDrawList wraps igGetWindowDrawList.
func GetWindowDrawList() DrawList {
	ret := C.igGetWindowDrawList()

	return DrawList(unsafe.Pointer(ret))
}

// GetWindowHeight wraps igGetWindowHeight.
func GetWindowHeight() float32 {
	ret := C.igGetWindowHeight()

	return float32(ret)
}

// GetWindowPos wraps igGetWindowPos.
func GetWindowPos(pOut Vec2) {
	C.igGetWindowPos((*C.ImVec2)(unsafe.Pointer(pOut)))
}

// GetWindowSize wraps igGetWindowSize.
func GetWindowSize(pOut Vec2) {
	C.igGetWindowSize((*C.ImVec2)(unsafe.Pointer(pOut)))
}

// GetWindowViewport wraps igGetWindowViewport.
func GetWindowViewport() GuiViewport {
	ret := C.igGetWindowViewport()

	return GuiViewport(unsafe.Pointer(ret))
}

// GetWindowWidth wraps igGetWindowWidth.
func GetWindowWidth() float32 {
	ret := C.igGetWindowWidth()

	return float32(ret)
}

// Age wraps igImage.
func Age(tex_ref TextureRef, image_size Vec2, uv0 Vec2, uv1 Vec2) {
	C.igImage(*(*C.ImTextureRef)(unsafe.Pointer(tex_ref)), *(*C.ImVec2)(unsafe.Pointer(image_size)), *(*C.ImVec2)(unsafe.Pointer(uv0)), *(*C.ImVec2)(unsafe.Pointer(uv1)))
}

// AgeButton wraps igImageButton.
func AgeButton(str_id ffi.CString, tex_ref TextureRef, image_size Vec2, uv0 Vec2, uv1 Vec2, bg_col Vec4, tint_col Vec4) bool {
	ret := C.igImageButton((*C.char)(str_id.Raw()), *(*C.ImTextureRef)(unsafe.Pointer(tex_ref)), *(*C.ImVec2)(unsafe.Pointer(image_size)), *(*C.ImVec2)(unsafe.Pointer(uv0)), *(*C.ImVec2)(unsafe.Pointer(uv1)), *(*C.ImVec4)(unsafe.Pointer(bg_col)), *(*C.ImVec4)(unsafe.Pointer(tint_col)))

	return bool(ret)
}

// AgeWithBg wraps igImageWithBg.
func AgeWithBg(tex_ref TextureRef, image_size Vec2, uv0 Vec2, uv1 Vec2, bg_col Vec4, tint_col Vec4) {
	C.igImageWithBg(*(*C.ImTextureRef)(unsafe.Pointer(tex_ref)), *(*C.ImVec2)(unsafe.Pointer(image_size)), *(*C.ImVec2)(unsafe.Pointer(uv0)), *(*C.ImVec2)(unsafe.Pointer(uv1)), *(*C.ImVec4)(unsafe.Pointer(bg_col)), *(*C.ImVec4)(unsafe.Pointer(tint_col)))
}

// Indent wraps igIndent.
func Indent(indent_w float32) {
	C.igIndent(C.float(indent_w))
}

// igInputDouble.v is unsupported: category unsupported.

// InputFloat wraps igInputFloat.
func InputFloat(label ffi.CString, v ffi.Ref[float32], step float32, step_fast float32, format ffi.CString, flags GuiInputTextFlags) bool {
	ret := C.igInputFloat((*C.char)(label.Raw()), (*C.float)(v.Raw()), C.float(step), C.float(step_fast), (*C.char)(format.Raw()), C.ImGuiInputTextFlags(flags))

	return bool(ret)
}

// igInputFloat2.v is unsupported: category unsupported.

// igInputFloat3.v is unsupported: category unsupported.

// igInputFloat4.v is unsupported: category unsupported.

// InputInt wraps igInputInt.
func InputInt(label ffi.CString, v ffi.Ref[int32], step int32, step_fast int32, flags GuiInputTextFlags) bool {
	ret := C.igInputInt((*C.char)(label.Raw()), (*C.int32_t)(v.Raw()), C.int32_t(step), C.int32_t(step_fast), C.ImGuiInputTextFlags(flags))

	return bool(ret)
}

// igInputInt2.v is unsupported: category unsupported.

// igInputInt3.v is unsupported: category unsupported.

// igInputInt4.v is unsupported: category unsupported.

// InputScalar wraps igInputScalar.
func InputScalar(label ffi.CString, data_type GuiDataType, p_data uintptr, p_step uintptr, p_step_fast uintptr, format ffi.CString, flags GuiInputTextFlags) bool {
	ret := C.igInputScalar((*C.char)(label.Raw()), C.ImGuiDataType(data_type), unsafe.Pointer(p_data), unsafe.Pointer(p_step), unsafe.Pointer(p_step_fast), (*C.char)(format.Raw()), C.ImGuiInputTextFlags(flags))

	return bool(ret)
}

// InputScalarN wraps igInputScalarN.
func InputScalarN(label ffi.CString, data_type GuiDataType, p_data uintptr, components int32, p_step uintptr, p_step_fast uintptr, format ffi.CString, flags GuiInputTextFlags) bool {
	ret := C.igInputScalarN((*C.char)(label.Raw()), C.ImGuiDataType(data_type), unsafe.Pointer(p_data), C.int32_t(components), unsafe.Pointer(p_step), unsafe.Pointer(p_step_fast), (*C.char)(format.Raw()), C.ImGuiInputTextFlags(flags))

	return bool(ret)
}

// igInputText.buf_size is unsupported: category unsupported.

// igInputTextMultiline.buf_size is unsupported: category unsupported.

// igInputTextWithHint.buf_size is unsupported: category unsupported.

// InvisibleButton wraps igInvisibleButton.
func InvisibleButton(str_id ffi.CString, size Vec2, flags GuiButtonFlags) bool {
	ret := C.igInvisibleButton((*C.char)(str_id.Raw()), *(*C.ImVec2)(unsafe.Pointer(size)), C.ImGuiButtonFlags(flags))

	return bool(ret)
}

// IsAnyItemActive wraps igIsAnyItemActive.
func IsAnyItemActive() bool {
	ret := C.igIsAnyItemActive()

	return bool(ret)
}

// IsAnyItemFocused wraps igIsAnyItemFocused.
func IsAnyItemFocused() bool {
	ret := C.igIsAnyItemFocused()

	return bool(ret)
}

// IsAnyItemHovered wraps igIsAnyItemHovered.
func IsAnyItemHovered() bool {
	ret := C.igIsAnyItemHovered()

	return bool(ret)
}

// IsAnyMouseDown wraps igIsAnyMouseDown.
func IsAnyMouseDown() bool {
	ret := C.igIsAnyMouseDown()

	return bool(ret)
}

// IsItemActivated wraps igIsItemActivated.
func IsItemActivated() bool {
	ret := C.igIsItemActivated()

	return bool(ret)
}

// IsItemActive wraps igIsItemActive.
func IsItemActive() bool {
	ret := C.igIsItemActive()

	return bool(ret)
}

// IsItemClicked wraps igIsItemClicked.
func IsItemClicked(mouse_button GuiMouseButton) bool {
	ret := C.igIsItemClicked(C.ImGuiMouseButton(mouse_button))

	return bool(ret)
}

// IsItemDeactivated wraps igIsItemDeactivated.
func IsItemDeactivated() bool {
	ret := C.igIsItemDeactivated()

	return bool(ret)
}

// IsItemDeactivatedAfterEdit wraps igIsItemDeactivatedAfterEdit.
func IsItemDeactivatedAfterEdit() bool {
	ret := C.igIsItemDeactivatedAfterEdit()

	return bool(ret)
}

// IsItemEdited wraps igIsItemEdited.
func IsItemEdited() bool {
	ret := C.igIsItemEdited()

	return bool(ret)
}

// IsItemFocused wraps igIsItemFocused.
func IsItemFocused() bool {
	ret := C.igIsItemFocused()

	return bool(ret)
}

// IsItemHovered wraps igIsItemHovered.
func IsItemHovered(flags GuiHoveredFlags) bool {
	ret := C.igIsItemHovered(C.ImGuiHoveredFlags(flags))

	return bool(ret)
}

// IsItemToggledOpen wraps igIsItemToggledOpen.
func IsItemToggledOpen() bool {
	ret := C.igIsItemToggledOpen()

	return bool(ret)
}

// IsItemToggledSelection wraps igIsItemToggledSelection.
func IsItemToggledSelection() bool {
	ret := C.igIsItemToggledSelection()

	return bool(ret)
}

// IsItemVisible wraps igIsItemVisible.
func IsItemVisible() bool {
	ret := C.igIsItemVisible()

	return bool(ret)
}

// igIsKeyChordPressed_Nil.key_chord is unsupported: category unsupported.

// IsKeyDown wraps igIsKeyDown_Nil.
func IsKeyDown(key GuiKey) bool {
	ret := C.igIsKeyDown_Nil(C.ImGuiKey(key))

	return bool(ret)
}

// IsKeyPressed wraps igIsKeyPressed_Bool.
func IsKeyPressed(key GuiKey, repeat bool) bool {
	ret := C.igIsKeyPressed_Bool(C.ImGuiKey(key), C.bool(repeat))

	return bool(ret)
}

// IsKeyReleased wraps igIsKeyReleased_Nil.
func IsKeyReleased(key GuiKey) bool {
	ret := C.igIsKeyReleased_Nil(C.ImGuiKey(key))

	return bool(ret)
}

// IsMouseClicked wraps igIsMouseClicked_Bool.
func IsMouseClicked(button GuiMouseButton, repeat bool) bool {
	ret := C.igIsMouseClicked_Bool(C.ImGuiMouseButton(button), C.bool(repeat))

	return bool(ret)
}

// IsMouseDoubleClicked wraps igIsMouseDoubleClicked_Nil.
func IsMouseDoubleClicked(button GuiMouseButton) bool {
	ret := C.igIsMouseDoubleClicked_Nil(C.ImGuiMouseButton(button))

	return bool(ret)
}

// IsMouseDown wraps igIsMouseDown_Nil.
func IsMouseDown(button GuiMouseButton) bool {
	ret := C.igIsMouseDown_Nil(C.ImGuiMouseButton(button))

	return bool(ret)
}

// IsMouseDragging wraps igIsMouseDragging.
func IsMouseDragging(button GuiMouseButton, lock_threshold float32) bool {
	ret := C.igIsMouseDragging(C.ImGuiMouseButton(button), C.float(lock_threshold))

	return bool(ret)
}

// IsMouseHoveringRect wraps igIsMouseHoveringRect.
func IsMouseHoveringRect(r_min Vec2, r_max Vec2, clip bool) bool {
	ret := C.igIsMouseHoveringRect(*(*C.ImVec2)(unsafe.Pointer(r_min)), *(*C.ImVec2)(unsafe.Pointer(r_max)), C.bool(clip))

	return bool(ret)
}

// IsMousePosValid wraps igIsMousePosValid.
func IsMousePosValid(mouse_pos Vec2) bool {
	ret := C.igIsMousePosValid((*C.ImVec2)(unsafe.Pointer(mouse_pos)))

	return bool(ret)
}

// IsMouseReleasedWithDelay wraps igIsMouseReleasedWithDelay.
func IsMouseReleasedWithDelay(button GuiMouseButton, delay float32) bool {
	ret := C.igIsMouseReleasedWithDelay(C.ImGuiMouseButton(button), C.float(delay))

	return bool(ret)
}

// IsMouseReleased wraps igIsMouseReleased_Nil.
func IsMouseReleased(button GuiMouseButton) bool {
	ret := C.igIsMouseReleased_Nil(C.ImGuiMouseButton(button))

	return bool(ret)
}

// IsPopupOpen wraps igIsPopupOpen_Str.
func IsPopupOpen(str_id ffi.CString, flags GuiPopupFlags) bool {
	ret := C.igIsPopupOpen_Str((*C.char)(str_id.Raw()), C.ImGuiPopupFlags(flags))

	return bool(ret)
}

// IsRectVisible_Nil wraps igIsRectVisible_Nil.
func IsRectVisible_Nil(size Vec2) bool {
	ret := C.igIsRectVisible_Nil(*(*C.ImVec2)(unsafe.Pointer(size)))

	return bool(ret)
}

// IsRectVisible_Vec2 wraps igIsRectVisible_Vec2.
func IsRectVisible_Vec2(rect_min Vec2, rect_max Vec2) bool {
	ret := C.igIsRectVisible_Vec2(*(*C.ImVec2)(unsafe.Pointer(rect_min)), *(*C.ImVec2)(unsafe.Pointer(rect_max)))

	return bool(ret)
}

// IsWindowAppearing wraps igIsWindowAppearing.
func IsWindowAppearing() bool {
	ret := C.igIsWindowAppearing()

	return bool(ret)
}

// IsWindowCollapsed wraps igIsWindowCollapsed.
func IsWindowCollapsed() bool {
	ret := C.igIsWindowCollapsed()

	return bool(ret)
}

// IsWindowDocked wraps igIsWindowDocked.
func IsWindowDocked() bool {
	ret := C.igIsWindowDocked()

	return bool(ret)
}

// IsWindowFocused wraps igIsWindowFocused.
func IsWindowFocused(flags GuiFocusedFlags) bool {
	ret := C.igIsWindowFocused(C.ImGuiFocusedFlags(flags))

	return bool(ret)
}

// IsWindowHovered wraps igIsWindowHovered.
func IsWindowHovered(flags GuiHoveredFlags) bool {
	ret := C.igIsWindowHovered(C.ImGuiHoveredFlags(flags))

	return bool(ret)
}

// igLabelText.... is unsupported: category unsupported.

// igLabelTextV.args is unsupported: category unsupported.

// igListBox_FnStrPtr.getter is unsupported: category unsupported.

// igListBox_Str_arr.items is unsupported: category unsupported.

// LoadIniSettingsFromDisk wraps igLoadIniSettingsFromDisk.
func LoadIniSettingsFromDisk(ini_filename ffi.CString) {
	C.igLoadIniSettingsFromDisk((*C.char)(ini_filename.Raw()))
}

// igLoadIniSettingsFromMemory.ini_size is unsupported: category unsupported.

// LogButtons wraps igLogButtons.
func LogButtons() {
	C.igLogButtons()
}

// LogFinish wraps igLogFinish.
func LogFinish() {
	C.igLogFinish()
}

// igLogText.... is unsupported: category unsupported.

// igLogTextV.args is unsupported: category unsupported.

// LogToClipboard wraps igLogToClipboard.
func LogToClipboard(auto_open_depth int32) {
	C.igLogToClipboard(C.int32_t(auto_open_depth))
}

// LogToFile wraps igLogToFile.
func LogToFile(auto_open_depth int32, filename ffi.CString) {
	C.igLogToFile(C.int32_t(auto_open_depth), (*C.char)(filename.Raw()))
}

// LogToTTY wraps igLogToTTY.
func LogToTTY(auto_open_depth int32) {
	C.igLogToTTY(C.int32_t(auto_open_depth))
}

// igMemAlloc is unsupported: return pointer type -> ?? void.

// MemFree wraps igMemFree.
func MemFree(ptr uintptr) {
	C.igMemFree(unsafe.Pointer(ptr))
}

// MenuItem_Bool wraps igMenuItem_Bool.
func MenuItem_Bool(label ffi.CString, shortcut ffi.CString, selected bool, enabled bool) bool {
	ret := C.igMenuItem_Bool((*C.char)(label.Raw()), (*C.char)(shortcut.Raw()), C.bool(selected), C.bool(enabled))

	return bool(ret)
}

// MenuItem_BoolPtr wraps igMenuItem_BoolPtr.
func MenuItem_BoolPtr(label ffi.CString, shortcut ffi.CString, p_selected ffi.Ref[bool], enabled bool) bool {
	ret := C.igMenuItem_BoolPtr((*C.char)(label.Raw()), (*C.char)(shortcut.Raw()), (*C.bool)(p_selected.Raw()), C.bool(enabled))

	return bool(ret)
}

// NewFrame wraps igNewFrame.
func NewFrame() {
	C.igNewFrame()
}

// NewLine wraps igNewLine.
func NewLine() {
	C.igNewLine()
}

// NextColumn wraps igNextColumn.
func NextColumn() {
	C.igNextColumn()
}

// OpenPopupOnItemClick wraps igOpenPopupOnItemClick.
func OpenPopupOnItemClick(str_id ffi.CString, popup_flags GuiPopupFlags) {
	C.igOpenPopupOnItemClick((*C.char)(str_id.Raw()), C.ImGuiPopupFlags(popup_flags))
}

// igOpenPopup_ID.id is unsupported: category unsupported.

// OpenPopup_Str wraps igOpenPopup_Str.
func OpenPopup_Str(str_id ffi.CString, popup_flags GuiPopupFlags) {
	C.igOpenPopup_Str((*C.char)(str_id.Raw()), C.ImGuiPopupFlags(popup_flags))
}

// PlotHistogram_FloatPtr wraps igPlotHistogram_FloatPtr.
func PlotHistogram_FloatPtr(label ffi.CString, values ffi.Ref[float32], values_count int32, values_offset int32, overlay_text ffi.CString, scale_min float32, scale_max float32, graph_size Vec2, stride int32) {
	C.igPlotHistogram_FloatPtr((*C.char)(label.Raw()), (*C.float)(values.Raw()), C.int32_t(values_count), C.int32_t(values_offset), (*C.char)(overlay_text.Raw()), C.float(scale_min), C.float(scale_max), *(*C.ImVec2)(unsafe.Pointer(graph_size)), C.int32_t(stride))
}

// igPlotHistogram_FnFloatPtr.values_getter is unsupported: category unsupported.

// PlotLines_FloatPtr wraps igPlotLines_FloatPtr.
func PlotLines_FloatPtr(label ffi.CString, values ffi.Ref[float32], values_count int32, values_offset int32, overlay_text ffi.CString, scale_min float32, scale_max float32, graph_size Vec2, stride int32) {
	C.igPlotLines_FloatPtr((*C.char)(label.Raw()), (*C.float)(values.Raw()), C.int32_t(values_count), C.int32_t(values_offset), (*C.char)(overlay_text.Raw()), C.float(scale_min), C.float(scale_max), *(*C.ImVec2)(unsafe.Pointer(graph_size)), C.int32_t(stride))
}

// igPlotLines_FnFloatPtr.values_getter is unsupported: category unsupported.

// PopClipRect wraps igPopClipRect.
func PopClipRect() {
	C.igPopClipRect()
}

// PopFont wraps igPopFont.
func PopFont() {
	C.igPopFont()
}

// PopID wraps igPopID.
func PopID() {
	C.igPopID()
}

// PopItemFlag wraps igPopItemFlag.
func PopItemFlag() {
	C.igPopItemFlag()
}

// PopItemWidth wraps igPopItemWidth.
func PopItemWidth() {
	C.igPopItemWidth()
}

// PopStyleColor wraps igPopStyleColor.
func PopStyleColor(count int32) {
	C.igPopStyleColor(C.int32_t(count))
}

// PopStyleVar wraps igPopStyleVar.
func PopStyleVar(count int32) {
	C.igPopStyleVar(C.int32_t(count))
}

// PopTextWrapPos wraps igPopTextWrapPos.
func PopTextWrapPos() {
	C.igPopTextWrapPos()
}

// ProgressBar wraps igProgressBar.
func ProgressBar(fraction float32, size_arg Vec2, overlay ffi.CString) {
	C.igProgressBar(C.float(fraction), *(*C.ImVec2)(unsafe.Pointer(size_arg)), (*C.char)(overlay.Raw()))
}

// PushClipRect wraps igPushClipRect.
func PushClipRect(clip_rect_min Vec2, clip_rect_max Vec2, intersect_with_current_clip_rect bool) {
	C.igPushClipRect(*(*C.ImVec2)(unsafe.Pointer(clip_rect_min)), *(*C.ImVec2)(unsafe.Pointer(clip_rect_max)), C.bool(intersect_with_current_clip_rect))
}

// PushFont wraps igPushFont.
func PushFont(font Font, font_size_base_unscaled float32) {
	C.igPushFont((*C.ImFont)(unsafe.Pointer(font)), C.float(font_size_base_unscaled))
}

// PushID_Int wraps igPushID_Int.
func PushID_Int(int_id int32) {
	C.igPushID_Int(C.int32_t(int_id))
}

// PushID_Ptr wraps igPushID_Ptr.
func PushID_Ptr(ptr_id uintptr) {
	C.igPushID_Ptr(unsafe.Pointer(ptr_id))
}

// PushID_Str wraps igPushID_Str.
func PushID_Str(str_id ffi.CString) {
	C.igPushID_Str((*C.char)(str_id.Raw()))
}

// PushID_StrStr wraps igPushID_StrStr.
func PushID_StrStr(str_id_begin ffi.CString, str_id_end ffi.CString) {
	C.igPushID_StrStr((*C.char)(str_id_begin.Raw()), (*C.char)(str_id_end.Raw()))
}

// PushItemFlag wraps igPushItemFlag.
func PushItemFlag(option GuiItemFlags, enabled bool) {
	C.igPushItemFlag(C.ImGuiItemFlags(option), C.bool(enabled))
}

// PushItemWidth wraps igPushItemWidth.
func PushItemWidth(item_width float32) {
	C.igPushItemWidth(C.float(item_width))
}

// PushStyleColor_U32 wraps igPushStyleColor_U32.
func PushStyleColor_U32(idx GuiCol, col uint32) {
	C.igPushStyleColor_U32(C.ImGuiCol(idx), C.uint32_t(col))
}

// PushStyleColor_Vec4 wraps igPushStyleColor_Vec4.
func PushStyleColor_Vec4(idx GuiCol, col Vec4) {
	C.igPushStyleColor_Vec4(C.ImGuiCol(idx), *(*C.ImVec4)(unsafe.Pointer(col)))
}

// PushStyleVarX wraps igPushStyleVarX.
func PushStyleVarX(idx GuiStyleVar, val_x float32) {
	C.igPushStyleVarX(C.ImGuiStyleVar(idx), C.float(val_x))
}

// PushStyleVarY wraps igPushStyleVarY.
func PushStyleVarY(idx GuiStyleVar, val_y float32) {
	C.igPushStyleVarY(C.ImGuiStyleVar(idx), C.float(val_y))
}

// PushStyleVar_Float wraps igPushStyleVar_Float.
func PushStyleVar_Float(idx GuiStyleVar, val float32) {
	C.igPushStyleVar_Float(C.ImGuiStyleVar(idx), C.float(val))
}

// PushStyleVar_Vec2 wraps igPushStyleVar_Vec2.
func PushStyleVar_Vec2(idx GuiStyleVar, val Vec2) {
	C.igPushStyleVar_Vec2(C.ImGuiStyleVar(idx), *(*C.ImVec2)(unsafe.Pointer(val)))
}

// PushTextWrapPos wraps igPushTextWrapPos.
func PushTextWrapPos(wrap_local_pos_x float32) {
	C.igPushTextWrapPos(C.float(wrap_local_pos_x))
}

// RadioButton_Bool wraps igRadioButton_Bool.
func RadioButton_Bool(label ffi.CString, active bool) bool {
	ret := C.igRadioButton_Bool((*C.char)(label.Raw()), C.bool(active))

	return bool(ret)
}

// RadioButton_IntPtr wraps igRadioButton_IntPtr.
func RadioButton_IntPtr(label ffi.CString, v ffi.Ref[int32], v_button int32) bool {
	ret := C.igRadioButton_IntPtr((*C.char)(label.Raw()), (*C.int32_t)(v.Raw()), C.int32_t(v_button))

	return bool(ret)
}

// Render wraps igRender.
func Render() {
	C.igRender()
}

// RenderPlatformWindowsDefault wraps igRenderPlatformWindowsDefault.
func RenderPlatformWindowsDefault(platform_render_arg uintptr, renderer_render_arg uintptr) {
	C.igRenderPlatformWindowsDefault(unsafe.Pointer(platform_render_arg), unsafe.Pointer(renderer_render_arg))
}

// ResetMouseDragDelta wraps igResetMouseDragDelta.
func ResetMouseDragDelta(button GuiMouseButton) {
	C.igResetMouseDragDelta(C.ImGuiMouseButton(button))
}

// SameLine wraps igSameLine.
func SameLine(offset_from_start_x float32, spacing float32) {
	C.igSameLine(C.float(offset_from_start_x), C.float(spacing))
}

// SaveIniSettingsToDisk wraps igSaveIniSettingsToDisk.
func SaveIniSettingsToDisk(ini_filename ffi.CString) {
	C.igSaveIniSettingsToDisk((*C.char)(ini_filename.Raw()))
}

// igSaveIniSettingsToMemory is unsupported: return pointer type -> ?? char.

// Selectable_Bool wraps igSelectable_Bool.
func Selectable_Bool(label ffi.CString, selected bool, flags GuiSelectableFlags, size Vec2) bool {
	ret := C.igSelectable_Bool((*C.char)(label.Raw()), C.bool(selected), C.ImGuiSelectableFlags(flags), *(*C.ImVec2)(unsafe.Pointer(size)))

	return bool(ret)
}

// Selectable_BoolPtr wraps igSelectable_BoolPtr.
func Selectable_BoolPtr(label ffi.CString, p_selected ffi.Ref[bool], flags GuiSelectableFlags, size Vec2) bool {
	ret := C.igSelectable_BoolPtr((*C.char)(label.Raw()), (*C.bool)(p_selected.Raw()), C.ImGuiSelectableFlags(flags), *(*C.ImVec2)(unsafe.Pointer(size)))

	return bool(ret)
}

// Separator wraps igSeparator.
func Separator() {
	C.igSeparator()
}

// SeparatorText wraps igSeparatorText.
func SeparatorText(label ffi.CString) {
	C.igSeparatorText((*C.char)(label.Raw()))
}

// igSetAllocatorFunctions.alloc_func is unsupported: category unsupported.

// SetClipboardText wraps igSetClipboardText.
func SetClipboardText(text ffi.CString) {
	C.igSetClipboardText((*C.char)(text.Raw()))
}

// SetColorEditOptions wraps igSetColorEditOptions.
func SetColorEditOptions(flags GuiColorEditFlags) {
	C.igSetColorEditOptions(C.ImGuiColorEditFlags(flags))
}

// SetColumnOffset wraps igSetColumnOffset.
func SetColumnOffset(column_index int32, offset_x float32) {
	C.igSetColumnOffset(C.int32_t(column_index), C.float(offset_x))
}

// SetColumnWidth wraps igSetColumnWidth.
func SetColumnWidth(column_index int32, width float32) {
	C.igSetColumnWidth(C.int32_t(column_index), C.float(width))
}

// SetCurrentContext wraps igSetCurrentContext.
func SetCurrentContext(ctx GuiContext) {
	C.igSetCurrentContext((*C.ImGuiContext)(unsafe.Pointer(ctx)))
}

// SetCursorPos wraps igSetCursorPos.
func SetCursorPos(local_pos Vec2) {
	C.igSetCursorPos(*(*C.ImVec2)(unsafe.Pointer(local_pos)))
}

// SetCursorPosX wraps igSetCursorPosX.
func SetCursorPosX(local_x float32) {
	C.igSetCursorPosX(C.float(local_x))
}

// SetCursorPosY wraps igSetCursorPosY.
func SetCursorPosY(local_y float32) {
	C.igSetCursorPosY(C.float(local_y))
}

// SetCursorScreenPos wraps igSetCursorScreenPos.
func SetCursorScreenPos(pos Vec2) {
	C.igSetCursorScreenPos(*(*C.ImVec2)(unsafe.Pointer(pos)))
}

// igSetDragDropPayload.sz is unsupported: category unsupported.

// SetItemDefaultFocus wraps igSetItemDefaultFocus.
func SetItemDefaultFocus() {
	C.igSetItemDefaultFocus()
}

// SetItemKeyOwner wraps igSetItemKeyOwner_Nil.
func SetItemKeyOwner(key GuiKey) {
	C.igSetItemKeyOwner_Nil(C.ImGuiKey(key))
}

// igSetItemTooltip.... is unsupported: category unsupported.

// igSetItemTooltipV.args is unsupported: category unsupported.

// SetKeyboardFocusHere wraps igSetKeyboardFocusHere.
func SetKeyboardFocusHere(offset int32) {
	C.igSetKeyboardFocusHere(C.int32_t(offset))
}

// SetMouseCursor wraps igSetMouseCursor.
func SetMouseCursor(cursor_type GuiMouseCursor) {
	C.igSetMouseCursor(C.ImGuiMouseCursor(cursor_type))
}

// SetNavCursorVisible wraps igSetNavCursorVisible.
func SetNavCursorVisible(visible bool) {
	C.igSetNavCursorVisible(C.bool(visible))
}

// SetNextFrameWantCaptureKeyboard wraps igSetNextFrameWantCaptureKeyboard.
func SetNextFrameWantCaptureKeyboard(want_capture_keyboard bool) {
	C.igSetNextFrameWantCaptureKeyboard(C.bool(want_capture_keyboard))
}

// SetNextFrameWantCaptureMouse wraps igSetNextFrameWantCaptureMouse.
func SetNextFrameWantCaptureMouse(want_capture_mouse bool) {
	C.igSetNextFrameWantCaptureMouse(C.bool(want_capture_mouse))
}

// SetNextItemAllowOverlap wraps igSetNextItemAllowOverlap.
func SetNextItemAllowOverlap() {
	C.igSetNextItemAllowOverlap()
}

// SetNextItemOpen wraps igSetNextItemOpen.
func SetNextItemOpen(is_open bool, cond GuiCond) {
	C.igSetNextItemOpen(C.bool(is_open), C.ImGuiCond(cond))
}

// igSetNextItemSelectionUserData.selection_user_data is unsupported: category unsupported.

// igSetNextItemShortcut.key_chord is unsupported: category unsupported.

// igSetNextItemStorageID.storage_id is unsupported: category unsupported.

// SetNextItemWidth wraps igSetNextItemWidth.
func SetNextItemWidth(item_width float32) {
	C.igSetNextItemWidth(C.float(item_width))
}

// SetNextWindowBgAlpha wraps igSetNextWindowBgAlpha.
func SetNextWindowBgAlpha(alpha float32) {
	C.igSetNextWindowBgAlpha(C.float(alpha))
}

// SetNextWindowClass wraps igSetNextWindowClass.
func SetNextWindowClass(window_class GuiWindowClass) {
	C.igSetNextWindowClass((*C.ImGuiWindowClass)(unsafe.Pointer(window_class)))
}

// SetNextWindowCollapsed wraps igSetNextWindowCollapsed.
func SetNextWindowCollapsed(collapsed bool, cond GuiCond) {
	C.igSetNextWindowCollapsed(C.bool(collapsed), C.ImGuiCond(cond))
}

// SetNextWindowContentSize wraps igSetNextWindowContentSize.
func SetNextWindowContentSize(size Vec2) {
	C.igSetNextWindowContentSize(*(*C.ImVec2)(unsafe.Pointer(size)))
}

// igSetNextWindowDockID.dock_id is unsupported: category unsupported.

// SetNextWindowFocus wraps igSetNextWindowFocus.
func SetNextWindowFocus() {
	C.igSetNextWindowFocus()
}

// SetNextWindowPos wraps igSetNextWindowPos.
func SetNextWindowPos(pos Vec2, cond GuiCond, pivot Vec2) {
	C.igSetNextWindowPos(*(*C.ImVec2)(unsafe.Pointer(pos)), C.ImGuiCond(cond), *(*C.ImVec2)(unsafe.Pointer(pivot)))
}

// SetNextWindowScroll wraps igSetNextWindowScroll.
func SetNextWindowScroll(scroll Vec2) {
	C.igSetNextWindowScroll(*(*C.ImVec2)(unsafe.Pointer(scroll)))
}

// SetNextWindowSize wraps igSetNextWindowSize.
func SetNextWindowSize(size Vec2, cond GuiCond) {
	C.igSetNextWindowSize(*(*C.ImVec2)(unsafe.Pointer(size)), C.ImGuiCond(cond))
}

// igSetNextWindowSizeConstraints.custom_callback is unsupported: category unsupported.

// igSetNextWindowViewport.viewport_id is unsupported: category unsupported.

// SetScrollFromPosX wraps igSetScrollFromPosX_Float.
func SetScrollFromPosX(local_x float32, center_x_ratio float32) {
	C.igSetScrollFromPosX_Float(C.float(local_x), C.float(center_x_ratio))
}

// SetScrollFromPosY wraps igSetScrollFromPosY_Float.
func SetScrollFromPosY(local_y float32, center_y_ratio float32) {
	C.igSetScrollFromPosY_Float(C.float(local_y), C.float(center_y_ratio))
}

// SetScrollHereX wraps igSetScrollHereX.
func SetScrollHereX(center_x_ratio float32) {
	C.igSetScrollHereX(C.float(center_x_ratio))
}

// SetScrollHereY wraps igSetScrollHereY.
func SetScrollHereY(center_y_ratio float32) {
	C.igSetScrollHereY(C.float(center_y_ratio))
}

// SetScrollX wraps igSetScrollX_Float.
func SetScrollX(scroll_x float32) {
	C.igSetScrollX_Float(C.float(scroll_x))
}

// SetScrollY wraps igSetScrollY_Float.
func SetScrollY(scroll_y float32) {
	C.igSetScrollY_Float(C.float(scroll_y))
}

// SetStateStorage wraps igSetStateStorage.
func SetStateStorage(storage GuiStorage) {
	C.igSetStateStorage((*C.ImGuiStorage)(unsafe.Pointer(storage)))
}

// SetTabItemClosed wraps igSetTabItemClosed.
func SetTabItemClosed(tab_or_docked_window_label ffi.CString) {
	C.igSetTabItemClosed((*C.char)(tab_or_docked_window_label.Raw()))
}

// igSetTooltip.... is unsupported: category unsupported.

// igSetTooltipV.args is unsupported: category unsupported.

// SetWindowCollapsed_Bool wraps igSetWindowCollapsed_Bool.
func SetWindowCollapsed_Bool(collapsed bool, cond GuiCond) {
	C.igSetWindowCollapsed_Bool(C.bool(collapsed), C.ImGuiCond(cond))
}

// SetWindowCollapsed_Str wraps igSetWindowCollapsed_Str.
func SetWindowCollapsed_Str(name ffi.CString, collapsed bool, cond GuiCond) {
	C.igSetWindowCollapsed_Str((*C.char)(name.Raw()), C.bool(collapsed), C.ImGuiCond(cond))
}

// SetWindowFocus_Nil wraps igSetWindowFocus_Nil.
func SetWindowFocus_Nil() {
	C.igSetWindowFocus_Nil()
}

// SetWindowFocus_Str wraps igSetWindowFocus_Str.
func SetWindowFocus_Str(name ffi.CString) {
	C.igSetWindowFocus_Str((*C.char)(name.Raw()))
}

// SetWindowPos_Str wraps igSetWindowPos_Str.
func SetWindowPos_Str(name ffi.CString, pos Vec2, cond GuiCond) {
	C.igSetWindowPos_Str((*C.char)(name.Raw()), *(*C.ImVec2)(unsafe.Pointer(pos)), C.ImGuiCond(cond))
}

// SetWindowPos_Vec2 wraps igSetWindowPos_Vec2.
func SetWindowPos_Vec2(pos Vec2, cond GuiCond) {
	C.igSetWindowPos_Vec2(*(*C.ImVec2)(unsafe.Pointer(pos)), C.ImGuiCond(cond))
}

// SetWindowSize_Str wraps igSetWindowSize_Str.
func SetWindowSize_Str(name ffi.CString, size Vec2, cond GuiCond) {
	C.igSetWindowSize_Str((*C.char)(name.Raw()), *(*C.ImVec2)(unsafe.Pointer(size)), C.ImGuiCond(cond))
}

// SetWindowSize_Vec2 wraps igSetWindowSize_Vec2.
func SetWindowSize_Vec2(size Vec2, cond GuiCond) {
	C.igSetWindowSize_Vec2(*(*C.ImVec2)(unsafe.Pointer(size)), C.ImGuiCond(cond))
}

// igShortcut_Nil.key_chord is unsupported: category unsupported.

// ShowAboutWindow wraps igShowAboutWindow.
func ShowAboutWindow(p_open ffi.Ref[bool]) {
	C.igShowAboutWindow((*C.bool)(p_open.Raw()))
}

// ShowDebugLogWindow wraps igShowDebugLogWindow.
func ShowDebugLogWindow(p_open ffi.Ref[bool]) {
	C.igShowDebugLogWindow((*C.bool)(p_open.Raw()))
}

// ShowDemoWindow wraps igShowDemoWindow.
func ShowDemoWindow(p_open ffi.Ref[bool]) {
	C.igShowDemoWindow((*C.bool)(p_open.Raw()))
}

// ShowFontSelector wraps igShowFontSelector.
func ShowFontSelector(label ffi.CString) {
	C.igShowFontSelector((*C.char)(label.Raw()))
}

// ShowIDStackToolWindow wraps igShowIDStackToolWindow.
func ShowIDStackToolWindow(p_open ffi.Ref[bool]) {
	C.igShowIDStackToolWindow((*C.bool)(p_open.Raw()))
}

// ShowMetricsWindow wraps igShowMetricsWindow.
func ShowMetricsWindow(p_open ffi.Ref[bool]) {
	C.igShowMetricsWindow((*C.bool)(p_open.Raw()))
}

// ShowStyleEditor wraps igShowStyleEditor.
func ShowStyleEditor(ref GuiStyle) {
	C.igShowStyleEditor((*C.ImGuiStyle)(unsafe.Pointer(ref)))
}

// ShowStyleSelector wraps igShowStyleSelector.
func ShowStyleSelector(label ffi.CString) bool {
	ret := C.igShowStyleSelector((*C.char)(label.Raw()))

	return bool(ret)
}

// ShowUserGuide wraps igShowUserGuide.
func ShowUserGuide() {
	C.igShowUserGuide()
}

// SliderAngle wraps igSliderAngle.
func SliderAngle(label ffi.CString, v_rad ffi.Ref[float32], v_degrees_min float32, v_degrees_max float32, format ffi.CString, flags GuiSliderFlags) bool {
	ret := C.igSliderAngle((*C.char)(label.Raw()), (*C.float)(v_rad.Raw()), C.float(v_degrees_min), C.float(v_degrees_max), (*C.char)(format.Raw()), C.ImGuiSliderFlags(flags))

	return bool(ret)
}

// SliderFloat wraps igSliderFloat.
func SliderFloat(label ffi.CString, v ffi.Ref[float32], v_min float32, v_max float32, format ffi.CString, flags GuiSliderFlags) bool {
	ret := C.igSliderFloat((*C.char)(label.Raw()), (*C.float)(v.Raw()), C.float(v_min), C.float(v_max), (*C.char)(format.Raw()), C.ImGuiSliderFlags(flags))

	return bool(ret)
}

// igSliderFloat2.v is unsupported: category unsupported.

// igSliderFloat3.v is unsupported: category unsupported.

// igSliderFloat4.v is unsupported: category unsupported.

// SliderInt wraps igSliderInt.
func SliderInt(label ffi.CString, v ffi.Ref[int32], v_min int32, v_max int32, format ffi.CString, flags GuiSliderFlags) bool {
	ret := C.igSliderInt((*C.char)(label.Raw()), (*C.int32_t)(v.Raw()), C.int32_t(v_min), C.int32_t(v_max), (*C.char)(format.Raw()), C.ImGuiSliderFlags(flags))

	return bool(ret)
}

// igSliderInt2.v is unsupported: category unsupported.

// igSliderInt3.v is unsupported: category unsupported.

// igSliderInt4.v is unsupported: category unsupported.

// SliderScalar wraps igSliderScalar.
func SliderScalar(label ffi.CString, data_type GuiDataType, p_data uintptr, p_min uintptr, p_max uintptr, format ffi.CString, flags GuiSliderFlags) bool {
	ret := C.igSliderScalar((*C.char)(label.Raw()), C.ImGuiDataType(data_type), unsafe.Pointer(p_data), unsafe.Pointer(p_min), unsafe.Pointer(p_max), (*C.char)(format.Raw()), C.ImGuiSliderFlags(flags))

	return bool(ret)
}

// SliderScalarN wraps igSliderScalarN.
func SliderScalarN(label ffi.CString, data_type GuiDataType, p_data uintptr, components int32, p_min uintptr, p_max uintptr, format ffi.CString, flags GuiSliderFlags) bool {
	ret := C.igSliderScalarN((*C.char)(label.Raw()), C.ImGuiDataType(data_type), unsafe.Pointer(p_data), C.int32_t(components), unsafe.Pointer(p_min), unsafe.Pointer(p_max), (*C.char)(format.Raw()), C.ImGuiSliderFlags(flags))

	return bool(ret)
}

// SmallButton wraps igSmallButton.
func SmallButton(label ffi.CString) bool {
	ret := C.igSmallButton((*C.char)(label.Raw()))

	return bool(ret)
}

// Spacing wraps igSpacing.
func Spacing() {
	C.igSpacing()
}

// StyleColorsClassic wraps igStyleColorsClassic.
func StyleColorsClassic(dst GuiStyle) {
	C.igStyleColorsClassic((*C.ImGuiStyle)(unsafe.Pointer(dst)))
}

// StyleColorsDark wraps igStyleColorsDark.
func StyleColorsDark(dst GuiStyle) {
	C.igStyleColorsDark((*C.ImGuiStyle)(unsafe.Pointer(dst)))
}

// StyleColorsLight wraps igStyleColorsLight.
func StyleColorsLight(dst GuiStyle) {
	C.igStyleColorsLight((*C.ImGuiStyle)(unsafe.Pointer(dst)))
}

// TabItemButton wraps igTabItemButton.
func TabItemButton(label ffi.CString, flags GuiTabItemFlags) bool {
	ret := C.igTabItemButton((*C.char)(label.Raw()), C.ImGuiTabItemFlags(flags))

	return bool(ret)
}

// TableAngledHeadersRow wraps igTableAngledHeadersRow.
func TableAngledHeadersRow() {
	C.igTableAngledHeadersRow()
}

// TableGetColumnCount wraps igTableGetColumnCount.
func TableGetColumnCount() int32 {
	ret := C.igTableGetColumnCount()

	return int32(ret)
}

// igTableGetColumnFlags is unsupported: unknown category bitmask.

// TableGetColumnIndex wraps igTableGetColumnIndex.
func TableGetColumnIndex() int32 {
	ret := C.igTableGetColumnIndex()

	return int32(ret)
}

// igTableGetColumnName_Int is unsupported: return pointer type -> ?? char.

// TableGetHoveredColumn wraps igTableGetHoveredColumn.
func TableGetHoveredColumn() int32 {
	ret := C.igTableGetHoveredColumn()

	return int32(ret)
}

// TableGetRowIndex wraps igTableGetRowIndex.
func TableGetRowIndex() int32 {
	ret := C.igTableGetRowIndex()

	return int32(ret)
}

// TableGetSortSpecs wraps igTableGetSortSpecs.
func TableGetSortSpecs() GuiTableSortSpecs {
	ret := C.igTableGetSortSpecs()

	return GuiTableSortSpecs(unsafe.Pointer(ret))
}

// TableHeader wraps igTableHeader.
func TableHeader(label ffi.CString) {
	C.igTableHeader((*C.char)(label.Raw()))
}

// TableHeadersRow wraps igTableHeadersRow.
func TableHeadersRow() {
	C.igTableHeadersRow()
}

// TableNextColumn wraps igTableNextColumn.
func TableNextColumn() bool {
	ret := C.igTableNextColumn()

	return bool(ret)
}

// TableNextRow wraps igTableNextRow.
func TableNextRow(row_flags GuiTableRowFlags, min_row_height float32) {
	C.igTableNextRow(C.ImGuiTableRowFlags(row_flags), C.float(min_row_height))
}

// TableSetBgColor wraps igTableSetBgColor.
func TableSetBgColor(target GuiTableBgTarget, color uint32, column_n int32) {
	C.igTableSetBgColor(C.ImGuiTableBgTarget(target), C.uint32_t(color), C.int32_t(column_n))
}

// TableSetColumnEnabled wraps igTableSetColumnEnabled.
func TableSetColumnEnabled(column_n int32, v bool) {
	C.igTableSetColumnEnabled(C.int32_t(column_n), C.bool(v))
}

// TableSetColumnIndex wraps igTableSetColumnIndex.
func TableSetColumnIndex(column_n int32) bool {
	ret := C.igTableSetColumnIndex(C.int32_t(column_n))

	return bool(ret)
}

// igTableSetupColumn.user_id is unsupported: category unsupported.

// TableSetupScrollFreeze wraps igTableSetupScrollFreeze.
func TableSetupScrollFreeze(cols int32, rows int32) {
	C.igTableSetupScrollFreeze(C.int32_t(cols), C.int32_t(rows))
}

// igText.... is unsupported: category unsupported.

// igTextColored.... is unsupported: category unsupported.

// igTextColoredV.args is unsupported: category unsupported.

// igTextDisabled.... is unsupported: category unsupported.

// igTextDisabledV.args is unsupported: category unsupported.

// TextLink wraps igTextLink.
func TextLink(label ffi.CString) bool {
	ret := C.igTextLink((*C.char)(label.Raw()))

	return bool(ret)
}

// TextLinkOpenURL wraps igTextLinkOpenURL.
func TextLinkOpenURL(label ffi.CString, url ffi.CString) bool {
	ret := C.igTextLinkOpenURL((*C.char)(label.Raw()), (*C.char)(url.Raw()))

	return bool(ret)
}

// TextUnformatted wraps igTextUnformatted.
func TextUnformatted(text ffi.CString, text_end ffi.CString) {
	C.igTextUnformatted((*C.char)(text.Raw()), (*C.char)(text_end.Raw()))
}

// igTextV.args is unsupported: category unsupported.

// igTextWrapped.... is unsupported: category unsupported.

// igTextWrappedV.args is unsupported: category unsupported.

// igTreeNodeExV_Ptr.args is unsupported: category unsupported.

// igTreeNodeExV_Str.args is unsupported: category unsupported.

// igTreeNodeEx_Ptr.... is unsupported: category unsupported.

// TreeNodeEx_Str wraps igTreeNodeEx_Str.
func TreeNodeEx_Str(label ffi.CString, flags GuiTreeNodeFlags) bool {
	ret := C.igTreeNodeEx_Str((*C.char)(label.Raw()), C.ImGuiTreeNodeFlags(flags))

	return bool(ret)
}

// igTreeNodeEx_StrStr.... is unsupported: category unsupported.

// igTreeNodeV_Ptr.args is unsupported: category unsupported.

// igTreeNodeV_Str.args is unsupported: category unsupported.

// igTreeNode_Ptr.... is unsupported: category unsupported.

// TreeNode_Str wraps igTreeNode_Str.
func TreeNode_Str(label ffi.CString) bool {
	ret := C.igTreeNode_Str((*C.char)(label.Raw()))

	return bool(ret)
}

// igTreeNode_StrStr.... is unsupported: category unsupported.

// TreePop wraps igTreePop.
func TreePop() {
	C.igTreePop()
}

// TreePush_Ptr wraps igTreePush_Ptr.
func TreePush_Ptr(ptr_id uintptr) {
	C.igTreePush_Ptr(unsafe.Pointer(ptr_id))
}

// TreePush_Str wraps igTreePush_Str.
func TreePush_Str(str_id ffi.CString) {
	C.igTreePush_Str((*C.char)(str_id.Raw()))
}

// Unindent wraps igUnindent.
func Unindent(indent_w float32) {
	C.igUnindent(C.float(indent_w))
}

// UpdatePlatformWindows wraps igUpdatePlatformWindows.
func UpdatePlatformWindows() {
	C.igUpdatePlatformWindows()
}

// VSliderFloat wraps igVSliderFloat.
func VSliderFloat(label ffi.CString, size Vec2, v ffi.Ref[float32], v_min float32, v_max float32, format ffi.CString, flags GuiSliderFlags) bool {
	ret := C.igVSliderFloat((*C.char)(label.Raw()), *(*C.ImVec2)(unsafe.Pointer(size)), (*C.float)(v.Raw()), C.float(v_min), C.float(v_max), (*C.char)(format.Raw()), C.ImGuiSliderFlags(flags))

	return bool(ret)
}

// VSliderInt wraps igVSliderInt.
func VSliderInt(label ffi.CString, size Vec2, v ffi.Ref[int32], v_min int32, v_max int32, format ffi.CString, flags GuiSliderFlags) bool {
	ret := C.igVSliderInt((*C.char)(label.Raw()), *(*C.ImVec2)(unsafe.Pointer(size)), (*C.int32_t)(v.Raw()), C.int32_t(v_min), C.int32_t(v_max), (*C.char)(format.Raw()), C.ImGuiSliderFlags(flags))

	return bool(ret)
}

// VSliderScalar wraps igVSliderScalar.
func VSliderScalar(label ffi.CString, size Vec2, data_type GuiDataType, p_data uintptr, p_min uintptr, p_max uintptr, format ffi.CString, flags GuiSliderFlags) bool {
	ret := C.igVSliderScalar((*C.char)(label.Raw()), *(*C.ImVec2)(unsafe.Pointer(size)), C.ImGuiDataType(data_type), unsafe.Pointer(p_data), unsafe.Pointer(p_min), unsafe.Pointer(p_max), (*C.char)(format.Raw()), C.ImGuiSliderFlags(flags))

	return bool(ret)
}

// Value_Bool wraps igValue_Bool.
func Value_Bool(prefix ffi.CString, b bool) {
	C.igValue_Bool((*C.char)(prefix.Raw()), C.bool(b))
}

// Value_Float wraps igValue_Float.
func Value_Float(prefix ffi.CString, v float32, float_format ffi.CString) {
	C.igValue_Float((*C.char)(prefix.Raw()), C.float(v), (*C.char)(float_format.Raw()))
}

// Value_Int wraps igValue_Int.
func Value_Int(prefix ffi.CString, v int32) {
	C.igValue_Int((*C.char)(prefix.Raw()), C.int32_t(v))
}

// Value_Uint wraps igValue_Uint.
func Value_Uint(prefix ffi.CString, v uint32) {
	C.igValue_Uint((*C.char)(prefix.Raw()), C.uint32_t(v))
}
