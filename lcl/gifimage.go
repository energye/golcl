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

type TGIFImage struct {
	IBitmap
	instance unsafe.Pointer
}

// NewGIFImage
//
// 创建一个新的对象。
//
// Create a new object.
func NewGIFImage() *TGIFImage {
	g := new(TGIFImage)
	g.instance = unsafe.Pointer(GIFImage_Create())
	setFinalizer(g, (*TGIFImage).Free)
	return g
}

// AsGIFImage
//
// 动态转换一个已存在的对象实例。
//
// Dynamically convert an existing object instance.
func AsGIFImage(obj interface{}) *TGIFImage {
	instance := getInstance(obj)
	if instance == nullptr {
		return nil
	}
	return &TGIFImage{instance: instance}
}

// Free
//
// 释放对象。
//
// Free object.
func (g *TGIFImage) Free() {
	if g.instance != nullptr {
		GIFImage_Free(g._instance())
		g.instance = nullptr
	}
}

func (g *TGIFImage) _instance() uintptr {
	return uintptr(g.instance)
}

// Instance
//
// 返回对象实例指针。
//
// Return object instance pointer.
func (g *TGIFImage) Instance() uintptr {
	return g._instance()
}

// UnsafeAddr
//
// 获取一个不安全的地址。
//
// Get an unsafe address.
func (g *TGIFImage) UnsafeAddr() unsafe.Pointer {
	return g.instance
}

// IsValid
//
// 检测地址是否为空。
//
// Check if the address is empty.
func (g *TGIFImage) IsValid() bool {
	return g.instance != nullptr
}

// Is
//
// 检测当前对象是否继承自目标对象。
//
// Checks whether the current object is inherited from the target object.
func (g *TGIFImage) Is() TIs {
	return TIs(g._instance())
}

// As
//
// 动态转换当前对象为目标对象。
//
// Dynamically convert the current object to the target object.
//func (g *TGIFImage) As() TAs {
//    return TAs(g._instance())
//}

// TGIFImageClass
//
// 获取类信息指针。
//
// Get class information pointer.
func TGIFImageClass() TClass {
	return GIFImage_StaticClassType()
}

// SaveToStream
//
// 保存至流。
func (g *TGIFImage) SaveToStream(Stream IStream) {
	GIFImage_SaveToStream(g._instance(), CheckPtr(Stream))
}

// LoadFromStream
//
// 文件流加载。
func (g *TGIFImage) LoadFromStream(Stream IStream) {
	GIFImage_LoadFromStream(g._instance(), CheckPtr(Stream))
}

// Clear
//
// 清除。
func (g *TGIFImage) Clear() {
	GIFImage_Clear(g._instance())
}

// Assign
//
// 复制一个对象，如果对象实现了此方法的话。
//
// Copy an object, if the object implements this method.
func (g *TGIFImage) Assign(Source IObject) {
	GIFImage_Assign(g._instance(), CheckPtr(Source))
}

// Equals
//
// 与一个对象进行比较。
//
// Compare with an object.
func (g *TGIFImage) Equals(Obj IObject) bool {
	return GIFImage_Equals(g._instance(), CheckPtr(Obj))
}

// LoadFromFile
//
// 从文件加载。
func (g *TGIFImage) LoadFromFile(Filename string) {
	GIFImage_LoadFromFile(g._instance(), Filename)
}

// SaveToFile
//
// 保存至文件。
func (g *TGIFImage) SaveToFile(Filename string) {
	GIFImage_SaveToFile(g._instance(), Filename)
}

func (g *TGIFImage) SetSize(AWidth int32, AHeight int32) {
	GIFImage_SetSize(g._instance(), AWidth, AHeight)
}

// GetNamePath
//
// 获取类名路径。
//
// Get the class name path.
func (g *TGIFImage) GetNamePath() string {
	return GIFImage_GetNamePath(g._instance())
}

// ClassType
//
// 获取类的类型信息。
//
// Get class type information.
func (g *TGIFImage) ClassType() TClass {
	return GIFImage_ClassType(g._instance())
}

// ClassName
//
// 获取当前对象类名称。
//
// Get the current object class name.
func (g *TGIFImage) ClassName() string {
	return GIFImage_ClassName(g._instance())
}

// InstanceSize
//
// 获取当前对象实例大小。
//
// Get the current object instance size.
func (g *TGIFImage) InstanceSize() int32 {
	return GIFImage_InstanceSize(g._instance())
}

// InheritsFrom
//
// 判断当前类是否继承自指定类。
//
// Determine whether the current class inherits from the specified class.
func (g *TGIFImage) InheritsFrom(AClass TClass) bool {
	return GIFImage_InheritsFrom(g._instance(), AClass)
}

// GetHashCode
//
// 获取类的哈希值。
//
// Get the hash value of the class.
func (g *TGIFImage) GetHashCode() int32 {
	return GIFImage_GetHashCode(g._instance())
}

// ToString
//
// 文本类信息。
//
// Text information.
func (g *TGIFImage) ToString() string {
	return GIFImage_ToString(g._instance())
}

func (g *TGIFImage) BitsPerPixel() int32 {
	return GIFImage_GetBitsPerPixel(g._instance())
}

func (g *TGIFImage) Empty() bool {
	return GIFImage_GetEmpty(g._instance())
}

// Height
//
// 获取高度。
//
// Get height.
func (g *TGIFImage) Height() int32 {
	return GIFImage_GetHeight(g._instance())
}

// SetHeight
//
// 设置高度。
//
// Set height.
func (g *TGIFImage) SetHeight(value int32) {
	GIFImage_SetHeight(g._instance(), value)
}

// Modified
//
// 获取修改。
//
// Get modified.
func (g *TGIFImage) Modified() bool {
	return GIFImage_GetModified(g._instance())
}

// SetModified
//
// 设置修改。
//
// Set modified.
func (g *TGIFImage) SetModified(value bool) {
	GIFImage_SetModified(g._instance(), value)
}

func (g *TGIFImage) Palette() HPALETTE {
	return GIFImage_GetPalette(g._instance())
}

func (g *TGIFImage) SetPalette(value HPALETTE) {
	GIFImage_SetPalette(g._instance(), value)
}

func (g *TGIFImage) PaletteModified() bool {
	return GIFImage_GetPaletteModified(g._instance())
}

func (g *TGIFImage) SetPaletteModified(value bool) {
	GIFImage_SetPaletteModified(g._instance(), value)
}

// Transparent
//
// 获取透明。
//
// Get transparent.
func (g *TGIFImage) Transparent() bool {
	return GIFImage_GetTransparent(g._instance())
}

// Width
//
// 获取宽度。
//
// Get width.
func (g *TGIFImage) Width() int32 {
	return GIFImage_GetWidth(g._instance())
}

// SetWidth
//
// 设置宽度。
//
// Set width.
func (g *TGIFImage) SetWidth(value int32) {
	GIFImage_SetWidth(g._instance(), value)
}

// SetOnChange
//
// 设置改变事件。
//
// Set changed event.
func (g *TGIFImage) SetOnChange(fn TNotifyEvent) {
	GIFImage_SetOnChange(g._instance(), fn)
}
