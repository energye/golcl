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

type TBitBtn struct {
	IWinControl
	instance uintptr
	// 特殊情况下使用，主要应对Go的GC问题，与LCL没有太多关系。
	ptr unsafe.Pointer
}

// 创建一个新的对象。
//
// Create a new object.
func NewBitBtn(owner IComponent) *TBitBtn {
	b := new(TBitBtn)
	b.instance = BitBtn_Create(CheckPtr(owner))
	b.ptr = unsafe.Pointer(b.instance)
	return b
}

// 动态转换一个已存在的对象实例。
//
// Dynamically convert an existing object instance.
func AsBitBtn(obj interface{}) *TBitBtn {
	instance, ptr := getInstance(obj)
	if instance == 0 {
		return nil
	}
	return &TBitBtn{instance: instance, ptr: ptr}
}

// -------------------------- Deprecated begin --------------------------
// 新建一个对象来自已经存在的对象实例指针。
//
// Create a new object from an existing object instance pointer.
// Deprecated: use AsBitBtn.
func BitBtnFromInst(inst uintptr) *TBitBtn {
	return AsBitBtn(inst)
}

// 新建一个对象来自已经存在的对象实例。
//
// Create a new object from an existing object instance.
// Deprecated: use AsBitBtn.
func BitBtnFromObj(obj IObject) *TBitBtn {
	return AsBitBtn(obj)
}

// 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
//
// Create a new object from an unsecured address. Note: Using this function may cause some unclear situations and be used with caution..
// Deprecated: use AsBitBtn.
func BitBtnFromUnsafePointer(ptr unsafe.Pointer) *TBitBtn {
	return AsBitBtn(ptr)
}

// -------------------------- Deprecated end --------------------------
// 释放对象。
//
// Free object.
func (b *TBitBtn) Free() {
	if b.instance != 0 {
		BitBtn_Free(b.instance)
		b.instance, b.ptr = 0, nullptr
	}
}

// 返回对象实例指针。
//
// Return object instance pointer.
func (b *TBitBtn) Instance() uintptr {
	return b.instance
}

// 获取一个不安全的地址。
//
// Get an unsafe address.
func (b *TBitBtn) UnsafeAddr() unsafe.Pointer {
	return b.ptr
}

// 检测地址是否为空。
//
// Check if the address is empty.
func (b *TBitBtn) IsValid() bool {
	return b.instance != 0
}

// 检测当前对象是否继承自目标对象。
//
// Checks whether the current object is inherited from the target object.
func (b *TBitBtn) Is() TIs {
	return TIs(b.instance)
}

// 动态转换当前对象为目标对象。
//
// Dynamically convert the current object to the target object.
//func (b *TBitBtn) As() TAs {
//    return TAs(b.instance)
//}

// 获取类信息指针。
//
// Get class information pointer.
func TBitBtnClass() TClass {
	return BitBtn_StaticClassType()
}

// 单击。
func (b *TBitBtn) Click() {
	BitBtn_Click(b.instance)
}

// 是否可以获得焦点。
func (b *TBitBtn) CanFocus() bool {
	return BitBtn_CanFocus(b.instance)
}

// 返回是否包含指定控件。
//
// it's contain a specified control.
func (b *TBitBtn) ContainsControl(Control IControl) bool {
	return BitBtn_ContainsControl(b.instance, CheckPtr(Control))
}

// 返回指定坐标及相关属性位置控件。
//
// Returns the specified coordinate and the relevant attribute position control..
func (b *TBitBtn) ControlAtPos(Pos TPoint, AllowDisabled bool, AllowWinControls bool, AllLevels bool) *TControl {
	return AsControl(BitBtn_ControlAtPos(b.instance, Pos, AllowDisabled, AllowWinControls, AllLevels))
}

// 禁用控件的对齐。
//
// Disable control alignment.
func (b *TBitBtn) DisableAlign() {
	BitBtn_DisableAlign(b.instance)
}

// 启用控件对齐。
//
// Enabled control alignment.
func (b *TBitBtn) EnableAlign() {
	BitBtn_EnableAlign(b.instance)
}

// 查找子控件。
//
// Find sub controls.
func (b *TBitBtn) FindChildControl(ControlName string) *TControl {
	return AsControl(BitBtn_FindChildControl(b.instance, ControlName))
}

func (b *TBitBtn) FlipChildren(AllLevels bool) {
	BitBtn_FlipChildren(b.instance, AllLevels)
}

// 返回是否获取焦点。
//
// Return to get focus.
func (b *TBitBtn) Focused() bool {
	return BitBtn_Focused(b.instance)
}

// 句柄是否已经分配。
//
// Is the handle already allocated.
func (b *TBitBtn) HandleAllocated() bool {
	return BitBtn_HandleAllocated(b.instance)
}

// 插入一个控件。
//
// Insert a control.
func (b *TBitBtn) InsertControl(AControl IControl) {
	BitBtn_InsertControl(b.instance, CheckPtr(AControl))
}

// 要求重绘。
//
// Redraw.
func (b *TBitBtn) Invalidate() {
	BitBtn_Invalidate(b.instance)
}

// 绘画至指定DC。
//
// Painting to the specified DC.
func (b *TBitBtn) PaintTo(DC HDC, X int32, Y int32) {
	BitBtn_PaintTo(b.instance, DC, X, Y)
}

// 移除一个控件。
//
// Remove a control.
func (b *TBitBtn) RemoveControl(AControl IControl) {
	BitBtn_RemoveControl(b.instance, CheckPtr(AControl))
}

// 重新对齐。
//
// Realign.
func (b *TBitBtn) Realign() {
	BitBtn_Realign(b.instance)
}

// 重绘。
//
// Repaint.
func (b *TBitBtn) Repaint() {
	BitBtn_Repaint(b.instance)
}

// 按比例缩放。
//
// Scale by.
func (b *TBitBtn) ScaleBy(M int32, D int32) {
	BitBtn_ScaleBy(b.instance, M, D)
}

// 滚动至指定位置。
//
// Scroll by.
func (b *TBitBtn) ScrollBy(DeltaX int32, DeltaY int32) {
	BitBtn_ScrollBy(b.instance, DeltaX, DeltaY)
}

// 设置组件边界。
//
// Set component boundaries.
func (b *TBitBtn) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
	BitBtn_SetBounds(b.instance, ALeft, ATop, AWidth, AHeight)
}

// 设置控件焦点。
//
// Set control focus.
func (b *TBitBtn) SetFocus() {
	BitBtn_SetFocus(b.instance)
}

// 控件更新。
//
// Update.
func (b *TBitBtn) Update() {
	BitBtn_Update(b.instance)
}

// 将控件置于最前。
//
// Bring the control to the front.
func (b *TBitBtn) BringToFront() {
	BitBtn_BringToFront(b.instance)
}

// 将客户端坐标转为绝对的屏幕坐标。
//
// Convert client coordinates to absolute screen coordinates.
func (b *TBitBtn) ClientToScreen(Point TPoint) TPoint {
	return BitBtn_ClientToScreen(b.instance, Point)
}

// 将客户端坐标转为父容器坐标。
//
// Convert client coordinates to parent container coordinates.
func (b *TBitBtn) ClientToParent(Point TPoint, AParent IWinControl) TPoint {
	return BitBtn_ClientToParent(b.instance, Point, CheckPtr(AParent))
}

// 是否在拖拽中。
//
// Is it in the middle of dragging.
func (b *TBitBtn) Dragging() bool {
	return BitBtn_Dragging(b.instance)
}

// 是否有父容器。
//
// Is there a parent container.
func (b *TBitBtn) HasParent() bool {
	return BitBtn_HasParent(b.instance)
}

// 隐藏控件。
//
// Hidden control.
func (b *TBitBtn) Hide() {
	BitBtn_Hide(b.instance)
}

// 发送一个消息。
//
// Send a message.
func (b *TBitBtn) Perform(Msg uint32, WParam uintptr, LParam int) int {
	return BitBtn_Perform(b.instance, Msg, WParam, LParam)
}

// 刷新控件。
//
// Refresh control.
func (b *TBitBtn) Refresh() {
	BitBtn_Refresh(b.instance)
}

// 将屏幕坐标转为客户端坐标。
//
// Convert screen coordinates to client coordinates.
func (b *TBitBtn) ScreenToClient(Point TPoint) TPoint {
	return BitBtn_ScreenToClient(b.instance, Point)
}

// 将父容器坐标转为客户端坐标。
//
// Convert parent container coordinates to client coordinates.
func (b *TBitBtn) ParentToClient(Point TPoint, AParent IWinControl) TPoint {
	return BitBtn_ParentToClient(b.instance, Point, CheckPtr(AParent))
}

// 控件至于最后面。
//
// The control is placed at the end.
func (b *TBitBtn) SendToBack() {
	BitBtn_SendToBack(b.instance)
}

// 显示控件。
//
// Show control.
func (b *TBitBtn) Show() {
	BitBtn_Show(b.instance)
}

// 获取控件的字符，如果有。
//
// Get the characters of the control, if any.
func (b *TBitBtn) GetTextBuf(Buffer *string, BufSize int32) int32 {
	return BitBtn_GetTextBuf(b.instance, Buffer, BufSize)
}

// 获取控件的字符长，如果有。
//
// Get the character length of the control, if any.
func (b *TBitBtn) GetTextLen() int32 {
	return BitBtn_GetTextLen(b.instance)
}

// 设置控件字符，如果有。
//
// Set control characters, if any.
func (b *TBitBtn) SetTextBuf(Buffer string) {
	BitBtn_SetTextBuf(b.instance, Buffer)
}

// 查找指定名称的组件。
//
// Find the component with the specified name.
func (b *TBitBtn) FindComponent(AName string) *TComponent {
	return AsComponent(BitBtn_FindComponent(b.instance, AName))
}

// 获取类名路径。
//
// Get the class name path.
func (b *TBitBtn) GetNamePath() string {
	return BitBtn_GetNamePath(b.instance)
}

// 复制一个对象，如果对象实现了此方法的话。
//
// Copy an object, if the object implements this method.
func (b *TBitBtn) Assign(Source IObject) {
	BitBtn_Assign(b.instance, CheckPtr(Source))
}

// 获取类的类型信息。
//
// Get class type information.
func (b *TBitBtn) ClassType() TClass {
	return BitBtn_ClassType(b.instance)
}

// 获取当前对象类名称。
//
// Get the current object class name.
func (b *TBitBtn) ClassName() string {
	return BitBtn_ClassName(b.instance)
}

// 获取当前对象实例大小。
//
// Get the current object instance size.
func (b *TBitBtn) InstanceSize() int32 {
	return BitBtn_InstanceSize(b.instance)
}

// 判断当前类是否继承自指定类。
//
// Determine whether the current class inherits from the specified class.
func (b *TBitBtn) InheritsFrom(AClass TClass) bool {
	return BitBtn_InheritsFrom(b.instance, AClass)
}

// 与一个对象进行比较。
//
// Compare with an object.
func (b *TBitBtn) Equals(Obj IObject) bool {
	return BitBtn_Equals(b.instance, CheckPtr(Obj))
}

// 获取类的哈希值。
//
// Get the hash value of the class.
func (b *TBitBtn) GetHashCode() int32 {
	return BitBtn_GetHashCode(b.instance)
}

// 文本类信息。
//
// Text information.
func (b *TBitBtn) ToString() string {
	return BitBtn_ToString(b.instance)
}

func (b *TBitBtn) AnchorToNeighbour(ASide TAnchorKind, ASpace int32, ASibling IControl) {
	BitBtn_AnchorToNeighbour(b.instance, ASide, ASpace, CheckPtr(ASibling))
}

func (b *TBitBtn) AnchorParallel(ASide TAnchorKind, ASpace int32, ASibling IControl) {
	BitBtn_AnchorParallel(b.instance, ASide, ASpace, CheckPtr(ASibling))
}

// 置于指定控件的横向中心。
func (b *TBitBtn) AnchorHorizontalCenterTo(ASibling IControl) {
	BitBtn_AnchorHorizontalCenterTo(b.instance, CheckPtr(ASibling))
}

// 置于指定控件的纵向中心。
func (b *TBitBtn) AnchorVerticalCenterTo(ASibling IControl) {
	BitBtn_AnchorVerticalCenterTo(b.instance, CheckPtr(ASibling))
}

func (b *TBitBtn) AnchorSame(ASide TAnchorKind, ASibling IControl) {
	BitBtn_AnchorSame(b.instance, ASide, CheckPtr(ASibling))
}

func (b *TBitBtn) AnchorAsAlign(ATheAlign TAlign, ASpace int32) {
	BitBtn_AnchorAsAlign(b.instance, ATheAlign, ASpace)
}

func (b *TBitBtn) AnchorClient(ASpace int32) {
	BitBtn_AnchorClient(b.instance, ASpace)
}

func (b *TBitBtn) ScaleDesignToForm(ASize int32) int32 {
	return BitBtn_ScaleDesignToForm(b.instance, ASize)
}

func (b *TBitBtn) ScaleFormToDesign(ASize int32) int32 {
	return BitBtn_ScaleFormToDesign(b.instance, ASize)
}

func (b *TBitBtn) Scale96ToForm(ASize int32) int32 {
	return BitBtn_Scale96ToForm(b.instance, ASize)
}

func (b *TBitBtn) ScaleFormTo96(ASize int32) int32 {
	return BitBtn_ScaleFormTo96(b.instance, ASize)
}

func (b *TBitBtn) Scale96ToFont(ASize int32) int32 {
	return BitBtn_Scale96ToFont(b.instance, ASize)
}

func (b *TBitBtn) ScaleFontTo96(ASize int32) int32 {
	return BitBtn_ScaleFontTo96(b.instance, ASize)
}

func (b *TBitBtn) ScaleScreenToFont(ASize int32) int32 {
	return BitBtn_ScaleScreenToFont(b.instance, ASize)
}

func (b *TBitBtn) ScaleFontToScreen(ASize int32) int32 {
	return BitBtn_ScaleFontToScreen(b.instance, ASize)
}

func (b *TBitBtn) Scale96ToScreen(ASize int32) int32 {
	return BitBtn_Scale96ToScreen(b.instance, ASize)
}

func (b *TBitBtn) ScaleScreenTo96(ASize int32) int32 {
	return BitBtn_ScaleScreenTo96(b.instance, ASize)
}

func (b *TBitBtn) AutoAdjustLayout(AMode TLayoutAdjustmentPolicy, AFromPPI int32, AToPPI int32, AOldFormWidth int32, ANewFormWidth int32) {
	BitBtn_AutoAdjustLayout(b.instance, AMode, AFromPPI, AToPPI, AOldFormWidth, ANewFormWidth)
}

func (b *TBitBtn) FixDesignFontsPPI(ADesignTimePPI int32) {
	BitBtn_FixDesignFontsPPI(b.instance, ADesignTimePPI)
}

func (b *TBitBtn) ScaleFontsPPI(AToPPI int32, AProportion float64) {
	BitBtn_ScaleFontsPPI(b.instance, AToPPI, AProportion)
}

func (b *TBitBtn) DefaultCaption() bool {
	return BitBtn_GetDefaultCaption(b.instance)
}

func (b *TBitBtn) SetDefaultCaption(value bool) {
	BitBtn_SetDefaultCaption(b.instance, value)
}

func (b *TBitBtn) GlyphShowMode() TGlyphShowMode {
	return BitBtn_GetGlyphShowMode(b.instance)
}

func (b *TBitBtn) SetGlyphShowMode(value TGlyphShowMode) {
	BitBtn_SetGlyphShowMode(b.instance, value)
}

// 获取图像在images中的索引。
func (b *TBitBtn) ImageIndex() int32 {
	return BitBtn_GetImageIndex(b.instance)
}

// 设置图像在images中的索引。
func (b *TBitBtn) SetImageIndex(value int32) {
	BitBtn_SetImageIndex(b.instance, value)
}

// 获取图标索引列表对象。
func (b *TBitBtn) Images() *TImageList {
	return AsImageList(BitBtn_GetImages(b.instance))
}

// 设置图标索引列表对象。
func (b *TBitBtn) SetImages(value IComponent) {
	BitBtn_SetImages(b.instance, CheckPtr(value))
}

func (b *TBitBtn) ImageWidth() int32 {
	return BitBtn_GetImageWidth(b.instance)
}

func (b *TBitBtn) SetImageWidth(value int32) {
	BitBtn_SetImageWidth(b.instance, value)
}

func (b *TBitBtn) Action() *TAction {
	return AsAction(BitBtn_GetAction(b.instance))
}

func (b *TBitBtn) SetAction(value IComponent) {
	BitBtn_SetAction(b.instance, CheckPtr(value))
}

// 获取控件自动调整。
//
// Get Control automatically adjusts.
func (b *TBitBtn) Align() TAlign {
	return BitBtn_GetAlign(b.instance)
}

// 设置控件自动调整。
//
// Set Control automatically adjusts.
func (b *TBitBtn) SetAlign(value TAlign) {
	BitBtn_SetAlign(b.instance, value)
}

// 获取四个角位置的锚点。
func (b *TBitBtn) Anchors() TAnchors {
	return BitBtn_GetAnchors(b.instance)
}

// 设置四个角位置的锚点。
func (b *TBitBtn) SetAnchors(value TAnchors) {
	BitBtn_SetAnchors(b.instance, value)
}

func (b *TBitBtn) BiDiMode() TBiDiMode {
	return BitBtn_GetBiDiMode(b.instance)
}

func (b *TBitBtn) SetBiDiMode(value TBiDiMode) {
	BitBtn_SetBiDiMode(b.instance, value)
}

func (b *TBitBtn) Cancel() bool {
	return BitBtn_GetCancel(b.instance)
}

func (b *TBitBtn) SetCancel(value bool) {
	BitBtn_SetCancel(b.instance, value)
}

// 获取控件标题。
//
// Get the control title.
func (b *TBitBtn) Caption() string {
	return BitBtn_GetCaption(b.instance)
}

// 设置控件标题。
//
// Set the control title.
func (b *TBitBtn) SetCaption(value string) {
	BitBtn_SetCaption(b.instance, value)
}

// 获取约束控件大小。
func (b *TBitBtn) Constraints() *TSizeConstraints {
	return AsSizeConstraints(BitBtn_GetConstraints(b.instance))
}

// 设置约束控件大小。
func (b *TBitBtn) SetConstraints(value *TSizeConstraints) {
	BitBtn_SetConstraints(b.instance, CheckPtr(value))
}

func (b *TBitBtn) Default() bool {
	return BitBtn_GetDefault(b.instance)
}

func (b *TBitBtn) SetDefault(value bool) {
	BitBtn_SetDefault(b.instance, value)
}

// 获取设置控件双缓冲。
//
// Get Set control double buffering.
func (b *TBitBtn) DoubleBuffered() bool {
	return BitBtn_GetDoubleBuffered(b.instance)
}

// 设置设置控件双缓冲。
//
// Set Set control double buffering.
func (b *TBitBtn) SetDoubleBuffered(value bool) {
	BitBtn_SetDoubleBuffered(b.instance, value)
}

// 获取控件启用。
//
// Get the control enabled.
func (b *TBitBtn) Enabled() bool {
	return BitBtn_GetEnabled(b.instance)
}

// 设置控件启用。
//
// Set the control enabled.
func (b *TBitBtn) SetEnabled(value bool) {
	BitBtn_SetEnabled(b.instance, value)
}

// 获取字体。
//
// Get Font.
func (b *TBitBtn) Font() *TFont {
	return AsFont(BitBtn_GetFont(b.instance))
}

// 设置字体。
//
// Set Font.
func (b *TBitBtn) SetFont(value *TFont) {
	BitBtn_SetFont(b.instance, CheckPtr(value))
}

func (b *TBitBtn) Glyph() *TBitmap {
	return AsBitmap(BitBtn_GetGlyph(b.instance))
}

func (b *TBitBtn) SetGlyph(value *TBitmap) {
	BitBtn_SetGlyph(b.instance, CheckPtr(value))
}

func (b *TBitBtn) Kind() TBitBtnKind {
	return BitBtn_GetKind(b.instance)
}

func (b *TBitBtn) SetKind(value TBitBtnKind) {
	BitBtn_SetKind(b.instance, value)
}

func (b *TBitBtn) Layout() TButtonLayout {
	return BitBtn_GetLayout(b.instance)
}

func (b *TBitBtn) SetLayout(value TButtonLayout) {
	BitBtn_SetLayout(b.instance, value)
}

// 获取模态对话框显示结果。
func (b *TBitBtn) ModalResult() TModalResult {
	return BitBtn_GetModalResult(b.instance)
}

// 设置模态对话框显示结果。
func (b *TBitBtn) SetModalResult(value TModalResult) {
	BitBtn_SetModalResult(b.instance, value)
}

func (b *TBitBtn) NumGlyphs() TNumGlyphs {
	return BitBtn_GetNumGlyphs(b.instance)
}

func (b *TBitBtn) SetNumGlyphs(value TNumGlyphs) {
	BitBtn_SetNumGlyphs(b.instance, value)
}

// 获取使用父容器双缓冲。
//
// Get Parent container double buffering.
func (b *TBitBtn) ParentDoubleBuffered() bool {
	return BitBtn_GetParentDoubleBuffered(b.instance)
}

// 设置使用父容器双缓冲。
//
// Set Parent container double buffering.
func (b *TBitBtn) SetParentDoubleBuffered(value bool) {
	BitBtn_SetParentDoubleBuffered(b.instance, value)
}

// 获取使用父容器字体。
//
// Get Parent container font.
func (b *TBitBtn) ParentFont() bool {
	return BitBtn_GetParentFont(b.instance)
}

// 设置使用父容器字体。
//
// Set Parent container font.
func (b *TBitBtn) SetParentFont(value bool) {
	BitBtn_SetParentFont(b.instance, value)
}

// 获取以父容器的ShowHint属性为准。
func (b *TBitBtn) ParentShowHint() bool {
	return BitBtn_GetParentShowHint(b.instance)
}

// 设置以父容器的ShowHint属性为准。
func (b *TBitBtn) SetParentShowHint(value bool) {
	BitBtn_SetParentShowHint(b.instance, value)
}

// 获取右键菜单。
//
// Get Right click menu.
func (b *TBitBtn) PopupMenu() *TPopupMenu {
	return AsPopupMenu(BitBtn_GetPopupMenu(b.instance))
}

// 设置右键菜单。
//
// Set Right click menu.
func (b *TBitBtn) SetPopupMenu(value IComponent) {
	BitBtn_SetPopupMenu(b.instance, CheckPtr(value))
}

// 获取显示鼠标悬停提示。
//
// Get Show mouseover tips.
func (b *TBitBtn) ShowHint() bool {
	return BitBtn_GetShowHint(b.instance)
}

// 设置显示鼠标悬停提示。
//
// Set Show mouseover tips.
func (b *TBitBtn) SetShowHint(value bool) {
	BitBtn_SetShowHint(b.instance, value)
}

func (b *TBitBtn) Spacing() int32 {
	return BitBtn_GetSpacing(b.instance)
}

func (b *TBitBtn) SetSpacing(value int32) {
	BitBtn_SetSpacing(b.instance, value)
}

// 获取Tab切换顺序序号。
//
// Get Tab switching sequence number.
func (b *TBitBtn) TabOrder() TTabOrder {
	return BitBtn_GetTabOrder(b.instance)
}

// 设置Tab切换顺序序号。
//
// Set Tab switching sequence number.
func (b *TBitBtn) SetTabOrder(value TTabOrder) {
	BitBtn_SetTabOrder(b.instance, value)
}

// 获取Tab可停留。
//
// Get Tab can stay.
func (b *TBitBtn) TabStop() bool {
	return BitBtn_GetTabStop(b.instance)
}

// 设置Tab可停留。
//
// Set Tab can stay.
func (b *TBitBtn) SetTabStop(value bool) {
	BitBtn_SetTabStop(b.instance, value)
}

// 获取控件可视。
//
// Get the control visible.
func (b *TBitBtn) Visible() bool {
	return BitBtn_GetVisible(b.instance)
}

// 设置控件可视。
//
// Set the control visible.
func (b *TBitBtn) SetVisible(value bool) {
	BitBtn_SetVisible(b.instance, value)
}

// 设置控件单击事件。
//
// Set control click event.
func (b *TBitBtn) SetOnClick(fn TNotifyEvent) {
	BitBtn_SetOnClick(b.instance, fn)
}

// 设置上下文弹出事件，一般是右键时弹出。
//
// Set Context popup event, usually pop up when right click.
func (b *TBitBtn) SetOnContextPopup(fn TContextPopupEvent) {
	BitBtn_SetOnContextPopup(b.instance, fn)
}

// 设置拖拽下落事件。
//
// Set Drag and drop event.
func (b *TBitBtn) SetOnDragDrop(fn TDragDropEvent) {
	BitBtn_SetOnDragDrop(b.instance, fn)
}

// 设置拖拽完成事件。
//
// Set Drag and drop completion event.
func (b *TBitBtn) SetOnDragOver(fn TDragOverEvent) {
	BitBtn_SetOnDragOver(b.instance, fn)
}

// 设置拖拽结束。
//
// Set End of drag.
func (b *TBitBtn) SetOnEndDrag(fn TEndDragEvent) {
	BitBtn_SetOnEndDrag(b.instance, fn)
}

// 设置焦点进入。
//
// Set Focus entry.
func (b *TBitBtn) SetOnEnter(fn TNotifyEvent) {
	BitBtn_SetOnEnter(b.instance, fn)
}

// 设置焦点退出。
//
// Set Focus exit.
func (b *TBitBtn) SetOnExit(fn TNotifyEvent) {
	BitBtn_SetOnExit(b.instance, fn)
}

// 设置键盘按键按下事件。
//
// Set Keyboard button press event.
func (b *TBitBtn) SetOnKeyDown(fn TKeyEvent) {
	BitBtn_SetOnKeyDown(b.instance, fn)
}

// 设置键键下事件。
func (b *TBitBtn) SetOnKeyPress(fn TKeyPressEvent) {
	BitBtn_SetOnKeyPress(b.instance, fn)
}

// 设置键盘按键抬起事件。
//
// Set Keyboard button lift event.
func (b *TBitBtn) SetOnKeyUp(fn TKeyEvent) {
	BitBtn_SetOnKeyUp(b.instance, fn)
}

// 设置鼠标按下事件。
//
// Set Mouse down event.
func (b *TBitBtn) SetOnMouseDown(fn TMouseEvent) {
	BitBtn_SetOnMouseDown(b.instance, fn)
}

// 设置鼠标进入事件。
//
// Set Mouse entry event.
func (b *TBitBtn) SetOnMouseEnter(fn TNotifyEvent) {
	BitBtn_SetOnMouseEnter(b.instance, fn)
}

// 设置鼠标离开事件。
//
// Set Mouse leave event.
func (b *TBitBtn) SetOnMouseLeave(fn TNotifyEvent) {
	BitBtn_SetOnMouseLeave(b.instance, fn)
}

// 设置鼠标移动事件。
func (b *TBitBtn) SetOnMouseMove(fn TMouseMoveEvent) {
	BitBtn_SetOnMouseMove(b.instance, fn)
}

// 设置鼠标抬起事件。
//
// Set Mouse lift event.
func (b *TBitBtn) SetOnMouseUp(fn TMouseEvent) {
	BitBtn_SetOnMouseUp(b.instance, fn)
}

// 获取依靠客户端总数。
func (b *TBitBtn) DockClientCount() int32 {
	return BitBtn_GetDockClientCount(b.instance)
}

// 获取停靠站点。
//
// Get Docking site.
func (b *TBitBtn) DockSite() bool {
	return BitBtn_GetDockSite(b.instance)
}

// 设置停靠站点。
//
// Set Docking site.
func (b *TBitBtn) SetDockSite(value bool) {
	BitBtn_SetDockSite(b.instance, value)
}

// 获取鼠标是否在客户端，仅lcl有效。
//
// Get Whether the mouse is on the client, only lcl is valid.
func (b *TBitBtn) MouseInClient() bool {
	return BitBtn_GetMouseInClient(b.instance)
}

// 获取当前停靠的可视总数。
//
// Get The total number of visible calls currently docked.
func (b *TBitBtn) VisibleDockClientCount() int32 {
	return BitBtn_GetVisibleDockClientCount(b.instance)
}

// 获取画刷对象。
//
// Get Brush.
func (b *TBitBtn) Brush() *TBrush {
	return AsBrush(BitBtn_GetBrush(b.instance))
}

// 获取子控件数。
//
// Get Number of child controls.
func (b *TBitBtn) ControlCount() int32 {
	return BitBtn_GetControlCount(b.instance)
}

// 获取控件句柄。
//
// Get Control handle.
func (b *TBitBtn) Handle() HWND {
	return BitBtn_GetHandle(b.instance)
}

// 获取父容器句柄。
//
// Get Parent container handle.
func (b *TBitBtn) ParentWindow() HWND {
	return BitBtn_GetParentWindow(b.instance)
}

// 设置父容器句柄。
//
// Set Parent container handle.
func (b *TBitBtn) SetParentWindow(value HWND) {
	BitBtn_SetParentWindow(b.instance, value)
}

func (b *TBitBtn) Showing() bool {
	return BitBtn_GetShowing(b.instance)
}

// 获取使用停靠管理。
func (b *TBitBtn) UseDockManager() bool {
	return BitBtn_GetUseDockManager(b.instance)
}

// 设置使用停靠管理。
func (b *TBitBtn) SetUseDockManager(value bool) {
	BitBtn_SetUseDockManager(b.instance, value)
}

func (b *TBitBtn) BoundsRect() TRect {
	return BitBtn_GetBoundsRect(b.instance)
}

func (b *TBitBtn) SetBoundsRect(value TRect) {
	BitBtn_SetBoundsRect(b.instance, value)
}

// 获取客户区高度。
//
// Get client height.
func (b *TBitBtn) ClientHeight() int32 {
	return BitBtn_GetClientHeight(b.instance)
}

// 设置客户区高度。
//
// Set client height.
func (b *TBitBtn) SetClientHeight(value int32) {
	BitBtn_SetClientHeight(b.instance, value)
}

func (b *TBitBtn) ClientOrigin() TPoint {
	return BitBtn_GetClientOrigin(b.instance)
}

// 获取客户区矩形。
//
// Get client rectangle.
func (b *TBitBtn) ClientRect() TRect {
	return BitBtn_GetClientRect(b.instance)
}

// 获取客户区宽度。
//
// Get client width.
func (b *TBitBtn) ClientWidth() int32 {
	return BitBtn_GetClientWidth(b.instance)
}

// 设置客户区宽度。
//
// Set client width.
func (b *TBitBtn) SetClientWidth(value int32) {
	BitBtn_SetClientWidth(b.instance, value)
}

// 获取控件状态。
//
// Get control state.
func (b *TBitBtn) ControlState() TControlState {
	return BitBtn_GetControlState(b.instance)
}

// 设置控件状态。
//
// Set control state.
func (b *TBitBtn) SetControlState(value TControlState) {
	BitBtn_SetControlState(b.instance, value)
}

// 获取控件样式。
//
// Get control style.
func (b *TBitBtn) ControlStyle() TControlStyle {
	return BitBtn_GetControlStyle(b.instance)
}

// 设置控件样式。
//
// Set control style.
func (b *TBitBtn) SetControlStyle(value TControlStyle) {
	BitBtn_SetControlStyle(b.instance, value)
}

func (b *TBitBtn) Floating() bool {
	return BitBtn_GetFloating(b.instance)
}

// 获取控件父容器。
//
// Get control parent container.
func (b *TBitBtn) Parent() *TWinControl {
	return AsWinControl(BitBtn_GetParent(b.instance))
}

// 设置控件父容器。
//
// Set control parent container.
func (b *TBitBtn) SetParent(value IWinControl) {
	BitBtn_SetParent(b.instance, CheckPtr(value))
}

// 获取左边位置。
//
// Get Left position.
func (b *TBitBtn) Left() int32 {
	return BitBtn_GetLeft(b.instance)
}

// 设置左边位置。
//
// Set Left position.
func (b *TBitBtn) SetLeft(value int32) {
	BitBtn_SetLeft(b.instance, value)
}

// 获取顶边位置。
//
// Get Top position.
func (b *TBitBtn) Top() int32 {
	return BitBtn_GetTop(b.instance)
}

// 设置顶边位置。
//
// Set Top position.
func (b *TBitBtn) SetTop(value int32) {
	BitBtn_SetTop(b.instance, value)
}

// 获取宽度。
//
// Get width.
func (b *TBitBtn) Width() int32 {
	return BitBtn_GetWidth(b.instance)
}

// 设置宽度。
//
// Set width.
func (b *TBitBtn) SetWidth(value int32) {
	BitBtn_SetWidth(b.instance, value)
}

// 获取高度。
//
// Get height.
func (b *TBitBtn) Height() int32 {
	return BitBtn_GetHeight(b.instance)
}

// 设置高度。
//
// Set height.
func (b *TBitBtn) SetHeight(value int32) {
	BitBtn_SetHeight(b.instance, value)
}

// 获取控件光标。
//
// Get control cursor.
func (b *TBitBtn) Cursor() TCursor {
	return BitBtn_GetCursor(b.instance)
}

// 设置控件光标。
//
// Set control cursor.
func (b *TBitBtn) SetCursor(value TCursor) {
	BitBtn_SetCursor(b.instance, value)
}

// 获取组件鼠标悬停提示。
//
// Get component mouse hints.
func (b *TBitBtn) Hint() string {
	return BitBtn_GetHint(b.instance)
}

// 设置组件鼠标悬停提示。
//
// Set component mouse hints.
func (b *TBitBtn) SetHint(value string) {
	BitBtn_SetHint(b.instance, value)
}

// 获取组件总数。
//
// Get the total number of components.
func (b *TBitBtn) ComponentCount() int32 {
	return BitBtn_GetComponentCount(b.instance)
}

// 获取组件索引。
//
// Get component index.
func (b *TBitBtn) ComponentIndex() int32 {
	return BitBtn_GetComponentIndex(b.instance)
}

// 设置组件索引。
//
// Set component index.
func (b *TBitBtn) SetComponentIndex(value int32) {
	BitBtn_SetComponentIndex(b.instance, value)
}

// 获取组件所有者。
//
// Get component owner.
func (b *TBitBtn) Owner() *TComponent {
	return AsComponent(BitBtn_GetOwner(b.instance))
}

// 获取组件名称。
//
// Get the component name.
func (b *TBitBtn) Name() string {
	return BitBtn_GetName(b.instance)
}

// 设置组件名称。
//
// Set the component name.
func (b *TBitBtn) SetName(value string) {
	BitBtn_SetName(b.instance, value)
}

// 获取对象标记。
//
// Get the control tag.
func (b *TBitBtn) Tag() int {
	return BitBtn_GetTag(b.instance)
}

// 设置对象标记。
//
// Set the control tag.
func (b *TBitBtn) SetTag(value int) {
	BitBtn_SetTag(b.instance, value)
}

// 获取左边锚点。
func (b *TBitBtn) AnchorSideLeft() *TAnchorSide {
	return AsAnchorSide(BitBtn_GetAnchorSideLeft(b.instance))
}

// 设置左边锚点。
func (b *TBitBtn) SetAnchorSideLeft(value *TAnchorSide) {
	BitBtn_SetAnchorSideLeft(b.instance, CheckPtr(value))
}

// 获取顶边锚点。
func (b *TBitBtn) AnchorSideTop() *TAnchorSide {
	return AsAnchorSide(BitBtn_GetAnchorSideTop(b.instance))
}

// 设置顶边锚点。
func (b *TBitBtn) SetAnchorSideTop(value *TAnchorSide) {
	BitBtn_SetAnchorSideTop(b.instance, CheckPtr(value))
}

// 获取右边锚点。
func (b *TBitBtn) AnchorSideRight() *TAnchorSide {
	return AsAnchorSide(BitBtn_GetAnchorSideRight(b.instance))
}

// 设置右边锚点。
func (b *TBitBtn) SetAnchorSideRight(value *TAnchorSide) {
	BitBtn_SetAnchorSideRight(b.instance, CheckPtr(value))
}

// 获取底边锚点。
func (b *TBitBtn) AnchorSideBottom() *TAnchorSide {
	return AsAnchorSide(BitBtn_GetAnchorSideBottom(b.instance))
}

// 设置底边锚点。
func (b *TBitBtn) SetAnchorSideBottom(value *TAnchorSide) {
	BitBtn_SetAnchorSideBottom(b.instance, CheckPtr(value))
}

func (b *TBitBtn) ChildSizing() *TControlChildSizing {
	return AsControlChildSizing(BitBtn_GetChildSizing(b.instance))
}

func (b *TBitBtn) SetChildSizing(value *TControlChildSizing) {
	BitBtn_SetChildSizing(b.instance, CheckPtr(value))
}

// 获取边框间距。
func (b *TBitBtn) BorderSpacing() *TControlBorderSpacing {
	return AsControlBorderSpacing(BitBtn_GetBorderSpacing(b.instance))
}

// 设置边框间距。
func (b *TBitBtn) SetBorderSpacing(value *TControlBorderSpacing) {
	BitBtn_SetBorderSpacing(b.instance, CheckPtr(value))
}

// 获取指定索引停靠客户端。
func (b *TBitBtn) DockClients(Index int32) *TControl {
	return AsControl(BitBtn_GetDockClients(b.instance, Index))
}

// 获取指定索引子控件。
func (b *TBitBtn) Controls(Index int32) *TControl {
	return AsControl(BitBtn_GetControls(b.instance, Index))
}

// 获取指定索引组件。
//
// Get the specified index component.
func (b *TBitBtn) Components(AIndex int32) *TComponent {
	return AsComponent(BitBtn_GetComponents(b.instance, AIndex))
}

// 获取锚侧面。
func (b *TBitBtn) AnchorSide(AKind TAnchorKind) *TAnchorSide {
	return AsAnchorSide(BitBtn_GetAnchorSide(b.instance, AKind))
}
