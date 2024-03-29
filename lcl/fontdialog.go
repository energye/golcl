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

type TFontDialog struct {
	IComponent
	instance unsafe.Pointer
}

// NewFontDialog
//
// 创建一个新的对象。
//
// Create a new object.
func NewFontDialog(owner IComponent) *TFontDialog {
	f := new(TFontDialog)
	f.instance = unsafe.Pointer(FontDialog_Create(CheckPtr(owner)))
	return f
}

// AsFontDialog
//
// 动态转换一个已存在的对象实例。
//
// Dynamically convert an existing object instance.
func AsFontDialog(obj interface{}) *TFontDialog {
	instance := getInstance(obj)
	if instance == nullptr {
		return nil
	}
	return &TFontDialog{instance: instance}
}

// Free
//
// 释放对象。
//
// Free object.
func (f *TFontDialog) Free() {
	if f.instance != nullptr {
		FontDialog_Free(f._instance())
		f.instance = nullptr
	}
}

func (f *TFontDialog) _instance() uintptr {
	return uintptr(f.instance)
}

// Instance
//
// 返回对象实例指针。
//
// Return object instance pointer.
func (f *TFontDialog) Instance() uintptr {
	return f._instance()
}

// UnsafeAddr
//
// 获取一个不安全的地址。
//
// Get an unsafe address.
func (f *TFontDialog) UnsafeAddr() unsafe.Pointer {
	return f.instance
}

// IsValid
//
// 检测地址是否为空。
//
// Check if the address is empty.
func (f *TFontDialog) IsValid() bool {
	return f.instance != nullptr
}

// Is
//
// 检测当前对象是否继承自目标对象。
//
// Checks whether the current object is inherited from the target object.
func (f *TFontDialog) Is() TIs {
	return TIs(f._instance())
}

// As
//
// 动态转换当前对象为目标对象。
//
// Dynamically convert the current object to the target object.
//func (f *TFontDialog) As() TAs {
//    return TAs(f._instance())
//}

// TFontDialogClass
//
// 获取类信息指针。
//
// Get class information pointer.
func TFontDialogClass() TClass {
	return FontDialog_StaticClassType()
}

// Execute
//
// 执行。
func (f *TFontDialog) Execute() bool {
	return FontDialog_Execute(f._instance())
}

// FindComponent
//
// 查找指定名称的组件。
//
// Find the component with the specified name.
func (f *TFontDialog) FindComponent(AName string) *TComponent {
	return AsComponent(FontDialog_FindComponent(f._instance(), AName))
}

// GetNamePath
//
// 获取类名路径。
//
// Get the class name path.
func (f *TFontDialog) GetNamePath() string {
	return FontDialog_GetNamePath(f._instance())
}

// HasParent
//
// 是否有父容器。
//
// Is there a parent container.
func (f *TFontDialog) HasParent() bool {
	return FontDialog_HasParent(f._instance())
}

// Assign
//
// 复制一个对象，如果对象实现了此方法的话。
//
// Copy an object, if the object implements this method.
func (f *TFontDialog) Assign(Source IObject) {
	FontDialog_Assign(f._instance(), CheckPtr(Source))
}

// ClassType
//
// 获取类的类型信息。
//
// Get class type information.
func (f *TFontDialog) ClassType() TClass {
	return FontDialog_ClassType(f._instance())
}

// ClassName
//
// 获取当前对象类名称。
//
// Get the current object class name.
func (f *TFontDialog) ClassName() string {
	return FontDialog_ClassName(f._instance())
}

// InstanceSize
//
// 获取当前对象实例大小。
//
// Get the current object instance size.
func (f *TFontDialog) InstanceSize() int32 {
	return FontDialog_InstanceSize(f._instance())
}

// InheritsFrom
//
// 判断当前类是否继承自指定类。
//
// Determine whether the current class inherits from the specified class.
func (f *TFontDialog) InheritsFrom(AClass TClass) bool {
	return FontDialog_InheritsFrom(f._instance(), AClass)
}

// Equals
//
// 与一个对象进行比较。
//
// Compare with an object.
func (f *TFontDialog) Equals(Obj IObject) bool {
	return FontDialog_Equals(f._instance(), CheckPtr(Obj))
}

// GetHashCode
//
// 获取类的哈希值。
//
// Get the hash value of the class.
func (f *TFontDialog) GetHashCode() int32 {
	return FontDialog_GetHashCode(f._instance())
}

// ToString
//
// 文本类信息。
//
// Text information.
func (f *TFontDialog) ToString() string {
	return FontDialog_ToString(f._instance())
}

// Font
//
// 获取字体。
//
// Get Font.
func (f *TFontDialog) Font() *TFont {
	return AsFont(FontDialog_GetFont(f._instance()))
}

// SetFont
//
// 设置字体。
//
// Set Font.
func (f *TFontDialog) SetFont(value *TFont) {
	FontDialog_SetFont(f._instance(), CheckPtr(value))
}

func (f *TFontDialog) Options() TFontDialogOptions {
	return FontDialog_GetOptions(f._instance())
}

func (f *TFontDialog) SetOptions(value TFontDialogOptions) {
	FontDialog_SetOptions(f._instance(), value)
}

// Handle
//
// 获取控件句柄。
//
// Get Control handle.
func (f *TFontDialog) Handle() HWND {
	return FontDialog_GetHandle(f._instance())
}

func (f *TFontDialog) SetOnClose(fn TNotifyEvent) {
	FontDialog_SetOnClose(f._instance(), fn)
}

// SetOnShow
//
// 设置显示事件。
func (f *TFontDialog) SetOnShow(fn TNotifyEvent) {
	FontDialog_SetOnShow(f._instance(), fn)
}

// ComponentCount
//
// 获取组件总数。
//
// Get the total number of components.
func (f *TFontDialog) ComponentCount() int32 {
	return FontDialog_GetComponentCount(f._instance())
}

// ComponentIndex
//
// 获取组件索引。
//
// Get component index.
func (f *TFontDialog) ComponentIndex() int32 {
	return FontDialog_GetComponentIndex(f._instance())
}

// SetComponentIndex
//
// 设置组件索引。
//
// Set component index.
func (f *TFontDialog) SetComponentIndex(value int32) {
	FontDialog_SetComponentIndex(f._instance(), value)
}

// Owner
//
// 获取组件所有者。
//
// Get component owner.
func (f *TFontDialog) Owner() *TComponent {
	return AsComponent(FontDialog_GetOwner(f._instance()))
}

// Name
//
// 获取组件名称。
//
// Get the component name.
func (f *TFontDialog) Name() string {
	return FontDialog_GetName(f._instance())
}

// SetName
//
// 设置组件名称。
//
// Set the component name.
func (f *TFontDialog) SetName(value string) {
	FontDialog_SetName(f._instance(), value)
}

// Tag
//
// 获取对象标记。
//
// Get the control tag.
func (f *TFontDialog) Tag() int {
	return FontDialog_GetTag(f._instance())
}

// SetTag
//
// 设置对象标记。
//
// Set the control tag.
func (f *TFontDialog) SetTag(value int) {
	FontDialog_SetTag(f._instance(), value)
}

// Components
//
// 获取指定索引组件。
//
// Get the specified index component.
func (f *TFontDialog) Components(AIndex int32) *TComponent {
	return AsComponent(FontDialog_GetComponents(f._instance(), AIndex))
}
