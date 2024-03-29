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

type TPen struct {
	IObject
	instance unsafe.Pointer
}

// NewPen
//
// 创建一个新的对象。
//
// Create a new object.
func NewPen() *TPen {
	p := new(TPen)
	p.instance = unsafe.Pointer(Pen_Create())
	setFinalizer(p, (*TPen).Free)
	return p
}

// AsPen
//
// 动态转换一个已存在的对象实例。
//
// Dynamically convert an existing object instance.
func AsPen(obj interface{}) *TPen {
	instance := getInstance(obj)
	if instance == nullptr {
		return nil
	}
	return &TPen{instance: instance}
}

// Free
//
// 释放对象。
//
// Free object.
func (p *TPen) Free() {
	if p.instance != nullptr {
		Pen_Free(p._instance())
		p.instance = nullptr
	}
}

func (p *TPen) _instance() uintptr {
	return uintptr(p.instance)
}

// Instance
//
// 返回对象实例指针。
//
// Return object instance pointer.
func (p *TPen) Instance() uintptr {
	return p._instance()
}

// UnsafeAddr
//
// 获取一个不安全的地址。
//
// Get an unsafe address.
func (p *TPen) UnsafeAddr() unsafe.Pointer {
	return p.instance
}

// IsValid
//
// 检测地址是否为空。
//
// Check if the address is empty.
func (p *TPen) IsValid() bool {
	return p.instance != nullptr
}

// Is
//
// 检测当前对象是否继承自目标对象。
//
// Checks whether the current object is inherited from the target object.
func (p *TPen) Is() TIs {
	return TIs(p._instance())
}

// As
//
// 动态转换当前对象为目标对象。
//
// Dynamically convert the current object to the target object.
//func (p *TPen) As() TAs {
//    return TAs(p._instance())
//}

// TPenClass
//
// 获取类信息指针。
//
// Get class information pointer.
func TPenClass() TClass {
	return Pen_StaticClassType()
}

// Assign
//
// 复制一个对象，如果对象实现了此方法的话。
//
// Copy an object, if the object implements this method.
func (p *TPen) Assign(Source IObject) {
	Pen_Assign(p._instance(), CheckPtr(Source))
}

// GetNamePath
//
// 获取类名路径。
//
// Get the class name path.
func (p *TPen) GetNamePath() string {
	return Pen_GetNamePath(p._instance())
}

// ClassType
//
// 获取类的类型信息。
//
// Get class type information.
func (p *TPen) ClassType() TClass {
	return Pen_ClassType(p._instance())
}

// ClassName
//
// 获取当前对象类名称。
//
// Get the current object class name.
func (p *TPen) ClassName() string {
	return Pen_ClassName(p._instance())
}

// InstanceSize
//
// 获取当前对象实例大小。
//
// Get the current object instance size.
func (p *TPen) InstanceSize() int32 {
	return Pen_InstanceSize(p._instance())
}

// InheritsFrom
//
// 判断当前类是否继承自指定类。
//
// Determine whether the current class inherits from the specified class.
func (p *TPen) InheritsFrom(AClass TClass) bool {
	return Pen_InheritsFrom(p._instance(), AClass)
}

// Equals
//
// 与一个对象进行比较。
//
// Compare with an object.
func (p *TPen) Equals(Obj IObject) bool {
	return Pen_Equals(p._instance(), CheckPtr(Obj))
}

// GetHashCode
//
// 获取类的哈希值。
//
// Get the hash value of the class.
func (p *TPen) GetHashCode() int32 {
	return Pen_GetHashCode(p._instance())
}

// ToString
//
// 文本类信息。
//
// Text information.
func (p *TPen) ToString() string {
	return Pen_ToString(p._instance())
}

// Handle
//
// 获取控件句柄。
//
// Get Control handle.
func (p *TPen) Handle() HPEN {
	return Pen_GetHandle(p._instance())
}

// SetHandle
//
// 设置控件句柄。
//
// Set Control handle.
func (p *TPen) SetHandle(value HPEN) {
	Pen_SetHandle(p._instance(), value)
}

// Color
//
// 获取颜色。
//
// Get color.
func (p *TPen) Color() TColor {
	return Pen_GetColor(p._instance())
}

// SetColor
//
// 设置颜色。
//
// Set color.
func (p *TPen) SetColor(value TColor) {
	Pen_SetColor(p._instance(), value)
}

func (p *TPen) Mode() TPenMode {
	return Pen_GetMode(p._instance())
}

func (p *TPen) SetMode(value TPenMode) {
	Pen_SetMode(p._instance(), value)
}

func (p *TPen) Style() TPenStyle {
	return Pen_GetStyle(p._instance())
}

func (p *TPen) SetStyle(value TPenStyle) {
	Pen_SetStyle(p._instance(), value)
}

// Width
//
// 获取宽度。
//
// Get width.
func (p *TPen) Width() int32 {
	return Pen_GetWidth(p._instance())
}

// SetWidth
//
// 设置宽度。
//
// Set width.
func (p *TPen) SetWidth(value int32) {
	Pen_SetWidth(p._instance(), value)
}

// SetOnChange
//
// 设置改变事件。
//
// Set changed event.
func (p *TPen) SetOnChange(fn TNotifyEvent) {
	Pen_SetOnChange(p._instance(), fn)
}
