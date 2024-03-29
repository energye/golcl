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

type TBrush struct {
	IObject
	instance unsafe.Pointer
}

// NewBrush
//
// 创建一个新的对象。
//
// Create a new object.
func NewBrush() *TBrush {
	b := new(TBrush)
	b.instance = unsafe.Pointer(Brush_Create())
	setFinalizer(b, (*TBrush).Free)
	return b
}

// AsBrush
//
// 动态转换一个已存在的对象实例。
//
// Dynamically convert an existing object instance.
func AsBrush(obj interface{}) *TBrush {
	instance := getInstance(obj)
	if instance == nullptr {
		return nil
	}
	return &TBrush{instance: instance}
}

// Free
//
// 释放对象。
//
// Free object.
func (b *TBrush) Free() {
	if b.instance != nullptr {
		Brush_Free(b._instance())
		b.instance = nullptr
	}
}

func (b *TBrush) _instance() uintptr {
	return uintptr(b.instance)
}

// Instance
//
// 返回对象实例指针。
//
// Return object instance pointer.
func (b *TBrush) Instance() uintptr {
	return b._instance()
}

// UnsafeAddr
//
// 获取一个不安全的地址。
//
// Get an unsafe address.
func (b *TBrush) UnsafeAddr() unsafe.Pointer {
	return b.instance
}

// IsValid
//
// 检测地址是否为空。
//
// Check if the address is empty.
func (b *TBrush) IsValid() bool {
	return b.instance != nullptr
}

// Is
//
// 检测当前对象是否继承自目标对象。
//
// Checks whether the current object is inherited from the target object.
func (b *TBrush) Is() TIs {
	return TIs(b._instance())
}

// As
//
// 动态转换当前对象为目标对象。
//
// Dynamically convert the current object to the target object.
//func (b *TBrush) As() TAs {
//    return TAs(b._instance())
//}

// TBrushClass
//
// 获取类信息指针。
//
// Get class information pointer.
func TBrushClass() TClass {
	return Brush_StaticClassType()
}

// Assign
//
// 复制一个对象，如果对象实现了此方法的话。
//
// Copy an object, if the object implements this method.
func (b *TBrush) Assign(Source IObject) {
	Brush_Assign(b._instance(), CheckPtr(Source))
}

// GetNamePath
//
// 获取类名路径。
//
// Get the class name path.
func (b *TBrush) GetNamePath() string {
	return Brush_GetNamePath(b._instance())
}

// ClassType
//
// 获取类的类型信息。
//
// Get class type information.
func (b *TBrush) ClassType() TClass {
	return Brush_ClassType(b._instance())
}

// ClassName
//
// 获取当前对象类名称。
//
// Get the current object class name.
func (b *TBrush) ClassName() string {
	return Brush_ClassName(b._instance())
}

// InstanceSize
//
// 获取当前对象实例大小。
//
// Get the current object instance size.
func (b *TBrush) InstanceSize() int32 {
	return Brush_InstanceSize(b._instance())
}

// InheritsFrom
//
// 判断当前类是否继承自指定类。
//
// Determine whether the current class inherits from the specified class.
func (b *TBrush) InheritsFrom(AClass TClass) bool {
	return Brush_InheritsFrom(b._instance(), AClass)
}

// Equals
//
// 与一个对象进行比较。
//
// Compare with an object.
func (b *TBrush) Equals(Obj IObject) bool {
	return Brush_Equals(b._instance(), CheckPtr(Obj))
}

// GetHashCode
//
// 获取类的哈希值。
//
// Get the hash value of the class.
func (b *TBrush) GetHashCode() int32 {
	return Brush_GetHashCode(b._instance())
}

// ToString
//
// 文本类信息。
//
// Text information.
func (b *TBrush) ToString() string {
	return Brush_ToString(b._instance())
}

func (b *TBrush) Bitmap() *TBitmap {
	return AsBitmap(Brush_GetBitmap(b._instance()))
}

func (b *TBrush) SetBitmap(value *TBitmap) {
	Brush_SetBitmap(b._instance(), CheckPtr(value))
}

// Handle
//
// 获取控件句柄。
//
// Get Control handle.
func (b *TBrush) Handle() HBRUSH {
	return Brush_GetHandle(b._instance())
}

// SetHandle
//
// 设置控件句柄。
//
// Set Control handle.
func (b *TBrush) SetHandle(value HBRUSH) {
	Brush_SetHandle(b._instance(), value)
}

// Color
//
// 获取颜色。
//
// Get color.
func (b *TBrush) Color() TColor {
	return Brush_GetColor(b._instance())
}

// SetColor
//
// 设置颜色。
//
// Set color.
func (b *TBrush) SetColor(value TColor) {
	Brush_SetColor(b._instance(), value)
}

func (b *TBrush) Style() TBrushStyle {
	return Brush_GetStyle(b._instance())
}

func (b *TBrush) SetStyle(value TBrushStyle) {
	Brush_SetStyle(b._instance(), value)
}

// SetOnChange
//
// 设置改变事件。
//
// Set changed event.
func (b *TBrush) SetOnChange(fn TNotifyEvent) {
	Brush_SetOnChange(b._instance(), fn)
}
