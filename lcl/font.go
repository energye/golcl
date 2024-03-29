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

type TFont struct {
	IObject
	instance unsafe.Pointer
}

// NewFont
//
// 创建一个新的对象。
//
// Create a new object.
func NewFont() *TFont {
	f := new(TFont)
	f.instance = unsafe.Pointer(Font_Create())
	setFinalizer(f, (*TFont).Free)
	return f
}

// AsFont
//
// 动态转换一个已存在的对象实例。
//
// Dynamically convert an existing object instance.
func AsFont(obj interface{}) *TFont {
	instance := getInstance(obj)
	if instance == nullptr {
		return nil
	}
	return &TFont{instance: instance}
}

// Free
//
// 释放对象。
//
// Free object.
func (f *TFont) Free() {
	if f.instance != nullptr {
		Font_Free(f._instance())
		f.instance = nullptr
	}
}

func (f *TFont) _instance() uintptr {
	return uintptr(f.instance)
}

// Instance
//
// 返回对象实例指针。
//
// Return object instance pointer.
func (f *TFont) Instance() uintptr {
	return f._instance()
}

// UnsafeAddr
//
// 获取一个不安全的地址。
//
// Get an unsafe address.
func (f *TFont) UnsafeAddr() unsafe.Pointer {
	return f.instance
}

// IsValid
//
// 检测地址是否为空。
//
// Check if the address is empty.
func (f *TFont) IsValid() bool {
	return f.instance != nullptr
}

// Is
//
// 检测当前对象是否继承自目标对象。
//
// Checks whether the current object is inherited from the target object.
func (f *TFont) Is() TIs {
	return TIs(f._instance())
}

// As
//
// 动态转换当前对象为目标对象。
//
// Dynamically convert the current object to the target object.
//func (f *TFont) As() TAs {
//    return TAs(f._instance())
//}

// TFontClass
//
// 获取类信息指针。
//
// Get class information pointer.
func TFontClass() TClass {
	return Font_StaticClassType()
}

// Assign
//
// 复制一个对象，如果对象实现了此方法的话。
//
// Copy an object, if the object implements this method.
func (f *TFont) Assign(Source IObject) {
	Font_Assign(f._instance(), CheckPtr(Source))
}

// HandleAllocated
//
// 句柄是否已经分配。
//
// Is the handle already allocated.
func (f *TFont) HandleAllocated() bool {
	return Font_HandleAllocated(f._instance())
}

// GetNamePath
//
// 获取类名路径。
//
// Get the class name path.
func (f *TFont) GetNamePath() string {
	return Font_GetNamePath(f._instance())
}

// ClassType
//
// 获取类的类型信息。
//
// Get class type information.
func (f *TFont) ClassType() TClass {
	return Font_ClassType(f._instance())
}

// ClassName
//
// 获取当前对象类名称。
//
// Get the current object class name.
func (f *TFont) ClassName() string {
	return Font_ClassName(f._instance())
}

// InstanceSize
//
// 获取当前对象实例大小。
//
// Get the current object instance size.
func (f *TFont) InstanceSize() int32 {
	return Font_InstanceSize(f._instance())
}

// InheritsFrom
//
// 判断当前类是否继承自指定类。
//
// Determine whether the current class inherits from the specified class.
func (f *TFont) InheritsFrom(AClass TClass) bool {
	return Font_InheritsFrom(f._instance(), AClass)
}

// Equals
//
// 与一个对象进行比较。
//
// Compare with an object.
func (f *TFont) Equals(Obj IObject) bool {
	return Font_Equals(f._instance(), CheckPtr(Obj))
}

// GetHashCode
//
// 获取类的哈希值。
//
// Get the hash value of the class.
func (f *TFont) GetHashCode() int32 {
	return Font_GetHashCode(f._instance())
}

// ToString
//
// 文本类信息。
//
// Text information.
func (f *TFont) ToString() string {
	return Font_ToString(f._instance())
}

// Handle
//
// 获取控件句柄。
//
// Get Control handle.
func (f *TFont) Handle() HFONT {
	return Font_GetHandle(f._instance())
}

// SetHandle
//
// 设置控件句柄。
//
// Set Control handle.
func (f *TFont) SetHandle(value HFONT) {
	Font_SetHandle(f._instance(), value)
}

func (f *TFont) PixelsPerInch() int32 {
	return Font_GetPixelsPerInch(f._instance())
}

func (f *TFont) SetPixelsPerInch(value int32) {
	Font_SetPixelsPerInch(f._instance(), value)
}

func (f *TFont) Charset() TFontCharset {
	return Font_GetCharset(f._instance())
}

func (f *TFont) SetCharset(value TFontCharset) {
	Font_SetCharset(f._instance(), value)
}

// Color
//
// 获取颜色。
//
// Get color.
func (f *TFont) Color() TColor {
	return Font_GetColor(f._instance())
}

// SetColor
//
// 设置颜色。
//
// Set color.
func (f *TFont) SetColor(value TColor) {
	Font_SetColor(f._instance(), value)
}

// Height
//
// 获取高度。
//
// Get height.
func (f *TFont) Height() int32 {
	return Font_GetHeight(f._instance())
}

// SetHeight
//
// 设置高度。
//
// Set height.
func (f *TFont) SetHeight(value int32) {
	Font_SetHeight(f._instance(), value)
}

// Name
//
// 获取组件名称。
//
// Get the component name.
func (f *TFont) Name() string {
	return Font_GetName(f._instance())
}

// SetName
//
// 设置组件名称。
//
// Set the component name.
func (f *TFont) SetName(value string) {
	Font_SetName(f._instance(), value)
}

func (f *TFont) Orientation() int32 {
	return Font_GetOrientation(f._instance())
}

func (f *TFont) SetOrientation(value int32) {
	Font_SetOrientation(f._instance(), value)
}

func (f *TFont) Pitch() TFontPitch {
	return Font_GetPitch(f._instance())
}

func (f *TFont) SetPitch(value TFontPitch) {
	Font_SetPitch(f._instance(), value)
}

func (f *TFont) Size() int32 {
	return Font_GetSize(f._instance())
}

func (f *TFont) SetSize(value int32) {
	Font_SetSize(f._instance(), value)
}

func (f *TFont) Style() TFontStyles {
	return Font_GetStyle(f._instance())
}

func (f *TFont) SetStyle(value TFontStyles) {
	Font_SetStyle(f._instance(), value)
}

func (f *TFont) Quality() TFontQuality {
	return Font_GetQuality(f._instance())
}

func (f *TFont) SetQuality(value TFontQuality) {
	Font_SetQuality(f._instance(), value)
}

// SetOnChange
//
// 设置改变事件。
//
// Set changed event.
func (f *TFont) SetOnChange(fn TNotifyEvent) {
	Font_SetOnChange(f._instance(), fn)
}
