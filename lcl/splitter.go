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

type TSplitter struct {
	IControl
	instance unsafe.Pointer
}

// NewSplitter
//
// 创建一个新的对象。
//
// Create a new object.
func NewSplitter(owner IComponent) *TSplitter {
	s := new(TSplitter)
	s.instance = unsafe.Pointer(Splitter_Create(CheckPtr(owner)))
	return s
}

// AsSplitter
//
// 动态转换一个已存在的对象实例。
//
// Dynamically convert an existing object instance.
func AsSplitter(obj interface{}) *TSplitter {
	instance := getInstance(obj)
	if instance == nullptr {
		return nil
	}
	return &TSplitter{instance: instance}
}

// Free
//
// 释放对象。
//
// Free object.
func (s *TSplitter) Free() {
	if s.instance != nullptr {
		Splitter_Free(s._instance())
		s.instance = nullptr
	}
}

func (s *TSplitter) _instance() uintptr {
	return uintptr(s.instance)
}

// Instance
//
// 返回对象实例指针。
//
// Return object instance pointer.
func (s *TSplitter) Instance() uintptr {
	return s._instance()
}

// UnsafeAddr
//
// 获取一个不安全的地址。
//
// Get an unsafe address.
func (s *TSplitter) UnsafeAddr() unsafe.Pointer {
	return s.instance
}

// IsValid
//
// 检测地址是否为空。
//
// Check if the address is empty.
func (s *TSplitter) IsValid() bool {
	return s.instance != nullptr
}

// Is
//
// 检测当前对象是否继承自目标对象。
//
// Checks whether the current object is inherited from the target object.
func (s *TSplitter) Is() TIs {
	return TIs(s._instance())
}

// As
//
// 动态转换当前对象为目标对象。
//
// Dynamically convert the current object to the target object.
//func (s *TSplitter) As() TAs {
//    return TAs(s._instance())
//}

// TSplitterClass
//
// 获取类信息指针。
//
// Get class information pointer.
func TSplitterClass() TClass {
	return Splitter_StaticClassType()
}

// BringToFront
//
// 将控件置于最前。
//
// Bring the control to the front.
func (s *TSplitter) BringToFront() {
	Splitter_BringToFront(s._instance())
}

// ClientToScreen
//
// 将客户端坐标转为绝对的屏幕坐标。
//
// Convert client coordinates to absolute screen coordinates.
func (s *TSplitter) ClientToScreen(Point TPoint) TPoint {
	return Splitter_ClientToScreen(s._instance(), Point)
}

// ClientToParent
//
// 将客户端坐标转为父容器坐标。
//
// Convert client coordinates to parent container coordinates.
func (s *TSplitter) ClientToParent(Point TPoint, AParent IWinControl) TPoint {
	return Splitter_ClientToParent(s._instance(), Point, CheckPtr(AParent))
}

// Dragging
//
// 是否在拖拽中。
//
// Is it in the middle of dragging.
func (s *TSplitter) Dragging() bool {
	return Splitter_Dragging(s._instance())
}

// HasParent
//
// 是否有父容器。
//
// Is there a parent container.
func (s *TSplitter) HasParent() bool {
	return Splitter_HasParent(s._instance())
}

// Hide
//
// 隐藏控件。
//
// Hidden control.
func (s *TSplitter) Hide() {
	Splitter_Hide(s._instance())
}

// Invalidate
//
// 要求重绘。
//
// Redraw.
func (s *TSplitter) Invalidate() {
	Splitter_Invalidate(s._instance())
}

// Perform
//
// 发送一个消息。
//
// Send a message.
func (s *TSplitter) Perform(Msg uint32, WParam uintptr, LParam int) int {
	return Splitter_Perform(s._instance(), Msg, WParam, LParam)
}

// Refresh
//
// 刷新控件。
//
// Refresh control.
func (s *TSplitter) Refresh() {
	Splitter_Refresh(s._instance())
}

// Repaint
//
// 重绘。
//
// Repaint.
func (s *TSplitter) Repaint() {
	Splitter_Repaint(s._instance())
}

// ScreenToClient
//
// 将屏幕坐标转为客户端坐标。
//
// Convert screen coordinates to client coordinates.
func (s *TSplitter) ScreenToClient(Point TPoint) TPoint {
	return Splitter_ScreenToClient(s._instance(), Point)
}

// ParentToClient
//
// 将父容器坐标转为客户端坐标。
//
// Convert parent container coordinates to client coordinates.
func (s *TSplitter) ParentToClient(Point TPoint, AParent IWinControl) TPoint {
	return Splitter_ParentToClient(s._instance(), Point, CheckPtr(AParent))
}

// SendToBack
//
// 控件至于最后面。
//
// The control is placed at the end.
func (s *TSplitter) SendToBack() {
	Splitter_SendToBack(s._instance())
}

// SetBounds
//
// 设置组件边界。
//
// Set component boundaries.
func (s *TSplitter) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
	Splitter_SetBounds(s._instance(), ALeft, ATop, AWidth, AHeight)
}

// Show
//
// 显示控件。
//
// Show control.
func (s *TSplitter) Show() {
	Splitter_Show(s._instance())
}

// Update
//
// 控件更新。
//
// Update.
func (s *TSplitter) Update() {
	Splitter_Update(s._instance())
}

// GetTextBuf
//
// 获取控件的字符，如果有。
//
// Get the characters of the control, if any.
func (s *TSplitter) GetTextBuf(Buffer *string, BufSize int32) int32 {
	return Splitter_GetTextBuf(s._instance(), Buffer, BufSize)
}

// GetTextLen
//
// 获取控件的字符长，如果有。
//
// Get the character length of the control, if any.
func (s *TSplitter) GetTextLen() int32 {
	return Splitter_GetTextLen(s._instance())
}

// SetTextBuf
//
// 设置控件字符，如果有。
//
// Set control characters, if any.
func (s *TSplitter) SetTextBuf(Buffer string) {
	Splitter_SetTextBuf(s._instance(), Buffer)
}

// FindComponent
//
// 查找指定名称的组件。
//
// Find the component with the specified name.
func (s *TSplitter) FindComponent(AName string) *TComponent {
	return AsComponent(Splitter_FindComponent(s._instance(), AName))
}

// GetNamePath
//
// 获取类名路径。
//
// Get the class name path.
func (s *TSplitter) GetNamePath() string {
	return Splitter_GetNamePath(s._instance())
}

// Assign
//
// 复制一个对象，如果对象实现了此方法的话。
//
// Copy an object, if the object implements this method.
func (s *TSplitter) Assign(Source IObject) {
	Splitter_Assign(s._instance(), CheckPtr(Source))
}

// ClassType
//
// 获取类的类型信息。
//
// Get class type information.
func (s *TSplitter) ClassType() TClass {
	return Splitter_ClassType(s._instance())
}

// ClassName
//
// 获取当前对象类名称。
//
// Get the current object class name.
func (s *TSplitter) ClassName() string {
	return Splitter_ClassName(s._instance())
}

// InstanceSize
//
// 获取当前对象实例大小。
//
// Get the current object instance size.
func (s *TSplitter) InstanceSize() int32 {
	return Splitter_InstanceSize(s._instance())
}

// InheritsFrom
//
// 判断当前类是否继承自指定类。
//
// Determine whether the current class inherits from the specified class.
func (s *TSplitter) InheritsFrom(AClass TClass) bool {
	return Splitter_InheritsFrom(s._instance(), AClass)
}

// Equals
//
// 与一个对象进行比较。
//
// Compare with an object.
func (s *TSplitter) Equals(Obj IObject) bool {
	return Splitter_Equals(s._instance(), CheckPtr(Obj))
}

// GetHashCode
//
// 获取类的哈希值。
//
// Get the hash value of the class.
func (s *TSplitter) GetHashCode() int32 {
	return Splitter_GetHashCode(s._instance())
}

// ToString
//
// 文本类信息。
//
// Text information.
func (s *TSplitter) ToString() string {
	return Splitter_ToString(s._instance())
}

func (s *TSplitter) AnchorToNeighbour(ASide TAnchorKind, ASpace int32, ASibling IControl) {
	Splitter_AnchorToNeighbour(s._instance(), ASide, ASpace, CheckPtr(ASibling))
}

func (s *TSplitter) AnchorParallel(ASide TAnchorKind, ASpace int32, ASibling IControl) {
	Splitter_AnchorParallel(s._instance(), ASide, ASpace, CheckPtr(ASibling))
}

// AnchorHorizontalCenterTo
//
// 置于指定控件的横向中心。
func (s *TSplitter) AnchorHorizontalCenterTo(ASibling IControl) {
	Splitter_AnchorHorizontalCenterTo(s._instance(), CheckPtr(ASibling))
}

// AnchorVerticalCenterTo
//
// 置于指定控件的纵向中心。
func (s *TSplitter) AnchorVerticalCenterTo(ASibling IControl) {
	Splitter_AnchorVerticalCenterTo(s._instance(), CheckPtr(ASibling))
}

func (s *TSplitter) AnchorSame(ASide TAnchorKind, ASibling IControl) {
	Splitter_AnchorSame(s._instance(), ASide, CheckPtr(ASibling))
}

func (s *TSplitter) AnchorAsAlign(ATheAlign TAlign, ASpace int32) {
	Splitter_AnchorAsAlign(s._instance(), ATheAlign, ASpace)
}

func (s *TSplitter) AnchorClient(ASpace int32) {
	Splitter_AnchorClient(s._instance(), ASpace)
}

func (s *TSplitter) ScaleDesignToForm(ASize int32) int32 {
	return Splitter_ScaleDesignToForm(s._instance(), ASize)
}

func (s *TSplitter) ScaleFormToDesign(ASize int32) int32 {
	return Splitter_ScaleFormToDesign(s._instance(), ASize)
}

func (s *TSplitter) Scale96ToForm(ASize int32) int32 {
	return Splitter_Scale96ToForm(s._instance(), ASize)
}

func (s *TSplitter) ScaleFormTo96(ASize int32) int32 {
	return Splitter_ScaleFormTo96(s._instance(), ASize)
}

func (s *TSplitter) Scale96ToFont(ASize int32) int32 {
	return Splitter_Scale96ToFont(s._instance(), ASize)
}

func (s *TSplitter) ScaleFontTo96(ASize int32) int32 {
	return Splitter_ScaleFontTo96(s._instance(), ASize)
}

func (s *TSplitter) ScaleScreenToFont(ASize int32) int32 {
	return Splitter_ScaleScreenToFont(s._instance(), ASize)
}

func (s *TSplitter) ScaleFontToScreen(ASize int32) int32 {
	return Splitter_ScaleFontToScreen(s._instance(), ASize)
}

func (s *TSplitter) Scale96ToScreen(ASize int32) int32 {
	return Splitter_Scale96ToScreen(s._instance(), ASize)
}

func (s *TSplitter) ScaleScreenTo96(ASize int32) int32 {
	return Splitter_ScaleScreenTo96(s._instance(), ASize)
}

func (s *TSplitter) AutoAdjustLayout(AMode TLayoutAdjustmentPolicy, AFromPPI int32, AToPPI int32, AOldFormWidth int32, ANewFormWidth int32) {
	Splitter_AutoAdjustLayout(s._instance(), AMode, AFromPPI, AToPPI, AOldFormWidth, ANewFormWidth)
}

func (s *TSplitter) FixDesignFontsPPI(ADesignTimePPI int32) {
	Splitter_FixDesignFontsPPI(s._instance(), ADesignTimePPI)
}

func (s *TSplitter) ScaleFontsPPI(AToPPI int32, AProportion float64) {
	Splitter_ScaleFontsPPI(s._instance(), AToPPI, AProportion)
}

func (s *TSplitter) ResizeAnchor() TAnchorKind {
	return Splitter_GetResizeAnchor(s._instance())
}

func (s *TSplitter) SetResizeAnchor(value TAnchorKind) {
	Splitter_SetResizeAnchor(s._instance(), value)
}

// Canvas
//
// 获取画布。
func (s *TSplitter) Canvas() *TCanvas {
	return AsCanvas(Splitter_GetCanvas(s._instance()))
}

// Align
//
// 获取控件自动调整。
//
// Get Control automatically adjusts.
func (s *TSplitter) Align() TAlign {
	return Splitter_GetAlign(s._instance())
}

// SetAlign
//
// 设置控件自动调整。
//
// Set Control automatically adjusts.
func (s *TSplitter) SetAlign(value TAlign) {
	Splitter_SetAlign(s._instance(), value)
}

// Color
//
// 获取颜色。
//
// Get color.
func (s *TSplitter) Color() TColor {
	return Splitter_GetColor(s._instance())
}

// SetColor
//
// 设置颜色。
//
// Set color.
func (s *TSplitter) SetColor(value TColor) {
	Splitter_SetColor(s._instance(), value)
}

// Cursor
//
// 获取控件光标。
//
// Get control cursor.
func (s *TSplitter) Cursor() TCursor {
	return Splitter_GetCursor(s._instance())
}

// SetCursor
//
// 设置控件光标。
//
// Set control cursor.
func (s *TSplitter) SetCursor(value TCursor) {
	Splitter_SetCursor(s._instance(), value)
}

// Constraints
//
// 获取约束控件大小。
func (s *TSplitter) Constraints() *TSizeConstraints {
	return AsSizeConstraints(Splitter_GetConstraints(s._instance()))
}

// SetConstraints
//
// 设置约束控件大小。
func (s *TSplitter) SetConstraints(value *TSizeConstraints) {
	Splitter_SetConstraints(s._instance(), CheckPtr(value))
}

func (s *TSplitter) MinSize() int32 {
	return Splitter_GetMinSize(s._instance())
}

func (s *TSplitter) SetMinSize(value int32) {
	Splitter_SetMinSize(s._instance(), value)
}

// ParentColor
//
// 获取使用父容器颜色。
//
// Get parent color.
func (s *TSplitter) ParentColor() bool {
	return Splitter_GetParentColor(s._instance())
}

// SetParentColor
//
// 设置使用父容器颜色。
//
// Set parent color.
func (s *TSplitter) SetParentColor(value bool) {
	Splitter_SetParentColor(s._instance(), value)
}

// Visible
//
// 获取控件可视。
//
// Get the control visible.
func (s *TSplitter) Visible() bool {
	return Splitter_GetVisible(s._instance())
}

// SetVisible
//
// 设置控件可视。
//
// Set the control visible.
func (s *TSplitter) SetVisible(value bool) {
	Splitter_SetVisible(s._instance(), value)
}

// Width
//
// 获取宽度。
//
// Get width.
func (s *TSplitter) Width() int32 {
	return Splitter_GetWidth(s._instance())
}

// SetWidth
//
// 设置宽度。
//
// Set width.
func (s *TSplitter) SetWidth(value int32) {
	Splitter_SetWidth(s._instance(), value)
}

// SetOnPaint
//
// 设置绘画事件。
func (s *TSplitter) SetOnPaint(fn TNotifyEvent) {
	Splitter_SetOnPaint(s._instance(), fn)
}

// Enabled
//
// 获取控件启用。
//
// Get the control enabled.
func (s *TSplitter) Enabled() bool {
	return Splitter_GetEnabled(s._instance())
}

// SetEnabled
//
// 设置控件启用。
//
// Set the control enabled.
func (s *TSplitter) SetEnabled(value bool) {
	Splitter_SetEnabled(s._instance(), value)
}

func (s *TSplitter) Action() *TAction {
	return AsAction(Splitter_GetAction(s._instance()))
}

func (s *TSplitter) SetAction(value IComponent) {
	Splitter_SetAction(s._instance(), CheckPtr(value))
}

// Anchors
//
// 获取四个角位置的锚点。
func (s *TSplitter) Anchors() TAnchors {
	return Splitter_GetAnchors(s._instance())
}

// SetAnchors
//
// 设置四个角位置的锚点。
func (s *TSplitter) SetAnchors(value TAnchors) {
	Splitter_SetAnchors(s._instance(), value)
}

func (s *TSplitter) BiDiMode() TBiDiMode {
	return Splitter_GetBiDiMode(s._instance())
}

func (s *TSplitter) SetBiDiMode(value TBiDiMode) {
	Splitter_SetBiDiMode(s._instance(), value)
}

func (s *TSplitter) BoundsRect() TRect {
	return Splitter_GetBoundsRect(s._instance())
}

func (s *TSplitter) SetBoundsRect(value TRect) {
	Splitter_SetBoundsRect(s._instance(), value)
}

// ClientHeight
//
// 获取客户区高度。
//
// Get client height.
func (s *TSplitter) ClientHeight() int32 {
	return Splitter_GetClientHeight(s._instance())
}

// SetClientHeight
//
// 设置客户区高度。
//
// Set client height.
func (s *TSplitter) SetClientHeight(value int32) {
	Splitter_SetClientHeight(s._instance(), value)
}

func (s *TSplitter) ClientOrigin() TPoint {
	return Splitter_GetClientOrigin(s._instance())
}

// ClientRect
//
// 获取客户区矩形。
//
// Get client rectangle.
func (s *TSplitter) ClientRect() TRect {
	return Splitter_GetClientRect(s._instance())
}

// ClientWidth
//
// 获取客户区宽度。
//
// Get client width.
func (s *TSplitter) ClientWidth() int32 {
	return Splitter_GetClientWidth(s._instance())
}

// SetClientWidth
//
// 设置客户区宽度。
//
// Set client width.
func (s *TSplitter) SetClientWidth(value int32) {
	Splitter_SetClientWidth(s._instance(), value)
}

// ControlState
//
// 获取控件状态。
//
// Get control state.
func (s *TSplitter) ControlState() TControlState {
	return Splitter_GetControlState(s._instance())
}

// SetControlState
//
// 设置控件状态。
//
// Set control state.
func (s *TSplitter) SetControlState(value TControlState) {
	Splitter_SetControlState(s._instance(), value)
}

// ControlStyle
//
// 获取控件样式。
//
// Get control style.
func (s *TSplitter) ControlStyle() TControlStyle {
	return Splitter_GetControlStyle(s._instance())
}

// SetControlStyle
//
// 设置控件样式。
//
// Set control style.
func (s *TSplitter) SetControlStyle(value TControlStyle) {
	Splitter_SetControlStyle(s._instance(), value)
}

func (s *TSplitter) Floating() bool {
	return Splitter_GetFloating(s._instance())
}

// ShowHint
//
// 获取显示鼠标悬停提示。
//
// Get Show mouseover tips.
func (s *TSplitter) ShowHint() bool {
	return Splitter_GetShowHint(s._instance())
}

// SetShowHint
//
// 设置显示鼠标悬停提示。
//
// Set Show mouseover tips.
func (s *TSplitter) SetShowHint(value bool) {
	Splitter_SetShowHint(s._instance(), value)
}

// Parent
//
// 获取控件父容器。
//
// Get control parent container.
func (s *TSplitter) Parent() *TWinControl {
	return AsWinControl(Splitter_GetParent(s._instance()))
}

// SetParent
//
// 设置控件父容器。
//
// Set control parent container.
func (s *TSplitter) SetParent(value IWinControl) {
	Splitter_SetParent(s._instance(), CheckPtr(value))
}

// Left
//
// 获取左边位置。
//
// Get Left position.
func (s *TSplitter) Left() int32 {
	return Splitter_GetLeft(s._instance())
}

// SetLeft
//
// 设置左边位置。
//
// Set Left position.
func (s *TSplitter) SetLeft(value int32) {
	Splitter_SetLeft(s._instance(), value)
}

// Top
//
// 获取顶边位置。
//
// Get Top position.
func (s *TSplitter) Top() int32 {
	return Splitter_GetTop(s._instance())
}

// SetTop
//
// 设置顶边位置。
//
// Set Top position.
func (s *TSplitter) SetTop(value int32) {
	Splitter_SetTop(s._instance(), value)
}

// Height
//
// 获取高度。
//
// Get height.
func (s *TSplitter) Height() int32 {
	return Splitter_GetHeight(s._instance())
}

// SetHeight
//
// 设置高度。
//
// Set height.
func (s *TSplitter) SetHeight(value int32) {
	Splitter_SetHeight(s._instance(), value)
}

// Hint
//
// 获取组件鼠标悬停提示。
//
// Get component mouse hints.
func (s *TSplitter) Hint() string {
	return Splitter_GetHint(s._instance())
}

// SetHint
//
// 设置组件鼠标悬停提示。
//
// Set component mouse hints.
func (s *TSplitter) SetHint(value string) {
	Splitter_SetHint(s._instance(), value)
}

// ComponentCount
//
// 获取组件总数。
//
// Get the total number of components.
func (s *TSplitter) ComponentCount() int32 {
	return Splitter_GetComponentCount(s._instance())
}

// ComponentIndex
//
// 获取组件索引。
//
// Get component index.
func (s *TSplitter) ComponentIndex() int32 {
	return Splitter_GetComponentIndex(s._instance())
}

// SetComponentIndex
//
// 设置组件索引。
//
// Set component index.
func (s *TSplitter) SetComponentIndex(value int32) {
	Splitter_SetComponentIndex(s._instance(), value)
}

// Owner
//
// 获取组件所有者。
//
// Get component owner.
func (s *TSplitter) Owner() *TComponent {
	return AsComponent(Splitter_GetOwner(s._instance()))
}

// Name
//
// 获取组件名称。
//
// Get the component name.
func (s *TSplitter) Name() string {
	return Splitter_GetName(s._instance())
}

// SetName
//
// 设置组件名称。
//
// Set the component name.
func (s *TSplitter) SetName(value string) {
	Splitter_SetName(s._instance(), value)
}

// Tag
//
// 获取对象标记。
//
// Get the control tag.
func (s *TSplitter) Tag() int {
	return Splitter_GetTag(s._instance())
}

// SetTag
//
// 设置对象标记。
//
// Set the control tag.
func (s *TSplitter) SetTag(value int) {
	Splitter_SetTag(s._instance(), value)
}

// AnchorSideLeft
//
// 获取左边锚点。
func (s *TSplitter) AnchorSideLeft() *TAnchorSide {
	return AsAnchorSide(Splitter_GetAnchorSideLeft(s._instance()))
}

// SetAnchorSideLeft
//
// 设置左边锚点。
func (s *TSplitter) SetAnchorSideLeft(value *TAnchorSide) {
	Splitter_SetAnchorSideLeft(s._instance(), CheckPtr(value))
}

// AnchorSideTop
//
// 获取顶边锚点。
func (s *TSplitter) AnchorSideTop() *TAnchorSide {
	return AsAnchorSide(Splitter_GetAnchorSideTop(s._instance()))
}

// SetAnchorSideTop
//
// 设置顶边锚点。
func (s *TSplitter) SetAnchorSideTop(value *TAnchorSide) {
	Splitter_SetAnchorSideTop(s._instance(), CheckPtr(value))
}

// AnchorSideRight
//
// 获取右边锚点。
func (s *TSplitter) AnchorSideRight() *TAnchorSide {
	return AsAnchorSide(Splitter_GetAnchorSideRight(s._instance()))
}

// SetAnchorSideRight
//
// 设置右边锚点。
func (s *TSplitter) SetAnchorSideRight(value *TAnchorSide) {
	Splitter_SetAnchorSideRight(s._instance(), CheckPtr(value))
}

// AnchorSideBottom
//
// 获取底边锚点。
func (s *TSplitter) AnchorSideBottom() *TAnchorSide {
	return AsAnchorSide(Splitter_GetAnchorSideBottom(s._instance()))
}

// SetAnchorSideBottom
//
// 设置底边锚点。
func (s *TSplitter) SetAnchorSideBottom(value *TAnchorSide) {
	Splitter_SetAnchorSideBottom(s._instance(), CheckPtr(value))
}

// BorderSpacing
//
// 获取边框间距。
func (s *TSplitter) BorderSpacing() *TControlBorderSpacing {
	return AsControlBorderSpacing(Splitter_GetBorderSpacing(s._instance()))
}

// SetBorderSpacing
//
// 设置边框间距。
func (s *TSplitter) SetBorderSpacing(value *TControlBorderSpacing) {
	Splitter_SetBorderSpacing(s._instance(), CheckPtr(value))
}

// Components
//
// 获取指定索引组件。
//
// Get the specified index component.
func (s *TSplitter) Components(AIndex int32) *TComponent {
	return AsComponent(Splitter_GetComponents(s._instance(), AIndex))
}

// AnchorSide
//
// 获取锚侧面。
func (s *TSplitter) AnchorSide(AKind TAnchorKind) *TAnchorSide {
	return AsAnchorSide(Splitter_GetAnchorSide(s._instance(), AKind))
}
