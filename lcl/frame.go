//----------------------------------------
// The code is automatically generated by the GenlibLcl tool.
// Copyright © sxm. All Rights Reserved.
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

type TFrame struct {
	IWinControl
	instance uintptr
	// 特殊情况下使用，主要应对Go的GC问题，与LCL没有太多关系。
	ptr unsafe.Pointer
}

// 创建一个新的对象。
//
// Create a new object.
func NewFrame(owner IComponent) *TFrame {
	f := new(TFrame)
	f.instance = Frame_Create(CheckPtr(owner))
	f.ptr = unsafe.Pointer(f.instance)
	return f
}

// 动态转换一个已存在的对象实例。
//
// Dynamically convert an existing object instance.
func AsFrame(obj interface{}) *TFrame {
	instance, ptr := getInstance(obj)
	if instance == 0 {
		return nil
	}
	return &TFrame{instance: instance, ptr: ptr}
}

// -------------------------- Deprecated begin --------------------------
// 新建一个对象来自已经存在的对象实例指针。
//
// Create a new object from an existing object instance pointer.
// Deprecated: use AsFrame.
func FrameFromInst(inst uintptr) *TFrame {
	return AsFrame(inst)
}

// 新建一个对象来自已经存在的对象实例。
//
// Create a new object from an existing object instance.
// Deprecated: use AsFrame.
func FrameFromObj(obj IObject) *TFrame {
	return AsFrame(obj)
}

// 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
//
// Create a new object from an unsecured address. Note: Using this function may cause some unclear situations and be used with caution..
// Deprecated: use AsFrame.
func FrameFromUnsafePointer(ptr unsafe.Pointer) *TFrame {
	return AsFrame(ptr)
}

// -------------------------- Deprecated end --------------------------
// 释放对象。
//
// Free object.
func (f *TFrame) Free() {
	if f.instance != 0 {
		Frame_Free(f.instance)
		f.instance, f.ptr = 0, nullptr
	}
}

// 返回对象实例指针。
//
// Return object instance pointer.
func (f *TFrame) Instance() uintptr {
	return f.instance
}

// 获取一个不安全的地址。
//
// Get an unsafe address.
func (f *TFrame) UnsafeAddr() unsafe.Pointer {
	return f.ptr
}

// 检测地址是否为空。
//
// Check if the address is empty.
func (f *TFrame) IsValid() bool {
	return f.instance != 0
}

// 检测当前对象是否继承自目标对象。
//
// Checks whether the current object is inherited from the target object.
func (f *TFrame) Is() TIs {
	return TIs(f.instance)
}

// 动态转换当前对象为目标对象。
//
// Dynamically convert the current object to the target object.
//func (f *TFrame) As() TAs {
//    return TAs(f.instance)
//}

// 获取类信息指针。
//
// Get class information pointer.
func TFrameClass() TClass {
	return Frame_StaticClassType()
}

func (f *TFrame) ScrollInView(AControl IControl) {
	Frame_ScrollInView(f.instance, CheckPtr(AControl))
}

// 是否可以获得焦点。
func (f *TFrame) CanFocus() bool {
	return Frame_CanFocus(f.instance)
}

// 返回是否包含指定控件。
//
// it's contain a specified control.
func (f *TFrame) ContainsControl(Control IControl) bool {
	return Frame_ContainsControl(f.instance, CheckPtr(Control))
}

// 返回指定坐标及相关属性位置控件。
//
// Returns the specified coordinate and the relevant attribute position control..
func (f *TFrame) ControlAtPos(Pos TPoint, AllowDisabled bool, AllowWinControls bool, AllLevels bool) *TControl {
	return AsControl(Frame_ControlAtPos(f.instance, Pos, AllowDisabled, AllowWinControls, AllLevels))
}

// 禁用控件的对齐。
//
// Disable control alignment.
func (f *TFrame) DisableAlign() {
	Frame_DisableAlign(f.instance)
}

// 启用控件对齐。
//
// Enabled control alignment.
func (f *TFrame) EnableAlign() {
	Frame_EnableAlign(f.instance)
}

// 查找子控件。
//
// Find sub controls.
func (f *TFrame) FindChildControl(ControlName string) *TControl {
	return AsControl(Frame_FindChildControl(f.instance, ControlName))
}

func (f *TFrame) FlipChildren(AllLevels bool) {
	Frame_FlipChildren(f.instance, AllLevels)
}

// 返回是否获取焦点。
//
// Return to get focus.
func (f *TFrame) Focused() bool {
	return Frame_Focused(f.instance)
}

// 句柄是否已经分配。
//
// Is the handle already allocated.
func (f *TFrame) HandleAllocated() bool {
	return Frame_HandleAllocated(f.instance)
}

// 插入一个控件。
//
// Insert a control.
func (f *TFrame) InsertControl(AControl IControl) {
	Frame_InsertControl(f.instance, CheckPtr(AControl))
}

// 要求重绘。
//
// Redraw.
func (f *TFrame) Invalidate() {
	Frame_Invalidate(f.instance)
}

// 绘画至指定DC。
//
// Painting to the specified DC.
func (f *TFrame) PaintTo(DC HDC, X int32, Y int32) {
	Frame_PaintTo(f.instance, DC, X, Y)
}

// 移除一个控件。
//
// Remove a control.
func (f *TFrame) RemoveControl(AControl IControl) {
	Frame_RemoveControl(f.instance, CheckPtr(AControl))
}

// 重新对齐。
//
// Realign.
func (f *TFrame) Realign() {
	Frame_Realign(f.instance)
}

// 重绘。
//
// Repaint.
func (f *TFrame) Repaint() {
	Frame_Repaint(f.instance)
}

// 按比例缩放。
//
// Scale by.
func (f *TFrame) ScaleBy(M int32, D int32) {
	Frame_ScaleBy(f.instance, M, D)
}

// 滚动至指定位置。
//
// Scroll by.
func (f *TFrame) ScrollBy(DeltaX int32, DeltaY int32) {
	Frame_ScrollBy(f.instance, DeltaX, DeltaY)
}

// 设置组件边界。
//
// Set component boundaries.
func (f *TFrame) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
	Frame_SetBounds(f.instance, ALeft, ATop, AWidth, AHeight)
}

// 设置控件焦点。
//
// Set control focus.
func (f *TFrame) SetFocus() {
	Frame_SetFocus(f.instance)
}

// 控件更新。
//
// Update.
func (f *TFrame) Update() {
	Frame_Update(f.instance)
}

// 将控件置于最前。
//
// Bring the control to the front.
func (f *TFrame) BringToFront() {
	Frame_BringToFront(f.instance)
}

// 将客户端坐标转为绝对的屏幕坐标。
//
// Convert client coordinates to absolute screen coordinates.
func (f *TFrame) ClientToScreen(Point TPoint) TPoint {
	return Frame_ClientToScreen(f.instance, Point)
}

// 将客户端坐标转为父容器坐标。
//
// Convert client coordinates to parent container coordinates.
func (f *TFrame) ClientToParent(Point TPoint, AParent IWinControl) TPoint {
	return Frame_ClientToParent(f.instance, Point, CheckPtr(AParent))
}

// 是否在拖拽中。
//
// Is it in the middle of dragging.
func (f *TFrame) Dragging() bool {
	return Frame_Dragging(f.instance)
}

// 是否有父容器。
//
// Is there a parent container.
func (f *TFrame) HasParent() bool {
	return Frame_HasParent(f.instance)
}

// 隐藏控件。
//
// Hidden control.
func (f *TFrame) Hide() {
	Frame_Hide(f.instance)
}

// 发送一个消息。
//
// Send a message.
func (f *TFrame) Perform(Msg uint32, WParam uintptr, LParam int) int {
	return Frame_Perform(f.instance, Msg, WParam, LParam)
}

// 刷新控件。
//
// Refresh control.
func (f *TFrame) Refresh() {
	Frame_Refresh(f.instance)
}

// 将屏幕坐标转为客户端坐标。
//
// Convert screen coordinates to client coordinates.
func (f *TFrame) ScreenToClient(Point TPoint) TPoint {
	return Frame_ScreenToClient(f.instance, Point)
}

// 将父容器坐标转为客户端坐标。
//
// Convert parent container coordinates to client coordinates.
func (f *TFrame) ParentToClient(Point TPoint, AParent IWinControl) TPoint {
	return Frame_ParentToClient(f.instance, Point, CheckPtr(AParent))
}

// 控件至于最后面。
//
// The control is placed at the end.
func (f *TFrame) SendToBack() {
	Frame_SendToBack(f.instance)
}

// 显示控件。
//
// Show control.
func (f *TFrame) Show() {
	Frame_Show(f.instance)
}

// 获取控件的字符，如果有。
//
// Get the characters of the control, if any.
func (f *TFrame) GetTextBuf(Buffer *string, BufSize int32) int32 {
	return Frame_GetTextBuf(f.instance, Buffer, BufSize)
}

// 获取控件的字符长，如果有。
//
// Get the character length of the control, if any.
func (f *TFrame) GetTextLen() int32 {
	return Frame_GetTextLen(f.instance)
}

// 设置控件字符，如果有。
//
// Set control characters, if any.
func (f *TFrame) SetTextBuf(Buffer string) {
	Frame_SetTextBuf(f.instance, Buffer)
}

// 查找指定名称的组件。
//
// Find the component with the specified name.
func (f *TFrame) FindComponent(AName string) *TComponent {
	return AsComponent(Frame_FindComponent(f.instance, AName))
}

// 获取类名路径。
//
// Get the class name path.
func (f *TFrame) GetNamePath() string {
	return Frame_GetNamePath(f.instance)
}

// 复制一个对象，如果对象实现了此方法的话。
//
// Copy an object, if the object implements this method.
func (f *TFrame) Assign(Source IObject) {
	Frame_Assign(f.instance, CheckPtr(Source))
}

// 获取类的类型信息。
//
// Get class type information.
func (f *TFrame) ClassType() TClass {
	return Frame_ClassType(f.instance)
}

// 获取当前对象类名称。
//
// Get the current object class name.
func (f *TFrame) ClassName() string {
	return Frame_ClassName(f.instance)
}

// 获取当前对象实例大小。
//
// Get the current object instance size.
func (f *TFrame) InstanceSize() int32 {
	return Frame_InstanceSize(f.instance)
}

// 判断当前类是否继承自指定类。
//
// Determine whether the current class inherits from the specified class.
func (f *TFrame) InheritsFrom(AClass TClass) bool {
	return Frame_InheritsFrom(f.instance, AClass)
}

// 与一个对象进行比较。
//
// Compare with an object.
func (f *TFrame) Equals(Obj IObject) bool {
	return Frame_Equals(f.instance, CheckPtr(Obj))
}

// 获取类的哈希值。
//
// Get the hash value of the class.
func (f *TFrame) GetHashCode() int32 {
	return Frame_GetHashCode(f.instance)
}

// 文本类信息。
//
// Text information.
func (f *TFrame) ToString() string {
	return Frame_ToString(f.instance)
}

func (f *TFrame) AnchorToNeighbour(ASide TAnchorKind, ASpace int32, ASibling IControl) {
	Frame_AnchorToNeighbour(f.instance, ASide, ASpace, CheckPtr(ASibling))
}

func (f *TFrame) AnchorParallel(ASide TAnchorKind, ASpace int32, ASibling IControl) {
	Frame_AnchorParallel(f.instance, ASide, ASpace, CheckPtr(ASibling))
}

// 置于指定控件的横向中心。
func (f *TFrame) AnchorHorizontalCenterTo(ASibling IControl) {
	Frame_AnchorHorizontalCenterTo(f.instance, CheckPtr(ASibling))
}

// 置于指定控件的纵向中心。
func (f *TFrame) AnchorVerticalCenterTo(ASibling IControl) {
	Frame_AnchorVerticalCenterTo(f.instance, CheckPtr(ASibling))
}

func (f *TFrame) AnchorSame(ASide TAnchorKind, ASibling IControl) {
	Frame_AnchorSame(f.instance, ASide, CheckPtr(ASibling))
}

func (f *TFrame) AnchorAsAlign(ATheAlign TAlign, ASpace int32) {
	Frame_AnchorAsAlign(f.instance, ATheAlign, ASpace)
}

func (f *TFrame) AnchorClient(ASpace int32) {
	Frame_AnchorClient(f.instance, ASpace)
}

func (f *TFrame) ScaleDesignToForm(ASize int32) int32 {
	return Frame_ScaleDesignToForm(f.instance, ASize)
}

func (f *TFrame) ScaleFormToDesign(ASize int32) int32 {
	return Frame_ScaleFormToDesign(f.instance, ASize)
}

func (f *TFrame) Scale96ToForm(ASize int32) int32 {
	return Frame_Scale96ToForm(f.instance, ASize)
}

func (f *TFrame) ScaleFormTo96(ASize int32) int32 {
	return Frame_ScaleFormTo96(f.instance, ASize)
}

func (f *TFrame) Scale96ToFont(ASize int32) int32 {
	return Frame_Scale96ToFont(f.instance, ASize)
}

func (f *TFrame) ScaleFontTo96(ASize int32) int32 {
	return Frame_ScaleFontTo96(f.instance, ASize)
}

func (f *TFrame) ScaleScreenToFont(ASize int32) int32 {
	return Frame_ScaleScreenToFont(f.instance, ASize)
}

func (f *TFrame) ScaleFontToScreen(ASize int32) int32 {
	return Frame_ScaleFontToScreen(f.instance, ASize)
}

func (f *TFrame) Scale96ToScreen(ASize int32) int32 {
	return Frame_Scale96ToScreen(f.instance, ASize)
}

func (f *TFrame) ScaleScreenTo96(ASize int32) int32 {
	return Frame_ScaleScreenTo96(f.instance, ASize)
}

func (f *TFrame) AutoAdjustLayout(AMode TLayoutAdjustmentPolicy, AFromPPI int32, AToPPI int32, AOldFormWidth int32, ANewFormWidth int32) {
	Frame_AutoAdjustLayout(f.instance, AMode, AFromPPI, AToPPI, AOldFormWidth, ANewFormWidth)
}

func (f *TFrame) FixDesignFontsPPI(ADesignTimePPI int32) {
	Frame_FixDesignFontsPPI(f.instance, ADesignTimePPI)
}

func (f *TFrame) ScaleFontsPPI(AToPPI int32, AProportion float64) {
	Frame_ScaleFontsPPI(f.instance, AToPPI, AProportion)
}

func (f *TFrame) DesignTimePPI() int32 {
	return Frame_GetDesignTimePPI(f.instance)
}

func (f *TFrame) SetDesignTimePPI(value int32) {
	Frame_SetDesignTimePPI(f.instance, value)
}

// 获取控件自动调整。
//
// Get Control automatically adjusts.
func (f *TFrame) Align() TAlign {
	return Frame_GetAlign(f.instance)
}

// 设置控件自动调整。
//
// Set Control automatically adjusts.
func (f *TFrame) SetAlign(value TAlign) {
	Frame_SetAlign(f.instance, value)
}

// 获取四个角位置的锚点。
func (f *TFrame) Anchors() TAnchors {
	return Frame_GetAnchors(f.instance)
}

// 设置四个角位置的锚点。
func (f *TFrame) SetAnchors(value TAnchors) {
	Frame_SetAnchors(f.instance, value)
}

func (f *TFrame) AutoScroll() bool {
	return Frame_GetAutoScroll(f.instance)
}

func (f *TFrame) SetAutoScroll(value bool) {
	Frame_SetAutoScroll(f.instance, value)
}

// 获取自动调整大小。
func (f *TFrame) AutoSize() bool {
	return Frame_GetAutoSize(f.instance)
}

// 设置自动调整大小。
func (f *TFrame) SetAutoSize(value bool) {
	Frame_SetAutoSize(f.instance, value)
}

func (f *TFrame) BiDiMode() TBiDiMode {
	return Frame_GetBiDiMode(f.instance)
}

func (f *TFrame) SetBiDiMode(value TBiDiMode) {
	Frame_SetBiDiMode(f.instance, value)
}

// 获取约束控件大小。
func (f *TFrame) Constraints() *TSizeConstraints {
	return AsSizeConstraints(Frame_GetConstraints(f.instance))
}

// 设置约束控件大小。
func (f *TFrame) SetConstraints(value *TSizeConstraints) {
	Frame_SetConstraints(f.instance, CheckPtr(value))
}

// 获取停靠站点。
//
// Get Docking site.
func (f *TFrame) DockSite() bool {
	return Frame_GetDockSite(f.instance)
}

// 设置停靠站点。
//
// Set Docking site.
func (f *TFrame) SetDockSite(value bool) {
	Frame_SetDockSite(f.instance, value)
}

// 获取设置控件双缓冲。
//
// Get Set control double buffering.
func (f *TFrame) DoubleBuffered() bool {
	return Frame_GetDoubleBuffered(f.instance)
}

// 设置设置控件双缓冲。
//
// Set Set control double buffering.
func (f *TFrame) SetDoubleBuffered(value bool) {
	Frame_SetDoubleBuffered(f.instance, value)
}

// 获取设置控件拖拽时的光标。
//
// Get Set the cursor when the control is dragged.
func (f *TFrame) DragCursor() TCursor {
	return Frame_GetDragCursor(f.instance)
}

// 设置设置控件拖拽时的光标。
//
// Set Set the cursor when the control is dragged.
func (f *TFrame) SetDragCursor(value TCursor) {
	Frame_SetDragCursor(f.instance, value)
}

// 获取拖拽方式。
//
// Get Drag and drop.
func (f *TFrame) DragKind() TDragKind {
	return Frame_GetDragKind(f.instance)
}

// 设置拖拽方式。
//
// Set Drag and drop.
func (f *TFrame) SetDragKind(value TDragKind) {
	Frame_SetDragKind(f.instance, value)
}

// 获取拖拽模式。
//
// Get Drag mode.
func (f *TFrame) DragMode() TDragMode {
	return Frame_GetDragMode(f.instance)
}

// 设置拖拽模式。
//
// Set Drag mode.
func (f *TFrame) SetDragMode(value TDragMode) {
	Frame_SetDragMode(f.instance, value)
}

// 获取控件启用。
//
// Get the control enabled.
func (f *TFrame) Enabled() bool {
	return Frame_GetEnabled(f.instance)
}

// 设置控件启用。
//
// Set the control enabled.
func (f *TFrame) SetEnabled(value bool) {
	Frame_SetEnabled(f.instance, value)
}

// 获取颜色。
//
// Get color.
func (f *TFrame) Color() TColor {
	return Frame_GetColor(f.instance)
}

// 设置颜色。
//
// Set color.
func (f *TFrame) SetColor(value TColor) {
	Frame_SetColor(f.instance, value)
}

// 获取字体。
//
// Get Font.
func (f *TFrame) Font() *TFont {
	return AsFont(Frame_GetFont(f.instance))
}

// 设置字体。
//
// Set Font.
func (f *TFrame) SetFont(value *TFont) {
	Frame_SetFont(f.instance, CheckPtr(value))
}

func (f *TFrame) ParentBackground() bool {
	return Frame_GetParentBackground(f.instance)
}

func (f *TFrame) SetParentBackground(value bool) {
	Frame_SetParentBackground(f.instance, value)
}

// 获取使用父容器颜色。
//
// Get parent color.
func (f *TFrame) ParentColor() bool {
	return Frame_GetParentColor(f.instance)
}

// 设置使用父容器颜色。
//
// Set parent color.
func (f *TFrame) SetParentColor(value bool) {
	Frame_SetParentColor(f.instance, value)
}

// 获取使用父容器双缓冲。
//
// Get Parent container double buffering.
func (f *TFrame) ParentDoubleBuffered() bool {
	return Frame_GetParentDoubleBuffered(f.instance)
}

// 设置使用父容器双缓冲。
//
// Set Parent container double buffering.
func (f *TFrame) SetParentDoubleBuffered(value bool) {
	Frame_SetParentDoubleBuffered(f.instance, value)
}

// 获取使用父容器字体。
//
// Get Parent container font.
func (f *TFrame) ParentFont() bool {
	return Frame_GetParentFont(f.instance)
}

// 设置使用父容器字体。
//
// Set Parent container font.
func (f *TFrame) SetParentFont(value bool) {
	Frame_SetParentFont(f.instance, value)
}

// 获取以父容器的ShowHint属性为准。
func (f *TFrame) ParentShowHint() bool {
	return Frame_GetParentShowHint(f.instance)
}

// 设置以父容器的ShowHint属性为准。
func (f *TFrame) SetParentShowHint(value bool) {
	Frame_SetParentShowHint(f.instance, value)
}

// 获取右键菜单。
//
// Get Right click menu.
func (f *TFrame) PopupMenu() *TPopupMenu {
	return AsPopupMenu(Frame_GetPopupMenu(f.instance))
}

// 设置右键菜单。
//
// Set Right click menu.
func (f *TFrame) SetPopupMenu(value IComponent) {
	Frame_SetPopupMenu(f.instance, CheckPtr(value))
}

// 获取显示鼠标悬停提示。
//
// Get Show mouseover tips.
func (f *TFrame) ShowHint() bool {
	return Frame_GetShowHint(f.instance)
}

// 设置显示鼠标悬停提示。
//
// Set Show mouseover tips.
func (f *TFrame) SetShowHint(value bool) {
	Frame_SetShowHint(f.instance, value)
}

// 获取Tab切换顺序序号。
//
// Get Tab switching sequence number.
func (f *TFrame) TabOrder() TTabOrder {
	return Frame_GetTabOrder(f.instance)
}

// 设置Tab切换顺序序号。
//
// Set Tab switching sequence number.
func (f *TFrame) SetTabOrder(value TTabOrder) {
	Frame_SetTabOrder(f.instance, value)
}

// 获取Tab可停留。
//
// Get Tab can stay.
func (f *TFrame) TabStop() bool {
	return Frame_GetTabStop(f.instance)
}

// 设置Tab可停留。
//
// Set Tab can stay.
func (f *TFrame) SetTabStop(value bool) {
	Frame_SetTabStop(f.instance, value)
}

// 获取控件可视。
//
// Get the control visible.
func (f *TFrame) Visible() bool {
	return Frame_GetVisible(f.instance)
}

// 设置控件可视。
//
// Set the control visible.
func (f *TFrame) SetVisible(value bool) {
	Frame_SetVisible(f.instance, value)
}

// 设置对齐位置事件，当Align为alCustom时Parent会收到这个消息。
func (f *TFrame) SetOnAlignPosition(fn TAlignPositionEvent) {
	Frame_SetOnAlignPosition(f.instance, fn)
}

// 设置控件单击事件。
//
// Set control click event.
func (f *TFrame) SetOnClick(fn TNotifyEvent) {
	Frame_SetOnClick(f.instance, fn)
}

func (f *TFrame) SetOnConstrainedResize(fn TConstrainedResizeEvent) {
	Frame_SetOnConstrainedResize(f.instance, fn)
}

// 设置上下文弹出事件，一般是右键时弹出。
//
// Set Context popup event, usually pop up when right click.
func (f *TFrame) SetOnContextPopup(fn TContextPopupEvent) {
	Frame_SetOnContextPopup(f.instance, fn)
}

// 设置双击事件。
func (f *TFrame) SetOnDblClick(fn TNotifyEvent) {
	Frame_SetOnDblClick(f.instance, fn)
}

func (f *TFrame) SetOnDockDrop(fn TDockDropEvent) {
	Frame_SetOnDockDrop(f.instance, fn)
}

// 设置拖拽下落事件。
//
// Set Drag and drop event.
func (f *TFrame) SetOnDragDrop(fn TDragDropEvent) {
	Frame_SetOnDragDrop(f.instance, fn)
}

// 设置拖拽完成事件。
//
// Set Drag and drop completion event.
func (f *TFrame) SetOnDragOver(fn TDragOverEvent) {
	Frame_SetOnDragOver(f.instance, fn)
}

// 设置停靠结束事件。
//
// Set Dock end event.
func (f *TFrame) SetOnEndDock(fn TEndDragEvent) {
	Frame_SetOnEndDock(f.instance, fn)
}

// 设置拖拽结束。
//
// Set End of drag.
func (f *TFrame) SetOnEndDrag(fn TEndDragEvent) {
	Frame_SetOnEndDrag(f.instance, fn)
}

// 设置焦点进入。
//
// Set Focus entry.
func (f *TFrame) SetOnEnter(fn TNotifyEvent) {
	Frame_SetOnEnter(f.instance, fn)
}

// 设置焦点退出。
//
// Set Focus exit.
func (f *TFrame) SetOnExit(fn TNotifyEvent) {
	Frame_SetOnExit(f.instance, fn)
}

func (f *TFrame) SetOnGetSiteInfo(fn TGetSiteInfoEvent) {
	Frame_SetOnGetSiteInfo(f.instance, fn)
}

// 设置鼠标按下事件。
//
// Set Mouse down event.
func (f *TFrame) SetOnMouseDown(fn TMouseEvent) {
	Frame_SetOnMouseDown(f.instance, fn)
}

// 设置鼠标进入事件。
//
// Set Mouse entry event.
func (f *TFrame) SetOnMouseEnter(fn TNotifyEvent) {
	Frame_SetOnMouseEnter(f.instance, fn)
}

// 设置鼠标离开事件。
//
// Set Mouse leave event.
func (f *TFrame) SetOnMouseLeave(fn TNotifyEvent) {
	Frame_SetOnMouseLeave(f.instance, fn)
}

// 设置鼠标移动事件。
func (f *TFrame) SetOnMouseMove(fn TMouseMoveEvent) {
	Frame_SetOnMouseMove(f.instance, fn)
}

// 设置鼠标抬起事件。
//
// Set Mouse lift event.
func (f *TFrame) SetOnMouseUp(fn TMouseEvent) {
	Frame_SetOnMouseUp(f.instance, fn)
}

// 设置鼠标滚轮事件。
func (f *TFrame) SetOnMouseWheel(fn TMouseWheelEvent) {
	Frame_SetOnMouseWheel(f.instance, fn)
}

// 设置鼠标滚轮按下事件。
func (f *TFrame) SetOnMouseWheelDown(fn TMouseWheelUpDownEvent) {
	Frame_SetOnMouseWheelDown(f.instance, fn)
}

// 设置鼠标滚轮抬起事件。
func (f *TFrame) SetOnMouseWheelUp(fn TMouseWheelUpDownEvent) {
	Frame_SetOnMouseWheelUp(f.instance, fn)
}

// 设置大小被改变事件。
func (f *TFrame) SetOnResize(fn TNotifyEvent) {
	Frame_SetOnResize(f.instance, fn)
}

// 设置启动停靠。
func (f *TFrame) SetOnStartDock(fn TStartDockEvent) {
	Frame_SetOnStartDock(f.instance, fn)
}

func (f *TFrame) SetOnUnDock(fn TUnDockEvent) {
	Frame_SetOnUnDock(f.instance, fn)
}

func (f *TFrame) HorzScrollBar() *TControlScrollBar {
	return AsControlScrollBar(Frame_GetHorzScrollBar(f.instance))
}

func (f *TFrame) SetHorzScrollBar(value *TControlScrollBar) {
	Frame_SetHorzScrollBar(f.instance, CheckPtr(value))
}

func (f *TFrame) VertScrollBar() *TControlScrollBar {
	return AsControlScrollBar(Frame_GetVertScrollBar(f.instance))
}

func (f *TFrame) SetVertScrollBar(value *TControlScrollBar) {
	Frame_SetVertScrollBar(f.instance, CheckPtr(value))
}

// 获取依靠客户端总数。
func (f *TFrame) DockClientCount() int32 {
	return Frame_GetDockClientCount(f.instance)
}

// 获取鼠标是否在客户端，仅VCL有效。
//
// Get Whether the mouse is on the client, only VCL is valid.
func (f *TFrame) MouseInClient() bool {
	return Frame_GetMouseInClient(f.instance)
}

// 获取当前停靠的可视总数。
//
// Get The total number of visible calls currently docked.
func (f *TFrame) VisibleDockClientCount() int32 {
	return Frame_GetVisibleDockClientCount(f.instance)
}

// 获取画刷对象。
//
// Get Brush.
func (f *TFrame) Brush() *TBrush {
	return AsBrush(Frame_GetBrush(f.instance))
}

// 获取子控件数。
//
// Get Number of child controls.
func (f *TFrame) ControlCount() int32 {
	return Frame_GetControlCount(f.instance)
}

// 获取控件句柄。
//
// Get Control handle.
func (f *TFrame) Handle() HWND {
	return Frame_GetHandle(f.instance)
}

// 获取父容器句柄。
//
// Get Parent container handle.
func (f *TFrame) ParentWindow() HWND {
	return Frame_GetParentWindow(f.instance)
}

// 设置父容器句柄。
//
// Set Parent container handle.
func (f *TFrame) SetParentWindow(value HWND) {
	Frame_SetParentWindow(f.instance, value)
}

func (f *TFrame) Showing() bool {
	return Frame_GetShowing(f.instance)
}

// 获取使用停靠管理。
func (f *TFrame) UseDockManager() bool {
	return Frame_GetUseDockManager(f.instance)
}

// 设置使用停靠管理。
func (f *TFrame) SetUseDockManager(value bool) {
	Frame_SetUseDockManager(f.instance, value)
}

func (f *TFrame) Action() *TAction {
	return AsAction(Frame_GetAction(f.instance))
}

func (f *TFrame) SetAction(value IComponent) {
	Frame_SetAction(f.instance, CheckPtr(value))
}

func (f *TFrame) BoundsRect() TRect {
	return Frame_GetBoundsRect(f.instance)
}

func (f *TFrame) SetBoundsRect(value TRect) {
	Frame_SetBoundsRect(f.instance, value)
}

// 获取客户区高度。
//
// Get client height.
func (f *TFrame) ClientHeight() int32 {
	return Frame_GetClientHeight(f.instance)
}

// 设置客户区高度。
//
// Set client height.
func (f *TFrame) SetClientHeight(value int32) {
	Frame_SetClientHeight(f.instance, value)
}

func (f *TFrame) ClientOrigin() TPoint {
	return Frame_GetClientOrigin(f.instance)
}

// 获取客户区矩形。
//
// Get client rectangle.
func (f *TFrame) ClientRect() TRect {
	return Frame_GetClientRect(f.instance)
}

// 获取客户区宽度。
//
// Get client width.
func (f *TFrame) ClientWidth() int32 {
	return Frame_GetClientWidth(f.instance)
}

// 设置客户区宽度。
//
// Set client width.
func (f *TFrame) SetClientWidth(value int32) {
	Frame_SetClientWidth(f.instance, value)
}

// 获取控件状态。
//
// Get control state.
func (f *TFrame) ControlState() TControlState {
	return Frame_GetControlState(f.instance)
}

// 设置控件状态。
//
// Set control state.
func (f *TFrame) SetControlState(value TControlState) {
	Frame_SetControlState(f.instance, value)
}

// 获取控件样式。
//
// Get control style.
func (f *TFrame) ControlStyle() TControlStyle {
	return Frame_GetControlStyle(f.instance)
}

// 设置控件样式。
//
// Set control style.
func (f *TFrame) SetControlStyle(value TControlStyle) {
	Frame_SetControlStyle(f.instance, value)
}

func (f *TFrame) Floating() bool {
	return Frame_GetFloating(f.instance)
}

// 获取控件父容器。
//
// Get control parent container.
func (f *TFrame) Parent() *TWinControl {
	return AsWinControl(Frame_GetParent(f.instance))
}

// 设置控件父容器。
//
// Set control parent container.
func (f *TFrame) SetParent(value IWinControl) {
	Frame_SetParent(f.instance, CheckPtr(value))
}

// 获取左边位置。
//
// Get Left position.
func (f *TFrame) Left() int32 {
	return Frame_GetLeft(f.instance)
}

// 设置左边位置。
//
// Set Left position.
func (f *TFrame) SetLeft(value int32) {
	Frame_SetLeft(f.instance, value)
}

// 获取顶边位置。
//
// Get Top position.
func (f *TFrame) Top() int32 {
	return Frame_GetTop(f.instance)
}

// 设置顶边位置。
//
// Set Top position.
func (f *TFrame) SetTop(value int32) {
	Frame_SetTop(f.instance, value)
}

// 获取宽度。
//
// Get width.
func (f *TFrame) Width() int32 {
	return Frame_GetWidth(f.instance)
}

// 设置宽度。
//
// Set width.
func (f *TFrame) SetWidth(value int32) {
	Frame_SetWidth(f.instance, value)
}

// 获取高度。
//
// Get height.
func (f *TFrame) Height() int32 {
	return Frame_GetHeight(f.instance)
}

// 设置高度。
//
// Set height.
func (f *TFrame) SetHeight(value int32) {
	Frame_SetHeight(f.instance, value)
}

// 获取控件光标。
//
// Get control cursor.
func (f *TFrame) Cursor() TCursor {
	return Frame_GetCursor(f.instance)
}

// 设置控件光标。
//
// Set control cursor.
func (f *TFrame) SetCursor(value TCursor) {
	Frame_SetCursor(f.instance, value)
}

// 获取组件鼠标悬停提示。
//
// Get component mouse hints.
func (f *TFrame) Hint() string {
	return Frame_GetHint(f.instance)
}

// 设置组件鼠标悬停提示。
//
// Set component mouse hints.
func (f *TFrame) SetHint(value string) {
	Frame_SetHint(f.instance, value)
}

// 获取组件总数。
//
// Get the total number of components.
func (f *TFrame) ComponentCount() int32 {
	return Frame_GetComponentCount(f.instance)
}

// 获取组件索引。
//
// Get component index.
func (f *TFrame) ComponentIndex() int32 {
	return Frame_GetComponentIndex(f.instance)
}

// 设置组件索引。
//
// Set component index.
func (f *TFrame) SetComponentIndex(value int32) {
	Frame_SetComponentIndex(f.instance, value)
}

// 获取组件所有者。
//
// Get component owner.
func (f *TFrame) Owner() *TComponent {
	return AsComponent(Frame_GetOwner(f.instance))
}

// 获取组件名称。
//
// Get the component name.
func (f *TFrame) Name() string {
	return Frame_GetName(f.instance)
}

// 设置组件名称。
//
// Set the component name.
func (f *TFrame) SetName(value string) {
	Frame_SetName(f.instance, value)
}

// 获取对象标记。
//
// Get the control tag.
func (f *TFrame) Tag() int {
	return Frame_GetTag(f.instance)
}

// 设置对象标记。
//
// Set the control tag.
func (f *TFrame) SetTag(value int) {
	Frame_SetTag(f.instance, value)
}

// 获取左边锚点。
func (f *TFrame) AnchorSideLeft() *TAnchorSide {
	return AsAnchorSide(Frame_GetAnchorSideLeft(f.instance))
}

// 设置左边锚点。
func (f *TFrame) SetAnchorSideLeft(value *TAnchorSide) {
	Frame_SetAnchorSideLeft(f.instance, CheckPtr(value))
}

// 获取顶边锚点。
func (f *TFrame) AnchorSideTop() *TAnchorSide {
	return AsAnchorSide(Frame_GetAnchorSideTop(f.instance))
}

// 设置顶边锚点。
func (f *TFrame) SetAnchorSideTop(value *TAnchorSide) {
	Frame_SetAnchorSideTop(f.instance, CheckPtr(value))
}

// 获取右边锚点。
func (f *TFrame) AnchorSideRight() *TAnchorSide {
	return AsAnchorSide(Frame_GetAnchorSideRight(f.instance))
}

// 设置右边锚点。
func (f *TFrame) SetAnchorSideRight(value *TAnchorSide) {
	Frame_SetAnchorSideRight(f.instance, CheckPtr(value))
}

// 获取底边锚点。
func (f *TFrame) AnchorSideBottom() *TAnchorSide {
	return AsAnchorSide(Frame_GetAnchorSideBottom(f.instance))
}

// 设置底边锚点。
func (f *TFrame) SetAnchorSideBottom(value *TAnchorSide) {
	Frame_SetAnchorSideBottom(f.instance, CheckPtr(value))
}

func (f *TFrame) ChildSizing() *TControlChildSizing {
	return AsControlChildSizing(Frame_GetChildSizing(f.instance))
}

func (f *TFrame) SetChildSizing(value *TControlChildSizing) {
	Frame_SetChildSizing(f.instance, CheckPtr(value))
}

// 获取边框间距。
func (f *TFrame) BorderSpacing() *TControlBorderSpacing {
	return AsControlBorderSpacing(Frame_GetBorderSpacing(f.instance))
}

// 设置边框间距。
func (f *TFrame) SetBorderSpacing(value *TControlBorderSpacing) {
	Frame_SetBorderSpacing(f.instance, CheckPtr(value))
}

// 获取指定索引停靠客户端。
func (f *TFrame) DockClients(Index int32) *TControl {
	return AsControl(Frame_GetDockClients(f.instance, Index))
}

// 获取指定索引子控件。
func (f *TFrame) Controls(Index int32) *TControl {
	return AsControl(Frame_GetControls(f.instance, Index))
}

// 获取指定索引组件。
//
// Get the specified index component.
func (f *TFrame) Components(AIndex int32) *TComponent {
	return AsComponent(Frame_GetComponents(f.instance, AIndex))
}

// 获取锚侧面。
func (f *TFrame) AnchorSide(AKind TAnchorKind) *TAnchorSide {
	return AsAnchorSide(Frame_GetAnchorSide(f.instance, AKind))
}