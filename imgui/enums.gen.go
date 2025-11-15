package imgui

import "fmt"

// #include "imgui_wrapper.h"
import "C"

// GuiAxis wraps the enum ImGuiAxis.
type GuiAxis int32

const (
	// GuiAxis_None wraps ImGuiAxis_None.
	GuiAxis_None GuiAxis = C.ImGuiAxis_None
	// GuiAxis_X wraps ImGuiAxis_X.
	GuiAxis_X GuiAxis = C.ImGuiAxis_X
	// GuiAxis_Y wraps ImGuiAxis_Y.
	GuiAxis_Y GuiAxis = C.ImGuiAxis_Y
)

func (e GuiAxis) String() string {
	return fmt.Sprintf("ImGuiAxis(%b)", e)
}

// GuiCol wraps the enum ImGuiCol.
type GuiCol int32

const (
	// GuiCol_Border wraps ImGuiCol_Border.
	GuiCol_Border GuiCol = C.ImGuiCol_Border
	// GuiCol_BorderShadow wraps ImGuiCol_BorderShadow.
	GuiCol_BorderShadow GuiCol = C.ImGuiCol_BorderShadow
	// GuiCol_Button wraps ImGuiCol_Button.
	GuiCol_Button GuiCol = C.ImGuiCol_Button
	// GuiCol_ButtonActive wraps ImGuiCol_ButtonActive.
	GuiCol_ButtonActive GuiCol = C.ImGuiCol_ButtonActive
	// GuiCol_ButtonHovered wraps ImGuiCol_ButtonHovered.
	GuiCol_ButtonHovered GuiCol = C.ImGuiCol_ButtonHovered
	// GuiCol_COUNT wraps ImGuiCol_COUNT.
	GuiCol_COUNT GuiCol = C.ImGuiCol_COUNT
	// GuiCol_CheckMark wraps ImGuiCol_CheckMark.
	GuiCol_CheckMark GuiCol = C.ImGuiCol_CheckMark
	// GuiCol_ChildBg wraps ImGuiCol_ChildBg.
	GuiCol_ChildBg GuiCol = C.ImGuiCol_ChildBg
	// GuiCol_DockingEmptyBg wraps ImGuiCol_DockingEmptyBg.
	GuiCol_DockingEmptyBg GuiCol = C.ImGuiCol_DockingEmptyBg
	// GuiCol_DockingPreview wraps ImGuiCol_DockingPreview.
	GuiCol_DockingPreview GuiCol = C.ImGuiCol_DockingPreview
	// GuiCol_DragDropTarget wraps ImGuiCol_DragDropTarget.
	GuiCol_DragDropTarget GuiCol = C.ImGuiCol_DragDropTarget
	// GuiCol_FrameBg wraps ImGuiCol_FrameBg.
	GuiCol_FrameBg GuiCol = C.ImGuiCol_FrameBg
	// GuiCol_FrameBgActive wraps ImGuiCol_FrameBgActive.
	GuiCol_FrameBgActive GuiCol = C.ImGuiCol_FrameBgActive
	// GuiCol_FrameBgHovered wraps ImGuiCol_FrameBgHovered.
	GuiCol_FrameBgHovered GuiCol = C.ImGuiCol_FrameBgHovered
	// GuiCol_Header wraps ImGuiCol_Header.
	GuiCol_Header GuiCol = C.ImGuiCol_Header
	// GuiCol_HeaderActive wraps ImGuiCol_HeaderActive.
	GuiCol_HeaderActive GuiCol = C.ImGuiCol_HeaderActive
	// GuiCol_HeaderHovered wraps ImGuiCol_HeaderHovered.
	GuiCol_HeaderHovered GuiCol = C.ImGuiCol_HeaderHovered
	// GuiCol_InputTextCursor wraps ImGuiCol_InputTextCursor.
	GuiCol_InputTextCursor GuiCol = C.ImGuiCol_InputTextCursor
	// GuiCol_MenuBarBg wraps ImGuiCol_MenuBarBg.
	GuiCol_MenuBarBg GuiCol = C.ImGuiCol_MenuBarBg
	// GuiCol_ModalWindowDimBg wraps ImGuiCol_ModalWindowDimBg.
	GuiCol_ModalWindowDimBg GuiCol = C.ImGuiCol_ModalWindowDimBg
	// GuiCol_NavCursor wraps ImGuiCol_NavCursor.
	GuiCol_NavCursor GuiCol = C.ImGuiCol_NavCursor
	// GuiCol_NavWindowingDimBg wraps ImGuiCol_NavWindowingDimBg.
	GuiCol_NavWindowingDimBg GuiCol = C.ImGuiCol_NavWindowingDimBg
	// GuiCol_NavWindowingHighlight wraps ImGuiCol_NavWindowingHighlight.
	GuiCol_NavWindowingHighlight GuiCol = C.ImGuiCol_NavWindowingHighlight
	// GuiCol_PlotHistogram wraps ImGuiCol_PlotHistogram.
	GuiCol_PlotHistogram GuiCol = C.ImGuiCol_PlotHistogram
	// GuiCol_PlotHistogramHovered wraps ImGuiCol_PlotHistogramHovered.
	GuiCol_PlotHistogramHovered GuiCol = C.ImGuiCol_PlotHistogramHovered
	// GuiCol_PlotLines wraps ImGuiCol_PlotLines.
	GuiCol_PlotLines GuiCol = C.ImGuiCol_PlotLines
	// GuiCol_PlotLinesHovered wraps ImGuiCol_PlotLinesHovered.
	GuiCol_PlotLinesHovered GuiCol = C.ImGuiCol_PlotLinesHovered
	// GuiCol_PopupBg wraps ImGuiCol_PopupBg.
	GuiCol_PopupBg GuiCol = C.ImGuiCol_PopupBg
	// GuiCol_ResizeGrip wraps ImGuiCol_ResizeGrip.
	GuiCol_ResizeGrip GuiCol = C.ImGuiCol_ResizeGrip
	// GuiCol_ResizeGripActive wraps ImGuiCol_ResizeGripActive.
	GuiCol_ResizeGripActive GuiCol = C.ImGuiCol_ResizeGripActive
	// GuiCol_ResizeGripHovered wraps ImGuiCol_ResizeGripHovered.
	GuiCol_ResizeGripHovered GuiCol = C.ImGuiCol_ResizeGripHovered
	// GuiCol_ScrollbarBg wraps ImGuiCol_ScrollbarBg.
	GuiCol_ScrollbarBg GuiCol = C.ImGuiCol_ScrollbarBg
	// GuiCol_ScrollbarGrab wraps ImGuiCol_ScrollbarGrab.
	GuiCol_ScrollbarGrab GuiCol = C.ImGuiCol_ScrollbarGrab
	// GuiCol_ScrollbarGrabActive wraps ImGuiCol_ScrollbarGrabActive.
	GuiCol_ScrollbarGrabActive GuiCol = C.ImGuiCol_ScrollbarGrabActive
	// GuiCol_ScrollbarGrabHovered wraps ImGuiCol_ScrollbarGrabHovered.
	GuiCol_ScrollbarGrabHovered GuiCol = C.ImGuiCol_ScrollbarGrabHovered
	// GuiCol_Separator wraps ImGuiCol_Separator.
	GuiCol_Separator GuiCol = C.ImGuiCol_Separator
	// GuiCol_SeparatorActive wraps ImGuiCol_SeparatorActive.
	GuiCol_SeparatorActive GuiCol = C.ImGuiCol_SeparatorActive
	// GuiCol_SeparatorHovered wraps ImGuiCol_SeparatorHovered.
	GuiCol_SeparatorHovered GuiCol = C.ImGuiCol_SeparatorHovered
	// GuiCol_SliderGrab wraps ImGuiCol_SliderGrab.
	GuiCol_SliderGrab GuiCol = C.ImGuiCol_SliderGrab
	// GuiCol_SliderGrabActive wraps ImGuiCol_SliderGrabActive.
	GuiCol_SliderGrabActive GuiCol = C.ImGuiCol_SliderGrabActive
	// GuiCol_Tab wraps ImGuiCol_Tab.
	GuiCol_Tab GuiCol = C.ImGuiCol_Tab
	// GuiCol_TabDimmed wraps ImGuiCol_TabDimmed.
	GuiCol_TabDimmed GuiCol = C.ImGuiCol_TabDimmed
	// GuiCol_TabDimmedSelected wraps ImGuiCol_TabDimmedSelected.
	GuiCol_TabDimmedSelected GuiCol = C.ImGuiCol_TabDimmedSelected
	// GuiCol_TabDimmedSelectedOverline wraps ImGuiCol_TabDimmedSelectedOverline.
	GuiCol_TabDimmedSelectedOverline GuiCol = C.ImGuiCol_TabDimmedSelectedOverline
	// GuiCol_TabHovered wraps ImGuiCol_TabHovered.
	GuiCol_TabHovered GuiCol = C.ImGuiCol_TabHovered
	// GuiCol_TabSelected wraps ImGuiCol_TabSelected.
	GuiCol_TabSelected GuiCol = C.ImGuiCol_TabSelected
	// GuiCol_TabSelectedOverline wraps ImGuiCol_TabSelectedOverline.
	GuiCol_TabSelectedOverline GuiCol = C.ImGuiCol_TabSelectedOverline
	// GuiCol_TableBorderLight wraps ImGuiCol_TableBorderLight.
	GuiCol_TableBorderLight GuiCol = C.ImGuiCol_TableBorderLight
	// GuiCol_TableBorderStrong wraps ImGuiCol_TableBorderStrong.
	GuiCol_TableBorderStrong GuiCol = C.ImGuiCol_TableBorderStrong
	// GuiCol_TableHeaderBg wraps ImGuiCol_TableHeaderBg.
	GuiCol_TableHeaderBg GuiCol = C.ImGuiCol_TableHeaderBg
	// GuiCol_TableRowBg wraps ImGuiCol_TableRowBg.
	GuiCol_TableRowBg GuiCol = C.ImGuiCol_TableRowBg
	// GuiCol_TableRowBgAlt wraps ImGuiCol_TableRowBgAlt.
	GuiCol_TableRowBgAlt GuiCol = C.ImGuiCol_TableRowBgAlt
	// GuiCol_Text wraps ImGuiCol_Text.
	GuiCol_Text GuiCol = C.ImGuiCol_Text
	// GuiCol_TextDisabled wraps ImGuiCol_TextDisabled.
	GuiCol_TextDisabled GuiCol = C.ImGuiCol_TextDisabled
	// GuiCol_TextLink wraps ImGuiCol_TextLink.
	GuiCol_TextLink GuiCol = C.ImGuiCol_TextLink
	// GuiCol_TextSelectedBg wraps ImGuiCol_TextSelectedBg.
	GuiCol_TextSelectedBg GuiCol = C.ImGuiCol_TextSelectedBg
	// GuiCol_TitleBg wraps ImGuiCol_TitleBg.
	GuiCol_TitleBg GuiCol = C.ImGuiCol_TitleBg
	// GuiCol_TitleBgActive wraps ImGuiCol_TitleBgActive.
	GuiCol_TitleBgActive GuiCol = C.ImGuiCol_TitleBgActive
	// GuiCol_TitleBgCollapsed wraps ImGuiCol_TitleBgCollapsed.
	GuiCol_TitleBgCollapsed GuiCol = C.ImGuiCol_TitleBgCollapsed
	// GuiCol_TreeLines wraps ImGuiCol_TreeLines.
	GuiCol_TreeLines GuiCol = C.ImGuiCol_TreeLines
	// GuiCol_UnsavedMarker wraps ImGuiCol_UnsavedMarker.
	GuiCol_UnsavedMarker GuiCol = C.ImGuiCol_UnsavedMarker
	// GuiCol_WindowBg wraps ImGuiCol_WindowBg.
	GuiCol_WindowBg GuiCol = C.ImGuiCol_WindowBg
)

func (e GuiCol) String() string {
	return fmt.Sprintf("ImGuiCol(%b)", e)
}

// GuiCond wraps the enum ImGuiCond.
type GuiCond int32

const (
	// GuiCond_Always wraps ImGuiCond_Always.
	GuiCond_Always GuiCond = C.ImGuiCond_Always
	// GuiCond_Appearing wraps ImGuiCond_Appearing.
	GuiCond_Appearing GuiCond = C.ImGuiCond_Appearing
	// GuiCond_FirstUseEver wraps ImGuiCond_FirstUseEver.
	GuiCond_FirstUseEver GuiCond = C.ImGuiCond_FirstUseEver
	// GuiCond_None wraps ImGuiCond_None.
	GuiCond_None GuiCond = C.ImGuiCond_None
	// GuiCond_Once wraps ImGuiCond_Once.
	GuiCond_Once GuiCond = C.ImGuiCond_Once
)

func (e GuiCond) String() string {
	return fmt.Sprintf("ImGuiCond(%b)", e)
}

// GuiContextHookType wraps the enum ImGuiContextHookType.
type GuiContextHookType int32

const (
	// GuiContextHookType_EndFramePost wraps ImGuiContextHookType_EndFramePost.
	GuiContextHookType_EndFramePost GuiContextHookType = C.ImGuiContextHookType_EndFramePost
	// GuiContextHookType_EndFramePre wraps ImGuiContextHookType_EndFramePre.
	GuiContextHookType_EndFramePre GuiContextHookType = C.ImGuiContextHookType_EndFramePre
	// GuiContextHookType_NewFramePost wraps ImGuiContextHookType_NewFramePost.
	GuiContextHookType_NewFramePost GuiContextHookType = C.ImGuiContextHookType_NewFramePost
	// GuiContextHookType_NewFramePre wraps ImGuiContextHookType_NewFramePre.
	GuiContextHookType_NewFramePre GuiContextHookType = C.ImGuiContextHookType_NewFramePre
	// GuiContextHookType_PendingRemoval wraps ImGuiContextHookType_PendingRemoval_.
	GuiContextHookType_PendingRemoval GuiContextHookType = C.ImGuiContextHookType_PendingRemoval_
	// GuiContextHookType_RenderPost wraps ImGuiContextHookType_RenderPost.
	GuiContextHookType_RenderPost GuiContextHookType = C.ImGuiContextHookType_RenderPost
	// GuiContextHookType_RenderPre wraps ImGuiContextHookType_RenderPre.
	GuiContextHookType_RenderPre GuiContextHookType = C.ImGuiContextHookType_RenderPre
	// GuiContextHookType_Shutdown wraps ImGuiContextHookType_Shutdown.
	GuiContextHookType_Shutdown GuiContextHookType = C.ImGuiContextHookType_Shutdown
)

func (e GuiContextHookType) String() string {
	return fmt.Sprintf("ImGuiContextHookType(%b)", e)
}

// GuiDataAuthority wraps the enum ImGuiDataAuthority.
type GuiDataAuthority int32

const (
	// GuiDataAuthority_Auto wraps ImGuiDataAuthority_Auto.
	GuiDataAuthority_Auto GuiDataAuthority = C.ImGuiDataAuthority_Auto
	// GuiDataAuthority_DockNode wraps ImGuiDataAuthority_DockNode.
	GuiDataAuthority_DockNode GuiDataAuthority = C.ImGuiDataAuthority_DockNode
	// GuiDataAuthority_Window wraps ImGuiDataAuthority_Window.
	GuiDataAuthority_Window GuiDataAuthority = C.ImGuiDataAuthority_Window
)

func (e GuiDataAuthority) String() string {
	return fmt.Sprintf("ImGuiDataAuthority(%b)", e)
}

// GuiDataType wraps the enum ImGuiDataType.
type GuiDataType int32

const (
	// GuiDataType_Bool wraps ImGuiDataType_Bool.
	GuiDataType_Bool GuiDataType = C.ImGuiDataType_Bool
	// GuiDataType_COUNT wraps ImGuiDataType_COUNT.
	GuiDataType_COUNT GuiDataType = C.ImGuiDataType_COUNT
	// GuiDataType_Double wraps ImGuiDataType_Double.
	GuiDataType_Double GuiDataType = C.ImGuiDataType_Double
	// GuiDataType_Float wraps ImGuiDataType_Float.
	GuiDataType_Float GuiDataType = C.ImGuiDataType_Float
	// GuiDataType_S16 wraps ImGuiDataType_S16.
	GuiDataType_S16 GuiDataType = C.ImGuiDataType_S16
	// GuiDataType_S32 wraps ImGuiDataType_S32.
	GuiDataType_S32 GuiDataType = C.ImGuiDataType_S32
	// GuiDataType_S64 wraps ImGuiDataType_S64.
	GuiDataType_S64 GuiDataType = C.ImGuiDataType_S64
	// GuiDataType_S8 wraps ImGuiDataType_S8.
	GuiDataType_S8 GuiDataType = C.ImGuiDataType_S8
	// GuiDataType_String wraps ImGuiDataType_String.
	GuiDataType_String GuiDataType = C.ImGuiDataType_String
	// GuiDataType_U16 wraps ImGuiDataType_U16.
	GuiDataType_U16 GuiDataType = C.ImGuiDataType_U16
	// GuiDataType_U32 wraps ImGuiDataType_U32.
	GuiDataType_U32 GuiDataType = C.ImGuiDataType_U32
	// GuiDataType_U64 wraps ImGuiDataType_U64.
	GuiDataType_U64 GuiDataType = C.ImGuiDataType_U64
	// GuiDataType_U8 wraps ImGuiDataType_U8.
	GuiDataType_U8 GuiDataType = C.ImGuiDataType_U8
)

func (e GuiDataType) String() string {
	return fmt.Sprintf("ImGuiDataType(%b)", e)
}

// GuiDataTypePrivate wraps the enum ImGuiDataTypePrivate.
type GuiDataTypePrivate int32

const (
	// GuiDataType_ID wraps ImGuiDataType_ID.
	GuiDataType_ID GuiDataTypePrivate = C.ImGuiDataType_ID
	// GuiDataType_Pointer wraps ImGuiDataType_Pointer.
	GuiDataType_Pointer GuiDataTypePrivate = C.ImGuiDataType_Pointer
)

func (e GuiDataTypePrivate) String() string {
	return fmt.Sprintf("ImGuiDataTypePrivate(%b)", e)
}

// GuiDir wraps the enum ImGuiDir.
type GuiDir int32

const (
	// GuiDir_COUNT wraps ImGuiDir_COUNT.
	GuiDir_COUNT GuiDir = C.ImGuiDir_COUNT
	// GuiDir_Down wraps ImGuiDir_Down.
	GuiDir_Down GuiDir = C.ImGuiDir_Down
	// GuiDir_Left wraps ImGuiDir_Left.
	GuiDir_Left GuiDir = C.ImGuiDir_Left
	// GuiDir_None wraps ImGuiDir_None.
	GuiDir_None GuiDir = C.ImGuiDir_None
	// GuiDir_Right wraps ImGuiDir_Right.
	GuiDir_Right GuiDir = C.ImGuiDir_Right
	// GuiDir_Up wraps ImGuiDir_Up.
	GuiDir_Up GuiDir = C.ImGuiDir_Up
)

func (e GuiDir) String() string {
	return fmt.Sprintf("ImGuiDir(%b)", e)
}

// GuiDockNodeState wraps the enum ImGuiDockNodeState.
type GuiDockNodeState int32

const (
	// GuiDockNodeState_HostWindowHiddenBecauseSingleWindow wraps ImGuiDockNodeState_HostWindowHiddenBecauseSingleWindow.
	GuiDockNodeState_HostWindowHiddenBecauseSingleWindow GuiDockNodeState = C.ImGuiDockNodeState_HostWindowHiddenBecauseSingleWindow
	// GuiDockNodeState_HostWindowHiddenBecauseWindowsAreResizing wraps ImGuiDockNodeState_HostWindowHiddenBecauseWindowsAreResizing.
	GuiDockNodeState_HostWindowHiddenBecauseWindowsAreResizing GuiDockNodeState = C.ImGuiDockNodeState_HostWindowHiddenBecauseWindowsAreResizing
	// GuiDockNodeState_HostWindowVisible wraps ImGuiDockNodeState_HostWindowVisible.
	GuiDockNodeState_HostWindowVisible GuiDockNodeState = C.ImGuiDockNodeState_HostWindowVisible
	// GuiDockNodeState_Unknown wraps ImGuiDockNodeState_Unknown.
	GuiDockNodeState_Unknown GuiDockNodeState = C.ImGuiDockNodeState_Unknown
)

func (e GuiDockNodeState) String() string {
	return fmt.Sprintf("ImGuiDockNodeState(%b)", e)
}

// GuiInputEventType wraps the enum ImGuiInputEventType.
type GuiInputEventType int32

const (
	// GuiInputEventType_COUNT wraps ImGuiInputEventType_COUNT.
	GuiInputEventType_COUNT GuiInputEventType = C.ImGuiInputEventType_COUNT
	// GuiInputEventType_Focus wraps ImGuiInputEventType_Focus.
	GuiInputEventType_Focus GuiInputEventType = C.ImGuiInputEventType_Focus
	// GuiInputEventType_Key wraps ImGuiInputEventType_Key.
	GuiInputEventType_Key GuiInputEventType = C.ImGuiInputEventType_Key
	// GuiInputEventType_MouseButton wraps ImGuiInputEventType_MouseButton.
	GuiInputEventType_MouseButton GuiInputEventType = C.ImGuiInputEventType_MouseButton
	// GuiInputEventType_MousePos wraps ImGuiInputEventType_MousePos.
	GuiInputEventType_MousePos GuiInputEventType = C.ImGuiInputEventType_MousePos
	// GuiInputEventType_MouseViewport wraps ImGuiInputEventType_MouseViewport.
	GuiInputEventType_MouseViewport GuiInputEventType = C.ImGuiInputEventType_MouseViewport
	// GuiInputEventType_MouseWheel wraps ImGuiInputEventType_MouseWheel.
	GuiInputEventType_MouseWheel GuiInputEventType = C.ImGuiInputEventType_MouseWheel
	// GuiInputEventType_None wraps ImGuiInputEventType_None.
	GuiInputEventType_None GuiInputEventType = C.ImGuiInputEventType_None
	// GuiInputEventType_Text wraps ImGuiInputEventType_Text.
	GuiInputEventType_Text GuiInputEventType = C.ImGuiInputEventType_Text
)

func (e GuiInputEventType) String() string {
	return fmt.Sprintf("ImGuiInputEventType(%b)", e)
}

// GuiInputSource wraps the enum ImGuiInputSource.
type GuiInputSource int32

const (
	// GuiInputSource_COUNT wraps ImGuiInputSource_COUNT.
	GuiInputSource_COUNT GuiInputSource = C.ImGuiInputSource_COUNT
	// GuiInputSource_Gamepad wraps ImGuiInputSource_Gamepad.
	GuiInputSource_Gamepad GuiInputSource = C.ImGuiInputSource_Gamepad
	// GuiInputSource_Keyboard wraps ImGuiInputSource_Keyboard.
	GuiInputSource_Keyboard GuiInputSource = C.ImGuiInputSource_Keyboard
	// GuiInputSource_Mouse wraps ImGuiInputSource_Mouse.
	GuiInputSource_Mouse GuiInputSource = C.ImGuiInputSource_Mouse
	// GuiInputSource_None wraps ImGuiInputSource_None.
	GuiInputSource_None GuiInputSource = C.ImGuiInputSource_None
)

func (e GuiInputSource) String() string {
	return fmt.Sprintf("ImGuiInputSource(%b)", e)
}

// GuiKey wraps the enum ImGuiKey.
type GuiKey int32

const (
	// GuiKey_0 wraps ImGuiKey_0.
	GuiKey_0 GuiKey = C.ImGuiKey_0
	// GuiKey_1 wraps ImGuiKey_1.
	GuiKey_1 GuiKey = C.ImGuiKey_1
	// GuiKey_2 wraps ImGuiKey_2.
	GuiKey_2 GuiKey = C.ImGuiKey_2
	// GuiKey_3 wraps ImGuiKey_3.
	GuiKey_3 GuiKey = C.ImGuiKey_3
	// GuiKey_4 wraps ImGuiKey_4.
	GuiKey_4 GuiKey = C.ImGuiKey_4
	// GuiKey_5 wraps ImGuiKey_5.
	GuiKey_5 GuiKey = C.ImGuiKey_5
	// GuiKey_6 wraps ImGuiKey_6.
	GuiKey_6 GuiKey = C.ImGuiKey_6
	// GuiKey_7 wraps ImGuiKey_7.
	GuiKey_7 GuiKey = C.ImGuiKey_7
	// GuiKey_8 wraps ImGuiKey_8.
	GuiKey_8 GuiKey = C.ImGuiKey_8
	// GuiKey_9 wraps ImGuiKey_9.
	GuiKey_9 GuiKey = C.ImGuiKey_9
	// GuiKey_A wraps ImGuiKey_A.
	GuiKey_A GuiKey = C.ImGuiKey_A
	// GuiKey_Apostrophe wraps ImGuiKey_Apostrophe.
	GuiKey_Apostrophe GuiKey = C.ImGuiKey_Apostrophe
	// GuiKey_AppBack wraps ImGuiKey_AppBack.
	GuiKey_AppBack GuiKey = C.ImGuiKey_AppBack
	// GuiKey_AppForward wraps ImGuiKey_AppForward.
	GuiKey_AppForward GuiKey = C.ImGuiKey_AppForward
	// GuiKey_B wraps ImGuiKey_B.
	GuiKey_B GuiKey = C.ImGuiKey_B
	// GuiKey_Backslash wraps ImGuiKey_Backslash.
	GuiKey_Backslash GuiKey = C.ImGuiKey_Backslash
	// GuiKey_Backspace wraps ImGuiKey_Backspace.
	GuiKey_Backspace GuiKey = C.ImGuiKey_Backspace
	// GuiKey_C wraps ImGuiKey_C.
	GuiKey_C GuiKey = C.ImGuiKey_C
	// GuiKey_CapsLock wraps ImGuiKey_CapsLock.
	GuiKey_CapsLock GuiKey = C.ImGuiKey_CapsLock
	// GuiKey_Comma wraps ImGuiKey_Comma.
	GuiKey_Comma GuiKey = C.ImGuiKey_Comma
	// GuiKey_D wraps ImGuiKey_D.
	GuiKey_D GuiKey = C.ImGuiKey_D
	// GuiKey_Delete wraps ImGuiKey_Delete.
	GuiKey_Delete GuiKey = C.ImGuiKey_Delete
	// GuiKey_DownArrow wraps ImGuiKey_DownArrow.
	GuiKey_DownArrow GuiKey = C.ImGuiKey_DownArrow
	// GuiKey_E wraps ImGuiKey_E.
	GuiKey_E GuiKey = C.ImGuiKey_E
	// GuiKey_End wraps ImGuiKey_End.
	GuiKey_End GuiKey = C.ImGuiKey_End
	// GuiKey_Enter wraps ImGuiKey_Enter.
	GuiKey_Enter GuiKey = C.ImGuiKey_Enter
	// GuiKey_Equal wraps ImGuiKey_Equal.
	GuiKey_Equal GuiKey = C.ImGuiKey_Equal
	// GuiKey_Escape wraps ImGuiKey_Escape.
	GuiKey_Escape GuiKey = C.ImGuiKey_Escape
	// GuiKey_F wraps ImGuiKey_F.
	GuiKey_F GuiKey = C.ImGuiKey_F
	// GuiKey_F1 wraps ImGuiKey_F1.
	GuiKey_F1 GuiKey = C.ImGuiKey_F1
	// GuiKey_F10 wraps ImGuiKey_F10.
	GuiKey_F10 GuiKey = C.ImGuiKey_F10
	// GuiKey_F11 wraps ImGuiKey_F11.
	GuiKey_F11 GuiKey = C.ImGuiKey_F11
	// GuiKey_F12 wraps ImGuiKey_F12.
	GuiKey_F12 GuiKey = C.ImGuiKey_F12
	// GuiKey_F13 wraps ImGuiKey_F13.
	GuiKey_F13 GuiKey = C.ImGuiKey_F13
	// GuiKey_F14 wraps ImGuiKey_F14.
	GuiKey_F14 GuiKey = C.ImGuiKey_F14
	// GuiKey_F15 wraps ImGuiKey_F15.
	GuiKey_F15 GuiKey = C.ImGuiKey_F15
	// GuiKey_F16 wraps ImGuiKey_F16.
	GuiKey_F16 GuiKey = C.ImGuiKey_F16
	// GuiKey_F17 wraps ImGuiKey_F17.
	GuiKey_F17 GuiKey = C.ImGuiKey_F17
	// GuiKey_F18 wraps ImGuiKey_F18.
	GuiKey_F18 GuiKey = C.ImGuiKey_F18
	// GuiKey_F19 wraps ImGuiKey_F19.
	GuiKey_F19 GuiKey = C.ImGuiKey_F19
	// GuiKey_F2 wraps ImGuiKey_F2.
	GuiKey_F2 GuiKey = C.ImGuiKey_F2
	// GuiKey_F20 wraps ImGuiKey_F20.
	GuiKey_F20 GuiKey = C.ImGuiKey_F20
	// GuiKey_F21 wraps ImGuiKey_F21.
	GuiKey_F21 GuiKey = C.ImGuiKey_F21
	// GuiKey_F22 wraps ImGuiKey_F22.
	GuiKey_F22 GuiKey = C.ImGuiKey_F22
	// GuiKey_F23 wraps ImGuiKey_F23.
	GuiKey_F23 GuiKey = C.ImGuiKey_F23
	// GuiKey_F24 wraps ImGuiKey_F24.
	GuiKey_F24 GuiKey = C.ImGuiKey_F24
	// GuiKey_F3 wraps ImGuiKey_F3.
	GuiKey_F3 GuiKey = C.ImGuiKey_F3
	// GuiKey_F4 wraps ImGuiKey_F4.
	GuiKey_F4 GuiKey = C.ImGuiKey_F4
	// GuiKey_F5 wraps ImGuiKey_F5.
	GuiKey_F5 GuiKey = C.ImGuiKey_F5
	// GuiKey_F6 wraps ImGuiKey_F6.
	GuiKey_F6 GuiKey = C.ImGuiKey_F6
	// GuiKey_F7 wraps ImGuiKey_F7.
	GuiKey_F7 GuiKey = C.ImGuiKey_F7
	// GuiKey_F8 wraps ImGuiKey_F8.
	GuiKey_F8 GuiKey = C.ImGuiKey_F8
	// GuiKey_F9 wraps ImGuiKey_F9.
	GuiKey_F9 GuiKey = C.ImGuiKey_F9
	// GuiKey_G wraps ImGuiKey_G.
	GuiKey_G GuiKey = C.ImGuiKey_G
	// GuiKey_GamepadBack wraps ImGuiKey_GamepadBack.
	GuiKey_GamepadBack GuiKey = C.ImGuiKey_GamepadBack
	// GuiKey_GamepadDpadDown wraps ImGuiKey_GamepadDpadDown.
	GuiKey_GamepadDpadDown GuiKey = C.ImGuiKey_GamepadDpadDown
	// GuiKey_GamepadDpadLeft wraps ImGuiKey_GamepadDpadLeft.
	GuiKey_GamepadDpadLeft GuiKey = C.ImGuiKey_GamepadDpadLeft
	// GuiKey_GamepadDpadRight wraps ImGuiKey_GamepadDpadRight.
	GuiKey_GamepadDpadRight GuiKey = C.ImGuiKey_GamepadDpadRight
	// GuiKey_GamepadDpadUp wraps ImGuiKey_GamepadDpadUp.
	GuiKey_GamepadDpadUp GuiKey = C.ImGuiKey_GamepadDpadUp
	// GuiKey_GamepadFaceDown wraps ImGuiKey_GamepadFaceDown.
	GuiKey_GamepadFaceDown GuiKey = C.ImGuiKey_GamepadFaceDown
	// GuiKey_GamepadFaceLeft wraps ImGuiKey_GamepadFaceLeft.
	GuiKey_GamepadFaceLeft GuiKey = C.ImGuiKey_GamepadFaceLeft
	// GuiKey_GamepadFaceRight wraps ImGuiKey_GamepadFaceRight.
	GuiKey_GamepadFaceRight GuiKey = C.ImGuiKey_GamepadFaceRight
	// GuiKey_GamepadFaceUp wraps ImGuiKey_GamepadFaceUp.
	GuiKey_GamepadFaceUp GuiKey = C.ImGuiKey_GamepadFaceUp
	// GuiKey_GamepadL1 wraps ImGuiKey_GamepadL1.
	GuiKey_GamepadL1 GuiKey = C.ImGuiKey_GamepadL1
	// GuiKey_GamepadL2 wraps ImGuiKey_GamepadL2.
	GuiKey_GamepadL2 GuiKey = C.ImGuiKey_GamepadL2
	// GuiKey_GamepadL3 wraps ImGuiKey_GamepadL3.
	GuiKey_GamepadL3 GuiKey = C.ImGuiKey_GamepadL3
	// GuiKey_GamepadLStickDown wraps ImGuiKey_GamepadLStickDown.
	GuiKey_GamepadLStickDown GuiKey = C.ImGuiKey_GamepadLStickDown
	// GuiKey_GamepadLStickLeft wraps ImGuiKey_GamepadLStickLeft.
	GuiKey_GamepadLStickLeft GuiKey = C.ImGuiKey_GamepadLStickLeft
	// GuiKey_GamepadLStickRight wraps ImGuiKey_GamepadLStickRight.
	GuiKey_GamepadLStickRight GuiKey = C.ImGuiKey_GamepadLStickRight
	// GuiKey_GamepadLStickUp wraps ImGuiKey_GamepadLStickUp.
	GuiKey_GamepadLStickUp GuiKey = C.ImGuiKey_GamepadLStickUp
	// GuiKey_GamepadR1 wraps ImGuiKey_GamepadR1.
	GuiKey_GamepadR1 GuiKey = C.ImGuiKey_GamepadR1
	// GuiKey_GamepadR2 wraps ImGuiKey_GamepadR2.
	GuiKey_GamepadR2 GuiKey = C.ImGuiKey_GamepadR2
	// GuiKey_GamepadR3 wraps ImGuiKey_GamepadR3.
	GuiKey_GamepadR3 GuiKey = C.ImGuiKey_GamepadR3
	// GuiKey_GamepadRStickDown wraps ImGuiKey_GamepadRStickDown.
	GuiKey_GamepadRStickDown GuiKey = C.ImGuiKey_GamepadRStickDown
	// GuiKey_GamepadRStickLeft wraps ImGuiKey_GamepadRStickLeft.
	GuiKey_GamepadRStickLeft GuiKey = C.ImGuiKey_GamepadRStickLeft
	// GuiKey_GamepadRStickRight wraps ImGuiKey_GamepadRStickRight.
	GuiKey_GamepadRStickRight GuiKey = C.ImGuiKey_GamepadRStickRight
	// GuiKey_GamepadRStickUp wraps ImGuiKey_GamepadRStickUp.
	GuiKey_GamepadRStickUp GuiKey = C.ImGuiKey_GamepadRStickUp
	// GuiKey_GamepadStart wraps ImGuiKey_GamepadStart.
	GuiKey_GamepadStart GuiKey = C.ImGuiKey_GamepadStart
	// GuiKey_GraveAccent wraps ImGuiKey_GraveAccent.
	GuiKey_GraveAccent GuiKey = C.ImGuiKey_GraveAccent
	// GuiKey_H wraps ImGuiKey_H.
	GuiKey_H GuiKey = C.ImGuiKey_H
	// GuiKey_Home wraps ImGuiKey_Home.
	GuiKey_Home GuiKey = C.ImGuiKey_Home
	// GuiKey_I wraps ImGuiKey_I.
	GuiKey_I GuiKey = C.ImGuiKey_I
	// GuiKey_Insert wraps ImGuiKey_Insert.
	GuiKey_Insert GuiKey = C.ImGuiKey_Insert
	// GuiKey_J wraps ImGuiKey_J.
	GuiKey_J GuiKey = C.ImGuiKey_J
	// GuiKey_K wraps ImGuiKey_K.
	GuiKey_K GuiKey = C.ImGuiKey_K
	// GuiKey_Keypad0 wraps ImGuiKey_Keypad0.
	GuiKey_Keypad0 GuiKey = C.ImGuiKey_Keypad0
	// GuiKey_Keypad1 wraps ImGuiKey_Keypad1.
	GuiKey_Keypad1 GuiKey = C.ImGuiKey_Keypad1
	// GuiKey_Keypad2 wraps ImGuiKey_Keypad2.
	GuiKey_Keypad2 GuiKey = C.ImGuiKey_Keypad2
	// GuiKey_Keypad3 wraps ImGuiKey_Keypad3.
	GuiKey_Keypad3 GuiKey = C.ImGuiKey_Keypad3
	// GuiKey_Keypad4 wraps ImGuiKey_Keypad4.
	GuiKey_Keypad4 GuiKey = C.ImGuiKey_Keypad4
	// GuiKey_Keypad5 wraps ImGuiKey_Keypad5.
	GuiKey_Keypad5 GuiKey = C.ImGuiKey_Keypad5
	// GuiKey_Keypad6 wraps ImGuiKey_Keypad6.
	GuiKey_Keypad6 GuiKey = C.ImGuiKey_Keypad6
	// GuiKey_Keypad7 wraps ImGuiKey_Keypad7.
	GuiKey_Keypad7 GuiKey = C.ImGuiKey_Keypad7
	// GuiKey_Keypad8 wraps ImGuiKey_Keypad8.
	GuiKey_Keypad8 GuiKey = C.ImGuiKey_Keypad8
	// GuiKey_Keypad9 wraps ImGuiKey_Keypad9.
	GuiKey_Keypad9 GuiKey = C.ImGuiKey_Keypad9
	// GuiKey_KeypadAdd wraps ImGuiKey_KeypadAdd.
	GuiKey_KeypadAdd GuiKey = C.ImGuiKey_KeypadAdd
	// GuiKey_KeypadDecimal wraps ImGuiKey_KeypadDecimal.
	GuiKey_KeypadDecimal GuiKey = C.ImGuiKey_KeypadDecimal
	// GuiKey_KeypadDivide wraps ImGuiKey_KeypadDivide.
	GuiKey_KeypadDivide GuiKey = C.ImGuiKey_KeypadDivide
	// GuiKey_KeypadEnter wraps ImGuiKey_KeypadEnter.
	GuiKey_KeypadEnter GuiKey = C.ImGuiKey_KeypadEnter
	// GuiKey_KeypadEqual wraps ImGuiKey_KeypadEqual.
	GuiKey_KeypadEqual GuiKey = C.ImGuiKey_KeypadEqual
	// GuiKey_KeypadMultiply wraps ImGuiKey_KeypadMultiply.
	GuiKey_KeypadMultiply GuiKey = C.ImGuiKey_KeypadMultiply
	// GuiKey_KeypadSubtract wraps ImGuiKey_KeypadSubtract.
	GuiKey_KeypadSubtract GuiKey = C.ImGuiKey_KeypadSubtract
	// GuiKey_L wraps ImGuiKey_L.
	GuiKey_L GuiKey = C.ImGuiKey_L
	// GuiKey_LeftAlt wraps ImGuiKey_LeftAlt.
	GuiKey_LeftAlt GuiKey = C.ImGuiKey_LeftAlt
	// GuiKey_LeftArrow wraps ImGuiKey_LeftArrow.
	GuiKey_LeftArrow GuiKey = C.ImGuiKey_LeftArrow
	// GuiKey_LeftBracket wraps ImGuiKey_LeftBracket.
	GuiKey_LeftBracket GuiKey = C.ImGuiKey_LeftBracket
	// GuiKey_LeftCtrl wraps ImGuiKey_LeftCtrl.
	GuiKey_LeftCtrl GuiKey = C.ImGuiKey_LeftCtrl
	// GuiKey_LeftShift wraps ImGuiKey_LeftShift.
	GuiKey_LeftShift GuiKey = C.ImGuiKey_LeftShift
	// GuiKey_LeftSuper wraps ImGuiKey_LeftSuper.
	GuiKey_LeftSuper GuiKey = C.ImGuiKey_LeftSuper
	// GuiKey_M wraps ImGuiKey_M.
	GuiKey_M GuiKey = C.ImGuiKey_M
	// GuiKey_Menu wraps ImGuiKey_Menu.
	GuiKey_Menu GuiKey = C.ImGuiKey_Menu
	// GuiKey_Minus wraps ImGuiKey_Minus.
	GuiKey_Minus GuiKey = C.ImGuiKey_Minus
	// GuiKey_MouseLeft wraps ImGuiKey_MouseLeft.
	GuiKey_MouseLeft GuiKey = C.ImGuiKey_MouseLeft
	// GuiKey_MouseMiddle wraps ImGuiKey_MouseMiddle.
	GuiKey_MouseMiddle GuiKey = C.ImGuiKey_MouseMiddle
	// GuiKey_MouseRight wraps ImGuiKey_MouseRight.
	GuiKey_MouseRight GuiKey = C.ImGuiKey_MouseRight
	// GuiKey_MouseWheelX wraps ImGuiKey_MouseWheelX.
	GuiKey_MouseWheelX GuiKey = C.ImGuiKey_MouseWheelX
	// GuiKey_MouseWheelY wraps ImGuiKey_MouseWheelY.
	GuiKey_MouseWheelY GuiKey = C.ImGuiKey_MouseWheelY
	// GuiKey_MouseX1 wraps ImGuiKey_MouseX1.
	GuiKey_MouseX1 GuiKey = C.ImGuiKey_MouseX1
	// GuiKey_MouseX2 wraps ImGuiKey_MouseX2.
	GuiKey_MouseX2 GuiKey = C.ImGuiKey_MouseX2
	// GuiKey_N wraps ImGuiKey_N.
	GuiKey_N GuiKey = C.ImGuiKey_N
	// GuiKey_NamedKey_BEGIN wraps ImGuiKey_NamedKey_BEGIN.
	GuiKey_NamedKey_BEGIN GuiKey = C.ImGuiKey_NamedKey_BEGIN
	// GuiKey_NamedKey_COUNT wraps ImGuiKey_NamedKey_COUNT.
	GuiKey_NamedKey_COUNT GuiKey = C.ImGuiKey_NamedKey_COUNT
	// GuiKey_NamedKey_END wraps ImGuiKey_NamedKey_END.
	GuiKey_NamedKey_END GuiKey = C.ImGuiKey_NamedKey_END
	// GuiKey_None wraps ImGuiKey_None.
	GuiKey_None GuiKey = C.ImGuiKey_None
	// GuiKey_NumLock wraps ImGuiKey_NumLock.
	GuiKey_NumLock GuiKey = C.ImGuiKey_NumLock
	// GuiKey_O wraps ImGuiKey_O.
	GuiKey_O GuiKey = C.ImGuiKey_O
	// GuiKey_Oem102 wraps ImGuiKey_Oem102.
	GuiKey_Oem102 GuiKey = C.ImGuiKey_Oem102
	// GuiKey_P wraps ImGuiKey_P.
	GuiKey_P GuiKey = C.ImGuiKey_P
	// GuiKey_PageDown wraps ImGuiKey_PageDown.
	GuiKey_PageDown GuiKey = C.ImGuiKey_PageDown
	// GuiKey_PageUp wraps ImGuiKey_PageUp.
	GuiKey_PageUp GuiKey = C.ImGuiKey_PageUp
	// GuiKey_Pause wraps ImGuiKey_Pause.
	GuiKey_Pause GuiKey = C.ImGuiKey_Pause
	// GuiKey_Period wraps ImGuiKey_Period.
	GuiKey_Period GuiKey = C.ImGuiKey_Period
	// GuiKey_PrintScreen wraps ImGuiKey_PrintScreen.
	GuiKey_PrintScreen GuiKey = C.ImGuiKey_PrintScreen
	// GuiKey_Q wraps ImGuiKey_Q.
	GuiKey_Q GuiKey = C.ImGuiKey_Q
	// GuiKey_R wraps ImGuiKey_R.
	GuiKey_R GuiKey = C.ImGuiKey_R
	// GuiKey_ReservedForModAlt wraps ImGuiKey_ReservedForModAlt.
	GuiKey_ReservedForModAlt GuiKey = C.ImGuiKey_ReservedForModAlt
	// GuiKey_ReservedForModCtrl wraps ImGuiKey_ReservedForModCtrl.
	GuiKey_ReservedForModCtrl GuiKey = C.ImGuiKey_ReservedForModCtrl
	// GuiKey_ReservedForModShift wraps ImGuiKey_ReservedForModShift.
	GuiKey_ReservedForModShift GuiKey = C.ImGuiKey_ReservedForModShift
	// GuiKey_ReservedForModSuper wraps ImGuiKey_ReservedForModSuper.
	GuiKey_ReservedForModSuper GuiKey = C.ImGuiKey_ReservedForModSuper
	// GuiKey_RightAlt wraps ImGuiKey_RightAlt.
	GuiKey_RightAlt GuiKey = C.ImGuiKey_RightAlt
	// GuiKey_RightArrow wraps ImGuiKey_RightArrow.
	GuiKey_RightArrow GuiKey = C.ImGuiKey_RightArrow
	// GuiKey_RightBracket wraps ImGuiKey_RightBracket.
	GuiKey_RightBracket GuiKey = C.ImGuiKey_RightBracket
	// GuiKey_RightCtrl wraps ImGuiKey_RightCtrl.
	GuiKey_RightCtrl GuiKey = C.ImGuiKey_RightCtrl
	// GuiKey_RightShift wraps ImGuiKey_RightShift.
	GuiKey_RightShift GuiKey = C.ImGuiKey_RightShift
	// GuiKey_RightSuper wraps ImGuiKey_RightSuper.
	GuiKey_RightSuper GuiKey = C.ImGuiKey_RightSuper
	// GuiKey_S wraps ImGuiKey_S.
	GuiKey_S GuiKey = C.ImGuiKey_S
	// GuiKey_ScrollLock wraps ImGuiKey_ScrollLock.
	GuiKey_ScrollLock GuiKey = C.ImGuiKey_ScrollLock
	// GuiKey_Semicolon wraps ImGuiKey_Semicolon.
	GuiKey_Semicolon GuiKey = C.ImGuiKey_Semicolon
	// GuiKey_Slash wraps ImGuiKey_Slash.
	GuiKey_Slash GuiKey = C.ImGuiKey_Slash
	// GuiKey_Space wraps ImGuiKey_Space.
	GuiKey_Space GuiKey = C.ImGuiKey_Space
	// GuiKey_T wraps ImGuiKey_T.
	GuiKey_T GuiKey = C.ImGuiKey_T
	// GuiKey_Tab wraps ImGuiKey_Tab.
	GuiKey_Tab GuiKey = C.ImGuiKey_Tab
	// GuiKey_U wraps ImGuiKey_U.
	GuiKey_U GuiKey = C.ImGuiKey_U
	// GuiKey_UpArrow wraps ImGuiKey_UpArrow.
	GuiKey_UpArrow GuiKey = C.ImGuiKey_UpArrow
	// GuiKey_V wraps ImGuiKey_V.
	GuiKey_V GuiKey = C.ImGuiKey_V
	// GuiKey_W wraps ImGuiKey_W.
	GuiKey_W GuiKey = C.ImGuiKey_W
	// GuiKey_X wraps ImGuiKey_X.
	GuiKey_X GuiKey = C.ImGuiKey_X
	// GuiKey_Y wraps ImGuiKey_Y.
	GuiKey_Y GuiKey = C.ImGuiKey_Y
	// GuiKey_Z wraps ImGuiKey_Z.
	GuiKey_Z GuiKey = C.ImGuiKey_Z
	// GuiMod_Alt wraps ImGuiMod_Alt.
	GuiMod_Alt GuiKey = C.ImGuiMod_Alt
	// GuiMod_Ctrl wraps ImGuiMod_Ctrl.
	GuiMod_Ctrl GuiKey = C.ImGuiMod_Ctrl
	// GuiMod_Mask wraps ImGuiMod_Mask_.
	GuiMod_Mask GuiKey = C.ImGuiMod_Mask_
	// GuiMod_None wraps ImGuiMod_None.
	GuiMod_None GuiKey = C.ImGuiMod_None
	// GuiMod_Shift wraps ImGuiMod_Shift.
	GuiMod_Shift GuiKey = C.ImGuiMod_Shift
	// GuiMod_Super wraps ImGuiMod_Super.
	GuiMod_Super GuiKey = C.ImGuiMod_Super
)

func (e GuiKey) String() string {
	return fmt.Sprintf("ImGuiKey(%b)", e)
}

// GuiLayoutType wraps the enum ImGuiLayoutType.
type GuiLayoutType int32

const (
	// GuiLayoutType_Horizontal wraps ImGuiLayoutType_Horizontal.
	GuiLayoutType_Horizontal GuiLayoutType = C.ImGuiLayoutType_Horizontal
	// GuiLayoutType_Vertical wraps ImGuiLayoutType_Vertical.
	GuiLayoutType_Vertical GuiLayoutType = C.ImGuiLayoutType_Vertical
)

func (e GuiLayoutType) String() string {
	return fmt.Sprintf("ImGuiLayoutType(%b)", e)
}

// GuiLocKey wraps the enum ImGuiLocKey.
type GuiLocKey int32

const (
	// GuiLocKey_COUNT wraps ImGuiLocKey_COUNT.
	GuiLocKey_COUNT GuiLocKey = C.ImGuiLocKey_COUNT
	// GuiLocKey_CopyLink wraps ImGuiLocKey_CopyLink.
	GuiLocKey_CopyLink GuiLocKey = C.ImGuiLocKey_CopyLink
	// GuiLocKey_DockingDragToUndockOrMoveNode wraps ImGuiLocKey_DockingDragToUndockOrMoveNode.
	GuiLocKey_DockingDragToUndockOrMoveNode GuiLocKey = C.ImGuiLocKey_DockingDragToUndockOrMoveNode
	// GuiLocKey_DockingHideTabBar wraps ImGuiLocKey_DockingHideTabBar.
	GuiLocKey_DockingHideTabBar GuiLocKey = C.ImGuiLocKey_DockingHideTabBar
	// GuiLocKey_DockingHoldShiftToDock wraps ImGuiLocKey_DockingHoldShiftToDock.
	GuiLocKey_DockingHoldShiftToDock GuiLocKey = C.ImGuiLocKey_DockingHoldShiftToDock
	// GuiLocKey_OpenLink_s wraps ImGuiLocKey_OpenLink_s.
	GuiLocKey_OpenLink_s GuiLocKey = C.ImGuiLocKey_OpenLink_s
	// GuiLocKey_TableResetOrder wraps ImGuiLocKey_TableResetOrder.
	GuiLocKey_TableResetOrder GuiLocKey = C.ImGuiLocKey_TableResetOrder
	// GuiLocKey_TableSizeAllDefault wraps ImGuiLocKey_TableSizeAllDefault.
	GuiLocKey_TableSizeAllDefault GuiLocKey = C.ImGuiLocKey_TableSizeAllDefault
	// GuiLocKey_TableSizeAllFit wraps ImGuiLocKey_TableSizeAllFit.
	GuiLocKey_TableSizeAllFit GuiLocKey = C.ImGuiLocKey_TableSizeAllFit
	// GuiLocKey_TableSizeOne wraps ImGuiLocKey_TableSizeOne.
	GuiLocKey_TableSizeOne GuiLocKey = C.ImGuiLocKey_TableSizeOne
	// GuiLocKey_VersionStr wraps ImGuiLocKey_VersionStr.
	GuiLocKey_VersionStr GuiLocKey = C.ImGuiLocKey_VersionStr
	// GuiLocKey_WindowingMainMenuBar wraps ImGuiLocKey_WindowingMainMenuBar.
	GuiLocKey_WindowingMainMenuBar GuiLocKey = C.ImGuiLocKey_WindowingMainMenuBar
	// GuiLocKey_WindowingPopup wraps ImGuiLocKey_WindowingPopup.
	GuiLocKey_WindowingPopup GuiLocKey = C.ImGuiLocKey_WindowingPopup
	// GuiLocKey_WindowingUntitled wraps ImGuiLocKey_WindowingUntitled.
	GuiLocKey_WindowingUntitled GuiLocKey = C.ImGuiLocKey_WindowingUntitled
)

func (e GuiLocKey) String() string {
	return fmt.Sprintf("ImGuiLocKey(%b)", e)
}

// GuiMouseButton wraps the enum ImGuiMouseButton.
type GuiMouseButton int32

const (
	// GuiMouseButton_COUNT wraps ImGuiMouseButton_COUNT.
	GuiMouseButton_COUNT GuiMouseButton = C.ImGuiMouseButton_COUNT
	// GuiMouseButton_Left wraps ImGuiMouseButton_Left.
	GuiMouseButton_Left GuiMouseButton = C.ImGuiMouseButton_Left
	// GuiMouseButton_Middle wraps ImGuiMouseButton_Middle.
	GuiMouseButton_Middle GuiMouseButton = C.ImGuiMouseButton_Middle
	// GuiMouseButton_Right wraps ImGuiMouseButton_Right.
	GuiMouseButton_Right GuiMouseButton = C.ImGuiMouseButton_Right
)

func (e GuiMouseButton) String() string {
	return fmt.Sprintf("ImGuiMouseButton(%b)", e)
}

// GuiMouseCursor wraps the enum ImGuiMouseCursor.
type GuiMouseCursor int32

const (
	// GuiMouseCursor_Arrow wraps ImGuiMouseCursor_Arrow.
	GuiMouseCursor_Arrow GuiMouseCursor = C.ImGuiMouseCursor_Arrow
	// GuiMouseCursor_COUNT wraps ImGuiMouseCursor_COUNT.
	GuiMouseCursor_COUNT GuiMouseCursor = C.ImGuiMouseCursor_COUNT
	// GuiMouseCursor_Hand wraps ImGuiMouseCursor_Hand.
	GuiMouseCursor_Hand GuiMouseCursor = C.ImGuiMouseCursor_Hand
	// GuiMouseCursor_None wraps ImGuiMouseCursor_None.
	GuiMouseCursor_None GuiMouseCursor = C.ImGuiMouseCursor_None
	// GuiMouseCursor_NotAllowed wraps ImGuiMouseCursor_NotAllowed.
	GuiMouseCursor_NotAllowed GuiMouseCursor = C.ImGuiMouseCursor_NotAllowed
	// GuiMouseCursor_Progress wraps ImGuiMouseCursor_Progress.
	GuiMouseCursor_Progress GuiMouseCursor = C.ImGuiMouseCursor_Progress
	// GuiMouseCursor_ResizeAll wraps ImGuiMouseCursor_ResizeAll.
	GuiMouseCursor_ResizeAll GuiMouseCursor = C.ImGuiMouseCursor_ResizeAll
	// GuiMouseCursor_ResizeEW wraps ImGuiMouseCursor_ResizeEW.
	GuiMouseCursor_ResizeEW GuiMouseCursor = C.ImGuiMouseCursor_ResizeEW
	// GuiMouseCursor_ResizeNESW wraps ImGuiMouseCursor_ResizeNESW.
	GuiMouseCursor_ResizeNESW GuiMouseCursor = C.ImGuiMouseCursor_ResizeNESW
	// GuiMouseCursor_ResizeNS wraps ImGuiMouseCursor_ResizeNS.
	GuiMouseCursor_ResizeNS GuiMouseCursor = C.ImGuiMouseCursor_ResizeNS
	// GuiMouseCursor_ResizeNWSE wraps ImGuiMouseCursor_ResizeNWSE.
	GuiMouseCursor_ResizeNWSE GuiMouseCursor = C.ImGuiMouseCursor_ResizeNWSE
	// GuiMouseCursor_TextInput wraps ImGuiMouseCursor_TextInput.
	GuiMouseCursor_TextInput GuiMouseCursor = C.ImGuiMouseCursor_TextInput
	// GuiMouseCursor_Wait wraps ImGuiMouseCursor_Wait.
	GuiMouseCursor_Wait GuiMouseCursor = C.ImGuiMouseCursor_Wait
)

func (e GuiMouseCursor) String() string {
	return fmt.Sprintf("ImGuiMouseCursor(%b)", e)
}

// GuiMouseSource wraps the enum ImGuiMouseSource.
type GuiMouseSource int32

const (
	// GuiMouseSource_COUNT wraps ImGuiMouseSource_COUNT.
	GuiMouseSource_COUNT GuiMouseSource = C.ImGuiMouseSource_COUNT
	// GuiMouseSource_Mouse wraps ImGuiMouseSource_Mouse.
	GuiMouseSource_Mouse GuiMouseSource = C.ImGuiMouseSource_Mouse
	// GuiMouseSource_Pen wraps ImGuiMouseSource_Pen.
	GuiMouseSource_Pen GuiMouseSource = C.ImGuiMouseSource_Pen
	// GuiMouseSource_TouchScreen wraps ImGuiMouseSource_TouchScreen.
	GuiMouseSource_TouchScreen GuiMouseSource = C.ImGuiMouseSource_TouchScreen
)

func (e GuiMouseSource) String() string {
	return fmt.Sprintf("ImGuiMouseSource(%b)", e)
}

// GuiNavLayer wraps the enum ImGuiNavLayer.
type GuiNavLayer int32

const (
	// GuiNavLayer_COUNT wraps ImGuiNavLayer_COUNT.
	GuiNavLayer_COUNT GuiNavLayer = C.ImGuiNavLayer_COUNT
	// GuiNavLayer_Main wraps ImGuiNavLayer_Main.
	GuiNavLayer_Main GuiNavLayer = C.ImGuiNavLayer_Main
	// GuiNavLayer_Menu wraps ImGuiNavLayer_Menu.
	GuiNavLayer_Menu GuiNavLayer = C.ImGuiNavLayer_Menu
)

func (e GuiNavLayer) String() string {
	return fmt.Sprintf("ImGuiNavLayer(%b)", e)
}

// GuiPlotType wraps the enum ImGuiPlotType.
type GuiPlotType int32

const (
	// GuiPlotType_Histogram wraps ImGuiPlotType_Histogram.
	GuiPlotType_Histogram GuiPlotType = C.ImGuiPlotType_Histogram
	// GuiPlotType_Lines wraps ImGuiPlotType_Lines.
	GuiPlotType_Lines GuiPlotType = C.ImGuiPlotType_Lines
)

func (e GuiPlotType) String() string {
	return fmt.Sprintf("ImGuiPlotType(%b)", e)
}

// GuiPopupPositionPolicy wraps the enum ImGuiPopupPositionPolicy.
type GuiPopupPositionPolicy int32

const (
	// GuiPopupPositionPolicy_ComboBox wraps ImGuiPopupPositionPolicy_ComboBox.
	GuiPopupPositionPolicy_ComboBox GuiPopupPositionPolicy = C.ImGuiPopupPositionPolicy_ComboBox
	// GuiPopupPositionPolicy_Default wraps ImGuiPopupPositionPolicy_Default.
	GuiPopupPositionPolicy_Default GuiPopupPositionPolicy = C.ImGuiPopupPositionPolicy_Default
	// GuiPopupPositionPolicy_Tooltip wraps ImGuiPopupPositionPolicy_Tooltip.
	GuiPopupPositionPolicy_Tooltip GuiPopupPositionPolicy = C.ImGuiPopupPositionPolicy_Tooltip
)

func (e GuiPopupPositionPolicy) String() string {
	return fmt.Sprintf("ImGuiPopupPositionPolicy(%b)", e)
}

// GuiSelectionRequestType wraps the enum ImGuiSelectionRequestType.
type GuiSelectionRequestType int32

const (
	// GuiSelectionRequestType_None wraps ImGuiSelectionRequestType_None.
	GuiSelectionRequestType_None GuiSelectionRequestType = C.ImGuiSelectionRequestType_None
	// GuiSelectionRequestType_SetAll wraps ImGuiSelectionRequestType_SetAll.
	GuiSelectionRequestType_SetAll GuiSelectionRequestType = C.ImGuiSelectionRequestType_SetAll
	// GuiSelectionRequestType_SetRange wraps ImGuiSelectionRequestType_SetRange.
	GuiSelectionRequestType_SetRange GuiSelectionRequestType = C.ImGuiSelectionRequestType_SetRange
)

func (e GuiSelectionRequestType) String() string {
	return fmt.Sprintf("ImGuiSelectionRequestType(%b)", e)
}

// GuiSortDirection wraps the enum ImGuiSortDirection.
type GuiSortDirection int32

const (
	// GuiSortDirection_Ascending wraps ImGuiSortDirection_Ascending.
	GuiSortDirection_Ascending GuiSortDirection = C.ImGuiSortDirection_Ascending
	// GuiSortDirection_Descending wraps ImGuiSortDirection_Descending.
	GuiSortDirection_Descending GuiSortDirection = C.ImGuiSortDirection_Descending
	// GuiSortDirection_None wraps ImGuiSortDirection_None.
	GuiSortDirection_None GuiSortDirection = C.ImGuiSortDirection_None
)

func (e GuiSortDirection) String() string {
	return fmt.Sprintf("ImGuiSortDirection(%b)", e)
}

// GuiStyleVar wraps the enum ImGuiStyleVar.
type GuiStyleVar int32

const (
	// GuiStyleVar_Alpha wraps ImGuiStyleVar_Alpha.
	GuiStyleVar_Alpha GuiStyleVar = C.ImGuiStyleVar_Alpha
	// GuiStyleVar_ButtonTextAlign wraps ImGuiStyleVar_ButtonTextAlign.
	GuiStyleVar_ButtonTextAlign GuiStyleVar = C.ImGuiStyleVar_ButtonTextAlign
	// GuiStyleVar_COUNT wraps ImGuiStyleVar_COUNT.
	GuiStyleVar_COUNT GuiStyleVar = C.ImGuiStyleVar_COUNT
	// GuiStyleVar_CellPadding wraps ImGuiStyleVar_CellPadding.
	GuiStyleVar_CellPadding GuiStyleVar = C.ImGuiStyleVar_CellPadding
	// GuiStyleVar_ChildBorderSize wraps ImGuiStyleVar_ChildBorderSize.
	GuiStyleVar_ChildBorderSize GuiStyleVar = C.ImGuiStyleVar_ChildBorderSize
	// GuiStyleVar_ChildRounding wraps ImGuiStyleVar_ChildRounding.
	GuiStyleVar_ChildRounding GuiStyleVar = C.ImGuiStyleVar_ChildRounding
	// GuiStyleVar_DisabledAlpha wraps ImGuiStyleVar_DisabledAlpha.
	GuiStyleVar_DisabledAlpha GuiStyleVar = C.ImGuiStyleVar_DisabledAlpha
	// GuiStyleVar_DockingSeparatorSize wraps ImGuiStyleVar_DockingSeparatorSize.
	GuiStyleVar_DockingSeparatorSize GuiStyleVar = C.ImGuiStyleVar_DockingSeparatorSize
	// GuiStyleVar_FrameBorderSize wraps ImGuiStyleVar_FrameBorderSize.
	GuiStyleVar_FrameBorderSize GuiStyleVar = C.ImGuiStyleVar_FrameBorderSize
	// GuiStyleVar_FramePadding wraps ImGuiStyleVar_FramePadding.
	GuiStyleVar_FramePadding GuiStyleVar = C.ImGuiStyleVar_FramePadding
	// GuiStyleVar_FrameRounding wraps ImGuiStyleVar_FrameRounding.
	GuiStyleVar_FrameRounding GuiStyleVar = C.ImGuiStyleVar_FrameRounding
	// GuiStyleVar_GrabMinSize wraps ImGuiStyleVar_GrabMinSize.
	GuiStyleVar_GrabMinSize GuiStyleVar = C.ImGuiStyleVar_GrabMinSize
	// GuiStyleVar_GrabRounding wraps ImGuiStyleVar_GrabRounding.
	GuiStyleVar_GrabRounding GuiStyleVar = C.ImGuiStyleVar_GrabRounding
	// GuiStyleVar_ImageBorderSize wraps ImGuiStyleVar_ImageBorderSize.
	GuiStyleVar_ImageBorderSize GuiStyleVar = C.ImGuiStyleVar_ImageBorderSize
	// GuiStyleVar_IndentSpacing wraps ImGuiStyleVar_IndentSpacing.
	GuiStyleVar_IndentSpacing GuiStyleVar = C.ImGuiStyleVar_IndentSpacing
	// GuiStyleVar_ItemInnerSpacing wraps ImGuiStyleVar_ItemInnerSpacing.
	GuiStyleVar_ItemInnerSpacing GuiStyleVar = C.ImGuiStyleVar_ItemInnerSpacing
	// GuiStyleVar_ItemSpacing wraps ImGuiStyleVar_ItemSpacing.
	GuiStyleVar_ItemSpacing GuiStyleVar = C.ImGuiStyleVar_ItemSpacing
	// GuiStyleVar_PopupBorderSize wraps ImGuiStyleVar_PopupBorderSize.
	GuiStyleVar_PopupBorderSize GuiStyleVar = C.ImGuiStyleVar_PopupBorderSize
	// GuiStyleVar_PopupRounding wraps ImGuiStyleVar_PopupRounding.
	GuiStyleVar_PopupRounding GuiStyleVar = C.ImGuiStyleVar_PopupRounding
	// GuiStyleVar_ScrollbarPadding wraps ImGuiStyleVar_ScrollbarPadding.
	GuiStyleVar_ScrollbarPadding GuiStyleVar = C.ImGuiStyleVar_ScrollbarPadding
	// GuiStyleVar_ScrollbarRounding wraps ImGuiStyleVar_ScrollbarRounding.
	GuiStyleVar_ScrollbarRounding GuiStyleVar = C.ImGuiStyleVar_ScrollbarRounding
	// GuiStyleVar_ScrollbarSize wraps ImGuiStyleVar_ScrollbarSize.
	GuiStyleVar_ScrollbarSize GuiStyleVar = C.ImGuiStyleVar_ScrollbarSize
	// GuiStyleVar_SelectableTextAlign wraps ImGuiStyleVar_SelectableTextAlign.
	GuiStyleVar_SelectableTextAlign GuiStyleVar = C.ImGuiStyleVar_SelectableTextAlign
	// GuiStyleVar_SeparatorTextAlign wraps ImGuiStyleVar_SeparatorTextAlign.
	GuiStyleVar_SeparatorTextAlign GuiStyleVar = C.ImGuiStyleVar_SeparatorTextAlign
	// GuiStyleVar_SeparatorTextBorderSize wraps ImGuiStyleVar_SeparatorTextBorderSize.
	GuiStyleVar_SeparatorTextBorderSize GuiStyleVar = C.ImGuiStyleVar_SeparatorTextBorderSize
	// GuiStyleVar_SeparatorTextPadding wraps ImGuiStyleVar_SeparatorTextPadding.
	GuiStyleVar_SeparatorTextPadding GuiStyleVar = C.ImGuiStyleVar_SeparatorTextPadding
	// GuiStyleVar_TabBarBorderSize wraps ImGuiStyleVar_TabBarBorderSize.
	GuiStyleVar_TabBarBorderSize GuiStyleVar = C.ImGuiStyleVar_TabBarBorderSize
	// GuiStyleVar_TabBarOverlineSize wraps ImGuiStyleVar_TabBarOverlineSize.
	GuiStyleVar_TabBarOverlineSize GuiStyleVar = C.ImGuiStyleVar_TabBarOverlineSize
	// GuiStyleVar_TabBorderSize wraps ImGuiStyleVar_TabBorderSize.
	GuiStyleVar_TabBorderSize GuiStyleVar = C.ImGuiStyleVar_TabBorderSize
	// GuiStyleVar_TabMinWidthBase wraps ImGuiStyleVar_TabMinWidthBase.
	GuiStyleVar_TabMinWidthBase GuiStyleVar = C.ImGuiStyleVar_TabMinWidthBase
	// GuiStyleVar_TabMinWidthShrink wraps ImGuiStyleVar_TabMinWidthShrink.
	GuiStyleVar_TabMinWidthShrink GuiStyleVar = C.ImGuiStyleVar_TabMinWidthShrink
	// GuiStyleVar_TabRounding wraps ImGuiStyleVar_TabRounding.
	GuiStyleVar_TabRounding GuiStyleVar = C.ImGuiStyleVar_TabRounding
	// GuiStyleVar_TableAngledHeadersAngle wraps ImGuiStyleVar_TableAngledHeadersAngle.
	GuiStyleVar_TableAngledHeadersAngle GuiStyleVar = C.ImGuiStyleVar_TableAngledHeadersAngle
	// GuiStyleVar_TableAngledHeadersTextAlign wraps ImGuiStyleVar_TableAngledHeadersTextAlign.
	GuiStyleVar_TableAngledHeadersTextAlign GuiStyleVar = C.ImGuiStyleVar_TableAngledHeadersTextAlign
	// GuiStyleVar_TreeLinesRounding wraps ImGuiStyleVar_TreeLinesRounding.
	GuiStyleVar_TreeLinesRounding GuiStyleVar = C.ImGuiStyleVar_TreeLinesRounding
	// GuiStyleVar_TreeLinesSize wraps ImGuiStyleVar_TreeLinesSize.
	GuiStyleVar_TreeLinesSize GuiStyleVar = C.ImGuiStyleVar_TreeLinesSize
	// GuiStyleVar_WindowBorderSize wraps ImGuiStyleVar_WindowBorderSize.
	GuiStyleVar_WindowBorderSize GuiStyleVar = C.ImGuiStyleVar_WindowBorderSize
	// GuiStyleVar_WindowMinSize wraps ImGuiStyleVar_WindowMinSize.
	GuiStyleVar_WindowMinSize GuiStyleVar = C.ImGuiStyleVar_WindowMinSize
	// GuiStyleVar_WindowPadding wraps ImGuiStyleVar_WindowPadding.
	GuiStyleVar_WindowPadding GuiStyleVar = C.ImGuiStyleVar_WindowPadding
	// GuiStyleVar_WindowRounding wraps ImGuiStyleVar_WindowRounding.
	GuiStyleVar_WindowRounding GuiStyleVar = C.ImGuiStyleVar_WindowRounding
	// GuiStyleVar_WindowTitleAlign wraps ImGuiStyleVar_WindowTitleAlign.
	GuiStyleVar_WindowTitleAlign GuiStyleVar = C.ImGuiStyleVar_WindowTitleAlign
)

func (e GuiStyleVar) String() string {
	return fmt.Sprintf("ImGuiStyleVar(%b)", e)
}

// GuiTableBgTarget wraps the enum ImGuiTableBgTarget.
type GuiTableBgTarget int32

const (
	// GuiTableBgTarget_CellBg wraps ImGuiTableBgTarget_CellBg.
	GuiTableBgTarget_CellBg GuiTableBgTarget = C.ImGuiTableBgTarget_CellBg
	// GuiTableBgTarget_None wraps ImGuiTableBgTarget_None.
	GuiTableBgTarget_None GuiTableBgTarget = C.ImGuiTableBgTarget_None
	// GuiTableBgTarget_RowBg0 wraps ImGuiTableBgTarget_RowBg0.
	GuiTableBgTarget_RowBg0 GuiTableBgTarget = C.ImGuiTableBgTarget_RowBg0
	// GuiTableBgTarget_RowBg1 wraps ImGuiTableBgTarget_RowBg1.
	GuiTableBgTarget_RowBg1 GuiTableBgTarget = C.ImGuiTableBgTarget_RowBg1
)

func (e GuiTableBgTarget) String() string {
	return fmt.Sprintf("ImGuiTableBgTarget(%b)", e)
}

// GuiWindowDockStyleCol wraps the enum ImGuiWindowDockStyleCol.
type GuiWindowDockStyleCol int32

const (
	// GuiWindowDockStyleCol_COUNT wraps ImGuiWindowDockStyleCol_COUNT.
	GuiWindowDockStyleCol_COUNT GuiWindowDockStyleCol = C.ImGuiWindowDockStyleCol_COUNT
	// GuiWindowDockStyleCol_TabDimmed wraps ImGuiWindowDockStyleCol_TabDimmed.
	GuiWindowDockStyleCol_TabDimmed GuiWindowDockStyleCol = C.ImGuiWindowDockStyleCol_TabDimmed
	// GuiWindowDockStyleCol_TabDimmedSelected wraps ImGuiWindowDockStyleCol_TabDimmedSelected.
	GuiWindowDockStyleCol_TabDimmedSelected GuiWindowDockStyleCol = C.ImGuiWindowDockStyleCol_TabDimmedSelected
	// GuiWindowDockStyleCol_TabDimmedSelectedOverline wraps ImGuiWindowDockStyleCol_TabDimmedSelectedOverline.
	GuiWindowDockStyleCol_TabDimmedSelectedOverline GuiWindowDockStyleCol = C.ImGuiWindowDockStyleCol_TabDimmedSelectedOverline
	// GuiWindowDockStyleCol_TabFocused wraps ImGuiWindowDockStyleCol_TabFocused.
	GuiWindowDockStyleCol_TabFocused GuiWindowDockStyleCol = C.ImGuiWindowDockStyleCol_TabFocused
	// GuiWindowDockStyleCol_TabHovered wraps ImGuiWindowDockStyleCol_TabHovered.
	GuiWindowDockStyleCol_TabHovered GuiWindowDockStyleCol = C.ImGuiWindowDockStyleCol_TabHovered
	// GuiWindowDockStyleCol_TabSelected wraps ImGuiWindowDockStyleCol_TabSelected.
	GuiWindowDockStyleCol_TabSelected GuiWindowDockStyleCol = C.ImGuiWindowDockStyleCol_TabSelected
	// GuiWindowDockStyleCol_TabSelectedOverline wraps ImGuiWindowDockStyleCol_TabSelectedOverline.
	GuiWindowDockStyleCol_TabSelectedOverline GuiWindowDockStyleCol = C.ImGuiWindowDockStyleCol_TabSelectedOverline
	// GuiWindowDockStyleCol_Text wraps ImGuiWindowDockStyleCol_Text.
	GuiWindowDockStyleCol_Text GuiWindowDockStyleCol = C.ImGuiWindowDockStyleCol_Text
)

func (e GuiWindowDockStyleCol) String() string {
	return fmt.Sprintf("ImGuiWindowDockStyleCol(%b)", e)
}

// TextureFormat wraps the enum ImTextureFormat.
type TextureFormat int32

const (
	// TextureFormat_Alpha8 wraps ImTextureFormat_Alpha8.
	TextureFormat_Alpha8 TextureFormat = C.ImTextureFormat_Alpha8
	// TextureFormat_RGBA32 wraps ImTextureFormat_RGBA32.
	TextureFormat_RGBA32 TextureFormat = C.ImTextureFormat_RGBA32
)

func (e TextureFormat) String() string {
	return fmt.Sprintf("ImTextureFormat(%b)", e)
}

// TextureStatus wraps the enum ImTextureStatus.
type TextureStatus int32

const (
	// TextureStatus_Destroyed wraps ImTextureStatus_Destroyed.
	TextureStatus_Destroyed TextureStatus = C.ImTextureStatus_Destroyed
	// TextureStatus_OK wraps ImTextureStatus_OK.
	TextureStatus_OK TextureStatus = C.ImTextureStatus_OK
	// TextureStatus_WantCreate wraps ImTextureStatus_WantCreate.
	TextureStatus_WantCreate TextureStatus = C.ImTextureStatus_WantCreate
	// TextureStatus_WantDestroy wraps ImTextureStatus_WantDestroy.
	TextureStatus_WantDestroy TextureStatus = C.ImTextureStatus_WantDestroy
	// TextureStatus_WantUpdates wraps ImTextureStatus_WantUpdates.
	TextureStatus_WantUpdates TextureStatus = C.ImTextureStatus_WantUpdates
)

func (e TextureStatus) String() string {
	return fmt.Sprintf("ImTextureStatus(%b)", e)
}
