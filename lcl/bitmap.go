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

type TBitmap struct {
	IBitmap
	instance unsafe.Pointer
}

// NewBitmap
//
// 创建一个新的对象。
//
// Create a new object.
func NewBitmap() *TBitmap {
	b := new(TBitmap)
	b.instance = unsafe.Pointer(Bitmap_Create())
	setFinalizer(b, (*TBitmap).Free)
	return b
}

// AsBitmap
//
// 动态转换一个已存在的对象实例。
//
// Dynamically convert an existing object instance.
func AsBitmap(obj interface{}) *TBitmap {
	instance := getInstance(obj)
	if instance == nullptr {
		return nil
	}
	return &TBitmap{instance: instance}
}

// Free
//
// 释放对象。
//
// Free object.
func (b *TBitmap) Free() {
	if b.instance != nullptr {
		Bitmap_Free(b._instance())
		b.instance = nullptr
	}
}

func (b *TBitmap) _instance() uintptr {
	return uintptr(b.instance)
}

// Instance
//
// 返回对象实例指针。
//
// Return object instance pointer.
func (b *TBitmap) Instance() uintptr {
	return b._instance()
}

// UnsafeAddr
//
// 获取一个不安全的地址。
//
// Get an unsafe address.
func (b *TBitmap) UnsafeAddr() unsafe.Pointer {
	return b.instance
}

// IsValid
//
// 检测地址是否为空。
//
// Check if the address is empty.
func (b *TBitmap) IsValid() bool {
	return b.instance != nullptr
}

// Is
//
// 检测当前对象是否继承自目标对象。
//
// Checks whether the current object is inherited from the target object.
func (b *TBitmap) Is() TIs {
	return TIs(b._instance())
}

// As
//
// 动态转换当前对象为目标对象。
//
// Dynamically convert the current object to the target object.
//func (b *TBitmap) As() TAs {
//    return TAs(b._instance())
//}

// TBitmapClass
//
// 获取类信息指针。
//
// Get class information pointer.
func TBitmapClass() TClass {
	return Bitmap_StaticClassType()
}

// LoadFromDevice
//
// 从设备驱动中加载Bitmap。
//
// Load the Bitmap from the device driver.
func (b *TBitmap) LoadFromDevice(ADc HDC) {
	Bitmap_LoadFromDevice(b._instance(), ADc)
}

// EndUpdate
//
// 用于ScanLine属性，aStreamIsValid 默认为 false。
//
// Used for ScanLine property, aStreamIsValid defaults to false.
func (b *TBitmap) EndUpdate(AStreamIsValid bool) {
	Bitmap_EndUpdate(b._instance(), AStreamIsValid)
}

// BeginUpdate
//
// 用于ScanLine属性，aCanvasOnly 默认为 false。
//
// Used for ScanLine properties, aCanvasOnly defaults to false.
func (b *TBitmap) BeginUpdate(ACanvasOnly bool) {
	Bitmap_BeginUpdate(b._instance(), ACanvasOnly)
}

// Clear
//
// 清除bitmap数据。
//
// Clear bitmap data.
func (b *TBitmap) Clear() {
	Bitmap_Clear(b._instance())
}

// Assign
//
// 复制一个对象，如果对象实现了此方法的话。
//
// Copy an object, if the object implements this method.
func (b *TBitmap) Assign(Source IObject) {
	Bitmap_Assign(b._instance(), CheckPtr(Source))
}

func (b *TBitmap) FreeImage() {
	Bitmap_FreeImage(b._instance())
}

// HandleAllocated
//
// 句柄是否已经分配。
//
// Is the handle already allocated.
func (b *TBitmap) HandleAllocated() bool {
	return Bitmap_HandleAllocated(b._instance())
}

// LoadFromStream
//
// 文件流加载。
func (b *TBitmap) LoadFromStream(Stream IStream) {
	Bitmap_LoadFromStream(b._instance(), CheckPtr(Stream))
}

// SaveToStream
//
// 保存至流。
func (b *TBitmap) SaveToStream(Stream IStream) {
	Bitmap_SaveToStream(b._instance(), CheckPtr(Stream))
}

func (b *TBitmap) SetSize(AWidth int32, AHeight int32) {
	Bitmap_SetSize(b._instance(), AWidth, AHeight)
}

func (b *TBitmap) LoadFromResourceName(Instance uintptr, ResName string) {
	Bitmap_LoadFromResourceName(b._instance(), Instance, ResName)
}

func (b *TBitmap) LoadFromResourceID(Instance uintptr, ResID int32) {
	Bitmap_LoadFromResourceID(b._instance(), Instance, ResID)
}

// Equals
//
// 与一个对象进行比较。
//
// Compare with an object.
func (b *TBitmap) Equals(Obj IObject) bool {
	return Bitmap_Equals(b._instance(), CheckPtr(Obj))
}

// LoadFromFile
//
// 从文件加载。
func (b *TBitmap) LoadFromFile(Filename string) {
	Bitmap_LoadFromFile(b._instance(), Filename)
}

// SaveToFile
//
// 保存至文件。
func (b *TBitmap) SaveToFile(Filename string) {
	Bitmap_SaveToFile(b._instance(), Filename)
}

// GetNamePath
//
// 获取类名路径。
//
// Get the class name path.
func (b *TBitmap) GetNamePath() string {
	return Bitmap_GetNamePath(b._instance())
}

// ClassType
//
// 获取类的类型信息。
//
// Get class type information.
func (b *TBitmap) ClassType() TClass {
	return Bitmap_ClassType(b._instance())
}

// ClassName
//
// 获取当前对象类名称。
//
// Get the current object class name.
func (b *TBitmap) ClassName() string {
	return Bitmap_ClassName(b._instance())
}

// InstanceSize
//
// 获取当前对象实例大小。
//
// Get the current object instance size.
func (b *TBitmap) InstanceSize() int32 {
	return Bitmap_InstanceSize(b._instance())
}

// InheritsFrom
//
// 判断当前类是否继承自指定类。
//
// Determine whether the current class inherits from the specified class.
func (b *TBitmap) InheritsFrom(AClass TClass) bool {
	return Bitmap_InheritsFrom(b._instance(), AClass)
}

// GetHashCode
//
// 获取类的哈希值。
//
// Get the hash value of the class.
func (b *TBitmap) GetHashCode() int32 {
	return Bitmap_GetHashCode(b._instance())
}

// ToString
//
// 文本类信息。
//
// Text information.
func (b *TBitmap) ToString() string {
	return Bitmap_ToString(b._instance())
}

// Canvas
//
// 获取画布。
func (b *TBitmap) Canvas() *TCanvas {
	return AsCanvas(Bitmap_GetCanvas(b._instance()))
}

// Handle
//
// 获取控件句柄。
//
// Get Control handle.
func (b *TBitmap) Handle() HBITMAP {
	return Bitmap_GetHandle(b._instance())
}

// SetHandle
//
// 设置控件句柄。
//
// Set Control handle.
func (b *TBitmap) SetHandle(value HBITMAP) {
	Bitmap_SetHandle(b._instance(), value)
}

func (b *TBitmap) HandleType() TBitmapHandleType {
	return Bitmap_GetHandleType(b._instance())
}

func (b *TBitmap) SetHandleType(value TBitmapHandleType) {
	Bitmap_SetHandleType(b._instance(), value)
}

func (b *TBitmap) MaskHandle() HBITMAP {
	return Bitmap_GetMaskHandle(b._instance())
}

func (b *TBitmap) SetMaskHandle(value HBITMAP) {
	Bitmap_SetMaskHandle(b._instance(), value)
}

func (b *TBitmap) PixelFormat() TPixelFormat {
	return Bitmap_GetPixelFormat(b._instance())
}

func (b *TBitmap) SetPixelFormat(value TPixelFormat) {
	Bitmap_SetPixelFormat(b._instance(), value)
}

func (b *TBitmap) TransparentMode() TTransparentMode {
	return Bitmap_GetTransparentMode(b._instance())
}

func (b *TBitmap) SetTransparentMode(value TTransparentMode) {
	Bitmap_SetTransparentMode(b._instance(), value)
}

func (b *TBitmap) Empty() bool {
	return Bitmap_GetEmpty(b._instance())
}

// Height
//
// 获取高度。
//
// Get height.
func (b *TBitmap) Height() int32 {
	return Bitmap_GetHeight(b._instance())
}

// SetHeight
//
// 设置高度。
//
// Set height.
func (b *TBitmap) SetHeight(value int32) {
	Bitmap_SetHeight(b._instance(), value)
}

// Modified
//
// 获取修改。
//
// Get modified.
func (b *TBitmap) Modified() bool {
	return Bitmap_GetModified(b._instance())
}

// SetModified
//
// 设置修改。
//
// Set modified.
func (b *TBitmap) SetModified(value bool) {
	Bitmap_SetModified(b._instance(), value)
}

func (b *TBitmap) Palette() HPALETTE {
	return Bitmap_GetPalette(b._instance())
}

func (b *TBitmap) SetPalette(value HPALETTE) {
	Bitmap_SetPalette(b._instance(), value)
}

func (b *TBitmap) PaletteModified() bool {
	return Bitmap_GetPaletteModified(b._instance())
}

func (b *TBitmap) SetPaletteModified(value bool) {
	Bitmap_SetPaletteModified(b._instance(), value)
}

// Transparent
//
// 获取透明。
//
// Get transparent.
func (b *TBitmap) Transparent() bool {
	return Bitmap_GetTransparent(b._instance())
}

// SetTransparent
//
// 设置透明。
//
// Set transparent.
func (b *TBitmap) SetTransparent(value bool) {
	Bitmap_SetTransparent(b._instance(), value)
}

// Width
//
// 获取宽度。
//
// Get width.
func (b *TBitmap) Width() int32 {
	return Bitmap_GetWidth(b._instance())
}

// SetWidth
//
// 设置宽度。
//
// Set width.
func (b *TBitmap) SetWidth(value int32) {
	Bitmap_SetWidth(b._instance(), value)
}

// SetOnChange
//
// 设置改变事件。
//
// Set changed event.
func (b *TBitmap) SetOnChange(fn TNotifyEvent) {
	Bitmap_SetOnChange(b._instance(), fn)
}

func (b *TBitmap) ScanLine(Row int32) uintptr {
	return Bitmap_GetScanLine(b._instance(), Row)
}
