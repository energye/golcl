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

type TCoolBands struct {
	IObject
	instance uintptr
	// 特殊情况下使用，主要应对Go的GC问题，与LCL没有太多关系。
	ptr unsafe.Pointer
}

// 创建一个新的对象。
//
// Create a new object.
func NewCoolBands(AOwner *TCoolBar) *TCoolBands {
	c := new(TCoolBands)
	c.instance = CoolBands_Create(CheckPtr(AOwner))
	c.ptr = unsafe.Pointer(c.instance)
	setFinalizer(c, (*TCoolBands).Free)
	return c
}

// 动态转换一个已存在的对象实例。
//
// Dynamically convert an existing object instance.
func AsCoolBands(obj interface{}) *TCoolBands {
	instance, ptr := getInstance(obj)
	if instance == 0 {
		return nil
	}
	return &TCoolBands{instance: instance, ptr: ptr}
}

// -------------------------- Deprecated begin --------------------------
// 新建一个对象来自已经存在的对象实例指针。
//
// Create a new object from an existing object instance pointer.
// Deprecated: use AsCoolBands.
func CoolBandsFromInst(inst uintptr) *TCoolBands {
	return AsCoolBands(inst)
}

// 新建一个对象来自已经存在的对象实例。
//
// Create a new object from an existing object instance.
// Deprecated: use AsCoolBands.
func CoolBandsFromObj(obj IObject) *TCoolBands {
	return AsCoolBands(obj)
}

// 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
//
// Create a new object from an unsecured address. Note: Using this function may cause some unclear situations and be used with caution..
// Deprecated: use AsCoolBands.
func CoolBandsFromUnsafePointer(ptr unsafe.Pointer) *TCoolBands {
	return AsCoolBands(ptr)
}

// -------------------------- Deprecated end --------------------------
// 释放对象。
//
// Free object.
func (c *TCoolBands) Free() {
	if c.instance != 0 {
		CoolBands_Free(c.instance)
		c.instance, c.ptr = 0, nullptr
	}
}

// 返回对象实例指针。
//
// Return object instance pointer.
func (c *TCoolBands) Instance() uintptr {
	return c.instance
}

// 获取一个不安全的地址。
//
// Get an unsafe address.
func (c *TCoolBands) UnsafeAddr() unsafe.Pointer {
	return c.ptr
}

// 检测地址是否为空。
//
// Check if the address is empty.
func (c *TCoolBands) IsValid() bool {
	return c.instance != 0
}

// 检测当前对象是否继承自目标对象。
//
// Checks whether the current object is inherited from the target object.
func (c *TCoolBands) Is() TIs {
	return TIs(c.instance)
}

// 动态转换当前对象为目标对象。
//
// Dynamically convert the current object to the target object.
//func (c *TCoolBands) As() TAs {
//    return TAs(c.instance)
//}

// 获取类信息指针。
//
// Get class information pointer.
func TCoolBandsClass() TClass {
	return CoolBands_StaticClassType()
}

func (c *TCoolBands) Add() *TCoolBand {
	return AsCoolBand(CoolBands_Add(c.instance))
}

func (c *TCoolBands) FindBand(AControl IControl) *TCoolBand {
	return AsCoolBand(CoolBands_FindBand(c.instance, CheckPtr(AControl)))
}

// 组件所有者。
//
// component owner.
func (c *TCoolBands) Owner() *TObject {
	return AsObject(CoolBands_Owner(c.instance))
}

// 复制一个对象，如果对象实现了此方法的话。
//
// Copy an object, if the object implements this method.
func (c *TCoolBands) Assign(Source IObject) {
	CoolBands_Assign(c.instance, CheckPtr(Source))
}

func (c *TCoolBands) BeginUpdate() {
	CoolBands_BeginUpdate(c.instance)
}

// 清除。
func (c *TCoolBands) Clear() {
	CoolBands_Clear(c.instance)
}

func (c *TCoolBands) Delete(Index int32) {
	CoolBands_Delete(c.instance, Index)
}

func (c *TCoolBands) EndUpdate() {
	CoolBands_EndUpdate(c.instance)
}

func (c *TCoolBands) FindItemID(ID int32) *TCollectionItem {
	return AsCollectionItem(CoolBands_FindItemID(c.instance, ID))
}

// 获取类名路径。
//
// Get the class name path.
func (c *TCoolBands) GetNamePath() string {
	return CoolBands_GetNamePath(c.instance)
}

func (c *TCoolBands) Insert(Index int32) *TCollectionItem {
	return AsCollectionItem(CoolBands_Insert(c.instance, Index))
}

// 获取类的类型信息。
//
// Get class type information.
func (c *TCoolBands) ClassType() TClass {
	return CoolBands_ClassType(c.instance)
}

// 获取当前对象类名称。
//
// Get the current object class name.
func (c *TCoolBands) ClassName() string {
	return CoolBands_ClassName(c.instance)
}

// 获取当前对象实例大小。
//
// Get the current object instance size.
func (c *TCoolBands) InstanceSize() int32 {
	return CoolBands_InstanceSize(c.instance)
}

// 判断当前类是否继承自指定类。
//
// Determine whether the current class inherits from the specified class.
func (c *TCoolBands) InheritsFrom(AClass TClass) bool {
	return CoolBands_InheritsFrom(c.instance, AClass)
}

// 与一个对象进行比较。
//
// Compare with an object.
func (c *TCoolBands) Equals(Obj IObject) bool {
	return CoolBands_Equals(c.instance, CheckPtr(Obj))
}

// 获取类的哈希值。
//
// Get the hash value of the class.
func (c *TCoolBands) GetHashCode() int32 {
	return CoolBands_GetHashCode(c.instance)
}

// 文本类信息。
//
// Text information.
func (c *TCoolBands) ToString() string {
	return CoolBands_ToString(c.instance)
}

func (c *TCoolBands) Count() int32 {
	return CoolBands_GetCount(c.instance)
}

func (c *TCoolBands) Items(Index int32) *TCoolBand {
	return AsCoolBand(CoolBands_GetItems(c.instance, Index))
}

func (c *TCoolBands) SetItems(Index int32, value *TCoolBand) {
	CoolBands_SetItems(c.instance, Index, CheckPtr(value))
}
