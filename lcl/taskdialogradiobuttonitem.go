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

type TTaskDialogRadioButtonItem struct {
	IObject
	instance unsafe.Pointer
}

// NewTaskDialogRadioButtonItem
//
// 创建一个新的对象。
//
// Create a new object.
func NewTaskDialogRadioButtonItem(AOwner *TCollection) *TTaskDialogRadioButtonItem {
	t := new(TTaskDialogRadioButtonItem)
	t.instance = unsafe.Pointer(TaskDialogRadioButtonItem_Create(CheckPtr(AOwner)))
	setFinalizer(t, (*TTaskDialogRadioButtonItem).Free)
	return t
}

// AsTaskDialogRadioButtonItem
//
// 动态转换一个已存在的对象实例。
//
// Dynamically convert an existing object instance.
func AsTaskDialogRadioButtonItem(obj interface{}) *TTaskDialogRadioButtonItem {
	instance := getInstance(obj)
	if instance == nullptr {
		return nil
	}
	return &TTaskDialogRadioButtonItem{instance: instance}
}

// Free
//
// 释放对象。
//
// Free object.
func (t *TTaskDialogRadioButtonItem) Free() {
	if t.instance != nullptr {
		TaskDialogRadioButtonItem_Free(t._instance())
		t.instance = nullptr
	}
}

func (t *TTaskDialogRadioButtonItem) _instance() uintptr {
	return uintptr(t.instance)
}

// Instance
//
// 返回对象实例指针。
//
// Return object instance pointer.
func (t *TTaskDialogRadioButtonItem) Instance() uintptr {
	return t._instance()
}

// UnsafeAddr
//
// 获取一个不安全的地址。
//
// Get an unsafe address.
func (t *TTaskDialogRadioButtonItem) UnsafeAddr() unsafe.Pointer {
	return t.instance
}

// IsValid
//
// 检测地址是否为空。
//
// Check if the address is empty.
func (t *TTaskDialogRadioButtonItem) IsValid() bool {
	return t.instance != nullptr
}

// Is
//
// 检测当前对象是否继承自目标对象。
//
// Checks whether the current object is inherited from the target object.
func (t *TTaskDialogRadioButtonItem) Is() TIs {
	return TIs(t._instance())
}

// As
//
// 动态转换当前对象为目标对象。
//
// Dynamically convert the current object to the target object.
//func (t *TTaskDialogRadioButtonItem) As() TAs {
//    return TAs(t._instance())
//}

// TTaskDialogRadioButtonItemClass
//
// 获取类信息指针。
//
// Get class information pointer.
func TTaskDialogRadioButtonItemClass() TClass {
	return TaskDialogRadioButtonItem_StaticClassType()
}

// GetNamePath
//
// 获取类名路径。
//
// Get the class name path.
func (t *TTaskDialogRadioButtonItem) GetNamePath() string {
	return TaskDialogRadioButtonItem_GetNamePath(t._instance())
}

// Assign
//
// 复制一个对象，如果对象实现了此方法的话。
//
// Copy an object, if the object implements this method.
func (t *TTaskDialogRadioButtonItem) Assign(Source IObject) {
	TaskDialogRadioButtonItem_Assign(t._instance(), CheckPtr(Source))
}

// ClassType
//
// 获取类的类型信息。
//
// Get class type information.
func (t *TTaskDialogRadioButtonItem) ClassType() TClass {
	return TaskDialogRadioButtonItem_ClassType(t._instance())
}

// ClassName
//
// 获取当前对象类名称。
//
// Get the current object class name.
func (t *TTaskDialogRadioButtonItem) ClassName() string {
	return TaskDialogRadioButtonItem_ClassName(t._instance())
}

// InstanceSize
//
// 获取当前对象实例大小。
//
// Get the current object instance size.
func (t *TTaskDialogRadioButtonItem) InstanceSize() int32 {
	return TaskDialogRadioButtonItem_InstanceSize(t._instance())
}

// InheritsFrom
//
// 判断当前类是否继承自指定类。
//
// Determine whether the current class inherits from the specified class.
func (t *TTaskDialogRadioButtonItem) InheritsFrom(AClass TClass) bool {
	return TaskDialogRadioButtonItem_InheritsFrom(t._instance(), AClass)
}

// Equals
//
// 与一个对象进行比较。
//
// Compare with an object.
func (t *TTaskDialogRadioButtonItem) Equals(Obj IObject) bool {
	return TaskDialogRadioButtonItem_Equals(t._instance(), CheckPtr(Obj))
}

// GetHashCode
//
// 获取类的哈希值。
//
// Get the hash value of the class.
func (t *TTaskDialogRadioButtonItem) GetHashCode() int32 {
	return TaskDialogRadioButtonItem_GetHashCode(t._instance())
}

// ToString
//
// 文本类信息。
//
// Text information.
func (t *TTaskDialogRadioButtonItem) ToString() string {
	return TaskDialogRadioButtonItem_ToString(t._instance())
}

// ModalResult
//
// 获取模态对话框显示结果。
func (t *TTaskDialogRadioButtonItem) ModalResult() TModalResult {
	return TaskDialogRadioButtonItem_GetModalResult(t._instance())
}

// SetModalResult
//
// 设置模态对话框显示结果。
func (t *TTaskDialogRadioButtonItem) SetModalResult(value TModalResult) {
	TaskDialogRadioButtonItem_SetModalResult(t._instance(), value)
}

// Caption
//
// 获取控件标题。
//
// Get the control title.
func (t *TTaskDialogRadioButtonItem) Caption() string {
	return TaskDialogRadioButtonItem_GetCaption(t._instance())
}

// SetCaption
//
// 设置控件标题。
//
// Set the control title.
func (t *TTaskDialogRadioButtonItem) SetCaption(value string) {
	TaskDialogRadioButtonItem_SetCaption(t._instance(), value)
}

func (t *TTaskDialogRadioButtonItem) Default() bool {
	return TaskDialogRadioButtonItem_GetDefault(t._instance())
}

func (t *TTaskDialogRadioButtonItem) SetDefault(value bool) {
	TaskDialogRadioButtonItem_SetDefault(t._instance(), value)
}

func (t *TTaskDialogRadioButtonItem) Collection() *TCollection {
	return AsCollection(TaskDialogRadioButtonItem_GetCollection(t._instance()))
}

func (t *TTaskDialogRadioButtonItem) SetCollection(value *TCollection) {
	TaskDialogRadioButtonItem_SetCollection(t._instance(), CheckPtr(value))
}

func (t *TTaskDialogRadioButtonItem) Index() int32 {
	return TaskDialogRadioButtonItem_GetIndex(t._instance())
}

func (t *TTaskDialogRadioButtonItem) SetIndex(value int32) {
	TaskDialogRadioButtonItem_SetIndex(t._instance(), value)
}

func (t *TTaskDialogRadioButtonItem) DisplayName() string {
	return TaskDialogRadioButtonItem_GetDisplayName(t._instance())
}

func (t *TTaskDialogRadioButtonItem) SetDisplayName(value string) {
	TaskDialogRadioButtonItem_SetDisplayName(t._instance(), value)
}
