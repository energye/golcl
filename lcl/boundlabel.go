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
	"unsafe"
)

type TBoundLabel struct {
	IControl
	instance unsafe.Pointer
}

// NewBoundLabel
//
// 创建一个新的对象。
//
// Create a new object.
func NewBoundLabel(owner IComponent) *TBoundLabel {
	b := new(TBoundLabel)
	b.instance = unsafe.Pointer(BoundLabel_Create(CheckPtr(owner)))
	return b
}

// AsBoundLabel
//
// 动态转换一个已存在的对象实例。
//
// Dynamically convert an existing object instance.
func AsBoundLabel(obj interface{}) *TBoundLabel {
	instance := getInstance(obj)
	if instance == nullptr {
		return nil
	}
	return &TBoundLabel{instance: instance}
}

// Free
//
// 释放对象。
//
// Free object.
func (b *TBoundLabel) Free() {
	if b.instance != nullptr {
		BoundLabel_Free(b._instance())
		b.instance = nullptr
	}
}

func (b *TBoundLabel) _instance() uintptr {
	return uintptr(b.instance)
}

// Instance
//
// 返回对象实例指针。
//
// Return object instance pointer.
func (b *TBoundLabel) Instance() uintptr {
	return b._instance()
}

// UnsafeAddr
//
// 获取一个不安全的地址。
//
// Get an unsafe address.
func (b *TBoundLabel) UnsafeAddr() unsafe.Pointer {
	return b.instance
}

// IsValid
//
// 检测地址是否为空。
//
// Check if the address is empty.
func (b *TBoundLabel) IsValid() bool {
	return b.instance != nullptr
}

// Is
//
// 检测当前对象是否继承自目标对象。
//
// Checks whether the current object is inherited from the target object.
func (b *TBoundLabel) Is() TIs {
	return TIs(b._instance())
}

// As
//
// 动态转换当前对象为目标对象。
//
// Dynamically convert the current object to the target object.
//func (b *TBoundLabel) As() TAs {
//    return TAs(b._instance())
//}

// TBoundLabelClass
//
// 获取类信息指针。
//
// Get class information pointer.
func TBoundLabelClass() TClass {
	return BoundLabel_StaticClassType()
}

// BringToFront
//
// 将控件置于最前。
//
// Bring the control to the front.
func (b *TBoundLabel) BringToFront() {
	BoundLabel_BringToFront(b._instance())
}

// ClientToScreen
//
// 将客户端坐标转为绝对的屏幕坐标。
//
// Convert client coordinates to absolute screen coordinates.
func (b *TBoundLabel) ClientToScreen(Point TPoint) TPoint {
	return BoundLabel_ClientToScreen(b._instance(), Point)
}

// ClientToParent
//
// 将客户端坐标转为父容器坐标。
//
// Convert client coordinates to parent container coordinates.
func (b *TBoundLabel) ClientToParent(Point TPoint, AParent IWinControl) TPoint {
	return BoundLabel_ClientToParent(b._instance(), Point, CheckPtr(AParent))
}

// Dragging
//
// 是否在拖拽中。
//
// Is it in the middle of dragging.
func (b *TBoundLabel) Dragging() bool {
	return BoundLabel_Dragging(b._instance())
}

// HasParent
//
// 是否有父容器。
//
// Is there a parent container.
func (b *TBoundLabel) HasParent() bool {
	return BoundLabel_HasParent(b._instance())
}

// Hide
//
// 隐藏控件。
//
// Hidden control.
func (b *TBoundLabel) Hide() {
	BoundLabel_Hide(b._instance())
}

// Invalidate
//
// 要求重绘。
//
// Redraw.
func (b *TBoundLabel) Invalidate() {
	BoundLabel_Invalidate(b._instance())
}

// Perform
//
// 发送一个消息。
//
// Send a message.
func (b *TBoundLabel) Perform(Msg uint32, WParam uintptr, LParam int) int {
	return BoundLabel_Perform(b._instance(), Msg, WParam, LParam)
}

// Refresh
//
// 刷新控件。
//
// Refresh control.
func (b *TBoundLabel) Refresh() {
	BoundLabel_Refresh(b._instance())
}

// Repaint
//
// 重绘。
//
// Repaint.
func (b *TBoundLabel) Repaint() {
	BoundLabel_Repaint(b._instance())
}

// ScreenToClient
//
// 将屏幕坐标转为客户端坐标。
//
// Convert screen coordinates to client coordinates.
func (b *TBoundLabel) ScreenToClient(Point TPoint) TPoint {
	return BoundLabel_ScreenToClient(b._instance(), Point)
}

// ParentToClient
//
// 将父容器坐标转为客户端坐标。
//
// Convert parent container coordinates to client coordinates.
func (b *TBoundLabel) ParentToClient(Point TPoint, AParent IWinControl) TPoint {
	return BoundLabel_ParentToClient(b._instance(), Point, CheckPtr(AParent))
}

// SendToBack
//
// 控件至于最后面。
//
// The control is placed at the end.
func (b *TBoundLabel) SendToBack() {
	BoundLabel_SendToBack(b._instance())
}

// SetBounds
//
// 设置组件边界。
//
// Set component boundaries.
func (b *TBoundLabel) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
	BoundLabel_SetBounds(b._instance(), ALeft, ATop, AWidth, AHeight)
}

// Show
//
// 显示控件。
//
// Show control.
func (b *TBoundLabel) Show() {
	BoundLabel_Show(b._instance())
}

// Update
//
// 控件更新。
//
// Update.
func (b *TBoundLabel) Update() {
	BoundLabel_Update(b._instance())
}

// GetTextBuf
//
// 获取控件的字符，如果有。
//
// Get the characters of the control, if any.
func (b *TBoundLabel) GetTextBuf(Buffer *string, BufSize int32) int32 {
	return BoundLabel_GetTextBuf(b._instance(), Buffer, BufSize)
}

// GetTextLen
//
// 获取控件的字符长，如果有。
//
// Get the character length of the control, if any.
func (b *TBoundLabel) GetTextLen() int32 {
	return BoundLabel_GetTextLen(b._instance())
}

// SetTextBuf
//
// 设置控件字符，如果有。
//
// Set control characters, if any.
func (b *TBoundLabel) SetTextBuf(Buffer string) {
	BoundLabel_SetTextBuf(b._instance(), Buffer)
}

// FindComponent
//
// 查找指定名称的组件。
//
// Find the component with the specified name.
func (b *TBoundLabel) FindComponent(AName string) *TComponent {
	return AsComponent(BoundLabel_FindComponent(b._instance(), AName))
}

// GetNamePath
//
// 获取类名路径。
//
// Get the class name path.
func (b *TBoundLabel) GetNamePath() string {
	return BoundLabel_GetNamePath(b._instance())
}

// Assign
//
// 复制一个对象，如果对象实现了此方法的话。
//
// Copy an object, if the object implements this method.
func (b *TBoundLabel) Assign(Source IObject) {
	BoundLabel_Assign(b._instance(), CheckPtr(Source))
}

// ClassType
//
// 获取类的类型信息。
//
// Get class type information.
func (b *TBoundLabel) ClassType() TClass {
	return BoundLabel_ClassType(b._instance())
}

// ClassName
//
// 获取当前对象类名称。
//
// Get the current object class name.
func (b *TBoundLabel) ClassName() string {
	return BoundLabel_ClassName(b._instance())
}

// InstanceSize
//
// 获取当前对象实例大小。
//
// Get the current object instance size.
func (b *TBoundLabel) InstanceSize() int32 {
	return BoundLabel_InstanceSize(b._instance())
}

// InheritsFrom
//
// 判断当前类是否继承自指定类。
//
// Determine whether the current class inherits from the specified class.
func (b *TBoundLabel) InheritsFrom(AClass TClass) bool {
	return BoundLabel_InheritsFrom(b._instance(), AClass)
}

// Equals
//
// 与一个对象进行比较。
//
// Compare with an object.
func (b *TBoundLabel) Equals(Obj IObject) bool {
	return BoundLabel_Equals(b._instance(), CheckPtr(Obj))
}

// GetHashCode
//
// 获取类的哈希值。
//
// Get the hash value of the class.
func (b *TBoundLabel) GetHashCode() int32 {
	return BoundLabel_GetHashCode(b._instance())
}

// ToString
//
// 文本类信息。
//
// Text information.
func (b *TBoundLabel) ToString() string {
	return BoundLabel_ToString(b._instance())
}

func (b *TBoundLabel) AnchorToNeighbour(ASide TAnchorKind, ASpace int32, ASibling IControl) {
	BoundLabel_AnchorToNeighbour(b._instance(), ASide, ASpace, CheckPtr(ASibling))
}

func (b *TBoundLabel) AnchorParallel(ASide TAnchorKind, ASpace int32, ASibling IControl) {
	BoundLabel_AnchorParallel(b._instance(), ASide, ASpace, CheckPtr(ASibling))
}

// AnchorHorizontalCenterTo
//
// 置于指定控件的横向中心。
func (b *TBoundLabel) AnchorHorizontalCenterTo(ASibling IControl) {
	BoundLabel_AnchorHorizontalCenterTo(b._instance(), CheckPtr(ASibling))
}

// AnchorVerticalCenterTo
//
// 置于指定控件的纵向中心。
func (b *TBoundLabel) AnchorVerticalCenterTo(ASibling IControl) {
	BoundLabel_AnchorVerticalCenterTo(b._instance(), CheckPtr(ASibling))
}

func (b *TBoundLabel) AnchorSame(ASide TAnchorKind, ASibling IControl) {
	BoundLabel_AnchorSame(b._instance(), ASide, CheckPtr(ASibling))
}

func (b *TBoundLabel) AnchorAsAlign(ATheAlign TAlign, ASpace int32) {
	BoundLabel_AnchorAsAlign(b._instance(), ATheAlign, ASpace)
}

func (b *TBoundLabel) AnchorClient(ASpace int32) {
	BoundLabel_AnchorClient(b._instance(), ASpace)
}

func (b *TBoundLabel) ScaleDesignToForm(ASize int32) int32 {
	return BoundLabel_ScaleDesignToForm(b._instance(), ASize)
}

func (b *TBoundLabel) ScaleFormToDesign(ASize int32) int32 {
	return BoundLabel_ScaleFormToDesign(b._instance(), ASize)
}

func (b *TBoundLabel) Scale96ToForm(ASize int32) int32 {
	return BoundLabel_Scale96ToForm(b._instance(), ASize)
}

func (b *TBoundLabel) ScaleFormTo96(ASize int32) int32 {
	return BoundLabel_ScaleFormTo96(b._instance(), ASize)
}

func (b *TBoundLabel) Scale96ToFont(ASize int32) int32 {
	return BoundLabel_Scale96ToFont(b._instance(), ASize)
}

func (b *TBoundLabel) ScaleFontTo96(ASize int32) int32 {
	return BoundLabel_ScaleFontTo96(b._instance(), ASize)
}

func (b *TBoundLabel) ScaleScreenToFont(ASize int32) int32 {
	return BoundLabel_ScaleScreenToFont(b._instance(), ASize)
}

func (b *TBoundLabel) ScaleFontToScreen(ASize int32) int32 {
	return BoundLabel_ScaleFontToScreen(b._instance(), ASize)
}

func (b *TBoundLabel) Scale96ToScreen(ASize int32) int32 {
	return BoundLabel_Scale96ToScreen(b._instance(), ASize)
}

func (b *TBoundLabel) ScaleScreenTo96(ASize int32) int32 {
	return BoundLabel_ScaleScreenTo96(b._instance(), ASize)
}

func (b *TBoundLabel) AutoAdjustLayout(AMode TLayoutAdjustmentPolicy, AFromPPI int32, AToPPI int32, AOldFormWidth int32, ANewFormWidth int32) {
	BoundLabel_AutoAdjustLayout(b._instance(), AMode, AFromPPI, AToPPI, AOldFormWidth, ANewFormWidth)
}

func (b *TBoundLabel) FixDesignFontsPPI(ADesignTimePPI int32) {
	BoundLabel_FixDesignFontsPPI(b._instance(), ADesignTimePPI)
}

func (b *TBoundLabel) ScaleFontsPPI(AToPPI int32, AProportion float64) {
	BoundLabel_ScaleFontsPPI(b._instance(), AToPPI, AProportion)
}

func (b *TBoundLabel) BiDiMode() TBiDiMode {
	return BoundLabel_GetBiDiMode(b._instance())
}

func (b *TBoundLabel) SetBiDiMode(value TBiDiMode) {
	BoundLabel_SetBiDiMode(b._instance(), value)
}

// Caption
//
// 获取控件标题。
//
// Get the control title.
func (b *TBoundLabel) Caption() string {
	return BoundLabel_GetCaption(b._instance())
}

// SetCaption
//
// 设置控件标题。
//
// Set the control title.
func (b *TBoundLabel) SetCaption(value string) {
	BoundLabel_SetCaption(b._instance(), value)
}

// Color
//
// 获取颜色。
//
// Get color.
func (b *TBoundLabel) Color() TColor {
	return BoundLabel_GetColor(b._instance())
}

// SetColor
//
// 设置颜色。
//
// Set color.
func (b *TBoundLabel) SetColor(value TColor) {
	BoundLabel_SetColor(b._instance(), value)
}

// DragCursor
//
// 获取设置控件拖拽时的光标。
//
// Get Set the cursor when the control is dragged.
func (b *TBoundLabel) DragCursor() TCursor {
	return BoundLabel_GetDragCursor(b._instance())
}

// SetDragCursor
//
// 设置设置控件拖拽时的光标。
//
// Set Set the cursor when the control is dragged.
func (b *TBoundLabel) SetDragCursor(value TCursor) {
	BoundLabel_SetDragCursor(b._instance(), value)
}

// DragMode
//
// 获取拖拽模式。
//
// Get Drag mode.
func (b *TBoundLabel) DragMode() TDragMode {
	return BoundLabel_GetDragMode(b._instance())
}

// SetDragMode
//
// 设置拖拽模式。
//
// Set Drag mode.
func (b *TBoundLabel) SetDragMode(value TDragMode) {
	BoundLabel_SetDragMode(b._instance(), value)
}

// Font
//
// 获取字体。
//
// Get Font.
func (b *TBoundLabel) Font() *TFont {
	return AsFont(BoundLabel_GetFont(b._instance()))
}

// SetFont
//
// 设置字体。
//
// Set Font.
func (b *TBoundLabel) SetFont(value *TFont) {
	BoundLabel_SetFont(b._instance(), CheckPtr(value))
}

// Height
//
// 获取高度。
//
// Get height.
func (b *TBoundLabel) Height() int32 {
	return BoundLabel_GetHeight(b._instance())
}

// SetHeight
//
// 设置高度。
//
// Set height.
func (b *TBoundLabel) SetHeight(value int32) {
	BoundLabel_SetHeight(b._instance(), value)
}

// Left
//
// 获取左边位置。
//
// Get Left position.
func (b *TBoundLabel) Left() int32 {
	return BoundLabel_GetLeft(b._instance())
}

// ParentColor
//
// 获取使用父容器颜色。
//
// Get parent color.
func (b *TBoundLabel) ParentColor() bool {
	return BoundLabel_GetParentColor(b._instance())
}

// SetParentColor
//
// 设置使用父容器颜色。
//
// Set parent color.
func (b *TBoundLabel) SetParentColor(value bool) {
	BoundLabel_SetParentColor(b._instance(), value)
}

// ParentFont
//
// 获取使用父容器字体。
//
// Get Parent container font.
func (b *TBoundLabel) ParentFont() bool {
	return BoundLabel_GetParentFont(b._instance())
}

// SetParentFont
//
// 设置使用父容器字体。
//
// Set Parent container font.
func (b *TBoundLabel) SetParentFont(value bool) {
	BoundLabel_SetParentFont(b._instance(), value)
}

// ParentShowHint
//
// 获取以父容器的ShowHint属性为准。
func (b *TBoundLabel) ParentShowHint() bool {
	return BoundLabel_GetParentShowHint(b._instance())
}

// SetParentShowHint
//
// 设置以父容器的ShowHint属性为准。
func (b *TBoundLabel) SetParentShowHint(value bool) {
	BoundLabel_SetParentShowHint(b._instance(), value)
}

// PopupMenu
//
// 获取右键菜单。
//
// Get Right click menu.
func (b *TBoundLabel) PopupMenu() *TPopupMenu {
	return AsPopupMenu(BoundLabel_GetPopupMenu(b._instance()))
}

// SetPopupMenu
//
// 设置右键菜单。
//
// Set Right click menu.
func (b *TBoundLabel) SetPopupMenu(value IComponent) {
	BoundLabel_SetPopupMenu(b._instance(), CheckPtr(value))
}

func (b *TBoundLabel) ShowAccelChar() bool {
	return BoundLabel_GetShowAccelChar(b._instance())
}

func (b *TBoundLabel) SetShowAccelChar(value bool) {
	BoundLabel_SetShowAccelChar(b._instance(), value)
}

// ShowHint
//
// 获取显示鼠标悬停提示。
//
// Get Show mouseover tips.
func (b *TBoundLabel) ShowHint() bool {
	return BoundLabel_GetShowHint(b._instance())
}

// SetShowHint
//
// 设置显示鼠标悬停提示。
//
// Set Show mouseover tips.
func (b *TBoundLabel) SetShowHint(value bool) {
	BoundLabel_SetShowHint(b._instance(), value)
}

// Top
//
// 获取顶边位置。
//
// Get Top position.
func (b *TBoundLabel) Top() int32 {
	return BoundLabel_GetTop(b._instance())
}

func (b *TBoundLabel) Layout() TTextLayout {
	return BoundLabel_GetLayout(b._instance())
}

func (b *TBoundLabel) SetLayout(value TTextLayout) {
	BoundLabel_SetLayout(b._instance(), value)
}

// WordWrap
//
// 获取自动换行。
//
// Get Automatic line break.
func (b *TBoundLabel) WordWrap() bool {
	return BoundLabel_GetWordWrap(b._instance())
}

// SetWordWrap
//
// 设置自动换行。
//
// Set Automatic line break.
func (b *TBoundLabel) SetWordWrap(value bool) {
	BoundLabel_SetWordWrap(b._instance(), value)
}

// Width
//
// 获取宽度。
//
// Get width.
func (b *TBoundLabel) Width() int32 {
	return BoundLabel_GetWidth(b._instance())
}

// SetWidth
//
// 设置宽度。
//
// Set width.
func (b *TBoundLabel) SetWidth(value int32) {
	BoundLabel_SetWidth(b._instance(), value)
}

// SetOnClick
//
// 设置控件单击事件。
//
// Set control click event.
func (b *TBoundLabel) SetOnClick(fn TNotifyEvent) {
	BoundLabel_SetOnClick(b._instance(), fn)
}

// SetOnDblClick
//
// 设置双击事件。
func (b *TBoundLabel) SetOnDblClick(fn TNotifyEvent) {
	BoundLabel_SetOnDblClick(b._instance(), fn)
}

// SetOnDragDrop
//
// 设置拖拽下落事件。
//
// Set Drag and drop event.
func (b *TBoundLabel) SetOnDragDrop(fn TDragDropEvent) {
	BoundLabel_SetOnDragDrop(b._instance(), fn)
}

// SetOnDragOver
//
// 设置拖拽完成事件。
//
// Set Drag and drop completion event.
func (b *TBoundLabel) SetOnDragOver(fn TDragOverEvent) {
	BoundLabel_SetOnDragOver(b._instance(), fn)
}

// SetOnEndDrag
//
// 设置拖拽结束。
//
// Set End of drag.
func (b *TBoundLabel) SetOnEndDrag(fn TEndDragEvent) {
	BoundLabel_SetOnEndDrag(b._instance(), fn)
}

// SetOnMouseDown
//
// 设置鼠标按下事件。
//
// Set Mouse down event.
func (b *TBoundLabel) SetOnMouseDown(fn TMouseEvent) {
	BoundLabel_SetOnMouseDown(b._instance(), fn)
}

// SetOnMouseMove
//
// 设置鼠标移动事件。
func (b *TBoundLabel) SetOnMouseMove(fn TMouseMoveEvent) {
	BoundLabel_SetOnMouseMove(b._instance(), fn)
}

// SetOnMouseUp
//
// 设置鼠标抬起事件。
//
// Set Mouse lift event.
func (b *TBoundLabel) SetOnMouseUp(fn TMouseEvent) {
	BoundLabel_SetOnMouseUp(b._instance(), fn)
}

// Canvas
//
// 获取画布。
func (b *TBoundLabel) Canvas() *TCanvas {
	return AsCanvas(BoundLabel_GetCanvas(b._instance()))
}

// Enabled
//
// 获取控件启用。
//
// Get the control enabled.
func (b *TBoundLabel) Enabled() bool {
	return BoundLabel_GetEnabled(b._instance())
}

// SetEnabled
//
// 设置控件启用。
//
// Set the control enabled.
func (b *TBoundLabel) SetEnabled(value bool) {
	BoundLabel_SetEnabled(b._instance(), value)
}

func (b *TBoundLabel) Action() *TAction {
	return AsAction(BoundLabel_GetAction(b._instance()))
}

func (b *TBoundLabel) SetAction(value IComponent) {
	BoundLabel_SetAction(b._instance(), CheckPtr(value))
}

// Align
//
// 获取控件自动调整。
//
// Get Control automatically adjusts.
func (b *TBoundLabel) Align() TAlign {
	return BoundLabel_GetAlign(b._instance())
}

// SetAlign
//
// 设置控件自动调整。
//
// Set Control automatically adjusts.
func (b *TBoundLabel) SetAlign(value TAlign) {
	BoundLabel_SetAlign(b._instance(), value)
}

// Anchors
//
// 获取四个角位置的锚点。
func (b *TBoundLabel) Anchors() TAnchors {
	return BoundLabel_GetAnchors(b._instance())
}

// SetAnchors
//
// 设置四个角位置的锚点。
func (b *TBoundLabel) SetAnchors(value TAnchors) {
	BoundLabel_SetAnchors(b._instance(), value)
}

func (b *TBoundLabel) BoundsRect() TRect {
	return BoundLabel_GetBoundsRect(b._instance())
}

func (b *TBoundLabel) SetBoundsRect(value TRect) {
	BoundLabel_SetBoundsRect(b._instance(), value)
}

// ClientHeight
//
// 获取客户区高度。
//
// Get client height.
func (b *TBoundLabel) ClientHeight() int32 {
	return BoundLabel_GetClientHeight(b._instance())
}

// SetClientHeight
//
// 设置客户区高度。
//
// Set client height.
func (b *TBoundLabel) SetClientHeight(value int32) {
	BoundLabel_SetClientHeight(b._instance(), value)
}

func (b *TBoundLabel) ClientOrigin() TPoint {
	return BoundLabel_GetClientOrigin(b._instance())
}

// ClientRect
//
// 获取客户区矩形。
//
// Get client rectangle.
func (b *TBoundLabel) ClientRect() TRect {
	return BoundLabel_GetClientRect(b._instance())
}

// ClientWidth
//
// 获取客户区宽度。
//
// Get client width.
func (b *TBoundLabel) ClientWidth() int32 {
	return BoundLabel_GetClientWidth(b._instance())
}

// SetClientWidth
//
// 设置客户区宽度。
//
// Set client width.
func (b *TBoundLabel) SetClientWidth(value int32) {
	BoundLabel_SetClientWidth(b._instance(), value)
}

// Constraints
//
// 获取约束控件大小。
func (b *TBoundLabel) Constraints() *TSizeConstraints {
	return AsSizeConstraints(BoundLabel_GetConstraints(b._instance()))
}

// SetConstraints
//
// 设置约束控件大小。
func (b *TBoundLabel) SetConstraints(value *TSizeConstraints) {
	BoundLabel_SetConstraints(b._instance(), CheckPtr(value))
}

// ControlState
//
// 获取控件状态。
//
// Get control state.
func (b *TBoundLabel) ControlState() TControlState {
	return BoundLabel_GetControlState(b._instance())
}

// SetControlState
//
// 设置控件状态。
//
// Set control state.
func (b *TBoundLabel) SetControlState(value TControlState) {
	BoundLabel_SetControlState(b._instance(), value)
}

// ControlStyle
//
// 获取控件样式。
//
// Get control style.
func (b *TBoundLabel) ControlStyle() TControlStyle {
	return BoundLabel_GetControlStyle(b._instance())
}

// SetControlStyle
//
// 设置控件样式。
//
// Set control style.
func (b *TBoundLabel) SetControlStyle(value TControlStyle) {
	BoundLabel_SetControlStyle(b._instance(), value)
}

func (b *TBoundLabel) Floating() bool {
	return BoundLabel_GetFloating(b._instance())
}

// Visible
//
// 获取控件可视。
//
// Get the control visible.
func (b *TBoundLabel) Visible() bool {
	return BoundLabel_GetVisible(b._instance())
}

// SetVisible
//
// 设置控件可视。
//
// Set the control visible.
func (b *TBoundLabel) SetVisible(value bool) {
	BoundLabel_SetVisible(b._instance(), value)
}

// Parent
//
// 获取控件父容器。
//
// Get control parent container.
func (b *TBoundLabel) Parent() *TWinControl {
	return AsWinControl(BoundLabel_GetParent(b._instance()))
}

// SetParent
//
// 设置控件父容器。
//
// Set control parent container.
func (b *TBoundLabel) SetParent(value IWinControl) {
	BoundLabel_SetParent(b._instance(), CheckPtr(value))
}

// Cursor
//
// 获取控件光标。
//
// Get control cursor.
func (b *TBoundLabel) Cursor() TCursor {
	return BoundLabel_GetCursor(b._instance())
}

// SetCursor
//
// 设置控件光标。
//
// Set control cursor.
func (b *TBoundLabel) SetCursor(value TCursor) {
	BoundLabel_SetCursor(b._instance(), value)
}

// Hint
//
// 获取组件鼠标悬停提示。
//
// Get component mouse hints.
func (b *TBoundLabel) Hint() string {
	return BoundLabel_GetHint(b._instance())
}

// SetHint
//
// 设置组件鼠标悬停提示。
//
// Set component mouse hints.
func (b *TBoundLabel) SetHint(value string) {
	BoundLabel_SetHint(b._instance(), value)
}

// ComponentCount
//
// 获取组件总数。
//
// Get the total number of components.
func (b *TBoundLabel) ComponentCount() int32 {
	return BoundLabel_GetComponentCount(b._instance())
}

// ComponentIndex
//
// 获取组件索引。
//
// Get component index.
func (b *TBoundLabel) ComponentIndex() int32 {
	return BoundLabel_GetComponentIndex(b._instance())
}

// SetComponentIndex
//
// 设置组件索引。
//
// Set component index.
func (b *TBoundLabel) SetComponentIndex(value int32) {
	BoundLabel_SetComponentIndex(b._instance(), value)
}

// Owner
//
// 获取组件所有者。
//
// Get component owner.
func (b *TBoundLabel) Owner() *TComponent {
	return AsComponent(BoundLabel_GetOwner(b._instance()))
}

// Name
//
// 获取组件名称。
//
// Get the component name.
func (b *TBoundLabel) Name() string {
	return BoundLabel_GetName(b._instance())
}

// SetName
//
// 设置组件名称。
//
// Set the component name.
func (b *TBoundLabel) SetName(value string) {
	BoundLabel_SetName(b._instance(), value)
}

// Tag
//
// 获取对象标记。
//
// Get the control tag.
func (b *TBoundLabel) Tag() int {
	return BoundLabel_GetTag(b._instance())
}

// SetTag
//
// 设置对象标记。
//
// Set the control tag.
func (b *TBoundLabel) SetTag(value int) {
	BoundLabel_SetTag(b._instance(), value)
}

// AnchorSideLeft
//
// 获取左边锚点。
func (b *TBoundLabel) AnchorSideLeft() *TAnchorSide {
	return AsAnchorSide(BoundLabel_GetAnchorSideLeft(b._instance()))
}

// SetAnchorSideLeft
//
// 设置左边锚点。
func (b *TBoundLabel) SetAnchorSideLeft(value *TAnchorSide) {
	BoundLabel_SetAnchorSideLeft(b._instance(), CheckPtr(value))
}

// AnchorSideTop
//
// 获取顶边锚点。
func (b *TBoundLabel) AnchorSideTop() *TAnchorSide {
	return AsAnchorSide(BoundLabel_GetAnchorSideTop(b._instance()))
}

// SetAnchorSideTop
//
// 设置顶边锚点。
func (b *TBoundLabel) SetAnchorSideTop(value *TAnchorSide) {
	BoundLabel_SetAnchorSideTop(b._instance(), CheckPtr(value))
}

// AnchorSideRight
//
// 获取右边锚点。
func (b *TBoundLabel) AnchorSideRight() *TAnchorSide {
	return AsAnchorSide(BoundLabel_GetAnchorSideRight(b._instance()))
}

// SetAnchorSideRight
//
// 设置右边锚点。
func (b *TBoundLabel) SetAnchorSideRight(value *TAnchorSide) {
	BoundLabel_SetAnchorSideRight(b._instance(), CheckPtr(value))
}

// AnchorSideBottom
//
// 获取底边锚点。
func (b *TBoundLabel) AnchorSideBottom() *TAnchorSide {
	return AsAnchorSide(BoundLabel_GetAnchorSideBottom(b._instance()))
}

// SetAnchorSideBottom
//
// 设置底边锚点。
func (b *TBoundLabel) SetAnchorSideBottom(value *TAnchorSide) {
	BoundLabel_SetAnchorSideBottom(b._instance(), CheckPtr(value))
}

// BorderSpacing
//
// 获取边框间距。
func (b *TBoundLabel) BorderSpacing() *TControlBorderSpacing {
	return AsControlBorderSpacing(BoundLabel_GetBorderSpacing(b._instance()))
}

// SetBorderSpacing
//
// 设置边框间距。
func (b *TBoundLabel) SetBorderSpacing(value *TControlBorderSpacing) {
	BoundLabel_SetBorderSpacing(b._instance(), CheckPtr(value))
}

// Components
//
// 获取指定索引组件。
//
// Get the specified index component.
func (b *TBoundLabel) Components(AIndex int32) *TComponent {
	return AsComponent(BoundLabel_GetComponents(b._instance(), AIndex))
}

// AnchorSide
//
// 获取锚侧面。
func (b *TBoundLabel) AnchorSide(AKind TAnchorKind) *TAnchorSide {
	return AsAnchorSide(BoundLabel_GetAnchorSide(b._instance(), AKind))
}
