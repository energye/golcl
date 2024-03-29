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
	"time"
	"unsafe"
)

type TRegistry struct {
	IObject
	instance unsafe.Pointer
}

// NewRegistry
//
// 创建一个新的对象。
//
// Create a new object.
func NewRegistry(aAccess uint32) *TRegistry {
	r := new(TRegistry)
	r.instance = unsafe.Pointer(Registry_Create(uintptr(aAccess)))
	setFinalizer(r, (*TRegistry).Free)
	return r
}

// AsRegistry
//
// 动态转换一个已存在的对象实例。
//
// Dynamically convert an existing object instance.
func AsRegistry(obj interface{}) *TRegistry {
	instance := getInstance(obj)
	if instance == nullptr {
		return nil
	}
	return &TRegistry{instance: instance}
}

// Free
//
// 释放对象。
//
// Free object.
func (r *TRegistry) Free() {
	if r.instance != nullptr {
		Registry_Free(r._instance())
		r.instance = nullptr
	}
}

func (r *TRegistry) _instance() uintptr {
	return uintptr(r.instance)
}

// Instance
//
// 返回对象实例指针。
//
// Return object instance pointer.
func (r *TRegistry) Instance() uintptr {
	return r._instance()
}

// UnsafeAddr
//
// 获取一个不安全的地址。
//
// Get an unsafe address.
func (r *TRegistry) UnsafeAddr() unsafe.Pointer {
	return r.instance
}

// IsValid
//
// 检测地址是否为空。
//
// Check if the address is empty.
func (r *TRegistry) IsValid() bool {
	return r.instance != nullptr
}

// Is
//
// 检测当前对象是否继承自目标对象。
//
// Checks whether the current object is inherited from the target object.
func (r *TRegistry) Is() TIs {
	return TIs(r._instance())
}

// As
//
// 动态转换当前对象为目标对象。
//
// Dynamically convert the current object to the target object.
//func (r *TRegistry) As() TAs {
//    return TAs(r._instance())
//}

// TRegistryClass
//
// 获取类信息指针。
//
// Get class information pointer.
func TRegistryClass() TClass {
	return Registry_StaticClassType()
}

func (r *TRegistry) CloseKey() {
	Registry_CloseKey(r._instance())
}

func (r *TRegistry) CreateKey(Key string) bool {
	return Registry_CreateKey(r._instance(), Key)
}

func (r *TRegistry) DeleteKey(Key string) bool {
	return Registry_DeleteKey(r._instance(), Key)
}

func (r *TRegistry) DeleteValue(Name string) bool {
	return Registry_DeleteValue(r._instance(), Name)
}

func (r *TRegistry) HasSubKeys() bool {
	return Registry_HasSubKeys(r._instance())
}

func (r *TRegistry) KeyExists(Key string) bool {
	return Registry_KeyExists(r._instance(), Key)
}

func (r *TRegistry) LoadKey(Key string, FileName string) bool {
	return Registry_LoadKey(r._instance(), Key, FileName)
}

func (r *TRegistry) MoveKey(OldName string, NewName string, Delete bool) {
	Registry_MoveKey(r._instance(), OldName, NewName, Delete)
}

func (r *TRegistry) OpenKey(Key string, CanCreate bool) bool {
	return Registry_OpenKey(r._instance(), Key, CanCreate)
}

func (r *TRegistry) OpenKeyReadOnly(Key string) bool {
	return Registry_OpenKeyReadOnly(r._instance(), Key)
}

func (r *TRegistry) ReadBool(Name string) bool {
	return Registry_ReadBool(r._instance(), Name)
}

func (r *TRegistry) ReadDate(Name string) time.Time {
	return Registry_ReadDate(r._instance(), Name)
}

func (r *TRegistry) ReadDateTime(Name string) time.Time {
	return Registry_ReadDateTime(r._instance(), Name)
}

func (r *TRegistry) ReadFloat(Name string) float64 {
	return Registry_ReadFloat(r._instance(), Name)
}

func (r *TRegistry) ReadInteger(Name string) int32 {
	return Registry_ReadInteger(r._instance(), Name)
}

func (r *TRegistry) ReadString(Name string) string {
	return Registry_ReadString(r._instance(), Name)
}

func (r *TRegistry) ReadTime(Name string) time.Time {
	return Registry_ReadTime(r._instance(), Name)
}

func (r *TRegistry) RegistryConnect(UNCName string) bool {
	return Registry_RegistryConnect(r._instance(), UNCName)
}

func (r *TRegistry) RenameValue(OldName string, NewName string) {
	Registry_RenameValue(r._instance(), OldName, NewName)
}

func (r *TRegistry) ReplaceKey(Key string, FileName string, BackUpFileName string) bool {
	return Registry_ReplaceKey(r._instance(), Key, FileName, BackUpFileName)
}

func (r *TRegistry) RestoreKey(Key string, FileName string) bool {
	return Registry_RestoreKey(r._instance(), Key, FileName)
}

func (r *TRegistry) SaveKey(Key string, FileName string) bool {
	return Registry_SaveKey(r._instance(), Key, FileName)
}

func (r *TRegistry) UnLoadKey(Key string) bool {
	return Registry_UnLoadKey(r._instance(), Key)
}

func (r *TRegistry) ValueExists(Name string) bool {
	return Registry_ValueExists(r._instance(), Name)
}

func (r *TRegistry) WriteBool(Name string, Value bool) {
	Registry_WriteBool(r._instance(), Name, Value)
}

func (r *TRegistry) WriteDate(Name string, Value time.Time) {
	Registry_WriteDate(r._instance(), Name, Value)
}

func (r *TRegistry) WriteDateTime(Name string, Value time.Time) {
	Registry_WriteDateTime(r._instance(), Name, Value)
}

func (r *TRegistry) WriteFloat(Name string, Value float64) {
	Registry_WriteFloat(r._instance(), Name, Value)
}

func (r *TRegistry) WriteInteger(Name string, Value int32) {
	Registry_WriteInteger(r._instance(), Name, Value)
}

func (r *TRegistry) WriteString(Name string, Value string) {
	Registry_WriteString(r._instance(), Name, Value)
}

func (r *TRegistry) WriteExpandString(Name string, Value string) {
	Registry_WriteExpandString(r._instance(), Name, Value)
}

func (r *TRegistry) WriteTime(Name string, Value time.Time) {
	Registry_WriteTime(r._instance(), Name, Value)
}

// ClassType
//
// 获取类的类型信息。
//
// Get class type information.
func (r *TRegistry) ClassType() TClass {
	return Registry_ClassType(r._instance())
}

// ClassName
//
// 获取当前对象类名称。
//
// Get the current object class name.
func (r *TRegistry) ClassName() string {
	return Registry_ClassName(r._instance())
}

// InstanceSize
//
// 获取当前对象实例大小。
//
// Get the current object instance size.
func (r *TRegistry) InstanceSize() int32 {
	return Registry_InstanceSize(r._instance())
}

// InheritsFrom
//
// 判断当前类是否继承自指定类。
//
// Determine whether the current class inherits from the specified class.
func (r *TRegistry) InheritsFrom(AClass TClass) bool {
	return Registry_InheritsFrom(r._instance(), AClass)
}

// Equals
//
// 与一个对象进行比较。
//
// Compare with an object.
func (r *TRegistry) Equals(Obj IObject) bool {
	return Registry_Equals(r._instance(), CheckPtr(Obj))
}

// GetHashCode
//
// 获取类的哈希值。
//
// Get the hash value of the class.
func (r *TRegistry) GetHashCode() int32 {
	return Registry_GetHashCode(r._instance())
}

// ToString
//
// 文本类信息。
//
// Text information.
func (r *TRegistry) ToString() string {
	return Registry_ToString(r._instance())
}

func (r *TRegistry) CurrentKey() HKEY {
	return Registry_GetCurrentKey(r._instance())
}

func (r *TRegistry) CurrentPath() string {
	return Registry_GetCurrentPath(r._instance())
}

func (r *TRegistry) LazyWrite() bool {
	return Registry_GetLazyWrite(r._instance())
}

func (r *TRegistry) SetLazyWrite(value bool) {
	Registry_SetLazyWrite(r._instance(), value)
}

func (r *TRegistry) LastError() int32 {
	return Registry_GetLastError(r._instance())
}

func (r *TRegistry) LastErrorMsg() string {
	return Registry_GetLastErrorMsg(r._instance())
}

func (r *TRegistry) RootKey() HKEY {
	return Registry_GetRootKey(r._instance())
}

func (r *TRegistry) SetRootKey(value HKEY) {
	Registry_SetRootKey(r._instance(), value)
}

func (r *TRegistry) Access() uint32 {
	return Registry_GetAccess(r._instance())
}

func (r *TRegistry) SetAccess(value uint32) {
	Registry_SetAccess(r._instance(), value)
}
