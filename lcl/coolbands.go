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

type TCoolBands struct {
	IObject
	instance unsafe.Pointer
}

// NewCoolBands
//
// 创建一个新的对象。
//
// Create a new object.
func NewCoolBands(AOwner *TCoolBar) *TCoolBands {
	c := new(TCoolBands)
	c.instance = unsafe.Pointer(CoolBands_Create(CheckPtr(AOwner)))
	setFinalizer(c, (*TCoolBands).Free)
	return c
}

// AsCoolBands
//
// 动态转换一个已存在的对象实例。
//
// Dynamically convert an existing object instance.
func AsCoolBands(obj interface{}) *TCoolBands {
	instance := getInstance(obj)
	if instance == nullptr {
		return nil
	}
	return &TCoolBands{instance: instance}
}

// Free
//
// 释放对象。
//
// Free object.
func (c *TCoolBands) Free() {
	if c.instance != nullptr {
		CoolBands_Free(c._instance())
		c.instance = nullptr
	}
}

func (c *TCoolBands) _instance() uintptr {
	return uintptr(c.instance)
}

// Instance
//
// 返回对象实例指针。
//
// Return object instance pointer.
func (c *TCoolBands) Instance() uintptr {
	return c._instance()
}

// UnsafeAddr
//
// 获取一个不安全的地址。
//
// Get an unsafe address.
func (c *TCoolBands) UnsafeAddr() unsafe.Pointer {
	return c.instance
}

// IsValid
//
// 检测地址是否为空。
//
// Check if the address is empty.
func (c *TCoolBands) IsValid() bool {
	return c.instance != nullptr
}

// Is
//
// 检测当前对象是否继承自目标对象。
//
// Checks whether the current object is inherited from the target object.
func (c *TCoolBands) Is() TIs {
	return TIs(c._instance())
}

// As
//
// 动态转换当前对象为目标对象。
//
// Dynamically convert the current object to the target object.
//func (c *TCoolBands) As() TAs {
//    return TAs(c._instance())
//}

// TCoolBandsClass
//
// 获取类信息指针。
//
// Get class information pointer.
func TCoolBandsClass() TClass {
	return CoolBands_StaticClassType()
}

func (c *TCoolBands) Add() *TCoolBand {
	return AsCoolBand(CoolBands_Add(c._instance()))
}

func (c *TCoolBands) FindBand(AControl IControl) *TCoolBand {
	return AsCoolBand(CoolBands_FindBand(c._instance(), CheckPtr(AControl)))
}

// Owner
//
// 组件所有者。
//
// component owner.
func (c *TCoolBands) Owner() *TObject {
	return AsObject(CoolBands_Owner(c._instance()))
}

// Assign
//
// 复制一个对象，如果对象实现了此方法的话。
//
// Copy an object, if the object implements this method.
func (c *TCoolBands) Assign(Source IObject) {
	CoolBands_Assign(c._instance(), CheckPtr(Source))
}

func (c *TCoolBands) BeginUpdate() {
	CoolBands_BeginUpdate(c._instance())
}

// Clear
//
// 清除。
func (c *TCoolBands) Clear() {
	CoolBands_Clear(c._instance())
}

func (c *TCoolBands) Delete(Index int32) {
	CoolBands_Delete(c._instance(), Index)
}

func (c *TCoolBands) EndUpdate() {
	CoolBands_EndUpdate(c._instance())
}

func (c *TCoolBands) FindItemID(ID int32) *TCollectionItem {
	return AsCollectionItem(CoolBands_FindItemID(c._instance(), ID))
}

// GetNamePath
//
// 获取类名路径。
//
// Get the class name path.
func (c *TCoolBands) GetNamePath() string {
	return CoolBands_GetNamePath(c._instance())
}

func (c *TCoolBands) Insert(Index int32) *TCollectionItem {
	return AsCollectionItem(CoolBands_Insert(c._instance(), Index))
}

// ClassType
//
// 获取类的类型信息。
//
// Get class type information.
func (c *TCoolBands) ClassType() TClass {
	return CoolBands_ClassType(c._instance())
}

// ClassName
//
// 获取当前对象类名称。
//
// Get the current object class name.
func (c *TCoolBands) ClassName() string {
	return CoolBands_ClassName(c._instance())
}

// InstanceSize
//
// 获取当前对象实例大小。
//
// Get the current object instance size.
func (c *TCoolBands) InstanceSize() int32 {
	return CoolBands_InstanceSize(c._instance())
}

// InheritsFrom
//
// 判断当前类是否继承自指定类。
//
// Determine whether the current class inherits from the specified class.
func (c *TCoolBands) InheritsFrom(AClass TClass) bool {
	return CoolBands_InheritsFrom(c._instance(), AClass)
}

// Equals
//
// 与一个对象进行比较。
//
// Compare with an object.
func (c *TCoolBands) Equals(Obj IObject) bool {
	return CoolBands_Equals(c._instance(), CheckPtr(Obj))
}

// GetHashCode
//
// 获取类的哈希值。
//
// Get the hash value of the class.
func (c *TCoolBands) GetHashCode() int32 {
	return CoolBands_GetHashCode(c._instance())
}

// ToString
//
// 文本类信息。
//
// Text information.
func (c *TCoolBands) ToString() string {
	return CoolBands_ToString(c._instance())
}

func (c *TCoolBands) Count() int32 {
	return CoolBands_GetCount(c._instance())
}

func (c *TCoolBands) Items(Index int32) *TCoolBand {
	return AsCoolBand(CoolBands_GetItems(c._instance(), Index))
}

func (c *TCoolBands) SetItems(Index int32, value *TCoolBand) {
	CoolBands_SetItems(c._instance(), Index, CheckPtr(value))
}
