//----------------------------------------
// The code is automatically generated by the GenlibLcl tool.
// Copyright © ying32. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package lcl

import (
	. "github.com/energye/golcl/lcl/api"
	. "github.com/energye/golcl/lcl/types"
	"time"
	"unsafe"
)

type TMonthCalendar struct {
	IWinControl
	instance unsafe.Pointer
}

// NewMonthCalendar
//
// 创建一个新的对象。
//
// Create a new object.
func NewMonthCalendar(owner IComponent) *TMonthCalendar {
	m := new(TMonthCalendar)
	m.instance = unsafe.Pointer(MonthCalendar_Create(CheckPtr(owner)))
	return m
}

// AsMonthCalendar
//
// 动态转换一个已存在的对象实例。
//
// Dynamically convert an existing object instance.
func AsMonthCalendar(obj interface{}) *TMonthCalendar {
	instance := getInstance(obj)
	if instance == nullptr {
		return nil
	}
	return &TMonthCalendar{instance: instance}
}

// Free
//
// 释放对象。
//
// Free object.
func (m *TMonthCalendar) Free() {
	if m.instance != nullptr {
		MonthCalendar_Free(m._instance())
		m.instance = nullptr
	}
}

func (m *TMonthCalendar) _instance() uintptr {
	return uintptr(m.instance)
}

// Instance
//
// 返回对象实例指针。
//
// Return object instance pointer.
func (m *TMonthCalendar) Instance() uintptr {
	return m._instance()
}

// UnsafeAddr
//
// 获取一个不安全的地址。
//
// Get an unsafe address.
func (m *TMonthCalendar) UnsafeAddr() unsafe.Pointer {
	return m.instance
}

// IsValid
//
// 检测地址是否为空。
//
// Check if the address is empty.
func (m *TMonthCalendar) IsValid() bool {
	return m.instance != nullptr
}

// Is
//
// 检测当前对象是否继承自目标对象。
//
// Checks whether the current object is inherited from the target object.
func (m *TMonthCalendar) Is() TIs {
	return TIs(m._instance())
}

// As
//
// 动态转换当前对象为目标对象。
//
// Dynamically convert the current object to the target object.
//func (m *TMonthCalendar) As() TAs {
//    return TAs(m._instance())
//}

// TMonthCalendarClass
//
// 获取类信息指针。
//
// Get class information pointer.
func TMonthCalendarClass() TClass {
	return MonthCalendar_StaticClassType()
}

// CanFocus
//
// 是否可以获得焦点。
func (m *TMonthCalendar) CanFocus() bool {
	return MonthCalendar_CanFocus(m._instance())
}

// ContainsControl
//
// 返回是否包含指定控件。
//
// it's contain a specified control.
func (m *TMonthCalendar) ContainsControl(Control IControl) bool {
	return MonthCalendar_ContainsControl(m._instance(), CheckPtr(Control))
}

// ControlAtPos
//
// 返回指定坐标及相关属性位置控件。
//
// Returns the specified coordinate and the relevant attribute position control..
func (m *TMonthCalendar) ControlAtPos(Pos TPoint, AllowDisabled bool, AllowWinControls bool, AllLevels bool) *TControl {
	return AsControl(MonthCalendar_ControlAtPos(m._instance(), Pos, AllowDisabled, AllowWinControls, AllLevels))
}

// DisableAlign
//
// 禁用控件的对齐。
//
// Disable control alignment.
func (m *TMonthCalendar) DisableAlign() {
	MonthCalendar_DisableAlign(m._instance())
}

// EnableAlign
//
// 启用控件对齐。
//
// Enabled control alignment.
func (m *TMonthCalendar) EnableAlign() {
	MonthCalendar_EnableAlign(m._instance())
}

// FindChildControl
//
// 查找子控件。
//
// Find sub controls.
func (m *TMonthCalendar) FindChildControl(ControlName string) *TControl {
	return AsControl(MonthCalendar_FindChildControl(m._instance(), ControlName))
}

func (m *TMonthCalendar) FlipChildren(AllLevels bool) {
	MonthCalendar_FlipChildren(m._instance(), AllLevels)
}

// Focused
//
// 返回是否获取焦点。
//
// Return to get focus.
func (m *TMonthCalendar) Focused() bool {
	return MonthCalendar_Focused(m._instance())
}

// HandleAllocated
//
// 句柄是否已经分配。
//
// Is the handle already allocated.
func (m *TMonthCalendar) HandleAllocated() bool {
	return MonthCalendar_HandleAllocated(m._instance())
}

// InsertControl
//
// 插入一个控件。
//
// Insert a control.
func (m *TMonthCalendar) InsertControl(AControl IControl) {
	MonthCalendar_InsertControl(m._instance(), CheckPtr(AControl))
}

// Invalidate
//
// 要求重绘。
//
// Redraw.
func (m *TMonthCalendar) Invalidate() {
	MonthCalendar_Invalidate(m._instance())
}

// PaintTo
//
// 绘画至指定DC。
//
// Painting to the specified DC.
func (m *TMonthCalendar) PaintTo(DC HDC, X int32, Y int32) {
	MonthCalendar_PaintTo(m._instance(), DC, X, Y)
}

// RemoveControl
//
// 移除一个控件。
//
// Remove a control.
func (m *TMonthCalendar) RemoveControl(AControl IControl) {
	MonthCalendar_RemoveControl(m._instance(), CheckPtr(AControl))
}

// Realign
//
// 重新对齐。
//
// Realign.
func (m *TMonthCalendar) Realign() {
	MonthCalendar_Realign(m._instance())
}

// Repaint
//
// 重绘。
//
// Repaint.
func (m *TMonthCalendar) Repaint() {
	MonthCalendar_Repaint(m._instance())
}

// ScaleBy
//
// 按比例缩放。
//
// Scale by.
func (m *TMonthCalendar) ScaleBy(M int32, D int32) {
	MonthCalendar_ScaleBy(m._instance(), M, D)
}

// ScrollBy
//
// 滚动至指定位置。
//
// Scroll by.
func (m *TMonthCalendar) ScrollBy(DeltaX int32, DeltaY int32) {
	MonthCalendar_ScrollBy(m._instance(), DeltaX, DeltaY)
}

// SetBounds
//
// 设置组件边界。
//
// Set component boundaries.
func (m *TMonthCalendar) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
	MonthCalendar_SetBounds(m._instance(), ALeft, ATop, AWidth, AHeight)
}

// SetFocus
//
// 设置控件焦点。
//
// Set control focus.
func (m *TMonthCalendar) SetFocus() {
	MonthCalendar_SetFocus(m._instance())
}

// Update
//
// 控件更新。
//
// Update.
func (m *TMonthCalendar) Update() {
	MonthCalendar_Update(m._instance())
}

// BringToFront
//
// 将控件置于最前。
//
// Bring the control to the front.
func (m *TMonthCalendar) BringToFront() {
	MonthCalendar_BringToFront(m._instance())
}

// ClientToScreen
//
// 将客户端坐标转为绝对的屏幕坐标。
//
// Convert client coordinates to absolute screen coordinates.
func (m *TMonthCalendar) ClientToScreen(Point TPoint) TPoint {
	return MonthCalendar_ClientToScreen(m._instance(), Point)
}

// ClientToParent
//
// 将客户端坐标转为父容器坐标。
//
// Convert client coordinates to parent container coordinates.
func (m *TMonthCalendar) ClientToParent(Point TPoint, AParent IWinControl) TPoint {
	return MonthCalendar_ClientToParent(m._instance(), Point, CheckPtr(AParent))
}

// Dragging
//
// 是否在拖拽中。
//
// Is it in the middle of dragging.
func (m *TMonthCalendar) Dragging() bool {
	return MonthCalendar_Dragging(m._instance())
}

// HasParent
//
// 是否有父容器。
//
// Is there a parent container.
func (m *TMonthCalendar) HasParent() bool {
	return MonthCalendar_HasParent(m._instance())
}

// Hide
//
// 隐藏控件。
//
// Hidden control.
func (m *TMonthCalendar) Hide() {
	MonthCalendar_Hide(m._instance())
}

// Perform
//
// 发送一个消息。
//
// Send a message.
func (m *TMonthCalendar) Perform(Msg uint32, WParam uintptr, LParam int) int {
	return MonthCalendar_Perform(m._instance(), Msg, WParam, LParam)
}

// Refresh
//
// 刷新控件。
//
// Refresh control.
func (m *TMonthCalendar) Refresh() {
	MonthCalendar_Refresh(m._instance())
}

// ScreenToClient
//
// 将屏幕坐标转为客户端坐标。
//
// Convert screen coordinates to client coordinates.
func (m *TMonthCalendar) ScreenToClient(Point TPoint) TPoint {
	return MonthCalendar_ScreenToClient(m._instance(), Point)
}

// ParentToClient
//
// 将父容器坐标转为客户端坐标。
//
// Convert parent container coordinates to client coordinates.
func (m *TMonthCalendar) ParentToClient(Point TPoint, AParent IWinControl) TPoint {
	return MonthCalendar_ParentToClient(m._instance(), Point, CheckPtr(AParent))
}

// SendToBack
//
// 控件至于最后面。
//
// The control is placed at the end.
func (m *TMonthCalendar) SendToBack() {
	MonthCalendar_SendToBack(m._instance())
}

// Show
//
// 显示控件。
//
// Show control.
func (m *TMonthCalendar) Show() {
	MonthCalendar_Show(m._instance())
}

// GetTextBuf
//
// 获取控件的字符，如果有。
//
// Get the characters of the control, if any.
func (m *TMonthCalendar) GetTextBuf(Buffer *string, BufSize int32) int32 {
	return MonthCalendar_GetTextBuf(m._instance(), Buffer, BufSize)
}

// GetTextLen
//
// 获取控件的字符长，如果有。
//
// Get the character length of the control, if any.
func (m *TMonthCalendar) GetTextLen() int32 {
	return MonthCalendar_GetTextLen(m._instance())
}

// SetTextBuf
//
// 设置控件字符，如果有。
//
// Set control characters, if any.
func (m *TMonthCalendar) SetTextBuf(Buffer string) {
	MonthCalendar_SetTextBuf(m._instance(), Buffer)
}

// FindComponent
//
// 查找指定名称的组件。
//
// Find the component with the specified name.
func (m *TMonthCalendar) FindComponent(AName string) *TComponent {
	return AsComponent(MonthCalendar_FindComponent(m._instance(), AName))
}

// GetNamePath
//
// 获取类名路径。
//
// Get the class name path.
func (m *TMonthCalendar) GetNamePath() string {
	return MonthCalendar_GetNamePath(m._instance())
}

// Assign
//
// 复制一个对象，如果对象实现了此方法的话。
//
// Copy an object, if the object implements this method.
func (m *TMonthCalendar) Assign(Source IObject) {
	MonthCalendar_Assign(m._instance(), CheckPtr(Source))
}

// ClassType
//
// 获取类的类型信息。
//
// Get class type information.
func (m *TMonthCalendar) ClassType() TClass {
	return MonthCalendar_ClassType(m._instance())
}

// ClassName
//
// 获取当前对象类名称。
//
// Get the current object class name.
func (m *TMonthCalendar) ClassName() string {
	return MonthCalendar_ClassName(m._instance())
}

// InstanceSize
//
// 获取当前对象实例大小。
//
// Get the current object instance size.
func (m *TMonthCalendar) InstanceSize() int32 {
	return MonthCalendar_InstanceSize(m._instance())
}

// InheritsFrom
//
// 判断当前类是否继承自指定类。
//
// Determine whether the current class inherits from the specified class.
func (m *TMonthCalendar) InheritsFrom(AClass TClass) bool {
	return MonthCalendar_InheritsFrom(m._instance(), AClass)
}

// Equals
//
// 与一个对象进行比较。
//
// Compare with an object.
func (m *TMonthCalendar) Equals(Obj IObject) bool {
	return MonthCalendar_Equals(m._instance(), CheckPtr(Obj))
}

// GetHashCode
//
// 获取类的哈希值。
//
// Get the hash value of the class.
func (m *TMonthCalendar) GetHashCode() int32 {
	return MonthCalendar_GetHashCode(m._instance())
}

// ToString
//
// 文本类信息。
//
// Text information.
func (m *TMonthCalendar) ToString() string {
	return MonthCalendar_ToString(m._instance())
}

func (m *TMonthCalendar) AnchorToNeighbour(ASide TAnchorKind, ASpace int32, ASibling IControl) {
	MonthCalendar_AnchorToNeighbour(m._instance(), ASide, ASpace, CheckPtr(ASibling))
}

func (m *TMonthCalendar) AnchorParallel(ASide TAnchorKind, ASpace int32, ASibling IControl) {
	MonthCalendar_AnchorParallel(m._instance(), ASide, ASpace, CheckPtr(ASibling))
}

// AnchorHorizontalCenterTo
//
// 置于指定控件的横向中心。
func (m *TMonthCalendar) AnchorHorizontalCenterTo(ASibling IControl) {
	MonthCalendar_AnchorHorizontalCenterTo(m._instance(), CheckPtr(ASibling))
}

// AnchorVerticalCenterTo
//
// 置于指定控件的纵向中心。
func (m *TMonthCalendar) AnchorVerticalCenterTo(ASibling IControl) {
	MonthCalendar_AnchorVerticalCenterTo(m._instance(), CheckPtr(ASibling))
}

func (m *TMonthCalendar) AnchorSame(ASide TAnchorKind, ASibling IControl) {
	MonthCalendar_AnchorSame(m._instance(), ASide, CheckPtr(ASibling))
}

func (m *TMonthCalendar) AnchorAsAlign(ATheAlign TAlign, ASpace int32) {
	MonthCalendar_AnchorAsAlign(m._instance(), ATheAlign, ASpace)
}

func (m *TMonthCalendar) AnchorClient(ASpace int32) {
	MonthCalendar_AnchorClient(m._instance(), ASpace)
}

func (m *TMonthCalendar) ScaleDesignToForm(ASize int32) int32 {
	return MonthCalendar_ScaleDesignToForm(m._instance(), ASize)
}

func (m *TMonthCalendar) ScaleFormToDesign(ASize int32) int32 {
	return MonthCalendar_ScaleFormToDesign(m._instance(), ASize)
}

func (m *TMonthCalendar) Scale96ToForm(ASize int32) int32 {
	return MonthCalendar_Scale96ToForm(m._instance(), ASize)
}

func (m *TMonthCalendar) ScaleFormTo96(ASize int32) int32 {
	return MonthCalendar_ScaleFormTo96(m._instance(), ASize)
}

func (m *TMonthCalendar) Scale96ToFont(ASize int32) int32 {
	return MonthCalendar_Scale96ToFont(m._instance(), ASize)
}

func (m *TMonthCalendar) ScaleFontTo96(ASize int32) int32 {
	return MonthCalendar_ScaleFontTo96(m._instance(), ASize)
}

func (m *TMonthCalendar) ScaleScreenToFont(ASize int32) int32 {
	return MonthCalendar_ScaleScreenToFont(m._instance(), ASize)
}

func (m *TMonthCalendar) ScaleFontToScreen(ASize int32) int32 {
	return MonthCalendar_ScaleFontToScreen(m._instance(), ASize)
}

func (m *TMonthCalendar) Scale96ToScreen(ASize int32) int32 {
	return MonthCalendar_Scale96ToScreen(m._instance(), ASize)
}

func (m *TMonthCalendar) ScaleScreenTo96(ASize int32) int32 {
	return MonthCalendar_ScaleScreenTo96(m._instance(), ASize)
}

func (m *TMonthCalendar) AutoAdjustLayout(AMode TLayoutAdjustmentPolicy, AFromPPI int32, AToPPI int32, AOldFormWidth int32, ANewFormWidth int32) {
	MonthCalendar_AutoAdjustLayout(m._instance(), AMode, AFromPPI, AToPPI, AOldFormWidth, ANewFormWidth)
}

func (m *TMonthCalendar) FixDesignFontsPPI(ADesignTimePPI int32) {
	MonthCalendar_FixDesignFontsPPI(m._instance(), ADesignTimePPI)
}

func (m *TMonthCalendar) ScaleFontsPPI(AToPPI int32, AProportion float64) {
	MonthCalendar_ScaleFontsPPI(m._instance(), AToPPI, AProportion)
}

func (m *TMonthCalendar) DateTime() time.Time {
	return MonthCalendar_GetDateTime(m._instance())
}

func (m *TMonthCalendar) SetDateTime(value time.Time) {
	MonthCalendar_SetDateTime(m._instance(), value)
}

// Align
//
// 获取控件自动调整。
//
// Get Control automatically adjusts.
func (m *TMonthCalendar) Align() TAlign {
	return MonthCalendar_GetAlign(m._instance())
}

// SetAlign
//
// 设置控件自动调整。
//
// Set Control automatically adjusts.
func (m *TMonthCalendar) SetAlign(value TAlign) {
	MonthCalendar_SetAlign(m._instance(), value)
}

// Anchors
//
// 获取四个角位置的锚点。
func (m *TMonthCalendar) Anchors() TAnchors {
	return MonthCalendar_GetAnchors(m._instance())
}

// SetAnchors
//
// 设置四个角位置的锚点。
func (m *TMonthCalendar) SetAnchors(value TAnchors) {
	MonthCalendar_SetAnchors(m._instance(), value)
}

// AutoSize
//
// 获取自动调整大小。
func (m *TMonthCalendar) AutoSize() bool {
	return MonthCalendar_GetAutoSize(m._instance())
}

// SetAutoSize
//
// 设置自动调整大小。
func (m *TMonthCalendar) SetAutoSize(value bool) {
	MonthCalendar_SetAutoSize(m._instance(), value)
}

// BorderWidth
//
// 获取边框的宽度。
func (m *TMonthCalendar) BorderWidth() int32 {
	return MonthCalendar_GetBorderWidth(m._instance())
}

// SetBorderWidth
//
// 设置边框的宽度。
func (m *TMonthCalendar) SetBorderWidth(value int32) {
	MonthCalendar_SetBorderWidth(m._instance(), value)
}

func (m *TMonthCalendar) BiDiMode() TBiDiMode {
	return MonthCalendar_GetBiDiMode(m._instance())
}

func (m *TMonthCalendar) SetBiDiMode(value TBiDiMode) {
	MonthCalendar_SetBiDiMode(m._instance(), value)
}

// Constraints
//
// 获取约束控件大小。
func (m *TMonthCalendar) Constraints() *TSizeConstraints {
	return AsSizeConstraints(MonthCalendar_GetConstraints(m._instance()))
}

// SetConstraints
//
// 设置约束控件大小。
func (m *TMonthCalendar) SetConstraints(value *TSizeConstraints) {
	MonthCalendar_SetConstraints(m._instance(), CheckPtr(value))
}

func (m *TMonthCalendar) Date() time.Time {
	return MonthCalendar_GetDate(m._instance())
}

func (m *TMonthCalendar) SetDate(value time.Time) {
	MonthCalendar_SetDate(m._instance(), value)
}

// DoubleBuffered
//
// 获取设置控件双缓冲。
//
// Get Set control double buffering.
func (m *TMonthCalendar) DoubleBuffered() bool {
	return MonthCalendar_GetDoubleBuffered(m._instance())
}

// SetDoubleBuffered
//
// 设置设置控件双缓冲。
//
// Set Set control double buffering.
func (m *TMonthCalendar) SetDoubleBuffered(value bool) {
	MonthCalendar_SetDoubleBuffered(m._instance(), value)
}

// DragCursor
//
// 获取设置控件拖拽时的光标。
//
// Get Set the cursor when the control is dragged.
func (m *TMonthCalendar) DragCursor() TCursor {
	return MonthCalendar_GetDragCursor(m._instance())
}

// SetDragCursor
//
// 设置设置控件拖拽时的光标。
//
// Set Set the cursor when the control is dragged.
func (m *TMonthCalendar) SetDragCursor(value TCursor) {
	MonthCalendar_SetDragCursor(m._instance(), value)
}

// DragKind
//
// 获取拖拽方式。
//
// Get Drag and drop.
func (m *TMonthCalendar) DragKind() TDragKind {
	return MonthCalendar_GetDragKind(m._instance())
}

// SetDragKind
//
// 设置拖拽方式。
//
// Set Drag and drop.
func (m *TMonthCalendar) SetDragKind(value TDragKind) {
	MonthCalendar_SetDragKind(m._instance(), value)
}

// DragMode
//
// 获取拖拽模式。
//
// Get Drag mode.
func (m *TMonthCalendar) DragMode() TDragMode {
	return MonthCalendar_GetDragMode(m._instance())
}

// SetDragMode
//
// 设置拖拽模式。
//
// Set Drag mode.
func (m *TMonthCalendar) SetDragMode(value TDragMode) {
	MonthCalendar_SetDragMode(m._instance(), value)
}

// Enabled
//
// 获取控件启用。
//
// Get the control enabled.
func (m *TMonthCalendar) Enabled() bool {
	return MonthCalendar_GetEnabled(m._instance())
}

// SetEnabled
//
// 设置控件启用。
//
// Set the control enabled.
func (m *TMonthCalendar) SetEnabled(value bool) {
	MonthCalendar_SetEnabled(m._instance(), value)
}

// Font
//
// 获取字体。
//
// Get Font.
func (m *TMonthCalendar) Font() *TFont {
	return AsFont(MonthCalendar_GetFont(m._instance()))
}

// SetFont
//
// 设置字体。
//
// Set Font.
func (m *TMonthCalendar) SetFont(value *TFont) {
	MonthCalendar_SetFont(m._instance(), CheckPtr(value))
}

// ParentDoubleBuffered
//
// 获取使用父容器双缓冲。
//
// Get Parent container double buffering.
func (m *TMonthCalendar) ParentDoubleBuffered() bool {
	return MonthCalendar_GetParentDoubleBuffered(m._instance())
}

// SetParentDoubleBuffered
//
// 设置使用父容器双缓冲。
//
// Set Parent container double buffering.
func (m *TMonthCalendar) SetParentDoubleBuffered(value bool) {
	MonthCalendar_SetParentDoubleBuffered(m._instance(), value)
}

// PopupMenu
//
// 获取右键菜单。
//
// Get Right click menu.
func (m *TMonthCalendar) PopupMenu() *TPopupMenu {
	return AsPopupMenu(MonthCalendar_GetPopupMenu(m._instance()))
}

// SetPopupMenu
//
// 设置右键菜单。
//
// Set Right click menu.
func (m *TMonthCalendar) SetPopupMenu(value IComponent) {
	MonthCalendar_SetPopupMenu(m._instance(), CheckPtr(value))
}

// ShowHint
//
// 获取显示鼠标悬停提示。
//
// Get Show mouseover tips.
func (m *TMonthCalendar) ShowHint() bool {
	return MonthCalendar_GetShowHint(m._instance())
}

// SetShowHint
//
// 设置显示鼠标悬停提示。
//
// Set Show mouseover tips.
func (m *TMonthCalendar) SetShowHint(value bool) {
	MonthCalendar_SetShowHint(m._instance(), value)
}

// TabOrder
//
// 获取Tab切换顺序序号。
//
// Get Tab switching sequence number.
func (m *TMonthCalendar) TabOrder() TTabOrder {
	return MonthCalendar_GetTabOrder(m._instance())
}

// SetTabOrder
//
// 设置Tab切换顺序序号。
//
// Set Tab switching sequence number.
func (m *TMonthCalendar) SetTabOrder(value TTabOrder) {
	MonthCalendar_SetTabOrder(m._instance(), value)
}

// TabStop
//
// 获取Tab可停留。
//
// Get Tab can stay.
func (m *TMonthCalendar) TabStop() bool {
	return MonthCalendar_GetTabStop(m._instance())
}

// SetTabStop
//
// 设置Tab可停留。
//
// Set Tab can stay.
func (m *TMonthCalendar) SetTabStop(value bool) {
	MonthCalendar_SetTabStop(m._instance(), value)
}

// Visible
//
// 获取控件可视。
//
// Get the control visible.
func (m *TMonthCalendar) Visible() bool {
	return MonthCalendar_GetVisible(m._instance())
}

// SetVisible
//
// 设置控件可视。
//
// Set the control visible.
func (m *TMonthCalendar) SetVisible(value bool) {
	MonthCalendar_SetVisible(m._instance(), value)
}

// SetOnClick
//
// 设置控件单击事件。
//
// Set control click event.
func (m *TMonthCalendar) SetOnClick(fn TNotifyEvent) {
	MonthCalendar_SetOnClick(m._instance(), fn)
}

// SetOnContextPopup
//
// 设置上下文弹出事件，一般是右键时弹出。
//
// Set Context popup event, usually pop up when right click.
func (m *TMonthCalendar) SetOnContextPopup(fn TContextPopupEvent) {
	MonthCalendar_SetOnContextPopup(m._instance(), fn)
}

// SetOnDblClick
//
// 设置双击事件。
func (m *TMonthCalendar) SetOnDblClick(fn TNotifyEvent) {
	MonthCalendar_SetOnDblClick(m._instance(), fn)
}

// SetOnDragDrop
//
// 设置拖拽下落事件。
//
// Set Drag and drop event.
func (m *TMonthCalendar) SetOnDragDrop(fn TDragDropEvent) {
	MonthCalendar_SetOnDragDrop(m._instance(), fn)
}

// SetOnDragOver
//
// 设置拖拽完成事件。
//
// Set Drag and drop completion event.
func (m *TMonthCalendar) SetOnDragOver(fn TDragOverEvent) {
	MonthCalendar_SetOnDragOver(m._instance(), fn)
}

// SetOnEndDock
//
// 设置停靠结束事件。
//
// Set Dock end event.
func (m *TMonthCalendar) SetOnEndDock(fn TEndDragEvent) {
	MonthCalendar_SetOnEndDock(m._instance(), fn)
}

// SetOnEndDrag
//
// 设置拖拽结束。
//
// Set End of drag.
func (m *TMonthCalendar) SetOnEndDrag(fn TEndDragEvent) {
	MonthCalendar_SetOnEndDrag(m._instance(), fn)
}

// SetOnEnter
//
// 设置焦点进入。
//
// Set Focus entry.
func (m *TMonthCalendar) SetOnEnter(fn TNotifyEvent) {
	MonthCalendar_SetOnEnter(m._instance(), fn)
}

// SetOnExit
//
// 设置焦点退出。
//
// Set Focus exit.
func (m *TMonthCalendar) SetOnExit(fn TNotifyEvent) {
	MonthCalendar_SetOnExit(m._instance(), fn)
}

// SetOnKeyDown
//
// 设置键盘按键按下事件。
//
// Set Keyboard button press event.
func (m *TMonthCalendar) SetOnKeyDown(fn TKeyEvent) {
	MonthCalendar_SetOnKeyDown(m._instance(), fn)
}

// SetOnKeyPress
//
// 设置键键下事件。
func (m *TMonthCalendar) SetOnKeyPress(fn TKeyPressEvent) {
	MonthCalendar_SetOnKeyPress(m._instance(), fn)
}

// SetOnKeyUp
//
// 设置键盘按键抬起事件。
//
// Set Keyboard button lift event.
func (m *TMonthCalendar) SetOnKeyUp(fn TKeyEvent) {
	MonthCalendar_SetOnKeyUp(m._instance(), fn)
}

// SetOnMouseEnter
//
// 设置鼠标进入事件。
//
// Set Mouse entry event.
func (m *TMonthCalendar) SetOnMouseEnter(fn TNotifyEvent) {
	MonthCalendar_SetOnMouseEnter(m._instance(), fn)
}

// SetOnMouseLeave
//
// 设置鼠标离开事件。
//
// Set Mouse leave event.
func (m *TMonthCalendar) SetOnMouseLeave(fn TNotifyEvent) {
	MonthCalendar_SetOnMouseLeave(m._instance(), fn)
}

// SetOnStartDock
//
// 设置启动停靠。
func (m *TMonthCalendar) SetOnStartDock(fn TStartDockEvent) {
	MonthCalendar_SetOnStartDock(m._instance(), fn)
}

// DockClientCount
//
// 获取依靠客户端总数。
func (m *TMonthCalendar) DockClientCount() int32 {
	return MonthCalendar_GetDockClientCount(m._instance())
}

// DockSite
//
// 获取停靠站点。
//
// Get Docking site.
func (m *TMonthCalendar) DockSite() bool {
	return MonthCalendar_GetDockSite(m._instance())
}

// SetDockSite
//
// 设置停靠站点。
//
// Set Docking site.
func (m *TMonthCalendar) SetDockSite(value bool) {
	MonthCalendar_SetDockSite(m._instance(), value)
}

// MouseInClient
//
// 获取鼠标是否在客户端，仅VCL有效。
//
// Get Whether the mouse is on the client, only VCL is valid.
func (m *TMonthCalendar) MouseInClient() bool {
	return MonthCalendar_GetMouseInClient(m._instance())
}

// VisibleDockClientCount
//
// 获取当前停靠的可视总数。
//
// Get The total number of visible calls currently docked.
func (m *TMonthCalendar) VisibleDockClientCount() int32 {
	return MonthCalendar_GetVisibleDockClientCount(m._instance())
}

// Brush
//
// 获取画刷对象。
//
// Get Brush.
func (m *TMonthCalendar) Brush() *TBrush {
	return AsBrush(MonthCalendar_GetBrush(m._instance()))
}

// ControlCount
//
// 获取子控件数。
//
// Get Number of child controls.
func (m *TMonthCalendar) ControlCount() int32 {
	return MonthCalendar_GetControlCount(m._instance())
}

// Handle
//
// 获取控件句柄。
//
// Get Control handle.
func (m *TMonthCalendar) Handle() HWND {
	return MonthCalendar_GetHandle(m._instance())
}

// ParentWindow
//
// 获取父容器句柄。
//
// Get Parent container handle.
func (m *TMonthCalendar) ParentWindow() HWND {
	return MonthCalendar_GetParentWindow(m._instance())
}

// SetParentWindow
//
// 设置父容器句柄。
//
// Set Parent container handle.
func (m *TMonthCalendar) SetParentWindow(value HWND) {
	MonthCalendar_SetParentWindow(m._instance(), value)
}

func (m *TMonthCalendar) Showing() bool {
	return MonthCalendar_GetShowing(m._instance())
}

// UseDockManager
//
// 获取使用停靠管理。
func (m *TMonthCalendar) UseDockManager() bool {
	return MonthCalendar_GetUseDockManager(m._instance())
}

// SetUseDockManager
//
// 设置使用停靠管理。
func (m *TMonthCalendar) SetUseDockManager(value bool) {
	MonthCalendar_SetUseDockManager(m._instance(), value)
}

func (m *TMonthCalendar) Action() *TAction {
	return AsAction(MonthCalendar_GetAction(m._instance()))
}

func (m *TMonthCalendar) SetAction(value IComponent) {
	MonthCalendar_SetAction(m._instance(), CheckPtr(value))
}

func (m *TMonthCalendar) BoundsRect() TRect {
	return MonthCalendar_GetBoundsRect(m._instance())
}

func (m *TMonthCalendar) SetBoundsRect(value TRect) {
	MonthCalendar_SetBoundsRect(m._instance(), value)
}

// ClientHeight
//
// 获取客户区高度。
//
// Get client height.
func (m *TMonthCalendar) ClientHeight() int32 {
	return MonthCalendar_GetClientHeight(m._instance())
}

// SetClientHeight
//
// 设置客户区高度。
//
// Set client height.
func (m *TMonthCalendar) SetClientHeight(value int32) {
	MonthCalendar_SetClientHeight(m._instance(), value)
}

func (m *TMonthCalendar) ClientOrigin() TPoint {
	return MonthCalendar_GetClientOrigin(m._instance())
}

// ClientRect
//
// 获取客户区矩形。
//
// Get client rectangle.
func (m *TMonthCalendar) ClientRect() TRect {
	return MonthCalendar_GetClientRect(m._instance())
}

// ClientWidth
//
// 获取客户区宽度。
//
// Get client width.
func (m *TMonthCalendar) ClientWidth() int32 {
	return MonthCalendar_GetClientWidth(m._instance())
}

// SetClientWidth
//
// 设置客户区宽度。
//
// Set client width.
func (m *TMonthCalendar) SetClientWidth(value int32) {
	MonthCalendar_SetClientWidth(m._instance(), value)
}

// ControlState
//
// 获取控件状态。
//
// Get control state.
func (m *TMonthCalendar) ControlState() TControlState {
	return MonthCalendar_GetControlState(m._instance())
}

// SetControlState
//
// 设置控件状态。
//
// Set control state.
func (m *TMonthCalendar) SetControlState(value TControlState) {
	MonthCalendar_SetControlState(m._instance(), value)
}

// ControlStyle
//
// 获取控件样式。
//
// Get control style.
func (m *TMonthCalendar) ControlStyle() TControlStyle {
	return MonthCalendar_GetControlStyle(m._instance())
}

// SetControlStyle
//
// 设置控件样式。
//
// Set control style.
func (m *TMonthCalendar) SetControlStyle(value TControlStyle) {
	MonthCalendar_SetControlStyle(m._instance(), value)
}

func (m *TMonthCalendar) Floating() bool {
	return MonthCalendar_GetFloating(m._instance())
}

// Parent
//
// 获取控件父容器。
//
// Get control parent container.
func (m *TMonthCalendar) Parent() *TWinControl {
	return AsWinControl(MonthCalendar_GetParent(m._instance()))
}

// SetParent
//
// 设置控件父容器。
//
// Set control parent container.
func (m *TMonthCalendar) SetParent(value IWinControl) {
	MonthCalendar_SetParent(m._instance(), CheckPtr(value))
}

// Left
//
// 获取左边位置。
//
// Get Left position.
func (m *TMonthCalendar) Left() int32 {
	return MonthCalendar_GetLeft(m._instance())
}

// SetLeft
//
// 设置左边位置。
//
// Set Left position.
func (m *TMonthCalendar) SetLeft(value int32) {
	MonthCalendar_SetLeft(m._instance(), value)
}

// Top
//
// 获取顶边位置。
//
// Get Top position.
func (m *TMonthCalendar) Top() int32 {
	return MonthCalendar_GetTop(m._instance())
}

// SetTop
//
// 设置顶边位置。
//
// Set Top position.
func (m *TMonthCalendar) SetTop(value int32) {
	MonthCalendar_SetTop(m._instance(), value)
}

// Width
//
// 获取宽度。
//
// Get width.
func (m *TMonthCalendar) Width() int32 {
	return MonthCalendar_GetWidth(m._instance())
}

// SetWidth
//
// 设置宽度。
//
// Set width.
func (m *TMonthCalendar) SetWidth(value int32) {
	MonthCalendar_SetWidth(m._instance(), value)
}

// Height
//
// 获取高度。
//
// Get height.
func (m *TMonthCalendar) Height() int32 {
	return MonthCalendar_GetHeight(m._instance())
}

// SetHeight
//
// 设置高度。
//
// Set height.
func (m *TMonthCalendar) SetHeight(value int32) {
	MonthCalendar_SetHeight(m._instance(), value)
}

// Cursor
//
// 获取控件光标。
//
// Get control cursor.
func (m *TMonthCalendar) Cursor() TCursor {
	return MonthCalendar_GetCursor(m._instance())
}

// SetCursor
//
// 设置控件光标。
//
// Set control cursor.
func (m *TMonthCalendar) SetCursor(value TCursor) {
	MonthCalendar_SetCursor(m._instance(), value)
}

// Hint
//
// 获取组件鼠标悬停提示。
//
// Get component mouse hints.
func (m *TMonthCalendar) Hint() string {
	return MonthCalendar_GetHint(m._instance())
}

// SetHint
//
// 设置组件鼠标悬停提示。
//
// Set component mouse hints.
func (m *TMonthCalendar) SetHint(value string) {
	MonthCalendar_SetHint(m._instance(), value)
}

// ComponentCount
//
// 获取组件总数。
//
// Get the total number of components.
func (m *TMonthCalendar) ComponentCount() int32 {
	return MonthCalendar_GetComponentCount(m._instance())
}

// ComponentIndex
//
// 获取组件索引。
//
// Get component index.
func (m *TMonthCalendar) ComponentIndex() int32 {
	return MonthCalendar_GetComponentIndex(m._instance())
}

// SetComponentIndex
//
// 设置组件索引。
//
// Set component index.
func (m *TMonthCalendar) SetComponentIndex(value int32) {
	MonthCalendar_SetComponentIndex(m._instance(), value)
}

// Owner
//
// 获取组件所有者。
//
// Get component owner.
func (m *TMonthCalendar) Owner() *TComponent {
	return AsComponent(MonthCalendar_GetOwner(m._instance()))
}

// Name
//
// 获取组件名称。
//
// Get the component name.
func (m *TMonthCalendar) Name() string {
	return MonthCalendar_GetName(m._instance())
}

// SetName
//
// 设置组件名称。
//
// Set the component name.
func (m *TMonthCalendar) SetName(value string) {
	MonthCalendar_SetName(m._instance(), value)
}

// Tag
//
// 获取对象标记。
//
// Get the control tag.
func (m *TMonthCalendar) Tag() int {
	return MonthCalendar_GetTag(m._instance())
}

// SetTag
//
// 设置对象标记。
//
// Set the control tag.
func (m *TMonthCalendar) SetTag(value int) {
	MonthCalendar_SetTag(m._instance(), value)
}

// AnchorSideLeft
//
// 获取左边锚点。
func (m *TMonthCalendar) AnchorSideLeft() *TAnchorSide {
	return AsAnchorSide(MonthCalendar_GetAnchorSideLeft(m._instance()))
}

// SetAnchorSideLeft
//
// 设置左边锚点。
func (m *TMonthCalendar) SetAnchorSideLeft(value *TAnchorSide) {
	MonthCalendar_SetAnchorSideLeft(m._instance(), CheckPtr(value))
}

// AnchorSideTop
//
// 获取顶边锚点。
func (m *TMonthCalendar) AnchorSideTop() *TAnchorSide {
	return AsAnchorSide(MonthCalendar_GetAnchorSideTop(m._instance()))
}

// SetAnchorSideTop
//
// 设置顶边锚点。
func (m *TMonthCalendar) SetAnchorSideTop(value *TAnchorSide) {
	MonthCalendar_SetAnchorSideTop(m._instance(), CheckPtr(value))
}

// AnchorSideRight
//
// 获取右边锚点。
func (m *TMonthCalendar) AnchorSideRight() *TAnchorSide {
	return AsAnchorSide(MonthCalendar_GetAnchorSideRight(m._instance()))
}

// SetAnchorSideRight
//
// 设置右边锚点。
func (m *TMonthCalendar) SetAnchorSideRight(value *TAnchorSide) {
	MonthCalendar_SetAnchorSideRight(m._instance(), CheckPtr(value))
}

// AnchorSideBottom
//
// 获取底边锚点。
func (m *TMonthCalendar) AnchorSideBottom() *TAnchorSide {
	return AsAnchorSide(MonthCalendar_GetAnchorSideBottom(m._instance()))
}

// SetAnchorSideBottom
//
// 设置底边锚点。
func (m *TMonthCalendar) SetAnchorSideBottom(value *TAnchorSide) {
	MonthCalendar_SetAnchorSideBottom(m._instance(), CheckPtr(value))
}

func (m *TMonthCalendar) ChildSizing() *TControlChildSizing {
	return AsControlChildSizing(MonthCalendar_GetChildSizing(m._instance()))
}

func (m *TMonthCalendar) SetChildSizing(value *TControlChildSizing) {
	MonthCalendar_SetChildSizing(m._instance(), CheckPtr(value))
}

// BorderSpacing
//
// 获取边框间距。
func (m *TMonthCalendar) BorderSpacing() *TControlBorderSpacing {
	return AsControlBorderSpacing(MonthCalendar_GetBorderSpacing(m._instance()))
}

// SetBorderSpacing
//
// 设置边框间距。
func (m *TMonthCalendar) SetBorderSpacing(value *TControlBorderSpacing) {
	MonthCalendar_SetBorderSpacing(m._instance(), CheckPtr(value))
}

// DockClients
//
// 获取指定索引停靠客户端。
func (m *TMonthCalendar) DockClients(Index int32) *TControl {
	return AsControl(MonthCalendar_GetDockClients(m._instance(), Index))
}

// Controls
//
// 获取指定索引子控件。
func (m *TMonthCalendar) Controls(Index int32) *TControl {
	return AsControl(MonthCalendar_GetControls(m._instance(), Index))
}

// Components
//
// 获取指定索引组件。
//
// Get the specified index component.
func (m *TMonthCalendar) Components(AIndex int32) *TComponent {
	return AsComponent(MonthCalendar_GetComponents(m._instance(), AIndex))
}

// AnchorSide
//
// 获取锚侧面。
func (m *TMonthCalendar) AnchorSide(AKind TAnchorKind) *TAnchorSide {
	return AsAnchorSide(MonthCalendar_GetAnchorSide(m._instance(), AKind))
}
