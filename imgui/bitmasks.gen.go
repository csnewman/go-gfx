package imgui

import "fmt"

// #include "imgui_wrapper.h"
import "C"

// DrawFlags wraps the bitmask ImDrawFlags.
type DrawFlags int32

const (
	// DrawFlags_Closed wraps ImDrawFlags_Closed.
	DrawFlags_Closed DrawFlags = C.ImDrawFlags_Closed
	// DrawFlags_None wraps ImDrawFlags_None.
	DrawFlags_None DrawFlags = C.ImDrawFlags_None
	// DrawFlags_RoundCornersAll wraps ImDrawFlags_RoundCornersAll.
	DrawFlags_RoundCornersAll DrawFlags = C.ImDrawFlags_RoundCornersAll
	// DrawFlags_RoundCornersBottom wraps ImDrawFlags_RoundCornersBottom.
	DrawFlags_RoundCornersBottom DrawFlags = C.ImDrawFlags_RoundCornersBottom
	// DrawFlags_RoundCornersBottomLeft wraps ImDrawFlags_RoundCornersBottomLeft.
	DrawFlags_RoundCornersBottomLeft DrawFlags = C.ImDrawFlags_RoundCornersBottomLeft
	// DrawFlags_RoundCornersBottomRight wraps ImDrawFlags_RoundCornersBottomRight.
	DrawFlags_RoundCornersBottomRight DrawFlags = C.ImDrawFlags_RoundCornersBottomRight
	// DrawFlags_RoundCornersDefault wraps ImDrawFlags_RoundCornersDefault_.
	DrawFlags_RoundCornersDefault DrawFlags = C.ImDrawFlags_RoundCornersDefault_
	// DrawFlags_RoundCornersLeft wraps ImDrawFlags_RoundCornersLeft.
	DrawFlags_RoundCornersLeft DrawFlags = C.ImDrawFlags_RoundCornersLeft
	// DrawFlags_RoundCornersMask wraps ImDrawFlags_RoundCornersMask_.
	DrawFlags_RoundCornersMask DrawFlags = C.ImDrawFlags_RoundCornersMask_
	// DrawFlags_RoundCornersNone wraps ImDrawFlags_RoundCornersNone.
	DrawFlags_RoundCornersNone DrawFlags = C.ImDrawFlags_RoundCornersNone
	// DrawFlags_RoundCornersRight wraps ImDrawFlags_RoundCornersRight.
	DrawFlags_RoundCornersRight DrawFlags = C.ImDrawFlags_RoundCornersRight
	// DrawFlags_RoundCornersTop wraps ImDrawFlags_RoundCornersTop.
	DrawFlags_RoundCornersTop DrawFlags = C.ImDrawFlags_RoundCornersTop
	// DrawFlags_RoundCornersTopLeft wraps ImDrawFlags_RoundCornersTopLeft.
	DrawFlags_RoundCornersTopLeft DrawFlags = C.ImDrawFlags_RoundCornersTopLeft
	// DrawFlags_RoundCornersTopRight wraps ImDrawFlags_RoundCornersTopRight.
	DrawFlags_RoundCornersTopRight DrawFlags = C.ImDrawFlags_RoundCornersTopRight
)

func (e DrawFlags) String() string {
	return fmt.Sprintf("ImDrawFlags(%b)", e)
}

// DrawListFlags wraps the bitmask ImDrawListFlags.
type DrawListFlags int32

const (
	// DrawListFlags_AllowVtxOffset wraps ImDrawListFlags_AllowVtxOffset.
	DrawListFlags_AllowVtxOffset DrawListFlags = C.ImDrawListFlags_AllowVtxOffset
	// DrawListFlags_AntiAliasedFill wraps ImDrawListFlags_AntiAliasedFill.
	DrawListFlags_AntiAliasedFill DrawListFlags = C.ImDrawListFlags_AntiAliasedFill
	// DrawListFlags_AntiAliasedLines wraps ImDrawListFlags_AntiAliasedLines.
	DrawListFlags_AntiAliasedLines DrawListFlags = C.ImDrawListFlags_AntiAliasedLines
	// DrawListFlags_AntiAliasedLinesUseTex wraps ImDrawListFlags_AntiAliasedLinesUseTex.
	DrawListFlags_AntiAliasedLinesUseTex DrawListFlags = C.ImDrawListFlags_AntiAliasedLinesUseTex
	// DrawListFlags_None wraps ImDrawListFlags_None.
	DrawListFlags_None DrawListFlags = C.ImDrawListFlags_None
)

func (e DrawListFlags) String() string {
	return fmt.Sprintf("ImDrawListFlags(%b)", e)
}

// DrawTextFlags wraps the bitmask ImDrawTextFlags.
type DrawTextFlags int32

const (
	// DrawTextFlags_CpuFineClip wraps ImDrawTextFlags_CpuFineClip.
	DrawTextFlags_CpuFineClip DrawTextFlags = C.ImDrawTextFlags_CpuFineClip
	// DrawTextFlags_None wraps ImDrawTextFlags_None.
	DrawTextFlags_None DrawTextFlags = C.ImDrawTextFlags_None
	// DrawTextFlags_StopOnNewLine wraps ImDrawTextFlags_StopOnNewLine.
	DrawTextFlags_StopOnNewLine DrawTextFlags = C.ImDrawTextFlags_StopOnNewLine
	// DrawTextFlags_WrapKeepBlanks wraps ImDrawTextFlags_WrapKeepBlanks.
	DrawTextFlags_WrapKeepBlanks DrawTextFlags = C.ImDrawTextFlags_WrapKeepBlanks
)

func (e DrawTextFlags) String() string {
	return fmt.Sprintf("ImDrawTextFlags(%b)", e)
}

// FontAtlasFlags wraps the bitmask ImFontAtlasFlags.
type FontAtlasFlags int32

const (
	// FontAtlasFlags_NoBakedLines wraps ImFontAtlasFlags_NoBakedLines.
	FontAtlasFlags_NoBakedLines FontAtlasFlags = C.ImFontAtlasFlags_NoBakedLines
	// FontAtlasFlags_NoMouseCursors wraps ImFontAtlasFlags_NoMouseCursors.
	FontAtlasFlags_NoMouseCursors FontAtlasFlags = C.ImFontAtlasFlags_NoMouseCursors
	// FontAtlasFlags_NoPowerOfTwoHeight wraps ImFontAtlasFlags_NoPowerOfTwoHeight.
	FontAtlasFlags_NoPowerOfTwoHeight FontAtlasFlags = C.ImFontAtlasFlags_NoPowerOfTwoHeight
	// FontAtlasFlags_None wraps ImFontAtlasFlags_None.
	FontAtlasFlags_None FontAtlasFlags = C.ImFontAtlasFlags_None
)

func (e FontAtlasFlags) String() string {
	return fmt.Sprintf("ImFontAtlasFlags(%b)", e)
}

// FontFlags wraps the bitmask ImFontFlags.
type FontFlags int32

const (
	// FontFlags_LockBakedSizes wraps ImFontFlags_LockBakedSizes.
	FontFlags_LockBakedSizes FontFlags = C.ImFontFlags_LockBakedSizes
	// FontFlags_NoLoadError wraps ImFontFlags_NoLoadError.
	FontFlags_NoLoadError FontFlags = C.ImFontFlags_NoLoadError
	// FontFlags_NoLoadGlyphs wraps ImFontFlags_NoLoadGlyphs.
	FontFlags_NoLoadGlyphs FontFlags = C.ImFontFlags_NoLoadGlyphs
	// FontFlags_None wraps ImFontFlags_None.
	FontFlags_None FontFlags = C.ImFontFlags_None
)

func (e FontFlags) String() string {
	return fmt.Sprintf("ImFontFlags(%b)", e)
}

// GuiActivateFlags wraps the bitmask ImGuiActivateFlags.
type GuiActivateFlags int32

const (
	// GuiActivateFlags_FromFocusApi wraps ImGuiActivateFlags_FromFocusApi.
	GuiActivateFlags_FromFocusApi GuiActivateFlags = C.ImGuiActivateFlags_FromFocusApi
	// GuiActivateFlags_FromShortcut wraps ImGuiActivateFlags_FromShortcut.
	GuiActivateFlags_FromShortcut GuiActivateFlags = C.ImGuiActivateFlags_FromShortcut
	// GuiActivateFlags_FromTabbing wraps ImGuiActivateFlags_FromTabbing.
	GuiActivateFlags_FromTabbing GuiActivateFlags = C.ImGuiActivateFlags_FromTabbing
	// GuiActivateFlags_None wraps ImGuiActivateFlags_None.
	GuiActivateFlags_None GuiActivateFlags = C.ImGuiActivateFlags_None
	// GuiActivateFlags_PreferInput wraps ImGuiActivateFlags_PreferInput.
	GuiActivateFlags_PreferInput GuiActivateFlags = C.ImGuiActivateFlags_PreferInput
	// GuiActivateFlags_PreferTweak wraps ImGuiActivateFlags_PreferTweak.
	GuiActivateFlags_PreferTweak GuiActivateFlags = C.ImGuiActivateFlags_PreferTweak
	// GuiActivateFlags_TryToPreserveState wraps ImGuiActivateFlags_TryToPreserveState.
	GuiActivateFlags_TryToPreserveState GuiActivateFlags = C.ImGuiActivateFlags_TryToPreserveState
)

func (e GuiActivateFlags) String() string {
	return fmt.Sprintf("ImGuiActivateFlags(%b)", e)
}

// GuiBackendFlags wraps the bitmask ImGuiBackendFlags.
type GuiBackendFlags int32

const (
	// GuiBackendFlags_HasGamepad wraps ImGuiBackendFlags_HasGamepad.
	GuiBackendFlags_HasGamepad GuiBackendFlags = C.ImGuiBackendFlags_HasGamepad
	// GuiBackendFlags_HasMouseCursors wraps ImGuiBackendFlags_HasMouseCursors.
	GuiBackendFlags_HasMouseCursors GuiBackendFlags = C.ImGuiBackendFlags_HasMouseCursors
	// GuiBackendFlags_HasMouseHoveredViewport wraps ImGuiBackendFlags_HasMouseHoveredViewport.
	GuiBackendFlags_HasMouseHoveredViewport GuiBackendFlags = C.ImGuiBackendFlags_HasMouseHoveredViewport
	// GuiBackendFlags_HasParentViewport wraps ImGuiBackendFlags_HasParentViewport.
	GuiBackendFlags_HasParentViewport GuiBackendFlags = C.ImGuiBackendFlags_HasParentViewport
	// GuiBackendFlags_HasSetMousePos wraps ImGuiBackendFlags_HasSetMousePos.
	GuiBackendFlags_HasSetMousePos GuiBackendFlags = C.ImGuiBackendFlags_HasSetMousePos
	// GuiBackendFlags_None wraps ImGuiBackendFlags_None.
	GuiBackendFlags_None GuiBackendFlags = C.ImGuiBackendFlags_None
	// GuiBackendFlags_PlatformHasViewports wraps ImGuiBackendFlags_PlatformHasViewports.
	GuiBackendFlags_PlatformHasViewports GuiBackendFlags = C.ImGuiBackendFlags_PlatformHasViewports
	// GuiBackendFlags_RendererHasTextures wraps ImGuiBackendFlags_RendererHasTextures.
	GuiBackendFlags_RendererHasTextures GuiBackendFlags = C.ImGuiBackendFlags_RendererHasTextures
	// GuiBackendFlags_RendererHasViewports wraps ImGuiBackendFlags_RendererHasViewports.
	GuiBackendFlags_RendererHasViewports GuiBackendFlags = C.ImGuiBackendFlags_RendererHasViewports
	// GuiBackendFlags_RendererHasVtxOffset wraps ImGuiBackendFlags_RendererHasVtxOffset.
	GuiBackendFlags_RendererHasVtxOffset GuiBackendFlags = C.ImGuiBackendFlags_RendererHasVtxOffset
)

func (e GuiBackendFlags) String() string {
	return fmt.Sprintf("ImGuiBackendFlags(%b)", e)
}

// GuiButtonFlags wraps the bitmask ImGuiButtonFlags.
type GuiButtonFlags int32

const (
	// GuiButtonFlags_EnableNav wraps ImGuiButtonFlags_EnableNav.
	GuiButtonFlags_EnableNav GuiButtonFlags = C.ImGuiButtonFlags_EnableNav
	// GuiButtonFlags_MouseButtonLeft wraps ImGuiButtonFlags_MouseButtonLeft.
	GuiButtonFlags_MouseButtonLeft GuiButtonFlags = C.ImGuiButtonFlags_MouseButtonLeft
	// GuiButtonFlags_MouseButtonMask wraps ImGuiButtonFlags_MouseButtonMask_.
	GuiButtonFlags_MouseButtonMask GuiButtonFlags = C.ImGuiButtonFlags_MouseButtonMask_
	// GuiButtonFlags_MouseButtonMiddle wraps ImGuiButtonFlags_MouseButtonMiddle.
	GuiButtonFlags_MouseButtonMiddle GuiButtonFlags = C.ImGuiButtonFlags_MouseButtonMiddle
	// GuiButtonFlags_MouseButtonRight wraps ImGuiButtonFlags_MouseButtonRight.
	GuiButtonFlags_MouseButtonRight GuiButtonFlags = C.ImGuiButtonFlags_MouseButtonRight
	// GuiButtonFlags_None wraps ImGuiButtonFlags_None.
	GuiButtonFlags_None GuiButtonFlags = C.ImGuiButtonFlags_None
)

func (e GuiButtonFlags) String() string {
	return fmt.Sprintf("ImGuiButtonFlags(%b)", e)
}

// GuiButtonFlagsPrivate wraps the bitmask ImGuiButtonFlagsPrivate.
type GuiButtonFlagsPrivate int32

const (
	// GuiButtonFlags_AlignTextBaseLine wraps ImGuiButtonFlags_AlignTextBaseLine.
	GuiButtonFlags_AlignTextBaseLine GuiButtonFlagsPrivate = C.ImGuiButtonFlags_AlignTextBaseLine
	// GuiButtonFlags_AllowOverlap wraps ImGuiButtonFlags_AllowOverlap.
	GuiButtonFlags_AllowOverlap GuiButtonFlagsPrivate = C.ImGuiButtonFlags_AllowOverlap
	// GuiButtonFlags_FlattenChildren wraps ImGuiButtonFlags_FlattenChildren.
	GuiButtonFlags_FlattenChildren GuiButtonFlagsPrivate = C.ImGuiButtonFlags_FlattenChildren
	// GuiButtonFlags_NoFocus wraps ImGuiButtonFlags_NoFocus.
	GuiButtonFlags_NoFocus GuiButtonFlagsPrivate = C.ImGuiButtonFlags_NoFocus
	// GuiButtonFlags_NoHoldingActiveId wraps ImGuiButtonFlags_NoHoldingActiveId.
	GuiButtonFlags_NoHoldingActiveId GuiButtonFlagsPrivate = C.ImGuiButtonFlags_NoHoldingActiveId
	// GuiButtonFlags_NoHoveredOnFocus wraps ImGuiButtonFlags_NoHoveredOnFocus.
	GuiButtonFlags_NoHoveredOnFocus GuiButtonFlagsPrivate = C.ImGuiButtonFlags_NoHoveredOnFocus
	// GuiButtonFlags_NoKeyModsAllowed wraps ImGuiButtonFlags_NoKeyModsAllowed.
	GuiButtonFlags_NoKeyModsAllowed GuiButtonFlagsPrivate = C.ImGuiButtonFlags_NoKeyModsAllowed
	// GuiButtonFlags_NoNavFocus wraps ImGuiButtonFlags_NoNavFocus.
	GuiButtonFlags_NoNavFocus GuiButtonFlagsPrivate = C.ImGuiButtonFlags_NoNavFocus
	// GuiButtonFlags_NoSetKeyOwner wraps ImGuiButtonFlags_NoSetKeyOwner.
	GuiButtonFlags_NoSetKeyOwner GuiButtonFlagsPrivate = C.ImGuiButtonFlags_NoSetKeyOwner
	// GuiButtonFlags_NoTestKeyOwner wraps ImGuiButtonFlags_NoTestKeyOwner.
	GuiButtonFlags_NoTestKeyOwner GuiButtonFlagsPrivate = C.ImGuiButtonFlags_NoTestKeyOwner
	// GuiButtonFlags_PressedOnClick wraps ImGuiButtonFlags_PressedOnClick.
	GuiButtonFlags_PressedOnClick GuiButtonFlagsPrivate = C.ImGuiButtonFlags_PressedOnClick
	// GuiButtonFlags_PressedOnClickRelease wraps ImGuiButtonFlags_PressedOnClickRelease.
	GuiButtonFlags_PressedOnClickRelease GuiButtonFlagsPrivate = C.ImGuiButtonFlags_PressedOnClickRelease
	// GuiButtonFlags_PressedOnClickReleaseAnywhere wraps ImGuiButtonFlags_PressedOnClickReleaseAnywhere.
	GuiButtonFlags_PressedOnClickReleaseAnywhere GuiButtonFlagsPrivate = C.ImGuiButtonFlags_PressedOnClickReleaseAnywhere
	// GuiButtonFlags_PressedOnDefault wraps ImGuiButtonFlags_PressedOnDefault_.
	GuiButtonFlags_PressedOnDefault GuiButtonFlagsPrivate = C.ImGuiButtonFlags_PressedOnDefault_
	// GuiButtonFlags_PressedOnDoubleClick wraps ImGuiButtonFlags_PressedOnDoubleClick.
	GuiButtonFlags_PressedOnDoubleClick GuiButtonFlagsPrivate = C.ImGuiButtonFlags_PressedOnDoubleClick
	// GuiButtonFlags_PressedOnDragDropHold wraps ImGuiButtonFlags_PressedOnDragDropHold.
	GuiButtonFlags_PressedOnDragDropHold GuiButtonFlagsPrivate = C.ImGuiButtonFlags_PressedOnDragDropHold
	// GuiButtonFlags_PressedOnMask wraps ImGuiButtonFlags_PressedOnMask_.
	GuiButtonFlags_PressedOnMask GuiButtonFlagsPrivate = C.ImGuiButtonFlags_PressedOnMask_
	// GuiButtonFlags_PressedOnRelease wraps ImGuiButtonFlags_PressedOnRelease.
	GuiButtonFlags_PressedOnRelease GuiButtonFlagsPrivate = C.ImGuiButtonFlags_PressedOnRelease
)

func (e GuiButtonFlagsPrivate) String() string {
	return fmt.Sprintf("ImGuiButtonFlagsPrivate(%b)", e)
}

// GuiChildFlags wraps the bitmask ImGuiChildFlags.
type GuiChildFlags int32

const (
	// GuiChildFlags_AlwaysAutoResize wraps ImGuiChildFlags_AlwaysAutoResize.
	GuiChildFlags_AlwaysAutoResize GuiChildFlags = C.ImGuiChildFlags_AlwaysAutoResize
	// GuiChildFlags_AlwaysUseWindowPadding wraps ImGuiChildFlags_AlwaysUseWindowPadding.
	GuiChildFlags_AlwaysUseWindowPadding GuiChildFlags = C.ImGuiChildFlags_AlwaysUseWindowPadding
	// GuiChildFlags_AutoResizeX wraps ImGuiChildFlags_AutoResizeX.
	GuiChildFlags_AutoResizeX GuiChildFlags = C.ImGuiChildFlags_AutoResizeX
	// GuiChildFlags_AutoResizeY wraps ImGuiChildFlags_AutoResizeY.
	GuiChildFlags_AutoResizeY GuiChildFlags = C.ImGuiChildFlags_AutoResizeY
	// GuiChildFlags_Borders wraps ImGuiChildFlags_Borders.
	GuiChildFlags_Borders GuiChildFlags = C.ImGuiChildFlags_Borders
	// GuiChildFlags_FrameStyle wraps ImGuiChildFlags_FrameStyle.
	GuiChildFlags_FrameStyle GuiChildFlags = C.ImGuiChildFlags_FrameStyle
	// GuiChildFlags_NavFlattened wraps ImGuiChildFlags_NavFlattened.
	GuiChildFlags_NavFlattened GuiChildFlags = C.ImGuiChildFlags_NavFlattened
	// GuiChildFlags_None wraps ImGuiChildFlags_None.
	GuiChildFlags_None GuiChildFlags = C.ImGuiChildFlags_None
	// GuiChildFlags_ResizeX wraps ImGuiChildFlags_ResizeX.
	GuiChildFlags_ResizeX GuiChildFlags = C.ImGuiChildFlags_ResizeX
	// GuiChildFlags_ResizeY wraps ImGuiChildFlags_ResizeY.
	GuiChildFlags_ResizeY GuiChildFlags = C.ImGuiChildFlags_ResizeY
)

func (e GuiChildFlags) String() string {
	return fmt.Sprintf("ImGuiChildFlags(%b)", e)
}

// GuiColorEditFlags wraps the bitmask ImGuiColorEditFlags.
type GuiColorEditFlags int32

const (
	// GuiColorEditFlags_AlphaBar wraps ImGuiColorEditFlags_AlphaBar.
	GuiColorEditFlags_AlphaBar GuiColorEditFlags = C.ImGuiColorEditFlags_AlphaBar
	// GuiColorEditFlags_AlphaMask wraps ImGuiColorEditFlags_AlphaMask_.
	GuiColorEditFlags_AlphaMask GuiColorEditFlags = C.ImGuiColorEditFlags_AlphaMask_
	// GuiColorEditFlags_AlphaNoBg wraps ImGuiColorEditFlags_AlphaNoBg.
	GuiColorEditFlags_AlphaNoBg GuiColorEditFlags = C.ImGuiColorEditFlags_AlphaNoBg
	// GuiColorEditFlags_AlphaOpaque wraps ImGuiColorEditFlags_AlphaOpaque.
	GuiColorEditFlags_AlphaOpaque GuiColorEditFlags = C.ImGuiColorEditFlags_AlphaOpaque
	// GuiColorEditFlags_AlphaPreviewHalf wraps ImGuiColorEditFlags_AlphaPreviewHalf.
	GuiColorEditFlags_AlphaPreviewHalf GuiColorEditFlags = C.ImGuiColorEditFlags_AlphaPreviewHalf
	// GuiColorEditFlags_DataTypeMask wraps ImGuiColorEditFlags_DataTypeMask_.
	GuiColorEditFlags_DataTypeMask GuiColorEditFlags = C.ImGuiColorEditFlags_DataTypeMask_
	// GuiColorEditFlags_DefaultOptions wraps ImGuiColorEditFlags_DefaultOptions_.
	GuiColorEditFlags_DefaultOptions GuiColorEditFlags = C.ImGuiColorEditFlags_DefaultOptions_
	// GuiColorEditFlags_DisplayHSV wraps ImGuiColorEditFlags_DisplayHSV.
	GuiColorEditFlags_DisplayHSV GuiColorEditFlags = C.ImGuiColorEditFlags_DisplayHSV
	// GuiColorEditFlags_DisplayHex wraps ImGuiColorEditFlags_DisplayHex.
	GuiColorEditFlags_DisplayHex GuiColorEditFlags = C.ImGuiColorEditFlags_DisplayHex
	// GuiColorEditFlags_DisplayMask wraps ImGuiColorEditFlags_DisplayMask_.
	GuiColorEditFlags_DisplayMask GuiColorEditFlags = C.ImGuiColorEditFlags_DisplayMask_
	// GuiColorEditFlags_DisplayRGB wraps ImGuiColorEditFlags_DisplayRGB.
	GuiColorEditFlags_DisplayRGB GuiColorEditFlags = C.ImGuiColorEditFlags_DisplayRGB
	// GuiColorEditFlags_Float wraps ImGuiColorEditFlags_Float.
	GuiColorEditFlags_Float GuiColorEditFlags = C.ImGuiColorEditFlags_Float
	// GuiColorEditFlags_HDR wraps ImGuiColorEditFlags_HDR.
	GuiColorEditFlags_HDR GuiColorEditFlags = C.ImGuiColorEditFlags_HDR
	// GuiColorEditFlags_InputHSV wraps ImGuiColorEditFlags_InputHSV.
	GuiColorEditFlags_InputHSV GuiColorEditFlags = C.ImGuiColorEditFlags_InputHSV
	// GuiColorEditFlags_InputMask wraps ImGuiColorEditFlags_InputMask_.
	GuiColorEditFlags_InputMask GuiColorEditFlags = C.ImGuiColorEditFlags_InputMask_
	// GuiColorEditFlags_InputRGB wraps ImGuiColorEditFlags_InputRGB.
	GuiColorEditFlags_InputRGB GuiColorEditFlags = C.ImGuiColorEditFlags_InputRGB
	// GuiColorEditFlags_NoAlpha wraps ImGuiColorEditFlags_NoAlpha.
	GuiColorEditFlags_NoAlpha GuiColorEditFlags = C.ImGuiColorEditFlags_NoAlpha
	// GuiColorEditFlags_NoBorder wraps ImGuiColorEditFlags_NoBorder.
	GuiColorEditFlags_NoBorder GuiColorEditFlags = C.ImGuiColorEditFlags_NoBorder
	// GuiColorEditFlags_NoDragDrop wraps ImGuiColorEditFlags_NoDragDrop.
	GuiColorEditFlags_NoDragDrop GuiColorEditFlags = C.ImGuiColorEditFlags_NoDragDrop
	// GuiColorEditFlags_NoInputs wraps ImGuiColorEditFlags_NoInputs.
	GuiColorEditFlags_NoInputs GuiColorEditFlags = C.ImGuiColorEditFlags_NoInputs
	// GuiColorEditFlags_NoLabel wraps ImGuiColorEditFlags_NoLabel.
	GuiColorEditFlags_NoLabel GuiColorEditFlags = C.ImGuiColorEditFlags_NoLabel
	// GuiColorEditFlags_NoOptions wraps ImGuiColorEditFlags_NoOptions.
	GuiColorEditFlags_NoOptions GuiColorEditFlags = C.ImGuiColorEditFlags_NoOptions
	// GuiColorEditFlags_NoPicker wraps ImGuiColorEditFlags_NoPicker.
	GuiColorEditFlags_NoPicker GuiColorEditFlags = C.ImGuiColorEditFlags_NoPicker
	// GuiColorEditFlags_NoSidePreview wraps ImGuiColorEditFlags_NoSidePreview.
	GuiColorEditFlags_NoSidePreview GuiColorEditFlags = C.ImGuiColorEditFlags_NoSidePreview
	// GuiColorEditFlags_NoSmallPreview wraps ImGuiColorEditFlags_NoSmallPreview.
	GuiColorEditFlags_NoSmallPreview GuiColorEditFlags = C.ImGuiColorEditFlags_NoSmallPreview
	// GuiColorEditFlags_NoTooltip wraps ImGuiColorEditFlags_NoTooltip.
	GuiColorEditFlags_NoTooltip GuiColorEditFlags = C.ImGuiColorEditFlags_NoTooltip
	// GuiColorEditFlags_None wraps ImGuiColorEditFlags_None.
	GuiColorEditFlags_None GuiColorEditFlags = C.ImGuiColorEditFlags_None
	// GuiColorEditFlags_PickerHueBar wraps ImGuiColorEditFlags_PickerHueBar.
	GuiColorEditFlags_PickerHueBar GuiColorEditFlags = C.ImGuiColorEditFlags_PickerHueBar
	// GuiColorEditFlags_PickerHueWheel wraps ImGuiColorEditFlags_PickerHueWheel.
	GuiColorEditFlags_PickerHueWheel GuiColorEditFlags = C.ImGuiColorEditFlags_PickerHueWheel
	// GuiColorEditFlags_PickerMask wraps ImGuiColorEditFlags_PickerMask_.
	GuiColorEditFlags_PickerMask GuiColorEditFlags = C.ImGuiColorEditFlags_PickerMask_
	// GuiColorEditFlags_Uint8 wraps ImGuiColorEditFlags_Uint8.
	GuiColorEditFlags_Uint8 GuiColorEditFlags = C.ImGuiColorEditFlags_Uint8
)

func (e GuiColorEditFlags) String() string {
	return fmt.Sprintf("ImGuiColorEditFlags(%b)", e)
}

// GuiComboFlags wraps the bitmask ImGuiComboFlags.
type GuiComboFlags int32

const (
	// GuiComboFlags_HeightLarge wraps ImGuiComboFlags_HeightLarge.
	GuiComboFlags_HeightLarge GuiComboFlags = C.ImGuiComboFlags_HeightLarge
	// GuiComboFlags_HeightLargest wraps ImGuiComboFlags_HeightLargest.
	GuiComboFlags_HeightLargest GuiComboFlags = C.ImGuiComboFlags_HeightLargest
	// GuiComboFlags_HeightMask wraps ImGuiComboFlags_HeightMask_.
	GuiComboFlags_HeightMask GuiComboFlags = C.ImGuiComboFlags_HeightMask_
	// GuiComboFlags_HeightRegular wraps ImGuiComboFlags_HeightRegular.
	GuiComboFlags_HeightRegular GuiComboFlags = C.ImGuiComboFlags_HeightRegular
	// GuiComboFlags_HeightSmall wraps ImGuiComboFlags_HeightSmall.
	GuiComboFlags_HeightSmall GuiComboFlags = C.ImGuiComboFlags_HeightSmall
	// GuiComboFlags_NoArrowButton wraps ImGuiComboFlags_NoArrowButton.
	GuiComboFlags_NoArrowButton GuiComboFlags = C.ImGuiComboFlags_NoArrowButton
	// GuiComboFlags_NoPreview wraps ImGuiComboFlags_NoPreview.
	GuiComboFlags_NoPreview GuiComboFlags = C.ImGuiComboFlags_NoPreview
	// GuiComboFlags_None wraps ImGuiComboFlags_None.
	GuiComboFlags_None GuiComboFlags = C.ImGuiComboFlags_None
	// GuiComboFlags_PopupAlignLeft wraps ImGuiComboFlags_PopupAlignLeft.
	GuiComboFlags_PopupAlignLeft GuiComboFlags = C.ImGuiComboFlags_PopupAlignLeft
	// GuiComboFlags_WidthFitPreview wraps ImGuiComboFlags_WidthFitPreview.
	GuiComboFlags_WidthFitPreview GuiComboFlags = C.ImGuiComboFlags_WidthFitPreview
)

func (e GuiComboFlags) String() string {
	return fmt.Sprintf("ImGuiComboFlags(%b)", e)
}

// GuiComboFlagsPrivate wraps the bitmask ImGuiComboFlagsPrivate.
type GuiComboFlagsPrivate int32

const (
	// GuiComboFlags_CustomPreview wraps ImGuiComboFlags_CustomPreview.
	GuiComboFlags_CustomPreview GuiComboFlagsPrivate = C.ImGuiComboFlags_CustomPreview
)

func (e GuiComboFlagsPrivate) String() string {
	return fmt.Sprintf("ImGuiComboFlagsPrivate(%b)", e)
}

// GuiConfigFlags wraps the bitmask ImGuiConfigFlags.
type GuiConfigFlags int32

const (
	// GuiConfigFlags_DockingEnable wraps ImGuiConfigFlags_DockingEnable.
	GuiConfigFlags_DockingEnable GuiConfigFlags = C.ImGuiConfigFlags_DockingEnable
	// GuiConfigFlags_IsSRGB wraps ImGuiConfigFlags_IsSRGB.
	GuiConfigFlags_IsSRGB GuiConfigFlags = C.ImGuiConfigFlags_IsSRGB
	// GuiConfigFlags_IsTouchScreen wraps ImGuiConfigFlags_IsTouchScreen.
	GuiConfigFlags_IsTouchScreen GuiConfigFlags = C.ImGuiConfigFlags_IsTouchScreen
	// GuiConfigFlags_NavEnableGamepad wraps ImGuiConfigFlags_NavEnableGamepad.
	GuiConfigFlags_NavEnableGamepad GuiConfigFlags = C.ImGuiConfigFlags_NavEnableGamepad
	// GuiConfigFlags_NavEnableKeyboard wraps ImGuiConfigFlags_NavEnableKeyboard.
	GuiConfigFlags_NavEnableKeyboard GuiConfigFlags = C.ImGuiConfigFlags_NavEnableKeyboard
	// GuiConfigFlags_NoKeyboard wraps ImGuiConfigFlags_NoKeyboard.
	GuiConfigFlags_NoKeyboard GuiConfigFlags = C.ImGuiConfigFlags_NoKeyboard
	// GuiConfigFlags_NoMouse wraps ImGuiConfigFlags_NoMouse.
	GuiConfigFlags_NoMouse GuiConfigFlags = C.ImGuiConfigFlags_NoMouse
	// GuiConfigFlags_NoMouseCursorChange wraps ImGuiConfigFlags_NoMouseCursorChange.
	GuiConfigFlags_NoMouseCursorChange GuiConfigFlags = C.ImGuiConfigFlags_NoMouseCursorChange
	// GuiConfigFlags_None wraps ImGuiConfigFlags_None.
	GuiConfigFlags_None GuiConfigFlags = C.ImGuiConfigFlags_None
	// GuiConfigFlags_ViewportsEnable wraps ImGuiConfigFlags_ViewportsEnable.
	GuiConfigFlags_ViewportsEnable GuiConfigFlags = C.ImGuiConfigFlags_ViewportsEnable
)

func (e GuiConfigFlags) String() string {
	return fmt.Sprintf("ImGuiConfigFlags(%b)", e)
}

// GuiDebugLogFlags wraps the bitmask ImGuiDebugLogFlags.
type GuiDebugLogFlags int32

const (
	// GuiDebugLogFlags_EventActiveId wraps ImGuiDebugLogFlags_EventActiveId.
	GuiDebugLogFlags_EventActiveId GuiDebugLogFlags = C.ImGuiDebugLogFlags_EventActiveId
	// GuiDebugLogFlags_EventClipper wraps ImGuiDebugLogFlags_EventClipper.
	GuiDebugLogFlags_EventClipper GuiDebugLogFlags = C.ImGuiDebugLogFlags_EventClipper
	// GuiDebugLogFlags_EventDocking wraps ImGuiDebugLogFlags_EventDocking.
	GuiDebugLogFlags_EventDocking GuiDebugLogFlags = C.ImGuiDebugLogFlags_EventDocking
	// GuiDebugLogFlags_EventError wraps ImGuiDebugLogFlags_EventError.
	GuiDebugLogFlags_EventError GuiDebugLogFlags = C.ImGuiDebugLogFlags_EventError
	// GuiDebugLogFlags_EventFocus wraps ImGuiDebugLogFlags_EventFocus.
	GuiDebugLogFlags_EventFocus GuiDebugLogFlags = C.ImGuiDebugLogFlags_EventFocus
	// GuiDebugLogFlags_EventFont wraps ImGuiDebugLogFlags_EventFont.
	GuiDebugLogFlags_EventFont GuiDebugLogFlags = C.ImGuiDebugLogFlags_EventFont
	// GuiDebugLogFlags_EventIO wraps ImGuiDebugLogFlags_EventIO.
	GuiDebugLogFlags_EventIO GuiDebugLogFlags = C.ImGuiDebugLogFlags_EventIO
	// GuiDebugLogFlags_EventInputRouting wraps ImGuiDebugLogFlags_EventInputRouting.
	GuiDebugLogFlags_EventInputRouting GuiDebugLogFlags = C.ImGuiDebugLogFlags_EventInputRouting
	// GuiDebugLogFlags_EventMask wraps ImGuiDebugLogFlags_EventMask_.
	GuiDebugLogFlags_EventMask GuiDebugLogFlags = C.ImGuiDebugLogFlags_EventMask_
	// GuiDebugLogFlags_EventNav wraps ImGuiDebugLogFlags_EventNav.
	GuiDebugLogFlags_EventNav GuiDebugLogFlags = C.ImGuiDebugLogFlags_EventNav
	// GuiDebugLogFlags_EventPopup wraps ImGuiDebugLogFlags_EventPopup.
	GuiDebugLogFlags_EventPopup GuiDebugLogFlags = C.ImGuiDebugLogFlags_EventPopup
	// GuiDebugLogFlags_EventSelection wraps ImGuiDebugLogFlags_EventSelection.
	GuiDebugLogFlags_EventSelection GuiDebugLogFlags = C.ImGuiDebugLogFlags_EventSelection
	// GuiDebugLogFlags_EventViewport wraps ImGuiDebugLogFlags_EventViewport.
	GuiDebugLogFlags_EventViewport GuiDebugLogFlags = C.ImGuiDebugLogFlags_EventViewport
	// GuiDebugLogFlags_None wraps ImGuiDebugLogFlags_None.
	GuiDebugLogFlags_None GuiDebugLogFlags = C.ImGuiDebugLogFlags_None
	// GuiDebugLogFlags_OutputToTTY wraps ImGuiDebugLogFlags_OutputToTTY.
	GuiDebugLogFlags_OutputToTTY GuiDebugLogFlags = C.ImGuiDebugLogFlags_OutputToTTY
	// GuiDebugLogFlags_OutputToTestEngine wraps ImGuiDebugLogFlags_OutputToTestEngine.
	GuiDebugLogFlags_OutputToTestEngine GuiDebugLogFlags = C.ImGuiDebugLogFlags_OutputToTestEngine
)

func (e GuiDebugLogFlags) String() string {
	return fmt.Sprintf("ImGuiDebugLogFlags(%b)", e)
}

// GuiDockNodeFlags wraps the bitmask ImGuiDockNodeFlags.
type GuiDockNodeFlags int32

const (
	// GuiDockNodeFlags_AutoHideTabBar wraps ImGuiDockNodeFlags_AutoHideTabBar.
	GuiDockNodeFlags_AutoHideTabBar GuiDockNodeFlags = C.ImGuiDockNodeFlags_AutoHideTabBar
	// GuiDockNodeFlags_KeepAliveOnly wraps ImGuiDockNodeFlags_KeepAliveOnly.
	GuiDockNodeFlags_KeepAliveOnly GuiDockNodeFlags = C.ImGuiDockNodeFlags_KeepAliveOnly
	// GuiDockNodeFlags_NoDockingOverCentralNode wraps ImGuiDockNodeFlags_NoDockingOverCentralNode.
	GuiDockNodeFlags_NoDockingOverCentralNode GuiDockNodeFlags = C.ImGuiDockNodeFlags_NoDockingOverCentralNode
	// GuiDockNodeFlags_NoDockingSplit wraps ImGuiDockNodeFlags_NoDockingSplit.
	GuiDockNodeFlags_NoDockingSplit GuiDockNodeFlags = C.ImGuiDockNodeFlags_NoDockingSplit
	// GuiDockNodeFlags_NoResize wraps ImGuiDockNodeFlags_NoResize.
	GuiDockNodeFlags_NoResize GuiDockNodeFlags = C.ImGuiDockNodeFlags_NoResize
	// GuiDockNodeFlags_NoUndocking wraps ImGuiDockNodeFlags_NoUndocking.
	GuiDockNodeFlags_NoUndocking GuiDockNodeFlags = C.ImGuiDockNodeFlags_NoUndocking
	// GuiDockNodeFlags_None wraps ImGuiDockNodeFlags_None.
	GuiDockNodeFlags_None GuiDockNodeFlags = C.ImGuiDockNodeFlags_None
	// GuiDockNodeFlags_PassthruCentralNode wraps ImGuiDockNodeFlags_PassthruCentralNode.
	GuiDockNodeFlags_PassthruCentralNode GuiDockNodeFlags = C.ImGuiDockNodeFlags_PassthruCentralNode
)

func (e GuiDockNodeFlags) String() string {
	return fmt.Sprintf("ImGuiDockNodeFlags(%b)", e)
}

// GuiDockNodeFlagsPrivate wraps the bitmask ImGuiDockNodeFlagsPrivate.
type GuiDockNodeFlagsPrivate int32

const (
	// GuiDockNodeFlags_CentralNode wraps ImGuiDockNodeFlags_CentralNode.
	GuiDockNodeFlags_CentralNode GuiDockNodeFlagsPrivate = C.ImGuiDockNodeFlags_CentralNode
	// GuiDockNodeFlags_DockSpace wraps ImGuiDockNodeFlags_DockSpace.
	GuiDockNodeFlags_DockSpace GuiDockNodeFlagsPrivate = C.ImGuiDockNodeFlags_DockSpace
	// GuiDockNodeFlags_DockedWindowsInFocusRoute wraps ImGuiDockNodeFlags_DockedWindowsInFocusRoute.
	GuiDockNodeFlags_DockedWindowsInFocusRoute GuiDockNodeFlagsPrivate = C.ImGuiDockNodeFlags_DockedWindowsInFocusRoute
	// GuiDockNodeFlags_HiddenTabBar wraps ImGuiDockNodeFlags_HiddenTabBar.
	GuiDockNodeFlags_HiddenTabBar GuiDockNodeFlagsPrivate = C.ImGuiDockNodeFlags_HiddenTabBar
	// GuiDockNodeFlags_LocalFlagsTransferMask wraps ImGuiDockNodeFlags_LocalFlagsTransferMask_.
	GuiDockNodeFlags_LocalFlagsTransferMask GuiDockNodeFlagsPrivate = C.ImGuiDockNodeFlags_LocalFlagsTransferMask_
	// GuiDockNodeFlags_NoCloseButton wraps ImGuiDockNodeFlags_NoCloseButton.
	GuiDockNodeFlags_NoCloseButton GuiDockNodeFlagsPrivate = C.ImGuiDockNodeFlags_NoCloseButton
	// GuiDockNodeFlags_NoDocking wraps ImGuiDockNodeFlags_NoDocking.
	GuiDockNodeFlags_NoDocking GuiDockNodeFlagsPrivate = C.ImGuiDockNodeFlags_NoDocking
	// GuiDockNodeFlags_NoDockingOverEmpty wraps ImGuiDockNodeFlags_NoDockingOverEmpty.
	GuiDockNodeFlags_NoDockingOverEmpty GuiDockNodeFlagsPrivate = C.ImGuiDockNodeFlags_NoDockingOverEmpty
	// GuiDockNodeFlags_NoDockingOverMe wraps ImGuiDockNodeFlags_NoDockingOverMe.
	GuiDockNodeFlags_NoDockingOverMe GuiDockNodeFlagsPrivate = C.ImGuiDockNodeFlags_NoDockingOverMe
	// GuiDockNodeFlags_NoDockingOverOther wraps ImGuiDockNodeFlags_NoDockingOverOther.
	GuiDockNodeFlags_NoDockingOverOther GuiDockNodeFlagsPrivate = C.ImGuiDockNodeFlags_NoDockingOverOther
	// GuiDockNodeFlags_NoDockingSplitOther wraps ImGuiDockNodeFlags_NoDockingSplitOther.
	GuiDockNodeFlags_NoDockingSplitOther GuiDockNodeFlagsPrivate = C.ImGuiDockNodeFlags_NoDockingSplitOther
	// GuiDockNodeFlags_NoResizeFlagsMask wraps ImGuiDockNodeFlags_NoResizeFlagsMask_.
	GuiDockNodeFlags_NoResizeFlagsMask GuiDockNodeFlagsPrivate = C.ImGuiDockNodeFlags_NoResizeFlagsMask_
	// GuiDockNodeFlags_NoResizeX wraps ImGuiDockNodeFlags_NoResizeX.
	GuiDockNodeFlags_NoResizeX GuiDockNodeFlagsPrivate = C.ImGuiDockNodeFlags_NoResizeX
	// GuiDockNodeFlags_NoResizeY wraps ImGuiDockNodeFlags_NoResizeY.
	GuiDockNodeFlags_NoResizeY GuiDockNodeFlagsPrivate = C.ImGuiDockNodeFlags_NoResizeY
	// GuiDockNodeFlags_NoTabBar wraps ImGuiDockNodeFlags_NoTabBar.
	GuiDockNodeFlags_NoTabBar GuiDockNodeFlagsPrivate = C.ImGuiDockNodeFlags_NoTabBar
	// GuiDockNodeFlags_NoWindowMenuButton wraps ImGuiDockNodeFlags_NoWindowMenuButton.
	GuiDockNodeFlags_NoWindowMenuButton GuiDockNodeFlagsPrivate = C.ImGuiDockNodeFlags_NoWindowMenuButton
	// GuiDockNodeFlags_SavedFlagsMask wraps ImGuiDockNodeFlags_SavedFlagsMask_.
	GuiDockNodeFlags_SavedFlagsMask GuiDockNodeFlagsPrivate = C.ImGuiDockNodeFlags_SavedFlagsMask_
	// GuiDockNodeFlags_SharedFlagsInheritMask wraps ImGuiDockNodeFlags_SharedFlagsInheritMask_.
	GuiDockNodeFlags_SharedFlagsInheritMask GuiDockNodeFlagsPrivate = C.ImGuiDockNodeFlags_SharedFlagsInheritMask_
)

func (e GuiDockNodeFlagsPrivate) String() string {
	return fmt.Sprintf("ImGuiDockNodeFlagsPrivate(%b)", e)
}

// GuiDragDropFlags wraps the bitmask ImGuiDragDropFlags.
type GuiDragDropFlags int32

const (
	// GuiDragDropFlags_AcceptBeforeDelivery wraps ImGuiDragDropFlags_AcceptBeforeDelivery.
	GuiDragDropFlags_AcceptBeforeDelivery GuiDragDropFlags = C.ImGuiDragDropFlags_AcceptBeforeDelivery
	// GuiDragDropFlags_AcceptNoDrawDefaultRect wraps ImGuiDragDropFlags_AcceptNoDrawDefaultRect.
	GuiDragDropFlags_AcceptNoDrawDefaultRect GuiDragDropFlags = C.ImGuiDragDropFlags_AcceptNoDrawDefaultRect
	// GuiDragDropFlags_AcceptNoPreviewTooltip wraps ImGuiDragDropFlags_AcceptNoPreviewTooltip.
	GuiDragDropFlags_AcceptNoPreviewTooltip GuiDragDropFlags = C.ImGuiDragDropFlags_AcceptNoPreviewTooltip
	// GuiDragDropFlags_AcceptPeekOnly wraps ImGuiDragDropFlags_AcceptPeekOnly.
	GuiDragDropFlags_AcceptPeekOnly GuiDragDropFlags = C.ImGuiDragDropFlags_AcceptPeekOnly
	// GuiDragDropFlags_None wraps ImGuiDragDropFlags_None.
	GuiDragDropFlags_None GuiDragDropFlags = C.ImGuiDragDropFlags_None
	// GuiDragDropFlags_PayloadAutoExpire wraps ImGuiDragDropFlags_PayloadAutoExpire.
	GuiDragDropFlags_PayloadAutoExpire GuiDragDropFlags = C.ImGuiDragDropFlags_PayloadAutoExpire
	// GuiDragDropFlags_PayloadNoCrossContext wraps ImGuiDragDropFlags_PayloadNoCrossContext.
	GuiDragDropFlags_PayloadNoCrossContext GuiDragDropFlags = C.ImGuiDragDropFlags_PayloadNoCrossContext
	// GuiDragDropFlags_PayloadNoCrossProcess wraps ImGuiDragDropFlags_PayloadNoCrossProcess.
	GuiDragDropFlags_PayloadNoCrossProcess GuiDragDropFlags = C.ImGuiDragDropFlags_PayloadNoCrossProcess
	// GuiDragDropFlags_SourceAllowNullID wraps ImGuiDragDropFlags_SourceAllowNullID.
	GuiDragDropFlags_SourceAllowNullID GuiDragDropFlags = C.ImGuiDragDropFlags_SourceAllowNullID
	// GuiDragDropFlags_SourceExtern wraps ImGuiDragDropFlags_SourceExtern.
	GuiDragDropFlags_SourceExtern GuiDragDropFlags = C.ImGuiDragDropFlags_SourceExtern
	// GuiDragDropFlags_SourceNoDisableHover wraps ImGuiDragDropFlags_SourceNoDisableHover.
	GuiDragDropFlags_SourceNoDisableHover GuiDragDropFlags = C.ImGuiDragDropFlags_SourceNoDisableHover
	// GuiDragDropFlags_SourceNoHoldToOpenOthers wraps ImGuiDragDropFlags_SourceNoHoldToOpenOthers.
	GuiDragDropFlags_SourceNoHoldToOpenOthers GuiDragDropFlags = C.ImGuiDragDropFlags_SourceNoHoldToOpenOthers
	// GuiDragDropFlags_SourceNoPreviewTooltip wraps ImGuiDragDropFlags_SourceNoPreviewTooltip.
	GuiDragDropFlags_SourceNoPreviewTooltip GuiDragDropFlags = C.ImGuiDragDropFlags_SourceNoPreviewTooltip
)

func (e GuiDragDropFlags) String() string {
	return fmt.Sprintf("ImGuiDragDropFlags(%b)", e)
}

// GuiFocusRequestFlags wraps the bitmask ImGuiFocusRequestFlags.
type GuiFocusRequestFlags int32

const (
	// GuiFocusRequestFlags_None wraps ImGuiFocusRequestFlags_None.
	GuiFocusRequestFlags_None GuiFocusRequestFlags = C.ImGuiFocusRequestFlags_None
	// GuiFocusRequestFlags_RestoreFocusedChild wraps ImGuiFocusRequestFlags_RestoreFocusedChild.
	GuiFocusRequestFlags_RestoreFocusedChild GuiFocusRequestFlags = C.ImGuiFocusRequestFlags_RestoreFocusedChild
	// GuiFocusRequestFlags_UnlessBelowModal wraps ImGuiFocusRequestFlags_UnlessBelowModal.
	GuiFocusRequestFlags_UnlessBelowModal GuiFocusRequestFlags = C.ImGuiFocusRequestFlags_UnlessBelowModal
)

func (e GuiFocusRequestFlags) String() string {
	return fmt.Sprintf("ImGuiFocusRequestFlags(%b)", e)
}

// GuiFocusedFlags wraps the bitmask ImGuiFocusedFlags.
type GuiFocusedFlags int32

const (
	// GuiFocusedFlags_AnyWindow wraps ImGuiFocusedFlags_AnyWindow.
	GuiFocusedFlags_AnyWindow GuiFocusedFlags = C.ImGuiFocusedFlags_AnyWindow
	// GuiFocusedFlags_ChildWindows wraps ImGuiFocusedFlags_ChildWindows.
	GuiFocusedFlags_ChildWindows GuiFocusedFlags = C.ImGuiFocusedFlags_ChildWindows
	// GuiFocusedFlags_DockHierarchy wraps ImGuiFocusedFlags_DockHierarchy.
	GuiFocusedFlags_DockHierarchy GuiFocusedFlags = C.ImGuiFocusedFlags_DockHierarchy
	// GuiFocusedFlags_NoPopupHierarchy wraps ImGuiFocusedFlags_NoPopupHierarchy.
	GuiFocusedFlags_NoPopupHierarchy GuiFocusedFlags = C.ImGuiFocusedFlags_NoPopupHierarchy
	// GuiFocusedFlags_None wraps ImGuiFocusedFlags_None.
	GuiFocusedFlags_None GuiFocusedFlags = C.ImGuiFocusedFlags_None
	// GuiFocusedFlags_RootAndChildWindows wraps ImGuiFocusedFlags_RootAndChildWindows.
	GuiFocusedFlags_RootAndChildWindows GuiFocusedFlags = C.ImGuiFocusedFlags_RootAndChildWindows
	// GuiFocusedFlags_RootWindow wraps ImGuiFocusedFlags_RootWindow.
	GuiFocusedFlags_RootWindow GuiFocusedFlags = C.ImGuiFocusedFlags_RootWindow
)

func (e GuiFocusedFlags) String() string {
	return fmt.Sprintf("ImGuiFocusedFlags(%b)", e)
}

// GuiHoveredFlags wraps the bitmask ImGuiHoveredFlags.
type GuiHoveredFlags int32

const (
	// GuiHoveredFlags_AllowWhenBlockedByActiveItem wraps ImGuiHoveredFlags_AllowWhenBlockedByActiveItem.
	GuiHoveredFlags_AllowWhenBlockedByActiveItem GuiHoveredFlags = C.ImGuiHoveredFlags_AllowWhenBlockedByActiveItem
	// GuiHoveredFlags_AllowWhenBlockedByPopup wraps ImGuiHoveredFlags_AllowWhenBlockedByPopup.
	GuiHoveredFlags_AllowWhenBlockedByPopup GuiHoveredFlags = C.ImGuiHoveredFlags_AllowWhenBlockedByPopup
	// GuiHoveredFlags_AllowWhenDisabled wraps ImGuiHoveredFlags_AllowWhenDisabled.
	GuiHoveredFlags_AllowWhenDisabled GuiHoveredFlags = C.ImGuiHoveredFlags_AllowWhenDisabled
	// GuiHoveredFlags_AllowWhenOverlapped wraps ImGuiHoveredFlags_AllowWhenOverlapped.
	GuiHoveredFlags_AllowWhenOverlapped GuiHoveredFlags = C.ImGuiHoveredFlags_AllowWhenOverlapped
	// GuiHoveredFlags_AllowWhenOverlappedByItem wraps ImGuiHoveredFlags_AllowWhenOverlappedByItem.
	GuiHoveredFlags_AllowWhenOverlappedByItem GuiHoveredFlags = C.ImGuiHoveredFlags_AllowWhenOverlappedByItem
	// GuiHoveredFlags_AllowWhenOverlappedByWindow wraps ImGuiHoveredFlags_AllowWhenOverlappedByWindow.
	GuiHoveredFlags_AllowWhenOverlappedByWindow GuiHoveredFlags = C.ImGuiHoveredFlags_AllowWhenOverlappedByWindow
	// GuiHoveredFlags_AnyWindow wraps ImGuiHoveredFlags_AnyWindow.
	GuiHoveredFlags_AnyWindow GuiHoveredFlags = C.ImGuiHoveredFlags_AnyWindow
	// GuiHoveredFlags_ChildWindows wraps ImGuiHoveredFlags_ChildWindows.
	GuiHoveredFlags_ChildWindows GuiHoveredFlags = C.ImGuiHoveredFlags_ChildWindows
	// GuiHoveredFlags_DelayNone wraps ImGuiHoveredFlags_DelayNone.
	GuiHoveredFlags_DelayNone GuiHoveredFlags = C.ImGuiHoveredFlags_DelayNone
	// GuiHoveredFlags_DelayNormal wraps ImGuiHoveredFlags_DelayNormal.
	GuiHoveredFlags_DelayNormal GuiHoveredFlags = C.ImGuiHoveredFlags_DelayNormal
	// GuiHoveredFlags_DelayShort wraps ImGuiHoveredFlags_DelayShort.
	GuiHoveredFlags_DelayShort GuiHoveredFlags = C.ImGuiHoveredFlags_DelayShort
	// GuiHoveredFlags_DockHierarchy wraps ImGuiHoveredFlags_DockHierarchy.
	GuiHoveredFlags_DockHierarchy GuiHoveredFlags = C.ImGuiHoveredFlags_DockHierarchy
	// GuiHoveredFlags_ForTooltip wraps ImGuiHoveredFlags_ForTooltip.
	GuiHoveredFlags_ForTooltip GuiHoveredFlags = C.ImGuiHoveredFlags_ForTooltip
	// GuiHoveredFlags_NoNavOverride wraps ImGuiHoveredFlags_NoNavOverride.
	GuiHoveredFlags_NoNavOverride GuiHoveredFlags = C.ImGuiHoveredFlags_NoNavOverride
	// GuiHoveredFlags_NoPopupHierarchy wraps ImGuiHoveredFlags_NoPopupHierarchy.
	GuiHoveredFlags_NoPopupHierarchy GuiHoveredFlags = C.ImGuiHoveredFlags_NoPopupHierarchy
	// GuiHoveredFlags_NoSharedDelay wraps ImGuiHoveredFlags_NoSharedDelay.
	GuiHoveredFlags_NoSharedDelay GuiHoveredFlags = C.ImGuiHoveredFlags_NoSharedDelay
	// GuiHoveredFlags_None wraps ImGuiHoveredFlags_None.
	GuiHoveredFlags_None GuiHoveredFlags = C.ImGuiHoveredFlags_None
	// GuiHoveredFlags_RectOnly wraps ImGuiHoveredFlags_RectOnly.
	GuiHoveredFlags_RectOnly GuiHoveredFlags = C.ImGuiHoveredFlags_RectOnly
	// GuiHoveredFlags_RootAndChildWindows wraps ImGuiHoveredFlags_RootAndChildWindows.
	GuiHoveredFlags_RootAndChildWindows GuiHoveredFlags = C.ImGuiHoveredFlags_RootAndChildWindows
	// GuiHoveredFlags_RootWindow wraps ImGuiHoveredFlags_RootWindow.
	GuiHoveredFlags_RootWindow GuiHoveredFlags = C.ImGuiHoveredFlags_RootWindow
	// GuiHoveredFlags_Stationary wraps ImGuiHoveredFlags_Stationary.
	GuiHoveredFlags_Stationary GuiHoveredFlags = C.ImGuiHoveredFlags_Stationary
)

func (e GuiHoveredFlags) String() string {
	return fmt.Sprintf("ImGuiHoveredFlags(%b)", e)
}

// GuiHoveredFlagsPrivate wraps the bitmask ImGuiHoveredFlagsPrivate.
type GuiHoveredFlagsPrivate int32

const (
	// GuiHoveredFlags_AllowedMaskForIsItemHovered wraps ImGuiHoveredFlags_AllowedMaskForIsItemHovered.
	GuiHoveredFlags_AllowedMaskForIsItemHovered GuiHoveredFlagsPrivate = C.ImGuiHoveredFlags_AllowedMaskForIsItemHovered
	// GuiHoveredFlags_AllowedMaskForIsWindowHovered wraps ImGuiHoveredFlags_AllowedMaskForIsWindowHovered.
	GuiHoveredFlags_AllowedMaskForIsWindowHovered GuiHoveredFlagsPrivate = C.ImGuiHoveredFlags_AllowedMaskForIsWindowHovered
	// GuiHoveredFlags_DelayMask wraps ImGuiHoveredFlags_DelayMask_.
	GuiHoveredFlags_DelayMask GuiHoveredFlagsPrivate = C.ImGuiHoveredFlags_DelayMask_
)

func (e GuiHoveredFlagsPrivate) String() string {
	return fmt.Sprintf("ImGuiHoveredFlagsPrivate(%b)", e)
}

// GuiInputFlags wraps the bitmask ImGuiInputFlags.
type GuiInputFlags int32

const (
	// GuiInputFlags_None wraps ImGuiInputFlags_None.
	GuiInputFlags_None GuiInputFlags = C.ImGuiInputFlags_None
	// GuiInputFlags_Repeat wraps ImGuiInputFlags_Repeat.
	GuiInputFlags_Repeat GuiInputFlags = C.ImGuiInputFlags_Repeat
	// GuiInputFlags_RouteActive wraps ImGuiInputFlags_RouteActive.
	GuiInputFlags_RouteActive GuiInputFlags = C.ImGuiInputFlags_RouteActive
	// GuiInputFlags_RouteAlways wraps ImGuiInputFlags_RouteAlways.
	GuiInputFlags_RouteAlways GuiInputFlags = C.ImGuiInputFlags_RouteAlways
	// GuiInputFlags_RouteFocused wraps ImGuiInputFlags_RouteFocused.
	GuiInputFlags_RouteFocused GuiInputFlags = C.ImGuiInputFlags_RouteFocused
	// GuiInputFlags_RouteFromRootWindow wraps ImGuiInputFlags_RouteFromRootWindow.
	GuiInputFlags_RouteFromRootWindow GuiInputFlags = C.ImGuiInputFlags_RouteFromRootWindow
	// GuiInputFlags_RouteGlobal wraps ImGuiInputFlags_RouteGlobal.
	GuiInputFlags_RouteGlobal GuiInputFlags = C.ImGuiInputFlags_RouteGlobal
	// GuiInputFlags_RouteOverActive wraps ImGuiInputFlags_RouteOverActive.
	GuiInputFlags_RouteOverActive GuiInputFlags = C.ImGuiInputFlags_RouteOverActive
	// GuiInputFlags_RouteOverFocused wraps ImGuiInputFlags_RouteOverFocused.
	GuiInputFlags_RouteOverFocused GuiInputFlags = C.ImGuiInputFlags_RouteOverFocused
	// GuiInputFlags_RouteUnlessBgFocused wraps ImGuiInputFlags_RouteUnlessBgFocused.
	GuiInputFlags_RouteUnlessBgFocused GuiInputFlags = C.ImGuiInputFlags_RouteUnlessBgFocused
	// GuiInputFlags_Tooltip wraps ImGuiInputFlags_Tooltip.
	GuiInputFlags_Tooltip GuiInputFlags = C.ImGuiInputFlags_Tooltip
)

func (e GuiInputFlags) String() string {
	return fmt.Sprintf("ImGuiInputFlags(%b)", e)
}

// GuiInputFlagsPrivate wraps the bitmask ImGuiInputFlagsPrivate.
type GuiInputFlagsPrivate int32

const (
	// GuiInputFlags_CondActive wraps ImGuiInputFlags_CondActive.
	GuiInputFlags_CondActive GuiInputFlagsPrivate = C.ImGuiInputFlags_CondActive
	// GuiInputFlags_CondDefault wraps ImGuiInputFlags_CondDefault_.
	GuiInputFlags_CondDefault GuiInputFlagsPrivate = C.ImGuiInputFlags_CondDefault_
	// GuiInputFlags_CondHovered wraps ImGuiInputFlags_CondHovered.
	GuiInputFlags_CondHovered GuiInputFlagsPrivate = C.ImGuiInputFlags_CondHovered
	// GuiInputFlags_CondMask wraps ImGuiInputFlags_CondMask_.
	GuiInputFlags_CondMask GuiInputFlagsPrivate = C.ImGuiInputFlags_CondMask_
	// GuiInputFlags_LockThisFrame wraps ImGuiInputFlags_LockThisFrame.
	GuiInputFlags_LockThisFrame GuiInputFlagsPrivate = C.ImGuiInputFlags_LockThisFrame
	// GuiInputFlags_LockUntilRelease wraps ImGuiInputFlags_LockUntilRelease.
	GuiInputFlags_LockUntilRelease GuiInputFlagsPrivate = C.ImGuiInputFlags_LockUntilRelease
	// GuiInputFlags_RepeatMask wraps ImGuiInputFlags_RepeatMask_.
	GuiInputFlags_RepeatMask GuiInputFlagsPrivate = C.ImGuiInputFlags_RepeatMask_
	// GuiInputFlags_RepeatRateDefault wraps ImGuiInputFlags_RepeatRateDefault.
	GuiInputFlags_RepeatRateDefault GuiInputFlagsPrivate = C.ImGuiInputFlags_RepeatRateDefault
	// GuiInputFlags_RepeatRateMask wraps ImGuiInputFlags_RepeatRateMask_.
	GuiInputFlags_RepeatRateMask GuiInputFlagsPrivate = C.ImGuiInputFlags_RepeatRateMask_
	// GuiInputFlags_RepeatRateNavMove wraps ImGuiInputFlags_RepeatRateNavMove.
	GuiInputFlags_RepeatRateNavMove GuiInputFlagsPrivate = C.ImGuiInputFlags_RepeatRateNavMove
	// GuiInputFlags_RepeatRateNavTweak wraps ImGuiInputFlags_RepeatRateNavTweak.
	GuiInputFlags_RepeatRateNavTweak GuiInputFlagsPrivate = C.ImGuiInputFlags_RepeatRateNavTweak
	// GuiInputFlags_RepeatUntilKeyModsChange wraps ImGuiInputFlags_RepeatUntilKeyModsChange.
	GuiInputFlags_RepeatUntilKeyModsChange GuiInputFlagsPrivate = C.ImGuiInputFlags_RepeatUntilKeyModsChange
	// GuiInputFlags_RepeatUntilKeyModsChangeFromNone wraps ImGuiInputFlags_RepeatUntilKeyModsChangeFromNone.
	GuiInputFlags_RepeatUntilKeyModsChangeFromNone GuiInputFlagsPrivate = C.ImGuiInputFlags_RepeatUntilKeyModsChangeFromNone
	// GuiInputFlags_RepeatUntilMask wraps ImGuiInputFlags_RepeatUntilMask_.
	GuiInputFlags_RepeatUntilMask GuiInputFlagsPrivate = C.ImGuiInputFlags_RepeatUntilMask_
	// GuiInputFlags_RepeatUntilOtherKeyPress wraps ImGuiInputFlags_RepeatUntilOtherKeyPress.
	GuiInputFlags_RepeatUntilOtherKeyPress GuiInputFlagsPrivate = C.ImGuiInputFlags_RepeatUntilOtherKeyPress
	// GuiInputFlags_RepeatUntilRelease wraps ImGuiInputFlags_RepeatUntilRelease.
	GuiInputFlags_RepeatUntilRelease GuiInputFlagsPrivate = C.ImGuiInputFlags_RepeatUntilRelease
	// GuiInputFlags_RouteOptionsMask wraps ImGuiInputFlags_RouteOptionsMask_.
	GuiInputFlags_RouteOptionsMask GuiInputFlagsPrivate = C.ImGuiInputFlags_RouteOptionsMask_
	// GuiInputFlags_RouteTypeMask wraps ImGuiInputFlags_RouteTypeMask_.
	GuiInputFlags_RouteTypeMask GuiInputFlagsPrivate = C.ImGuiInputFlags_RouteTypeMask_
	// GuiInputFlags_SupportedByIsKeyPressed wraps ImGuiInputFlags_SupportedByIsKeyPressed.
	GuiInputFlags_SupportedByIsKeyPressed GuiInputFlagsPrivate = C.ImGuiInputFlags_SupportedByIsKeyPressed
	// GuiInputFlags_SupportedByIsMouseClicked wraps ImGuiInputFlags_SupportedByIsMouseClicked.
	GuiInputFlags_SupportedByIsMouseClicked GuiInputFlagsPrivate = C.ImGuiInputFlags_SupportedByIsMouseClicked
	// GuiInputFlags_SupportedBySetItemKeyOwner wraps ImGuiInputFlags_SupportedBySetItemKeyOwner.
	GuiInputFlags_SupportedBySetItemKeyOwner GuiInputFlagsPrivate = C.ImGuiInputFlags_SupportedBySetItemKeyOwner
	// GuiInputFlags_SupportedBySetKeyOwner wraps ImGuiInputFlags_SupportedBySetKeyOwner.
	GuiInputFlags_SupportedBySetKeyOwner GuiInputFlagsPrivate = C.ImGuiInputFlags_SupportedBySetKeyOwner
	// GuiInputFlags_SupportedBySetNextItemShortcut wraps ImGuiInputFlags_SupportedBySetNextItemShortcut.
	GuiInputFlags_SupportedBySetNextItemShortcut GuiInputFlagsPrivate = C.ImGuiInputFlags_SupportedBySetNextItemShortcut
	// GuiInputFlags_SupportedByShortcut wraps ImGuiInputFlags_SupportedByShortcut.
	GuiInputFlags_SupportedByShortcut GuiInputFlagsPrivate = C.ImGuiInputFlags_SupportedByShortcut
)

func (e GuiInputFlagsPrivate) String() string {
	return fmt.Sprintf("ImGuiInputFlagsPrivate(%b)", e)
}

// GuiInputTextFlags wraps the bitmask ImGuiInputTextFlags.
type GuiInputTextFlags int32

const (
	// GuiInputTextFlags_AllowTabInput wraps ImGuiInputTextFlags_AllowTabInput.
	GuiInputTextFlags_AllowTabInput GuiInputTextFlags = C.ImGuiInputTextFlags_AllowTabInput
	// GuiInputTextFlags_AlwaysOverwrite wraps ImGuiInputTextFlags_AlwaysOverwrite.
	GuiInputTextFlags_AlwaysOverwrite GuiInputTextFlags = C.ImGuiInputTextFlags_AlwaysOverwrite
	// GuiInputTextFlags_AutoSelectAll wraps ImGuiInputTextFlags_AutoSelectAll.
	GuiInputTextFlags_AutoSelectAll GuiInputTextFlags = C.ImGuiInputTextFlags_AutoSelectAll
	// GuiInputTextFlags_CallbackAlways wraps ImGuiInputTextFlags_CallbackAlways.
	GuiInputTextFlags_CallbackAlways GuiInputTextFlags = C.ImGuiInputTextFlags_CallbackAlways
	// GuiInputTextFlags_CallbackCharFilter wraps ImGuiInputTextFlags_CallbackCharFilter.
	GuiInputTextFlags_CallbackCharFilter GuiInputTextFlags = C.ImGuiInputTextFlags_CallbackCharFilter
	// GuiInputTextFlags_CallbackCompletion wraps ImGuiInputTextFlags_CallbackCompletion.
	GuiInputTextFlags_CallbackCompletion GuiInputTextFlags = C.ImGuiInputTextFlags_CallbackCompletion
	// GuiInputTextFlags_CallbackEdit wraps ImGuiInputTextFlags_CallbackEdit.
	GuiInputTextFlags_CallbackEdit GuiInputTextFlags = C.ImGuiInputTextFlags_CallbackEdit
	// GuiInputTextFlags_CallbackHistory wraps ImGuiInputTextFlags_CallbackHistory.
	GuiInputTextFlags_CallbackHistory GuiInputTextFlags = C.ImGuiInputTextFlags_CallbackHistory
	// GuiInputTextFlags_CallbackResize wraps ImGuiInputTextFlags_CallbackResize.
	GuiInputTextFlags_CallbackResize GuiInputTextFlags = C.ImGuiInputTextFlags_CallbackResize
	// GuiInputTextFlags_CharsDecimal wraps ImGuiInputTextFlags_CharsDecimal.
	GuiInputTextFlags_CharsDecimal GuiInputTextFlags = C.ImGuiInputTextFlags_CharsDecimal
	// GuiInputTextFlags_CharsHexadecimal wraps ImGuiInputTextFlags_CharsHexadecimal.
	GuiInputTextFlags_CharsHexadecimal GuiInputTextFlags = C.ImGuiInputTextFlags_CharsHexadecimal
	// GuiInputTextFlags_CharsNoBlank wraps ImGuiInputTextFlags_CharsNoBlank.
	GuiInputTextFlags_CharsNoBlank GuiInputTextFlags = C.ImGuiInputTextFlags_CharsNoBlank
	// GuiInputTextFlags_CharsScientific wraps ImGuiInputTextFlags_CharsScientific.
	GuiInputTextFlags_CharsScientific GuiInputTextFlags = C.ImGuiInputTextFlags_CharsScientific
	// GuiInputTextFlags_CharsUppercase wraps ImGuiInputTextFlags_CharsUppercase.
	GuiInputTextFlags_CharsUppercase GuiInputTextFlags = C.ImGuiInputTextFlags_CharsUppercase
	// GuiInputTextFlags_CtrlEnterForNewLine wraps ImGuiInputTextFlags_CtrlEnterForNewLine.
	GuiInputTextFlags_CtrlEnterForNewLine GuiInputTextFlags = C.ImGuiInputTextFlags_CtrlEnterForNewLine
	// GuiInputTextFlags_DisplayEmptyRefVal wraps ImGuiInputTextFlags_DisplayEmptyRefVal.
	GuiInputTextFlags_DisplayEmptyRefVal GuiInputTextFlags = C.ImGuiInputTextFlags_DisplayEmptyRefVal
	// GuiInputTextFlags_ElideLeft wraps ImGuiInputTextFlags_ElideLeft.
	GuiInputTextFlags_ElideLeft GuiInputTextFlags = C.ImGuiInputTextFlags_ElideLeft
	// GuiInputTextFlags_EnterReturnsTrue wraps ImGuiInputTextFlags_EnterReturnsTrue.
	GuiInputTextFlags_EnterReturnsTrue GuiInputTextFlags = C.ImGuiInputTextFlags_EnterReturnsTrue
	// GuiInputTextFlags_EscapeClearsAll wraps ImGuiInputTextFlags_EscapeClearsAll.
	GuiInputTextFlags_EscapeClearsAll GuiInputTextFlags = C.ImGuiInputTextFlags_EscapeClearsAll
	// GuiInputTextFlags_NoHorizontalScroll wraps ImGuiInputTextFlags_NoHorizontalScroll.
	GuiInputTextFlags_NoHorizontalScroll GuiInputTextFlags = C.ImGuiInputTextFlags_NoHorizontalScroll
	// GuiInputTextFlags_NoUndoRedo wraps ImGuiInputTextFlags_NoUndoRedo.
	GuiInputTextFlags_NoUndoRedo GuiInputTextFlags = C.ImGuiInputTextFlags_NoUndoRedo
	// GuiInputTextFlags_None wraps ImGuiInputTextFlags_None.
	GuiInputTextFlags_None GuiInputTextFlags = C.ImGuiInputTextFlags_None
	// GuiInputTextFlags_ParseEmptyRefVal wraps ImGuiInputTextFlags_ParseEmptyRefVal.
	GuiInputTextFlags_ParseEmptyRefVal GuiInputTextFlags = C.ImGuiInputTextFlags_ParseEmptyRefVal
	// GuiInputTextFlags_Password wraps ImGuiInputTextFlags_Password.
	GuiInputTextFlags_Password GuiInputTextFlags = C.ImGuiInputTextFlags_Password
	// GuiInputTextFlags_ReadOnly wraps ImGuiInputTextFlags_ReadOnly.
	GuiInputTextFlags_ReadOnly GuiInputTextFlags = C.ImGuiInputTextFlags_ReadOnly
	// GuiInputTextFlags_WordWrap wraps ImGuiInputTextFlags_WordWrap.
	GuiInputTextFlags_WordWrap GuiInputTextFlags = C.ImGuiInputTextFlags_WordWrap
)

func (e GuiInputTextFlags) String() string {
	return fmt.Sprintf("ImGuiInputTextFlags(%b)", e)
}

// GuiInputTextFlagsPrivate wraps the bitmask ImGuiInputTextFlagsPrivate.
type GuiInputTextFlagsPrivate int32

const (
	// GuiInputTextFlags_LocalizeDecimalPoint wraps ImGuiInputTextFlags_LocalizeDecimalPoint.
	GuiInputTextFlags_LocalizeDecimalPoint GuiInputTextFlagsPrivate = C.ImGuiInputTextFlags_LocalizeDecimalPoint
	// GuiInputTextFlags_MergedItem wraps ImGuiInputTextFlags_MergedItem.
	GuiInputTextFlags_MergedItem GuiInputTextFlagsPrivate = C.ImGuiInputTextFlags_MergedItem
	// GuiInputTextFlags_Multiline wraps ImGuiInputTextFlags_Multiline.
	GuiInputTextFlags_Multiline GuiInputTextFlagsPrivate = C.ImGuiInputTextFlags_Multiline
)

func (e GuiInputTextFlagsPrivate) String() string {
	return fmt.Sprintf("ImGuiInputTextFlagsPrivate(%b)", e)
}

// GuiItemFlags wraps the bitmask ImGuiItemFlags.
type GuiItemFlags int32

const (
	// GuiItemFlags_AllowDuplicateId wraps ImGuiItemFlags_AllowDuplicateId.
	GuiItemFlags_AllowDuplicateId GuiItemFlags = C.ImGuiItemFlags_AllowDuplicateId
	// GuiItemFlags_AutoClosePopups wraps ImGuiItemFlags_AutoClosePopups.
	GuiItemFlags_AutoClosePopups GuiItemFlags = C.ImGuiItemFlags_AutoClosePopups
	// GuiItemFlags_ButtonRepeat wraps ImGuiItemFlags_ButtonRepeat.
	GuiItemFlags_ButtonRepeat GuiItemFlags = C.ImGuiItemFlags_ButtonRepeat
	// GuiItemFlags_NoNav wraps ImGuiItemFlags_NoNav.
	GuiItemFlags_NoNav GuiItemFlags = C.ImGuiItemFlags_NoNav
	// GuiItemFlags_NoNavDefaultFocus wraps ImGuiItemFlags_NoNavDefaultFocus.
	GuiItemFlags_NoNavDefaultFocus GuiItemFlags = C.ImGuiItemFlags_NoNavDefaultFocus
	// GuiItemFlags_NoTabStop wraps ImGuiItemFlags_NoTabStop.
	GuiItemFlags_NoTabStop GuiItemFlags = C.ImGuiItemFlags_NoTabStop
	// GuiItemFlags_None wraps ImGuiItemFlags_None.
	GuiItemFlags_None GuiItemFlags = C.ImGuiItemFlags_None
)

func (e GuiItemFlags) String() string {
	return fmt.Sprintf("ImGuiItemFlags(%b)", e)
}

// GuiItemFlagsPrivate wraps the bitmask ImGuiItemFlagsPrivate.
type GuiItemFlagsPrivate int32

const (
	// GuiItemFlags_AllowOverlap wraps ImGuiItemFlags_AllowOverlap.
	GuiItemFlags_AllowOverlap GuiItemFlagsPrivate = C.ImGuiItemFlags_AllowOverlap
	// GuiItemFlags_Default wraps ImGuiItemFlags_Default_.
	GuiItemFlags_Default GuiItemFlagsPrivate = C.ImGuiItemFlags_Default_
	// GuiItemFlags_Disabled wraps ImGuiItemFlags_Disabled.
	GuiItemFlags_Disabled GuiItemFlagsPrivate = C.ImGuiItemFlags_Disabled
	// GuiItemFlags_HasSelectionUserData wraps ImGuiItemFlags_HasSelectionUserData.
	GuiItemFlags_HasSelectionUserData GuiItemFlagsPrivate = C.ImGuiItemFlags_HasSelectionUserData
	// GuiItemFlags_Inputable wraps ImGuiItemFlags_Inputable.
	GuiItemFlags_Inputable GuiItemFlagsPrivate = C.ImGuiItemFlags_Inputable
	// GuiItemFlags_IsMultiSelect wraps ImGuiItemFlags_IsMultiSelect.
	GuiItemFlags_IsMultiSelect GuiItemFlagsPrivate = C.ImGuiItemFlags_IsMultiSelect
	// GuiItemFlags_MixedValue wraps ImGuiItemFlags_MixedValue.
	GuiItemFlags_MixedValue GuiItemFlagsPrivate = C.ImGuiItemFlags_MixedValue
	// GuiItemFlags_NoFocus wraps ImGuiItemFlags_NoFocus.
	GuiItemFlags_NoFocus GuiItemFlagsPrivate = C.ImGuiItemFlags_NoFocus
	// GuiItemFlags_NoMarkEdited wraps ImGuiItemFlags_NoMarkEdited.
	GuiItemFlags_NoMarkEdited GuiItemFlagsPrivate = C.ImGuiItemFlags_NoMarkEdited
	// GuiItemFlags_NoNavDisableMouseHover wraps ImGuiItemFlags_NoNavDisableMouseHover.
	GuiItemFlags_NoNavDisableMouseHover GuiItemFlagsPrivate = C.ImGuiItemFlags_NoNavDisableMouseHover
	// GuiItemFlags_NoWindowHoverableCheck wraps ImGuiItemFlags_NoWindowHoverableCheck.
	GuiItemFlags_NoWindowHoverableCheck GuiItemFlagsPrivate = C.ImGuiItemFlags_NoWindowHoverableCheck
	// GuiItemFlags_ReadOnly wraps ImGuiItemFlags_ReadOnly.
	GuiItemFlags_ReadOnly GuiItemFlagsPrivate = C.ImGuiItemFlags_ReadOnly
)

func (e GuiItemFlagsPrivate) String() string {
	return fmt.Sprintf("ImGuiItemFlagsPrivate(%b)", e)
}

// GuiItemStatusFlags wraps the bitmask ImGuiItemStatusFlags.
type GuiItemStatusFlags int32

const (
	// GuiItemStatusFlags_Deactivated wraps ImGuiItemStatusFlags_Deactivated.
	GuiItemStatusFlags_Deactivated GuiItemStatusFlags = C.ImGuiItemStatusFlags_Deactivated
	// GuiItemStatusFlags_Edited wraps ImGuiItemStatusFlags_Edited.
	GuiItemStatusFlags_Edited GuiItemStatusFlags = C.ImGuiItemStatusFlags_Edited
	// GuiItemStatusFlags_HasClipRect wraps ImGuiItemStatusFlags_HasClipRect.
	GuiItemStatusFlags_HasClipRect GuiItemStatusFlags = C.ImGuiItemStatusFlags_HasClipRect
	// GuiItemStatusFlags_HasDeactivated wraps ImGuiItemStatusFlags_HasDeactivated.
	GuiItemStatusFlags_HasDeactivated GuiItemStatusFlags = C.ImGuiItemStatusFlags_HasDeactivated
	// GuiItemStatusFlags_HasDisplayRect wraps ImGuiItemStatusFlags_HasDisplayRect.
	GuiItemStatusFlags_HasDisplayRect GuiItemStatusFlags = C.ImGuiItemStatusFlags_HasDisplayRect
	// GuiItemStatusFlags_HasShortcut wraps ImGuiItemStatusFlags_HasShortcut.
	GuiItemStatusFlags_HasShortcut GuiItemStatusFlags = C.ImGuiItemStatusFlags_HasShortcut
	// GuiItemStatusFlags_HoveredRect wraps ImGuiItemStatusFlags_HoveredRect.
	GuiItemStatusFlags_HoveredRect GuiItemStatusFlags = C.ImGuiItemStatusFlags_HoveredRect
	// GuiItemStatusFlags_HoveredWindow wraps ImGuiItemStatusFlags_HoveredWindow.
	GuiItemStatusFlags_HoveredWindow GuiItemStatusFlags = C.ImGuiItemStatusFlags_HoveredWindow
	// GuiItemStatusFlags_None wraps ImGuiItemStatusFlags_None.
	GuiItemStatusFlags_None GuiItemStatusFlags = C.ImGuiItemStatusFlags_None
	// GuiItemStatusFlags_ToggledOpen wraps ImGuiItemStatusFlags_ToggledOpen.
	GuiItemStatusFlags_ToggledOpen GuiItemStatusFlags = C.ImGuiItemStatusFlags_ToggledOpen
	// GuiItemStatusFlags_ToggledSelection wraps ImGuiItemStatusFlags_ToggledSelection.
	GuiItemStatusFlags_ToggledSelection GuiItemStatusFlags = C.ImGuiItemStatusFlags_ToggledSelection
	// GuiItemStatusFlags_Visible wraps ImGuiItemStatusFlags_Visible.
	GuiItemStatusFlags_Visible GuiItemStatusFlags = C.ImGuiItemStatusFlags_Visible
)

func (e GuiItemStatusFlags) String() string {
	return fmt.Sprintf("ImGuiItemStatusFlags(%b)", e)
}

// GuiListClipperFlags wraps the bitmask ImGuiListClipperFlags.
type GuiListClipperFlags int32

const (
	// GuiListClipperFlags_NoSetTableRowCounters wraps ImGuiListClipperFlags_NoSetTableRowCounters.
	GuiListClipperFlags_NoSetTableRowCounters GuiListClipperFlags = C.ImGuiListClipperFlags_NoSetTableRowCounters
	// GuiListClipperFlags_None wraps ImGuiListClipperFlags_None.
	GuiListClipperFlags_None GuiListClipperFlags = C.ImGuiListClipperFlags_None
)

func (e GuiListClipperFlags) String() string {
	return fmt.Sprintf("ImGuiListClipperFlags(%b)", e)
}

// GuiLogFlags wraps the bitmask ImGuiLogFlags.
type GuiLogFlags int32

const (
	// GuiLogFlags_None wraps ImGuiLogFlags_None.
	GuiLogFlags_None GuiLogFlags = C.ImGuiLogFlags_None
	// GuiLogFlags_OutputBuffer wraps ImGuiLogFlags_OutputBuffer.
	GuiLogFlags_OutputBuffer GuiLogFlags = C.ImGuiLogFlags_OutputBuffer
	// GuiLogFlags_OutputClipboard wraps ImGuiLogFlags_OutputClipboard.
	GuiLogFlags_OutputClipboard GuiLogFlags = C.ImGuiLogFlags_OutputClipboard
	// GuiLogFlags_OutputFile wraps ImGuiLogFlags_OutputFile.
	GuiLogFlags_OutputFile GuiLogFlags = C.ImGuiLogFlags_OutputFile
	// GuiLogFlags_OutputMask wraps ImGuiLogFlags_OutputMask_.
	GuiLogFlags_OutputMask GuiLogFlags = C.ImGuiLogFlags_OutputMask_
	// GuiLogFlags_OutputTTY wraps ImGuiLogFlags_OutputTTY.
	GuiLogFlags_OutputTTY GuiLogFlags = C.ImGuiLogFlags_OutputTTY
)

func (e GuiLogFlags) String() string {
	return fmt.Sprintf("ImGuiLogFlags(%b)", e)
}

// GuiMultiSelectFlags wraps the bitmask ImGuiMultiSelectFlags.
type GuiMultiSelectFlags int32

const (
	// GuiMultiSelectFlags_BoxSelect1d wraps ImGuiMultiSelectFlags_BoxSelect1d.
	GuiMultiSelectFlags_BoxSelect1d GuiMultiSelectFlags = C.ImGuiMultiSelectFlags_BoxSelect1d
	// GuiMultiSelectFlags_BoxSelect2d wraps ImGuiMultiSelectFlags_BoxSelect2d.
	GuiMultiSelectFlags_BoxSelect2d GuiMultiSelectFlags = C.ImGuiMultiSelectFlags_BoxSelect2d
	// GuiMultiSelectFlags_BoxSelectNoScroll wraps ImGuiMultiSelectFlags_BoxSelectNoScroll.
	GuiMultiSelectFlags_BoxSelectNoScroll GuiMultiSelectFlags = C.ImGuiMultiSelectFlags_BoxSelectNoScroll
	// GuiMultiSelectFlags_ClearOnClickVoid wraps ImGuiMultiSelectFlags_ClearOnClickVoid.
	GuiMultiSelectFlags_ClearOnClickVoid GuiMultiSelectFlags = C.ImGuiMultiSelectFlags_ClearOnClickVoid
	// GuiMultiSelectFlags_ClearOnEscape wraps ImGuiMultiSelectFlags_ClearOnEscape.
	GuiMultiSelectFlags_ClearOnEscape GuiMultiSelectFlags = C.ImGuiMultiSelectFlags_ClearOnEscape
	// GuiMultiSelectFlags_NavWrapX wraps ImGuiMultiSelectFlags_NavWrapX.
	GuiMultiSelectFlags_NavWrapX GuiMultiSelectFlags = C.ImGuiMultiSelectFlags_NavWrapX
	// GuiMultiSelectFlags_NoAutoClear wraps ImGuiMultiSelectFlags_NoAutoClear.
	GuiMultiSelectFlags_NoAutoClear GuiMultiSelectFlags = C.ImGuiMultiSelectFlags_NoAutoClear
	// GuiMultiSelectFlags_NoAutoClearOnReselect wraps ImGuiMultiSelectFlags_NoAutoClearOnReselect.
	GuiMultiSelectFlags_NoAutoClearOnReselect GuiMultiSelectFlags = C.ImGuiMultiSelectFlags_NoAutoClearOnReselect
	// GuiMultiSelectFlags_NoAutoSelect wraps ImGuiMultiSelectFlags_NoAutoSelect.
	GuiMultiSelectFlags_NoAutoSelect GuiMultiSelectFlags = C.ImGuiMultiSelectFlags_NoAutoSelect
	// GuiMultiSelectFlags_NoRangeSelect wraps ImGuiMultiSelectFlags_NoRangeSelect.
	GuiMultiSelectFlags_NoRangeSelect GuiMultiSelectFlags = C.ImGuiMultiSelectFlags_NoRangeSelect
	// GuiMultiSelectFlags_NoSelectAll wraps ImGuiMultiSelectFlags_NoSelectAll.
	GuiMultiSelectFlags_NoSelectAll GuiMultiSelectFlags = C.ImGuiMultiSelectFlags_NoSelectAll
	// GuiMultiSelectFlags_None wraps ImGuiMultiSelectFlags_None.
	GuiMultiSelectFlags_None GuiMultiSelectFlags = C.ImGuiMultiSelectFlags_None
	// GuiMultiSelectFlags_ScopeRect wraps ImGuiMultiSelectFlags_ScopeRect.
	GuiMultiSelectFlags_ScopeRect GuiMultiSelectFlags = C.ImGuiMultiSelectFlags_ScopeRect
	// GuiMultiSelectFlags_ScopeWindow wraps ImGuiMultiSelectFlags_ScopeWindow.
	GuiMultiSelectFlags_ScopeWindow GuiMultiSelectFlags = C.ImGuiMultiSelectFlags_ScopeWindow
	// GuiMultiSelectFlags_SelectOnClick wraps ImGuiMultiSelectFlags_SelectOnClick.
	GuiMultiSelectFlags_SelectOnClick GuiMultiSelectFlags = C.ImGuiMultiSelectFlags_SelectOnClick
	// GuiMultiSelectFlags_SelectOnClickRelease wraps ImGuiMultiSelectFlags_SelectOnClickRelease.
	GuiMultiSelectFlags_SelectOnClickRelease GuiMultiSelectFlags = C.ImGuiMultiSelectFlags_SelectOnClickRelease
	// GuiMultiSelectFlags_SingleSelect wraps ImGuiMultiSelectFlags_SingleSelect.
	GuiMultiSelectFlags_SingleSelect GuiMultiSelectFlags = C.ImGuiMultiSelectFlags_SingleSelect
)

func (e GuiMultiSelectFlags) String() string {
	return fmt.Sprintf("ImGuiMultiSelectFlags(%b)", e)
}

// GuiNavMoveFlags wraps the bitmask ImGuiNavMoveFlags.
type GuiNavMoveFlags int32

const (
	// GuiNavMoveFlags_Activate wraps ImGuiNavMoveFlags_Activate.
	GuiNavMoveFlags_Activate GuiNavMoveFlags = C.ImGuiNavMoveFlags_Activate
	// GuiNavMoveFlags_AllowCurrentNavId wraps ImGuiNavMoveFlags_AllowCurrentNavId.
	GuiNavMoveFlags_AllowCurrentNavId GuiNavMoveFlags = C.ImGuiNavMoveFlags_AllowCurrentNavId
	// GuiNavMoveFlags_AlsoScoreVisibleSet wraps ImGuiNavMoveFlags_AlsoScoreVisibleSet.
	GuiNavMoveFlags_AlsoScoreVisibleSet GuiNavMoveFlags = C.ImGuiNavMoveFlags_AlsoScoreVisibleSet
	// GuiNavMoveFlags_DebugNoResult wraps ImGuiNavMoveFlags_DebugNoResult.
	GuiNavMoveFlags_DebugNoResult GuiNavMoveFlags = C.ImGuiNavMoveFlags_DebugNoResult
	// GuiNavMoveFlags_FocusApi wraps ImGuiNavMoveFlags_FocusApi.
	GuiNavMoveFlags_FocusApi GuiNavMoveFlags = C.ImGuiNavMoveFlags_FocusApi
	// GuiNavMoveFlags_Forwarded wraps ImGuiNavMoveFlags_Forwarded.
	GuiNavMoveFlags_Forwarded GuiNavMoveFlags = C.ImGuiNavMoveFlags_Forwarded
	// GuiNavMoveFlags_IsPageMove wraps ImGuiNavMoveFlags_IsPageMove.
	GuiNavMoveFlags_IsPageMove GuiNavMoveFlags = C.ImGuiNavMoveFlags_IsPageMove
	// GuiNavMoveFlags_IsTabbing wraps ImGuiNavMoveFlags_IsTabbing.
	GuiNavMoveFlags_IsTabbing GuiNavMoveFlags = C.ImGuiNavMoveFlags_IsTabbing
	// GuiNavMoveFlags_LoopX wraps ImGuiNavMoveFlags_LoopX.
	GuiNavMoveFlags_LoopX GuiNavMoveFlags = C.ImGuiNavMoveFlags_LoopX
	// GuiNavMoveFlags_LoopY wraps ImGuiNavMoveFlags_LoopY.
	GuiNavMoveFlags_LoopY GuiNavMoveFlags = C.ImGuiNavMoveFlags_LoopY
	// GuiNavMoveFlags_NoClearActiveId wraps ImGuiNavMoveFlags_NoClearActiveId.
	GuiNavMoveFlags_NoClearActiveId GuiNavMoveFlags = C.ImGuiNavMoveFlags_NoClearActiveId
	// GuiNavMoveFlags_NoSelect wraps ImGuiNavMoveFlags_NoSelect.
	GuiNavMoveFlags_NoSelect GuiNavMoveFlags = C.ImGuiNavMoveFlags_NoSelect
	// GuiNavMoveFlags_NoSetNavCursorVisible wraps ImGuiNavMoveFlags_NoSetNavCursorVisible.
	GuiNavMoveFlags_NoSetNavCursorVisible GuiNavMoveFlags = C.ImGuiNavMoveFlags_NoSetNavCursorVisible
	// GuiNavMoveFlags_None wraps ImGuiNavMoveFlags_None.
	GuiNavMoveFlags_None GuiNavMoveFlags = C.ImGuiNavMoveFlags_None
	// GuiNavMoveFlags_ScrollToEdgeY wraps ImGuiNavMoveFlags_ScrollToEdgeY.
	GuiNavMoveFlags_ScrollToEdgeY GuiNavMoveFlags = C.ImGuiNavMoveFlags_ScrollToEdgeY
	// GuiNavMoveFlags_WrapMask wraps ImGuiNavMoveFlags_WrapMask_.
	GuiNavMoveFlags_WrapMask GuiNavMoveFlags = C.ImGuiNavMoveFlags_WrapMask_
	// GuiNavMoveFlags_WrapX wraps ImGuiNavMoveFlags_WrapX.
	GuiNavMoveFlags_WrapX GuiNavMoveFlags = C.ImGuiNavMoveFlags_WrapX
	// GuiNavMoveFlags_WrapY wraps ImGuiNavMoveFlags_WrapY.
	GuiNavMoveFlags_WrapY GuiNavMoveFlags = C.ImGuiNavMoveFlags_WrapY
)

func (e GuiNavMoveFlags) String() string {
	return fmt.Sprintf("ImGuiNavMoveFlags(%b)", e)
}

// GuiNavRenderCursorFlags wraps the bitmask ImGuiNavRenderCursorFlags.
type GuiNavRenderCursorFlags int32

const (
	// GuiNavRenderCursorFlags_AlwaysDraw wraps ImGuiNavRenderCursorFlags_AlwaysDraw.
	GuiNavRenderCursorFlags_AlwaysDraw GuiNavRenderCursorFlags = C.ImGuiNavRenderCursorFlags_AlwaysDraw
	// GuiNavRenderCursorFlags_Compact wraps ImGuiNavRenderCursorFlags_Compact.
	GuiNavRenderCursorFlags_Compact GuiNavRenderCursorFlags = C.ImGuiNavRenderCursorFlags_Compact
	// GuiNavRenderCursorFlags_NoRounding wraps ImGuiNavRenderCursorFlags_NoRounding.
	GuiNavRenderCursorFlags_NoRounding GuiNavRenderCursorFlags = C.ImGuiNavRenderCursorFlags_NoRounding
	// GuiNavRenderCursorFlags_None wraps ImGuiNavRenderCursorFlags_None.
	GuiNavRenderCursorFlags_None GuiNavRenderCursorFlags = C.ImGuiNavRenderCursorFlags_None
)

func (e GuiNavRenderCursorFlags) String() string {
	return fmt.Sprintf("ImGuiNavRenderCursorFlags(%b)", e)
}

// GuiNextItemDataFlags wraps the bitmask ImGuiNextItemDataFlags.
type GuiNextItemDataFlags int32

const (
	// GuiNextItemDataFlags_HasOpen wraps ImGuiNextItemDataFlags_HasOpen.
	GuiNextItemDataFlags_HasOpen GuiNextItemDataFlags = C.ImGuiNextItemDataFlags_HasOpen
	// GuiNextItemDataFlags_HasRefVal wraps ImGuiNextItemDataFlags_HasRefVal.
	GuiNextItemDataFlags_HasRefVal GuiNextItemDataFlags = C.ImGuiNextItemDataFlags_HasRefVal
	// GuiNextItemDataFlags_HasShortcut wraps ImGuiNextItemDataFlags_HasShortcut.
	GuiNextItemDataFlags_HasShortcut GuiNextItemDataFlags = C.ImGuiNextItemDataFlags_HasShortcut
	// GuiNextItemDataFlags_HasStorageID wraps ImGuiNextItemDataFlags_HasStorageID.
	GuiNextItemDataFlags_HasStorageID GuiNextItemDataFlags = C.ImGuiNextItemDataFlags_HasStorageID
	// GuiNextItemDataFlags_HasWidth wraps ImGuiNextItemDataFlags_HasWidth.
	GuiNextItemDataFlags_HasWidth GuiNextItemDataFlags = C.ImGuiNextItemDataFlags_HasWidth
	// GuiNextItemDataFlags_None wraps ImGuiNextItemDataFlags_None.
	GuiNextItemDataFlags_None GuiNextItemDataFlags = C.ImGuiNextItemDataFlags_None
)

func (e GuiNextItemDataFlags) String() string {
	return fmt.Sprintf("ImGuiNextItemDataFlags(%b)", e)
}

// GuiNextWindowDataFlags wraps the bitmask ImGuiNextWindowDataFlags.
type GuiNextWindowDataFlags int32

const (
	// GuiNextWindowDataFlags_HasBgAlpha wraps ImGuiNextWindowDataFlags_HasBgAlpha.
	GuiNextWindowDataFlags_HasBgAlpha GuiNextWindowDataFlags = C.ImGuiNextWindowDataFlags_HasBgAlpha
	// GuiNextWindowDataFlags_HasChildFlags wraps ImGuiNextWindowDataFlags_HasChildFlags.
	GuiNextWindowDataFlags_HasChildFlags GuiNextWindowDataFlags = C.ImGuiNextWindowDataFlags_HasChildFlags
	// GuiNextWindowDataFlags_HasCollapsed wraps ImGuiNextWindowDataFlags_HasCollapsed.
	GuiNextWindowDataFlags_HasCollapsed GuiNextWindowDataFlags = C.ImGuiNextWindowDataFlags_HasCollapsed
	// GuiNextWindowDataFlags_HasContentSize wraps ImGuiNextWindowDataFlags_HasContentSize.
	GuiNextWindowDataFlags_HasContentSize GuiNextWindowDataFlags = C.ImGuiNextWindowDataFlags_HasContentSize
	// GuiNextWindowDataFlags_HasDock wraps ImGuiNextWindowDataFlags_HasDock.
	GuiNextWindowDataFlags_HasDock GuiNextWindowDataFlags = C.ImGuiNextWindowDataFlags_HasDock
	// GuiNextWindowDataFlags_HasFocus wraps ImGuiNextWindowDataFlags_HasFocus.
	GuiNextWindowDataFlags_HasFocus GuiNextWindowDataFlags = C.ImGuiNextWindowDataFlags_HasFocus
	// GuiNextWindowDataFlags_HasPos wraps ImGuiNextWindowDataFlags_HasPos.
	GuiNextWindowDataFlags_HasPos GuiNextWindowDataFlags = C.ImGuiNextWindowDataFlags_HasPos
	// GuiNextWindowDataFlags_HasRefreshPolicy wraps ImGuiNextWindowDataFlags_HasRefreshPolicy.
	GuiNextWindowDataFlags_HasRefreshPolicy GuiNextWindowDataFlags = C.ImGuiNextWindowDataFlags_HasRefreshPolicy
	// GuiNextWindowDataFlags_HasScroll wraps ImGuiNextWindowDataFlags_HasScroll.
	GuiNextWindowDataFlags_HasScroll GuiNextWindowDataFlags = C.ImGuiNextWindowDataFlags_HasScroll
	// GuiNextWindowDataFlags_HasSize wraps ImGuiNextWindowDataFlags_HasSize.
	GuiNextWindowDataFlags_HasSize GuiNextWindowDataFlags = C.ImGuiNextWindowDataFlags_HasSize
	// GuiNextWindowDataFlags_HasSizeConstraint wraps ImGuiNextWindowDataFlags_HasSizeConstraint.
	GuiNextWindowDataFlags_HasSizeConstraint GuiNextWindowDataFlags = C.ImGuiNextWindowDataFlags_HasSizeConstraint
	// GuiNextWindowDataFlags_HasViewport wraps ImGuiNextWindowDataFlags_HasViewport.
	GuiNextWindowDataFlags_HasViewport GuiNextWindowDataFlags = C.ImGuiNextWindowDataFlags_HasViewport
	// GuiNextWindowDataFlags_HasWindowClass wraps ImGuiNextWindowDataFlags_HasWindowClass.
	GuiNextWindowDataFlags_HasWindowClass GuiNextWindowDataFlags = C.ImGuiNextWindowDataFlags_HasWindowClass
	// GuiNextWindowDataFlags_HasWindowFlags wraps ImGuiNextWindowDataFlags_HasWindowFlags.
	GuiNextWindowDataFlags_HasWindowFlags GuiNextWindowDataFlags = C.ImGuiNextWindowDataFlags_HasWindowFlags
	// GuiNextWindowDataFlags_None wraps ImGuiNextWindowDataFlags_None.
	GuiNextWindowDataFlags_None GuiNextWindowDataFlags = C.ImGuiNextWindowDataFlags_None
)

func (e GuiNextWindowDataFlags) String() string {
	return fmt.Sprintf("ImGuiNextWindowDataFlags(%b)", e)
}

// GuiOldColumnFlags wraps the bitmask ImGuiOldColumnFlags.
type GuiOldColumnFlags int32

const (
	// GuiOldColumnFlags_GrowParentContentsSize wraps ImGuiOldColumnFlags_GrowParentContentsSize.
	GuiOldColumnFlags_GrowParentContentsSize GuiOldColumnFlags = C.ImGuiOldColumnFlags_GrowParentContentsSize
	// GuiOldColumnFlags_NoBorder wraps ImGuiOldColumnFlags_NoBorder.
	GuiOldColumnFlags_NoBorder GuiOldColumnFlags = C.ImGuiOldColumnFlags_NoBorder
	// GuiOldColumnFlags_NoForceWithinWindow wraps ImGuiOldColumnFlags_NoForceWithinWindow.
	GuiOldColumnFlags_NoForceWithinWindow GuiOldColumnFlags = C.ImGuiOldColumnFlags_NoForceWithinWindow
	// GuiOldColumnFlags_NoPreserveWidths wraps ImGuiOldColumnFlags_NoPreserveWidths.
	GuiOldColumnFlags_NoPreserveWidths GuiOldColumnFlags = C.ImGuiOldColumnFlags_NoPreserveWidths
	// GuiOldColumnFlags_NoResize wraps ImGuiOldColumnFlags_NoResize.
	GuiOldColumnFlags_NoResize GuiOldColumnFlags = C.ImGuiOldColumnFlags_NoResize
	// GuiOldColumnFlags_None wraps ImGuiOldColumnFlags_None.
	GuiOldColumnFlags_None GuiOldColumnFlags = C.ImGuiOldColumnFlags_None
)

func (e GuiOldColumnFlags) String() string {
	return fmt.Sprintf("ImGuiOldColumnFlags(%b)", e)
}

// GuiPopupFlags wraps the bitmask ImGuiPopupFlags.
type GuiPopupFlags int32

const (
	// GuiPopupFlags_AnyPopup wraps ImGuiPopupFlags_AnyPopup.
	GuiPopupFlags_AnyPopup GuiPopupFlags = C.ImGuiPopupFlags_AnyPopup
	// GuiPopupFlags_AnyPopupId wraps ImGuiPopupFlags_AnyPopupId.
	GuiPopupFlags_AnyPopupId GuiPopupFlags = C.ImGuiPopupFlags_AnyPopupId
	// GuiPopupFlags_AnyPopupLevel wraps ImGuiPopupFlags_AnyPopupLevel.
	GuiPopupFlags_AnyPopupLevel GuiPopupFlags = C.ImGuiPopupFlags_AnyPopupLevel
	// GuiPopupFlags_MouseButtonDefault wraps ImGuiPopupFlags_MouseButtonDefault_.
	GuiPopupFlags_MouseButtonDefault GuiPopupFlags = C.ImGuiPopupFlags_MouseButtonDefault_
	// GuiPopupFlags_MouseButtonLeft wraps ImGuiPopupFlags_MouseButtonLeft.
	GuiPopupFlags_MouseButtonLeft GuiPopupFlags = C.ImGuiPopupFlags_MouseButtonLeft
	// GuiPopupFlags_MouseButtonMask wraps ImGuiPopupFlags_MouseButtonMask_.
	GuiPopupFlags_MouseButtonMask GuiPopupFlags = C.ImGuiPopupFlags_MouseButtonMask_
	// GuiPopupFlags_MouseButtonMiddle wraps ImGuiPopupFlags_MouseButtonMiddle.
	GuiPopupFlags_MouseButtonMiddle GuiPopupFlags = C.ImGuiPopupFlags_MouseButtonMiddle
	// GuiPopupFlags_MouseButtonRight wraps ImGuiPopupFlags_MouseButtonRight.
	GuiPopupFlags_MouseButtonRight GuiPopupFlags = C.ImGuiPopupFlags_MouseButtonRight
	// GuiPopupFlags_NoOpenOverExistingPopup wraps ImGuiPopupFlags_NoOpenOverExistingPopup.
	GuiPopupFlags_NoOpenOverExistingPopup GuiPopupFlags = C.ImGuiPopupFlags_NoOpenOverExistingPopup
	// GuiPopupFlags_NoOpenOverItems wraps ImGuiPopupFlags_NoOpenOverItems.
	GuiPopupFlags_NoOpenOverItems GuiPopupFlags = C.ImGuiPopupFlags_NoOpenOverItems
	// GuiPopupFlags_NoReopen wraps ImGuiPopupFlags_NoReopen.
	GuiPopupFlags_NoReopen GuiPopupFlags = C.ImGuiPopupFlags_NoReopen
	// GuiPopupFlags_None wraps ImGuiPopupFlags_None.
	GuiPopupFlags_None GuiPopupFlags = C.ImGuiPopupFlags_None
)

func (e GuiPopupFlags) String() string {
	return fmt.Sprintf("ImGuiPopupFlags(%b)", e)
}

// GuiScrollFlags wraps the bitmask ImGuiScrollFlags.
type GuiScrollFlags int32

const (
	// GuiScrollFlags_AlwaysCenterX wraps ImGuiScrollFlags_AlwaysCenterX.
	GuiScrollFlags_AlwaysCenterX GuiScrollFlags = C.ImGuiScrollFlags_AlwaysCenterX
	// GuiScrollFlags_AlwaysCenterY wraps ImGuiScrollFlags_AlwaysCenterY.
	GuiScrollFlags_AlwaysCenterY GuiScrollFlags = C.ImGuiScrollFlags_AlwaysCenterY
	// GuiScrollFlags_KeepVisibleCenterX wraps ImGuiScrollFlags_KeepVisibleCenterX.
	GuiScrollFlags_KeepVisibleCenterX GuiScrollFlags = C.ImGuiScrollFlags_KeepVisibleCenterX
	// GuiScrollFlags_KeepVisibleCenterY wraps ImGuiScrollFlags_KeepVisibleCenterY.
	GuiScrollFlags_KeepVisibleCenterY GuiScrollFlags = C.ImGuiScrollFlags_KeepVisibleCenterY
	// GuiScrollFlags_KeepVisibleEdgeX wraps ImGuiScrollFlags_KeepVisibleEdgeX.
	GuiScrollFlags_KeepVisibleEdgeX GuiScrollFlags = C.ImGuiScrollFlags_KeepVisibleEdgeX
	// GuiScrollFlags_KeepVisibleEdgeY wraps ImGuiScrollFlags_KeepVisibleEdgeY.
	GuiScrollFlags_KeepVisibleEdgeY GuiScrollFlags = C.ImGuiScrollFlags_KeepVisibleEdgeY
	// GuiScrollFlags_MaskX wraps ImGuiScrollFlags_MaskX_.
	GuiScrollFlags_MaskX GuiScrollFlags = C.ImGuiScrollFlags_MaskX_
	// GuiScrollFlags_MaskY wraps ImGuiScrollFlags_MaskY_.
	GuiScrollFlags_MaskY GuiScrollFlags = C.ImGuiScrollFlags_MaskY_
	// GuiScrollFlags_NoScrollParent wraps ImGuiScrollFlags_NoScrollParent.
	GuiScrollFlags_NoScrollParent GuiScrollFlags = C.ImGuiScrollFlags_NoScrollParent
	// GuiScrollFlags_None wraps ImGuiScrollFlags_None.
	GuiScrollFlags_None GuiScrollFlags = C.ImGuiScrollFlags_None
)

func (e GuiScrollFlags) String() string {
	return fmt.Sprintf("ImGuiScrollFlags(%b)", e)
}

// GuiSelectableFlags wraps the bitmask ImGuiSelectableFlags.
type GuiSelectableFlags int32

const (
	// GuiSelectableFlags_AllowDoubleClick wraps ImGuiSelectableFlags_AllowDoubleClick.
	GuiSelectableFlags_AllowDoubleClick GuiSelectableFlags = C.ImGuiSelectableFlags_AllowDoubleClick
	// GuiSelectableFlags_AllowOverlap wraps ImGuiSelectableFlags_AllowOverlap.
	GuiSelectableFlags_AllowOverlap GuiSelectableFlags = C.ImGuiSelectableFlags_AllowOverlap
	// GuiSelectableFlags_Disabled wraps ImGuiSelectableFlags_Disabled.
	GuiSelectableFlags_Disabled GuiSelectableFlags = C.ImGuiSelectableFlags_Disabled
	// GuiSelectableFlags_Highlight wraps ImGuiSelectableFlags_Highlight.
	GuiSelectableFlags_Highlight GuiSelectableFlags = C.ImGuiSelectableFlags_Highlight
	// GuiSelectableFlags_NoAutoClosePopups wraps ImGuiSelectableFlags_NoAutoClosePopups.
	GuiSelectableFlags_NoAutoClosePopups GuiSelectableFlags = C.ImGuiSelectableFlags_NoAutoClosePopups
	// GuiSelectableFlags_None wraps ImGuiSelectableFlags_None.
	GuiSelectableFlags_None GuiSelectableFlags = C.ImGuiSelectableFlags_None
	// GuiSelectableFlags_SelectOnNav wraps ImGuiSelectableFlags_SelectOnNav.
	GuiSelectableFlags_SelectOnNav GuiSelectableFlags = C.ImGuiSelectableFlags_SelectOnNav
	// GuiSelectableFlags_SpanAllColumns wraps ImGuiSelectableFlags_SpanAllColumns.
	GuiSelectableFlags_SpanAllColumns GuiSelectableFlags = C.ImGuiSelectableFlags_SpanAllColumns
)

func (e GuiSelectableFlags) String() string {
	return fmt.Sprintf("ImGuiSelectableFlags(%b)", e)
}

// GuiSelectableFlagsPrivate wraps the bitmask ImGuiSelectableFlagsPrivate.
type GuiSelectableFlagsPrivate int32

const (
	// GuiSelectableFlags_NoHoldingActiveID wraps ImGuiSelectableFlags_NoHoldingActiveID.
	GuiSelectableFlags_NoHoldingActiveID GuiSelectableFlagsPrivate = C.ImGuiSelectableFlags_NoHoldingActiveID
	// GuiSelectableFlags_NoPadWithHalfSpacing wraps ImGuiSelectableFlags_NoPadWithHalfSpacing.
	GuiSelectableFlags_NoPadWithHalfSpacing GuiSelectableFlagsPrivate = C.ImGuiSelectableFlags_NoPadWithHalfSpacing
	// GuiSelectableFlags_NoSetKeyOwner wraps ImGuiSelectableFlags_NoSetKeyOwner.
	GuiSelectableFlags_NoSetKeyOwner GuiSelectableFlagsPrivate = C.ImGuiSelectableFlags_NoSetKeyOwner
	// GuiSelectableFlags_SelectOnClick wraps ImGuiSelectableFlags_SelectOnClick.
	GuiSelectableFlags_SelectOnClick GuiSelectableFlagsPrivate = C.ImGuiSelectableFlags_SelectOnClick
	// GuiSelectableFlags_SelectOnRelease wraps ImGuiSelectableFlags_SelectOnRelease.
	GuiSelectableFlags_SelectOnRelease GuiSelectableFlagsPrivate = C.ImGuiSelectableFlags_SelectOnRelease
	// GuiSelectableFlags_SetNavIdOnHover wraps ImGuiSelectableFlags_SetNavIdOnHover.
	GuiSelectableFlags_SetNavIdOnHover GuiSelectableFlagsPrivate = C.ImGuiSelectableFlags_SetNavIdOnHover
	// GuiSelectableFlags_SpanAvailWidth wraps ImGuiSelectableFlags_SpanAvailWidth.
	GuiSelectableFlags_SpanAvailWidth GuiSelectableFlagsPrivate = C.ImGuiSelectableFlags_SpanAvailWidth
)

func (e GuiSelectableFlagsPrivate) String() string {
	return fmt.Sprintf("ImGuiSelectableFlagsPrivate(%b)", e)
}

// GuiSeparatorFlags wraps the bitmask ImGuiSeparatorFlags.
type GuiSeparatorFlags int32

const (
	// GuiSeparatorFlags_Horizontal wraps ImGuiSeparatorFlags_Horizontal.
	GuiSeparatorFlags_Horizontal GuiSeparatorFlags = C.ImGuiSeparatorFlags_Horizontal
	// GuiSeparatorFlags_None wraps ImGuiSeparatorFlags_None.
	GuiSeparatorFlags_None GuiSeparatorFlags = C.ImGuiSeparatorFlags_None
	// GuiSeparatorFlags_SpanAllColumns wraps ImGuiSeparatorFlags_SpanAllColumns.
	GuiSeparatorFlags_SpanAllColumns GuiSeparatorFlags = C.ImGuiSeparatorFlags_SpanAllColumns
	// GuiSeparatorFlags_Vertical wraps ImGuiSeparatorFlags_Vertical.
	GuiSeparatorFlags_Vertical GuiSeparatorFlags = C.ImGuiSeparatorFlags_Vertical
)

func (e GuiSeparatorFlags) String() string {
	return fmt.Sprintf("ImGuiSeparatorFlags(%b)", e)
}

// GuiSliderFlags wraps the bitmask ImGuiSliderFlags.
type GuiSliderFlags int32

const (
	// GuiSliderFlags_AlwaysClamp wraps ImGuiSliderFlags_AlwaysClamp.
	GuiSliderFlags_AlwaysClamp GuiSliderFlags = C.ImGuiSliderFlags_AlwaysClamp
	// GuiSliderFlags_ClampOnInput wraps ImGuiSliderFlags_ClampOnInput.
	GuiSliderFlags_ClampOnInput GuiSliderFlags = C.ImGuiSliderFlags_ClampOnInput
	// GuiSliderFlags_ClampZeroRange wraps ImGuiSliderFlags_ClampZeroRange.
	GuiSliderFlags_ClampZeroRange GuiSliderFlags = C.ImGuiSliderFlags_ClampZeroRange
	// GuiSliderFlags_InvalidMask wraps ImGuiSliderFlags_InvalidMask_.
	GuiSliderFlags_InvalidMask GuiSliderFlags = C.ImGuiSliderFlags_InvalidMask_
	// GuiSliderFlags_Logarithmic wraps ImGuiSliderFlags_Logarithmic.
	GuiSliderFlags_Logarithmic GuiSliderFlags = C.ImGuiSliderFlags_Logarithmic
	// GuiSliderFlags_NoInput wraps ImGuiSliderFlags_NoInput.
	GuiSliderFlags_NoInput GuiSliderFlags = C.ImGuiSliderFlags_NoInput
	// GuiSliderFlags_NoRoundToFormat wraps ImGuiSliderFlags_NoRoundToFormat.
	GuiSliderFlags_NoRoundToFormat GuiSliderFlags = C.ImGuiSliderFlags_NoRoundToFormat
	// GuiSliderFlags_NoSpeedTweaks wraps ImGuiSliderFlags_NoSpeedTweaks.
	GuiSliderFlags_NoSpeedTweaks GuiSliderFlags = C.ImGuiSliderFlags_NoSpeedTweaks
	// GuiSliderFlags_None wraps ImGuiSliderFlags_None.
	GuiSliderFlags_None GuiSliderFlags = C.ImGuiSliderFlags_None
	// GuiSliderFlags_WrapAround wraps ImGuiSliderFlags_WrapAround.
	GuiSliderFlags_WrapAround GuiSliderFlags = C.ImGuiSliderFlags_WrapAround
)

func (e GuiSliderFlags) String() string {
	return fmt.Sprintf("ImGuiSliderFlags(%b)", e)
}

// GuiSliderFlagsPrivate wraps the bitmask ImGuiSliderFlagsPrivate.
type GuiSliderFlagsPrivate int32

const (
	// GuiSliderFlags_ReadOnly wraps ImGuiSliderFlags_ReadOnly.
	GuiSliderFlags_ReadOnly GuiSliderFlagsPrivate = C.ImGuiSliderFlags_ReadOnly
	// GuiSliderFlags_Vertical wraps ImGuiSliderFlags_Vertical.
	GuiSliderFlags_Vertical GuiSliderFlagsPrivate = C.ImGuiSliderFlags_Vertical
)

func (e GuiSliderFlagsPrivate) String() string {
	return fmt.Sprintf("ImGuiSliderFlagsPrivate(%b)", e)
}

// GuiTabBarFlags wraps the bitmask ImGuiTabBarFlags.
type GuiTabBarFlags int32

const (
	// GuiTabBarFlags_AutoSelectNewTabs wraps ImGuiTabBarFlags_AutoSelectNewTabs.
	GuiTabBarFlags_AutoSelectNewTabs GuiTabBarFlags = C.ImGuiTabBarFlags_AutoSelectNewTabs
	// GuiTabBarFlags_DrawSelectedOverline wraps ImGuiTabBarFlags_DrawSelectedOverline.
	GuiTabBarFlags_DrawSelectedOverline GuiTabBarFlags = C.ImGuiTabBarFlags_DrawSelectedOverline
	// GuiTabBarFlags_FittingPolicyDefault wraps ImGuiTabBarFlags_FittingPolicyDefault_.
	GuiTabBarFlags_FittingPolicyDefault GuiTabBarFlags = C.ImGuiTabBarFlags_FittingPolicyDefault_
	// GuiTabBarFlags_FittingPolicyMask wraps ImGuiTabBarFlags_FittingPolicyMask_.
	GuiTabBarFlags_FittingPolicyMask GuiTabBarFlags = C.ImGuiTabBarFlags_FittingPolicyMask_
	// GuiTabBarFlags_FittingPolicyMixed wraps ImGuiTabBarFlags_FittingPolicyMixed.
	GuiTabBarFlags_FittingPolicyMixed GuiTabBarFlags = C.ImGuiTabBarFlags_FittingPolicyMixed
	// GuiTabBarFlags_FittingPolicyScroll wraps ImGuiTabBarFlags_FittingPolicyScroll.
	GuiTabBarFlags_FittingPolicyScroll GuiTabBarFlags = C.ImGuiTabBarFlags_FittingPolicyScroll
	// GuiTabBarFlags_FittingPolicyShrink wraps ImGuiTabBarFlags_FittingPolicyShrink.
	GuiTabBarFlags_FittingPolicyShrink GuiTabBarFlags = C.ImGuiTabBarFlags_FittingPolicyShrink
	// GuiTabBarFlags_NoCloseWithMiddleMouseButton wraps ImGuiTabBarFlags_NoCloseWithMiddleMouseButton.
	GuiTabBarFlags_NoCloseWithMiddleMouseButton GuiTabBarFlags = C.ImGuiTabBarFlags_NoCloseWithMiddleMouseButton
	// GuiTabBarFlags_NoTabListScrollingButtons wraps ImGuiTabBarFlags_NoTabListScrollingButtons.
	GuiTabBarFlags_NoTabListScrollingButtons GuiTabBarFlags = C.ImGuiTabBarFlags_NoTabListScrollingButtons
	// GuiTabBarFlags_NoTooltip wraps ImGuiTabBarFlags_NoTooltip.
	GuiTabBarFlags_NoTooltip GuiTabBarFlags = C.ImGuiTabBarFlags_NoTooltip
	// GuiTabBarFlags_None wraps ImGuiTabBarFlags_None.
	GuiTabBarFlags_None GuiTabBarFlags = C.ImGuiTabBarFlags_None
	// GuiTabBarFlags_Reorderable wraps ImGuiTabBarFlags_Reorderable.
	GuiTabBarFlags_Reorderable GuiTabBarFlags = C.ImGuiTabBarFlags_Reorderable
	// GuiTabBarFlags_TabListPopupButton wraps ImGuiTabBarFlags_TabListPopupButton.
	GuiTabBarFlags_TabListPopupButton GuiTabBarFlags = C.ImGuiTabBarFlags_TabListPopupButton
)

func (e GuiTabBarFlags) String() string {
	return fmt.Sprintf("ImGuiTabBarFlags(%b)", e)
}

// GuiTabBarFlagsPrivate wraps the bitmask ImGuiTabBarFlagsPrivate.
type GuiTabBarFlagsPrivate int32

const (
	// GuiTabBarFlags_DockNode wraps ImGuiTabBarFlags_DockNode.
	GuiTabBarFlags_DockNode GuiTabBarFlagsPrivate = C.ImGuiTabBarFlags_DockNode
	// GuiTabBarFlags_IsFocused wraps ImGuiTabBarFlags_IsFocused.
	GuiTabBarFlags_IsFocused GuiTabBarFlagsPrivate = C.ImGuiTabBarFlags_IsFocused
	// GuiTabBarFlags_SaveSettings wraps ImGuiTabBarFlags_SaveSettings.
	GuiTabBarFlags_SaveSettings GuiTabBarFlagsPrivate = C.ImGuiTabBarFlags_SaveSettings
)

func (e GuiTabBarFlagsPrivate) String() string {
	return fmt.Sprintf("ImGuiTabBarFlagsPrivate(%b)", e)
}

// GuiTabItemFlags wraps the bitmask ImGuiTabItemFlags.
type GuiTabItemFlags int32

const (
	// GuiTabItemFlags_Leading wraps ImGuiTabItemFlags_Leading.
	GuiTabItemFlags_Leading GuiTabItemFlags = C.ImGuiTabItemFlags_Leading
	// GuiTabItemFlags_NoAssumedClosure wraps ImGuiTabItemFlags_NoAssumedClosure.
	GuiTabItemFlags_NoAssumedClosure GuiTabItemFlags = C.ImGuiTabItemFlags_NoAssumedClosure
	// GuiTabItemFlags_NoCloseWithMiddleMouseButton wraps ImGuiTabItemFlags_NoCloseWithMiddleMouseButton.
	GuiTabItemFlags_NoCloseWithMiddleMouseButton GuiTabItemFlags = C.ImGuiTabItemFlags_NoCloseWithMiddleMouseButton
	// GuiTabItemFlags_NoPushId wraps ImGuiTabItemFlags_NoPushId.
	GuiTabItemFlags_NoPushId GuiTabItemFlags = C.ImGuiTabItemFlags_NoPushId
	// GuiTabItemFlags_NoReorder wraps ImGuiTabItemFlags_NoReorder.
	GuiTabItemFlags_NoReorder GuiTabItemFlags = C.ImGuiTabItemFlags_NoReorder
	// GuiTabItemFlags_NoTooltip wraps ImGuiTabItemFlags_NoTooltip.
	GuiTabItemFlags_NoTooltip GuiTabItemFlags = C.ImGuiTabItemFlags_NoTooltip
	// GuiTabItemFlags_None wraps ImGuiTabItemFlags_None.
	GuiTabItemFlags_None GuiTabItemFlags = C.ImGuiTabItemFlags_None
	// GuiTabItemFlags_SetSelected wraps ImGuiTabItemFlags_SetSelected.
	GuiTabItemFlags_SetSelected GuiTabItemFlags = C.ImGuiTabItemFlags_SetSelected
	// GuiTabItemFlags_Trailing wraps ImGuiTabItemFlags_Trailing.
	GuiTabItemFlags_Trailing GuiTabItemFlags = C.ImGuiTabItemFlags_Trailing
	// GuiTabItemFlags_UnsavedDocument wraps ImGuiTabItemFlags_UnsavedDocument.
	GuiTabItemFlags_UnsavedDocument GuiTabItemFlags = C.ImGuiTabItemFlags_UnsavedDocument
)

func (e GuiTabItemFlags) String() string {
	return fmt.Sprintf("ImGuiTabItemFlags(%b)", e)
}

// GuiTabItemFlagsPrivate wraps the bitmask ImGuiTabItemFlagsPrivate.
type GuiTabItemFlagsPrivate int32

const (
	// GuiTabItemFlags_Button wraps ImGuiTabItemFlags_Button.
	GuiTabItemFlags_Button GuiTabItemFlagsPrivate = C.ImGuiTabItemFlags_Button
	// GuiTabItemFlags_Invisible wraps ImGuiTabItemFlags_Invisible.
	GuiTabItemFlags_Invisible GuiTabItemFlagsPrivate = C.ImGuiTabItemFlags_Invisible
	// GuiTabItemFlags_NoCloseButton wraps ImGuiTabItemFlags_NoCloseButton.
	GuiTabItemFlags_NoCloseButton GuiTabItemFlagsPrivate = C.ImGuiTabItemFlags_NoCloseButton
	// GuiTabItemFlags_SectionMask wraps ImGuiTabItemFlags_SectionMask_.
	GuiTabItemFlags_SectionMask GuiTabItemFlagsPrivate = C.ImGuiTabItemFlags_SectionMask_
	// GuiTabItemFlags_Unsorted wraps ImGuiTabItemFlags_Unsorted.
	GuiTabItemFlags_Unsorted GuiTabItemFlagsPrivate = C.ImGuiTabItemFlags_Unsorted
)

func (e GuiTabItemFlagsPrivate) String() string {
	return fmt.Sprintf("ImGuiTabItemFlagsPrivate(%b)", e)
}

// GuiTableColumnFlags wraps the bitmask ImGuiTableColumnFlags.
type GuiTableColumnFlags int32

const (
	// GuiTableColumnFlags_AngledHeader wraps ImGuiTableColumnFlags_AngledHeader.
	GuiTableColumnFlags_AngledHeader GuiTableColumnFlags = C.ImGuiTableColumnFlags_AngledHeader
	// GuiTableColumnFlags_DefaultHide wraps ImGuiTableColumnFlags_DefaultHide.
	GuiTableColumnFlags_DefaultHide GuiTableColumnFlags = C.ImGuiTableColumnFlags_DefaultHide
	// GuiTableColumnFlags_DefaultSort wraps ImGuiTableColumnFlags_DefaultSort.
	GuiTableColumnFlags_DefaultSort GuiTableColumnFlags = C.ImGuiTableColumnFlags_DefaultSort
	// GuiTableColumnFlags_Disabled wraps ImGuiTableColumnFlags_Disabled.
	GuiTableColumnFlags_Disabled GuiTableColumnFlags = C.ImGuiTableColumnFlags_Disabled
	// GuiTableColumnFlags_IndentDisable wraps ImGuiTableColumnFlags_IndentDisable.
	GuiTableColumnFlags_IndentDisable GuiTableColumnFlags = C.ImGuiTableColumnFlags_IndentDisable
	// GuiTableColumnFlags_IndentEnable wraps ImGuiTableColumnFlags_IndentEnable.
	GuiTableColumnFlags_IndentEnable GuiTableColumnFlags = C.ImGuiTableColumnFlags_IndentEnable
	// GuiTableColumnFlags_IndentMask wraps ImGuiTableColumnFlags_IndentMask_.
	GuiTableColumnFlags_IndentMask GuiTableColumnFlags = C.ImGuiTableColumnFlags_IndentMask_
	// GuiTableColumnFlags_IsEnabled wraps ImGuiTableColumnFlags_IsEnabled.
	GuiTableColumnFlags_IsEnabled GuiTableColumnFlags = C.ImGuiTableColumnFlags_IsEnabled
	// GuiTableColumnFlags_IsHovered wraps ImGuiTableColumnFlags_IsHovered.
	GuiTableColumnFlags_IsHovered GuiTableColumnFlags = C.ImGuiTableColumnFlags_IsHovered
	// GuiTableColumnFlags_IsSorted wraps ImGuiTableColumnFlags_IsSorted.
	GuiTableColumnFlags_IsSorted GuiTableColumnFlags = C.ImGuiTableColumnFlags_IsSorted
	// GuiTableColumnFlags_IsVisible wraps ImGuiTableColumnFlags_IsVisible.
	GuiTableColumnFlags_IsVisible GuiTableColumnFlags = C.ImGuiTableColumnFlags_IsVisible
	// GuiTableColumnFlags_NoClip wraps ImGuiTableColumnFlags_NoClip.
	GuiTableColumnFlags_NoClip GuiTableColumnFlags = C.ImGuiTableColumnFlags_NoClip
	// GuiTableColumnFlags_NoDirectResize wraps ImGuiTableColumnFlags_NoDirectResize_.
	GuiTableColumnFlags_NoDirectResize GuiTableColumnFlags = C.ImGuiTableColumnFlags_NoDirectResize_
	// GuiTableColumnFlags_NoHeaderLabel wraps ImGuiTableColumnFlags_NoHeaderLabel.
	GuiTableColumnFlags_NoHeaderLabel GuiTableColumnFlags = C.ImGuiTableColumnFlags_NoHeaderLabel
	// GuiTableColumnFlags_NoHeaderWidth wraps ImGuiTableColumnFlags_NoHeaderWidth.
	GuiTableColumnFlags_NoHeaderWidth GuiTableColumnFlags = C.ImGuiTableColumnFlags_NoHeaderWidth
	// GuiTableColumnFlags_NoHide wraps ImGuiTableColumnFlags_NoHide.
	GuiTableColumnFlags_NoHide GuiTableColumnFlags = C.ImGuiTableColumnFlags_NoHide
	// GuiTableColumnFlags_NoReorder wraps ImGuiTableColumnFlags_NoReorder.
	GuiTableColumnFlags_NoReorder GuiTableColumnFlags = C.ImGuiTableColumnFlags_NoReorder
	// GuiTableColumnFlags_NoResize wraps ImGuiTableColumnFlags_NoResize.
	GuiTableColumnFlags_NoResize GuiTableColumnFlags = C.ImGuiTableColumnFlags_NoResize
	// GuiTableColumnFlags_NoSort wraps ImGuiTableColumnFlags_NoSort.
	GuiTableColumnFlags_NoSort GuiTableColumnFlags = C.ImGuiTableColumnFlags_NoSort
	// GuiTableColumnFlags_NoSortAscending wraps ImGuiTableColumnFlags_NoSortAscending.
	GuiTableColumnFlags_NoSortAscending GuiTableColumnFlags = C.ImGuiTableColumnFlags_NoSortAscending
	// GuiTableColumnFlags_NoSortDescending wraps ImGuiTableColumnFlags_NoSortDescending.
	GuiTableColumnFlags_NoSortDescending GuiTableColumnFlags = C.ImGuiTableColumnFlags_NoSortDescending
	// GuiTableColumnFlags_None wraps ImGuiTableColumnFlags_None.
	GuiTableColumnFlags_None GuiTableColumnFlags = C.ImGuiTableColumnFlags_None
	// GuiTableColumnFlags_PreferSortAscending wraps ImGuiTableColumnFlags_PreferSortAscending.
	GuiTableColumnFlags_PreferSortAscending GuiTableColumnFlags = C.ImGuiTableColumnFlags_PreferSortAscending
	// GuiTableColumnFlags_PreferSortDescending wraps ImGuiTableColumnFlags_PreferSortDescending.
	GuiTableColumnFlags_PreferSortDescending GuiTableColumnFlags = C.ImGuiTableColumnFlags_PreferSortDescending
	// GuiTableColumnFlags_StatusMask wraps ImGuiTableColumnFlags_StatusMask_.
	GuiTableColumnFlags_StatusMask GuiTableColumnFlags = C.ImGuiTableColumnFlags_StatusMask_
	// GuiTableColumnFlags_WidthFixed wraps ImGuiTableColumnFlags_WidthFixed.
	GuiTableColumnFlags_WidthFixed GuiTableColumnFlags = C.ImGuiTableColumnFlags_WidthFixed
	// GuiTableColumnFlags_WidthMask wraps ImGuiTableColumnFlags_WidthMask_.
	GuiTableColumnFlags_WidthMask GuiTableColumnFlags = C.ImGuiTableColumnFlags_WidthMask_
	// GuiTableColumnFlags_WidthStretch wraps ImGuiTableColumnFlags_WidthStretch.
	GuiTableColumnFlags_WidthStretch GuiTableColumnFlags = C.ImGuiTableColumnFlags_WidthStretch
)

func (e GuiTableColumnFlags) String() string {
	return fmt.Sprintf("ImGuiTableColumnFlags(%b)", e)
}

// GuiTableFlags wraps the bitmask ImGuiTableFlags.
type GuiTableFlags int32

const (
	// GuiTableFlags_Borders wraps ImGuiTableFlags_Borders.
	GuiTableFlags_Borders GuiTableFlags = C.ImGuiTableFlags_Borders
	// GuiTableFlags_BordersH wraps ImGuiTableFlags_BordersH.
	GuiTableFlags_BordersH GuiTableFlags = C.ImGuiTableFlags_BordersH
	// GuiTableFlags_BordersInner wraps ImGuiTableFlags_BordersInner.
	GuiTableFlags_BordersInner GuiTableFlags = C.ImGuiTableFlags_BordersInner
	// GuiTableFlags_BordersInnerH wraps ImGuiTableFlags_BordersInnerH.
	GuiTableFlags_BordersInnerH GuiTableFlags = C.ImGuiTableFlags_BordersInnerH
	// GuiTableFlags_BordersInnerV wraps ImGuiTableFlags_BordersInnerV.
	GuiTableFlags_BordersInnerV GuiTableFlags = C.ImGuiTableFlags_BordersInnerV
	// GuiTableFlags_BordersOuter wraps ImGuiTableFlags_BordersOuter.
	GuiTableFlags_BordersOuter GuiTableFlags = C.ImGuiTableFlags_BordersOuter
	// GuiTableFlags_BordersOuterH wraps ImGuiTableFlags_BordersOuterH.
	GuiTableFlags_BordersOuterH GuiTableFlags = C.ImGuiTableFlags_BordersOuterH
	// GuiTableFlags_BordersOuterV wraps ImGuiTableFlags_BordersOuterV.
	GuiTableFlags_BordersOuterV GuiTableFlags = C.ImGuiTableFlags_BordersOuterV
	// GuiTableFlags_BordersV wraps ImGuiTableFlags_BordersV.
	GuiTableFlags_BordersV GuiTableFlags = C.ImGuiTableFlags_BordersV
	// GuiTableFlags_ContextMenuInBody wraps ImGuiTableFlags_ContextMenuInBody.
	GuiTableFlags_ContextMenuInBody GuiTableFlags = C.ImGuiTableFlags_ContextMenuInBody
	// GuiTableFlags_Hideable wraps ImGuiTableFlags_Hideable.
	GuiTableFlags_Hideable GuiTableFlags = C.ImGuiTableFlags_Hideable
	// GuiTableFlags_HighlightHoveredColumn wraps ImGuiTableFlags_HighlightHoveredColumn.
	GuiTableFlags_HighlightHoveredColumn GuiTableFlags = C.ImGuiTableFlags_HighlightHoveredColumn
	// GuiTableFlags_NoBordersInBody wraps ImGuiTableFlags_NoBordersInBody.
	GuiTableFlags_NoBordersInBody GuiTableFlags = C.ImGuiTableFlags_NoBordersInBody
	// GuiTableFlags_NoBordersInBodyUntilResize wraps ImGuiTableFlags_NoBordersInBodyUntilResize.
	GuiTableFlags_NoBordersInBodyUntilResize GuiTableFlags = C.ImGuiTableFlags_NoBordersInBodyUntilResize
	// GuiTableFlags_NoClip wraps ImGuiTableFlags_NoClip.
	GuiTableFlags_NoClip GuiTableFlags = C.ImGuiTableFlags_NoClip
	// GuiTableFlags_NoHostExtendX wraps ImGuiTableFlags_NoHostExtendX.
	GuiTableFlags_NoHostExtendX GuiTableFlags = C.ImGuiTableFlags_NoHostExtendX
	// GuiTableFlags_NoHostExtendY wraps ImGuiTableFlags_NoHostExtendY.
	GuiTableFlags_NoHostExtendY GuiTableFlags = C.ImGuiTableFlags_NoHostExtendY
	// GuiTableFlags_NoKeepColumnsVisible wraps ImGuiTableFlags_NoKeepColumnsVisible.
	GuiTableFlags_NoKeepColumnsVisible GuiTableFlags = C.ImGuiTableFlags_NoKeepColumnsVisible
	// GuiTableFlags_NoPadInnerX wraps ImGuiTableFlags_NoPadInnerX.
	GuiTableFlags_NoPadInnerX GuiTableFlags = C.ImGuiTableFlags_NoPadInnerX
	// GuiTableFlags_NoPadOuterX wraps ImGuiTableFlags_NoPadOuterX.
	GuiTableFlags_NoPadOuterX GuiTableFlags = C.ImGuiTableFlags_NoPadOuterX
	// GuiTableFlags_NoSavedSettings wraps ImGuiTableFlags_NoSavedSettings.
	GuiTableFlags_NoSavedSettings GuiTableFlags = C.ImGuiTableFlags_NoSavedSettings
	// GuiTableFlags_None wraps ImGuiTableFlags_None.
	GuiTableFlags_None GuiTableFlags = C.ImGuiTableFlags_None
	// GuiTableFlags_PadOuterX wraps ImGuiTableFlags_PadOuterX.
	GuiTableFlags_PadOuterX GuiTableFlags = C.ImGuiTableFlags_PadOuterX
	// GuiTableFlags_PreciseWidths wraps ImGuiTableFlags_PreciseWidths.
	GuiTableFlags_PreciseWidths GuiTableFlags = C.ImGuiTableFlags_PreciseWidths
	// GuiTableFlags_Reorderable wraps ImGuiTableFlags_Reorderable.
	GuiTableFlags_Reorderable GuiTableFlags = C.ImGuiTableFlags_Reorderable
	// GuiTableFlags_Resizable wraps ImGuiTableFlags_Resizable.
	GuiTableFlags_Resizable GuiTableFlags = C.ImGuiTableFlags_Resizable
	// GuiTableFlags_RowBg wraps ImGuiTableFlags_RowBg.
	GuiTableFlags_RowBg GuiTableFlags = C.ImGuiTableFlags_RowBg
	// GuiTableFlags_ScrollX wraps ImGuiTableFlags_ScrollX.
	GuiTableFlags_ScrollX GuiTableFlags = C.ImGuiTableFlags_ScrollX
	// GuiTableFlags_ScrollY wraps ImGuiTableFlags_ScrollY.
	GuiTableFlags_ScrollY GuiTableFlags = C.ImGuiTableFlags_ScrollY
	// GuiTableFlags_SizingFixedFit wraps ImGuiTableFlags_SizingFixedFit.
	GuiTableFlags_SizingFixedFit GuiTableFlags = C.ImGuiTableFlags_SizingFixedFit
	// GuiTableFlags_SizingFixedSame wraps ImGuiTableFlags_SizingFixedSame.
	GuiTableFlags_SizingFixedSame GuiTableFlags = C.ImGuiTableFlags_SizingFixedSame
	// GuiTableFlags_SizingMask wraps ImGuiTableFlags_SizingMask_.
	GuiTableFlags_SizingMask GuiTableFlags = C.ImGuiTableFlags_SizingMask_
	// GuiTableFlags_SizingStretchProp wraps ImGuiTableFlags_SizingStretchProp.
	GuiTableFlags_SizingStretchProp GuiTableFlags = C.ImGuiTableFlags_SizingStretchProp
	// GuiTableFlags_SizingStretchSame wraps ImGuiTableFlags_SizingStretchSame.
	GuiTableFlags_SizingStretchSame GuiTableFlags = C.ImGuiTableFlags_SizingStretchSame
	// GuiTableFlags_SortMulti wraps ImGuiTableFlags_SortMulti.
	GuiTableFlags_SortMulti GuiTableFlags = C.ImGuiTableFlags_SortMulti
	// GuiTableFlags_SortTristate wraps ImGuiTableFlags_SortTristate.
	GuiTableFlags_SortTristate GuiTableFlags = C.ImGuiTableFlags_SortTristate
	// GuiTableFlags_Sortable wraps ImGuiTableFlags_Sortable.
	GuiTableFlags_Sortable GuiTableFlags = C.ImGuiTableFlags_Sortable
)

func (e GuiTableFlags) String() string {
	return fmt.Sprintf("ImGuiTableFlags(%b)", e)
}

// GuiTableRowFlags wraps the bitmask ImGuiTableRowFlags.
type GuiTableRowFlags int32

const (
	// GuiTableRowFlags_Headers wraps ImGuiTableRowFlags_Headers.
	GuiTableRowFlags_Headers GuiTableRowFlags = C.ImGuiTableRowFlags_Headers
	// GuiTableRowFlags_None wraps ImGuiTableRowFlags_None.
	GuiTableRowFlags_None GuiTableRowFlags = C.ImGuiTableRowFlags_None
)

func (e GuiTableRowFlags) String() string {
	return fmt.Sprintf("ImGuiTableRowFlags(%b)", e)
}

// GuiTextFlags wraps the bitmask ImGuiTextFlags.
type GuiTextFlags int32

const (
	// GuiTextFlags_NoWidthForLargeClippedText wraps ImGuiTextFlags_NoWidthForLargeClippedText.
	GuiTextFlags_NoWidthForLargeClippedText GuiTextFlags = C.ImGuiTextFlags_NoWidthForLargeClippedText
	// GuiTextFlags_None wraps ImGuiTextFlags_None.
	GuiTextFlags_None GuiTextFlags = C.ImGuiTextFlags_None
)

func (e GuiTextFlags) String() string {
	return fmt.Sprintf("ImGuiTextFlags(%b)", e)
}

// GuiTooltipFlags wraps the bitmask ImGuiTooltipFlags.
type GuiTooltipFlags int32

const (
	// GuiTooltipFlags_None wraps ImGuiTooltipFlags_None.
	GuiTooltipFlags_None GuiTooltipFlags = C.ImGuiTooltipFlags_None
	// GuiTooltipFlags_OverridePrevious wraps ImGuiTooltipFlags_OverridePrevious.
	GuiTooltipFlags_OverridePrevious GuiTooltipFlags = C.ImGuiTooltipFlags_OverridePrevious
)

func (e GuiTooltipFlags) String() string {
	return fmt.Sprintf("ImGuiTooltipFlags(%b)", e)
}

// GuiTreeNodeFlags wraps the bitmask ImGuiTreeNodeFlags.
type GuiTreeNodeFlags int32

const (
	// GuiTreeNodeFlags_AllowOverlap wraps ImGuiTreeNodeFlags_AllowOverlap.
	GuiTreeNodeFlags_AllowOverlap GuiTreeNodeFlags = C.ImGuiTreeNodeFlags_AllowOverlap
	// GuiTreeNodeFlags_Bullet wraps ImGuiTreeNodeFlags_Bullet.
	GuiTreeNodeFlags_Bullet GuiTreeNodeFlags = C.ImGuiTreeNodeFlags_Bullet
	// GuiTreeNodeFlags_CollapsingHeader wraps ImGuiTreeNodeFlags_CollapsingHeader.
	GuiTreeNodeFlags_CollapsingHeader GuiTreeNodeFlags = C.ImGuiTreeNodeFlags_CollapsingHeader
	// GuiTreeNodeFlags_DefaultOpen wraps ImGuiTreeNodeFlags_DefaultOpen.
	GuiTreeNodeFlags_DefaultOpen GuiTreeNodeFlags = C.ImGuiTreeNodeFlags_DefaultOpen
	// GuiTreeNodeFlags_DrawLinesFull wraps ImGuiTreeNodeFlags_DrawLinesFull.
	GuiTreeNodeFlags_DrawLinesFull GuiTreeNodeFlags = C.ImGuiTreeNodeFlags_DrawLinesFull
	// GuiTreeNodeFlags_DrawLinesNone wraps ImGuiTreeNodeFlags_DrawLinesNone.
	GuiTreeNodeFlags_DrawLinesNone GuiTreeNodeFlags = C.ImGuiTreeNodeFlags_DrawLinesNone
	// GuiTreeNodeFlags_DrawLinesToNodes wraps ImGuiTreeNodeFlags_DrawLinesToNodes.
	GuiTreeNodeFlags_DrawLinesToNodes GuiTreeNodeFlags = C.ImGuiTreeNodeFlags_DrawLinesToNodes
	// GuiTreeNodeFlags_FramePadding wraps ImGuiTreeNodeFlags_FramePadding.
	GuiTreeNodeFlags_FramePadding GuiTreeNodeFlags = C.ImGuiTreeNodeFlags_FramePadding
	// GuiTreeNodeFlags_Framed wraps ImGuiTreeNodeFlags_Framed.
	GuiTreeNodeFlags_Framed GuiTreeNodeFlags = C.ImGuiTreeNodeFlags_Framed
	// GuiTreeNodeFlags_LabelSpanAllColumns wraps ImGuiTreeNodeFlags_LabelSpanAllColumns.
	GuiTreeNodeFlags_LabelSpanAllColumns GuiTreeNodeFlags = C.ImGuiTreeNodeFlags_LabelSpanAllColumns
	// GuiTreeNodeFlags_Leaf wraps ImGuiTreeNodeFlags_Leaf.
	GuiTreeNodeFlags_Leaf GuiTreeNodeFlags = C.ImGuiTreeNodeFlags_Leaf
	// GuiTreeNodeFlags_NavLeftJumpsToParent wraps ImGuiTreeNodeFlags_NavLeftJumpsToParent.
	GuiTreeNodeFlags_NavLeftJumpsToParent GuiTreeNodeFlags = C.ImGuiTreeNodeFlags_NavLeftJumpsToParent
	// GuiTreeNodeFlags_NoAutoOpenOnLog wraps ImGuiTreeNodeFlags_NoAutoOpenOnLog.
	GuiTreeNodeFlags_NoAutoOpenOnLog GuiTreeNodeFlags = C.ImGuiTreeNodeFlags_NoAutoOpenOnLog
	// GuiTreeNodeFlags_NoTreePushOnOpen wraps ImGuiTreeNodeFlags_NoTreePushOnOpen.
	GuiTreeNodeFlags_NoTreePushOnOpen GuiTreeNodeFlags = C.ImGuiTreeNodeFlags_NoTreePushOnOpen
	// GuiTreeNodeFlags_None wraps ImGuiTreeNodeFlags_None.
	GuiTreeNodeFlags_None GuiTreeNodeFlags = C.ImGuiTreeNodeFlags_None
	// GuiTreeNodeFlags_OpenOnArrow wraps ImGuiTreeNodeFlags_OpenOnArrow.
	GuiTreeNodeFlags_OpenOnArrow GuiTreeNodeFlags = C.ImGuiTreeNodeFlags_OpenOnArrow
	// GuiTreeNodeFlags_OpenOnDoubleClick wraps ImGuiTreeNodeFlags_OpenOnDoubleClick.
	GuiTreeNodeFlags_OpenOnDoubleClick GuiTreeNodeFlags = C.ImGuiTreeNodeFlags_OpenOnDoubleClick
	// GuiTreeNodeFlags_Selected wraps ImGuiTreeNodeFlags_Selected.
	GuiTreeNodeFlags_Selected GuiTreeNodeFlags = C.ImGuiTreeNodeFlags_Selected
	// GuiTreeNodeFlags_SpanAllColumns wraps ImGuiTreeNodeFlags_SpanAllColumns.
	GuiTreeNodeFlags_SpanAllColumns GuiTreeNodeFlags = C.ImGuiTreeNodeFlags_SpanAllColumns
	// GuiTreeNodeFlags_SpanAvailWidth wraps ImGuiTreeNodeFlags_SpanAvailWidth.
	GuiTreeNodeFlags_SpanAvailWidth GuiTreeNodeFlags = C.ImGuiTreeNodeFlags_SpanAvailWidth
	// GuiTreeNodeFlags_SpanFullWidth wraps ImGuiTreeNodeFlags_SpanFullWidth.
	GuiTreeNodeFlags_SpanFullWidth GuiTreeNodeFlags = C.ImGuiTreeNodeFlags_SpanFullWidth
	// GuiTreeNodeFlags_SpanLabelWidth wraps ImGuiTreeNodeFlags_SpanLabelWidth.
	GuiTreeNodeFlags_SpanLabelWidth GuiTreeNodeFlags = C.ImGuiTreeNodeFlags_SpanLabelWidth
)

func (e GuiTreeNodeFlags) String() string {
	return fmt.Sprintf("ImGuiTreeNodeFlags(%b)", e)
}

// GuiTreeNodeFlagsPrivate wraps the bitmask ImGuiTreeNodeFlagsPrivate.
type GuiTreeNodeFlagsPrivate int32

const (
	// GuiTreeNodeFlags_ClipLabelForTrailingButton wraps ImGuiTreeNodeFlags_ClipLabelForTrailingButton.
	GuiTreeNodeFlags_ClipLabelForTrailingButton GuiTreeNodeFlagsPrivate = C.ImGuiTreeNodeFlags_ClipLabelForTrailingButton
	// GuiTreeNodeFlags_DrawLinesMask wraps ImGuiTreeNodeFlags_DrawLinesMask_.
	GuiTreeNodeFlags_DrawLinesMask GuiTreeNodeFlagsPrivate = C.ImGuiTreeNodeFlags_DrawLinesMask_
	// GuiTreeNodeFlags_NoNavFocus wraps ImGuiTreeNodeFlags_NoNavFocus.
	GuiTreeNodeFlags_NoNavFocus GuiTreeNodeFlagsPrivate = C.ImGuiTreeNodeFlags_NoNavFocus
	// GuiTreeNodeFlags_OpenOnMask wraps ImGuiTreeNodeFlags_OpenOnMask_.
	GuiTreeNodeFlags_OpenOnMask GuiTreeNodeFlagsPrivate = C.ImGuiTreeNodeFlags_OpenOnMask_
	// GuiTreeNodeFlags_UpsideDownArrow wraps ImGuiTreeNodeFlags_UpsideDownArrow.
	GuiTreeNodeFlags_UpsideDownArrow GuiTreeNodeFlagsPrivate = C.ImGuiTreeNodeFlags_UpsideDownArrow
)

func (e GuiTreeNodeFlagsPrivate) String() string {
	return fmt.Sprintf("ImGuiTreeNodeFlagsPrivate(%b)", e)
}

// GuiTypingSelectFlags wraps the bitmask ImGuiTypingSelectFlags.
type GuiTypingSelectFlags int32

const (
	// GuiTypingSelectFlags_AllowBackspace wraps ImGuiTypingSelectFlags_AllowBackspace.
	GuiTypingSelectFlags_AllowBackspace GuiTypingSelectFlags = C.ImGuiTypingSelectFlags_AllowBackspace
	// GuiTypingSelectFlags_AllowSingleCharMode wraps ImGuiTypingSelectFlags_AllowSingleCharMode.
	GuiTypingSelectFlags_AllowSingleCharMode GuiTypingSelectFlags = C.ImGuiTypingSelectFlags_AllowSingleCharMode
	// GuiTypingSelectFlags_None wraps ImGuiTypingSelectFlags_None.
	GuiTypingSelectFlags_None GuiTypingSelectFlags = C.ImGuiTypingSelectFlags_None
)

func (e GuiTypingSelectFlags) String() string {
	return fmt.Sprintf("ImGuiTypingSelectFlags(%b)", e)
}

// GuiViewportFlags wraps the bitmask ImGuiViewportFlags.
type GuiViewportFlags int32

const (
	// GuiViewportFlags_CanHostOtherWindows wraps ImGuiViewportFlags_CanHostOtherWindows.
	GuiViewportFlags_CanHostOtherWindows GuiViewportFlags = C.ImGuiViewportFlags_CanHostOtherWindows
	// GuiViewportFlags_IsFocused wraps ImGuiViewportFlags_IsFocused.
	GuiViewportFlags_IsFocused GuiViewportFlags = C.ImGuiViewportFlags_IsFocused
	// GuiViewportFlags_IsMinimized wraps ImGuiViewportFlags_IsMinimized.
	GuiViewportFlags_IsMinimized GuiViewportFlags = C.ImGuiViewportFlags_IsMinimized
	// GuiViewportFlags_IsPlatformMonitor wraps ImGuiViewportFlags_IsPlatformMonitor.
	GuiViewportFlags_IsPlatformMonitor GuiViewportFlags = C.ImGuiViewportFlags_IsPlatformMonitor
	// GuiViewportFlags_IsPlatformWindow wraps ImGuiViewportFlags_IsPlatformWindow.
	GuiViewportFlags_IsPlatformWindow GuiViewportFlags = C.ImGuiViewportFlags_IsPlatformWindow
	// GuiViewportFlags_NoAutoMerge wraps ImGuiViewportFlags_NoAutoMerge.
	GuiViewportFlags_NoAutoMerge GuiViewportFlags = C.ImGuiViewportFlags_NoAutoMerge
	// GuiViewportFlags_NoDecoration wraps ImGuiViewportFlags_NoDecoration.
	GuiViewportFlags_NoDecoration GuiViewportFlags = C.ImGuiViewportFlags_NoDecoration
	// GuiViewportFlags_NoFocusOnAppearing wraps ImGuiViewportFlags_NoFocusOnAppearing.
	GuiViewportFlags_NoFocusOnAppearing GuiViewportFlags = C.ImGuiViewportFlags_NoFocusOnAppearing
	// GuiViewportFlags_NoFocusOnClick wraps ImGuiViewportFlags_NoFocusOnClick.
	GuiViewportFlags_NoFocusOnClick GuiViewportFlags = C.ImGuiViewportFlags_NoFocusOnClick
	// GuiViewportFlags_NoInputs wraps ImGuiViewportFlags_NoInputs.
	GuiViewportFlags_NoInputs GuiViewportFlags = C.ImGuiViewportFlags_NoInputs
	// GuiViewportFlags_NoRendererClear wraps ImGuiViewportFlags_NoRendererClear.
	GuiViewportFlags_NoRendererClear GuiViewportFlags = C.ImGuiViewportFlags_NoRendererClear
	// GuiViewportFlags_NoTaskBarIcon wraps ImGuiViewportFlags_NoTaskBarIcon.
	GuiViewportFlags_NoTaskBarIcon GuiViewportFlags = C.ImGuiViewportFlags_NoTaskBarIcon
	// GuiViewportFlags_None wraps ImGuiViewportFlags_None.
	GuiViewportFlags_None GuiViewportFlags = C.ImGuiViewportFlags_None
	// GuiViewportFlags_OwnedByApp wraps ImGuiViewportFlags_OwnedByApp.
	GuiViewportFlags_OwnedByApp GuiViewportFlags = C.ImGuiViewportFlags_OwnedByApp
	// GuiViewportFlags_TopMost wraps ImGuiViewportFlags_TopMost.
	GuiViewportFlags_TopMost GuiViewportFlags = C.ImGuiViewportFlags_TopMost
)

func (e GuiViewportFlags) String() string {
	return fmt.Sprintf("ImGuiViewportFlags(%b)", e)
}

// GuiWindowFlags wraps the bitmask ImGuiWindowFlags.
type GuiWindowFlags int32

const (
	// GuiWindowFlags_AlwaysAutoResize wraps ImGuiWindowFlags_AlwaysAutoResize.
	GuiWindowFlags_AlwaysAutoResize GuiWindowFlags = C.ImGuiWindowFlags_AlwaysAutoResize
	// GuiWindowFlags_AlwaysHorizontalScrollbar wraps ImGuiWindowFlags_AlwaysHorizontalScrollbar.
	GuiWindowFlags_AlwaysHorizontalScrollbar GuiWindowFlags = C.ImGuiWindowFlags_AlwaysHorizontalScrollbar
	// GuiWindowFlags_AlwaysVerticalScrollbar wraps ImGuiWindowFlags_AlwaysVerticalScrollbar.
	GuiWindowFlags_AlwaysVerticalScrollbar GuiWindowFlags = C.ImGuiWindowFlags_AlwaysVerticalScrollbar
	// GuiWindowFlags_ChildMenu wraps ImGuiWindowFlags_ChildMenu.
	GuiWindowFlags_ChildMenu GuiWindowFlags = C.ImGuiWindowFlags_ChildMenu
	// GuiWindowFlags_ChildWindow wraps ImGuiWindowFlags_ChildWindow.
	GuiWindowFlags_ChildWindow GuiWindowFlags = C.ImGuiWindowFlags_ChildWindow
	// GuiWindowFlags_DockNodeHost wraps ImGuiWindowFlags_DockNodeHost.
	GuiWindowFlags_DockNodeHost GuiWindowFlags = C.ImGuiWindowFlags_DockNodeHost
	// GuiWindowFlags_HorizontalScrollbar wraps ImGuiWindowFlags_HorizontalScrollbar.
	GuiWindowFlags_HorizontalScrollbar GuiWindowFlags = C.ImGuiWindowFlags_HorizontalScrollbar
	// GuiWindowFlags_MenuBar wraps ImGuiWindowFlags_MenuBar.
	GuiWindowFlags_MenuBar GuiWindowFlags = C.ImGuiWindowFlags_MenuBar
	// GuiWindowFlags_Modal wraps ImGuiWindowFlags_Modal.
	GuiWindowFlags_Modal GuiWindowFlags = C.ImGuiWindowFlags_Modal
	// GuiWindowFlags_NoBackground wraps ImGuiWindowFlags_NoBackground.
	GuiWindowFlags_NoBackground GuiWindowFlags = C.ImGuiWindowFlags_NoBackground
	// GuiWindowFlags_NoBringToFrontOnFocus wraps ImGuiWindowFlags_NoBringToFrontOnFocus.
	GuiWindowFlags_NoBringToFrontOnFocus GuiWindowFlags = C.ImGuiWindowFlags_NoBringToFrontOnFocus
	// GuiWindowFlags_NoCollapse wraps ImGuiWindowFlags_NoCollapse.
	GuiWindowFlags_NoCollapse GuiWindowFlags = C.ImGuiWindowFlags_NoCollapse
	// GuiWindowFlags_NoDecoration wraps ImGuiWindowFlags_NoDecoration.
	GuiWindowFlags_NoDecoration GuiWindowFlags = C.ImGuiWindowFlags_NoDecoration
	// GuiWindowFlags_NoDocking wraps ImGuiWindowFlags_NoDocking.
	GuiWindowFlags_NoDocking GuiWindowFlags = C.ImGuiWindowFlags_NoDocking
	// GuiWindowFlags_NoFocusOnAppearing wraps ImGuiWindowFlags_NoFocusOnAppearing.
	GuiWindowFlags_NoFocusOnAppearing GuiWindowFlags = C.ImGuiWindowFlags_NoFocusOnAppearing
	// GuiWindowFlags_NoInputs wraps ImGuiWindowFlags_NoInputs.
	GuiWindowFlags_NoInputs GuiWindowFlags = C.ImGuiWindowFlags_NoInputs
	// GuiWindowFlags_NoMouseInputs wraps ImGuiWindowFlags_NoMouseInputs.
	GuiWindowFlags_NoMouseInputs GuiWindowFlags = C.ImGuiWindowFlags_NoMouseInputs
	// GuiWindowFlags_NoMove wraps ImGuiWindowFlags_NoMove.
	GuiWindowFlags_NoMove GuiWindowFlags = C.ImGuiWindowFlags_NoMove
	// GuiWindowFlags_NoNav wraps ImGuiWindowFlags_NoNav.
	GuiWindowFlags_NoNav GuiWindowFlags = C.ImGuiWindowFlags_NoNav
	// GuiWindowFlags_NoNavFocus wraps ImGuiWindowFlags_NoNavFocus.
	GuiWindowFlags_NoNavFocus GuiWindowFlags = C.ImGuiWindowFlags_NoNavFocus
	// GuiWindowFlags_NoNavInputs wraps ImGuiWindowFlags_NoNavInputs.
	GuiWindowFlags_NoNavInputs GuiWindowFlags = C.ImGuiWindowFlags_NoNavInputs
	// GuiWindowFlags_NoResize wraps ImGuiWindowFlags_NoResize.
	GuiWindowFlags_NoResize GuiWindowFlags = C.ImGuiWindowFlags_NoResize
	// GuiWindowFlags_NoSavedSettings wraps ImGuiWindowFlags_NoSavedSettings.
	GuiWindowFlags_NoSavedSettings GuiWindowFlags = C.ImGuiWindowFlags_NoSavedSettings
	// GuiWindowFlags_NoScrollWithMouse wraps ImGuiWindowFlags_NoScrollWithMouse.
	GuiWindowFlags_NoScrollWithMouse GuiWindowFlags = C.ImGuiWindowFlags_NoScrollWithMouse
	// GuiWindowFlags_NoScrollbar wraps ImGuiWindowFlags_NoScrollbar.
	GuiWindowFlags_NoScrollbar GuiWindowFlags = C.ImGuiWindowFlags_NoScrollbar
	// GuiWindowFlags_NoTitleBar wraps ImGuiWindowFlags_NoTitleBar.
	GuiWindowFlags_NoTitleBar GuiWindowFlags = C.ImGuiWindowFlags_NoTitleBar
	// GuiWindowFlags_None wraps ImGuiWindowFlags_None.
	GuiWindowFlags_None GuiWindowFlags = C.ImGuiWindowFlags_None
	// GuiWindowFlags_Popup wraps ImGuiWindowFlags_Popup.
	GuiWindowFlags_Popup GuiWindowFlags = C.ImGuiWindowFlags_Popup
	// GuiWindowFlags_Tooltip wraps ImGuiWindowFlags_Tooltip.
	GuiWindowFlags_Tooltip GuiWindowFlags = C.ImGuiWindowFlags_Tooltip
	// GuiWindowFlags_UnsavedDocument wraps ImGuiWindowFlags_UnsavedDocument.
	GuiWindowFlags_UnsavedDocument GuiWindowFlags = C.ImGuiWindowFlags_UnsavedDocument
)

func (e GuiWindowFlags) String() string {
	return fmt.Sprintf("ImGuiWindowFlags(%b)", e)
}

// GuiWindowRefreshFlags wraps the bitmask ImGuiWindowRefreshFlags.
type GuiWindowRefreshFlags int32

const (
	// GuiWindowRefreshFlags_None wraps ImGuiWindowRefreshFlags_None.
	GuiWindowRefreshFlags_None GuiWindowRefreshFlags = C.ImGuiWindowRefreshFlags_None
	// GuiWindowRefreshFlags_RefreshOnFocus wraps ImGuiWindowRefreshFlags_RefreshOnFocus.
	GuiWindowRefreshFlags_RefreshOnFocus GuiWindowRefreshFlags = C.ImGuiWindowRefreshFlags_RefreshOnFocus
	// GuiWindowRefreshFlags_RefreshOnHover wraps ImGuiWindowRefreshFlags_RefreshOnHover.
	GuiWindowRefreshFlags_RefreshOnHover GuiWindowRefreshFlags = C.ImGuiWindowRefreshFlags_RefreshOnHover
	// GuiWindowRefreshFlags_TryToAvoidRefresh wraps ImGuiWindowRefreshFlags_TryToAvoidRefresh.
	GuiWindowRefreshFlags_TryToAvoidRefresh GuiWindowRefreshFlags = C.ImGuiWindowRefreshFlags_TryToAvoidRefresh
)

func (e GuiWindowRefreshFlags) String() string {
	return fmt.Sprintf("ImGuiWindowRefreshFlags(%b)", e)
}
