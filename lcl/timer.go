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

type TTimer struct {
	IComponent
	instance uintptr
	// 特殊情况下使用，主要应对Go的GC问题，与LCL没有太多关系。
	ptr unsafe.Pointer
}

// 创建一个新的对象。
//
// Create a new object.
func NewTimer(owner IComponent) *TTimer {
	t := new(TTimer)
	t.instance = Timer_Create(CheckPtr(owner))
	t.ptr = unsafe.Pointer(t.instance)
	return t
}

// 动态转换一个已存在的对象实例。
//
// Dynamically convert an existing object instance.
func AsTimer(obj interface{}) *TTimer {
	instance, ptr := getInstance(obj)
	if instance == 0 {
		return nil
	}
	return &TTimer{instance: instance, ptr: ptr}
}

// -------------------------- Deprecated begin --------------------------
// 新建一个对象来自已经存在的对象实例指针。
//
// Create a new object from an existing object instance pointer.
// Deprecated: use AsTimer.
func TimerFromInst(inst uintptr) *TTimer {
	return AsTimer(inst)
}

// 新建一个对象来自已经存在的对象实例。
//
// Create a new object from an existing object instance.
// Deprecated: use AsTimer.
func TimerFromObj(obj IObject) *TTimer {
	return AsTimer(obj)
}

// 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
//
// Create a new object from an unsecured address. Note: Using this function may cause some unclear situations and be used with caution..
// Deprecated: use AsTimer.
func TimerFromUnsafePointer(ptr unsafe.Pointer) *TTimer {
	return AsTimer(ptr)
}

// -------------------------- Deprecated end --------------------------
// 释放对象。
//
// Free object.
func (t *TTimer) Free() {
	if t.instance != 0 {
		Timer_Free(t.instance)
		t.instance, t.ptr = 0, nullptr
	}
}

// 返回对象实例指针。
//
// Return object instance pointer.
func (t *TTimer) Instance() uintptr {
	return t.instance
}

// 获取一个不安全的地址。
//
// Get an unsafe address.
func (t *TTimer) UnsafeAddr() unsafe.Pointer {
	return t.ptr
}

// 检测地址是否为空。
//
// Check if the address is empty.
func (t *TTimer) IsValid() bool {
	return t.instance != 0
}

// 检测当前对象是否继承自目标对象。
//
// Checks whether the current object is inherited from the target object.
func (t *TTimer) Is() TIs {
	return TIs(t.instance)
}

// 动态转换当前对象为目标对象。
//
// Dynamically convert the current object to the target object.
//func (t *TTimer) As() TAs {
//    return TAs(t.instance)
//}

// 获取类信息指针。
//
// Get class information pointer.
func TTimerClass() TClass {
	return Timer_StaticClassType()
}

// 查找指定名称的组件。
//
// Find the component with the specified name.
func (t *TTimer) FindComponent(AName string) *TComponent {
	return AsComponent(Timer_FindComponent(t.instance, AName))
}

// 获取类名路径。
//
// Get the class name path.
func (t *TTimer) GetNamePath() string {
	return Timer_GetNamePath(t.instance)
}

// 是否有父容器。
//
// Is there a parent container.
func (t *TTimer) HasParent() bool {
	return Timer_HasParent(t.instance)
}

// 复制一个对象，如果对象实现了此方法的话。
//
// Copy an object, if the object implements this method.
func (t *TTimer) Assign(Source IObject) {
	Timer_Assign(t.instance, CheckPtr(Source))
}

// 获取类的类型信息。
//
// Get class type information.
func (t *TTimer) ClassType() TClass {
	return Timer_ClassType(t.instance)
}

// 获取当前对象类名称。
//
// Get the current object class name.
func (t *TTimer) ClassName() string {
	return Timer_ClassName(t.instance)
}

// 获取当前对象实例大小。
//
// Get the current object instance size.
func (t *TTimer) InstanceSize() int32 {
	return Timer_InstanceSize(t.instance)
}

// 判断当前类是否继承自指定类。
//
// Determine whether the current class inherits from the specified class.
func (t *TTimer) InheritsFrom(AClass TClass) bool {
	return Timer_InheritsFrom(t.instance, AClass)
}

// 与一个对象进行比较。
//
// Compare with an object.
func (t *TTimer) Equals(Obj IObject) bool {
	return Timer_Equals(t.instance, CheckPtr(Obj))
}

// 获取类的哈希值。
//
// Get the hash value of the class.
func (t *TTimer) GetHashCode() int32 {
	return Timer_GetHashCode(t.instance)
}

// 文本类信息。
//
// Text information.
func (t *TTimer) ToString() string {
	return Timer_ToString(t.instance)
}

// 获取控件启用。
//
// Get the control enabled.
func (t *TTimer) Enabled() bool {
	return Timer_GetEnabled(t.instance)
}

// 设置控件启用。
//
// Set the control enabled.
func (t *TTimer) SetEnabled(value bool) {
	Timer_SetEnabled(t.instance, value)
}

// 获取时钟每次跳动间隔时间，ms。
func (t *TTimer) Interval() uint32 {
	return Timer_GetInterval(t.instance)
}

// 设置时钟每次跳动间隔时间，ms。
func (t *TTimer) SetInterval(value uint32) {
	Timer_SetInterval(t.instance, value)
}

// 设置时钟事件。
func (t *TTimer) SetOnTimer(fn TNotifyEvent) {
	Timer_SetOnTimer(t.instance, fn)
}

// 获取组件总数。
//
// Get the total number of components.
func (t *TTimer) ComponentCount() int32 {
	return Timer_GetComponentCount(t.instance)
}

// 获取组件索引。
//
// Get component index.
func (t *TTimer) ComponentIndex() int32 {
	return Timer_GetComponentIndex(t.instance)
}

// 设置组件索引。
//
// Set component index.
func (t *TTimer) SetComponentIndex(value int32) {
	Timer_SetComponentIndex(t.instance, value)
}

// 获取组件所有者。
//
// Get component owner.
func (t *TTimer) Owner() *TComponent {
	return AsComponent(Timer_GetOwner(t.instance))
}

// 获取组件名称。
//
// Get the component name.
func (t *TTimer) Name() string {
	return Timer_GetName(t.instance)
}

// 设置组件名称。
//
// Set the component name.
func (t *TTimer) SetName(value string) {
	Timer_SetName(t.instance, value)
}

// 获取对象标记。
//
// Get the control tag.
func (t *TTimer) Tag() int {
	return Timer_GetTag(t.instance)
}

// 设置对象标记。
//
// Set the control tag.
func (t *TTimer) SetTag(value int) {
	Timer_SetTag(t.instance, value)
}

// 获取指定索引组件。
//
// Get the specified index component.
func (t *TTimer) Components(AIndex int32) *TComponent {
	return AsComponent(Timer_GetComponents(t.instance, AIndex))
}
