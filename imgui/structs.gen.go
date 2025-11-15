package imgui

import (
	ffi "github.com/csnewman/go-gfx/ffi"
	"unsafe"
)

// #include "imgui_wrapper.h"
// static const int offsetof_ImBitVector_Storage = offsetof(ImBitVector, Storage);
// static const int offsetof_ImChunkStream_ImGuiTableSettings_Buf = offsetof(ImChunkStream_ImGuiTableSettings, Buf);
// static const int offsetof_ImChunkStream_ImGuiWindowSettings_Buf = offsetof(ImChunkStream_ImGuiWindowSettings, Buf);
// static const int offsetof_ImColor_Value = offsetof(ImColor, Value);
// static const int offsetof_ImDrawChannel__CmdBuffer = offsetof(ImDrawChannel, _CmdBuffer);
// static const int offsetof_ImDrawChannel__IdxBuffer = offsetof(ImDrawChannel, _IdxBuffer);
// static const int offsetof_ImDrawCmd_ClipRect = offsetof(ImDrawCmd, ClipRect);
// static const int offsetof_ImDrawCmd_TexRef = offsetof(ImDrawCmd, TexRef);
// static const int offsetof_ImDrawCmdHeader_ClipRect = offsetof(ImDrawCmdHeader, ClipRect);
// static const int offsetof_ImDrawCmdHeader_TexRef = offsetof(ImDrawCmdHeader, TexRef);
// static const int offsetof_ImDrawData_CmdLists = offsetof(ImDrawData, CmdLists);
// static const int offsetof_ImDrawData_DisplayPos = offsetof(ImDrawData, DisplayPos);
// static const int offsetof_ImDrawData_DisplaySize = offsetof(ImDrawData, DisplaySize);
// static const int offsetof_ImDrawData_FramebufferScale = offsetof(ImDrawData, FramebufferScale);
// static const int offsetof_ImDrawDataBuilder_LayerData1 = offsetof(ImDrawDataBuilder, LayerData1);
// static const int offsetof_ImDrawList_CmdBuffer = offsetof(ImDrawList, CmdBuffer);
// static const int offsetof_ImDrawList_IdxBuffer = offsetof(ImDrawList, IdxBuffer);
// static const int offsetof_ImDrawList_VtxBuffer = offsetof(ImDrawList, VtxBuffer);
// static const int offsetof_ImDrawList__CallbacksDataBuf = offsetof(ImDrawList, _CallbacksDataBuf);
// static const int offsetof_ImDrawList__ClipRectStack = offsetof(ImDrawList, _ClipRectStack);
// static const int offsetof_ImDrawList__CmdHeader = offsetof(ImDrawList, _CmdHeader);
// static const int offsetof_ImDrawList__Path = offsetof(ImDrawList, _Path);
// static const int offsetof_ImDrawList__Splitter = offsetof(ImDrawList, _Splitter);
// static const int offsetof_ImDrawList__TextureStack = offsetof(ImDrawList, _TextureStack);
// static const int offsetof_ImDrawListSharedData_ClipRectFullscreen = offsetof(ImDrawListSharedData, ClipRectFullscreen);
// static const int offsetof_ImDrawListSharedData_DrawLists = offsetof(ImDrawListSharedData, DrawLists);
// static const int offsetof_ImDrawListSharedData_TempBuffer = offsetof(ImDrawListSharedData, TempBuffer);
// static const int offsetof_ImDrawListSharedData_TexUvWhitePixel = offsetof(ImDrawListSharedData, TexUvWhitePixel);
// static const int offsetof_ImDrawListSplitter__Channels = offsetof(ImDrawListSplitter, _Channels);
// static const int offsetof_ImDrawVert_pos = offsetof(ImDrawVert, pos);
// static const int offsetof_ImDrawVert_uv = offsetof(ImDrawVert, uv);
// static const int offsetof_ImFont_RemapPairs = offsetof(ImFont, RemapPairs);
// static const int offsetof_ImFont_Sources = offsetof(ImFont, Sources);
// static const int offsetof_ImFontAtlas_DrawListSharedDatas = offsetof(ImFontAtlas, DrawListSharedDatas);
// static const int offsetof_ImFontAtlas_Fonts = offsetof(ImFontAtlas, Fonts);
// static const int offsetof_ImFontAtlas_Sources = offsetof(ImFontAtlas, Sources);
// static const int offsetof_ImFontAtlas_TexList = offsetof(ImFontAtlas, TexList);
// static const int offsetof_ImFontAtlas_TexRef = offsetof(ImFontAtlas, TexRef);
// static const int offsetof_ImFontAtlas_TexUvScale = offsetof(ImFontAtlas, TexUvScale);
// static const int offsetof_ImFontAtlas_TexUvWhitePixel = offsetof(ImFontAtlas, TexUvWhitePixel);
// static const int offsetof_ImFontAtlasBuilder_BakedMap = offsetof(ImFontAtlasBuilder, BakedMap);
// static const int offsetof_ImFontAtlasBuilder_MaxRectBounds = offsetof(ImFontAtlasBuilder, MaxRectBounds);
// static const int offsetof_ImFontAtlasBuilder_MaxRectSize = offsetof(ImFontAtlasBuilder, MaxRectSize);
// static const int offsetof_ImFontAtlasBuilder_PackNodes = offsetof(ImFontAtlasBuilder, PackNodes);
// static const int offsetof_ImFontAtlasBuilder_Rects = offsetof(ImFontAtlasBuilder, Rects);
// static const int offsetof_ImFontAtlasBuilder_RectsIndex = offsetof(ImFontAtlasBuilder, RectsIndex);
// static const int offsetof_ImFontAtlasRect_uv0 = offsetof(ImFontAtlasRect, uv0);
// static const int offsetof_ImFontAtlasRect_uv1 = offsetof(ImFontAtlasRect, uv1);
// static const int offsetof_ImFontBaked_Glyphs = offsetof(ImFontBaked, Glyphs);
// static const int offsetof_ImFontBaked_IndexAdvanceX = offsetof(ImFontBaked, IndexAdvanceX);
// static const int offsetof_ImFontBaked_IndexLookup = offsetof(ImFontBaked, IndexLookup);
// static const int offsetof_ImFontConfig_GlyphOffset = offsetof(ImFontConfig, GlyphOffset);
// static const int offsetof_ImFontGlyphRangesBuilder_UsedChars = offsetof(ImFontGlyphRangesBuilder, UsedChars);
// static const int offsetof_ImGuiBoxSelectState_BoxSelectRectCurr = offsetof(ImGuiBoxSelectState, BoxSelectRectCurr);
// static const int offsetof_ImGuiBoxSelectState_BoxSelectRectPrev = offsetof(ImGuiBoxSelectState, BoxSelectRectPrev);
// static const int offsetof_ImGuiBoxSelectState_EndPosRel = offsetof(ImGuiBoxSelectState, EndPosRel);
// static const int offsetof_ImGuiBoxSelectState_ScrollAccum = offsetof(ImGuiBoxSelectState, ScrollAccum);
// static const int offsetof_ImGuiBoxSelectState_StartPosRel = offsetof(ImGuiBoxSelectState, StartPosRel);
// static const int offsetof_ImGuiBoxSelectState_UnclipRect = offsetof(ImGuiBoxSelectState, UnclipRect);
// static const int offsetof_ImGuiColorMod_BackupValue = offsetof(ImGuiColorMod, BackupValue);
// static const int offsetof_ImGuiComboPreviewData_BackupCursorMaxPos = offsetof(ImGuiComboPreviewData, BackupCursorMaxPos);
// static const int offsetof_ImGuiComboPreviewData_BackupCursorPos = offsetof(ImGuiComboPreviewData, BackupCursorPos);
// static const int offsetof_ImGuiComboPreviewData_BackupCursorPosPrevLine = offsetof(ImGuiComboPreviewData, BackupCursorPosPrevLine);
// static const int offsetof_ImGuiComboPreviewData_PreviewRect = offsetof(ImGuiComboPreviewData, PreviewRect);
// static const int offsetof_ImGuiContext_ActiveIdClickOffset = offsetof(ImGuiContext, ActiveIdClickOffset);
// static const int offsetof_ImGuiContext_ActiveIdValueOnActivation = offsetof(ImGuiContext, ActiveIdValueOnActivation);
// static const int offsetof_ImGuiContext_BeginPopupStack = offsetof(ImGuiContext, BeginPopupStack);
// static const int offsetof_ImGuiContext_BoxSelectState = offsetof(ImGuiContext, BoxSelectState);
// static const int offsetof_ImGuiContext_ClipboardHandlerData = offsetof(ImGuiContext, ClipboardHandlerData);
// static const int offsetof_ImGuiContext_ClipperTempData = offsetof(ImGuiContext, ClipperTempData);
// static const int offsetof_ImGuiContext_ColorPickerRef = offsetof(ImGuiContext, ColorPickerRef);
// static const int offsetof_ImGuiContext_ColorStack = offsetof(ImGuiContext, ColorStack);
// static const int offsetof_ImGuiContext_ComboPreviewData = offsetof(ImGuiContext, ComboPreviewData);
// static const int offsetof_ImGuiContext_CurrentTabBarStack = offsetof(ImGuiContext, CurrentTabBarStack);
// static const int offsetof_ImGuiContext_CurrentWindowStack = offsetof(ImGuiContext, CurrentWindowStack);
// static const int offsetof_ImGuiContext_DataTypeZeroValue = offsetof(ImGuiContext, DataTypeZeroValue);
// static const int offsetof_ImGuiContext_DeactivatedItemData = offsetof(ImGuiContext, DeactivatedItemData);
// static const int offsetof_ImGuiContext_DebugAllocInfo = offsetof(ImGuiContext, DebugAllocInfo);
// static const int offsetof_ImGuiContext_DebugFlashStyleColorBackup = offsetof(ImGuiContext, DebugFlashStyleColorBackup);
// static const int offsetof_ImGuiContext_DebugIDStackTool = offsetof(ImGuiContext, DebugIDStackTool);
// static const int offsetof_ImGuiContext_DebugLogBuf = offsetof(ImGuiContext, DebugLogBuf);
// static const int offsetof_ImGuiContext_DebugLogIndex = offsetof(ImGuiContext, DebugLogIndex);
// static const int offsetof_ImGuiContext_DebugMetricsConfig = offsetof(ImGuiContext, DebugMetricsConfig);
// static const int offsetof_ImGuiContext_DockContext = offsetof(ImGuiContext, DockContext);
// static const int offsetof_ImGuiContext_DragDropPayload = offsetof(ImGuiContext, DragDropPayload);
// static const int offsetof_ImGuiContext_DragDropTargetClipRect = offsetof(ImGuiContext, DragDropTargetClipRect);
// static const int offsetof_ImGuiContext_DragDropTargetRect = offsetof(ImGuiContext, DragDropTargetRect);
// static const int offsetof_ImGuiContext_DrawChannelsTempMergeBuffer = offsetof(ImGuiContext, DrawChannelsTempMergeBuffer);
// static const int offsetof_ImGuiContext_DrawListSharedData = offsetof(ImGuiContext, DrawListSharedData);
// static const int offsetof_ImGuiContext_ErrorTooltipLockedPos = offsetof(ImGuiContext, ErrorTooltipLockedPos);
// static const int offsetof_ImGuiContext_FallbackMonitor = offsetof(ImGuiContext, FallbackMonitor);
// static const int offsetof_ImGuiContext_FocusScopeStack = offsetof(ImGuiContext, FocusScopeStack);
// static const int offsetof_ImGuiContext_FontAtlases = offsetof(ImGuiContext, FontAtlases);
// static const int offsetof_ImGuiContext_FontStack = offsetof(ImGuiContext, FontStack);
// static const int offsetof_ImGuiContext_GroupStack = offsetof(ImGuiContext, GroupStack);
// static const int offsetof_ImGuiContext_Hooks = offsetof(ImGuiContext, Hooks);
// static const int offsetof_ImGuiContext_IO = offsetof(ImGuiContext, IO);
// static const int offsetof_ImGuiContext_InputEventsQueue = offsetof(ImGuiContext, InputEventsQueue);
// static const int offsetof_ImGuiContext_InputEventsTrail = offsetof(ImGuiContext, InputEventsTrail);
// static const int offsetof_ImGuiContext_InputTextDeactivatedState = offsetof(ImGuiContext, InputTextDeactivatedState);
// static const int offsetof_ImGuiContext_InputTextLineIndex = offsetof(ImGuiContext, InputTextLineIndex);
// static const int offsetof_ImGuiContext_InputTextPasswordFontBackupBaked = offsetof(ImGuiContext, InputTextPasswordFontBackupBaked);
// static const int offsetof_ImGuiContext_InputTextState = offsetof(ImGuiContext, InputTextState);
// static const int offsetof_ImGuiContext_ItemFlagsStack = offsetof(ImGuiContext, ItemFlagsStack);
// static const int offsetof_ImGuiContext_KeysRoutingTable = offsetof(ImGuiContext, KeysRoutingTable);
// static const int offsetof_ImGuiContext_LastItemData = offsetof(ImGuiContext, LastItemData);
// static const int offsetof_ImGuiContext_LogBuffer = offsetof(ImGuiContext, LogBuffer);
// static const int offsetof_ImGuiContext_MenusIdSubmittedThisFrame = offsetof(ImGuiContext, MenusIdSubmittedThisFrame);
// static const int offsetof_ImGuiContext_MouseLastValidPos = offsetof(ImGuiContext, MouseLastValidPos);
// static const int offsetof_ImGuiContext_MultiSelectStorage = offsetof(ImGuiContext, MultiSelectStorage);
// static const int offsetof_ImGuiContext_MultiSelectTempData = offsetof(ImGuiContext, MultiSelectTempData);
// static const int offsetof_ImGuiContext_NavFocusRoute = offsetof(ImGuiContext, NavFocusRoute);
// static const int offsetof_ImGuiContext_NavInitResult = offsetof(ImGuiContext, NavInitResult);
// static const int offsetof_ImGuiContext_NavMoveResultLocal = offsetof(ImGuiContext, NavMoveResultLocal);
// static const int offsetof_ImGuiContext_NavMoveResultLocalVisible = offsetof(ImGuiContext, NavMoveResultLocalVisible);
// static const int offsetof_ImGuiContext_NavMoveResultOther = offsetof(ImGuiContext, NavMoveResultOther);
// static const int offsetof_ImGuiContext_NavScoringNoClipRect = offsetof(ImGuiContext, NavScoringNoClipRect);
// static const int offsetof_ImGuiContext_NavScoringRect = offsetof(ImGuiContext, NavScoringRect);
// static const int offsetof_ImGuiContext_NavTabbingResultFirst = offsetof(ImGuiContext, NavTabbingResultFirst);
// static const int offsetof_ImGuiContext_NavWindowingAccumDeltaPos = offsetof(ImGuiContext, NavWindowingAccumDeltaPos);
// static const int offsetof_ImGuiContext_NavWindowingAccumDeltaSize = offsetof(ImGuiContext, NavWindowingAccumDeltaSize);
// static const int offsetof_ImGuiContext_NextItemData = offsetof(ImGuiContext, NextItemData);
// static const int offsetof_ImGuiContext_NextWindowData = offsetof(ImGuiContext, NextWindowData);
// static const int offsetof_ImGuiContext_OpenPopupStack = offsetof(ImGuiContext, OpenPopupStack);
// static const int offsetof_ImGuiContext_PlatformIO = offsetof(ImGuiContext, PlatformIO);
// static const int offsetof_ImGuiContext_PlatformImeData = offsetof(ImGuiContext, PlatformImeData);
// static const int offsetof_ImGuiContext_PlatformImeDataPrev = offsetof(ImGuiContext, PlatformImeDataPrev);
// static const int offsetof_ImGuiContext_PlatformMonitorsFullWorkRect = offsetof(ImGuiContext, PlatformMonitorsFullWorkRect);
// static const int offsetof_ImGuiContext_SettingsHandlers = offsetof(ImGuiContext, SettingsHandlers);
// static const int offsetof_ImGuiContext_SettingsIniData = offsetof(ImGuiContext, SettingsIniData);
// static const int offsetof_ImGuiContext_SettingsTables = offsetof(ImGuiContext, SettingsTables);
// static const int offsetof_ImGuiContext_SettingsWindows = offsetof(ImGuiContext, SettingsWindows);
// static const int offsetof_ImGuiContext_ShrinkWidthBuffer = offsetof(ImGuiContext, ShrinkWidthBuffer);
// static const int offsetof_ImGuiContext_StackSizesInNewFrame = offsetof(ImGuiContext, StackSizesInNewFrame);
// static const int offsetof_ImGuiContext_Style = offsetof(ImGuiContext, Style);
// static const int offsetof_ImGuiContext_StyleVarStack = offsetof(ImGuiContext, StyleVarStack);
// static const int offsetof_ImGuiContext_TabBars = offsetof(ImGuiContext, TabBars);
// static const int offsetof_ImGuiContext_Tables = offsetof(ImGuiContext, Tables);
// static const int offsetof_ImGuiContext_TablesLastTimeActive = offsetof(ImGuiContext, TablesLastTimeActive);
// static const int offsetof_ImGuiContext_TablesTempData = offsetof(ImGuiContext, TablesTempData);
// static const int offsetof_ImGuiContext_TempBuffer = offsetof(ImGuiContext, TempBuffer);
// static const int offsetof_ImGuiContext_TreeNodeStack = offsetof(ImGuiContext, TreeNodeStack);
// static const int offsetof_ImGuiContext_TypingSelectState = offsetof(ImGuiContext, TypingSelectState);
// static const int offsetof_ImGuiContext_UserTextures = offsetof(ImGuiContext, UserTextures);
// static const int offsetof_ImGuiContext_Viewports = offsetof(ImGuiContext, Viewports);
// static const int offsetof_ImGuiContext_WheelingAxisAvg = offsetof(ImGuiContext, WheelingAxisAvg);
// static const int offsetof_ImGuiContext_WheelingWindowRefMousePos = offsetof(ImGuiContext, WheelingWindowRefMousePos);
// static const int offsetof_ImGuiContext_WheelingWindowWheelRemainder = offsetof(ImGuiContext, WheelingWindowWheelRemainder);
// static const int offsetof_ImGuiContext_WindowResizeBorderExpectedRect = offsetof(ImGuiContext, WindowResizeBorderExpectedRect);
// static const int offsetof_ImGuiContext_Windows = offsetof(ImGuiContext, Windows);
// static const int offsetof_ImGuiContext_WindowsById = offsetof(ImGuiContext, WindowsById);
// static const int offsetof_ImGuiContext_WindowsFocusOrder = offsetof(ImGuiContext, WindowsFocusOrder);
// static const int offsetof_ImGuiContext_WindowsTempSortBuffer = offsetof(ImGuiContext, WindowsTempSortBuffer);
// static const int offsetof_ImGuiDockContext_Nodes = offsetof(ImGuiDockContext, Nodes);
// static const int offsetof_ImGuiDockContext_NodesSettings = offsetof(ImGuiDockContext, NodesSettings);
// static const int offsetof_ImGuiDockContext_Requests = offsetof(ImGuiDockContext, Requests);
// static const int offsetof_ImGuiDockNode_Pos = offsetof(ImGuiDockNode, Pos);
// static const int offsetof_ImGuiDockNode_Size = offsetof(ImGuiDockNode, Size);
// static const int offsetof_ImGuiDockNode_SizeRef = offsetof(ImGuiDockNode, SizeRef);
// static const int offsetof_ImGuiDockNode_WindowClass = offsetof(ImGuiDockNode, WindowClass);
// static const int offsetof_ImGuiDockNode_Windows = offsetof(ImGuiDockNode, Windows);
// static const int offsetof_ImGuiGroupData_BackupCurrLineSize = offsetof(ImGuiGroupData, BackupCurrLineSize);
// static const int offsetof_ImGuiGroupData_BackupCursorMaxPos = offsetof(ImGuiGroupData, BackupCursorMaxPos);
// static const int offsetof_ImGuiGroupData_BackupCursorPos = offsetof(ImGuiGroupData, BackupCursorPos);
// static const int offsetof_ImGuiGroupData_BackupCursorPosPrevLine = offsetof(ImGuiGroupData, BackupCursorPosPrevLine);
// static const int offsetof_ImGuiGroupData_BackupGroupOffset = offsetof(ImGuiGroupData, BackupGroupOffset);
// static const int offsetof_ImGuiGroupData_BackupIndent = offsetof(ImGuiGroupData, BackupIndent);
// static const int offsetof_ImGuiIDStackTool_ResultPathsBuf = offsetof(ImGuiIDStackTool, ResultPathsBuf);
// static const int offsetof_ImGuiIDStackTool_ResultTempBuf = offsetof(ImGuiIDStackTool, ResultTempBuf);
// static const int offsetof_ImGuiIDStackTool_Results = offsetof(ImGuiIDStackTool, Results);
// static const int offsetof_ImGuiIO_DisplayFramebufferScale = offsetof(ImGuiIO, DisplayFramebufferScale);
// static const int offsetof_ImGuiIO_DisplaySize = offsetof(ImGuiIO, DisplaySize);
// static const int offsetof_ImGuiIO_InputQueueCharacters = offsetof(ImGuiIO, InputQueueCharacters);
// static const int offsetof_ImGuiIO_MouseDelta = offsetof(ImGuiIO, MouseDelta);
// static const int offsetof_ImGuiIO_MousePos = offsetof(ImGuiIO, MousePos);
// static const int offsetof_ImGuiIO_MousePosPrev = offsetof(ImGuiIO, MousePosPrev);
// static const int offsetof_ImGuiInputTextDeactivatedState_TextA = offsetof(ImGuiInputTextDeactivatedState, TextA);
// static const int offsetof_ImGuiInputTextState_CallbackTextBackup = offsetof(ImGuiInputTextState, CallbackTextBackup);
// static const int offsetof_ImGuiInputTextState_Scroll = offsetof(ImGuiInputTextState, Scroll);
// static const int offsetof_ImGuiInputTextState_TextA = offsetof(ImGuiInputTextState, TextA);
// static const int offsetof_ImGuiInputTextState_TextToRevertTo = offsetof(ImGuiInputTextState, TextToRevertTo);
// static const int offsetof_ImGuiKeyRoutingTable_Entries = offsetof(ImGuiKeyRoutingTable, Entries);
// static const int offsetof_ImGuiKeyRoutingTable_EntriesNext = offsetof(ImGuiKeyRoutingTable, EntriesNext);
// static const int offsetof_ImGuiLastItemData_ClipRect = offsetof(ImGuiLastItemData, ClipRect);
// static const int offsetof_ImGuiLastItemData_DisplayRect = offsetof(ImGuiLastItemData, DisplayRect);
// static const int offsetof_ImGuiLastItemData_NavRect = offsetof(ImGuiLastItemData, NavRect);
// static const int offsetof_ImGuiLastItemData_Rect = offsetof(ImGuiLastItemData, Rect);
// static const int offsetof_ImGuiListClipperData_Ranges = offsetof(ImGuiListClipperData, Ranges);
// static const int offsetof_ImGuiMultiSelectIO_Requests = offsetof(ImGuiMultiSelectIO, Requests);
// static const int offsetof_ImGuiMultiSelectTempData_BackupCursorMaxPos = offsetof(ImGuiMultiSelectTempData, BackupCursorMaxPos);
// static const int offsetof_ImGuiMultiSelectTempData_IO = offsetof(ImGuiMultiSelectTempData, IO);
// static const int offsetof_ImGuiMultiSelectTempData_ScopeRectMin = offsetof(ImGuiMultiSelectTempData, ScopeRectMin);
// static const int offsetof_ImGuiNavItemData_RectRel = offsetof(ImGuiNavItemData, RectRel);
// static const int offsetof_ImGuiNextItemData_RefVal = offsetof(ImGuiNextItemData, RefVal);
// static const int offsetof_ImGuiNextWindowData_ContentSizeVal = offsetof(ImGuiNextWindowData, ContentSizeVal);
// static const int offsetof_ImGuiNextWindowData_MenuBarOffsetMinVal = offsetof(ImGuiNextWindowData, MenuBarOffsetMinVal);
// static const int offsetof_ImGuiNextWindowData_PosPivotVal = offsetof(ImGuiNextWindowData, PosPivotVal);
// static const int offsetof_ImGuiNextWindowData_PosVal = offsetof(ImGuiNextWindowData, PosVal);
// static const int offsetof_ImGuiNextWindowData_ScrollVal = offsetof(ImGuiNextWindowData, ScrollVal);
// static const int offsetof_ImGuiNextWindowData_SizeConstraintRect = offsetof(ImGuiNextWindowData, SizeConstraintRect);
// static const int offsetof_ImGuiNextWindowData_SizeVal = offsetof(ImGuiNextWindowData, SizeVal);
// static const int offsetof_ImGuiNextWindowData_WindowClass = offsetof(ImGuiNextWindowData, WindowClass);
// static const int offsetof_ImGuiOldColumnData_ClipRect = offsetof(ImGuiOldColumnData, ClipRect);
// static const int offsetof_ImGuiOldColumns_Columns = offsetof(ImGuiOldColumns, Columns);
// static const int offsetof_ImGuiOldColumns_HostBackupClipRect = offsetof(ImGuiOldColumns, HostBackupClipRect);
// static const int offsetof_ImGuiOldColumns_HostBackupParentWorkRect = offsetof(ImGuiOldColumns, HostBackupParentWorkRect);
// static const int offsetof_ImGuiOldColumns_HostInitialClipRect = offsetof(ImGuiOldColumns, HostInitialClipRect);
// static const int offsetof_ImGuiOldColumns_Splitter = offsetof(ImGuiOldColumns, Splitter);
// static const int offsetof_ImGuiPlatformIO_Monitors = offsetof(ImGuiPlatformIO, Monitors);
// static const int offsetof_ImGuiPlatformIO_Textures = offsetof(ImGuiPlatformIO, Textures);
// static const int offsetof_ImGuiPlatformIO_Viewports = offsetof(ImGuiPlatformIO, Viewports);
// static const int offsetof_ImGuiPlatformImeData_InputPos = offsetof(ImGuiPlatformImeData, InputPos);
// static const int offsetof_ImGuiPlatformMonitor_MainPos = offsetof(ImGuiPlatformMonitor, MainPos);
// static const int offsetof_ImGuiPlatformMonitor_MainSize = offsetof(ImGuiPlatformMonitor, MainSize);
// static const int offsetof_ImGuiPlatformMonitor_WorkPos = offsetof(ImGuiPlatformMonitor, WorkPos);
// static const int offsetof_ImGuiPlatformMonitor_WorkSize = offsetof(ImGuiPlatformMonitor, WorkSize);
// static const int offsetof_ImGuiPopupData_OpenMousePos = offsetof(ImGuiPopupData, OpenMousePos);
// static const int offsetof_ImGuiPopupData_OpenPopupPos = offsetof(ImGuiPopupData, OpenPopupPos);
// static const int offsetof_ImGuiSelectionBasicStorage__Storage = offsetof(ImGuiSelectionBasicStorage, _Storage);
// static const int offsetof_ImGuiSizeCallbackData_CurrentSize = offsetof(ImGuiSizeCallbackData, CurrentSize);
// static const int offsetof_ImGuiSizeCallbackData_DesiredSize = offsetof(ImGuiSizeCallbackData, DesiredSize);
// static const int offsetof_ImGuiSizeCallbackData_Pos = offsetof(ImGuiSizeCallbackData, Pos);
// static const int offsetof_ImGuiStorage_Data = offsetof(ImGuiStorage, Data);
// static const int offsetof_ImGuiStyle_ButtonTextAlign = offsetof(ImGuiStyle, ButtonTextAlign);
// static const int offsetof_ImGuiStyle_CellPadding = offsetof(ImGuiStyle, CellPadding);
// static const int offsetof_ImGuiStyle_DisplaySafeAreaPadding = offsetof(ImGuiStyle, DisplaySafeAreaPadding);
// static const int offsetof_ImGuiStyle_DisplayWindowPadding = offsetof(ImGuiStyle, DisplayWindowPadding);
// static const int offsetof_ImGuiStyle_FramePadding = offsetof(ImGuiStyle, FramePadding);
// static const int offsetof_ImGuiStyle_ItemInnerSpacing = offsetof(ImGuiStyle, ItemInnerSpacing);
// static const int offsetof_ImGuiStyle_ItemSpacing = offsetof(ImGuiStyle, ItemSpacing);
// static const int offsetof_ImGuiStyle_SelectableTextAlign = offsetof(ImGuiStyle, SelectableTextAlign);
// static const int offsetof_ImGuiStyle_SeparatorTextAlign = offsetof(ImGuiStyle, SeparatorTextAlign);
// static const int offsetof_ImGuiStyle_SeparatorTextPadding = offsetof(ImGuiStyle, SeparatorTextPadding);
// static const int offsetof_ImGuiStyle_TableAngledHeadersTextAlign = offsetof(ImGuiStyle, TableAngledHeadersTextAlign);
// static const int offsetof_ImGuiStyle_TouchExtraPadding = offsetof(ImGuiStyle, TouchExtraPadding);
// static const int offsetof_ImGuiStyle_WindowMinSize = offsetof(ImGuiStyle, WindowMinSize);
// static const int offsetof_ImGuiStyle_WindowPadding = offsetof(ImGuiStyle, WindowPadding);
// static const int offsetof_ImGuiStyle_WindowTitleAlign = offsetof(ImGuiStyle, WindowTitleAlign);
// static const int offsetof_ImGuiTabBar_BackupCursorPos = offsetof(ImGuiTabBar, BackupCursorPos);
// static const int offsetof_ImGuiTabBar_BarRect = offsetof(ImGuiTabBar, BarRect);
// static const int offsetof_ImGuiTabBar_FramePadding = offsetof(ImGuiTabBar, FramePadding);
// static const int offsetof_ImGuiTabBar_Tabs = offsetof(ImGuiTabBar, Tabs);
// static const int offsetof_ImGuiTabBar_TabsNames = offsetof(ImGuiTabBar, TabsNames);
// static const int offsetof_ImGuiTable_Bg0ClipRectForDrawCmd = offsetof(ImGuiTable, Bg0ClipRectForDrawCmd);
// static const int offsetof_ImGuiTable_Bg2ClipRectForDrawCmd = offsetof(ImGuiTable, Bg2ClipRectForDrawCmd);
// static const int offsetof_ImGuiTable_BgClipRect = offsetof(ImGuiTable, BgClipRect);
// static const int offsetof_ImGuiTable_Columns = offsetof(ImGuiTable, Columns);
// static const int offsetof_ImGuiTable_ColumnsNames = offsetof(ImGuiTable, ColumnsNames);
// static const int offsetof_ImGuiTable_DisplayOrderToIndex = offsetof(ImGuiTable, DisplayOrderToIndex);
// static const int offsetof_ImGuiTable_HostBackupInnerClipRect = offsetof(ImGuiTable, HostBackupInnerClipRect);
// static const int offsetof_ImGuiTable_HostClipRect = offsetof(ImGuiTable, HostClipRect);
// static const int offsetof_ImGuiTable_InnerClipRect = offsetof(ImGuiTable, InnerClipRect);
// static const int offsetof_ImGuiTable_InnerRect = offsetof(ImGuiTable, InnerRect);
// static const int offsetof_ImGuiTable_InstanceDataExtra = offsetof(ImGuiTable, InstanceDataExtra);
// static const int offsetof_ImGuiTable_InstanceDataFirst = offsetof(ImGuiTable, InstanceDataFirst);
// static const int offsetof_ImGuiTable_OuterRect = offsetof(ImGuiTable, OuterRect);
// static const int offsetof_ImGuiTable_RowCellData = offsetof(ImGuiTable, RowCellData);
// static const int offsetof_ImGuiTable_SortSpecs = offsetof(ImGuiTable, SortSpecs);
// static const int offsetof_ImGuiTable_SortSpecsMulti = offsetof(ImGuiTable, SortSpecsMulti);
// static const int offsetof_ImGuiTable_SortSpecsSingle = offsetof(ImGuiTable, SortSpecsSingle);
// static const int offsetof_ImGuiTable_WorkRect = offsetof(ImGuiTable, WorkRect);
// static const int offsetof_ImGuiTableColumn_ClipRect = offsetof(ImGuiTableColumn, ClipRect);
// static const int offsetof_ImGuiTableTempData_AngledHeadersRequests = offsetof(ImGuiTableTempData, AngledHeadersRequests);
// static const int offsetof_ImGuiTableTempData_DrawSplitter = offsetof(ImGuiTableTempData, DrawSplitter);
// static const int offsetof_ImGuiTableTempData_HostBackupColumnsOffset = offsetof(ImGuiTableTempData, HostBackupColumnsOffset);
// static const int offsetof_ImGuiTableTempData_HostBackupCurrLineSize = offsetof(ImGuiTableTempData, HostBackupCurrLineSize);
// static const int offsetof_ImGuiTableTempData_HostBackupCursorMaxPos = offsetof(ImGuiTableTempData, HostBackupCursorMaxPos);
// static const int offsetof_ImGuiTableTempData_HostBackupParentWorkRect = offsetof(ImGuiTableTempData, HostBackupParentWorkRect);
// static const int offsetof_ImGuiTableTempData_HostBackupPrevLineSize = offsetof(ImGuiTableTempData, HostBackupPrevLineSize);
// static const int offsetof_ImGuiTableTempData_HostBackupWorkRect = offsetof(ImGuiTableTempData, HostBackupWorkRect);
// static const int offsetof_ImGuiTableTempData_UserOuterSize = offsetof(ImGuiTableTempData, UserOuterSize);
// static const int offsetof_ImGuiTextBuffer_Buf = offsetof(ImGuiTextBuffer, Buf);
// static const int offsetof_ImGuiTextFilter_Filters = offsetof(ImGuiTextFilter, Filters);
// static const int offsetof_ImGuiTextIndex_Offsets = offsetof(ImGuiTextIndex, Offsets);
// static const int offsetof_ImGuiTreeNodeStackData_NavRect = offsetof(ImGuiTreeNodeStackData, NavRect);
// static const int offsetof_ImGuiTypingSelectState_Request = offsetof(ImGuiTypingSelectState, Request);
// static const int offsetof_ImGuiViewport_FramebufferScale = offsetof(ImGuiViewport, FramebufferScale);
// static const int offsetof_ImGuiViewport_Pos = offsetof(ImGuiViewport, Pos);
// static const int offsetof_ImGuiViewport_Size = offsetof(ImGuiViewport, Size);
// static const int offsetof_ImGuiViewport_WorkPos = offsetof(ImGuiViewport, WorkPos);
// static const int offsetof_ImGuiViewport_WorkSize = offsetof(ImGuiViewport, WorkSize);
// static const int offsetof_ImGuiViewportP_BuildWorkInsetMax = offsetof(ImGuiViewportP, BuildWorkInsetMax);
// static const int offsetof_ImGuiViewportP_BuildWorkInsetMin = offsetof(ImGuiViewportP, BuildWorkInsetMin);
// static const int offsetof_ImGuiViewportP_DrawDataBuilder = offsetof(ImGuiViewportP, DrawDataBuilder);
// static const int offsetof_ImGuiViewportP_DrawDataP = offsetof(ImGuiViewportP, DrawDataP);
// static const int offsetof_ImGuiViewportP_LastPlatformPos = offsetof(ImGuiViewportP, LastPlatformPos);
// static const int offsetof_ImGuiViewportP_LastPlatformSize = offsetof(ImGuiViewportP, LastPlatformSize);
// static const int offsetof_ImGuiViewportP_LastPos = offsetof(ImGuiViewportP, LastPos);
// static const int offsetof_ImGuiViewportP_LastRendererSize = offsetof(ImGuiViewportP, LastRendererSize);
// static const int offsetof_ImGuiViewportP_LastSize = offsetof(ImGuiViewportP, LastSize);
// static const int offsetof_ImGuiViewportP_WorkInsetMax = offsetof(ImGuiViewportP, WorkInsetMax);
// static const int offsetof_ImGuiViewportP_WorkInsetMin = offsetof(ImGuiViewportP, WorkInsetMin);
// static const int offsetof_ImGuiViewportP__ImGuiViewport = offsetof(ImGuiViewportP, _ImGuiViewport);
// static const int offsetof_ImGuiWindow_ClipRect = offsetof(ImGuiWindow, ClipRect);
// static const int offsetof_ImGuiWindow_ColumnsStorage = offsetof(ImGuiWindow, ColumnsStorage);
// static const int offsetof_ImGuiWindow_ContentRegionRect = offsetof(ImGuiWindow, ContentRegionRect);
// static const int offsetof_ImGuiWindow_ContentSize = offsetof(ImGuiWindow, ContentSize);
// static const int offsetof_ImGuiWindow_ContentSizeExplicit = offsetof(ImGuiWindow, ContentSizeExplicit);
// static const int offsetof_ImGuiWindow_ContentSizeIdeal = offsetof(ImGuiWindow, ContentSizeIdeal);
// static const int offsetof_ImGuiWindow_DC = offsetof(ImGuiWindow, DC);
// static const int offsetof_ImGuiWindow_DockStyle = offsetof(ImGuiWindow, DockStyle);
// static const int offsetof_ImGuiWindow_DrawListInst = offsetof(ImGuiWindow, DrawListInst);
// static const int offsetof_ImGuiWindow_HitTestHoleOffset = offsetof(ImGuiWindow, HitTestHoleOffset);
// static const int offsetof_ImGuiWindow_HitTestHoleSize = offsetof(ImGuiWindow, HitTestHoleSize);
// static const int offsetof_ImGuiWindow_IDStack = offsetof(ImGuiWindow, IDStack);
// static const int offsetof_ImGuiWindow_InnerClipRect = offsetof(ImGuiWindow, InnerClipRect);
// static const int offsetof_ImGuiWindow_InnerRect = offsetof(ImGuiWindow, InnerRect);
// static const int offsetof_ImGuiWindow_OuterRectClipped = offsetof(ImGuiWindow, OuterRectClipped);
// static const int offsetof_ImGuiWindow_ParentWorkRect = offsetof(ImGuiWindow, ParentWorkRect);
// static const int offsetof_ImGuiWindow_Pos = offsetof(ImGuiWindow, Pos);
// static const int offsetof_ImGuiWindow_Scroll = offsetof(ImGuiWindow, Scroll);
// static const int offsetof_ImGuiWindow_ScrollMax = offsetof(ImGuiWindow, ScrollMax);
// static const int offsetof_ImGuiWindow_ScrollTarget = offsetof(ImGuiWindow, ScrollTarget);
// static const int offsetof_ImGuiWindow_ScrollTargetCenterRatio = offsetof(ImGuiWindow, ScrollTargetCenterRatio);
// static const int offsetof_ImGuiWindow_ScrollTargetEdgeSnapDist = offsetof(ImGuiWindow, ScrollTargetEdgeSnapDist);
// static const int offsetof_ImGuiWindow_ScrollbarSizes = offsetof(ImGuiWindow, ScrollbarSizes);
// static const int offsetof_ImGuiWindow_SetWindowPosPivot = offsetof(ImGuiWindow, SetWindowPosPivot);
// static const int offsetof_ImGuiWindow_SetWindowPosVal = offsetof(ImGuiWindow, SetWindowPosVal);
// static const int offsetof_ImGuiWindow_Size = offsetof(ImGuiWindow, Size);
// static const int offsetof_ImGuiWindow_SizeFull = offsetof(ImGuiWindow, SizeFull);
// static const int offsetof_ImGuiWindow_StateStorage = offsetof(ImGuiWindow, StateStorage);
// static const int offsetof_ImGuiWindow_ViewportPos = offsetof(ImGuiWindow, ViewportPos);
// static const int offsetof_ImGuiWindow_WindowClass = offsetof(ImGuiWindow, WindowClass);
// static const int offsetof_ImGuiWindow_WindowPadding = offsetof(ImGuiWindow, WindowPadding);
// static const int offsetof_ImGuiWindow_WorkRect = offsetof(ImGuiWindow, WorkRect);
// static const int offsetof_ImGuiWindowSettings_Pos = offsetof(ImGuiWindowSettings, Pos);
// static const int offsetof_ImGuiWindowSettings_Size = offsetof(ImGuiWindowSettings, Size);
// static const int offsetof_ImGuiWindowSettings_ViewportPos = offsetof(ImGuiWindowSettings, ViewportPos);
// static const int offsetof_ImGuiWindowStackData_ParentLastItemDataBackup = offsetof(ImGuiWindowStackData, ParentLastItemDataBackup);
// static const int offsetof_ImGuiWindowStackData_StackSizesInBegin = offsetof(ImGuiWindowStackData, StackSizesInBegin);
// static const int offsetof_ImGuiWindowTempData_ChildWindows = offsetof(ImGuiWindowTempData, ChildWindows);
// static const int offsetof_ImGuiWindowTempData_ColumnsOffset = offsetof(ImGuiWindowTempData, ColumnsOffset);
// static const int offsetof_ImGuiWindowTempData_CurrLineSize = offsetof(ImGuiWindowTempData, CurrLineSize);
// static const int offsetof_ImGuiWindowTempData_CursorMaxPos = offsetof(ImGuiWindowTempData, CursorMaxPos);
// static const int offsetof_ImGuiWindowTempData_CursorPos = offsetof(ImGuiWindowTempData, CursorPos);
// static const int offsetof_ImGuiWindowTempData_CursorPosPrevLine = offsetof(ImGuiWindowTempData, CursorPosPrevLine);
// static const int offsetof_ImGuiWindowTempData_CursorStartPos = offsetof(ImGuiWindowTempData, CursorStartPos);
// static const int offsetof_ImGuiWindowTempData_CursorStartPosLossyness = offsetof(ImGuiWindowTempData, CursorStartPosLossyness);
// static const int offsetof_ImGuiWindowTempData_DockTabItemRect = offsetof(ImGuiWindowTempData, DockTabItemRect);
// static const int offsetof_ImGuiWindowTempData_GroupOffset = offsetof(ImGuiWindowTempData, GroupOffset);
// static const int offsetof_ImGuiWindowTempData_IdealMaxPos = offsetof(ImGuiWindowTempData, IdealMaxPos);
// static const int offsetof_ImGuiWindowTempData_Indent = offsetof(ImGuiWindowTempData, Indent);
// static const int offsetof_ImGuiWindowTempData_ItemWidthStack = offsetof(ImGuiWindowTempData, ItemWidthStack);
// static const int offsetof_ImGuiWindowTempData_MenuBarOffset = offsetof(ImGuiWindowTempData, MenuBarOffset);
// static const int offsetof_ImGuiWindowTempData_MenuColumns = offsetof(ImGuiWindowTempData, MenuColumns);
// static const int offsetof_ImGuiWindowTempData_PrevLineSize = offsetof(ImGuiWindowTempData, PrevLineSize);
// static const int offsetof_ImGuiWindowTempData_TextWrapPosStack = offsetof(ImGuiWindowTempData, TextWrapPosStack);
// static const int offsetof_ImPool_ImGuiMultiSelectState_Map = offsetof(ImPool_ImGuiMultiSelectState, Map);
// static const int offsetof_ImPool_ImGuiTabBar_Map = offsetof(ImPool_ImGuiTabBar, Map);
// static const int offsetof_ImPool_ImGuiTable_Map = offsetof(ImPool_ImGuiTable, Map);
// static const int offsetof_ImRect_Max = offsetof(ImRect, Max);
// static const int offsetof_ImRect_Min = offsetof(ImRect, Min);
// static const int offsetof_ImTextureData_UpdateRect = offsetof(ImTextureData, UpdateRect);
// static const int offsetof_ImTextureData_Updates = offsetof(ImTextureData, Updates);
// static const int offsetof_ImTextureData_UsedRect = offsetof(ImTextureData, UsedRect);
import "C"

// BitVector wraps struct ImBitVector.
type BitVector uintptr

// BitVectorNil is a null pointer.
var BitVectorNil BitVector

// BitVectorSizeOf is the byte size of ImBitVector.
const BitVectorSizeOf = int(C.sizeof_ImBitVector)

// BitVectorAlloc allocates a continuous block of BitVector.
func BitVectorAlloc(alloc ffi.Allocator, count int) BitVector {
	ptr := alloc.Allocate(BitVectorSizeOf * count)
	return BitVector(ptr)
}

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p BitVector) Offset(offset int) BitVector {
	return p + BitVector(offset*BitVectorSizeOf)
}

// RefStorage returns pointer to the Storage field.
func (p BitVector) RefStorage() Vector_ImU32 {
	return Vector_ImU32(p + BitVector(C.offsetof_ImBitVector_Storage))
}

// ChunkStream_ImGuiTableSettings wraps struct ImChunkStream_ImGuiTableSettings.
type ChunkStream_ImGuiTableSettings uintptr

// ChunkStream_ImGuiTableSettingsNil is a null pointer.
var ChunkStream_ImGuiTableSettingsNil ChunkStream_ImGuiTableSettings

// ChunkStream_ImGuiTableSettingsSizeOf is the byte size of ImChunkStream_ImGuiTableSettings.
const ChunkStream_ImGuiTableSettingsSizeOf = int(C.sizeof_ImChunkStream_ImGuiTableSettings)

// ChunkStream_ImGuiTableSettingsAlloc allocates a continuous block of ChunkStream_ImGuiTableSettings.
func ChunkStream_ImGuiTableSettingsAlloc(alloc ffi.Allocator, count int) ChunkStream_ImGuiTableSettings {
	ptr := alloc.Allocate(ChunkStream_ImGuiTableSettingsSizeOf * count)
	return ChunkStream_ImGuiTableSettings(ptr)
}

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p ChunkStream_ImGuiTableSettings) Offset(offset int) ChunkStream_ImGuiTableSettings {
	return p + ChunkStream_ImGuiTableSettings(offset*ChunkStream_ImGuiTableSettingsSizeOf)
}

// RefBuf returns pointer to the Buf field.
func (p ChunkStream_ImGuiTableSettings) RefBuf() Vector_char {
	return Vector_char(p + ChunkStream_ImGuiTableSettings(C.offsetof_ImChunkStream_ImGuiTableSettings_Buf))
}

// ChunkStream_ImGuiWindowSettings wraps struct ImChunkStream_ImGuiWindowSettings.
type ChunkStream_ImGuiWindowSettings uintptr

// ChunkStream_ImGuiWindowSettingsNil is a null pointer.
var ChunkStream_ImGuiWindowSettingsNil ChunkStream_ImGuiWindowSettings

// ChunkStream_ImGuiWindowSettingsSizeOf is the byte size of ImChunkStream_ImGuiWindowSettings.
const ChunkStream_ImGuiWindowSettingsSizeOf = int(C.sizeof_ImChunkStream_ImGuiWindowSettings)

// ChunkStream_ImGuiWindowSettingsAlloc allocates a continuous block of ChunkStream_ImGuiWindowSettings.
func ChunkStream_ImGuiWindowSettingsAlloc(alloc ffi.Allocator, count int) ChunkStream_ImGuiWindowSettings {
	ptr := alloc.Allocate(ChunkStream_ImGuiWindowSettingsSizeOf * count)
	return ChunkStream_ImGuiWindowSettings(ptr)
}

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p ChunkStream_ImGuiWindowSettings) Offset(offset int) ChunkStream_ImGuiWindowSettings {
	return p + ChunkStream_ImGuiWindowSettings(offset*ChunkStream_ImGuiWindowSettingsSizeOf)
}

// RefBuf returns pointer to the Buf field.
func (p ChunkStream_ImGuiWindowSettings) RefBuf() Vector_char {
	return Vector_char(p + ChunkStream_ImGuiWindowSettings(C.offsetof_ImChunkStream_ImGuiWindowSettings_Buf))
}

// Color wraps struct ImColor.
type Color uintptr

// ColorNil is a null pointer.
var ColorNil Color

// ColorSizeOf is the byte size of ImColor.
const ColorSizeOf = int(C.sizeof_ImColor)

// ImColor allocator is suppressed.

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p Color) Offset(offset int) Color {
	return p + Color(offset*ColorSizeOf)
}

// RefValue returns pointer to the Value field.
func (p Color) RefValue() Vec4 {
	return Vec4(p + Color(C.offsetof_ImColor_Value))
}

// DrawChannel wraps struct ImDrawChannel.
type DrawChannel uintptr

// DrawChannelNil is a null pointer.
var DrawChannelNil DrawChannel

// DrawChannelSizeOf is the byte size of ImDrawChannel.
const DrawChannelSizeOf = int(C.sizeof_ImDrawChannel)

// DrawChannelAlloc allocates a continuous block of DrawChannel.
func DrawChannelAlloc(alloc ffi.Allocator, count int) DrawChannel {
	ptr := alloc.Allocate(DrawChannelSizeOf * count)
	return DrawChannel(ptr)
}

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p DrawChannel) Offset(offset int) DrawChannel {
	return p + DrawChannel(offset*DrawChannelSizeOf)
}

// RefCmdBuffer returns pointer to the _CmdBuffer field.
func (p DrawChannel) RefCmdBuffer() Vector_ImDrawCmd {
	return Vector_ImDrawCmd(p + DrawChannel(C.offsetof_ImDrawChannel__CmdBuffer))
}

// RefIdxBuffer returns pointer to the _IdxBuffer field.
func (p DrawChannel) RefIdxBuffer() Vector_ImDrawIdx {
	return Vector_ImDrawIdx(p + DrawChannel(C.offsetof_ImDrawChannel__IdxBuffer))
}

// DrawCmd wraps struct ImDrawCmd.
type DrawCmd uintptr

// DrawCmdNil is a null pointer.
var DrawCmdNil DrawCmd

// DrawCmdSizeOf is the byte size of ImDrawCmd.
const DrawCmdSizeOf = int(C.sizeof_ImDrawCmd)

// ImDrawCmd allocator is suppressed.

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p DrawCmd) Offset(offset int) DrawCmd {
	return p + DrawCmd(offset*DrawCmdSizeOf)
}

// RefClipRect returns pointer to the ClipRect field.
func (p DrawCmd) RefClipRect() Vec4 {
	return Vec4(p + DrawCmd(C.offsetof_ImDrawCmd_ClipRect))
}

// GetElemCount returns the value in ElemCount.
func (p DrawCmd) GetElemCount() uint32 {
	ptr := (*C.ImDrawCmd)(unsafe.Pointer(p))
	return uint32(ptr.ElemCount)
}

// SetElemCount sets the value in ElemCount.
func (p DrawCmd) SetElemCount(value uint32) {
	ptr := (*C.ImDrawCmd)(unsafe.Pointer(p))
	ptr.ElemCount = (C.uint32_t)(value)
}

// GetIdxOffset returns the value in IdxOffset.
func (p DrawCmd) GetIdxOffset() uint32 {
	ptr := (*C.ImDrawCmd)(unsafe.Pointer(p))
	return uint32(ptr.IdxOffset)
}

// SetIdxOffset sets the value in IdxOffset.
func (p DrawCmd) SetIdxOffset(value uint32) {
	ptr := (*C.ImDrawCmd)(unsafe.Pointer(p))
	ptr.IdxOffset = (C.uint32_t)(value)
}

// RefTexRef returns pointer to the TexRef field.
func (p DrawCmd) RefTexRef() TextureRef {
	return TextureRef(p + DrawCmd(C.offsetof_ImDrawCmd_TexRef))
}

// DrawCmd.UserCallback is unsupported: category unsupported.

// GetUserCallbackData returns the value in UserCallbackData.
func (p DrawCmd) GetUserCallbackData() uintptr {
	ptr := (*C.ImDrawCmd)(unsafe.Pointer(p))
	return uintptr(unsafe.Pointer(ptr.UserCallbackData))
}

// SetUserCallbackData sets the value in UserCallbackData.
func (p DrawCmd) SetUserCallbackData(value uintptr) {
	ptr := (*C.ImDrawCmd)(unsafe.Pointer(p))
	ptr.UserCallbackData = unsafe.Pointer(value)
}

// GetUserCallbackDataOffset returns the value in UserCallbackDataOffset.
func (p DrawCmd) GetUserCallbackDataOffset() int32 {
	ptr := (*C.ImDrawCmd)(unsafe.Pointer(p))
	return int32(ptr.UserCallbackDataOffset)
}

// SetUserCallbackDataOffset sets the value in UserCallbackDataOffset.
func (p DrawCmd) SetUserCallbackDataOffset(value int32) {
	ptr := (*C.ImDrawCmd)(unsafe.Pointer(p))
	ptr.UserCallbackDataOffset = (C.int32_t)(value)
}

// GetUserCallbackDataSize returns the value in UserCallbackDataSize.
func (p DrawCmd) GetUserCallbackDataSize() int32 {
	ptr := (*C.ImDrawCmd)(unsafe.Pointer(p))
	return int32(ptr.UserCallbackDataSize)
}

// SetUserCallbackDataSize sets the value in UserCallbackDataSize.
func (p DrawCmd) SetUserCallbackDataSize(value int32) {
	ptr := (*C.ImDrawCmd)(unsafe.Pointer(p))
	ptr.UserCallbackDataSize = (C.int32_t)(value)
}

// GetVtxOffset returns the value in VtxOffset.
func (p DrawCmd) GetVtxOffset() uint32 {
	ptr := (*C.ImDrawCmd)(unsafe.Pointer(p))
	return uint32(ptr.VtxOffset)
}

// SetVtxOffset sets the value in VtxOffset.
func (p DrawCmd) SetVtxOffset(value uint32) {
	ptr := (*C.ImDrawCmd)(unsafe.Pointer(p))
	ptr.VtxOffset = (C.uint32_t)(value)
}

// DrawCmdHeader wraps struct ImDrawCmdHeader.
type DrawCmdHeader uintptr

// DrawCmdHeaderNil is a null pointer.
var DrawCmdHeaderNil DrawCmdHeader

// DrawCmdHeaderSizeOf is the byte size of ImDrawCmdHeader.
const DrawCmdHeaderSizeOf = int(C.sizeof_ImDrawCmdHeader)

// DrawCmdHeaderAlloc allocates a continuous block of DrawCmdHeader.
func DrawCmdHeaderAlloc(alloc ffi.Allocator, count int) DrawCmdHeader {
	ptr := alloc.Allocate(DrawCmdHeaderSizeOf * count)
	return DrawCmdHeader(ptr)
}

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p DrawCmdHeader) Offset(offset int) DrawCmdHeader {
	return p + DrawCmdHeader(offset*DrawCmdHeaderSizeOf)
}

// RefClipRect returns pointer to the ClipRect field.
func (p DrawCmdHeader) RefClipRect() Vec4 {
	return Vec4(p + DrawCmdHeader(C.offsetof_ImDrawCmdHeader_ClipRect))
}

// RefTexRef returns pointer to the TexRef field.
func (p DrawCmdHeader) RefTexRef() TextureRef {
	return TextureRef(p + DrawCmdHeader(C.offsetof_ImDrawCmdHeader_TexRef))
}

// GetVtxOffset returns the value in VtxOffset.
func (p DrawCmdHeader) GetVtxOffset() uint32 {
	ptr := (*C.ImDrawCmdHeader)(unsafe.Pointer(p))
	return uint32(ptr.VtxOffset)
}

// SetVtxOffset sets the value in VtxOffset.
func (p DrawCmdHeader) SetVtxOffset(value uint32) {
	ptr := (*C.ImDrawCmdHeader)(unsafe.Pointer(p))
	ptr.VtxOffset = (C.uint32_t)(value)
}

// DrawData wraps struct ImDrawData.
type DrawData uintptr

// DrawDataNil is a null pointer.
var DrawDataNil DrawData

// DrawDataSizeOf is the byte size of ImDrawData.
const DrawDataSizeOf = int(C.sizeof_ImDrawData)

// ImDrawData allocator is suppressed.

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p DrawData) Offset(offset int) DrawData {
	return p + DrawData(offset*DrawDataSizeOf)
}

// RefCmdLists returns pointer to the CmdLists field.
func (p DrawData) RefCmdLists() Vector_ImDrawListPtr {
	return Vector_ImDrawListPtr(p + DrawData(C.offsetof_ImDrawData_CmdLists))
}

// GetCmdListsCount returns the value in CmdListsCount.
func (p DrawData) GetCmdListsCount() int32 {
	ptr := (*C.ImDrawData)(unsafe.Pointer(p))
	return int32(ptr.CmdListsCount)
}

// SetCmdListsCount sets the value in CmdListsCount.
func (p DrawData) SetCmdListsCount(value int32) {
	ptr := (*C.ImDrawData)(unsafe.Pointer(p))
	ptr.CmdListsCount = (C.int32_t)(value)
}

// RefDisplayPos returns pointer to the DisplayPos field.
func (p DrawData) RefDisplayPos() Vec2 {
	return Vec2(p + DrawData(C.offsetof_ImDrawData_DisplayPos))
}

// RefDisplaySize returns pointer to the DisplaySize field.
func (p DrawData) RefDisplaySize() Vec2 {
	return Vec2(p + DrawData(C.offsetof_ImDrawData_DisplaySize))
}

// RefFramebufferScale returns pointer to the FramebufferScale field.
func (p DrawData) RefFramebufferScale() Vec2 {
	return Vec2(p + DrawData(C.offsetof_ImDrawData_FramebufferScale))
}

// GetOwnerViewport returns the value in OwnerViewport.
func (p DrawData) GetOwnerViewport() GuiViewport {
	ptr := (*C.ImDrawData)(unsafe.Pointer(p))
	return GuiViewport(unsafe.Pointer(ptr.OwnerViewport))
}

// SetOwnerViewport sets the value in OwnerViewport.
func (p DrawData) SetOwnerViewport(value GuiViewport) {
	ptr := (*C.ImDrawData)(unsafe.Pointer(p))
	ptr.OwnerViewport = (*C.ImGuiViewport)(unsafe.Pointer(value))
}

// GetTextures returns the value in Textures.
func (p DrawData) GetTextures() Vector_ImTextureDataPtr {
	ptr := (*C.ImDrawData)(unsafe.Pointer(p))
	return Vector_ImTextureDataPtr(unsafe.Pointer(ptr.Textures))
}

// SetTextures sets the value in Textures.
func (p DrawData) SetTextures(value Vector_ImTextureDataPtr) {
	ptr := (*C.ImDrawData)(unsafe.Pointer(p))
	ptr.Textures = (*C.ImVector_ImTextureDataPtr)(unsafe.Pointer(value))
}

// GetTotalIdxCount returns the value in TotalIdxCount.
func (p DrawData) GetTotalIdxCount() int32 {
	ptr := (*C.ImDrawData)(unsafe.Pointer(p))
	return int32(ptr.TotalIdxCount)
}

// SetTotalIdxCount sets the value in TotalIdxCount.
func (p DrawData) SetTotalIdxCount(value int32) {
	ptr := (*C.ImDrawData)(unsafe.Pointer(p))
	ptr.TotalIdxCount = (C.int32_t)(value)
}

// GetTotalVtxCount returns the value in TotalVtxCount.
func (p DrawData) GetTotalVtxCount() int32 {
	ptr := (*C.ImDrawData)(unsafe.Pointer(p))
	return int32(ptr.TotalVtxCount)
}

// SetTotalVtxCount sets the value in TotalVtxCount.
func (p DrawData) SetTotalVtxCount(value int32) {
	ptr := (*C.ImDrawData)(unsafe.Pointer(p))
	ptr.TotalVtxCount = (C.int32_t)(value)
}

// GetValid returns the value in Valid.
func (p DrawData) GetValid() bool {
	ptr := (*C.ImDrawData)(unsafe.Pointer(p))
	return bool(ptr.Valid)
}

// SetValid sets the value in Valid.
func (p DrawData) SetValid(value bool) {
	ptr := (*C.ImDrawData)(unsafe.Pointer(p))
	ptr.Valid = (C.bool)(value)
}

// DrawDataBuilder wraps struct ImDrawDataBuilder.
type DrawDataBuilder uintptr

// DrawDataBuilderNil is a null pointer.
var DrawDataBuilderNil DrawDataBuilder

// DrawDataBuilderSizeOf is the byte size of ImDrawDataBuilder.
const DrawDataBuilderSizeOf = int(C.sizeof_ImDrawDataBuilder)

// ImDrawDataBuilder allocator is suppressed.

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p DrawDataBuilder) Offset(offset int) DrawDataBuilder {
	return p + DrawDataBuilder(offset*DrawDataBuilderSizeOf)
}

// RefLayerData1 returns pointer to the LayerData1 field.
func (p DrawDataBuilder) RefLayerData1() Vector_ImDrawListPtr {
	return Vector_ImDrawListPtr(p + DrawDataBuilder(C.offsetof_ImDrawDataBuilder_LayerData1))
}

// DrawDataBuilder.Layers[2] is unsupported: category unsupported.

// DrawList wraps struct ImDrawList.
type DrawList uintptr

// DrawListNil is a null pointer.
var DrawListNil DrawList

// DrawListSizeOf is the byte size of ImDrawList.
const DrawListSizeOf = int(C.sizeof_ImDrawList)

// ImDrawList allocator is suppressed.

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p DrawList) Offset(offset int) DrawList {
	return p + DrawList(offset*DrawListSizeOf)
}

// RefCmdBuffer returns pointer to the CmdBuffer field.
func (p DrawList) RefCmdBuffer() Vector_ImDrawCmd {
	return Vector_ImDrawCmd(p + DrawList(C.offsetof_ImDrawList_CmdBuffer))
}

// GetFlags returns the value in Flags.
func (p DrawList) GetFlags() DrawListFlags {
	ptr := (*C.ImDrawList)(unsafe.Pointer(p))
	return DrawListFlags(ptr.Flags)
}

// SetFlags sets the value in Flags.
func (p DrawList) SetFlags(value DrawListFlags) {
	ptr := (*C.ImDrawList)(unsafe.Pointer(p))
	ptr.Flags = (C.ImDrawListFlags)(value)
}

// RefIdxBuffer returns pointer to the IdxBuffer field.
func (p DrawList) RefIdxBuffer() Vector_ImDrawIdx {
	return Vector_ImDrawIdx(p + DrawList(C.offsetof_ImDrawList_IdxBuffer))
}

// RefVtxBuffer returns pointer to the VtxBuffer field.
func (p DrawList) RefVtxBuffer() Vector_ImDrawVert {
	return Vector_ImDrawVert(p + DrawList(C.offsetof_ImDrawList_VtxBuffer))
}

// RefCallbacksDataBuf returns pointer to the _CallbacksDataBuf field.
func (p DrawList) RefCallbacksDataBuf() Vector_ImU8 {
	return Vector_ImU8(p + DrawList(C.offsetof_ImDrawList__CallbacksDataBuf))
}

// RefClipRectStack returns pointer to the _ClipRectStack field.
func (p DrawList) RefClipRectStack() Vector_ImVec4 {
	return Vector_ImVec4(p + DrawList(C.offsetof_ImDrawList__ClipRectStack))
}

// RefCmdHeader returns pointer to the _CmdHeader field.
func (p DrawList) RefCmdHeader() DrawCmdHeader {
	return DrawCmdHeader(p + DrawList(C.offsetof_ImDrawList__CmdHeader))
}

// GetData returns the value in _Data.
func (p DrawList) GetData() DrawListSharedData {
	ptr := (*C.ImDrawList)(unsafe.Pointer(p))
	return DrawListSharedData(unsafe.Pointer(ptr._Data))
}

// SetData sets the value in _Data.
func (p DrawList) SetData(value DrawListSharedData) {
	ptr := (*C.ImDrawList)(unsafe.Pointer(p))
	ptr._Data = (*C.ImDrawListSharedData)(unsafe.Pointer(value))
}

// GetFringeScale returns the value in _FringeScale.
func (p DrawList) GetFringeScale() float32 {
	ptr := (*C.ImDrawList)(unsafe.Pointer(p))
	return float32(ptr._FringeScale)
}

// SetFringeScale sets the value in _FringeScale.
func (p DrawList) SetFringeScale(value float32) {
	ptr := (*C.ImDrawList)(unsafe.Pointer(p))
	ptr._FringeScale = (C.float)(value)
}

// GetIdxWritePtr returns the value in _IdxWritePtr.
func (p DrawList) GetIdxWritePtr() ffi.Ref[DrawIdx] {
	ptr := (*C.ImDrawList)(unsafe.Pointer(p))
	return ffi.RefFromPtr[DrawIdx](unsafe.Pointer(ptr._IdxWritePtr))
}

// SetIdxWritePtr sets the value in _IdxWritePtr.
func (p DrawList) SetIdxWritePtr(value ffi.Ref[DrawIdx]) {
	ptr := (*C.ImDrawList)(unsafe.Pointer(p))
	ptr._IdxWritePtr = (*C.ImDrawIdx)(value.Raw())
}

// GetOwnerName returns the value in _OwnerName.
func (p DrawList) GetOwnerName() ffi.CString {
	ptr := (*C.ImDrawList)(unsafe.Pointer(p))
	return ffi.CStringFromPtr((unsafe.Pointer)(ptr._OwnerName))
}

// SetOwnerName sets the value in _OwnerName.
func (p DrawList) SetOwnerName(value ffi.CString) {
	ptr := (*C.ImDrawList)(unsafe.Pointer(p))
	ptr._OwnerName = (*C.char)(value.Raw())
}

// RefPath returns pointer to the _Path field.
func (p DrawList) RefPath() Vector_ImVec2 {
	return Vector_ImVec2(p + DrawList(C.offsetof_ImDrawList__Path))
}

// RefSplitter returns pointer to the _Splitter field.
func (p DrawList) RefSplitter() DrawListSplitter {
	return DrawListSplitter(p + DrawList(C.offsetof_ImDrawList__Splitter))
}

// RefTextureStack returns pointer to the _TextureStack field.
func (p DrawList) RefTextureStack() Vector_ImTextureRef {
	return Vector_ImTextureRef(p + DrawList(C.offsetof_ImDrawList__TextureStack))
}

// GetVtxCurrentIdx returns the value in _VtxCurrentIdx.
func (p DrawList) GetVtxCurrentIdx() uint32 {
	ptr := (*C.ImDrawList)(unsafe.Pointer(p))
	return uint32(ptr._VtxCurrentIdx)
}

// SetVtxCurrentIdx sets the value in _VtxCurrentIdx.
func (p DrawList) SetVtxCurrentIdx(value uint32) {
	ptr := (*C.ImDrawList)(unsafe.Pointer(p))
	ptr._VtxCurrentIdx = (C.uint32_t)(value)
}

// GetVtxWritePtr returns the value in _VtxWritePtr.
func (p DrawList) GetVtxWritePtr() DrawVert {
	ptr := (*C.ImDrawList)(unsafe.Pointer(p))
	return DrawVert(unsafe.Pointer(ptr._VtxWritePtr))
}

// SetVtxWritePtr sets the value in _VtxWritePtr.
func (p DrawList) SetVtxWritePtr(value DrawVert) {
	ptr := (*C.ImDrawList)(unsafe.Pointer(p))
	ptr._VtxWritePtr = (*C.ImDrawVert)(unsafe.Pointer(value))
}

// DrawListSharedData wraps struct ImDrawListSharedData.
type DrawListSharedData uintptr

// DrawListSharedDataNil is a null pointer.
var DrawListSharedDataNil DrawListSharedData

// DrawListSharedDataSizeOf is the byte size of ImDrawListSharedData.
const DrawListSharedDataSizeOf = int(C.sizeof_ImDrawListSharedData)

// ImDrawListSharedData allocator is suppressed.

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p DrawListSharedData) Offset(offset int) DrawListSharedData {
	return p + DrawListSharedData(offset*DrawListSharedDataSizeOf)
}

// GetArcFastRadiusCutoff returns the value in ArcFastRadiusCutoff.
func (p DrawListSharedData) GetArcFastRadiusCutoff() float32 {
	ptr := (*C.ImDrawListSharedData)(unsafe.Pointer(p))
	return float32(ptr.ArcFastRadiusCutoff)
}

// SetArcFastRadiusCutoff sets the value in ArcFastRadiusCutoff.
func (p DrawListSharedData) SetArcFastRadiusCutoff(value float32) {
	ptr := (*C.ImDrawListSharedData)(unsafe.Pointer(p))
	ptr.ArcFastRadiusCutoff = (C.float)(value)
}

// DrawListSharedData.ArcFastVtx[48] is unsupported: category unsupported.

// DrawListSharedData.CircleSegmentCounts[64] is unsupported: category unsupported.

// GetCircleSegmentMaxError returns the value in CircleSegmentMaxError.
func (p DrawListSharedData) GetCircleSegmentMaxError() float32 {
	ptr := (*C.ImDrawListSharedData)(unsafe.Pointer(p))
	return float32(ptr.CircleSegmentMaxError)
}

// SetCircleSegmentMaxError sets the value in CircleSegmentMaxError.
func (p DrawListSharedData) SetCircleSegmentMaxError(value float32) {
	ptr := (*C.ImDrawListSharedData)(unsafe.Pointer(p))
	ptr.CircleSegmentMaxError = (C.float)(value)
}

// RefClipRectFullscreen returns pointer to the ClipRectFullscreen field.
func (p DrawListSharedData) RefClipRectFullscreen() Vec4 {
	return Vec4(p + DrawListSharedData(C.offsetof_ImDrawListSharedData_ClipRectFullscreen))
}

// GetContext returns the value in Context.
func (p DrawListSharedData) GetContext() GuiContext {
	ptr := (*C.ImDrawListSharedData)(unsafe.Pointer(p))
	return GuiContext(unsafe.Pointer(ptr.Context))
}

// SetContext sets the value in Context.
func (p DrawListSharedData) SetContext(value GuiContext) {
	ptr := (*C.ImDrawListSharedData)(unsafe.Pointer(p))
	ptr.Context = (*C.ImGuiContext)(unsafe.Pointer(value))
}

// GetCurveTessellationTol returns the value in CurveTessellationTol.
func (p DrawListSharedData) GetCurveTessellationTol() float32 {
	ptr := (*C.ImDrawListSharedData)(unsafe.Pointer(p))
	return float32(ptr.CurveTessellationTol)
}

// SetCurveTessellationTol sets the value in CurveTessellationTol.
func (p DrawListSharedData) SetCurveTessellationTol(value float32) {
	ptr := (*C.ImDrawListSharedData)(unsafe.Pointer(p))
	ptr.CurveTessellationTol = (C.float)(value)
}

// RefDrawLists returns pointer to the DrawLists field.
func (p DrawListSharedData) RefDrawLists() Vector_ImDrawListPtr {
	return Vector_ImDrawListPtr(p + DrawListSharedData(C.offsetof_ImDrawListSharedData_DrawLists))
}

// GetFont returns the value in Font.
func (p DrawListSharedData) GetFont() Font {
	ptr := (*C.ImDrawListSharedData)(unsafe.Pointer(p))
	return Font(unsafe.Pointer(ptr.Font))
}

// SetFont sets the value in Font.
func (p DrawListSharedData) SetFont(value Font) {
	ptr := (*C.ImDrawListSharedData)(unsafe.Pointer(p))
	ptr.Font = (*C.ImFont)(unsafe.Pointer(value))
}

// GetFontAtlas returns the value in FontAtlas.
func (p DrawListSharedData) GetFontAtlas() FontAtlas {
	ptr := (*C.ImDrawListSharedData)(unsafe.Pointer(p))
	return FontAtlas(unsafe.Pointer(ptr.FontAtlas))
}

// SetFontAtlas sets the value in FontAtlas.
func (p DrawListSharedData) SetFontAtlas(value FontAtlas) {
	ptr := (*C.ImDrawListSharedData)(unsafe.Pointer(p))
	ptr.FontAtlas = (*C.ImFontAtlas)(unsafe.Pointer(value))
}

// GetFontScale returns the value in FontScale.
func (p DrawListSharedData) GetFontScale() float32 {
	ptr := (*C.ImDrawListSharedData)(unsafe.Pointer(p))
	return float32(ptr.FontScale)
}

// SetFontScale sets the value in FontScale.
func (p DrawListSharedData) SetFontScale(value float32) {
	ptr := (*C.ImDrawListSharedData)(unsafe.Pointer(p))
	ptr.FontScale = (C.float)(value)
}

// GetFontSize returns the value in FontSize.
func (p DrawListSharedData) GetFontSize() float32 {
	ptr := (*C.ImDrawListSharedData)(unsafe.Pointer(p))
	return float32(ptr.FontSize)
}

// SetFontSize sets the value in FontSize.
func (p DrawListSharedData) SetFontSize(value float32) {
	ptr := (*C.ImDrawListSharedData)(unsafe.Pointer(p))
	ptr.FontSize = (C.float)(value)
}

// GetInitialFlags returns the value in InitialFlags.
func (p DrawListSharedData) GetInitialFlags() DrawListFlags {
	ptr := (*C.ImDrawListSharedData)(unsafe.Pointer(p))
	return DrawListFlags(ptr.InitialFlags)
}

// SetInitialFlags sets the value in InitialFlags.
func (p DrawListSharedData) SetInitialFlags(value DrawListFlags) {
	ptr := (*C.ImDrawListSharedData)(unsafe.Pointer(p))
	ptr.InitialFlags = (C.ImDrawListFlags)(value)
}

// GetInitialFringeScale returns the value in InitialFringeScale.
func (p DrawListSharedData) GetInitialFringeScale() float32 {
	ptr := (*C.ImDrawListSharedData)(unsafe.Pointer(p))
	return float32(ptr.InitialFringeScale)
}

// SetInitialFringeScale sets the value in InitialFringeScale.
func (p DrawListSharedData) SetInitialFringeScale(value float32) {
	ptr := (*C.ImDrawListSharedData)(unsafe.Pointer(p))
	ptr.InitialFringeScale = (C.float)(value)
}

// RefTempBuffer returns pointer to the TempBuffer field.
func (p DrawListSharedData) RefTempBuffer() Vector_ImVec2 {
	return Vector_ImVec2(p + DrawListSharedData(C.offsetof_ImDrawListSharedData_TempBuffer))
}

// GetTexUvLines returns the value in TexUvLines.
func (p DrawListSharedData) GetTexUvLines() Vec4 {
	ptr := (*C.ImDrawListSharedData)(unsafe.Pointer(p))
	return Vec4(unsafe.Pointer(ptr.TexUvLines))
}

// SetTexUvLines sets the value in TexUvLines.
func (p DrawListSharedData) SetTexUvLines(value Vec4) {
	ptr := (*C.ImDrawListSharedData)(unsafe.Pointer(p))
	ptr.TexUvLines = (*C.ImVec4)(unsafe.Pointer(value))
}

// RefTexUvWhitePixel returns pointer to the TexUvWhitePixel field.
func (p DrawListSharedData) RefTexUvWhitePixel() Vec2 {
	return Vec2(p + DrawListSharedData(C.offsetof_ImDrawListSharedData_TexUvWhitePixel))
}

// DrawListSplitter wraps struct ImDrawListSplitter.
type DrawListSplitter uintptr

// DrawListSplitterNil is a null pointer.
var DrawListSplitterNil DrawListSplitter

// DrawListSplitterSizeOf is the byte size of ImDrawListSplitter.
const DrawListSplitterSizeOf = int(C.sizeof_ImDrawListSplitter)

// ImDrawListSplitter allocator is suppressed.

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p DrawListSplitter) Offset(offset int) DrawListSplitter {
	return p + DrawListSplitter(offset*DrawListSplitterSizeOf)
}

// RefChannels returns pointer to the _Channels field.
func (p DrawListSplitter) RefChannels() Vector_ImDrawChannel {
	return Vector_ImDrawChannel(p + DrawListSplitter(C.offsetof_ImDrawListSplitter__Channels))
}

// GetCount returns the value in _Count.
func (p DrawListSplitter) GetCount() int32 {
	ptr := (*C.ImDrawListSplitter)(unsafe.Pointer(p))
	return int32(ptr._Count)
}

// SetCount sets the value in _Count.
func (p DrawListSplitter) SetCount(value int32) {
	ptr := (*C.ImDrawListSplitter)(unsafe.Pointer(p))
	ptr._Count = (C.int32_t)(value)
}

// GetCurrent returns the value in _Current.
func (p DrawListSplitter) GetCurrent() int32 {
	ptr := (*C.ImDrawListSplitter)(unsafe.Pointer(p))
	return int32(ptr._Current)
}

// SetCurrent sets the value in _Current.
func (p DrawListSplitter) SetCurrent(value int32) {
	ptr := (*C.ImDrawListSplitter)(unsafe.Pointer(p))
	ptr._Current = (C.int32_t)(value)
}

// DrawVert wraps struct ImDrawVert.
type DrawVert uintptr

// DrawVertNil is a null pointer.
var DrawVertNil DrawVert

// DrawVertSizeOf is the byte size of ImDrawVert.
const DrawVertSizeOf = int(C.sizeof_ImDrawVert)

// DrawVertAlloc allocates a continuous block of DrawVert.
func DrawVertAlloc(alloc ffi.Allocator, count int) DrawVert {
	ptr := alloc.Allocate(DrawVertSizeOf * count)
	return DrawVert(ptr)
}

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p DrawVert) Offset(offset int) DrawVert {
	return p + DrawVert(offset*DrawVertSizeOf)
}

// GetCol returns the value in col.
func (p DrawVert) GetCol() uint32 {
	ptr := (*C.ImDrawVert)(unsafe.Pointer(p))
	return uint32(ptr.col)
}

// SetCol sets the value in col.
func (p DrawVert) SetCol(value uint32) {
	ptr := (*C.ImDrawVert)(unsafe.Pointer(p))
	ptr.col = (C.uint32_t)(value)
}

// RefPos returns pointer to the pos field.
func (p DrawVert) RefPos() Vec2 {
	return Vec2(p + DrawVert(C.offsetof_ImDrawVert_pos))
}

// RefUv returns pointer to the uv field.
func (p DrawVert) RefUv() Vec2 {
	return Vec2(p + DrawVert(C.offsetof_ImDrawVert_uv))
}

// Font wraps struct ImFont.
type Font uintptr

// FontNil is a null pointer.
var FontNil Font

// FontSizeOf is the byte size of ImFont.
const FontSizeOf = int(C.sizeof_ImFont)

// ImFont allocator is suppressed.

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p Font) Offset(offset int) Font {
	return p + Font(offset*FontSizeOf)
}

// GetContainerAtlas returns the value in ContainerAtlas.
func (p Font) GetContainerAtlas() FontAtlas {
	ptr := (*C.ImFont)(unsafe.Pointer(p))
	return FontAtlas(unsafe.Pointer(ptr.ContainerAtlas))
}

// SetContainerAtlas sets the value in ContainerAtlas.
func (p Font) SetContainerAtlas(value FontAtlas) {
	ptr := (*C.ImFont)(unsafe.Pointer(p))
	ptr.ContainerAtlas = (*C.ImFontAtlas)(unsafe.Pointer(value))
}

// GetCurrentRasterizerDensity returns the value in CurrentRasterizerDensity.
func (p Font) GetCurrentRasterizerDensity() float32 {
	ptr := (*C.ImFont)(unsafe.Pointer(p))
	return float32(ptr.CurrentRasterizerDensity)
}

// SetCurrentRasterizerDensity sets the value in CurrentRasterizerDensity.
func (p Font) SetCurrentRasterizerDensity(value float32) {
	ptr := (*C.ImFont)(unsafe.Pointer(p))
	ptr.CurrentRasterizerDensity = (C.float)(value)
}

// GetEllipsisAutoBake returns the value in EllipsisAutoBake.
func (p Font) GetEllipsisAutoBake() bool {
	ptr := (*C.ImFont)(unsafe.Pointer(p))
	return bool(ptr.EllipsisAutoBake)
}

// SetEllipsisAutoBake sets the value in EllipsisAutoBake.
func (p Font) SetEllipsisAutoBake(value bool) {
	ptr := (*C.ImFont)(unsafe.Pointer(p))
	ptr.EllipsisAutoBake = (C.bool)(value)
}

// Font.EllipsisChar is unsupported: category unsupported.

// Font.FallbackChar is unsupported: category unsupported.

// GetFlags returns the value in Flags.
func (p Font) GetFlags() FontFlags {
	ptr := (*C.ImFont)(unsafe.Pointer(p))
	return FontFlags(ptr.Flags)
}

// SetFlags sets the value in Flags.
func (p Font) SetFlags(value FontFlags) {
	ptr := (*C.ImFont)(unsafe.Pointer(p))
	ptr.Flags = (C.ImFontFlags)(value)
}

// Font.FontId is unsupported: category unsupported.

// GetLastBaked returns the value in LastBaked.
func (p Font) GetLastBaked() FontBaked {
	ptr := (*C.ImFont)(unsafe.Pointer(p))
	return FontBaked(unsafe.Pointer(ptr.LastBaked))
}

// SetLastBaked sets the value in LastBaked.
func (p Font) SetLastBaked(value FontBaked) {
	ptr := (*C.ImFont)(unsafe.Pointer(p))
	ptr.LastBaked = (*C.ImFontBaked)(unsafe.Pointer(value))
}

// GetLegacySize returns the value in LegacySize.
func (p Font) GetLegacySize() float32 {
	ptr := (*C.ImFont)(unsafe.Pointer(p))
	return float32(ptr.LegacySize)
}

// SetLegacySize sets the value in LegacySize.
func (p Font) SetLegacySize(value float32) {
	ptr := (*C.ImFont)(unsafe.Pointer(p))
	ptr.LegacySize = (C.float)(value)
}

// RefRemapPairs returns pointer to the RemapPairs field.
func (p Font) RefRemapPairs() GuiStorage {
	return GuiStorage(p + Font(C.offsetof_ImFont_RemapPairs))
}

// RefSources returns pointer to the Sources field.
func (p Font) RefSources() Vector_ImFontConfigPtr {
	return Vector_ImFontConfigPtr(p + Font(C.offsetof_ImFont_Sources))
}

// Font.Used8kPagesMap[(0xFFFF+1)/8192/8] is unsupported: category unsupported.

// FontAtlas wraps struct ImFontAtlas.
type FontAtlas uintptr

// FontAtlasNil is a null pointer.
var FontAtlasNil FontAtlas

// FontAtlasSizeOf is the byte size of ImFontAtlas.
const FontAtlasSizeOf = int(C.sizeof_ImFontAtlas)

// ImFontAtlas allocator is suppressed.

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p FontAtlas) Offset(offset int) FontAtlas {
	return p + FontAtlas(offset*FontAtlasSizeOf)
}

// GetBuilder returns the value in Builder.
func (p FontAtlas) GetBuilder() FontAtlasBuilder {
	ptr := (*C.ImFontAtlas)(unsafe.Pointer(p))
	return FontAtlasBuilder(unsafe.Pointer(ptr.Builder))
}

// SetBuilder sets the value in Builder.
func (p FontAtlas) SetBuilder(value FontAtlasBuilder) {
	ptr := (*C.ImFontAtlas)(unsafe.Pointer(p))
	ptr.Builder = (*C.ImFontAtlasBuilder)(unsafe.Pointer(value))
}

// RefDrawListSharedDatas returns pointer to the DrawListSharedDatas field.
func (p FontAtlas) RefDrawListSharedDatas() Vector_ImDrawListSharedDataPtr {
	return Vector_ImDrawListSharedDataPtr(p + FontAtlas(C.offsetof_ImFontAtlas_DrawListSharedDatas))
}

// GetFlags returns the value in Flags.
func (p FontAtlas) GetFlags() FontAtlasFlags {
	ptr := (*C.ImFontAtlas)(unsafe.Pointer(p))
	return FontAtlasFlags(ptr.Flags)
}

// SetFlags sets the value in Flags.
func (p FontAtlas) SetFlags(value FontAtlasFlags) {
	ptr := (*C.ImFontAtlas)(unsafe.Pointer(p))
	ptr.Flags = (C.ImFontAtlasFlags)(value)
}

// GetFontLoader returns the value in FontLoader.
func (p FontAtlas) GetFontLoader() FontLoader {
	ptr := (*C.ImFontAtlas)(unsafe.Pointer(p))
	return FontLoader(unsafe.Pointer(ptr.FontLoader))
}

// FontAtlas.FontLoader getter is suppressed.

// GetFontLoaderData returns the value in FontLoaderData.
func (p FontAtlas) GetFontLoaderData() uintptr {
	ptr := (*C.ImFontAtlas)(unsafe.Pointer(p))
	return uintptr(unsafe.Pointer(ptr.FontLoaderData))
}

// SetFontLoaderData sets the value in FontLoaderData.
func (p FontAtlas) SetFontLoaderData(value uintptr) {
	ptr := (*C.ImFontAtlas)(unsafe.Pointer(p))
	ptr.FontLoaderData = unsafe.Pointer(value)
}

// GetFontLoaderFlags returns the value in FontLoaderFlags.
func (p FontAtlas) GetFontLoaderFlags() uint32 {
	ptr := (*C.ImFontAtlas)(unsafe.Pointer(p))
	return uint32(ptr.FontLoaderFlags)
}

// SetFontLoaderFlags sets the value in FontLoaderFlags.
func (p FontAtlas) SetFontLoaderFlags(value uint32) {
	ptr := (*C.ImFontAtlas)(unsafe.Pointer(p))
	ptr.FontLoaderFlags = (C.uint32_t)(value)
}

// GetFontLoaderName returns the value in FontLoaderName.
func (p FontAtlas) GetFontLoaderName() ffi.CString {
	ptr := (*C.ImFontAtlas)(unsafe.Pointer(p))
	return ffi.CStringFromPtr((unsafe.Pointer)(ptr.FontLoaderName))
}

// SetFontLoaderName sets the value in FontLoaderName.
func (p FontAtlas) SetFontLoaderName(value ffi.CString) {
	ptr := (*C.ImFontAtlas)(unsafe.Pointer(p))
	ptr.FontLoaderName = (*C.char)(value.Raw())
}

// GetFontNextUniqueID returns the value in FontNextUniqueID.
func (p FontAtlas) GetFontNextUniqueID() int32 {
	ptr := (*C.ImFontAtlas)(unsafe.Pointer(p))
	return int32(ptr.FontNextUniqueID)
}

// SetFontNextUniqueID sets the value in FontNextUniqueID.
func (p FontAtlas) SetFontNextUniqueID(value int32) {
	ptr := (*C.ImFontAtlas)(unsafe.Pointer(p))
	ptr.FontNextUniqueID = (C.int32_t)(value)
}

// RefFonts returns pointer to the Fonts field.
func (p FontAtlas) RefFonts() Vector_ImFontPtr {
	return Vector_ImFontPtr(p + FontAtlas(C.offsetof_ImFontAtlas_Fonts))
}

// GetLocked returns the value in Locked.
func (p FontAtlas) GetLocked() bool {
	ptr := (*C.ImFontAtlas)(unsafe.Pointer(p))
	return bool(ptr.Locked)
}

// SetLocked sets the value in Locked.
func (p FontAtlas) SetLocked(value bool) {
	ptr := (*C.ImFontAtlas)(unsafe.Pointer(p))
	ptr.Locked = (C.bool)(value)
}

// GetOwnerContext returns the value in OwnerContext.
func (p FontAtlas) GetOwnerContext() GuiContext {
	ptr := (*C.ImFontAtlas)(unsafe.Pointer(p))
	return GuiContext(unsafe.Pointer(ptr.OwnerContext))
}

// SetOwnerContext sets the value in OwnerContext.
func (p FontAtlas) SetOwnerContext(value GuiContext) {
	ptr := (*C.ImFontAtlas)(unsafe.Pointer(p))
	ptr.OwnerContext = (*C.ImGuiContext)(unsafe.Pointer(value))
}

// GetRefCount returns the value in RefCount.
func (p FontAtlas) GetRefCount() int32 {
	ptr := (*C.ImFontAtlas)(unsafe.Pointer(p))
	return int32(ptr.RefCount)
}

// SetRefCount sets the value in RefCount.
func (p FontAtlas) SetRefCount(value int32) {
	ptr := (*C.ImFontAtlas)(unsafe.Pointer(p))
	ptr.RefCount = (C.int32_t)(value)
}

// GetRendererHasTextures returns the value in RendererHasTextures.
func (p FontAtlas) GetRendererHasTextures() bool {
	ptr := (*C.ImFontAtlas)(unsafe.Pointer(p))
	return bool(ptr.RendererHasTextures)
}

// SetRendererHasTextures sets the value in RendererHasTextures.
func (p FontAtlas) SetRendererHasTextures(value bool) {
	ptr := (*C.ImFontAtlas)(unsafe.Pointer(p))
	ptr.RendererHasTextures = (C.bool)(value)
}

// RefSources returns pointer to the Sources field.
func (p FontAtlas) RefSources() Vector_ImFontConfig {
	return Vector_ImFontConfig(p + FontAtlas(C.offsetof_ImFontAtlas_Sources))
}

// GetTexData returns the value in TexData.
func (p FontAtlas) GetTexData() TextureData {
	ptr := (*C.ImFontAtlas)(unsafe.Pointer(p))
	return TextureData(unsafe.Pointer(ptr.TexData))
}

// SetTexData sets the value in TexData.
func (p FontAtlas) SetTexData(value TextureData) {
	ptr := (*C.ImFontAtlas)(unsafe.Pointer(p))
	ptr.TexData = (*C.ImTextureData)(unsafe.Pointer(value))
}

// GetTexDesiredFormat returns the value in TexDesiredFormat.
func (p FontAtlas) GetTexDesiredFormat() TextureFormat {
	ptr := (*C.ImFontAtlas)(unsafe.Pointer(p))
	return TextureFormat(ptr.TexDesiredFormat)
}

// SetTexDesiredFormat sets the value in TexDesiredFormat.
func (p FontAtlas) SetTexDesiredFormat(value TextureFormat) {
	ptr := (*C.ImFontAtlas)(unsafe.Pointer(p))
	ptr.TexDesiredFormat = (C.ImTextureFormat)(value)
}

// GetTexGlyphPadding returns the value in TexGlyphPadding.
func (p FontAtlas) GetTexGlyphPadding() int32 {
	ptr := (*C.ImFontAtlas)(unsafe.Pointer(p))
	return int32(ptr.TexGlyphPadding)
}

// SetTexGlyphPadding sets the value in TexGlyphPadding.
func (p FontAtlas) SetTexGlyphPadding(value int32) {
	ptr := (*C.ImFontAtlas)(unsafe.Pointer(p))
	ptr.TexGlyphPadding = (C.int32_t)(value)
}

// GetTexIsBuilt returns the value in TexIsBuilt.
func (p FontAtlas) GetTexIsBuilt() bool {
	ptr := (*C.ImFontAtlas)(unsafe.Pointer(p))
	return bool(ptr.TexIsBuilt)
}

// SetTexIsBuilt sets the value in TexIsBuilt.
func (p FontAtlas) SetTexIsBuilt(value bool) {
	ptr := (*C.ImFontAtlas)(unsafe.Pointer(p))
	ptr.TexIsBuilt = (C.bool)(value)
}

// RefTexList returns pointer to the TexList field.
func (p FontAtlas) RefTexList() Vector_ImTextureDataPtr {
	return Vector_ImTextureDataPtr(p + FontAtlas(C.offsetof_ImFontAtlas_TexList))
}

// GetTexMaxHeight returns the value in TexMaxHeight.
func (p FontAtlas) GetTexMaxHeight() int32 {
	ptr := (*C.ImFontAtlas)(unsafe.Pointer(p))
	return int32(ptr.TexMaxHeight)
}

// SetTexMaxHeight sets the value in TexMaxHeight.
func (p FontAtlas) SetTexMaxHeight(value int32) {
	ptr := (*C.ImFontAtlas)(unsafe.Pointer(p))
	ptr.TexMaxHeight = (C.int32_t)(value)
}

// GetTexMaxWidth returns the value in TexMaxWidth.
func (p FontAtlas) GetTexMaxWidth() int32 {
	ptr := (*C.ImFontAtlas)(unsafe.Pointer(p))
	return int32(ptr.TexMaxWidth)
}

// SetTexMaxWidth sets the value in TexMaxWidth.
func (p FontAtlas) SetTexMaxWidth(value int32) {
	ptr := (*C.ImFontAtlas)(unsafe.Pointer(p))
	ptr.TexMaxWidth = (C.int32_t)(value)
}

// GetTexMinHeight returns the value in TexMinHeight.
func (p FontAtlas) GetTexMinHeight() int32 {
	ptr := (*C.ImFontAtlas)(unsafe.Pointer(p))
	return int32(ptr.TexMinHeight)
}

// SetTexMinHeight sets the value in TexMinHeight.
func (p FontAtlas) SetTexMinHeight(value int32) {
	ptr := (*C.ImFontAtlas)(unsafe.Pointer(p))
	ptr.TexMinHeight = (C.int32_t)(value)
}

// GetTexMinWidth returns the value in TexMinWidth.
func (p FontAtlas) GetTexMinWidth() int32 {
	ptr := (*C.ImFontAtlas)(unsafe.Pointer(p))
	return int32(ptr.TexMinWidth)
}

// SetTexMinWidth sets the value in TexMinWidth.
func (p FontAtlas) SetTexMinWidth(value int32) {
	ptr := (*C.ImFontAtlas)(unsafe.Pointer(p))
	ptr.TexMinWidth = (C.int32_t)(value)
}

// GetTexNextUniqueID returns the value in TexNextUniqueID.
func (p FontAtlas) GetTexNextUniqueID() int32 {
	ptr := (*C.ImFontAtlas)(unsafe.Pointer(p))
	return int32(ptr.TexNextUniqueID)
}

// SetTexNextUniqueID sets the value in TexNextUniqueID.
func (p FontAtlas) SetTexNextUniqueID(value int32) {
	ptr := (*C.ImFontAtlas)(unsafe.Pointer(p))
	ptr.TexNextUniqueID = (C.int32_t)(value)
}

// GetTexPixelsUseColors returns the value in TexPixelsUseColors.
func (p FontAtlas) GetTexPixelsUseColors() bool {
	ptr := (*C.ImFontAtlas)(unsafe.Pointer(p))
	return bool(ptr.TexPixelsUseColors)
}

// SetTexPixelsUseColors sets the value in TexPixelsUseColors.
func (p FontAtlas) SetTexPixelsUseColors(value bool) {
	ptr := (*C.ImFontAtlas)(unsafe.Pointer(p))
	ptr.TexPixelsUseColors = (C.bool)(value)
}

// RefTexRef returns pointer to the TexRef field.
func (p FontAtlas) RefTexRef() TextureRef {
	return TextureRef(p + FontAtlas(C.offsetof_ImFontAtlas_TexRef))
}

// FontAtlas.TexUvLines[(32)+1] is unsupported: category unsupported.

// RefTexUvScale returns pointer to the TexUvScale field.
func (p FontAtlas) RefTexUvScale() Vec2 {
	return Vec2(p + FontAtlas(C.offsetof_ImFontAtlas_TexUvScale))
}

// RefTexUvWhitePixel returns pointer to the TexUvWhitePixel field.
func (p FontAtlas) RefTexUvWhitePixel() Vec2 {
	return Vec2(p + FontAtlas(C.offsetof_ImFontAtlas_TexUvWhitePixel))
}

// GetUserData returns the value in UserData.
func (p FontAtlas) GetUserData() uintptr {
	ptr := (*C.ImFontAtlas)(unsafe.Pointer(p))
	return uintptr(unsafe.Pointer(ptr.UserData))
}

// SetUserData sets the value in UserData.
func (p FontAtlas) SetUserData(value uintptr) {
	ptr := (*C.ImFontAtlas)(unsafe.Pointer(p))
	ptr.UserData = unsafe.Pointer(value)
}

// FontAtlasBuilder wraps struct ImFontAtlasBuilder.
type FontAtlasBuilder uintptr

// FontAtlasBuilderNil is a null pointer.
var FontAtlasBuilderNil FontAtlasBuilder

// FontAtlasBuilderSizeOf is the byte size of ImFontAtlasBuilder.
const FontAtlasBuilderSizeOf = int(C.sizeof_ImFontAtlasBuilder)

// ImFontAtlasBuilder allocator is suppressed.

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p FontAtlasBuilder) Offset(offset int) FontAtlasBuilder {
	return p + FontAtlasBuilder(offset*FontAtlasBuilderSizeOf)
}

// GetBakedDiscardedCount returns the value in BakedDiscardedCount.
func (p FontAtlasBuilder) GetBakedDiscardedCount() int32 {
	ptr := (*C.ImFontAtlasBuilder)(unsafe.Pointer(p))
	return int32(ptr.BakedDiscardedCount)
}

// SetBakedDiscardedCount sets the value in BakedDiscardedCount.
func (p FontAtlasBuilder) SetBakedDiscardedCount(value int32) {
	ptr := (*C.ImFontAtlasBuilder)(unsafe.Pointer(p))
	ptr.BakedDiscardedCount = (C.int32_t)(value)
}

// RefBakedMap returns pointer to the BakedMap field.
func (p FontAtlasBuilder) RefBakedMap() GuiStorage {
	return GuiStorage(p + FontAtlasBuilder(C.offsetof_ImFontAtlasBuilder_BakedMap))
}

// FontAtlasBuilder.BakedPool is unsupported: category unsupported.

// GetFrameCount returns the value in FrameCount.
func (p FontAtlasBuilder) GetFrameCount() int32 {
	ptr := (*C.ImFontAtlasBuilder)(unsafe.Pointer(p))
	return int32(ptr.FrameCount)
}

// SetFrameCount sets the value in FrameCount.
func (p FontAtlasBuilder) SetFrameCount(value int32) {
	ptr := (*C.ImFontAtlasBuilder)(unsafe.Pointer(p))
	ptr.FrameCount = (C.int32_t)(value)
}

// GetLockDisableResize returns the value in LockDisableResize.
func (p FontAtlasBuilder) GetLockDisableResize() bool {
	ptr := (*C.ImFontAtlasBuilder)(unsafe.Pointer(p))
	return bool(ptr.LockDisableResize)
}

// SetLockDisableResize sets the value in LockDisableResize.
func (p FontAtlasBuilder) SetLockDisableResize(value bool) {
	ptr := (*C.ImFontAtlasBuilder)(unsafe.Pointer(p))
	ptr.LockDisableResize = (C.bool)(value)
}

// RefMaxRectBounds returns pointer to the MaxRectBounds field.
func (p FontAtlasBuilder) RefMaxRectBounds() Vec2i {
	return Vec2i(p + FontAtlasBuilder(C.offsetof_ImFontAtlasBuilder_MaxRectBounds))
}

// RefMaxRectSize returns pointer to the MaxRectSize field.
func (p FontAtlasBuilder) RefMaxRectSize() Vec2i {
	return Vec2i(p + FontAtlasBuilder(C.offsetof_ImFontAtlasBuilder_MaxRectSize))
}

// FontAtlasBuilder.PackContext is unsupported: unknown type stbrp_context_opaque.

// GetPackIdLinesTexData returns the value in PackIdLinesTexData.
func (p FontAtlasBuilder) GetPackIdLinesTexData() FontAtlasRectId {
	ptr := (*C.ImFontAtlasBuilder)(unsafe.Pointer(p))
	return FontAtlasRectId(ptr.PackIdLinesTexData)
}

// SetPackIdLinesTexData sets the value in PackIdLinesTexData.
func (p FontAtlasBuilder) SetPackIdLinesTexData(value FontAtlasRectId) {
	ptr := (*C.ImFontAtlasBuilder)(unsafe.Pointer(p))
	ptr.PackIdLinesTexData = (C.ImFontAtlasRectId)(value)
}

// GetPackIdMouseCursors returns the value in PackIdMouseCursors.
func (p FontAtlasBuilder) GetPackIdMouseCursors() FontAtlasRectId {
	ptr := (*C.ImFontAtlasBuilder)(unsafe.Pointer(p))
	return FontAtlasRectId(ptr.PackIdMouseCursors)
}

// SetPackIdMouseCursors sets the value in PackIdMouseCursors.
func (p FontAtlasBuilder) SetPackIdMouseCursors(value FontAtlasRectId) {
	ptr := (*C.ImFontAtlasBuilder)(unsafe.Pointer(p))
	ptr.PackIdMouseCursors = (C.ImFontAtlasRectId)(value)
}

// RefPackNodes returns pointer to the PackNodes field.
func (p FontAtlasBuilder) RefPackNodes() Vector_stbrp_node_im {
	return Vector_stbrp_node_im(p + FontAtlasBuilder(C.offsetof_ImFontAtlasBuilder_PackNodes))
}

// GetPreloadedAllGlyphsRanges returns the value in PreloadedAllGlyphsRanges.
func (p FontAtlasBuilder) GetPreloadedAllGlyphsRanges() bool {
	ptr := (*C.ImFontAtlasBuilder)(unsafe.Pointer(p))
	return bool(ptr.PreloadedAllGlyphsRanges)
}

// SetPreloadedAllGlyphsRanges sets the value in PreloadedAllGlyphsRanges.
func (p FontAtlasBuilder) SetPreloadedAllGlyphsRanges(value bool) {
	ptr := (*C.ImFontAtlasBuilder)(unsafe.Pointer(p))
	ptr.PreloadedAllGlyphsRanges = (C.bool)(value)
}

// RefRects returns pointer to the Rects field.
func (p FontAtlasBuilder) RefRects() Vector_ImTextureRect {
	return Vector_ImTextureRect(p + FontAtlasBuilder(C.offsetof_ImFontAtlasBuilder_Rects))
}

// GetRectsDiscardedCount returns the value in RectsDiscardedCount.
func (p FontAtlasBuilder) GetRectsDiscardedCount() int32 {
	ptr := (*C.ImFontAtlasBuilder)(unsafe.Pointer(p))
	return int32(ptr.RectsDiscardedCount)
}

// SetRectsDiscardedCount sets the value in RectsDiscardedCount.
func (p FontAtlasBuilder) SetRectsDiscardedCount(value int32) {
	ptr := (*C.ImFontAtlasBuilder)(unsafe.Pointer(p))
	ptr.RectsDiscardedCount = (C.int32_t)(value)
}

// GetRectsDiscardedSurface returns the value in RectsDiscardedSurface.
func (p FontAtlasBuilder) GetRectsDiscardedSurface() int32 {
	ptr := (*C.ImFontAtlasBuilder)(unsafe.Pointer(p))
	return int32(ptr.RectsDiscardedSurface)
}

// SetRectsDiscardedSurface sets the value in RectsDiscardedSurface.
func (p FontAtlasBuilder) SetRectsDiscardedSurface(value int32) {
	ptr := (*C.ImFontAtlasBuilder)(unsafe.Pointer(p))
	ptr.RectsDiscardedSurface = (C.int32_t)(value)
}

// RefRectsIndex returns pointer to the RectsIndex field.
func (p FontAtlasBuilder) RefRectsIndex() Vector_ImFontAtlasRectEntry {
	return Vector_ImFontAtlasRectEntry(p + FontAtlasBuilder(C.offsetof_ImFontAtlasBuilder_RectsIndex))
}

// GetRectsIndexFreeListStart returns the value in RectsIndexFreeListStart.
func (p FontAtlasBuilder) GetRectsIndexFreeListStart() int32 {
	ptr := (*C.ImFontAtlasBuilder)(unsafe.Pointer(p))
	return int32(ptr.RectsIndexFreeListStart)
}

// SetRectsIndexFreeListStart sets the value in RectsIndexFreeListStart.
func (p FontAtlasBuilder) SetRectsIndexFreeListStart(value int32) {
	ptr := (*C.ImFontAtlasBuilder)(unsafe.Pointer(p))
	ptr.RectsIndexFreeListStart = (C.int32_t)(value)
}

// GetRectsPackedCount returns the value in RectsPackedCount.
func (p FontAtlasBuilder) GetRectsPackedCount() int32 {
	ptr := (*C.ImFontAtlasBuilder)(unsafe.Pointer(p))
	return int32(ptr.RectsPackedCount)
}

// SetRectsPackedCount sets the value in RectsPackedCount.
func (p FontAtlasBuilder) SetRectsPackedCount(value int32) {
	ptr := (*C.ImFontAtlasBuilder)(unsafe.Pointer(p))
	ptr.RectsPackedCount = (C.int32_t)(value)
}

// GetRectsPackedSurface returns the value in RectsPackedSurface.
func (p FontAtlasBuilder) GetRectsPackedSurface() int32 {
	ptr := (*C.ImFontAtlasBuilder)(unsafe.Pointer(p))
	return int32(ptr.RectsPackedSurface)
}

// SetRectsPackedSurface sets the value in RectsPackedSurface.
func (p FontAtlasBuilder) SetRectsPackedSurface(value int32) {
	ptr := (*C.ImFontAtlasBuilder)(unsafe.Pointer(p))
	ptr.RectsPackedSurface = (C.int32_t)(value)
}

// FontAtlasBuilder.TempBuffer is unsupported: category unsupported.

// FontAtlasPostProcessData wraps struct ImFontAtlasPostProcessData.
type FontAtlasPostProcessData uintptr

// FontAtlasPostProcessDataNil is a null pointer.
var FontAtlasPostProcessDataNil FontAtlasPostProcessData

// FontAtlasPostProcessDataSizeOf is the byte size of ImFontAtlasPostProcessData.
const FontAtlasPostProcessDataSizeOf = int(C.sizeof_ImFontAtlasPostProcessData)

// FontAtlasPostProcessDataAlloc allocates a continuous block of FontAtlasPostProcessData.
func FontAtlasPostProcessDataAlloc(alloc ffi.Allocator, count int) FontAtlasPostProcessData {
	ptr := alloc.Allocate(FontAtlasPostProcessDataSizeOf * count)
	return FontAtlasPostProcessData(ptr)
}

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p FontAtlasPostProcessData) Offset(offset int) FontAtlasPostProcessData {
	return p + FontAtlasPostProcessData(offset*FontAtlasPostProcessDataSizeOf)
}

// GetFont returns the value in Font.
func (p FontAtlasPostProcessData) GetFont() Font {
	ptr := (*C.ImFontAtlasPostProcessData)(unsafe.Pointer(p))
	return Font(unsafe.Pointer(ptr.Font))
}

// SetFont sets the value in Font.
func (p FontAtlasPostProcessData) SetFont(value Font) {
	ptr := (*C.ImFontAtlasPostProcessData)(unsafe.Pointer(p))
	ptr.Font = (*C.ImFont)(unsafe.Pointer(value))
}

// GetFontAtlas returns the value in FontAtlas.
func (p FontAtlasPostProcessData) GetFontAtlas() FontAtlas {
	ptr := (*C.ImFontAtlasPostProcessData)(unsafe.Pointer(p))
	return FontAtlas(unsafe.Pointer(ptr.FontAtlas))
}

// SetFontAtlas sets the value in FontAtlas.
func (p FontAtlasPostProcessData) SetFontAtlas(value FontAtlas) {
	ptr := (*C.ImFontAtlasPostProcessData)(unsafe.Pointer(p))
	ptr.FontAtlas = (*C.ImFontAtlas)(unsafe.Pointer(value))
}

// GetFontBaked returns the value in FontBaked.
func (p FontAtlasPostProcessData) GetFontBaked() FontBaked {
	ptr := (*C.ImFontAtlasPostProcessData)(unsafe.Pointer(p))
	return FontBaked(unsafe.Pointer(ptr.FontBaked))
}

// SetFontBaked sets the value in FontBaked.
func (p FontAtlasPostProcessData) SetFontBaked(value FontBaked) {
	ptr := (*C.ImFontAtlasPostProcessData)(unsafe.Pointer(p))
	ptr.FontBaked = (*C.ImFontBaked)(unsafe.Pointer(value))
}

// GetFontSrc returns the value in FontSrc.
func (p FontAtlasPostProcessData) GetFontSrc() FontConfig {
	ptr := (*C.ImFontAtlasPostProcessData)(unsafe.Pointer(p))
	return FontConfig(unsafe.Pointer(ptr.FontSrc))
}

// SetFontSrc sets the value in FontSrc.
func (p FontAtlasPostProcessData) SetFontSrc(value FontConfig) {
	ptr := (*C.ImFontAtlasPostProcessData)(unsafe.Pointer(p))
	ptr.FontSrc = (*C.ImFontConfig)(unsafe.Pointer(value))
}

// GetFormat returns the value in Format.
func (p FontAtlasPostProcessData) GetFormat() TextureFormat {
	ptr := (*C.ImFontAtlasPostProcessData)(unsafe.Pointer(p))
	return TextureFormat(ptr.Format)
}

// SetFormat sets the value in Format.
func (p FontAtlasPostProcessData) SetFormat(value TextureFormat) {
	ptr := (*C.ImFontAtlasPostProcessData)(unsafe.Pointer(p))
	ptr.Format = (C.ImTextureFormat)(value)
}

// GetGlyph returns the value in Glyph.
func (p FontAtlasPostProcessData) GetGlyph() FontGlyph {
	ptr := (*C.ImFontAtlasPostProcessData)(unsafe.Pointer(p))
	return FontGlyph(unsafe.Pointer(ptr.Glyph))
}

// SetGlyph sets the value in Glyph.
func (p FontAtlasPostProcessData) SetGlyph(value FontGlyph) {
	ptr := (*C.ImFontAtlasPostProcessData)(unsafe.Pointer(p))
	ptr.Glyph = (*C.ImFontGlyph)(unsafe.Pointer(value))
}

// GetHeight returns the value in Height.
func (p FontAtlasPostProcessData) GetHeight() int32 {
	ptr := (*C.ImFontAtlasPostProcessData)(unsafe.Pointer(p))
	return int32(ptr.Height)
}

// SetHeight sets the value in Height.
func (p FontAtlasPostProcessData) SetHeight(value int32) {
	ptr := (*C.ImFontAtlasPostProcessData)(unsafe.Pointer(p))
	ptr.Height = (C.int32_t)(value)
}

// GetPitch returns the value in Pitch.
func (p FontAtlasPostProcessData) GetPitch() int32 {
	ptr := (*C.ImFontAtlasPostProcessData)(unsafe.Pointer(p))
	return int32(ptr.Pitch)
}

// SetPitch sets the value in Pitch.
func (p FontAtlasPostProcessData) SetPitch(value int32) {
	ptr := (*C.ImFontAtlasPostProcessData)(unsafe.Pointer(p))
	ptr.Pitch = (C.int32_t)(value)
}

// GetPixels returns the value in Pixels.
func (p FontAtlasPostProcessData) GetPixels() uintptr {
	ptr := (*C.ImFontAtlasPostProcessData)(unsafe.Pointer(p))
	return uintptr(unsafe.Pointer(ptr.Pixels))
}

// SetPixels sets the value in Pixels.
func (p FontAtlasPostProcessData) SetPixels(value uintptr) {
	ptr := (*C.ImFontAtlasPostProcessData)(unsafe.Pointer(p))
	ptr.Pixels = unsafe.Pointer(value)
}

// GetWidth returns the value in Width.
func (p FontAtlasPostProcessData) GetWidth() int32 {
	ptr := (*C.ImFontAtlasPostProcessData)(unsafe.Pointer(p))
	return int32(ptr.Width)
}

// SetWidth sets the value in Width.
func (p FontAtlasPostProcessData) SetWidth(value int32) {
	ptr := (*C.ImFontAtlasPostProcessData)(unsafe.Pointer(p))
	ptr.Width = (C.int32_t)(value)
}

// FontAtlasRect wraps struct ImFontAtlasRect.
type FontAtlasRect uintptr

// FontAtlasRectNil is a null pointer.
var FontAtlasRectNil FontAtlasRect

// FontAtlasRectSizeOf is the byte size of ImFontAtlasRect.
const FontAtlasRectSizeOf = int(C.sizeof_ImFontAtlasRect)

// ImFontAtlasRect allocator is suppressed.

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p FontAtlasRect) Offset(offset int) FontAtlasRect {
	return p + FontAtlasRect(offset*FontAtlasRectSizeOf)
}

// GetH returns the value in h.
func (p FontAtlasRect) GetH() uint16 {
	ptr := (*C.ImFontAtlasRect)(unsafe.Pointer(p))
	return uint16(ptr.h)
}

// SetH sets the value in h.
func (p FontAtlasRect) SetH(value uint16) {
	ptr := (*C.ImFontAtlasRect)(unsafe.Pointer(p))
	ptr.h = (C.uint16_t)(value)
}

// RefUv0 returns pointer to the uv0 field.
func (p FontAtlasRect) RefUv0() Vec2 {
	return Vec2(p + FontAtlasRect(C.offsetof_ImFontAtlasRect_uv0))
}

// RefUv1 returns pointer to the uv1 field.
func (p FontAtlasRect) RefUv1() Vec2 {
	return Vec2(p + FontAtlasRect(C.offsetof_ImFontAtlasRect_uv1))
}

// GetW returns the value in w.
func (p FontAtlasRect) GetW() uint16 {
	ptr := (*C.ImFontAtlasRect)(unsafe.Pointer(p))
	return uint16(ptr.w)
}

// SetW sets the value in w.
func (p FontAtlasRect) SetW(value uint16) {
	ptr := (*C.ImFontAtlasRect)(unsafe.Pointer(p))
	ptr.w = (C.uint16_t)(value)
}

// GetX returns the value in x.
func (p FontAtlasRect) GetX() uint16 {
	ptr := (*C.ImFontAtlasRect)(unsafe.Pointer(p))
	return uint16(ptr.x)
}

// SetX sets the value in x.
func (p FontAtlasRect) SetX(value uint16) {
	ptr := (*C.ImFontAtlasRect)(unsafe.Pointer(p))
	ptr.x = (C.uint16_t)(value)
}

// GetY returns the value in y.
func (p FontAtlasRect) GetY() uint16 {
	ptr := (*C.ImFontAtlasRect)(unsafe.Pointer(p))
	return uint16(ptr.y)
}

// SetY sets the value in y.
func (p FontAtlasRect) SetY(value uint16) {
	ptr := (*C.ImFontAtlasRect)(unsafe.Pointer(p))
	ptr.y = (C.uint16_t)(value)
}

// FontAtlasRectEntry wraps struct ImFontAtlasRectEntry.
type FontAtlasRectEntry uintptr

// FontAtlasRectEntryNil is a null pointer.
var FontAtlasRectEntryNil FontAtlasRectEntry

// FontAtlasRectEntrySizeOf is the byte size of ImFontAtlasRectEntry.
const FontAtlasRectEntrySizeOf = int(C.sizeof_ImFontAtlasRectEntry)

// FontAtlasRectEntryAlloc allocates a continuous block of FontAtlasRectEntry.
func FontAtlasRectEntryAlloc(alloc ffi.Allocator, count int) FontAtlasRectEntry {
	ptr := alloc.Allocate(FontAtlasRectEntrySizeOf * count)
	return FontAtlasRectEntry(ptr)
}

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p FontAtlasRectEntry) Offset(offset int) FontAtlasRectEntry {
	return p + FontAtlasRectEntry(offset*FontAtlasRectEntrySizeOf)
}

// FontAtlasRectEntry.Generation is unsupported: category unsupported.

// FontAtlasRectEntry.IsUsed is unsupported: category unsupported.

// FontAtlasRectEntry.TargetIndex is unsupported: category unsupported.

// FontBaked wraps struct ImFontBaked.
type FontBaked uintptr

// FontBakedNil is a null pointer.
var FontBakedNil FontBaked

// FontBakedSizeOf is the byte size of ImFontBaked.
const FontBakedSizeOf = int(C.sizeof_ImFontBaked)

// ImFontBaked allocator is suppressed.

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p FontBaked) Offset(offset int) FontBaked {
	return p + FontBaked(offset*FontBakedSizeOf)
}

// GetAscent returns the value in Ascent.
func (p FontBaked) GetAscent() float32 {
	ptr := (*C.ImFontBaked)(unsafe.Pointer(p))
	return float32(ptr.Ascent)
}

// SetAscent sets the value in Ascent.
func (p FontBaked) SetAscent(value float32) {
	ptr := (*C.ImFontBaked)(unsafe.Pointer(p))
	ptr.Ascent = (C.float)(value)
}

// FontBaked.BakedId is unsupported: category unsupported.

// GetContainerFont returns the value in ContainerFont.
func (p FontBaked) GetContainerFont() Font {
	ptr := (*C.ImFontBaked)(unsafe.Pointer(p))
	return Font(unsafe.Pointer(ptr.ContainerFont))
}

// SetContainerFont sets the value in ContainerFont.
func (p FontBaked) SetContainerFont(value Font) {
	ptr := (*C.ImFontBaked)(unsafe.Pointer(p))
	ptr.ContainerFont = (*C.ImFont)(unsafe.Pointer(value))
}

// GetDescent returns the value in Descent.
func (p FontBaked) GetDescent() float32 {
	ptr := (*C.ImFontBaked)(unsafe.Pointer(p))
	return float32(ptr.Descent)
}

// SetDescent sets the value in Descent.
func (p FontBaked) SetDescent(value float32) {
	ptr := (*C.ImFontBaked)(unsafe.Pointer(p))
	ptr.Descent = (C.float)(value)
}

// GetFallbackAdvanceX returns the value in FallbackAdvanceX.
func (p FontBaked) GetFallbackAdvanceX() float32 {
	ptr := (*C.ImFontBaked)(unsafe.Pointer(p))
	return float32(ptr.FallbackAdvanceX)
}

// SetFallbackAdvanceX sets the value in FallbackAdvanceX.
func (p FontBaked) SetFallbackAdvanceX(value float32) {
	ptr := (*C.ImFontBaked)(unsafe.Pointer(p))
	ptr.FallbackAdvanceX = (C.float)(value)
}

// GetFallbackGlyphIndex returns the value in FallbackGlyphIndex.
func (p FontBaked) GetFallbackGlyphIndex() int32 {
	ptr := (*C.ImFontBaked)(unsafe.Pointer(p))
	return int32(ptr.FallbackGlyphIndex)
}

// SetFallbackGlyphIndex sets the value in FallbackGlyphIndex.
func (p FontBaked) SetFallbackGlyphIndex(value int32) {
	ptr := (*C.ImFontBaked)(unsafe.Pointer(p))
	ptr.FallbackGlyphIndex = (C.int32_t)(value)
}

// GetFontLoaderDatas returns the value in FontLoaderDatas.
func (p FontBaked) GetFontLoaderDatas() uintptr {
	ptr := (*C.ImFontBaked)(unsafe.Pointer(p))
	return uintptr(unsafe.Pointer(ptr.FontLoaderDatas))
}

// SetFontLoaderDatas sets the value in FontLoaderDatas.
func (p FontBaked) SetFontLoaderDatas(value uintptr) {
	ptr := (*C.ImFontBaked)(unsafe.Pointer(p))
	ptr.FontLoaderDatas = unsafe.Pointer(value)
}

// RefGlyphs returns pointer to the Glyphs field.
func (p FontBaked) RefGlyphs() Vector_ImFontGlyph {
	return Vector_ImFontGlyph(p + FontBaked(C.offsetof_ImFontBaked_Glyphs))
}

// RefIndexAdvanceX returns pointer to the IndexAdvanceX field.
func (p FontBaked) RefIndexAdvanceX() Vector_float {
	return Vector_float(p + FontBaked(C.offsetof_ImFontBaked_IndexAdvanceX))
}

// RefIndexLookup returns pointer to the IndexLookup field.
func (p FontBaked) RefIndexLookup() Vector_ImU16 {
	return Vector_ImU16(p + FontBaked(C.offsetof_ImFontBaked_IndexLookup))
}

// GetLastUsedFrame returns the value in LastUsedFrame.
func (p FontBaked) GetLastUsedFrame() int32 {
	ptr := (*C.ImFontBaked)(unsafe.Pointer(p))
	return int32(ptr.LastUsedFrame)
}

// SetLastUsedFrame sets the value in LastUsedFrame.
func (p FontBaked) SetLastUsedFrame(value int32) {
	ptr := (*C.ImFontBaked)(unsafe.Pointer(p))
	ptr.LastUsedFrame = (C.int32_t)(value)
}

// FontBaked.LoadNoFallback is unsupported: category unsupported.

// FontBaked.LoadNoRenderOnLayout is unsupported: category unsupported.

// FontBaked.MetricsTotalSurface is unsupported: category unsupported.

// GetRasterizerDensity returns the value in RasterizerDensity.
func (p FontBaked) GetRasterizerDensity() float32 {
	ptr := (*C.ImFontBaked)(unsafe.Pointer(p))
	return float32(ptr.RasterizerDensity)
}

// SetRasterizerDensity sets the value in RasterizerDensity.
func (p FontBaked) SetRasterizerDensity(value float32) {
	ptr := (*C.ImFontBaked)(unsafe.Pointer(p))
	ptr.RasterizerDensity = (C.float)(value)
}

// GetSize returns the value in Size.
func (p FontBaked) GetSize() float32 {
	ptr := (*C.ImFontBaked)(unsafe.Pointer(p))
	return float32(ptr.Size)
}

// SetSize sets the value in Size.
func (p FontBaked) SetSize(value float32) {
	ptr := (*C.ImFontBaked)(unsafe.Pointer(p))
	ptr.Size = (C.float)(value)
}

// FontBaked.WantDestroy is unsupported: category unsupported.

// FontConfig wraps struct ImFontConfig.
type FontConfig uintptr

// FontConfigNil is a null pointer.
var FontConfigNil FontConfig

// FontConfigSizeOf is the byte size of ImFontConfig.
const FontConfigSizeOf = int(C.sizeof_ImFontConfig)

// ImFontConfig allocator is suppressed.

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p FontConfig) Offset(offset int) FontConfig {
	return p + FontConfig(offset*FontConfigSizeOf)
}

// GetDstFont returns the value in DstFont.
func (p FontConfig) GetDstFont() Font {
	ptr := (*C.ImFontConfig)(unsafe.Pointer(p))
	return Font(unsafe.Pointer(ptr.DstFont))
}

// SetDstFont sets the value in DstFont.
func (p FontConfig) SetDstFont(value Font) {
	ptr := (*C.ImFontConfig)(unsafe.Pointer(p))
	ptr.DstFont = (*C.ImFont)(unsafe.Pointer(value))
}

// FontConfig.EllipsisChar is unsupported: category unsupported.

// GetFlags returns the value in Flags.
func (p FontConfig) GetFlags() FontFlags {
	ptr := (*C.ImFontConfig)(unsafe.Pointer(p))
	return FontFlags(ptr.Flags)
}

// SetFlags sets the value in Flags.
func (p FontConfig) SetFlags(value FontFlags) {
	ptr := (*C.ImFontConfig)(unsafe.Pointer(p))
	ptr.Flags = (C.ImFontFlags)(value)
}

// GetFontData returns the value in FontData.
func (p FontConfig) GetFontData() uintptr {
	ptr := (*C.ImFontConfig)(unsafe.Pointer(p))
	return uintptr(unsafe.Pointer(ptr.FontData))
}

// SetFontData sets the value in FontData.
func (p FontConfig) SetFontData(value uintptr) {
	ptr := (*C.ImFontConfig)(unsafe.Pointer(p))
	ptr.FontData = unsafe.Pointer(value)
}

// GetFontDataOwnedByAtlas returns the value in FontDataOwnedByAtlas.
func (p FontConfig) GetFontDataOwnedByAtlas() bool {
	ptr := (*C.ImFontConfig)(unsafe.Pointer(p))
	return bool(ptr.FontDataOwnedByAtlas)
}

// SetFontDataOwnedByAtlas sets the value in FontDataOwnedByAtlas.
func (p FontConfig) SetFontDataOwnedByAtlas(value bool) {
	ptr := (*C.ImFontConfig)(unsafe.Pointer(p))
	ptr.FontDataOwnedByAtlas = (C.bool)(value)
}

// GetFontDataSize returns the value in FontDataSize.
func (p FontConfig) GetFontDataSize() int32 {
	ptr := (*C.ImFontConfig)(unsafe.Pointer(p))
	return int32(ptr.FontDataSize)
}

// SetFontDataSize sets the value in FontDataSize.
func (p FontConfig) SetFontDataSize(value int32) {
	ptr := (*C.ImFontConfig)(unsafe.Pointer(p))
	ptr.FontDataSize = (C.int32_t)(value)
}

// GetFontLoader returns the value in FontLoader.
func (p FontConfig) GetFontLoader() FontLoader {
	ptr := (*C.ImFontConfig)(unsafe.Pointer(p))
	return FontLoader(unsafe.Pointer(ptr.FontLoader))
}

// SetFontLoader sets the value in FontLoader.
func (p FontConfig) SetFontLoader(value FontLoader) {
	ptr := (*C.ImFontConfig)(unsafe.Pointer(p))
	ptr.FontLoader = (*C.ImFontLoader)(unsafe.Pointer(value))
}

// GetFontLoaderData returns the value in FontLoaderData.
func (p FontConfig) GetFontLoaderData() uintptr {
	ptr := (*C.ImFontConfig)(unsafe.Pointer(p))
	return uintptr(unsafe.Pointer(ptr.FontLoaderData))
}

// SetFontLoaderData sets the value in FontLoaderData.
func (p FontConfig) SetFontLoaderData(value uintptr) {
	ptr := (*C.ImFontConfig)(unsafe.Pointer(p))
	ptr.FontLoaderData = unsafe.Pointer(value)
}

// GetFontLoaderFlags returns the value in FontLoaderFlags.
func (p FontConfig) GetFontLoaderFlags() uint32 {
	ptr := (*C.ImFontConfig)(unsafe.Pointer(p))
	return uint32(ptr.FontLoaderFlags)
}

// SetFontLoaderFlags sets the value in FontLoaderFlags.
func (p FontConfig) SetFontLoaderFlags(value uint32) {
	ptr := (*C.ImFontConfig)(unsafe.Pointer(p))
	ptr.FontLoaderFlags = (C.uint32_t)(value)
}

// GetFontNo returns the value in FontNo.
func (p FontConfig) GetFontNo() uint32 {
	ptr := (*C.ImFontConfig)(unsafe.Pointer(p))
	return uint32(ptr.FontNo)
}

// SetFontNo sets the value in FontNo.
func (p FontConfig) SetFontNo(value uint32) {
	ptr := (*C.ImFontConfig)(unsafe.Pointer(p))
	ptr.FontNo = (C.uint32_t)(value)
}

// FontConfig.GlyphExcludeRanges is unsupported: category unsupported.

// GetGlyphExtraAdvanceX returns the value in GlyphExtraAdvanceX.
func (p FontConfig) GetGlyphExtraAdvanceX() float32 {
	ptr := (*C.ImFontConfig)(unsafe.Pointer(p))
	return float32(ptr.GlyphExtraAdvanceX)
}

// SetGlyphExtraAdvanceX sets the value in GlyphExtraAdvanceX.
func (p FontConfig) SetGlyphExtraAdvanceX(value float32) {
	ptr := (*C.ImFontConfig)(unsafe.Pointer(p))
	ptr.GlyphExtraAdvanceX = (C.float)(value)
}

// GetGlyphMaxAdvanceX returns the value in GlyphMaxAdvanceX.
func (p FontConfig) GetGlyphMaxAdvanceX() float32 {
	ptr := (*C.ImFontConfig)(unsafe.Pointer(p))
	return float32(ptr.GlyphMaxAdvanceX)
}

// SetGlyphMaxAdvanceX sets the value in GlyphMaxAdvanceX.
func (p FontConfig) SetGlyphMaxAdvanceX(value float32) {
	ptr := (*C.ImFontConfig)(unsafe.Pointer(p))
	ptr.GlyphMaxAdvanceX = (C.float)(value)
}

// GetGlyphMinAdvanceX returns the value in GlyphMinAdvanceX.
func (p FontConfig) GetGlyphMinAdvanceX() float32 {
	ptr := (*C.ImFontConfig)(unsafe.Pointer(p))
	return float32(ptr.GlyphMinAdvanceX)
}

// SetGlyphMinAdvanceX sets the value in GlyphMinAdvanceX.
func (p FontConfig) SetGlyphMinAdvanceX(value float32) {
	ptr := (*C.ImFontConfig)(unsafe.Pointer(p))
	ptr.GlyphMinAdvanceX = (C.float)(value)
}

// RefGlyphOffset returns pointer to the GlyphOffset field.
func (p FontConfig) RefGlyphOffset() Vec2 {
	return Vec2(p + FontConfig(C.offsetof_ImFontConfig_GlyphOffset))
}

// FontConfig.GlyphRanges is unsupported: category unsupported.

// GetMergeMode returns the value in MergeMode.
func (p FontConfig) GetMergeMode() bool {
	ptr := (*C.ImFontConfig)(unsafe.Pointer(p))
	return bool(ptr.MergeMode)
}

// SetMergeMode sets the value in MergeMode.
func (p FontConfig) SetMergeMode(value bool) {
	ptr := (*C.ImFontConfig)(unsafe.Pointer(p))
	ptr.MergeMode = (C.bool)(value)
}

// FontConfig.Name[40] is unsupported: category unsupported.

// FontConfig.OversampleH is unsupported: category unsupported.

// FontConfig.OversampleV is unsupported: category unsupported.

// GetPixelSnapH returns the value in PixelSnapH.
func (p FontConfig) GetPixelSnapH() bool {
	ptr := (*C.ImFontConfig)(unsafe.Pointer(p))
	return bool(ptr.PixelSnapH)
}

// SetPixelSnapH sets the value in PixelSnapH.
func (p FontConfig) SetPixelSnapH(value bool) {
	ptr := (*C.ImFontConfig)(unsafe.Pointer(p))
	ptr.PixelSnapH = (C.bool)(value)
}

// GetPixelSnapV returns the value in PixelSnapV.
func (p FontConfig) GetPixelSnapV() bool {
	ptr := (*C.ImFontConfig)(unsafe.Pointer(p))
	return bool(ptr.PixelSnapV)
}

// SetPixelSnapV sets the value in PixelSnapV.
func (p FontConfig) SetPixelSnapV(value bool) {
	ptr := (*C.ImFontConfig)(unsafe.Pointer(p))
	ptr.PixelSnapV = (C.bool)(value)
}

// GetRasterizerDensity returns the value in RasterizerDensity.
func (p FontConfig) GetRasterizerDensity() float32 {
	ptr := (*C.ImFontConfig)(unsafe.Pointer(p))
	return float32(ptr.RasterizerDensity)
}

// SetRasterizerDensity sets the value in RasterizerDensity.
func (p FontConfig) SetRasterizerDensity(value float32) {
	ptr := (*C.ImFontConfig)(unsafe.Pointer(p))
	ptr.RasterizerDensity = (C.float)(value)
}

// GetRasterizerMultiply returns the value in RasterizerMultiply.
func (p FontConfig) GetRasterizerMultiply() float32 {
	ptr := (*C.ImFontConfig)(unsafe.Pointer(p))
	return float32(ptr.RasterizerMultiply)
}

// SetRasterizerMultiply sets the value in RasterizerMultiply.
func (p FontConfig) SetRasterizerMultiply(value float32) {
	ptr := (*C.ImFontConfig)(unsafe.Pointer(p))
	ptr.RasterizerMultiply = (C.float)(value)
}

// GetSizePixels returns the value in SizePixels.
func (p FontConfig) GetSizePixels() float32 {
	ptr := (*C.ImFontConfig)(unsafe.Pointer(p))
	return float32(ptr.SizePixels)
}

// SetSizePixels sets the value in SizePixels.
func (p FontConfig) SetSizePixels(value float32) {
	ptr := (*C.ImFontConfig)(unsafe.Pointer(p))
	ptr.SizePixels = (C.float)(value)
}

// FontGlyph wraps struct ImFontGlyph.
type FontGlyph uintptr

// FontGlyphNil is a null pointer.
var FontGlyphNil FontGlyph

// FontGlyphSizeOf is the byte size of ImFontGlyph.
const FontGlyphSizeOf = int(C.sizeof_ImFontGlyph)

// ImFontGlyph allocator is suppressed.

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p FontGlyph) Offset(offset int) FontGlyph {
	return p + FontGlyph(offset*FontGlyphSizeOf)
}

// GetAdvanceX returns the value in AdvanceX.
func (p FontGlyph) GetAdvanceX() float32 {
	ptr := (*C.ImFontGlyph)(unsafe.Pointer(p))
	return float32(ptr.AdvanceX)
}

// SetAdvanceX sets the value in AdvanceX.
func (p FontGlyph) SetAdvanceX(value float32) {
	ptr := (*C.ImFontGlyph)(unsafe.Pointer(p))
	ptr.AdvanceX = (C.float)(value)
}

// FontGlyph.Codepoint is unsupported: category unsupported.

// FontGlyph.Colored is unsupported: category unsupported.

// GetPackId returns the value in PackId.
func (p FontGlyph) GetPackId() int32 {
	ptr := (*C.ImFontGlyph)(unsafe.Pointer(p))
	return int32(ptr.PackId)
}

// SetPackId sets the value in PackId.
func (p FontGlyph) SetPackId(value int32) {
	ptr := (*C.ImFontGlyph)(unsafe.Pointer(p))
	ptr.PackId = (C.int32_t)(value)
}

// FontGlyph.SourceIdx is unsupported: category unsupported.

// GetU0 returns the value in U0.
func (p FontGlyph) GetU0() float32 {
	ptr := (*C.ImFontGlyph)(unsafe.Pointer(p))
	return float32(ptr.U0)
}

// SetU0 sets the value in U0.
func (p FontGlyph) SetU0(value float32) {
	ptr := (*C.ImFontGlyph)(unsafe.Pointer(p))
	ptr.U0 = (C.float)(value)
}

// GetU1 returns the value in U1.
func (p FontGlyph) GetU1() float32 {
	ptr := (*C.ImFontGlyph)(unsafe.Pointer(p))
	return float32(ptr.U1)
}

// SetU1 sets the value in U1.
func (p FontGlyph) SetU1(value float32) {
	ptr := (*C.ImFontGlyph)(unsafe.Pointer(p))
	ptr.U1 = (C.float)(value)
}

// GetV0 returns the value in V0.
func (p FontGlyph) GetV0() float32 {
	ptr := (*C.ImFontGlyph)(unsafe.Pointer(p))
	return float32(ptr.V0)
}

// SetV0 sets the value in V0.
func (p FontGlyph) SetV0(value float32) {
	ptr := (*C.ImFontGlyph)(unsafe.Pointer(p))
	ptr.V0 = (C.float)(value)
}

// GetV1 returns the value in V1.
func (p FontGlyph) GetV1() float32 {
	ptr := (*C.ImFontGlyph)(unsafe.Pointer(p))
	return float32(ptr.V1)
}

// SetV1 sets the value in V1.
func (p FontGlyph) SetV1(value float32) {
	ptr := (*C.ImFontGlyph)(unsafe.Pointer(p))
	ptr.V1 = (C.float)(value)
}

// FontGlyph.Visible is unsupported: category unsupported.

// GetX0 returns the value in X0.
func (p FontGlyph) GetX0() float32 {
	ptr := (*C.ImFontGlyph)(unsafe.Pointer(p))
	return float32(ptr.X0)
}

// SetX0 sets the value in X0.
func (p FontGlyph) SetX0(value float32) {
	ptr := (*C.ImFontGlyph)(unsafe.Pointer(p))
	ptr.X0 = (C.float)(value)
}

// GetX1 returns the value in X1.
func (p FontGlyph) GetX1() float32 {
	ptr := (*C.ImFontGlyph)(unsafe.Pointer(p))
	return float32(ptr.X1)
}

// SetX1 sets the value in X1.
func (p FontGlyph) SetX1(value float32) {
	ptr := (*C.ImFontGlyph)(unsafe.Pointer(p))
	ptr.X1 = (C.float)(value)
}

// GetY0 returns the value in Y0.
func (p FontGlyph) GetY0() float32 {
	ptr := (*C.ImFontGlyph)(unsafe.Pointer(p))
	return float32(ptr.Y0)
}

// SetY0 sets the value in Y0.
func (p FontGlyph) SetY0(value float32) {
	ptr := (*C.ImFontGlyph)(unsafe.Pointer(p))
	ptr.Y0 = (C.float)(value)
}

// GetY1 returns the value in Y1.
func (p FontGlyph) GetY1() float32 {
	ptr := (*C.ImFontGlyph)(unsafe.Pointer(p))
	return float32(ptr.Y1)
}

// SetY1 sets the value in Y1.
func (p FontGlyph) SetY1(value float32) {
	ptr := (*C.ImFontGlyph)(unsafe.Pointer(p))
	ptr.Y1 = (C.float)(value)
}

// FontGlyphRangesBuilder wraps struct ImFontGlyphRangesBuilder.
type FontGlyphRangesBuilder uintptr

// FontGlyphRangesBuilderNil is a null pointer.
var FontGlyphRangesBuilderNil FontGlyphRangesBuilder

// FontGlyphRangesBuilderSizeOf is the byte size of ImFontGlyphRangesBuilder.
const FontGlyphRangesBuilderSizeOf = int(C.sizeof_ImFontGlyphRangesBuilder)

// ImFontGlyphRangesBuilder allocator is suppressed.

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p FontGlyphRangesBuilder) Offset(offset int) FontGlyphRangesBuilder {
	return p + FontGlyphRangesBuilder(offset*FontGlyphRangesBuilderSizeOf)
}

// RefUsedChars returns pointer to the UsedChars field.
func (p FontGlyphRangesBuilder) RefUsedChars() Vector_ImU32 {
	return Vector_ImU32(p + FontGlyphRangesBuilder(C.offsetof_ImFontGlyphRangesBuilder_UsedChars))
}

// FontLoader wraps struct ImFontLoader.
type FontLoader uintptr

// FontLoaderNil is a null pointer.
var FontLoaderNil FontLoader

// FontLoaderSizeOf is the byte size of ImFontLoader.
const FontLoaderSizeOf = int(C.sizeof_ImFontLoader)

// ImFontLoader allocator is suppressed.

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p FontLoader) Offset(offset int) FontLoader {
	return p + FontLoader(offset*FontLoaderSizeOf)
}

// FontLoader.FontBakedDestroy is unsupported: category unsupported.

// FontLoader.FontBakedInit is unsupported: category unsupported.

// FontLoader.FontBakedLoadGlyph is unsupported: category unsupported.

// FontLoader.FontBakedSrcLoaderDataSize is unsupported: category unsupported.

// FontLoader.FontSrcContainsGlyph is unsupported: category unsupported.

// FontLoader.FontSrcDestroy is unsupported: category unsupported.

// FontLoader.FontSrcInit is unsupported: category unsupported.

// FontLoader.LoaderInit is unsupported: category unsupported.

// FontLoader.LoaderShutdown is unsupported: category unsupported.

// GetName returns the value in Name.
func (p FontLoader) GetName() ffi.CString {
	ptr := (*C.ImFontLoader)(unsafe.Pointer(p))
	return ffi.CStringFromPtr((unsafe.Pointer)(ptr.Name))
}

// SetName sets the value in Name.
func (p FontLoader) SetName(value ffi.CString) {
	ptr := (*C.ImFontLoader)(unsafe.Pointer(p))
	ptr.Name = (*C.char)(value.Raw())
}

// FontStackData wraps struct ImFontStackData.
type FontStackData uintptr

// FontStackDataNil is a null pointer.
var FontStackDataNil FontStackData

// FontStackDataSizeOf is the byte size of ImFontStackData.
const FontStackDataSizeOf = int(C.sizeof_ImFontStackData)

// FontStackDataAlloc allocates a continuous block of FontStackData.
func FontStackDataAlloc(alloc ffi.Allocator, count int) FontStackData {
	ptr := alloc.Allocate(FontStackDataSizeOf * count)
	return FontStackData(ptr)
}

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p FontStackData) Offset(offset int) FontStackData {
	return p + FontStackData(offset*FontStackDataSizeOf)
}

// GetFont returns the value in Font.
func (p FontStackData) GetFont() Font {
	ptr := (*C.ImFontStackData)(unsafe.Pointer(p))
	return Font(unsafe.Pointer(ptr.Font))
}

// SetFont sets the value in Font.
func (p FontStackData) SetFont(value Font) {
	ptr := (*C.ImFontStackData)(unsafe.Pointer(p))
	ptr.Font = (*C.ImFont)(unsafe.Pointer(value))
}

// GetFontSizeAfterScaling returns the value in FontSizeAfterScaling.
func (p FontStackData) GetFontSizeAfterScaling() float32 {
	ptr := (*C.ImFontStackData)(unsafe.Pointer(p))
	return float32(ptr.FontSizeAfterScaling)
}

// SetFontSizeAfterScaling sets the value in FontSizeAfterScaling.
func (p FontStackData) SetFontSizeAfterScaling(value float32) {
	ptr := (*C.ImFontStackData)(unsafe.Pointer(p))
	ptr.FontSizeAfterScaling = (C.float)(value)
}

// GetFontSizeBeforeScaling returns the value in FontSizeBeforeScaling.
func (p FontStackData) GetFontSizeBeforeScaling() float32 {
	ptr := (*C.ImFontStackData)(unsafe.Pointer(p))
	return float32(ptr.FontSizeBeforeScaling)
}

// SetFontSizeBeforeScaling sets the value in FontSizeBeforeScaling.
func (p FontStackData) SetFontSizeBeforeScaling(value float32) {
	ptr := (*C.ImFontStackData)(unsafe.Pointer(p))
	ptr.FontSizeBeforeScaling = (C.float)(value)
}

// GuiBoxSelectState wraps struct ImGuiBoxSelectState.
type GuiBoxSelectState uintptr

// GuiBoxSelectStateNil is a null pointer.
var GuiBoxSelectStateNil GuiBoxSelectState

// GuiBoxSelectStateSizeOf is the byte size of ImGuiBoxSelectState.
const GuiBoxSelectStateSizeOf = int(C.sizeof_ImGuiBoxSelectState)

// ImGuiBoxSelectState allocator is suppressed.

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p GuiBoxSelectState) Offset(offset int) GuiBoxSelectState {
	return p + GuiBoxSelectState(offset*GuiBoxSelectStateSizeOf)
}

// RefBoxSelectRectCurr returns pointer to the BoxSelectRectCurr field.
func (p GuiBoxSelectState) RefBoxSelectRectCurr() Rect {
	return Rect(p + GuiBoxSelectState(C.offsetof_ImGuiBoxSelectState_BoxSelectRectCurr))
}

// RefBoxSelectRectPrev returns pointer to the BoxSelectRectPrev field.
func (p GuiBoxSelectState) RefBoxSelectRectPrev() Rect {
	return Rect(p + GuiBoxSelectState(C.offsetof_ImGuiBoxSelectState_BoxSelectRectPrev))
}

// RefEndPosRel returns pointer to the EndPosRel field.
func (p GuiBoxSelectState) RefEndPosRel() Vec2 {
	return Vec2(p + GuiBoxSelectState(C.offsetof_ImGuiBoxSelectState_EndPosRel))
}

// GuiBoxSelectState.ID is unsupported: category unsupported.

// GetIsActive returns the value in IsActive.
func (p GuiBoxSelectState) GetIsActive() bool {
	ptr := (*C.ImGuiBoxSelectState)(unsafe.Pointer(p))
	return bool(ptr.IsActive)
}

// SetIsActive sets the value in IsActive.
func (p GuiBoxSelectState) SetIsActive(value bool) {
	ptr := (*C.ImGuiBoxSelectState)(unsafe.Pointer(p))
	ptr.IsActive = (C.bool)(value)
}

// GetIsStartedFromVoid returns the value in IsStartedFromVoid.
func (p GuiBoxSelectState) GetIsStartedFromVoid() bool {
	ptr := (*C.ImGuiBoxSelectState)(unsafe.Pointer(p))
	return bool(ptr.IsStartedFromVoid)
}

// SetIsStartedFromVoid sets the value in IsStartedFromVoid.
func (p GuiBoxSelectState) SetIsStartedFromVoid(value bool) {
	ptr := (*C.ImGuiBoxSelectState)(unsafe.Pointer(p))
	ptr.IsStartedFromVoid = (C.bool)(value)
}

// GetIsStartedSetNavIdOnce returns the value in IsStartedSetNavIdOnce.
func (p GuiBoxSelectState) GetIsStartedSetNavIdOnce() bool {
	ptr := (*C.ImGuiBoxSelectState)(unsafe.Pointer(p))
	return bool(ptr.IsStartedSetNavIdOnce)
}

// SetIsStartedSetNavIdOnce sets the value in IsStartedSetNavIdOnce.
func (p GuiBoxSelectState) SetIsStartedSetNavIdOnce(value bool) {
	ptr := (*C.ImGuiBoxSelectState)(unsafe.Pointer(p))
	ptr.IsStartedSetNavIdOnce = (C.bool)(value)
}

// GetIsStarting returns the value in IsStarting.
func (p GuiBoxSelectState) GetIsStarting() bool {
	ptr := (*C.ImGuiBoxSelectState)(unsafe.Pointer(p))
	return bool(ptr.IsStarting)
}

// SetIsStarting sets the value in IsStarting.
func (p GuiBoxSelectState) SetIsStarting(value bool) {
	ptr := (*C.ImGuiBoxSelectState)(unsafe.Pointer(p))
	ptr.IsStarting = (C.bool)(value)
}

// GuiBoxSelectState.KeyMods is unsupported: category unsupported.

// GetRequestClear returns the value in RequestClear.
func (p GuiBoxSelectState) GetRequestClear() bool {
	ptr := (*C.ImGuiBoxSelectState)(unsafe.Pointer(p))
	return bool(ptr.RequestClear)
}

// SetRequestClear sets the value in RequestClear.
func (p GuiBoxSelectState) SetRequestClear(value bool) {
	ptr := (*C.ImGuiBoxSelectState)(unsafe.Pointer(p))
	ptr.RequestClear = (C.bool)(value)
}

// RefScrollAccum returns pointer to the ScrollAccum field.
func (p GuiBoxSelectState) RefScrollAccum() Vec2 {
	return Vec2(p + GuiBoxSelectState(C.offsetof_ImGuiBoxSelectState_ScrollAccum))
}

// RefStartPosRel returns pointer to the StartPosRel field.
func (p GuiBoxSelectState) RefStartPosRel() Vec2 {
	return Vec2(p + GuiBoxSelectState(C.offsetof_ImGuiBoxSelectState_StartPosRel))
}

// GetUnclipMode returns the value in UnclipMode.
func (p GuiBoxSelectState) GetUnclipMode() bool {
	ptr := (*C.ImGuiBoxSelectState)(unsafe.Pointer(p))
	return bool(ptr.UnclipMode)
}

// SetUnclipMode sets the value in UnclipMode.
func (p GuiBoxSelectState) SetUnclipMode(value bool) {
	ptr := (*C.ImGuiBoxSelectState)(unsafe.Pointer(p))
	ptr.UnclipMode = (C.bool)(value)
}

// RefUnclipRect returns pointer to the UnclipRect field.
func (p GuiBoxSelectState) RefUnclipRect() Rect {
	return Rect(p + GuiBoxSelectState(C.offsetof_ImGuiBoxSelectState_UnclipRect))
}

// GetWindow returns the value in Window.
func (p GuiBoxSelectState) GetWindow() GuiWindow {
	ptr := (*C.ImGuiBoxSelectState)(unsafe.Pointer(p))
	return GuiWindow(unsafe.Pointer(ptr.Window))
}

// SetWindow sets the value in Window.
func (p GuiBoxSelectState) SetWindow(value GuiWindow) {
	ptr := (*C.ImGuiBoxSelectState)(unsafe.Pointer(p))
	ptr.Window = (*C.ImGuiWindow)(unsafe.Pointer(value))
}

// GuiColorMod wraps struct ImGuiColorMod.
type GuiColorMod uintptr

// GuiColorModNil is a null pointer.
var GuiColorModNil GuiColorMod

// GuiColorModSizeOf is the byte size of ImGuiColorMod.
const GuiColorModSizeOf = int(C.sizeof_ImGuiColorMod)

// GuiColorModAlloc allocates a continuous block of GuiColorMod.
func GuiColorModAlloc(alloc ffi.Allocator, count int) GuiColorMod {
	ptr := alloc.Allocate(GuiColorModSizeOf * count)
	return GuiColorMod(ptr)
}

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p GuiColorMod) Offset(offset int) GuiColorMod {
	return p + GuiColorMod(offset*GuiColorModSizeOf)
}

// RefBackupValue returns pointer to the BackupValue field.
func (p GuiColorMod) RefBackupValue() Vec4 {
	return Vec4(p + GuiColorMod(C.offsetof_ImGuiColorMod_BackupValue))
}

// GetCol returns the value in Col.
func (p GuiColorMod) GetCol() GuiCol {
	ptr := (*C.ImGuiColorMod)(unsafe.Pointer(p))
	return GuiCol(ptr.Col)
}

// SetCol sets the value in Col.
func (p GuiColorMod) SetCol(value GuiCol) {
	ptr := (*C.ImGuiColorMod)(unsafe.Pointer(p))
	ptr.Col = (C.ImGuiCol)(value)
}

// GuiComboPreviewData wraps struct ImGuiComboPreviewData.
type GuiComboPreviewData uintptr

// GuiComboPreviewDataNil is a null pointer.
var GuiComboPreviewDataNil GuiComboPreviewData

// GuiComboPreviewDataSizeOf is the byte size of ImGuiComboPreviewData.
const GuiComboPreviewDataSizeOf = int(C.sizeof_ImGuiComboPreviewData)

// ImGuiComboPreviewData allocator is suppressed.

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p GuiComboPreviewData) Offset(offset int) GuiComboPreviewData {
	return p + GuiComboPreviewData(offset*GuiComboPreviewDataSizeOf)
}

// RefBackupCursorMaxPos returns pointer to the BackupCursorMaxPos field.
func (p GuiComboPreviewData) RefBackupCursorMaxPos() Vec2 {
	return Vec2(p + GuiComboPreviewData(C.offsetof_ImGuiComboPreviewData_BackupCursorMaxPos))
}

// RefBackupCursorPos returns pointer to the BackupCursorPos field.
func (p GuiComboPreviewData) RefBackupCursorPos() Vec2 {
	return Vec2(p + GuiComboPreviewData(C.offsetof_ImGuiComboPreviewData_BackupCursorPos))
}

// RefBackupCursorPosPrevLine returns pointer to the BackupCursorPosPrevLine field.
func (p GuiComboPreviewData) RefBackupCursorPosPrevLine() Vec2 {
	return Vec2(p + GuiComboPreviewData(C.offsetof_ImGuiComboPreviewData_BackupCursorPosPrevLine))
}

// GetBackupLayout returns the value in BackupLayout.
func (p GuiComboPreviewData) GetBackupLayout() GuiLayoutType {
	ptr := (*C.ImGuiComboPreviewData)(unsafe.Pointer(p))
	return GuiLayoutType(ptr.BackupLayout)
}

// SetBackupLayout sets the value in BackupLayout.
func (p GuiComboPreviewData) SetBackupLayout(value GuiLayoutType) {
	ptr := (*C.ImGuiComboPreviewData)(unsafe.Pointer(p))
	ptr.BackupLayout = (C.ImGuiLayoutType)(value)
}

// GetBackupPrevLineTextBaseOffset returns the value in BackupPrevLineTextBaseOffset.
func (p GuiComboPreviewData) GetBackupPrevLineTextBaseOffset() float32 {
	ptr := (*C.ImGuiComboPreviewData)(unsafe.Pointer(p))
	return float32(ptr.BackupPrevLineTextBaseOffset)
}

// SetBackupPrevLineTextBaseOffset sets the value in BackupPrevLineTextBaseOffset.
func (p GuiComboPreviewData) SetBackupPrevLineTextBaseOffset(value float32) {
	ptr := (*C.ImGuiComboPreviewData)(unsafe.Pointer(p))
	ptr.BackupPrevLineTextBaseOffset = (C.float)(value)
}

// RefPreviewRect returns pointer to the PreviewRect field.
func (p GuiComboPreviewData) RefPreviewRect() Rect {
	return Rect(p + GuiComboPreviewData(C.offsetof_ImGuiComboPreviewData_PreviewRect))
}

// GuiContext wraps struct ImGuiContext.
type GuiContext uintptr

// GuiContextNil is a null pointer.
var GuiContextNil GuiContext

// GuiContextSizeOf is the byte size of ImGuiContext.
const GuiContextSizeOf = int(C.sizeof_ImGuiContext)

// ImGuiContext allocator is suppressed.

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p GuiContext) Offset(offset int) GuiContext {
	return p + GuiContext(offset*GuiContextSizeOf)
}

// GuiContext.ActiveId is unsupported: category unsupported.

// GetActiveIdAllowOverlap returns the value in ActiveIdAllowOverlap.
func (p GuiContext) GetActiveIdAllowOverlap() bool {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	return bool(ptr.ActiveIdAllowOverlap)
}

// SetActiveIdAllowOverlap sets the value in ActiveIdAllowOverlap.
func (p GuiContext) SetActiveIdAllowOverlap(value bool) {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	ptr.ActiveIdAllowOverlap = (C.bool)(value)
}

// RefActiveIdClickOffset returns pointer to the ActiveIdClickOffset field.
func (p GuiContext) RefActiveIdClickOffset() Vec2 {
	return Vec2(p + GuiContext(C.offsetof_ImGuiContext_ActiveIdClickOffset))
}

// GuiContext.ActiveIdDisabledId is unsupported: category unsupported.

// GetActiveIdFromShortcut returns the value in ActiveIdFromShortcut.
func (p GuiContext) GetActiveIdFromShortcut() bool {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	return bool(ptr.ActiveIdFromShortcut)
}

// SetActiveIdFromShortcut sets the value in ActiveIdFromShortcut.
func (p GuiContext) SetActiveIdFromShortcut(value bool) {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	ptr.ActiveIdFromShortcut = (C.bool)(value)
}

// GetActiveIdHasBeenEditedBefore returns the value in ActiveIdHasBeenEditedBefore.
func (p GuiContext) GetActiveIdHasBeenEditedBefore() bool {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	return bool(ptr.ActiveIdHasBeenEditedBefore)
}

// SetActiveIdHasBeenEditedBefore sets the value in ActiveIdHasBeenEditedBefore.
func (p GuiContext) SetActiveIdHasBeenEditedBefore(value bool) {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	ptr.ActiveIdHasBeenEditedBefore = (C.bool)(value)
}

// GetActiveIdHasBeenEditedThisFrame returns the value in ActiveIdHasBeenEditedThisFrame.
func (p GuiContext) GetActiveIdHasBeenEditedThisFrame() bool {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	return bool(ptr.ActiveIdHasBeenEditedThisFrame)
}

// SetActiveIdHasBeenEditedThisFrame sets the value in ActiveIdHasBeenEditedThisFrame.
func (p GuiContext) SetActiveIdHasBeenEditedThisFrame(value bool) {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	ptr.ActiveIdHasBeenEditedThisFrame = (C.bool)(value)
}

// GetActiveIdHasBeenPressedBefore returns the value in ActiveIdHasBeenPressedBefore.
func (p GuiContext) GetActiveIdHasBeenPressedBefore() bool {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	return bool(ptr.ActiveIdHasBeenPressedBefore)
}

// SetActiveIdHasBeenPressedBefore sets the value in ActiveIdHasBeenPressedBefore.
func (p GuiContext) SetActiveIdHasBeenPressedBefore(value bool) {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	ptr.ActiveIdHasBeenPressedBefore = (C.bool)(value)
}

// GuiContext.ActiveIdIsAlive is unsupported: category unsupported.

// GetActiveIdIsJustActivated returns the value in ActiveIdIsJustActivated.
func (p GuiContext) GetActiveIdIsJustActivated() bool {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	return bool(ptr.ActiveIdIsJustActivated)
}

// SetActiveIdIsJustActivated sets the value in ActiveIdIsJustActivated.
func (p GuiContext) SetActiveIdIsJustActivated(value bool) {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	ptr.ActiveIdIsJustActivated = (C.bool)(value)
}

// GuiContext.ActiveIdMouseButton is unsupported: category unsupported.

// GetActiveIdNoClearOnFocusLoss returns the value in ActiveIdNoClearOnFocusLoss.
func (p GuiContext) GetActiveIdNoClearOnFocusLoss() bool {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	return bool(ptr.ActiveIdNoClearOnFocusLoss)
}

// SetActiveIdNoClearOnFocusLoss sets the value in ActiveIdNoClearOnFocusLoss.
func (p GuiContext) SetActiveIdNoClearOnFocusLoss(value bool) {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	ptr.ActiveIdNoClearOnFocusLoss = (C.bool)(value)
}

// GuiContext.ActiveIdPreviousFrame is unsupported: category unsupported.

// GetActiveIdSource returns the value in ActiveIdSource.
func (p GuiContext) GetActiveIdSource() GuiInputSource {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	return GuiInputSource(ptr.ActiveIdSource)
}

// SetActiveIdSource sets the value in ActiveIdSource.
func (p GuiContext) SetActiveIdSource(value GuiInputSource) {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	ptr.ActiveIdSource = (C.ImGuiInputSource)(value)
}

// GetActiveIdTimer returns the value in ActiveIdTimer.
func (p GuiContext) GetActiveIdTimer() float32 {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	return float32(ptr.ActiveIdTimer)
}

// SetActiveIdTimer sets the value in ActiveIdTimer.
func (p GuiContext) SetActiveIdTimer(value float32) {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	ptr.ActiveIdTimer = (C.float)(value)
}

// GetActiveIdUsingAllKeyboardKeys returns the value in ActiveIdUsingAllKeyboardKeys.
func (p GuiContext) GetActiveIdUsingAllKeyboardKeys() bool {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	return bool(ptr.ActiveIdUsingAllKeyboardKeys)
}

// SetActiveIdUsingAllKeyboardKeys sets the value in ActiveIdUsingAllKeyboardKeys.
func (p GuiContext) SetActiveIdUsingAllKeyboardKeys(value bool) {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	ptr.ActiveIdUsingAllKeyboardKeys = (C.bool)(value)
}

// GetActiveIdUsingNavDirMask returns the value in ActiveIdUsingNavDirMask.
func (p GuiContext) GetActiveIdUsingNavDirMask() uint32 {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	return uint32(ptr.ActiveIdUsingNavDirMask)
}

// SetActiveIdUsingNavDirMask sets the value in ActiveIdUsingNavDirMask.
func (p GuiContext) SetActiveIdUsingNavDirMask(value uint32) {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	ptr.ActiveIdUsingNavDirMask = (C.uint32_t)(value)
}

// RefActiveIdValueOnActivation returns pointer to the ActiveIdValueOnActivation field.
func (p GuiContext) RefActiveIdValueOnActivation() GuiDataTypeStorage {
	return GuiDataTypeStorage(p + GuiContext(C.offsetof_ImGuiContext_ActiveIdValueOnActivation))
}

// GetActiveIdWindow returns the value in ActiveIdWindow.
func (p GuiContext) GetActiveIdWindow() GuiWindow {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	return GuiWindow(unsafe.Pointer(ptr.ActiveIdWindow))
}

// SetActiveIdWindow sets the value in ActiveIdWindow.
func (p GuiContext) SetActiveIdWindow(value GuiWindow) {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	ptr.ActiveIdWindow = (*C.ImGuiWindow)(unsafe.Pointer(value))
}

// GetBeginComboDepth returns the value in BeginComboDepth.
func (p GuiContext) GetBeginComboDepth() int32 {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	return int32(ptr.BeginComboDepth)
}

// SetBeginComboDepth sets the value in BeginComboDepth.
func (p GuiContext) SetBeginComboDepth(value int32) {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	ptr.BeginComboDepth = (C.int32_t)(value)
}

// GetBeginMenuDepth returns the value in BeginMenuDepth.
func (p GuiContext) GetBeginMenuDepth() int32 {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	return int32(ptr.BeginMenuDepth)
}

// SetBeginMenuDepth sets the value in BeginMenuDepth.
func (p GuiContext) SetBeginMenuDepth(value int32) {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	ptr.BeginMenuDepth = (C.int32_t)(value)
}

// RefBeginPopupStack returns pointer to the BeginPopupStack field.
func (p GuiContext) RefBeginPopupStack() Vector_ImGuiPopupData {
	return Vector_ImGuiPopupData(p + GuiContext(C.offsetof_ImGuiContext_BeginPopupStack))
}

// RefBoxSelectState returns pointer to the BoxSelectState field.
func (p GuiContext) RefBoxSelectState() GuiBoxSelectState {
	return GuiBoxSelectState(p + GuiContext(C.offsetof_ImGuiContext_BoxSelectState))
}

// RefClipboardHandlerData returns pointer to the ClipboardHandlerData field.
func (p GuiContext) RefClipboardHandlerData() Vector_char {
	return Vector_char(p + GuiContext(C.offsetof_ImGuiContext_ClipboardHandlerData))
}

// RefClipperTempData returns pointer to the ClipperTempData field.
func (p GuiContext) RefClipperTempData() Vector_ImGuiListClipperData {
	return Vector_ImGuiListClipperData(p + GuiContext(C.offsetof_ImGuiContext_ClipperTempData))
}

// GetClipperTempDataStacked returns the value in ClipperTempDataStacked.
func (p GuiContext) GetClipperTempDataStacked() int32 {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	return int32(ptr.ClipperTempDataStacked)
}

// SetClipperTempDataStacked sets the value in ClipperTempDataStacked.
func (p GuiContext) SetClipperTempDataStacked(value int32) {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	ptr.ClipperTempDataStacked = (C.int32_t)(value)
}

// GuiContext.ColorEditCurrentID is unsupported: category unsupported.

// GetColorEditOptions returns the value in ColorEditOptions.
func (p GuiContext) GetColorEditOptions() GuiColorEditFlags {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	return GuiColorEditFlags(ptr.ColorEditOptions)
}

// SetColorEditOptions sets the value in ColorEditOptions.
func (p GuiContext) SetColorEditOptions(value GuiColorEditFlags) {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	ptr.ColorEditOptions = (C.ImGuiColorEditFlags)(value)
}

// GetColorEditSavedColor returns the value in ColorEditSavedColor.
func (p GuiContext) GetColorEditSavedColor() uint32 {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	return uint32(ptr.ColorEditSavedColor)
}

// SetColorEditSavedColor sets the value in ColorEditSavedColor.
func (p GuiContext) SetColorEditSavedColor(value uint32) {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	ptr.ColorEditSavedColor = (C.uint32_t)(value)
}

// GetColorEditSavedHue returns the value in ColorEditSavedHue.
func (p GuiContext) GetColorEditSavedHue() float32 {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	return float32(ptr.ColorEditSavedHue)
}

// SetColorEditSavedHue sets the value in ColorEditSavedHue.
func (p GuiContext) SetColorEditSavedHue(value float32) {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	ptr.ColorEditSavedHue = (C.float)(value)
}

// GuiContext.ColorEditSavedID is unsupported: category unsupported.

// GetColorEditSavedSat returns the value in ColorEditSavedSat.
func (p GuiContext) GetColorEditSavedSat() float32 {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	return float32(ptr.ColorEditSavedSat)
}

// SetColorEditSavedSat sets the value in ColorEditSavedSat.
func (p GuiContext) SetColorEditSavedSat(value float32) {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	ptr.ColorEditSavedSat = (C.float)(value)
}

// RefColorPickerRef returns pointer to the ColorPickerRef field.
func (p GuiContext) RefColorPickerRef() Vec4 {
	return Vec4(p + GuiContext(C.offsetof_ImGuiContext_ColorPickerRef))
}

// RefColorStack returns pointer to the ColorStack field.
func (p GuiContext) RefColorStack() Vector_ImGuiColorMod {
	return Vector_ImGuiColorMod(p + GuiContext(C.offsetof_ImGuiContext_ColorStack))
}

// RefComboPreviewData returns pointer to the ComboPreviewData field.
func (p GuiContext) RefComboPreviewData() GuiComboPreviewData {
	return GuiComboPreviewData(p + GuiContext(C.offsetof_ImGuiContext_ComboPreviewData))
}

// GetConfigFlagsCurrFrame returns the value in ConfigFlagsCurrFrame.
func (p GuiContext) GetConfigFlagsCurrFrame() GuiConfigFlags {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	return GuiConfigFlags(ptr.ConfigFlagsCurrFrame)
}

// SetConfigFlagsCurrFrame sets the value in ConfigFlagsCurrFrame.
func (p GuiContext) SetConfigFlagsCurrFrame(value GuiConfigFlags) {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	ptr.ConfigFlagsCurrFrame = (C.ImGuiConfigFlags)(value)
}

// GetConfigFlagsLastFrame returns the value in ConfigFlagsLastFrame.
func (p GuiContext) GetConfigFlagsLastFrame() GuiConfigFlags {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	return GuiConfigFlags(ptr.ConfigFlagsLastFrame)
}

// SetConfigFlagsLastFrame sets the value in ConfigFlagsLastFrame.
func (p GuiContext) SetConfigFlagsLastFrame(value GuiConfigFlags) {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	ptr.ConfigFlagsLastFrame = (C.ImGuiConfigFlags)(value)
}

// GuiContext.ConfigNavWindowingKeyNext is unsupported: category unsupported.

// GuiContext.ConfigNavWindowingKeyPrev is unsupported: category unsupported.

// GetConfigNavWindowingWithGamepad returns the value in ConfigNavWindowingWithGamepad.
func (p GuiContext) GetConfigNavWindowingWithGamepad() bool {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	return bool(ptr.ConfigNavWindowingWithGamepad)
}

// SetConfigNavWindowingWithGamepad sets the value in ConfigNavWindowingWithGamepad.
func (p GuiContext) SetConfigNavWindowingWithGamepad(value bool) {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	ptr.ConfigNavWindowingWithGamepad = (C.bool)(value)
}

// GuiContext.ContextName[16] is unsupported: category unsupported.

// GetCurrentDpiScale returns the value in CurrentDpiScale.
func (p GuiContext) GetCurrentDpiScale() float32 {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	return float32(ptr.CurrentDpiScale)
}

// SetCurrentDpiScale sets the value in CurrentDpiScale.
func (p GuiContext) SetCurrentDpiScale(value float32) {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	ptr.CurrentDpiScale = (C.float)(value)
}

// GuiContext.CurrentFocusScopeId is unsupported: category unsupported.

// GetCurrentItemFlags returns the value in CurrentItemFlags.
func (p GuiContext) GetCurrentItemFlags() GuiItemFlags {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	return GuiItemFlags(ptr.CurrentItemFlags)
}

// SetCurrentItemFlags sets the value in CurrentItemFlags.
func (p GuiContext) SetCurrentItemFlags(value GuiItemFlags) {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	ptr.CurrentItemFlags = (C.ImGuiItemFlags)(value)
}

// GetCurrentMultiSelect returns the value in CurrentMultiSelect.
func (p GuiContext) GetCurrentMultiSelect() GuiMultiSelectTempData {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	return GuiMultiSelectTempData(unsafe.Pointer(ptr.CurrentMultiSelect))
}

// SetCurrentMultiSelect sets the value in CurrentMultiSelect.
func (p GuiContext) SetCurrentMultiSelect(value GuiMultiSelectTempData) {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	ptr.CurrentMultiSelect = (*C.ImGuiMultiSelectTempData)(unsafe.Pointer(value))
}

// GetCurrentTabBar returns the value in CurrentTabBar.
func (p GuiContext) GetCurrentTabBar() GuiTabBar {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	return GuiTabBar(unsafe.Pointer(ptr.CurrentTabBar))
}

// SetCurrentTabBar sets the value in CurrentTabBar.
func (p GuiContext) SetCurrentTabBar(value GuiTabBar) {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	ptr.CurrentTabBar = (*C.ImGuiTabBar)(unsafe.Pointer(value))
}

// RefCurrentTabBarStack returns pointer to the CurrentTabBarStack field.
func (p GuiContext) RefCurrentTabBarStack() Vector_ImGuiPtrOrIndex {
	return Vector_ImGuiPtrOrIndex(p + GuiContext(C.offsetof_ImGuiContext_CurrentTabBarStack))
}

// GetCurrentTable returns the value in CurrentTable.
func (p GuiContext) GetCurrentTable() GuiTable {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	return GuiTable(unsafe.Pointer(ptr.CurrentTable))
}

// SetCurrentTable sets the value in CurrentTable.
func (p GuiContext) SetCurrentTable(value GuiTable) {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	ptr.CurrentTable = (*C.ImGuiTable)(unsafe.Pointer(value))
}

// GetCurrentViewport returns the value in CurrentViewport.
func (p GuiContext) GetCurrentViewport() GuiViewportP {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	return GuiViewportP(unsafe.Pointer(ptr.CurrentViewport))
}

// SetCurrentViewport sets the value in CurrentViewport.
func (p GuiContext) SetCurrentViewport(value GuiViewportP) {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	ptr.CurrentViewport = (*C.ImGuiViewportP)(unsafe.Pointer(value))
}

// GetCurrentWindow returns the value in CurrentWindow.
func (p GuiContext) GetCurrentWindow() GuiWindow {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	return GuiWindow(unsafe.Pointer(ptr.CurrentWindow))
}

// SetCurrentWindow sets the value in CurrentWindow.
func (p GuiContext) SetCurrentWindow(value GuiWindow) {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	ptr.CurrentWindow = (*C.ImGuiWindow)(unsafe.Pointer(value))
}

// RefCurrentWindowStack returns pointer to the CurrentWindowStack field.
func (p GuiContext) RefCurrentWindowStack() Vector_ImGuiWindowStackData {
	return Vector_ImGuiWindowStackData(p + GuiContext(C.offsetof_ImGuiContext_CurrentWindowStack))
}

// RefDataTypeZeroValue returns pointer to the DataTypeZeroValue field.
func (p GuiContext) RefDataTypeZeroValue() GuiDataTypeStorage {
	return GuiDataTypeStorage(p + GuiContext(C.offsetof_ImGuiContext_DataTypeZeroValue))
}

// RefDeactivatedItemData returns pointer to the DeactivatedItemData field.
func (p GuiContext) RefDeactivatedItemData() GuiDeactivatedItemData {
	return GuiDeactivatedItemData(p + GuiContext(C.offsetof_ImGuiContext_DeactivatedItemData))
}

// RefDebugAllocInfo returns pointer to the DebugAllocInfo field.
func (p GuiContext) RefDebugAllocInfo() GuiDebugAllocInfo {
	return GuiDebugAllocInfo(p + GuiContext(C.offsetof_ImGuiContext_DebugAllocInfo))
}

// GuiContext.DebugBeginReturnValueCullDepth is unsupported: category unsupported.

// GetDebugBreakInLocateId returns the value in DebugBreakInLocateId.
func (p GuiContext) GetDebugBreakInLocateId() bool {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	return bool(ptr.DebugBreakInLocateId)
}

// SetDebugBreakInLocateId sets the value in DebugBreakInLocateId.
func (p GuiContext) SetDebugBreakInLocateId(value bool) {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	ptr.DebugBreakInLocateId = (C.bool)(value)
}

// GuiContext.DebugBreakInShortcutRouting is unsupported: category unsupported.

// GuiContext.DebugBreakInTable is unsupported: category unsupported.

// GuiContext.DebugBreakInWindow is unsupported: category unsupported.

// GuiContext.DebugBreakKeyChord is unsupported: category unsupported.

// GetDebugDrawIdConflictsCount returns the value in DebugDrawIdConflictsCount.
func (p GuiContext) GetDebugDrawIdConflictsCount() int32 {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	return int32(ptr.DebugDrawIdConflictsCount)
}

// SetDebugDrawIdConflictsCount sets the value in DebugDrawIdConflictsCount.
func (p GuiContext) SetDebugDrawIdConflictsCount(value int32) {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	ptr.DebugDrawIdConflictsCount = (C.int32_t)(value)
}

// GuiContext.DebugDrawIdConflictsId is unsupported: category unsupported.

// RefDebugFlashStyleColorBackup returns pointer to the DebugFlashStyleColorBackup field.
func (p GuiContext) RefDebugFlashStyleColorBackup() Vec4 {
	return Vec4(p + GuiContext(C.offsetof_ImGuiContext_DebugFlashStyleColorBackup))
}

// GetDebugFlashStyleColorIdx returns the value in DebugFlashStyleColorIdx.
func (p GuiContext) GetDebugFlashStyleColorIdx() GuiCol {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	return GuiCol(ptr.DebugFlashStyleColorIdx)
}

// SetDebugFlashStyleColorIdx sets the value in DebugFlashStyleColorIdx.
func (p GuiContext) SetDebugFlashStyleColorIdx(value GuiCol) {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	ptr.DebugFlashStyleColorIdx = (C.ImGuiCol)(value)
}

// GetDebugFlashStyleColorTime returns the value in DebugFlashStyleColorTime.
func (p GuiContext) GetDebugFlashStyleColorTime() float32 {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	return float32(ptr.DebugFlashStyleColorTime)
}

// SetDebugFlashStyleColorTime sets the value in DebugFlashStyleColorTime.
func (p GuiContext) SetDebugFlashStyleColorTime(value float32) {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	ptr.DebugFlashStyleColorTime = (C.float)(value)
}

// GuiContext.DebugHookIdInfoId is unsupported: category unsupported.

// GetDebugHoveredDockNode returns the value in DebugHoveredDockNode.
func (p GuiContext) GetDebugHoveredDockNode() GuiDockNode {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	return GuiDockNode(unsafe.Pointer(ptr.DebugHoveredDockNode))
}

// SetDebugHoveredDockNode sets the value in DebugHoveredDockNode.
func (p GuiContext) SetDebugHoveredDockNode(value GuiDockNode) {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	ptr.DebugHoveredDockNode = (*C.ImGuiDockNode)(unsafe.Pointer(value))
}

// RefDebugIDStackTool returns pointer to the DebugIDStackTool field.
func (p GuiContext) RefDebugIDStackTool() GuiIDStackTool {
	return GuiIDStackTool(p + GuiContext(C.offsetof_ImGuiContext_DebugIDStackTool))
}

// GetDebugItemPickerActive returns the value in DebugItemPickerActive.
func (p GuiContext) GetDebugItemPickerActive() bool {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	return bool(ptr.DebugItemPickerActive)
}

// SetDebugItemPickerActive sets the value in DebugItemPickerActive.
func (p GuiContext) SetDebugItemPickerActive(value bool) {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	ptr.DebugItemPickerActive = (C.bool)(value)
}

// GuiContext.DebugItemPickerBreakId is unsupported: category unsupported.

// GetDebugItemPickerMouseButton returns the value in DebugItemPickerMouseButton.
func (p GuiContext) GetDebugItemPickerMouseButton() uint8 {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	return uint8(ptr.DebugItemPickerMouseButton)
}

// SetDebugItemPickerMouseButton sets the value in DebugItemPickerMouseButton.
func (p GuiContext) SetDebugItemPickerMouseButton(value uint8) {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	ptr.DebugItemPickerMouseButton = (C.uint8_t)(value)
}

// GetDebugLocateFrames returns the value in DebugLocateFrames.
func (p GuiContext) GetDebugLocateFrames() uint8 {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	return uint8(ptr.DebugLocateFrames)
}

// SetDebugLocateFrames sets the value in DebugLocateFrames.
func (p GuiContext) SetDebugLocateFrames(value uint8) {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	ptr.DebugLocateFrames = (C.uint8_t)(value)
}

// GuiContext.DebugLocateId is unsupported: category unsupported.

// GetDebugLogAutoDisableFlags returns the value in DebugLogAutoDisableFlags.
func (p GuiContext) GetDebugLogAutoDisableFlags() GuiDebugLogFlags {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	return GuiDebugLogFlags(ptr.DebugLogAutoDisableFlags)
}

// SetDebugLogAutoDisableFlags sets the value in DebugLogAutoDisableFlags.
func (p GuiContext) SetDebugLogAutoDisableFlags(value GuiDebugLogFlags) {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	ptr.DebugLogAutoDisableFlags = (C.ImGuiDebugLogFlags)(value)
}

// GetDebugLogAutoDisableFrames returns the value in DebugLogAutoDisableFrames.
func (p GuiContext) GetDebugLogAutoDisableFrames() uint8 {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	return uint8(ptr.DebugLogAutoDisableFrames)
}

// SetDebugLogAutoDisableFrames sets the value in DebugLogAutoDisableFrames.
func (p GuiContext) SetDebugLogAutoDisableFrames(value uint8) {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	ptr.DebugLogAutoDisableFrames = (C.uint8_t)(value)
}

// RefDebugLogBuf returns pointer to the DebugLogBuf field.
func (p GuiContext) RefDebugLogBuf() GuiTextBuffer {
	return GuiTextBuffer(p + GuiContext(C.offsetof_ImGuiContext_DebugLogBuf))
}

// GetDebugLogFlags returns the value in DebugLogFlags.
func (p GuiContext) GetDebugLogFlags() GuiDebugLogFlags {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	return GuiDebugLogFlags(ptr.DebugLogFlags)
}

// SetDebugLogFlags sets the value in DebugLogFlags.
func (p GuiContext) SetDebugLogFlags(value GuiDebugLogFlags) {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	ptr.DebugLogFlags = (C.ImGuiDebugLogFlags)(value)
}

// RefDebugLogIndex returns pointer to the DebugLogIndex field.
func (p GuiContext) RefDebugLogIndex() GuiTextIndex {
	return GuiTextIndex(p + GuiContext(C.offsetof_ImGuiContext_DebugLogIndex))
}

// GetDebugLogSkippedErrors returns the value in DebugLogSkippedErrors.
func (p GuiContext) GetDebugLogSkippedErrors() int32 {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	return int32(ptr.DebugLogSkippedErrors)
}

// SetDebugLogSkippedErrors sets the value in DebugLogSkippedErrors.
func (p GuiContext) SetDebugLogSkippedErrors(value int32) {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	ptr.DebugLogSkippedErrors = (C.int32_t)(value)
}

// RefDebugMetricsConfig returns pointer to the DebugMetricsConfig field.
func (p GuiContext) RefDebugMetricsConfig() GuiMetricsConfig {
	return GuiMetricsConfig(p + GuiContext(C.offsetof_ImGuiContext_DebugMetricsConfig))
}

// GetDebugShowGroupRects returns the value in DebugShowGroupRects.
func (p GuiContext) GetDebugShowGroupRects() bool {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	return bool(ptr.DebugShowGroupRects)
}

// SetDebugShowGroupRects sets the value in DebugShowGroupRects.
func (p GuiContext) SetDebugShowGroupRects(value bool) {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	ptr.DebugShowGroupRects = (C.bool)(value)
}

// GetDimBgRatio returns the value in DimBgRatio.
func (p GuiContext) GetDimBgRatio() float32 {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	return float32(ptr.DimBgRatio)
}

// SetDimBgRatio sets the value in DimBgRatio.
func (p GuiContext) SetDimBgRatio(value float32) {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	ptr.DimBgRatio = (C.float)(value)
}

// GetDisabledAlphaBackup returns the value in DisabledAlphaBackup.
func (p GuiContext) GetDisabledAlphaBackup() float32 {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	return float32(ptr.DisabledAlphaBackup)
}

// SetDisabledAlphaBackup sets the value in DisabledAlphaBackup.
func (p GuiContext) SetDisabledAlphaBackup(value float32) {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	ptr.DisabledAlphaBackup = (C.float)(value)
}

// GuiContext.DisabledStackSize is unsupported: category unsupported.

// RefDockContext returns pointer to the DockContext field.
func (p GuiContext) RefDockContext() GuiDockContext {
	return GuiDockContext(p + GuiContext(C.offsetof_ImGuiContext_DockContext))
}

// GuiContext.DockNodeWindowMenuHandler is unsupported: category unsupported.

// GetDragCurrentAccum returns the value in DragCurrentAccum.
func (p GuiContext) GetDragCurrentAccum() float32 {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	return float32(ptr.DragCurrentAccum)
}

// SetDragCurrentAccum sets the value in DragCurrentAccum.
func (p GuiContext) SetDragCurrentAccum(value float32) {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	ptr.DragCurrentAccum = (C.float)(value)
}

// GetDragCurrentAccumDirty returns the value in DragCurrentAccumDirty.
func (p GuiContext) GetDragCurrentAccumDirty() bool {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	return bool(ptr.DragCurrentAccumDirty)
}

// SetDragCurrentAccumDirty sets the value in DragCurrentAccumDirty.
func (p GuiContext) SetDragCurrentAccumDirty(value bool) {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	ptr.DragCurrentAccumDirty = (C.bool)(value)
}

// GetDragDropAcceptFlags returns the value in DragDropAcceptFlags.
func (p GuiContext) GetDragDropAcceptFlags() GuiDragDropFlags {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	return GuiDragDropFlags(ptr.DragDropAcceptFlags)
}

// SetDragDropAcceptFlags sets the value in DragDropAcceptFlags.
func (p GuiContext) SetDragDropAcceptFlags(value GuiDragDropFlags) {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	ptr.DragDropAcceptFlags = (C.ImGuiDragDropFlags)(value)
}

// GetDragDropAcceptFrameCount returns the value in DragDropAcceptFrameCount.
func (p GuiContext) GetDragDropAcceptFrameCount() int32 {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	return int32(ptr.DragDropAcceptFrameCount)
}

// SetDragDropAcceptFrameCount sets the value in DragDropAcceptFrameCount.
func (p GuiContext) SetDragDropAcceptFrameCount(value int32) {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	ptr.DragDropAcceptFrameCount = (C.int32_t)(value)
}

// GuiContext.DragDropAcceptIdCurr is unsupported: category unsupported.

// GetDragDropAcceptIdCurrRectSurface returns the value in DragDropAcceptIdCurrRectSurface.
func (p GuiContext) GetDragDropAcceptIdCurrRectSurface() float32 {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	return float32(ptr.DragDropAcceptIdCurrRectSurface)
}

// SetDragDropAcceptIdCurrRectSurface sets the value in DragDropAcceptIdCurrRectSurface.
func (p GuiContext) SetDragDropAcceptIdCurrRectSurface(value float32) {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	ptr.DragDropAcceptIdCurrRectSurface = (C.float)(value)
}

// GuiContext.DragDropAcceptIdPrev is unsupported: category unsupported.

// GetDragDropActive returns the value in DragDropActive.
func (p GuiContext) GetDragDropActive() bool {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	return bool(ptr.DragDropActive)
}

// SetDragDropActive sets the value in DragDropActive.
func (p GuiContext) SetDragDropActive(value bool) {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	ptr.DragDropActive = (C.bool)(value)
}

// GuiContext.DragDropHoldJustPressedId is unsupported: category unsupported.

// GetDragDropMouseButton returns the value in DragDropMouseButton.
func (p GuiContext) GetDragDropMouseButton() int32 {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	return int32(ptr.DragDropMouseButton)
}

// SetDragDropMouseButton sets the value in DragDropMouseButton.
func (p GuiContext) SetDragDropMouseButton(value int32) {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	ptr.DragDropMouseButton = (C.int32_t)(value)
}

// RefDragDropPayload returns pointer to the DragDropPayload field.
func (p GuiContext) RefDragDropPayload() GuiPayload {
	return GuiPayload(p + GuiContext(C.offsetof_ImGuiContext_DragDropPayload))
}

// GuiContext.DragDropPayloadBufHeap is unsupported: category unsupported.

// GuiContext.DragDropPayloadBufLocal[16] is unsupported: category unsupported.

// GetDragDropSourceFlags returns the value in DragDropSourceFlags.
func (p GuiContext) GetDragDropSourceFlags() GuiDragDropFlags {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	return GuiDragDropFlags(ptr.DragDropSourceFlags)
}

// SetDragDropSourceFlags sets the value in DragDropSourceFlags.
func (p GuiContext) SetDragDropSourceFlags(value GuiDragDropFlags) {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	ptr.DragDropSourceFlags = (C.ImGuiDragDropFlags)(value)
}

// GetDragDropSourceFrameCount returns the value in DragDropSourceFrameCount.
func (p GuiContext) GetDragDropSourceFrameCount() int32 {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	return int32(ptr.DragDropSourceFrameCount)
}

// SetDragDropSourceFrameCount sets the value in DragDropSourceFrameCount.
func (p GuiContext) SetDragDropSourceFrameCount(value int32) {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	ptr.DragDropSourceFrameCount = (C.int32_t)(value)
}

// RefDragDropTargetClipRect returns pointer to the DragDropTargetClipRect field.
func (p GuiContext) RefDragDropTargetClipRect() Rect {
	return Rect(p + GuiContext(C.offsetof_ImGuiContext_DragDropTargetClipRect))
}

// GuiContext.DragDropTargetFullViewport is unsupported: category unsupported.

// GuiContext.DragDropTargetId is unsupported: category unsupported.

// RefDragDropTargetRect returns pointer to the DragDropTargetRect field.
func (p GuiContext) RefDragDropTargetRect() Rect {
	return Rect(p + GuiContext(C.offsetof_ImGuiContext_DragDropTargetRect))
}

// GetDragDropWithinSource returns the value in DragDropWithinSource.
func (p GuiContext) GetDragDropWithinSource() bool {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	return bool(ptr.DragDropWithinSource)
}

// SetDragDropWithinSource sets the value in DragDropWithinSource.
func (p GuiContext) SetDragDropWithinSource(value bool) {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	ptr.DragDropWithinSource = (C.bool)(value)
}

// GetDragDropWithinTarget returns the value in DragDropWithinTarget.
func (p GuiContext) GetDragDropWithinTarget() bool {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	return bool(ptr.DragDropWithinTarget)
}

// SetDragDropWithinTarget sets the value in DragDropWithinTarget.
func (p GuiContext) SetDragDropWithinTarget(value bool) {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	ptr.DragDropWithinTarget = (C.bool)(value)
}

// GetDragSpeedDefaultRatio returns the value in DragSpeedDefaultRatio.
func (p GuiContext) GetDragSpeedDefaultRatio() float32 {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	return float32(ptr.DragSpeedDefaultRatio)
}

// SetDragSpeedDefaultRatio sets the value in DragSpeedDefaultRatio.
func (p GuiContext) SetDragSpeedDefaultRatio(value float32) {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	ptr.DragSpeedDefaultRatio = (C.float)(value)
}

// RefDrawChannelsTempMergeBuffer returns pointer to the DrawChannelsTempMergeBuffer field.
func (p GuiContext) RefDrawChannelsTempMergeBuffer() Vector_ImDrawChannel {
	return Vector_ImDrawChannel(p + GuiContext(C.offsetof_ImGuiContext_DrawChannelsTempMergeBuffer))
}

// RefDrawListSharedData returns pointer to the DrawListSharedData field.
func (p GuiContext) RefDrawListSharedData() DrawListSharedData {
	return DrawListSharedData(p + GuiContext(C.offsetof_ImGuiContext_DrawListSharedData))
}

// GuiContext.ErrorCallback is unsupported: category unsupported.

// GetErrorCallbackUserData returns the value in ErrorCallbackUserData.
func (p GuiContext) GetErrorCallbackUserData() uintptr {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	return uintptr(unsafe.Pointer(ptr.ErrorCallbackUserData))
}

// SetErrorCallbackUserData sets the value in ErrorCallbackUserData.
func (p GuiContext) SetErrorCallbackUserData(value uintptr) {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	ptr.ErrorCallbackUserData = unsafe.Pointer(value)
}

// GetErrorCountCurrentFrame returns the value in ErrorCountCurrentFrame.
func (p GuiContext) GetErrorCountCurrentFrame() int32 {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	return int32(ptr.ErrorCountCurrentFrame)
}

// SetErrorCountCurrentFrame sets the value in ErrorCountCurrentFrame.
func (p GuiContext) SetErrorCountCurrentFrame(value int32) {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	ptr.ErrorCountCurrentFrame = (C.int32_t)(value)
}

// GetErrorFirst returns the value in ErrorFirst.
func (p GuiContext) GetErrorFirst() bool {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	return bool(ptr.ErrorFirst)
}

// SetErrorFirst sets the value in ErrorFirst.
func (p GuiContext) SetErrorFirst(value bool) {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	ptr.ErrorFirst = (C.bool)(value)
}

// RefErrorTooltipLockedPos returns pointer to the ErrorTooltipLockedPos field.
func (p GuiContext) RefErrorTooltipLockedPos() Vec2 {
	return Vec2(p + GuiContext(C.offsetof_ImGuiContext_ErrorTooltipLockedPos))
}

// RefFallbackMonitor returns pointer to the FallbackMonitor field.
func (p GuiContext) RefFallbackMonitor() GuiPlatformMonitor {
	return GuiPlatformMonitor(p + GuiContext(C.offsetof_ImGuiContext_FallbackMonitor))
}

// RefFocusScopeStack returns pointer to the FocusScopeStack field.
func (p GuiContext) RefFocusScopeStack() Vector_ImGuiFocusScopeData {
	return Vector_ImGuiFocusScopeData(p + GuiContext(C.offsetof_ImGuiContext_FocusScopeStack))
}

// GetFont returns the value in Font.
func (p GuiContext) GetFont() Font {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	return Font(unsafe.Pointer(ptr.Font))
}

// SetFont sets the value in Font.
func (p GuiContext) SetFont(value Font) {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	ptr.Font = (*C.ImFont)(unsafe.Pointer(value))
}

// RefFontAtlases returns pointer to the FontAtlases field.
func (p GuiContext) RefFontAtlases() Vector_ImFontAtlasPtr {
	return Vector_ImFontAtlasPtr(p + GuiContext(C.offsetof_ImGuiContext_FontAtlases))
}

// GetFontBaked returns the value in FontBaked.
func (p GuiContext) GetFontBaked() FontBaked {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	return FontBaked(unsafe.Pointer(ptr.FontBaked))
}

// SetFontBaked sets the value in FontBaked.
func (p GuiContext) SetFontBaked(value FontBaked) {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	ptr.FontBaked = (*C.ImFontBaked)(unsafe.Pointer(value))
}

// GetFontBakedScale returns the value in FontBakedScale.
func (p GuiContext) GetFontBakedScale() float32 {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	return float32(ptr.FontBakedScale)
}

// SetFontBakedScale sets the value in FontBakedScale.
func (p GuiContext) SetFontBakedScale(value float32) {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	ptr.FontBakedScale = (C.float)(value)
}

// GetFontRasterizerDensity returns the value in FontRasterizerDensity.
func (p GuiContext) GetFontRasterizerDensity() float32 {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	return float32(ptr.FontRasterizerDensity)
}

// SetFontRasterizerDensity sets the value in FontRasterizerDensity.
func (p GuiContext) SetFontRasterizerDensity(value float32) {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	ptr.FontRasterizerDensity = (C.float)(value)
}

// GetFontSize returns the value in FontSize.
func (p GuiContext) GetFontSize() float32 {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	return float32(ptr.FontSize)
}

// SetFontSize sets the value in FontSize.
func (p GuiContext) SetFontSize(value float32) {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	ptr.FontSize = (C.float)(value)
}

// GetFontSizeBase returns the value in FontSizeBase.
func (p GuiContext) GetFontSizeBase() float32 {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	return float32(ptr.FontSizeBase)
}

// SetFontSizeBase sets the value in FontSizeBase.
func (p GuiContext) SetFontSizeBase(value float32) {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	ptr.FontSizeBase = (C.float)(value)
}

// RefFontStack returns pointer to the FontStack field.
func (p GuiContext) RefFontStack() Vector_ImFontStackData {
	return Vector_ImFontStackData(p + GuiContext(C.offsetof_ImGuiContext_FontStack))
}

// GetFrameCount returns the value in FrameCount.
func (p GuiContext) GetFrameCount() int32 {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	return int32(ptr.FrameCount)
}

// SetFrameCount sets the value in FrameCount.
func (p GuiContext) SetFrameCount(value int32) {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	ptr.FrameCount = (C.int32_t)(value)
}

// GetFrameCountEnded returns the value in FrameCountEnded.
func (p GuiContext) GetFrameCountEnded() int32 {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	return int32(ptr.FrameCountEnded)
}

// SetFrameCountEnded sets the value in FrameCountEnded.
func (p GuiContext) SetFrameCountEnded(value int32) {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	ptr.FrameCountEnded = (C.int32_t)(value)
}

// GetFrameCountPlatformEnded returns the value in FrameCountPlatformEnded.
func (p GuiContext) GetFrameCountPlatformEnded() int32 {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	return int32(ptr.FrameCountPlatformEnded)
}

// SetFrameCountPlatformEnded sets the value in FrameCountPlatformEnded.
func (p GuiContext) SetFrameCountPlatformEnded(value int32) {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	ptr.FrameCountPlatformEnded = (C.int32_t)(value)
}

// GetFrameCountRendered returns the value in FrameCountRendered.
func (p GuiContext) GetFrameCountRendered() int32 {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	return int32(ptr.FrameCountRendered)
}

// SetFrameCountRendered sets the value in FrameCountRendered.
func (p GuiContext) SetFrameCountRendered(value int32) {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	ptr.FrameCountRendered = (C.int32_t)(value)
}

// GetFramerateSecPerFrameAccum returns the value in FramerateSecPerFrameAccum.
func (p GuiContext) GetFramerateSecPerFrameAccum() float32 {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	return float32(ptr.FramerateSecPerFrameAccum)
}

// SetFramerateSecPerFrameAccum sets the value in FramerateSecPerFrameAccum.
func (p GuiContext) SetFramerateSecPerFrameAccum(value float32) {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	ptr.FramerateSecPerFrameAccum = (C.float)(value)
}

// GetFramerateSecPerFrameCount returns the value in FramerateSecPerFrameCount.
func (p GuiContext) GetFramerateSecPerFrameCount() int32 {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	return int32(ptr.FramerateSecPerFrameCount)
}

// SetFramerateSecPerFrameCount sets the value in FramerateSecPerFrameCount.
func (p GuiContext) SetFramerateSecPerFrameCount(value int32) {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	ptr.FramerateSecPerFrameCount = (C.int32_t)(value)
}

// GetFramerateSecPerFrameIdx returns the value in FramerateSecPerFrameIdx.
func (p GuiContext) GetFramerateSecPerFrameIdx() int32 {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	return int32(ptr.FramerateSecPerFrameIdx)
}

// SetFramerateSecPerFrameIdx sets the value in FramerateSecPerFrameIdx.
func (p GuiContext) SetFramerateSecPerFrameIdx(value int32) {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	ptr.FramerateSecPerFrameIdx = (C.int32_t)(value)
}

// GuiContext.FramerateSecPerFrame[60] is unsupported: category unsupported.

// GetGcCompactAll returns the value in GcCompactAll.
func (p GuiContext) GetGcCompactAll() bool {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	return bool(ptr.GcCompactAll)
}

// SetGcCompactAll sets the value in GcCompactAll.
func (p GuiContext) SetGcCompactAll(value bool) {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	ptr.GcCompactAll = (C.bool)(value)
}

// RefGroupStack returns pointer to the GroupStack field.
func (p GuiContext) RefGroupStack() Vector_ImGuiGroupData {
	return Vector_ImGuiGroupData(p + GuiContext(C.offsetof_ImGuiContext_GroupStack))
}

// GuiContext.HookIdNext is unsupported: category unsupported.

// RefHooks returns pointer to the Hooks field.
func (p GuiContext) RefHooks() Vector_ImGuiContextHook {
	return Vector_ImGuiContextHook(p + GuiContext(C.offsetof_ImGuiContext_Hooks))
}

// GetHoverItemDelayClearTimer returns the value in HoverItemDelayClearTimer.
func (p GuiContext) GetHoverItemDelayClearTimer() float32 {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	return float32(ptr.HoverItemDelayClearTimer)
}

// SetHoverItemDelayClearTimer sets the value in HoverItemDelayClearTimer.
func (p GuiContext) SetHoverItemDelayClearTimer(value float32) {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	ptr.HoverItemDelayClearTimer = (C.float)(value)
}

// GuiContext.HoverItemDelayId is unsupported: category unsupported.

// GuiContext.HoverItemDelayIdPreviousFrame is unsupported: category unsupported.

// GetHoverItemDelayTimer returns the value in HoverItemDelayTimer.
func (p GuiContext) GetHoverItemDelayTimer() float32 {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	return float32(ptr.HoverItemDelayTimer)
}

// SetHoverItemDelayTimer sets the value in HoverItemDelayTimer.
func (p GuiContext) SetHoverItemDelayTimer(value float32) {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	ptr.HoverItemDelayTimer = (C.float)(value)
}

// GuiContext.HoverItemUnlockedStationaryId is unsupported: category unsupported.

// GuiContext.HoverWindowUnlockedStationaryId is unsupported: category unsupported.

// GuiContext.HoveredId is unsupported: category unsupported.

// GetHoveredIdAllowOverlap returns the value in HoveredIdAllowOverlap.
func (p GuiContext) GetHoveredIdAllowOverlap() bool {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	return bool(ptr.HoveredIdAllowOverlap)
}

// SetHoveredIdAllowOverlap sets the value in HoveredIdAllowOverlap.
func (p GuiContext) SetHoveredIdAllowOverlap(value bool) {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	ptr.HoveredIdAllowOverlap = (C.bool)(value)
}

// GetHoveredIdIsDisabled returns the value in HoveredIdIsDisabled.
func (p GuiContext) GetHoveredIdIsDisabled() bool {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	return bool(ptr.HoveredIdIsDisabled)
}

// SetHoveredIdIsDisabled sets the value in HoveredIdIsDisabled.
func (p GuiContext) SetHoveredIdIsDisabled(value bool) {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	ptr.HoveredIdIsDisabled = (C.bool)(value)
}

// GetHoveredIdNotActiveTimer returns the value in HoveredIdNotActiveTimer.
func (p GuiContext) GetHoveredIdNotActiveTimer() float32 {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	return float32(ptr.HoveredIdNotActiveTimer)
}

// SetHoveredIdNotActiveTimer sets the value in HoveredIdNotActiveTimer.
func (p GuiContext) SetHoveredIdNotActiveTimer(value float32) {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	ptr.HoveredIdNotActiveTimer = (C.float)(value)
}

// GuiContext.HoveredIdPreviousFrame is unsupported: category unsupported.

// GetHoveredIdPreviousFrameItemCount returns the value in HoveredIdPreviousFrameItemCount.
func (p GuiContext) GetHoveredIdPreviousFrameItemCount() int32 {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	return int32(ptr.HoveredIdPreviousFrameItemCount)
}

// SetHoveredIdPreviousFrameItemCount sets the value in HoveredIdPreviousFrameItemCount.
func (p GuiContext) SetHoveredIdPreviousFrameItemCount(value int32) {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	ptr.HoveredIdPreviousFrameItemCount = (C.int32_t)(value)
}

// GetHoveredIdTimer returns the value in HoveredIdTimer.
func (p GuiContext) GetHoveredIdTimer() float32 {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	return float32(ptr.HoveredIdTimer)
}

// SetHoveredIdTimer sets the value in HoveredIdTimer.
func (p GuiContext) SetHoveredIdTimer(value float32) {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	ptr.HoveredIdTimer = (C.float)(value)
}

// GetHoveredWindow returns the value in HoveredWindow.
func (p GuiContext) GetHoveredWindow() GuiWindow {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	return GuiWindow(unsafe.Pointer(ptr.HoveredWindow))
}

// SetHoveredWindow sets the value in HoveredWindow.
func (p GuiContext) SetHoveredWindow(value GuiWindow) {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	ptr.HoveredWindow = (*C.ImGuiWindow)(unsafe.Pointer(value))
}

// GetHoveredWindowBeforeClear returns the value in HoveredWindowBeforeClear.
func (p GuiContext) GetHoveredWindowBeforeClear() GuiWindow {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	return GuiWindow(unsafe.Pointer(ptr.HoveredWindowBeforeClear))
}

// SetHoveredWindowBeforeClear sets the value in HoveredWindowBeforeClear.
func (p GuiContext) SetHoveredWindowBeforeClear(value GuiWindow) {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	ptr.HoveredWindowBeforeClear = (*C.ImGuiWindow)(unsafe.Pointer(value))
}

// GetHoveredWindowUnderMovingWindow returns the value in HoveredWindowUnderMovingWindow.
func (p GuiContext) GetHoveredWindowUnderMovingWindow() GuiWindow {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	return GuiWindow(unsafe.Pointer(ptr.HoveredWindowUnderMovingWindow))
}

// SetHoveredWindowUnderMovingWindow sets the value in HoveredWindowUnderMovingWindow.
func (p GuiContext) SetHoveredWindowUnderMovingWindow(value GuiWindow) {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	ptr.HoveredWindowUnderMovingWindow = (*C.ImGuiWindow)(unsafe.Pointer(value))
}

// RefIO returns pointer to the IO field.
func (p GuiContext) RefIO() GuiIO {
	return GuiIO(p + GuiContext(C.offsetof_ImGuiContext_IO))
}

// GetInitialized returns the value in Initialized.
func (p GuiContext) GetInitialized() bool {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	return bool(ptr.Initialized)
}

// SetInitialized sets the value in Initialized.
func (p GuiContext) SetInitialized(value bool) {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	ptr.Initialized = (C.bool)(value)
}

// GetInputEventsNextEventId returns the value in InputEventsNextEventId.
func (p GuiContext) GetInputEventsNextEventId() uint32 {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	return uint32(ptr.InputEventsNextEventId)
}

// SetInputEventsNextEventId sets the value in InputEventsNextEventId.
func (p GuiContext) SetInputEventsNextEventId(value uint32) {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	ptr.InputEventsNextEventId = (C.uint32_t)(value)
}

// GetInputEventsNextMouseSource returns the value in InputEventsNextMouseSource.
func (p GuiContext) GetInputEventsNextMouseSource() GuiMouseSource {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	return GuiMouseSource(ptr.InputEventsNextMouseSource)
}

// SetInputEventsNextMouseSource sets the value in InputEventsNextMouseSource.
func (p GuiContext) SetInputEventsNextMouseSource(value GuiMouseSource) {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	ptr.InputEventsNextMouseSource = (C.ImGuiMouseSource)(value)
}

// RefInputEventsQueue returns pointer to the InputEventsQueue field.
func (p GuiContext) RefInputEventsQueue() Vector_ImGuiInputEvent {
	return Vector_ImGuiInputEvent(p + GuiContext(C.offsetof_ImGuiContext_InputEventsQueue))
}

// RefInputEventsTrail returns pointer to the InputEventsTrail field.
func (p GuiContext) RefInputEventsTrail() Vector_ImGuiInputEvent {
	return Vector_ImGuiInputEvent(p + GuiContext(C.offsetof_ImGuiContext_InputEventsTrail))
}

// RefInputTextDeactivatedState returns pointer to the InputTextDeactivatedState field.
func (p GuiContext) RefInputTextDeactivatedState() GuiInputTextDeactivatedState {
	return GuiInputTextDeactivatedState(p + GuiContext(C.offsetof_ImGuiContext_InputTextDeactivatedState))
}

// RefInputTextLineIndex returns pointer to the InputTextLineIndex field.
func (p GuiContext) RefInputTextLineIndex() GuiTextIndex {
	return GuiTextIndex(p + GuiContext(C.offsetof_ImGuiContext_InputTextLineIndex))
}

// RefInputTextPasswordFontBackupBaked returns pointer to the InputTextPasswordFontBackupBaked field.
func (p GuiContext) RefInputTextPasswordFontBackupBaked() FontBaked {
	return FontBaked(p + GuiContext(C.offsetof_ImGuiContext_InputTextPasswordFontBackupBaked))
}

// GetInputTextPasswordFontBackupFlags returns the value in InputTextPasswordFontBackupFlags.
func (p GuiContext) GetInputTextPasswordFontBackupFlags() FontFlags {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	return FontFlags(ptr.InputTextPasswordFontBackupFlags)
}

// SetInputTextPasswordFontBackupFlags sets the value in InputTextPasswordFontBackupFlags.
func (p GuiContext) SetInputTextPasswordFontBackupFlags(value FontFlags) {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	ptr.InputTextPasswordFontBackupFlags = (C.ImFontFlags)(value)
}

// RefInputTextState returns pointer to the InputTextState field.
func (p GuiContext) RefInputTextState() GuiInputTextState {
	return GuiInputTextState(p + GuiContext(C.offsetof_ImGuiContext_InputTextState))
}

// RefItemFlagsStack returns pointer to the ItemFlagsStack field.
func (p GuiContext) RefItemFlagsStack() Vector_ImGuiItemFlags {
	return Vector_ImGuiItemFlags(p + GuiContext(C.offsetof_ImGuiContext_ItemFlagsStack))
}

// GetItemUnclipByLog returns the value in ItemUnclipByLog.
func (p GuiContext) GetItemUnclipByLog() bool {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	return bool(ptr.ItemUnclipByLog)
}

// SetItemUnclipByLog sets the value in ItemUnclipByLog.
func (p GuiContext) SetItemUnclipByLog(value bool) {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	ptr.ItemUnclipByLog = (C.bool)(value)
}

// GuiContext.KeysMayBeCharInput is unsupported: category unsupported.

// GuiContext.KeysOwnerData[ImGuiKey_NamedKey_COUNT] is unsupported: category unsupported.

// RefKeysRoutingTable returns pointer to the KeysRoutingTable field.
func (p GuiContext) RefKeysRoutingTable() GuiKeyRoutingTable {
	return GuiKeyRoutingTable(p + GuiContext(C.offsetof_ImGuiContext_KeysRoutingTable))
}

// GuiContext.LastActiveId is unsupported: category unsupported.

// GetLastActiveIdTimer returns the value in LastActiveIdTimer.
func (p GuiContext) GetLastActiveIdTimer() float32 {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	return float32(ptr.LastActiveIdTimer)
}

// SetLastActiveIdTimer sets the value in LastActiveIdTimer.
func (p GuiContext) SetLastActiveIdTimer(value float32) {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	ptr.LastActiveIdTimer = (C.float)(value)
}

// RefLastItemData returns pointer to the LastItemData field.
func (p GuiContext) RefLastItemData() GuiLastItemData {
	return GuiLastItemData(p + GuiContext(C.offsetof_ImGuiContext_LastItemData))
}

// GuiContext.LastKeyModsChangeFromNoneTime is unsupported: category unsupported.

// GuiContext.LastKeyModsChangeTime is unsupported: category unsupported.

// GuiContext.LastKeyboardKeyPressTime is unsupported: category unsupported.

// GuiContext.LocalizationTable[ImGuiLocKey_COUNT] is unsupported: category unsupported.

// RefLogBuffer returns pointer to the LogBuffer field.
func (p GuiContext) RefLogBuffer() GuiTextBuffer {
	return GuiTextBuffer(p + GuiContext(C.offsetof_ImGuiContext_LogBuffer))
}

// GetLogDepthRef returns the value in LogDepthRef.
func (p GuiContext) GetLogDepthRef() int32 {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	return int32(ptr.LogDepthRef)
}

// SetLogDepthRef sets the value in LogDepthRef.
func (p GuiContext) SetLogDepthRef(value int32) {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	ptr.LogDepthRef = (C.int32_t)(value)
}

// GetLogDepthToExpand returns the value in LogDepthToExpand.
func (p GuiContext) GetLogDepthToExpand() int32 {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	return int32(ptr.LogDepthToExpand)
}

// SetLogDepthToExpand sets the value in LogDepthToExpand.
func (p GuiContext) SetLogDepthToExpand(value int32) {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	ptr.LogDepthToExpand = (C.int32_t)(value)
}

// GetLogDepthToExpandDefault returns the value in LogDepthToExpandDefault.
func (p GuiContext) GetLogDepthToExpandDefault() int32 {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	return int32(ptr.LogDepthToExpandDefault)
}

// SetLogDepthToExpandDefault sets the value in LogDepthToExpandDefault.
func (p GuiContext) SetLogDepthToExpandDefault(value int32) {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	ptr.LogDepthToExpandDefault = (C.int32_t)(value)
}

// GetLogEnabled returns the value in LogEnabled.
func (p GuiContext) GetLogEnabled() bool {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	return bool(ptr.LogEnabled)
}

// SetLogEnabled sets the value in LogEnabled.
func (p GuiContext) SetLogEnabled(value bool) {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	ptr.LogEnabled = (C.bool)(value)
}

// GuiContext.LogFile is unsupported: category unsupported.

// GetLogFlags returns the value in LogFlags.
func (p GuiContext) GetLogFlags() GuiLogFlags {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	return GuiLogFlags(ptr.LogFlags)
}

// SetLogFlags sets the value in LogFlags.
func (p GuiContext) SetLogFlags(value GuiLogFlags) {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	ptr.LogFlags = (C.ImGuiLogFlags)(value)
}

// GetLogLineFirstItem returns the value in LogLineFirstItem.
func (p GuiContext) GetLogLineFirstItem() bool {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	return bool(ptr.LogLineFirstItem)
}

// SetLogLineFirstItem sets the value in LogLineFirstItem.
func (p GuiContext) SetLogLineFirstItem(value bool) {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	ptr.LogLineFirstItem = (C.bool)(value)
}

// GetLogLinePosY returns the value in LogLinePosY.
func (p GuiContext) GetLogLinePosY() float32 {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	return float32(ptr.LogLinePosY)
}

// SetLogLinePosY sets the value in LogLinePosY.
func (p GuiContext) SetLogLinePosY(value float32) {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	ptr.LogLinePosY = (C.float)(value)
}

// GetLogNextPrefix returns the value in LogNextPrefix.
func (p GuiContext) GetLogNextPrefix() ffi.CString {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	return ffi.CStringFromPtr((unsafe.Pointer)(ptr.LogNextPrefix))
}

// SetLogNextPrefix sets the value in LogNextPrefix.
func (p GuiContext) SetLogNextPrefix(value ffi.CString) {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	ptr.LogNextPrefix = (*C.char)(value.Raw())
}

// GetLogNextSuffix returns the value in LogNextSuffix.
func (p GuiContext) GetLogNextSuffix() ffi.CString {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	return ffi.CStringFromPtr((unsafe.Pointer)(ptr.LogNextSuffix))
}

// SetLogNextSuffix sets the value in LogNextSuffix.
func (p GuiContext) SetLogNextSuffix(value ffi.CString) {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	ptr.LogNextSuffix = (*C.char)(value.Raw())
}

// GetLogWindow returns the value in LogWindow.
func (p GuiContext) GetLogWindow() GuiWindow {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	return GuiWindow(unsafe.Pointer(ptr.LogWindow))
}

// SetLogWindow sets the value in LogWindow.
func (p GuiContext) SetLogWindow(value GuiWindow) {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	ptr.LogWindow = (*C.ImGuiWindow)(unsafe.Pointer(value))
}

// RefMenusIdSubmittedThisFrame returns pointer to the MenusIdSubmittedThisFrame field.
func (p GuiContext) RefMenusIdSubmittedThisFrame() Vector_ImGuiID {
	return Vector_ImGuiID(p + GuiContext(C.offsetof_ImGuiContext_MenusIdSubmittedThisFrame))
}

// GetMouseCursor returns the value in MouseCursor.
func (p GuiContext) GetMouseCursor() GuiMouseCursor {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	return GuiMouseCursor(ptr.MouseCursor)
}

// SetMouseCursor sets the value in MouseCursor.
func (p GuiContext) SetMouseCursor(value GuiMouseCursor) {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	ptr.MouseCursor = (C.ImGuiMouseCursor)(value)
}

// GetMouseLastHoveredViewport returns the value in MouseLastHoveredViewport.
func (p GuiContext) GetMouseLastHoveredViewport() GuiViewportP {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	return GuiViewportP(unsafe.Pointer(ptr.MouseLastHoveredViewport))
}

// SetMouseLastHoveredViewport sets the value in MouseLastHoveredViewport.
func (p GuiContext) SetMouseLastHoveredViewport(value GuiViewportP) {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	ptr.MouseLastHoveredViewport = (*C.ImGuiViewportP)(unsafe.Pointer(value))
}

// RefMouseLastValidPos returns pointer to the MouseLastValidPos field.
func (p GuiContext) RefMouseLastValidPos() Vec2 {
	return Vec2(p + GuiContext(C.offsetof_ImGuiContext_MouseLastValidPos))
}

// GetMouseStationaryTimer returns the value in MouseStationaryTimer.
func (p GuiContext) GetMouseStationaryTimer() float32 {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	return float32(ptr.MouseStationaryTimer)
}

// SetMouseStationaryTimer sets the value in MouseStationaryTimer.
func (p GuiContext) SetMouseStationaryTimer(value float32) {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	ptr.MouseStationaryTimer = (C.float)(value)
}

// GetMouseViewport returns the value in MouseViewport.
func (p GuiContext) GetMouseViewport() GuiViewportP {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	return GuiViewportP(unsafe.Pointer(ptr.MouseViewport))
}

// SetMouseViewport sets the value in MouseViewport.
func (p GuiContext) SetMouseViewport(value GuiViewportP) {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	ptr.MouseViewport = (*C.ImGuiViewportP)(unsafe.Pointer(value))
}

// GetMovingWindow returns the value in MovingWindow.
func (p GuiContext) GetMovingWindow() GuiWindow {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	return GuiWindow(unsafe.Pointer(ptr.MovingWindow))
}

// SetMovingWindow sets the value in MovingWindow.
func (p GuiContext) SetMovingWindow(value GuiWindow) {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	ptr.MovingWindow = (*C.ImGuiWindow)(unsafe.Pointer(value))
}

// RefMultiSelectStorage returns pointer to the MultiSelectStorage field.
func (p GuiContext) RefMultiSelectStorage() Pool_ImGuiMultiSelectState {
	return Pool_ImGuiMultiSelectState(p + GuiContext(C.offsetof_ImGuiContext_MultiSelectStorage))
}

// RefMultiSelectTempData returns pointer to the MultiSelectTempData field.
func (p GuiContext) RefMultiSelectTempData() Vector_ImGuiMultiSelectTempData {
	return Vector_ImGuiMultiSelectTempData(p + GuiContext(C.offsetof_ImGuiContext_MultiSelectTempData))
}

// GetMultiSelectTempDataStacked returns the value in MultiSelectTempDataStacked.
func (p GuiContext) GetMultiSelectTempDataStacked() int32 {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	return int32(ptr.MultiSelectTempDataStacked)
}

// SetMultiSelectTempDataStacked sets the value in MultiSelectTempDataStacked.
func (p GuiContext) SetMultiSelectTempDataStacked(value int32) {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	ptr.MultiSelectTempDataStacked = (C.int32_t)(value)
}

// GuiContext.NavActivateDownId is unsupported: category unsupported.

// GetNavActivateFlags returns the value in NavActivateFlags.
func (p GuiContext) GetNavActivateFlags() GuiActivateFlags {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	return GuiActivateFlags(ptr.NavActivateFlags)
}

// SetNavActivateFlags sets the value in NavActivateFlags.
func (p GuiContext) SetNavActivateFlags(value GuiActivateFlags) {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	ptr.NavActivateFlags = (C.ImGuiActivateFlags)(value)
}

// GuiContext.NavActivateId is unsupported: category unsupported.

// GuiContext.NavActivatePressedId is unsupported: category unsupported.

// GetNavAnyRequest returns the value in NavAnyRequest.
func (p GuiContext) GetNavAnyRequest() bool {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	return bool(ptr.NavAnyRequest)
}

// SetNavAnyRequest sets the value in NavAnyRequest.
func (p GuiContext) SetNavAnyRequest(value bool) {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	ptr.NavAnyRequest = (C.bool)(value)
}

// GuiContext.NavCursorHideFrames is unsupported: category unsupported.

// GetNavCursorVisible returns the value in NavCursorVisible.
func (p GuiContext) GetNavCursorVisible() bool {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	return bool(ptr.NavCursorVisible)
}

// SetNavCursorVisible sets the value in NavCursorVisible.
func (p GuiContext) SetNavCursorVisible(value bool) {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	ptr.NavCursorVisible = (C.bool)(value)
}

// RefNavFocusRoute returns pointer to the NavFocusRoute field.
func (p GuiContext) RefNavFocusRoute() Vector_ImGuiFocusScopeData {
	return Vector_ImGuiFocusScopeData(p + GuiContext(C.offsetof_ImGuiContext_NavFocusRoute))
}

// GuiContext.NavFocusScopeId is unsupported: category unsupported.

// GuiContext.NavHighlightActivatedId is unsupported: category unsupported.

// GetNavHighlightActivatedTimer returns the value in NavHighlightActivatedTimer.
func (p GuiContext) GetNavHighlightActivatedTimer() float32 {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	return float32(ptr.NavHighlightActivatedTimer)
}

// SetNavHighlightActivatedTimer sets the value in NavHighlightActivatedTimer.
func (p GuiContext) SetNavHighlightActivatedTimer(value float32) {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	ptr.NavHighlightActivatedTimer = (C.float)(value)
}

// GetNavHighlightItemUnderNav returns the value in NavHighlightItemUnderNav.
func (p GuiContext) GetNavHighlightItemUnderNav() bool {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	return bool(ptr.NavHighlightItemUnderNav)
}

// SetNavHighlightItemUnderNav sets the value in NavHighlightItemUnderNav.
func (p GuiContext) SetNavHighlightItemUnderNav(value bool) {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	ptr.NavHighlightItemUnderNav = (C.bool)(value)
}

// GuiContext.NavId is unsupported: category unsupported.

// GetNavIdIsAlive returns the value in NavIdIsAlive.
func (p GuiContext) GetNavIdIsAlive() bool {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	return bool(ptr.NavIdIsAlive)
}

// SetNavIdIsAlive sets the value in NavIdIsAlive.
func (p GuiContext) SetNavIdIsAlive(value bool) {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	ptr.NavIdIsAlive = (C.bool)(value)
}

// GetNavInitRequest returns the value in NavInitRequest.
func (p GuiContext) GetNavInitRequest() bool {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	return bool(ptr.NavInitRequest)
}

// SetNavInitRequest sets the value in NavInitRequest.
func (p GuiContext) SetNavInitRequest(value bool) {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	ptr.NavInitRequest = (C.bool)(value)
}

// GetNavInitRequestFromMove returns the value in NavInitRequestFromMove.
func (p GuiContext) GetNavInitRequestFromMove() bool {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	return bool(ptr.NavInitRequestFromMove)
}

// SetNavInitRequestFromMove sets the value in NavInitRequestFromMove.
func (p GuiContext) SetNavInitRequestFromMove(value bool) {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	ptr.NavInitRequestFromMove = (C.bool)(value)
}

// RefNavInitResult returns pointer to the NavInitResult field.
func (p GuiContext) RefNavInitResult() GuiNavItemData {
	return GuiNavItemData(p + GuiContext(C.offsetof_ImGuiContext_NavInitResult))
}

// GetNavInputSource returns the value in NavInputSource.
func (p GuiContext) GetNavInputSource() GuiInputSource {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	return GuiInputSource(ptr.NavInputSource)
}

// SetNavInputSource sets the value in NavInputSource.
func (p GuiContext) SetNavInputSource(value GuiInputSource) {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	ptr.NavInputSource = (C.ImGuiInputSource)(value)
}

// GuiContext.NavJustMovedFromFocusScopeId is unsupported: category unsupported.

// GuiContext.NavJustMovedToFocusScopeId is unsupported: category unsupported.

// GetNavJustMovedToHasSelectionData returns the value in NavJustMovedToHasSelectionData.
func (p GuiContext) GetNavJustMovedToHasSelectionData() bool {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	return bool(ptr.NavJustMovedToHasSelectionData)
}

// SetNavJustMovedToHasSelectionData sets the value in NavJustMovedToHasSelectionData.
func (p GuiContext) SetNavJustMovedToHasSelectionData(value bool) {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	ptr.NavJustMovedToHasSelectionData = (C.bool)(value)
}

// GuiContext.NavJustMovedToId is unsupported: category unsupported.

// GetNavJustMovedToIsTabbing returns the value in NavJustMovedToIsTabbing.
func (p GuiContext) GetNavJustMovedToIsTabbing() bool {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	return bool(ptr.NavJustMovedToIsTabbing)
}

// SetNavJustMovedToIsTabbing sets the value in NavJustMovedToIsTabbing.
func (p GuiContext) SetNavJustMovedToIsTabbing(value bool) {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	ptr.NavJustMovedToIsTabbing = (C.bool)(value)
}

// GuiContext.NavJustMovedToKeyMods is unsupported: category unsupported.

// GuiContext.NavLastValidSelectionUserData is unsupported: category unsupported.

// GetNavLayer returns the value in NavLayer.
func (p GuiContext) GetNavLayer() GuiNavLayer {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	return GuiNavLayer(ptr.NavLayer)
}

// SetNavLayer sets the value in NavLayer.
func (p GuiContext) SetNavLayer(value GuiNavLayer) {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	ptr.NavLayer = (C.ImGuiNavLayer)(value)
}

// GetNavMousePosDirty returns the value in NavMousePosDirty.
func (p GuiContext) GetNavMousePosDirty() bool {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	return bool(ptr.NavMousePosDirty)
}

// SetNavMousePosDirty sets the value in NavMousePosDirty.
func (p GuiContext) SetNavMousePosDirty(value bool) {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	ptr.NavMousePosDirty = (C.bool)(value)
}

// GetNavMoveClipDir returns the value in NavMoveClipDir.
func (p GuiContext) GetNavMoveClipDir() GuiDir {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	return GuiDir(ptr.NavMoveClipDir)
}

// SetNavMoveClipDir sets the value in NavMoveClipDir.
func (p GuiContext) SetNavMoveClipDir(value GuiDir) {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	ptr.NavMoveClipDir = (C.ImGuiDir)(value)
}

// GetNavMoveDir returns the value in NavMoveDir.
func (p GuiContext) GetNavMoveDir() GuiDir {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	return GuiDir(ptr.NavMoveDir)
}

// SetNavMoveDir sets the value in NavMoveDir.
func (p GuiContext) SetNavMoveDir(value GuiDir) {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	ptr.NavMoveDir = (C.ImGuiDir)(value)
}

// GetNavMoveDirForDebug returns the value in NavMoveDirForDebug.
func (p GuiContext) GetNavMoveDirForDebug() GuiDir {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	return GuiDir(ptr.NavMoveDirForDebug)
}

// SetNavMoveDirForDebug sets the value in NavMoveDirForDebug.
func (p GuiContext) SetNavMoveDirForDebug(value GuiDir) {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	ptr.NavMoveDirForDebug = (C.ImGuiDir)(value)
}

// GetNavMoveFlags returns the value in NavMoveFlags.
func (p GuiContext) GetNavMoveFlags() GuiNavMoveFlags {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	return GuiNavMoveFlags(ptr.NavMoveFlags)
}

// SetNavMoveFlags sets the value in NavMoveFlags.
func (p GuiContext) SetNavMoveFlags(value GuiNavMoveFlags) {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	ptr.NavMoveFlags = (C.ImGuiNavMoveFlags)(value)
}

// GetNavMoveForwardToNextFrame returns the value in NavMoveForwardToNextFrame.
func (p GuiContext) GetNavMoveForwardToNextFrame() bool {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	return bool(ptr.NavMoveForwardToNextFrame)
}

// SetNavMoveForwardToNextFrame sets the value in NavMoveForwardToNextFrame.
func (p GuiContext) SetNavMoveForwardToNextFrame(value bool) {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	ptr.NavMoveForwardToNextFrame = (C.bool)(value)
}

// GuiContext.NavMoveKeyMods is unsupported: category unsupported.

// RefNavMoveResultLocal returns pointer to the NavMoveResultLocal field.
func (p GuiContext) RefNavMoveResultLocal() GuiNavItemData {
	return GuiNavItemData(p + GuiContext(C.offsetof_ImGuiContext_NavMoveResultLocal))
}

// RefNavMoveResultLocalVisible returns pointer to the NavMoveResultLocalVisible field.
func (p GuiContext) RefNavMoveResultLocalVisible() GuiNavItemData {
	return GuiNavItemData(p + GuiContext(C.offsetof_ImGuiContext_NavMoveResultLocalVisible))
}

// RefNavMoveResultOther returns pointer to the NavMoveResultOther field.
func (p GuiContext) RefNavMoveResultOther() GuiNavItemData {
	return GuiNavItemData(p + GuiContext(C.offsetof_ImGuiContext_NavMoveResultOther))
}

// GetNavMoveScoringItems returns the value in NavMoveScoringItems.
func (p GuiContext) GetNavMoveScoringItems() bool {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	return bool(ptr.NavMoveScoringItems)
}

// SetNavMoveScoringItems sets the value in NavMoveScoringItems.
func (p GuiContext) SetNavMoveScoringItems(value bool) {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	ptr.NavMoveScoringItems = (C.bool)(value)
}

// GetNavMoveScrollFlags returns the value in NavMoveScrollFlags.
func (p GuiContext) GetNavMoveScrollFlags() GuiScrollFlags {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	return GuiScrollFlags(ptr.NavMoveScrollFlags)
}

// SetNavMoveScrollFlags sets the value in NavMoveScrollFlags.
func (p GuiContext) SetNavMoveScrollFlags(value GuiScrollFlags) {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	ptr.NavMoveScrollFlags = (C.ImGuiScrollFlags)(value)
}

// GetNavMoveSubmitted returns the value in NavMoveSubmitted.
func (p GuiContext) GetNavMoveSubmitted() bool {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	return bool(ptr.NavMoveSubmitted)
}

// SetNavMoveSubmitted sets the value in NavMoveSubmitted.
func (p GuiContext) SetNavMoveSubmitted(value bool) {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	ptr.NavMoveSubmitted = (C.bool)(value)
}

// GetNavNextActivateFlags returns the value in NavNextActivateFlags.
func (p GuiContext) GetNavNextActivateFlags() GuiActivateFlags {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	return GuiActivateFlags(ptr.NavNextActivateFlags)
}

// SetNavNextActivateFlags sets the value in NavNextActivateFlags.
func (p GuiContext) SetNavNextActivateFlags(value GuiActivateFlags) {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	ptr.NavNextActivateFlags = (C.ImGuiActivateFlags)(value)
}

// GuiContext.NavNextActivateId is unsupported: category unsupported.

// GetNavScoringDebugCount returns the value in NavScoringDebugCount.
func (p GuiContext) GetNavScoringDebugCount() int32 {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	return int32(ptr.NavScoringDebugCount)
}

// SetNavScoringDebugCount sets the value in NavScoringDebugCount.
func (p GuiContext) SetNavScoringDebugCount(value int32) {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	ptr.NavScoringDebugCount = (C.int32_t)(value)
}

// RefNavScoringNoClipRect returns pointer to the NavScoringNoClipRect field.
func (p GuiContext) RefNavScoringNoClipRect() Rect {
	return Rect(p + GuiContext(C.offsetof_ImGuiContext_NavScoringNoClipRect))
}

// RefNavScoringRect returns pointer to the NavScoringRect field.
func (p GuiContext) RefNavScoringRect() Rect {
	return Rect(p + GuiContext(C.offsetof_ImGuiContext_NavScoringRect))
}

// GetNavTabbingCounter returns the value in NavTabbingCounter.
func (p GuiContext) GetNavTabbingCounter() int32 {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	return int32(ptr.NavTabbingCounter)
}

// SetNavTabbingCounter sets the value in NavTabbingCounter.
func (p GuiContext) SetNavTabbingCounter(value int32) {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	ptr.NavTabbingCounter = (C.int32_t)(value)
}

// GetNavTabbingDir returns the value in NavTabbingDir.
func (p GuiContext) GetNavTabbingDir() int32 {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	return int32(ptr.NavTabbingDir)
}

// SetNavTabbingDir sets the value in NavTabbingDir.
func (p GuiContext) SetNavTabbingDir(value int32) {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	ptr.NavTabbingDir = (C.int32_t)(value)
}

// RefNavTabbingResultFirst returns pointer to the NavTabbingResultFirst field.
func (p GuiContext) RefNavTabbingResultFirst() GuiNavItemData {
	return GuiNavItemData(p + GuiContext(C.offsetof_ImGuiContext_NavTabbingResultFirst))
}

// GetNavWindow returns the value in NavWindow.
func (p GuiContext) GetNavWindow() GuiWindow {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	return GuiWindow(unsafe.Pointer(ptr.NavWindow))
}

// SetNavWindow sets the value in NavWindow.
func (p GuiContext) SetNavWindow(value GuiWindow) {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	ptr.NavWindow = (*C.ImGuiWindow)(unsafe.Pointer(value))
}

// RefNavWindowingAccumDeltaPos returns pointer to the NavWindowingAccumDeltaPos field.
func (p GuiContext) RefNavWindowingAccumDeltaPos() Vec2 {
	return Vec2(p + GuiContext(C.offsetof_ImGuiContext_NavWindowingAccumDeltaPos))
}

// RefNavWindowingAccumDeltaSize returns pointer to the NavWindowingAccumDeltaSize field.
func (p GuiContext) RefNavWindowingAccumDeltaSize() Vec2 {
	return Vec2(p + GuiContext(C.offsetof_ImGuiContext_NavWindowingAccumDeltaSize))
}

// GetNavWindowingHighlightAlpha returns the value in NavWindowingHighlightAlpha.
func (p GuiContext) GetNavWindowingHighlightAlpha() float32 {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	return float32(ptr.NavWindowingHighlightAlpha)
}

// SetNavWindowingHighlightAlpha sets the value in NavWindowingHighlightAlpha.
func (p GuiContext) SetNavWindowingHighlightAlpha(value float32) {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	ptr.NavWindowingHighlightAlpha = (C.float)(value)
}

// GetNavWindowingInputSource returns the value in NavWindowingInputSource.
func (p GuiContext) GetNavWindowingInputSource() GuiInputSource {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	return GuiInputSource(ptr.NavWindowingInputSource)
}

// SetNavWindowingInputSource sets the value in NavWindowingInputSource.
func (p GuiContext) SetNavWindowingInputSource(value GuiInputSource) {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	ptr.NavWindowingInputSource = (C.ImGuiInputSource)(value)
}

// GetNavWindowingListWindow returns the value in NavWindowingListWindow.
func (p GuiContext) GetNavWindowingListWindow() GuiWindow {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	return GuiWindow(unsafe.Pointer(ptr.NavWindowingListWindow))
}

// SetNavWindowingListWindow sets the value in NavWindowingListWindow.
func (p GuiContext) SetNavWindowingListWindow(value GuiWindow) {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	ptr.NavWindowingListWindow = (*C.ImGuiWindow)(unsafe.Pointer(value))
}

// GetNavWindowingTarget returns the value in NavWindowingTarget.
func (p GuiContext) GetNavWindowingTarget() GuiWindow {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	return GuiWindow(unsafe.Pointer(ptr.NavWindowingTarget))
}

// SetNavWindowingTarget sets the value in NavWindowingTarget.
func (p GuiContext) SetNavWindowingTarget(value GuiWindow) {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	ptr.NavWindowingTarget = (*C.ImGuiWindow)(unsafe.Pointer(value))
}

// GetNavWindowingTargetAnim returns the value in NavWindowingTargetAnim.
func (p GuiContext) GetNavWindowingTargetAnim() GuiWindow {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	return GuiWindow(unsafe.Pointer(ptr.NavWindowingTargetAnim))
}

// SetNavWindowingTargetAnim sets the value in NavWindowingTargetAnim.
func (p GuiContext) SetNavWindowingTargetAnim(value GuiWindow) {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	ptr.NavWindowingTargetAnim = (*C.ImGuiWindow)(unsafe.Pointer(value))
}

// GetNavWindowingTimer returns the value in NavWindowingTimer.
func (p GuiContext) GetNavWindowingTimer() float32 {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	return float32(ptr.NavWindowingTimer)
}

// SetNavWindowingTimer sets the value in NavWindowingTimer.
func (p GuiContext) SetNavWindowingTimer(value float32) {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	ptr.NavWindowingTimer = (C.float)(value)
}

// GetNavWindowingToggleKey returns the value in NavWindowingToggleKey.
func (p GuiContext) GetNavWindowingToggleKey() GuiKey {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	return GuiKey(ptr.NavWindowingToggleKey)
}

// SetNavWindowingToggleKey sets the value in NavWindowingToggleKey.
func (p GuiContext) SetNavWindowingToggleKey(value GuiKey) {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	ptr.NavWindowingToggleKey = (C.ImGuiKey)(value)
}

// GetNavWindowingToggleLayer returns the value in NavWindowingToggleLayer.
func (p GuiContext) GetNavWindowingToggleLayer() bool {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	return bool(ptr.NavWindowingToggleLayer)
}

// SetNavWindowingToggleLayer sets the value in NavWindowingToggleLayer.
func (p GuiContext) SetNavWindowingToggleLayer(value bool) {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	ptr.NavWindowingToggleLayer = (C.bool)(value)
}

// RefNextItemData returns pointer to the NextItemData field.
func (p GuiContext) RefNextItemData() GuiNextItemData {
	return GuiNextItemData(p + GuiContext(C.offsetof_ImGuiContext_NextItemData))
}

// RefNextWindowData returns pointer to the NextWindowData field.
func (p GuiContext) RefNextWindowData() GuiNextWindowData {
	return GuiNextWindowData(p + GuiContext(C.offsetof_ImGuiContext_NextWindowData))
}

// RefOpenPopupStack returns pointer to the OpenPopupStack field.
func (p GuiContext) RefOpenPopupStack() Vector_ImGuiPopupData {
	return Vector_ImGuiPopupData(p + GuiContext(C.offsetof_ImGuiContext_OpenPopupStack))
}

// RefPlatformIO returns pointer to the PlatformIO field.
func (p GuiContext) RefPlatformIO() GuiPlatformIO {
	return GuiPlatformIO(p + GuiContext(C.offsetof_ImGuiContext_PlatformIO))
}

// RefPlatformImeData returns pointer to the PlatformImeData field.
func (p GuiContext) RefPlatformImeData() GuiPlatformImeData {
	return GuiPlatformImeData(p + GuiContext(C.offsetof_ImGuiContext_PlatformImeData))
}

// RefPlatformImeDataPrev returns pointer to the PlatformImeDataPrev field.
func (p GuiContext) RefPlatformImeDataPrev() GuiPlatformImeData {
	return GuiPlatformImeData(p + GuiContext(C.offsetof_ImGuiContext_PlatformImeDataPrev))
}

// GuiContext.PlatformLastFocusedViewportId is unsupported: category unsupported.

// RefPlatformMonitorsFullWorkRect returns pointer to the PlatformMonitorsFullWorkRect field.
func (p GuiContext) RefPlatformMonitorsFullWorkRect() Rect {
	return Rect(p + GuiContext(C.offsetof_ImGuiContext_PlatformMonitorsFullWorkRect))
}

// GetPlatformWindowsCreatedCount returns the value in PlatformWindowsCreatedCount.
func (p GuiContext) GetPlatformWindowsCreatedCount() int32 {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	return int32(ptr.PlatformWindowsCreatedCount)
}

// SetPlatformWindowsCreatedCount sets the value in PlatformWindowsCreatedCount.
func (p GuiContext) SetPlatformWindowsCreatedCount(value int32) {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	ptr.PlatformWindowsCreatedCount = (C.int32_t)(value)
}

// GetScrollbarClickDeltaToGrabCenter returns the value in ScrollbarClickDeltaToGrabCenter.
func (p GuiContext) GetScrollbarClickDeltaToGrabCenter() float32 {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	return float32(ptr.ScrollbarClickDeltaToGrabCenter)
}

// SetScrollbarClickDeltaToGrabCenter sets the value in ScrollbarClickDeltaToGrabCenter.
func (p GuiContext) SetScrollbarClickDeltaToGrabCenter(value float32) {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	ptr.ScrollbarClickDeltaToGrabCenter = (C.float)(value)
}

// GuiContext.ScrollbarSeekMode is unsupported: category unsupported.

// GetSettingsDirtyTimer returns the value in SettingsDirtyTimer.
func (p GuiContext) GetSettingsDirtyTimer() float32 {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	return float32(ptr.SettingsDirtyTimer)
}

// SetSettingsDirtyTimer sets the value in SettingsDirtyTimer.
func (p GuiContext) SetSettingsDirtyTimer(value float32) {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	ptr.SettingsDirtyTimer = (C.float)(value)
}

// RefSettingsHandlers returns pointer to the SettingsHandlers field.
func (p GuiContext) RefSettingsHandlers() Vector_ImGuiSettingsHandler {
	return Vector_ImGuiSettingsHandler(p + GuiContext(C.offsetof_ImGuiContext_SettingsHandlers))
}

// RefSettingsIniData returns pointer to the SettingsIniData field.
func (p GuiContext) RefSettingsIniData() GuiTextBuffer {
	return GuiTextBuffer(p + GuiContext(C.offsetof_ImGuiContext_SettingsIniData))
}

// GetSettingsLoaded returns the value in SettingsLoaded.
func (p GuiContext) GetSettingsLoaded() bool {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	return bool(ptr.SettingsLoaded)
}

// SetSettingsLoaded sets the value in SettingsLoaded.
func (p GuiContext) SetSettingsLoaded(value bool) {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	ptr.SettingsLoaded = (C.bool)(value)
}

// RefSettingsTables returns pointer to the SettingsTables field.
func (p GuiContext) RefSettingsTables() ChunkStream_ImGuiTableSettings {
	return ChunkStream_ImGuiTableSettings(p + GuiContext(C.offsetof_ImGuiContext_SettingsTables))
}

// RefSettingsWindows returns pointer to the SettingsWindows field.
func (p GuiContext) RefSettingsWindows() ChunkStream_ImGuiWindowSettings {
	return ChunkStream_ImGuiWindowSettings(p + GuiContext(C.offsetof_ImGuiContext_SettingsWindows))
}

// RefShrinkWidthBuffer returns pointer to the ShrinkWidthBuffer field.
func (p GuiContext) RefShrinkWidthBuffer() Vector_ImGuiShrinkWidthItem {
	return Vector_ImGuiShrinkWidthItem(p + GuiContext(C.offsetof_ImGuiContext_ShrinkWidthBuffer))
}

// GetSliderCurrentAccum returns the value in SliderCurrentAccum.
func (p GuiContext) GetSliderCurrentAccum() float32 {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	return float32(ptr.SliderCurrentAccum)
}

// SetSliderCurrentAccum sets the value in SliderCurrentAccum.
func (p GuiContext) SetSliderCurrentAccum(value float32) {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	ptr.SliderCurrentAccum = (C.float)(value)
}

// GetSliderCurrentAccumDirty returns the value in SliderCurrentAccumDirty.
func (p GuiContext) GetSliderCurrentAccumDirty() bool {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	return bool(ptr.SliderCurrentAccumDirty)
}

// SetSliderCurrentAccumDirty sets the value in SliderCurrentAccumDirty.
func (p GuiContext) SetSliderCurrentAccumDirty(value bool) {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	ptr.SliderCurrentAccumDirty = (C.bool)(value)
}

// GetSliderGrabClickOffset returns the value in SliderGrabClickOffset.
func (p GuiContext) GetSliderGrabClickOffset() float32 {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	return float32(ptr.SliderGrabClickOffset)
}

// SetSliderGrabClickOffset sets the value in SliderGrabClickOffset.
func (p GuiContext) SetSliderGrabClickOffset(value float32) {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	ptr.SliderGrabClickOffset = (C.float)(value)
}

// GetStackSizesInBeginForCurrentWindow returns the value in StackSizesInBeginForCurrentWindow.
func (p GuiContext) GetStackSizesInBeginForCurrentWindow() GuiErrorRecoveryState {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	return GuiErrorRecoveryState(unsafe.Pointer(ptr.StackSizesInBeginForCurrentWindow))
}

// SetStackSizesInBeginForCurrentWindow sets the value in StackSizesInBeginForCurrentWindow.
func (p GuiContext) SetStackSizesInBeginForCurrentWindow(value GuiErrorRecoveryState) {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	ptr.StackSizesInBeginForCurrentWindow = (*C.ImGuiErrorRecoveryState)(unsafe.Pointer(value))
}

// RefStackSizesInNewFrame returns pointer to the StackSizesInNewFrame field.
func (p GuiContext) RefStackSizesInNewFrame() GuiErrorRecoveryState {
	return GuiErrorRecoveryState(p + GuiContext(C.offsetof_ImGuiContext_StackSizesInNewFrame))
}

// RefStyle returns pointer to the Style field.
func (p GuiContext) RefStyle() GuiStyle {
	return GuiStyle(p + GuiContext(C.offsetof_ImGuiContext_Style))
}

// RefStyleVarStack returns pointer to the StyleVarStack field.
func (p GuiContext) RefStyleVarStack() Vector_ImGuiStyleMod {
	return Vector_ImGuiStyleMod(p + GuiContext(C.offsetof_ImGuiContext_StyleVarStack))
}

// RefTabBars returns pointer to the TabBars field.
func (p GuiContext) RefTabBars() Pool_ImGuiTabBar {
	return Pool_ImGuiTabBar(p + GuiContext(C.offsetof_ImGuiContext_TabBars))
}

// RefTables returns pointer to the Tables field.
func (p GuiContext) RefTables() Pool_ImGuiTable {
	return Pool_ImGuiTable(p + GuiContext(C.offsetof_ImGuiContext_Tables))
}

// RefTablesLastTimeActive returns pointer to the TablesLastTimeActive field.
func (p GuiContext) RefTablesLastTimeActive() Vector_float {
	return Vector_float(p + GuiContext(C.offsetof_ImGuiContext_TablesLastTimeActive))
}

// RefTablesTempData returns pointer to the TablesTempData field.
func (p GuiContext) RefTablesTempData() Vector_ImGuiTableTempData {
	return Vector_ImGuiTableTempData(p + GuiContext(C.offsetof_ImGuiContext_TablesTempData))
}

// GetTablesTempDataStacked returns the value in TablesTempDataStacked.
func (p GuiContext) GetTablesTempDataStacked() int32 {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	return int32(ptr.TablesTempDataStacked)
}

// SetTablesTempDataStacked sets the value in TablesTempDataStacked.
func (p GuiContext) SetTablesTempDataStacked(value int32) {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	ptr.TablesTempDataStacked = (C.int32_t)(value)
}

// RefTempBuffer returns pointer to the TempBuffer field.
func (p GuiContext) RefTempBuffer() Vector_char {
	return Vector_char(p + GuiContext(C.offsetof_ImGuiContext_TempBuffer))
}

// GuiContext.TempInputId is unsupported: category unsupported.

// GuiContext.TempKeychordName[64] is unsupported: category unsupported.

// GetTestEngine returns the value in TestEngine.
func (p GuiContext) GetTestEngine() uintptr {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	return uintptr(unsafe.Pointer(ptr.TestEngine))
}

// SetTestEngine sets the value in TestEngine.
func (p GuiContext) SetTestEngine(value uintptr) {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	ptr.TestEngine = unsafe.Pointer(value)
}

// GetTestEngineHookItems returns the value in TestEngineHookItems.
func (p GuiContext) GetTestEngineHookItems() bool {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	return bool(ptr.TestEngineHookItems)
}

// SetTestEngineHookItems sets the value in TestEngineHookItems.
func (p GuiContext) SetTestEngineHookItems(value bool) {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	ptr.TestEngineHookItems = (C.bool)(value)
}

// GuiContext.Time is unsupported: category unsupported.

// GuiContext.TooltipOverrideCount is unsupported: category unsupported.

// GetTooltipPreviousWindow returns the value in TooltipPreviousWindow.
func (p GuiContext) GetTooltipPreviousWindow() GuiWindow {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	return GuiWindow(unsafe.Pointer(ptr.TooltipPreviousWindow))
}

// SetTooltipPreviousWindow sets the value in TooltipPreviousWindow.
func (p GuiContext) SetTooltipPreviousWindow(value GuiWindow) {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	ptr.TooltipPreviousWindow = (*C.ImGuiWindow)(unsafe.Pointer(value))
}

// RefTreeNodeStack returns pointer to the TreeNodeStack field.
func (p GuiContext) RefTreeNodeStack() Vector_ImGuiTreeNodeStackData {
	return Vector_ImGuiTreeNodeStackData(p + GuiContext(C.offsetof_ImGuiContext_TreeNodeStack))
}

// RefTypingSelectState returns pointer to the TypingSelectState field.
func (p GuiContext) RefTypingSelectState() GuiTypingSelectState {
	return GuiTypingSelectState(p + GuiContext(C.offsetof_ImGuiContext_TypingSelectState))
}

// RefUserTextures returns pointer to the UserTextures field.
func (p GuiContext) RefUserTextures() Vector_ImTextureDataPtr {
	return Vector_ImTextureDataPtr(p + GuiContext(C.offsetof_ImGuiContext_UserTextures))
}

// GetViewportCreatedCount returns the value in ViewportCreatedCount.
func (p GuiContext) GetViewportCreatedCount() int32 {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	return int32(ptr.ViewportCreatedCount)
}

// SetViewportCreatedCount sets the value in ViewportCreatedCount.
func (p GuiContext) SetViewportCreatedCount(value int32) {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	ptr.ViewportCreatedCount = (C.int32_t)(value)
}

// GetViewportFocusedStampCount returns the value in ViewportFocusedStampCount.
func (p GuiContext) GetViewportFocusedStampCount() int32 {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	return int32(ptr.ViewportFocusedStampCount)
}

// SetViewportFocusedStampCount sets the value in ViewportFocusedStampCount.
func (p GuiContext) SetViewportFocusedStampCount(value int32) {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	ptr.ViewportFocusedStampCount = (C.int32_t)(value)
}

// RefViewports returns pointer to the Viewports field.
func (p GuiContext) RefViewports() Vector_ImGuiViewportPPtr {
	return Vector_ImGuiViewportPPtr(p + GuiContext(C.offsetof_ImGuiContext_Viewports))
}

// GetWantCaptureKeyboardNextFrame returns the value in WantCaptureKeyboardNextFrame.
func (p GuiContext) GetWantCaptureKeyboardNextFrame() int32 {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	return int32(ptr.WantCaptureKeyboardNextFrame)
}

// SetWantCaptureKeyboardNextFrame sets the value in WantCaptureKeyboardNextFrame.
func (p GuiContext) SetWantCaptureKeyboardNextFrame(value int32) {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	ptr.WantCaptureKeyboardNextFrame = (C.int32_t)(value)
}

// GetWantCaptureMouseNextFrame returns the value in WantCaptureMouseNextFrame.
func (p GuiContext) GetWantCaptureMouseNextFrame() int32 {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	return int32(ptr.WantCaptureMouseNextFrame)
}

// SetWantCaptureMouseNextFrame sets the value in WantCaptureMouseNextFrame.
func (p GuiContext) SetWantCaptureMouseNextFrame(value int32) {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	ptr.WantCaptureMouseNextFrame = (C.int32_t)(value)
}

// GetWantTextInputNextFrame returns the value in WantTextInputNextFrame.
func (p GuiContext) GetWantTextInputNextFrame() int32 {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	return int32(ptr.WantTextInputNextFrame)
}

// SetWantTextInputNextFrame sets the value in WantTextInputNextFrame.
func (p GuiContext) SetWantTextInputNextFrame(value int32) {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	ptr.WantTextInputNextFrame = (C.int32_t)(value)
}

// RefWheelingAxisAvg returns pointer to the WheelingAxisAvg field.
func (p GuiContext) RefWheelingAxisAvg() Vec2 {
	return Vec2(p + GuiContext(C.offsetof_ImGuiContext_WheelingAxisAvg))
}

// GetWheelingWindow returns the value in WheelingWindow.
func (p GuiContext) GetWheelingWindow() GuiWindow {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	return GuiWindow(unsafe.Pointer(ptr.WheelingWindow))
}

// SetWheelingWindow sets the value in WheelingWindow.
func (p GuiContext) SetWheelingWindow(value GuiWindow) {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	ptr.WheelingWindow = (*C.ImGuiWindow)(unsafe.Pointer(value))
}

// RefWheelingWindowRefMousePos returns pointer to the WheelingWindowRefMousePos field.
func (p GuiContext) RefWheelingWindowRefMousePos() Vec2 {
	return Vec2(p + GuiContext(C.offsetof_ImGuiContext_WheelingWindowRefMousePos))
}

// GetWheelingWindowReleaseTimer returns the value in WheelingWindowReleaseTimer.
func (p GuiContext) GetWheelingWindowReleaseTimer() float32 {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	return float32(ptr.WheelingWindowReleaseTimer)
}

// SetWheelingWindowReleaseTimer sets the value in WheelingWindowReleaseTimer.
func (p GuiContext) SetWheelingWindowReleaseTimer(value float32) {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	ptr.WheelingWindowReleaseTimer = (C.float)(value)
}

// GetWheelingWindowScrolledFrame returns the value in WheelingWindowScrolledFrame.
func (p GuiContext) GetWheelingWindowScrolledFrame() int32 {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	return int32(ptr.WheelingWindowScrolledFrame)
}

// SetWheelingWindowScrolledFrame sets the value in WheelingWindowScrolledFrame.
func (p GuiContext) SetWheelingWindowScrolledFrame(value int32) {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	ptr.WheelingWindowScrolledFrame = (C.int32_t)(value)
}

// GetWheelingWindowStartFrame returns the value in WheelingWindowStartFrame.
func (p GuiContext) GetWheelingWindowStartFrame() int32 {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	return int32(ptr.WheelingWindowStartFrame)
}

// SetWheelingWindowStartFrame sets the value in WheelingWindowStartFrame.
func (p GuiContext) SetWheelingWindowStartFrame(value int32) {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	ptr.WheelingWindowStartFrame = (C.int32_t)(value)
}

// RefWheelingWindowWheelRemainder returns pointer to the WheelingWindowWheelRemainder field.
func (p GuiContext) RefWheelingWindowWheelRemainder() Vec2 {
	return Vec2(p + GuiContext(C.offsetof_ImGuiContext_WheelingWindowWheelRemainder))
}

// RefWindowResizeBorderExpectedRect returns pointer to the WindowResizeBorderExpectedRect field.
func (p GuiContext) RefWindowResizeBorderExpectedRect() Rect {
	return Rect(p + GuiContext(C.offsetof_ImGuiContext_WindowResizeBorderExpectedRect))
}

// GetWindowResizeRelativeMode returns the value in WindowResizeRelativeMode.
func (p GuiContext) GetWindowResizeRelativeMode() bool {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	return bool(ptr.WindowResizeRelativeMode)
}

// SetWindowResizeRelativeMode sets the value in WindowResizeRelativeMode.
func (p GuiContext) SetWindowResizeRelativeMode(value bool) {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	ptr.WindowResizeRelativeMode = (C.bool)(value)
}

// RefWindows returns pointer to the Windows field.
func (p GuiContext) RefWindows() Vector_ImGuiWindowPtr {
	return Vector_ImGuiWindowPtr(p + GuiContext(C.offsetof_ImGuiContext_Windows))
}

// GetWindowsActiveCount returns the value in WindowsActiveCount.
func (p GuiContext) GetWindowsActiveCount() int32 {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	return int32(ptr.WindowsActiveCount)
}

// SetWindowsActiveCount sets the value in WindowsActiveCount.
func (p GuiContext) SetWindowsActiveCount(value int32) {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	ptr.WindowsActiveCount = (C.int32_t)(value)
}

// GetWindowsBorderHoverPadding returns the value in WindowsBorderHoverPadding.
func (p GuiContext) GetWindowsBorderHoverPadding() float32 {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	return float32(ptr.WindowsBorderHoverPadding)
}

// SetWindowsBorderHoverPadding sets the value in WindowsBorderHoverPadding.
func (p GuiContext) SetWindowsBorderHoverPadding(value float32) {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	ptr.WindowsBorderHoverPadding = (C.float)(value)
}

// RefWindowsById returns pointer to the WindowsById field.
func (p GuiContext) RefWindowsById() GuiStorage {
	return GuiStorage(p + GuiContext(C.offsetof_ImGuiContext_WindowsById))
}

// RefWindowsFocusOrder returns pointer to the WindowsFocusOrder field.
func (p GuiContext) RefWindowsFocusOrder() Vector_ImGuiWindowPtr {
	return Vector_ImGuiWindowPtr(p + GuiContext(C.offsetof_ImGuiContext_WindowsFocusOrder))
}

// RefWindowsTempSortBuffer returns pointer to the WindowsTempSortBuffer field.
func (p GuiContext) RefWindowsTempSortBuffer() Vector_ImGuiWindowPtr {
	return Vector_ImGuiWindowPtr(p + GuiContext(C.offsetof_ImGuiContext_WindowsTempSortBuffer))
}

// GuiContext.WithinEndChildID is unsupported: category unsupported.

// GetWithinFrameScope returns the value in WithinFrameScope.
func (p GuiContext) GetWithinFrameScope() bool {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	return bool(ptr.WithinFrameScope)
}

// SetWithinFrameScope sets the value in WithinFrameScope.
func (p GuiContext) SetWithinFrameScope(value bool) {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	ptr.WithinFrameScope = (C.bool)(value)
}

// GetWithinFrameScopeWithImplicitWindow returns the value in WithinFrameScopeWithImplicitWindow.
func (p GuiContext) GetWithinFrameScopeWithImplicitWindow() bool {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	return bool(ptr.WithinFrameScopeWithImplicitWindow)
}

// SetWithinFrameScopeWithImplicitWindow sets the value in WithinFrameScopeWithImplicitWindow.
func (p GuiContext) SetWithinFrameScopeWithImplicitWindow(value bool) {
	ptr := (*C.ImGuiContext)(unsafe.Pointer(p))
	ptr.WithinFrameScopeWithImplicitWindow = (C.bool)(value)
}

// GuiContextHook wraps struct ImGuiContextHook.
type GuiContextHook uintptr

// GuiContextHookNil is a null pointer.
var GuiContextHookNil GuiContextHook

// GuiContextHookSizeOf is the byte size of ImGuiContextHook.
const GuiContextHookSizeOf = int(C.sizeof_ImGuiContextHook)

// ImGuiContextHook allocator is suppressed.

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p GuiContextHook) Offset(offset int) GuiContextHook {
	return p + GuiContextHook(offset*GuiContextHookSizeOf)
}

// GuiContextHook.Callback is unsupported: category unsupported.

// GuiContextHook.HookId is unsupported: category unsupported.

// GuiContextHook.Owner is unsupported: category unsupported.

// GetType returns the value in Type.
func (p GuiContextHook) GetType() GuiContextHookType {
	ptr := (*C.ImGuiContextHook)(unsafe.Pointer(p))
	return GuiContextHookType(ptr.Type)
}

// SetType sets the value in Type.
func (p GuiContextHook) SetType(value GuiContextHookType) {
	ptr := (*C.ImGuiContextHook)(unsafe.Pointer(p))
	ptr.Type = (C.ImGuiContextHookType)(value)
}

// GetUserData returns the value in UserData.
func (p GuiContextHook) GetUserData() uintptr {
	ptr := (*C.ImGuiContextHook)(unsafe.Pointer(p))
	return uintptr(unsafe.Pointer(ptr.UserData))
}

// SetUserData sets the value in UserData.
func (p GuiContextHook) SetUserData(value uintptr) {
	ptr := (*C.ImGuiContextHook)(unsafe.Pointer(p))
	ptr.UserData = unsafe.Pointer(value)
}

// GuiDataTypeInfo wraps struct ImGuiDataTypeInfo.
type GuiDataTypeInfo uintptr

// GuiDataTypeInfoNil is a null pointer.
var GuiDataTypeInfoNil GuiDataTypeInfo

// GuiDataTypeInfoSizeOf is the byte size of ImGuiDataTypeInfo.
const GuiDataTypeInfoSizeOf = int(C.sizeof_ImGuiDataTypeInfo)

// GuiDataTypeInfoAlloc allocates a continuous block of GuiDataTypeInfo.
func GuiDataTypeInfoAlloc(alloc ffi.Allocator, count int) GuiDataTypeInfo {
	ptr := alloc.Allocate(GuiDataTypeInfoSizeOf * count)
	return GuiDataTypeInfo(ptr)
}

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p GuiDataTypeInfo) Offset(offset int) GuiDataTypeInfo {
	return p + GuiDataTypeInfo(offset*GuiDataTypeInfoSizeOf)
}

// GetName returns the value in Name.
func (p GuiDataTypeInfo) GetName() ffi.CString {
	ptr := (*C.ImGuiDataTypeInfo)(unsafe.Pointer(p))
	return ffi.CStringFromPtr((unsafe.Pointer)(ptr.Name))
}

// SetName sets the value in Name.
func (p GuiDataTypeInfo) SetName(value ffi.CString) {
	ptr := (*C.ImGuiDataTypeInfo)(unsafe.Pointer(p))
	ptr.Name = (*C.char)(value.Raw())
}

// GetPrintFmt returns the value in PrintFmt.
func (p GuiDataTypeInfo) GetPrintFmt() ffi.CString {
	ptr := (*C.ImGuiDataTypeInfo)(unsafe.Pointer(p))
	return ffi.CStringFromPtr((unsafe.Pointer)(ptr.PrintFmt))
}

// SetPrintFmt sets the value in PrintFmt.
func (p GuiDataTypeInfo) SetPrintFmt(value ffi.CString) {
	ptr := (*C.ImGuiDataTypeInfo)(unsafe.Pointer(p))
	ptr.PrintFmt = (*C.char)(value.Raw())
}

// GetScanFmt returns the value in ScanFmt.
func (p GuiDataTypeInfo) GetScanFmt() ffi.CString {
	ptr := (*C.ImGuiDataTypeInfo)(unsafe.Pointer(p))
	return ffi.CStringFromPtr((unsafe.Pointer)(ptr.ScanFmt))
}

// SetScanFmt sets the value in ScanFmt.
func (p GuiDataTypeInfo) SetScanFmt(value ffi.CString) {
	ptr := (*C.ImGuiDataTypeInfo)(unsafe.Pointer(p))
	ptr.ScanFmt = (*C.char)(value.Raw())
}

// GuiDataTypeInfo.Size is unsupported: category unsupported.

// GuiDataTypeStorage wraps struct ImGuiDataTypeStorage.
type GuiDataTypeStorage uintptr

// GuiDataTypeStorageNil is a null pointer.
var GuiDataTypeStorageNil GuiDataTypeStorage

// GuiDataTypeStorageSizeOf is the byte size of ImGuiDataTypeStorage.
const GuiDataTypeStorageSizeOf = int(C.sizeof_ImGuiDataTypeStorage)

// GuiDataTypeStorageAlloc allocates a continuous block of GuiDataTypeStorage.
func GuiDataTypeStorageAlloc(alloc ffi.Allocator, count int) GuiDataTypeStorage {
	ptr := alloc.Allocate(GuiDataTypeStorageSizeOf * count)
	return GuiDataTypeStorage(ptr)
}

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p GuiDataTypeStorage) Offset(offset int) GuiDataTypeStorage {
	return p + GuiDataTypeStorage(offset*GuiDataTypeStorageSizeOf)
}

// GuiDataTypeStorage.Data[8] is unsupported: category unsupported.

// GuiDeactivatedItemData wraps struct ImGuiDeactivatedItemData.
type GuiDeactivatedItemData uintptr

// GuiDeactivatedItemDataNil is a null pointer.
var GuiDeactivatedItemDataNil GuiDeactivatedItemData

// GuiDeactivatedItemDataSizeOf is the byte size of ImGuiDeactivatedItemData.
const GuiDeactivatedItemDataSizeOf = int(C.sizeof_ImGuiDeactivatedItemData)

// GuiDeactivatedItemDataAlloc allocates a continuous block of GuiDeactivatedItemData.
func GuiDeactivatedItemDataAlloc(alloc ffi.Allocator, count int) GuiDeactivatedItemData {
	ptr := alloc.Allocate(GuiDeactivatedItemDataSizeOf * count)
	return GuiDeactivatedItemData(ptr)
}

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p GuiDeactivatedItemData) Offset(offset int) GuiDeactivatedItemData {
	return p + GuiDeactivatedItemData(offset*GuiDeactivatedItemDataSizeOf)
}

// GetElapseFrame returns the value in ElapseFrame.
func (p GuiDeactivatedItemData) GetElapseFrame() int32 {
	ptr := (*C.ImGuiDeactivatedItemData)(unsafe.Pointer(p))
	return int32(ptr.ElapseFrame)
}

// SetElapseFrame sets the value in ElapseFrame.
func (p GuiDeactivatedItemData) SetElapseFrame(value int32) {
	ptr := (*C.ImGuiDeactivatedItemData)(unsafe.Pointer(p))
	ptr.ElapseFrame = (C.int32_t)(value)
}

// GetHasBeenEditedBefore returns the value in HasBeenEditedBefore.
func (p GuiDeactivatedItemData) GetHasBeenEditedBefore() bool {
	ptr := (*C.ImGuiDeactivatedItemData)(unsafe.Pointer(p))
	return bool(ptr.HasBeenEditedBefore)
}

// SetHasBeenEditedBefore sets the value in HasBeenEditedBefore.
func (p GuiDeactivatedItemData) SetHasBeenEditedBefore(value bool) {
	ptr := (*C.ImGuiDeactivatedItemData)(unsafe.Pointer(p))
	ptr.HasBeenEditedBefore = (C.bool)(value)
}

// GuiDeactivatedItemData.ID is unsupported: category unsupported.

// GetIsAlive returns the value in IsAlive.
func (p GuiDeactivatedItemData) GetIsAlive() bool {
	ptr := (*C.ImGuiDeactivatedItemData)(unsafe.Pointer(p))
	return bool(ptr.IsAlive)
}

// SetIsAlive sets the value in IsAlive.
func (p GuiDeactivatedItemData) SetIsAlive(value bool) {
	ptr := (*C.ImGuiDeactivatedItemData)(unsafe.Pointer(p))
	ptr.IsAlive = (C.bool)(value)
}

// GuiDebugAllocEntry wraps struct ImGuiDebugAllocEntry.
type GuiDebugAllocEntry uintptr

// GuiDebugAllocEntryNil is a null pointer.
var GuiDebugAllocEntryNil GuiDebugAllocEntry

// GuiDebugAllocEntrySizeOf is the byte size of ImGuiDebugAllocEntry.
const GuiDebugAllocEntrySizeOf = int(C.sizeof_ImGuiDebugAllocEntry)

// GuiDebugAllocEntryAlloc allocates a continuous block of GuiDebugAllocEntry.
func GuiDebugAllocEntryAlloc(alloc ffi.Allocator, count int) GuiDebugAllocEntry {
	ptr := alloc.Allocate(GuiDebugAllocEntrySizeOf * count)
	return GuiDebugAllocEntry(ptr)
}

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p GuiDebugAllocEntry) Offset(offset int) GuiDebugAllocEntry {
	return p + GuiDebugAllocEntry(offset*GuiDebugAllocEntrySizeOf)
}

// GuiDebugAllocEntry.AllocCount is unsupported: category unsupported.

// GetFrameCount returns the value in FrameCount.
func (p GuiDebugAllocEntry) GetFrameCount() int32 {
	ptr := (*C.ImGuiDebugAllocEntry)(unsafe.Pointer(p))
	return int32(ptr.FrameCount)
}

// SetFrameCount sets the value in FrameCount.
func (p GuiDebugAllocEntry) SetFrameCount(value int32) {
	ptr := (*C.ImGuiDebugAllocEntry)(unsafe.Pointer(p))
	ptr.FrameCount = (C.int32_t)(value)
}

// GuiDebugAllocEntry.FreeCount is unsupported: category unsupported.

// GuiDebugAllocInfo wraps struct ImGuiDebugAllocInfo.
type GuiDebugAllocInfo uintptr

// GuiDebugAllocInfoNil is a null pointer.
var GuiDebugAllocInfoNil GuiDebugAllocInfo

// GuiDebugAllocInfoSizeOf is the byte size of ImGuiDebugAllocInfo.
const GuiDebugAllocInfoSizeOf = int(C.sizeof_ImGuiDebugAllocInfo)

// ImGuiDebugAllocInfo allocator is suppressed.

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p GuiDebugAllocInfo) Offset(offset int) GuiDebugAllocInfo {
	return p + GuiDebugAllocInfo(offset*GuiDebugAllocInfoSizeOf)
}

// GuiDebugAllocInfo.LastEntriesBuf[6] is unsupported: category unsupported.

// GuiDebugAllocInfo.LastEntriesIdx is unsupported: category unsupported.

// GetTotalAllocCount returns the value in TotalAllocCount.
func (p GuiDebugAllocInfo) GetTotalAllocCount() int32 {
	ptr := (*C.ImGuiDebugAllocInfo)(unsafe.Pointer(p))
	return int32(ptr.TotalAllocCount)
}

// SetTotalAllocCount sets the value in TotalAllocCount.
func (p GuiDebugAllocInfo) SetTotalAllocCount(value int32) {
	ptr := (*C.ImGuiDebugAllocInfo)(unsafe.Pointer(p))
	ptr.TotalAllocCount = (C.int32_t)(value)
}

// GetTotalFreeCount returns the value in TotalFreeCount.
func (p GuiDebugAllocInfo) GetTotalFreeCount() int32 {
	ptr := (*C.ImGuiDebugAllocInfo)(unsafe.Pointer(p))
	return int32(ptr.TotalFreeCount)
}

// SetTotalFreeCount sets the value in TotalFreeCount.
func (p GuiDebugAllocInfo) SetTotalFreeCount(value int32) {
	ptr := (*C.ImGuiDebugAllocInfo)(unsafe.Pointer(p))
	ptr.TotalFreeCount = (C.int32_t)(value)
}

// GuiDockContext wraps struct ImGuiDockContext.
type GuiDockContext uintptr

// GuiDockContextNil is a null pointer.
var GuiDockContextNil GuiDockContext

// GuiDockContextSizeOf is the byte size of ImGuiDockContext.
const GuiDockContextSizeOf = int(C.sizeof_ImGuiDockContext)

// ImGuiDockContext allocator is suppressed.

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p GuiDockContext) Offset(offset int) GuiDockContext {
	return p + GuiDockContext(offset*GuiDockContextSizeOf)
}

// RefNodes returns pointer to the Nodes field.
func (p GuiDockContext) RefNodes() GuiStorage {
	return GuiStorage(p + GuiDockContext(C.offsetof_ImGuiDockContext_Nodes))
}

// RefNodesSettings returns pointer to the NodesSettings field.
func (p GuiDockContext) RefNodesSettings() Vector_ImGuiDockNodeSettings {
	return Vector_ImGuiDockNodeSettings(p + GuiDockContext(C.offsetof_ImGuiDockContext_NodesSettings))
}

// RefRequests returns pointer to the Requests field.
func (p GuiDockContext) RefRequests() Vector_ImGuiDockRequest {
	return Vector_ImGuiDockRequest(p + GuiDockContext(C.offsetof_ImGuiDockContext_Requests))
}

// GetWantFullRebuild returns the value in WantFullRebuild.
func (p GuiDockContext) GetWantFullRebuild() bool {
	ptr := (*C.ImGuiDockContext)(unsafe.Pointer(p))
	return bool(ptr.WantFullRebuild)
}

// SetWantFullRebuild sets the value in WantFullRebuild.
func (p GuiDockContext) SetWantFullRebuild(value bool) {
	ptr := (*C.ImGuiDockContext)(unsafe.Pointer(p))
	ptr.WantFullRebuild = (C.bool)(value)
}

// GuiDockNode wraps struct ImGuiDockNode.
type GuiDockNode uintptr

// GuiDockNodeNil is a null pointer.
var GuiDockNodeNil GuiDockNode

// GuiDockNodeSizeOf is the byte size of ImGuiDockNode.
const GuiDockNodeSizeOf = int(C.sizeof_ImGuiDockNode)

// ImGuiDockNode allocator is suppressed.

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p GuiDockNode) Offset(offset int) GuiDockNode {
	return p + GuiDockNode(offset*GuiDockNodeSizeOf)
}

// GuiDockNode.AuthorityForPos is unsupported: category unsupported.

// GuiDockNode.AuthorityForSize is unsupported: category unsupported.

// GuiDockNode.AuthorityForViewport is unsupported: category unsupported.

// GetCentralNode returns the value in CentralNode.
func (p GuiDockNode) GetCentralNode() GuiDockNode {
	ptr := (*C.ImGuiDockNode)(unsafe.Pointer(p))
	return GuiDockNode(unsafe.Pointer(ptr.CentralNode))
}

// SetCentralNode sets the value in CentralNode.
func (p GuiDockNode) SetCentralNode(value GuiDockNode) {
	ptr := (*C.ImGuiDockNode)(unsafe.Pointer(p))
	ptr.CentralNode = (*C.ImGuiDockNode)(unsafe.Pointer(value))
}

// GuiDockNode.ChildNodes[2] is unsupported: category unsupported.

// GetCountNodeWithWindows returns the value in CountNodeWithWindows.
func (p GuiDockNode) GetCountNodeWithWindows() int32 {
	ptr := (*C.ImGuiDockNode)(unsafe.Pointer(p))
	return int32(ptr.CountNodeWithWindows)
}

// SetCountNodeWithWindows sets the value in CountNodeWithWindows.
func (p GuiDockNode) SetCountNodeWithWindows(value int32) {
	ptr := (*C.ImGuiDockNode)(unsafe.Pointer(p))
	ptr.CountNodeWithWindows = (C.int32_t)(value)
}

// GuiDockNode.HasCentralNodeChild is unsupported: category unsupported.

// GuiDockNode.HasCloseButton is unsupported: category unsupported.

// GuiDockNode.HasWindowMenuButton is unsupported: category unsupported.

// GetHostWindow returns the value in HostWindow.
func (p GuiDockNode) GetHostWindow() GuiWindow {
	ptr := (*C.ImGuiDockNode)(unsafe.Pointer(p))
	return GuiWindow(unsafe.Pointer(ptr.HostWindow))
}

// SetHostWindow sets the value in HostWindow.
func (p GuiDockNode) SetHostWindow(value GuiWindow) {
	ptr := (*C.ImGuiDockNode)(unsafe.Pointer(p))
	ptr.HostWindow = (*C.ImGuiWindow)(unsafe.Pointer(value))
}

// GuiDockNode.ID is unsupported: category unsupported.

// GuiDockNode.IsBgDrawnThisFrame is unsupported: category unsupported.

// GuiDockNode.IsFocused is unsupported: category unsupported.

// GuiDockNode.IsVisible is unsupported: category unsupported.

// GetLastBgColor returns the value in LastBgColor.
func (p GuiDockNode) GetLastBgColor() uint32 {
	ptr := (*C.ImGuiDockNode)(unsafe.Pointer(p))
	return uint32(ptr.LastBgColor)
}

// SetLastBgColor sets the value in LastBgColor.
func (p GuiDockNode) SetLastBgColor(value uint32) {
	ptr := (*C.ImGuiDockNode)(unsafe.Pointer(p))
	ptr.LastBgColor = (C.uint32_t)(value)
}

// GuiDockNode.LastFocusedNodeId is unsupported: category unsupported.

// GetLastFrameActive returns the value in LastFrameActive.
func (p GuiDockNode) GetLastFrameActive() int32 {
	ptr := (*C.ImGuiDockNode)(unsafe.Pointer(p))
	return int32(ptr.LastFrameActive)
}

// SetLastFrameActive sets the value in LastFrameActive.
func (p GuiDockNode) SetLastFrameActive(value int32) {
	ptr := (*C.ImGuiDockNode)(unsafe.Pointer(p))
	ptr.LastFrameActive = (C.int32_t)(value)
}

// GetLastFrameAlive returns the value in LastFrameAlive.
func (p GuiDockNode) GetLastFrameAlive() int32 {
	ptr := (*C.ImGuiDockNode)(unsafe.Pointer(p))
	return int32(ptr.LastFrameAlive)
}

// SetLastFrameAlive sets the value in LastFrameAlive.
func (p GuiDockNode) SetLastFrameAlive(value int32) {
	ptr := (*C.ImGuiDockNode)(unsafe.Pointer(p))
	ptr.LastFrameAlive = (C.int32_t)(value)
}

// GetLastFrameFocused returns the value in LastFrameFocused.
func (p GuiDockNode) GetLastFrameFocused() int32 {
	ptr := (*C.ImGuiDockNode)(unsafe.Pointer(p))
	return int32(ptr.LastFrameFocused)
}

// SetLastFrameFocused sets the value in LastFrameFocused.
func (p GuiDockNode) SetLastFrameFocused(value int32) {
	ptr := (*C.ImGuiDockNode)(unsafe.Pointer(p))
	ptr.LastFrameFocused = (C.int32_t)(value)
}

// GetLocalFlags returns the value in LocalFlags.
func (p GuiDockNode) GetLocalFlags() GuiDockNodeFlags {
	ptr := (*C.ImGuiDockNode)(unsafe.Pointer(p))
	return GuiDockNodeFlags(ptr.LocalFlags)
}

// SetLocalFlags sets the value in LocalFlags.
func (p GuiDockNode) SetLocalFlags(value GuiDockNodeFlags) {
	ptr := (*C.ImGuiDockNode)(unsafe.Pointer(p))
	ptr.LocalFlags = (C.ImGuiDockNodeFlags)(value)
}

// GetLocalFlagsInWindows returns the value in LocalFlagsInWindows.
func (p GuiDockNode) GetLocalFlagsInWindows() GuiDockNodeFlags {
	ptr := (*C.ImGuiDockNode)(unsafe.Pointer(p))
	return GuiDockNodeFlags(ptr.LocalFlagsInWindows)
}

// SetLocalFlagsInWindows sets the value in LocalFlagsInWindows.
func (p GuiDockNode) SetLocalFlagsInWindows(value GuiDockNodeFlags) {
	ptr := (*C.ImGuiDockNode)(unsafe.Pointer(p))
	ptr.LocalFlagsInWindows = (C.ImGuiDockNodeFlags)(value)
}

// GetMergedFlags returns the value in MergedFlags.
func (p GuiDockNode) GetMergedFlags() GuiDockNodeFlags {
	ptr := (*C.ImGuiDockNode)(unsafe.Pointer(p))
	return GuiDockNodeFlags(ptr.MergedFlags)
}

// SetMergedFlags sets the value in MergedFlags.
func (p GuiDockNode) SetMergedFlags(value GuiDockNodeFlags) {
	ptr := (*C.ImGuiDockNode)(unsafe.Pointer(p))
	ptr.MergedFlags = (C.ImGuiDockNodeFlags)(value)
}

// GetOnlyNodeWithWindows returns the value in OnlyNodeWithWindows.
func (p GuiDockNode) GetOnlyNodeWithWindows() GuiDockNode {
	ptr := (*C.ImGuiDockNode)(unsafe.Pointer(p))
	return GuiDockNode(unsafe.Pointer(ptr.OnlyNodeWithWindows))
}

// SetOnlyNodeWithWindows sets the value in OnlyNodeWithWindows.
func (p GuiDockNode) SetOnlyNodeWithWindows(value GuiDockNode) {
	ptr := (*C.ImGuiDockNode)(unsafe.Pointer(p))
	ptr.OnlyNodeWithWindows = (*C.ImGuiDockNode)(unsafe.Pointer(value))
}

// GetParentNode returns the value in ParentNode.
func (p GuiDockNode) GetParentNode() GuiDockNode {
	ptr := (*C.ImGuiDockNode)(unsafe.Pointer(p))
	return GuiDockNode(unsafe.Pointer(ptr.ParentNode))
}

// SetParentNode sets the value in ParentNode.
func (p GuiDockNode) SetParentNode(value GuiDockNode) {
	ptr := (*C.ImGuiDockNode)(unsafe.Pointer(p))
	ptr.ParentNode = (*C.ImGuiDockNode)(unsafe.Pointer(value))
}

// RefPos returns pointer to the Pos field.
func (p GuiDockNode) RefPos() Vec2 {
	return Vec2(p + GuiDockNode(C.offsetof_ImGuiDockNode_Pos))
}

// GuiDockNode.RefViewportId is unsupported: category unsupported.

// GuiDockNode.SelectedTabId is unsupported: category unsupported.

// GetSharedFlags returns the value in SharedFlags.
func (p GuiDockNode) GetSharedFlags() GuiDockNodeFlags {
	ptr := (*C.ImGuiDockNode)(unsafe.Pointer(p))
	return GuiDockNodeFlags(ptr.SharedFlags)
}

// SetSharedFlags sets the value in SharedFlags.
func (p GuiDockNode) SetSharedFlags(value GuiDockNodeFlags) {
	ptr := (*C.ImGuiDockNode)(unsafe.Pointer(p))
	ptr.SharedFlags = (C.ImGuiDockNodeFlags)(value)
}

// RefSize returns pointer to the Size field.
func (p GuiDockNode) RefSize() Vec2 {
	return Vec2(p + GuiDockNode(C.offsetof_ImGuiDockNode_Size))
}

// RefSizeRef returns pointer to the SizeRef field.
func (p GuiDockNode) RefSizeRef() Vec2 {
	return Vec2(p + GuiDockNode(C.offsetof_ImGuiDockNode_SizeRef))
}

// GetSplitAxis returns the value in SplitAxis.
func (p GuiDockNode) GetSplitAxis() GuiAxis {
	ptr := (*C.ImGuiDockNode)(unsafe.Pointer(p))
	return GuiAxis(ptr.SplitAxis)
}

// SetSplitAxis sets the value in SplitAxis.
func (p GuiDockNode) SetSplitAxis(value GuiAxis) {
	ptr := (*C.ImGuiDockNode)(unsafe.Pointer(p))
	ptr.SplitAxis = (C.ImGuiAxis)(value)
}

// GetState returns the value in State.
func (p GuiDockNode) GetState() GuiDockNodeState {
	ptr := (*C.ImGuiDockNode)(unsafe.Pointer(p))
	return GuiDockNodeState(ptr.State)
}

// SetState sets the value in State.
func (p GuiDockNode) SetState(value GuiDockNodeState) {
	ptr := (*C.ImGuiDockNode)(unsafe.Pointer(p))
	ptr.State = (C.ImGuiDockNodeState)(value)
}

// GetTabBar returns the value in TabBar.
func (p GuiDockNode) GetTabBar() GuiTabBar {
	ptr := (*C.ImGuiDockNode)(unsafe.Pointer(p))
	return GuiTabBar(unsafe.Pointer(ptr.TabBar))
}

// SetTabBar sets the value in TabBar.
func (p GuiDockNode) SetTabBar(value GuiTabBar) {
	ptr := (*C.ImGuiDockNode)(unsafe.Pointer(p))
	ptr.TabBar = (*C.ImGuiTabBar)(unsafe.Pointer(value))
}

// GetVisibleWindow returns the value in VisibleWindow.
func (p GuiDockNode) GetVisibleWindow() GuiWindow {
	ptr := (*C.ImGuiDockNode)(unsafe.Pointer(p))
	return GuiWindow(unsafe.Pointer(ptr.VisibleWindow))
}

// SetVisibleWindow sets the value in VisibleWindow.
func (p GuiDockNode) SetVisibleWindow(value GuiWindow) {
	ptr := (*C.ImGuiDockNode)(unsafe.Pointer(p))
	ptr.VisibleWindow = (*C.ImGuiWindow)(unsafe.Pointer(value))
}

// GuiDockNode.WantCloseAll is unsupported: category unsupported.

// GuiDockNode.WantCloseTabId is unsupported: category unsupported.

// GuiDockNode.WantHiddenTabBarToggle is unsupported: category unsupported.

// GuiDockNode.WantHiddenTabBarUpdate is unsupported: category unsupported.

// GuiDockNode.WantLockSizeOnce is unsupported: category unsupported.

// GuiDockNode.WantMouseMove is unsupported: category unsupported.

// RefWindowClass returns pointer to the WindowClass field.
func (p GuiDockNode) RefWindowClass() GuiWindowClass {
	return GuiWindowClass(p + GuiDockNode(C.offsetof_ImGuiDockNode_WindowClass))
}

// RefWindows returns pointer to the Windows field.
func (p GuiDockNode) RefWindows() Vector_ImGuiWindowPtr {
	return Vector_ImGuiWindowPtr(p + GuiDockNode(C.offsetof_ImGuiDockNode_Windows))
}

// GuiErrorRecoveryState wraps struct ImGuiErrorRecoveryState.
type GuiErrorRecoveryState uintptr

// GuiErrorRecoveryStateNil is a null pointer.
var GuiErrorRecoveryStateNil GuiErrorRecoveryState

// GuiErrorRecoveryStateSizeOf is the byte size of ImGuiErrorRecoveryState.
const GuiErrorRecoveryStateSizeOf = int(C.sizeof_ImGuiErrorRecoveryState)

// ImGuiErrorRecoveryState allocator is suppressed.

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p GuiErrorRecoveryState) Offset(offset int) GuiErrorRecoveryState {
	return p + GuiErrorRecoveryState(offset*GuiErrorRecoveryStateSizeOf)
}

// GuiErrorRecoveryState.SizeOfBeginPopupStack is unsupported: category unsupported.

// GuiErrorRecoveryState.SizeOfColorStack is unsupported: category unsupported.

// GuiErrorRecoveryState.SizeOfDisabledStack is unsupported: category unsupported.

// GuiErrorRecoveryState.SizeOfFocusScopeStack is unsupported: category unsupported.

// GuiErrorRecoveryState.SizeOfFontStack is unsupported: category unsupported.

// GuiErrorRecoveryState.SizeOfGroupStack is unsupported: category unsupported.

// GuiErrorRecoveryState.SizeOfIDStack is unsupported: category unsupported.

// GuiErrorRecoveryState.SizeOfItemFlagsStack is unsupported: category unsupported.

// GuiErrorRecoveryState.SizeOfStyleVarStack is unsupported: category unsupported.

// GuiErrorRecoveryState.SizeOfTreeStack is unsupported: category unsupported.

// GuiErrorRecoveryState.SizeOfWindowStack is unsupported: category unsupported.

// GuiFocusScopeData wraps struct ImGuiFocusScopeData.
type GuiFocusScopeData uintptr

// GuiFocusScopeDataNil is a null pointer.
var GuiFocusScopeDataNil GuiFocusScopeData

// GuiFocusScopeDataSizeOf is the byte size of ImGuiFocusScopeData.
const GuiFocusScopeDataSizeOf = int(C.sizeof_ImGuiFocusScopeData)

// GuiFocusScopeDataAlloc allocates a continuous block of GuiFocusScopeData.
func GuiFocusScopeDataAlloc(alloc ffi.Allocator, count int) GuiFocusScopeData {
	ptr := alloc.Allocate(GuiFocusScopeDataSizeOf * count)
	return GuiFocusScopeData(ptr)
}

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p GuiFocusScopeData) Offset(offset int) GuiFocusScopeData {
	return p + GuiFocusScopeData(offset*GuiFocusScopeDataSizeOf)
}

// GuiFocusScopeData.ID is unsupported: category unsupported.

// GuiFocusScopeData.WindowID is unsupported: category unsupported.

// GuiGroupData wraps struct ImGuiGroupData.
type GuiGroupData uintptr

// GuiGroupDataNil is a null pointer.
var GuiGroupDataNil GuiGroupData

// GuiGroupDataSizeOf is the byte size of ImGuiGroupData.
const GuiGroupDataSizeOf = int(C.sizeof_ImGuiGroupData)

// GuiGroupDataAlloc allocates a continuous block of GuiGroupData.
func GuiGroupDataAlloc(alloc ffi.Allocator, count int) GuiGroupData {
	ptr := alloc.Allocate(GuiGroupDataSizeOf * count)
	return GuiGroupData(ptr)
}

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p GuiGroupData) Offset(offset int) GuiGroupData {
	return p + GuiGroupData(offset*GuiGroupDataSizeOf)
}

// GuiGroupData.BackupActiveIdIsAlive is unsupported: category unsupported.

// RefBackupCurrLineSize returns pointer to the BackupCurrLineSize field.
func (p GuiGroupData) RefBackupCurrLineSize() Vec2 {
	return Vec2(p + GuiGroupData(C.offsetof_ImGuiGroupData_BackupCurrLineSize))
}

// GetBackupCurrLineTextBaseOffset returns the value in BackupCurrLineTextBaseOffset.
func (p GuiGroupData) GetBackupCurrLineTextBaseOffset() float32 {
	ptr := (*C.ImGuiGroupData)(unsafe.Pointer(p))
	return float32(ptr.BackupCurrLineTextBaseOffset)
}

// SetBackupCurrLineTextBaseOffset sets the value in BackupCurrLineTextBaseOffset.
func (p GuiGroupData) SetBackupCurrLineTextBaseOffset(value float32) {
	ptr := (*C.ImGuiGroupData)(unsafe.Pointer(p))
	ptr.BackupCurrLineTextBaseOffset = (C.float)(value)
}

// RefBackupCursorMaxPos returns pointer to the BackupCursorMaxPos field.
func (p GuiGroupData) RefBackupCursorMaxPos() Vec2 {
	return Vec2(p + GuiGroupData(C.offsetof_ImGuiGroupData_BackupCursorMaxPos))
}

// RefBackupCursorPos returns pointer to the BackupCursorPos field.
func (p GuiGroupData) RefBackupCursorPos() Vec2 {
	return Vec2(p + GuiGroupData(C.offsetof_ImGuiGroupData_BackupCursorPos))
}

// RefBackupCursorPosPrevLine returns pointer to the BackupCursorPosPrevLine field.
func (p GuiGroupData) RefBackupCursorPosPrevLine() Vec2 {
	return Vec2(p + GuiGroupData(C.offsetof_ImGuiGroupData_BackupCursorPosPrevLine))
}

// GetBackupDeactivatedIdIsAlive returns the value in BackupDeactivatedIdIsAlive.
func (p GuiGroupData) GetBackupDeactivatedIdIsAlive() bool {
	ptr := (*C.ImGuiGroupData)(unsafe.Pointer(p))
	return bool(ptr.BackupDeactivatedIdIsAlive)
}

// SetBackupDeactivatedIdIsAlive sets the value in BackupDeactivatedIdIsAlive.
func (p GuiGroupData) SetBackupDeactivatedIdIsAlive(value bool) {
	ptr := (*C.ImGuiGroupData)(unsafe.Pointer(p))
	ptr.BackupDeactivatedIdIsAlive = (C.bool)(value)
}

// RefBackupGroupOffset returns pointer to the BackupGroupOffset field.
func (p GuiGroupData) RefBackupGroupOffset() Vec1 {
	return Vec1(p + GuiGroupData(C.offsetof_ImGuiGroupData_BackupGroupOffset))
}

// GetBackupHoveredIdIsAlive returns the value in BackupHoveredIdIsAlive.
func (p GuiGroupData) GetBackupHoveredIdIsAlive() bool {
	ptr := (*C.ImGuiGroupData)(unsafe.Pointer(p))
	return bool(ptr.BackupHoveredIdIsAlive)
}

// SetBackupHoveredIdIsAlive sets the value in BackupHoveredIdIsAlive.
func (p GuiGroupData) SetBackupHoveredIdIsAlive(value bool) {
	ptr := (*C.ImGuiGroupData)(unsafe.Pointer(p))
	ptr.BackupHoveredIdIsAlive = (C.bool)(value)
}

// RefBackupIndent returns pointer to the BackupIndent field.
func (p GuiGroupData) RefBackupIndent() Vec1 {
	return Vec1(p + GuiGroupData(C.offsetof_ImGuiGroupData_BackupIndent))
}

// GetBackupIsSameLine returns the value in BackupIsSameLine.
func (p GuiGroupData) GetBackupIsSameLine() bool {
	ptr := (*C.ImGuiGroupData)(unsafe.Pointer(p))
	return bool(ptr.BackupIsSameLine)
}

// SetBackupIsSameLine sets the value in BackupIsSameLine.
func (p GuiGroupData) SetBackupIsSameLine(value bool) {
	ptr := (*C.ImGuiGroupData)(unsafe.Pointer(p))
	ptr.BackupIsSameLine = (C.bool)(value)
}

// GetEmitItem returns the value in EmitItem.
func (p GuiGroupData) GetEmitItem() bool {
	ptr := (*C.ImGuiGroupData)(unsafe.Pointer(p))
	return bool(ptr.EmitItem)
}

// SetEmitItem sets the value in EmitItem.
func (p GuiGroupData) SetEmitItem(value bool) {
	ptr := (*C.ImGuiGroupData)(unsafe.Pointer(p))
	ptr.EmitItem = (C.bool)(value)
}

// GuiGroupData.WindowID is unsupported: category unsupported.

// GuiIDStackTool wraps struct ImGuiIDStackTool.
type GuiIDStackTool uintptr

// GuiIDStackToolNil is a null pointer.
var GuiIDStackToolNil GuiIDStackTool

// GuiIDStackToolSizeOf is the byte size of ImGuiIDStackTool.
const GuiIDStackToolSizeOf = int(C.sizeof_ImGuiIDStackTool)

// ImGuiIDStackTool allocator is suppressed.

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p GuiIDStackTool) Offset(offset int) GuiIDStackTool {
	return p + GuiIDStackTool(offset*GuiIDStackToolSizeOf)
}

// GetCopyToClipboardLastTime returns the value in CopyToClipboardLastTime.
func (p GuiIDStackTool) GetCopyToClipboardLastTime() float32 {
	ptr := (*C.ImGuiIDStackTool)(unsafe.Pointer(p))
	return float32(ptr.CopyToClipboardLastTime)
}

// SetCopyToClipboardLastTime sets the value in CopyToClipboardLastTime.
func (p GuiIDStackTool) SetCopyToClipboardLastTime(value float32) {
	ptr := (*C.ImGuiIDStackTool)(unsafe.Pointer(p))
	ptr.CopyToClipboardLastTime = (C.float)(value)
}

// GetLastActiveFrame returns the value in LastActiveFrame.
func (p GuiIDStackTool) GetLastActiveFrame() int32 {
	ptr := (*C.ImGuiIDStackTool)(unsafe.Pointer(p))
	return int32(ptr.LastActiveFrame)
}

// SetLastActiveFrame sets the value in LastActiveFrame.
func (p GuiIDStackTool) SetLastActiveFrame(value int32) {
	ptr := (*C.ImGuiIDStackTool)(unsafe.Pointer(p))
	ptr.LastActiveFrame = (C.int32_t)(value)
}

// GetOptCopyToClipboardOnCtrlC returns the value in OptCopyToClipboardOnCtrlC.
func (p GuiIDStackTool) GetOptCopyToClipboardOnCtrlC() bool {
	ptr := (*C.ImGuiIDStackTool)(unsafe.Pointer(p))
	return bool(ptr.OptCopyToClipboardOnCtrlC)
}

// SetOptCopyToClipboardOnCtrlC sets the value in OptCopyToClipboardOnCtrlC.
func (p GuiIDStackTool) SetOptCopyToClipboardOnCtrlC(value bool) {
	ptr := (*C.ImGuiIDStackTool)(unsafe.Pointer(p))
	ptr.OptCopyToClipboardOnCtrlC = (C.bool)(value)
}

// GetOptHexEncodeNonAsciiChars returns the value in OptHexEncodeNonAsciiChars.
func (p GuiIDStackTool) GetOptHexEncodeNonAsciiChars() bool {
	ptr := (*C.ImGuiIDStackTool)(unsafe.Pointer(p))
	return bool(ptr.OptHexEncodeNonAsciiChars)
}

// SetOptHexEncodeNonAsciiChars sets the value in OptHexEncodeNonAsciiChars.
func (p GuiIDStackTool) SetOptHexEncodeNonAsciiChars(value bool) {
	ptr := (*C.ImGuiIDStackTool)(unsafe.Pointer(p))
	ptr.OptHexEncodeNonAsciiChars = (C.bool)(value)
}

// GetQueryHookActive returns the value in QueryHookActive.
func (p GuiIDStackTool) GetQueryHookActive() bool {
	ptr := (*C.ImGuiIDStackTool)(unsafe.Pointer(p))
	return bool(ptr.QueryHookActive)
}

// SetQueryHookActive sets the value in QueryHookActive.
func (p GuiIDStackTool) SetQueryHookActive(value bool) {
	ptr := (*C.ImGuiIDStackTool)(unsafe.Pointer(p))
	ptr.QueryHookActive = (C.bool)(value)
}

// GuiIDStackTool.QueryMainId is unsupported: category unsupported.

// RefResultPathsBuf returns pointer to the ResultPathsBuf field.
func (p GuiIDStackTool) RefResultPathsBuf() GuiTextBuffer {
	return GuiTextBuffer(p + GuiIDStackTool(C.offsetof_ImGuiIDStackTool_ResultPathsBuf))
}

// RefResultTempBuf returns pointer to the ResultTempBuf field.
func (p GuiIDStackTool) RefResultTempBuf() GuiTextBuffer {
	return GuiTextBuffer(p + GuiIDStackTool(C.offsetof_ImGuiIDStackTool_ResultTempBuf))
}

// RefResults returns pointer to the Results field.
func (p GuiIDStackTool) RefResults() Vector_ImGuiStackLevelInfo {
	return Vector_ImGuiStackLevelInfo(p + GuiIDStackTool(C.offsetof_ImGuiIDStackTool_Results))
}

// GetStackLevel returns the value in StackLevel.
func (p GuiIDStackTool) GetStackLevel() int32 {
	ptr := (*C.ImGuiIDStackTool)(unsafe.Pointer(p))
	return int32(ptr.StackLevel)
}

// SetStackLevel sets the value in StackLevel.
func (p GuiIDStackTool) SetStackLevel(value int32) {
	ptr := (*C.ImGuiIDStackTool)(unsafe.Pointer(p))
	ptr.StackLevel = (C.int32_t)(value)
}

// GuiIO wraps struct ImGuiIO.
type GuiIO uintptr

// GuiIONil is a null pointer.
var GuiIONil GuiIO

// GuiIOSizeOf is the byte size of ImGuiIO.
const GuiIOSizeOf = int(C.sizeof_ImGuiIO)

// ImGuiIO allocator is suppressed.

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p GuiIO) Offset(offset int) GuiIO {
	return p + GuiIO(offset*GuiIOSizeOf)
}

// GetAppAcceptingEvents returns the value in AppAcceptingEvents.
func (p GuiIO) GetAppAcceptingEvents() bool {
	ptr := (*C.ImGuiIO)(unsafe.Pointer(p))
	return bool(ptr.AppAcceptingEvents)
}

// GuiIO.AppAcceptingEvents getter is suppressed.

// GetAppFocusLost returns the value in AppFocusLost.
func (p GuiIO) GetAppFocusLost() bool {
	ptr := (*C.ImGuiIO)(unsafe.Pointer(p))
	return bool(ptr.AppFocusLost)
}

// SetAppFocusLost sets the value in AppFocusLost.
func (p GuiIO) SetAppFocusLost(value bool) {
	ptr := (*C.ImGuiIO)(unsafe.Pointer(p))
	ptr.AppFocusLost = (C.bool)(value)
}

// GetBackendFlags returns the value in BackendFlags.
func (p GuiIO) GetBackendFlags() GuiBackendFlags {
	ptr := (*C.ImGuiIO)(unsafe.Pointer(p))
	return GuiBackendFlags(ptr.BackendFlags)
}

// SetBackendFlags sets the value in BackendFlags.
func (p GuiIO) SetBackendFlags(value GuiBackendFlags) {
	ptr := (*C.ImGuiIO)(unsafe.Pointer(p))
	ptr.BackendFlags = (C.ImGuiBackendFlags)(value)
}

// GetBackendLanguageUserData returns the value in BackendLanguageUserData.
func (p GuiIO) GetBackendLanguageUserData() uintptr {
	ptr := (*C.ImGuiIO)(unsafe.Pointer(p))
	return uintptr(unsafe.Pointer(ptr.BackendLanguageUserData))
}

// SetBackendLanguageUserData sets the value in BackendLanguageUserData.
func (p GuiIO) SetBackendLanguageUserData(value uintptr) {
	ptr := (*C.ImGuiIO)(unsafe.Pointer(p))
	ptr.BackendLanguageUserData = unsafe.Pointer(value)
}

// GetBackendPlatformName returns the value in BackendPlatformName.
func (p GuiIO) GetBackendPlatformName() ffi.CString {
	ptr := (*C.ImGuiIO)(unsafe.Pointer(p))
	return ffi.CStringFromPtr((unsafe.Pointer)(ptr.BackendPlatformName))
}

// SetBackendPlatformName sets the value in BackendPlatformName.
func (p GuiIO) SetBackendPlatformName(value ffi.CString) {
	ptr := (*C.ImGuiIO)(unsafe.Pointer(p))
	ptr.BackendPlatformName = (*C.char)(value.Raw())
}

// GetBackendPlatformUserData returns the value in BackendPlatformUserData.
func (p GuiIO) GetBackendPlatformUserData() uintptr {
	ptr := (*C.ImGuiIO)(unsafe.Pointer(p))
	return uintptr(unsafe.Pointer(ptr.BackendPlatformUserData))
}

// SetBackendPlatformUserData sets the value in BackendPlatformUserData.
func (p GuiIO) SetBackendPlatformUserData(value uintptr) {
	ptr := (*C.ImGuiIO)(unsafe.Pointer(p))
	ptr.BackendPlatformUserData = unsafe.Pointer(value)
}

// GetBackendRendererName returns the value in BackendRendererName.
func (p GuiIO) GetBackendRendererName() ffi.CString {
	ptr := (*C.ImGuiIO)(unsafe.Pointer(p))
	return ffi.CStringFromPtr((unsafe.Pointer)(ptr.BackendRendererName))
}

// SetBackendRendererName sets the value in BackendRendererName.
func (p GuiIO) SetBackendRendererName(value ffi.CString) {
	ptr := (*C.ImGuiIO)(unsafe.Pointer(p))
	ptr.BackendRendererName = (*C.char)(value.Raw())
}

// GetBackendRendererUserData returns the value in BackendRendererUserData.
func (p GuiIO) GetBackendRendererUserData() uintptr {
	ptr := (*C.ImGuiIO)(unsafe.Pointer(p))
	return uintptr(unsafe.Pointer(ptr.BackendRendererUserData))
}

// SetBackendRendererUserData sets the value in BackendRendererUserData.
func (p GuiIO) SetBackendRendererUserData(value uintptr) {
	ptr := (*C.ImGuiIO)(unsafe.Pointer(p))
	ptr.BackendRendererUserData = unsafe.Pointer(value)
}

// GetConfigDebugBeginReturnValueLoop returns the value in ConfigDebugBeginReturnValueLoop.
func (p GuiIO) GetConfigDebugBeginReturnValueLoop() bool {
	ptr := (*C.ImGuiIO)(unsafe.Pointer(p))
	return bool(ptr.ConfigDebugBeginReturnValueLoop)
}

// SetConfigDebugBeginReturnValueLoop sets the value in ConfigDebugBeginReturnValueLoop.
func (p GuiIO) SetConfigDebugBeginReturnValueLoop(value bool) {
	ptr := (*C.ImGuiIO)(unsafe.Pointer(p))
	ptr.ConfigDebugBeginReturnValueLoop = (C.bool)(value)
}

// GetConfigDebugBeginReturnValueOnce returns the value in ConfigDebugBeginReturnValueOnce.
func (p GuiIO) GetConfigDebugBeginReturnValueOnce() bool {
	ptr := (*C.ImGuiIO)(unsafe.Pointer(p))
	return bool(ptr.ConfigDebugBeginReturnValueOnce)
}

// SetConfigDebugBeginReturnValueOnce sets the value in ConfigDebugBeginReturnValueOnce.
func (p GuiIO) SetConfigDebugBeginReturnValueOnce(value bool) {
	ptr := (*C.ImGuiIO)(unsafe.Pointer(p))
	ptr.ConfigDebugBeginReturnValueOnce = (C.bool)(value)
}

// GetConfigDebugHighlightIdConflicts returns the value in ConfigDebugHighlightIdConflicts.
func (p GuiIO) GetConfigDebugHighlightIdConflicts() bool {
	ptr := (*C.ImGuiIO)(unsafe.Pointer(p))
	return bool(ptr.ConfigDebugHighlightIdConflicts)
}

// SetConfigDebugHighlightIdConflicts sets the value in ConfigDebugHighlightIdConflicts.
func (p GuiIO) SetConfigDebugHighlightIdConflicts(value bool) {
	ptr := (*C.ImGuiIO)(unsafe.Pointer(p))
	ptr.ConfigDebugHighlightIdConflicts = (C.bool)(value)
}

// GetConfigDebugHighlightIdConflictsShowItemPicker returns the value in ConfigDebugHighlightIdConflictsShowItemPicker.
func (p GuiIO) GetConfigDebugHighlightIdConflictsShowItemPicker() bool {
	ptr := (*C.ImGuiIO)(unsafe.Pointer(p))
	return bool(ptr.ConfigDebugHighlightIdConflictsShowItemPicker)
}

// SetConfigDebugHighlightIdConflictsShowItemPicker sets the value in ConfigDebugHighlightIdConflictsShowItemPicker.
func (p GuiIO) SetConfigDebugHighlightIdConflictsShowItemPicker(value bool) {
	ptr := (*C.ImGuiIO)(unsafe.Pointer(p))
	ptr.ConfigDebugHighlightIdConflictsShowItemPicker = (C.bool)(value)
}

// GetConfigDebugIgnoreFocusLoss returns the value in ConfigDebugIgnoreFocusLoss.
func (p GuiIO) GetConfigDebugIgnoreFocusLoss() bool {
	ptr := (*C.ImGuiIO)(unsafe.Pointer(p))
	return bool(ptr.ConfigDebugIgnoreFocusLoss)
}

// SetConfigDebugIgnoreFocusLoss sets the value in ConfigDebugIgnoreFocusLoss.
func (p GuiIO) SetConfigDebugIgnoreFocusLoss(value bool) {
	ptr := (*C.ImGuiIO)(unsafe.Pointer(p))
	ptr.ConfigDebugIgnoreFocusLoss = (C.bool)(value)
}

// GetConfigDebugIniSettings returns the value in ConfigDebugIniSettings.
func (p GuiIO) GetConfigDebugIniSettings() bool {
	ptr := (*C.ImGuiIO)(unsafe.Pointer(p))
	return bool(ptr.ConfigDebugIniSettings)
}

// SetConfigDebugIniSettings sets the value in ConfigDebugIniSettings.
func (p GuiIO) SetConfigDebugIniSettings(value bool) {
	ptr := (*C.ImGuiIO)(unsafe.Pointer(p))
	ptr.ConfigDebugIniSettings = (C.bool)(value)
}

// GetConfigDebugIsDebuggerPresent returns the value in ConfigDebugIsDebuggerPresent.
func (p GuiIO) GetConfigDebugIsDebuggerPresent() bool {
	ptr := (*C.ImGuiIO)(unsafe.Pointer(p))
	return bool(ptr.ConfigDebugIsDebuggerPresent)
}

// SetConfigDebugIsDebuggerPresent sets the value in ConfigDebugIsDebuggerPresent.
func (p GuiIO) SetConfigDebugIsDebuggerPresent(value bool) {
	ptr := (*C.ImGuiIO)(unsafe.Pointer(p))
	ptr.ConfigDebugIsDebuggerPresent = (C.bool)(value)
}

// GetConfigDockingAlwaysTabBar returns the value in ConfigDockingAlwaysTabBar.
func (p GuiIO) GetConfigDockingAlwaysTabBar() bool {
	ptr := (*C.ImGuiIO)(unsafe.Pointer(p))
	return bool(ptr.ConfigDockingAlwaysTabBar)
}

// SetConfigDockingAlwaysTabBar sets the value in ConfigDockingAlwaysTabBar.
func (p GuiIO) SetConfigDockingAlwaysTabBar(value bool) {
	ptr := (*C.ImGuiIO)(unsafe.Pointer(p))
	ptr.ConfigDockingAlwaysTabBar = (C.bool)(value)
}

// GetConfigDockingNoSplit returns the value in ConfigDockingNoSplit.
func (p GuiIO) GetConfigDockingNoSplit() bool {
	ptr := (*C.ImGuiIO)(unsafe.Pointer(p))
	return bool(ptr.ConfigDockingNoSplit)
}

// SetConfigDockingNoSplit sets the value in ConfigDockingNoSplit.
func (p GuiIO) SetConfigDockingNoSplit(value bool) {
	ptr := (*C.ImGuiIO)(unsafe.Pointer(p))
	ptr.ConfigDockingNoSplit = (C.bool)(value)
}

// GetConfigDockingTransparentPayload returns the value in ConfigDockingTransparentPayload.
func (p GuiIO) GetConfigDockingTransparentPayload() bool {
	ptr := (*C.ImGuiIO)(unsafe.Pointer(p))
	return bool(ptr.ConfigDockingTransparentPayload)
}

// SetConfigDockingTransparentPayload sets the value in ConfigDockingTransparentPayload.
func (p GuiIO) SetConfigDockingTransparentPayload(value bool) {
	ptr := (*C.ImGuiIO)(unsafe.Pointer(p))
	ptr.ConfigDockingTransparentPayload = (C.bool)(value)
}

// GetConfigDockingWithShift returns the value in ConfigDockingWithShift.
func (p GuiIO) GetConfigDockingWithShift() bool {
	ptr := (*C.ImGuiIO)(unsafe.Pointer(p))
	return bool(ptr.ConfigDockingWithShift)
}

// SetConfigDockingWithShift sets the value in ConfigDockingWithShift.
func (p GuiIO) SetConfigDockingWithShift(value bool) {
	ptr := (*C.ImGuiIO)(unsafe.Pointer(p))
	ptr.ConfigDockingWithShift = (C.bool)(value)
}

// GetConfigDpiScaleFonts returns the value in ConfigDpiScaleFonts.
func (p GuiIO) GetConfigDpiScaleFonts() bool {
	ptr := (*C.ImGuiIO)(unsafe.Pointer(p))
	return bool(ptr.ConfigDpiScaleFonts)
}

// SetConfigDpiScaleFonts sets the value in ConfigDpiScaleFonts.
func (p GuiIO) SetConfigDpiScaleFonts(value bool) {
	ptr := (*C.ImGuiIO)(unsafe.Pointer(p))
	ptr.ConfigDpiScaleFonts = (C.bool)(value)
}

// GetConfigDpiScaleViewports returns the value in ConfigDpiScaleViewports.
func (p GuiIO) GetConfigDpiScaleViewports() bool {
	ptr := (*C.ImGuiIO)(unsafe.Pointer(p))
	return bool(ptr.ConfigDpiScaleViewports)
}

// SetConfigDpiScaleViewports sets the value in ConfigDpiScaleViewports.
func (p GuiIO) SetConfigDpiScaleViewports(value bool) {
	ptr := (*C.ImGuiIO)(unsafe.Pointer(p))
	ptr.ConfigDpiScaleViewports = (C.bool)(value)
}

// GetConfigDragClickToInputText returns the value in ConfigDragClickToInputText.
func (p GuiIO) GetConfigDragClickToInputText() bool {
	ptr := (*C.ImGuiIO)(unsafe.Pointer(p))
	return bool(ptr.ConfigDragClickToInputText)
}

// SetConfigDragClickToInputText sets the value in ConfigDragClickToInputText.
func (p GuiIO) SetConfigDragClickToInputText(value bool) {
	ptr := (*C.ImGuiIO)(unsafe.Pointer(p))
	ptr.ConfigDragClickToInputText = (C.bool)(value)
}

// GetConfigErrorRecovery returns the value in ConfigErrorRecovery.
func (p GuiIO) GetConfigErrorRecovery() bool {
	ptr := (*C.ImGuiIO)(unsafe.Pointer(p))
	return bool(ptr.ConfigErrorRecovery)
}

// SetConfigErrorRecovery sets the value in ConfigErrorRecovery.
func (p GuiIO) SetConfigErrorRecovery(value bool) {
	ptr := (*C.ImGuiIO)(unsafe.Pointer(p))
	ptr.ConfigErrorRecovery = (C.bool)(value)
}

// GetConfigErrorRecoveryEnableAssert returns the value in ConfigErrorRecoveryEnableAssert.
func (p GuiIO) GetConfigErrorRecoveryEnableAssert() bool {
	ptr := (*C.ImGuiIO)(unsafe.Pointer(p))
	return bool(ptr.ConfigErrorRecoveryEnableAssert)
}

// SetConfigErrorRecoveryEnableAssert sets the value in ConfigErrorRecoveryEnableAssert.
func (p GuiIO) SetConfigErrorRecoveryEnableAssert(value bool) {
	ptr := (*C.ImGuiIO)(unsafe.Pointer(p))
	ptr.ConfigErrorRecoveryEnableAssert = (C.bool)(value)
}

// GetConfigErrorRecoveryEnableDebugLog returns the value in ConfigErrorRecoveryEnableDebugLog.
func (p GuiIO) GetConfigErrorRecoveryEnableDebugLog() bool {
	ptr := (*C.ImGuiIO)(unsafe.Pointer(p))
	return bool(ptr.ConfigErrorRecoveryEnableDebugLog)
}

// SetConfigErrorRecoveryEnableDebugLog sets the value in ConfigErrorRecoveryEnableDebugLog.
func (p GuiIO) SetConfigErrorRecoveryEnableDebugLog(value bool) {
	ptr := (*C.ImGuiIO)(unsafe.Pointer(p))
	ptr.ConfigErrorRecoveryEnableDebugLog = (C.bool)(value)
}

// GetConfigErrorRecoveryEnableTooltip returns the value in ConfigErrorRecoveryEnableTooltip.
func (p GuiIO) GetConfigErrorRecoveryEnableTooltip() bool {
	ptr := (*C.ImGuiIO)(unsafe.Pointer(p))
	return bool(ptr.ConfigErrorRecoveryEnableTooltip)
}

// SetConfigErrorRecoveryEnableTooltip sets the value in ConfigErrorRecoveryEnableTooltip.
func (p GuiIO) SetConfigErrorRecoveryEnableTooltip(value bool) {
	ptr := (*C.ImGuiIO)(unsafe.Pointer(p))
	ptr.ConfigErrorRecoveryEnableTooltip = (C.bool)(value)
}

// GetConfigFlags returns the value in ConfigFlags.
func (p GuiIO) GetConfigFlags() GuiConfigFlags {
	ptr := (*C.ImGuiIO)(unsafe.Pointer(p))
	return GuiConfigFlags(ptr.ConfigFlags)
}

// SetConfigFlags sets the value in ConfigFlags.
func (p GuiIO) SetConfigFlags(value GuiConfigFlags) {
	ptr := (*C.ImGuiIO)(unsafe.Pointer(p))
	ptr.ConfigFlags = (C.ImGuiConfigFlags)(value)
}

// GetConfigInputTextCursorBlink returns the value in ConfigInputTextCursorBlink.
func (p GuiIO) GetConfigInputTextCursorBlink() bool {
	ptr := (*C.ImGuiIO)(unsafe.Pointer(p))
	return bool(ptr.ConfigInputTextCursorBlink)
}

// SetConfigInputTextCursorBlink sets the value in ConfigInputTextCursorBlink.
func (p GuiIO) SetConfigInputTextCursorBlink(value bool) {
	ptr := (*C.ImGuiIO)(unsafe.Pointer(p))
	ptr.ConfigInputTextCursorBlink = (C.bool)(value)
}

// GetConfigInputTextEnterKeepActive returns the value in ConfigInputTextEnterKeepActive.
func (p GuiIO) GetConfigInputTextEnterKeepActive() bool {
	ptr := (*C.ImGuiIO)(unsafe.Pointer(p))
	return bool(ptr.ConfigInputTextEnterKeepActive)
}

// SetConfigInputTextEnterKeepActive sets the value in ConfigInputTextEnterKeepActive.
func (p GuiIO) SetConfigInputTextEnterKeepActive(value bool) {
	ptr := (*C.ImGuiIO)(unsafe.Pointer(p))
	ptr.ConfigInputTextEnterKeepActive = (C.bool)(value)
}

// GetConfigInputTrickleEventQueue returns the value in ConfigInputTrickleEventQueue.
func (p GuiIO) GetConfigInputTrickleEventQueue() bool {
	ptr := (*C.ImGuiIO)(unsafe.Pointer(p))
	return bool(ptr.ConfigInputTrickleEventQueue)
}

// SetConfigInputTrickleEventQueue sets the value in ConfigInputTrickleEventQueue.
func (p GuiIO) SetConfigInputTrickleEventQueue(value bool) {
	ptr := (*C.ImGuiIO)(unsafe.Pointer(p))
	ptr.ConfigInputTrickleEventQueue = (C.bool)(value)
}

// GetConfigMacOSXBehaviors returns the value in ConfigMacOSXBehaviors.
func (p GuiIO) GetConfigMacOSXBehaviors() bool {
	ptr := (*C.ImGuiIO)(unsafe.Pointer(p))
	return bool(ptr.ConfigMacOSXBehaviors)
}

// SetConfigMacOSXBehaviors sets the value in ConfigMacOSXBehaviors.
func (p GuiIO) SetConfigMacOSXBehaviors(value bool) {
	ptr := (*C.ImGuiIO)(unsafe.Pointer(p))
	ptr.ConfigMacOSXBehaviors = (C.bool)(value)
}

// GetConfigMemoryCompactTimer returns the value in ConfigMemoryCompactTimer.
func (p GuiIO) GetConfigMemoryCompactTimer() float32 {
	ptr := (*C.ImGuiIO)(unsafe.Pointer(p))
	return float32(ptr.ConfigMemoryCompactTimer)
}

// SetConfigMemoryCompactTimer sets the value in ConfigMemoryCompactTimer.
func (p GuiIO) SetConfigMemoryCompactTimer(value float32) {
	ptr := (*C.ImGuiIO)(unsafe.Pointer(p))
	ptr.ConfigMemoryCompactTimer = (C.float)(value)
}

// GetConfigNavCaptureKeyboard returns the value in ConfigNavCaptureKeyboard.
func (p GuiIO) GetConfigNavCaptureKeyboard() bool {
	ptr := (*C.ImGuiIO)(unsafe.Pointer(p))
	return bool(ptr.ConfigNavCaptureKeyboard)
}

// SetConfigNavCaptureKeyboard sets the value in ConfigNavCaptureKeyboard.
func (p GuiIO) SetConfigNavCaptureKeyboard(value bool) {
	ptr := (*C.ImGuiIO)(unsafe.Pointer(p))
	ptr.ConfigNavCaptureKeyboard = (C.bool)(value)
}

// GetConfigNavCursorVisibleAlways returns the value in ConfigNavCursorVisibleAlways.
func (p GuiIO) GetConfigNavCursorVisibleAlways() bool {
	ptr := (*C.ImGuiIO)(unsafe.Pointer(p))
	return bool(ptr.ConfigNavCursorVisibleAlways)
}

// SetConfigNavCursorVisibleAlways sets the value in ConfigNavCursorVisibleAlways.
func (p GuiIO) SetConfigNavCursorVisibleAlways(value bool) {
	ptr := (*C.ImGuiIO)(unsafe.Pointer(p))
	ptr.ConfigNavCursorVisibleAlways = (C.bool)(value)
}

// GetConfigNavCursorVisibleAuto returns the value in ConfigNavCursorVisibleAuto.
func (p GuiIO) GetConfigNavCursorVisibleAuto() bool {
	ptr := (*C.ImGuiIO)(unsafe.Pointer(p))
	return bool(ptr.ConfigNavCursorVisibleAuto)
}

// SetConfigNavCursorVisibleAuto sets the value in ConfigNavCursorVisibleAuto.
func (p GuiIO) SetConfigNavCursorVisibleAuto(value bool) {
	ptr := (*C.ImGuiIO)(unsafe.Pointer(p))
	ptr.ConfigNavCursorVisibleAuto = (C.bool)(value)
}

// GetConfigNavEscapeClearFocusItem returns the value in ConfigNavEscapeClearFocusItem.
func (p GuiIO) GetConfigNavEscapeClearFocusItem() bool {
	ptr := (*C.ImGuiIO)(unsafe.Pointer(p))
	return bool(ptr.ConfigNavEscapeClearFocusItem)
}

// SetConfigNavEscapeClearFocusItem sets the value in ConfigNavEscapeClearFocusItem.
func (p GuiIO) SetConfigNavEscapeClearFocusItem(value bool) {
	ptr := (*C.ImGuiIO)(unsafe.Pointer(p))
	ptr.ConfigNavEscapeClearFocusItem = (C.bool)(value)
}

// GetConfigNavEscapeClearFocusWindow returns the value in ConfigNavEscapeClearFocusWindow.
func (p GuiIO) GetConfigNavEscapeClearFocusWindow() bool {
	ptr := (*C.ImGuiIO)(unsafe.Pointer(p))
	return bool(ptr.ConfigNavEscapeClearFocusWindow)
}

// SetConfigNavEscapeClearFocusWindow sets the value in ConfigNavEscapeClearFocusWindow.
func (p GuiIO) SetConfigNavEscapeClearFocusWindow(value bool) {
	ptr := (*C.ImGuiIO)(unsafe.Pointer(p))
	ptr.ConfigNavEscapeClearFocusWindow = (C.bool)(value)
}

// GetConfigNavMoveSetMousePos returns the value in ConfigNavMoveSetMousePos.
func (p GuiIO) GetConfigNavMoveSetMousePos() bool {
	ptr := (*C.ImGuiIO)(unsafe.Pointer(p))
	return bool(ptr.ConfigNavMoveSetMousePos)
}

// SetConfigNavMoveSetMousePos sets the value in ConfigNavMoveSetMousePos.
func (p GuiIO) SetConfigNavMoveSetMousePos(value bool) {
	ptr := (*C.ImGuiIO)(unsafe.Pointer(p))
	ptr.ConfigNavMoveSetMousePos = (C.bool)(value)
}

// GetConfigNavSwapGamepadButtons returns the value in ConfigNavSwapGamepadButtons.
func (p GuiIO) GetConfigNavSwapGamepadButtons() bool {
	ptr := (*C.ImGuiIO)(unsafe.Pointer(p))
	return bool(ptr.ConfigNavSwapGamepadButtons)
}

// SetConfigNavSwapGamepadButtons sets the value in ConfigNavSwapGamepadButtons.
func (p GuiIO) SetConfigNavSwapGamepadButtons(value bool) {
	ptr := (*C.ImGuiIO)(unsafe.Pointer(p))
	ptr.ConfigNavSwapGamepadButtons = (C.bool)(value)
}

// GetConfigScrollbarScrollByPage returns the value in ConfigScrollbarScrollByPage.
func (p GuiIO) GetConfigScrollbarScrollByPage() bool {
	ptr := (*C.ImGuiIO)(unsafe.Pointer(p))
	return bool(ptr.ConfigScrollbarScrollByPage)
}

// SetConfigScrollbarScrollByPage sets the value in ConfigScrollbarScrollByPage.
func (p GuiIO) SetConfigScrollbarScrollByPage(value bool) {
	ptr := (*C.ImGuiIO)(unsafe.Pointer(p))
	ptr.ConfigScrollbarScrollByPage = (C.bool)(value)
}

// GetConfigViewportsNoAutoMerge returns the value in ConfigViewportsNoAutoMerge.
func (p GuiIO) GetConfigViewportsNoAutoMerge() bool {
	ptr := (*C.ImGuiIO)(unsafe.Pointer(p))
	return bool(ptr.ConfigViewportsNoAutoMerge)
}

// SetConfigViewportsNoAutoMerge sets the value in ConfigViewportsNoAutoMerge.
func (p GuiIO) SetConfigViewportsNoAutoMerge(value bool) {
	ptr := (*C.ImGuiIO)(unsafe.Pointer(p))
	ptr.ConfigViewportsNoAutoMerge = (C.bool)(value)
}

// GetConfigViewportsNoDecoration returns the value in ConfigViewportsNoDecoration.
func (p GuiIO) GetConfigViewportsNoDecoration() bool {
	ptr := (*C.ImGuiIO)(unsafe.Pointer(p))
	return bool(ptr.ConfigViewportsNoDecoration)
}

// SetConfigViewportsNoDecoration sets the value in ConfigViewportsNoDecoration.
func (p GuiIO) SetConfigViewportsNoDecoration(value bool) {
	ptr := (*C.ImGuiIO)(unsafe.Pointer(p))
	ptr.ConfigViewportsNoDecoration = (C.bool)(value)
}

// GetConfigViewportsNoDefaultParent returns the value in ConfigViewportsNoDefaultParent.
func (p GuiIO) GetConfigViewportsNoDefaultParent() bool {
	ptr := (*C.ImGuiIO)(unsafe.Pointer(p))
	return bool(ptr.ConfigViewportsNoDefaultParent)
}

// SetConfigViewportsNoDefaultParent sets the value in ConfigViewportsNoDefaultParent.
func (p GuiIO) SetConfigViewportsNoDefaultParent(value bool) {
	ptr := (*C.ImGuiIO)(unsafe.Pointer(p))
	ptr.ConfigViewportsNoDefaultParent = (C.bool)(value)
}

// GetConfigViewportsNoTaskBarIcon returns the value in ConfigViewportsNoTaskBarIcon.
func (p GuiIO) GetConfigViewportsNoTaskBarIcon() bool {
	ptr := (*C.ImGuiIO)(unsafe.Pointer(p))
	return bool(ptr.ConfigViewportsNoTaskBarIcon)
}

// SetConfigViewportsNoTaskBarIcon sets the value in ConfigViewportsNoTaskBarIcon.
func (p GuiIO) SetConfigViewportsNoTaskBarIcon(value bool) {
	ptr := (*C.ImGuiIO)(unsafe.Pointer(p))
	ptr.ConfigViewportsNoTaskBarIcon = (C.bool)(value)
}

// GetConfigViewportsPlatformFocusSetsImGuiFocus returns the value in ConfigViewportsPlatformFocusSetsImGuiFocus.
func (p GuiIO) GetConfigViewportsPlatformFocusSetsImGuiFocus() bool {
	ptr := (*C.ImGuiIO)(unsafe.Pointer(p))
	return bool(ptr.ConfigViewportsPlatformFocusSetsImGuiFocus)
}

// SetConfigViewportsPlatformFocusSetsImGuiFocus sets the value in ConfigViewportsPlatformFocusSetsImGuiFocus.
func (p GuiIO) SetConfigViewportsPlatformFocusSetsImGuiFocus(value bool) {
	ptr := (*C.ImGuiIO)(unsafe.Pointer(p))
	ptr.ConfigViewportsPlatformFocusSetsImGuiFocus = (C.bool)(value)
}

// GetConfigWindowsCopyContentsWithCtrlC returns the value in ConfigWindowsCopyContentsWithCtrlC.
func (p GuiIO) GetConfigWindowsCopyContentsWithCtrlC() bool {
	ptr := (*C.ImGuiIO)(unsafe.Pointer(p))
	return bool(ptr.ConfigWindowsCopyContentsWithCtrlC)
}

// SetConfigWindowsCopyContentsWithCtrlC sets the value in ConfigWindowsCopyContentsWithCtrlC.
func (p GuiIO) SetConfigWindowsCopyContentsWithCtrlC(value bool) {
	ptr := (*C.ImGuiIO)(unsafe.Pointer(p))
	ptr.ConfigWindowsCopyContentsWithCtrlC = (C.bool)(value)
}

// GetConfigWindowsMoveFromTitleBarOnly returns the value in ConfigWindowsMoveFromTitleBarOnly.
func (p GuiIO) GetConfigWindowsMoveFromTitleBarOnly() bool {
	ptr := (*C.ImGuiIO)(unsafe.Pointer(p))
	return bool(ptr.ConfigWindowsMoveFromTitleBarOnly)
}

// SetConfigWindowsMoveFromTitleBarOnly sets the value in ConfigWindowsMoveFromTitleBarOnly.
func (p GuiIO) SetConfigWindowsMoveFromTitleBarOnly(value bool) {
	ptr := (*C.ImGuiIO)(unsafe.Pointer(p))
	ptr.ConfigWindowsMoveFromTitleBarOnly = (C.bool)(value)
}

// GetConfigWindowsResizeFromEdges returns the value in ConfigWindowsResizeFromEdges.
func (p GuiIO) GetConfigWindowsResizeFromEdges() bool {
	ptr := (*C.ImGuiIO)(unsafe.Pointer(p))
	return bool(ptr.ConfigWindowsResizeFromEdges)
}

// SetConfigWindowsResizeFromEdges sets the value in ConfigWindowsResizeFromEdges.
func (p GuiIO) SetConfigWindowsResizeFromEdges(value bool) {
	ptr := (*C.ImGuiIO)(unsafe.Pointer(p))
	ptr.ConfigWindowsResizeFromEdges = (C.bool)(value)
}

// GetCtx returns the value in Ctx.
func (p GuiIO) GetCtx() GuiContext {
	ptr := (*C.ImGuiIO)(unsafe.Pointer(p))
	return GuiContext(unsafe.Pointer(ptr.Ctx))
}

// SetCtx sets the value in Ctx.
func (p GuiIO) SetCtx(value GuiContext) {
	ptr := (*C.ImGuiIO)(unsafe.Pointer(p))
	ptr.Ctx = (*C.ImGuiContext)(unsafe.Pointer(value))
}

// GetDeltaTime returns the value in DeltaTime.
func (p GuiIO) GetDeltaTime() float32 {
	ptr := (*C.ImGuiIO)(unsafe.Pointer(p))
	return float32(ptr.DeltaTime)
}

// SetDeltaTime sets the value in DeltaTime.
func (p GuiIO) SetDeltaTime(value float32) {
	ptr := (*C.ImGuiIO)(unsafe.Pointer(p))
	ptr.DeltaTime = (C.float)(value)
}

// RefDisplayFramebufferScale returns pointer to the DisplayFramebufferScale field.
func (p GuiIO) RefDisplayFramebufferScale() Vec2 {
	return Vec2(p + GuiIO(C.offsetof_ImGuiIO_DisplayFramebufferScale))
}

// RefDisplaySize returns pointer to the DisplaySize field.
func (p GuiIO) RefDisplaySize() Vec2 {
	return Vec2(p + GuiIO(C.offsetof_ImGuiIO_DisplaySize))
}

// GetFontAllowUserScaling returns the value in FontAllowUserScaling.
func (p GuiIO) GetFontAllowUserScaling() bool {
	ptr := (*C.ImGuiIO)(unsafe.Pointer(p))
	return bool(ptr.FontAllowUserScaling)
}

// SetFontAllowUserScaling sets the value in FontAllowUserScaling.
func (p GuiIO) SetFontAllowUserScaling(value bool) {
	ptr := (*C.ImGuiIO)(unsafe.Pointer(p))
	ptr.FontAllowUserScaling = (C.bool)(value)
}

// GetFontDefault returns the value in FontDefault.
func (p GuiIO) GetFontDefault() Font {
	ptr := (*C.ImGuiIO)(unsafe.Pointer(p))
	return Font(unsafe.Pointer(ptr.FontDefault))
}

// SetFontDefault sets the value in FontDefault.
func (p GuiIO) SetFontDefault(value Font) {
	ptr := (*C.ImGuiIO)(unsafe.Pointer(p))
	ptr.FontDefault = (*C.ImFont)(unsafe.Pointer(value))
}

// GetFonts returns the value in Fonts.
func (p GuiIO) GetFonts() FontAtlas {
	ptr := (*C.ImGuiIO)(unsafe.Pointer(p))
	return FontAtlas(unsafe.Pointer(ptr.Fonts))
}

// SetFonts sets the value in Fonts.
func (p GuiIO) SetFonts(value FontAtlas) {
	ptr := (*C.ImGuiIO)(unsafe.Pointer(p))
	ptr.Fonts = (*C.ImFontAtlas)(unsafe.Pointer(value))
}

// GetFramerate returns the value in Framerate.
func (p GuiIO) GetFramerate() float32 {
	ptr := (*C.ImGuiIO)(unsafe.Pointer(p))
	return float32(ptr.Framerate)
}

// SetFramerate sets the value in Framerate.
func (p GuiIO) SetFramerate(value float32) {
	ptr := (*C.ImGuiIO)(unsafe.Pointer(p))
	ptr.Framerate = (C.float)(value)
}

// GetIniFilename returns the value in IniFilename.
func (p GuiIO) GetIniFilename() ffi.CString {
	ptr := (*C.ImGuiIO)(unsafe.Pointer(p))
	return ffi.CStringFromPtr((unsafe.Pointer)(ptr.IniFilename))
}

// SetIniFilename sets the value in IniFilename.
func (p GuiIO) SetIniFilename(value ffi.CString) {
	ptr := (*C.ImGuiIO)(unsafe.Pointer(p))
	ptr.IniFilename = (*C.char)(value.Raw())
}

// GetIniSavingRate returns the value in IniSavingRate.
func (p GuiIO) GetIniSavingRate() float32 {
	ptr := (*C.ImGuiIO)(unsafe.Pointer(p))
	return float32(ptr.IniSavingRate)
}

// SetIniSavingRate sets the value in IniSavingRate.
func (p GuiIO) SetIniSavingRate(value float32) {
	ptr := (*C.ImGuiIO)(unsafe.Pointer(p))
	ptr.IniSavingRate = (C.float)(value)
}

// RefInputQueueCharacters returns pointer to the InputQueueCharacters field.
func (p GuiIO) RefInputQueueCharacters() Vector_ImWchar {
	return Vector_ImWchar(p + GuiIO(C.offsetof_ImGuiIO_InputQueueCharacters))
}

// GuiIO.InputQueueSurrogate is unsupported: category unsupported.

// GetKeyAlt returns the value in KeyAlt.
func (p GuiIO) GetKeyAlt() bool {
	ptr := (*C.ImGuiIO)(unsafe.Pointer(p))
	return bool(ptr.KeyAlt)
}

// SetKeyAlt sets the value in KeyAlt.
func (p GuiIO) SetKeyAlt(value bool) {
	ptr := (*C.ImGuiIO)(unsafe.Pointer(p))
	ptr.KeyAlt = (C.bool)(value)
}

// GetKeyCtrl returns the value in KeyCtrl.
func (p GuiIO) GetKeyCtrl() bool {
	ptr := (*C.ImGuiIO)(unsafe.Pointer(p))
	return bool(ptr.KeyCtrl)
}

// SetKeyCtrl sets the value in KeyCtrl.
func (p GuiIO) SetKeyCtrl(value bool) {
	ptr := (*C.ImGuiIO)(unsafe.Pointer(p))
	ptr.KeyCtrl = (C.bool)(value)
}

// GuiIO.KeyMods is unsupported: category unsupported.

// GetKeyRepeatDelay returns the value in KeyRepeatDelay.
func (p GuiIO) GetKeyRepeatDelay() float32 {
	ptr := (*C.ImGuiIO)(unsafe.Pointer(p))
	return float32(ptr.KeyRepeatDelay)
}

// SetKeyRepeatDelay sets the value in KeyRepeatDelay.
func (p GuiIO) SetKeyRepeatDelay(value float32) {
	ptr := (*C.ImGuiIO)(unsafe.Pointer(p))
	ptr.KeyRepeatDelay = (C.float)(value)
}

// GetKeyRepeatRate returns the value in KeyRepeatRate.
func (p GuiIO) GetKeyRepeatRate() float32 {
	ptr := (*C.ImGuiIO)(unsafe.Pointer(p))
	return float32(ptr.KeyRepeatRate)
}

// SetKeyRepeatRate sets the value in KeyRepeatRate.
func (p GuiIO) SetKeyRepeatRate(value float32) {
	ptr := (*C.ImGuiIO)(unsafe.Pointer(p))
	ptr.KeyRepeatRate = (C.float)(value)
}

// GetKeyShift returns the value in KeyShift.
func (p GuiIO) GetKeyShift() bool {
	ptr := (*C.ImGuiIO)(unsafe.Pointer(p))
	return bool(ptr.KeyShift)
}

// SetKeyShift sets the value in KeyShift.
func (p GuiIO) SetKeyShift(value bool) {
	ptr := (*C.ImGuiIO)(unsafe.Pointer(p))
	ptr.KeyShift = (C.bool)(value)
}

// GetKeySuper returns the value in KeySuper.
func (p GuiIO) GetKeySuper() bool {
	ptr := (*C.ImGuiIO)(unsafe.Pointer(p))
	return bool(ptr.KeySuper)
}

// SetKeySuper sets the value in KeySuper.
func (p GuiIO) SetKeySuper(value bool) {
	ptr := (*C.ImGuiIO)(unsafe.Pointer(p))
	ptr.KeySuper = (C.bool)(value)
}

// GuiIO.KeysData[ImGuiKey_NamedKey_COUNT] is unsupported: category unsupported.

// GetLogFilename returns the value in LogFilename.
func (p GuiIO) GetLogFilename() ffi.CString {
	ptr := (*C.ImGuiIO)(unsafe.Pointer(p))
	return ffi.CStringFromPtr((unsafe.Pointer)(ptr.LogFilename))
}

// SetLogFilename sets the value in LogFilename.
func (p GuiIO) SetLogFilename(value ffi.CString) {
	ptr := (*C.ImGuiIO)(unsafe.Pointer(p))
	ptr.LogFilename = (*C.char)(value.Raw())
}

// GetMetricsActiveWindows returns the value in MetricsActiveWindows.
func (p GuiIO) GetMetricsActiveWindows() int32 {
	ptr := (*C.ImGuiIO)(unsafe.Pointer(p))
	return int32(ptr.MetricsActiveWindows)
}

// SetMetricsActiveWindows sets the value in MetricsActiveWindows.
func (p GuiIO) SetMetricsActiveWindows(value int32) {
	ptr := (*C.ImGuiIO)(unsafe.Pointer(p))
	ptr.MetricsActiveWindows = (C.int32_t)(value)
}

// GetMetricsRenderIndices returns the value in MetricsRenderIndices.
func (p GuiIO) GetMetricsRenderIndices() int32 {
	ptr := (*C.ImGuiIO)(unsafe.Pointer(p))
	return int32(ptr.MetricsRenderIndices)
}

// SetMetricsRenderIndices sets the value in MetricsRenderIndices.
func (p GuiIO) SetMetricsRenderIndices(value int32) {
	ptr := (*C.ImGuiIO)(unsafe.Pointer(p))
	ptr.MetricsRenderIndices = (C.int32_t)(value)
}

// GetMetricsRenderVertices returns the value in MetricsRenderVertices.
func (p GuiIO) GetMetricsRenderVertices() int32 {
	ptr := (*C.ImGuiIO)(unsafe.Pointer(p))
	return int32(ptr.MetricsRenderVertices)
}

// SetMetricsRenderVertices sets the value in MetricsRenderVertices.
func (p GuiIO) SetMetricsRenderVertices(value int32) {
	ptr := (*C.ImGuiIO)(unsafe.Pointer(p))
	ptr.MetricsRenderVertices = (C.int32_t)(value)
}

// GetMetricsRenderWindows returns the value in MetricsRenderWindows.
func (p GuiIO) GetMetricsRenderWindows() int32 {
	ptr := (*C.ImGuiIO)(unsafe.Pointer(p))
	return int32(ptr.MetricsRenderWindows)
}

// SetMetricsRenderWindows sets the value in MetricsRenderWindows.
func (p GuiIO) SetMetricsRenderWindows(value int32) {
	ptr := (*C.ImGuiIO)(unsafe.Pointer(p))
	ptr.MetricsRenderWindows = (C.int32_t)(value)
}

// GuiIO.MouseClickedCount[5] is unsupported: category unsupported.

// GuiIO.MouseClickedLastCount[5] is unsupported: category unsupported.

// GuiIO.MouseClickedPos[5] is unsupported: category unsupported.

// GuiIO.MouseClickedTime[5] is unsupported: category unsupported.

// GuiIO.MouseClicked[5] is unsupported: category unsupported.

// GetMouseCtrlLeftAsRightClick returns the value in MouseCtrlLeftAsRightClick.
func (p GuiIO) GetMouseCtrlLeftAsRightClick() bool {
	ptr := (*C.ImGuiIO)(unsafe.Pointer(p))
	return bool(ptr.MouseCtrlLeftAsRightClick)
}

// SetMouseCtrlLeftAsRightClick sets the value in MouseCtrlLeftAsRightClick.
func (p GuiIO) SetMouseCtrlLeftAsRightClick(value bool) {
	ptr := (*C.ImGuiIO)(unsafe.Pointer(p))
	ptr.MouseCtrlLeftAsRightClick = (C.bool)(value)
}

// RefMouseDelta returns pointer to the MouseDelta field.
func (p GuiIO) RefMouseDelta() Vec2 {
	return Vec2(p + GuiIO(C.offsetof_ImGuiIO_MouseDelta))
}

// GetMouseDoubleClickMaxDist returns the value in MouseDoubleClickMaxDist.
func (p GuiIO) GetMouseDoubleClickMaxDist() float32 {
	ptr := (*C.ImGuiIO)(unsafe.Pointer(p))
	return float32(ptr.MouseDoubleClickMaxDist)
}

// SetMouseDoubleClickMaxDist sets the value in MouseDoubleClickMaxDist.
func (p GuiIO) SetMouseDoubleClickMaxDist(value float32) {
	ptr := (*C.ImGuiIO)(unsafe.Pointer(p))
	ptr.MouseDoubleClickMaxDist = (C.float)(value)
}

// GetMouseDoubleClickTime returns the value in MouseDoubleClickTime.
func (p GuiIO) GetMouseDoubleClickTime() float32 {
	ptr := (*C.ImGuiIO)(unsafe.Pointer(p))
	return float32(ptr.MouseDoubleClickTime)
}

// SetMouseDoubleClickTime sets the value in MouseDoubleClickTime.
func (p GuiIO) SetMouseDoubleClickTime(value float32) {
	ptr := (*C.ImGuiIO)(unsafe.Pointer(p))
	ptr.MouseDoubleClickTime = (C.float)(value)
}

// GuiIO.MouseDoubleClicked[5] is unsupported: category unsupported.

// GuiIO.MouseDownDurationPrev[5] is unsupported: category unsupported.

// GuiIO.MouseDownDuration[5] is unsupported: category unsupported.

// GuiIO.MouseDownOwnedUnlessPopupClose[5] is unsupported: category unsupported.

// GuiIO.MouseDownOwned[5] is unsupported: category unsupported.

// GuiIO.MouseDown[5] is unsupported: category unsupported.

// GuiIO.MouseDragMaxDistanceAbs[5] is unsupported: category unsupported.

// GuiIO.MouseDragMaxDistanceSqr[5] is unsupported: category unsupported.

// GetMouseDragThreshold returns the value in MouseDragThreshold.
func (p GuiIO) GetMouseDragThreshold() float32 {
	ptr := (*C.ImGuiIO)(unsafe.Pointer(p))
	return float32(ptr.MouseDragThreshold)
}

// SetMouseDragThreshold sets the value in MouseDragThreshold.
func (p GuiIO) SetMouseDragThreshold(value float32) {
	ptr := (*C.ImGuiIO)(unsafe.Pointer(p))
	ptr.MouseDragThreshold = (C.float)(value)
}

// GetMouseDrawCursor returns the value in MouseDrawCursor.
func (p GuiIO) GetMouseDrawCursor() bool {
	ptr := (*C.ImGuiIO)(unsafe.Pointer(p))
	return bool(ptr.MouseDrawCursor)
}

// SetMouseDrawCursor sets the value in MouseDrawCursor.
func (p GuiIO) SetMouseDrawCursor(value bool) {
	ptr := (*C.ImGuiIO)(unsafe.Pointer(p))
	ptr.MouseDrawCursor = (C.bool)(value)
}

// GuiIO.MouseHoveredViewport is unsupported: category unsupported.

// RefMousePos returns pointer to the MousePos field.
func (p GuiIO) RefMousePos() Vec2 {
	return Vec2(p + GuiIO(C.offsetof_ImGuiIO_MousePos))
}

// RefMousePosPrev returns pointer to the MousePosPrev field.
func (p GuiIO) RefMousePosPrev() Vec2 {
	return Vec2(p + GuiIO(C.offsetof_ImGuiIO_MousePosPrev))
}

// GuiIO.MouseReleasedTime[5] is unsupported: category unsupported.

// GuiIO.MouseReleased[5] is unsupported: category unsupported.

// GetMouseSource returns the value in MouseSource.
func (p GuiIO) GetMouseSource() GuiMouseSource {
	ptr := (*C.ImGuiIO)(unsafe.Pointer(p))
	return GuiMouseSource(ptr.MouseSource)
}

// SetMouseSource sets the value in MouseSource.
func (p GuiIO) SetMouseSource(value GuiMouseSource) {
	ptr := (*C.ImGuiIO)(unsafe.Pointer(p))
	ptr.MouseSource = (C.ImGuiMouseSource)(value)
}

// GetMouseWheel returns the value in MouseWheel.
func (p GuiIO) GetMouseWheel() float32 {
	ptr := (*C.ImGuiIO)(unsafe.Pointer(p))
	return float32(ptr.MouseWheel)
}

// SetMouseWheel sets the value in MouseWheel.
func (p GuiIO) SetMouseWheel(value float32) {
	ptr := (*C.ImGuiIO)(unsafe.Pointer(p))
	ptr.MouseWheel = (C.float)(value)
}

// GetMouseWheelH returns the value in MouseWheelH.
func (p GuiIO) GetMouseWheelH() float32 {
	ptr := (*C.ImGuiIO)(unsafe.Pointer(p))
	return float32(ptr.MouseWheelH)
}

// SetMouseWheelH sets the value in MouseWheelH.
func (p GuiIO) SetMouseWheelH(value float32) {
	ptr := (*C.ImGuiIO)(unsafe.Pointer(p))
	ptr.MouseWheelH = (C.float)(value)
}

// GetMouseWheelRequestAxisSwap returns the value in MouseWheelRequestAxisSwap.
func (p GuiIO) GetMouseWheelRequestAxisSwap() bool {
	ptr := (*C.ImGuiIO)(unsafe.Pointer(p))
	return bool(ptr.MouseWheelRequestAxisSwap)
}

// SetMouseWheelRequestAxisSwap sets the value in MouseWheelRequestAxisSwap.
func (p GuiIO) SetMouseWheelRequestAxisSwap(value bool) {
	ptr := (*C.ImGuiIO)(unsafe.Pointer(p))
	ptr.MouseWheelRequestAxisSwap = (C.bool)(value)
}

// GetNavActive returns the value in NavActive.
func (p GuiIO) GetNavActive() bool {
	ptr := (*C.ImGuiIO)(unsafe.Pointer(p))
	return bool(ptr.NavActive)
}

// SetNavActive sets the value in NavActive.
func (p GuiIO) SetNavActive(value bool) {
	ptr := (*C.ImGuiIO)(unsafe.Pointer(p))
	ptr.NavActive = (C.bool)(value)
}

// GetNavVisible returns the value in NavVisible.
func (p GuiIO) GetNavVisible() bool {
	ptr := (*C.ImGuiIO)(unsafe.Pointer(p))
	return bool(ptr.NavVisible)
}

// SetNavVisible sets the value in NavVisible.
func (p GuiIO) SetNavVisible(value bool) {
	ptr := (*C.ImGuiIO)(unsafe.Pointer(p))
	ptr.NavVisible = (C.bool)(value)
}

// GetPenPressure returns the value in PenPressure.
func (p GuiIO) GetPenPressure() float32 {
	ptr := (*C.ImGuiIO)(unsafe.Pointer(p))
	return float32(ptr.PenPressure)
}

// SetPenPressure sets the value in PenPressure.
func (p GuiIO) SetPenPressure(value float32) {
	ptr := (*C.ImGuiIO)(unsafe.Pointer(p))
	ptr.PenPressure = (C.float)(value)
}

// GetUserData returns the value in UserData.
func (p GuiIO) GetUserData() uintptr {
	ptr := (*C.ImGuiIO)(unsafe.Pointer(p))
	return uintptr(unsafe.Pointer(ptr.UserData))
}

// SetUserData sets the value in UserData.
func (p GuiIO) SetUserData(value uintptr) {
	ptr := (*C.ImGuiIO)(unsafe.Pointer(p))
	ptr.UserData = unsafe.Pointer(value)
}

// GetWantCaptureKeyboard returns the value in WantCaptureKeyboard.
func (p GuiIO) GetWantCaptureKeyboard() bool {
	ptr := (*C.ImGuiIO)(unsafe.Pointer(p))
	return bool(ptr.WantCaptureKeyboard)
}

// SetWantCaptureKeyboard sets the value in WantCaptureKeyboard.
func (p GuiIO) SetWantCaptureKeyboard(value bool) {
	ptr := (*C.ImGuiIO)(unsafe.Pointer(p))
	ptr.WantCaptureKeyboard = (C.bool)(value)
}

// GetWantCaptureMouse returns the value in WantCaptureMouse.
func (p GuiIO) GetWantCaptureMouse() bool {
	ptr := (*C.ImGuiIO)(unsafe.Pointer(p))
	return bool(ptr.WantCaptureMouse)
}

// SetWantCaptureMouse sets the value in WantCaptureMouse.
func (p GuiIO) SetWantCaptureMouse(value bool) {
	ptr := (*C.ImGuiIO)(unsafe.Pointer(p))
	ptr.WantCaptureMouse = (C.bool)(value)
}

// GetWantCaptureMouseUnlessPopupClose returns the value in WantCaptureMouseUnlessPopupClose.
func (p GuiIO) GetWantCaptureMouseUnlessPopupClose() bool {
	ptr := (*C.ImGuiIO)(unsafe.Pointer(p))
	return bool(ptr.WantCaptureMouseUnlessPopupClose)
}

// SetWantCaptureMouseUnlessPopupClose sets the value in WantCaptureMouseUnlessPopupClose.
func (p GuiIO) SetWantCaptureMouseUnlessPopupClose(value bool) {
	ptr := (*C.ImGuiIO)(unsafe.Pointer(p))
	ptr.WantCaptureMouseUnlessPopupClose = (C.bool)(value)
}

// GetWantSaveIniSettings returns the value in WantSaveIniSettings.
func (p GuiIO) GetWantSaveIniSettings() bool {
	ptr := (*C.ImGuiIO)(unsafe.Pointer(p))
	return bool(ptr.WantSaveIniSettings)
}

// SetWantSaveIniSettings sets the value in WantSaveIniSettings.
func (p GuiIO) SetWantSaveIniSettings(value bool) {
	ptr := (*C.ImGuiIO)(unsafe.Pointer(p))
	ptr.WantSaveIniSettings = (C.bool)(value)
}

// GetWantSetMousePos returns the value in WantSetMousePos.
func (p GuiIO) GetWantSetMousePos() bool {
	ptr := (*C.ImGuiIO)(unsafe.Pointer(p))
	return bool(ptr.WantSetMousePos)
}

// SetWantSetMousePos sets the value in WantSetMousePos.
func (p GuiIO) SetWantSetMousePos(value bool) {
	ptr := (*C.ImGuiIO)(unsafe.Pointer(p))
	ptr.WantSetMousePos = (C.bool)(value)
}

// GetWantTextInput returns the value in WantTextInput.
func (p GuiIO) GetWantTextInput() bool {
	ptr := (*C.ImGuiIO)(unsafe.Pointer(p))
	return bool(ptr.WantTextInput)
}

// SetWantTextInput sets the value in WantTextInput.
func (p GuiIO) SetWantTextInput(value bool) {
	ptr := (*C.ImGuiIO)(unsafe.Pointer(p))
	ptr.WantTextInput = (C.bool)(value)
}

// GuiInputEvent wraps struct ImGuiInputEvent.
type GuiInputEvent uintptr

// GuiInputEventNil is a null pointer.
var GuiInputEventNil GuiInputEvent

// GuiInputEventSizeOf is the byte size of ImGuiInputEvent.
const GuiInputEventSizeOf = int(C.sizeof_ImGuiInputEvent)

// ImGuiInputEvent allocator is suppressed.

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p GuiInputEvent) Offset(offset int) GuiInputEvent {
	return p + GuiInputEvent(offset*GuiInputEventSizeOf)
}

// GuiInputEvent. is unsupported: category unsupported.

// GetAddedByTestEngine returns the value in AddedByTestEngine.
func (p GuiInputEvent) GetAddedByTestEngine() bool {
	ptr := (*C.ImGuiInputEvent)(unsafe.Pointer(p))
	return bool(ptr.AddedByTestEngine)
}

// SetAddedByTestEngine sets the value in AddedByTestEngine.
func (p GuiInputEvent) SetAddedByTestEngine(value bool) {
	ptr := (*C.ImGuiInputEvent)(unsafe.Pointer(p))
	ptr.AddedByTestEngine = (C.bool)(value)
}

// GetEventId returns the value in EventId.
func (p GuiInputEvent) GetEventId() uint32 {
	ptr := (*C.ImGuiInputEvent)(unsafe.Pointer(p))
	return uint32(ptr.EventId)
}

// SetEventId sets the value in EventId.
func (p GuiInputEvent) SetEventId(value uint32) {
	ptr := (*C.ImGuiInputEvent)(unsafe.Pointer(p))
	ptr.EventId = (C.uint32_t)(value)
}

// GetSource returns the value in Source.
func (p GuiInputEvent) GetSource() GuiInputSource {
	ptr := (*C.ImGuiInputEvent)(unsafe.Pointer(p))
	return GuiInputSource(ptr.Source)
}

// SetSource sets the value in Source.
func (p GuiInputEvent) SetSource(value GuiInputSource) {
	ptr := (*C.ImGuiInputEvent)(unsafe.Pointer(p))
	ptr.Source = (C.ImGuiInputSource)(value)
}

// GetType returns the value in Type.
func (p GuiInputEvent) GetType() GuiInputEventType {
	ptr := (*C.ImGuiInputEvent)(unsafe.Pointer(p))
	return GuiInputEventType(ptr.Type)
}

// SetType sets the value in Type.
func (p GuiInputEvent) SetType(value GuiInputEventType) {
	ptr := (*C.ImGuiInputEvent)(unsafe.Pointer(p))
	ptr.Type = (C.ImGuiInputEventType)(value)
}

// GuiInputEventAppFocused wraps struct ImGuiInputEventAppFocused.
type GuiInputEventAppFocused uintptr

// GuiInputEventAppFocusedNil is a null pointer.
var GuiInputEventAppFocusedNil GuiInputEventAppFocused

// GuiInputEventAppFocusedSizeOf is the byte size of ImGuiInputEventAppFocused.
const GuiInputEventAppFocusedSizeOf = int(C.sizeof_ImGuiInputEventAppFocused)

// GuiInputEventAppFocusedAlloc allocates a continuous block of GuiInputEventAppFocused.
func GuiInputEventAppFocusedAlloc(alloc ffi.Allocator, count int) GuiInputEventAppFocused {
	ptr := alloc.Allocate(GuiInputEventAppFocusedSizeOf * count)
	return GuiInputEventAppFocused(ptr)
}

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p GuiInputEventAppFocused) Offset(offset int) GuiInputEventAppFocused {
	return p + GuiInputEventAppFocused(offset*GuiInputEventAppFocusedSizeOf)
}

// GetFocused returns the value in Focused.
func (p GuiInputEventAppFocused) GetFocused() bool {
	ptr := (*C.ImGuiInputEventAppFocused)(unsafe.Pointer(p))
	return bool(ptr.Focused)
}

// SetFocused sets the value in Focused.
func (p GuiInputEventAppFocused) SetFocused(value bool) {
	ptr := (*C.ImGuiInputEventAppFocused)(unsafe.Pointer(p))
	ptr.Focused = (C.bool)(value)
}

// GuiInputEventKey wraps struct ImGuiInputEventKey.
type GuiInputEventKey uintptr

// GuiInputEventKeyNil is a null pointer.
var GuiInputEventKeyNil GuiInputEventKey

// GuiInputEventKeySizeOf is the byte size of ImGuiInputEventKey.
const GuiInputEventKeySizeOf = int(C.sizeof_ImGuiInputEventKey)

// GuiInputEventKeyAlloc allocates a continuous block of GuiInputEventKey.
func GuiInputEventKeyAlloc(alloc ffi.Allocator, count int) GuiInputEventKey {
	ptr := alloc.Allocate(GuiInputEventKeySizeOf * count)
	return GuiInputEventKey(ptr)
}

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p GuiInputEventKey) Offset(offset int) GuiInputEventKey {
	return p + GuiInputEventKey(offset*GuiInputEventKeySizeOf)
}

// GetAnalogValue returns the value in AnalogValue.
func (p GuiInputEventKey) GetAnalogValue() float32 {
	ptr := (*C.ImGuiInputEventKey)(unsafe.Pointer(p))
	return float32(ptr.AnalogValue)
}

// SetAnalogValue sets the value in AnalogValue.
func (p GuiInputEventKey) SetAnalogValue(value float32) {
	ptr := (*C.ImGuiInputEventKey)(unsafe.Pointer(p))
	ptr.AnalogValue = (C.float)(value)
}

// GetDown returns the value in Down.
func (p GuiInputEventKey) GetDown() bool {
	ptr := (*C.ImGuiInputEventKey)(unsafe.Pointer(p))
	return bool(ptr.Down)
}

// SetDown sets the value in Down.
func (p GuiInputEventKey) SetDown(value bool) {
	ptr := (*C.ImGuiInputEventKey)(unsafe.Pointer(p))
	ptr.Down = (C.bool)(value)
}

// GetKey returns the value in Key.
func (p GuiInputEventKey) GetKey() GuiKey {
	ptr := (*C.ImGuiInputEventKey)(unsafe.Pointer(p))
	return GuiKey(ptr.Key)
}

// SetKey sets the value in Key.
func (p GuiInputEventKey) SetKey(value GuiKey) {
	ptr := (*C.ImGuiInputEventKey)(unsafe.Pointer(p))
	ptr.Key = (C.ImGuiKey)(value)
}

// GuiInputEventMouseButton wraps struct ImGuiInputEventMouseButton.
type GuiInputEventMouseButton uintptr

// GuiInputEventMouseButtonNil is a null pointer.
var GuiInputEventMouseButtonNil GuiInputEventMouseButton

// GuiInputEventMouseButtonSizeOf is the byte size of ImGuiInputEventMouseButton.
const GuiInputEventMouseButtonSizeOf = int(C.sizeof_ImGuiInputEventMouseButton)

// GuiInputEventMouseButtonAlloc allocates a continuous block of GuiInputEventMouseButton.
func GuiInputEventMouseButtonAlloc(alloc ffi.Allocator, count int) GuiInputEventMouseButton {
	ptr := alloc.Allocate(GuiInputEventMouseButtonSizeOf * count)
	return GuiInputEventMouseButton(ptr)
}

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p GuiInputEventMouseButton) Offset(offset int) GuiInputEventMouseButton {
	return p + GuiInputEventMouseButton(offset*GuiInputEventMouseButtonSizeOf)
}

// GetButton returns the value in Button.
func (p GuiInputEventMouseButton) GetButton() int32 {
	ptr := (*C.ImGuiInputEventMouseButton)(unsafe.Pointer(p))
	return int32(ptr.Button)
}

// SetButton sets the value in Button.
func (p GuiInputEventMouseButton) SetButton(value int32) {
	ptr := (*C.ImGuiInputEventMouseButton)(unsafe.Pointer(p))
	ptr.Button = (C.int32_t)(value)
}

// GetDown returns the value in Down.
func (p GuiInputEventMouseButton) GetDown() bool {
	ptr := (*C.ImGuiInputEventMouseButton)(unsafe.Pointer(p))
	return bool(ptr.Down)
}

// SetDown sets the value in Down.
func (p GuiInputEventMouseButton) SetDown(value bool) {
	ptr := (*C.ImGuiInputEventMouseButton)(unsafe.Pointer(p))
	ptr.Down = (C.bool)(value)
}

// GetMouseSource returns the value in MouseSource.
func (p GuiInputEventMouseButton) GetMouseSource() GuiMouseSource {
	ptr := (*C.ImGuiInputEventMouseButton)(unsafe.Pointer(p))
	return GuiMouseSource(ptr.MouseSource)
}

// SetMouseSource sets the value in MouseSource.
func (p GuiInputEventMouseButton) SetMouseSource(value GuiMouseSource) {
	ptr := (*C.ImGuiInputEventMouseButton)(unsafe.Pointer(p))
	ptr.MouseSource = (C.ImGuiMouseSource)(value)
}

// GuiInputEventMousePos wraps struct ImGuiInputEventMousePos.
type GuiInputEventMousePos uintptr

// GuiInputEventMousePosNil is a null pointer.
var GuiInputEventMousePosNil GuiInputEventMousePos

// GuiInputEventMousePosSizeOf is the byte size of ImGuiInputEventMousePos.
const GuiInputEventMousePosSizeOf = int(C.sizeof_ImGuiInputEventMousePos)

// GuiInputEventMousePosAlloc allocates a continuous block of GuiInputEventMousePos.
func GuiInputEventMousePosAlloc(alloc ffi.Allocator, count int) GuiInputEventMousePos {
	ptr := alloc.Allocate(GuiInputEventMousePosSizeOf * count)
	return GuiInputEventMousePos(ptr)
}

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p GuiInputEventMousePos) Offset(offset int) GuiInputEventMousePos {
	return p + GuiInputEventMousePos(offset*GuiInputEventMousePosSizeOf)
}

// GetMouseSource returns the value in MouseSource.
func (p GuiInputEventMousePos) GetMouseSource() GuiMouseSource {
	ptr := (*C.ImGuiInputEventMousePos)(unsafe.Pointer(p))
	return GuiMouseSource(ptr.MouseSource)
}

// SetMouseSource sets the value in MouseSource.
func (p GuiInputEventMousePos) SetMouseSource(value GuiMouseSource) {
	ptr := (*C.ImGuiInputEventMousePos)(unsafe.Pointer(p))
	ptr.MouseSource = (C.ImGuiMouseSource)(value)
}

// GetPosX returns the value in PosX.
func (p GuiInputEventMousePos) GetPosX() float32 {
	ptr := (*C.ImGuiInputEventMousePos)(unsafe.Pointer(p))
	return float32(ptr.PosX)
}

// SetPosX sets the value in PosX.
func (p GuiInputEventMousePos) SetPosX(value float32) {
	ptr := (*C.ImGuiInputEventMousePos)(unsafe.Pointer(p))
	ptr.PosX = (C.float)(value)
}

// GetPosY returns the value in PosY.
func (p GuiInputEventMousePos) GetPosY() float32 {
	ptr := (*C.ImGuiInputEventMousePos)(unsafe.Pointer(p))
	return float32(ptr.PosY)
}

// SetPosY sets the value in PosY.
func (p GuiInputEventMousePos) SetPosY(value float32) {
	ptr := (*C.ImGuiInputEventMousePos)(unsafe.Pointer(p))
	ptr.PosY = (C.float)(value)
}

// GuiInputEventMouseViewport wraps struct ImGuiInputEventMouseViewport.
type GuiInputEventMouseViewport uintptr

// GuiInputEventMouseViewportNil is a null pointer.
var GuiInputEventMouseViewportNil GuiInputEventMouseViewport

// GuiInputEventMouseViewportSizeOf is the byte size of ImGuiInputEventMouseViewport.
const GuiInputEventMouseViewportSizeOf = int(C.sizeof_ImGuiInputEventMouseViewport)

// GuiInputEventMouseViewportAlloc allocates a continuous block of GuiInputEventMouseViewport.
func GuiInputEventMouseViewportAlloc(alloc ffi.Allocator, count int) GuiInputEventMouseViewport {
	ptr := alloc.Allocate(GuiInputEventMouseViewportSizeOf * count)
	return GuiInputEventMouseViewport(ptr)
}

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p GuiInputEventMouseViewport) Offset(offset int) GuiInputEventMouseViewport {
	return p + GuiInputEventMouseViewport(offset*GuiInputEventMouseViewportSizeOf)
}

// GuiInputEventMouseViewport.HoveredViewportID is unsupported: category unsupported.

// GuiInputEventMouseWheel wraps struct ImGuiInputEventMouseWheel.
type GuiInputEventMouseWheel uintptr

// GuiInputEventMouseWheelNil is a null pointer.
var GuiInputEventMouseWheelNil GuiInputEventMouseWheel

// GuiInputEventMouseWheelSizeOf is the byte size of ImGuiInputEventMouseWheel.
const GuiInputEventMouseWheelSizeOf = int(C.sizeof_ImGuiInputEventMouseWheel)

// GuiInputEventMouseWheelAlloc allocates a continuous block of GuiInputEventMouseWheel.
func GuiInputEventMouseWheelAlloc(alloc ffi.Allocator, count int) GuiInputEventMouseWheel {
	ptr := alloc.Allocate(GuiInputEventMouseWheelSizeOf * count)
	return GuiInputEventMouseWheel(ptr)
}

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p GuiInputEventMouseWheel) Offset(offset int) GuiInputEventMouseWheel {
	return p + GuiInputEventMouseWheel(offset*GuiInputEventMouseWheelSizeOf)
}

// GetMouseSource returns the value in MouseSource.
func (p GuiInputEventMouseWheel) GetMouseSource() GuiMouseSource {
	ptr := (*C.ImGuiInputEventMouseWheel)(unsafe.Pointer(p))
	return GuiMouseSource(ptr.MouseSource)
}

// SetMouseSource sets the value in MouseSource.
func (p GuiInputEventMouseWheel) SetMouseSource(value GuiMouseSource) {
	ptr := (*C.ImGuiInputEventMouseWheel)(unsafe.Pointer(p))
	ptr.MouseSource = (C.ImGuiMouseSource)(value)
}

// GetWheelX returns the value in WheelX.
func (p GuiInputEventMouseWheel) GetWheelX() float32 {
	ptr := (*C.ImGuiInputEventMouseWheel)(unsafe.Pointer(p))
	return float32(ptr.WheelX)
}

// SetWheelX sets the value in WheelX.
func (p GuiInputEventMouseWheel) SetWheelX(value float32) {
	ptr := (*C.ImGuiInputEventMouseWheel)(unsafe.Pointer(p))
	ptr.WheelX = (C.float)(value)
}

// GetWheelY returns the value in WheelY.
func (p GuiInputEventMouseWheel) GetWheelY() float32 {
	ptr := (*C.ImGuiInputEventMouseWheel)(unsafe.Pointer(p))
	return float32(ptr.WheelY)
}

// SetWheelY sets the value in WheelY.
func (p GuiInputEventMouseWheel) SetWheelY(value float32) {
	ptr := (*C.ImGuiInputEventMouseWheel)(unsafe.Pointer(p))
	ptr.WheelY = (C.float)(value)
}

// GuiInputEventText wraps struct ImGuiInputEventText.
type GuiInputEventText uintptr

// GuiInputEventTextNil is a null pointer.
var GuiInputEventTextNil GuiInputEventText

// GuiInputEventTextSizeOf is the byte size of ImGuiInputEventText.
const GuiInputEventTextSizeOf = int(C.sizeof_ImGuiInputEventText)

// GuiInputEventTextAlloc allocates a continuous block of GuiInputEventText.
func GuiInputEventTextAlloc(alloc ffi.Allocator, count int) GuiInputEventText {
	ptr := alloc.Allocate(GuiInputEventTextSizeOf * count)
	return GuiInputEventText(ptr)
}

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p GuiInputEventText) Offset(offset int) GuiInputEventText {
	return p + GuiInputEventText(offset*GuiInputEventTextSizeOf)
}

// GetChar returns the value in Char.
func (p GuiInputEventText) GetChar() uint32 {
	ptr := (*C.ImGuiInputEventText)(unsafe.Pointer(p))
	return uint32(ptr.Char)
}

// SetChar sets the value in Char.
func (p GuiInputEventText) SetChar(value uint32) {
	ptr := (*C.ImGuiInputEventText)(unsafe.Pointer(p))
	ptr.Char = (C.uint32_t)(value)
}

// GuiInputTextCallbackData wraps struct ImGuiInputTextCallbackData.
type GuiInputTextCallbackData uintptr

// GuiInputTextCallbackDataNil is a null pointer.
var GuiInputTextCallbackDataNil GuiInputTextCallbackData

// GuiInputTextCallbackDataSizeOf is the byte size of ImGuiInputTextCallbackData.
const GuiInputTextCallbackDataSizeOf = int(C.sizeof_ImGuiInputTextCallbackData)

// ImGuiInputTextCallbackData allocator is suppressed.

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p GuiInputTextCallbackData) Offset(offset int) GuiInputTextCallbackData {
	return p + GuiInputTextCallbackData(offset*GuiInputTextCallbackDataSizeOf)
}

// GetBuf returns the value in Buf.
func (p GuiInputTextCallbackData) GetBuf() ffi.CString {
	ptr := (*C.ImGuiInputTextCallbackData)(unsafe.Pointer(p))
	return ffi.CStringFromPtr((unsafe.Pointer)(ptr.Buf))
}

// SetBuf sets the value in Buf.
func (p GuiInputTextCallbackData) SetBuf(value ffi.CString) {
	ptr := (*C.ImGuiInputTextCallbackData)(unsafe.Pointer(p))
	ptr.Buf = (*C.char)(value.Raw())
}

// GetBufDirty returns the value in BufDirty.
func (p GuiInputTextCallbackData) GetBufDirty() bool {
	ptr := (*C.ImGuiInputTextCallbackData)(unsafe.Pointer(p))
	return bool(ptr.BufDirty)
}

// SetBufDirty sets the value in BufDirty.
func (p GuiInputTextCallbackData) SetBufDirty(value bool) {
	ptr := (*C.ImGuiInputTextCallbackData)(unsafe.Pointer(p))
	ptr.BufDirty = (C.bool)(value)
}

// GetBufSize returns the value in BufSize.
func (p GuiInputTextCallbackData) GetBufSize() int32 {
	ptr := (*C.ImGuiInputTextCallbackData)(unsafe.Pointer(p))
	return int32(ptr.BufSize)
}

// SetBufSize sets the value in BufSize.
func (p GuiInputTextCallbackData) SetBufSize(value int32) {
	ptr := (*C.ImGuiInputTextCallbackData)(unsafe.Pointer(p))
	ptr.BufSize = (C.int32_t)(value)
}

// GetBufTextLen returns the value in BufTextLen.
func (p GuiInputTextCallbackData) GetBufTextLen() int32 {
	ptr := (*C.ImGuiInputTextCallbackData)(unsafe.Pointer(p))
	return int32(ptr.BufTextLen)
}

// SetBufTextLen sets the value in BufTextLen.
func (p GuiInputTextCallbackData) SetBufTextLen(value int32) {
	ptr := (*C.ImGuiInputTextCallbackData)(unsafe.Pointer(p))
	ptr.BufTextLen = (C.int32_t)(value)
}

// GetCtx returns the value in Ctx.
func (p GuiInputTextCallbackData) GetCtx() GuiContext {
	ptr := (*C.ImGuiInputTextCallbackData)(unsafe.Pointer(p))
	return GuiContext(unsafe.Pointer(ptr.Ctx))
}

// SetCtx sets the value in Ctx.
func (p GuiInputTextCallbackData) SetCtx(value GuiContext) {
	ptr := (*C.ImGuiInputTextCallbackData)(unsafe.Pointer(p))
	ptr.Ctx = (*C.ImGuiContext)(unsafe.Pointer(value))
}

// GetCursorPos returns the value in CursorPos.
func (p GuiInputTextCallbackData) GetCursorPos() int32 {
	ptr := (*C.ImGuiInputTextCallbackData)(unsafe.Pointer(p))
	return int32(ptr.CursorPos)
}

// SetCursorPos sets the value in CursorPos.
func (p GuiInputTextCallbackData) SetCursorPos(value int32) {
	ptr := (*C.ImGuiInputTextCallbackData)(unsafe.Pointer(p))
	ptr.CursorPos = (C.int32_t)(value)
}

// GuiInputTextCallbackData.EventChar is unsupported: category unsupported.

// GetEventFlag returns the value in EventFlag.
func (p GuiInputTextCallbackData) GetEventFlag() GuiInputTextFlags {
	ptr := (*C.ImGuiInputTextCallbackData)(unsafe.Pointer(p))
	return GuiInputTextFlags(ptr.EventFlag)
}

// SetEventFlag sets the value in EventFlag.
func (p GuiInputTextCallbackData) SetEventFlag(value GuiInputTextFlags) {
	ptr := (*C.ImGuiInputTextCallbackData)(unsafe.Pointer(p))
	ptr.EventFlag = (C.ImGuiInputTextFlags)(value)
}

// GetEventKey returns the value in EventKey.
func (p GuiInputTextCallbackData) GetEventKey() GuiKey {
	ptr := (*C.ImGuiInputTextCallbackData)(unsafe.Pointer(p))
	return GuiKey(ptr.EventKey)
}

// SetEventKey sets the value in EventKey.
func (p GuiInputTextCallbackData) SetEventKey(value GuiKey) {
	ptr := (*C.ImGuiInputTextCallbackData)(unsafe.Pointer(p))
	ptr.EventKey = (C.ImGuiKey)(value)
}

// GetFlags returns the value in Flags.
func (p GuiInputTextCallbackData) GetFlags() GuiInputTextFlags {
	ptr := (*C.ImGuiInputTextCallbackData)(unsafe.Pointer(p))
	return GuiInputTextFlags(ptr.Flags)
}

// SetFlags sets the value in Flags.
func (p GuiInputTextCallbackData) SetFlags(value GuiInputTextFlags) {
	ptr := (*C.ImGuiInputTextCallbackData)(unsafe.Pointer(p))
	ptr.Flags = (C.ImGuiInputTextFlags)(value)
}

// GetSelectionEnd returns the value in SelectionEnd.
func (p GuiInputTextCallbackData) GetSelectionEnd() int32 {
	ptr := (*C.ImGuiInputTextCallbackData)(unsafe.Pointer(p))
	return int32(ptr.SelectionEnd)
}

// SetSelectionEnd sets the value in SelectionEnd.
func (p GuiInputTextCallbackData) SetSelectionEnd(value int32) {
	ptr := (*C.ImGuiInputTextCallbackData)(unsafe.Pointer(p))
	ptr.SelectionEnd = (C.int32_t)(value)
}

// GetSelectionStart returns the value in SelectionStart.
func (p GuiInputTextCallbackData) GetSelectionStart() int32 {
	ptr := (*C.ImGuiInputTextCallbackData)(unsafe.Pointer(p))
	return int32(ptr.SelectionStart)
}

// SetSelectionStart sets the value in SelectionStart.
func (p GuiInputTextCallbackData) SetSelectionStart(value int32) {
	ptr := (*C.ImGuiInputTextCallbackData)(unsafe.Pointer(p))
	ptr.SelectionStart = (C.int32_t)(value)
}

// GetUserData returns the value in UserData.
func (p GuiInputTextCallbackData) GetUserData() uintptr {
	ptr := (*C.ImGuiInputTextCallbackData)(unsafe.Pointer(p))
	return uintptr(unsafe.Pointer(ptr.UserData))
}

// SetUserData sets the value in UserData.
func (p GuiInputTextCallbackData) SetUserData(value uintptr) {
	ptr := (*C.ImGuiInputTextCallbackData)(unsafe.Pointer(p))
	ptr.UserData = unsafe.Pointer(value)
}

// GuiInputTextDeactivatedState wraps struct ImGuiInputTextDeactivatedState.
type GuiInputTextDeactivatedState uintptr

// GuiInputTextDeactivatedStateNil is a null pointer.
var GuiInputTextDeactivatedStateNil GuiInputTextDeactivatedState

// GuiInputTextDeactivatedStateSizeOf is the byte size of ImGuiInputTextDeactivatedState.
const GuiInputTextDeactivatedStateSizeOf = int(C.sizeof_ImGuiInputTextDeactivatedState)

// ImGuiInputTextDeactivatedState allocator is suppressed.

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p GuiInputTextDeactivatedState) Offset(offset int) GuiInputTextDeactivatedState {
	return p + GuiInputTextDeactivatedState(offset*GuiInputTextDeactivatedStateSizeOf)
}

// GuiInputTextDeactivatedState.ID is unsupported: category unsupported.

// RefTextA returns pointer to the TextA field.
func (p GuiInputTextDeactivatedState) RefTextA() Vector_char {
	return Vector_char(p + GuiInputTextDeactivatedState(C.offsetof_ImGuiInputTextDeactivatedState_TextA))
}

// GuiInputTextState wraps struct ImGuiInputTextState.
type GuiInputTextState uintptr

// GuiInputTextStateNil is a null pointer.
var GuiInputTextStateNil GuiInputTextState

// GuiInputTextStateSizeOf is the byte size of ImGuiInputTextState.
const GuiInputTextStateSizeOf = int(C.sizeof_ImGuiInputTextState)

// ImGuiInputTextState allocator is suppressed.

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p GuiInputTextState) Offset(offset int) GuiInputTextState {
	return p + GuiInputTextState(offset*GuiInputTextStateSizeOf)
}

// GetBufCapacity returns the value in BufCapacity.
func (p GuiInputTextState) GetBufCapacity() int32 {
	ptr := (*C.ImGuiInputTextState)(unsafe.Pointer(p))
	return int32(ptr.BufCapacity)
}

// SetBufCapacity sets the value in BufCapacity.
func (p GuiInputTextState) SetBufCapacity(value int32) {
	ptr := (*C.ImGuiInputTextState)(unsafe.Pointer(p))
	ptr.BufCapacity = (C.int32_t)(value)
}

// RefCallbackTextBackup returns pointer to the CallbackTextBackup field.
func (p GuiInputTextState) RefCallbackTextBackup() Vector_char {
	return Vector_char(p + GuiInputTextState(C.offsetof_ImGuiInputTextState_CallbackTextBackup))
}

// GetCtx returns the value in Ctx.
func (p GuiInputTextState) GetCtx() GuiContext {
	ptr := (*C.ImGuiInputTextState)(unsafe.Pointer(p))
	return GuiContext(unsafe.Pointer(ptr.Ctx))
}

// SetCtx sets the value in Ctx.
func (p GuiInputTextState) SetCtx(value GuiContext) {
	ptr := (*C.ImGuiInputTextState)(unsafe.Pointer(p))
	ptr.Ctx = (*C.ImGuiContext)(unsafe.Pointer(value))
}

// GetCursorAnim returns the value in CursorAnim.
func (p GuiInputTextState) GetCursorAnim() float32 {
	ptr := (*C.ImGuiInputTextState)(unsafe.Pointer(p))
	return float32(ptr.CursorAnim)
}

// SetCursorAnim sets the value in CursorAnim.
func (p GuiInputTextState) SetCursorAnim(value float32) {
	ptr := (*C.ImGuiInputTextState)(unsafe.Pointer(p))
	ptr.CursorAnim = (C.float)(value)
}

// GetCursorCenterY returns the value in CursorCenterY.
func (p GuiInputTextState) GetCursorCenterY() bool {
	ptr := (*C.ImGuiInputTextState)(unsafe.Pointer(p))
	return bool(ptr.CursorCenterY)
}

// SetCursorCenterY sets the value in CursorCenterY.
func (p GuiInputTextState) SetCursorCenterY(value bool) {
	ptr := (*C.ImGuiInputTextState)(unsafe.Pointer(p))
	ptr.CursorCenterY = (C.bool)(value)
}

// GetCursorFollow returns the value in CursorFollow.
func (p GuiInputTextState) GetCursorFollow() bool {
	ptr := (*C.ImGuiInputTextState)(unsafe.Pointer(p))
	return bool(ptr.CursorFollow)
}

// SetCursorFollow sets the value in CursorFollow.
func (p GuiInputTextState) SetCursorFollow(value bool) {
	ptr := (*C.ImGuiInputTextState)(unsafe.Pointer(p))
	ptr.CursorFollow = (C.bool)(value)
}

// GetEdited returns the value in Edited.
func (p GuiInputTextState) GetEdited() bool {
	ptr := (*C.ImGuiInputTextState)(unsafe.Pointer(p))
	return bool(ptr.Edited)
}

// SetEdited sets the value in Edited.
func (p GuiInputTextState) SetEdited(value bool) {
	ptr := (*C.ImGuiInputTextState)(unsafe.Pointer(p))
	ptr.Edited = (C.bool)(value)
}

// GetFlags returns the value in Flags.
func (p GuiInputTextState) GetFlags() GuiInputTextFlags {
	ptr := (*C.ImGuiInputTextState)(unsafe.Pointer(p))
	return GuiInputTextFlags(ptr.Flags)
}

// SetFlags sets the value in Flags.
func (p GuiInputTextState) SetFlags(value GuiInputTextFlags) {
	ptr := (*C.ImGuiInputTextState)(unsafe.Pointer(p))
	ptr.Flags = (C.ImGuiInputTextFlags)(value)
}

// GuiInputTextState.ID is unsupported: category unsupported.

// GuiInputTextState.LastMoveDirectionLR is unsupported: category unsupported.

// GetLineCount returns the value in LineCount.
func (p GuiInputTextState) GetLineCount() int32 {
	ptr := (*C.ImGuiInputTextState)(unsafe.Pointer(p))
	return int32(ptr.LineCount)
}

// SetLineCount sets the value in LineCount.
func (p GuiInputTextState) SetLineCount(value int32) {
	ptr := (*C.ImGuiInputTextState)(unsafe.Pointer(p))
	ptr.LineCount = (C.int32_t)(value)
}

// GetReloadSelectionEnd returns the value in ReloadSelectionEnd.
func (p GuiInputTextState) GetReloadSelectionEnd() int32 {
	ptr := (*C.ImGuiInputTextState)(unsafe.Pointer(p))
	return int32(ptr.ReloadSelectionEnd)
}

// SetReloadSelectionEnd sets the value in ReloadSelectionEnd.
func (p GuiInputTextState) SetReloadSelectionEnd(value int32) {
	ptr := (*C.ImGuiInputTextState)(unsafe.Pointer(p))
	ptr.ReloadSelectionEnd = (C.int32_t)(value)
}

// GetReloadSelectionStart returns the value in ReloadSelectionStart.
func (p GuiInputTextState) GetReloadSelectionStart() int32 {
	ptr := (*C.ImGuiInputTextState)(unsafe.Pointer(p))
	return int32(ptr.ReloadSelectionStart)
}

// SetReloadSelectionStart sets the value in ReloadSelectionStart.
func (p GuiInputTextState) SetReloadSelectionStart(value int32) {
	ptr := (*C.ImGuiInputTextState)(unsafe.Pointer(p))
	ptr.ReloadSelectionStart = (C.int32_t)(value)
}

// RefScroll returns pointer to the Scroll field.
func (p GuiInputTextState) RefScroll() Vec2 {
	return Vec2(p + GuiInputTextState(C.offsetof_ImGuiInputTextState_Scroll))
}

// GetSelectedAllMouseLock returns the value in SelectedAllMouseLock.
func (p GuiInputTextState) GetSelectedAllMouseLock() bool {
	ptr := (*C.ImGuiInputTextState)(unsafe.Pointer(p))
	return bool(ptr.SelectedAllMouseLock)
}

// SetSelectedAllMouseLock sets the value in SelectedAllMouseLock.
func (p GuiInputTextState) SetSelectedAllMouseLock(value bool) {
	ptr := (*C.ImGuiInputTextState)(unsafe.Pointer(p))
	ptr.SelectedAllMouseLock = (C.bool)(value)
}

// GuiInputTextState.Stb is unsupported: category unsupported.

// RefTextA returns pointer to the TextA field.
func (p GuiInputTextState) RefTextA() Vector_char {
	return Vector_char(p + GuiInputTextState(C.offsetof_ImGuiInputTextState_TextA))
}

// GetTextLen returns the value in TextLen.
func (p GuiInputTextState) GetTextLen() int32 {
	ptr := (*C.ImGuiInputTextState)(unsafe.Pointer(p))
	return int32(ptr.TextLen)
}

// SetTextLen sets the value in TextLen.
func (p GuiInputTextState) SetTextLen(value int32) {
	ptr := (*C.ImGuiInputTextState)(unsafe.Pointer(p))
	ptr.TextLen = (C.int32_t)(value)
}

// GetTextSrc returns the value in TextSrc.
func (p GuiInputTextState) GetTextSrc() ffi.CString {
	ptr := (*C.ImGuiInputTextState)(unsafe.Pointer(p))
	return ffi.CStringFromPtr((unsafe.Pointer)(ptr.TextSrc))
}

// SetTextSrc sets the value in TextSrc.
func (p GuiInputTextState) SetTextSrc(value ffi.CString) {
	ptr := (*C.ImGuiInputTextState)(unsafe.Pointer(p))
	ptr.TextSrc = (*C.char)(value.Raw())
}

// RefTextToRevertTo returns pointer to the TextToRevertTo field.
func (p GuiInputTextState) RefTextToRevertTo() Vector_char {
	return Vector_char(p + GuiInputTextState(C.offsetof_ImGuiInputTextState_TextToRevertTo))
}

// GetWantReloadUserBuf returns the value in WantReloadUserBuf.
func (p GuiInputTextState) GetWantReloadUserBuf() bool {
	ptr := (*C.ImGuiInputTextState)(unsafe.Pointer(p))
	return bool(ptr.WantReloadUserBuf)
}

// SetWantReloadUserBuf sets the value in WantReloadUserBuf.
func (p GuiInputTextState) SetWantReloadUserBuf(value bool) {
	ptr := (*C.ImGuiInputTextState)(unsafe.Pointer(p))
	ptr.WantReloadUserBuf = (C.bool)(value)
}

// GetWrapWidth returns the value in WrapWidth.
func (p GuiInputTextState) GetWrapWidth() float32 {
	ptr := (*C.ImGuiInputTextState)(unsafe.Pointer(p))
	return float32(ptr.WrapWidth)
}

// SetWrapWidth sets the value in WrapWidth.
func (p GuiInputTextState) SetWrapWidth(value float32) {
	ptr := (*C.ImGuiInputTextState)(unsafe.Pointer(p))
	ptr.WrapWidth = (C.float)(value)
}

// GuiKeyData wraps struct ImGuiKeyData.
type GuiKeyData uintptr

// GuiKeyDataNil is a null pointer.
var GuiKeyDataNil GuiKeyData

// GuiKeyDataSizeOf is the byte size of ImGuiKeyData.
const GuiKeyDataSizeOf = int(C.sizeof_ImGuiKeyData)

// GuiKeyDataAlloc allocates a continuous block of GuiKeyData.
func GuiKeyDataAlloc(alloc ffi.Allocator, count int) GuiKeyData {
	ptr := alloc.Allocate(GuiKeyDataSizeOf * count)
	return GuiKeyData(ptr)
}

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p GuiKeyData) Offset(offset int) GuiKeyData {
	return p + GuiKeyData(offset*GuiKeyDataSizeOf)
}

// GetAnalogValue returns the value in AnalogValue.
func (p GuiKeyData) GetAnalogValue() float32 {
	ptr := (*C.ImGuiKeyData)(unsafe.Pointer(p))
	return float32(ptr.AnalogValue)
}

// SetAnalogValue sets the value in AnalogValue.
func (p GuiKeyData) SetAnalogValue(value float32) {
	ptr := (*C.ImGuiKeyData)(unsafe.Pointer(p))
	ptr.AnalogValue = (C.float)(value)
}

// GetDown returns the value in Down.
func (p GuiKeyData) GetDown() bool {
	ptr := (*C.ImGuiKeyData)(unsafe.Pointer(p))
	return bool(ptr.Down)
}

// SetDown sets the value in Down.
func (p GuiKeyData) SetDown(value bool) {
	ptr := (*C.ImGuiKeyData)(unsafe.Pointer(p))
	ptr.Down = (C.bool)(value)
}

// GetDownDuration returns the value in DownDuration.
func (p GuiKeyData) GetDownDuration() float32 {
	ptr := (*C.ImGuiKeyData)(unsafe.Pointer(p))
	return float32(ptr.DownDuration)
}

// SetDownDuration sets the value in DownDuration.
func (p GuiKeyData) SetDownDuration(value float32) {
	ptr := (*C.ImGuiKeyData)(unsafe.Pointer(p))
	ptr.DownDuration = (C.float)(value)
}

// GetDownDurationPrev returns the value in DownDurationPrev.
func (p GuiKeyData) GetDownDurationPrev() float32 {
	ptr := (*C.ImGuiKeyData)(unsafe.Pointer(p))
	return float32(ptr.DownDurationPrev)
}

// SetDownDurationPrev sets the value in DownDurationPrev.
func (p GuiKeyData) SetDownDurationPrev(value float32) {
	ptr := (*C.ImGuiKeyData)(unsafe.Pointer(p))
	ptr.DownDurationPrev = (C.float)(value)
}

// GuiKeyOwnerData wraps struct ImGuiKeyOwnerData.
type GuiKeyOwnerData uintptr

// GuiKeyOwnerDataNil is a null pointer.
var GuiKeyOwnerDataNil GuiKeyOwnerData

// GuiKeyOwnerDataSizeOf is the byte size of ImGuiKeyOwnerData.
const GuiKeyOwnerDataSizeOf = int(C.sizeof_ImGuiKeyOwnerData)

// ImGuiKeyOwnerData allocator is suppressed.

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p GuiKeyOwnerData) Offset(offset int) GuiKeyOwnerData {
	return p + GuiKeyOwnerData(offset*GuiKeyOwnerDataSizeOf)
}

// GetLockThisFrame returns the value in LockThisFrame.
func (p GuiKeyOwnerData) GetLockThisFrame() bool {
	ptr := (*C.ImGuiKeyOwnerData)(unsafe.Pointer(p))
	return bool(ptr.LockThisFrame)
}

// SetLockThisFrame sets the value in LockThisFrame.
func (p GuiKeyOwnerData) SetLockThisFrame(value bool) {
	ptr := (*C.ImGuiKeyOwnerData)(unsafe.Pointer(p))
	ptr.LockThisFrame = (C.bool)(value)
}

// GetLockUntilRelease returns the value in LockUntilRelease.
func (p GuiKeyOwnerData) GetLockUntilRelease() bool {
	ptr := (*C.ImGuiKeyOwnerData)(unsafe.Pointer(p))
	return bool(ptr.LockUntilRelease)
}

// SetLockUntilRelease sets the value in LockUntilRelease.
func (p GuiKeyOwnerData) SetLockUntilRelease(value bool) {
	ptr := (*C.ImGuiKeyOwnerData)(unsafe.Pointer(p))
	ptr.LockUntilRelease = (C.bool)(value)
}

// GuiKeyOwnerData.OwnerCurr is unsupported: category unsupported.

// GuiKeyOwnerData.OwnerNext is unsupported: category unsupported.

// GuiKeyRoutingData wraps struct ImGuiKeyRoutingData.
type GuiKeyRoutingData uintptr

// GuiKeyRoutingDataNil is a null pointer.
var GuiKeyRoutingDataNil GuiKeyRoutingData

// GuiKeyRoutingDataSizeOf is the byte size of ImGuiKeyRoutingData.
const GuiKeyRoutingDataSizeOf = int(C.sizeof_ImGuiKeyRoutingData)

// ImGuiKeyRoutingData allocator is suppressed.

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p GuiKeyRoutingData) Offset(offset int) GuiKeyRoutingData {
	return p + GuiKeyRoutingData(offset*GuiKeyRoutingDataSizeOf)
}

// GetMods returns the value in Mods.
func (p GuiKeyRoutingData) GetMods() uint16 {
	ptr := (*C.ImGuiKeyRoutingData)(unsafe.Pointer(p))
	return uint16(ptr.Mods)
}

// SetMods sets the value in Mods.
func (p GuiKeyRoutingData) SetMods(value uint16) {
	ptr := (*C.ImGuiKeyRoutingData)(unsafe.Pointer(p))
	ptr.Mods = (C.uint16_t)(value)
}

// GuiKeyRoutingData.NextEntryIndex is unsupported: category unsupported.

// GuiKeyRoutingData.RoutingCurr is unsupported: category unsupported.

// GetRoutingCurrScore returns the value in RoutingCurrScore.
func (p GuiKeyRoutingData) GetRoutingCurrScore() uint16 {
	ptr := (*C.ImGuiKeyRoutingData)(unsafe.Pointer(p))
	return uint16(ptr.RoutingCurrScore)
}

// SetRoutingCurrScore sets the value in RoutingCurrScore.
func (p GuiKeyRoutingData) SetRoutingCurrScore(value uint16) {
	ptr := (*C.ImGuiKeyRoutingData)(unsafe.Pointer(p))
	ptr.RoutingCurrScore = (C.uint16_t)(value)
}

// GuiKeyRoutingData.RoutingNext is unsupported: category unsupported.

// GetRoutingNextScore returns the value in RoutingNextScore.
func (p GuiKeyRoutingData) GetRoutingNextScore() uint16 {
	ptr := (*C.ImGuiKeyRoutingData)(unsafe.Pointer(p))
	return uint16(ptr.RoutingNextScore)
}

// SetRoutingNextScore sets the value in RoutingNextScore.
func (p GuiKeyRoutingData) SetRoutingNextScore(value uint16) {
	ptr := (*C.ImGuiKeyRoutingData)(unsafe.Pointer(p))
	ptr.RoutingNextScore = (C.uint16_t)(value)
}

// GuiKeyRoutingTable wraps struct ImGuiKeyRoutingTable.
type GuiKeyRoutingTable uintptr

// GuiKeyRoutingTableNil is a null pointer.
var GuiKeyRoutingTableNil GuiKeyRoutingTable

// GuiKeyRoutingTableSizeOf is the byte size of ImGuiKeyRoutingTable.
const GuiKeyRoutingTableSizeOf = int(C.sizeof_ImGuiKeyRoutingTable)

// ImGuiKeyRoutingTable allocator is suppressed.

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p GuiKeyRoutingTable) Offset(offset int) GuiKeyRoutingTable {
	return p + GuiKeyRoutingTable(offset*GuiKeyRoutingTableSizeOf)
}

// RefEntries returns pointer to the Entries field.
func (p GuiKeyRoutingTable) RefEntries() Vector_ImGuiKeyRoutingData {
	return Vector_ImGuiKeyRoutingData(p + GuiKeyRoutingTable(C.offsetof_ImGuiKeyRoutingTable_Entries))
}

// RefEntriesNext returns pointer to the EntriesNext field.
func (p GuiKeyRoutingTable) RefEntriesNext() Vector_ImGuiKeyRoutingData {
	return Vector_ImGuiKeyRoutingData(p + GuiKeyRoutingTable(C.offsetof_ImGuiKeyRoutingTable_EntriesNext))
}

// GuiKeyRoutingTable.Index[ImGuiKey_NamedKey_COUNT] is unsupported: category unsupported.

// GuiLastItemData wraps struct ImGuiLastItemData.
type GuiLastItemData uintptr

// GuiLastItemDataNil is a null pointer.
var GuiLastItemDataNil GuiLastItemData

// GuiLastItemDataSizeOf is the byte size of ImGuiLastItemData.
const GuiLastItemDataSizeOf = int(C.sizeof_ImGuiLastItemData)

// ImGuiLastItemData allocator is suppressed.

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p GuiLastItemData) Offset(offset int) GuiLastItemData {
	return p + GuiLastItemData(offset*GuiLastItemDataSizeOf)
}

// RefClipRect returns pointer to the ClipRect field.
func (p GuiLastItemData) RefClipRect() Rect {
	return Rect(p + GuiLastItemData(C.offsetof_ImGuiLastItemData_ClipRect))
}

// RefDisplayRect returns pointer to the DisplayRect field.
func (p GuiLastItemData) RefDisplayRect() Rect {
	return Rect(p + GuiLastItemData(C.offsetof_ImGuiLastItemData_DisplayRect))
}

// GuiLastItemData.ID is unsupported: category unsupported.

// GetItemFlags returns the value in ItemFlags.
func (p GuiLastItemData) GetItemFlags() GuiItemFlags {
	ptr := (*C.ImGuiLastItemData)(unsafe.Pointer(p))
	return GuiItemFlags(ptr.ItemFlags)
}

// SetItemFlags sets the value in ItemFlags.
func (p GuiLastItemData) SetItemFlags(value GuiItemFlags) {
	ptr := (*C.ImGuiLastItemData)(unsafe.Pointer(p))
	ptr.ItemFlags = (C.ImGuiItemFlags)(value)
}

// RefNavRect returns pointer to the NavRect field.
func (p GuiLastItemData) RefNavRect() Rect {
	return Rect(p + GuiLastItemData(C.offsetof_ImGuiLastItemData_NavRect))
}

// RefRect returns pointer to the Rect field.
func (p GuiLastItemData) RefRect() Rect {
	return Rect(p + GuiLastItemData(C.offsetof_ImGuiLastItemData_Rect))
}

// GuiLastItemData.Shortcut is unsupported: category unsupported.

// GetStatusFlags returns the value in StatusFlags.
func (p GuiLastItemData) GetStatusFlags() GuiItemStatusFlags {
	ptr := (*C.ImGuiLastItemData)(unsafe.Pointer(p))
	return GuiItemStatusFlags(ptr.StatusFlags)
}

// SetStatusFlags sets the value in StatusFlags.
func (p GuiLastItemData) SetStatusFlags(value GuiItemStatusFlags) {
	ptr := (*C.ImGuiLastItemData)(unsafe.Pointer(p))
	ptr.StatusFlags = (C.ImGuiItemStatusFlags)(value)
}

// GuiListClipper wraps struct ImGuiListClipper.
type GuiListClipper uintptr

// GuiListClipperNil is a null pointer.
var GuiListClipperNil GuiListClipper

// GuiListClipperSizeOf is the byte size of ImGuiListClipper.
const GuiListClipperSizeOf = int(C.sizeof_ImGuiListClipper)

// ImGuiListClipper allocator is suppressed.

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p GuiListClipper) Offset(offset int) GuiListClipper {
	return p + GuiListClipper(offset*GuiListClipperSizeOf)
}

// GetCtx returns the value in Ctx.
func (p GuiListClipper) GetCtx() GuiContext {
	ptr := (*C.ImGuiListClipper)(unsafe.Pointer(p))
	return GuiContext(unsafe.Pointer(ptr.Ctx))
}

// SetCtx sets the value in Ctx.
func (p GuiListClipper) SetCtx(value GuiContext) {
	ptr := (*C.ImGuiListClipper)(unsafe.Pointer(p))
	ptr.Ctx = (*C.ImGuiContext)(unsafe.Pointer(value))
}

// GetDisplayEnd returns the value in DisplayEnd.
func (p GuiListClipper) GetDisplayEnd() int32 {
	ptr := (*C.ImGuiListClipper)(unsafe.Pointer(p))
	return int32(ptr.DisplayEnd)
}

// SetDisplayEnd sets the value in DisplayEnd.
func (p GuiListClipper) SetDisplayEnd(value int32) {
	ptr := (*C.ImGuiListClipper)(unsafe.Pointer(p))
	ptr.DisplayEnd = (C.int32_t)(value)
}

// GetDisplayStart returns the value in DisplayStart.
func (p GuiListClipper) GetDisplayStart() int32 {
	ptr := (*C.ImGuiListClipper)(unsafe.Pointer(p))
	return int32(ptr.DisplayStart)
}

// SetDisplayStart sets the value in DisplayStart.
func (p GuiListClipper) SetDisplayStart(value int32) {
	ptr := (*C.ImGuiListClipper)(unsafe.Pointer(p))
	ptr.DisplayStart = (C.int32_t)(value)
}

// GetFlags returns the value in Flags.
func (p GuiListClipper) GetFlags() GuiListClipperFlags {
	ptr := (*C.ImGuiListClipper)(unsafe.Pointer(p))
	return GuiListClipperFlags(ptr.Flags)
}

// SetFlags sets the value in Flags.
func (p GuiListClipper) SetFlags(value GuiListClipperFlags) {
	ptr := (*C.ImGuiListClipper)(unsafe.Pointer(p))
	ptr.Flags = (C.ImGuiListClipperFlags)(value)
}

// GetItemsCount returns the value in ItemsCount.
func (p GuiListClipper) GetItemsCount() int32 {
	ptr := (*C.ImGuiListClipper)(unsafe.Pointer(p))
	return int32(ptr.ItemsCount)
}

// SetItemsCount sets the value in ItemsCount.
func (p GuiListClipper) SetItemsCount(value int32) {
	ptr := (*C.ImGuiListClipper)(unsafe.Pointer(p))
	ptr.ItemsCount = (C.int32_t)(value)
}

// GetItemsHeight returns the value in ItemsHeight.
func (p GuiListClipper) GetItemsHeight() float32 {
	ptr := (*C.ImGuiListClipper)(unsafe.Pointer(p))
	return float32(ptr.ItemsHeight)
}

// SetItemsHeight sets the value in ItemsHeight.
func (p GuiListClipper) SetItemsHeight(value float32) {
	ptr := (*C.ImGuiListClipper)(unsafe.Pointer(p))
	ptr.ItemsHeight = (C.float)(value)
}

// GuiListClipper.StartPosY is unsupported: category unsupported.

// GuiListClipper.StartSeekOffsetY is unsupported: category unsupported.

// GetTempData returns the value in TempData.
func (p GuiListClipper) GetTempData() uintptr {
	ptr := (*C.ImGuiListClipper)(unsafe.Pointer(p))
	return uintptr(unsafe.Pointer(ptr.TempData))
}

// SetTempData sets the value in TempData.
func (p GuiListClipper) SetTempData(value uintptr) {
	ptr := (*C.ImGuiListClipper)(unsafe.Pointer(p))
	ptr.TempData = unsafe.Pointer(value)
}

// GuiListClipperData wraps struct ImGuiListClipperData.
type GuiListClipperData uintptr

// GuiListClipperDataNil is a null pointer.
var GuiListClipperDataNil GuiListClipperData

// GuiListClipperDataSizeOf is the byte size of ImGuiListClipperData.
const GuiListClipperDataSizeOf = int(C.sizeof_ImGuiListClipperData)

// ImGuiListClipperData allocator is suppressed.

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p GuiListClipperData) Offset(offset int) GuiListClipperData {
	return p + GuiListClipperData(offset*GuiListClipperDataSizeOf)
}

// GetItemsFrozen returns the value in ItemsFrozen.
func (p GuiListClipperData) GetItemsFrozen() int32 {
	ptr := (*C.ImGuiListClipperData)(unsafe.Pointer(p))
	return int32(ptr.ItemsFrozen)
}

// SetItemsFrozen sets the value in ItemsFrozen.
func (p GuiListClipperData) SetItemsFrozen(value int32) {
	ptr := (*C.ImGuiListClipperData)(unsafe.Pointer(p))
	ptr.ItemsFrozen = (C.int32_t)(value)
}

// GetListClipper returns the value in ListClipper.
func (p GuiListClipperData) GetListClipper() GuiListClipper {
	ptr := (*C.ImGuiListClipperData)(unsafe.Pointer(p))
	return GuiListClipper(unsafe.Pointer(ptr.ListClipper))
}

// SetListClipper sets the value in ListClipper.
func (p GuiListClipperData) SetListClipper(value GuiListClipper) {
	ptr := (*C.ImGuiListClipperData)(unsafe.Pointer(p))
	ptr.ListClipper = (*C.ImGuiListClipper)(unsafe.Pointer(value))
}

// GetLossynessOffset returns the value in LossynessOffset.
func (p GuiListClipperData) GetLossynessOffset() float32 {
	ptr := (*C.ImGuiListClipperData)(unsafe.Pointer(p))
	return float32(ptr.LossynessOffset)
}

// SetLossynessOffset sets the value in LossynessOffset.
func (p GuiListClipperData) SetLossynessOffset(value float32) {
	ptr := (*C.ImGuiListClipperData)(unsafe.Pointer(p))
	ptr.LossynessOffset = (C.float)(value)
}

// RefRanges returns pointer to the Ranges field.
func (p GuiListClipperData) RefRanges() Vector_ImGuiListClipperRange {
	return Vector_ImGuiListClipperRange(p + GuiListClipperData(C.offsetof_ImGuiListClipperData_Ranges))
}

// GetStepNo returns the value in StepNo.
func (p GuiListClipperData) GetStepNo() int32 {
	ptr := (*C.ImGuiListClipperData)(unsafe.Pointer(p))
	return int32(ptr.StepNo)
}

// SetStepNo sets the value in StepNo.
func (p GuiListClipperData) SetStepNo(value int32) {
	ptr := (*C.ImGuiListClipperData)(unsafe.Pointer(p))
	ptr.StepNo = (C.int32_t)(value)
}

// GuiListClipperRange wraps struct ImGuiListClipperRange.
type GuiListClipperRange uintptr

// GuiListClipperRangeNil is a null pointer.
var GuiListClipperRangeNil GuiListClipperRange

// GuiListClipperRangeSizeOf is the byte size of ImGuiListClipperRange.
const GuiListClipperRangeSizeOf = int(C.sizeof_ImGuiListClipperRange)

// GuiListClipperRangeAlloc allocates a continuous block of GuiListClipperRange.
func GuiListClipperRangeAlloc(alloc ffi.Allocator, count int) GuiListClipperRange {
	ptr := alloc.Allocate(GuiListClipperRangeSizeOf * count)
	return GuiListClipperRange(ptr)
}

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p GuiListClipperRange) Offset(offset int) GuiListClipperRange {
	return p + GuiListClipperRange(offset*GuiListClipperRangeSizeOf)
}

// GetMax returns the value in Max.
func (p GuiListClipperRange) GetMax() int32 {
	ptr := (*C.ImGuiListClipperRange)(unsafe.Pointer(p))
	return int32(ptr.Max)
}

// SetMax sets the value in Max.
func (p GuiListClipperRange) SetMax(value int32) {
	ptr := (*C.ImGuiListClipperRange)(unsafe.Pointer(p))
	ptr.Max = (C.int32_t)(value)
}

// GetMin returns the value in Min.
func (p GuiListClipperRange) GetMin() int32 {
	ptr := (*C.ImGuiListClipperRange)(unsafe.Pointer(p))
	return int32(ptr.Min)
}

// SetMin sets the value in Min.
func (p GuiListClipperRange) SetMin(value int32) {
	ptr := (*C.ImGuiListClipperRange)(unsafe.Pointer(p))
	ptr.Min = (C.int32_t)(value)
}

// GetPosToIndexConvert returns the value in PosToIndexConvert.
func (p GuiListClipperRange) GetPosToIndexConvert() bool {
	ptr := (*C.ImGuiListClipperRange)(unsafe.Pointer(p))
	return bool(ptr.PosToIndexConvert)
}

// SetPosToIndexConvert sets the value in PosToIndexConvert.
func (p GuiListClipperRange) SetPosToIndexConvert(value bool) {
	ptr := (*C.ImGuiListClipperRange)(unsafe.Pointer(p))
	ptr.PosToIndexConvert = (C.bool)(value)
}

// GuiListClipperRange.PosToIndexOffsetMax is unsupported: category unsupported.

// GuiListClipperRange.PosToIndexOffsetMin is unsupported: category unsupported.

// GuiLocEntry wraps struct ImGuiLocEntry.
type GuiLocEntry uintptr

// GuiLocEntryNil is a null pointer.
var GuiLocEntryNil GuiLocEntry

// GuiLocEntrySizeOf is the byte size of ImGuiLocEntry.
const GuiLocEntrySizeOf = int(C.sizeof_ImGuiLocEntry)

// GuiLocEntryAlloc allocates a continuous block of GuiLocEntry.
func GuiLocEntryAlloc(alloc ffi.Allocator, count int) GuiLocEntry {
	ptr := alloc.Allocate(GuiLocEntrySizeOf * count)
	return GuiLocEntry(ptr)
}

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p GuiLocEntry) Offset(offset int) GuiLocEntry {
	return p + GuiLocEntry(offset*GuiLocEntrySizeOf)
}

// GetKey returns the value in Key.
func (p GuiLocEntry) GetKey() GuiLocKey {
	ptr := (*C.ImGuiLocEntry)(unsafe.Pointer(p))
	return GuiLocKey(ptr.Key)
}

// SetKey sets the value in Key.
func (p GuiLocEntry) SetKey(value GuiLocKey) {
	ptr := (*C.ImGuiLocEntry)(unsafe.Pointer(p))
	ptr.Key = (C.ImGuiLocKey)(value)
}

// GetText returns the value in Text.
func (p GuiLocEntry) GetText() ffi.CString {
	ptr := (*C.ImGuiLocEntry)(unsafe.Pointer(p))
	return ffi.CStringFromPtr((unsafe.Pointer)(ptr.Text))
}

// SetText sets the value in Text.
func (p GuiLocEntry) SetText(value ffi.CString) {
	ptr := (*C.ImGuiLocEntry)(unsafe.Pointer(p))
	ptr.Text = (*C.char)(value.Raw())
}

// GuiMenuColumns wraps struct ImGuiMenuColumns.
type GuiMenuColumns uintptr

// GuiMenuColumnsNil is a null pointer.
var GuiMenuColumnsNil GuiMenuColumns

// GuiMenuColumnsSizeOf is the byte size of ImGuiMenuColumns.
const GuiMenuColumnsSizeOf = int(C.sizeof_ImGuiMenuColumns)

// ImGuiMenuColumns allocator is suppressed.

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p GuiMenuColumns) Offset(offset int) GuiMenuColumns {
	return p + GuiMenuColumns(offset*GuiMenuColumnsSizeOf)
}

// GetNextTotalWidth returns the value in NextTotalWidth.
func (p GuiMenuColumns) GetNextTotalWidth() uint32 {
	ptr := (*C.ImGuiMenuColumns)(unsafe.Pointer(p))
	return uint32(ptr.NextTotalWidth)
}

// SetNextTotalWidth sets the value in NextTotalWidth.
func (p GuiMenuColumns) SetNextTotalWidth(value uint32) {
	ptr := (*C.ImGuiMenuColumns)(unsafe.Pointer(p))
	ptr.NextTotalWidth = (C.uint32_t)(value)
}

// GetOffsetIcon returns the value in OffsetIcon.
func (p GuiMenuColumns) GetOffsetIcon() uint16 {
	ptr := (*C.ImGuiMenuColumns)(unsafe.Pointer(p))
	return uint16(ptr.OffsetIcon)
}

// SetOffsetIcon sets the value in OffsetIcon.
func (p GuiMenuColumns) SetOffsetIcon(value uint16) {
	ptr := (*C.ImGuiMenuColumns)(unsafe.Pointer(p))
	ptr.OffsetIcon = (C.uint16_t)(value)
}

// GetOffsetLabel returns the value in OffsetLabel.
func (p GuiMenuColumns) GetOffsetLabel() uint16 {
	ptr := (*C.ImGuiMenuColumns)(unsafe.Pointer(p))
	return uint16(ptr.OffsetLabel)
}

// SetOffsetLabel sets the value in OffsetLabel.
func (p GuiMenuColumns) SetOffsetLabel(value uint16) {
	ptr := (*C.ImGuiMenuColumns)(unsafe.Pointer(p))
	ptr.OffsetLabel = (C.uint16_t)(value)
}

// GetOffsetMark returns the value in OffsetMark.
func (p GuiMenuColumns) GetOffsetMark() uint16 {
	ptr := (*C.ImGuiMenuColumns)(unsafe.Pointer(p))
	return uint16(ptr.OffsetMark)
}

// SetOffsetMark sets the value in OffsetMark.
func (p GuiMenuColumns) SetOffsetMark(value uint16) {
	ptr := (*C.ImGuiMenuColumns)(unsafe.Pointer(p))
	ptr.OffsetMark = (C.uint16_t)(value)
}

// GetOffsetShortcut returns the value in OffsetShortcut.
func (p GuiMenuColumns) GetOffsetShortcut() uint16 {
	ptr := (*C.ImGuiMenuColumns)(unsafe.Pointer(p))
	return uint16(ptr.OffsetShortcut)
}

// SetOffsetShortcut sets the value in OffsetShortcut.
func (p GuiMenuColumns) SetOffsetShortcut(value uint16) {
	ptr := (*C.ImGuiMenuColumns)(unsafe.Pointer(p))
	ptr.OffsetShortcut = (C.uint16_t)(value)
}

// GetSpacing returns the value in Spacing.
func (p GuiMenuColumns) GetSpacing() uint16 {
	ptr := (*C.ImGuiMenuColumns)(unsafe.Pointer(p))
	return uint16(ptr.Spacing)
}

// SetSpacing sets the value in Spacing.
func (p GuiMenuColumns) SetSpacing(value uint16) {
	ptr := (*C.ImGuiMenuColumns)(unsafe.Pointer(p))
	ptr.Spacing = (C.uint16_t)(value)
}

// GetTotalWidth returns the value in TotalWidth.
func (p GuiMenuColumns) GetTotalWidth() uint32 {
	ptr := (*C.ImGuiMenuColumns)(unsafe.Pointer(p))
	return uint32(ptr.TotalWidth)
}

// SetTotalWidth sets the value in TotalWidth.
func (p GuiMenuColumns) SetTotalWidth(value uint32) {
	ptr := (*C.ImGuiMenuColumns)(unsafe.Pointer(p))
	ptr.TotalWidth = (C.uint32_t)(value)
}

// GuiMenuColumns.Widths[4] is unsupported: category unsupported.

// GuiMetricsConfig wraps struct ImGuiMetricsConfig.
type GuiMetricsConfig uintptr

// GuiMetricsConfigNil is a null pointer.
var GuiMetricsConfigNil GuiMetricsConfig

// GuiMetricsConfigSizeOf is the byte size of ImGuiMetricsConfig.
const GuiMetricsConfigSizeOf = int(C.sizeof_ImGuiMetricsConfig)

// GuiMetricsConfigAlloc allocates a continuous block of GuiMetricsConfig.
func GuiMetricsConfigAlloc(alloc ffi.Allocator, count int) GuiMetricsConfig {
	ptr := alloc.Allocate(GuiMetricsConfigSizeOf * count)
	return GuiMetricsConfig(ptr)
}

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p GuiMetricsConfig) Offset(offset int) GuiMetricsConfig {
	return p + GuiMetricsConfig(offset*GuiMetricsConfigSizeOf)
}

// GetHighlightMonitorIdx returns the value in HighlightMonitorIdx.
func (p GuiMetricsConfig) GetHighlightMonitorIdx() int32 {
	ptr := (*C.ImGuiMetricsConfig)(unsafe.Pointer(p))
	return int32(ptr.HighlightMonitorIdx)
}

// SetHighlightMonitorIdx sets the value in HighlightMonitorIdx.
func (p GuiMetricsConfig) SetHighlightMonitorIdx(value int32) {
	ptr := (*C.ImGuiMetricsConfig)(unsafe.Pointer(p))
	ptr.HighlightMonitorIdx = (C.int32_t)(value)
}

// GuiMetricsConfig.HighlightViewportID is unsupported: category unsupported.

// GetShowDebugLog returns the value in ShowDebugLog.
func (p GuiMetricsConfig) GetShowDebugLog() bool {
	ptr := (*C.ImGuiMetricsConfig)(unsafe.Pointer(p))
	return bool(ptr.ShowDebugLog)
}

// SetShowDebugLog sets the value in ShowDebugLog.
func (p GuiMetricsConfig) SetShowDebugLog(value bool) {
	ptr := (*C.ImGuiMetricsConfig)(unsafe.Pointer(p))
	ptr.ShowDebugLog = (C.bool)(value)
}

// GetShowDockingNodes returns the value in ShowDockingNodes.
func (p GuiMetricsConfig) GetShowDockingNodes() bool {
	ptr := (*C.ImGuiMetricsConfig)(unsafe.Pointer(p))
	return bool(ptr.ShowDockingNodes)
}

// SetShowDockingNodes sets the value in ShowDockingNodes.
func (p GuiMetricsConfig) SetShowDockingNodes(value bool) {
	ptr := (*C.ImGuiMetricsConfig)(unsafe.Pointer(p))
	ptr.ShowDockingNodes = (C.bool)(value)
}

// GetShowDrawCmdBoundingBoxes returns the value in ShowDrawCmdBoundingBoxes.
func (p GuiMetricsConfig) GetShowDrawCmdBoundingBoxes() bool {
	ptr := (*C.ImGuiMetricsConfig)(unsafe.Pointer(p))
	return bool(ptr.ShowDrawCmdBoundingBoxes)
}

// SetShowDrawCmdBoundingBoxes sets the value in ShowDrawCmdBoundingBoxes.
func (p GuiMetricsConfig) SetShowDrawCmdBoundingBoxes(value bool) {
	ptr := (*C.ImGuiMetricsConfig)(unsafe.Pointer(p))
	ptr.ShowDrawCmdBoundingBoxes = (C.bool)(value)
}

// GetShowDrawCmdMesh returns the value in ShowDrawCmdMesh.
func (p GuiMetricsConfig) GetShowDrawCmdMesh() bool {
	ptr := (*C.ImGuiMetricsConfig)(unsafe.Pointer(p))
	return bool(ptr.ShowDrawCmdMesh)
}

// SetShowDrawCmdMesh sets the value in ShowDrawCmdMesh.
func (p GuiMetricsConfig) SetShowDrawCmdMesh(value bool) {
	ptr := (*C.ImGuiMetricsConfig)(unsafe.Pointer(p))
	ptr.ShowDrawCmdMesh = (C.bool)(value)
}

// GetShowFontPreview returns the value in ShowFontPreview.
func (p GuiMetricsConfig) GetShowFontPreview() bool {
	ptr := (*C.ImGuiMetricsConfig)(unsafe.Pointer(p))
	return bool(ptr.ShowFontPreview)
}

// SetShowFontPreview sets the value in ShowFontPreview.
func (p GuiMetricsConfig) SetShowFontPreview(value bool) {
	ptr := (*C.ImGuiMetricsConfig)(unsafe.Pointer(p))
	ptr.ShowFontPreview = (C.bool)(value)
}

// GetShowIDStackTool returns the value in ShowIDStackTool.
func (p GuiMetricsConfig) GetShowIDStackTool() bool {
	ptr := (*C.ImGuiMetricsConfig)(unsafe.Pointer(p))
	return bool(ptr.ShowIDStackTool)
}

// SetShowIDStackTool sets the value in ShowIDStackTool.
func (p GuiMetricsConfig) SetShowIDStackTool(value bool) {
	ptr := (*C.ImGuiMetricsConfig)(unsafe.Pointer(p))
	ptr.ShowIDStackTool = (C.bool)(value)
}

// GetShowTablesRects returns the value in ShowTablesRects.
func (p GuiMetricsConfig) GetShowTablesRects() bool {
	ptr := (*C.ImGuiMetricsConfig)(unsafe.Pointer(p))
	return bool(ptr.ShowTablesRects)
}

// SetShowTablesRects sets the value in ShowTablesRects.
func (p GuiMetricsConfig) SetShowTablesRects(value bool) {
	ptr := (*C.ImGuiMetricsConfig)(unsafe.Pointer(p))
	ptr.ShowTablesRects = (C.bool)(value)
}

// GetShowTablesRectsType returns the value in ShowTablesRectsType.
func (p GuiMetricsConfig) GetShowTablesRectsType() int32 {
	ptr := (*C.ImGuiMetricsConfig)(unsafe.Pointer(p))
	return int32(ptr.ShowTablesRectsType)
}

// SetShowTablesRectsType sets the value in ShowTablesRectsType.
func (p GuiMetricsConfig) SetShowTablesRectsType(value int32) {
	ptr := (*C.ImGuiMetricsConfig)(unsafe.Pointer(p))
	ptr.ShowTablesRectsType = (C.int32_t)(value)
}

// GetShowTextEncodingViewer returns the value in ShowTextEncodingViewer.
func (p GuiMetricsConfig) GetShowTextEncodingViewer() bool {
	ptr := (*C.ImGuiMetricsConfig)(unsafe.Pointer(p))
	return bool(ptr.ShowTextEncodingViewer)
}

// SetShowTextEncodingViewer sets the value in ShowTextEncodingViewer.
func (p GuiMetricsConfig) SetShowTextEncodingViewer(value bool) {
	ptr := (*C.ImGuiMetricsConfig)(unsafe.Pointer(p))
	ptr.ShowTextEncodingViewer = (C.bool)(value)
}

// GetShowTextureUsedRect returns the value in ShowTextureUsedRect.
func (p GuiMetricsConfig) GetShowTextureUsedRect() bool {
	ptr := (*C.ImGuiMetricsConfig)(unsafe.Pointer(p))
	return bool(ptr.ShowTextureUsedRect)
}

// SetShowTextureUsedRect sets the value in ShowTextureUsedRect.
func (p GuiMetricsConfig) SetShowTextureUsedRect(value bool) {
	ptr := (*C.ImGuiMetricsConfig)(unsafe.Pointer(p))
	ptr.ShowTextureUsedRect = (C.bool)(value)
}

// GetShowWindowsBeginOrder returns the value in ShowWindowsBeginOrder.
func (p GuiMetricsConfig) GetShowWindowsBeginOrder() bool {
	ptr := (*C.ImGuiMetricsConfig)(unsafe.Pointer(p))
	return bool(ptr.ShowWindowsBeginOrder)
}

// SetShowWindowsBeginOrder sets the value in ShowWindowsBeginOrder.
func (p GuiMetricsConfig) SetShowWindowsBeginOrder(value bool) {
	ptr := (*C.ImGuiMetricsConfig)(unsafe.Pointer(p))
	ptr.ShowWindowsBeginOrder = (C.bool)(value)
}

// GetShowWindowsRects returns the value in ShowWindowsRects.
func (p GuiMetricsConfig) GetShowWindowsRects() bool {
	ptr := (*C.ImGuiMetricsConfig)(unsafe.Pointer(p))
	return bool(ptr.ShowWindowsRects)
}

// SetShowWindowsRects sets the value in ShowWindowsRects.
func (p GuiMetricsConfig) SetShowWindowsRects(value bool) {
	ptr := (*C.ImGuiMetricsConfig)(unsafe.Pointer(p))
	ptr.ShowWindowsRects = (C.bool)(value)
}

// GetShowWindowsRectsType returns the value in ShowWindowsRectsType.
func (p GuiMetricsConfig) GetShowWindowsRectsType() int32 {
	ptr := (*C.ImGuiMetricsConfig)(unsafe.Pointer(p))
	return int32(ptr.ShowWindowsRectsType)
}

// SetShowWindowsRectsType sets the value in ShowWindowsRectsType.
func (p GuiMetricsConfig) SetShowWindowsRectsType(value int32) {
	ptr := (*C.ImGuiMetricsConfig)(unsafe.Pointer(p))
	ptr.ShowWindowsRectsType = (C.int32_t)(value)
}

// GuiMultiSelectIO wraps struct ImGuiMultiSelectIO.
type GuiMultiSelectIO uintptr

// GuiMultiSelectIONil is a null pointer.
var GuiMultiSelectIONil GuiMultiSelectIO

// GuiMultiSelectIOSizeOf is the byte size of ImGuiMultiSelectIO.
const GuiMultiSelectIOSizeOf = int(C.sizeof_ImGuiMultiSelectIO)

// GuiMultiSelectIOAlloc allocates a continuous block of GuiMultiSelectIO.
func GuiMultiSelectIOAlloc(alloc ffi.Allocator, count int) GuiMultiSelectIO {
	ptr := alloc.Allocate(GuiMultiSelectIOSizeOf * count)
	return GuiMultiSelectIO(ptr)
}

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p GuiMultiSelectIO) Offset(offset int) GuiMultiSelectIO {
	return p + GuiMultiSelectIO(offset*GuiMultiSelectIOSizeOf)
}

// GetItemsCount returns the value in ItemsCount.
func (p GuiMultiSelectIO) GetItemsCount() int32 {
	ptr := (*C.ImGuiMultiSelectIO)(unsafe.Pointer(p))
	return int32(ptr.ItemsCount)
}

// SetItemsCount sets the value in ItemsCount.
func (p GuiMultiSelectIO) SetItemsCount(value int32) {
	ptr := (*C.ImGuiMultiSelectIO)(unsafe.Pointer(p))
	ptr.ItemsCount = (C.int32_t)(value)
}

// GuiMultiSelectIO.NavIdItem is unsupported: category unsupported.

// GetNavIdSelected returns the value in NavIdSelected.
func (p GuiMultiSelectIO) GetNavIdSelected() bool {
	ptr := (*C.ImGuiMultiSelectIO)(unsafe.Pointer(p))
	return bool(ptr.NavIdSelected)
}

// SetNavIdSelected sets the value in NavIdSelected.
func (p GuiMultiSelectIO) SetNavIdSelected(value bool) {
	ptr := (*C.ImGuiMultiSelectIO)(unsafe.Pointer(p))
	ptr.NavIdSelected = (C.bool)(value)
}

// GuiMultiSelectIO.RangeSrcItem is unsupported: category unsupported.

// GetRangeSrcReset returns the value in RangeSrcReset.
func (p GuiMultiSelectIO) GetRangeSrcReset() bool {
	ptr := (*C.ImGuiMultiSelectIO)(unsafe.Pointer(p))
	return bool(ptr.RangeSrcReset)
}

// SetRangeSrcReset sets the value in RangeSrcReset.
func (p GuiMultiSelectIO) SetRangeSrcReset(value bool) {
	ptr := (*C.ImGuiMultiSelectIO)(unsafe.Pointer(p))
	ptr.RangeSrcReset = (C.bool)(value)
}

// RefRequests returns pointer to the Requests field.
func (p GuiMultiSelectIO) RefRequests() Vector_ImGuiSelectionRequest {
	return Vector_ImGuiSelectionRequest(p + GuiMultiSelectIO(C.offsetof_ImGuiMultiSelectIO_Requests))
}

// GuiMultiSelectState wraps struct ImGuiMultiSelectState.
type GuiMultiSelectState uintptr

// GuiMultiSelectStateNil is a null pointer.
var GuiMultiSelectStateNil GuiMultiSelectState

// GuiMultiSelectStateSizeOf is the byte size of ImGuiMultiSelectState.
const GuiMultiSelectStateSizeOf = int(C.sizeof_ImGuiMultiSelectState)

// ImGuiMultiSelectState allocator is suppressed.

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p GuiMultiSelectState) Offset(offset int) GuiMultiSelectState {
	return p + GuiMultiSelectState(offset*GuiMultiSelectStateSizeOf)
}

// GuiMultiSelectState.ID is unsupported: category unsupported.

// GetLastFrameActive returns the value in LastFrameActive.
func (p GuiMultiSelectState) GetLastFrameActive() int32 {
	ptr := (*C.ImGuiMultiSelectState)(unsafe.Pointer(p))
	return int32(ptr.LastFrameActive)
}

// SetLastFrameActive sets the value in LastFrameActive.
func (p GuiMultiSelectState) SetLastFrameActive(value int32) {
	ptr := (*C.ImGuiMultiSelectState)(unsafe.Pointer(p))
	ptr.LastFrameActive = (C.int32_t)(value)
}

// GetLastSelectionSize returns the value in LastSelectionSize.
func (p GuiMultiSelectState) GetLastSelectionSize() int32 {
	ptr := (*C.ImGuiMultiSelectState)(unsafe.Pointer(p))
	return int32(ptr.LastSelectionSize)
}

// SetLastSelectionSize sets the value in LastSelectionSize.
func (p GuiMultiSelectState) SetLastSelectionSize(value int32) {
	ptr := (*C.ImGuiMultiSelectState)(unsafe.Pointer(p))
	ptr.LastSelectionSize = (C.int32_t)(value)
}

// GuiMultiSelectState.NavIdItem is unsupported: category unsupported.

// GuiMultiSelectState.NavIdSelected is unsupported: category unsupported.

// GuiMultiSelectState.RangeSelected is unsupported: category unsupported.

// GuiMultiSelectState.RangeSrcItem is unsupported: category unsupported.

// GetWindow returns the value in Window.
func (p GuiMultiSelectState) GetWindow() GuiWindow {
	ptr := (*C.ImGuiMultiSelectState)(unsafe.Pointer(p))
	return GuiWindow(unsafe.Pointer(ptr.Window))
}

// SetWindow sets the value in Window.
func (p GuiMultiSelectState) SetWindow(value GuiWindow) {
	ptr := (*C.ImGuiMultiSelectState)(unsafe.Pointer(p))
	ptr.Window = (*C.ImGuiWindow)(unsafe.Pointer(value))
}

// GuiMultiSelectTempData wraps struct ImGuiMultiSelectTempData.
type GuiMultiSelectTempData uintptr

// GuiMultiSelectTempDataNil is a null pointer.
var GuiMultiSelectTempDataNil GuiMultiSelectTempData

// GuiMultiSelectTempDataSizeOf is the byte size of ImGuiMultiSelectTempData.
const GuiMultiSelectTempDataSizeOf = int(C.sizeof_ImGuiMultiSelectTempData)

// ImGuiMultiSelectTempData allocator is suppressed.

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p GuiMultiSelectTempData) Offset(offset int) GuiMultiSelectTempData {
	return p + GuiMultiSelectTempData(offset*GuiMultiSelectTempDataSizeOf)
}

// RefBackupCursorMaxPos returns pointer to the BackupCursorMaxPos field.
func (p GuiMultiSelectTempData) RefBackupCursorMaxPos() Vec2 {
	return Vec2(p + GuiMultiSelectTempData(C.offsetof_ImGuiMultiSelectTempData_BackupCursorMaxPos))
}

// GuiMultiSelectTempData.BoxSelectId is unsupported: category unsupported.

// GetFlags returns the value in Flags.
func (p GuiMultiSelectTempData) GetFlags() GuiMultiSelectFlags {
	ptr := (*C.ImGuiMultiSelectTempData)(unsafe.Pointer(p))
	return GuiMultiSelectFlags(ptr.Flags)
}

// SetFlags sets the value in Flags.
func (p GuiMultiSelectTempData) SetFlags(value GuiMultiSelectFlags) {
	ptr := (*C.ImGuiMultiSelectTempData)(unsafe.Pointer(p))
	ptr.Flags = (C.ImGuiMultiSelectFlags)(value)
}

// GuiMultiSelectTempData.FocusScopeId is unsupported: category unsupported.

// RefIO returns pointer to the IO field.
func (p GuiMultiSelectTempData) RefIO() GuiMultiSelectIO {
	return GuiMultiSelectIO(p + GuiMultiSelectTempData(C.offsetof_ImGuiMultiSelectTempData_IO))
}

// GetIsEndIO returns the value in IsEndIO.
func (p GuiMultiSelectTempData) GetIsEndIO() bool {
	ptr := (*C.ImGuiMultiSelectTempData)(unsafe.Pointer(p))
	return bool(ptr.IsEndIO)
}

// SetIsEndIO sets the value in IsEndIO.
func (p GuiMultiSelectTempData) SetIsEndIO(value bool) {
	ptr := (*C.ImGuiMultiSelectTempData)(unsafe.Pointer(p))
	ptr.IsEndIO = (C.bool)(value)
}

// GetIsFocused returns the value in IsFocused.
func (p GuiMultiSelectTempData) GetIsFocused() bool {
	ptr := (*C.ImGuiMultiSelectTempData)(unsafe.Pointer(p))
	return bool(ptr.IsFocused)
}

// SetIsFocused sets the value in IsFocused.
func (p GuiMultiSelectTempData) SetIsFocused(value bool) {
	ptr := (*C.ImGuiMultiSelectTempData)(unsafe.Pointer(p))
	ptr.IsFocused = (C.bool)(value)
}

// GetIsKeyboardSetRange returns the value in IsKeyboardSetRange.
func (p GuiMultiSelectTempData) GetIsKeyboardSetRange() bool {
	ptr := (*C.ImGuiMultiSelectTempData)(unsafe.Pointer(p))
	return bool(ptr.IsKeyboardSetRange)
}

// SetIsKeyboardSetRange sets the value in IsKeyboardSetRange.
func (p GuiMultiSelectTempData) SetIsKeyboardSetRange(value bool) {
	ptr := (*C.ImGuiMultiSelectTempData)(unsafe.Pointer(p))
	ptr.IsKeyboardSetRange = (C.bool)(value)
}

// GuiMultiSelectTempData.KeyMods is unsupported: category unsupported.

// GuiMultiSelectTempData.LastSubmittedItem is unsupported: category unsupported.

// GuiMultiSelectTempData.LoopRequestSetAll is unsupported: category unsupported.

// GetNavIdPassedBy returns the value in NavIdPassedBy.
func (p GuiMultiSelectTempData) GetNavIdPassedBy() bool {
	ptr := (*C.ImGuiMultiSelectTempData)(unsafe.Pointer(p))
	return bool(ptr.NavIdPassedBy)
}

// SetNavIdPassedBy sets the value in NavIdPassedBy.
func (p GuiMultiSelectTempData) SetNavIdPassedBy(value bool) {
	ptr := (*C.ImGuiMultiSelectTempData)(unsafe.Pointer(p))
	ptr.NavIdPassedBy = (C.bool)(value)
}

// GetRangeDstPassedBy returns the value in RangeDstPassedBy.
func (p GuiMultiSelectTempData) GetRangeDstPassedBy() bool {
	ptr := (*C.ImGuiMultiSelectTempData)(unsafe.Pointer(p))
	return bool(ptr.RangeDstPassedBy)
}

// SetRangeDstPassedBy sets the value in RangeDstPassedBy.
func (p GuiMultiSelectTempData) SetRangeDstPassedBy(value bool) {
	ptr := (*C.ImGuiMultiSelectTempData)(unsafe.Pointer(p))
	ptr.RangeDstPassedBy = (C.bool)(value)
}

// GetRangeSrcPassedBy returns the value in RangeSrcPassedBy.
func (p GuiMultiSelectTempData) GetRangeSrcPassedBy() bool {
	ptr := (*C.ImGuiMultiSelectTempData)(unsafe.Pointer(p))
	return bool(ptr.RangeSrcPassedBy)
}

// SetRangeSrcPassedBy sets the value in RangeSrcPassedBy.
func (p GuiMultiSelectTempData) SetRangeSrcPassedBy(value bool) {
	ptr := (*C.ImGuiMultiSelectTempData)(unsafe.Pointer(p))
	ptr.RangeSrcPassedBy = (C.bool)(value)
}

// RefScopeRectMin returns pointer to the ScopeRectMin field.
func (p GuiMultiSelectTempData) RefScopeRectMin() Vec2 {
	return Vec2(p + GuiMultiSelectTempData(C.offsetof_ImGuiMultiSelectTempData_ScopeRectMin))
}

// GetStorage returns the value in Storage.
func (p GuiMultiSelectTempData) GetStorage() GuiMultiSelectState {
	ptr := (*C.ImGuiMultiSelectTempData)(unsafe.Pointer(p))
	return GuiMultiSelectState(unsafe.Pointer(ptr.Storage))
}

// SetStorage sets the value in Storage.
func (p GuiMultiSelectTempData) SetStorage(value GuiMultiSelectState) {
	ptr := (*C.ImGuiMultiSelectTempData)(unsafe.Pointer(p))
	ptr.Storage = (*C.ImGuiMultiSelectState)(unsafe.Pointer(value))
}

// GuiNavItemData wraps struct ImGuiNavItemData.
type GuiNavItemData uintptr

// GuiNavItemDataNil is a null pointer.
var GuiNavItemDataNil GuiNavItemData

// GuiNavItemDataSizeOf is the byte size of ImGuiNavItemData.
const GuiNavItemDataSizeOf = int(C.sizeof_ImGuiNavItemData)

// ImGuiNavItemData allocator is suppressed.

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p GuiNavItemData) Offset(offset int) GuiNavItemData {
	return p + GuiNavItemData(offset*GuiNavItemDataSizeOf)
}

// GetDistAxial returns the value in DistAxial.
func (p GuiNavItemData) GetDistAxial() float32 {
	ptr := (*C.ImGuiNavItemData)(unsafe.Pointer(p))
	return float32(ptr.DistAxial)
}

// SetDistAxial sets the value in DistAxial.
func (p GuiNavItemData) SetDistAxial(value float32) {
	ptr := (*C.ImGuiNavItemData)(unsafe.Pointer(p))
	ptr.DistAxial = (C.float)(value)
}

// GetDistBox returns the value in DistBox.
func (p GuiNavItemData) GetDistBox() float32 {
	ptr := (*C.ImGuiNavItemData)(unsafe.Pointer(p))
	return float32(ptr.DistBox)
}

// SetDistBox sets the value in DistBox.
func (p GuiNavItemData) SetDistBox(value float32) {
	ptr := (*C.ImGuiNavItemData)(unsafe.Pointer(p))
	ptr.DistBox = (C.float)(value)
}

// GetDistCenter returns the value in DistCenter.
func (p GuiNavItemData) GetDistCenter() float32 {
	ptr := (*C.ImGuiNavItemData)(unsafe.Pointer(p))
	return float32(ptr.DistCenter)
}

// SetDistCenter sets the value in DistCenter.
func (p GuiNavItemData) SetDistCenter(value float32) {
	ptr := (*C.ImGuiNavItemData)(unsafe.Pointer(p))
	ptr.DistCenter = (C.float)(value)
}

// GuiNavItemData.FocusScopeId is unsupported: category unsupported.

// GuiNavItemData.ID is unsupported: category unsupported.

// GetItemFlags returns the value in ItemFlags.
func (p GuiNavItemData) GetItemFlags() GuiItemFlags {
	ptr := (*C.ImGuiNavItemData)(unsafe.Pointer(p))
	return GuiItemFlags(ptr.ItemFlags)
}

// SetItemFlags sets the value in ItemFlags.
func (p GuiNavItemData) SetItemFlags(value GuiItemFlags) {
	ptr := (*C.ImGuiNavItemData)(unsafe.Pointer(p))
	ptr.ItemFlags = (C.ImGuiItemFlags)(value)
}

// RefRectRel returns pointer to the RectRel field.
func (p GuiNavItemData) RefRectRel() Rect {
	return Rect(p + GuiNavItemData(C.offsetof_ImGuiNavItemData_RectRel))
}

// GuiNavItemData.SelectionUserData is unsupported: category unsupported.

// GetWindow returns the value in Window.
func (p GuiNavItemData) GetWindow() GuiWindow {
	ptr := (*C.ImGuiNavItemData)(unsafe.Pointer(p))
	return GuiWindow(unsafe.Pointer(ptr.Window))
}

// SetWindow sets the value in Window.
func (p GuiNavItemData) SetWindow(value GuiWindow) {
	ptr := (*C.ImGuiNavItemData)(unsafe.Pointer(p))
	ptr.Window = (*C.ImGuiWindow)(unsafe.Pointer(value))
}

// GuiNextItemData wraps struct ImGuiNextItemData.
type GuiNextItemData uintptr

// GuiNextItemDataNil is a null pointer.
var GuiNextItemDataNil GuiNextItemData

// GuiNextItemDataSizeOf is the byte size of ImGuiNextItemData.
const GuiNextItemDataSizeOf = int(C.sizeof_ImGuiNextItemData)

// ImGuiNextItemData allocator is suppressed.

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p GuiNextItemData) Offset(offset int) GuiNextItemData {
	return p + GuiNextItemData(offset*GuiNextItemDataSizeOf)
}

// GuiNextItemData.FocusScopeId is unsupported: category unsupported.

// GetHasFlags returns the value in HasFlags.
func (p GuiNextItemData) GetHasFlags() GuiNextItemDataFlags {
	ptr := (*C.ImGuiNextItemData)(unsafe.Pointer(p))
	return GuiNextItemDataFlags(ptr.HasFlags)
}

// SetHasFlags sets the value in HasFlags.
func (p GuiNextItemData) SetHasFlags(value GuiNextItemDataFlags) {
	ptr := (*C.ImGuiNextItemData)(unsafe.Pointer(p))
	ptr.HasFlags = (C.ImGuiNextItemDataFlags)(value)
}

// GetItemFlags returns the value in ItemFlags.
func (p GuiNextItemData) GetItemFlags() GuiItemFlags {
	ptr := (*C.ImGuiNextItemData)(unsafe.Pointer(p))
	return GuiItemFlags(ptr.ItemFlags)
}

// SetItemFlags sets the value in ItemFlags.
func (p GuiNextItemData) SetItemFlags(value GuiItemFlags) {
	ptr := (*C.ImGuiNextItemData)(unsafe.Pointer(p))
	ptr.ItemFlags = (C.ImGuiItemFlags)(value)
}

// GetOpenCond returns the value in OpenCond.
func (p GuiNextItemData) GetOpenCond() uint8 {
	ptr := (*C.ImGuiNextItemData)(unsafe.Pointer(p))
	return uint8(ptr.OpenCond)
}

// SetOpenCond sets the value in OpenCond.
func (p GuiNextItemData) SetOpenCond(value uint8) {
	ptr := (*C.ImGuiNextItemData)(unsafe.Pointer(p))
	ptr.OpenCond = (C.uint8_t)(value)
}

// GetOpenVal returns the value in OpenVal.
func (p GuiNextItemData) GetOpenVal() bool {
	ptr := (*C.ImGuiNextItemData)(unsafe.Pointer(p))
	return bool(ptr.OpenVal)
}

// SetOpenVal sets the value in OpenVal.
func (p GuiNextItemData) SetOpenVal(value bool) {
	ptr := (*C.ImGuiNextItemData)(unsafe.Pointer(p))
	ptr.OpenVal = (C.bool)(value)
}

// RefRefVal returns pointer to the RefVal field.
func (p GuiNextItemData) RefRefVal() GuiDataTypeStorage {
	return GuiDataTypeStorage(p + GuiNextItemData(C.offsetof_ImGuiNextItemData_RefVal))
}

// GuiNextItemData.SelectionUserData is unsupported: category unsupported.

// GuiNextItemData.Shortcut is unsupported: category unsupported.

// GetShortcutFlags returns the value in ShortcutFlags.
func (p GuiNextItemData) GetShortcutFlags() GuiInputFlags {
	ptr := (*C.ImGuiNextItemData)(unsafe.Pointer(p))
	return GuiInputFlags(ptr.ShortcutFlags)
}

// SetShortcutFlags sets the value in ShortcutFlags.
func (p GuiNextItemData) SetShortcutFlags(value GuiInputFlags) {
	ptr := (*C.ImGuiNextItemData)(unsafe.Pointer(p))
	ptr.ShortcutFlags = (C.ImGuiInputFlags)(value)
}

// GuiNextItemData.StorageId is unsupported: category unsupported.

// GetWidth returns the value in Width.
func (p GuiNextItemData) GetWidth() float32 {
	ptr := (*C.ImGuiNextItemData)(unsafe.Pointer(p))
	return float32(ptr.Width)
}

// SetWidth sets the value in Width.
func (p GuiNextItemData) SetWidth(value float32) {
	ptr := (*C.ImGuiNextItemData)(unsafe.Pointer(p))
	ptr.Width = (C.float)(value)
}

// GuiNextWindowData wraps struct ImGuiNextWindowData.
type GuiNextWindowData uintptr

// GuiNextWindowDataNil is a null pointer.
var GuiNextWindowDataNil GuiNextWindowData

// GuiNextWindowDataSizeOf is the byte size of ImGuiNextWindowData.
const GuiNextWindowDataSizeOf = int(C.sizeof_ImGuiNextWindowData)

// ImGuiNextWindowData allocator is suppressed.

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p GuiNextWindowData) Offset(offset int) GuiNextWindowData {
	return p + GuiNextWindowData(offset*GuiNextWindowDataSizeOf)
}

// GetBgAlphaVal returns the value in BgAlphaVal.
func (p GuiNextWindowData) GetBgAlphaVal() float32 {
	ptr := (*C.ImGuiNextWindowData)(unsafe.Pointer(p))
	return float32(ptr.BgAlphaVal)
}

// SetBgAlphaVal sets the value in BgAlphaVal.
func (p GuiNextWindowData) SetBgAlphaVal(value float32) {
	ptr := (*C.ImGuiNextWindowData)(unsafe.Pointer(p))
	ptr.BgAlphaVal = (C.float)(value)
}

// GetChildFlags returns the value in ChildFlags.
func (p GuiNextWindowData) GetChildFlags() GuiChildFlags {
	ptr := (*C.ImGuiNextWindowData)(unsafe.Pointer(p))
	return GuiChildFlags(ptr.ChildFlags)
}

// SetChildFlags sets the value in ChildFlags.
func (p GuiNextWindowData) SetChildFlags(value GuiChildFlags) {
	ptr := (*C.ImGuiNextWindowData)(unsafe.Pointer(p))
	ptr.ChildFlags = (C.ImGuiChildFlags)(value)
}

// GetCollapsedCond returns the value in CollapsedCond.
func (p GuiNextWindowData) GetCollapsedCond() GuiCond {
	ptr := (*C.ImGuiNextWindowData)(unsafe.Pointer(p))
	return GuiCond(ptr.CollapsedCond)
}

// SetCollapsedCond sets the value in CollapsedCond.
func (p GuiNextWindowData) SetCollapsedCond(value GuiCond) {
	ptr := (*C.ImGuiNextWindowData)(unsafe.Pointer(p))
	ptr.CollapsedCond = (C.ImGuiCond)(value)
}

// GetCollapsedVal returns the value in CollapsedVal.
func (p GuiNextWindowData) GetCollapsedVal() bool {
	ptr := (*C.ImGuiNextWindowData)(unsafe.Pointer(p))
	return bool(ptr.CollapsedVal)
}

// SetCollapsedVal sets the value in CollapsedVal.
func (p GuiNextWindowData) SetCollapsedVal(value bool) {
	ptr := (*C.ImGuiNextWindowData)(unsafe.Pointer(p))
	ptr.CollapsedVal = (C.bool)(value)
}

// RefContentSizeVal returns pointer to the ContentSizeVal field.
func (p GuiNextWindowData) RefContentSizeVal() Vec2 {
	return Vec2(p + GuiNextWindowData(C.offsetof_ImGuiNextWindowData_ContentSizeVal))
}

// GetDockCond returns the value in DockCond.
func (p GuiNextWindowData) GetDockCond() GuiCond {
	ptr := (*C.ImGuiNextWindowData)(unsafe.Pointer(p))
	return GuiCond(ptr.DockCond)
}

// SetDockCond sets the value in DockCond.
func (p GuiNextWindowData) SetDockCond(value GuiCond) {
	ptr := (*C.ImGuiNextWindowData)(unsafe.Pointer(p))
	ptr.DockCond = (C.ImGuiCond)(value)
}

// GuiNextWindowData.DockId is unsupported: category unsupported.

// GetHasFlags returns the value in HasFlags.
func (p GuiNextWindowData) GetHasFlags() GuiNextWindowDataFlags {
	ptr := (*C.ImGuiNextWindowData)(unsafe.Pointer(p))
	return GuiNextWindowDataFlags(ptr.HasFlags)
}

// SetHasFlags sets the value in HasFlags.
func (p GuiNextWindowData) SetHasFlags(value GuiNextWindowDataFlags) {
	ptr := (*C.ImGuiNextWindowData)(unsafe.Pointer(p))
	ptr.HasFlags = (C.ImGuiNextWindowDataFlags)(value)
}

// RefMenuBarOffsetMinVal returns pointer to the MenuBarOffsetMinVal field.
func (p GuiNextWindowData) RefMenuBarOffsetMinVal() Vec2 {
	return Vec2(p + GuiNextWindowData(C.offsetof_ImGuiNextWindowData_MenuBarOffsetMinVal))
}

// GetPosCond returns the value in PosCond.
func (p GuiNextWindowData) GetPosCond() GuiCond {
	ptr := (*C.ImGuiNextWindowData)(unsafe.Pointer(p))
	return GuiCond(ptr.PosCond)
}

// SetPosCond sets the value in PosCond.
func (p GuiNextWindowData) SetPosCond(value GuiCond) {
	ptr := (*C.ImGuiNextWindowData)(unsafe.Pointer(p))
	ptr.PosCond = (C.ImGuiCond)(value)
}

// RefPosPivotVal returns pointer to the PosPivotVal field.
func (p GuiNextWindowData) RefPosPivotVal() Vec2 {
	return Vec2(p + GuiNextWindowData(C.offsetof_ImGuiNextWindowData_PosPivotVal))
}

// GetPosUndock returns the value in PosUndock.
func (p GuiNextWindowData) GetPosUndock() bool {
	ptr := (*C.ImGuiNextWindowData)(unsafe.Pointer(p))
	return bool(ptr.PosUndock)
}

// SetPosUndock sets the value in PosUndock.
func (p GuiNextWindowData) SetPosUndock(value bool) {
	ptr := (*C.ImGuiNextWindowData)(unsafe.Pointer(p))
	ptr.PosUndock = (C.bool)(value)
}

// RefPosVal returns pointer to the PosVal field.
func (p GuiNextWindowData) RefPosVal() Vec2 {
	return Vec2(p + GuiNextWindowData(C.offsetof_ImGuiNextWindowData_PosVal))
}

// GetRefreshFlagsVal returns the value in RefreshFlagsVal.
func (p GuiNextWindowData) GetRefreshFlagsVal() GuiWindowRefreshFlags {
	ptr := (*C.ImGuiNextWindowData)(unsafe.Pointer(p))
	return GuiWindowRefreshFlags(ptr.RefreshFlagsVal)
}

// SetRefreshFlagsVal sets the value in RefreshFlagsVal.
func (p GuiNextWindowData) SetRefreshFlagsVal(value GuiWindowRefreshFlags) {
	ptr := (*C.ImGuiNextWindowData)(unsafe.Pointer(p))
	ptr.RefreshFlagsVal = (C.ImGuiWindowRefreshFlags)(value)
}

// RefScrollVal returns pointer to the ScrollVal field.
func (p GuiNextWindowData) RefScrollVal() Vec2 {
	return Vec2(p + GuiNextWindowData(C.offsetof_ImGuiNextWindowData_ScrollVal))
}

// GuiNextWindowData.SizeCallback is unsupported: category unsupported.

// GetSizeCallbackUserData returns the value in SizeCallbackUserData.
func (p GuiNextWindowData) GetSizeCallbackUserData() uintptr {
	ptr := (*C.ImGuiNextWindowData)(unsafe.Pointer(p))
	return uintptr(unsafe.Pointer(ptr.SizeCallbackUserData))
}

// SetSizeCallbackUserData sets the value in SizeCallbackUserData.
func (p GuiNextWindowData) SetSizeCallbackUserData(value uintptr) {
	ptr := (*C.ImGuiNextWindowData)(unsafe.Pointer(p))
	ptr.SizeCallbackUserData = unsafe.Pointer(value)
}

// GetSizeCond returns the value in SizeCond.
func (p GuiNextWindowData) GetSizeCond() GuiCond {
	ptr := (*C.ImGuiNextWindowData)(unsafe.Pointer(p))
	return GuiCond(ptr.SizeCond)
}

// SetSizeCond sets the value in SizeCond.
func (p GuiNextWindowData) SetSizeCond(value GuiCond) {
	ptr := (*C.ImGuiNextWindowData)(unsafe.Pointer(p))
	ptr.SizeCond = (C.ImGuiCond)(value)
}

// RefSizeConstraintRect returns pointer to the SizeConstraintRect field.
func (p GuiNextWindowData) RefSizeConstraintRect() Rect {
	return Rect(p + GuiNextWindowData(C.offsetof_ImGuiNextWindowData_SizeConstraintRect))
}

// RefSizeVal returns pointer to the SizeVal field.
func (p GuiNextWindowData) RefSizeVal() Vec2 {
	return Vec2(p + GuiNextWindowData(C.offsetof_ImGuiNextWindowData_SizeVal))
}

// GuiNextWindowData.ViewportId is unsupported: category unsupported.

// RefWindowClass returns pointer to the WindowClass field.
func (p GuiNextWindowData) RefWindowClass() GuiWindowClass {
	return GuiWindowClass(p + GuiNextWindowData(C.offsetof_ImGuiNextWindowData_WindowClass))
}

// GetWindowFlags returns the value in WindowFlags.
func (p GuiNextWindowData) GetWindowFlags() GuiWindowFlags {
	ptr := (*C.ImGuiNextWindowData)(unsafe.Pointer(p))
	return GuiWindowFlags(ptr.WindowFlags)
}

// SetWindowFlags sets the value in WindowFlags.
func (p GuiNextWindowData) SetWindowFlags(value GuiWindowFlags) {
	ptr := (*C.ImGuiNextWindowData)(unsafe.Pointer(p))
	ptr.WindowFlags = (C.ImGuiWindowFlags)(value)
}

// GuiOldColumnData wraps struct ImGuiOldColumnData.
type GuiOldColumnData uintptr

// GuiOldColumnDataNil is a null pointer.
var GuiOldColumnDataNil GuiOldColumnData

// GuiOldColumnDataSizeOf is the byte size of ImGuiOldColumnData.
const GuiOldColumnDataSizeOf = int(C.sizeof_ImGuiOldColumnData)

// ImGuiOldColumnData allocator is suppressed.

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p GuiOldColumnData) Offset(offset int) GuiOldColumnData {
	return p + GuiOldColumnData(offset*GuiOldColumnDataSizeOf)
}

// RefClipRect returns pointer to the ClipRect field.
func (p GuiOldColumnData) RefClipRect() Rect {
	return Rect(p + GuiOldColumnData(C.offsetof_ImGuiOldColumnData_ClipRect))
}

// GetFlags returns the value in Flags.
func (p GuiOldColumnData) GetFlags() GuiOldColumnFlags {
	ptr := (*C.ImGuiOldColumnData)(unsafe.Pointer(p))
	return GuiOldColumnFlags(ptr.Flags)
}

// SetFlags sets the value in Flags.
func (p GuiOldColumnData) SetFlags(value GuiOldColumnFlags) {
	ptr := (*C.ImGuiOldColumnData)(unsafe.Pointer(p))
	ptr.Flags = (C.ImGuiOldColumnFlags)(value)
}

// GetOffsetNorm returns the value in OffsetNorm.
func (p GuiOldColumnData) GetOffsetNorm() float32 {
	ptr := (*C.ImGuiOldColumnData)(unsafe.Pointer(p))
	return float32(ptr.OffsetNorm)
}

// SetOffsetNorm sets the value in OffsetNorm.
func (p GuiOldColumnData) SetOffsetNorm(value float32) {
	ptr := (*C.ImGuiOldColumnData)(unsafe.Pointer(p))
	ptr.OffsetNorm = (C.float)(value)
}

// GetOffsetNormBeforeResize returns the value in OffsetNormBeforeResize.
func (p GuiOldColumnData) GetOffsetNormBeforeResize() float32 {
	ptr := (*C.ImGuiOldColumnData)(unsafe.Pointer(p))
	return float32(ptr.OffsetNormBeforeResize)
}

// SetOffsetNormBeforeResize sets the value in OffsetNormBeforeResize.
func (p GuiOldColumnData) SetOffsetNormBeforeResize(value float32) {
	ptr := (*C.ImGuiOldColumnData)(unsafe.Pointer(p))
	ptr.OffsetNormBeforeResize = (C.float)(value)
}

// GuiOldColumns wraps struct ImGuiOldColumns.
type GuiOldColumns uintptr

// GuiOldColumnsNil is a null pointer.
var GuiOldColumnsNil GuiOldColumns

// GuiOldColumnsSizeOf is the byte size of ImGuiOldColumns.
const GuiOldColumnsSizeOf = int(C.sizeof_ImGuiOldColumns)

// ImGuiOldColumns allocator is suppressed.

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p GuiOldColumns) Offset(offset int) GuiOldColumns {
	return p + GuiOldColumns(offset*GuiOldColumnsSizeOf)
}

// RefColumns returns pointer to the Columns field.
func (p GuiOldColumns) RefColumns() Vector_ImGuiOldColumnData {
	return Vector_ImGuiOldColumnData(p + GuiOldColumns(C.offsetof_ImGuiOldColumns_Columns))
}

// GetCount returns the value in Count.
func (p GuiOldColumns) GetCount() int32 {
	ptr := (*C.ImGuiOldColumns)(unsafe.Pointer(p))
	return int32(ptr.Count)
}

// SetCount sets the value in Count.
func (p GuiOldColumns) SetCount(value int32) {
	ptr := (*C.ImGuiOldColumns)(unsafe.Pointer(p))
	ptr.Count = (C.int32_t)(value)
}

// GetCurrent returns the value in Current.
func (p GuiOldColumns) GetCurrent() int32 {
	ptr := (*C.ImGuiOldColumns)(unsafe.Pointer(p))
	return int32(ptr.Current)
}

// SetCurrent sets the value in Current.
func (p GuiOldColumns) SetCurrent(value int32) {
	ptr := (*C.ImGuiOldColumns)(unsafe.Pointer(p))
	ptr.Current = (C.int32_t)(value)
}

// GetFlags returns the value in Flags.
func (p GuiOldColumns) GetFlags() GuiOldColumnFlags {
	ptr := (*C.ImGuiOldColumns)(unsafe.Pointer(p))
	return GuiOldColumnFlags(ptr.Flags)
}

// SetFlags sets the value in Flags.
func (p GuiOldColumns) SetFlags(value GuiOldColumnFlags) {
	ptr := (*C.ImGuiOldColumns)(unsafe.Pointer(p))
	ptr.Flags = (C.ImGuiOldColumnFlags)(value)
}

// RefHostBackupClipRect returns pointer to the HostBackupClipRect field.
func (p GuiOldColumns) RefHostBackupClipRect() Rect {
	return Rect(p + GuiOldColumns(C.offsetof_ImGuiOldColumns_HostBackupClipRect))
}

// RefHostBackupParentWorkRect returns pointer to the HostBackupParentWorkRect field.
func (p GuiOldColumns) RefHostBackupParentWorkRect() Rect {
	return Rect(p + GuiOldColumns(C.offsetof_ImGuiOldColumns_HostBackupParentWorkRect))
}

// GetHostCursorMaxPosX returns the value in HostCursorMaxPosX.
func (p GuiOldColumns) GetHostCursorMaxPosX() float32 {
	ptr := (*C.ImGuiOldColumns)(unsafe.Pointer(p))
	return float32(ptr.HostCursorMaxPosX)
}

// SetHostCursorMaxPosX sets the value in HostCursorMaxPosX.
func (p GuiOldColumns) SetHostCursorMaxPosX(value float32) {
	ptr := (*C.ImGuiOldColumns)(unsafe.Pointer(p))
	ptr.HostCursorMaxPosX = (C.float)(value)
}

// GetHostCursorPosY returns the value in HostCursorPosY.
func (p GuiOldColumns) GetHostCursorPosY() float32 {
	ptr := (*C.ImGuiOldColumns)(unsafe.Pointer(p))
	return float32(ptr.HostCursorPosY)
}

// SetHostCursorPosY sets the value in HostCursorPosY.
func (p GuiOldColumns) SetHostCursorPosY(value float32) {
	ptr := (*C.ImGuiOldColumns)(unsafe.Pointer(p))
	ptr.HostCursorPosY = (C.float)(value)
}

// RefHostInitialClipRect returns pointer to the HostInitialClipRect field.
func (p GuiOldColumns) RefHostInitialClipRect() Rect {
	return Rect(p + GuiOldColumns(C.offsetof_ImGuiOldColumns_HostInitialClipRect))
}

// GuiOldColumns.ID is unsupported: category unsupported.

// GetIsBeingResized returns the value in IsBeingResized.
func (p GuiOldColumns) GetIsBeingResized() bool {
	ptr := (*C.ImGuiOldColumns)(unsafe.Pointer(p))
	return bool(ptr.IsBeingResized)
}

// SetIsBeingResized sets the value in IsBeingResized.
func (p GuiOldColumns) SetIsBeingResized(value bool) {
	ptr := (*C.ImGuiOldColumns)(unsafe.Pointer(p))
	ptr.IsBeingResized = (C.bool)(value)
}

// GetIsFirstFrame returns the value in IsFirstFrame.
func (p GuiOldColumns) GetIsFirstFrame() bool {
	ptr := (*C.ImGuiOldColumns)(unsafe.Pointer(p))
	return bool(ptr.IsFirstFrame)
}

// SetIsFirstFrame sets the value in IsFirstFrame.
func (p GuiOldColumns) SetIsFirstFrame(value bool) {
	ptr := (*C.ImGuiOldColumns)(unsafe.Pointer(p))
	ptr.IsFirstFrame = (C.bool)(value)
}

// GetLineMaxY returns the value in LineMaxY.
func (p GuiOldColumns) GetLineMaxY() float32 {
	ptr := (*C.ImGuiOldColumns)(unsafe.Pointer(p))
	return float32(ptr.LineMaxY)
}

// SetLineMaxY sets the value in LineMaxY.
func (p GuiOldColumns) SetLineMaxY(value float32) {
	ptr := (*C.ImGuiOldColumns)(unsafe.Pointer(p))
	ptr.LineMaxY = (C.float)(value)
}

// GetLineMinY returns the value in LineMinY.
func (p GuiOldColumns) GetLineMinY() float32 {
	ptr := (*C.ImGuiOldColumns)(unsafe.Pointer(p))
	return float32(ptr.LineMinY)
}

// SetLineMinY sets the value in LineMinY.
func (p GuiOldColumns) SetLineMinY(value float32) {
	ptr := (*C.ImGuiOldColumns)(unsafe.Pointer(p))
	ptr.LineMinY = (C.float)(value)
}

// GetOffMaxX returns the value in OffMaxX.
func (p GuiOldColumns) GetOffMaxX() float32 {
	ptr := (*C.ImGuiOldColumns)(unsafe.Pointer(p))
	return float32(ptr.OffMaxX)
}

// SetOffMaxX sets the value in OffMaxX.
func (p GuiOldColumns) SetOffMaxX(value float32) {
	ptr := (*C.ImGuiOldColumns)(unsafe.Pointer(p))
	ptr.OffMaxX = (C.float)(value)
}

// GetOffMinX returns the value in OffMinX.
func (p GuiOldColumns) GetOffMinX() float32 {
	ptr := (*C.ImGuiOldColumns)(unsafe.Pointer(p))
	return float32(ptr.OffMinX)
}

// SetOffMinX sets the value in OffMinX.
func (p GuiOldColumns) SetOffMinX(value float32) {
	ptr := (*C.ImGuiOldColumns)(unsafe.Pointer(p))
	ptr.OffMinX = (C.float)(value)
}

// RefSplitter returns pointer to the Splitter field.
func (p GuiOldColumns) RefSplitter() DrawListSplitter {
	return DrawListSplitter(p + GuiOldColumns(C.offsetof_ImGuiOldColumns_Splitter))
}

// GuiOnceUponAFrame wraps struct ImGuiOnceUponAFrame.
type GuiOnceUponAFrame uintptr

// GuiOnceUponAFrameNil is a null pointer.
var GuiOnceUponAFrameNil GuiOnceUponAFrame

// GuiOnceUponAFrameSizeOf is the byte size of ImGuiOnceUponAFrame.
const GuiOnceUponAFrameSizeOf = int(C.sizeof_ImGuiOnceUponAFrame)

// ImGuiOnceUponAFrame allocator is suppressed.

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p GuiOnceUponAFrame) Offset(offset int) GuiOnceUponAFrame {
	return p + GuiOnceUponAFrame(offset*GuiOnceUponAFrameSizeOf)
}

// GetRefFrame returns the value in RefFrame.
func (p GuiOnceUponAFrame) GetRefFrame() int32 {
	ptr := (*C.ImGuiOnceUponAFrame)(unsafe.Pointer(p))
	return int32(ptr.RefFrame)
}

// SetRefFrame sets the value in RefFrame.
func (p GuiOnceUponAFrame) SetRefFrame(value int32) {
	ptr := (*C.ImGuiOnceUponAFrame)(unsafe.Pointer(p))
	ptr.RefFrame = (C.int32_t)(value)
}

// GuiPayload wraps struct ImGuiPayload.
type GuiPayload uintptr

// GuiPayloadNil is a null pointer.
var GuiPayloadNil GuiPayload

// GuiPayloadSizeOf is the byte size of ImGuiPayload.
const GuiPayloadSizeOf = int(C.sizeof_ImGuiPayload)

// ImGuiPayload allocator is suppressed.

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p GuiPayload) Offset(offset int) GuiPayload {
	return p + GuiPayload(offset*GuiPayloadSizeOf)
}

// GetData returns the value in Data.
func (p GuiPayload) GetData() uintptr {
	ptr := (*C.ImGuiPayload)(unsafe.Pointer(p))
	return uintptr(unsafe.Pointer(ptr.Data))
}

// SetData sets the value in Data.
func (p GuiPayload) SetData(value uintptr) {
	ptr := (*C.ImGuiPayload)(unsafe.Pointer(p))
	ptr.Data = unsafe.Pointer(value)
}

// GetDataFrameCount returns the value in DataFrameCount.
func (p GuiPayload) GetDataFrameCount() int32 {
	ptr := (*C.ImGuiPayload)(unsafe.Pointer(p))
	return int32(ptr.DataFrameCount)
}

// SetDataFrameCount sets the value in DataFrameCount.
func (p GuiPayload) SetDataFrameCount(value int32) {
	ptr := (*C.ImGuiPayload)(unsafe.Pointer(p))
	ptr.DataFrameCount = (C.int32_t)(value)
}

// GetDataSize returns the value in DataSize.
func (p GuiPayload) GetDataSize() int32 {
	ptr := (*C.ImGuiPayload)(unsafe.Pointer(p))
	return int32(ptr.DataSize)
}

// SetDataSize sets the value in DataSize.
func (p GuiPayload) SetDataSize(value int32) {
	ptr := (*C.ImGuiPayload)(unsafe.Pointer(p))
	ptr.DataSize = (C.int32_t)(value)
}

// GuiPayload.DataType[32+1] is unsupported: category unsupported.

// GetDelivery returns the value in Delivery.
func (p GuiPayload) GetDelivery() bool {
	ptr := (*C.ImGuiPayload)(unsafe.Pointer(p))
	return bool(ptr.Delivery)
}

// SetDelivery sets the value in Delivery.
func (p GuiPayload) SetDelivery(value bool) {
	ptr := (*C.ImGuiPayload)(unsafe.Pointer(p))
	ptr.Delivery = (C.bool)(value)
}

// GetPreview returns the value in Preview.
func (p GuiPayload) GetPreview() bool {
	ptr := (*C.ImGuiPayload)(unsafe.Pointer(p))
	return bool(ptr.Preview)
}

// SetPreview sets the value in Preview.
func (p GuiPayload) SetPreview(value bool) {
	ptr := (*C.ImGuiPayload)(unsafe.Pointer(p))
	ptr.Preview = (C.bool)(value)
}

// GuiPayload.SourceId is unsupported: category unsupported.

// GuiPayload.SourceParentId is unsupported: category unsupported.

// GuiPlatformIO wraps struct ImGuiPlatformIO.
type GuiPlatformIO uintptr

// GuiPlatformIONil is a null pointer.
var GuiPlatformIONil GuiPlatformIO

// GuiPlatformIOSizeOf is the byte size of ImGuiPlatformIO.
const GuiPlatformIOSizeOf = int(C.sizeof_ImGuiPlatformIO)

// ImGuiPlatformIO allocator is suppressed.

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p GuiPlatformIO) Offset(offset int) GuiPlatformIO {
	return p + GuiPlatformIO(offset*GuiPlatformIOSizeOf)
}

// RefMonitors returns pointer to the Monitors field.
func (p GuiPlatformIO) RefMonitors() Vector_ImGuiPlatformMonitor {
	return Vector_ImGuiPlatformMonitor(p + GuiPlatformIO(C.offsetof_ImGuiPlatformIO_Monitors))
}

// GetPlatform_ClipboardUserData returns the value in Platform_ClipboardUserData.
func (p GuiPlatformIO) GetPlatform_ClipboardUserData() uintptr {
	ptr := (*C.ImGuiPlatformIO)(unsafe.Pointer(p))
	return uintptr(unsafe.Pointer(ptr.Platform_ClipboardUserData))
}

// SetPlatform_ClipboardUserData sets the value in Platform_ClipboardUserData.
func (p GuiPlatformIO) SetPlatform_ClipboardUserData(value uintptr) {
	ptr := (*C.ImGuiPlatformIO)(unsafe.Pointer(p))
	ptr.Platform_ClipboardUserData = unsafe.Pointer(value)
}

// GuiPlatformIO.Platform_CreateVkSurface is unsupported: category unsupported.

// GuiPlatformIO.Platform_CreateWindow is unsupported: category unsupported.

// GuiPlatformIO.Platform_DestroyWindow is unsupported: category unsupported.

// GuiPlatformIO.Platform_GetClipboardTextFn is unsupported: category unsupported.

// GuiPlatformIO.Platform_GetWindowDpiScale is unsupported: category unsupported.

// GuiPlatformIO.Platform_GetWindowFocus is unsupported: category unsupported.

// GuiPlatformIO.Platform_GetWindowFramebufferScale is unsupported: category unsupported.

// GuiPlatformIO.Platform_GetWindowMinimized is unsupported: category unsupported.

// GuiPlatformIO.Platform_GetWindowPos is unsupported: category unsupported.

// GuiPlatformIO.Platform_GetWindowSize is unsupported: category unsupported.

// GuiPlatformIO.Platform_GetWindowWorkAreaInsets is unsupported: category unsupported.

// GetPlatform_ImeUserData returns the value in Platform_ImeUserData.
func (p GuiPlatformIO) GetPlatform_ImeUserData() uintptr {
	ptr := (*C.ImGuiPlatformIO)(unsafe.Pointer(p))
	return uintptr(unsafe.Pointer(ptr.Platform_ImeUserData))
}

// SetPlatform_ImeUserData sets the value in Platform_ImeUserData.
func (p GuiPlatformIO) SetPlatform_ImeUserData(value uintptr) {
	ptr := (*C.ImGuiPlatformIO)(unsafe.Pointer(p))
	ptr.Platform_ImeUserData = unsafe.Pointer(value)
}

// GuiPlatformIO.Platform_LocaleDecimalPoint is unsupported: category unsupported.

// GuiPlatformIO.Platform_OnChangedViewport is unsupported: category unsupported.

// GuiPlatformIO.Platform_OpenInShellFn is unsupported: category unsupported.

// GetPlatform_OpenInShellUserData returns the value in Platform_OpenInShellUserData.
func (p GuiPlatformIO) GetPlatform_OpenInShellUserData() uintptr {
	ptr := (*C.ImGuiPlatformIO)(unsafe.Pointer(p))
	return uintptr(unsafe.Pointer(ptr.Platform_OpenInShellUserData))
}

// SetPlatform_OpenInShellUserData sets the value in Platform_OpenInShellUserData.
func (p GuiPlatformIO) SetPlatform_OpenInShellUserData(value uintptr) {
	ptr := (*C.ImGuiPlatformIO)(unsafe.Pointer(p))
	ptr.Platform_OpenInShellUserData = unsafe.Pointer(value)
}

// GuiPlatformIO.Platform_RenderWindow is unsupported: category unsupported.

// GuiPlatformIO.Platform_SetClipboardTextFn is unsupported: category unsupported.

// GuiPlatformIO.Platform_SetImeDataFn is unsupported: category unsupported.

// GuiPlatformIO.Platform_SetWindowAlpha is unsupported: category unsupported.

// GuiPlatformIO.Platform_SetWindowFocus is unsupported: category unsupported.

// GuiPlatformIO.Platform_SetWindowPos is unsupported: category unsupported.

// GuiPlatformIO.Platform_SetWindowSize is unsupported: category unsupported.

// GuiPlatformIO.Platform_SetWindowTitle is unsupported: category unsupported.

// GuiPlatformIO.Platform_ShowWindow is unsupported: category unsupported.

// GuiPlatformIO.Platform_SwapBuffers is unsupported: category unsupported.

// GuiPlatformIO.Platform_UpdateWindow is unsupported: category unsupported.

// GuiPlatformIO.Renderer_CreateWindow is unsupported: category unsupported.

// GuiPlatformIO.Renderer_DestroyWindow is unsupported: category unsupported.

// GetRenderer_RenderState returns the value in Renderer_RenderState.
func (p GuiPlatformIO) GetRenderer_RenderState() uintptr {
	ptr := (*C.ImGuiPlatformIO)(unsafe.Pointer(p))
	return uintptr(unsafe.Pointer(ptr.Renderer_RenderState))
}

// SetRenderer_RenderState sets the value in Renderer_RenderState.
func (p GuiPlatformIO) SetRenderer_RenderState(value uintptr) {
	ptr := (*C.ImGuiPlatformIO)(unsafe.Pointer(p))
	ptr.Renderer_RenderState = unsafe.Pointer(value)
}

// GuiPlatformIO.Renderer_RenderWindow is unsupported: category unsupported.

// GuiPlatformIO.Renderer_SetWindowSize is unsupported: category unsupported.

// GuiPlatformIO.Renderer_SwapBuffers is unsupported: category unsupported.

// GetRenderer_TextureMaxHeight returns the value in Renderer_TextureMaxHeight.
func (p GuiPlatformIO) GetRenderer_TextureMaxHeight() int32 {
	ptr := (*C.ImGuiPlatformIO)(unsafe.Pointer(p))
	return int32(ptr.Renderer_TextureMaxHeight)
}

// SetRenderer_TextureMaxHeight sets the value in Renderer_TextureMaxHeight.
func (p GuiPlatformIO) SetRenderer_TextureMaxHeight(value int32) {
	ptr := (*C.ImGuiPlatformIO)(unsafe.Pointer(p))
	ptr.Renderer_TextureMaxHeight = (C.int32_t)(value)
}

// GetRenderer_TextureMaxWidth returns the value in Renderer_TextureMaxWidth.
func (p GuiPlatformIO) GetRenderer_TextureMaxWidth() int32 {
	ptr := (*C.ImGuiPlatformIO)(unsafe.Pointer(p))
	return int32(ptr.Renderer_TextureMaxWidth)
}

// SetRenderer_TextureMaxWidth sets the value in Renderer_TextureMaxWidth.
func (p GuiPlatformIO) SetRenderer_TextureMaxWidth(value int32) {
	ptr := (*C.ImGuiPlatformIO)(unsafe.Pointer(p))
	ptr.Renderer_TextureMaxWidth = (C.int32_t)(value)
}

// RefTextures returns pointer to the Textures field.
func (p GuiPlatformIO) RefTextures() Vector_ImTextureDataPtr {
	return Vector_ImTextureDataPtr(p + GuiPlatformIO(C.offsetof_ImGuiPlatformIO_Textures))
}

// RefViewports returns pointer to the Viewports field.
func (p GuiPlatformIO) RefViewports() Vector_ImGuiViewportPtr {
	return Vector_ImGuiViewportPtr(p + GuiPlatformIO(C.offsetof_ImGuiPlatformIO_Viewports))
}

// GuiPlatformImeData wraps struct ImGuiPlatformImeData.
type GuiPlatformImeData uintptr

// GuiPlatformImeDataNil is a null pointer.
var GuiPlatformImeDataNil GuiPlatformImeData

// GuiPlatformImeDataSizeOf is the byte size of ImGuiPlatformImeData.
const GuiPlatformImeDataSizeOf = int(C.sizeof_ImGuiPlatformImeData)

// ImGuiPlatformImeData allocator is suppressed.

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p GuiPlatformImeData) Offset(offset int) GuiPlatformImeData {
	return p + GuiPlatformImeData(offset*GuiPlatformImeDataSizeOf)
}

// GetInputLineHeight returns the value in InputLineHeight.
func (p GuiPlatformImeData) GetInputLineHeight() float32 {
	ptr := (*C.ImGuiPlatformImeData)(unsafe.Pointer(p))
	return float32(ptr.InputLineHeight)
}

// SetInputLineHeight sets the value in InputLineHeight.
func (p GuiPlatformImeData) SetInputLineHeight(value float32) {
	ptr := (*C.ImGuiPlatformImeData)(unsafe.Pointer(p))
	ptr.InputLineHeight = (C.float)(value)
}

// RefInputPos returns pointer to the InputPos field.
func (p GuiPlatformImeData) RefInputPos() Vec2 {
	return Vec2(p + GuiPlatformImeData(C.offsetof_ImGuiPlatformImeData_InputPos))
}

// GuiPlatformImeData.ViewportId is unsupported: category unsupported.

// GetWantTextInput returns the value in WantTextInput.
func (p GuiPlatformImeData) GetWantTextInput() bool {
	ptr := (*C.ImGuiPlatformImeData)(unsafe.Pointer(p))
	return bool(ptr.WantTextInput)
}

// SetWantTextInput sets the value in WantTextInput.
func (p GuiPlatformImeData) SetWantTextInput(value bool) {
	ptr := (*C.ImGuiPlatformImeData)(unsafe.Pointer(p))
	ptr.WantTextInput = (C.bool)(value)
}

// GetWantVisible returns the value in WantVisible.
func (p GuiPlatformImeData) GetWantVisible() bool {
	ptr := (*C.ImGuiPlatformImeData)(unsafe.Pointer(p))
	return bool(ptr.WantVisible)
}

// SetWantVisible sets the value in WantVisible.
func (p GuiPlatformImeData) SetWantVisible(value bool) {
	ptr := (*C.ImGuiPlatformImeData)(unsafe.Pointer(p))
	ptr.WantVisible = (C.bool)(value)
}

// GuiPlatformMonitor wraps struct ImGuiPlatformMonitor.
type GuiPlatformMonitor uintptr

// GuiPlatformMonitorNil is a null pointer.
var GuiPlatformMonitorNil GuiPlatformMonitor

// GuiPlatformMonitorSizeOf is the byte size of ImGuiPlatformMonitor.
const GuiPlatformMonitorSizeOf = int(C.sizeof_ImGuiPlatformMonitor)

// ImGuiPlatformMonitor allocator is suppressed.

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p GuiPlatformMonitor) Offset(offset int) GuiPlatformMonitor {
	return p + GuiPlatformMonitor(offset*GuiPlatformMonitorSizeOf)
}

// GetDpiScale returns the value in DpiScale.
func (p GuiPlatformMonitor) GetDpiScale() float32 {
	ptr := (*C.ImGuiPlatformMonitor)(unsafe.Pointer(p))
	return float32(ptr.DpiScale)
}

// SetDpiScale sets the value in DpiScale.
func (p GuiPlatformMonitor) SetDpiScale(value float32) {
	ptr := (*C.ImGuiPlatformMonitor)(unsafe.Pointer(p))
	ptr.DpiScale = (C.float)(value)
}

// RefMainPos returns pointer to the MainPos field.
func (p GuiPlatformMonitor) RefMainPos() Vec2 {
	return Vec2(p + GuiPlatformMonitor(C.offsetof_ImGuiPlatformMonitor_MainPos))
}

// RefMainSize returns pointer to the MainSize field.
func (p GuiPlatformMonitor) RefMainSize() Vec2 {
	return Vec2(p + GuiPlatformMonitor(C.offsetof_ImGuiPlatformMonitor_MainSize))
}

// GetPlatformHandle returns the value in PlatformHandle.
func (p GuiPlatformMonitor) GetPlatformHandle() uintptr {
	ptr := (*C.ImGuiPlatformMonitor)(unsafe.Pointer(p))
	return uintptr(unsafe.Pointer(ptr.PlatformHandle))
}

// SetPlatformHandle sets the value in PlatformHandle.
func (p GuiPlatformMonitor) SetPlatformHandle(value uintptr) {
	ptr := (*C.ImGuiPlatformMonitor)(unsafe.Pointer(p))
	ptr.PlatformHandle = unsafe.Pointer(value)
}

// RefWorkPos returns pointer to the WorkPos field.
func (p GuiPlatformMonitor) RefWorkPos() Vec2 {
	return Vec2(p + GuiPlatformMonitor(C.offsetof_ImGuiPlatformMonitor_WorkPos))
}

// RefWorkSize returns pointer to the WorkSize field.
func (p GuiPlatformMonitor) RefWorkSize() Vec2 {
	return Vec2(p + GuiPlatformMonitor(C.offsetof_ImGuiPlatformMonitor_WorkSize))
}

// GuiPopupData wraps struct ImGuiPopupData.
type GuiPopupData uintptr

// GuiPopupDataNil is a null pointer.
var GuiPopupDataNil GuiPopupData

// GuiPopupDataSizeOf is the byte size of ImGuiPopupData.
const GuiPopupDataSizeOf = int(C.sizeof_ImGuiPopupData)

// ImGuiPopupData allocator is suppressed.

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p GuiPopupData) Offset(offset int) GuiPopupData {
	return p + GuiPopupData(offset*GuiPopupDataSizeOf)
}

// GetOpenFrameCount returns the value in OpenFrameCount.
func (p GuiPopupData) GetOpenFrameCount() int32 {
	ptr := (*C.ImGuiPopupData)(unsafe.Pointer(p))
	return int32(ptr.OpenFrameCount)
}

// SetOpenFrameCount sets the value in OpenFrameCount.
func (p GuiPopupData) SetOpenFrameCount(value int32) {
	ptr := (*C.ImGuiPopupData)(unsafe.Pointer(p))
	ptr.OpenFrameCount = (C.int32_t)(value)
}

// RefOpenMousePos returns pointer to the OpenMousePos field.
func (p GuiPopupData) RefOpenMousePos() Vec2 {
	return Vec2(p + GuiPopupData(C.offsetof_ImGuiPopupData_OpenMousePos))
}

// GuiPopupData.OpenParentId is unsupported: category unsupported.

// RefOpenPopupPos returns pointer to the OpenPopupPos field.
func (p GuiPopupData) RefOpenPopupPos() Vec2 {
	return Vec2(p + GuiPopupData(C.offsetof_ImGuiPopupData_OpenPopupPos))
}

// GetParentNavLayer returns the value in ParentNavLayer.
func (p GuiPopupData) GetParentNavLayer() int32 {
	ptr := (*C.ImGuiPopupData)(unsafe.Pointer(p))
	return int32(ptr.ParentNavLayer)
}

// SetParentNavLayer sets the value in ParentNavLayer.
func (p GuiPopupData) SetParentNavLayer(value int32) {
	ptr := (*C.ImGuiPopupData)(unsafe.Pointer(p))
	ptr.ParentNavLayer = (C.int32_t)(value)
}

// GuiPopupData.PopupId is unsupported: category unsupported.

// GetRestoreNavWindow returns the value in RestoreNavWindow.
func (p GuiPopupData) GetRestoreNavWindow() GuiWindow {
	ptr := (*C.ImGuiPopupData)(unsafe.Pointer(p))
	return GuiWindow(unsafe.Pointer(ptr.RestoreNavWindow))
}

// SetRestoreNavWindow sets the value in RestoreNavWindow.
func (p GuiPopupData) SetRestoreNavWindow(value GuiWindow) {
	ptr := (*C.ImGuiPopupData)(unsafe.Pointer(p))
	ptr.RestoreNavWindow = (*C.ImGuiWindow)(unsafe.Pointer(value))
}

// GetWindow returns the value in Window.
func (p GuiPopupData) GetWindow() GuiWindow {
	ptr := (*C.ImGuiPopupData)(unsafe.Pointer(p))
	return GuiWindow(unsafe.Pointer(ptr.Window))
}

// SetWindow sets the value in Window.
func (p GuiPopupData) SetWindow(value GuiWindow) {
	ptr := (*C.ImGuiPopupData)(unsafe.Pointer(p))
	ptr.Window = (*C.ImGuiWindow)(unsafe.Pointer(value))
}

// GuiPtrOrIndex wraps struct ImGuiPtrOrIndex.
type GuiPtrOrIndex uintptr

// GuiPtrOrIndexNil is a null pointer.
var GuiPtrOrIndexNil GuiPtrOrIndex

// GuiPtrOrIndexSizeOf is the byte size of ImGuiPtrOrIndex.
const GuiPtrOrIndexSizeOf = int(C.sizeof_ImGuiPtrOrIndex)

// ImGuiPtrOrIndex allocator is suppressed.

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p GuiPtrOrIndex) Offset(offset int) GuiPtrOrIndex {
	return p + GuiPtrOrIndex(offset*GuiPtrOrIndexSizeOf)
}

// GetIndex returns the value in Index.
func (p GuiPtrOrIndex) GetIndex() int32 {
	ptr := (*C.ImGuiPtrOrIndex)(unsafe.Pointer(p))
	return int32(ptr.Index)
}

// SetIndex sets the value in Index.
func (p GuiPtrOrIndex) SetIndex(value int32) {
	ptr := (*C.ImGuiPtrOrIndex)(unsafe.Pointer(p))
	ptr.Index = (C.int32_t)(value)
}

// GetPtr returns the value in Ptr.
func (p GuiPtrOrIndex) GetPtr() uintptr {
	ptr := (*C.ImGuiPtrOrIndex)(unsafe.Pointer(p))
	return uintptr(unsafe.Pointer(ptr.Ptr))
}

// SetPtr sets the value in Ptr.
func (p GuiPtrOrIndex) SetPtr(value uintptr) {
	ptr := (*C.ImGuiPtrOrIndex)(unsafe.Pointer(p))
	ptr.Ptr = unsafe.Pointer(value)
}

// GuiSelectionBasicStorage wraps struct ImGuiSelectionBasicStorage.
type GuiSelectionBasicStorage uintptr

// GuiSelectionBasicStorageNil is a null pointer.
var GuiSelectionBasicStorageNil GuiSelectionBasicStorage

// GuiSelectionBasicStorageSizeOf is the byte size of ImGuiSelectionBasicStorage.
const GuiSelectionBasicStorageSizeOf = int(C.sizeof_ImGuiSelectionBasicStorage)

// ImGuiSelectionBasicStorage allocator is suppressed.

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p GuiSelectionBasicStorage) Offset(offset int) GuiSelectionBasicStorage {
	return p + GuiSelectionBasicStorage(offset*GuiSelectionBasicStorageSizeOf)
}

// GuiSelectionBasicStorage.AdapterIndexToStorageId is unsupported: category unsupported.

// GetPreserveOrder returns the value in PreserveOrder.
func (p GuiSelectionBasicStorage) GetPreserveOrder() bool {
	ptr := (*C.ImGuiSelectionBasicStorage)(unsafe.Pointer(p))
	return bool(ptr.PreserveOrder)
}

// SetPreserveOrder sets the value in PreserveOrder.
func (p GuiSelectionBasicStorage) SetPreserveOrder(value bool) {
	ptr := (*C.ImGuiSelectionBasicStorage)(unsafe.Pointer(p))
	ptr.PreserveOrder = (C.bool)(value)
}

// GetSize returns the value in Size.
func (p GuiSelectionBasicStorage) GetSize() int32 {
	ptr := (*C.ImGuiSelectionBasicStorage)(unsafe.Pointer(p))
	return int32(ptr.Size)
}

// SetSize sets the value in Size.
func (p GuiSelectionBasicStorage) SetSize(value int32) {
	ptr := (*C.ImGuiSelectionBasicStorage)(unsafe.Pointer(p))
	ptr.Size = (C.int32_t)(value)
}

// GetUserData returns the value in UserData.
func (p GuiSelectionBasicStorage) GetUserData() uintptr {
	ptr := (*C.ImGuiSelectionBasicStorage)(unsafe.Pointer(p))
	return uintptr(unsafe.Pointer(ptr.UserData))
}

// SetUserData sets the value in UserData.
func (p GuiSelectionBasicStorage) SetUserData(value uintptr) {
	ptr := (*C.ImGuiSelectionBasicStorage)(unsafe.Pointer(p))
	ptr.UserData = unsafe.Pointer(value)
}

// GetSelectionOrder returns the value in _SelectionOrder.
func (p GuiSelectionBasicStorage) GetSelectionOrder() int32 {
	ptr := (*C.ImGuiSelectionBasicStorage)(unsafe.Pointer(p))
	return int32(ptr._SelectionOrder)
}

// SetSelectionOrder sets the value in _SelectionOrder.
func (p GuiSelectionBasicStorage) SetSelectionOrder(value int32) {
	ptr := (*C.ImGuiSelectionBasicStorage)(unsafe.Pointer(p))
	ptr._SelectionOrder = (C.int32_t)(value)
}

// RefStorage returns pointer to the _Storage field.
func (p GuiSelectionBasicStorage) RefStorage() GuiStorage {
	return GuiStorage(p + GuiSelectionBasicStorage(C.offsetof_ImGuiSelectionBasicStorage__Storage))
}

// GuiSelectionExternalStorage wraps struct ImGuiSelectionExternalStorage.
type GuiSelectionExternalStorage uintptr

// GuiSelectionExternalStorageNil is a null pointer.
var GuiSelectionExternalStorageNil GuiSelectionExternalStorage

// GuiSelectionExternalStorageSizeOf is the byte size of ImGuiSelectionExternalStorage.
const GuiSelectionExternalStorageSizeOf = int(C.sizeof_ImGuiSelectionExternalStorage)

// ImGuiSelectionExternalStorage allocator is suppressed.

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p GuiSelectionExternalStorage) Offset(offset int) GuiSelectionExternalStorage {
	return p + GuiSelectionExternalStorage(offset*GuiSelectionExternalStorageSizeOf)
}

// GuiSelectionExternalStorage.AdapterSetItemSelected is unsupported: category unsupported.

// GetUserData returns the value in UserData.
func (p GuiSelectionExternalStorage) GetUserData() uintptr {
	ptr := (*C.ImGuiSelectionExternalStorage)(unsafe.Pointer(p))
	return uintptr(unsafe.Pointer(ptr.UserData))
}

// SetUserData sets the value in UserData.
func (p GuiSelectionExternalStorage) SetUserData(value uintptr) {
	ptr := (*C.ImGuiSelectionExternalStorage)(unsafe.Pointer(p))
	ptr.UserData = unsafe.Pointer(value)
}

// GuiSelectionRequest wraps struct ImGuiSelectionRequest.
type GuiSelectionRequest uintptr

// GuiSelectionRequestNil is a null pointer.
var GuiSelectionRequestNil GuiSelectionRequest

// GuiSelectionRequestSizeOf is the byte size of ImGuiSelectionRequest.
const GuiSelectionRequestSizeOf = int(C.sizeof_ImGuiSelectionRequest)

// GuiSelectionRequestAlloc allocates a continuous block of GuiSelectionRequest.
func GuiSelectionRequestAlloc(alloc ffi.Allocator, count int) GuiSelectionRequest {
	ptr := alloc.Allocate(GuiSelectionRequestSizeOf * count)
	return GuiSelectionRequest(ptr)
}

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p GuiSelectionRequest) Offset(offset int) GuiSelectionRequest {
	return p + GuiSelectionRequest(offset*GuiSelectionRequestSizeOf)
}

// GuiSelectionRequest.RangeDirection is unsupported: category unsupported.

// GuiSelectionRequest.RangeFirstItem is unsupported: category unsupported.

// GuiSelectionRequest.RangeLastItem is unsupported: category unsupported.

// GetSelected returns the value in Selected.
func (p GuiSelectionRequest) GetSelected() bool {
	ptr := (*C.ImGuiSelectionRequest)(unsafe.Pointer(p))
	return bool(ptr.Selected)
}

// SetSelected sets the value in Selected.
func (p GuiSelectionRequest) SetSelected(value bool) {
	ptr := (*C.ImGuiSelectionRequest)(unsafe.Pointer(p))
	ptr.Selected = (C.bool)(value)
}

// GetType returns the value in Type.
func (p GuiSelectionRequest) GetType() GuiSelectionRequestType {
	ptr := (*C.ImGuiSelectionRequest)(unsafe.Pointer(p))
	return GuiSelectionRequestType(ptr.Type)
}

// SetType sets the value in Type.
func (p GuiSelectionRequest) SetType(value GuiSelectionRequestType) {
	ptr := (*C.ImGuiSelectionRequest)(unsafe.Pointer(p))
	ptr.Type = (C.ImGuiSelectionRequestType)(value)
}

// GuiSettingsHandler wraps struct ImGuiSettingsHandler.
type GuiSettingsHandler uintptr

// GuiSettingsHandlerNil is a null pointer.
var GuiSettingsHandlerNil GuiSettingsHandler

// GuiSettingsHandlerSizeOf is the byte size of ImGuiSettingsHandler.
const GuiSettingsHandlerSizeOf = int(C.sizeof_ImGuiSettingsHandler)

// ImGuiSettingsHandler allocator is suppressed.

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p GuiSettingsHandler) Offset(offset int) GuiSettingsHandler {
	return p + GuiSettingsHandler(offset*GuiSettingsHandlerSizeOf)
}

// GuiSettingsHandler.ApplyAllFn is unsupported: category unsupported.

// GuiSettingsHandler.ClearAllFn is unsupported: category unsupported.

// GuiSettingsHandler.ReadInitFn is unsupported: category unsupported.

// GuiSettingsHandler.ReadLineFn is unsupported: category unsupported.

// GuiSettingsHandler.ReadOpenFn is unsupported: category unsupported.

// GuiSettingsHandler.TypeHash is unsupported: category unsupported.

// GetTypeName returns the value in TypeName.
func (p GuiSettingsHandler) GetTypeName() ffi.CString {
	ptr := (*C.ImGuiSettingsHandler)(unsafe.Pointer(p))
	return ffi.CStringFromPtr((unsafe.Pointer)(ptr.TypeName))
}

// SetTypeName sets the value in TypeName.
func (p GuiSettingsHandler) SetTypeName(value ffi.CString) {
	ptr := (*C.ImGuiSettingsHandler)(unsafe.Pointer(p))
	ptr.TypeName = (*C.char)(value.Raw())
}

// GetUserData returns the value in UserData.
func (p GuiSettingsHandler) GetUserData() uintptr {
	ptr := (*C.ImGuiSettingsHandler)(unsafe.Pointer(p))
	return uintptr(unsafe.Pointer(ptr.UserData))
}

// SetUserData sets the value in UserData.
func (p GuiSettingsHandler) SetUserData(value uintptr) {
	ptr := (*C.ImGuiSettingsHandler)(unsafe.Pointer(p))
	ptr.UserData = unsafe.Pointer(value)
}

// GuiSettingsHandler.WriteAllFn is unsupported: category unsupported.

// GuiShrinkWidthItem wraps struct ImGuiShrinkWidthItem.
type GuiShrinkWidthItem uintptr

// GuiShrinkWidthItemNil is a null pointer.
var GuiShrinkWidthItemNil GuiShrinkWidthItem

// GuiShrinkWidthItemSizeOf is the byte size of ImGuiShrinkWidthItem.
const GuiShrinkWidthItemSizeOf = int(C.sizeof_ImGuiShrinkWidthItem)

// GuiShrinkWidthItemAlloc allocates a continuous block of GuiShrinkWidthItem.
func GuiShrinkWidthItemAlloc(alloc ffi.Allocator, count int) GuiShrinkWidthItem {
	ptr := alloc.Allocate(GuiShrinkWidthItemSizeOf * count)
	return GuiShrinkWidthItem(ptr)
}

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p GuiShrinkWidthItem) Offset(offset int) GuiShrinkWidthItem {
	return p + GuiShrinkWidthItem(offset*GuiShrinkWidthItemSizeOf)
}

// GetIndex returns the value in Index.
func (p GuiShrinkWidthItem) GetIndex() int32 {
	ptr := (*C.ImGuiShrinkWidthItem)(unsafe.Pointer(p))
	return int32(ptr.Index)
}

// SetIndex sets the value in Index.
func (p GuiShrinkWidthItem) SetIndex(value int32) {
	ptr := (*C.ImGuiShrinkWidthItem)(unsafe.Pointer(p))
	ptr.Index = (C.int32_t)(value)
}

// GetInitialWidth returns the value in InitialWidth.
func (p GuiShrinkWidthItem) GetInitialWidth() float32 {
	ptr := (*C.ImGuiShrinkWidthItem)(unsafe.Pointer(p))
	return float32(ptr.InitialWidth)
}

// SetInitialWidth sets the value in InitialWidth.
func (p GuiShrinkWidthItem) SetInitialWidth(value float32) {
	ptr := (*C.ImGuiShrinkWidthItem)(unsafe.Pointer(p))
	ptr.InitialWidth = (C.float)(value)
}

// GetWidth returns the value in Width.
func (p GuiShrinkWidthItem) GetWidth() float32 {
	ptr := (*C.ImGuiShrinkWidthItem)(unsafe.Pointer(p))
	return float32(ptr.Width)
}

// SetWidth sets the value in Width.
func (p GuiShrinkWidthItem) SetWidth(value float32) {
	ptr := (*C.ImGuiShrinkWidthItem)(unsafe.Pointer(p))
	ptr.Width = (C.float)(value)
}

// GuiSizeCallbackData wraps struct ImGuiSizeCallbackData.
type GuiSizeCallbackData uintptr

// GuiSizeCallbackDataNil is a null pointer.
var GuiSizeCallbackDataNil GuiSizeCallbackData

// GuiSizeCallbackDataSizeOf is the byte size of ImGuiSizeCallbackData.
const GuiSizeCallbackDataSizeOf = int(C.sizeof_ImGuiSizeCallbackData)

// GuiSizeCallbackDataAlloc allocates a continuous block of GuiSizeCallbackData.
func GuiSizeCallbackDataAlloc(alloc ffi.Allocator, count int) GuiSizeCallbackData {
	ptr := alloc.Allocate(GuiSizeCallbackDataSizeOf * count)
	return GuiSizeCallbackData(ptr)
}

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p GuiSizeCallbackData) Offset(offset int) GuiSizeCallbackData {
	return p + GuiSizeCallbackData(offset*GuiSizeCallbackDataSizeOf)
}

// RefCurrentSize returns pointer to the CurrentSize field.
func (p GuiSizeCallbackData) RefCurrentSize() Vec2 {
	return Vec2(p + GuiSizeCallbackData(C.offsetof_ImGuiSizeCallbackData_CurrentSize))
}

// RefDesiredSize returns pointer to the DesiredSize field.
func (p GuiSizeCallbackData) RefDesiredSize() Vec2 {
	return Vec2(p + GuiSizeCallbackData(C.offsetof_ImGuiSizeCallbackData_DesiredSize))
}

// RefPos returns pointer to the Pos field.
func (p GuiSizeCallbackData) RefPos() Vec2 {
	return Vec2(p + GuiSizeCallbackData(C.offsetof_ImGuiSizeCallbackData_Pos))
}

// GetUserData returns the value in UserData.
func (p GuiSizeCallbackData) GetUserData() uintptr {
	ptr := (*C.ImGuiSizeCallbackData)(unsafe.Pointer(p))
	return uintptr(unsafe.Pointer(ptr.UserData))
}

// SetUserData sets the value in UserData.
func (p GuiSizeCallbackData) SetUserData(value uintptr) {
	ptr := (*C.ImGuiSizeCallbackData)(unsafe.Pointer(p))
	ptr.UserData = unsafe.Pointer(value)
}

// GuiStackLevelInfo wraps struct ImGuiStackLevelInfo.
type GuiStackLevelInfo uintptr

// GuiStackLevelInfoNil is a null pointer.
var GuiStackLevelInfoNil GuiStackLevelInfo

// GuiStackLevelInfoSizeOf is the byte size of ImGuiStackLevelInfo.
const GuiStackLevelInfoSizeOf = int(C.sizeof_ImGuiStackLevelInfo)

// ImGuiStackLevelInfo allocator is suppressed.

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p GuiStackLevelInfo) Offset(offset int) GuiStackLevelInfo {
	return p + GuiStackLevelInfo(offset*GuiStackLevelInfoSizeOf)
}

// GuiStackLevelInfo.DataType is unsupported: category unsupported.

// GetDescOffset returns the value in DescOffset.
func (p GuiStackLevelInfo) GetDescOffset() int32 {
	ptr := (*C.ImGuiStackLevelInfo)(unsafe.Pointer(p))
	return int32(ptr.DescOffset)
}

// SetDescOffset sets the value in DescOffset.
func (p GuiStackLevelInfo) SetDescOffset(value int32) {
	ptr := (*C.ImGuiStackLevelInfo)(unsafe.Pointer(p))
	ptr.DescOffset = (C.int32_t)(value)
}

// GuiStackLevelInfo.ID is unsupported: category unsupported.

// GuiStackLevelInfo.QueryFrameCount is unsupported: category unsupported.

// GetQuerySuccess returns the value in QuerySuccess.
func (p GuiStackLevelInfo) GetQuerySuccess() bool {
	ptr := (*C.ImGuiStackLevelInfo)(unsafe.Pointer(p))
	return bool(ptr.QuerySuccess)
}

// SetQuerySuccess sets the value in QuerySuccess.
func (p GuiStackLevelInfo) SetQuerySuccess(value bool) {
	ptr := (*C.ImGuiStackLevelInfo)(unsafe.Pointer(p))
	ptr.QuerySuccess = (C.bool)(value)
}

// GuiStorage wraps struct ImGuiStorage.
type GuiStorage uintptr

// GuiStorageNil is a null pointer.
var GuiStorageNil GuiStorage

// GuiStorageSizeOf is the byte size of ImGuiStorage.
const GuiStorageSizeOf = int(C.sizeof_ImGuiStorage)

// GuiStorageAlloc allocates a continuous block of GuiStorage.
func GuiStorageAlloc(alloc ffi.Allocator, count int) GuiStorage {
	ptr := alloc.Allocate(GuiStorageSizeOf * count)
	return GuiStorage(ptr)
}

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p GuiStorage) Offset(offset int) GuiStorage {
	return p + GuiStorage(offset*GuiStorageSizeOf)
}

// RefData returns pointer to the Data field.
func (p GuiStorage) RefData() Vector_ImGuiStoragePair {
	return Vector_ImGuiStoragePair(p + GuiStorage(C.offsetof_ImGuiStorage_Data))
}

// GuiStoragePair wraps struct ImGuiStoragePair.
type GuiStoragePair uintptr

// GuiStoragePairNil is a null pointer.
var GuiStoragePairNil GuiStoragePair

// GuiStoragePairSizeOf is the byte size of ImGuiStoragePair.
const GuiStoragePairSizeOf = int(C.sizeof_ImGuiStoragePair)

// ImGuiStoragePair allocator is suppressed.

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p GuiStoragePair) Offset(offset int) GuiStoragePair {
	return p + GuiStoragePair(offset*GuiStoragePairSizeOf)
}

// GuiStoragePair. is unsupported: category unsupported.

// GuiStoragePair.key is unsupported: category unsupported.

// GuiStyle wraps struct ImGuiStyle.
type GuiStyle uintptr

// GuiStyleNil is a null pointer.
var GuiStyleNil GuiStyle

// GuiStyleSizeOf is the byte size of ImGuiStyle.
const GuiStyleSizeOf = int(C.sizeof_ImGuiStyle)

// ImGuiStyle allocator is suppressed.

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p GuiStyle) Offset(offset int) GuiStyle {
	return p + GuiStyle(offset*GuiStyleSizeOf)
}

// GetAlpha returns the value in Alpha.
func (p GuiStyle) GetAlpha() float32 {
	ptr := (*C.ImGuiStyle)(unsafe.Pointer(p))
	return float32(ptr.Alpha)
}

// SetAlpha sets the value in Alpha.
func (p GuiStyle) SetAlpha(value float32) {
	ptr := (*C.ImGuiStyle)(unsafe.Pointer(p))
	ptr.Alpha = (C.float)(value)
}

// GetAntiAliasedFill returns the value in AntiAliasedFill.
func (p GuiStyle) GetAntiAliasedFill() bool {
	ptr := (*C.ImGuiStyle)(unsafe.Pointer(p))
	return bool(ptr.AntiAliasedFill)
}

// SetAntiAliasedFill sets the value in AntiAliasedFill.
func (p GuiStyle) SetAntiAliasedFill(value bool) {
	ptr := (*C.ImGuiStyle)(unsafe.Pointer(p))
	ptr.AntiAliasedFill = (C.bool)(value)
}

// GetAntiAliasedLines returns the value in AntiAliasedLines.
func (p GuiStyle) GetAntiAliasedLines() bool {
	ptr := (*C.ImGuiStyle)(unsafe.Pointer(p))
	return bool(ptr.AntiAliasedLines)
}

// SetAntiAliasedLines sets the value in AntiAliasedLines.
func (p GuiStyle) SetAntiAliasedLines(value bool) {
	ptr := (*C.ImGuiStyle)(unsafe.Pointer(p))
	ptr.AntiAliasedLines = (C.bool)(value)
}

// GetAntiAliasedLinesUseTex returns the value in AntiAliasedLinesUseTex.
func (p GuiStyle) GetAntiAliasedLinesUseTex() bool {
	ptr := (*C.ImGuiStyle)(unsafe.Pointer(p))
	return bool(ptr.AntiAliasedLinesUseTex)
}

// SetAntiAliasedLinesUseTex sets the value in AntiAliasedLinesUseTex.
func (p GuiStyle) SetAntiAliasedLinesUseTex(value bool) {
	ptr := (*C.ImGuiStyle)(unsafe.Pointer(p))
	ptr.AntiAliasedLinesUseTex = (C.bool)(value)
}

// RefButtonTextAlign returns pointer to the ButtonTextAlign field.
func (p GuiStyle) RefButtonTextAlign() Vec2 {
	return Vec2(p + GuiStyle(C.offsetof_ImGuiStyle_ButtonTextAlign))
}

// RefCellPadding returns pointer to the CellPadding field.
func (p GuiStyle) RefCellPadding() Vec2 {
	return Vec2(p + GuiStyle(C.offsetof_ImGuiStyle_CellPadding))
}

// GetChildBorderSize returns the value in ChildBorderSize.
func (p GuiStyle) GetChildBorderSize() float32 {
	ptr := (*C.ImGuiStyle)(unsafe.Pointer(p))
	return float32(ptr.ChildBorderSize)
}

// SetChildBorderSize sets the value in ChildBorderSize.
func (p GuiStyle) SetChildBorderSize(value float32) {
	ptr := (*C.ImGuiStyle)(unsafe.Pointer(p))
	ptr.ChildBorderSize = (C.float)(value)
}

// GetChildRounding returns the value in ChildRounding.
func (p GuiStyle) GetChildRounding() float32 {
	ptr := (*C.ImGuiStyle)(unsafe.Pointer(p))
	return float32(ptr.ChildRounding)
}

// SetChildRounding sets the value in ChildRounding.
func (p GuiStyle) SetChildRounding(value float32) {
	ptr := (*C.ImGuiStyle)(unsafe.Pointer(p))
	ptr.ChildRounding = (C.float)(value)
}

// GetCircleTessellationMaxError returns the value in CircleTessellationMaxError.
func (p GuiStyle) GetCircleTessellationMaxError() float32 {
	ptr := (*C.ImGuiStyle)(unsafe.Pointer(p))
	return float32(ptr.CircleTessellationMaxError)
}

// SetCircleTessellationMaxError sets the value in CircleTessellationMaxError.
func (p GuiStyle) SetCircleTessellationMaxError(value float32) {
	ptr := (*C.ImGuiStyle)(unsafe.Pointer(p))
	ptr.CircleTessellationMaxError = (C.float)(value)
}

// GetColorButtonPosition returns the value in ColorButtonPosition.
func (p GuiStyle) GetColorButtonPosition() GuiDir {
	ptr := (*C.ImGuiStyle)(unsafe.Pointer(p))
	return GuiDir(ptr.ColorButtonPosition)
}

// SetColorButtonPosition sets the value in ColorButtonPosition.
func (p GuiStyle) SetColorButtonPosition(value GuiDir) {
	ptr := (*C.ImGuiStyle)(unsafe.Pointer(p))
	ptr.ColorButtonPosition = (C.ImGuiDir)(value)
}

// GuiStyle.Colors[ImGuiCol_COUNT] is unsupported: category unsupported.

// GetColumnsMinSpacing returns the value in ColumnsMinSpacing.
func (p GuiStyle) GetColumnsMinSpacing() float32 {
	ptr := (*C.ImGuiStyle)(unsafe.Pointer(p))
	return float32(ptr.ColumnsMinSpacing)
}

// SetColumnsMinSpacing sets the value in ColumnsMinSpacing.
func (p GuiStyle) SetColumnsMinSpacing(value float32) {
	ptr := (*C.ImGuiStyle)(unsafe.Pointer(p))
	ptr.ColumnsMinSpacing = (C.float)(value)
}

// GetCurveTessellationTol returns the value in CurveTessellationTol.
func (p GuiStyle) GetCurveTessellationTol() float32 {
	ptr := (*C.ImGuiStyle)(unsafe.Pointer(p))
	return float32(ptr.CurveTessellationTol)
}

// SetCurveTessellationTol sets the value in CurveTessellationTol.
func (p GuiStyle) SetCurveTessellationTol(value float32) {
	ptr := (*C.ImGuiStyle)(unsafe.Pointer(p))
	ptr.CurveTessellationTol = (C.float)(value)
}

// GetDisabledAlpha returns the value in DisabledAlpha.
func (p GuiStyle) GetDisabledAlpha() float32 {
	ptr := (*C.ImGuiStyle)(unsafe.Pointer(p))
	return float32(ptr.DisabledAlpha)
}

// SetDisabledAlpha sets the value in DisabledAlpha.
func (p GuiStyle) SetDisabledAlpha(value float32) {
	ptr := (*C.ImGuiStyle)(unsafe.Pointer(p))
	ptr.DisabledAlpha = (C.float)(value)
}

// RefDisplaySafeAreaPadding returns pointer to the DisplaySafeAreaPadding field.
func (p GuiStyle) RefDisplaySafeAreaPadding() Vec2 {
	return Vec2(p + GuiStyle(C.offsetof_ImGuiStyle_DisplaySafeAreaPadding))
}

// RefDisplayWindowPadding returns pointer to the DisplayWindowPadding field.
func (p GuiStyle) RefDisplayWindowPadding() Vec2 {
	return Vec2(p + GuiStyle(C.offsetof_ImGuiStyle_DisplayWindowPadding))
}

// GetDockingNodeHasCloseButton returns the value in DockingNodeHasCloseButton.
func (p GuiStyle) GetDockingNodeHasCloseButton() bool {
	ptr := (*C.ImGuiStyle)(unsafe.Pointer(p))
	return bool(ptr.DockingNodeHasCloseButton)
}

// SetDockingNodeHasCloseButton sets the value in DockingNodeHasCloseButton.
func (p GuiStyle) SetDockingNodeHasCloseButton(value bool) {
	ptr := (*C.ImGuiStyle)(unsafe.Pointer(p))
	ptr.DockingNodeHasCloseButton = (C.bool)(value)
}

// GetDockingSeparatorSize returns the value in DockingSeparatorSize.
func (p GuiStyle) GetDockingSeparatorSize() float32 {
	ptr := (*C.ImGuiStyle)(unsafe.Pointer(p))
	return float32(ptr.DockingSeparatorSize)
}

// SetDockingSeparatorSize sets the value in DockingSeparatorSize.
func (p GuiStyle) SetDockingSeparatorSize(value float32) {
	ptr := (*C.ImGuiStyle)(unsafe.Pointer(p))
	ptr.DockingSeparatorSize = (C.float)(value)
}

// GetFontScaleDpi returns the value in FontScaleDpi.
func (p GuiStyle) GetFontScaleDpi() float32 {
	ptr := (*C.ImGuiStyle)(unsafe.Pointer(p))
	return float32(ptr.FontScaleDpi)
}

// SetFontScaleDpi sets the value in FontScaleDpi.
func (p GuiStyle) SetFontScaleDpi(value float32) {
	ptr := (*C.ImGuiStyle)(unsafe.Pointer(p))
	ptr.FontScaleDpi = (C.float)(value)
}

// GetFontScaleMain returns the value in FontScaleMain.
func (p GuiStyle) GetFontScaleMain() float32 {
	ptr := (*C.ImGuiStyle)(unsafe.Pointer(p))
	return float32(ptr.FontScaleMain)
}

// SetFontScaleMain sets the value in FontScaleMain.
func (p GuiStyle) SetFontScaleMain(value float32) {
	ptr := (*C.ImGuiStyle)(unsafe.Pointer(p))
	ptr.FontScaleMain = (C.float)(value)
}

// GetFontSizeBase returns the value in FontSizeBase.
func (p GuiStyle) GetFontSizeBase() float32 {
	ptr := (*C.ImGuiStyle)(unsafe.Pointer(p))
	return float32(ptr.FontSizeBase)
}

// SetFontSizeBase sets the value in FontSizeBase.
func (p GuiStyle) SetFontSizeBase(value float32) {
	ptr := (*C.ImGuiStyle)(unsafe.Pointer(p))
	ptr.FontSizeBase = (C.float)(value)
}

// GetFrameBorderSize returns the value in FrameBorderSize.
func (p GuiStyle) GetFrameBorderSize() float32 {
	ptr := (*C.ImGuiStyle)(unsafe.Pointer(p))
	return float32(ptr.FrameBorderSize)
}

// SetFrameBorderSize sets the value in FrameBorderSize.
func (p GuiStyle) SetFrameBorderSize(value float32) {
	ptr := (*C.ImGuiStyle)(unsafe.Pointer(p))
	ptr.FrameBorderSize = (C.float)(value)
}

// RefFramePadding returns pointer to the FramePadding field.
func (p GuiStyle) RefFramePadding() Vec2 {
	return Vec2(p + GuiStyle(C.offsetof_ImGuiStyle_FramePadding))
}

// GetFrameRounding returns the value in FrameRounding.
func (p GuiStyle) GetFrameRounding() float32 {
	ptr := (*C.ImGuiStyle)(unsafe.Pointer(p))
	return float32(ptr.FrameRounding)
}

// SetFrameRounding sets the value in FrameRounding.
func (p GuiStyle) SetFrameRounding(value float32) {
	ptr := (*C.ImGuiStyle)(unsafe.Pointer(p))
	ptr.FrameRounding = (C.float)(value)
}

// GetGrabMinSize returns the value in GrabMinSize.
func (p GuiStyle) GetGrabMinSize() float32 {
	ptr := (*C.ImGuiStyle)(unsafe.Pointer(p))
	return float32(ptr.GrabMinSize)
}

// SetGrabMinSize sets the value in GrabMinSize.
func (p GuiStyle) SetGrabMinSize(value float32) {
	ptr := (*C.ImGuiStyle)(unsafe.Pointer(p))
	ptr.GrabMinSize = (C.float)(value)
}

// GetGrabRounding returns the value in GrabRounding.
func (p GuiStyle) GetGrabRounding() float32 {
	ptr := (*C.ImGuiStyle)(unsafe.Pointer(p))
	return float32(ptr.GrabRounding)
}

// SetGrabRounding sets the value in GrabRounding.
func (p GuiStyle) SetGrabRounding(value float32) {
	ptr := (*C.ImGuiStyle)(unsafe.Pointer(p))
	ptr.GrabRounding = (C.float)(value)
}

// GetHoverDelayNormal returns the value in HoverDelayNormal.
func (p GuiStyle) GetHoverDelayNormal() float32 {
	ptr := (*C.ImGuiStyle)(unsafe.Pointer(p))
	return float32(ptr.HoverDelayNormal)
}

// SetHoverDelayNormal sets the value in HoverDelayNormal.
func (p GuiStyle) SetHoverDelayNormal(value float32) {
	ptr := (*C.ImGuiStyle)(unsafe.Pointer(p))
	ptr.HoverDelayNormal = (C.float)(value)
}

// GetHoverDelayShort returns the value in HoverDelayShort.
func (p GuiStyle) GetHoverDelayShort() float32 {
	ptr := (*C.ImGuiStyle)(unsafe.Pointer(p))
	return float32(ptr.HoverDelayShort)
}

// SetHoverDelayShort sets the value in HoverDelayShort.
func (p GuiStyle) SetHoverDelayShort(value float32) {
	ptr := (*C.ImGuiStyle)(unsafe.Pointer(p))
	ptr.HoverDelayShort = (C.float)(value)
}

// GetHoverFlagsForTooltipMouse returns the value in HoverFlagsForTooltipMouse.
func (p GuiStyle) GetHoverFlagsForTooltipMouse() GuiHoveredFlags {
	ptr := (*C.ImGuiStyle)(unsafe.Pointer(p))
	return GuiHoveredFlags(ptr.HoverFlagsForTooltipMouse)
}

// SetHoverFlagsForTooltipMouse sets the value in HoverFlagsForTooltipMouse.
func (p GuiStyle) SetHoverFlagsForTooltipMouse(value GuiHoveredFlags) {
	ptr := (*C.ImGuiStyle)(unsafe.Pointer(p))
	ptr.HoverFlagsForTooltipMouse = (C.ImGuiHoveredFlags)(value)
}

// GetHoverFlagsForTooltipNav returns the value in HoverFlagsForTooltipNav.
func (p GuiStyle) GetHoverFlagsForTooltipNav() GuiHoveredFlags {
	ptr := (*C.ImGuiStyle)(unsafe.Pointer(p))
	return GuiHoveredFlags(ptr.HoverFlagsForTooltipNav)
}

// SetHoverFlagsForTooltipNav sets the value in HoverFlagsForTooltipNav.
func (p GuiStyle) SetHoverFlagsForTooltipNav(value GuiHoveredFlags) {
	ptr := (*C.ImGuiStyle)(unsafe.Pointer(p))
	ptr.HoverFlagsForTooltipNav = (C.ImGuiHoveredFlags)(value)
}

// GetHoverStationaryDelay returns the value in HoverStationaryDelay.
func (p GuiStyle) GetHoverStationaryDelay() float32 {
	ptr := (*C.ImGuiStyle)(unsafe.Pointer(p))
	return float32(ptr.HoverStationaryDelay)
}

// SetHoverStationaryDelay sets the value in HoverStationaryDelay.
func (p GuiStyle) SetHoverStationaryDelay(value float32) {
	ptr := (*C.ImGuiStyle)(unsafe.Pointer(p))
	ptr.HoverStationaryDelay = (C.float)(value)
}

// GetImageBorderSize returns the value in ImageBorderSize.
func (p GuiStyle) GetImageBorderSize() float32 {
	ptr := (*C.ImGuiStyle)(unsafe.Pointer(p))
	return float32(ptr.ImageBorderSize)
}

// SetImageBorderSize sets the value in ImageBorderSize.
func (p GuiStyle) SetImageBorderSize(value float32) {
	ptr := (*C.ImGuiStyle)(unsafe.Pointer(p))
	ptr.ImageBorderSize = (C.float)(value)
}

// GetIndentSpacing returns the value in IndentSpacing.
func (p GuiStyle) GetIndentSpacing() float32 {
	ptr := (*C.ImGuiStyle)(unsafe.Pointer(p))
	return float32(ptr.IndentSpacing)
}

// SetIndentSpacing sets the value in IndentSpacing.
func (p GuiStyle) SetIndentSpacing(value float32) {
	ptr := (*C.ImGuiStyle)(unsafe.Pointer(p))
	ptr.IndentSpacing = (C.float)(value)
}

// RefItemInnerSpacing returns pointer to the ItemInnerSpacing field.
func (p GuiStyle) RefItemInnerSpacing() Vec2 {
	return Vec2(p + GuiStyle(C.offsetof_ImGuiStyle_ItemInnerSpacing))
}

// RefItemSpacing returns pointer to the ItemSpacing field.
func (p GuiStyle) RefItemSpacing() Vec2 {
	return Vec2(p + GuiStyle(C.offsetof_ImGuiStyle_ItemSpacing))
}

// GetLogSliderDeadzone returns the value in LogSliderDeadzone.
func (p GuiStyle) GetLogSliderDeadzone() float32 {
	ptr := (*C.ImGuiStyle)(unsafe.Pointer(p))
	return float32(ptr.LogSliderDeadzone)
}

// SetLogSliderDeadzone sets the value in LogSliderDeadzone.
func (p GuiStyle) SetLogSliderDeadzone(value float32) {
	ptr := (*C.ImGuiStyle)(unsafe.Pointer(p))
	ptr.LogSliderDeadzone = (C.float)(value)
}

// GetMouseCursorScale returns the value in MouseCursorScale.
func (p GuiStyle) GetMouseCursorScale() float32 {
	ptr := (*C.ImGuiStyle)(unsafe.Pointer(p))
	return float32(ptr.MouseCursorScale)
}

// SetMouseCursorScale sets the value in MouseCursorScale.
func (p GuiStyle) SetMouseCursorScale(value float32) {
	ptr := (*C.ImGuiStyle)(unsafe.Pointer(p))
	ptr.MouseCursorScale = (C.float)(value)
}

// GetPopupBorderSize returns the value in PopupBorderSize.
func (p GuiStyle) GetPopupBorderSize() float32 {
	ptr := (*C.ImGuiStyle)(unsafe.Pointer(p))
	return float32(ptr.PopupBorderSize)
}

// SetPopupBorderSize sets the value in PopupBorderSize.
func (p GuiStyle) SetPopupBorderSize(value float32) {
	ptr := (*C.ImGuiStyle)(unsafe.Pointer(p))
	ptr.PopupBorderSize = (C.float)(value)
}

// GetPopupRounding returns the value in PopupRounding.
func (p GuiStyle) GetPopupRounding() float32 {
	ptr := (*C.ImGuiStyle)(unsafe.Pointer(p))
	return float32(ptr.PopupRounding)
}

// SetPopupRounding sets the value in PopupRounding.
func (p GuiStyle) SetPopupRounding(value float32) {
	ptr := (*C.ImGuiStyle)(unsafe.Pointer(p))
	ptr.PopupRounding = (C.float)(value)
}

// GetScrollbarPadding returns the value in ScrollbarPadding.
func (p GuiStyle) GetScrollbarPadding() float32 {
	ptr := (*C.ImGuiStyle)(unsafe.Pointer(p))
	return float32(ptr.ScrollbarPadding)
}

// SetScrollbarPadding sets the value in ScrollbarPadding.
func (p GuiStyle) SetScrollbarPadding(value float32) {
	ptr := (*C.ImGuiStyle)(unsafe.Pointer(p))
	ptr.ScrollbarPadding = (C.float)(value)
}

// GetScrollbarRounding returns the value in ScrollbarRounding.
func (p GuiStyle) GetScrollbarRounding() float32 {
	ptr := (*C.ImGuiStyle)(unsafe.Pointer(p))
	return float32(ptr.ScrollbarRounding)
}

// SetScrollbarRounding sets the value in ScrollbarRounding.
func (p GuiStyle) SetScrollbarRounding(value float32) {
	ptr := (*C.ImGuiStyle)(unsafe.Pointer(p))
	ptr.ScrollbarRounding = (C.float)(value)
}

// GetScrollbarSize returns the value in ScrollbarSize.
func (p GuiStyle) GetScrollbarSize() float32 {
	ptr := (*C.ImGuiStyle)(unsafe.Pointer(p))
	return float32(ptr.ScrollbarSize)
}

// SetScrollbarSize sets the value in ScrollbarSize.
func (p GuiStyle) SetScrollbarSize(value float32) {
	ptr := (*C.ImGuiStyle)(unsafe.Pointer(p))
	ptr.ScrollbarSize = (C.float)(value)
}

// RefSelectableTextAlign returns pointer to the SelectableTextAlign field.
func (p GuiStyle) RefSelectableTextAlign() Vec2 {
	return Vec2(p + GuiStyle(C.offsetof_ImGuiStyle_SelectableTextAlign))
}

// RefSeparatorTextAlign returns pointer to the SeparatorTextAlign field.
func (p GuiStyle) RefSeparatorTextAlign() Vec2 {
	return Vec2(p + GuiStyle(C.offsetof_ImGuiStyle_SeparatorTextAlign))
}

// GetSeparatorTextBorderSize returns the value in SeparatorTextBorderSize.
func (p GuiStyle) GetSeparatorTextBorderSize() float32 {
	ptr := (*C.ImGuiStyle)(unsafe.Pointer(p))
	return float32(ptr.SeparatorTextBorderSize)
}

// SetSeparatorTextBorderSize sets the value in SeparatorTextBorderSize.
func (p GuiStyle) SetSeparatorTextBorderSize(value float32) {
	ptr := (*C.ImGuiStyle)(unsafe.Pointer(p))
	ptr.SeparatorTextBorderSize = (C.float)(value)
}

// RefSeparatorTextPadding returns pointer to the SeparatorTextPadding field.
func (p GuiStyle) RefSeparatorTextPadding() Vec2 {
	return Vec2(p + GuiStyle(C.offsetof_ImGuiStyle_SeparatorTextPadding))
}

// GetTabBarBorderSize returns the value in TabBarBorderSize.
func (p GuiStyle) GetTabBarBorderSize() float32 {
	ptr := (*C.ImGuiStyle)(unsafe.Pointer(p))
	return float32(ptr.TabBarBorderSize)
}

// SetTabBarBorderSize sets the value in TabBarBorderSize.
func (p GuiStyle) SetTabBarBorderSize(value float32) {
	ptr := (*C.ImGuiStyle)(unsafe.Pointer(p))
	ptr.TabBarBorderSize = (C.float)(value)
}

// GetTabBarOverlineSize returns the value in TabBarOverlineSize.
func (p GuiStyle) GetTabBarOverlineSize() float32 {
	ptr := (*C.ImGuiStyle)(unsafe.Pointer(p))
	return float32(ptr.TabBarOverlineSize)
}

// SetTabBarOverlineSize sets the value in TabBarOverlineSize.
func (p GuiStyle) SetTabBarOverlineSize(value float32) {
	ptr := (*C.ImGuiStyle)(unsafe.Pointer(p))
	ptr.TabBarOverlineSize = (C.float)(value)
}

// GetTabBorderSize returns the value in TabBorderSize.
func (p GuiStyle) GetTabBorderSize() float32 {
	ptr := (*C.ImGuiStyle)(unsafe.Pointer(p))
	return float32(ptr.TabBorderSize)
}

// SetTabBorderSize sets the value in TabBorderSize.
func (p GuiStyle) SetTabBorderSize(value float32) {
	ptr := (*C.ImGuiStyle)(unsafe.Pointer(p))
	ptr.TabBorderSize = (C.float)(value)
}

// GetTabCloseButtonMinWidthSelected returns the value in TabCloseButtonMinWidthSelected.
func (p GuiStyle) GetTabCloseButtonMinWidthSelected() float32 {
	ptr := (*C.ImGuiStyle)(unsafe.Pointer(p))
	return float32(ptr.TabCloseButtonMinWidthSelected)
}

// SetTabCloseButtonMinWidthSelected sets the value in TabCloseButtonMinWidthSelected.
func (p GuiStyle) SetTabCloseButtonMinWidthSelected(value float32) {
	ptr := (*C.ImGuiStyle)(unsafe.Pointer(p))
	ptr.TabCloseButtonMinWidthSelected = (C.float)(value)
}

// GetTabCloseButtonMinWidthUnselected returns the value in TabCloseButtonMinWidthUnselected.
func (p GuiStyle) GetTabCloseButtonMinWidthUnselected() float32 {
	ptr := (*C.ImGuiStyle)(unsafe.Pointer(p))
	return float32(ptr.TabCloseButtonMinWidthUnselected)
}

// SetTabCloseButtonMinWidthUnselected sets the value in TabCloseButtonMinWidthUnselected.
func (p GuiStyle) SetTabCloseButtonMinWidthUnselected(value float32) {
	ptr := (*C.ImGuiStyle)(unsafe.Pointer(p))
	ptr.TabCloseButtonMinWidthUnselected = (C.float)(value)
}

// GetTabMinWidthBase returns the value in TabMinWidthBase.
func (p GuiStyle) GetTabMinWidthBase() float32 {
	ptr := (*C.ImGuiStyle)(unsafe.Pointer(p))
	return float32(ptr.TabMinWidthBase)
}

// SetTabMinWidthBase sets the value in TabMinWidthBase.
func (p GuiStyle) SetTabMinWidthBase(value float32) {
	ptr := (*C.ImGuiStyle)(unsafe.Pointer(p))
	ptr.TabMinWidthBase = (C.float)(value)
}

// GetTabMinWidthShrink returns the value in TabMinWidthShrink.
func (p GuiStyle) GetTabMinWidthShrink() float32 {
	ptr := (*C.ImGuiStyle)(unsafe.Pointer(p))
	return float32(ptr.TabMinWidthShrink)
}

// SetTabMinWidthShrink sets the value in TabMinWidthShrink.
func (p GuiStyle) SetTabMinWidthShrink(value float32) {
	ptr := (*C.ImGuiStyle)(unsafe.Pointer(p))
	ptr.TabMinWidthShrink = (C.float)(value)
}

// GetTabRounding returns the value in TabRounding.
func (p GuiStyle) GetTabRounding() float32 {
	ptr := (*C.ImGuiStyle)(unsafe.Pointer(p))
	return float32(ptr.TabRounding)
}

// SetTabRounding sets the value in TabRounding.
func (p GuiStyle) SetTabRounding(value float32) {
	ptr := (*C.ImGuiStyle)(unsafe.Pointer(p))
	ptr.TabRounding = (C.float)(value)
}

// GetTableAngledHeadersAngle returns the value in TableAngledHeadersAngle.
func (p GuiStyle) GetTableAngledHeadersAngle() float32 {
	ptr := (*C.ImGuiStyle)(unsafe.Pointer(p))
	return float32(ptr.TableAngledHeadersAngle)
}

// SetTableAngledHeadersAngle sets the value in TableAngledHeadersAngle.
func (p GuiStyle) SetTableAngledHeadersAngle(value float32) {
	ptr := (*C.ImGuiStyle)(unsafe.Pointer(p))
	ptr.TableAngledHeadersAngle = (C.float)(value)
}

// RefTableAngledHeadersTextAlign returns pointer to the TableAngledHeadersTextAlign field.
func (p GuiStyle) RefTableAngledHeadersTextAlign() Vec2 {
	return Vec2(p + GuiStyle(C.offsetof_ImGuiStyle_TableAngledHeadersTextAlign))
}

// RefTouchExtraPadding returns pointer to the TouchExtraPadding field.
func (p GuiStyle) RefTouchExtraPadding() Vec2 {
	return Vec2(p + GuiStyle(C.offsetof_ImGuiStyle_TouchExtraPadding))
}

// GetTreeLinesFlags returns the value in TreeLinesFlags.
func (p GuiStyle) GetTreeLinesFlags() GuiTreeNodeFlags {
	ptr := (*C.ImGuiStyle)(unsafe.Pointer(p))
	return GuiTreeNodeFlags(ptr.TreeLinesFlags)
}

// SetTreeLinesFlags sets the value in TreeLinesFlags.
func (p GuiStyle) SetTreeLinesFlags(value GuiTreeNodeFlags) {
	ptr := (*C.ImGuiStyle)(unsafe.Pointer(p))
	ptr.TreeLinesFlags = (C.ImGuiTreeNodeFlags)(value)
}

// GetTreeLinesRounding returns the value in TreeLinesRounding.
func (p GuiStyle) GetTreeLinesRounding() float32 {
	ptr := (*C.ImGuiStyle)(unsafe.Pointer(p))
	return float32(ptr.TreeLinesRounding)
}

// SetTreeLinesRounding sets the value in TreeLinesRounding.
func (p GuiStyle) SetTreeLinesRounding(value float32) {
	ptr := (*C.ImGuiStyle)(unsafe.Pointer(p))
	ptr.TreeLinesRounding = (C.float)(value)
}

// GetTreeLinesSize returns the value in TreeLinesSize.
func (p GuiStyle) GetTreeLinesSize() float32 {
	ptr := (*C.ImGuiStyle)(unsafe.Pointer(p))
	return float32(ptr.TreeLinesSize)
}

// SetTreeLinesSize sets the value in TreeLinesSize.
func (p GuiStyle) SetTreeLinesSize(value float32) {
	ptr := (*C.ImGuiStyle)(unsafe.Pointer(p))
	ptr.TreeLinesSize = (C.float)(value)
}

// GetWindowBorderHoverPadding returns the value in WindowBorderHoverPadding.
func (p GuiStyle) GetWindowBorderHoverPadding() float32 {
	ptr := (*C.ImGuiStyle)(unsafe.Pointer(p))
	return float32(ptr.WindowBorderHoverPadding)
}

// SetWindowBorderHoverPadding sets the value in WindowBorderHoverPadding.
func (p GuiStyle) SetWindowBorderHoverPadding(value float32) {
	ptr := (*C.ImGuiStyle)(unsafe.Pointer(p))
	ptr.WindowBorderHoverPadding = (C.float)(value)
}

// GetWindowBorderSize returns the value in WindowBorderSize.
func (p GuiStyle) GetWindowBorderSize() float32 {
	ptr := (*C.ImGuiStyle)(unsafe.Pointer(p))
	return float32(ptr.WindowBorderSize)
}

// SetWindowBorderSize sets the value in WindowBorderSize.
func (p GuiStyle) SetWindowBorderSize(value float32) {
	ptr := (*C.ImGuiStyle)(unsafe.Pointer(p))
	ptr.WindowBorderSize = (C.float)(value)
}

// GetWindowMenuButtonPosition returns the value in WindowMenuButtonPosition.
func (p GuiStyle) GetWindowMenuButtonPosition() GuiDir {
	ptr := (*C.ImGuiStyle)(unsafe.Pointer(p))
	return GuiDir(ptr.WindowMenuButtonPosition)
}

// SetWindowMenuButtonPosition sets the value in WindowMenuButtonPosition.
func (p GuiStyle) SetWindowMenuButtonPosition(value GuiDir) {
	ptr := (*C.ImGuiStyle)(unsafe.Pointer(p))
	ptr.WindowMenuButtonPosition = (C.ImGuiDir)(value)
}

// RefWindowMinSize returns pointer to the WindowMinSize field.
func (p GuiStyle) RefWindowMinSize() Vec2 {
	return Vec2(p + GuiStyle(C.offsetof_ImGuiStyle_WindowMinSize))
}

// RefWindowPadding returns pointer to the WindowPadding field.
func (p GuiStyle) RefWindowPadding() Vec2 {
	return Vec2(p + GuiStyle(C.offsetof_ImGuiStyle_WindowPadding))
}

// GetWindowRounding returns the value in WindowRounding.
func (p GuiStyle) GetWindowRounding() float32 {
	ptr := (*C.ImGuiStyle)(unsafe.Pointer(p))
	return float32(ptr.WindowRounding)
}

// SetWindowRounding sets the value in WindowRounding.
func (p GuiStyle) SetWindowRounding(value float32) {
	ptr := (*C.ImGuiStyle)(unsafe.Pointer(p))
	ptr.WindowRounding = (C.float)(value)
}

// RefWindowTitleAlign returns pointer to the WindowTitleAlign field.
func (p GuiStyle) RefWindowTitleAlign() Vec2 {
	return Vec2(p + GuiStyle(C.offsetof_ImGuiStyle_WindowTitleAlign))
}

// GetMainScale returns the value in _MainScale.
func (p GuiStyle) GetMainScale() float32 {
	ptr := (*C.ImGuiStyle)(unsafe.Pointer(p))
	return float32(ptr._MainScale)
}

// SetMainScale sets the value in _MainScale.
func (p GuiStyle) SetMainScale(value float32) {
	ptr := (*C.ImGuiStyle)(unsafe.Pointer(p))
	ptr._MainScale = (C.float)(value)
}

// GetNextFrameFontSizeBase returns the value in _NextFrameFontSizeBase.
func (p GuiStyle) GetNextFrameFontSizeBase() float32 {
	ptr := (*C.ImGuiStyle)(unsafe.Pointer(p))
	return float32(ptr._NextFrameFontSizeBase)
}

// SetNextFrameFontSizeBase sets the value in _NextFrameFontSizeBase.
func (p GuiStyle) SetNextFrameFontSizeBase(value float32) {
	ptr := (*C.ImGuiStyle)(unsafe.Pointer(p))
	ptr._NextFrameFontSizeBase = (C.float)(value)
}

// GuiStyleMod wraps struct ImGuiStyleMod.
type GuiStyleMod uintptr

// GuiStyleModNil is a null pointer.
var GuiStyleModNil GuiStyleMod

// GuiStyleModSizeOf is the byte size of ImGuiStyleMod.
const GuiStyleModSizeOf = int(C.sizeof_ImGuiStyleMod)

// ImGuiStyleMod allocator is suppressed.

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p GuiStyleMod) Offset(offset int) GuiStyleMod {
	return p + GuiStyleMod(offset*GuiStyleModSizeOf)
}

// GuiStyleMod. is unsupported: category unsupported.

// GetVarIdx returns the value in VarIdx.
func (p GuiStyleMod) GetVarIdx() GuiStyleVar {
	ptr := (*C.ImGuiStyleMod)(unsafe.Pointer(p))
	return GuiStyleVar(ptr.VarIdx)
}

// SetVarIdx sets the value in VarIdx.
func (p GuiStyleMod) SetVarIdx(value GuiStyleVar) {
	ptr := (*C.ImGuiStyleMod)(unsafe.Pointer(p))
	ptr.VarIdx = (C.ImGuiStyleVar)(value)
}

// GuiStyleVarInfo wraps struct ImGuiStyleVarInfo.
type GuiStyleVarInfo uintptr

// GuiStyleVarInfoNil is a null pointer.
var GuiStyleVarInfoNil GuiStyleVarInfo

// GuiStyleVarInfoSizeOf is the byte size of ImGuiStyleVarInfo.
const GuiStyleVarInfoSizeOf = int(C.sizeof_ImGuiStyleVarInfo)

// GuiStyleVarInfoAlloc allocates a continuous block of GuiStyleVarInfo.
func GuiStyleVarInfoAlloc(alloc ffi.Allocator, count int) GuiStyleVarInfo {
	ptr := alloc.Allocate(GuiStyleVarInfoSizeOf * count)
	return GuiStyleVarInfo(ptr)
}

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p GuiStyleVarInfo) Offset(offset int) GuiStyleVarInfo {
	return p + GuiStyleVarInfo(offset*GuiStyleVarInfoSizeOf)
}

// GuiStyleVarInfo.Count is unsupported: category unsupported.

// GuiStyleVarInfo.DataType is unsupported: category unsupported.

// GuiStyleVarInfo.Offset is unsupported: category unsupported.

// GuiTabBar wraps struct ImGuiTabBar.
type GuiTabBar uintptr

// GuiTabBarNil is a null pointer.
var GuiTabBarNil GuiTabBar

// GuiTabBarSizeOf is the byte size of ImGuiTabBar.
const GuiTabBarSizeOf = int(C.sizeof_ImGuiTabBar)

// ImGuiTabBar allocator is suppressed.

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p GuiTabBar) Offset(offset int) GuiTabBar {
	return p + GuiTabBar(offset*GuiTabBarSizeOf)
}

// RefBackupCursorPos returns pointer to the BackupCursorPos field.
func (p GuiTabBar) RefBackupCursorPos() Vec2 {
	return Vec2(p + GuiTabBar(C.offsetof_ImGuiTabBar_BackupCursorPos))
}

// RefBarRect returns pointer to the BarRect field.
func (p GuiTabBar) RefBarRect() Rect {
	return Rect(p + GuiTabBar(C.offsetof_ImGuiTabBar_BarRect))
}

// GetBarRectPrevWidth returns the value in BarRectPrevWidth.
func (p GuiTabBar) GetBarRectPrevWidth() float32 {
	ptr := (*C.ImGuiTabBar)(unsafe.Pointer(p))
	return float32(ptr.BarRectPrevWidth)
}

// SetBarRectPrevWidth sets the value in BarRectPrevWidth.
func (p GuiTabBar) SetBarRectPrevWidth(value float32) {
	ptr := (*C.ImGuiTabBar)(unsafe.Pointer(p))
	ptr.BarRectPrevWidth = (C.float)(value)
}

// GuiTabBar.BeginCount is unsupported: category unsupported.

// GetCurrFrameVisible returns the value in CurrFrameVisible.
func (p GuiTabBar) GetCurrFrameVisible() int32 {
	ptr := (*C.ImGuiTabBar)(unsafe.Pointer(p))
	return int32(ptr.CurrFrameVisible)
}

// SetCurrFrameVisible sets the value in CurrFrameVisible.
func (p GuiTabBar) SetCurrFrameVisible(value int32) {
	ptr := (*C.ImGuiTabBar)(unsafe.Pointer(p))
	ptr.CurrFrameVisible = (C.int32_t)(value)
}

// GetCurrTabsContentsHeight returns the value in CurrTabsContentsHeight.
func (p GuiTabBar) GetCurrTabsContentsHeight() float32 {
	ptr := (*C.ImGuiTabBar)(unsafe.Pointer(p))
	return float32(ptr.CurrTabsContentsHeight)
}

// SetCurrTabsContentsHeight sets the value in CurrTabsContentsHeight.
func (p GuiTabBar) SetCurrTabsContentsHeight(value float32) {
	ptr := (*C.ImGuiTabBar)(unsafe.Pointer(p))
	ptr.CurrTabsContentsHeight = (C.float)(value)
}

// GetFlags returns the value in Flags.
func (p GuiTabBar) GetFlags() GuiTabBarFlags {
	ptr := (*C.ImGuiTabBar)(unsafe.Pointer(p))
	return GuiTabBarFlags(ptr.Flags)
}

// SetFlags sets the value in Flags.
func (p GuiTabBar) SetFlags(value GuiTabBarFlags) {
	ptr := (*C.ImGuiTabBar)(unsafe.Pointer(p))
	ptr.Flags = (C.ImGuiTabBarFlags)(value)
}

// RefFramePadding returns pointer to the FramePadding field.
func (p GuiTabBar) RefFramePadding() Vec2 {
	return Vec2(p + GuiTabBar(C.offsetof_ImGuiTabBar_FramePadding))
}

// GuiTabBar.ID is unsupported: category unsupported.

// GetItemSpacingY returns the value in ItemSpacingY.
func (p GuiTabBar) GetItemSpacingY() float32 {
	ptr := (*C.ImGuiTabBar)(unsafe.Pointer(p))
	return float32(ptr.ItemSpacingY)
}

// SetItemSpacingY sets the value in ItemSpacingY.
func (p GuiTabBar) SetItemSpacingY(value float32) {
	ptr := (*C.ImGuiTabBar)(unsafe.Pointer(p))
	ptr.ItemSpacingY = (C.float)(value)
}

// GuiTabBar.LastTabItemIdx is unsupported: category unsupported.

// GuiTabBar.NextSelectedTabId is unsupported: category unsupported.

// GetPrevFrameVisible returns the value in PrevFrameVisible.
func (p GuiTabBar) GetPrevFrameVisible() int32 {
	ptr := (*C.ImGuiTabBar)(unsafe.Pointer(p))
	return int32(ptr.PrevFrameVisible)
}

// SetPrevFrameVisible sets the value in PrevFrameVisible.
func (p GuiTabBar) SetPrevFrameVisible(value int32) {
	ptr := (*C.ImGuiTabBar)(unsafe.Pointer(p))
	ptr.PrevFrameVisible = (C.int32_t)(value)
}

// GetPrevTabsContentsHeight returns the value in PrevTabsContentsHeight.
func (p GuiTabBar) GetPrevTabsContentsHeight() float32 {
	ptr := (*C.ImGuiTabBar)(unsafe.Pointer(p))
	return float32(ptr.PrevTabsContentsHeight)
}

// SetPrevTabsContentsHeight sets the value in PrevTabsContentsHeight.
func (p GuiTabBar) SetPrevTabsContentsHeight(value float32) {
	ptr := (*C.ImGuiTabBar)(unsafe.Pointer(p))
	ptr.PrevTabsContentsHeight = (C.float)(value)
}

// GuiTabBar.ReorderRequestOffset is unsupported: category unsupported.

// GuiTabBar.ReorderRequestTabId is unsupported: category unsupported.

// GetScrollButtonEnabled returns the value in ScrollButtonEnabled.
func (p GuiTabBar) GetScrollButtonEnabled() bool {
	ptr := (*C.ImGuiTabBar)(unsafe.Pointer(p))
	return bool(ptr.ScrollButtonEnabled)
}

// SetScrollButtonEnabled sets the value in ScrollButtonEnabled.
func (p GuiTabBar) SetScrollButtonEnabled(value bool) {
	ptr := (*C.ImGuiTabBar)(unsafe.Pointer(p))
	ptr.ScrollButtonEnabled = (C.bool)(value)
}

// GetScrollingAnim returns the value in ScrollingAnim.
func (p GuiTabBar) GetScrollingAnim() float32 {
	ptr := (*C.ImGuiTabBar)(unsafe.Pointer(p))
	return float32(ptr.ScrollingAnim)
}

// SetScrollingAnim sets the value in ScrollingAnim.
func (p GuiTabBar) SetScrollingAnim(value float32) {
	ptr := (*C.ImGuiTabBar)(unsafe.Pointer(p))
	ptr.ScrollingAnim = (C.float)(value)
}

// GetScrollingRectMaxX returns the value in ScrollingRectMaxX.
func (p GuiTabBar) GetScrollingRectMaxX() float32 {
	ptr := (*C.ImGuiTabBar)(unsafe.Pointer(p))
	return float32(ptr.ScrollingRectMaxX)
}

// SetScrollingRectMaxX sets the value in ScrollingRectMaxX.
func (p GuiTabBar) SetScrollingRectMaxX(value float32) {
	ptr := (*C.ImGuiTabBar)(unsafe.Pointer(p))
	ptr.ScrollingRectMaxX = (C.float)(value)
}

// GetScrollingRectMinX returns the value in ScrollingRectMinX.
func (p GuiTabBar) GetScrollingRectMinX() float32 {
	ptr := (*C.ImGuiTabBar)(unsafe.Pointer(p))
	return float32(ptr.ScrollingRectMinX)
}

// SetScrollingRectMinX sets the value in ScrollingRectMinX.
func (p GuiTabBar) SetScrollingRectMinX(value float32) {
	ptr := (*C.ImGuiTabBar)(unsafe.Pointer(p))
	ptr.ScrollingRectMinX = (C.float)(value)
}

// GetScrollingSpeed returns the value in ScrollingSpeed.
func (p GuiTabBar) GetScrollingSpeed() float32 {
	ptr := (*C.ImGuiTabBar)(unsafe.Pointer(p))
	return float32(ptr.ScrollingSpeed)
}

// SetScrollingSpeed sets the value in ScrollingSpeed.
func (p GuiTabBar) SetScrollingSpeed(value float32) {
	ptr := (*C.ImGuiTabBar)(unsafe.Pointer(p))
	ptr.ScrollingSpeed = (C.float)(value)
}

// GetScrollingTarget returns the value in ScrollingTarget.
func (p GuiTabBar) GetScrollingTarget() float32 {
	ptr := (*C.ImGuiTabBar)(unsafe.Pointer(p))
	return float32(ptr.ScrollingTarget)
}

// SetScrollingTarget sets the value in ScrollingTarget.
func (p GuiTabBar) SetScrollingTarget(value float32) {
	ptr := (*C.ImGuiTabBar)(unsafe.Pointer(p))
	ptr.ScrollingTarget = (C.float)(value)
}

// GetScrollingTargetDistToVisibility returns the value in ScrollingTargetDistToVisibility.
func (p GuiTabBar) GetScrollingTargetDistToVisibility() float32 {
	ptr := (*C.ImGuiTabBar)(unsafe.Pointer(p))
	return float32(ptr.ScrollingTargetDistToVisibility)
}

// SetScrollingTargetDistToVisibility sets the value in ScrollingTargetDistToVisibility.
func (p GuiTabBar) SetScrollingTargetDistToVisibility(value float32) {
	ptr := (*C.ImGuiTabBar)(unsafe.Pointer(p))
	ptr.ScrollingTargetDistToVisibility = (C.float)(value)
}

// GuiTabBar.SelectedTabId is unsupported: category unsupported.

// GetSeparatorMaxX returns the value in SeparatorMaxX.
func (p GuiTabBar) GetSeparatorMaxX() float32 {
	ptr := (*C.ImGuiTabBar)(unsafe.Pointer(p))
	return float32(ptr.SeparatorMaxX)
}

// SetSeparatorMaxX sets the value in SeparatorMaxX.
func (p GuiTabBar) SetSeparatorMaxX(value float32) {
	ptr := (*C.ImGuiTabBar)(unsafe.Pointer(p))
	ptr.SeparatorMaxX = (C.float)(value)
}

// GetSeparatorMinX returns the value in SeparatorMinX.
func (p GuiTabBar) GetSeparatorMinX() float32 {
	ptr := (*C.ImGuiTabBar)(unsafe.Pointer(p))
	return float32(ptr.SeparatorMinX)
}

// SetSeparatorMinX sets the value in SeparatorMinX.
func (p GuiTabBar) SetSeparatorMinX(value float32) {
	ptr := (*C.ImGuiTabBar)(unsafe.Pointer(p))
	ptr.SeparatorMinX = (C.float)(value)
}

// RefTabs returns pointer to the Tabs field.
func (p GuiTabBar) RefTabs() Vector_ImGuiTabItem {
	return Vector_ImGuiTabItem(p + GuiTabBar(C.offsetof_ImGuiTabBar_Tabs))
}

// GuiTabBar.TabsActiveCount is unsupported: category unsupported.

// GetTabsAddedNew returns the value in TabsAddedNew.
func (p GuiTabBar) GetTabsAddedNew() bool {
	ptr := (*C.ImGuiTabBar)(unsafe.Pointer(p))
	return bool(ptr.TabsAddedNew)
}

// SetTabsAddedNew sets the value in TabsAddedNew.
func (p GuiTabBar) SetTabsAddedNew(value bool) {
	ptr := (*C.ImGuiTabBar)(unsafe.Pointer(p))
	ptr.TabsAddedNew = (C.bool)(value)
}

// RefTabsNames returns pointer to the TabsNames field.
func (p GuiTabBar) RefTabsNames() GuiTextBuffer {
	return GuiTextBuffer(p + GuiTabBar(C.offsetof_ImGuiTabBar_TabsNames))
}

// GuiTabBar.VisibleTabId is unsupported: category unsupported.

// GetVisibleTabWasSubmitted returns the value in VisibleTabWasSubmitted.
func (p GuiTabBar) GetVisibleTabWasSubmitted() bool {
	ptr := (*C.ImGuiTabBar)(unsafe.Pointer(p))
	return bool(ptr.VisibleTabWasSubmitted)
}

// SetVisibleTabWasSubmitted sets the value in VisibleTabWasSubmitted.
func (p GuiTabBar) SetVisibleTabWasSubmitted(value bool) {
	ptr := (*C.ImGuiTabBar)(unsafe.Pointer(p))
	ptr.VisibleTabWasSubmitted = (C.bool)(value)
}

// GetWantLayout returns the value in WantLayout.
func (p GuiTabBar) GetWantLayout() bool {
	ptr := (*C.ImGuiTabBar)(unsafe.Pointer(p))
	return bool(ptr.WantLayout)
}

// SetWantLayout sets the value in WantLayout.
func (p GuiTabBar) SetWantLayout(value bool) {
	ptr := (*C.ImGuiTabBar)(unsafe.Pointer(p))
	ptr.WantLayout = (C.bool)(value)
}

// GetWidthAllTabs returns the value in WidthAllTabs.
func (p GuiTabBar) GetWidthAllTabs() float32 {
	ptr := (*C.ImGuiTabBar)(unsafe.Pointer(p))
	return float32(ptr.WidthAllTabs)
}

// SetWidthAllTabs sets the value in WidthAllTabs.
func (p GuiTabBar) SetWidthAllTabs(value float32) {
	ptr := (*C.ImGuiTabBar)(unsafe.Pointer(p))
	ptr.WidthAllTabs = (C.float)(value)
}

// GetWidthAllTabsIdeal returns the value in WidthAllTabsIdeal.
func (p GuiTabBar) GetWidthAllTabsIdeal() float32 {
	ptr := (*C.ImGuiTabBar)(unsafe.Pointer(p))
	return float32(ptr.WidthAllTabsIdeal)
}

// SetWidthAllTabsIdeal sets the value in WidthAllTabsIdeal.
func (p GuiTabBar) SetWidthAllTabsIdeal(value float32) {
	ptr := (*C.ImGuiTabBar)(unsafe.Pointer(p))
	ptr.WidthAllTabsIdeal = (C.float)(value)
}

// GetWindow returns the value in Window.
func (p GuiTabBar) GetWindow() GuiWindow {
	ptr := (*C.ImGuiTabBar)(unsafe.Pointer(p))
	return GuiWindow(unsafe.Pointer(ptr.Window))
}

// SetWindow sets the value in Window.
func (p GuiTabBar) SetWindow(value GuiWindow) {
	ptr := (*C.ImGuiTabBar)(unsafe.Pointer(p))
	ptr.Window = (*C.ImGuiWindow)(unsafe.Pointer(value))
}

// GuiTabItem wraps struct ImGuiTabItem.
type GuiTabItem uintptr

// GuiTabItemNil is a null pointer.
var GuiTabItemNil GuiTabItem

// GuiTabItemSizeOf is the byte size of ImGuiTabItem.
const GuiTabItemSizeOf = int(C.sizeof_ImGuiTabItem)

// ImGuiTabItem allocator is suppressed.

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p GuiTabItem) Offset(offset int) GuiTabItem {
	return p + GuiTabItem(offset*GuiTabItemSizeOf)
}

// GuiTabItem.BeginOrder is unsupported: category unsupported.

// GetContentWidth returns the value in ContentWidth.
func (p GuiTabItem) GetContentWidth() float32 {
	ptr := (*C.ImGuiTabItem)(unsafe.Pointer(p))
	return float32(ptr.ContentWidth)
}

// SetContentWidth sets the value in ContentWidth.
func (p GuiTabItem) SetContentWidth(value float32) {
	ptr := (*C.ImGuiTabItem)(unsafe.Pointer(p))
	ptr.ContentWidth = (C.float)(value)
}

// GetFlags returns the value in Flags.
func (p GuiTabItem) GetFlags() GuiTabItemFlags {
	ptr := (*C.ImGuiTabItem)(unsafe.Pointer(p))
	return GuiTabItemFlags(ptr.Flags)
}

// SetFlags sets the value in Flags.
func (p GuiTabItem) SetFlags(value GuiTabItemFlags) {
	ptr := (*C.ImGuiTabItem)(unsafe.Pointer(p))
	ptr.Flags = (C.ImGuiTabItemFlags)(value)
}

// GuiTabItem.ID is unsupported: category unsupported.

// GuiTabItem.IndexDuringLayout is unsupported: category unsupported.

// GetLastFrameSelected returns the value in LastFrameSelected.
func (p GuiTabItem) GetLastFrameSelected() int32 {
	ptr := (*C.ImGuiTabItem)(unsafe.Pointer(p))
	return int32(ptr.LastFrameSelected)
}

// SetLastFrameSelected sets the value in LastFrameSelected.
func (p GuiTabItem) SetLastFrameSelected(value int32) {
	ptr := (*C.ImGuiTabItem)(unsafe.Pointer(p))
	ptr.LastFrameSelected = (C.int32_t)(value)
}

// GetLastFrameVisible returns the value in LastFrameVisible.
func (p GuiTabItem) GetLastFrameVisible() int32 {
	ptr := (*C.ImGuiTabItem)(unsafe.Pointer(p))
	return int32(ptr.LastFrameVisible)
}

// SetLastFrameVisible sets the value in LastFrameVisible.
func (p GuiTabItem) SetLastFrameVisible(value int32) {
	ptr := (*C.ImGuiTabItem)(unsafe.Pointer(p))
	ptr.LastFrameVisible = (C.int32_t)(value)
}

// GuiTabItem.NameOffset is unsupported: category unsupported.

// GetOffset returns the value in Offset.
func (p GuiTabItem) GetOffset() float32 {
	ptr := (*C.ImGuiTabItem)(unsafe.Pointer(p))
	return float32(ptr.Offset)
}

// SetOffset sets the value in Offset.
func (p GuiTabItem) SetOffset(value float32) {
	ptr := (*C.ImGuiTabItem)(unsafe.Pointer(p))
	ptr.Offset = (C.float)(value)
}

// GetRequestedWidth returns the value in RequestedWidth.
func (p GuiTabItem) GetRequestedWidth() float32 {
	ptr := (*C.ImGuiTabItem)(unsafe.Pointer(p))
	return float32(ptr.RequestedWidth)
}

// SetRequestedWidth sets the value in RequestedWidth.
func (p GuiTabItem) SetRequestedWidth(value float32) {
	ptr := (*C.ImGuiTabItem)(unsafe.Pointer(p))
	ptr.RequestedWidth = (C.float)(value)
}

// GetWantClose returns the value in WantClose.
func (p GuiTabItem) GetWantClose() bool {
	ptr := (*C.ImGuiTabItem)(unsafe.Pointer(p))
	return bool(ptr.WantClose)
}

// SetWantClose sets the value in WantClose.
func (p GuiTabItem) SetWantClose(value bool) {
	ptr := (*C.ImGuiTabItem)(unsafe.Pointer(p))
	ptr.WantClose = (C.bool)(value)
}

// GetWidth returns the value in Width.
func (p GuiTabItem) GetWidth() float32 {
	ptr := (*C.ImGuiTabItem)(unsafe.Pointer(p))
	return float32(ptr.Width)
}

// SetWidth sets the value in Width.
func (p GuiTabItem) SetWidth(value float32) {
	ptr := (*C.ImGuiTabItem)(unsafe.Pointer(p))
	ptr.Width = (C.float)(value)
}

// GetWindow returns the value in Window.
func (p GuiTabItem) GetWindow() GuiWindow {
	ptr := (*C.ImGuiTabItem)(unsafe.Pointer(p))
	return GuiWindow(unsafe.Pointer(ptr.Window))
}

// SetWindow sets the value in Window.
func (p GuiTabItem) SetWindow(value GuiWindow) {
	ptr := (*C.ImGuiTabItem)(unsafe.Pointer(p))
	ptr.Window = (*C.ImGuiWindow)(unsafe.Pointer(value))
}

// GuiTable wraps struct ImGuiTable.
type GuiTable uintptr

// GuiTableNil is a null pointer.
var GuiTableNil GuiTable

// GuiTableSizeOf is the byte size of ImGuiTable.
const GuiTableSizeOf = int(C.sizeof_ImGuiTable)

// ImGuiTable allocator is suppressed.

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p GuiTable) Offset(offset int) GuiTable {
	return p + GuiTable(offset*GuiTableSizeOf)
}

// GuiTable.AngledHeadersCount is unsupported: category unsupported.

// GetAngledHeadersHeight returns the value in AngledHeadersHeight.
func (p GuiTable) GetAngledHeadersHeight() float32 {
	ptr := (*C.ImGuiTable)(unsafe.Pointer(p))
	return float32(ptr.AngledHeadersHeight)
}

// SetAngledHeadersHeight sets the value in AngledHeadersHeight.
func (p GuiTable) SetAngledHeadersHeight(value float32) {
	ptr := (*C.ImGuiTable)(unsafe.Pointer(p))
	ptr.AngledHeadersHeight = (C.float)(value)
}

// GetAngledHeadersSlope returns the value in AngledHeadersSlope.
func (p GuiTable) GetAngledHeadersSlope() float32 {
	ptr := (*C.ImGuiTable)(unsafe.Pointer(p))
	return float32(ptr.AngledHeadersSlope)
}

// SetAngledHeadersSlope sets the value in AngledHeadersSlope.
func (p GuiTable) SetAngledHeadersSlope(value float32) {
	ptr := (*C.ImGuiTable)(unsafe.Pointer(p))
	ptr.AngledHeadersSlope = (C.float)(value)
}

// GuiTable.AutoFitSingleColumn is unsupported: category unsupported.

// RefBg0ClipRectForDrawCmd returns pointer to the Bg0ClipRectForDrawCmd field.
func (p GuiTable) RefBg0ClipRectForDrawCmd() Rect {
	return Rect(p + GuiTable(C.offsetof_ImGuiTable_Bg0ClipRectForDrawCmd))
}

// RefBg2ClipRectForDrawCmd returns pointer to the Bg2ClipRectForDrawCmd field.
func (p GuiTable) RefBg2ClipRectForDrawCmd() Rect {
	return Rect(p + GuiTable(C.offsetof_ImGuiTable_Bg2ClipRectForDrawCmd))
}

// GuiTable.Bg2DrawChannelCurrent is unsupported: category unsupported.

// GuiTable.Bg2DrawChannelUnfrozen is unsupported: category unsupported.

// RefBgClipRect returns pointer to the BgClipRect field.
func (p GuiTable) RefBgClipRect() Rect {
	return Rect(p + GuiTable(C.offsetof_ImGuiTable_BgClipRect))
}

// GetBorderColorLight returns the value in BorderColorLight.
func (p GuiTable) GetBorderColorLight() uint32 {
	ptr := (*C.ImGuiTable)(unsafe.Pointer(p))
	return uint32(ptr.BorderColorLight)
}

// SetBorderColorLight sets the value in BorderColorLight.
func (p GuiTable) SetBorderColorLight(value uint32) {
	ptr := (*C.ImGuiTable)(unsafe.Pointer(p))
	ptr.BorderColorLight = (C.uint32_t)(value)
}

// GetBorderColorStrong returns the value in BorderColorStrong.
func (p GuiTable) GetBorderColorStrong() uint32 {
	ptr := (*C.ImGuiTable)(unsafe.Pointer(p))
	return uint32(ptr.BorderColorStrong)
}

// SetBorderColorStrong sets the value in BorderColorStrong.
func (p GuiTable) SetBorderColorStrong(value uint32) {
	ptr := (*C.ImGuiTable)(unsafe.Pointer(p))
	ptr.BorderColorStrong = (C.uint32_t)(value)
}

// GetBorderX1 returns the value in BorderX1.
func (p GuiTable) GetBorderX1() float32 {
	ptr := (*C.ImGuiTable)(unsafe.Pointer(p))
	return float32(ptr.BorderX1)
}

// SetBorderX1 sets the value in BorderX1.
func (p GuiTable) SetBorderX1(value float32) {
	ptr := (*C.ImGuiTable)(unsafe.Pointer(p))
	ptr.BorderX1 = (C.float)(value)
}

// GetBorderX2 returns the value in BorderX2.
func (p GuiTable) GetBorderX2() float32 {
	ptr := (*C.ImGuiTable)(unsafe.Pointer(p))
	return float32(ptr.BorderX2)
}

// SetBorderX2 sets the value in BorderX2.
func (p GuiTable) SetBorderX2(value float32) {
	ptr := (*C.ImGuiTable)(unsafe.Pointer(p))
	ptr.BorderX2 = (C.float)(value)
}

// GetCellPaddingX returns the value in CellPaddingX.
func (p GuiTable) GetCellPaddingX() float32 {
	ptr := (*C.ImGuiTable)(unsafe.Pointer(p))
	return float32(ptr.CellPaddingX)
}

// SetCellPaddingX sets the value in CellPaddingX.
func (p GuiTable) SetCellPaddingX(value float32) {
	ptr := (*C.ImGuiTable)(unsafe.Pointer(p))
	ptr.CellPaddingX = (C.float)(value)
}

// GetCellSpacingX1 returns the value in CellSpacingX1.
func (p GuiTable) GetCellSpacingX1() float32 {
	ptr := (*C.ImGuiTable)(unsafe.Pointer(p))
	return float32(ptr.CellSpacingX1)
}

// SetCellSpacingX1 sets the value in CellSpacingX1.
func (p GuiTable) SetCellSpacingX1(value float32) {
	ptr := (*C.ImGuiTable)(unsafe.Pointer(p))
	ptr.CellSpacingX1 = (C.float)(value)
}

// GetCellSpacingX2 returns the value in CellSpacingX2.
func (p GuiTable) GetCellSpacingX2() float32 {
	ptr := (*C.ImGuiTable)(unsafe.Pointer(p))
	return float32(ptr.CellSpacingX2)
}

// SetCellSpacingX2 sets the value in CellSpacingX2.
func (p GuiTable) SetCellSpacingX2(value float32) {
	ptr := (*C.ImGuiTable)(unsafe.Pointer(p))
	ptr.CellSpacingX2 = (C.float)(value)
}

// RefColumns returns pointer to the Columns field.
func (p GuiTable) RefColumns() Span_ImGuiTableColumn {
	return Span_ImGuiTableColumn(p + GuiTable(C.offsetof_ImGuiTable_Columns))
}

// GetColumnsAutoFitWidth returns the value in ColumnsAutoFitWidth.
func (p GuiTable) GetColumnsAutoFitWidth() float32 {
	ptr := (*C.ImGuiTable)(unsafe.Pointer(p))
	return float32(ptr.ColumnsAutoFitWidth)
}

// SetColumnsAutoFitWidth sets the value in ColumnsAutoFitWidth.
func (p GuiTable) SetColumnsAutoFitWidth(value float32) {
	ptr := (*C.ImGuiTable)(unsafe.Pointer(p))
	ptr.ColumnsAutoFitWidth = (C.float)(value)
}

// GetColumnsCount returns the value in ColumnsCount.
func (p GuiTable) GetColumnsCount() int32 {
	ptr := (*C.ImGuiTable)(unsafe.Pointer(p))
	return int32(ptr.ColumnsCount)
}

// SetColumnsCount sets the value in ColumnsCount.
func (p GuiTable) SetColumnsCount(value int32) {
	ptr := (*C.ImGuiTable)(unsafe.Pointer(p))
	ptr.ColumnsCount = (C.int32_t)(value)
}

// GuiTable.ColumnsEnabledCount is unsupported: category unsupported.

// GuiTable.ColumnsEnabledFixedCount is unsupported: category unsupported.

// GetColumnsGivenWidth returns the value in ColumnsGivenWidth.
func (p GuiTable) GetColumnsGivenWidth() float32 {
	ptr := (*C.ImGuiTable)(unsafe.Pointer(p))
	return float32(ptr.ColumnsGivenWidth)
}

// SetColumnsGivenWidth sets the value in ColumnsGivenWidth.
func (p GuiTable) SetColumnsGivenWidth(value float32) {
	ptr := (*C.ImGuiTable)(unsafe.Pointer(p))
	ptr.ColumnsGivenWidth = (C.float)(value)
}

// RefColumnsNames returns pointer to the ColumnsNames field.
func (p GuiTable) RefColumnsNames() GuiTextBuffer {
	return GuiTextBuffer(p + GuiTable(C.offsetof_ImGuiTable_ColumnsNames))
}

// GetColumnsStretchSumWeights returns the value in ColumnsStretchSumWeights.
func (p GuiTable) GetColumnsStretchSumWeights() float32 {
	ptr := (*C.ImGuiTable)(unsafe.Pointer(p))
	return float32(ptr.ColumnsStretchSumWeights)
}

// SetColumnsStretchSumWeights sets the value in ColumnsStretchSumWeights.
func (p GuiTable) SetColumnsStretchSumWeights(value float32) {
	ptr := (*C.ImGuiTable)(unsafe.Pointer(p))
	ptr.ColumnsStretchSumWeights = (C.float)(value)
}

// GuiTable.ContextPopupColumn is unsupported: category unsupported.

// GetCurrentColumn returns the value in CurrentColumn.
func (p GuiTable) GetCurrentColumn() int32 {
	ptr := (*C.ImGuiTable)(unsafe.Pointer(p))
	return int32(ptr.CurrentColumn)
}

// SetCurrentColumn sets the value in CurrentColumn.
func (p GuiTable) SetCurrentColumn(value int32) {
	ptr := (*C.ImGuiTable)(unsafe.Pointer(p))
	ptr.CurrentColumn = (C.int32_t)(value)
}

// GetCurrentRow returns the value in CurrentRow.
func (p GuiTable) GetCurrentRow() int32 {
	ptr := (*C.ImGuiTable)(unsafe.Pointer(p))
	return int32(ptr.CurrentRow)
}

// SetCurrentRow sets the value in CurrentRow.
func (p GuiTable) SetCurrentRow(value int32) {
	ptr := (*C.ImGuiTable)(unsafe.Pointer(p))
	ptr.CurrentRow = (C.int32_t)(value)
}

// GuiTable.DeclColumnsCount is unsupported: category unsupported.

// GetDisableDefaultContextMenu returns the value in DisableDefaultContextMenu.
func (p GuiTable) GetDisableDefaultContextMenu() bool {
	ptr := (*C.ImGuiTable)(unsafe.Pointer(p))
	return bool(ptr.DisableDefaultContextMenu)
}

// SetDisableDefaultContextMenu sets the value in DisableDefaultContextMenu.
func (p GuiTable) SetDisableDefaultContextMenu(value bool) {
	ptr := (*C.ImGuiTable)(unsafe.Pointer(p))
	ptr.DisableDefaultContextMenu = (C.bool)(value)
}

// RefDisplayOrderToIndex returns pointer to the DisplayOrderToIndex field.
func (p GuiTable) RefDisplayOrderToIndex() Span_ImGuiTableColumnIdx {
	return Span_ImGuiTableColumnIdx(p + GuiTable(C.offsetof_ImGuiTable_DisplayOrderToIndex))
}

// GetDrawSplitter returns the value in DrawSplitter.
func (p GuiTable) GetDrawSplitter() DrawListSplitter {
	ptr := (*C.ImGuiTable)(unsafe.Pointer(p))
	return DrawListSplitter(unsafe.Pointer(ptr.DrawSplitter))
}

// SetDrawSplitter sets the value in DrawSplitter.
func (p GuiTable) SetDrawSplitter(value DrawListSplitter) {
	ptr := (*C.ImGuiTable)(unsafe.Pointer(p))
	ptr.DrawSplitter = (*C.ImDrawListSplitter)(unsafe.Pointer(value))
}

// GuiTable.DummyDrawChannel is unsupported: category unsupported.

// GuiTable.EnabledMaskByDisplayOrder is unsupported: category unsupported.

// GuiTable.EnabledMaskByIndex is unsupported: category unsupported.

// GetFlags returns the value in Flags.
func (p GuiTable) GetFlags() GuiTableFlags {
	ptr := (*C.ImGuiTable)(unsafe.Pointer(p))
	return GuiTableFlags(ptr.Flags)
}

// SetFlags sets the value in Flags.
func (p GuiTable) SetFlags(value GuiTableFlags) {
	ptr := (*C.ImGuiTable)(unsafe.Pointer(p))
	ptr.Flags = (C.ImGuiTableFlags)(value)
}

// GuiTable.FreezeColumnsCount is unsupported: category unsupported.

// GuiTable.FreezeColumnsRequest is unsupported: category unsupported.

// GuiTable.FreezeRowsCount is unsupported: category unsupported.

// GuiTable.FreezeRowsRequest is unsupported: category unsupported.

// GetHasScrollbarYCurr returns the value in HasScrollbarYCurr.
func (p GuiTable) GetHasScrollbarYCurr() bool {
	ptr := (*C.ImGuiTable)(unsafe.Pointer(p))
	return bool(ptr.HasScrollbarYCurr)
}

// SetHasScrollbarYCurr sets the value in HasScrollbarYCurr.
func (p GuiTable) SetHasScrollbarYCurr(value bool) {
	ptr := (*C.ImGuiTable)(unsafe.Pointer(p))
	ptr.HasScrollbarYCurr = (C.bool)(value)
}

// GetHasScrollbarYPrev returns the value in HasScrollbarYPrev.
func (p GuiTable) GetHasScrollbarYPrev() bool {
	ptr := (*C.ImGuiTable)(unsafe.Pointer(p))
	return bool(ptr.HasScrollbarYPrev)
}

// SetHasScrollbarYPrev sets the value in HasScrollbarYPrev.
func (p GuiTable) SetHasScrollbarYPrev(value bool) {
	ptr := (*C.ImGuiTable)(unsafe.Pointer(p))
	ptr.HasScrollbarYPrev = (C.bool)(value)
}

// GuiTable.HeldHeaderColumn is unsupported: category unsupported.

// GuiTable.HighlightColumnHeader is unsupported: category unsupported.

// RefHostBackupInnerClipRect returns pointer to the HostBackupInnerClipRect field.
func (p GuiTable) RefHostBackupInnerClipRect() Rect {
	return Rect(p + GuiTable(C.offsetof_ImGuiTable_HostBackupInnerClipRect))
}

// RefHostClipRect returns pointer to the HostClipRect field.
func (p GuiTable) RefHostClipRect() Rect {
	return Rect(p + GuiTable(C.offsetof_ImGuiTable_HostClipRect))
}

// GetHostIndentX returns the value in HostIndentX.
func (p GuiTable) GetHostIndentX() float32 {
	ptr := (*C.ImGuiTable)(unsafe.Pointer(p))
	return float32(ptr.HostIndentX)
}

// SetHostIndentX sets the value in HostIndentX.
func (p GuiTable) SetHostIndentX(value float32) {
	ptr := (*C.ImGuiTable)(unsafe.Pointer(p))
	ptr.HostIndentX = (C.float)(value)
}

// GetHostSkipItems returns the value in HostSkipItems.
func (p GuiTable) GetHostSkipItems() bool {
	ptr := (*C.ImGuiTable)(unsafe.Pointer(p))
	return bool(ptr.HostSkipItems)
}

// SetHostSkipItems sets the value in HostSkipItems.
func (p GuiTable) SetHostSkipItems(value bool) {
	ptr := (*C.ImGuiTable)(unsafe.Pointer(p))
	ptr.HostSkipItems = (C.bool)(value)
}

// GuiTable.HoveredColumnBody is unsupported: category unsupported.

// GuiTable.HoveredColumnBorder is unsupported: category unsupported.

// GuiTable.ID is unsupported: category unsupported.

// RefInnerClipRect returns pointer to the InnerClipRect field.
func (p GuiTable) RefInnerClipRect() Rect {
	return Rect(p + GuiTable(C.offsetof_ImGuiTable_InnerClipRect))
}

// RefInnerRect returns pointer to the InnerRect field.
func (p GuiTable) RefInnerRect() Rect {
	return Rect(p + GuiTable(C.offsetof_ImGuiTable_InnerRect))
}

// GetInnerWidth returns the value in InnerWidth.
func (p GuiTable) GetInnerWidth() float32 {
	ptr := (*C.ImGuiTable)(unsafe.Pointer(p))
	return float32(ptr.InnerWidth)
}

// SetInnerWidth sets the value in InnerWidth.
func (p GuiTable) SetInnerWidth(value float32) {
	ptr := (*C.ImGuiTable)(unsafe.Pointer(p))
	ptr.InnerWidth = (C.float)(value)
}

// GetInnerWindow returns the value in InnerWindow.
func (p GuiTable) GetInnerWindow() GuiWindow {
	ptr := (*C.ImGuiTable)(unsafe.Pointer(p))
	return GuiWindow(unsafe.Pointer(ptr.InnerWindow))
}

// SetInnerWindow sets the value in InnerWindow.
func (p GuiTable) SetInnerWindow(value GuiWindow) {
	ptr := (*C.ImGuiTable)(unsafe.Pointer(p))
	ptr.InnerWindow = (*C.ImGuiWindow)(unsafe.Pointer(value))
}

// GuiTable.InstanceCurrent is unsupported: category unsupported.

// RefInstanceDataExtra returns pointer to the InstanceDataExtra field.
func (p GuiTable) RefInstanceDataExtra() Vector_ImGuiTableInstanceData {
	return Vector_ImGuiTableInstanceData(p + GuiTable(C.offsetof_ImGuiTable_InstanceDataExtra))
}

// RefInstanceDataFirst returns pointer to the InstanceDataFirst field.
func (p GuiTable) RefInstanceDataFirst() GuiTableInstanceData {
	return GuiTableInstanceData(p + GuiTable(C.offsetof_ImGuiTable_InstanceDataFirst))
}

// GuiTable.InstanceInteracted is unsupported: category unsupported.

// GetIsActiveIdAliveBeforeTable returns the value in IsActiveIdAliveBeforeTable.
func (p GuiTable) GetIsActiveIdAliveBeforeTable() bool {
	ptr := (*C.ImGuiTable)(unsafe.Pointer(p))
	return bool(ptr.IsActiveIdAliveBeforeTable)
}

// SetIsActiveIdAliveBeforeTable sets the value in IsActiveIdAliveBeforeTable.
func (p GuiTable) SetIsActiveIdAliveBeforeTable(value bool) {
	ptr := (*C.ImGuiTable)(unsafe.Pointer(p))
	ptr.IsActiveIdAliveBeforeTable = (C.bool)(value)
}

// GetIsActiveIdInTable returns the value in IsActiveIdInTable.
func (p GuiTable) GetIsActiveIdInTable() bool {
	ptr := (*C.ImGuiTable)(unsafe.Pointer(p))
	return bool(ptr.IsActiveIdInTable)
}

// SetIsActiveIdInTable sets the value in IsActiveIdInTable.
func (p GuiTable) SetIsActiveIdInTable(value bool) {
	ptr := (*C.ImGuiTable)(unsafe.Pointer(p))
	ptr.IsActiveIdInTable = (C.bool)(value)
}

// GetIsContextPopupOpen returns the value in IsContextPopupOpen.
func (p GuiTable) GetIsContextPopupOpen() bool {
	ptr := (*C.ImGuiTable)(unsafe.Pointer(p))
	return bool(ptr.IsContextPopupOpen)
}

// SetIsContextPopupOpen sets the value in IsContextPopupOpen.
func (p GuiTable) SetIsContextPopupOpen(value bool) {
	ptr := (*C.ImGuiTable)(unsafe.Pointer(p))
	ptr.IsContextPopupOpen = (C.bool)(value)
}

// GetIsDefaultDisplayOrder returns the value in IsDefaultDisplayOrder.
func (p GuiTable) GetIsDefaultDisplayOrder() bool {
	ptr := (*C.ImGuiTable)(unsafe.Pointer(p))
	return bool(ptr.IsDefaultDisplayOrder)
}

// SetIsDefaultDisplayOrder sets the value in IsDefaultDisplayOrder.
func (p GuiTable) SetIsDefaultDisplayOrder(value bool) {
	ptr := (*C.ImGuiTable)(unsafe.Pointer(p))
	ptr.IsDefaultDisplayOrder = (C.bool)(value)
}

// GetIsDefaultSizingPolicy returns the value in IsDefaultSizingPolicy.
func (p GuiTable) GetIsDefaultSizingPolicy() bool {
	ptr := (*C.ImGuiTable)(unsafe.Pointer(p))
	return bool(ptr.IsDefaultSizingPolicy)
}

// SetIsDefaultSizingPolicy sets the value in IsDefaultSizingPolicy.
func (p GuiTable) SetIsDefaultSizingPolicy(value bool) {
	ptr := (*C.ImGuiTable)(unsafe.Pointer(p))
	ptr.IsDefaultSizingPolicy = (C.bool)(value)
}

// GetIsInitializing returns the value in IsInitializing.
func (p GuiTable) GetIsInitializing() bool {
	ptr := (*C.ImGuiTable)(unsafe.Pointer(p))
	return bool(ptr.IsInitializing)
}

// SetIsInitializing sets the value in IsInitializing.
func (p GuiTable) SetIsInitializing(value bool) {
	ptr := (*C.ImGuiTable)(unsafe.Pointer(p))
	ptr.IsInitializing = (C.bool)(value)
}

// GetIsInsideRow returns the value in IsInsideRow.
func (p GuiTable) GetIsInsideRow() bool {
	ptr := (*C.ImGuiTable)(unsafe.Pointer(p))
	return bool(ptr.IsInsideRow)
}

// SetIsInsideRow sets the value in IsInsideRow.
func (p GuiTable) SetIsInsideRow(value bool) {
	ptr := (*C.ImGuiTable)(unsafe.Pointer(p))
	ptr.IsInsideRow = (C.bool)(value)
}

// GetIsLayoutLocked returns the value in IsLayoutLocked.
func (p GuiTable) GetIsLayoutLocked() bool {
	ptr := (*C.ImGuiTable)(unsafe.Pointer(p))
	return bool(ptr.IsLayoutLocked)
}

// SetIsLayoutLocked sets the value in IsLayoutLocked.
func (p GuiTable) SetIsLayoutLocked(value bool) {
	ptr := (*C.ImGuiTable)(unsafe.Pointer(p))
	ptr.IsLayoutLocked = (C.bool)(value)
}

// GetIsResetAllRequest returns the value in IsResetAllRequest.
func (p GuiTable) GetIsResetAllRequest() bool {
	ptr := (*C.ImGuiTable)(unsafe.Pointer(p))
	return bool(ptr.IsResetAllRequest)
}

// SetIsResetAllRequest sets the value in IsResetAllRequest.
func (p GuiTable) SetIsResetAllRequest(value bool) {
	ptr := (*C.ImGuiTable)(unsafe.Pointer(p))
	ptr.IsResetAllRequest = (C.bool)(value)
}

// GetIsResetDisplayOrderRequest returns the value in IsResetDisplayOrderRequest.
func (p GuiTable) GetIsResetDisplayOrderRequest() bool {
	ptr := (*C.ImGuiTable)(unsafe.Pointer(p))
	return bool(ptr.IsResetDisplayOrderRequest)
}

// SetIsResetDisplayOrderRequest sets the value in IsResetDisplayOrderRequest.
func (p GuiTable) SetIsResetDisplayOrderRequest(value bool) {
	ptr := (*C.ImGuiTable)(unsafe.Pointer(p))
	ptr.IsResetDisplayOrderRequest = (C.bool)(value)
}

// GetIsSettingsDirty returns the value in IsSettingsDirty.
func (p GuiTable) GetIsSettingsDirty() bool {
	ptr := (*C.ImGuiTable)(unsafe.Pointer(p))
	return bool(ptr.IsSettingsDirty)
}

// SetIsSettingsDirty sets the value in IsSettingsDirty.
func (p GuiTable) SetIsSettingsDirty(value bool) {
	ptr := (*C.ImGuiTable)(unsafe.Pointer(p))
	ptr.IsSettingsDirty = (C.bool)(value)
}

// GetIsSettingsRequestLoad returns the value in IsSettingsRequestLoad.
func (p GuiTable) GetIsSettingsRequestLoad() bool {
	ptr := (*C.ImGuiTable)(unsafe.Pointer(p))
	return bool(ptr.IsSettingsRequestLoad)
}

// SetIsSettingsRequestLoad sets the value in IsSettingsRequestLoad.
func (p GuiTable) SetIsSettingsRequestLoad(value bool) {
	ptr := (*C.ImGuiTable)(unsafe.Pointer(p))
	ptr.IsSettingsRequestLoad = (C.bool)(value)
}

// GetIsSortSpecsDirty returns the value in IsSortSpecsDirty.
func (p GuiTable) GetIsSortSpecsDirty() bool {
	ptr := (*C.ImGuiTable)(unsafe.Pointer(p))
	return bool(ptr.IsSortSpecsDirty)
}

// SetIsSortSpecsDirty sets the value in IsSortSpecsDirty.
func (p GuiTable) SetIsSortSpecsDirty(value bool) {
	ptr := (*C.ImGuiTable)(unsafe.Pointer(p))
	ptr.IsSortSpecsDirty = (C.bool)(value)
}

// GetIsUnfrozenRows returns the value in IsUnfrozenRows.
func (p GuiTable) GetIsUnfrozenRows() bool {
	ptr := (*C.ImGuiTable)(unsafe.Pointer(p))
	return bool(ptr.IsUnfrozenRows)
}

// SetIsUnfrozenRows sets the value in IsUnfrozenRows.
func (p GuiTable) SetIsUnfrozenRows(value bool) {
	ptr := (*C.ImGuiTable)(unsafe.Pointer(p))
	ptr.IsUnfrozenRows = (C.bool)(value)
}

// GetIsUsingHeaders returns the value in IsUsingHeaders.
func (p GuiTable) GetIsUsingHeaders() bool {
	ptr := (*C.ImGuiTable)(unsafe.Pointer(p))
	return bool(ptr.IsUsingHeaders)
}

// SetIsUsingHeaders sets the value in IsUsingHeaders.
func (p GuiTable) SetIsUsingHeaders(value bool) {
	ptr := (*C.ImGuiTable)(unsafe.Pointer(p))
	ptr.IsUsingHeaders = (C.bool)(value)
}

// GetLastFrameActive returns the value in LastFrameActive.
func (p GuiTable) GetLastFrameActive() int32 {
	ptr := (*C.ImGuiTable)(unsafe.Pointer(p))
	return int32(ptr.LastFrameActive)
}

// SetLastFrameActive sets the value in LastFrameActive.
func (p GuiTable) SetLastFrameActive(value int32) {
	ptr := (*C.ImGuiTable)(unsafe.Pointer(p))
	ptr.LastFrameActive = (C.int32_t)(value)
}

// GuiTable.LastResizedColumn is unsupported: category unsupported.

// GuiTable.LastRowFlags is unsupported: category unsupported.

// GuiTable.LeftMostEnabledColumn is unsupported: category unsupported.

// GuiTable.LeftMostStretchedColumn is unsupported: category unsupported.

// GetMemoryCompacted returns the value in MemoryCompacted.
func (p GuiTable) GetMemoryCompacted() bool {
	ptr := (*C.ImGuiTable)(unsafe.Pointer(p))
	return bool(ptr.MemoryCompacted)
}

// SetMemoryCompacted sets the value in MemoryCompacted.
func (p GuiTable) SetMemoryCompacted(value bool) {
	ptr := (*C.ImGuiTable)(unsafe.Pointer(p))
	ptr.MemoryCompacted = (C.bool)(value)
}

// GetMinColumnWidth returns the value in MinColumnWidth.
func (p GuiTable) GetMinColumnWidth() float32 {
	ptr := (*C.ImGuiTable)(unsafe.Pointer(p))
	return float32(ptr.MinColumnWidth)
}

// SetMinColumnWidth sets the value in MinColumnWidth.
func (p GuiTable) SetMinColumnWidth(value float32) {
	ptr := (*C.ImGuiTable)(unsafe.Pointer(p))
	ptr.MinColumnWidth = (C.float)(value)
}

// GuiTable.NavLayer is unsupported: category unsupported.

// GetOuterPaddingX returns the value in OuterPaddingX.
func (p GuiTable) GetOuterPaddingX() float32 {
	ptr := (*C.ImGuiTable)(unsafe.Pointer(p))
	return float32(ptr.OuterPaddingX)
}

// SetOuterPaddingX sets the value in OuterPaddingX.
func (p GuiTable) SetOuterPaddingX(value float32) {
	ptr := (*C.ImGuiTable)(unsafe.Pointer(p))
	ptr.OuterPaddingX = (C.float)(value)
}

// RefOuterRect returns pointer to the OuterRect field.
func (p GuiTable) RefOuterRect() Rect {
	return Rect(p + GuiTable(C.offsetof_ImGuiTable_OuterRect))
}

// GetOuterWindow returns the value in OuterWindow.
func (p GuiTable) GetOuterWindow() GuiWindow {
	ptr := (*C.ImGuiTable)(unsafe.Pointer(p))
	return GuiWindow(unsafe.Pointer(ptr.OuterWindow))
}

// SetOuterWindow sets the value in OuterWindow.
func (p GuiTable) SetOuterWindow(value GuiWindow) {
	ptr := (*C.ImGuiTable)(unsafe.Pointer(p))
	ptr.OuterWindow = (*C.ImGuiWindow)(unsafe.Pointer(value))
}

// GetRawData returns the value in RawData.
func (p GuiTable) GetRawData() uintptr {
	ptr := (*C.ImGuiTable)(unsafe.Pointer(p))
	return uintptr(unsafe.Pointer(ptr.RawData))
}

// SetRawData sets the value in RawData.
func (p GuiTable) SetRawData(value uintptr) {
	ptr := (*C.ImGuiTable)(unsafe.Pointer(p))
	ptr.RawData = unsafe.Pointer(value)
}

// GetRefScale returns the value in RefScale.
func (p GuiTable) GetRefScale() float32 {
	ptr := (*C.ImGuiTable)(unsafe.Pointer(p))
	return float32(ptr.RefScale)
}

// SetRefScale sets the value in RefScale.
func (p GuiTable) SetRefScale(value float32) {
	ptr := (*C.ImGuiTable)(unsafe.Pointer(p))
	ptr.RefScale = (C.float)(value)
}

// GuiTable.ReorderColumn is unsupported: category unsupported.

// GuiTable.ReorderColumnDir is unsupported: category unsupported.

// GetResizeLockMinContentsX2 returns the value in ResizeLockMinContentsX2.
func (p GuiTable) GetResizeLockMinContentsX2() float32 {
	ptr := (*C.ImGuiTable)(unsafe.Pointer(p))
	return float32(ptr.ResizeLockMinContentsX2)
}

// SetResizeLockMinContentsX2 sets the value in ResizeLockMinContentsX2.
func (p GuiTable) SetResizeLockMinContentsX2(value float32) {
	ptr := (*C.ImGuiTable)(unsafe.Pointer(p))
	ptr.ResizeLockMinContentsX2 = (C.float)(value)
}

// GuiTable.ResizedColumn is unsupported: category unsupported.

// GetResizedColumnNextWidth returns the value in ResizedColumnNextWidth.
func (p GuiTable) GetResizedColumnNextWidth() float32 {
	ptr := (*C.ImGuiTable)(unsafe.Pointer(p))
	return float32(ptr.ResizedColumnNextWidth)
}

// SetResizedColumnNextWidth sets the value in ResizedColumnNextWidth.
func (p GuiTable) SetResizedColumnNextWidth(value float32) {
	ptr := (*C.ImGuiTable)(unsafe.Pointer(p))
	ptr.ResizedColumnNextWidth = (C.float)(value)
}

// GuiTable.RightMostEnabledColumn is unsupported: category unsupported.

// GuiTable.RightMostStretchedColumn is unsupported: category unsupported.

// GetRowBgColorCounter returns the value in RowBgColorCounter.
func (p GuiTable) GetRowBgColorCounter() int32 {
	ptr := (*C.ImGuiTable)(unsafe.Pointer(p))
	return int32(ptr.RowBgColorCounter)
}

// SetRowBgColorCounter sets the value in RowBgColorCounter.
func (p GuiTable) SetRowBgColorCounter(value int32) {
	ptr := (*C.ImGuiTable)(unsafe.Pointer(p))
	ptr.RowBgColorCounter = (C.int32_t)(value)
}

// GuiTable.RowBgColor[2] is unsupported: category unsupported.

// RefRowCellData returns pointer to the RowCellData field.
func (p GuiTable) RefRowCellData() Span_ImGuiTableCellData {
	return Span_ImGuiTableCellData(p + GuiTable(C.offsetof_ImGuiTable_RowCellData))
}

// GuiTable.RowCellDataCurrent is unsupported: category unsupported.

// GetRowCellPaddingY returns the value in RowCellPaddingY.
func (p GuiTable) GetRowCellPaddingY() float32 {
	ptr := (*C.ImGuiTable)(unsafe.Pointer(p))
	return float32(ptr.RowCellPaddingY)
}

// SetRowCellPaddingY sets the value in RowCellPaddingY.
func (p GuiTable) SetRowCellPaddingY(value float32) {
	ptr := (*C.ImGuiTable)(unsafe.Pointer(p))
	ptr.RowCellPaddingY = (C.float)(value)
}

// GuiTable.RowFlags is unsupported: category unsupported.

// GetRowIndentOffsetX returns the value in RowIndentOffsetX.
func (p GuiTable) GetRowIndentOffsetX() float32 {
	ptr := (*C.ImGuiTable)(unsafe.Pointer(p))
	return float32(ptr.RowIndentOffsetX)
}

// SetRowIndentOffsetX sets the value in RowIndentOffsetX.
func (p GuiTable) SetRowIndentOffsetX(value float32) {
	ptr := (*C.ImGuiTable)(unsafe.Pointer(p))
	ptr.RowIndentOffsetX = (C.float)(value)
}

// GetRowMinHeight returns the value in RowMinHeight.
func (p GuiTable) GetRowMinHeight() float32 {
	ptr := (*C.ImGuiTable)(unsafe.Pointer(p))
	return float32(ptr.RowMinHeight)
}

// SetRowMinHeight sets the value in RowMinHeight.
func (p GuiTable) SetRowMinHeight(value float32) {
	ptr := (*C.ImGuiTable)(unsafe.Pointer(p))
	ptr.RowMinHeight = (C.float)(value)
}

// GetRowPosY1 returns the value in RowPosY1.
func (p GuiTable) GetRowPosY1() float32 {
	ptr := (*C.ImGuiTable)(unsafe.Pointer(p))
	return float32(ptr.RowPosY1)
}

// SetRowPosY1 sets the value in RowPosY1.
func (p GuiTable) SetRowPosY1(value float32) {
	ptr := (*C.ImGuiTable)(unsafe.Pointer(p))
	ptr.RowPosY1 = (C.float)(value)
}

// GetRowPosY2 returns the value in RowPosY2.
func (p GuiTable) GetRowPosY2() float32 {
	ptr := (*C.ImGuiTable)(unsafe.Pointer(p))
	return float32(ptr.RowPosY2)
}

// SetRowPosY2 sets the value in RowPosY2.
func (p GuiTable) SetRowPosY2(value float32) {
	ptr := (*C.ImGuiTable)(unsafe.Pointer(p))
	ptr.RowPosY2 = (C.float)(value)
}

// GetRowTextBaseline returns the value in RowTextBaseline.
func (p GuiTable) GetRowTextBaseline() float32 {
	ptr := (*C.ImGuiTable)(unsafe.Pointer(p))
	return float32(ptr.RowTextBaseline)
}

// SetRowTextBaseline sets the value in RowTextBaseline.
func (p GuiTable) SetRowTextBaseline(value float32) {
	ptr := (*C.ImGuiTable)(unsafe.Pointer(p))
	ptr.RowTextBaseline = (C.float)(value)
}

// GetSettingsLoadedFlags returns the value in SettingsLoadedFlags.
func (p GuiTable) GetSettingsLoadedFlags() GuiTableFlags {
	ptr := (*C.ImGuiTable)(unsafe.Pointer(p))
	return GuiTableFlags(ptr.SettingsLoadedFlags)
}

// SetSettingsLoadedFlags sets the value in SettingsLoadedFlags.
func (p GuiTable) SetSettingsLoadedFlags(value GuiTableFlags) {
	ptr := (*C.ImGuiTable)(unsafe.Pointer(p))
	ptr.SettingsLoadedFlags = (C.ImGuiTableFlags)(value)
}

// GetSettingsOffset returns the value in SettingsOffset.
func (p GuiTable) GetSettingsOffset() int32 {
	ptr := (*C.ImGuiTable)(unsafe.Pointer(p))
	return int32(ptr.SettingsOffset)
}

// SetSettingsOffset sets the value in SettingsOffset.
func (p GuiTable) SetSettingsOffset(value int32) {
	ptr := (*C.ImGuiTable)(unsafe.Pointer(p))
	ptr.SettingsOffset = (C.int32_t)(value)
}

// RefSortSpecs returns pointer to the SortSpecs field.
func (p GuiTable) RefSortSpecs() GuiTableSortSpecs {
	return GuiTableSortSpecs(p + GuiTable(C.offsetof_ImGuiTable_SortSpecs))
}

// GuiTable.SortSpecsCount is unsupported: category unsupported.

// RefSortSpecsMulti returns pointer to the SortSpecsMulti field.
func (p GuiTable) RefSortSpecsMulti() Vector_ImGuiTableColumnSortSpecs {
	return Vector_ImGuiTableColumnSortSpecs(p + GuiTable(C.offsetof_ImGuiTable_SortSpecsMulti))
}

// RefSortSpecsSingle returns pointer to the SortSpecsSingle field.
func (p GuiTable) RefSortSpecsSingle() GuiTableColumnSortSpecs {
	return GuiTableColumnSortSpecs(p + GuiTable(C.offsetof_ImGuiTable_SortSpecsSingle))
}

// GetTempData returns the value in TempData.
func (p GuiTable) GetTempData() GuiTableTempData {
	ptr := (*C.ImGuiTable)(unsafe.Pointer(p))
	return GuiTableTempData(unsafe.Pointer(ptr.TempData))
}

// SetTempData sets the value in TempData.
func (p GuiTable) SetTempData(value GuiTableTempData) {
	ptr := (*C.ImGuiTable)(unsafe.Pointer(p))
	ptr.TempData = (*C.ImGuiTableTempData)(unsafe.Pointer(value))
}

// GuiTable.VisibleMaskByIndex is unsupported: category unsupported.

// RefWorkRect returns pointer to the WorkRect field.
func (p GuiTable) RefWorkRect() Rect {
	return Rect(p + GuiTable(C.offsetof_ImGuiTable_WorkRect))
}

// GuiTableCellData wraps struct ImGuiTableCellData.
type GuiTableCellData uintptr

// GuiTableCellDataNil is a null pointer.
var GuiTableCellDataNil GuiTableCellData

// GuiTableCellDataSizeOf is the byte size of ImGuiTableCellData.
const GuiTableCellDataSizeOf = int(C.sizeof_ImGuiTableCellData)

// GuiTableCellDataAlloc allocates a continuous block of GuiTableCellData.
func GuiTableCellDataAlloc(alloc ffi.Allocator, count int) GuiTableCellData {
	ptr := alloc.Allocate(GuiTableCellDataSizeOf * count)
	return GuiTableCellData(ptr)
}

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p GuiTableCellData) Offset(offset int) GuiTableCellData {
	return p + GuiTableCellData(offset*GuiTableCellDataSizeOf)
}

// GetBgColor returns the value in BgColor.
func (p GuiTableCellData) GetBgColor() uint32 {
	ptr := (*C.ImGuiTableCellData)(unsafe.Pointer(p))
	return uint32(ptr.BgColor)
}

// SetBgColor sets the value in BgColor.
func (p GuiTableCellData) SetBgColor(value uint32) {
	ptr := (*C.ImGuiTableCellData)(unsafe.Pointer(p))
	ptr.BgColor = (C.uint32_t)(value)
}

// GuiTableCellData.Column is unsupported: category unsupported.

// GuiTableColumn wraps struct ImGuiTableColumn.
type GuiTableColumn uintptr

// GuiTableColumnNil is a null pointer.
var GuiTableColumnNil GuiTableColumn

// GuiTableColumnSizeOf is the byte size of ImGuiTableColumn.
const GuiTableColumnSizeOf = int(C.sizeof_ImGuiTableColumn)

// ImGuiTableColumn allocator is suppressed.

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p GuiTableColumn) Offset(offset int) GuiTableColumn {
	return p + GuiTableColumn(offset*GuiTableColumnSizeOf)
}

// GetAutoFitQueue returns the value in AutoFitQueue.
func (p GuiTableColumn) GetAutoFitQueue() uint8 {
	ptr := (*C.ImGuiTableColumn)(unsafe.Pointer(p))
	return uint8(ptr.AutoFitQueue)
}

// SetAutoFitQueue sets the value in AutoFitQueue.
func (p GuiTableColumn) SetAutoFitQueue(value uint8) {
	ptr := (*C.ImGuiTableColumn)(unsafe.Pointer(p))
	ptr.AutoFitQueue = (C.uint8_t)(value)
}

// GetCannotSkipItemsQueue returns the value in CannotSkipItemsQueue.
func (p GuiTableColumn) GetCannotSkipItemsQueue() uint8 {
	ptr := (*C.ImGuiTableColumn)(unsafe.Pointer(p))
	return uint8(ptr.CannotSkipItemsQueue)
}

// SetCannotSkipItemsQueue sets the value in CannotSkipItemsQueue.
func (p GuiTableColumn) SetCannotSkipItemsQueue(value uint8) {
	ptr := (*C.ImGuiTableColumn)(unsafe.Pointer(p))
	ptr.CannotSkipItemsQueue = (C.uint8_t)(value)
}

// RefClipRect returns pointer to the ClipRect field.
func (p GuiTableColumn) RefClipRect() Rect {
	return Rect(p + GuiTableColumn(C.offsetof_ImGuiTableColumn_ClipRect))
}

// GetContentMaxXFrozen returns the value in ContentMaxXFrozen.
func (p GuiTableColumn) GetContentMaxXFrozen() float32 {
	ptr := (*C.ImGuiTableColumn)(unsafe.Pointer(p))
	return float32(ptr.ContentMaxXFrozen)
}

// SetContentMaxXFrozen sets the value in ContentMaxXFrozen.
func (p GuiTableColumn) SetContentMaxXFrozen(value float32) {
	ptr := (*C.ImGuiTableColumn)(unsafe.Pointer(p))
	ptr.ContentMaxXFrozen = (C.float)(value)
}

// GetContentMaxXHeadersIdeal returns the value in ContentMaxXHeadersIdeal.
func (p GuiTableColumn) GetContentMaxXHeadersIdeal() float32 {
	ptr := (*C.ImGuiTableColumn)(unsafe.Pointer(p))
	return float32(ptr.ContentMaxXHeadersIdeal)
}

// SetContentMaxXHeadersIdeal sets the value in ContentMaxXHeadersIdeal.
func (p GuiTableColumn) SetContentMaxXHeadersIdeal(value float32) {
	ptr := (*C.ImGuiTableColumn)(unsafe.Pointer(p))
	ptr.ContentMaxXHeadersIdeal = (C.float)(value)
}

// GetContentMaxXHeadersUsed returns the value in ContentMaxXHeadersUsed.
func (p GuiTableColumn) GetContentMaxXHeadersUsed() float32 {
	ptr := (*C.ImGuiTableColumn)(unsafe.Pointer(p))
	return float32(ptr.ContentMaxXHeadersUsed)
}

// SetContentMaxXHeadersUsed sets the value in ContentMaxXHeadersUsed.
func (p GuiTableColumn) SetContentMaxXHeadersUsed(value float32) {
	ptr := (*C.ImGuiTableColumn)(unsafe.Pointer(p))
	ptr.ContentMaxXHeadersUsed = (C.float)(value)
}

// GetContentMaxXUnfrozen returns the value in ContentMaxXUnfrozen.
func (p GuiTableColumn) GetContentMaxXUnfrozen() float32 {
	ptr := (*C.ImGuiTableColumn)(unsafe.Pointer(p))
	return float32(ptr.ContentMaxXUnfrozen)
}

// SetContentMaxXUnfrozen sets the value in ContentMaxXUnfrozen.
func (p GuiTableColumn) SetContentMaxXUnfrozen(value float32) {
	ptr := (*C.ImGuiTableColumn)(unsafe.Pointer(p))
	ptr.ContentMaxXUnfrozen = (C.float)(value)
}

// GuiTableColumn.DisplayOrder is unsupported: category unsupported.

// GuiTableColumn.DrawChannelCurrent is unsupported: category unsupported.

// GuiTableColumn.DrawChannelFrozen is unsupported: category unsupported.

// GuiTableColumn.DrawChannelUnfrozen is unsupported: category unsupported.

// GetFlags returns the value in Flags.
func (p GuiTableColumn) GetFlags() GuiTableColumnFlags {
	ptr := (*C.ImGuiTableColumn)(unsafe.Pointer(p))
	return GuiTableColumnFlags(ptr.Flags)
}

// SetFlags sets the value in Flags.
func (p GuiTableColumn) SetFlags(value GuiTableColumnFlags) {
	ptr := (*C.ImGuiTableColumn)(unsafe.Pointer(p))
	ptr.Flags = (C.ImGuiTableColumnFlags)(value)
}

// GuiTableColumn.IndexWithinEnabledSet is unsupported: category unsupported.

// GetInitStretchWeightOrWidth returns the value in InitStretchWeightOrWidth.
func (p GuiTableColumn) GetInitStretchWeightOrWidth() float32 {
	ptr := (*C.ImGuiTableColumn)(unsafe.Pointer(p))
	return float32(ptr.InitStretchWeightOrWidth)
}

// SetInitStretchWeightOrWidth sets the value in InitStretchWeightOrWidth.
func (p GuiTableColumn) SetInitStretchWeightOrWidth(value float32) {
	ptr := (*C.ImGuiTableColumn)(unsafe.Pointer(p))
	ptr.InitStretchWeightOrWidth = (C.float)(value)
}

// GetIsEnabled returns the value in IsEnabled.
func (p GuiTableColumn) GetIsEnabled() bool {
	ptr := (*C.ImGuiTableColumn)(unsafe.Pointer(p))
	return bool(ptr.IsEnabled)
}

// SetIsEnabled sets the value in IsEnabled.
func (p GuiTableColumn) SetIsEnabled(value bool) {
	ptr := (*C.ImGuiTableColumn)(unsafe.Pointer(p))
	ptr.IsEnabled = (C.bool)(value)
}

// GetIsPreserveWidthAuto returns the value in IsPreserveWidthAuto.
func (p GuiTableColumn) GetIsPreserveWidthAuto() bool {
	ptr := (*C.ImGuiTableColumn)(unsafe.Pointer(p))
	return bool(ptr.IsPreserveWidthAuto)
}

// SetIsPreserveWidthAuto sets the value in IsPreserveWidthAuto.
func (p GuiTableColumn) SetIsPreserveWidthAuto(value bool) {
	ptr := (*C.ImGuiTableColumn)(unsafe.Pointer(p))
	ptr.IsPreserveWidthAuto = (C.bool)(value)
}

// GetIsRequestOutput returns the value in IsRequestOutput.
func (p GuiTableColumn) GetIsRequestOutput() bool {
	ptr := (*C.ImGuiTableColumn)(unsafe.Pointer(p))
	return bool(ptr.IsRequestOutput)
}

// SetIsRequestOutput sets the value in IsRequestOutput.
func (p GuiTableColumn) SetIsRequestOutput(value bool) {
	ptr := (*C.ImGuiTableColumn)(unsafe.Pointer(p))
	ptr.IsRequestOutput = (C.bool)(value)
}

// GetIsSkipItems returns the value in IsSkipItems.
func (p GuiTableColumn) GetIsSkipItems() bool {
	ptr := (*C.ImGuiTableColumn)(unsafe.Pointer(p))
	return bool(ptr.IsSkipItems)
}

// SetIsSkipItems sets the value in IsSkipItems.
func (p GuiTableColumn) SetIsSkipItems(value bool) {
	ptr := (*C.ImGuiTableColumn)(unsafe.Pointer(p))
	ptr.IsSkipItems = (C.bool)(value)
}

// GetIsUserEnabled returns the value in IsUserEnabled.
func (p GuiTableColumn) GetIsUserEnabled() bool {
	ptr := (*C.ImGuiTableColumn)(unsafe.Pointer(p))
	return bool(ptr.IsUserEnabled)
}

// SetIsUserEnabled sets the value in IsUserEnabled.
func (p GuiTableColumn) SetIsUserEnabled(value bool) {
	ptr := (*C.ImGuiTableColumn)(unsafe.Pointer(p))
	ptr.IsUserEnabled = (C.bool)(value)
}

// GetIsUserEnabledNextFrame returns the value in IsUserEnabledNextFrame.
func (p GuiTableColumn) GetIsUserEnabledNextFrame() bool {
	ptr := (*C.ImGuiTableColumn)(unsafe.Pointer(p))
	return bool(ptr.IsUserEnabledNextFrame)
}

// SetIsUserEnabledNextFrame sets the value in IsUserEnabledNextFrame.
func (p GuiTableColumn) SetIsUserEnabledNextFrame(value bool) {
	ptr := (*C.ImGuiTableColumn)(unsafe.Pointer(p))
	ptr.IsUserEnabledNextFrame = (C.bool)(value)
}

// GetIsVisibleX returns the value in IsVisibleX.
func (p GuiTableColumn) GetIsVisibleX() bool {
	ptr := (*C.ImGuiTableColumn)(unsafe.Pointer(p))
	return bool(ptr.IsVisibleX)
}

// SetIsVisibleX sets the value in IsVisibleX.
func (p GuiTableColumn) SetIsVisibleX(value bool) {
	ptr := (*C.ImGuiTableColumn)(unsafe.Pointer(p))
	ptr.IsVisibleX = (C.bool)(value)
}

// GetIsVisibleY returns the value in IsVisibleY.
func (p GuiTableColumn) GetIsVisibleY() bool {
	ptr := (*C.ImGuiTableColumn)(unsafe.Pointer(p))
	return bool(ptr.IsVisibleY)
}

// SetIsVisibleY sets the value in IsVisibleY.
func (p GuiTableColumn) SetIsVisibleY(value bool) {
	ptr := (*C.ImGuiTableColumn)(unsafe.Pointer(p))
	ptr.IsVisibleY = (C.bool)(value)
}

// GetItemWidth returns the value in ItemWidth.
func (p GuiTableColumn) GetItemWidth() float32 {
	ptr := (*C.ImGuiTableColumn)(unsafe.Pointer(p))
	return float32(ptr.ItemWidth)
}

// SetItemWidth sets the value in ItemWidth.
func (p GuiTableColumn) SetItemWidth(value float32) {
	ptr := (*C.ImGuiTableColumn)(unsafe.Pointer(p))
	ptr.ItemWidth = (C.float)(value)
}

// GetMaxX returns the value in MaxX.
func (p GuiTableColumn) GetMaxX() float32 {
	ptr := (*C.ImGuiTableColumn)(unsafe.Pointer(p))
	return float32(ptr.MaxX)
}

// SetMaxX sets the value in MaxX.
func (p GuiTableColumn) SetMaxX(value float32) {
	ptr := (*C.ImGuiTableColumn)(unsafe.Pointer(p))
	ptr.MaxX = (C.float)(value)
}

// GetMinX returns the value in MinX.
func (p GuiTableColumn) GetMinX() float32 {
	ptr := (*C.ImGuiTableColumn)(unsafe.Pointer(p))
	return float32(ptr.MinX)
}

// SetMinX sets the value in MinX.
func (p GuiTableColumn) SetMinX(value float32) {
	ptr := (*C.ImGuiTableColumn)(unsafe.Pointer(p))
	ptr.MinX = (C.float)(value)
}

// GuiTableColumn.NameOffset is unsupported: category unsupported.

// GuiTableColumn.NavLayerCurrent is unsupported: category unsupported.

// GuiTableColumn.NextEnabledColumn is unsupported: category unsupported.

// GuiTableColumn.PrevEnabledColumn is unsupported: category unsupported.

// GuiTableColumn.SortDirection is unsupported: category unsupported.

// GuiTableColumn.SortDirectionsAvailCount is unsupported: category unsupported.

// GetSortDirectionsAvailList returns the value in SortDirectionsAvailList.
func (p GuiTableColumn) GetSortDirectionsAvailList() uint8 {
	ptr := (*C.ImGuiTableColumn)(unsafe.Pointer(p))
	return uint8(ptr.SortDirectionsAvailList)
}

// SetSortDirectionsAvailList sets the value in SortDirectionsAvailList.
func (p GuiTableColumn) SetSortDirectionsAvailList(value uint8) {
	ptr := (*C.ImGuiTableColumn)(unsafe.Pointer(p))
	ptr.SortDirectionsAvailList = (C.uint8_t)(value)
}

// GuiTableColumn.SortDirectionsAvailMask is unsupported: category unsupported.

// GuiTableColumn.SortOrder is unsupported: category unsupported.

// GetStretchWeight returns the value in StretchWeight.
func (p GuiTableColumn) GetStretchWeight() float32 {
	ptr := (*C.ImGuiTableColumn)(unsafe.Pointer(p))
	return float32(ptr.StretchWeight)
}

// SetStretchWeight sets the value in StretchWeight.
func (p GuiTableColumn) SetStretchWeight(value float32) {
	ptr := (*C.ImGuiTableColumn)(unsafe.Pointer(p))
	ptr.StretchWeight = (C.float)(value)
}

// GuiTableColumn.UserID is unsupported: category unsupported.

// GetWidthAuto returns the value in WidthAuto.
func (p GuiTableColumn) GetWidthAuto() float32 {
	ptr := (*C.ImGuiTableColumn)(unsafe.Pointer(p))
	return float32(ptr.WidthAuto)
}

// SetWidthAuto sets the value in WidthAuto.
func (p GuiTableColumn) SetWidthAuto(value float32) {
	ptr := (*C.ImGuiTableColumn)(unsafe.Pointer(p))
	ptr.WidthAuto = (C.float)(value)
}

// GetWidthGiven returns the value in WidthGiven.
func (p GuiTableColumn) GetWidthGiven() float32 {
	ptr := (*C.ImGuiTableColumn)(unsafe.Pointer(p))
	return float32(ptr.WidthGiven)
}

// SetWidthGiven sets the value in WidthGiven.
func (p GuiTableColumn) SetWidthGiven(value float32) {
	ptr := (*C.ImGuiTableColumn)(unsafe.Pointer(p))
	ptr.WidthGiven = (C.float)(value)
}

// GetWidthMax returns the value in WidthMax.
func (p GuiTableColumn) GetWidthMax() float32 {
	ptr := (*C.ImGuiTableColumn)(unsafe.Pointer(p))
	return float32(ptr.WidthMax)
}

// SetWidthMax sets the value in WidthMax.
func (p GuiTableColumn) SetWidthMax(value float32) {
	ptr := (*C.ImGuiTableColumn)(unsafe.Pointer(p))
	ptr.WidthMax = (C.float)(value)
}

// GetWidthRequest returns the value in WidthRequest.
func (p GuiTableColumn) GetWidthRequest() float32 {
	ptr := (*C.ImGuiTableColumn)(unsafe.Pointer(p))
	return float32(ptr.WidthRequest)
}

// SetWidthRequest sets the value in WidthRequest.
func (p GuiTableColumn) SetWidthRequest(value float32) {
	ptr := (*C.ImGuiTableColumn)(unsafe.Pointer(p))
	ptr.WidthRequest = (C.float)(value)
}

// GetWorkMaxX returns the value in WorkMaxX.
func (p GuiTableColumn) GetWorkMaxX() float32 {
	ptr := (*C.ImGuiTableColumn)(unsafe.Pointer(p))
	return float32(ptr.WorkMaxX)
}

// SetWorkMaxX sets the value in WorkMaxX.
func (p GuiTableColumn) SetWorkMaxX(value float32) {
	ptr := (*C.ImGuiTableColumn)(unsafe.Pointer(p))
	ptr.WorkMaxX = (C.float)(value)
}

// GetWorkMinX returns the value in WorkMinX.
func (p GuiTableColumn) GetWorkMinX() float32 {
	ptr := (*C.ImGuiTableColumn)(unsafe.Pointer(p))
	return float32(ptr.WorkMinX)
}

// SetWorkMinX sets the value in WorkMinX.
func (p GuiTableColumn) SetWorkMinX(value float32) {
	ptr := (*C.ImGuiTableColumn)(unsafe.Pointer(p))
	ptr.WorkMinX = (C.float)(value)
}

// GuiTableColumnSettings wraps struct ImGuiTableColumnSettings.
type GuiTableColumnSettings uintptr

// GuiTableColumnSettingsNil is a null pointer.
var GuiTableColumnSettingsNil GuiTableColumnSettings

// GuiTableColumnSettingsSizeOf is the byte size of ImGuiTableColumnSettings.
const GuiTableColumnSettingsSizeOf = int(C.sizeof_ImGuiTableColumnSettings)

// ImGuiTableColumnSettings allocator is suppressed.

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p GuiTableColumnSettings) Offset(offset int) GuiTableColumnSettings {
	return p + GuiTableColumnSettings(offset*GuiTableColumnSettingsSizeOf)
}

// GuiTableColumnSettings.DisplayOrder is unsupported: category unsupported.

// GuiTableColumnSettings.Index is unsupported: category unsupported.

// GuiTableColumnSettings.IsEnabled is unsupported: category unsupported.

// GuiTableColumnSettings.IsStretch is unsupported: category unsupported.

// GuiTableColumnSettings.SortDirection is unsupported: category unsupported.

// GuiTableColumnSettings.SortOrder is unsupported: category unsupported.

// GuiTableColumnSettings.UserID is unsupported: category unsupported.

// GetWidthOrWeight returns the value in WidthOrWeight.
func (p GuiTableColumnSettings) GetWidthOrWeight() float32 {
	ptr := (*C.ImGuiTableColumnSettings)(unsafe.Pointer(p))
	return float32(ptr.WidthOrWeight)
}

// SetWidthOrWeight sets the value in WidthOrWeight.
func (p GuiTableColumnSettings) SetWidthOrWeight(value float32) {
	ptr := (*C.ImGuiTableColumnSettings)(unsafe.Pointer(p))
	ptr.WidthOrWeight = (C.float)(value)
}

// GuiTableColumnSortSpecs wraps struct ImGuiTableColumnSortSpecs.
type GuiTableColumnSortSpecs uintptr

// GuiTableColumnSortSpecsNil is a null pointer.
var GuiTableColumnSortSpecsNil GuiTableColumnSortSpecs

// GuiTableColumnSortSpecsSizeOf is the byte size of ImGuiTableColumnSortSpecs.
const GuiTableColumnSortSpecsSizeOf = int(C.sizeof_ImGuiTableColumnSortSpecs)

// ImGuiTableColumnSortSpecs allocator is suppressed.

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p GuiTableColumnSortSpecs) Offset(offset int) GuiTableColumnSortSpecs {
	return p + GuiTableColumnSortSpecs(offset*GuiTableColumnSortSpecsSizeOf)
}

// GuiTableColumnSortSpecs.ColumnIndex is unsupported: category unsupported.

// GuiTableColumnSortSpecs.ColumnUserID is unsupported: category unsupported.

// GetSortDirection returns the value in SortDirection.
func (p GuiTableColumnSortSpecs) GetSortDirection() GuiSortDirection {
	ptr := (*C.ImGuiTableColumnSortSpecs)(unsafe.Pointer(p))
	return GuiSortDirection(ptr.SortDirection)
}

// SetSortDirection sets the value in SortDirection.
func (p GuiTableColumnSortSpecs) SetSortDirection(value GuiSortDirection) {
	ptr := (*C.ImGuiTableColumnSortSpecs)(unsafe.Pointer(p))
	ptr.SortDirection = (C.ImGuiSortDirection)(value)
}

// GuiTableColumnSortSpecs.SortOrder is unsupported: category unsupported.

// GuiTableHeaderData wraps struct ImGuiTableHeaderData.
type GuiTableHeaderData uintptr

// GuiTableHeaderDataNil is a null pointer.
var GuiTableHeaderDataNil GuiTableHeaderData

// GuiTableHeaderDataSizeOf is the byte size of ImGuiTableHeaderData.
const GuiTableHeaderDataSizeOf = int(C.sizeof_ImGuiTableHeaderData)

// GuiTableHeaderDataAlloc allocates a continuous block of GuiTableHeaderData.
func GuiTableHeaderDataAlloc(alloc ffi.Allocator, count int) GuiTableHeaderData {
	ptr := alloc.Allocate(GuiTableHeaderDataSizeOf * count)
	return GuiTableHeaderData(ptr)
}

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p GuiTableHeaderData) Offset(offset int) GuiTableHeaderData {
	return p + GuiTableHeaderData(offset*GuiTableHeaderDataSizeOf)
}

// GetBgColor0 returns the value in BgColor0.
func (p GuiTableHeaderData) GetBgColor0() uint32 {
	ptr := (*C.ImGuiTableHeaderData)(unsafe.Pointer(p))
	return uint32(ptr.BgColor0)
}

// SetBgColor0 sets the value in BgColor0.
func (p GuiTableHeaderData) SetBgColor0(value uint32) {
	ptr := (*C.ImGuiTableHeaderData)(unsafe.Pointer(p))
	ptr.BgColor0 = (C.uint32_t)(value)
}

// GetBgColor1 returns the value in BgColor1.
func (p GuiTableHeaderData) GetBgColor1() uint32 {
	ptr := (*C.ImGuiTableHeaderData)(unsafe.Pointer(p))
	return uint32(ptr.BgColor1)
}

// SetBgColor1 sets the value in BgColor1.
func (p GuiTableHeaderData) SetBgColor1(value uint32) {
	ptr := (*C.ImGuiTableHeaderData)(unsafe.Pointer(p))
	ptr.BgColor1 = (C.uint32_t)(value)
}

// GuiTableHeaderData.Index is unsupported: category unsupported.

// GetTextColor returns the value in TextColor.
func (p GuiTableHeaderData) GetTextColor() uint32 {
	ptr := (*C.ImGuiTableHeaderData)(unsafe.Pointer(p))
	return uint32(ptr.TextColor)
}

// SetTextColor sets the value in TextColor.
func (p GuiTableHeaderData) SetTextColor(value uint32) {
	ptr := (*C.ImGuiTableHeaderData)(unsafe.Pointer(p))
	ptr.TextColor = (C.uint32_t)(value)
}

// GuiTableInstanceData wraps struct ImGuiTableInstanceData.
type GuiTableInstanceData uintptr

// GuiTableInstanceDataNil is a null pointer.
var GuiTableInstanceDataNil GuiTableInstanceData

// GuiTableInstanceDataSizeOf is the byte size of ImGuiTableInstanceData.
const GuiTableInstanceDataSizeOf = int(C.sizeof_ImGuiTableInstanceData)

// ImGuiTableInstanceData allocator is suppressed.

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p GuiTableInstanceData) Offset(offset int) GuiTableInstanceData {
	return p + GuiTableInstanceData(offset*GuiTableInstanceDataSizeOf)
}

// GetHoveredRowLast returns the value in HoveredRowLast.
func (p GuiTableInstanceData) GetHoveredRowLast() int32 {
	ptr := (*C.ImGuiTableInstanceData)(unsafe.Pointer(p))
	return int32(ptr.HoveredRowLast)
}

// SetHoveredRowLast sets the value in HoveredRowLast.
func (p GuiTableInstanceData) SetHoveredRowLast(value int32) {
	ptr := (*C.ImGuiTableInstanceData)(unsafe.Pointer(p))
	ptr.HoveredRowLast = (C.int32_t)(value)
}

// GetHoveredRowNext returns the value in HoveredRowNext.
func (p GuiTableInstanceData) GetHoveredRowNext() int32 {
	ptr := (*C.ImGuiTableInstanceData)(unsafe.Pointer(p))
	return int32(ptr.HoveredRowNext)
}

// SetHoveredRowNext sets the value in HoveredRowNext.
func (p GuiTableInstanceData) SetHoveredRowNext(value int32) {
	ptr := (*C.ImGuiTableInstanceData)(unsafe.Pointer(p))
	ptr.HoveredRowNext = (C.int32_t)(value)
}

// GetLastFrozenHeight returns the value in LastFrozenHeight.
func (p GuiTableInstanceData) GetLastFrozenHeight() float32 {
	ptr := (*C.ImGuiTableInstanceData)(unsafe.Pointer(p))
	return float32(ptr.LastFrozenHeight)
}

// SetLastFrozenHeight sets the value in LastFrozenHeight.
func (p GuiTableInstanceData) SetLastFrozenHeight(value float32) {
	ptr := (*C.ImGuiTableInstanceData)(unsafe.Pointer(p))
	ptr.LastFrozenHeight = (C.float)(value)
}

// GetLastOuterHeight returns the value in LastOuterHeight.
func (p GuiTableInstanceData) GetLastOuterHeight() float32 {
	ptr := (*C.ImGuiTableInstanceData)(unsafe.Pointer(p))
	return float32(ptr.LastOuterHeight)
}

// SetLastOuterHeight sets the value in LastOuterHeight.
func (p GuiTableInstanceData) SetLastOuterHeight(value float32) {
	ptr := (*C.ImGuiTableInstanceData)(unsafe.Pointer(p))
	ptr.LastOuterHeight = (C.float)(value)
}

// GetLastTopHeadersRowHeight returns the value in LastTopHeadersRowHeight.
func (p GuiTableInstanceData) GetLastTopHeadersRowHeight() float32 {
	ptr := (*C.ImGuiTableInstanceData)(unsafe.Pointer(p))
	return float32(ptr.LastTopHeadersRowHeight)
}

// SetLastTopHeadersRowHeight sets the value in LastTopHeadersRowHeight.
func (p GuiTableInstanceData) SetLastTopHeadersRowHeight(value float32) {
	ptr := (*C.ImGuiTableInstanceData)(unsafe.Pointer(p))
	ptr.LastTopHeadersRowHeight = (C.float)(value)
}

// GuiTableInstanceData.TableInstanceID is unsupported: category unsupported.

// GuiTableSettings wraps struct ImGuiTableSettings.
type GuiTableSettings uintptr

// GuiTableSettingsNil is a null pointer.
var GuiTableSettingsNil GuiTableSettings

// GuiTableSettingsSizeOf is the byte size of ImGuiTableSettings.
const GuiTableSettingsSizeOf = int(C.sizeof_ImGuiTableSettings)

// ImGuiTableSettings allocator is suppressed.

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p GuiTableSettings) Offset(offset int) GuiTableSettings {
	return p + GuiTableSettings(offset*GuiTableSettingsSizeOf)
}

// GuiTableSettings.ColumnsCount is unsupported: category unsupported.

// GuiTableSettings.ColumnsCountMax is unsupported: category unsupported.

// GuiTableSettings.ID is unsupported: category unsupported.

// GetRefScale returns the value in RefScale.
func (p GuiTableSettings) GetRefScale() float32 {
	ptr := (*C.ImGuiTableSettings)(unsafe.Pointer(p))
	return float32(ptr.RefScale)
}

// SetRefScale sets the value in RefScale.
func (p GuiTableSettings) SetRefScale(value float32) {
	ptr := (*C.ImGuiTableSettings)(unsafe.Pointer(p))
	ptr.RefScale = (C.float)(value)
}

// GetSaveFlags returns the value in SaveFlags.
func (p GuiTableSettings) GetSaveFlags() GuiTableFlags {
	ptr := (*C.ImGuiTableSettings)(unsafe.Pointer(p))
	return GuiTableFlags(ptr.SaveFlags)
}

// SetSaveFlags sets the value in SaveFlags.
func (p GuiTableSettings) SetSaveFlags(value GuiTableFlags) {
	ptr := (*C.ImGuiTableSettings)(unsafe.Pointer(p))
	ptr.SaveFlags = (C.ImGuiTableFlags)(value)
}

// GetWantApply returns the value in WantApply.
func (p GuiTableSettings) GetWantApply() bool {
	ptr := (*C.ImGuiTableSettings)(unsafe.Pointer(p))
	return bool(ptr.WantApply)
}

// SetWantApply sets the value in WantApply.
func (p GuiTableSettings) SetWantApply(value bool) {
	ptr := (*C.ImGuiTableSettings)(unsafe.Pointer(p))
	ptr.WantApply = (C.bool)(value)
}

// GuiTableSortSpecs wraps struct ImGuiTableSortSpecs.
type GuiTableSortSpecs uintptr

// GuiTableSortSpecsNil is a null pointer.
var GuiTableSortSpecsNil GuiTableSortSpecs

// GuiTableSortSpecsSizeOf is the byte size of ImGuiTableSortSpecs.
const GuiTableSortSpecsSizeOf = int(C.sizeof_ImGuiTableSortSpecs)

// ImGuiTableSortSpecs allocator is suppressed.

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p GuiTableSortSpecs) Offset(offset int) GuiTableSortSpecs {
	return p + GuiTableSortSpecs(offset*GuiTableSortSpecsSizeOf)
}

// GetSpecs returns the value in Specs.
func (p GuiTableSortSpecs) GetSpecs() GuiTableColumnSortSpecs {
	ptr := (*C.ImGuiTableSortSpecs)(unsafe.Pointer(p))
	return GuiTableColumnSortSpecs(unsafe.Pointer(ptr.Specs))
}

// SetSpecs sets the value in Specs.
func (p GuiTableSortSpecs) SetSpecs(value GuiTableColumnSortSpecs) {
	ptr := (*C.ImGuiTableSortSpecs)(unsafe.Pointer(p))
	ptr.Specs = (*C.ImGuiTableColumnSortSpecs)(unsafe.Pointer(value))
}

// GetSpecsCount returns the value in SpecsCount.
func (p GuiTableSortSpecs) GetSpecsCount() int32 {
	ptr := (*C.ImGuiTableSortSpecs)(unsafe.Pointer(p))
	return int32(ptr.SpecsCount)
}

// SetSpecsCount sets the value in SpecsCount.
func (p GuiTableSortSpecs) SetSpecsCount(value int32) {
	ptr := (*C.ImGuiTableSortSpecs)(unsafe.Pointer(p))
	ptr.SpecsCount = (C.int32_t)(value)
}

// GetSpecsDirty returns the value in SpecsDirty.
func (p GuiTableSortSpecs) GetSpecsDirty() bool {
	ptr := (*C.ImGuiTableSortSpecs)(unsafe.Pointer(p))
	return bool(ptr.SpecsDirty)
}

// SetSpecsDirty sets the value in SpecsDirty.
func (p GuiTableSortSpecs) SetSpecsDirty(value bool) {
	ptr := (*C.ImGuiTableSortSpecs)(unsafe.Pointer(p))
	ptr.SpecsDirty = (C.bool)(value)
}

// GuiTableTempData wraps struct ImGuiTableTempData.
type GuiTableTempData uintptr

// GuiTableTempDataNil is a null pointer.
var GuiTableTempDataNil GuiTableTempData

// GuiTableTempDataSizeOf is the byte size of ImGuiTableTempData.
const GuiTableTempDataSizeOf = int(C.sizeof_ImGuiTableTempData)

// ImGuiTableTempData allocator is suppressed.

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p GuiTableTempData) Offset(offset int) GuiTableTempData {
	return p + GuiTableTempData(offset*GuiTableTempDataSizeOf)
}

// GetAngledHeadersExtraWidth returns the value in AngledHeadersExtraWidth.
func (p GuiTableTempData) GetAngledHeadersExtraWidth() float32 {
	ptr := (*C.ImGuiTableTempData)(unsafe.Pointer(p))
	return float32(ptr.AngledHeadersExtraWidth)
}

// SetAngledHeadersExtraWidth sets the value in AngledHeadersExtraWidth.
func (p GuiTableTempData) SetAngledHeadersExtraWidth(value float32) {
	ptr := (*C.ImGuiTableTempData)(unsafe.Pointer(p))
	ptr.AngledHeadersExtraWidth = (C.float)(value)
}

// RefAngledHeadersRequests returns pointer to the AngledHeadersRequests field.
func (p GuiTableTempData) RefAngledHeadersRequests() Vector_ImGuiTableHeaderData {
	return Vector_ImGuiTableHeaderData(p + GuiTableTempData(C.offsetof_ImGuiTableTempData_AngledHeadersRequests))
}

// RefDrawSplitter returns pointer to the DrawSplitter field.
func (p GuiTableTempData) RefDrawSplitter() DrawListSplitter {
	return DrawListSplitter(p + GuiTableTempData(C.offsetof_ImGuiTableTempData_DrawSplitter))
}

// RefHostBackupColumnsOffset returns pointer to the HostBackupColumnsOffset field.
func (p GuiTableTempData) RefHostBackupColumnsOffset() Vec1 {
	return Vec1(p + GuiTableTempData(C.offsetof_ImGuiTableTempData_HostBackupColumnsOffset))
}

// RefHostBackupCurrLineSize returns pointer to the HostBackupCurrLineSize field.
func (p GuiTableTempData) RefHostBackupCurrLineSize() Vec2 {
	return Vec2(p + GuiTableTempData(C.offsetof_ImGuiTableTempData_HostBackupCurrLineSize))
}

// RefHostBackupCursorMaxPos returns pointer to the HostBackupCursorMaxPos field.
func (p GuiTableTempData) RefHostBackupCursorMaxPos() Vec2 {
	return Vec2(p + GuiTableTempData(C.offsetof_ImGuiTableTempData_HostBackupCursorMaxPos))
}

// GetHostBackupItemWidth returns the value in HostBackupItemWidth.
func (p GuiTableTempData) GetHostBackupItemWidth() float32 {
	ptr := (*C.ImGuiTableTempData)(unsafe.Pointer(p))
	return float32(ptr.HostBackupItemWidth)
}

// SetHostBackupItemWidth sets the value in HostBackupItemWidth.
func (p GuiTableTempData) SetHostBackupItemWidth(value float32) {
	ptr := (*C.ImGuiTableTempData)(unsafe.Pointer(p))
	ptr.HostBackupItemWidth = (C.float)(value)
}

// GetHostBackupItemWidthStackSize returns the value in HostBackupItemWidthStackSize.
func (p GuiTableTempData) GetHostBackupItemWidthStackSize() int32 {
	ptr := (*C.ImGuiTableTempData)(unsafe.Pointer(p))
	return int32(ptr.HostBackupItemWidthStackSize)
}

// SetHostBackupItemWidthStackSize sets the value in HostBackupItemWidthStackSize.
func (p GuiTableTempData) SetHostBackupItemWidthStackSize(value int32) {
	ptr := (*C.ImGuiTableTempData)(unsafe.Pointer(p))
	ptr.HostBackupItemWidthStackSize = (C.int32_t)(value)
}

// RefHostBackupParentWorkRect returns pointer to the HostBackupParentWorkRect field.
func (p GuiTableTempData) RefHostBackupParentWorkRect() Rect {
	return Rect(p + GuiTableTempData(C.offsetof_ImGuiTableTempData_HostBackupParentWorkRect))
}

// RefHostBackupPrevLineSize returns pointer to the HostBackupPrevLineSize field.
func (p GuiTableTempData) RefHostBackupPrevLineSize() Vec2 {
	return Vec2(p + GuiTableTempData(C.offsetof_ImGuiTableTempData_HostBackupPrevLineSize))
}

// RefHostBackupWorkRect returns pointer to the HostBackupWorkRect field.
func (p GuiTableTempData) RefHostBackupWorkRect() Rect {
	return Rect(p + GuiTableTempData(C.offsetof_ImGuiTableTempData_HostBackupWorkRect))
}

// GetLastTimeActive returns the value in LastTimeActive.
func (p GuiTableTempData) GetLastTimeActive() float32 {
	ptr := (*C.ImGuiTableTempData)(unsafe.Pointer(p))
	return float32(ptr.LastTimeActive)
}

// SetLastTimeActive sets the value in LastTimeActive.
func (p GuiTableTempData) SetLastTimeActive(value float32) {
	ptr := (*C.ImGuiTableTempData)(unsafe.Pointer(p))
	ptr.LastTimeActive = (C.float)(value)
}

// GetTableIndex returns the value in TableIndex.
func (p GuiTableTempData) GetTableIndex() int32 {
	ptr := (*C.ImGuiTableTempData)(unsafe.Pointer(p))
	return int32(ptr.TableIndex)
}

// SetTableIndex sets the value in TableIndex.
func (p GuiTableTempData) SetTableIndex(value int32) {
	ptr := (*C.ImGuiTableTempData)(unsafe.Pointer(p))
	ptr.TableIndex = (C.int32_t)(value)
}

// RefUserOuterSize returns pointer to the UserOuterSize field.
func (p GuiTableTempData) RefUserOuterSize() Vec2 {
	return Vec2(p + GuiTableTempData(C.offsetof_ImGuiTableTempData_UserOuterSize))
}

// GuiTextBuffer wraps struct ImGuiTextBuffer.
type GuiTextBuffer uintptr

// GuiTextBufferNil is a null pointer.
var GuiTextBufferNil GuiTextBuffer

// GuiTextBufferSizeOf is the byte size of ImGuiTextBuffer.
const GuiTextBufferSizeOf = int(C.sizeof_ImGuiTextBuffer)

// ImGuiTextBuffer allocator is suppressed.

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p GuiTextBuffer) Offset(offset int) GuiTextBuffer {
	return p + GuiTextBuffer(offset*GuiTextBufferSizeOf)
}

// RefBuf returns pointer to the Buf field.
func (p GuiTextBuffer) RefBuf() Vector_char {
	return Vector_char(p + GuiTextBuffer(C.offsetof_ImGuiTextBuffer_Buf))
}

// GuiTextFilter wraps struct ImGuiTextFilter.
type GuiTextFilter uintptr

// GuiTextFilterNil is a null pointer.
var GuiTextFilterNil GuiTextFilter

// GuiTextFilterSizeOf is the byte size of ImGuiTextFilter.
const GuiTextFilterSizeOf = int(C.sizeof_ImGuiTextFilter)

// ImGuiTextFilter allocator is suppressed.

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p GuiTextFilter) Offset(offset int) GuiTextFilter {
	return p + GuiTextFilter(offset*GuiTextFilterSizeOf)
}

// GetCountGrep returns the value in CountGrep.
func (p GuiTextFilter) GetCountGrep() int32 {
	ptr := (*C.ImGuiTextFilter)(unsafe.Pointer(p))
	return int32(ptr.CountGrep)
}

// SetCountGrep sets the value in CountGrep.
func (p GuiTextFilter) SetCountGrep(value int32) {
	ptr := (*C.ImGuiTextFilter)(unsafe.Pointer(p))
	ptr.CountGrep = (C.int32_t)(value)
}

// RefFilters returns pointer to the Filters field.
func (p GuiTextFilter) RefFilters() Vector_ImGuiTextRange {
	return Vector_ImGuiTextRange(p + GuiTextFilter(C.offsetof_ImGuiTextFilter_Filters))
}

// GuiTextFilter.InputBuf[256] is unsupported: category unsupported.

// GuiTextIndex wraps struct ImGuiTextIndex.
type GuiTextIndex uintptr

// GuiTextIndexNil is a null pointer.
var GuiTextIndexNil GuiTextIndex

// GuiTextIndexSizeOf is the byte size of ImGuiTextIndex.
const GuiTextIndexSizeOf = int(C.sizeof_ImGuiTextIndex)

// GuiTextIndexAlloc allocates a continuous block of GuiTextIndex.
func GuiTextIndexAlloc(alloc ffi.Allocator, count int) GuiTextIndex {
	ptr := alloc.Allocate(GuiTextIndexSizeOf * count)
	return GuiTextIndex(ptr)
}

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p GuiTextIndex) Offset(offset int) GuiTextIndex {
	return p + GuiTextIndex(offset*GuiTextIndexSizeOf)
}

// GetEndOffset returns the value in EndOffset.
func (p GuiTextIndex) GetEndOffset() int32 {
	ptr := (*C.ImGuiTextIndex)(unsafe.Pointer(p))
	return int32(ptr.EndOffset)
}

// SetEndOffset sets the value in EndOffset.
func (p GuiTextIndex) SetEndOffset(value int32) {
	ptr := (*C.ImGuiTextIndex)(unsafe.Pointer(p))
	ptr.EndOffset = (C.int32_t)(value)
}

// RefOffsets returns pointer to the Offsets field.
func (p GuiTextIndex) RefOffsets() Vector_int {
	return Vector_int(p + GuiTextIndex(C.offsetof_ImGuiTextIndex_Offsets))
}

// GuiTextRange wraps struct ImGuiTextRange.
type GuiTextRange uintptr

// GuiTextRangeNil is a null pointer.
var GuiTextRangeNil GuiTextRange

// GuiTextRangeSizeOf is the byte size of ImGuiTextRange.
const GuiTextRangeSizeOf = int(C.sizeof_ImGuiTextRange)

// ImGuiTextRange allocator is suppressed.

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p GuiTextRange) Offset(offset int) GuiTextRange {
	return p + GuiTextRange(offset*GuiTextRangeSizeOf)
}

// GetB returns the value in b.
func (p GuiTextRange) GetB() ffi.CString {
	ptr := (*C.ImGuiTextRange)(unsafe.Pointer(p))
	return ffi.CStringFromPtr((unsafe.Pointer)(ptr.b))
}

// SetB sets the value in b.
func (p GuiTextRange) SetB(value ffi.CString) {
	ptr := (*C.ImGuiTextRange)(unsafe.Pointer(p))
	ptr.b = (*C.char)(value.Raw())
}

// GetE returns the value in e.
func (p GuiTextRange) GetE() ffi.CString {
	ptr := (*C.ImGuiTextRange)(unsafe.Pointer(p))
	return ffi.CStringFromPtr((unsafe.Pointer)(ptr.e))
}

// SetE sets the value in e.
func (p GuiTextRange) SetE(value ffi.CString) {
	ptr := (*C.ImGuiTextRange)(unsafe.Pointer(p))
	ptr.e = (*C.char)(value.Raw())
}

// GuiTreeNodeStackData wraps struct ImGuiTreeNodeStackData.
type GuiTreeNodeStackData uintptr

// GuiTreeNodeStackDataNil is a null pointer.
var GuiTreeNodeStackDataNil GuiTreeNodeStackData

// GuiTreeNodeStackDataSizeOf is the byte size of ImGuiTreeNodeStackData.
const GuiTreeNodeStackDataSizeOf = int(C.sizeof_ImGuiTreeNodeStackData)

// GuiTreeNodeStackDataAlloc allocates a continuous block of GuiTreeNodeStackData.
func GuiTreeNodeStackDataAlloc(alloc ffi.Allocator, count int) GuiTreeNodeStackData {
	ptr := alloc.Allocate(GuiTreeNodeStackDataSizeOf * count)
	return GuiTreeNodeStackData(ptr)
}

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p GuiTreeNodeStackData) Offset(offset int) GuiTreeNodeStackData {
	return p + GuiTreeNodeStackData(offset*GuiTreeNodeStackDataSizeOf)
}

// GuiTreeNodeStackData.DrawLinesTableColumn is unsupported: category unsupported.

// GetDrawLinesToNodesY2 returns the value in DrawLinesToNodesY2.
func (p GuiTreeNodeStackData) GetDrawLinesToNodesY2() float32 {
	ptr := (*C.ImGuiTreeNodeStackData)(unsafe.Pointer(p))
	return float32(ptr.DrawLinesToNodesY2)
}

// SetDrawLinesToNodesY2 sets the value in DrawLinesToNodesY2.
func (p GuiTreeNodeStackData) SetDrawLinesToNodesY2(value float32) {
	ptr := (*C.ImGuiTreeNodeStackData)(unsafe.Pointer(p))
	ptr.DrawLinesToNodesY2 = (C.float)(value)
}

// GetDrawLinesX1 returns the value in DrawLinesX1.
func (p GuiTreeNodeStackData) GetDrawLinesX1() float32 {
	ptr := (*C.ImGuiTreeNodeStackData)(unsafe.Pointer(p))
	return float32(ptr.DrawLinesX1)
}

// SetDrawLinesX1 sets the value in DrawLinesX1.
func (p GuiTreeNodeStackData) SetDrawLinesX1(value float32) {
	ptr := (*C.ImGuiTreeNodeStackData)(unsafe.Pointer(p))
	ptr.DrawLinesX1 = (C.float)(value)
}

// GuiTreeNodeStackData.ID is unsupported: category unsupported.

// GetItemFlags returns the value in ItemFlags.
func (p GuiTreeNodeStackData) GetItemFlags() GuiItemFlags {
	ptr := (*C.ImGuiTreeNodeStackData)(unsafe.Pointer(p))
	return GuiItemFlags(ptr.ItemFlags)
}

// SetItemFlags sets the value in ItemFlags.
func (p GuiTreeNodeStackData) SetItemFlags(value GuiItemFlags) {
	ptr := (*C.ImGuiTreeNodeStackData)(unsafe.Pointer(p))
	ptr.ItemFlags = (C.ImGuiItemFlags)(value)
}

// RefNavRect returns pointer to the NavRect field.
func (p GuiTreeNodeStackData) RefNavRect() Rect {
	return Rect(p + GuiTreeNodeStackData(C.offsetof_ImGuiTreeNodeStackData_NavRect))
}

// GetTreeFlags returns the value in TreeFlags.
func (p GuiTreeNodeStackData) GetTreeFlags() GuiTreeNodeFlags {
	ptr := (*C.ImGuiTreeNodeStackData)(unsafe.Pointer(p))
	return GuiTreeNodeFlags(ptr.TreeFlags)
}

// SetTreeFlags sets the value in TreeFlags.
func (p GuiTreeNodeStackData) SetTreeFlags(value GuiTreeNodeFlags) {
	ptr := (*C.ImGuiTreeNodeStackData)(unsafe.Pointer(p))
	ptr.TreeFlags = (C.ImGuiTreeNodeFlags)(value)
}

// GuiTypingSelectRequest wraps struct ImGuiTypingSelectRequest.
type GuiTypingSelectRequest uintptr

// GuiTypingSelectRequestNil is a null pointer.
var GuiTypingSelectRequestNil GuiTypingSelectRequest

// GuiTypingSelectRequestSizeOf is the byte size of ImGuiTypingSelectRequest.
const GuiTypingSelectRequestSizeOf = int(C.sizeof_ImGuiTypingSelectRequest)

// GuiTypingSelectRequestAlloc allocates a continuous block of GuiTypingSelectRequest.
func GuiTypingSelectRequestAlloc(alloc ffi.Allocator, count int) GuiTypingSelectRequest {
	ptr := alloc.Allocate(GuiTypingSelectRequestSizeOf * count)
	return GuiTypingSelectRequest(ptr)
}

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p GuiTypingSelectRequest) Offset(offset int) GuiTypingSelectRequest {
	return p + GuiTypingSelectRequest(offset*GuiTypingSelectRequestSizeOf)
}

// GetFlags returns the value in Flags.
func (p GuiTypingSelectRequest) GetFlags() GuiTypingSelectFlags {
	ptr := (*C.ImGuiTypingSelectRequest)(unsafe.Pointer(p))
	return GuiTypingSelectFlags(ptr.Flags)
}

// SetFlags sets the value in Flags.
func (p GuiTypingSelectRequest) SetFlags(value GuiTypingSelectFlags) {
	ptr := (*C.ImGuiTypingSelectRequest)(unsafe.Pointer(p))
	ptr.Flags = (C.ImGuiTypingSelectFlags)(value)
}

// GetSearchBuffer returns the value in SearchBuffer.
func (p GuiTypingSelectRequest) GetSearchBuffer() ffi.CString {
	ptr := (*C.ImGuiTypingSelectRequest)(unsafe.Pointer(p))
	return ffi.CStringFromPtr((unsafe.Pointer)(ptr.SearchBuffer))
}

// SetSearchBuffer sets the value in SearchBuffer.
func (p GuiTypingSelectRequest) SetSearchBuffer(value ffi.CString) {
	ptr := (*C.ImGuiTypingSelectRequest)(unsafe.Pointer(p))
	ptr.SearchBuffer = (*C.char)(value.Raw())
}

// GetSearchBufferLen returns the value in SearchBufferLen.
func (p GuiTypingSelectRequest) GetSearchBufferLen() int32 {
	ptr := (*C.ImGuiTypingSelectRequest)(unsafe.Pointer(p))
	return int32(ptr.SearchBufferLen)
}

// SetSearchBufferLen sets the value in SearchBufferLen.
func (p GuiTypingSelectRequest) SetSearchBufferLen(value int32) {
	ptr := (*C.ImGuiTypingSelectRequest)(unsafe.Pointer(p))
	ptr.SearchBufferLen = (C.int32_t)(value)
}

// GetSelectRequest returns the value in SelectRequest.
func (p GuiTypingSelectRequest) GetSelectRequest() bool {
	ptr := (*C.ImGuiTypingSelectRequest)(unsafe.Pointer(p))
	return bool(ptr.SelectRequest)
}

// SetSelectRequest sets the value in SelectRequest.
func (p GuiTypingSelectRequest) SetSelectRequest(value bool) {
	ptr := (*C.ImGuiTypingSelectRequest)(unsafe.Pointer(p))
	ptr.SelectRequest = (C.bool)(value)
}

// GetSingleCharMode returns the value in SingleCharMode.
func (p GuiTypingSelectRequest) GetSingleCharMode() bool {
	ptr := (*C.ImGuiTypingSelectRequest)(unsafe.Pointer(p))
	return bool(ptr.SingleCharMode)
}

// SetSingleCharMode sets the value in SingleCharMode.
func (p GuiTypingSelectRequest) SetSingleCharMode(value bool) {
	ptr := (*C.ImGuiTypingSelectRequest)(unsafe.Pointer(p))
	ptr.SingleCharMode = (C.bool)(value)
}

// GuiTypingSelectRequest.SingleCharSize is unsupported: category unsupported.

// GuiTypingSelectState wraps struct ImGuiTypingSelectState.
type GuiTypingSelectState uintptr

// GuiTypingSelectStateNil is a null pointer.
var GuiTypingSelectStateNil GuiTypingSelectState

// GuiTypingSelectStateSizeOf is the byte size of ImGuiTypingSelectState.
const GuiTypingSelectStateSizeOf = int(C.sizeof_ImGuiTypingSelectState)

// ImGuiTypingSelectState allocator is suppressed.

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p GuiTypingSelectState) Offset(offset int) GuiTypingSelectState {
	return p + GuiTypingSelectState(offset*GuiTypingSelectStateSizeOf)
}

// GuiTypingSelectState.FocusScope is unsupported: category unsupported.

// GetLastRequestFrame returns the value in LastRequestFrame.
func (p GuiTypingSelectState) GetLastRequestFrame() int32 {
	ptr := (*C.ImGuiTypingSelectState)(unsafe.Pointer(p))
	return int32(ptr.LastRequestFrame)
}

// SetLastRequestFrame sets the value in LastRequestFrame.
func (p GuiTypingSelectState) SetLastRequestFrame(value int32) {
	ptr := (*C.ImGuiTypingSelectState)(unsafe.Pointer(p))
	ptr.LastRequestFrame = (C.int32_t)(value)
}

// GetLastRequestTime returns the value in LastRequestTime.
func (p GuiTypingSelectState) GetLastRequestTime() float32 {
	ptr := (*C.ImGuiTypingSelectState)(unsafe.Pointer(p))
	return float32(ptr.LastRequestTime)
}

// SetLastRequestTime sets the value in LastRequestTime.
func (p GuiTypingSelectState) SetLastRequestTime(value float32) {
	ptr := (*C.ImGuiTypingSelectState)(unsafe.Pointer(p))
	ptr.LastRequestTime = (C.float)(value)
}

// RefRequest returns pointer to the Request field.
func (p GuiTypingSelectState) RefRequest() GuiTypingSelectRequest {
	return GuiTypingSelectRequest(p + GuiTypingSelectState(C.offsetof_ImGuiTypingSelectState_Request))
}

// GuiTypingSelectState.SearchBuffer[64] is unsupported: category unsupported.

// GetSingleCharModeLock returns the value in SingleCharModeLock.
func (p GuiTypingSelectState) GetSingleCharModeLock() bool {
	ptr := (*C.ImGuiTypingSelectState)(unsafe.Pointer(p))
	return bool(ptr.SingleCharModeLock)
}

// SetSingleCharModeLock sets the value in SingleCharModeLock.
func (p GuiTypingSelectState) SetSingleCharModeLock(value bool) {
	ptr := (*C.ImGuiTypingSelectState)(unsafe.Pointer(p))
	ptr.SingleCharModeLock = (C.bool)(value)
}

// GuiViewport wraps struct ImGuiViewport.
type GuiViewport uintptr

// GuiViewportNil is a null pointer.
var GuiViewportNil GuiViewport

// GuiViewportSizeOf is the byte size of ImGuiViewport.
const GuiViewportSizeOf = int(C.sizeof_ImGuiViewport)

// ImGuiViewport allocator is suppressed.

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p GuiViewport) Offset(offset int) GuiViewport {
	return p + GuiViewport(offset*GuiViewportSizeOf)
}

// GetDpiScale returns the value in DpiScale.
func (p GuiViewport) GetDpiScale() float32 {
	ptr := (*C.ImGuiViewport)(unsafe.Pointer(p))
	return float32(ptr.DpiScale)
}

// SetDpiScale sets the value in DpiScale.
func (p GuiViewport) SetDpiScale(value float32) {
	ptr := (*C.ImGuiViewport)(unsafe.Pointer(p))
	ptr.DpiScale = (C.float)(value)
}

// GetDrawData returns the value in DrawData.
func (p GuiViewport) GetDrawData() DrawData {
	ptr := (*C.ImGuiViewport)(unsafe.Pointer(p))
	return DrawData(unsafe.Pointer(ptr.DrawData))
}

// SetDrawData sets the value in DrawData.
func (p GuiViewport) SetDrawData(value DrawData) {
	ptr := (*C.ImGuiViewport)(unsafe.Pointer(p))
	ptr.DrawData = (*C.ImDrawData)(unsafe.Pointer(value))
}

// GetFlags returns the value in Flags.
func (p GuiViewport) GetFlags() GuiViewportFlags {
	ptr := (*C.ImGuiViewport)(unsafe.Pointer(p))
	return GuiViewportFlags(ptr.Flags)
}

// SetFlags sets the value in Flags.
func (p GuiViewport) SetFlags(value GuiViewportFlags) {
	ptr := (*C.ImGuiViewport)(unsafe.Pointer(p))
	ptr.Flags = (C.ImGuiViewportFlags)(value)
}

// RefFramebufferScale returns pointer to the FramebufferScale field.
func (p GuiViewport) RefFramebufferScale() Vec2 {
	return Vec2(p + GuiViewport(C.offsetof_ImGuiViewport_FramebufferScale))
}

// GuiViewport.ID is unsupported: category unsupported.

// GetParentViewport returns the value in ParentViewport.
func (p GuiViewport) GetParentViewport() GuiViewport {
	ptr := (*C.ImGuiViewport)(unsafe.Pointer(p))
	return GuiViewport(unsafe.Pointer(ptr.ParentViewport))
}

// SetParentViewport sets the value in ParentViewport.
func (p GuiViewport) SetParentViewport(value GuiViewport) {
	ptr := (*C.ImGuiViewport)(unsafe.Pointer(p))
	ptr.ParentViewport = (*C.ImGuiViewport)(unsafe.Pointer(value))
}

// GuiViewport.ParentViewportId is unsupported: category unsupported.

// GetPlatformHandle returns the value in PlatformHandle.
func (p GuiViewport) GetPlatformHandle() uintptr {
	ptr := (*C.ImGuiViewport)(unsafe.Pointer(p))
	return uintptr(unsafe.Pointer(ptr.PlatformHandle))
}

// SetPlatformHandle sets the value in PlatformHandle.
func (p GuiViewport) SetPlatformHandle(value uintptr) {
	ptr := (*C.ImGuiViewport)(unsafe.Pointer(p))
	ptr.PlatformHandle = unsafe.Pointer(value)
}

// GetPlatformHandleRaw returns the value in PlatformHandleRaw.
func (p GuiViewport) GetPlatformHandleRaw() uintptr {
	ptr := (*C.ImGuiViewport)(unsafe.Pointer(p))
	return uintptr(unsafe.Pointer(ptr.PlatformHandleRaw))
}

// SetPlatformHandleRaw sets the value in PlatformHandleRaw.
func (p GuiViewport) SetPlatformHandleRaw(value uintptr) {
	ptr := (*C.ImGuiViewport)(unsafe.Pointer(p))
	ptr.PlatformHandleRaw = unsafe.Pointer(value)
}

// GetPlatformRequestClose returns the value in PlatformRequestClose.
func (p GuiViewport) GetPlatformRequestClose() bool {
	ptr := (*C.ImGuiViewport)(unsafe.Pointer(p))
	return bool(ptr.PlatformRequestClose)
}

// SetPlatformRequestClose sets the value in PlatformRequestClose.
func (p GuiViewport) SetPlatformRequestClose(value bool) {
	ptr := (*C.ImGuiViewport)(unsafe.Pointer(p))
	ptr.PlatformRequestClose = (C.bool)(value)
}

// GetPlatformRequestMove returns the value in PlatformRequestMove.
func (p GuiViewport) GetPlatformRequestMove() bool {
	ptr := (*C.ImGuiViewport)(unsafe.Pointer(p))
	return bool(ptr.PlatformRequestMove)
}

// SetPlatformRequestMove sets the value in PlatformRequestMove.
func (p GuiViewport) SetPlatformRequestMove(value bool) {
	ptr := (*C.ImGuiViewport)(unsafe.Pointer(p))
	ptr.PlatformRequestMove = (C.bool)(value)
}

// GetPlatformRequestResize returns the value in PlatformRequestResize.
func (p GuiViewport) GetPlatformRequestResize() bool {
	ptr := (*C.ImGuiViewport)(unsafe.Pointer(p))
	return bool(ptr.PlatformRequestResize)
}

// SetPlatformRequestResize sets the value in PlatformRequestResize.
func (p GuiViewport) SetPlatformRequestResize(value bool) {
	ptr := (*C.ImGuiViewport)(unsafe.Pointer(p))
	ptr.PlatformRequestResize = (C.bool)(value)
}

// GetPlatformUserData returns the value in PlatformUserData.
func (p GuiViewport) GetPlatformUserData() uintptr {
	ptr := (*C.ImGuiViewport)(unsafe.Pointer(p))
	return uintptr(unsafe.Pointer(ptr.PlatformUserData))
}

// SetPlatformUserData sets the value in PlatformUserData.
func (p GuiViewport) SetPlatformUserData(value uintptr) {
	ptr := (*C.ImGuiViewport)(unsafe.Pointer(p))
	ptr.PlatformUserData = unsafe.Pointer(value)
}

// GetPlatformWindowCreated returns the value in PlatformWindowCreated.
func (p GuiViewport) GetPlatformWindowCreated() bool {
	ptr := (*C.ImGuiViewport)(unsafe.Pointer(p))
	return bool(ptr.PlatformWindowCreated)
}

// SetPlatformWindowCreated sets the value in PlatformWindowCreated.
func (p GuiViewport) SetPlatformWindowCreated(value bool) {
	ptr := (*C.ImGuiViewport)(unsafe.Pointer(p))
	ptr.PlatformWindowCreated = (C.bool)(value)
}

// RefPos returns pointer to the Pos field.
func (p GuiViewport) RefPos() Vec2 {
	return Vec2(p + GuiViewport(C.offsetof_ImGuiViewport_Pos))
}

// GetRendererUserData returns the value in RendererUserData.
func (p GuiViewport) GetRendererUserData() uintptr {
	ptr := (*C.ImGuiViewport)(unsafe.Pointer(p))
	return uintptr(unsafe.Pointer(ptr.RendererUserData))
}

// SetRendererUserData sets the value in RendererUserData.
func (p GuiViewport) SetRendererUserData(value uintptr) {
	ptr := (*C.ImGuiViewport)(unsafe.Pointer(p))
	ptr.RendererUserData = unsafe.Pointer(value)
}

// RefSize returns pointer to the Size field.
func (p GuiViewport) RefSize() Vec2 {
	return Vec2(p + GuiViewport(C.offsetof_ImGuiViewport_Size))
}

// RefWorkPos returns pointer to the WorkPos field.
func (p GuiViewport) RefWorkPos() Vec2 {
	return Vec2(p + GuiViewport(C.offsetof_ImGuiViewport_WorkPos))
}

// RefWorkSize returns pointer to the WorkSize field.
func (p GuiViewport) RefWorkSize() Vec2 {
	return Vec2(p + GuiViewport(C.offsetof_ImGuiViewport_WorkSize))
}

// GuiViewportP wraps struct ImGuiViewportP.
type GuiViewportP uintptr

// GuiViewportPNil is a null pointer.
var GuiViewportPNil GuiViewportP

// GuiViewportPSizeOf is the byte size of ImGuiViewportP.
const GuiViewportPSizeOf = int(C.sizeof_ImGuiViewportP)

// ImGuiViewportP allocator is suppressed.

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p GuiViewportP) Offset(offset int) GuiViewportP {
	return p + GuiViewportP(offset*GuiViewportPSizeOf)
}

// GetAlpha returns the value in Alpha.
func (p GuiViewportP) GetAlpha() float32 {
	ptr := (*C.ImGuiViewportP)(unsafe.Pointer(p))
	return float32(ptr.Alpha)
}

// SetAlpha sets the value in Alpha.
func (p GuiViewportP) SetAlpha(value float32) {
	ptr := (*C.ImGuiViewportP)(unsafe.Pointer(p))
	ptr.Alpha = (C.float)(value)
}

// GuiViewportP.BgFgDrawListsLastFrame[2] is unsupported: category unsupported.

// GuiViewportP.BgFgDrawLists[2] is unsupported: category unsupported.

// RefBuildWorkInsetMax returns pointer to the BuildWorkInsetMax field.
func (p GuiViewportP) RefBuildWorkInsetMax() Vec2 {
	return Vec2(p + GuiViewportP(C.offsetof_ImGuiViewportP_BuildWorkInsetMax))
}

// RefBuildWorkInsetMin returns pointer to the BuildWorkInsetMin field.
func (p GuiViewportP) RefBuildWorkInsetMin() Vec2 {
	return Vec2(p + GuiViewportP(C.offsetof_ImGuiViewportP_BuildWorkInsetMin))
}

// RefDrawDataBuilder returns pointer to the DrawDataBuilder field.
func (p GuiViewportP) RefDrawDataBuilder() DrawDataBuilder {
	return DrawDataBuilder(p + GuiViewportP(C.offsetof_ImGuiViewportP_DrawDataBuilder))
}

// RefDrawDataP returns pointer to the DrawDataP field.
func (p GuiViewportP) RefDrawDataP() DrawData {
	return DrawData(p + GuiViewportP(C.offsetof_ImGuiViewportP_DrawDataP))
}

// GetIdx returns the value in Idx.
func (p GuiViewportP) GetIdx() int32 {
	ptr := (*C.ImGuiViewportP)(unsafe.Pointer(p))
	return int32(ptr.Idx)
}

// SetIdx sets the value in Idx.
func (p GuiViewportP) SetIdx(value int32) {
	ptr := (*C.ImGuiViewportP)(unsafe.Pointer(p))
	ptr.Idx = (C.int32_t)(value)
}

// GetLastAlpha returns the value in LastAlpha.
func (p GuiViewportP) GetLastAlpha() float32 {
	ptr := (*C.ImGuiViewportP)(unsafe.Pointer(p))
	return float32(ptr.LastAlpha)
}

// SetLastAlpha sets the value in LastAlpha.
func (p GuiViewportP) SetLastAlpha(value float32) {
	ptr := (*C.ImGuiViewportP)(unsafe.Pointer(p))
	ptr.LastAlpha = (C.float)(value)
}

// GetLastFocusedHadNavWindow returns the value in LastFocusedHadNavWindow.
func (p GuiViewportP) GetLastFocusedHadNavWindow() bool {
	ptr := (*C.ImGuiViewportP)(unsafe.Pointer(p))
	return bool(ptr.LastFocusedHadNavWindow)
}

// SetLastFocusedHadNavWindow sets the value in LastFocusedHadNavWindow.
func (p GuiViewportP) SetLastFocusedHadNavWindow(value bool) {
	ptr := (*C.ImGuiViewportP)(unsafe.Pointer(p))
	ptr.LastFocusedHadNavWindow = (C.bool)(value)
}

// GetLastFocusedStampCount returns the value in LastFocusedStampCount.
func (p GuiViewportP) GetLastFocusedStampCount() int32 {
	ptr := (*C.ImGuiViewportP)(unsafe.Pointer(p))
	return int32(ptr.LastFocusedStampCount)
}

// SetLastFocusedStampCount sets the value in LastFocusedStampCount.
func (p GuiViewportP) SetLastFocusedStampCount(value int32) {
	ptr := (*C.ImGuiViewportP)(unsafe.Pointer(p))
	ptr.LastFocusedStampCount = (C.int32_t)(value)
}

// GetLastFrameActive returns the value in LastFrameActive.
func (p GuiViewportP) GetLastFrameActive() int32 {
	ptr := (*C.ImGuiViewportP)(unsafe.Pointer(p))
	return int32(ptr.LastFrameActive)
}

// SetLastFrameActive sets the value in LastFrameActive.
func (p GuiViewportP) SetLastFrameActive(value int32) {
	ptr := (*C.ImGuiViewportP)(unsafe.Pointer(p))
	ptr.LastFrameActive = (C.int32_t)(value)
}

// GuiViewportP.LastNameHash is unsupported: category unsupported.

// RefLastPlatformPos returns pointer to the LastPlatformPos field.
func (p GuiViewportP) RefLastPlatformPos() Vec2 {
	return Vec2(p + GuiViewportP(C.offsetof_ImGuiViewportP_LastPlatformPos))
}

// RefLastPlatformSize returns pointer to the LastPlatformSize field.
func (p GuiViewportP) RefLastPlatformSize() Vec2 {
	return Vec2(p + GuiViewportP(C.offsetof_ImGuiViewportP_LastPlatformSize))
}

// RefLastPos returns pointer to the LastPos field.
func (p GuiViewportP) RefLastPos() Vec2 {
	return Vec2(p + GuiViewportP(C.offsetof_ImGuiViewportP_LastPos))
}

// RefLastRendererSize returns pointer to the LastRendererSize field.
func (p GuiViewportP) RefLastRendererSize() Vec2 {
	return Vec2(p + GuiViewportP(C.offsetof_ImGuiViewportP_LastRendererSize))
}

// RefLastSize returns pointer to the LastSize field.
func (p GuiViewportP) RefLastSize() Vec2 {
	return Vec2(p + GuiViewportP(C.offsetof_ImGuiViewportP_LastSize))
}

// GuiViewportP.PlatformMonitor is unsupported: category unsupported.

// GetWindow returns the value in Window.
func (p GuiViewportP) GetWindow() GuiWindow {
	ptr := (*C.ImGuiViewportP)(unsafe.Pointer(p))
	return GuiWindow(unsafe.Pointer(ptr.Window))
}

// SetWindow sets the value in Window.
func (p GuiViewportP) SetWindow(value GuiWindow) {
	ptr := (*C.ImGuiViewportP)(unsafe.Pointer(p))
	ptr.Window = (*C.ImGuiWindow)(unsafe.Pointer(value))
}

// RefWorkInsetMax returns pointer to the WorkInsetMax field.
func (p GuiViewportP) RefWorkInsetMax() Vec2 {
	return Vec2(p + GuiViewportP(C.offsetof_ImGuiViewportP_WorkInsetMax))
}

// RefWorkInsetMin returns pointer to the WorkInsetMin field.
func (p GuiViewportP) RefWorkInsetMin() Vec2 {
	return Vec2(p + GuiViewportP(C.offsetof_ImGuiViewportP_WorkInsetMin))
}

// RefImGuiViewport returns pointer to the _ImGuiViewport field.
func (p GuiViewportP) RefImGuiViewport() GuiViewport {
	return GuiViewport(p + GuiViewportP(C.offsetof_ImGuiViewportP__ImGuiViewport))
}

// GuiWindow wraps struct ImGuiWindow.
type GuiWindow uintptr

// GuiWindowNil is a null pointer.
var GuiWindowNil GuiWindow

// GuiWindowSizeOf is the byte size of ImGuiWindow.
const GuiWindowSizeOf = int(C.sizeof_ImGuiWindow)

// ImGuiWindow allocator is suppressed.

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p GuiWindow) Offset(offset int) GuiWindow {
	return p + GuiWindow(offset*GuiWindowSizeOf)
}

// GetActive returns the value in Active.
func (p GuiWindow) GetActive() bool {
	ptr := (*C.ImGuiWindow)(unsafe.Pointer(p))
	return bool(ptr.Active)
}

// SetActive sets the value in Active.
func (p GuiWindow) SetActive(value bool) {
	ptr := (*C.ImGuiWindow)(unsafe.Pointer(p))
	ptr.Active = (C.bool)(value)
}

// GetAppearing returns the value in Appearing.
func (p GuiWindow) GetAppearing() bool {
	ptr := (*C.ImGuiWindow)(unsafe.Pointer(p))
	return bool(ptr.Appearing)
}

// SetAppearing sets the value in Appearing.
func (p GuiWindow) SetAppearing(value bool) {
	ptr := (*C.ImGuiWindow)(unsafe.Pointer(p))
	ptr.Appearing = (C.bool)(value)
}

// GuiWindow.AutoFitFramesX is unsupported: category unsupported.

// GuiWindow.AutoFitFramesY is unsupported: category unsupported.

// GetAutoFitOnlyGrows returns the value in AutoFitOnlyGrows.
func (p GuiWindow) GetAutoFitOnlyGrows() bool {
	ptr := (*C.ImGuiWindow)(unsafe.Pointer(p))
	return bool(ptr.AutoFitOnlyGrows)
}

// SetAutoFitOnlyGrows sets the value in AutoFitOnlyGrows.
func (p GuiWindow) SetAutoFitOnlyGrows(value bool) {
	ptr := (*C.ImGuiWindow)(unsafe.Pointer(p))
	ptr.AutoFitOnlyGrows = (C.bool)(value)
}

// GetAutoPosLastDirection returns the value in AutoPosLastDirection.
func (p GuiWindow) GetAutoPosLastDirection() GuiDir {
	ptr := (*C.ImGuiWindow)(unsafe.Pointer(p))
	return GuiDir(ptr.AutoPosLastDirection)
}

// SetAutoPosLastDirection sets the value in AutoPosLastDirection.
func (p GuiWindow) SetAutoPosLastDirection(value GuiDir) {
	ptr := (*C.ImGuiWindow)(unsafe.Pointer(p))
	ptr.AutoPosLastDirection = (C.ImGuiDir)(value)
}

// GuiWindow.BeginCount is unsupported: category unsupported.

// GuiWindow.BeginCountPreviousFrame is unsupported: category unsupported.

// GuiWindow.BeginOrderWithinContext is unsupported: category unsupported.

// GuiWindow.BeginOrderWithinParent is unsupported: category unsupported.

// GetChildFlags returns the value in ChildFlags.
func (p GuiWindow) GetChildFlags() GuiChildFlags {
	ptr := (*C.ImGuiWindow)(unsafe.Pointer(p))
	return GuiChildFlags(ptr.ChildFlags)
}

// SetChildFlags sets the value in ChildFlags.
func (p GuiWindow) SetChildFlags(value GuiChildFlags) {
	ptr := (*C.ImGuiWindow)(unsafe.Pointer(p))
	ptr.ChildFlags = (C.ImGuiChildFlags)(value)
}

// GuiWindow.ChildId is unsupported: category unsupported.

// RefClipRect returns pointer to the ClipRect field.
func (p GuiWindow) RefClipRect() Rect {
	return Rect(p + GuiWindow(C.offsetof_ImGuiWindow_ClipRect))
}

// GetCollapsed returns the value in Collapsed.
func (p GuiWindow) GetCollapsed() bool {
	ptr := (*C.ImGuiWindow)(unsafe.Pointer(p))
	return bool(ptr.Collapsed)
}

// SetCollapsed sets the value in Collapsed.
func (p GuiWindow) SetCollapsed(value bool) {
	ptr := (*C.ImGuiWindow)(unsafe.Pointer(p))
	ptr.Collapsed = (C.bool)(value)
}

// RefColumnsStorage returns pointer to the ColumnsStorage field.
func (p GuiWindow) RefColumnsStorage() Vector_ImGuiOldColumns {
	return Vector_ImGuiOldColumns(p + GuiWindow(C.offsetof_ImGuiWindow_ColumnsStorage))
}

// RefContentRegionRect returns pointer to the ContentRegionRect field.
func (p GuiWindow) RefContentRegionRect() Rect {
	return Rect(p + GuiWindow(C.offsetof_ImGuiWindow_ContentRegionRect))
}

// RefContentSize returns pointer to the ContentSize field.
func (p GuiWindow) RefContentSize() Vec2 {
	return Vec2(p + GuiWindow(C.offsetof_ImGuiWindow_ContentSize))
}

// RefContentSizeExplicit returns pointer to the ContentSizeExplicit field.
func (p GuiWindow) RefContentSizeExplicit() Vec2 {
	return Vec2(p + GuiWindow(C.offsetof_ImGuiWindow_ContentSizeExplicit))
}

// RefContentSizeIdeal returns pointer to the ContentSizeIdeal field.
func (p GuiWindow) RefContentSizeIdeal() Vec2 {
	return Vec2(p + GuiWindow(C.offsetof_ImGuiWindow_ContentSizeIdeal))
}

// GetCtx returns the value in Ctx.
func (p GuiWindow) GetCtx() GuiContext {
	ptr := (*C.ImGuiWindow)(unsafe.Pointer(p))
	return GuiContext(unsafe.Pointer(ptr.Ctx))
}

// SetCtx sets the value in Ctx.
func (p GuiWindow) SetCtx(value GuiContext) {
	ptr := (*C.ImGuiWindow)(unsafe.Pointer(p))
	ptr.Ctx = (*C.ImGuiContext)(unsafe.Pointer(value))
}

// RefDC returns pointer to the DC field.
func (p GuiWindow) RefDC() GuiWindowTempData {
	return GuiWindowTempData(p + GuiWindow(C.offsetof_ImGuiWindow_DC))
}

// GetDecoInnerSizeX1 returns the value in DecoInnerSizeX1.
func (p GuiWindow) GetDecoInnerSizeX1() float32 {
	ptr := (*C.ImGuiWindow)(unsafe.Pointer(p))
	return float32(ptr.DecoInnerSizeX1)
}

// SetDecoInnerSizeX1 sets the value in DecoInnerSizeX1.
func (p GuiWindow) SetDecoInnerSizeX1(value float32) {
	ptr := (*C.ImGuiWindow)(unsafe.Pointer(p))
	ptr.DecoInnerSizeX1 = (C.float)(value)
}

// GetDecoInnerSizeY1 returns the value in DecoInnerSizeY1.
func (p GuiWindow) GetDecoInnerSizeY1() float32 {
	ptr := (*C.ImGuiWindow)(unsafe.Pointer(p))
	return float32(ptr.DecoInnerSizeY1)
}

// SetDecoInnerSizeY1 sets the value in DecoInnerSizeY1.
func (p GuiWindow) SetDecoInnerSizeY1(value float32) {
	ptr := (*C.ImGuiWindow)(unsafe.Pointer(p))
	ptr.DecoInnerSizeY1 = (C.float)(value)
}

// GetDecoOuterSizeX1 returns the value in DecoOuterSizeX1.
func (p GuiWindow) GetDecoOuterSizeX1() float32 {
	ptr := (*C.ImGuiWindow)(unsafe.Pointer(p))
	return float32(ptr.DecoOuterSizeX1)
}

// SetDecoOuterSizeX1 sets the value in DecoOuterSizeX1.
func (p GuiWindow) SetDecoOuterSizeX1(value float32) {
	ptr := (*C.ImGuiWindow)(unsafe.Pointer(p))
	ptr.DecoOuterSizeX1 = (C.float)(value)
}

// GetDecoOuterSizeX2 returns the value in DecoOuterSizeX2.
func (p GuiWindow) GetDecoOuterSizeX2() float32 {
	ptr := (*C.ImGuiWindow)(unsafe.Pointer(p))
	return float32(ptr.DecoOuterSizeX2)
}

// SetDecoOuterSizeX2 sets the value in DecoOuterSizeX2.
func (p GuiWindow) SetDecoOuterSizeX2(value float32) {
	ptr := (*C.ImGuiWindow)(unsafe.Pointer(p))
	ptr.DecoOuterSizeX2 = (C.float)(value)
}

// GetDecoOuterSizeY1 returns the value in DecoOuterSizeY1.
func (p GuiWindow) GetDecoOuterSizeY1() float32 {
	ptr := (*C.ImGuiWindow)(unsafe.Pointer(p))
	return float32(ptr.DecoOuterSizeY1)
}

// SetDecoOuterSizeY1 sets the value in DecoOuterSizeY1.
func (p GuiWindow) SetDecoOuterSizeY1(value float32) {
	ptr := (*C.ImGuiWindow)(unsafe.Pointer(p))
	ptr.DecoOuterSizeY1 = (C.float)(value)
}

// GetDecoOuterSizeY2 returns the value in DecoOuterSizeY2.
func (p GuiWindow) GetDecoOuterSizeY2() float32 {
	ptr := (*C.ImGuiWindow)(unsafe.Pointer(p))
	return float32(ptr.DecoOuterSizeY2)
}

// SetDecoOuterSizeY2 sets the value in DecoOuterSizeY2.
func (p GuiWindow) SetDecoOuterSizeY2(value float32) {
	ptr := (*C.ImGuiWindow)(unsafe.Pointer(p))
	ptr.DecoOuterSizeY2 = (C.float)(value)
}

// GuiWindow.DisableInputsFrames is unsupported: category unsupported.

// GuiWindow.DockId is unsupported: category unsupported.

// GuiWindow.DockIsActive is unsupported: category unsupported.

// GetDockNode returns the value in DockNode.
func (p GuiWindow) GetDockNode() GuiDockNode {
	ptr := (*C.ImGuiWindow)(unsafe.Pointer(p))
	return GuiDockNode(unsafe.Pointer(ptr.DockNode))
}

// SetDockNode sets the value in DockNode.
func (p GuiWindow) SetDockNode(value GuiDockNode) {
	ptr := (*C.ImGuiWindow)(unsafe.Pointer(p))
	ptr.DockNode = (*C.ImGuiDockNode)(unsafe.Pointer(value))
}

// GetDockNodeAsHost returns the value in DockNodeAsHost.
func (p GuiWindow) GetDockNodeAsHost() GuiDockNode {
	ptr := (*C.ImGuiWindow)(unsafe.Pointer(p))
	return GuiDockNode(unsafe.Pointer(ptr.DockNodeAsHost))
}

// SetDockNodeAsHost sets the value in DockNodeAsHost.
func (p GuiWindow) SetDockNodeAsHost(value GuiDockNode) {
	ptr := (*C.ImGuiWindow)(unsafe.Pointer(p))
	ptr.DockNodeAsHost = (*C.ImGuiDockNode)(unsafe.Pointer(value))
}

// GuiWindow.DockNodeIsVisible is unsupported: category unsupported.

// GuiWindow.DockOrder is unsupported: category unsupported.

// RefDockStyle returns pointer to the DockStyle field.
func (p GuiWindow) RefDockStyle() GuiWindowDockStyle {
	return GuiWindowDockStyle(p + GuiWindow(C.offsetof_ImGuiWindow_DockStyle))
}

// GuiWindow.DockTabIsVisible is unsupported: category unsupported.

// GuiWindow.DockTabWantClose is unsupported: category unsupported.

// GetDrawList returns the value in DrawList.
func (p GuiWindow) GetDrawList() DrawList {
	ptr := (*C.ImGuiWindow)(unsafe.Pointer(p))
	return DrawList(unsafe.Pointer(ptr.DrawList))
}

// SetDrawList sets the value in DrawList.
func (p GuiWindow) SetDrawList(value DrawList) {
	ptr := (*C.ImGuiWindow)(unsafe.Pointer(p))
	ptr.DrawList = (*C.ImDrawList)(unsafe.Pointer(value))
}

// RefDrawListInst returns pointer to the DrawListInst field.
func (p GuiWindow) RefDrawListInst() DrawList {
	return DrawList(p + GuiWindow(C.offsetof_ImGuiWindow_DrawListInst))
}

// GetFlags returns the value in Flags.
func (p GuiWindow) GetFlags() GuiWindowFlags {
	ptr := (*C.ImGuiWindow)(unsafe.Pointer(p))
	return GuiWindowFlags(ptr.Flags)
}

// SetFlags sets the value in Flags.
func (p GuiWindow) SetFlags(value GuiWindowFlags) {
	ptr := (*C.ImGuiWindow)(unsafe.Pointer(p))
	ptr.Flags = (C.ImGuiWindowFlags)(value)
}

// GetFlagsPreviousFrame returns the value in FlagsPreviousFrame.
func (p GuiWindow) GetFlagsPreviousFrame() GuiWindowFlags {
	ptr := (*C.ImGuiWindow)(unsafe.Pointer(p))
	return GuiWindowFlags(ptr.FlagsPreviousFrame)
}

// SetFlagsPreviousFrame sets the value in FlagsPreviousFrame.
func (p GuiWindow) SetFlagsPreviousFrame(value GuiWindowFlags) {
	ptr := (*C.ImGuiWindow)(unsafe.Pointer(p))
	ptr.FlagsPreviousFrame = (C.ImGuiWindowFlags)(value)
}

// GuiWindow.FocusOrder is unsupported: category unsupported.

// GetFontRefSize returns the value in FontRefSize.
func (p GuiWindow) GetFontRefSize() float32 {
	ptr := (*C.ImGuiWindow)(unsafe.Pointer(p))
	return float32(ptr.FontRefSize)
}

// SetFontRefSize sets the value in FontRefSize.
func (p GuiWindow) SetFontRefSize(value float32) {
	ptr := (*C.ImGuiWindow)(unsafe.Pointer(p))
	ptr.FontRefSize = (C.float)(value)
}

// GetFontWindowScale returns the value in FontWindowScale.
func (p GuiWindow) GetFontWindowScale() float32 {
	ptr := (*C.ImGuiWindow)(unsafe.Pointer(p))
	return float32(ptr.FontWindowScale)
}

// SetFontWindowScale sets the value in FontWindowScale.
func (p GuiWindow) SetFontWindowScale(value float32) {
	ptr := (*C.ImGuiWindow)(unsafe.Pointer(p))
	ptr.FontWindowScale = (C.float)(value)
}

// GetFontWindowScaleParents returns the value in FontWindowScaleParents.
func (p GuiWindow) GetFontWindowScaleParents() float32 {
	ptr := (*C.ImGuiWindow)(unsafe.Pointer(p))
	return float32(ptr.FontWindowScaleParents)
}

// SetFontWindowScaleParents sets the value in FontWindowScaleParents.
func (p GuiWindow) SetFontWindowScaleParents(value float32) {
	ptr := (*C.ImGuiWindow)(unsafe.Pointer(p))
	ptr.FontWindowScaleParents = (C.float)(value)
}

// GetHasCloseButton returns the value in HasCloseButton.
func (p GuiWindow) GetHasCloseButton() bool {
	ptr := (*C.ImGuiWindow)(unsafe.Pointer(p))
	return bool(ptr.HasCloseButton)
}

// SetHasCloseButton sets the value in HasCloseButton.
func (p GuiWindow) SetHasCloseButton(value bool) {
	ptr := (*C.ImGuiWindow)(unsafe.Pointer(p))
	ptr.HasCloseButton = (C.bool)(value)
}

// GetHidden returns the value in Hidden.
func (p GuiWindow) GetHidden() bool {
	ptr := (*C.ImGuiWindow)(unsafe.Pointer(p))
	return bool(ptr.Hidden)
}

// SetHidden sets the value in Hidden.
func (p GuiWindow) SetHidden(value bool) {
	ptr := (*C.ImGuiWindow)(unsafe.Pointer(p))
	ptr.Hidden = (C.bool)(value)
}

// GuiWindow.HiddenFramesCanSkipItems is unsupported: category unsupported.

// GuiWindow.HiddenFramesCannotSkipItems is unsupported: category unsupported.

// GuiWindow.HiddenFramesForRenderOnly is unsupported: category unsupported.

// RefHitTestHoleOffset returns pointer to the HitTestHoleOffset field.
func (p GuiWindow) RefHitTestHoleOffset() Vec2ih {
	return Vec2ih(p + GuiWindow(C.offsetof_ImGuiWindow_HitTestHoleOffset))
}

// RefHitTestHoleSize returns pointer to the HitTestHoleSize field.
func (p GuiWindow) RefHitTestHoleSize() Vec2ih {
	return Vec2ih(p + GuiWindow(C.offsetof_ImGuiWindow_HitTestHoleSize))
}

// GuiWindow.ID is unsupported: category unsupported.

// RefIDStack returns pointer to the IDStack field.
func (p GuiWindow) RefIDStack() Vector_ImGuiID {
	return Vector_ImGuiID(p + GuiWindow(C.offsetof_ImGuiWindow_IDStack))
}

// RefInnerClipRect returns pointer to the InnerClipRect field.
func (p GuiWindow) RefInnerClipRect() Rect {
	return Rect(p + GuiWindow(C.offsetof_ImGuiWindow_InnerClipRect))
}

// RefInnerRect returns pointer to the InnerRect field.
func (p GuiWindow) RefInnerRect() Rect {
	return Rect(p + GuiWindow(C.offsetof_ImGuiWindow_InnerRect))
}

// GetIsExplicitChild returns the value in IsExplicitChild.
func (p GuiWindow) GetIsExplicitChild() bool {
	ptr := (*C.ImGuiWindow)(unsafe.Pointer(p))
	return bool(ptr.IsExplicitChild)
}

// SetIsExplicitChild sets the value in IsExplicitChild.
func (p GuiWindow) SetIsExplicitChild(value bool) {
	ptr := (*C.ImGuiWindow)(unsafe.Pointer(p))
	ptr.IsExplicitChild = (C.bool)(value)
}

// GetIsFallbackWindow returns the value in IsFallbackWindow.
func (p GuiWindow) GetIsFallbackWindow() bool {
	ptr := (*C.ImGuiWindow)(unsafe.Pointer(p))
	return bool(ptr.IsFallbackWindow)
}

// SetIsFallbackWindow sets the value in IsFallbackWindow.
func (p GuiWindow) SetIsFallbackWindow(value bool) {
	ptr := (*C.ImGuiWindow)(unsafe.Pointer(p))
	ptr.IsFallbackWindow = (C.bool)(value)
}

// GetItemWidthDefault returns the value in ItemWidthDefault.
func (p GuiWindow) GetItemWidthDefault() float32 {
	ptr := (*C.ImGuiWindow)(unsafe.Pointer(p))
	return float32(ptr.ItemWidthDefault)
}

// SetItemWidthDefault sets the value in ItemWidthDefault.
func (p GuiWindow) SetItemWidthDefault(value float32) {
	ptr := (*C.ImGuiWindow)(unsafe.Pointer(p))
	ptr.ItemWidthDefault = (C.float)(value)
}

// GetLastFrameActive returns the value in LastFrameActive.
func (p GuiWindow) GetLastFrameActive() int32 {
	ptr := (*C.ImGuiWindow)(unsafe.Pointer(p))
	return int32(ptr.LastFrameActive)
}

// SetLastFrameActive sets the value in LastFrameActive.
func (p GuiWindow) SetLastFrameActive(value int32) {
	ptr := (*C.ImGuiWindow)(unsafe.Pointer(p))
	ptr.LastFrameActive = (C.int32_t)(value)
}

// GetLastFrameJustFocused returns the value in LastFrameJustFocused.
func (p GuiWindow) GetLastFrameJustFocused() int32 {
	ptr := (*C.ImGuiWindow)(unsafe.Pointer(p))
	return int32(ptr.LastFrameJustFocused)
}

// SetLastFrameJustFocused sets the value in LastFrameJustFocused.
func (p GuiWindow) SetLastFrameJustFocused(value int32) {
	ptr := (*C.ImGuiWindow)(unsafe.Pointer(p))
	ptr.LastFrameJustFocused = (C.int32_t)(value)
}

// GetLastTimeActive returns the value in LastTimeActive.
func (p GuiWindow) GetLastTimeActive() float32 {
	ptr := (*C.ImGuiWindow)(unsafe.Pointer(p))
	return float32(ptr.LastTimeActive)
}

// SetLastTimeActive sets the value in LastTimeActive.
func (p GuiWindow) SetLastTimeActive(value float32) {
	ptr := (*C.ImGuiWindow)(unsafe.Pointer(p))
	ptr.LastTimeActive = (C.float)(value)
}

// GetMemoryCompacted returns the value in MemoryCompacted.
func (p GuiWindow) GetMemoryCompacted() bool {
	ptr := (*C.ImGuiWindow)(unsafe.Pointer(p))
	return bool(ptr.MemoryCompacted)
}

// SetMemoryCompacted sets the value in MemoryCompacted.
func (p GuiWindow) SetMemoryCompacted(value bool) {
	ptr := (*C.ImGuiWindow)(unsafe.Pointer(p))
	ptr.MemoryCompacted = (C.bool)(value)
}

// GetMemoryDrawListIdxCapacity returns the value in MemoryDrawListIdxCapacity.
func (p GuiWindow) GetMemoryDrawListIdxCapacity() int32 {
	ptr := (*C.ImGuiWindow)(unsafe.Pointer(p))
	return int32(ptr.MemoryDrawListIdxCapacity)
}

// SetMemoryDrawListIdxCapacity sets the value in MemoryDrawListIdxCapacity.
func (p GuiWindow) SetMemoryDrawListIdxCapacity(value int32) {
	ptr := (*C.ImGuiWindow)(unsafe.Pointer(p))
	ptr.MemoryDrawListIdxCapacity = (C.int32_t)(value)
}

// GetMemoryDrawListVtxCapacity returns the value in MemoryDrawListVtxCapacity.
func (p GuiWindow) GetMemoryDrawListVtxCapacity() int32 {
	ptr := (*C.ImGuiWindow)(unsafe.Pointer(p))
	return int32(ptr.MemoryDrawListVtxCapacity)
}

// SetMemoryDrawListVtxCapacity sets the value in MemoryDrawListVtxCapacity.
func (p GuiWindow) SetMemoryDrawListVtxCapacity(value int32) {
	ptr := (*C.ImGuiWindow)(unsafe.Pointer(p))
	ptr.MemoryDrawListVtxCapacity = (C.int32_t)(value)
}

// GetMenuBarHeight returns the value in MenuBarHeight.
func (p GuiWindow) GetMenuBarHeight() float32 {
	ptr := (*C.ImGuiWindow)(unsafe.Pointer(p))
	return float32(ptr.MenuBarHeight)
}

// SetMenuBarHeight sets the value in MenuBarHeight.
func (p GuiWindow) SetMenuBarHeight(value float32) {
	ptr := (*C.ImGuiWindow)(unsafe.Pointer(p))
	ptr.MenuBarHeight = (C.float)(value)
}

// GuiWindow.MoveId is unsupported: category unsupported.

// GetName returns the value in Name.
func (p GuiWindow) GetName() ffi.CString {
	ptr := (*C.ImGuiWindow)(unsafe.Pointer(p))
	return ffi.CStringFromPtr((unsafe.Pointer)(ptr.Name))
}

// SetName sets the value in Name.
func (p GuiWindow) SetName(value ffi.CString) {
	ptr := (*C.ImGuiWindow)(unsafe.Pointer(p))
	ptr.Name = (*C.char)(value.Raw())
}

// GetNameBufLen returns the value in NameBufLen.
func (p GuiWindow) GetNameBufLen() int32 {
	ptr := (*C.ImGuiWindow)(unsafe.Pointer(p))
	return int32(ptr.NameBufLen)
}

// SetNameBufLen sets the value in NameBufLen.
func (p GuiWindow) SetNameBufLen(value int32) {
	ptr := (*C.ImGuiWindow)(unsafe.Pointer(p))
	ptr.NameBufLen = (C.int32_t)(value)
}

// GetNavLastChildNavWindow returns the value in NavLastChildNavWindow.
func (p GuiWindow) GetNavLastChildNavWindow() GuiWindow {
	ptr := (*C.ImGuiWindow)(unsafe.Pointer(p))
	return GuiWindow(unsafe.Pointer(ptr.NavLastChildNavWindow))
}

// SetNavLastChildNavWindow sets the value in NavLastChildNavWindow.
func (p GuiWindow) SetNavLastChildNavWindow(value GuiWindow) {
	ptr := (*C.ImGuiWindow)(unsafe.Pointer(p))
	ptr.NavLastChildNavWindow = (*C.ImGuiWindow)(unsafe.Pointer(value))
}

// GuiWindow.NavLastIds[ImGuiNavLayer_COUNT] is unsupported: category unsupported.

// GuiWindow.NavPreferredScoringPosRel[ImGuiNavLayer_COUNT] is unsupported: category unsupported.

// GuiWindow.NavRectRel[ImGuiNavLayer_COUNT] is unsupported: category unsupported.

// GuiWindow.NavRootFocusScopeId is unsupported: category unsupported.

// RefOuterRectClipped returns pointer to the OuterRectClipped field.
func (p GuiWindow) RefOuterRectClipped() Rect {
	return Rect(p + GuiWindow(C.offsetof_ImGuiWindow_OuterRectClipped))
}

// GetParentWindow returns the value in ParentWindow.
func (p GuiWindow) GetParentWindow() GuiWindow {
	ptr := (*C.ImGuiWindow)(unsafe.Pointer(p))
	return GuiWindow(unsafe.Pointer(ptr.ParentWindow))
}

// SetParentWindow sets the value in ParentWindow.
func (p GuiWindow) SetParentWindow(value GuiWindow) {
	ptr := (*C.ImGuiWindow)(unsafe.Pointer(p))
	ptr.ParentWindow = (*C.ImGuiWindow)(unsafe.Pointer(value))
}

// GetParentWindowForFocusRoute returns the value in ParentWindowForFocusRoute.
func (p GuiWindow) GetParentWindowForFocusRoute() GuiWindow {
	ptr := (*C.ImGuiWindow)(unsafe.Pointer(p))
	return GuiWindow(unsafe.Pointer(ptr.ParentWindowForFocusRoute))
}

// SetParentWindowForFocusRoute sets the value in ParentWindowForFocusRoute.
func (p GuiWindow) SetParentWindowForFocusRoute(value GuiWindow) {
	ptr := (*C.ImGuiWindow)(unsafe.Pointer(p))
	ptr.ParentWindowForFocusRoute = (*C.ImGuiWindow)(unsafe.Pointer(value))
}

// GetParentWindowInBeginStack returns the value in ParentWindowInBeginStack.
func (p GuiWindow) GetParentWindowInBeginStack() GuiWindow {
	ptr := (*C.ImGuiWindow)(unsafe.Pointer(p))
	return GuiWindow(unsafe.Pointer(ptr.ParentWindowInBeginStack))
}

// SetParentWindowInBeginStack sets the value in ParentWindowInBeginStack.
func (p GuiWindow) SetParentWindowInBeginStack(value GuiWindow) {
	ptr := (*C.ImGuiWindow)(unsafe.Pointer(p))
	ptr.ParentWindowInBeginStack = (*C.ImGuiWindow)(unsafe.Pointer(value))
}

// RefParentWorkRect returns pointer to the ParentWorkRect field.
func (p GuiWindow) RefParentWorkRect() Rect {
	return Rect(p + GuiWindow(C.offsetof_ImGuiWindow_ParentWorkRect))
}

// GuiWindow.PopupId is unsupported: category unsupported.

// RefPos returns pointer to the Pos field.
func (p GuiWindow) RefPos() Vec2 {
	return Vec2(p + GuiWindow(C.offsetof_ImGuiWindow_Pos))
}

// GuiWindow.ResizeBorderHeld is unsupported: category unsupported.

// GuiWindow.ResizeBorderHovered is unsupported: category unsupported.

// GetRootWindow returns the value in RootWindow.
func (p GuiWindow) GetRootWindow() GuiWindow {
	ptr := (*C.ImGuiWindow)(unsafe.Pointer(p))
	return GuiWindow(unsafe.Pointer(ptr.RootWindow))
}

// SetRootWindow sets the value in RootWindow.
func (p GuiWindow) SetRootWindow(value GuiWindow) {
	ptr := (*C.ImGuiWindow)(unsafe.Pointer(p))
	ptr.RootWindow = (*C.ImGuiWindow)(unsafe.Pointer(value))
}

// GetRootWindowDockTree returns the value in RootWindowDockTree.
func (p GuiWindow) GetRootWindowDockTree() GuiWindow {
	ptr := (*C.ImGuiWindow)(unsafe.Pointer(p))
	return GuiWindow(unsafe.Pointer(ptr.RootWindowDockTree))
}

// SetRootWindowDockTree sets the value in RootWindowDockTree.
func (p GuiWindow) SetRootWindowDockTree(value GuiWindow) {
	ptr := (*C.ImGuiWindow)(unsafe.Pointer(p))
	ptr.RootWindowDockTree = (*C.ImGuiWindow)(unsafe.Pointer(value))
}

// GetRootWindowForNav returns the value in RootWindowForNav.
func (p GuiWindow) GetRootWindowForNav() GuiWindow {
	ptr := (*C.ImGuiWindow)(unsafe.Pointer(p))
	return GuiWindow(unsafe.Pointer(ptr.RootWindowForNav))
}

// SetRootWindowForNav sets the value in RootWindowForNav.
func (p GuiWindow) SetRootWindowForNav(value GuiWindow) {
	ptr := (*C.ImGuiWindow)(unsafe.Pointer(p))
	ptr.RootWindowForNav = (*C.ImGuiWindow)(unsafe.Pointer(value))
}

// GetRootWindowForTitleBarHighlight returns the value in RootWindowForTitleBarHighlight.
func (p GuiWindow) GetRootWindowForTitleBarHighlight() GuiWindow {
	ptr := (*C.ImGuiWindow)(unsafe.Pointer(p))
	return GuiWindow(unsafe.Pointer(ptr.RootWindowForTitleBarHighlight))
}

// SetRootWindowForTitleBarHighlight sets the value in RootWindowForTitleBarHighlight.
func (p GuiWindow) SetRootWindowForTitleBarHighlight(value GuiWindow) {
	ptr := (*C.ImGuiWindow)(unsafe.Pointer(p))
	ptr.RootWindowForTitleBarHighlight = (*C.ImGuiWindow)(unsafe.Pointer(value))
}

// GetRootWindowPopupTree returns the value in RootWindowPopupTree.
func (p GuiWindow) GetRootWindowPopupTree() GuiWindow {
	ptr := (*C.ImGuiWindow)(unsafe.Pointer(p))
	return GuiWindow(unsafe.Pointer(ptr.RootWindowPopupTree))
}

// SetRootWindowPopupTree sets the value in RootWindowPopupTree.
func (p GuiWindow) SetRootWindowPopupTree(value GuiWindow) {
	ptr := (*C.ImGuiWindow)(unsafe.Pointer(p))
	ptr.RootWindowPopupTree = (*C.ImGuiWindow)(unsafe.Pointer(value))
}

// RefScroll returns pointer to the Scroll field.
func (p GuiWindow) RefScroll() Vec2 {
	return Vec2(p + GuiWindow(C.offsetof_ImGuiWindow_Scroll))
}

// RefScrollMax returns pointer to the ScrollMax field.
func (p GuiWindow) RefScrollMax() Vec2 {
	return Vec2(p + GuiWindow(C.offsetof_ImGuiWindow_ScrollMax))
}

// RefScrollTarget returns pointer to the ScrollTarget field.
func (p GuiWindow) RefScrollTarget() Vec2 {
	return Vec2(p + GuiWindow(C.offsetof_ImGuiWindow_ScrollTarget))
}

// RefScrollTargetCenterRatio returns pointer to the ScrollTargetCenterRatio field.
func (p GuiWindow) RefScrollTargetCenterRatio() Vec2 {
	return Vec2(p + GuiWindow(C.offsetof_ImGuiWindow_ScrollTargetCenterRatio))
}

// RefScrollTargetEdgeSnapDist returns pointer to the ScrollTargetEdgeSnapDist field.
func (p GuiWindow) RefScrollTargetEdgeSnapDist() Vec2 {
	return Vec2(p + GuiWindow(C.offsetof_ImGuiWindow_ScrollTargetEdgeSnapDist))
}

// RefScrollbarSizes returns pointer to the ScrollbarSizes field.
func (p GuiWindow) RefScrollbarSizes() Vec2 {
	return Vec2(p + GuiWindow(C.offsetof_ImGuiWindow_ScrollbarSizes))
}

// GetScrollbarX returns the value in ScrollbarX.
func (p GuiWindow) GetScrollbarX() bool {
	ptr := (*C.ImGuiWindow)(unsafe.Pointer(p))
	return bool(ptr.ScrollbarX)
}

// SetScrollbarX sets the value in ScrollbarX.
func (p GuiWindow) SetScrollbarX(value bool) {
	ptr := (*C.ImGuiWindow)(unsafe.Pointer(p))
	ptr.ScrollbarX = (C.bool)(value)
}

// GetScrollbarXStabilizeEnabled returns the value in ScrollbarXStabilizeEnabled.
func (p GuiWindow) GetScrollbarXStabilizeEnabled() bool {
	ptr := (*C.ImGuiWindow)(unsafe.Pointer(p))
	return bool(ptr.ScrollbarXStabilizeEnabled)
}

// SetScrollbarXStabilizeEnabled sets the value in ScrollbarXStabilizeEnabled.
func (p GuiWindow) SetScrollbarXStabilizeEnabled(value bool) {
	ptr := (*C.ImGuiWindow)(unsafe.Pointer(p))
	ptr.ScrollbarXStabilizeEnabled = (C.bool)(value)
}

// GetScrollbarXStabilizeToggledHistory returns the value in ScrollbarXStabilizeToggledHistory.
func (p GuiWindow) GetScrollbarXStabilizeToggledHistory() uint8 {
	ptr := (*C.ImGuiWindow)(unsafe.Pointer(p))
	return uint8(ptr.ScrollbarXStabilizeToggledHistory)
}

// SetScrollbarXStabilizeToggledHistory sets the value in ScrollbarXStabilizeToggledHistory.
func (p GuiWindow) SetScrollbarXStabilizeToggledHistory(value uint8) {
	ptr := (*C.ImGuiWindow)(unsafe.Pointer(p))
	ptr.ScrollbarXStabilizeToggledHistory = (C.uint8_t)(value)
}

// GetScrollbarY returns the value in ScrollbarY.
func (p GuiWindow) GetScrollbarY() bool {
	ptr := (*C.ImGuiWindow)(unsafe.Pointer(p))
	return bool(ptr.ScrollbarY)
}

// SetScrollbarY sets the value in ScrollbarY.
func (p GuiWindow) SetScrollbarY(value bool) {
	ptr := (*C.ImGuiWindow)(unsafe.Pointer(p))
	ptr.ScrollbarY = (C.bool)(value)
}

// GuiWindow.SetWindowCollapsedAllowFlags is unsupported: category unsupported.

// GuiWindow.SetWindowDockAllowFlags is unsupported: category unsupported.

// GuiWindow.SetWindowPosAllowFlags is unsupported: category unsupported.

// RefSetWindowPosPivot returns pointer to the SetWindowPosPivot field.
func (p GuiWindow) RefSetWindowPosPivot() Vec2 {
	return Vec2(p + GuiWindow(C.offsetof_ImGuiWindow_SetWindowPosPivot))
}

// RefSetWindowPosVal returns pointer to the SetWindowPosVal field.
func (p GuiWindow) RefSetWindowPosVal() Vec2 {
	return Vec2(p + GuiWindow(C.offsetof_ImGuiWindow_SetWindowPosVal))
}

// GuiWindow.SetWindowSizeAllowFlags is unsupported: category unsupported.

// GetSettingsOffset returns the value in SettingsOffset.
func (p GuiWindow) GetSettingsOffset() int32 {
	ptr := (*C.ImGuiWindow)(unsafe.Pointer(p))
	return int32(ptr.SettingsOffset)
}

// SetSettingsOffset sets the value in SettingsOffset.
func (p GuiWindow) SetSettingsOffset(value int32) {
	ptr := (*C.ImGuiWindow)(unsafe.Pointer(p))
	ptr.SettingsOffset = (C.int32_t)(value)
}

// RefSize returns pointer to the Size field.
func (p GuiWindow) RefSize() Vec2 {
	return Vec2(p + GuiWindow(C.offsetof_ImGuiWindow_Size))
}

// RefSizeFull returns pointer to the SizeFull field.
func (p GuiWindow) RefSizeFull() Vec2 {
	return Vec2(p + GuiWindow(C.offsetof_ImGuiWindow_SizeFull))
}

// GetSkipItems returns the value in SkipItems.
func (p GuiWindow) GetSkipItems() bool {
	ptr := (*C.ImGuiWindow)(unsafe.Pointer(p))
	return bool(ptr.SkipItems)
}

// SetSkipItems sets the value in SkipItems.
func (p GuiWindow) SetSkipItems(value bool) {
	ptr := (*C.ImGuiWindow)(unsafe.Pointer(p))
	ptr.SkipItems = (C.bool)(value)
}

// GetSkipRefresh returns the value in SkipRefresh.
func (p GuiWindow) GetSkipRefresh() bool {
	ptr := (*C.ImGuiWindow)(unsafe.Pointer(p))
	return bool(ptr.SkipRefresh)
}

// SetSkipRefresh sets the value in SkipRefresh.
func (p GuiWindow) SetSkipRefresh(value bool) {
	ptr := (*C.ImGuiWindow)(unsafe.Pointer(p))
	ptr.SkipRefresh = (C.bool)(value)
}

// RefStateStorage returns pointer to the StateStorage field.
func (p GuiWindow) RefStateStorage() GuiStorage {
	return GuiStorage(p + GuiWindow(C.offsetof_ImGuiWindow_StateStorage))
}

// GuiWindow.TabId is unsupported: category unsupported.

// GetTitleBarHeight returns the value in TitleBarHeight.
func (p GuiWindow) GetTitleBarHeight() float32 {
	ptr := (*C.ImGuiWindow)(unsafe.Pointer(p))
	return float32(ptr.TitleBarHeight)
}

// SetTitleBarHeight sets the value in TitleBarHeight.
func (p GuiWindow) SetTitleBarHeight(value float32) {
	ptr := (*C.ImGuiWindow)(unsafe.Pointer(p))
	ptr.TitleBarHeight = (C.float)(value)
}

// GetViewport returns the value in Viewport.
func (p GuiWindow) GetViewport() GuiViewportP {
	ptr := (*C.ImGuiWindow)(unsafe.Pointer(p))
	return GuiViewportP(unsafe.Pointer(ptr.Viewport))
}

// SetViewport sets the value in Viewport.
func (p GuiWindow) SetViewport(value GuiViewportP) {
	ptr := (*C.ImGuiWindow)(unsafe.Pointer(p))
	ptr.Viewport = (*C.ImGuiViewportP)(unsafe.Pointer(value))
}

// GetViewportAllowPlatformMonitorExtend returns the value in ViewportAllowPlatformMonitorExtend.
func (p GuiWindow) GetViewportAllowPlatformMonitorExtend() int32 {
	ptr := (*C.ImGuiWindow)(unsafe.Pointer(p))
	return int32(ptr.ViewportAllowPlatformMonitorExtend)
}

// SetViewportAllowPlatformMonitorExtend sets the value in ViewportAllowPlatformMonitorExtend.
func (p GuiWindow) SetViewportAllowPlatformMonitorExtend(value int32) {
	ptr := (*C.ImGuiWindow)(unsafe.Pointer(p))
	ptr.ViewportAllowPlatformMonitorExtend = (C.int32_t)(value)
}

// GuiWindow.ViewportId is unsupported: category unsupported.

// GetViewportOwned returns the value in ViewportOwned.
func (p GuiWindow) GetViewportOwned() bool {
	ptr := (*C.ImGuiWindow)(unsafe.Pointer(p))
	return bool(ptr.ViewportOwned)
}

// SetViewportOwned sets the value in ViewportOwned.
func (p GuiWindow) SetViewportOwned(value bool) {
	ptr := (*C.ImGuiWindow)(unsafe.Pointer(p))
	ptr.ViewportOwned = (C.bool)(value)
}

// RefViewportPos returns pointer to the ViewportPos field.
func (p GuiWindow) RefViewportPos() Vec2 {
	return Vec2(p + GuiWindow(C.offsetof_ImGuiWindow_ViewportPos))
}

// GetWantCollapseToggle returns the value in WantCollapseToggle.
func (p GuiWindow) GetWantCollapseToggle() bool {
	ptr := (*C.ImGuiWindow)(unsafe.Pointer(p))
	return bool(ptr.WantCollapseToggle)
}

// SetWantCollapseToggle sets the value in WantCollapseToggle.
func (p GuiWindow) SetWantCollapseToggle(value bool) {
	ptr := (*C.ImGuiWindow)(unsafe.Pointer(p))
	ptr.WantCollapseToggle = (C.bool)(value)
}

// GetWasActive returns the value in WasActive.
func (p GuiWindow) GetWasActive() bool {
	ptr := (*C.ImGuiWindow)(unsafe.Pointer(p))
	return bool(ptr.WasActive)
}

// SetWasActive sets the value in WasActive.
func (p GuiWindow) SetWasActive(value bool) {
	ptr := (*C.ImGuiWindow)(unsafe.Pointer(p))
	ptr.WasActive = (C.bool)(value)
}

// GetWindowBorderSize returns the value in WindowBorderSize.
func (p GuiWindow) GetWindowBorderSize() float32 {
	ptr := (*C.ImGuiWindow)(unsafe.Pointer(p))
	return float32(ptr.WindowBorderSize)
}

// SetWindowBorderSize sets the value in WindowBorderSize.
func (p GuiWindow) SetWindowBorderSize(value float32) {
	ptr := (*C.ImGuiWindow)(unsafe.Pointer(p))
	ptr.WindowBorderSize = (C.float)(value)
}

// RefWindowClass returns pointer to the WindowClass field.
func (p GuiWindow) RefWindowClass() GuiWindowClass {
	return GuiWindowClass(p + GuiWindow(C.offsetof_ImGuiWindow_WindowClass))
}

// RefWindowPadding returns pointer to the WindowPadding field.
func (p GuiWindow) RefWindowPadding() Vec2 {
	return Vec2(p + GuiWindow(C.offsetof_ImGuiWindow_WindowPadding))
}

// GetWindowRounding returns the value in WindowRounding.
func (p GuiWindow) GetWindowRounding() float32 {
	ptr := (*C.ImGuiWindow)(unsafe.Pointer(p))
	return float32(ptr.WindowRounding)
}

// SetWindowRounding sets the value in WindowRounding.
func (p GuiWindow) SetWindowRounding(value float32) {
	ptr := (*C.ImGuiWindow)(unsafe.Pointer(p))
	ptr.WindowRounding = (C.float)(value)
}

// RefWorkRect returns pointer to the WorkRect field.
func (p GuiWindow) RefWorkRect() Rect {
	return Rect(p + GuiWindow(C.offsetof_ImGuiWindow_WorkRect))
}

// GetWriteAccessed returns the value in WriteAccessed.
func (p GuiWindow) GetWriteAccessed() bool {
	ptr := (*C.ImGuiWindow)(unsafe.Pointer(p))
	return bool(ptr.WriteAccessed)
}

// SetWriteAccessed sets the value in WriteAccessed.
func (p GuiWindow) SetWriteAccessed(value bool) {
	ptr := (*C.ImGuiWindow)(unsafe.Pointer(p))
	ptr.WriteAccessed = (C.bool)(value)
}

// GuiWindowClass wraps struct ImGuiWindowClass.
type GuiWindowClass uintptr

// GuiWindowClassNil is a null pointer.
var GuiWindowClassNil GuiWindowClass

// GuiWindowClassSizeOf is the byte size of ImGuiWindowClass.
const GuiWindowClassSizeOf = int(C.sizeof_ImGuiWindowClass)

// ImGuiWindowClass allocator is suppressed.

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p GuiWindowClass) Offset(offset int) GuiWindowClass {
	return p + GuiWindowClass(offset*GuiWindowClassSizeOf)
}

// GuiWindowClass.ClassId is unsupported: category unsupported.

// GetDockNodeFlagsOverrideSet returns the value in DockNodeFlagsOverrideSet.
func (p GuiWindowClass) GetDockNodeFlagsOverrideSet() GuiDockNodeFlags {
	ptr := (*C.ImGuiWindowClass)(unsafe.Pointer(p))
	return GuiDockNodeFlags(ptr.DockNodeFlagsOverrideSet)
}

// SetDockNodeFlagsOverrideSet sets the value in DockNodeFlagsOverrideSet.
func (p GuiWindowClass) SetDockNodeFlagsOverrideSet(value GuiDockNodeFlags) {
	ptr := (*C.ImGuiWindowClass)(unsafe.Pointer(p))
	ptr.DockNodeFlagsOverrideSet = (C.ImGuiDockNodeFlags)(value)
}

// GetDockingAllowUnclassed returns the value in DockingAllowUnclassed.
func (p GuiWindowClass) GetDockingAllowUnclassed() bool {
	ptr := (*C.ImGuiWindowClass)(unsafe.Pointer(p))
	return bool(ptr.DockingAllowUnclassed)
}

// SetDockingAllowUnclassed sets the value in DockingAllowUnclassed.
func (p GuiWindowClass) SetDockingAllowUnclassed(value bool) {
	ptr := (*C.ImGuiWindowClass)(unsafe.Pointer(p))
	ptr.DockingAllowUnclassed = (C.bool)(value)
}

// GetDockingAlwaysTabBar returns the value in DockingAlwaysTabBar.
func (p GuiWindowClass) GetDockingAlwaysTabBar() bool {
	ptr := (*C.ImGuiWindowClass)(unsafe.Pointer(p))
	return bool(ptr.DockingAlwaysTabBar)
}

// SetDockingAlwaysTabBar sets the value in DockingAlwaysTabBar.
func (p GuiWindowClass) SetDockingAlwaysTabBar(value bool) {
	ptr := (*C.ImGuiWindowClass)(unsafe.Pointer(p))
	ptr.DockingAlwaysTabBar = (C.bool)(value)
}

// GuiWindowClass.FocusRouteParentWindowId is unsupported: category unsupported.

// GuiWindowClass.ParentViewportId is unsupported: category unsupported.

// GetTabItemFlagsOverrideSet returns the value in TabItemFlagsOverrideSet.
func (p GuiWindowClass) GetTabItemFlagsOverrideSet() GuiTabItemFlags {
	ptr := (*C.ImGuiWindowClass)(unsafe.Pointer(p))
	return GuiTabItemFlags(ptr.TabItemFlagsOverrideSet)
}

// SetTabItemFlagsOverrideSet sets the value in TabItemFlagsOverrideSet.
func (p GuiWindowClass) SetTabItemFlagsOverrideSet(value GuiTabItemFlags) {
	ptr := (*C.ImGuiWindowClass)(unsafe.Pointer(p))
	ptr.TabItemFlagsOverrideSet = (C.ImGuiTabItemFlags)(value)
}

// GetViewportFlagsOverrideClear returns the value in ViewportFlagsOverrideClear.
func (p GuiWindowClass) GetViewportFlagsOverrideClear() GuiViewportFlags {
	ptr := (*C.ImGuiWindowClass)(unsafe.Pointer(p))
	return GuiViewportFlags(ptr.ViewportFlagsOverrideClear)
}

// SetViewportFlagsOverrideClear sets the value in ViewportFlagsOverrideClear.
func (p GuiWindowClass) SetViewportFlagsOverrideClear(value GuiViewportFlags) {
	ptr := (*C.ImGuiWindowClass)(unsafe.Pointer(p))
	ptr.ViewportFlagsOverrideClear = (C.ImGuiViewportFlags)(value)
}

// GetViewportFlagsOverrideSet returns the value in ViewportFlagsOverrideSet.
func (p GuiWindowClass) GetViewportFlagsOverrideSet() GuiViewportFlags {
	ptr := (*C.ImGuiWindowClass)(unsafe.Pointer(p))
	return GuiViewportFlags(ptr.ViewportFlagsOverrideSet)
}

// SetViewportFlagsOverrideSet sets the value in ViewportFlagsOverrideSet.
func (p GuiWindowClass) SetViewportFlagsOverrideSet(value GuiViewportFlags) {
	ptr := (*C.ImGuiWindowClass)(unsafe.Pointer(p))
	ptr.ViewportFlagsOverrideSet = (C.ImGuiViewportFlags)(value)
}

// GuiWindowDockStyle wraps struct ImGuiWindowDockStyle.
type GuiWindowDockStyle uintptr

// GuiWindowDockStyleNil is a null pointer.
var GuiWindowDockStyleNil GuiWindowDockStyle

// GuiWindowDockStyleSizeOf is the byte size of ImGuiWindowDockStyle.
const GuiWindowDockStyleSizeOf = int(C.sizeof_ImGuiWindowDockStyle)

// GuiWindowDockStyleAlloc allocates a continuous block of GuiWindowDockStyle.
func GuiWindowDockStyleAlloc(alloc ffi.Allocator, count int) GuiWindowDockStyle {
	ptr := alloc.Allocate(GuiWindowDockStyleSizeOf * count)
	return GuiWindowDockStyle(ptr)
}

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p GuiWindowDockStyle) Offset(offset int) GuiWindowDockStyle {
	return p + GuiWindowDockStyle(offset*GuiWindowDockStyleSizeOf)
}

// GuiWindowDockStyle.Colors[ImGuiWindowDockStyleCol_COUNT] is unsupported: category unsupported.

// GuiWindowSettings wraps struct ImGuiWindowSettings.
type GuiWindowSettings uintptr

// GuiWindowSettingsNil is a null pointer.
var GuiWindowSettingsNil GuiWindowSettings

// GuiWindowSettingsSizeOf is the byte size of ImGuiWindowSettings.
const GuiWindowSettingsSizeOf = int(C.sizeof_ImGuiWindowSettings)

// ImGuiWindowSettings allocator is suppressed.

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p GuiWindowSettings) Offset(offset int) GuiWindowSettings {
	return p + GuiWindowSettings(offset*GuiWindowSettingsSizeOf)
}

// GuiWindowSettings.ClassId is unsupported: category unsupported.

// GetCollapsed returns the value in Collapsed.
func (p GuiWindowSettings) GetCollapsed() bool {
	ptr := (*C.ImGuiWindowSettings)(unsafe.Pointer(p))
	return bool(ptr.Collapsed)
}

// SetCollapsed sets the value in Collapsed.
func (p GuiWindowSettings) SetCollapsed(value bool) {
	ptr := (*C.ImGuiWindowSettings)(unsafe.Pointer(p))
	ptr.Collapsed = (C.bool)(value)
}

// GuiWindowSettings.DockId is unsupported: category unsupported.

// GuiWindowSettings.DockOrder is unsupported: category unsupported.

// GuiWindowSettings.ID is unsupported: category unsupported.

// GetIsChild returns the value in IsChild.
func (p GuiWindowSettings) GetIsChild() bool {
	ptr := (*C.ImGuiWindowSettings)(unsafe.Pointer(p))
	return bool(ptr.IsChild)
}

// SetIsChild sets the value in IsChild.
func (p GuiWindowSettings) SetIsChild(value bool) {
	ptr := (*C.ImGuiWindowSettings)(unsafe.Pointer(p))
	ptr.IsChild = (C.bool)(value)
}

// RefPos returns pointer to the Pos field.
func (p GuiWindowSettings) RefPos() Vec2ih {
	return Vec2ih(p + GuiWindowSettings(C.offsetof_ImGuiWindowSettings_Pos))
}

// RefSize returns pointer to the Size field.
func (p GuiWindowSettings) RefSize() Vec2ih {
	return Vec2ih(p + GuiWindowSettings(C.offsetof_ImGuiWindowSettings_Size))
}

// GuiWindowSettings.ViewportId is unsupported: category unsupported.

// RefViewportPos returns pointer to the ViewportPos field.
func (p GuiWindowSettings) RefViewportPos() Vec2ih {
	return Vec2ih(p + GuiWindowSettings(C.offsetof_ImGuiWindowSettings_ViewportPos))
}

// GetWantApply returns the value in WantApply.
func (p GuiWindowSettings) GetWantApply() bool {
	ptr := (*C.ImGuiWindowSettings)(unsafe.Pointer(p))
	return bool(ptr.WantApply)
}

// SetWantApply sets the value in WantApply.
func (p GuiWindowSettings) SetWantApply(value bool) {
	ptr := (*C.ImGuiWindowSettings)(unsafe.Pointer(p))
	ptr.WantApply = (C.bool)(value)
}

// GetWantDelete returns the value in WantDelete.
func (p GuiWindowSettings) GetWantDelete() bool {
	ptr := (*C.ImGuiWindowSettings)(unsafe.Pointer(p))
	return bool(ptr.WantDelete)
}

// SetWantDelete sets the value in WantDelete.
func (p GuiWindowSettings) SetWantDelete(value bool) {
	ptr := (*C.ImGuiWindowSettings)(unsafe.Pointer(p))
	ptr.WantDelete = (C.bool)(value)
}

// GuiWindowStackData wraps struct ImGuiWindowStackData.
type GuiWindowStackData uintptr

// GuiWindowStackDataNil is a null pointer.
var GuiWindowStackDataNil GuiWindowStackData

// GuiWindowStackDataSizeOf is the byte size of ImGuiWindowStackData.
const GuiWindowStackDataSizeOf = int(C.sizeof_ImGuiWindowStackData)

// GuiWindowStackDataAlloc allocates a continuous block of GuiWindowStackData.
func GuiWindowStackDataAlloc(alloc ffi.Allocator, count int) GuiWindowStackData {
	ptr := alloc.Allocate(GuiWindowStackDataSizeOf * count)
	return GuiWindowStackData(ptr)
}

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p GuiWindowStackData) Offset(offset int) GuiWindowStackData {
	return p + GuiWindowStackData(offset*GuiWindowStackDataSizeOf)
}

// GetDisabledOverrideReenable returns the value in DisabledOverrideReenable.
func (p GuiWindowStackData) GetDisabledOverrideReenable() bool {
	ptr := (*C.ImGuiWindowStackData)(unsafe.Pointer(p))
	return bool(ptr.DisabledOverrideReenable)
}

// SetDisabledOverrideReenable sets the value in DisabledOverrideReenable.
func (p GuiWindowStackData) SetDisabledOverrideReenable(value bool) {
	ptr := (*C.ImGuiWindowStackData)(unsafe.Pointer(p))
	ptr.DisabledOverrideReenable = (C.bool)(value)
}

// GetDisabledOverrideReenableAlphaBackup returns the value in DisabledOverrideReenableAlphaBackup.
func (p GuiWindowStackData) GetDisabledOverrideReenableAlphaBackup() float32 {
	ptr := (*C.ImGuiWindowStackData)(unsafe.Pointer(p))
	return float32(ptr.DisabledOverrideReenableAlphaBackup)
}

// SetDisabledOverrideReenableAlphaBackup sets the value in DisabledOverrideReenableAlphaBackup.
func (p GuiWindowStackData) SetDisabledOverrideReenableAlphaBackup(value float32) {
	ptr := (*C.ImGuiWindowStackData)(unsafe.Pointer(p))
	ptr.DisabledOverrideReenableAlphaBackup = (C.float)(value)
}

// RefParentLastItemDataBackup returns pointer to the ParentLastItemDataBackup field.
func (p GuiWindowStackData) RefParentLastItemDataBackup() GuiLastItemData {
	return GuiLastItemData(p + GuiWindowStackData(C.offsetof_ImGuiWindowStackData_ParentLastItemDataBackup))
}

// RefStackSizesInBegin returns pointer to the StackSizesInBegin field.
func (p GuiWindowStackData) RefStackSizesInBegin() GuiErrorRecoveryState {
	return GuiErrorRecoveryState(p + GuiWindowStackData(C.offsetof_ImGuiWindowStackData_StackSizesInBegin))
}

// GetWindow returns the value in Window.
func (p GuiWindowStackData) GetWindow() GuiWindow {
	ptr := (*C.ImGuiWindowStackData)(unsafe.Pointer(p))
	return GuiWindow(unsafe.Pointer(ptr.Window))
}

// SetWindow sets the value in Window.
func (p GuiWindowStackData) SetWindow(value GuiWindow) {
	ptr := (*C.ImGuiWindowStackData)(unsafe.Pointer(p))
	ptr.Window = (*C.ImGuiWindow)(unsafe.Pointer(value))
}

// GuiWindowTempData wraps struct ImGuiWindowTempData.
type GuiWindowTempData uintptr

// GuiWindowTempDataNil is a null pointer.
var GuiWindowTempDataNil GuiWindowTempData

// GuiWindowTempDataSizeOf is the byte size of ImGuiWindowTempData.
const GuiWindowTempDataSizeOf = int(C.sizeof_ImGuiWindowTempData)

// GuiWindowTempDataAlloc allocates a continuous block of GuiWindowTempData.
func GuiWindowTempDataAlloc(alloc ffi.Allocator, count int) GuiWindowTempData {
	ptr := alloc.Allocate(GuiWindowTempDataSizeOf * count)
	return GuiWindowTempData(ptr)
}

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p GuiWindowTempData) Offset(offset int) GuiWindowTempData {
	return p + GuiWindowTempData(offset*GuiWindowTempDataSizeOf)
}

// GetChildItemStatusFlags returns the value in ChildItemStatusFlags.
func (p GuiWindowTempData) GetChildItemStatusFlags() GuiItemStatusFlags {
	ptr := (*C.ImGuiWindowTempData)(unsafe.Pointer(p))
	return GuiItemStatusFlags(ptr.ChildItemStatusFlags)
}

// SetChildItemStatusFlags sets the value in ChildItemStatusFlags.
func (p GuiWindowTempData) SetChildItemStatusFlags(value GuiItemStatusFlags) {
	ptr := (*C.ImGuiWindowTempData)(unsafe.Pointer(p))
	ptr.ChildItemStatusFlags = (C.ImGuiItemStatusFlags)(value)
}

// RefChildWindows returns pointer to the ChildWindows field.
func (p GuiWindowTempData) RefChildWindows() Vector_ImGuiWindowPtr {
	return Vector_ImGuiWindowPtr(p + GuiWindowTempData(C.offsetof_ImGuiWindowTempData_ChildWindows))
}

// RefColumnsOffset returns pointer to the ColumnsOffset field.
func (p GuiWindowTempData) RefColumnsOffset() Vec1 {
	return Vec1(p + GuiWindowTempData(C.offsetof_ImGuiWindowTempData_ColumnsOffset))
}

// RefCurrLineSize returns pointer to the CurrLineSize field.
func (p GuiWindowTempData) RefCurrLineSize() Vec2 {
	return Vec2(p + GuiWindowTempData(C.offsetof_ImGuiWindowTempData_CurrLineSize))
}

// GetCurrLineTextBaseOffset returns the value in CurrLineTextBaseOffset.
func (p GuiWindowTempData) GetCurrLineTextBaseOffset() float32 {
	ptr := (*C.ImGuiWindowTempData)(unsafe.Pointer(p))
	return float32(ptr.CurrLineTextBaseOffset)
}

// SetCurrLineTextBaseOffset sets the value in CurrLineTextBaseOffset.
func (p GuiWindowTempData) SetCurrLineTextBaseOffset(value float32) {
	ptr := (*C.ImGuiWindowTempData)(unsafe.Pointer(p))
	ptr.CurrLineTextBaseOffset = (C.float)(value)
}

// GetCurrentColumns returns the value in CurrentColumns.
func (p GuiWindowTempData) GetCurrentColumns() GuiOldColumns {
	ptr := (*C.ImGuiWindowTempData)(unsafe.Pointer(p))
	return GuiOldColumns(unsafe.Pointer(ptr.CurrentColumns))
}

// SetCurrentColumns sets the value in CurrentColumns.
func (p GuiWindowTempData) SetCurrentColumns(value GuiOldColumns) {
	ptr := (*C.ImGuiWindowTempData)(unsafe.Pointer(p))
	ptr.CurrentColumns = (*C.ImGuiOldColumns)(unsafe.Pointer(value))
}

// GetCurrentTableIdx returns the value in CurrentTableIdx.
func (p GuiWindowTempData) GetCurrentTableIdx() int32 {
	ptr := (*C.ImGuiWindowTempData)(unsafe.Pointer(p))
	return int32(ptr.CurrentTableIdx)
}

// SetCurrentTableIdx sets the value in CurrentTableIdx.
func (p GuiWindowTempData) SetCurrentTableIdx(value int32) {
	ptr := (*C.ImGuiWindowTempData)(unsafe.Pointer(p))
	ptr.CurrentTableIdx = (C.int32_t)(value)
}

// RefCursorMaxPos returns pointer to the CursorMaxPos field.
func (p GuiWindowTempData) RefCursorMaxPos() Vec2 {
	return Vec2(p + GuiWindowTempData(C.offsetof_ImGuiWindowTempData_CursorMaxPos))
}

// RefCursorPos returns pointer to the CursorPos field.
func (p GuiWindowTempData) RefCursorPos() Vec2 {
	return Vec2(p + GuiWindowTempData(C.offsetof_ImGuiWindowTempData_CursorPos))
}

// RefCursorPosPrevLine returns pointer to the CursorPosPrevLine field.
func (p GuiWindowTempData) RefCursorPosPrevLine() Vec2 {
	return Vec2(p + GuiWindowTempData(C.offsetof_ImGuiWindowTempData_CursorPosPrevLine))
}

// RefCursorStartPos returns pointer to the CursorStartPos field.
func (p GuiWindowTempData) RefCursorStartPos() Vec2 {
	return Vec2(p + GuiWindowTempData(C.offsetof_ImGuiWindowTempData_CursorStartPos))
}

// RefCursorStartPosLossyness returns pointer to the CursorStartPosLossyness field.
func (p GuiWindowTempData) RefCursorStartPosLossyness() Vec2 {
	return Vec2(p + GuiWindowTempData(C.offsetof_ImGuiWindowTempData_CursorStartPosLossyness))
}

// RefDockTabItemRect returns pointer to the DockTabItemRect field.
func (p GuiWindowTempData) RefDockTabItemRect() Rect {
	return Rect(p + GuiWindowTempData(C.offsetof_ImGuiWindowTempData_DockTabItemRect))
}

// GetDockTabItemStatusFlags returns the value in DockTabItemStatusFlags.
func (p GuiWindowTempData) GetDockTabItemStatusFlags() GuiItemStatusFlags {
	ptr := (*C.ImGuiWindowTempData)(unsafe.Pointer(p))
	return GuiItemStatusFlags(ptr.DockTabItemStatusFlags)
}

// SetDockTabItemStatusFlags sets the value in DockTabItemStatusFlags.
func (p GuiWindowTempData) SetDockTabItemStatusFlags(value GuiItemStatusFlags) {
	ptr := (*C.ImGuiWindowTempData)(unsafe.Pointer(p))
	ptr.DockTabItemStatusFlags = (C.ImGuiItemStatusFlags)(value)
}

// RefGroupOffset returns pointer to the GroupOffset field.
func (p GuiWindowTempData) RefGroupOffset() Vec1 {
	return Vec1(p + GuiWindowTempData(C.offsetof_ImGuiWindowTempData_GroupOffset))
}

// RefIdealMaxPos returns pointer to the IdealMaxPos field.
func (p GuiWindowTempData) RefIdealMaxPos() Vec2 {
	return Vec2(p + GuiWindowTempData(C.offsetof_ImGuiWindowTempData_IdealMaxPos))
}

// RefIndent returns pointer to the Indent field.
func (p GuiWindowTempData) RefIndent() Vec1 {
	return Vec1(p + GuiWindowTempData(C.offsetof_ImGuiWindowTempData_Indent))
}

// GetIsSameLine returns the value in IsSameLine.
func (p GuiWindowTempData) GetIsSameLine() bool {
	ptr := (*C.ImGuiWindowTempData)(unsafe.Pointer(p))
	return bool(ptr.IsSameLine)
}

// SetIsSameLine sets the value in IsSameLine.
func (p GuiWindowTempData) SetIsSameLine(value bool) {
	ptr := (*C.ImGuiWindowTempData)(unsafe.Pointer(p))
	ptr.IsSameLine = (C.bool)(value)
}

// GetIsSetPos returns the value in IsSetPos.
func (p GuiWindowTempData) GetIsSetPos() bool {
	ptr := (*C.ImGuiWindowTempData)(unsafe.Pointer(p))
	return bool(ptr.IsSetPos)
}

// SetIsSetPos sets the value in IsSetPos.
func (p GuiWindowTempData) SetIsSetPos(value bool) {
	ptr := (*C.ImGuiWindowTempData)(unsafe.Pointer(p))
	ptr.IsSetPos = (C.bool)(value)
}

// GetItemWidth returns the value in ItemWidth.
func (p GuiWindowTempData) GetItemWidth() float32 {
	ptr := (*C.ImGuiWindowTempData)(unsafe.Pointer(p))
	return float32(ptr.ItemWidth)
}

// SetItemWidth sets the value in ItemWidth.
func (p GuiWindowTempData) SetItemWidth(value float32) {
	ptr := (*C.ImGuiWindowTempData)(unsafe.Pointer(p))
	ptr.ItemWidth = (C.float)(value)
}

// RefItemWidthStack returns pointer to the ItemWidthStack field.
func (p GuiWindowTempData) RefItemWidthStack() Vector_float {
	return Vector_float(p + GuiWindowTempData(C.offsetof_ImGuiWindowTempData_ItemWidthStack))
}

// GetLayoutType returns the value in LayoutType.
func (p GuiWindowTempData) GetLayoutType() GuiLayoutType {
	ptr := (*C.ImGuiWindowTempData)(unsafe.Pointer(p))
	return GuiLayoutType(ptr.LayoutType)
}

// SetLayoutType sets the value in LayoutType.
func (p GuiWindowTempData) SetLayoutType(value GuiLayoutType) {
	ptr := (*C.ImGuiWindowTempData)(unsafe.Pointer(p))
	ptr.LayoutType = (C.ImGuiLayoutType)(value)
}

// GetMenuBarAppending returns the value in MenuBarAppending.
func (p GuiWindowTempData) GetMenuBarAppending() bool {
	ptr := (*C.ImGuiWindowTempData)(unsafe.Pointer(p))
	return bool(ptr.MenuBarAppending)
}

// SetMenuBarAppending sets the value in MenuBarAppending.
func (p GuiWindowTempData) SetMenuBarAppending(value bool) {
	ptr := (*C.ImGuiWindowTempData)(unsafe.Pointer(p))
	ptr.MenuBarAppending = (C.bool)(value)
}

// RefMenuBarOffset returns pointer to the MenuBarOffset field.
func (p GuiWindowTempData) RefMenuBarOffset() Vec2 {
	return Vec2(p + GuiWindowTempData(C.offsetof_ImGuiWindowTempData_MenuBarOffset))
}

// RefMenuColumns returns pointer to the MenuColumns field.
func (p GuiWindowTempData) RefMenuColumns() GuiMenuColumns {
	return GuiMenuColumns(p + GuiWindowTempData(C.offsetof_ImGuiWindowTempData_MenuColumns))
}

// GetModalDimBgColor returns the value in ModalDimBgColor.
func (p GuiWindowTempData) GetModalDimBgColor() uint32 {
	ptr := (*C.ImGuiWindowTempData)(unsafe.Pointer(p))
	return uint32(ptr.ModalDimBgColor)
}

// SetModalDimBgColor sets the value in ModalDimBgColor.
func (p GuiWindowTempData) SetModalDimBgColor(value uint32) {
	ptr := (*C.ImGuiWindowTempData)(unsafe.Pointer(p))
	ptr.ModalDimBgColor = (C.uint32_t)(value)
}

// GetNavHideHighlightOneFrame returns the value in NavHideHighlightOneFrame.
func (p GuiWindowTempData) GetNavHideHighlightOneFrame() bool {
	ptr := (*C.ImGuiWindowTempData)(unsafe.Pointer(p))
	return bool(ptr.NavHideHighlightOneFrame)
}

// SetNavHideHighlightOneFrame sets the value in NavHideHighlightOneFrame.
func (p GuiWindowTempData) SetNavHideHighlightOneFrame(value bool) {
	ptr := (*C.ImGuiWindowTempData)(unsafe.Pointer(p))
	ptr.NavHideHighlightOneFrame = (C.bool)(value)
}

// GetNavIsScrollPushableX returns the value in NavIsScrollPushableX.
func (p GuiWindowTempData) GetNavIsScrollPushableX() bool {
	ptr := (*C.ImGuiWindowTempData)(unsafe.Pointer(p))
	return bool(ptr.NavIsScrollPushableX)
}

// SetNavIsScrollPushableX sets the value in NavIsScrollPushableX.
func (p GuiWindowTempData) SetNavIsScrollPushableX(value bool) {
	ptr := (*C.ImGuiWindowTempData)(unsafe.Pointer(p))
	ptr.NavIsScrollPushableX = (C.bool)(value)
}

// GetNavLayerCurrent returns the value in NavLayerCurrent.
func (p GuiWindowTempData) GetNavLayerCurrent() GuiNavLayer {
	ptr := (*C.ImGuiWindowTempData)(unsafe.Pointer(p))
	return GuiNavLayer(ptr.NavLayerCurrent)
}

// SetNavLayerCurrent sets the value in NavLayerCurrent.
func (p GuiWindowTempData) SetNavLayerCurrent(value GuiNavLayer) {
	ptr := (*C.ImGuiWindowTempData)(unsafe.Pointer(p))
	ptr.NavLayerCurrent = (C.ImGuiNavLayer)(value)
}

// GuiWindowTempData.NavLayersActiveMask is unsupported: category unsupported.

// GuiWindowTempData.NavLayersActiveMaskNext is unsupported: category unsupported.

// GetNavWindowHasScrollY returns the value in NavWindowHasScrollY.
func (p GuiWindowTempData) GetNavWindowHasScrollY() bool {
	ptr := (*C.ImGuiWindowTempData)(unsafe.Pointer(p))
	return bool(ptr.NavWindowHasScrollY)
}

// SetNavWindowHasScrollY sets the value in NavWindowHasScrollY.
func (p GuiWindowTempData) SetNavWindowHasScrollY(value bool) {
	ptr := (*C.ImGuiWindowTempData)(unsafe.Pointer(p))
	ptr.NavWindowHasScrollY = (C.bool)(value)
}

// GetParentLayoutType returns the value in ParentLayoutType.
func (p GuiWindowTempData) GetParentLayoutType() GuiLayoutType {
	ptr := (*C.ImGuiWindowTempData)(unsafe.Pointer(p))
	return GuiLayoutType(ptr.ParentLayoutType)
}

// SetParentLayoutType sets the value in ParentLayoutType.
func (p GuiWindowTempData) SetParentLayoutType(value GuiLayoutType) {
	ptr := (*C.ImGuiWindowTempData)(unsafe.Pointer(p))
	ptr.ParentLayoutType = (C.ImGuiLayoutType)(value)
}

// RefPrevLineSize returns pointer to the PrevLineSize field.
func (p GuiWindowTempData) RefPrevLineSize() Vec2 {
	return Vec2(p + GuiWindowTempData(C.offsetof_ImGuiWindowTempData_PrevLineSize))
}

// GetPrevLineTextBaseOffset returns the value in PrevLineTextBaseOffset.
func (p GuiWindowTempData) GetPrevLineTextBaseOffset() float32 {
	ptr := (*C.ImGuiWindowTempData)(unsafe.Pointer(p))
	return float32(ptr.PrevLineTextBaseOffset)
}

// SetPrevLineTextBaseOffset sets the value in PrevLineTextBaseOffset.
func (p GuiWindowTempData) SetPrevLineTextBaseOffset(value float32) {
	ptr := (*C.ImGuiWindowTempData)(unsafe.Pointer(p))
	ptr.PrevLineTextBaseOffset = (C.float)(value)
}

// GetStateStorage returns the value in StateStorage.
func (p GuiWindowTempData) GetStateStorage() GuiStorage {
	ptr := (*C.ImGuiWindowTempData)(unsafe.Pointer(p))
	return GuiStorage(unsafe.Pointer(ptr.StateStorage))
}

// SetStateStorage sets the value in StateStorage.
func (p GuiWindowTempData) SetStateStorage(value GuiStorage) {
	ptr := (*C.ImGuiWindowTempData)(unsafe.Pointer(p))
	ptr.StateStorage = (*C.ImGuiStorage)(unsafe.Pointer(value))
}

// GetTextWrapPos returns the value in TextWrapPos.
func (p GuiWindowTempData) GetTextWrapPos() float32 {
	ptr := (*C.ImGuiWindowTempData)(unsafe.Pointer(p))
	return float32(ptr.TextWrapPos)
}

// SetTextWrapPos sets the value in TextWrapPos.
func (p GuiWindowTempData) SetTextWrapPos(value float32) {
	ptr := (*C.ImGuiWindowTempData)(unsafe.Pointer(p))
	ptr.TextWrapPos = (C.float)(value)
}

// RefTextWrapPosStack returns pointer to the TextWrapPosStack field.
func (p GuiWindowTempData) RefTextWrapPosStack() Vector_float {
	return Vector_float(p + GuiWindowTempData(C.offsetof_ImGuiWindowTempData_TextWrapPosStack))
}

// GetTreeDepth returns the value in TreeDepth.
func (p GuiWindowTempData) GetTreeDepth() int32 {
	ptr := (*C.ImGuiWindowTempData)(unsafe.Pointer(p))
	return int32(ptr.TreeDepth)
}

// SetTreeDepth sets the value in TreeDepth.
func (p GuiWindowTempData) SetTreeDepth(value int32) {
	ptr := (*C.ImGuiWindowTempData)(unsafe.Pointer(p))
	ptr.TreeDepth = (C.int32_t)(value)
}

// GetTreeHasStackDataDepthMask returns the value in TreeHasStackDataDepthMask.
func (p GuiWindowTempData) GetTreeHasStackDataDepthMask() uint32 {
	ptr := (*C.ImGuiWindowTempData)(unsafe.Pointer(p))
	return uint32(ptr.TreeHasStackDataDepthMask)
}

// SetTreeHasStackDataDepthMask sets the value in TreeHasStackDataDepthMask.
func (p GuiWindowTempData) SetTreeHasStackDataDepthMask(value uint32) {
	ptr := (*C.ImGuiWindowTempData)(unsafe.Pointer(p))
	ptr.TreeHasStackDataDepthMask = (C.uint32_t)(value)
}

// GetTreeRecordsClippedNodesY2Mask returns the value in TreeRecordsClippedNodesY2Mask.
func (p GuiWindowTempData) GetTreeRecordsClippedNodesY2Mask() uint32 {
	ptr := (*C.ImGuiWindowTempData)(unsafe.Pointer(p))
	return uint32(ptr.TreeRecordsClippedNodesY2Mask)
}

// SetTreeRecordsClippedNodesY2Mask sets the value in TreeRecordsClippedNodesY2Mask.
func (p GuiWindowTempData) SetTreeRecordsClippedNodesY2Mask(value uint32) {
	ptr := (*C.ImGuiWindowTempData)(unsafe.Pointer(p))
	ptr.TreeRecordsClippedNodesY2Mask = (C.uint32_t)(value)
}

// GetWindowItemStatusFlags returns the value in WindowItemStatusFlags.
func (p GuiWindowTempData) GetWindowItemStatusFlags() GuiItemStatusFlags {
	ptr := (*C.ImGuiWindowTempData)(unsafe.Pointer(p))
	return GuiItemStatusFlags(ptr.WindowItemStatusFlags)
}

// SetWindowItemStatusFlags sets the value in WindowItemStatusFlags.
func (p GuiWindowTempData) SetWindowItemStatusFlags(value GuiItemStatusFlags) {
	ptr := (*C.ImGuiWindowTempData)(unsafe.Pointer(p))
	ptr.WindowItemStatusFlags = (C.ImGuiItemStatusFlags)(value)
}

// Pool_ImGuiMultiSelectState wraps struct ImPool_ImGuiMultiSelectState.
type Pool_ImGuiMultiSelectState uintptr

// Pool_ImGuiMultiSelectStateNil is a null pointer.
var Pool_ImGuiMultiSelectStateNil Pool_ImGuiMultiSelectState

// Pool_ImGuiMultiSelectStateSizeOf is the byte size of ImPool_ImGuiMultiSelectState.
const Pool_ImGuiMultiSelectStateSizeOf = int(C.sizeof_ImPool_ImGuiMultiSelectState)

// Pool_ImGuiMultiSelectStateAlloc allocates a continuous block of Pool_ImGuiMultiSelectState.
func Pool_ImGuiMultiSelectStateAlloc(alloc ffi.Allocator, count int) Pool_ImGuiMultiSelectState {
	ptr := alloc.Allocate(Pool_ImGuiMultiSelectStateSizeOf * count)
	return Pool_ImGuiMultiSelectState(ptr)
}

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p Pool_ImGuiMultiSelectState) Offset(offset int) Pool_ImGuiMultiSelectState {
	return p + Pool_ImGuiMultiSelectState(offset*Pool_ImGuiMultiSelectStateSizeOf)
}

// Pool_ImGuiMultiSelectState.AliveCount is unsupported: category unsupported.

// Pool_ImGuiMultiSelectState.Buf is unsupported: category unsupported.

// Pool_ImGuiMultiSelectState.FreeIdx is unsupported: category unsupported.

// RefMap returns pointer to the Map field.
func (p Pool_ImGuiMultiSelectState) RefMap() GuiStorage {
	return GuiStorage(p + Pool_ImGuiMultiSelectState(C.offsetof_ImPool_ImGuiMultiSelectState_Map))
}

// Pool_ImGuiTabBar wraps struct ImPool_ImGuiTabBar.
type Pool_ImGuiTabBar uintptr

// Pool_ImGuiTabBarNil is a null pointer.
var Pool_ImGuiTabBarNil Pool_ImGuiTabBar

// Pool_ImGuiTabBarSizeOf is the byte size of ImPool_ImGuiTabBar.
const Pool_ImGuiTabBarSizeOf = int(C.sizeof_ImPool_ImGuiTabBar)

// Pool_ImGuiTabBarAlloc allocates a continuous block of Pool_ImGuiTabBar.
func Pool_ImGuiTabBarAlloc(alloc ffi.Allocator, count int) Pool_ImGuiTabBar {
	ptr := alloc.Allocate(Pool_ImGuiTabBarSizeOf * count)
	return Pool_ImGuiTabBar(ptr)
}

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p Pool_ImGuiTabBar) Offset(offset int) Pool_ImGuiTabBar {
	return p + Pool_ImGuiTabBar(offset*Pool_ImGuiTabBarSizeOf)
}

// Pool_ImGuiTabBar.AliveCount is unsupported: category unsupported.

// Pool_ImGuiTabBar.Buf is unsupported: category unsupported.

// Pool_ImGuiTabBar.FreeIdx is unsupported: category unsupported.

// RefMap returns pointer to the Map field.
func (p Pool_ImGuiTabBar) RefMap() GuiStorage {
	return GuiStorage(p + Pool_ImGuiTabBar(C.offsetof_ImPool_ImGuiTabBar_Map))
}

// Pool_ImGuiTable wraps struct ImPool_ImGuiTable.
type Pool_ImGuiTable uintptr

// Pool_ImGuiTableNil is a null pointer.
var Pool_ImGuiTableNil Pool_ImGuiTable

// Pool_ImGuiTableSizeOf is the byte size of ImPool_ImGuiTable.
const Pool_ImGuiTableSizeOf = int(C.sizeof_ImPool_ImGuiTable)

// Pool_ImGuiTableAlloc allocates a continuous block of Pool_ImGuiTable.
func Pool_ImGuiTableAlloc(alloc ffi.Allocator, count int) Pool_ImGuiTable {
	ptr := alloc.Allocate(Pool_ImGuiTableSizeOf * count)
	return Pool_ImGuiTable(ptr)
}

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p Pool_ImGuiTable) Offset(offset int) Pool_ImGuiTable {
	return p + Pool_ImGuiTable(offset*Pool_ImGuiTableSizeOf)
}

// Pool_ImGuiTable.AliveCount is unsupported: category unsupported.

// Pool_ImGuiTable.Buf is unsupported: category unsupported.

// Pool_ImGuiTable.FreeIdx is unsupported: category unsupported.

// RefMap returns pointer to the Map field.
func (p Pool_ImGuiTable) RefMap() GuiStorage {
	return GuiStorage(p + Pool_ImGuiTable(C.offsetof_ImPool_ImGuiTable_Map))
}

// Rect wraps struct ImRect.
type Rect uintptr

// RectNil is a null pointer.
var RectNil Rect

// RectSizeOf is the byte size of ImRect.
const RectSizeOf = int(C.sizeof_ImRect)

// ImRect allocator is suppressed.

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p Rect) Offset(offset int) Rect {
	return p + Rect(offset*RectSizeOf)
}

// RefMax returns pointer to the Max field.
func (p Rect) RefMax() Vec2 {
	return Vec2(p + Rect(C.offsetof_ImRect_Max))
}

// RefMin returns pointer to the Min field.
func (p Rect) RefMin() Vec2 {
	return Vec2(p + Rect(C.offsetof_ImRect_Min))
}

// Span_ImGuiTableCellData wraps struct ImSpan_ImGuiTableCellData.
type Span_ImGuiTableCellData uintptr

// Span_ImGuiTableCellDataNil is a null pointer.
var Span_ImGuiTableCellDataNil Span_ImGuiTableCellData

// Span_ImGuiTableCellDataSizeOf is the byte size of ImSpan_ImGuiTableCellData.
const Span_ImGuiTableCellDataSizeOf = int(C.sizeof_ImSpan_ImGuiTableCellData)

// Span_ImGuiTableCellDataAlloc allocates a continuous block of Span_ImGuiTableCellData.
func Span_ImGuiTableCellDataAlloc(alloc ffi.Allocator, count int) Span_ImGuiTableCellData {
	ptr := alloc.Allocate(Span_ImGuiTableCellDataSizeOf * count)
	return Span_ImGuiTableCellData(ptr)
}

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p Span_ImGuiTableCellData) Offset(offset int) Span_ImGuiTableCellData {
	return p + Span_ImGuiTableCellData(offset*Span_ImGuiTableCellDataSizeOf)
}

// GetData returns the value in Data.
func (p Span_ImGuiTableCellData) GetData() GuiTableCellData {
	ptr := (*C.ImSpan_ImGuiTableCellData)(unsafe.Pointer(p))
	return GuiTableCellData(unsafe.Pointer(ptr.Data))
}

// SetData sets the value in Data.
func (p Span_ImGuiTableCellData) SetData(value GuiTableCellData) {
	ptr := (*C.ImSpan_ImGuiTableCellData)(unsafe.Pointer(p))
	ptr.Data = (*C.ImGuiTableCellData)(unsafe.Pointer(value))
}

// GetDataEnd returns the value in DataEnd.
func (p Span_ImGuiTableCellData) GetDataEnd() GuiTableCellData {
	ptr := (*C.ImSpan_ImGuiTableCellData)(unsafe.Pointer(p))
	return GuiTableCellData(unsafe.Pointer(ptr.DataEnd))
}

// SetDataEnd sets the value in DataEnd.
func (p Span_ImGuiTableCellData) SetDataEnd(value GuiTableCellData) {
	ptr := (*C.ImSpan_ImGuiTableCellData)(unsafe.Pointer(p))
	ptr.DataEnd = (*C.ImGuiTableCellData)(unsafe.Pointer(value))
}

// Span_ImGuiTableColumn wraps struct ImSpan_ImGuiTableColumn.
type Span_ImGuiTableColumn uintptr

// Span_ImGuiTableColumnNil is a null pointer.
var Span_ImGuiTableColumnNil Span_ImGuiTableColumn

// Span_ImGuiTableColumnSizeOf is the byte size of ImSpan_ImGuiTableColumn.
const Span_ImGuiTableColumnSizeOf = int(C.sizeof_ImSpan_ImGuiTableColumn)

// Span_ImGuiTableColumnAlloc allocates a continuous block of Span_ImGuiTableColumn.
func Span_ImGuiTableColumnAlloc(alloc ffi.Allocator, count int) Span_ImGuiTableColumn {
	ptr := alloc.Allocate(Span_ImGuiTableColumnSizeOf * count)
	return Span_ImGuiTableColumn(ptr)
}

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p Span_ImGuiTableColumn) Offset(offset int) Span_ImGuiTableColumn {
	return p + Span_ImGuiTableColumn(offset*Span_ImGuiTableColumnSizeOf)
}

// GetData returns the value in Data.
func (p Span_ImGuiTableColumn) GetData() GuiTableColumn {
	ptr := (*C.ImSpan_ImGuiTableColumn)(unsafe.Pointer(p))
	return GuiTableColumn(unsafe.Pointer(ptr.Data))
}

// SetData sets the value in Data.
func (p Span_ImGuiTableColumn) SetData(value GuiTableColumn) {
	ptr := (*C.ImSpan_ImGuiTableColumn)(unsafe.Pointer(p))
	ptr.Data = (*C.ImGuiTableColumn)(unsafe.Pointer(value))
}

// GetDataEnd returns the value in DataEnd.
func (p Span_ImGuiTableColumn) GetDataEnd() GuiTableColumn {
	ptr := (*C.ImSpan_ImGuiTableColumn)(unsafe.Pointer(p))
	return GuiTableColumn(unsafe.Pointer(ptr.DataEnd))
}

// SetDataEnd sets the value in DataEnd.
func (p Span_ImGuiTableColumn) SetDataEnd(value GuiTableColumn) {
	ptr := (*C.ImSpan_ImGuiTableColumn)(unsafe.Pointer(p))
	ptr.DataEnd = (*C.ImGuiTableColumn)(unsafe.Pointer(value))
}

// Span_ImGuiTableColumnIdx wraps struct ImSpan_ImGuiTableColumnIdx.
type Span_ImGuiTableColumnIdx uintptr

// Span_ImGuiTableColumnIdxNil is a null pointer.
var Span_ImGuiTableColumnIdxNil Span_ImGuiTableColumnIdx

// Span_ImGuiTableColumnIdxSizeOf is the byte size of ImSpan_ImGuiTableColumnIdx.
const Span_ImGuiTableColumnIdxSizeOf = int(C.sizeof_ImSpan_ImGuiTableColumnIdx)

// Span_ImGuiTableColumnIdxAlloc allocates a continuous block of Span_ImGuiTableColumnIdx.
func Span_ImGuiTableColumnIdxAlloc(alloc ffi.Allocator, count int) Span_ImGuiTableColumnIdx {
	ptr := alloc.Allocate(Span_ImGuiTableColumnIdxSizeOf * count)
	return Span_ImGuiTableColumnIdx(ptr)
}

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p Span_ImGuiTableColumnIdx) Offset(offset int) Span_ImGuiTableColumnIdx {
	return p + Span_ImGuiTableColumnIdx(offset*Span_ImGuiTableColumnIdxSizeOf)
}

// Span_ImGuiTableColumnIdx.Data is unsupported: category unsupported.

// Span_ImGuiTableColumnIdx.DataEnd is unsupported: category unsupported.

// TextureData wraps struct ImTextureData.
type TextureData uintptr

// TextureDataNil is a null pointer.
var TextureDataNil TextureData

// TextureDataSizeOf is the byte size of ImTextureData.
const TextureDataSizeOf = int(C.sizeof_ImTextureData)

// ImTextureData allocator is suppressed.

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p TextureData) Offset(offset int) TextureData {
	return p + TextureData(offset*TextureDataSizeOf)
}

// GetBackendUserData returns the value in BackendUserData.
func (p TextureData) GetBackendUserData() uintptr {
	ptr := (*C.ImTextureData)(unsafe.Pointer(p))
	return uintptr(unsafe.Pointer(ptr.BackendUserData))
}

// SetBackendUserData sets the value in BackendUserData.
func (p TextureData) SetBackendUserData(value uintptr) {
	ptr := (*C.ImTextureData)(unsafe.Pointer(p))
	ptr.BackendUserData = unsafe.Pointer(value)
}

// GetBytesPerPixel returns the value in BytesPerPixel.
func (p TextureData) GetBytesPerPixel() int32 {
	ptr := (*C.ImTextureData)(unsafe.Pointer(p))
	return int32(ptr.BytesPerPixel)
}

// SetBytesPerPixel sets the value in BytesPerPixel.
func (p TextureData) SetBytesPerPixel(value int32) {
	ptr := (*C.ImTextureData)(unsafe.Pointer(p))
	ptr.BytesPerPixel = (C.int32_t)(value)
}

// GetFormat returns the value in Format.
func (p TextureData) GetFormat() TextureFormat {
	ptr := (*C.ImTextureData)(unsafe.Pointer(p))
	return TextureFormat(ptr.Format)
}

// SetFormat sets the value in Format.
func (p TextureData) SetFormat(value TextureFormat) {
	ptr := (*C.ImTextureData)(unsafe.Pointer(p))
	ptr.Format = (C.ImTextureFormat)(value)
}

// GetHeight returns the value in Height.
func (p TextureData) GetHeight() int32 {
	ptr := (*C.ImTextureData)(unsafe.Pointer(p))
	return int32(ptr.Height)
}

// SetHeight sets the value in Height.
func (p TextureData) SetHeight(value int32) {
	ptr := (*C.ImTextureData)(unsafe.Pointer(p))
	ptr.Height = (C.int32_t)(value)
}

// GetPixels returns the value in Pixels.
func (p TextureData) GetPixels() ffi.Ref[uint8] {
	ptr := (*C.ImTextureData)(unsafe.Pointer(p))
	return ffi.RefFromPtr[uint8](unsafe.Pointer(ptr.Pixels))
}

// SetPixels sets the value in Pixels.
func (p TextureData) SetPixels(value ffi.Ref[uint8]) {
	ptr := (*C.ImTextureData)(unsafe.Pointer(p))
	ptr.Pixels = (*C.uint8_t)(value.Raw())
}

// GetRefCount returns the value in RefCount.
func (p TextureData) GetRefCount() uint16 {
	ptr := (*C.ImTextureData)(unsafe.Pointer(p))
	return uint16(ptr.RefCount)
}

// SetRefCount sets the value in RefCount.
func (p TextureData) SetRefCount(value uint16) {
	ptr := (*C.ImTextureData)(unsafe.Pointer(p))
	ptr.RefCount = (C.uint16_t)(value)
}

// GetStatus returns the value in Status.
func (p TextureData) GetStatus() TextureStatus {
	ptr := (*C.ImTextureData)(unsafe.Pointer(p))
	return TextureStatus(ptr.Status)
}

// TextureData.Status getter is suppressed.

// TextureData.TexID is unsupported: category unsupported.

// GetUniqueID returns the value in UniqueID.
func (p TextureData) GetUniqueID() int32 {
	ptr := (*C.ImTextureData)(unsafe.Pointer(p))
	return int32(ptr.UniqueID)
}

// SetUniqueID sets the value in UniqueID.
func (p TextureData) SetUniqueID(value int32) {
	ptr := (*C.ImTextureData)(unsafe.Pointer(p))
	ptr.UniqueID = (C.int32_t)(value)
}

// GetUnusedFrames returns the value in UnusedFrames.
func (p TextureData) GetUnusedFrames() int32 {
	ptr := (*C.ImTextureData)(unsafe.Pointer(p))
	return int32(ptr.UnusedFrames)
}

// SetUnusedFrames sets the value in UnusedFrames.
func (p TextureData) SetUnusedFrames(value int32) {
	ptr := (*C.ImTextureData)(unsafe.Pointer(p))
	ptr.UnusedFrames = (C.int32_t)(value)
}

// RefUpdateRect returns pointer to the UpdateRect field.
func (p TextureData) RefUpdateRect() TextureRect {
	return TextureRect(p + TextureData(C.offsetof_ImTextureData_UpdateRect))
}

// RefUpdates returns pointer to the Updates field.
func (p TextureData) RefUpdates() Vector_ImTextureRect {
	return Vector_ImTextureRect(p + TextureData(C.offsetof_ImTextureData_Updates))
}

// GetUseColors returns the value in UseColors.
func (p TextureData) GetUseColors() bool {
	ptr := (*C.ImTextureData)(unsafe.Pointer(p))
	return bool(ptr.UseColors)
}

// SetUseColors sets the value in UseColors.
func (p TextureData) SetUseColors(value bool) {
	ptr := (*C.ImTextureData)(unsafe.Pointer(p))
	ptr.UseColors = (C.bool)(value)
}

// RefUsedRect returns pointer to the UsedRect field.
func (p TextureData) RefUsedRect() TextureRect {
	return TextureRect(p + TextureData(C.offsetof_ImTextureData_UsedRect))
}

// GetWantDestroyNextFrame returns the value in WantDestroyNextFrame.
func (p TextureData) GetWantDestroyNextFrame() bool {
	ptr := (*C.ImTextureData)(unsafe.Pointer(p))
	return bool(ptr.WantDestroyNextFrame)
}

// SetWantDestroyNextFrame sets the value in WantDestroyNextFrame.
func (p TextureData) SetWantDestroyNextFrame(value bool) {
	ptr := (*C.ImTextureData)(unsafe.Pointer(p))
	ptr.WantDestroyNextFrame = (C.bool)(value)
}

// GetWidth returns the value in Width.
func (p TextureData) GetWidth() int32 {
	ptr := (*C.ImTextureData)(unsafe.Pointer(p))
	return int32(ptr.Width)
}

// SetWidth sets the value in Width.
func (p TextureData) SetWidth(value int32) {
	ptr := (*C.ImTextureData)(unsafe.Pointer(p))
	ptr.Width = (C.int32_t)(value)
}

// TextureRect wraps struct ImTextureRect.
type TextureRect uintptr

// TextureRectNil is a null pointer.
var TextureRectNil TextureRect

// TextureRectSizeOf is the byte size of ImTextureRect.
const TextureRectSizeOf = int(C.sizeof_ImTextureRect)

// TextureRectAlloc allocates a continuous block of TextureRect.
func TextureRectAlloc(alloc ffi.Allocator, count int) TextureRect {
	ptr := alloc.Allocate(TextureRectSizeOf * count)
	return TextureRect(ptr)
}

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p TextureRect) Offset(offset int) TextureRect {
	return p + TextureRect(offset*TextureRectSizeOf)
}

// GetH returns the value in h.
func (p TextureRect) GetH() uint16 {
	ptr := (*C.ImTextureRect)(unsafe.Pointer(p))
	return uint16(ptr.h)
}

// SetH sets the value in h.
func (p TextureRect) SetH(value uint16) {
	ptr := (*C.ImTextureRect)(unsafe.Pointer(p))
	ptr.h = (C.uint16_t)(value)
}

// GetW returns the value in w.
func (p TextureRect) GetW() uint16 {
	ptr := (*C.ImTextureRect)(unsafe.Pointer(p))
	return uint16(ptr.w)
}

// SetW sets the value in w.
func (p TextureRect) SetW(value uint16) {
	ptr := (*C.ImTextureRect)(unsafe.Pointer(p))
	ptr.w = (C.uint16_t)(value)
}

// GetX returns the value in x.
func (p TextureRect) GetX() uint16 {
	ptr := (*C.ImTextureRect)(unsafe.Pointer(p))
	return uint16(ptr.x)
}

// SetX sets the value in x.
func (p TextureRect) SetX(value uint16) {
	ptr := (*C.ImTextureRect)(unsafe.Pointer(p))
	ptr.x = (C.uint16_t)(value)
}

// GetY returns the value in y.
func (p TextureRect) GetY() uint16 {
	ptr := (*C.ImTextureRect)(unsafe.Pointer(p))
	return uint16(ptr.y)
}

// SetY sets the value in y.
func (p TextureRect) SetY(value uint16) {
	ptr := (*C.ImTextureRect)(unsafe.Pointer(p))
	ptr.y = (C.uint16_t)(value)
}

// TextureRef wraps struct ImTextureRef.
type TextureRef uintptr

// TextureRefNil is a null pointer.
var TextureRefNil TextureRef

// TextureRefSizeOf is the byte size of ImTextureRef.
const TextureRefSizeOf = int(C.sizeof_ImTextureRef)

// ImTextureRef allocator is suppressed.

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p TextureRef) Offset(offset int) TextureRef {
	return p + TextureRef(offset*TextureRefSizeOf)
}

// GetTexData returns the value in _TexData.
func (p TextureRef) GetTexData() TextureData {
	ptr := (*C.ImTextureRef)(unsafe.Pointer(p))
	return TextureData(unsafe.Pointer(ptr._TexData))
}

// SetTexData sets the value in _TexData.
func (p TextureRef) SetTexData(value TextureData) {
	ptr := (*C.ImTextureRef)(unsafe.Pointer(p))
	ptr._TexData = (*C.ImTextureData)(unsafe.Pointer(value))
}

// TextureRef._TexID is unsupported: category unsupported.

// Vec1 wraps struct ImVec1.
type Vec1 uintptr

// Vec1Nil is a null pointer.
var Vec1Nil Vec1

// Vec1SizeOf is the byte size of ImVec1.
const Vec1SizeOf = int(C.sizeof_ImVec1)

// ImVec1 allocator is suppressed.

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p Vec1) Offset(offset int) Vec1 {
	return p + Vec1(offset*Vec1SizeOf)
}

// GetX returns the value in x.
func (p Vec1) GetX() float32 {
	ptr := (*C.ImVec1)(unsafe.Pointer(p))
	return float32(ptr.x)
}

// SetX sets the value in x.
func (p Vec1) SetX(value float32) {
	ptr := (*C.ImVec1)(unsafe.Pointer(p))
	ptr.x = (C.float)(value)
}

// Vec2 wraps struct ImVec2.
type Vec2 uintptr

// Vec2Nil is a null pointer.
var Vec2Nil Vec2

// Vec2SizeOf is the byte size of ImVec2.
const Vec2SizeOf = int(C.sizeof_ImVec2)

// ImVec2 allocator is suppressed.

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p Vec2) Offset(offset int) Vec2 {
	return p + Vec2(offset*Vec2SizeOf)
}

// GetX returns the value in x.
func (p Vec2) GetX() float32 {
	ptr := (*C.ImVec2)(unsafe.Pointer(p))
	return float32(ptr.x)
}

// SetX sets the value in x.
func (p Vec2) SetX(value float32) {
	ptr := (*C.ImVec2)(unsafe.Pointer(p))
	ptr.x = (C.float)(value)
}

// GetY returns the value in y.
func (p Vec2) GetY() float32 {
	ptr := (*C.ImVec2)(unsafe.Pointer(p))
	return float32(ptr.y)
}

// SetY sets the value in y.
func (p Vec2) SetY(value float32) {
	ptr := (*C.ImVec2)(unsafe.Pointer(p))
	ptr.y = (C.float)(value)
}

// Vec2i wraps struct ImVec2i.
type Vec2i uintptr

// Vec2iNil is a null pointer.
var Vec2iNil Vec2i

// Vec2iSizeOf is the byte size of ImVec2i.
const Vec2iSizeOf = int(C.sizeof_ImVec2i)

// ImVec2i allocator is suppressed.

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p Vec2i) Offset(offset int) Vec2i {
	return p + Vec2i(offset*Vec2iSizeOf)
}

// GetX returns the value in x.
func (p Vec2i) GetX() int32 {
	ptr := (*C.ImVec2i)(unsafe.Pointer(p))
	return int32(ptr.x)
}

// SetX sets the value in x.
func (p Vec2i) SetX(value int32) {
	ptr := (*C.ImVec2i)(unsafe.Pointer(p))
	ptr.x = (C.int32_t)(value)
}

// GetY returns the value in y.
func (p Vec2i) GetY() int32 {
	ptr := (*C.ImVec2i)(unsafe.Pointer(p))
	return int32(ptr.y)
}

// SetY sets the value in y.
func (p Vec2i) SetY(value int32) {
	ptr := (*C.ImVec2i)(unsafe.Pointer(p))
	ptr.y = (C.int32_t)(value)
}

// Vec2ih wraps struct ImVec2ih.
type Vec2ih uintptr

// Vec2ihNil is a null pointer.
var Vec2ihNil Vec2ih

// Vec2ihSizeOf is the byte size of ImVec2ih.
const Vec2ihSizeOf = int(C.sizeof_ImVec2ih)

// ImVec2ih allocator is suppressed.

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p Vec2ih) Offset(offset int) Vec2ih {
	return p + Vec2ih(offset*Vec2ihSizeOf)
}

// Vec2ih.x is unsupported: category unsupported.

// Vec2ih.y is unsupported: category unsupported.

// Vec4 wraps struct ImVec4.
type Vec4 uintptr

// Vec4Nil is a null pointer.
var Vec4Nil Vec4

// Vec4SizeOf is the byte size of ImVec4.
const Vec4SizeOf = int(C.sizeof_ImVec4)

// ImVec4 allocator is suppressed.

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p Vec4) Offset(offset int) Vec4 {
	return p + Vec4(offset*Vec4SizeOf)
}

// GetW returns the value in w.
func (p Vec4) GetW() float32 {
	ptr := (*C.ImVec4)(unsafe.Pointer(p))
	return float32(ptr.w)
}

// SetW sets the value in w.
func (p Vec4) SetW(value float32) {
	ptr := (*C.ImVec4)(unsafe.Pointer(p))
	ptr.w = (C.float)(value)
}

// GetX returns the value in x.
func (p Vec4) GetX() float32 {
	ptr := (*C.ImVec4)(unsafe.Pointer(p))
	return float32(ptr.x)
}

// SetX sets the value in x.
func (p Vec4) SetX(value float32) {
	ptr := (*C.ImVec4)(unsafe.Pointer(p))
	ptr.x = (C.float)(value)
}

// GetY returns the value in y.
func (p Vec4) GetY() float32 {
	ptr := (*C.ImVec4)(unsafe.Pointer(p))
	return float32(ptr.y)
}

// SetY sets the value in y.
func (p Vec4) SetY(value float32) {
	ptr := (*C.ImVec4)(unsafe.Pointer(p))
	ptr.y = (C.float)(value)
}

// GetZ returns the value in z.
func (p Vec4) GetZ() float32 {
	ptr := (*C.ImVec4)(unsafe.Pointer(p))
	return float32(ptr.z)
}

// SetZ sets the value in z.
func (p Vec4) SetZ(value float32) {
	ptr := (*C.ImVec4)(unsafe.Pointer(p))
	ptr.z = (C.float)(value)
}

// Vector_ImDrawChannel wraps struct ImVector_ImDrawChannel.
type Vector_ImDrawChannel uintptr

// Vector_ImDrawChannelNil is a null pointer.
var Vector_ImDrawChannelNil Vector_ImDrawChannel

// Vector_ImDrawChannelSizeOf is the byte size of ImVector_ImDrawChannel.
const Vector_ImDrawChannelSizeOf = int(C.sizeof_ImVector_ImDrawChannel)

// Vector_ImDrawChannelAlloc allocates a continuous block of Vector_ImDrawChannel.
func Vector_ImDrawChannelAlloc(alloc ffi.Allocator, count int) Vector_ImDrawChannel {
	ptr := alloc.Allocate(Vector_ImDrawChannelSizeOf * count)
	return Vector_ImDrawChannel(ptr)
}

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p Vector_ImDrawChannel) Offset(offset int) Vector_ImDrawChannel {
	return p + Vector_ImDrawChannel(offset*Vector_ImDrawChannelSizeOf)
}

// GetCapacity returns the value in Capacity.
func (p Vector_ImDrawChannel) GetCapacity() int32 {
	ptr := (*C.ImVector_ImDrawChannel)(unsafe.Pointer(p))
	return int32(ptr.Capacity)
}

// SetCapacity sets the value in Capacity.
func (p Vector_ImDrawChannel) SetCapacity(value int32) {
	ptr := (*C.ImVector_ImDrawChannel)(unsafe.Pointer(p))
	ptr.Capacity = (C.int32_t)(value)
}

// GetData returns the value in Data.
func (p Vector_ImDrawChannel) GetData() DrawChannel {
	ptr := (*C.ImVector_ImDrawChannel)(unsafe.Pointer(p))
	return DrawChannel(unsafe.Pointer(ptr.Data))
}

// SetData sets the value in Data.
func (p Vector_ImDrawChannel) SetData(value DrawChannel) {
	ptr := (*C.ImVector_ImDrawChannel)(unsafe.Pointer(p))
	ptr.Data = (*C.ImDrawChannel)(unsafe.Pointer(value))
}

// GetSize returns the value in Size.
func (p Vector_ImDrawChannel) GetSize() int32 {
	ptr := (*C.ImVector_ImDrawChannel)(unsafe.Pointer(p))
	return int32(ptr.Size)
}

// SetSize sets the value in Size.
func (p Vector_ImDrawChannel) SetSize(value int32) {
	ptr := (*C.ImVector_ImDrawChannel)(unsafe.Pointer(p))
	ptr.Size = (C.int32_t)(value)
}

// Vector_ImDrawCmd wraps struct ImVector_ImDrawCmd.
type Vector_ImDrawCmd uintptr

// Vector_ImDrawCmdNil is a null pointer.
var Vector_ImDrawCmdNil Vector_ImDrawCmd

// Vector_ImDrawCmdSizeOf is the byte size of ImVector_ImDrawCmd.
const Vector_ImDrawCmdSizeOf = int(C.sizeof_ImVector_ImDrawCmd)

// Vector_ImDrawCmdAlloc allocates a continuous block of Vector_ImDrawCmd.
func Vector_ImDrawCmdAlloc(alloc ffi.Allocator, count int) Vector_ImDrawCmd {
	ptr := alloc.Allocate(Vector_ImDrawCmdSizeOf * count)
	return Vector_ImDrawCmd(ptr)
}

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p Vector_ImDrawCmd) Offset(offset int) Vector_ImDrawCmd {
	return p + Vector_ImDrawCmd(offset*Vector_ImDrawCmdSizeOf)
}

// GetCapacity returns the value in Capacity.
func (p Vector_ImDrawCmd) GetCapacity() int32 {
	ptr := (*C.ImVector_ImDrawCmd)(unsafe.Pointer(p))
	return int32(ptr.Capacity)
}

// SetCapacity sets the value in Capacity.
func (p Vector_ImDrawCmd) SetCapacity(value int32) {
	ptr := (*C.ImVector_ImDrawCmd)(unsafe.Pointer(p))
	ptr.Capacity = (C.int32_t)(value)
}

// GetData returns the value in Data.
func (p Vector_ImDrawCmd) GetData() DrawCmd {
	ptr := (*C.ImVector_ImDrawCmd)(unsafe.Pointer(p))
	return DrawCmd(unsafe.Pointer(ptr.Data))
}

// SetData sets the value in Data.
func (p Vector_ImDrawCmd) SetData(value DrawCmd) {
	ptr := (*C.ImVector_ImDrawCmd)(unsafe.Pointer(p))
	ptr.Data = (*C.ImDrawCmd)(unsafe.Pointer(value))
}

// GetSize returns the value in Size.
func (p Vector_ImDrawCmd) GetSize() int32 {
	ptr := (*C.ImVector_ImDrawCmd)(unsafe.Pointer(p))
	return int32(ptr.Size)
}

// SetSize sets the value in Size.
func (p Vector_ImDrawCmd) SetSize(value int32) {
	ptr := (*C.ImVector_ImDrawCmd)(unsafe.Pointer(p))
	ptr.Size = (C.int32_t)(value)
}

// Vector_ImDrawIdx wraps struct ImVector_ImDrawIdx.
type Vector_ImDrawIdx uintptr

// Vector_ImDrawIdxNil is a null pointer.
var Vector_ImDrawIdxNil Vector_ImDrawIdx

// Vector_ImDrawIdxSizeOf is the byte size of ImVector_ImDrawIdx.
const Vector_ImDrawIdxSizeOf = int(C.sizeof_ImVector_ImDrawIdx)

// Vector_ImDrawIdxAlloc allocates a continuous block of Vector_ImDrawIdx.
func Vector_ImDrawIdxAlloc(alloc ffi.Allocator, count int) Vector_ImDrawIdx {
	ptr := alloc.Allocate(Vector_ImDrawIdxSizeOf * count)
	return Vector_ImDrawIdx(ptr)
}

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p Vector_ImDrawIdx) Offset(offset int) Vector_ImDrawIdx {
	return p + Vector_ImDrawIdx(offset*Vector_ImDrawIdxSizeOf)
}

// GetCapacity returns the value in Capacity.
func (p Vector_ImDrawIdx) GetCapacity() int32 {
	ptr := (*C.ImVector_ImDrawIdx)(unsafe.Pointer(p))
	return int32(ptr.Capacity)
}

// SetCapacity sets the value in Capacity.
func (p Vector_ImDrawIdx) SetCapacity(value int32) {
	ptr := (*C.ImVector_ImDrawIdx)(unsafe.Pointer(p))
	ptr.Capacity = (C.int32_t)(value)
}

// GetData returns the value in Data.
func (p Vector_ImDrawIdx) GetData() ffi.Ref[DrawIdx] {
	ptr := (*C.ImVector_ImDrawIdx)(unsafe.Pointer(p))
	return ffi.RefFromPtr[DrawIdx](unsafe.Pointer(ptr.Data))
}

// SetData sets the value in Data.
func (p Vector_ImDrawIdx) SetData(value ffi.Ref[DrawIdx]) {
	ptr := (*C.ImVector_ImDrawIdx)(unsafe.Pointer(p))
	ptr.Data = (*C.ImDrawIdx)(value.Raw())
}

// GetSize returns the value in Size.
func (p Vector_ImDrawIdx) GetSize() int32 {
	ptr := (*C.ImVector_ImDrawIdx)(unsafe.Pointer(p))
	return int32(ptr.Size)
}

// SetSize sets the value in Size.
func (p Vector_ImDrawIdx) SetSize(value int32) {
	ptr := (*C.ImVector_ImDrawIdx)(unsafe.Pointer(p))
	ptr.Size = (C.int32_t)(value)
}

// Vector_ImDrawListPtr wraps struct ImVector_ImDrawListPtr.
type Vector_ImDrawListPtr uintptr

// Vector_ImDrawListPtrNil is a null pointer.
var Vector_ImDrawListPtrNil Vector_ImDrawListPtr

// Vector_ImDrawListPtrSizeOf is the byte size of ImVector_ImDrawListPtr.
const Vector_ImDrawListPtrSizeOf = int(C.sizeof_ImVector_ImDrawListPtr)

// Vector_ImDrawListPtrAlloc allocates a continuous block of Vector_ImDrawListPtr.
func Vector_ImDrawListPtrAlloc(alloc ffi.Allocator, count int) Vector_ImDrawListPtr {
	ptr := alloc.Allocate(Vector_ImDrawListPtrSizeOf * count)
	return Vector_ImDrawListPtr(ptr)
}

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p Vector_ImDrawListPtr) Offset(offset int) Vector_ImDrawListPtr {
	return p + Vector_ImDrawListPtr(offset*Vector_ImDrawListPtrSizeOf)
}

// GetCapacity returns the value in Capacity.
func (p Vector_ImDrawListPtr) GetCapacity() int32 {
	ptr := (*C.ImVector_ImDrawListPtr)(unsafe.Pointer(p))
	return int32(ptr.Capacity)
}

// SetCapacity sets the value in Capacity.
func (p Vector_ImDrawListPtr) SetCapacity(value int32) {
	ptr := (*C.ImVector_ImDrawListPtr)(unsafe.Pointer(p))
	ptr.Capacity = (C.int32_t)(value)
}

// GetData returns the value in Data.
func (p Vector_ImDrawListPtr) GetData() ffi.Ref[DrawList] {
	ptr := (*C.ImVector_ImDrawListPtr)(unsafe.Pointer(p))
	return ffi.RefFromPtr[DrawList](unsafe.Pointer(ptr.Data))
}

// SetData sets the value in Data.
func (p Vector_ImDrawListPtr) SetData(value ffi.Ref[DrawList]) {
	ptr := (*C.ImVector_ImDrawListPtr)(unsafe.Pointer(p))
	ptr.Data = (**C.ImDrawList)(value.Raw())
}

// GetSize returns the value in Size.
func (p Vector_ImDrawListPtr) GetSize() int32 {
	ptr := (*C.ImVector_ImDrawListPtr)(unsafe.Pointer(p))
	return int32(ptr.Size)
}

// SetSize sets the value in Size.
func (p Vector_ImDrawListPtr) SetSize(value int32) {
	ptr := (*C.ImVector_ImDrawListPtr)(unsafe.Pointer(p))
	ptr.Size = (C.int32_t)(value)
}

// Vector_ImDrawListSharedDataPtr wraps struct ImVector_ImDrawListSharedDataPtr.
type Vector_ImDrawListSharedDataPtr uintptr

// Vector_ImDrawListSharedDataPtrNil is a null pointer.
var Vector_ImDrawListSharedDataPtrNil Vector_ImDrawListSharedDataPtr

// Vector_ImDrawListSharedDataPtrSizeOf is the byte size of ImVector_ImDrawListSharedDataPtr.
const Vector_ImDrawListSharedDataPtrSizeOf = int(C.sizeof_ImVector_ImDrawListSharedDataPtr)

// Vector_ImDrawListSharedDataPtrAlloc allocates a continuous block of Vector_ImDrawListSharedDataPtr.
func Vector_ImDrawListSharedDataPtrAlloc(alloc ffi.Allocator, count int) Vector_ImDrawListSharedDataPtr {
	ptr := alloc.Allocate(Vector_ImDrawListSharedDataPtrSizeOf * count)
	return Vector_ImDrawListSharedDataPtr(ptr)
}

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p Vector_ImDrawListSharedDataPtr) Offset(offset int) Vector_ImDrawListSharedDataPtr {
	return p + Vector_ImDrawListSharedDataPtr(offset*Vector_ImDrawListSharedDataPtrSizeOf)
}

// GetCapacity returns the value in Capacity.
func (p Vector_ImDrawListSharedDataPtr) GetCapacity() int32 {
	ptr := (*C.ImVector_ImDrawListSharedDataPtr)(unsafe.Pointer(p))
	return int32(ptr.Capacity)
}

// SetCapacity sets the value in Capacity.
func (p Vector_ImDrawListSharedDataPtr) SetCapacity(value int32) {
	ptr := (*C.ImVector_ImDrawListSharedDataPtr)(unsafe.Pointer(p))
	ptr.Capacity = (C.int32_t)(value)
}

// GetData returns the value in Data.
func (p Vector_ImDrawListSharedDataPtr) GetData() ffi.Ref[DrawListSharedData] {
	ptr := (*C.ImVector_ImDrawListSharedDataPtr)(unsafe.Pointer(p))
	return ffi.RefFromPtr[DrawListSharedData](unsafe.Pointer(ptr.Data))
}

// SetData sets the value in Data.
func (p Vector_ImDrawListSharedDataPtr) SetData(value ffi.Ref[DrawListSharedData]) {
	ptr := (*C.ImVector_ImDrawListSharedDataPtr)(unsafe.Pointer(p))
	ptr.Data = (**C.ImDrawListSharedData)(value.Raw())
}

// GetSize returns the value in Size.
func (p Vector_ImDrawListSharedDataPtr) GetSize() int32 {
	ptr := (*C.ImVector_ImDrawListSharedDataPtr)(unsafe.Pointer(p))
	return int32(ptr.Size)
}

// SetSize sets the value in Size.
func (p Vector_ImDrawListSharedDataPtr) SetSize(value int32) {
	ptr := (*C.ImVector_ImDrawListSharedDataPtr)(unsafe.Pointer(p))
	ptr.Size = (C.int32_t)(value)
}

// Vector_ImDrawVert wraps struct ImVector_ImDrawVert.
type Vector_ImDrawVert uintptr

// Vector_ImDrawVertNil is a null pointer.
var Vector_ImDrawVertNil Vector_ImDrawVert

// Vector_ImDrawVertSizeOf is the byte size of ImVector_ImDrawVert.
const Vector_ImDrawVertSizeOf = int(C.sizeof_ImVector_ImDrawVert)

// Vector_ImDrawVertAlloc allocates a continuous block of Vector_ImDrawVert.
func Vector_ImDrawVertAlloc(alloc ffi.Allocator, count int) Vector_ImDrawVert {
	ptr := alloc.Allocate(Vector_ImDrawVertSizeOf * count)
	return Vector_ImDrawVert(ptr)
}

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p Vector_ImDrawVert) Offset(offset int) Vector_ImDrawVert {
	return p + Vector_ImDrawVert(offset*Vector_ImDrawVertSizeOf)
}

// GetCapacity returns the value in Capacity.
func (p Vector_ImDrawVert) GetCapacity() int32 {
	ptr := (*C.ImVector_ImDrawVert)(unsafe.Pointer(p))
	return int32(ptr.Capacity)
}

// SetCapacity sets the value in Capacity.
func (p Vector_ImDrawVert) SetCapacity(value int32) {
	ptr := (*C.ImVector_ImDrawVert)(unsafe.Pointer(p))
	ptr.Capacity = (C.int32_t)(value)
}

// GetData returns the value in Data.
func (p Vector_ImDrawVert) GetData() DrawVert {
	ptr := (*C.ImVector_ImDrawVert)(unsafe.Pointer(p))
	return DrawVert(unsafe.Pointer(ptr.Data))
}

// SetData sets the value in Data.
func (p Vector_ImDrawVert) SetData(value DrawVert) {
	ptr := (*C.ImVector_ImDrawVert)(unsafe.Pointer(p))
	ptr.Data = (*C.ImDrawVert)(unsafe.Pointer(value))
}

// GetSize returns the value in Size.
func (p Vector_ImDrawVert) GetSize() int32 {
	ptr := (*C.ImVector_ImDrawVert)(unsafe.Pointer(p))
	return int32(ptr.Size)
}

// SetSize sets the value in Size.
func (p Vector_ImDrawVert) SetSize(value int32) {
	ptr := (*C.ImVector_ImDrawVert)(unsafe.Pointer(p))
	ptr.Size = (C.int32_t)(value)
}

// Vector_ImFontAtlasPtr wraps struct ImVector_ImFontAtlasPtr.
type Vector_ImFontAtlasPtr uintptr

// Vector_ImFontAtlasPtrNil is a null pointer.
var Vector_ImFontAtlasPtrNil Vector_ImFontAtlasPtr

// Vector_ImFontAtlasPtrSizeOf is the byte size of ImVector_ImFontAtlasPtr.
const Vector_ImFontAtlasPtrSizeOf = int(C.sizeof_ImVector_ImFontAtlasPtr)

// Vector_ImFontAtlasPtrAlloc allocates a continuous block of Vector_ImFontAtlasPtr.
func Vector_ImFontAtlasPtrAlloc(alloc ffi.Allocator, count int) Vector_ImFontAtlasPtr {
	ptr := alloc.Allocate(Vector_ImFontAtlasPtrSizeOf * count)
	return Vector_ImFontAtlasPtr(ptr)
}

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p Vector_ImFontAtlasPtr) Offset(offset int) Vector_ImFontAtlasPtr {
	return p + Vector_ImFontAtlasPtr(offset*Vector_ImFontAtlasPtrSizeOf)
}

// GetCapacity returns the value in Capacity.
func (p Vector_ImFontAtlasPtr) GetCapacity() int32 {
	ptr := (*C.ImVector_ImFontAtlasPtr)(unsafe.Pointer(p))
	return int32(ptr.Capacity)
}

// SetCapacity sets the value in Capacity.
func (p Vector_ImFontAtlasPtr) SetCapacity(value int32) {
	ptr := (*C.ImVector_ImFontAtlasPtr)(unsafe.Pointer(p))
	ptr.Capacity = (C.int32_t)(value)
}

// GetData returns the value in Data.
func (p Vector_ImFontAtlasPtr) GetData() ffi.Ref[FontAtlas] {
	ptr := (*C.ImVector_ImFontAtlasPtr)(unsafe.Pointer(p))
	return ffi.RefFromPtr[FontAtlas](unsafe.Pointer(ptr.Data))
}

// SetData sets the value in Data.
func (p Vector_ImFontAtlasPtr) SetData(value ffi.Ref[FontAtlas]) {
	ptr := (*C.ImVector_ImFontAtlasPtr)(unsafe.Pointer(p))
	ptr.Data = (**C.ImFontAtlas)(value.Raw())
}

// GetSize returns the value in Size.
func (p Vector_ImFontAtlasPtr) GetSize() int32 {
	ptr := (*C.ImVector_ImFontAtlasPtr)(unsafe.Pointer(p))
	return int32(ptr.Size)
}

// SetSize sets the value in Size.
func (p Vector_ImFontAtlasPtr) SetSize(value int32) {
	ptr := (*C.ImVector_ImFontAtlasPtr)(unsafe.Pointer(p))
	ptr.Size = (C.int32_t)(value)
}

// Vector_ImFontAtlasRectEntry wraps struct ImVector_ImFontAtlasRectEntry.
type Vector_ImFontAtlasRectEntry uintptr

// Vector_ImFontAtlasRectEntryNil is a null pointer.
var Vector_ImFontAtlasRectEntryNil Vector_ImFontAtlasRectEntry

// Vector_ImFontAtlasRectEntrySizeOf is the byte size of ImVector_ImFontAtlasRectEntry.
const Vector_ImFontAtlasRectEntrySizeOf = int(C.sizeof_ImVector_ImFontAtlasRectEntry)

// Vector_ImFontAtlasRectEntryAlloc allocates a continuous block of Vector_ImFontAtlasRectEntry.
func Vector_ImFontAtlasRectEntryAlloc(alloc ffi.Allocator, count int) Vector_ImFontAtlasRectEntry {
	ptr := alloc.Allocate(Vector_ImFontAtlasRectEntrySizeOf * count)
	return Vector_ImFontAtlasRectEntry(ptr)
}

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p Vector_ImFontAtlasRectEntry) Offset(offset int) Vector_ImFontAtlasRectEntry {
	return p + Vector_ImFontAtlasRectEntry(offset*Vector_ImFontAtlasRectEntrySizeOf)
}

// GetCapacity returns the value in Capacity.
func (p Vector_ImFontAtlasRectEntry) GetCapacity() int32 {
	ptr := (*C.ImVector_ImFontAtlasRectEntry)(unsafe.Pointer(p))
	return int32(ptr.Capacity)
}

// SetCapacity sets the value in Capacity.
func (p Vector_ImFontAtlasRectEntry) SetCapacity(value int32) {
	ptr := (*C.ImVector_ImFontAtlasRectEntry)(unsafe.Pointer(p))
	ptr.Capacity = (C.int32_t)(value)
}

// GetData returns the value in Data.
func (p Vector_ImFontAtlasRectEntry) GetData() FontAtlasRectEntry {
	ptr := (*C.ImVector_ImFontAtlasRectEntry)(unsafe.Pointer(p))
	return FontAtlasRectEntry(unsafe.Pointer(ptr.Data))
}

// SetData sets the value in Data.
func (p Vector_ImFontAtlasRectEntry) SetData(value FontAtlasRectEntry) {
	ptr := (*C.ImVector_ImFontAtlasRectEntry)(unsafe.Pointer(p))
	ptr.Data = (*C.ImFontAtlasRectEntry)(unsafe.Pointer(value))
}

// GetSize returns the value in Size.
func (p Vector_ImFontAtlasRectEntry) GetSize() int32 {
	ptr := (*C.ImVector_ImFontAtlasRectEntry)(unsafe.Pointer(p))
	return int32(ptr.Size)
}

// SetSize sets the value in Size.
func (p Vector_ImFontAtlasRectEntry) SetSize(value int32) {
	ptr := (*C.ImVector_ImFontAtlasRectEntry)(unsafe.Pointer(p))
	ptr.Size = (C.int32_t)(value)
}

// Vector_ImFontBakedPtr wraps struct ImVector_ImFontBakedPtr.
type Vector_ImFontBakedPtr uintptr

// Vector_ImFontBakedPtrNil is a null pointer.
var Vector_ImFontBakedPtrNil Vector_ImFontBakedPtr

// Vector_ImFontBakedPtrSizeOf is the byte size of ImVector_ImFontBakedPtr.
const Vector_ImFontBakedPtrSizeOf = int(C.sizeof_ImVector_ImFontBakedPtr)

// Vector_ImFontBakedPtrAlloc allocates a continuous block of Vector_ImFontBakedPtr.
func Vector_ImFontBakedPtrAlloc(alloc ffi.Allocator, count int) Vector_ImFontBakedPtr {
	ptr := alloc.Allocate(Vector_ImFontBakedPtrSizeOf * count)
	return Vector_ImFontBakedPtr(ptr)
}

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p Vector_ImFontBakedPtr) Offset(offset int) Vector_ImFontBakedPtr {
	return p + Vector_ImFontBakedPtr(offset*Vector_ImFontBakedPtrSizeOf)
}

// GetCapacity returns the value in Capacity.
func (p Vector_ImFontBakedPtr) GetCapacity() int32 {
	ptr := (*C.ImVector_ImFontBakedPtr)(unsafe.Pointer(p))
	return int32(ptr.Capacity)
}

// SetCapacity sets the value in Capacity.
func (p Vector_ImFontBakedPtr) SetCapacity(value int32) {
	ptr := (*C.ImVector_ImFontBakedPtr)(unsafe.Pointer(p))
	ptr.Capacity = (C.int32_t)(value)
}

// GetData returns the value in Data.
func (p Vector_ImFontBakedPtr) GetData() ffi.Ref[FontBaked] {
	ptr := (*C.ImVector_ImFontBakedPtr)(unsafe.Pointer(p))
	return ffi.RefFromPtr[FontBaked](unsafe.Pointer(ptr.Data))
}

// SetData sets the value in Data.
func (p Vector_ImFontBakedPtr) SetData(value ffi.Ref[FontBaked]) {
	ptr := (*C.ImVector_ImFontBakedPtr)(unsafe.Pointer(p))
	ptr.Data = (**C.ImFontBaked)(value.Raw())
}

// GetSize returns the value in Size.
func (p Vector_ImFontBakedPtr) GetSize() int32 {
	ptr := (*C.ImVector_ImFontBakedPtr)(unsafe.Pointer(p))
	return int32(ptr.Size)
}

// SetSize sets the value in Size.
func (p Vector_ImFontBakedPtr) SetSize(value int32) {
	ptr := (*C.ImVector_ImFontBakedPtr)(unsafe.Pointer(p))
	ptr.Size = (C.int32_t)(value)
}

// Vector_ImFontConfig wraps struct ImVector_ImFontConfig.
type Vector_ImFontConfig uintptr

// Vector_ImFontConfigNil is a null pointer.
var Vector_ImFontConfigNil Vector_ImFontConfig

// Vector_ImFontConfigSizeOf is the byte size of ImVector_ImFontConfig.
const Vector_ImFontConfigSizeOf = int(C.sizeof_ImVector_ImFontConfig)

// Vector_ImFontConfigAlloc allocates a continuous block of Vector_ImFontConfig.
func Vector_ImFontConfigAlloc(alloc ffi.Allocator, count int) Vector_ImFontConfig {
	ptr := alloc.Allocate(Vector_ImFontConfigSizeOf * count)
	return Vector_ImFontConfig(ptr)
}

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p Vector_ImFontConfig) Offset(offset int) Vector_ImFontConfig {
	return p + Vector_ImFontConfig(offset*Vector_ImFontConfigSizeOf)
}

// GetCapacity returns the value in Capacity.
func (p Vector_ImFontConfig) GetCapacity() int32 {
	ptr := (*C.ImVector_ImFontConfig)(unsafe.Pointer(p))
	return int32(ptr.Capacity)
}

// SetCapacity sets the value in Capacity.
func (p Vector_ImFontConfig) SetCapacity(value int32) {
	ptr := (*C.ImVector_ImFontConfig)(unsafe.Pointer(p))
	ptr.Capacity = (C.int32_t)(value)
}

// GetData returns the value in Data.
func (p Vector_ImFontConfig) GetData() FontConfig {
	ptr := (*C.ImVector_ImFontConfig)(unsafe.Pointer(p))
	return FontConfig(unsafe.Pointer(ptr.Data))
}

// SetData sets the value in Data.
func (p Vector_ImFontConfig) SetData(value FontConfig) {
	ptr := (*C.ImVector_ImFontConfig)(unsafe.Pointer(p))
	ptr.Data = (*C.ImFontConfig)(unsafe.Pointer(value))
}

// GetSize returns the value in Size.
func (p Vector_ImFontConfig) GetSize() int32 {
	ptr := (*C.ImVector_ImFontConfig)(unsafe.Pointer(p))
	return int32(ptr.Size)
}

// SetSize sets the value in Size.
func (p Vector_ImFontConfig) SetSize(value int32) {
	ptr := (*C.ImVector_ImFontConfig)(unsafe.Pointer(p))
	ptr.Size = (C.int32_t)(value)
}

// Vector_ImFontConfigPtr wraps struct ImVector_ImFontConfigPtr.
type Vector_ImFontConfigPtr uintptr

// Vector_ImFontConfigPtrNil is a null pointer.
var Vector_ImFontConfigPtrNil Vector_ImFontConfigPtr

// Vector_ImFontConfigPtrSizeOf is the byte size of ImVector_ImFontConfigPtr.
const Vector_ImFontConfigPtrSizeOf = int(C.sizeof_ImVector_ImFontConfigPtr)

// Vector_ImFontConfigPtrAlloc allocates a continuous block of Vector_ImFontConfigPtr.
func Vector_ImFontConfigPtrAlloc(alloc ffi.Allocator, count int) Vector_ImFontConfigPtr {
	ptr := alloc.Allocate(Vector_ImFontConfigPtrSizeOf * count)
	return Vector_ImFontConfigPtr(ptr)
}

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p Vector_ImFontConfigPtr) Offset(offset int) Vector_ImFontConfigPtr {
	return p + Vector_ImFontConfigPtr(offset*Vector_ImFontConfigPtrSizeOf)
}

// GetCapacity returns the value in Capacity.
func (p Vector_ImFontConfigPtr) GetCapacity() int32 {
	ptr := (*C.ImVector_ImFontConfigPtr)(unsafe.Pointer(p))
	return int32(ptr.Capacity)
}

// SetCapacity sets the value in Capacity.
func (p Vector_ImFontConfigPtr) SetCapacity(value int32) {
	ptr := (*C.ImVector_ImFontConfigPtr)(unsafe.Pointer(p))
	ptr.Capacity = (C.int32_t)(value)
}

// GetData returns the value in Data.
func (p Vector_ImFontConfigPtr) GetData() ffi.Ref[FontConfig] {
	ptr := (*C.ImVector_ImFontConfigPtr)(unsafe.Pointer(p))
	return ffi.RefFromPtr[FontConfig](unsafe.Pointer(ptr.Data))
}

// SetData sets the value in Data.
func (p Vector_ImFontConfigPtr) SetData(value ffi.Ref[FontConfig]) {
	ptr := (*C.ImVector_ImFontConfigPtr)(unsafe.Pointer(p))
	ptr.Data = (**C.ImFontConfig)(value.Raw())
}

// GetSize returns the value in Size.
func (p Vector_ImFontConfigPtr) GetSize() int32 {
	ptr := (*C.ImVector_ImFontConfigPtr)(unsafe.Pointer(p))
	return int32(ptr.Size)
}

// SetSize sets the value in Size.
func (p Vector_ImFontConfigPtr) SetSize(value int32) {
	ptr := (*C.ImVector_ImFontConfigPtr)(unsafe.Pointer(p))
	ptr.Size = (C.int32_t)(value)
}

// Vector_ImFontGlyph wraps struct ImVector_ImFontGlyph.
type Vector_ImFontGlyph uintptr

// Vector_ImFontGlyphNil is a null pointer.
var Vector_ImFontGlyphNil Vector_ImFontGlyph

// Vector_ImFontGlyphSizeOf is the byte size of ImVector_ImFontGlyph.
const Vector_ImFontGlyphSizeOf = int(C.sizeof_ImVector_ImFontGlyph)

// Vector_ImFontGlyphAlloc allocates a continuous block of Vector_ImFontGlyph.
func Vector_ImFontGlyphAlloc(alloc ffi.Allocator, count int) Vector_ImFontGlyph {
	ptr := alloc.Allocate(Vector_ImFontGlyphSizeOf * count)
	return Vector_ImFontGlyph(ptr)
}

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p Vector_ImFontGlyph) Offset(offset int) Vector_ImFontGlyph {
	return p + Vector_ImFontGlyph(offset*Vector_ImFontGlyphSizeOf)
}

// GetCapacity returns the value in Capacity.
func (p Vector_ImFontGlyph) GetCapacity() int32 {
	ptr := (*C.ImVector_ImFontGlyph)(unsafe.Pointer(p))
	return int32(ptr.Capacity)
}

// SetCapacity sets the value in Capacity.
func (p Vector_ImFontGlyph) SetCapacity(value int32) {
	ptr := (*C.ImVector_ImFontGlyph)(unsafe.Pointer(p))
	ptr.Capacity = (C.int32_t)(value)
}

// GetData returns the value in Data.
func (p Vector_ImFontGlyph) GetData() FontGlyph {
	ptr := (*C.ImVector_ImFontGlyph)(unsafe.Pointer(p))
	return FontGlyph(unsafe.Pointer(ptr.Data))
}

// SetData sets the value in Data.
func (p Vector_ImFontGlyph) SetData(value FontGlyph) {
	ptr := (*C.ImVector_ImFontGlyph)(unsafe.Pointer(p))
	ptr.Data = (*C.ImFontGlyph)(unsafe.Pointer(value))
}

// GetSize returns the value in Size.
func (p Vector_ImFontGlyph) GetSize() int32 {
	ptr := (*C.ImVector_ImFontGlyph)(unsafe.Pointer(p))
	return int32(ptr.Size)
}

// SetSize sets the value in Size.
func (p Vector_ImFontGlyph) SetSize(value int32) {
	ptr := (*C.ImVector_ImFontGlyph)(unsafe.Pointer(p))
	ptr.Size = (C.int32_t)(value)
}

// Vector_ImFontPtr wraps struct ImVector_ImFontPtr.
type Vector_ImFontPtr uintptr

// Vector_ImFontPtrNil is a null pointer.
var Vector_ImFontPtrNil Vector_ImFontPtr

// Vector_ImFontPtrSizeOf is the byte size of ImVector_ImFontPtr.
const Vector_ImFontPtrSizeOf = int(C.sizeof_ImVector_ImFontPtr)

// Vector_ImFontPtrAlloc allocates a continuous block of Vector_ImFontPtr.
func Vector_ImFontPtrAlloc(alloc ffi.Allocator, count int) Vector_ImFontPtr {
	ptr := alloc.Allocate(Vector_ImFontPtrSizeOf * count)
	return Vector_ImFontPtr(ptr)
}

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p Vector_ImFontPtr) Offset(offset int) Vector_ImFontPtr {
	return p + Vector_ImFontPtr(offset*Vector_ImFontPtrSizeOf)
}

// GetCapacity returns the value in Capacity.
func (p Vector_ImFontPtr) GetCapacity() int32 {
	ptr := (*C.ImVector_ImFontPtr)(unsafe.Pointer(p))
	return int32(ptr.Capacity)
}

// SetCapacity sets the value in Capacity.
func (p Vector_ImFontPtr) SetCapacity(value int32) {
	ptr := (*C.ImVector_ImFontPtr)(unsafe.Pointer(p))
	ptr.Capacity = (C.int32_t)(value)
}

// GetData returns the value in Data.
func (p Vector_ImFontPtr) GetData() ffi.Ref[Font] {
	ptr := (*C.ImVector_ImFontPtr)(unsafe.Pointer(p))
	return ffi.RefFromPtr[Font](unsafe.Pointer(ptr.Data))
}

// SetData sets the value in Data.
func (p Vector_ImFontPtr) SetData(value ffi.Ref[Font]) {
	ptr := (*C.ImVector_ImFontPtr)(unsafe.Pointer(p))
	ptr.Data = (**C.ImFont)(value.Raw())
}

// GetSize returns the value in Size.
func (p Vector_ImFontPtr) GetSize() int32 {
	ptr := (*C.ImVector_ImFontPtr)(unsafe.Pointer(p))
	return int32(ptr.Size)
}

// SetSize sets the value in Size.
func (p Vector_ImFontPtr) SetSize(value int32) {
	ptr := (*C.ImVector_ImFontPtr)(unsafe.Pointer(p))
	ptr.Size = (C.int32_t)(value)
}

// Vector_ImFontStackData wraps struct ImVector_ImFontStackData.
type Vector_ImFontStackData uintptr

// Vector_ImFontStackDataNil is a null pointer.
var Vector_ImFontStackDataNil Vector_ImFontStackData

// Vector_ImFontStackDataSizeOf is the byte size of ImVector_ImFontStackData.
const Vector_ImFontStackDataSizeOf = int(C.sizeof_ImVector_ImFontStackData)

// Vector_ImFontStackDataAlloc allocates a continuous block of Vector_ImFontStackData.
func Vector_ImFontStackDataAlloc(alloc ffi.Allocator, count int) Vector_ImFontStackData {
	ptr := alloc.Allocate(Vector_ImFontStackDataSizeOf * count)
	return Vector_ImFontStackData(ptr)
}

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p Vector_ImFontStackData) Offset(offset int) Vector_ImFontStackData {
	return p + Vector_ImFontStackData(offset*Vector_ImFontStackDataSizeOf)
}

// GetCapacity returns the value in Capacity.
func (p Vector_ImFontStackData) GetCapacity() int32 {
	ptr := (*C.ImVector_ImFontStackData)(unsafe.Pointer(p))
	return int32(ptr.Capacity)
}

// SetCapacity sets the value in Capacity.
func (p Vector_ImFontStackData) SetCapacity(value int32) {
	ptr := (*C.ImVector_ImFontStackData)(unsafe.Pointer(p))
	ptr.Capacity = (C.int32_t)(value)
}

// GetData returns the value in Data.
func (p Vector_ImFontStackData) GetData() FontStackData {
	ptr := (*C.ImVector_ImFontStackData)(unsafe.Pointer(p))
	return FontStackData(unsafe.Pointer(ptr.Data))
}

// SetData sets the value in Data.
func (p Vector_ImFontStackData) SetData(value FontStackData) {
	ptr := (*C.ImVector_ImFontStackData)(unsafe.Pointer(p))
	ptr.Data = (*C.ImFontStackData)(unsafe.Pointer(value))
}

// GetSize returns the value in Size.
func (p Vector_ImFontStackData) GetSize() int32 {
	ptr := (*C.ImVector_ImFontStackData)(unsafe.Pointer(p))
	return int32(ptr.Size)
}

// SetSize sets the value in Size.
func (p Vector_ImFontStackData) SetSize(value int32) {
	ptr := (*C.ImVector_ImFontStackData)(unsafe.Pointer(p))
	ptr.Size = (C.int32_t)(value)
}

// Vector_ImGuiColorMod wraps struct ImVector_ImGuiColorMod.
type Vector_ImGuiColorMod uintptr

// Vector_ImGuiColorModNil is a null pointer.
var Vector_ImGuiColorModNil Vector_ImGuiColorMod

// Vector_ImGuiColorModSizeOf is the byte size of ImVector_ImGuiColorMod.
const Vector_ImGuiColorModSizeOf = int(C.sizeof_ImVector_ImGuiColorMod)

// Vector_ImGuiColorModAlloc allocates a continuous block of Vector_ImGuiColorMod.
func Vector_ImGuiColorModAlloc(alloc ffi.Allocator, count int) Vector_ImGuiColorMod {
	ptr := alloc.Allocate(Vector_ImGuiColorModSizeOf * count)
	return Vector_ImGuiColorMod(ptr)
}

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p Vector_ImGuiColorMod) Offset(offset int) Vector_ImGuiColorMod {
	return p + Vector_ImGuiColorMod(offset*Vector_ImGuiColorModSizeOf)
}

// GetCapacity returns the value in Capacity.
func (p Vector_ImGuiColorMod) GetCapacity() int32 {
	ptr := (*C.ImVector_ImGuiColorMod)(unsafe.Pointer(p))
	return int32(ptr.Capacity)
}

// SetCapacity sets the value in Capacity.
func (p Vector_ImGuiColorMod) SetCapacity(value int32) {
	ptr := (*C.ImVector_ImGuiColorMod)(unsafe.Pointer(p))
	ptr.Capacity = (C.int32_t)(value)
}

// GetData returns the value in Data.
func (p Vector_ImGuiColorMod) GetData() GuiColorMod {
	ptr := (*C.ImVector_ImGuiColorMod)(unsafe.Pointer(p))
	return GuiColorMod(unsafe.Pointer(ptr.Data))
}

// SetData sets the value in Data.
func (p Vector_ImGuiColorMod) SetData(value GuiColorMod) {
	ptr := (*C.ImVector_ImGuiColorMod)(unsafe.Pointer(p))
	ptr.Data = (*C.ImGuiColorMod)(unsafe.Pointer(value))
}

// GetSize returns the value in Size.
func (p Vector_ImGuiColorMod) GetSize() int32 {
	ptr := (*C.ImVector_ImGuiColorMod)(unsafe.Pointer(p))
	return int32(ptr.Size)
}

// SetSize sets the value in Size.
func (p Vector_ImGuiColorMod) SetSize(value int32) {
	ptr := (*C.ImVector_ImGuiColorMod)(unsafe.Pointer(p))
	ptr.Size = (C.int32_t)(value)
}

// Vector_ImGuiContextHook wraps struct ImVector_ImGuiContextHook.
type Vector_ImGuiContextHook uintptr

// Vector_ImGuiContextHookNil is a null pointer.
var Vector_ImGuiContextHookNil Vector_ImGuiContextHook

// Vector_ImGuiContextHookSizeOf is the byte size of ImVector_ImGuiContextHook.
const Vector_ImGuiContextHookSizeOf = int(C.sizeof_ImVector_ImGuiContextHook)

// Vector_ImGuiContextHookAlloc allocates a continuous block of Vector_ImGuiContextHook.
func Vector_ImGuiContextHookAlloc(alloc ffi.Allocator, count int) Vector_ImGuiContextHook {
	ptr := alloc.Allocate(Vector_ImGuiContextHookSizeOf * count)
	return Vector_ImGuiContextHook(ptr)
}

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p Vector_ImGuiContextHook) Offset(offset int) Vector_ImGuiContextHook {
	return p + Vector_ImGuiContextHook(offset*Vector_ImGuiContextHookSizeOf)
}

// GetCapacity returns the value in Capacity.
func (p Vector_ImGuiContextHook) GetCapacity() int32 {
	ptr := (*C.ImVector_ImGuiContextHook)(unsafe.Pointer(p))
	return int32(ptr.Capacity)
}

// SetCapacity sets the value in Capacity.
func (p Vector_ImGuiContextHook) SetCapacity(value int32) {
	ptr := (*C.ImVector_ImGuiContextHook)(unsafe.Pointer(p))
	ptr.Capacity = (C.int32_t)(value)
}

// GetData returns the value in Data.
func (p Vector_ImGuiContextHook) GetData() GuiContextHook {
	ptr := (*C.ImVector_ImGuiContextHook)(unsafe.Pointer(p))
	return GuiContextHook(unsafe.Pointer(ptr.Data))
}

// SetData sets the value in Data.
func (p Vector_ImGuiContextHook) SetData(value GuiContextHook) {
	ptr := (*C.ImVector_ImGuiContextHook)(unsafe.Pointer(p))
	ptr.Data = (*C.ImGuiContextHook)(unsafe.Pointer(value))
}

// GetSize returns the value in Size.
func (p Vector_ImGuiContextHook) GetSize() int32 {
	ptr := (*C.ImVector_ImGuiContextHook)(unsafe.Pointer(p))
	return int32(ptr.Size)
}

// SetSize sets the value in Size.
func (p Vector_ImGuiContextHook) SetSize(value int32) {
	ptr := (*C.ImVector_ImGuiContextHook)(unsafe.Pointer(p))
	ptr.Size = (C.int32_t)(value)
}

// Vector_ImGuiDockNodeSettings wraps struct ImVector_ImGuiDockNodeSettings.
type Vector_ImGuiDockNodeSettings uintptr

// Vector_ImGuiDockNodeSettingsNil is a null pointer.
var Vector_ImGuiDockNodeSettingsNil Vector_ImGuiDockNodeSettings

// Vector_ImGuiDockNodeSettingsSizeOf is the byte size of ImVector_ImGuiDockNodeSettings.
const Vector_ImGuiDockNodeSettingsSizeOf = int(C.sizeof_ImVector_ImGuiDockNodeSettings)

// Vector_ImGuiDockNodeSettingsAlloc allocates a continuous block of Vector_ImGuiDockNodeSettings.
func Vector_ImGuiDockNodeSettingsAlloc(alloc ffi.Allocator, count int) Vector_ImGuiDockNodeSettings {
	ptr := alloc.Allocate(Vector_ImGuiDockNodeSettingsSizeOf * count)
	return Vector_ImGuiDockNodeSettings(ptr)
}

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p Vector_ImGuiDockNodeSettings) Offset(offset int) Vector_ImGuiDockNodeSettings {
	return p + Vector_ImGuiDockNodeSettings(offset*Vector_ImGuiDockNodeSettingsSizeOf)
}

// GetCapacity returns the value in Capacity.
func (p Vector_ImGuiDockNodeSettings) GetCapacity() int32 {
	ptr := (*C.ImVector_ImGuiDockNodeSettings)(unsafe.Pointer(p))
	return int32(ptr.Capacity)
}

// SetCapacity sets the value in Capacity.
func (p Vector_ImGuiDockNodeSettings) SetCapacity(value int32) {
	ptr := (*C.ImVector_ImGuiDockNodeSettings)(unsafe.Pointer(p))
	ptr.Capacity = (C.int32_t)(value)
}

// Vector_ImGuiDockNodeSettings.Data is unsupported: category unsupported.

// GetSize returns the value in Size.
func (p Vector_ImGuiDockNodeSettings) GetSize() int32 {
	ptr := (*C.ImVector_ImGuiDockNodeSettings)(unsafe.Pointer(p))
	return int32(ptr.Size)
}

// SetSize sets the value in Size.
func (p Vector_ImGuiDockNodeSettings) SetSize(value int32) {
	ptr := (*C.ImVector_ImGuiDockNodeSettings)(unsafe.Pointer(p))
	ptr.Size = (C.int32_t)(value)
}

// Vector_ImGuiDockRequest wraps struct ImVector_ImGuiDockRequest.
type Vector_ImGuiDockRequest uintptr

// Vector_ImGuiDockRequestNil is a null pointer.
var Vector_ImGuiDockRequestNil Vector_ImGuiDockRequest

// Vector_ImGuiDockRequestSizeOf is the byte size of ImVector_ImGuiDockRequest.
const Vector_ImGuiDockRequestSizeOf = int(C.sizeof_ImVector_ImGuiDockRequest)

// Vector_ImGuiDockRequestAlloc allocates a continuous block of Vector_ImGuiDockRequest.
func Vector_ImGuiDockRequestAlloc(alloc ffi.Allocator, count int) Vector_ImGuiDockRequest {
	ptr := alloc.Allocate(Vector_ImGuiDockRequestSizeOf * count)
	return Vector_ImGuiDockRequest(ptr)
}

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p Vector_ImGuiDockRequest) Offset(offset int) Vector_ImGuiDockRequest {
	return p + Vector_ImGuiDockRequest(offset*Vector_ImGuiDockRequestSizeOf)
}

// GetCapacity returns the value in Capacity.
func (p Vector_ImGuiDockRequest) GetCapacity() int32 {
	ptr := (*C.ImVector_ImGuiDockRequest)(unsafe.Pointer(p))
	return int32(ptr.Capacity)
}

// SetCapacity sets the value in Capacity.
func (p Vector_ImGuiDockRequest) SetCapacity(value int32) {
	ptr := (*C.ImVector_ImGuiDockRequest)(unsafe.Pointer(p))
	ptr.Capacity = (C.int32_t)(value)
}

// Vector_ImGuiDockRequest.Data is unsupported: category unsupported.

// GetSize returns the value in Size.
func (p Vector_ImGuiDockRequest) GetSize() int32 {
	ptr := (*C.ImVector_ImGuiDockRequest)(unsafe.Pointer(p))
	return int32(ptr.Size)
}

// SetSize sets the value in Size.
func (p Vector_ImGuiDockRequest) SetSize(value int32) {
	ptr := (*C.ImVector_ImGuiDockRequest)(unsafe.Pointer(p))
	ptr.Size = (C.int32_t)(value)
}

// Vector_ImGuiFocusScopeData wraps struct ImVector_ImGuiFocusScopeData.
type Vector_ImGuiFocusScopeData uintptr

// Vector_ImGuiFocusScopeDataNil is a null pointer.
var Vector_ImGuiFocusScopeDataNil Vector_ImGuiFocusScopeData

// Vector_ImGuiFocusScopeDataSizeOf is the byte size of ImVector_ImGuiFocusScopeData.
const Vector_ImGuiFocusScopeDataSizeOf = int(C.sizeof_ImVector_ImGuiFocusScopeData)

// Vector_ImGuiFocusScopeDataAlloc allocates a continuous block of Vector_ImGuiFocusScopeData.
func Vector_ImGuiFocusScopeDataAlloc(alloc ffi.Allocator, count int) Vector_ImGuiFocusScopeData {
	ptr := alloc.Allocate(Vector_ImGuiFocusScopeDataSizeOf * count)
	return Vector_ImGuiFocusScopeData(ptr)
}

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p Vector_ImGuiFocusScopeData) Offset(offset int) Vector_ImGuiFocusScopeData {
	return p + Vector_ImGuiFocusScopeData(offset*Vector_ImGuiFocusScopeDataSizeOf)
}

// GetCapacity returns the value in Capacity.
func (p Vector_ImGuiFocusScopeData) GetCapacity() int32 {
	ptr := (*C.ImVector_ImGuiFocusScopeData)(unsafe.Pointer(p))
	return int32(ptr.Capacity)
}

// SetCapacity sets the value in Capacity.
func (p Vector_ImGuiFocusScopeData) SetCapacity(value int32) {
	ptr := (*C.ImVector_ImGuiFocusScopeData)(unsafe.Pointer(p))
	ptr.Capacity = (C.int32_t)(value)
}

// GetData returns the value in Data.
func (p Vector_ImGuiFocusScopeData) GetData() GuiFocusScopeData {
	ptr := (*C.ImVector_ImGuiFocusScopeData)(unsafe.Pointer(p))
	return GuiFocusScopeData(unsafe.Pointer(ptr.Data))
}

// SetData sets the value in Data.
func (p Vector_ImGuiFocusScopeData) SetData(value GuiFocusScopeData) {
	ptr := (*C.ImVector_ImGuiFocusScopeData)(unsafe.Pointer(p))
	ptr.Data = (*C.ImGuiFocusScopeData)(unsafe.Pointer(value))
}

// GetSize returns the value in Size.
func (p Vector_ImGuiFocusScopeData) GetSize() int32 {
	ptr := (*C.ImVector_ImGuiFocusScopeData)(unsafe.Pointer(p))
	return int32(ptr.Size)
}

// SetSize sets the value in Size.
func (p Vector_ImGuiFocusScopeData) SetSize(value int32) {
	ptr := (*C.ImVector_ImGuiFocusScopeData)(unsafe.Pointer(p))
	ptr.Size = (C.int32_t)(value)
}

// Vector_ImGuiGroupData wraps struct ImVector_ImGuiGroupData.
type Vector_ImGuiGroupData uintptr

// Vector_ImGuiGroupDataNil is a null pointer.
var Vector_ImGuiGroupDataNil Vector_ImGuiGroupData

// Vector_ImGuiGroupDataSizeOf is the byte size of ImVector_ImGuiGroupData.
const Vector_ImGuiGroupDataSizeOf = int(C.sizeof_ImVector_ImGuiGroupData)

// Vector_ImGuiGroupDataAlloc allocates a continuous block of Vector_ImGuiGroupData.
func Vector_ImGuiGroupDataAlloc(alloc ffi.Allocator, count int) Vector_ImGuiGroupData {
	ptr := alloc.Allocate(Vector_ImGuiGroupDataSizeOf * count)
	return Vector_ImGuiGroupData(ptr)
}

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p Vector_ImGuiGroupData) Offset(offset int) Vector_ImGuiGroupData {
	return p + Vector_ImGuiGroupData(offset*Vector_ImGuiGroupDataSizeOf)
}

// GetCapacity returns the value in Capacity.
func (p Vector_ImGuiGroupData) GetCapacity() int32 {
	ptr := (*C.ImVector_ImGuiGroupData)(unsafe.Pointer(p))
	return int32(ptr.Capacity)
}

// SetCapacity sets the value in Capacity.
func (p Vector_ImGuiGroupData) SetCapacity(value int32) {
	ptr := (*C.ImVector_ImGuiGroupData)(unsafe.Pointer(p))
	ptr.Capacity = (C.int32_t)(value)
}

// GetData returns the value in Data.
func (p Vector_ImGuiGroupData) GetData() GuiGroupData {
	ptr := (*C.ImVector_ImGuiGroupData)(unsafe.Pointer(p))
	return GuiGroupData(unsafe.Pointer(ptr.Data))
}

// SetData sets the value in Data.
func (p Vector_ImGuiGroupData) SetData(value GuiGroupData) {
	ptr := (*C.ImVector_ImGuiGroupData)(unsafe.Pointer(p))
	ptr.Data = (*C.ImGuiGroupData)(unsafe.Pointer(value))
}

// GetSize returns the value in Size.
func (p Vector_ImGuiGroupData) GetSize() int32 {
	ptr := (*C.ImVector_ImGuiGroupData)(unsafe.Pointer(p))
	return int32(ptr.Size)
}

// SetSize sets the value in Size.
func (p Vector_ImGuiGroupData) SetSize(value int32) {
	ptr := (*C.ImVector_ImGuiGroupData)(unsafe.Pointer(p))
	ptr.Size = (C.int32_t)(value)
}

// Vector_ImGuiID wraps struct ImVector_ImGuiID.
type Vector_ImGuiID uintptr

// Vector_ImGuiIDNil is a null pointer.
var Vector_ImGuiIDNil Vector_ImGuiID

// Vector_ImGuiIDSizeOf is the byte size of ImVector_ImGuiID.
const Vector_ImGuiIDSizeOf = int(C.sizeof_ImVector_ImGuiID)

// Vector_ImGuiIDAlloc allocates a continuous block of Vector_ImGuiID.
func Vector_ImGuiIDAlloc(alloc ffi.Allocator, count int) Vector_ImGuiID {
	ptr := alloc.Allocate(Vector_ImGuiIDSizeOf * count)
	return Vector_ImGuiID(ptr)
}

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p Vector_ImGuiID) Offset(offset int) Vector_ImGuiID {
	return p + Vector_ImGuiID(offset*Vector_ImGuiIDSizeOf)
}

// GetCapacity returns the value in Capacity.
func (p Vector_ImGuiID) GetCapacity() int32 {
	ptr := (*C.ImVector_ImGuiID)(unsafe.Pointer(p))
	return int32(ptr.Capacity)
}

// SetCapacity sets the value in Capacity.
func (p Vector_ImGuiID) SetCapacity(value int32) {
	ptr := (*C.ImVector_ImGuiID)(unsafe.Pointer(p))
	ptr.Capacity = (C.int32_t)(value)
}

// Vector_ImGuiID.Data is unsupported: category unsupported.

// GetSize returns the value in Size.
func (p Vector_ImGuiID) GetSize() int32 {
	ptr := (*C.ImVector_ImGuiID)(unsafe.Pointer(p))
	return int32(ptr.Size)
}

// SetSize sets the value in Size.
func (p Vector_ImGuiID) SetSize(value int32) {
	ptr := (*C.ImVector_ImGuiID)(unsafe.Pointer(p))
	ptr.Size = (C.int32_t)(value)
}

// Vector_ImGuiInputEvent wraps struct ImVector_ImGuiInputEvent.
type Vector_ImGuiInputEvent uintptr

// Vector_ImGuiInputEventNil is a null pointer.
var Vector_ImGuiInputEventNil Vector_ImGuiInputEvent

// Vector_ImGuiInputEventSizeOf is the byte size of ImVector_ImGuiInputEvent.
const Vector_ImGuiInputEventSizeOf = int(C.sizeof_ImVector_ImGuiInputEvent)

// Vector_ImGuiInputEventAlloc allocates a continuous block of Vector_ImGuiInputEvent.
func Vector_ImGuiInputEventAlloc(alloc ffi.Allocator, count int) Vector_ImGuiInputEvent {
	ptr := alloc.Allocate(Vector_ImGuiInputEventSizeOf * count)
	return Vector_ImGuiInputEvent(ptr)
}

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p Vector_ImGuiInputEvent) Offset(offset int) Vector_ImGuiInputEvent {
	return p + Vector_ImGuiInputEvent(offset*Vector_ImGuiInputEventSizeOf)
}

// GetCapacity returns the value in Capacity.
func (p Vector_ImGuiInputEvent) GetCapacity() int32 {
	ptr := (*C.ImVector_ImGuiInputEvent)(unsafe.Pointer(p))
	return int32(ptr.Capacity)
}

// SetCapacity sets the value in Capacity.
func (p Vector_ImGuiInputEvent) SetCapacity(value int32) {
	ptr := (*C.ImVector_ImGuiInputEvent)(unsafe.Pointer(p))
	ptr.Capacity = (C.int32_t)(value)
}

// GetData returns the value in Data.
func (p Vector_ImGuiInputEvent) GetData() GuiInputEvent {
	ptr := (*C.ImVector_ImGuiInputEvent)(unsafe.Pointer(p))
	return GuiInputEvent(unsafe.Pointer(ptr.Data))
}

// SetData sets the value in Data.
func (p Vector_ImGuiInputEvent) SetData(value GuiInputEvent) {
	ptr := (*C.ImVector_ImGuiInputEvent)(unsafe.Pointer(p))
	ptr.Data = (*C.ImGuiInputEvent)(unsafe.Pointer(value))
}

// GetSize returns the value in Size.
func (p Vector_ImGuiInputEvent) GetSize() int32 {
	ptr := (*C.ImVector_ImGuiInputEvent)(unsafe.Pointer(p))
	return int32(ptr.Size)
}

// SetSize sets the value in Size.
func (p Vector_ImGuiInputEvent) SetSize(value int32) {
	ptr := (*C.ImVector_ImGuiInputEvent)(unsafe.Pointer(p))
	ptr.Size = (C.int32_t)(value)
}

// Vector_ImGuiItemFlags wraps struct ImVector_ImGuiItemFlags.
type Vector_ImGuiItemFlags uintptr

// Vector_ImGuiItemFlagsNil is a null pointer.
var Vector_ImGuiItemFlagsNil Vector_ImGuiItemFlags

// Vector_ImGuiItemFlagsSizeOf is the byte size of ImVector_ImGuiItemFlags.
const Vector_ImGuiItemFlagsSizeOf = int(C.sizeof_ImVector_ImGuiItemFlags)

// Vector_ImGuiItemFlagsAlloc allocates a continuous block of Vector_ImGuiItemFlags.
func Vector_ImGuiItemFlagsAlloc(alloc ffi.Allocator, count int) Vector_ImGuiItemFlags {
	ptr := alloc.Allocate(Vector_ImGuiItemFlagsSizeOf * count)
	return Vector_ImGuiItemFlags(ptr)
}

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p Vector_ImGuiItemFlags) Offset(offset int) Vector_ImGuiItemFlags {
	return p + Vector_ImGuiItemFlags(offset*Vector_ImGuiItemFlagsSizeOf)
}

// GetCapacity returns the value in Capacity.
func (p Vector_ImGuiItemFlags) GetCapacity() int32 {
	ptr := (*C.ImVector_ImGuiItemFlags)(unsafe.Pointer(p))
	return int32(ptr.Capacity)
}

// SetCapacity sets the value in Capacity.
func (p Vector_ImGuiItemFlags) SetCapacity(value int32) {
	ptr := (*C.ImVector_ImGuiItemFlags)(unsafe.Pointer(p))
	ptr.Capacity = (C.int32_t)(value)
}

// GetData returns the value in Data.
func (p Vector_ImGuiItemFlags) GetData() ffi.Ref[GuiItemFlags] {
	ptr := (*C.ImVector_ImGuiItemFlags)(unsafe.Pointer(p))
	return ffi.RefFromPtr[GuiItemFlags](unsafe.Pointer(ptr.Data))
}

// SetData sets the value in Data.
func (p Vector_ImGuiItemFlags) SetData(value ffi.Ref[GuiItemFlags]) {
	ptr := (*C.ImVector_ImGuiItemFlags)(unsafe.Pointer(p))
	ptr.Data = (*C.ImGuiItemFlags)(value.Raw())
}

// GetSize returns the value in Size.
func (p Vector_ImGuiItemFlags) GetSize() int32 {
	ptr := (*C.ImVector_ImGuiItemFlags)(unsafe.Pointer(p))
	return int32(ptr.Size)
}

// SetSize sets the value in Size.
func (p Vector_ImGuiItemFlags) SetSize(value int32) {
	ptr := (*C.ImVector_ImGuiItemFlags)(unsafe.Pointer(p))
	ptr.Size = (C.int32_t)(value)
}

// Vector_ImGuiKeyRoutingData wraps struct ImVector_ImGuiKeyRoutingData.
type Vector_ImGuiKeyRoutingData uintptr

// Vector_ImGuiKeyRoutingDataNil is a null pointer.
var Vector_ImGuiKeyRoutingDataNil Vector_ImGuiKeyRoutingData

// Vector_ImGuiKeyRoutingDataSizeOf is the byte size of ImVector_ImGuiKeyRoutingData.
const Vector_ImGuiKeyRoutingDataSizeOf = int(C.sizeof_ImVector_ImGuiKeyRoutingData)

// Vector_ImGuiKeyRoutingDataAlloc allocates a continuous block of Vector_ImGuiKeyRoutingData.
func Vector_ImGuiKeyRoutingDataAlloc(alloc ffi.Allocator, count int) Vector_ImGuiKeyRoutingData {
	ptr := alloc.Allocate(Vector_ImGuiKeyRoutingDataSizeOf * count)
	return Vector_ImGuiKeyRoutingData(ptr)
}

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p Vector_ImGuiKeyRoutingData) Offset(offset int) Vector_ImGuiKeyRoutingData {
	return p + Vector_ImGuiKeyRoutingData(offset*Vector_ImGuiKeyRoutingDataSizeOf)
}

// GetCapacity returns the value in Capacity.
func (p Vector_ImGuiKeyRoutingData) GetCapacity() int32 {
	ptr := (*C.ImVector_ImGuiKeyRoutingData)(unsafe.Pointer(p))
	return int32(ptr.Capacity)
}

// SetCapacity sets the value in Capacity.
func (p Vector_ImGuiKeyRoutingData) SetCapacity(value int32) {
	ptr := (*C.ImVector_ImGuiKeyRoutingData)(unsafe.Pointer(p))
	ptr.Capacity = (C.int32_t)(value)
}

// GetData returns the value in Data.
func (p Vector_ImGuiKeyRoutingData) GetData() GuiKeyRoutingData {
	ptr := (*C.ImVector_ImGuiKeyRoutingData)(unsafe.Pointer(p))
	return GuiKeyRoutingData(unsafe.Pointer(ptr.Data))
}

// SetData sets the value in Data.
func (p Vector_ImGuiKeyRoutingData) SetData(value GuiKeyRoutingData) {
	ptr := (*C.ImVector_ImGuiKeyRoutingData)(unsafe.Pointer(p))
	ptr.Data = (*C.ImGuiKeyRoutingData)(unsafe.Pointer(value))
}

// GetSize returns the value in Size.
func (p Vector_ImGuiKeyRoutingData) GetSize() int32 {
	ptr := (*C.ImVector_ImGuiKeyRoutingData)(unsafe.Pointer(p))
	return int32(ptr.Size)
}

// SetSize sets the value in Size.
func (p Vector_ImGuiKeyRoutingData) SetSize(value int32) {
	ptr := (*C.ImVector_ImGuiKeyRoutingData)(unsafe.Pointer(p))
	ptr.Size = (C.int32_t)(value)
}

// Vector_ImGuiListClipperData wraps struct ImVector_ImGuiListClipperData.
type Vector_ImGuiListClipperData uintptr

// Vector_ImGuiListClipperDataNil is a null pointer.
var Vector_ImGuiListClipperDataNil Vector_ImGuiListClipperData

// Vector_ImGuiListClipperDataSizeOf is the byte size of ImVector_ImGuiListClipperData.
const Vector_ImGuiListClipperDataSizeOf = int(C.sizeof_ImVector_ImGuiListClipperData)

// Vector_ImGuiListClipperDataAlloc allocates a continuous block of Vector_ImGuiListClipperData.
func Vector_ImGuiListClipperDataAlloc(alloc ffi.Allocator, count int) Vector_ImGuiListClipperData {
	ptr := alloc.Allocate(Vector_ImGuiListClipperDataSizeOf * count)
	return Vector_ImGuiListClipperData(ptr)
}

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p Vector_ImGuiListClipperData) Offset(offset int) Vector_ImGuiListClipperData {
	return p + Vector_ImGuiListClipperData(offset*Vector_ImGuiListClipperDataSizeOf)
}

// GetCapacity returns the value in Capacity.
func (p Vector_ImGuiListClipperData) GetCapacity() int32 {
	ptr := (*C.ImVector_ImGuiListClipperData)(unsafe.Pointer(p))
	return int32(ptr.Capacity)
}

// SetCapacity sets the value in Capacity.
func (p Vector_ImGuiListClipperData) SetCapacity(value int32) {
	ptr := (*C.ImVector_ImGuiListClipperData)(unsafe.Pointer(p))
	ptr.Capacity = (C.int32_t)(value)
}

// GetData returns the value in Data.
func (p Vector_ImGuiListClipperData) GetData() GuiListClipperData {
	ptr := (*C.ImVector_ImGuiListClipperData)(unsafe.Pointer(p))
	return GuiListClipperData(unsafe.Pointer(ptr.Data))
}

// SetData sets the value in Data.
func (p Vector_ImGuiListClipperData) SetData(value GuiListClipperData) {
	ptr := (*C.ImVector_ImGuiListClipperData)(unsafe.Pointer(p))
	ptr.Data = (*C.ImGuiListClipperData)(unsafe.Pointer(value))
}

// GetSize returns the value in Size.
func (p Vector_ImGuiListClipperData) GetSize() int32 {
	ptr := (*C.ImVector_ImGuiListClipperData)(unsafe.Pointer(p))
	return int32(ptr.Size)
}

// SetSize sets the value in Size.
func (p Vector_ImGuiListClipperData) SetSize(value int32) {
	ptr := (*C.ImVector_ImGuiListClipperData)(unsafe.Pointer(p))
	ptr.Size = (C.int32_t)(value)
}

// Vector_ImGuiListClipperRange wraps struct ImVector_ImGuiListClipperRange.
type Vector_ImGuiListClipperRange uintptr

// Vector_ImGuiListClipperRangeNil is a null pointer.
var Vector_ImGuiListClipperRangeNil Vector_ImGuiListClipperRange

// Vector_ImGuiListClipperRangeSizeOf is the byte size of ImVector_ImGuiListClipperRange.
const Vector_ImGuiListClipperRangeSizeOf = int(C.sizeof_ImVector_ImGuiListClipperRange)

// Vector_ImGuiListClipperRangeAlloc allocates a continuous block of Vector_ImGuiListClipperRange.
func Vector_ImGuiListClipperRangeAlloc(alloc ffi.Allocator, count int) Vector_ImGuiListClipperRange {
	ptr := alloc.Allocate(Vector_ImGuiListClipperRangeSizeOf * count)
	return Vector_ImGuiListClipperRange(ptr)
}

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p Vector_ImGuiListClipperRange) Offset(offset int) Vector_ImGuiListClipperRange {
	return p + Vector_ImGuiListClipperRange(offset*Vector_ImGuiListClipperRangeSizeOf)
}

// GetCapacity returns the value in Capacity.
func (p Vector_ImGuiListClipperRange) GetCapacity() int32 {
	ptr := (*C.ImVector_ImGuiListClipperRange)(unsafe.Pointer(p))
	return int32(ptr.Capacity)
}

// SetCapacity sets the value in Capacity.
func (p Vector_ImGuiListClipperRange) SetCapacity(value int32) {
	ptr := (*C.ImVector_ImGuiListClipperRange)(unsafe.Pointer(p))
	ptr.Capacity = (C.int32_t)(value)
}

// GetData returns the value in Data.
func (p Vector_ImGuiListClipperRange) GetData() GuiListClipperRange {
	ptr := (*C.ImVector_ImGuiListClipperRange)(unsafe.Pointer(p))
	return GuiListClipperRange(unsafe.Pointer(ptr.Data))
}

// SetData sets the value in Data.
func (p Vector_ImGuiListClipperRange) SetData(value GuiListClipperRange) {
	ptr := (*C.ImVector_ImGuiListClipperRange)(unsafe.Pointer(p))
	ptr.Data = (*C.ImGuiListClipperRange)(unsafe.Pointer(value))
}

// GetSize returns the value in Size.
func (p Vector_ImGuiListClipperRange) GetSize() int32 {
	ptr := (*C.ImVector_ImGuiListClipperRange)(unsafe.Pointer(p))
	return int32(ptr.Size)
}

// SetSize sets the value in Size.
func (p Vector_ImGuiListClipperRange) SetSize(value int32) {
	ptr := (*C.ImVector_ImGuiListClipperRange)(unsafe.Pointer(p))
	ptr.Size = (C.int32_t)(value)
}

// Vector_ImGuiMultiSelectState wraps struct ImVector_ImGuiMultiSelectState.
type Vector_ImGuiMultiSelectState uintptr

// Vector_ImGuiMultiSelectStateNil is a null pointer.
var Vector_ImGuiMultiSelectStateNil Vector_ImGuiMultiSelectState

// Vector_ImGuiMultiSelectStateSizeOf is the byte size of ImVector_ImGuiMultiSelectState.
const Vector_ImGuiMultiSelectStateSizeOf = int(C.sizeof_ImVector_ImGuiMultiSelectState)

// Vector_ImGuiMultiSelectStateAlloc allocates a continuous block of Vector_ImGuiMultiSelectState.
func Vector_ImGuiMultiSelectStateAlloc(alloc ffi.Allocator, count int) Vector_ImGuiMultiSelectState {
	ptr := alloc.Allocate(Vector_ImGuiMultiSelectStateSizeOf * count)
	return Vector_ImGuiMultiSelectState(ptr)
}

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p Vector_ImGuiMultiSelectState) Offset(offset int) Vector_ImGuiMultiSelectState {
	return p + Vector_ImGuiMultiSelectState(offset*Vector_ImGuiMultiSelectStateSizeOf)
}

// GetCapacity returns the value in Capacity.
func (p Vector_ImGuiMultiSelectState) GetCapacity() int32 {
	ptr := (*C.ImVector_ImGuiMultiSelectState)(unsafe.Pointer(p))
	return int32(ptr.Capacity)
}

// SetCapacity sets the value in Capacity.
func (p Vector_ImGuiMultiSelectState) SetCapacity(value int32) {
	ptr := (*C.ImVector_ImGuiMultiSelectState)(unsafe.Pointer(p))
	ptr.Capacity = (C.int32_t)(value)
}

// GetData returns the value in Data.
func (p Vector_ImGuiMultiSelectState) GetData() GuiMultiSelectState {
	ptr := (*C.ImVector_ImGuiMultiSelectState)(unsafe.Pointer(p))
	return GuiMultiSelectState(unsafe.Pointer(ptr.Data))
}

// SetData sets the value in Data.
func (p Vector_ImGuiMultiSelectState) SetData(value GuiMultiSelectState) {
	ptr := (*C.ImVector_ImGuiMultiSelectState)(unsafe.Pointer(p))
	ptr.Data = (*C.ImGuiMultiSelectState)(unsafe.Pointer(value))
}

// GetSize returns the value in Size.
func (p Vector_ImGuiMultiSelectState) GetSize() int32 {
	ptr := (*C.ImVector_ImGuiMultiSelectState)(unsafe.Pointer(p))
	return int32(ptr.Size)
}

// SetSize sets the value in Size.
func (p Vector_ImGuiMultiSelectState) SetSize(value int32) {
	ptr := (*C.ImVector_ImGuiMultiSelectState)(unsafe.Pointer(p))
	ptr.Size = (C.int32_t)(value)
}

// Vector_ImGuiMultiSelectTempData wraps struct ImVector_ImGuiMultiSelectTempData.
type Vector_ImGuiMultiSelectTempData uintptr

// Vector_ImGuiMultiSelectTempDataNil is a null pointer.
var Vector_ImGuiMultiSelectTempDataNil Vector_ImGuiMultiSelectTempData

// Vector_ImGuiMultiSelectTempDataSizeOf is the byte size of ImVector_ImGuiMultiSelectTempData.
const Vector_ImGuiMultiSelectTempDataSizeOf = int(C.sizeof_ImVector_ImGuiMultiSelectTempData)

// Vector_ImGuiMultiSelectTempDataAlloc allocates a continuous block of Vector_ImGuiMultiSelectTempData.
func Vector_ImGuiMultiSelectTempDataAlloc(alloc ffi.Allocator, count int) Vector_ImGuiMultiSelectTempData {
	ptr := alloc.Allocate(Vector_ImGuiMultiSelectTempDataSizeOf * count)
	return Vector_ImGuiMultiSelectTempData(ptr)
}

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p Vector_ImGuiMultiSelectTempData) Offset(offset int) Vector_ImGuiMultiSelectTempData {
	return p + Vector_ImGuiMultiSelectTempData(offset*Vector_ImGuiMultiSelectTempDataSizeOf)
}

// GetCapacity returns the value in Capacity.
func (p Vector_ImGuiMultiSelectTempData) GetCapacity() int32 {
	ptr := (*C.ImVector_ImGuiMultiSelectTempData)(unsafe.Pointer(p))
	return int32(ptr.Capacity)
}

// SetCapacity sets the value in Capacity.
func (p Vector_ImGuiMultiSelectTempData) SetCapacity(value int32) {
	ptr := (*C.ImVector_ImGuiMultiSelectTempData)(unsafe.Pointer(p))
	ptr.Capacity = (C.int32_t)(value)
}

// GetData returns the value in Data.
func (p Vector_ImGuiMultiSelectTempData) GetData() GuiMultiSelectTempData {
	ptr := (*C.ImVector_ImGuiMultiSelectTempData)(unsafe.Pointer(p))
	return GuiMultiSelectTempData(unsafe.Pointer(ptr.Data))
}

// SetData sets the value in Data.
func (p Vector_ImGuiMultiSelectTempData) SetData(value GuiMultiSelectTempData) {
	ptr := (*C.ImVector_ImGuiMultiSelectTempData)(unsafe.Pointer(p))
	ptr.Data = (*C.ImGuiMultiSelectTempData)(unsafe.Pointer(value))
}

// GetSize returns the value in Size.
func (p Vector_ImGuiMultiSelectTempData) GetSize() int32 {
	ptr := (*C.ImVector_ImGuiMultiSelectTempData)(unsafe.Pointer(p))
	return int32(ptr.Size)
}

// SetSize sets the value in Size.
func (p Vector_ImGuiMultiSelectTempData) SetSize(value int32) {
	ptr := (*C.ImVector_ImGuiMultiSelectTempData)(unsafe.Pointer(p))
	ptr.Size = (C.int32_t)(value)
}

// Vector_ImGuiOldColumnData wraps struct ImVector_ImGuiOldColumnData.
type Vector_ImGuiOldColumnData uintptr

// Vector_ImGuiOldColumnDataNil is a null pointer.
var Vector_ImGuiOldColumnDataNil Vector_ImGuiOldColumnData

// Vector_ImGuiOldColumnDataSizeOf is the byte size of ImVector_ImGuiOldColumnData.
const Vector_ImGuiOldColumnDataSizeOf = int(C.sizeof_ImVector_ImGuiOldColumnData)

// Vector_ImGuiOldColumnDataAlloc allocates a continuous block of Vector_ImGuiOldColumnData.
func Vector_ImGuiOldColumnDataAlloc(alloc ffi.Allocator, count int) Vector_ImGuiOldColumnData {
	ptr := alloc.Allocate(Vector_ImGuiOldColumnDataSizeOf * count)
	return Vector_ImGuiOldColumnData(ptr)
}

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p Vector_ImGuiOldColumnData) Offset(offset int) Vector_ImGuiOldColumnData {
	return p + Vector_ImGuiOldColumnData(offset*Vector_ImGuiOldColumnDataSizeOf)
}

// GetCapacity returns the value in Capacity.
func (p Vector_ImGuiOldColumnData) GetCapacity() int32 {
	ptr := (*C.ImVector_ImGuiOldColumnData)(unsafe.Pointer(p))
	return int32(ptr.Capacity)
}

// SetCapacity sets the value in Capacity.
func (p Vector_ImGuiOldColumnData) SetCapacity(value int32) {
	ptr := (*C.ImVector_ImGuiOldColumnData)(unsafe.Pointer(p))
	ptr.Capacity = (C.int32_t)(value)
}

// GetData returns the value in Data.
func (p Vector_ImGuiOldColumnData) GetData() GuiOldColumnData {
	ptr := (*C.ImVector_ImGuiOldColumnData)(unsafe.Pointer(p))
	return GuiOldColumnData(unsafe.Pointer(ptr.Data))
}

// SetData sets the value in Data.
func (p Vector_ImGuiOldColumnData) SetData(value GuiOldColumnData) {
	ptr := (*C.ImVector_ImGuiOldColumnData)(unsafe.Pointer(p))
	ptr.Data = (*C.ImGuiOldColumnData)(unsafe.Pointer(value))
}

// GetSize returns the value in Size.
func (p Vector_ImGuiOldColumnData) GetSize() int32 {
	ptr := (*C.ImVector_ImGuiOldColumnData)(unsafe.Pointer(p))
	return int32(ptr.Size)
}

// SetSize sets the value in Size.
func (p Vector_ImGuiOldColumnData) SetSize(value int32) {
	ptr := (*C.ImVector_ImGuiOldColumnData)(unsafe.Pointer(p))
	ptr.Size = (C.int32_t)(value)
}

// Vector_ImGuiOldColumns wraps struct ImVector_ImGuiOldColumns.
type Vector_ImGuiOldColumns uintptr

// Vector_ImGuiOldColumnsNil is a null pointer.
var Vector_ImGuiOldColumnsNil Vector_ImGuiOldColumns

// Vector_ImGuiOldColumnsSizeOf is the byte size of ImVector_ImGuiOldColumns.
const Vector_ImGuiOldColumnsSizeOf = int(C.sizeof_ImVector_ImGuiOldColumns)

// Vector_ImGuiOldColumnsAlloc allocates a continuous block of Vector_ImGuiOldColumns.
func Vector_ImGuiOldColumnsAlloc(alloc ffi.Allocator, count int) Vector_ImGuiOldColumns {
	ptr := alloc.Allocate(Vector_ImGuiOldColumnsSizeOf * count)
	return Vector_ImGuiOldColumns(ptr)
}

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p Vector_ImGuiOldColumns) Offset(offset int) Vector_ImGuiOldColumns {
	return p + Vector_ImGuiOldColumns(offset*Vector_ImGuiOldColumnsSizeOf)
}

// GetCapacity returns the value in Capacity.
func (p Vector_ImGuiOldColumns) GetCapacity() int32 {
	ptr := (*C.ImVector_ImGuiOldColumns)(unsafe.Pointer(p))
	return int32(ptr.Capacity)
}

// SetCapacity sets the value in Capacity.
func (p Vector_ImGuiOldColumns) SetCapacity(value int32) {
	ptr := (*C.ImVector_ImGuiOldColumns)(unsafe.Pointer(p))
	ptr.Capacity = (C.int32_t)(value)
}

// GetData returns the value in Data.
func (p Vector_ImGuiOldColumns) GetData() GuiOldColumns {
	ptr := (*C.ImVector_ImGuiOldColumns)(unsafe.Pointer(p))
	return GuiOldColumns(unsafe.Pointer(ptr.Data))
}

// SetData sets the value in Data.
func (p Vector_ImGuiOldColumns) SetData(value GuiOldColumns) {
	ptr := (*C.ImVector_ImGuiOldColumns)(unsafe.Pointer(p))
	ptr.Data = (*C.ImGuiOldColumns)(unsafe.Pointer(value))
}

// GetSize returns the value in Size.
func (p Vector_ImGuiOldColumns) GetSize() int32 {
	ptr := (*C.ImVector_ImGuiOldColumns)(unsafe.Pointer(p))
	return int32(ptr.Size)
}

// SetSize sets the value in Size.
func (p Vector_ImGuiOldColumns) SetSize(value int32) {
	ptr := (*C.ImVector_ImGuiOldColumns)(unsafe.Pointer(p))
	ptr.Size = (C.int32_t)(value)
}

// Vector_ImGuiPlatformMonitor wraps struct ImVector_ImGuiPlatformMonitor.
type Vector_ImGuiPlatformMonitor uintptr

// Vector_ImGuiPlatformMonitorNil is a null pointer.
var Vector_ImGuiPlatformMonitorNil Vector_ImGuiPlatformMonitor

// Vector_ImGuiPlatformMonitorSizeOf is the byte size of ImVector_ImGuiPlatformMonitor.
const Vector_ImGuiPlatformMonitorSizeOf = int(C.sizeof_ImVector_ImGuiPlatformMonitor)

// Vector_ImGuiPlatformMonitorAlloc allocates a continuous block of Vector_ImGuiPlatformMonitor.
func Vector_ImGuiPlatformMonitorAlloc(alloc ffi.Allocator, count int) Vector_ImGuiPlatformMonitor {
	ptr := alloc.Allocate(Vector_ImGuiPlatformMonitorSizeOf * count)
	return Vector_ImGuiPlatformMonitor(ptr)
}

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p Vector_ImGuiPlatformMonitor) Offset(offset int) Vector_ImGuiPlatformMonitor {
	return p + Vector_ImGuiPlatformMonitor(offset*Vector_ImGuiPlatformMonitorSizeOf)
}

// GetCapacity returns the value in Capacity.
func (p Vector_ImGuiPlatformMonitor) GetCapacity() int32 {
	ptr := (*C.ImVector_ImGuiPlatformMonitor)(unsafe.Pointer(p))
	return int32(ptr.Capacity)
}

// SetCapacity sets the value in Capacity.
func (p Vector_ImGuiPlatformMonitor) SetCapacity(value int32) {
	ptr := (*C.ImVector_ImGuiPlatformMonitor)(unsafe.Pointer(p))
	ptr.Capacity = (C.int32_t)(value)
}

// GetData returns the value in Data.
func (p Vector_ImGuiPlatformMonitor) GetData() GuiPlatformMonitor {
	ptr := (*C.ImVector_ImGuiPlatformMonitor)(unsafe.Pointer(p))
	return GuiPlatformMonitor(unsafe.Pointer(ptr.Data))
}

// SetData sets the value in Data.
func (p Vector_ImGuiPlatformMonitor) SetData(value GuiPlatformMonitor) {
	ptr := (*C.ImVector_ImGuiPlatformMonitor)(unsafe.Pointer(p))
	ptr.Data = (*C.ImGuiPlatformMonitor)(unsafe.Pointer(value))
}

// GetSize returns the value in Size.
func (p Vector_ImGuiPlatformMonitor) GetSize() int32 {
	ptr := (*C.ImVector_ImGuiPlatformMonitor)(unsafe.Pointer(p))
	return int32(ptr.Size)
}

// SetSize sets the value in Size.
func (p Vector_ImGuiPlatformMonitor) SetSize(value int32) {
	ptr := (*C.ImVector_ImGuiPlatformMonitor)(unsafe.Pointer(p))
	ptr.Size = (C.int32_t)(value)
}

// Vector_ImGuiPopupData wraps struct ImVector_ImGuiPopupData.
type Vector_ImGuiPopupData uintptr

// Vector_ImGuiPopupDataNil is a null pointer.
var Vector_ImGuiPopupDataNil Vector_ImGuiPopupData

// Vector_ImGuiPopupDataSizeOf is the byte size of ImVector_ImGuiPopupData.
const Vector_ImGuiPopupDataSizeOf = int(C.sizeof_ImVector_ImGuiPopupData)

// Vector_ImGuiPopupDataAlloc allocates a continuous block of Vector_ImGuiPopupData.
func Vector_ImGuiPopupDataAlloc(alloc ffi.Allocator, count int) Vector_ImGuiPopupData {
	ptr := alloc.Allocate(Vector_ImGuiPopupDataSizeOf * count)
	return Vector_ImGuiPopupData(ptr)
}

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p Vector_ImGuiPopupData) Offset(offset int) Vector_ImGuiPopupData {
	return p + Vector_ImGuiPopupData(offset*Vector_ImGuiPopupDataSizeOf)
}

// GetCapacity returns the value in Capacity.
func (p Vector_ImGuiPopupData) GetCapacity() int32 {
	ptr := (*C.ImVector_ImGuiPopupData)(unsafe.Pointer(p))
	return int32(ptr.Capacity)
}

// SetCapacity sets the value in Capacity.
func (p Vector_ImGuiPopupData) SetCapacity(value int32) {
	ptr := (*C.ImVector_ImGuiPopupData)(unsafe.Pointer(p))
	ptr.Capacity = (C.int32_t)(value)
}

// GetData returns the value in Data.
func (p Vector_ImGuiPopupData) GetData() GuiPopupData {
	ptr := (*C.ImVector_ImGuiPopupData)(unsafe.Pointer(p))
	return GuiPopupData(unsafe.Pointer(ptr.Data))
}

// SetData sets the value in Data.
func (p Vector_ImGuiPopupData) SetData(value GuiPopupData) {
	ptr := (*C.ImVector_ImGuiPopupData)(unsafe.Pointer(p))
	ptr.Data = (*C.ImGuiPopupData)(unsafe.Pointer(value))
}

// GetSize returns the value in Size.
func (p Vector_ImGuiPopupData) GetSize() int32 {
	ptr := (*C.ImVector_ImGuiPopupData)(unsafe.Pointer(p))
	return int32(ptr.Size)
}

// SetSize sets the value in Size.
func (p Vector_ImGuiPopupData) SetSize(value int32) {
	ptr := (*C.ImVector_ImGuiPopupData)(unsafe.Pointer(p))
	ptr.Size = (C.int32_t)(value)
}

// Vector_ImGuiPtrOrIndex wraps struct ImVector_ImGuiPtrOrIndex.
type Vector_ImGuiPtrOrIndex uintptr

// Vector_ImGuiPtrOrIndexNil is a null pointer.
var Vector_ImGuiPtrOrIndexNil Vector_ImGuiPtrOrIndex

// Vector_ImGuiPtrOrIndexSizeOf is the byte size of ImVector_ImGuiPtrOrIndex.
const Vector_ImGuiPtrOrIndexSizeOf = int(C.sizeof_ImVector_ImGuiPtrOrIndex)

// Vector_ImGuiPtrOrIndexAlloc allocates a continuous block of Vector_ImGuiPtrOrIndex.
func Vector_ImGuiPtrOrIndexAlloc(alloc ffi.Allocator, count int) Vector_ImGuiPtrOrIndex {
	ptr := alloc.Allocate(Vector_ImGuiPtrOrIndexSizeOf * count)
	return Vector_ImGuiPtrOrIndex(ptr)
}

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p Vector_ImGuiPtrOrIndex) Offset(offset int) Vector_ImGuiPtrOrIndex {
	return p + Vector_ImGuiPtrOrIndex(offset*Vector_ImGuiPtrOrIndexSizeOf)
}

// GetCapacity returns the value in Capacity.
func (p Vector_ImGuiPtrOrIndex) GetCapacity() int32 {
	ptr := (*C.ImVector_ImGuiPtrOrIndex)(unsafe.Pointer(p))
	return int32(ptr.Capacity)
}

// SetCapacity sets the value in Capacity.
func (p Vector_ImGuiPtrOrIndex) SetCapacity(value int32) {
	ptr := (*C.ImVector_ImGuiPtrOrIndex)(unsafe.Pointer(p))
	ptr.Capacity = (C.int32_t)(value)
}

// GetData returns the value in Data.
func (p Vector_ImGuiPtrOrIndex) GetData() GuiPtrOrIndex {
	ptr := (*C.ImVector_ImGuiPtrOrIndex)(unsafe.Pointer(p))
	return GuiPtrOrIndex(unsafe.Pointer(ptr.Data))
}

// SetData sets the value in Data.
func (p Vector_ImGuiPtrOrIndex) SetData(value GuiPtrOrIndex) {
	ptr := (*C.ImVector_ImGuiPtrOrIndex)(unsafe.Pointer(p))
	ptr.Data = (*C.ImGuiPtrOrIndex)(unsafe.Pointer(value))
}

// GetSize returns the value in Size.
func (p Vector_ImGuiPtrOrIndex) GetSize() int32 {
	ptr := (*C.ImVector_ImGuiPtrOrIndex)(unsafe.Pointer(p))
	return int32(ptr.Size)
}

// SetSize sets the value in Size.
func (p Vector_ImGuiPtrOrIndex) SetSize(value int32) {
	ptr := (*C.ImVector_ImGuiPtrOrIndex)(unsafe.Pointer(p))
	ptr.Size = (C.int32_t)(value)
}

// Vector_ImGuiSelectionRequest wraps struct ImVector_ImGuiSelectionRequest.
type Vector_ImGuiSelectionRequest uintptr

// Vector_ImGuiSelectionRequestNil is a null pointer.
var Vector_ImGuiSelectionRequestNil Vector_ImGuiSelectionRequest

// Vector_ImGuiSelectionRequestSizeOf is the byte size of ImVector_ImGuiSelectionRequest.
const Vector_ImGuiSelectionRequestSizeOf = int(C.sizeof_ImVector_ImGuiSelectionRequest)

// Vector_ImGuiSelectionRequestAlloc allocates a continuous block of Vector_ImGuiSelectionRequest.
func Vector_ImGuiSelectionRequestAlloc(alloc ffi.Allocator, count int) Vector_ImGuiSelectionRequest {
	ptr := alloc.Allocate(Vector_ImGuiSelectionRequestSizeOf * count)
	return Vector_ImGuiSelectionRequest(ptr)
}

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p Vector_ImGuiSelectionRequest) Offset(offset int) Vector_ImGuiSelectionRequest {
	return p + Vector_ImGuiSelectionRequest(offset*Vector_ImGuiSelectionRequestSizeOf)
}

// GetCapacity returns the value in Capacity.
func (p Vector_ImGuiSelectionRequest) GetCapacity() int32 {
	ptr := (*C.ImVector_ImGuiSelectionRequest)(unsafe.Pointer(p))
	return int32(ptr.Capacity)
}

// SetCapacity sets the value in Capacity.
func (p Vector_ImGuiSelectionRequest) SetCapacity(value int32) {
	ptr := (*C.ImVector_ImGuiSelectionRequest)(unsafe.Pointer(p))
	ptr.Capacity = (C.int32_t)(value)
}

// GetData returns the value in Data.
func (p Vector_ImGuiSelectionRequest) GetData() GuiSelectionRequest {
	ptr := (*C.ImVector_ImGuiSelectionRequest)(unsafe.Pointer(p))
	return GuiSelectionRequest(unsafe.Pointer(ptr.Data))
}

// SetData sets the value in Data.
func (p Vector_ImGuiSelectionRequest) SetData(value GuiSelectionRequest) {
	ptr := (*C.ImVector_ImGuiSelectionRequest)(unsafe.Pointer(p))
	ptr.Data = (*C.ImGuiSelectionRequest)(unsafe.Pointer(value))
}

// GetSize returns the value in Size.
func (p Vector_ImGuiSelectionRequest) GetSize() int32 {
	ptr := (*C.ImVector_ImGuiSelectionRequest)(unsafe.Pointer(p))
	return int32(ptr.Size)
}

// SetSize sets the value in Size.
func (p Vector_ImGuiSelectionRequest) SetSize(value int32) {
	ptr := (*C.ImVector_ImGuiSelectionRequest)(unsafe.Pointer(p))
	ptr.Size = (C.int32_t)(value)
}

// Vector_ImGuiSettingsHandler wraps struct ImVector_ImGuiSettingsHandler.
type Vector_ImGuiSettingsHandler uintptr

// Vector_ImGuiSettingsHandlerNil is a null pointer.
var Vector_ImGuiSettingsHandlerNil Vector_ImGuiSettingsHandler

// Vector_ImGuiSettingsHandlerSizeOf is the byte size of ImVector_ImGuiSettingsHandler.
const Vector_ImGuiSettingsHandlerSizeOf = int(C.sizeof_ImVector_ImGuiSettingsHandler)

// Vector_ImGuiSettingsHandlerAlloc allocates a continuous block of Vector_ImGuiSettingsHandler.
func Vector_ImGuiSettingsHandlerAlloc(alloc ffi.Allocator, count int) Vector_ImGuiSettingsHandler {
	ptr := alloc.Allocate(Vector_ImGuiSettingsHandlerSizeOf * count)
	return Vector_ImGuiSettingsHandler(ptr)
}

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p Vector_ImGuiSettingsHandler) Offset(offset int) Vector_ImGuiSettingsHandler {
	return p + Vector_ImGuiSettingsHandler(offset*Vector_ImGuiSettingsHandlerSizeOf)
}

// GetCapacity returns the value in Capacity.
func (p Vector_ImGuiSettingsHandler) GetCapacity() int32 {
	ptr := (*C.ImVector_ImGuiSettingsHandler)(unsafe.Pointer(p))
	return int32(ptr.Capacity)
}

// SetCapacity sets the value in Capacity.
func (p Vector_ImGuiSettingsHandler) SetCapacity(value int32) {
	ptr := (*C.ImVector_ImGuiSettingsHandler)(unsafe.Pointer(p))
	ptr.Capacity = (C.int32_t)(value)
}

// GetData returns the value in Data.
func (p Vector_ImGuiSettingsHandler) GetData() GuiSettingsHandler {
	ptr := (*C.ImVector_ImGuiSettingsHandler)(unsafe.Pointer(p))
	return GuiSettingsHandler(unsafe.Pointer(ptr.Data))
}

// SetData sets the value in Data.
func (p Vector_ImGuiSettingsHandler) SetData(value GuiSettingsHandler) {
	ptr := (*C.ImVector_ImGuiSettingsHandler)(unsafe.Pointer(p))
	ptr.Data = (*C.ImGuiSettingsHandler)(unsafe.Pointer(value))
}

// GetSize returns the value in Size.
func (p Vector_ImGuiSettingsHandler) GetSize() int32 {
	ptr := (*C.ImVector_ImGuiSettingsHandler)(unsafe.Pointer(p))
	return int32(ptr.Size)
}

// SetSize sets the value in Size.
func (p Vector_ImGuiSettingsHandler) SetSize(value int32) {
	ptr := (*C.ImVector_ImGuiSettingsHandler)(unsafe.Pointer(p))
	ptr.Size = (C.int32_t)(value)
}

// Vector_ImGuiShrinkWidthItem wraps struct ImVector_ImGuiShrinkWidthItem.
type Vector_ImGuiShrinkWidthItem uintptr

// Vector_ImGuiShrinkWidthItemNil is a null pointer.
var Vector_ImGuiShrinkWidthItemNil Vector_ImGuiShrinkWidthItem

// Vector_ImGuiShrinkWidthItemSizeOf is the byte size of ImVector_ImGuiShrinkWidthItem.
const Vector_ImGuiShrinkWidthItemSizeOf = int(C.sizeof_ImVector_ImGuiShrinkWidthItem)

// Vector_ImGuiShrinkWidthItemAlloc allocates a continuous block of Vector_ImGuiShrinkWidthItem.
func Vector_ImGuiShrinkWidthItemAlloc(alloc ffi.Allocator, count int) Vector_ImGuiShrinkWidthItem {
	ptr := alloc.Allocate(Vector_ImGuiShrinkWidthItemSizeOf * count)
	return Vector_ImGuiShrinkWidthItem(ptr)
}

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p Vector_ImGuiShrinkWidthItem) Offset(offset int) Vector_ImGuiShrinkWidthItem {
	return p + Vector_ImGuiShrinkWidthItem(offset*Vector_ImGuiShrinkWidthItemSizeOf)
}

// GetCapacity returns the value in Capacity.
func (p Vector_ImGuiShrinkWidthItem) GetCapacity() int32 {
	ptr := (*C.ImVector_ImGuiShrinkWidthItem)(unsafe.Pointer(p))
	return int32(ptr.Capacity)
}

// SetCapacity sets the value in Capacity.
func (p Vector_ImGuiShrinkWidthItem) SetCapacity(value int32) {
	ptr := (*C.ImVector_ImGuiShrinkWidthItem)(unsafe.Pointer(p))
	ptr.Capacity = (C.int32_t)(value)
}

// GetData returns the value in Data.
func (p Vector_ImGuiShrinkWidthItem) GetData() GuiShrinkWidthItem {
	ptr := (*C.ImVector_ImGuiShrinkWidthItem)(unsafe.Pointer(p))
	return GuiShrinkWidthItem(unsafe.Pointer(ptr.Data))
}

// SetData sets the value in Data.
func (p Vector_ImGuiShrinkWidthItem) SetData(value GuiShrinkWidthItem) {
	ptr := (*C.ImVector_ImGuiShrinkWidthItem)(unsafe.Pointer(p))
	ptr.Data = (*C.ImGuiShrinkWidthItem)(unsafe.Pointer(value))
}

// GetSize returns the value in Size.
func (p Vector_ImGuiShrinkWidthItem) GetSize() int32 {
	ptr := (*C.ImVector_ImGuiShrinkWidthItem)(unsafe.Pointer(p))
	return int32(ptr.Size)
}

// SetSize sets the value in Size.
func (p Vector_ImGuiShrinkWidthItem) SetSize(value int32) {
	ptr := (*C.ImVector_ImGuiShrinkWidthItem)(unsafe.Pointer(p))
	ptr.Size = (C.int32_t)(value)
}

// Vector_ImGuiStackLevelInfo wraps struct ImVector_ImGuiStackLevelInfo.
type Vector_ImGuiStackLevelInfo uintptr

// Vector_ImGuiStackLevelInfoNil is a null pointer.
var Vector_ImGuiStackLevelInfoNil Vector_ImGuiStackLevelInfo

// Vector_ImGuiStackLevelInfoSizeOf is the byte size of ImVector_ImGuiStackLevelInfo.
const Vector_ImGuiStackLevelInfoSizeOf = int(C.sizeof_ImVector_ImGuiStackLevelInfo)

// Vector_ImGuiStackLevelInfoAlloc allocates a continuous block of Vector_ImGuiStackLevelInfo.
func Vector_ImGuiStackLevelInfoAlloc(alloc ffi.Allocator, count int) Vector_ImGuiStackLevelInfo {
	ptr := alloc.Allocate(Vector_ImGuiStackLevelInfoSizeOf * count)
	return Vector_ImGuiStackLevelInfo(ptr)
}

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p Vector_ImGuiStackLevelInfo) Offset(offset int) Vector_ImGuiStackLevelInfo {
	return p + Vector_ImGuiStackLevelInfo(offset*Vector_ImGuiStackLevelInfoSizeOf)
}

// GetCapacity returns the value in Capacity.
func (p Vector_ImGuiStackLevelInfo) GetCapacity() int32 {
	ptr := (*C.ImVector_ImGuiStackLevelInfo)(unsafe.Pointer(p))
	return int32(ptr.Capacity)
}

// SetCapacity sets the value in Capacity.
func (p Vector_ImGuiStackLevelInfo) SetCapacity(value int32) {
	ptr := (*C.ImVector_ImGuiStackLevelInfo)(unsafe.Pointer(p))
	ptr.Capacity = (C.int32_t)(value)
}

// GetData returns the value in Data.
func (p Vector_ImGuiStackLevelInfo) GetData() GuiStackLevelInfo {
	ptr := (*C.ImVector_ImGuiStackLevelInfo)(unsafe.Pointer(p))
	return GuiStackLevelInfo(unsafe.Pointer(ptr.Data))
}

// SetData sets the value in Data.
func (p Vector_ImGuiStackLevelInfo) SetData(value GuiStackLevelInfo) {
	ptr := (*C.ImVector_ImGuiStackLevelInfo)(unsafe.Pointer(p))
	ptr.Data = (*C.ImGuiStackLevelInfo)(unsafe.Pointer(value))
}

// GetSize returns the value in Size.
func (p Vector_ImGuiStackLevelInfo) GetSize() int32 {
	ptr := (*C.ImVector_ImGuiStackLevelInfo)(unsafe.Pointer(p))
	return int32(ptr.Size)
}

// SetSize sets the value in Size.
func (p Vector_ImGuiStackLevelInfo) SetSize(value int32) {
	ptr := (*C.ImVector_ImGuiStackLevelInfo)(unsafe.Pointer(p))
	ptr.Size = (C.int32_t)(value)
}

// Vector_ImGuiStoragePair wraps struct ImVector_ImGuiStoragePair.
type Vector_ImGuiStoragePair uintptr

// Vector_ImGuiStoragePairNil is a null pointer.
var Vector_ImGuiStoragePairNil Vector_ImGuiStoragePair

// Vector_ImGuiStoragePairSizeOf is the byte size of ImVector_ImGuiStoragePair.
const Vector_ImGuiStoragePairSizeOf = int(C.sizeof_ImVector_ImGuiStoragePair)

// Vector_ImGuiStoragePairAlloc allocates a continuous block of Vector_ImGuiStoragePair.
func Vector_ImGuiStoragePairAlloc(alloc ffi.Allocator, count int) Vector_ImGuiStoragePair {
	ptr := alloc.Allocate(Vector_ImGuiStoragePairSizeOf * count)
	return Vector_ImGuiStoragePair(ptr)
}

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p Vector_ImGuiStoragePair) Offset(offset int) Vector_ImGuiStoragePair {
	return p + Vector_ImGuiStoragePair(offset*Vector_ImGuiStoragePairSizeOf)
}

// GetCapacity returns the value in Capacity.
func (p Vector_ImGuiStoragePair) GetCapacity() int32 {
	ptr := (*C.ImVector_ImGuiStoragePair)(unsafe.Pointer(p))
	return int32(ptr.Capacity)
}

// SetCapacity sets the value in Capacity.
func (p Vector_ImGuiStoragePair) SetCapacity(value int32) {
	ptr := (*C.ImVector_ImGuiStoragePair)(unsafe.Pointer(p))
	ptr.Capacity = (C.int32_t)(value)
}

// GetData returns the value in Data.
func (p Vector_ImGuiStoragePair) GetData() GuiStoragePair {
	ptr := (*C.ImVector_ImGuiStoragePair)(unsafe.Pointer(p))
	return GuiStoragePair(unsafe.Pointer(ptr.Data))
}

// SetData sets the value in Data.
func (p Vector_ImGuiStoragePair) SetData(value GuiStoragePair) {
	ptr := (*C.ImVector_ImGuiStoragePair)(unsafe.Pointer(p))
	ptr.Data = (*C.ImGuiStoragePair)(unsafe.Pointer(value))
}

// GetSize returns the value in Size.
func (p Vector_ImGuiStoragePair) GetSize() int32 {
	ptr := (*C.ImVector_ImGuiStoragePair)(unsafe.Pointer(p))
	return int32(ptr.Size)
}

// SetSize sets the value in Size.
func (p Vector_ImGuiStoragePair) SetSize(value int32) {
	ptr := (*C.ImVector_ImGuiStoragePair)(unsafe.Pointer(p))
	ptr.Size = (C.int32_t)(value)
}

// Vector_ImGuiStyleMod wraps struct ImVector_ImGuiStyleMod.
type Vector_ImGuiStyleMod uintptr

// Vector_ImGuiStyleModNil is a null pointer.
var Vector_ImGuiStyleModNil Vector_ImGuiStyleMod

// Vector_ImGuiStyleModSizeOf is the byte size of ImVector_ImGuiStyleMod.
const Vector_ImGuiStyleModSizeOf = int(C.sizeof_ImVector_ImGuiStyleMod)

// Vector_ImGuiStyleModAlloc allocates a continuous block of Vector_ImGuiStyleMod.
func Vector_ImGuiStyleModAlloc(alloc ffi.Allocator, count int) Vector_ImGuiStyleMod {
	ptr := alloc.Allocate(Vector_ImGuiStyleModSizeOf * count)
	return Vector_ImGuiStyleMod(ptr)
}

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p Vector_ImGuiStyleMod) Offset(offset int) Vector_ImGuiStyleMod {
	return p + Vector_ImGuiStyleMod(offset*Vector_ImGuiStyleModSizeOf)
}

// GetCapacity returns the value in Capacity.
func (p Vector_ImGuiStyleMod) GetCapacity() int32 {
	ptr := (*C.ImVector_ImGuiStyleMod)(unsafe.Pointer(p))
	return int32(ptr.Capacity)
}

// SetCapacity sets the value in Capacity.
func (p Vector_ImGuiStyleMod) SetCapacity(value int32) {
	ptr := (*C.ImVector_ImGuiStyleMod)(unsafe.Pointer(p))
	ptr.Capacity = (C.int32_t)(value)
}

// GetData returns the value in Data.
func (p Vector_ImGuiStyleMod) GetData() GuiStyleMod {
	ptr := (*C.ImVector_ImGuiStyleMod)(unsafe.Pointer(p))
	return GuiStyleMod(unsafe.Pointer(ptr.Data))
}

// SetData sets the value in Data.
func (p Vector_ImGuiStyleMod) SetData(value GuiStyleMod) {
	ptr := (*C.ImVector_ImGuiStyleMod)(unsafe.Pointer(p))
	ptr.Data = (*C.ImGuiStyleMod)(unsafe.Pointer(value))
}

// GetSize returns the value in Size.
func (p Vector_ImGuiStyleMod) GetSize() int32 {
	ptr := (*C.ImVector_ImGuiStyleMod)(unsafe.Pointer(p))
	return int32(ptr.Size)
}

// SetSize sets the value in Size.
func (p Vector_ImGuiStyleMod) SetSize(value int32) {
	ptr := (*C.ImVector_ImGuiStyleMod)(unsafe.Pointer(p))
	ptr.Size = (C.int32_t)(value)
}

// Vector_ImGuiTabBar wraps struct ImVector_ImGuiTabBar.
type Vector_ImGuiTabBar uintptr

// Vector_ImGuiTabBarNil is a null pointer.
var Vector_ImGuiTabBarNil Vector_ImGuiTabBar

// Vector_ImGuiTabBarSizeOf is the byte size of ImVector_ImGuiTabBar.
const Vector_ImGuiTabBarSizeOf = int(C.sizeof_ImVector_ImGuiTabBar)

// Vector_ImGuiTabBarAlloc allocates a continuous block of Vector_ImGuiTabBar.
func Vector_ImGuiTabBarAlloc(alloc ffi.Allocator, count int) Vector_ImGuiTabBar {
	ptr := alloc.Allocate(Vector_ImGuiTabBarSizeOf * count)
	return Vector_ImGuiTabBar(ptr)
}

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p Vector_ImGuiTabBar) Offset(offset int) Vector_ImGuiTabBar {
	return p + Vector_ImGuiTabBar(offset*Vector_ImGuiTabBarSizeOf)
}

// GetCapacity returns the value in Capacity.
func (p Vector_ImGuiTabBar) GetCapacity() int32 {
	ptr := (*C.ImVector_ImGuiTabBar)(unsafe.Pointer(p))
	return int32(ptr.Capacity)
}

// SetCapacity sets the value in Capacity.
func (p Vector_ImGuiTabBar) SetCapacity(value int32) {
	ptr := (*C.ImVector_ImGuiTabBar)(unsafe.Pointer(p))
	ptr.Capacity = (C.int32_t)(value)
}

// GetData returns the value in Data.
func (p Vector_ImGuiTabBar) GetData() GuiTabBar {
	ptr := (*C.ImVector_ImGuiTabBar)(unsafe.Pointer(p))
	return GuiTabBar(unsafe.Pointer(ptr.Data))
}

// SetData sets the value in Data.
func (p Vector_ImGuiTabBar) SetData(value GuiTabBar) {
	ptr := (*C.ImVector_ImGuiTabBar)(unsafe.Pointer(p))
	ptr.Data = (*C.ImGuiTabBar)(unsafe.Pointer(value))
}

// GetSize returns the value in Size.
func (p Vector_ImGuiTabBar) GetSize() int32 {
	ptr := (*C.ImVector_ImGuiTabBar)(unsafe.Pointer(p))
	return int32(ptr.Size)
}

// SetSize sets the value in Size.
func (p Vector_ImGuiTabBar) SetSize(value int32) {
	ptr := (*C.ImVector_ImGuiTabBar)(unsafe.Pointer(p))
	ptr.Size = (C.int32_t)(value)
}

// Vector_ImGuiTabItem wraps struct ImVector_ImGuiTabItem.
type Vector_ImGuiTabItem uintptr

// Vector_ImGuiTabItemNil is a null pointer.
var Vector_ImGuiTabItemNil Vector_ImGuiTabItem

// Vector_ImGuiTabItemSizeOf is the byte size of ImVector_ImGuiTabItem.
const Vector_ImGuiTabItemSizeOf = int(C.sizeof_ImVector_ImGuiTabItem)

// Vector_ImGuiTabItemAlloc allocates a continuous block of Vector_ImGuiTabItem.
func Vector_ImGuiTabItemAlloc(alloc ffi.Allocator, count int) Vector_ImGuiTabItem {
	ptr := alloc.Allocate(Vector_ImGuiTabItemSizeOf * count)
	return Vector_ImGuiTabItem(ptr)
}

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p Vector_ImGuiTabItem) Offset(offset int) Vector_ImGuiTabItem {
	return p + Vector_ImGuiTabItem(offset*Vector_ImGuiTabItemSizeOf)
}

// GetCapacity returns the value in Capacity.
func (p Vector_ImGuiTabItem) GetCapacity() int32 {
	ptr := (*C.ImVector_ImGuiTabItem)(unsafe.Pointer(p))
	return int32(ptr.Capacity)
}

// SetCapacity sets the value in Capacity.
func (p Vector_ImGuiTabItem) SetCapacity(value int32) {
	ptr := (*C.ImVector_ImGuiTabItem)(unsafe.Pointer(p))
	ptr.Capacity = (C.int32_t)(value)
}

// GetData returns the value in Data.
func (p Vector_ImGuiTabItem) GetData() GuiTabItem {
	ptr := (*C.ImVector_ImGuiTabItem)(unsafe.Pointer(p))
	return GuiTabItem(unsafe.Pointer(ptr.Data))
}

// SetData sets the value in Data.
func (p Vector_ImGuiTabItem) SetData(value GuiTabItem) {
	ptr := (*C.ImVector_ImGuiTabItem)(unsafe.Pointer(p))
	ptr.Data = (*C.ImGuiTabItem)(unsafe.Pointer(value))
}

// GetSize returns the value in Size.
func (p Vector_ImGuiTabItem) GetSize() int32 {
	ptr := (*C.ImVector_ImGuiTabItem)(unsafe.Pointer(p))
	return int32(ptr.Size)
}

// SetSize sets the value in Size.
func (p Vector_ImGuiTabItem) SetSize(value int32) {
	ptr := (*C.ImVector_ImGuiTabItem)(unsafe.Pointer(p))
	ptr.Size = (C.int32_t)(value)
}

// Vector_ImGuiTable wraps struct ImVector_ImGuiTable.
type Vector_ImGuiTable uintptr

// Vector_ImGuiTableNil is a null pointer.
var Vector_ImGuiTableNil Vector_ImGuiTable

// Vector_ImGuiTableSizeOf is the byte size of ImVector_ImGuiTable.
const Vector_ImGuiTableSizeOf = int(C.sizeof_ImVector_ImGuiTable)

// Vector_ImGuiTableAlloc allocates a continuous block of Vector_ImGuiTable.
func Vector_ImGuiTableAlloc(alloc ffi.Allocator, count int) Vector_ImGuiTable {
	ptr := alloc.Allocate(Vector_ImGuiTableSizeOf * count)
	return Vector_ImGuiTable(ptr)
}

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p Vector_ImGuiTable) Offset(offset int) Vector_ImGuiTable {
	return p + Vector_ImGuiTable(offset*Vector_ImGuiTableSizeOf)
}

// GetCapacity returns the value in Capacity.
func (p Vector_ImGuiTable) GetCapacity() int32 {
	ptr := (*C.ImVector_ImGuiTable)(unsafe.Pointer(p))
	return int32(ptr.Capacity)
}

// SetCapacity sets the value in Capacity.
func (p Vector_ImGuiTable) SetCapacity(value int32) {
	ptr := (*C.ImVector_ImGuiTable)(unsafe.Pointer(p))
	ptr.Capacity = (C.int32_t)(value)
}

// GetData returns the value in Data.
func (p Vector_ImGuiTable) GetData() GuiTable {
	ptr := (*C.ImVector_ImGuiTable)(unsafe.Pointer(p))
	return GuiTable(unsafe.Pointer(ptr.Data))
}

// SetData sets the value in Data.
func (p Vector_ImGuiTable) SetData(value GuiTable) {
	ptr := (*C.ImVector_ImGuiTable)(unsafe.Pointer(p))
	ptr.Data = (*C.ImGuiTable)(unsafe.Pointer(value))
}

// GetSize returns the value in Size.
func (p Vector_ImGuiTable) GetSize() int32 {
	ptr := (*C.ImVector_ImGuiTable)(unsafe.Pointer(p))
	return int32(ptr.Size)
}

// SetSize sets the value in Size.
func (p Vector_ImGuiTable) SetSize(value int32) {
	ptr := (*C.ImVector_ImGuiTable)(unsafe.Pointer(p))
	ptr.Size = (C.int32_t)(value)
}

// Vector_ImGuiTableColumnSortSpecs wraps struct ImVector_ImGuiTableColumnSortSpecs.
type Vector_ImGuiTableColumnSortSpecs uintptr

// Vector_ImGuiTableColumnSortSpecsNil is a null pointer.
var Vector_ImGuiTableColumnSortSpecsNil Vector_ImGuiTableColumnSortSpecs

// Vector_ImGuiTableColumnSortSpecsSizeOf is the byte size of ImVector_ImGuiTableColumnSortSpecs.
const Vector_ImGuiTableColumnSortSpecsSizeOf = int(C.sizeof_ImVector_ImGuiTableColumnSortSpecs)

// Vector_ImGuiTableColumnSortSpecsAlloc allocates a continuous block of Vector_ImGuiTableColumnSortSpecs.
func Vector_ImGuiTableColumnSortSpecsAlloc(alloc ffi.Allocator, count int) Vector_ImGuiTableColumnSortSpecs {
	ptr := alloc.Allocate(Vector_ImGuiTableColumnSortSpecsSizeOf * count)
	return Vector_ImGuiTableColumnSortSpecs(ptr)
}

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p Vector_ImGuiTableColumnSortSpecs) Offset(offset int) Vector_ImGuiTableColumnSortSpecs {
	return p + Vector_ImGuiTableColumnSortSpecs(offset*Vector_ImGuiTableColumnSortSpecsSizeOf)
}

// GetCapacity returns the value in Capacity.
func (p Vector_ImGuiTableColumnSortSpecs) GetCapacity() int32 {
	ptr := (*C.ImVector_ImGuiTableColumnSortSpecs)(unsafe.Pointer(p))
	return int32(ptr.Capacity)
}

// SetCapacity sets the value in Capacity.
func (p Vector_ImGuiTableColumnSortSpecs) SetCapacity(value int32) {
	ptr := (*C.ImVector_ImGuiTableColumnSortSpecs)(unsafe.Pointer(p))
	ptr.Capacity = (C.int32_t)(value)
}

// GetData returns the value in Data.
func (p Vector_ImGuiTableColumnSortSpecs) GetData() GuiTableColumnSortSpecs {
	ptr := (*C.ImVector_ImGuiTableColumnSortSpecs)(unsafe.Pointer(p))
	return GuiTableColumnSortSpecs(unsafe.Pointer(ptr.Data))
}

// SetData sets the value in Data.
func (p Vector_ImGuiTableColumnSortSpecs) SetData(value GuiTableColumnSortSpecs) {
	ptr := (*C.ImVector_ImGuiTableColumnSortSpecs)(unsafe.Pointer(p))
	ptr.Data = (*C.ImGuiTableColumnSortSpecs)(unsafe.Pointer(value))
}

// GetSize returns the value in Size.
func (p Vector_ImGuiTableColumnSortSpecs) GetSize() int32 {
	ptr := (*C.ImVector_ImGuiTableColumnSortSpecs)(unsafe.Pointer(p))
	return int32(ptr.Size)
}

// SetSize sets the value in Size.
func (p Vector_ImGuiTableColumnSortSpecs) SetSize(value int32) {
	ptr := (*C.ImVector_ImGuiTableColumnSortSpecs)(unsafe.Pointer(p))
	ptr.Size = (C.int32_t)(value)
}

// Vector_ImGuiTableHeaderData wraps struct ImVector_ImGuiTableHeaderData.
type Vector_ImGuiTableHeaderData uintptr

// Vector_ImGuiTableHeaderDataNil is a null pointer.
var Vector_ImGuiTableHeaderDataNil Vector_ImGuiTableHeaderData

// Vector_ImGuiTableHeaderDataSizeOf is the byte size of ImVector_ImGuiTableHeaderData.
const Vector_ImGuiTableHeaderDataSizeOf = int(C.sizeof_ImVector_ImGuiTableHeaderData)

// Vector_ImGuiTableHeaderDataAlloc allocates a continuous block of Vector_ImGuiTableHeaderData.
func Vector_ImGuiTableHeaderDataAlloc(alloc ffi.Allocator, count int) Vector_ImGuiTableHeaderData {
	ptr := alloc.Allocate(Vector_ImGuiTableHeaderDataSizeOf * count)
	return Vector_ImGuiTableHeaderData(ptr)
}

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p Vector_ImGuiTableHeaderData) Offset(offset int) Vector_ImGuiTableHeaderData {
	return p + Vector_ImGuiTableHeaderData(offset*Vector_ImGuiTableHeaderDataSizeOf)
}

// GetCapacity returns the value in Capacity.
func (p Vector_ImGuiTableHeaderData) GetCapacity() int32 {
	ptr := (*C.ImVector_ImGuiTableHeaderData)(unsafe.Pointer(p))
	return int32(ptr.Capacity)
}

// SetCapacity sets the value in Capacity.
func (p Vector_ImGuiTableHeaderData) SetCapacity(value int32) {
	ptr := (*C.ImVector_ImGuiTableHeaderData)(unsafe.Pointer(p))
	ptr.Capacity = (C.int32_t)(value)
}

// GetData returns the value in Data.
func (p Vector_ImGuiTableHeaderData) GetData() GuiTableHeaderData {
	ptr := (*C.ImVector_ImGuiTableHeaderData)(unsafe.Pointer(p))
	return GuiTableHeaderData(unsafe.Pointer(ptr.Data))
}

// SetData sets the value in Data.
func (p Vector_ImGuiTableHeaderData) SetData(value GuiTableHeaderData) {
	ptr := (*C.ImVector_ImGuiTableHeaderData)(unsafe.Pointer(p))
	ptr.Data = (*C.ImGuiTableHeaderData)(unsafe.Pointer(value))
}

// GetSize returns the value in Size.
func (p Vector_ImGuiTableHeaderData) GetSize() int32 {
	ptr := (*C.ImVector_ImGuiTableHeaderData)(unsafe.Pointer(p))
	return int32(ptr.Size)
}

// SetSize sets the value in Size.
func (p Vector_ImGuiTableHeaderData) SetSize(value int32) {
	ptr := (*C.ImVector_ImGuiTableHeaderData)(unsafe.Pointer(p))
	ptr.Size = (C.int32_t)(value)
}

// Vector_ImGuiTableInstanceData wraps struct ImVector_ImGuiTableInstanceData.
type Vector_ImGuiTableInstanceData uintptr

// Vector_ImGuiTableInstanceDataNil is a null pointer.
var Vector_ImGuiTableInstanceDataNil Vector_ImGuiTableInstanceData

// Vector_ImGuiTableInstanceDataSizeOf is the byte size of ImVector_ImGuiTableInstanceData.
const Vector_ImGuiTableInstanceDataSizeOf = int(C.sizeof_ImVector_ImGuiTableInstanceData)

// Vector_ImGuiTableInstanceDataAlloc allocates a continuous block of Vector_ImGuiTableInstanceData.
func Vector_ImGuiTableInstanceDataAlloc(alloc ffi.Allocator, count int) Vector_ImGuiTableInstanceData {
	ptr := alloc.Allocate(Vector_ImGuiTableInstanceDataSizeOf * count)
	return Vector_ImGuiTableInstanceData(ptr)
}

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p Vector_ImGuiTableInstanceData) Offset(offset int) Vector_ImGuiTableInstanceData {
	return p + Vector_ImGuiTableInstanceData(offset*Vector_ImGuiTableInstanceDataSizeOf)
}

// GetCapacity returns the value in Capacity.
func (p Vector_ImGuiTableInstanceData) GetCapacity() int32 {
	ptr := (*C.ImVector_ImGuiTableInstanceData)(unsafe.Pointer(p))
	return int32(ptr.Capacity)
}

// SetCapacity sets the value in Capacity.
func (p Vector_ImGuiTableInstanceData) SetCapacity(value int32) {
	ptr := (*C.ImVector_ImGuiTableInstanceData)(unsafe.Pointer(p))
	ptr.Capacity = (C.int32_t)(value)
}

// GetData returns the value in Data.
func (p Vector_ImGuiTableInstanceData) GetData() GuiTableInstanceData {
	ptr := (*C.ImVector_ImGuiTableInstanceData)(unsafe.Pointer(p))
	return GuiTableInstanceData(unsafe.Pointer(ptr.Data))
}

// SetData sets the value in Data.
func (p Vector_ImGuiTableInstanceData) SetData(value GuiTableInstanceData) {
	ptr := (*C.ImVector_ImGuiTableInstanceData)(unsafe.Pointer(p))
	ptr.Data = (*C.ImGuiTableInstanceData)(unsafe.Pointer(value))
}

// GetSize returns the value in Size.
func (p Vector_ImGuiTableInstanceData) GetSize() int32 {
	ptr := (*C.ImVector_ImGuiTableInstanceData)(unsafe.Pointer(p))
	return int32(ptr.Size)
}

// SetSize sets the value in Size.
func (p Vector_ImGuiTableInstanceData) SetSize(value int32) {
	ptr := (*C.ImVector_ImGuiTableInstanceData)(unsafe.Pointer(p))
	ptr.Size = (C.int32_t)(value)
}

// Vector_ImGuiTableTempData wraps struct ImVector_ImGuiTableTempData.
type Vector_ImGuiTableTempData uintptr

// Vector_ImGuiTableTempDataNil is a null pointer.
var Vector_ImGuiTableTempDataNil Vector_ImGuiTableTempData

// Vector_ImGuiTableTempDataSizeOf is the byte size of ImVector_ImGuiTableTempData.
const Vector_ImGuiTableTempDataSizeOf = int(C.sizeof_ImVector_ImGuiTableTempData)

// Vector_ImGuiTableTempDataAlloc allocates a continuous block of Vector_ImGuiTableTempData.
func Vector_ImGuiTableTempDataAlloc(alloc ffi.Allocator, count int) Vector_ImGuiTableTempData {
	ptr := alloc.Allocate(Vector_ImGuiTableTempDataSizeOf * count)
	return Vector_ImGuiTableTempData(ptr)
}

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p Vector_ImGuiTableTempData) Offset(offset int) Vector_ImGuiTableTempData {
	return p + Vector_ImGuiTableTempData(offset*Vector_ImGuiTableTempDataSizeOf)
}

// GetCapacity returns the value in Capacity.
func (p Vector_ImGuiTableTempData) GetCapacity() int32 {
	ptr := (*C.ImVector_ImGuiTableTempData)(unsafe.Pointer(p))
	return int32(ptr.Capacity)
}

// SetCapacity sets the value in Capacity.
func (p Vector_ImGuiTableTempData) SetCapacity(value int32) {
	ptr := (*C.ImVector_ImGuiTableTempData)(unsafe.Pointer(p))
	ptr.Capacity = (C.int32_t)(value)
}

// GetData returns the value in Data.
func (p Vector_ImGuiTableTempData) GetData() GuiTableTempData {
	ptr := (*C.ImVector_ImGuiTableTempData)(unsafe.Pointer(p))
	return GuiTableTempData(unsafe.Pointer(ptr.Data))
}

// SetData sets the value in Data.
func (p Vector_ImGuiTableTempData) SetData(value GuiTableTempData) {
	ptr := (*C.ImVector_ImGuiTableTempData)(unsafe.Pointer(p))
	ptr.Data = (*C.ImGuiTableTempData)(unsafe.Pointer(value))
}

// GetSize returns the value in Size.
func (p Vector_ImGuiTableTempData) GetSize() int32 {
	ptr := (*C.ImVector_ImGuiTableTempData)(unsafe.Pointer(p))
	return int32(ptr.Size)
}

// SetSize sets the value in Size.
func (p Vector_ImGuiTableTempData) SetSize(value int32) {
	ptr := (*C.ImVector_ImGuiTableTempData)(unsafe.Pointer(p))
	ptr.Size = (C.int32_t)(value)
}

// Vector_ImGuiTextRange wraps struct ImVector_ImGuiTextRange.
type Vector_ImGuiTextRange uintptr

// Vector_ImGuiTextRangeNil is a null pointer.
var Vector_ImGuiTextRangeNil Vector_ImGuiTextRange

// Vector_ImGuiTextRangeSizeOf is the byte size of ImVector_ImGuiTextRange.
const Vector_ImGuiTextRangeSizeOf = int(C.sizeof_ImVector_ImGuiTextRange)

// Vector_ImGuiTextRangeAlloc allocates a continuous block of Vector_ImGuiTextRange.
func Vector_ImGuiTextRangeAlloc(alloc ffi.Allocator, count int) Vector_ImGuiTextRange {
	ptr := alloc.Allocate(Vector_ImGuiTextRangeSizeOf * count)
	return Vector_ImGuiTextRange(ptr)
}

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p Vector_ImGuiTextRange) Offset(offset int) Vector_ImGuiTextRange {
	return p + Vector_ImGuiTextRange(offset*Vector_ImGuiTextRangeSizeOf)
}

// GetCapacity returns the value in Capacity.
func (p Vector_ImGuiTextRange) GetCapacity() int32 {
	ptr := (*C.ImVector_ImGuiTextRange)(unsafe.Pointer(p))
	return int32(ptr.Capacity)
}

// SetCapacity sets the value in Capacity.
func (p Vector_ImGuiTextRange) SetCapacity(value int32) {
	ptr := (*C.ImVector_ImGuiTextRange)(unsafe.Pointer(p))
	ptr.Capacity = (C.int32_t)(value)
}

// GetData returns the value in Data.
func (p Vector_ImGuiTextRange) GetData() GuiTextRange {
	ptr := (*C.ImVector_ImGuiTextRange)(unsafe.Pointer(p))
	return GuiTextRange(unsafe.Pointer(ptr.Data))
}

// SetData sets the value in Data.
func (p Vector_ImGuiTextRange) SetData(value GuiTextRange) {
	ptr := (*C.ImVector_ImGuiTextRange)(unsafe.Pointer(p))
	ptr.Data = (*C.ImGuiTextRange)(unsafe.Pointer(value))
}

// GetSize returns the value in Size.
func (p Vector_ImGuiTextRange) GetSize() int32 {
	ptr := (*C.ImVector_ImGuiTextRange)(unsafe.Pointer(p))
	return int32(ptr.Size)
}

// SetSize sets the value in Size.
func (p Vector_ImGuiTextRange) SetSize(value int32) {
	ptr := (*C.ImVector_ImGuiTextRange)(unsafe.Pointer(p))
	ptr.Size = (C.int32_t)(value)
}

// Vector_ImGuiTreeNodeStackData wraps struct ImVector_ImGuiTreeNodeStackData.
type Vector_ImGuiTreeNodeStackData uintptr

// Vector_ImGuiTreeNodeStackDataNil is a null pointer.
var Vector_ImGuiTreeNodeStackDataNil Vector_ImGuiTreeNodeStackData

// Vector_ImGuiTreeNodeStackDataSizeOf is the byte size of ImVector_ImGuiTreeNodeStackData.
const Vector_ImGuiTreeNodeStackDataSizeOf = int(C.sizeof_ImVector_ImGuiTreeNodeStackData)

// Vector_ImGuiTreeNodeStackDataAlloc allocates a continuous block of Vector_ImGuiTreeNodeStackData.
func Vector_ImGuiTreeNodeStackDataAlloc(alloc ffi.Allocator, count int) Vector_ImGuiTreeNodeStackData {
	ptr := alloc.Allocate(Vector_ImGuiTreeNodeStackDataSizeOf * count)
	return Vector_ImGuiTreeNodeStackData(ptr)
}

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p Vector_ImGuiTreeNodeStackData) Offset(offset int) Vector_ImGuiTreeNodeStackData {
	return p + Vector_ImGuiTreeNodeStackData(offset*Vector_ImGuiTreeNodeStackDataSizeOf)
}

// GetCapacity returns the value in Capacity.
func (p Vector_ImGuiTreeNodeStackData) GetCapacity() int32 {
	ptr := (*C.ImVector_ImGuiTreeNodeStackData)(unsafe.Pointer(p))
	return int32(ptr.Capacity)
}

// SetCapacity sets the value in Capacity.
func (p Vector_ImGuiTreeNodeStackData) SetCapacity(value int32) {
	ptr := (*C.ImVector_ImGuiTreeNodeStackData)(unsafe.Pointer(p))
	ptr.Capacity = (C.int32_t)(value)
}

// GetData returns the value in Data.
func (p Vector_ImGuiTreeNodeStackData) GetData() GuiTreeNodeStackData {
	ptr := (*C.ImVector_ImGuiTreeNodeStackData)(unsafe.Pointer(p))
	return GuiTreeNodeStackData(unsafe.Pointer(ptr.Data))
}

// SetData sets the value in Data.
func (p Vector_ImGuiTreeNodeStackData) SetData(value GuiTreeNodeStackData) {
	ptr := (*C.ImVector_ImGuiTreeNodeStackData)(unsafe.Pointer(p))
	ptr.Data = (*C.ImGuiTreeNodeStackData)(unsafe.Pointer(value))
}

// GetSize returns the value in Size.
func (p Vector_ImGuiTreeNodeStackData) GetSize() int32 {
	ptr := (*C.ImVector_ImGuiTreeNodeStackData)(unsafe.Pointer(p))
	return int32(ptr.Size)
}

// SetSize sets the value in Size.
func (p Vector_ImGuiTreeNodeStackData) SetSize(value int32) {
	ptr := (*C.ImVector_ImGuiTreeNodeStackData)(unsafe.Pointer(p))
	ptr.Size = (C.int32_t)(value)
}

// Vector_ImGuiViewportPPtr wraps struct ImVector_ImGuiViewportPPtr.
type Vector_ImGuiViewportPPtr uintptr

// Vector_ImGuiViewportPPtrNil is a null pointer.
var Vector_ImGuiViewportPPtrNil Vector_ImGuiViewportPPtr

// Vector_ImGuiViewportPPtrSizeOf is the byte size of ImVector_ImGuiViewportPPtr.
const Vector_ImGuiViewportPPtrSizeOf = int(C.sizeof_ImVector_ImGuiViewportPPtr)

// Vector_ImGuiViewportPPtrAlloc allocates a continuous block of Vector_ImGuiViewportPPtr.
func Vector_ImGuiViewportPPtrAlloc(alloc ffi.Allocator, count int) Vector_ImGuiViewportPPtr {
	ptr := alloc.Allocate(Vector_ImGuiViewportPPtrSizeOf * count)
	return Vector_ImGuiViewportPPtr(ptr)
}

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p Vector_ImGuiViewportPPtr) Offset(offset int) Vector_ImGuiViewportPPtr {
	return p + Vector_ImGuiViewportPPtr(offset*Vector_ImGuiViewportPPtrSizeOf)
}

// GetCapacity returns the value in Capacity.
func (p Vector_ImGuiViewportPPtr) GetCapacity() int32 {
	ptr := (*C.ImVector_ImGuiViewportPPtr)(unsafe.Pointer(p))
	return int32(ptr.Capacity)
}

// SetCapacity sets the value in Capacity.
func (p Vector_ImGuiViewportPPtr) SetCapacity(value int32) {
	ptr := (*C.ImVector_ImGuiViewportPPtr)(unsafe.Pointer(p))
	ptr.Capacity = (C.int32_t)(value)
}

// GetData returns the value in Data.
func (p Vector_ImGuiViewportPPtr) GetData() ffi.Ref[GuiViewportP] {
	ptr := (*C.ImVector_ImGuiViewportPPtr)(unsafe.Pointer(p))
	return ffi.RefFromPtr[GuiViewportP](unsafe.Pointer(ptr.Data))
}

// SetData sets the value in Data.
func (p Vector_ImGuiViewportPPtr) SetData(value ffi.Ref[GuiViewportP]) {
	ptr := (*C.ImVector_ImGuiViewportPPtr)(unsafe.Pointer(p))
	ptr.Data = (**C.ImGuiViewportP)(value.Raw())
}

// GetSize returns the value in Size.
func (p Vector_ImGuiViewportPPtr) GetSize() int32 {
	ptr := (*C.ImVector_ImGuiViewportPPtr)(unsafe.Pointer(p))
	return int32(ptr.Size)
}

// SetSize sets the value in Size.
func (p Vector_ImGuiViewportPPtr) SetSize(value int32) {
	ptr := (*C.ImVector_ImGuiViewportPPtr)(unsafe.Pointer(p))
	ptr.Size = (C.int32_t)(value)
}

// Vector_ImGuiViewportPtr wraps struct ImVector_ImGuiViewportPtr.
type Vector_ImGuiViewportPtr uintptr

// Vector_ImGuiViewportPtrNil is a null pointer.
var Vector_ImGuiViewportPtrNil Vector_ImGuiViewportPtr

// Vector_ImGuiViewportPtrSizeOf is the byte size of ImVector_ImGuiViewportPtr.
const Vector_ImGuiViewportPtrSizeOf = int(C.sizeof_ImVector_ImGuiViewportPtr)

// Vector_ImGuiViewportPtrAlloc allocates a continuous block of Vector_ImGuiViewportPtr.
func Vector_ImGuiViewportPtrAlloc(alloc ffi.Allocator, count int) Vector_ImGuiViewportPtr {
	ptr := alloc.Allocate(Vector_ImGuiViewportPtrSizeOf * count)
	return Vector_ImGuiViewportPtr(ptr)
}

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p Vector_ImGuiViewportPtr) Offset(offset int) Vector_ImGuiViewportPtr {
	return p + Vector_ImGuiViewportPtr(offset*Vector_ImGuiViewportPtrSizeOf)
}

// GetCapacity returns the value in Capacity.
func (p Vector_ImGuiViewportPtr) GetCapacity() int32 {
	ptr := (*C.ImVector_ImGuiViewportPtr)(unsafe.Pointer(p))
	return int32(ptr.Capacity)
}

// SetCapacity sets the value in Capacity.
func (p Vector_ImGuiViewportPtr) SetCapacity(value int32) {
	ptr := (*C.ImVector_ImGuiViewportPtr)(unsafe.Pointer(p))
	ptr.Capacity = (C.int32_t)(value)
}

// GetData returns the value in Data.
func (p Vector_ImGuiViewportPtr) GetData() ffi.Ref[GuiViewport] {
	ptr := (*C.ImVector_ImGuiViewportPtr)(unsafe.Pointer(p))
	return ffi.RefFromPtr[GuiViewport](unsafe.Pointer(ptr.Data))
}

// SetData sets the value in Data.
func (p Vector_ImGuiViewportPtr) SetData(value ffi.Ref[GuiViewport]) {
	ptr := (*C.ImVector_ImGuiViewportPtr)(unsafe.Pointer(p))
	ptr.Data = (**C.ImGuiViewport)(value.Raw())
}

// GetSize returns the value in Size.
func (p Vector_ImGuiViewportPtr) GetSize() int32 {
	ptr := (*C.ImVector_ImGuiViewportPtr)(unsafe.Pointer(p))
	return int32(ptr.Size)
}

// SetSize sets the value in Size.
func (p Vector_ImGuiViewportPtr) SetSize(value int32) {
	ptr := (*C.ImVector_ImGuiViewportPtr)(unsafe.Pointer(p))
	ptr.Size = (C.int32_t)(value)
}

// Vector_ImGuiWindowPtr wraps struct ImVector_ImGuiWindowPtr.
type Vector_ImGuiWindowPtr uintptr

// Vector_ImGuiWindowPtrNil is a null pointer.
var Vector_ImGuiWindowPtrNil Vector_ImGuiWindowPtr

// Vector_ImGuiWindowPtrSizeOf is the byte size of ImVector_ImGuiWindowPtr.
const Vector_ImGuiWindowPtrSizeOf = int(C.sizeof_ImVector_ImGuiWindowPtr)

// Vector_ImGuiWindowPtrAlloc allocates a continuous block of Vector_ImGuiWindowPtr.
func Vector_ImGuiWindowPtrAlloc(alloc ffi.Allocator, count int) Vector_ImGuiWindowPtr {
	ptr := alloc.Allocate(Vector_ImGuiWindowPtrSizeOf * count)
	return Vector_ImGuiWindowPtr(ptr)
}

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p Vector_ImGuiWindowPtr) Offset(offset int) Vector_ImGuiWindowPtr {
	return p + Vector_ImGuiWindowPtr(offset*Vector_ImGuiWindowPtrSizeOf)
}

// GetCapacity returns the value in Capacity.
func (p Vector_ImGuiWindowPtr) GetCapacity() int32 {
	ptr := (*C.ImVector_ImGuiWindowPtr)(unsafe.Pointer(p))
	return int32(ptr.Capacity)
}

// SetCapacity sets the value in Capacity.
func (p Vector_ImGuiWindowPtr) SetCapacity(value int32) {
	ptr := (*C.ImVector_ImGuiWindowPtr)(unsafe.Pointer(p))
	ptr.Capacity = (C.int32_t)(value)
}

// GetData returns the value in Data.
func (p Vector_ImGuiWindowPtr) GetData() ffi.Ref[GuiWindow] {
	ptr := (*C.ImVector_ImGuiWindowPtr)(unsafe.Pointer(p))
	return ffi.RefFromPtr[GuiWindow](unsafe.Pointer(ptr.Data))
}

// SetData sets the value in Data.
func (p Vector_ImGuiWindowPtr) SetData(value ffi.Ref[GuiWindow]) {
	ptr := (*C.ImVector_ImGuiWindowPtr)(unsafe.Pointer(p))
	ptr.Data = (**C.ImGuiWindow)(value.Raw())
}

// GetSize returns the value in Size.
func (p Vector_ImGuiWindowPtr) GetSize() int32 {
	ptr := (*C.ImVector_ImGuiWindowPtr)(unsafe.Pointer(p))
	return int32(ptr.Size)
}

// SetSize sets the value in Size.
func (p Vector_ImGuiWindowPtr) SetSize(value int32) {
	ptr := (*C.ImVector_ImGuiWindowPtr)(unsafe.Pointer(p))
	ptr.Size = (C.int32_t)(value)
}

// Vector_ImGuiWindowStackData wraps struct ImVector_ImGuiWindowStackData.
type Vector_ImGuiWindowStackData uintptr

// Vector_ImGuiWindowStackDataNil is a null pointer.
var Vector_ImGuiWindowStackDataNil Vector_ImGuiWindowStackData

// Vector_ImGuiWindowStackDataSizeOf is the byte size of ImVector_ImGuiWindowStackData.
const Vector_ImGuiWindowStackDataSizeOf = int(C.sizeof_ImVector_ImGuiWindowStackData)

// Vector_ImGuiWindowStackDataAlloc allocates a continuous block of Vector_ImGuiWindowStackData.
func Vector_ImGuiWindowStackDataAlloc(alloc ffi.Allocator, count int) Vector_ImGuiWindowStackData {
	ptr := alloc.Allocate(Vector_ImGuiWindowStackDataSizeOf * count)
	return Vector_ImGuiWindowStackData(ptr)
}

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p Vector_ImGuiWindowStackData) Offset(offset int) Vector_ImGuiWindowStackData {
	return p + Vector_ImGuiWindowStackData(offset*Vector_ImGuiWindowStackDataSizeOf)
}

// GetCapacity returns the value in Capacity.
func (p Vector_ImGuiWindowStackData) GetCapacity() int32 {
	ptr := (*C.ImVector_ImGuiWindowStackData)(unsafe.Pointer(p))
	return int32(ptr.Capacity)
}

// SetCapacity sets the value in Capacity.
func (p Vector_ImGuiWindowStackData) SetCapacity(value int32) {
	ptr := (*C.ImVector_ImGuiWindowStackData)(unsafe.Pointer(p))
	ptr.Capacity = (C.int32_t)(value)
}

// GetData returns the value in Data.
func (p Vector_ImGuiWindowStackData) GetData() GuiWindowStackData {
	ptr := (*C.ImVector_ImGuiWindowStackData)(unsafe.Pointer(p))
	return GuiWindowStackData(unsafe.Pointer(ptr.Data))
}

// SetData sets the value in Data.
func (p Vector_ImGuiWindowStackData) SetData(value GuiWindowStackData) {
	ptr := (*C.ImVector_ImGuiWindowStackData)(unsafe.Pointer(p))
	ptr.Data = (*C.ImGuiWindowStackData)(unsafe.Pointer(value))
}

// GetSize returns the value in Size.
func (p Vector_ImGuiWindowStackData) GetSize() int32 {
	ptr := (*C.ImVector_ImGuiWindowStackData)(unsafe.Pointer(p))
	return int32(ptr.Size)
}

// SetSize sets the value in Size.
func (p Vector_ImGuiWindowStackData) SetSize(value int32) {
	ptr := (*C.ImVector_ImGuiWindowStackData)(unsafe.Pointer(p))
	ptr.Size = (C.int32_t)(value)
}

// Vector_ImTextureDataPtr wraps struct ImVector_ImTextureDataPtr.
type Vector_ImTextureDataPtr uintptr

// Vector_ImTextureDataPtrNil is a null pointer.
var Vector_ImTextureDataPtrNil Vector_ImTextureDataPtr

// Vector_ImTextureDataPtrSizeOf is the byte size of ImVector_ImTextureDataPtr.
const Vector_ImTextureDataPtrSizeOf = int(C.sizeof_ImVector_ImTextureDataPtr)

// Vector_ImTextureDataPtrAlloc allocates a continuous block of Vector_ImTextureDataPtr.
func Vector_ImTextureDataPtrAlloc(alloc ffi.Allocator, count int) Vector_ImTextureDataPtr {
	ptr := alloc.Allocate(Vector_ImTextureDataPtrSizeOf * count)
	return Vector_ImTextureDataPtr(ptr)
}

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p Vector_ImTextureDataPtr) Offset(offset int) Vector_ImTextureDataPtr {
	return p + Vector_ImTextureDataPtr(offset*Vector_ImTextureDataPtrSizeOf)
}

// GetCapacity returns the value in Capacity.
func (p Vector_ImTextureDataPtr) GetCapacity() int32 {
	ptr := (*C.ImVector_ImTextureDataPtr)(unsafe.Pointer(p))
	return int32(ptr.Capacity)
}

// SetCapacity sets the value in Capacity.
func (p Vector_ImTextureDataPtr) SetCapacity(value int32) {
	ptr := (*C.ImVector_ImTextureDataPtr)(unsafe.Pointer(p))
	ptr.Capacity = (C.int32_t)(value)
}

// GetData returns the value in Data.
func (p Vector_ImTextureDataPtr) GetData() ffi.Ref[TextureData] {
	ptr := (*C.ImVector_ImTextureDataPtr)(unsafe.Pointer(p))
	return ffi.RefFromPtr[TextureData](unsafe.Pointer(ptr.Data))
}

// SetData sets the value in Data.
func (p Vector_ImTextureDataPtr) SetData(value ffi.Ref[TextureData]) {
	ptr := (*C.ImVector_ImTextureDataPtr)(unsafe.Pointer(p))
	ptr.Data = (**C.ImTextureData)(value.Raw())
}

// GetSize returns the value in Size.
func (p Vector_ImTextureDataPtr) GetSize() int32 {
	ptr := (*C.ImVector_ImTextureDataPtr)(unsafe.Pointer(p))
	return int32(ptr.Size)
}

// SetSize sets the value in Size.
func (p Vector_ImTextureDataPtr) SetSize(value int32) {
	ptr := (*C.ImVector_ImTextureDataPtr)(unsafe.Pointer(p))
	ptr.Size = (C.int32_t)(value)
}

// Vector_ImTextureRect wraps struct ImVector_ImTextureRect.
type Vector_ImTextureRect uintptr

// Vector_ImTextureRectNil is a null pointer.
var Vector_ImTextureRectNil Vector_ImTextureRect

// Vector_ImTextureRectSizeOf is the byte size of ImVector_ImTextureRect.
const Vector_ImTextureRectSizeOf = int(C.sizeof_ImVector_ImTextureRect)

// Vector_ImTextureRectAlloc allocates a continuous block of Vector_ImTextureRect.
func Vector_ImTextureRectAlloc(alloc ffi.Allocator, count int) Vector_ImTextureRect {
	ptr := alloc.Allocate(Vector_ImTextureRectSizeOf * count)
	return Vector_ImTextureRect(ptr)
}

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p Vector_ImTextureRect) Offset(offset int) Vector_ImTextureRect {
	return p + Vector_ImTextureRect(offset*Vector_ImTextureRectSizeOf)
}

// GetCapacity returns the value in Capacity.
func (p Vector_ImTextureRect) GetCapacity() int32 {
	ptr := (*C.ImVector_ImTextureRect)(unsafe.Pointer(p))
	return int32(ptr.Capacity)
}

// SetCapacity sets the value in Capacity.
func (p Vector_ImTextureRect) SetCapacity(value int32) {
	ptr := (*C.ImVector_ImTextureRect)(unsafe.Pointer(p))
	ptr.Capacity = (C.int32_t)(value)
}

// GetData returns the value in Data.
func (p Vector_ImTextureRect) GetData() TextureRect {
	ptr := (*C.ImVector_ImTextureRect)(unsafe.Pointer(p))
	return TextureRect(unsafe.Pointer(ptr.Data))
}

// SetData sets the value in Data.
func (p Vector_ImTextureRect) SetData(value TextureRect) {
	ptr := (*C.ImVector_ImTextureRect)(unsafe.Pointer(p))
	ptr.Data = (*C.ImTextureRect)(unsafe.Pointer(value))
}

// GetSize returns the value in Size.
func (p Vector_ImTextureRect) GetSize() int32 {
	ptr := (*C.ImVector_ImTextureRect)(unsafe.Pointer(p))
	return int32(ptr.Size)
}

// SetSize sets the value in Size.
func (p Vector_ImTextureRect) SetSize(value int32) {
	ptr := (*C.ImVector_ImTextureRect)(unsafe.Pointer(p))
	ptr.Size = (C.int32_t)(value)
}

// Vector_ImTextureRef wraps struct ImVector_ImTextureRef.
type Vector_ImTextureRef uintptr

// Vector_ImTextureRefNil is a null pointer.
var Vector_ImTextureRefNil Vector_ImTextureRef

// Vector_ImTextureRefSizeOf is the byte size of ImVector_ImTextureRef.
const Vector_ImTextureRefSizeOf = int(C.sizeof_ImVector_ImTextureRef)

// Vector_ImTextureRefAlloc allocates a continuous block of Vector_ImTextureRef.
func Vector_ImTextureRefAlloc(alloc ffi.Allocator, count int) Vector_ImTextureRef {
	ptr := alloc.Allocate(Vector_ImTextureRefSizeOf * count)
	return Vector_ImTextureRef(ptr)
}

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p Vector_ImTextureRef) Offset(offset int) Vector_ImTextureRef {
	return p + Vector_ImTextureRef(offset*Vector_ImTextureRefSizeOf)
}

// GetCapacity returns the value in Capacity.
func (p Vector_ImTextureRef) GetCapacity() int32 {
	ptr := (*C.ImVector_ImTextureRef)(unsafe.Pointer(p))
	return int32(ptr.Capacity)
}

// SetCapacity sets the value in Capacity.
func (p Vector_ImTextureRef) SetCapacity(value int32) {
	ptr := (*C.ImVector_ImTextureRef)(unsafe.Pointer(p))
	ptr.Capacity = (C.int32_t)(value)
}

// GetData returns the value in Data.
func (p Vector_ImTextureRef) GetData() TextureRef {
	ptr := (*C.ImVector_ImTextureRef)(unsafe.Pointer(p))
	return TextureRef(unsafe.Pointer(ptr.Data))
}

// SetData sets the value in Data.
func (p Vector_ImTextureRef) SetData(value TextureRef) {
	ptr := (*C.ImVector_ImTextureRef)(unsafe.Pointer(p))
	ptr.Data = (*C.ImTextureRef)(unsafe.Pointer(value))
}

// GetSize returns the value in Size.
func (p Vector_ImTextureRef) GetSize() int32 {
	ptr := (*C.ImVector_ImTextureRef)(unsafe.Pointer(p))
	return int32(ptr.Size)
}

// SetSize sets the value in Size.
func (p Vector_ImTextureRef) SetSize(value int32) {
	ptr := (*C.ImVector_ImTextureRef)(unsafe.Pointer(p))
	ptr.Size = (C.int32_t)(value)
}

// Vector_ImU16 wraps struct ImVector_ImU16.
type Vector_ImU16 uintptr

// Vector_ImU16Nil is a null pointer.
var Vector_ImU16Nil Vector_ImU16

// Vector_ImU16SizeOf is the byte size of ImVector_ImU16.
const Vector_ImU16SizeOf = int(C.sizeof_ImVector_ImU16)

// Vector_ImU16Alloc allocates a continuous block of Vector_ImU16.
func Vector_ImU16Alloc(alloc ffi.Allocator, count int) Vector_ImU16 {
	ptr := alloc.Allocate(Vector_ImU16SizeOf * count)
	return Vector_ImU16(ptr)
}

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p Vector_ImU16) Offset(offset int) Vector_ImU16 {
	return p + Vector_ImU16(offset*Vector_ImU16SizeOf)
}

// GetCapacity returns the value in Capacity.
func (p Vector_ImU16) GetCapacity() int32 {
	ptr := (*C.ImVector_ImU16)(unsafe.Pointer(p))
	return int32(ptr.Capacity)
}

// SetCapacity sets the value in Capacity.
func (p Vector_ImU16) SetCapacity(value int32) {
	ptr := (*C.ImVector_ImU16)(unsafe.Pointer(p))
	ptr.Capacity = (C.int32_t)(value)
}

// GetData returns the value in Data.
func (p Vector_ImU16) GetData() ffi.Ref[uint16] {
	ptr := (*C.ImVector_ImU16)(unsafe.Pointer(p))
	return ffi.RefFromPtr[uint16](unsafe.Pointer(ptr.Data))
}

// SetData sets the value in Data.
func (p Vector_ImU16) SetData(value ffi.Ref[uint16]) {
	ptr := (*C.ImVector_ImU16)(unsafe.Pointer(p))
	ptr.Data = (*C.uint16_t)(value.Raw())
}

// GetSize returns the value in Size.
func (p Vector_ImU16) GetSize() int32 {
	ptr := (*C.ImVector_ImU16)(unsafe.Pointer(p))
	return int32(ptr.Size)
}

// SetSize sets the value in Size.
func (p Vector_ImU16) SetSize(value int32) {
	ptr := (*C.ImVector_ImU16)(unsafe.Pointer(p))
	ptr.Size = (C.int32_t)(value)
}

// Vector_ImU32 wraps struct ImVector_ImU32.
type Vector_ImU32 uintptr

// Vector_ImU32Nil is a null pointer.
var Vector_ImU32Nil Vector_ImU32

// Vector_ImU32SizeOf is the byte size of ImVector_ImU32.
const Vector_ImU32SizeOf = int(C.sizeof_ImVector_ImU32)

// Vector_ImU32Alloc allocates a continuous block of Vector_ImU32.
func Vector_ImU32Alloc(alloc ffi.Allocator, count int) Vector_ImU32 {
	ptr := alloc.Allocate(Vector_ImU32SizeOf * count)
	return Vector_ImU32(ptr)
}

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p Vector_ImU32) Offset(offset int) Vector_ImU32 {
	return p + Vector_ImU32(offset*Vector_ImU32SizeOf)
}

// GetCapacity returns the value in Capacity.
func (p Vector_ImU32) GetCapacity() int32 {
	ptr := (*C.ImVector_ImU32)(unsafe.Pointer(p))
	return int32(ptr.Capacity)
}

// SetCapacity sets the value in Capacity.
func (p Vector_ImU32) SetCapacity(value int32) {
	ptr := (*C.ImVector_ImU32)(unsafe.Pointer(p))
	ptr.Capacity = (C.int32_t)(value)
}

// GetData returns the value in Data.
func (p Vector_ImU32) GetData() ffi.Ref[uint32] {
	ptr := (*C.ImVector_ImU32)(unsafe.Pointer(p))
	return ffi.RefFromPtr[uint32](unsafe.Pointer(ptr.Data))
}

// SetData sets the value in Data.
func (p Vector_ImU32) SetData(value ffi.Ref[uint32]) {
	ptr := (*C.ImVector_ImU32)(unsafe.Pointer(p))
	ptr.Data = (*C.uint32_t)(value.Raw())
}

// GetSize returns the value in Size.
func (p Vector_ImU32) GetSize() int32 {
	ptr := (*C.ImVector_ImU32)(unsafe.Pointer(p))
	return int32(ptr.Size)
}

// SetSize sets the value in Size.
func (p Vector_ImU32) SetSize(value int32) {
	ptr := (*C.ImVector_ImU32)(unsafe.Pointer(p))
	ptr.Size = (C.int32_t)(value)
}

// Vector_ImU8 wraps struct ImVector_ImU8.
type Vector_ImU8 uintptr

// Vector_ImU8Nil is a null pointer.
var Vector_ImU8Nil Vector_ImU8

// Vector_ImU8SizeOf is the byte size of ImVector_ImU8.
const Vector_ImU8SizeOf = int(C.sizeof_ImVector_ImU8)

// Vector_ImU8Alloc allocates a continuous block of Vector_ImU8.
func Vector_ImU8Alloc(alloc ffi.Allocator, count int) Vector_ImU8 {
	ptr := alloc.Allocate(Vector_ImU8SizeOf * count)
	return Vector_ImU8(ptr)
}

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p Vector_ImU8) Offset(offset int) Vector_ImU8 {
	return p + Vector_ImU8(offset*Vector_ImU8SizeOf)
}

// GetCapacity returns the value in Capacity.
func (p Vector_ImU8) GetCapacity() int32 {
	ptr := (*C.ImVector_ImU8)(unsafe.Pointer(p))
	return int32(ptr.Capacity)
}

// SetCapacity sets the value in Capacity.
func (p Vector_ImU8) SetCapacity(value int32) {
	ptr := (*C.ImVector_ImU8)(unsafe.Pointer(p))
	ptr.Capacity = (C.int32_t)(value)
}

// GetData returns the value in Data.
func (p Vector_ImU8) GetData() ffi.Ref[uint8] {
	ptr := (*C.ImVector_ImU8)(unsafe.Pointer(p))
	return ffi.RefFromPtr[uint8](unsafe.Pointer(ptr.Data))
}

// SetData sets the value in Data.
func (p Vector_ImU8) SetData(value ffi.Ref[uint8]) {
	ptr := (*C.ImVector_ImU8)(unsafe.Pointer(p))
	ptr.Data = (*C.uint8_t)(value.Raw())
}

// GetSize returns the value in Size.
func (p Vector_ImU8) GetSize() int32 {
	ptr := (*C.ImVector_ImU8)(unsafe.Pointer(p))
	return int32(ptr.Size)
}

// SetSize sets the value in Size.
func (p Vector_ImU8) SetSize(value int32) {
	ptr := (*C.ImVector_ImU8)(unsafe.Pointer(p))
	ptr.Size = (C.int32_t)(value)
}

// Vector_ImVec2 wraps struct ImVector_ImVec2.
type Vector_ImVec2 uintptr

// Vector_ImVec2Nil is a null pointer.
var Vector_ImVec2Nil Vector_ImVec2

// Vector_ImVec2SizeOf is the byte size of ImVector_ImVec2.
const Vector_ImVec2SizeOf = int(C.sizeof_ImVector_ImVec2)

// Vector_ImVec2Alloc allocates a continuous block of Vector_ImVec2.
func Vector_ImVec2Alloc(alloc ffi.Allocator, count int) Vector_ImVec2 {
	ptr := alloc.Allocate(Vector_ImVec2SizeOf * count)
	return Vector_ImVec2(ptr)
}

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p Vector_ImVec2) Offset(offset int) Vector_ImVec2 {
	return p + Vector_ImVec2(offset*Vector_ImVec2SizeOf)
}

// GetCapacity returns the value in Capacity.
func (p Vector_ImVec2) GetCapacity() int32 {
	ptr := (*C.ImVector_ImVec2)(unsafe.Pointer(p))
	return int32(ptr.Capacity)
}

// SetCapacity sets the value in Capacity.
func (p Vector_ImVec2) SetCapacity(value int32) {
	ptr := (*C.ImVector_ImVec2)(unsafe.Pointer(p))
	ptr.Capacity = (C.int32_t)(value)
}

// GetData returns the value in Data.
func (p Vector_ImVec2) GetData() Vec2 {
	ptr := (*C.ImVector_ImVec2)(unsafe.Pointer(p))
	return Vec2(unsafe.Pointer(ptr.Data))
}

// SetData sets the value in Data.
func (p Vector_ImVec2) SetData(value Vec2) {
	ptr := (*C.ImVector_ImVec2)(unsafe.Pointer(p))
	ptr.Data = (*C.ImVec2)(unsafe.Pointer(value))
}

// GetSize returns the value in Size.
func (p Vector_ImVec2) GetSize() int32 {
	ptr := (*C.ImVector_ImVec2)(unsafe.Pointer(p))
	return int32(ptr.Size)
}

// SetSize sets the value in Size.
func (p Vector_ImVec2) SetSize(value int32) {
	ptr := (*C.ImVector_ImVec2)(unsafe.Pointer(p))
	ptr.Size = (C.int32_t)(value)
}

// Vector_ImVec4 wraps struct ImVector_ImVec4.
type Vector_ImVec4 uintptr

// Vector_ImVec4Nil is a null pointer.
var Vector_ImVec4Nil Vector_ImVec4

// Vector_ImVec4SizeOf is the byte size of ImVector_ImVec4.
const Vector_ImVec4SizeOf = int(C.sizeof_ImVector_ImVec4)

// Vector_ImVec4Alloc allocates a continuous block of Vector_ImVec4.
func Vector_ImVec4Alloc(alloc ffi.Allocator, count int) Vector_ImVec4 {
	ptr := alloc.Allocate(Vector_ImVec4SizeOf * count)
	return Vector_ImVec4(ptr)
}

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p Vector_ImVec4) Offset(offset int) Vector_ImVec4 {
	return p + Vector_ImVec4(offset*Vector_ImVec4SizeOf)
}

// GetCapacity returns the value in Capacity.
func (p Vector_ImVec4) GetCapacity() int32 {
	ptr := (*C.ImVector_ImVec4)(unsafe.Pointer(p))
	return int32(ptr.Capacity)
}

// SetCapacity sets the value in Capacity.
func (p Vector_ImVec4) SetCapacity(value int32) {
	ptr := (*C.ImVector_ImVec4)(unsafe.Pointer(p))
	ptr.Capacity = (C.int32_t)(value)
}

// GetData returns the value in Data.
func (p Vector_ImVec4) GetData() Vec4 {
	ptr := (*C.ImVector_ImVec4)(unsafe.Pointer(p))
	return Vec4(unsafe.Pointer(ptr.Data))
}

// SetData sets the value in Data.
func (p Vector_ImVec4) SetData(value Vec4) {
	ptr := (*C.ImVector_ImVec4)(unsafe.Pointer(p))
	ptr.Data = (*C.ImVec4)(unsafe.Pointer(value))
}

// GetSize returns the value in Size.
func (p Vector_ImVec4) GetSize() int32 {
	ptr := (*C.ImVector_ImVec4)(unsafe.Pointer(p))
	return int32(ptr.Size)
}

// SetSize sets the value in Size.
func (p Vector_ImVec4) SetSize(value int32) {
	ptr := (*C.ImVector_ImVec4)(unsafe.Pointer(p))
	ptr.Size = (C.int32_t)(value)
}

// Vector_ImWchar wraps struct ImVector_ImWchar.
type Vector_ImWchar uintptr

// Vector_ImWcharNil is a null pointer.
var Vector_ImWcharNil Vector_ImWchar

// Vector_ImWcharSizeOf is the byte size of ImVector_ImWchar.
const Vector_ImWcharSizeOf = int(C.sizeof_ImVector_ImWchar)

// Vector_ImWcharAlloc allocates a continuous block of Vector_ImWchar.
func Vector_ImWcharAlloc(alloc ffi.Allocator, count int) Vector_ImWchar {
	ptr := alloc.Allocate(Vector_ImWcharSizeOf * count)
	return Vector_ImWchar(ptr)
}

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p Vector_ImWchar) Offset(offset int) Vector_ImWchar {
	return p + Vector_ImWchar(offset*Vector_ImWcharSizeOf)
}

// GetCapacity returns the value in Capacity.
func (p Vector_ImWchar) GetCapacity() int32 {
	ptr := (*C.ImVector_ImWchar)(unsafe.Pointer(p))
	return int32(ptr.Capacity)
}

// SetCapacity sets the value in Capacity.
func (p Vector_ImWchar) SetCapacity(value int32) {
	ptr := (*C.ImVector_ImWchar)(unsafe.Pointer(p))
	ptr.Capacity = (C.int32_t)(value)
}

// Vector_ImWchar.Data is unsupported: category unsupported.

// GetSize returns the value in Size.
func (p Vector_ImWchar) GetSize() int32 {
	ptr := (*C.ImVector_ImWchar)(unsafe.Pointer(p))
	return int32(ptr.Size)
}

// SetSize sets the value in Size.
func (p Vector_ImWchar) SetSize(value int32) {
	ptr := (*C.ImVector_ImWchar)(unsafe.Pointer(p))
	ptr.Size = (C.int32_t)(value)
}

// Vector_char wraps struct ImVector_char.
type Vector_char uintptr

// Vector_charNil is a null pointer.
var Vector_charNil Vector_char

// Vector_charSizeOf is the byte size of ImVector_char.
const Vector_charSizeOf = int(C.sizeof_ImVector_char)

// Vector_charAlloc allocates a continuous block of Vector_char.
func Vector_charAlloc(alloc ffi.Allocator, count int) Vector_char {
	ptr := alloc.Allocate(Vector_charSizeOf * count)
	return Vector_char(ptr)
}

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p Vector_char) Offset(offset int) Vector_char {
	return p + Vector_char(offset*Vector_charSizeOf)
}

// GetCapacity returns the value in Capacity.
func (p Vector_char) GetCapacity() int32 {
	ptr := (*C.ImVector_char)(unsafe.Pointer(p))
	return int32(ptr.Capacity)
}

// SetCapacity sets the value in Capacity.
func (p Vector_char) SetCapacity(value int32) {
	ptr := (*C.ImVector_char)(unsafe.Pointer(p))
	ptr.Capacity = (C.int32_t)(value)
}

// GetData returns the value in Data.
func (p Vector_char) GetData() ffi.CString {
	ptr := (*C.ImVector_char)(unsafe.Pointer(p))
	return ffi.CStringFromPtr((unsafe.Pointer)(ptr.Data))
}

// SetData sets the value in Data.
func (p Vector_char) SetData(value ffi.CString) {
	ptr := (*C.ImVector_char)(unsafe.Pointer(p))
	ptr.Data = (*C.char)(value.Raw())
}

// GetSize returns the value in Size.
func (p Vector_char) GetSize() int32 {
	ptr := (*C.ImVector_char)(unsafe.Pointer(p))
	return int32(ptr.Size)
}

// SetSize sets the value in Size.
func (p Vector_char) SetSize(value int32) {
	ptr := (*C.ImVector_char)(unsafe.Pointer(p))
	ptr.Size = (C.int32_t)(value)
}

// Vector_const_charPtr wraps struct ImVector_const_charPtr.
type Vector_const_charPtr uintptr

// Vector_const_charPtrNil is a null pointer.
var Vector_const_charPtrNil Vector_const_charPtr

// Vector_const_charPtrSizeOf is the byte size of ImVector_const_charPtr.
const Vector_const_charPtrSizeOf = int(C.sizeof_ImVector_const_charPtr)

// Vector_const_charPtrAlloc allocates a continuous block of Vector_const_charPtr.
func Vector_const_charPtrAlloc(alloc ffi.Allocator, count int) Vector_const_charPtr {
	ptr := alloc.Allocate(Vector_const_charPtrSizeOf * count)
	return Vector_const_charPtr(ptr)
}

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p Vector_const_charPtr) Offset(offset int) Vector_const_charPtr {
	return p + Vector_const_charPtr(offset*Vector_const_charPtrSizeOf)
}

// GetCapacity returns the value in Capacity.
func (p Vector_const_charPtr) GetCapacity() int32 {
	ptr := (*C.ImVector_const_charPtr)(unsafe.Pointer(p))
	return int32(ptr.Capacity)
}

// SetCapacity sets the value in Capacity.
func (p Vector_const_charPtr) SetCapacity(value int32) {
	ptr := (*C.ImVector_const_charPtr)(unsafe.Pointer(p))
	ptr.Capacity = (C.int32_t)(value)
}

// GetData returns the value in Data.
func (p Vector_const_charPtr) GetData() ffi.Ref[ffi.CString] {
	ptr := (*C.ImVector_const_charPtr)(unsafe.Pointer(p))
	return ffi.RefFromPtr[ffi.CString](unsafe.Pointer(ptr.Data))
}

// SetData sets the value in Data.
func (p Vector_const_charPtr) SetData(value ffi.Ref[ffi.CString]) {
	ptr := (*C.ImVector_const_charPtr)(unsafe.Pointer(p))
	ptr.Data = (**C.char)(value.Raw())
}

// GetSize returns the value in Size.
func (p Vector_const_charPtr) GetSize() int32 {
	ptr := (*C.ImVector_const_charPtr)(unsafe.Pointer(p))
	return int32(ptr.Size)
}

// SetSize sets the value in Size.
func (p Vector_const_charPtr) SetSize(value int32) {
	ptr := (*C.ImVector_const_charPtr)(unsafe.Pointer(p))
	ptr.Size = (C.int32_t)(value)
}

// Vector_float wraps struct ImVector_float.
type Vector_float uintptr

// Vector_floatNil is a null pointer.
var Vector_floatNil Vector_float

// Vector_floatSizeOf is the byte size of ImVector_float.
const Vector_floatSizeOf = int(C.sizeof_ImVector_float)

// Vector_floatAlloc allocates a continuous block of Vector_float.
func Vector_floatAlloc(alloc ffi.Allocator, count int) Vector_float {
	ptr := alloc.Allocate(Vector_floatSizeOf * count)
	return Vector_float(ptr)
}

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p Vector_float) Offset(offset int) Vector_float {
	return p + Vector_float(offset*Vector_floatSizeOf)
}

// GetCapacity returns the value in Capacity.
func (p Vector_float) GetCapacity() int32 {
	ptr := (*C.ImVector_float)(unsafe.Pointer(p))
	return int32(ptr.Capacity)
}

// SetCapacity sets the value in Capacity.
func (p Vector_float) SetCapacity(value int32) {
	ptr := (*C.ImVector_float)(unsafe.Pointer(p))
	ptr.Capacity = (C.int32_t)(value)
}

// GetData returns the value in Data.
func (p Vector_float) GetData() ffi.Ref[float32] {
	ptr := (*C.ImVector_float)(unsafe.Pointer(p))
	return ffi.RefFromPtr[float32](unsafe.Pointer(ptr.Data))
}

// SetData sets the value in Data.
func (p Vector_float) SetData(value ffi.Ref[float32]) {
	ptr := (*C.ImVector_float)(unsafe.Pointer(p))
	ptr.Data = (*C.float)(value.Raw())
}

// GetSize returns the value in Size.
func (p Vector_float) GetSize() int32 {
	ptr := (*C.ImVector_float)(unsafe.Pointer(p))
	return int32(ptr.Size)
}

// SetSize sets the value in Size.
func (p Vector_float) SetSize(value int32) {
	ptr := (*C.ImVector_float)(unsafe.Pointer(p))
	ptr.Size = (C.int32_t)(value)
}

// Vector_int wraps struct ImVector_int.
type Vector_int uintptr

// Vector_intNil is a null pointer.
var Vector_intNil Vector_int

// Vector_intSizeOf is the byte size of ImVector_int.
const Vector_intSizeOf = int(C.sizeof_ImVector_int)

// Vector_intAlloc allocates a continuous block of Vector_int.
func Vector_intAlloc(alloc ffi.Allocator, count int) Vector_int {
	ptr := alloc.Allocate(Vector_intSizeOf * count)
	return Vector_int(ptr)
}

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p Vector_int) Offset(offset int) Vector_int {
	return p + Vector_int(offset*Vector_intSizeOf)
}

// GetCapacity returns the value in Capacity.
func (p Vector_int) GetCapacity() int32 {
	ptr := (*C.ImVector_int)(unsafe.Pointer(p))
	return int32(ptr.Capacity)
}

// SetCapacity sets the value in Capacity.
func (p Vector_int) SetCapacity(value int32) {
	ptr := (*C.ImVector_int)(unsafe.Pointer(p))
	ptr.Capacity = (C.int32_t)(value)
}

// GetData returns the value in Data.
func (p Vector_int) GetData() ffi.Ref[int32] {
	ptr := (*C.ImVector_int)(unsafe.Pointer(p))
	return ffi.RefFromPtr[int32](unsafe.Pointer(ptr.Data))
}

// SetData sets the value in Data.
func (p Vector_int) SetData(value ffi.Ref[int32]) {
	ptr := (*C.ImVector_int)(unsafe.Pointer(p))
	ptr.Data = (*C.int32_t)(value.Raw())
}

// GetSize returns the value in Size.
func (p Vector_int) GetSize() int32 {
	ptr := (*C.ImVector_int)(unsafe.Pointer(p))
	return int32(ptr.Size)
}

// SetSize sets the value in Size.
func (p Vector_int) SetSize(value int32) {
	ptr := (*C.ImVector_int)(unsafe.Pointer(p))
	ptr.Size = (C.int32_t)(value)
}

// Vector_stbrp_node_im wraps struct ImVector_stbrp_node_im.
type Vector_stbrp_node_im uintptr

// Vector_stbrp_node_imNil is a null pointer.
var Vector_stbrp_node_imNil Vector_stbrp_node_im

// Vector_stbrp_node_imSizeOf is the byte size of ImVector_stbrp_node_im.
const Vector_stbrp_node_imSizeOf = int(C.sizeof_ImVector_stbrp_node_im)

// Vector_stbrp_node_imAlloc allocates a continuous block of Vector_stbrp_node_im.
func Vector_stbrp_node_imAlloc(alloc ffi.Allocator, count int) Vector_stbrp_node_im {
	ptr := alloc.Allocate(Vector_stbrp_node_imSizeOf * count)
	return Vector_stbrp_node_im(ptr)
}

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p Vector_stbrp_node_im) Offset(offset int) Vector_stbrp_node_im {
	return p + Vector_stbrp_node_im(offset*Vector_stbrp_node_imSizeOf)
}

// GetCapacity returns the value in Capacity.
func (p Vector_stbrp_node_im) GetCapacity() int32 {
	ptr := (*C.ImVector_stbrp_node_im)(unsafe.Pointer(p))
	return int32(ptr.Capacity)
}

// SetCapacity sets the value in Capacity.
func (p Vector_stbrp_node_im) SetCapacity(value int32) {
	ptr := (*C.ImVector_stbrp_node_im)(unsafe.Pointer(p))
	ptr.Capacity = (C.int32_t)(value)
}

// Vector_stbrp_node_im.Data is unsupported: category unsupported.

// GetSize returns the value in Size.
func (p Vector_stbrp_node_im) GetSize() int32 {
	ptr := (*C.ImVector_stbrp_node_im)(unsafe.Pointer(p))
	return int32(ptr.Size)
}

// SetSize sets the value in Size.
func (p Vector_stbrp_node_im) SetSize(value int32) {
	ptr := (*C.ImVector_stbrp_node_im)(unsafe.Pointer(p))
	ptr.Size = (C.int32_t)(value)
}

// Vector_unsigned_char wraps struct ImVector_unsigned_char.
type Vector_unsigned_char uintptr

// Vector_unsigned_charNil is a null pointer.
var Vector_unsigned_charNil Vector_unsigned_char

// Vector_unsigned_charSizeOf is the byte size of ImVector_unsigned_char.
const Vector_unsigned_charSizeOf = int(C.sizeof_ImVector_unsigned_char)

// Vector_unsigned_charAlloc allocates a continuous block of Vector_unsigned_char.
func Vector_unsigned_charAlloc(alloc ffi.Allocator, count int) Vector_unsigned_char {
	ptr := alloc.Allocate(Vector_unsigned_charSizeOf * count)
	return Vector_unsigned_char(ptr)
}

// Offset returns an offset pointer by the size of the struct and the provided multiple.
func (p Vector_unsigned_char) Offset(offset int) Vector_unsigned_char {
	return p + Vector_unsigned_char(offset*Vector_unsigned_charSizeOf)
}

// GetCapacity returns the value in Capacity.
func (p Vector_unsigned_char) GetCapacity() int32 {
	ptr := (*C.ImVector_unsigned_char)(unsafe.Pointer(p))
	return int32(ptr.Capacity)
}

// SetCapacity sets the value in Capacity.
func (p Vector_unsigned_char) SetCapacity(value int32) {
	ptr := (*C.ImVector_unsigned_char)(unsafe.Pointer(p))
	ptr.Capacity = (C.int32_t)(value)
}

// GetData returns the value in Data.
func (p Vector_unsigned_char) GetData() ffi.Ref[uint8] {
	ptr := (*C.ImVector_unsigned_char)(unsafe.Pointer(p))
	return ffi.RefFromPtr[uint8](unsafe.Pointer(ptr.Data))
}

// SetData sets the value in Data.
func (p Vector_unsigned_char) SetData(value ffi.Ref[uint8]) {
	ptr := (*C.ImVector_unsigned_char)(unsafe.Pointer(p))
	ptr.Data = (*C.uint8_t)(value.Raw())
}

// GetSize returns the value in Size.
func (p Vector_unsigned_char) GetSize() int32 {
	ptr := (*C.ImVector_unsigned_char)(unsafe.Pointer(p))
	return int32(ptr.Size)
}

// SetSize sets the value in Size.
func (p Vector_unsigned_char) SetSize(value int32) {
	ptr := (*C.ImVector_unsigned_char)(unsafe.Pointer(p))
	ptr.Size = (C.int32_t)(value)
}
