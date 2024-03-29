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

type TStrings struct {
	IStrings
	instance unsafe.Pointer
}

// NewStrings
//
// 创建一个新的对象。
//
// Create a new object.
func NewStrings() *TStrings {
	s := new(TStrings)
	s.instance = unsafe.Pointer(Strings_Create())
	setFinalizer(s, (*TStrings).Free)
	return s
}

// AsStrings
//
// 动态转换一个已存在的对象实例。
//
// Dynamically convert an existing object instance.
func AsStrings(obj interface{}) *TStrings {
	instance := getInstance(obj)
	if instance == nullptr {
		return nil
	}
	return &TStrings{instance: instance}
}

// Free
//
// 释放对象。
//
// Free object.
func (s *TStrings) Free() {
	if s.instance != nullptr {
		Strings_Free(s._instance())
		s.instance = nullptr
	}
}

func (s *TStrings) _instance() uintptr {
	return uintptr(s.instance)
}

// Instance
//
// 返回对象实例指针。
//
// Return object instance pointer.
func (s *TStrings) Instance() uintptr {
	return s._instance()
}

// UnsafeAddr
//
// 获取一个不安全的地址。
//
// Get an unsafe address.
func (s *TStrings) UnsafeAddr() unsafe.Pointer {
	return s.instance
}

// IsValid
//
// 检测地址是否为空。
//
// Check if the address is empty.
func (s *TStrings) IsValid() bool {
	return s.instance != nullptr
}

// Is
//
// 检测当前对象是否继承自目标对象。
//
// Checks whether the current object is inherited from the target object.
func (s *TStrings) Is() TIs {
	return TIs(s._instance())
}

// As
//
// 动态转换当前对象为目标对象。
//
// Dynamically convert the current object to the target object.
//func (s *TStrings) As() TAs {
//    return TAs(s._instance())
//}

// TStringsClass
//
// 获取类信息指针。
//
// Get class information pointer.
func TStringsClass() TClass {
	return Strings_StaticClassType()
}

// S
//
// Strings()的别名。
//
// Alias of Strings().
func (s *TStrings) S(Index int32) string {
	return s.Strings(Index)
}

// SetS
//
// SetStrings()的别名。
//
// Alias of SetStrings().
func (s *TStrings) SetS(Index int32, value string) {
	s.SetStrings(Index, value)
}

func (s *TStrings) Add(S string) int32 {
	return Strings_Add(s._instance(), S)
}

func (s *TStrings) AddObject(S string, AObject IObject) int32 {
	return Strings_AddObject(s._instance(), S, CheckPtr(AObject))
}

func (s *TStrings) Append(S string) {
	Strings_Append(s._instance(), S)
}

// Assign
//
// 复制一个对象，如果对象实现了此方法的话。
//
// Copy an object, if the object implements this method.
func (s *TStrings) Assign(Source IObject) {
	Strings_Assign(s._instance(), CheckPtr(Source))
}

func (s *TStrings) BeginUpdate() {
	Strings_BeginUpdate(s._instance())
}

// Clear
//
// 清除。
func (s *TStrings) Clear() {
	Strings_Clear(s._instance())
}

func (s *TStrings) Delete(Index int32) {
	Strings_Delete(s._instance(), Index)
}

func (s *TStrings) EndUpdate() {
	Strings_EndUpdate(s._instance())
}

// Equals
//
// 与一个对象进行比较。
//
// Compare with an object.
func (s *TStrings) Equals(Strings IObject) bool {
	return Strings_Equals(s._instance(), CheckPtr(Strings))
}

func (s *TStrings) Exchange(Index1 int32, Index2 int32) {
	Strings_Exchange(s._instance(), Index1, Index2)
}

func (s *TStrings) IndexOf(S string) int32 {
	return Strings_IndexOf(s._instance(), S)
}

func (s *TStrings) IndexOfName(Name string) int32 {
	return Strings_IndexOfName(s._instance(), Name)
}

func (s *TStrings) IndexOfObject(AObject IObject) int32 {
	return Strings_IndexOfObject(s._instance(), CheckPtr(AObject))
}

func (s *TStrings) Insert(Index int32, S string) {
	Strings_Insert(s._instance(), Index, S)
}

func (s *TStrings) InsertObject(Index int32, S string, AObject IObject) {
	Strings_InsertObject(s._instance(), Index, S, CheckPtr(AObject))
}

// LoadFromFile
//
// 从文件加载。
func (s *TStrings) LoadFromFile(FileName string) {
	Strings_LoadFromFile(s._instance(), FileName)
}

// LoadFromStream
//
// 文件流加载。
func (s *TStrings) LoadFromStream(Stream IStream) {
	Strings_LoadFromStream(s._instance(), CheckPtr(Stream))
}

func (s *TStrings) Move(CurIndex int32, NewIndex int32) {
	Strings_Move(s._instance(), CurIndex, NewIndex)
}

// SaveToFile
//
// 保存至文件。
func (s *TStrings) SaveToFile(FileName string) {
	Strings_SaveToFile(s._instance(), FileName)
}

// SaveToStream
//
// 保存至流。
func (s *TStrings) SaveToStream(Stream IStream) {
	Strings_SaveToStream(s._instance(), CheckPtr(Stream))
}

// GetNamePath
//
// 获取类名路径。
//
// Get the class name path.
func (s *TStrings) GetNamePath() string {
	return Strings_GetNamePath(s._instance())
}

// ClassType
//
// 获取类的类型信息。
//
// Get class type information.
func (s *TStrings) ClassType() TClass {
	return Strings_ClassType(s._instance())
}

// ClassName
//
// 获取当前对象类名称。
//
// Get the current object class name.
func (s *TStrings) ClassName() string {
	return Strings_ClassName(s._instance())
}

// InstanceSize
//
// 获取当前对象实例大小。
//
// Get the current object instance size.
func (s *TStrings) InstanceSize() int32 {
	return Strings_InstanceSize(s._instance())
}

// InheritsFrom
//
// 判断当前类是否继承自指定类。
//
// Determine whether the current class inherits from the specified class.
func (s *TStrings) InheritsFrom(AClass TClass) bool {
	return Strings_InheritsFrom(s._instance(), AClass)
}

// GetHashCode
//
// 获取类的哈希值。
//
// Get the hash value of the class.
func (s *TStrings) GetHashCode() int32 {
	return Strings_GetHashCode(s._instance())
}

// ToString
//
// 文本类信息。
//
// Text information.
func (s *TStrings) ToString() string {
	return Strings_ToString(s._instance())
}

func (s *TStrings) Capacity() int32 {
	return Strings_GetCapacity(s._instance())
}

func (s *TStrings) SetCapacity(value int32) {
	Strings_SetCapacity(s._instance(), value)
}

func (s *TStrings) CommaText() string {
	return Strings_GetCommaText(s._instance())
}

func (s *TStrings) SetCommaText(value string) {
	Strings_SetCommaText(s._instance(), value)
}

func (s *TStrings) Count() int32 {
	return Strings_GetCount(s._instance())
}

func (s *TStrings) Delimiter() uint16 {
	return Strings_GetDelimiter(s._instance())
}

func (s *TStrings) SetDelimiter(value uint16) {
	Strings_SetDelimiter(s._instance(), value)
}

func (s *TStrings) NameValueSeparator() uint16 {
	return Strings_GetNameValueSeparator(s._instance())
}

func (s *TStrings) SetNameValueSeparator(value uint16) {
	Strings_SetNameValueSeparator(s._instance(), value)
}

// Text
//
// 获取文本。
func (s *TStrings) Text() string {
	return Strings_GetText(s._instance())
}

// SetText
//
// 设置文本。
func (s *TStrings) SetText(value string) {
	Strings_SetText(s._instance(), value)
}

func (s *TStrings) Objects(Index int32) *TObject {
	return AsObject(Strings_GetObjects(s._instance(), Index))
}

func (s *TStrings) SetObjects(Index int32, value IObject) {
	Strings_SetObjects(s._instance(), Index, CheckPtr(value))
}

func (s *TStrings) Values(Name string) string {
	return Strings_GetValues(s._instance(), Name)
}

func (s *TStrings) SetValues(Name string, value string) {
	Strings_SetValues(s._instance(), Name, value)
}

func (s *TStrings) ValueFromIndex(Index int32) string {
	return Strings_GetValueFromIndex(s._instance(), Index)
}

func (s *TStrings) SetValueFromIndex(Index int32, value string) {
	Strings_SetValueFromIndex(s._instance(), Index, value)
}

func (s *TStrings) Strings(Index int32) string {
	return Strings_GetStrings(s._instance(), Index)
}

func (s *TStrings) SetStrings(Index int32, value string) {
	Strings_SetStrings(s._instance(), Index, value)
}
