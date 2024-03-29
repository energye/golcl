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

type TPrinter struct {
	IObject
	instance unsafe.Pointer
}

// NewPrinter
//
// 创建一个新的对象。
//
// Create a new object.
func NewPrinter() *TPrinter {
	p := new(TPrinter)
	p.instance = unsafe.Pointer(Printer_Create())
	setFinalizer(p, (*TPrinter).Free)
	return p
}

// AsPrinter
//
// 动态转换一个已存在的对象实例。
//
// Dynamically convert an existing object instance.
func AsPrinter(obj interface{}) *TPrinter {
	instance := getInstance(obj)
	if instance == nullptr {
		return nil
	}
	return &TPrinter{instance: instance}
}

// Free
//
// 释放对象。
//
// Free object.
func (p *TPrinter) Free() {
	if p.instance != nullptr {
		Printer_Free(p._instance())
		p.instance = nullptr
	}
}

func (p *TPrinter) _instance() uintptr {
	return uintptr(p.instance)
}

// Instance
//
// 返回对象实例指针。
//
// Return object instance pointer.
func (p *TPrinter) Instance() uintptr {
	return p._instance()
}

// UnsafeAddr
//
// 获取一个不安全的地址。
//
// Get an unsafe address.
func (p *TPrinter) UnsafeAddr() unsafe.Pointer {
	return p.instance
}

// IsValid
//
// 检测地址是否为空。
//
// Check if the address is empty.
func (p *TPrinter) IsValid() bool {
	return p.instance != nullptr
}

// Is
//
// 检测当前对象是否继承自目标对象。
//
// Checks whether the current object is inherited from the target object.
func (p *TPrinter) Is() TIs {
	return TIs(p._instance())
}

// As
//
// 动态转换当前对象为目标对象。
//
// Dynamically convert the current object to the target object.
//func (p *TPrinter) As() TAs {
//    return TAs(p._instance())
//}

// TPrinterClass
//
// 获取类信息指针。
//
// Get class information pointer.
func TPrinterClass() TClass {
	return Printer_StaticClassType()
}

func (p *TPrinter) Abort() {
	Printer_Abort(p._instance())
}

func (p *TPrinter) BeginDoc() {
	Printer_BeginDoc(p._instance())
}

func (p *TPrinter) EndDoc() {
	Printer_EndDoc(p._instance())
}

func (p *TPrinter) NewPage() {
	Printer_NewPage(p._instance())
}

// Refresh
//
// 刷新控件。
//
// Refresh control.
func (p *TPrinter) Refresh() {
	Printer_Refresh(p._instance())
}

// ClassType
//
// 获取类的类型信息。
//
// Get class type information.
func (p *TPrinter) ClassType() TClass {
	return Printer_ClassType(p._instance())
}

// ClassName
//
// 获取当前对象类名称。
//
// Get the current object class name.
func (p *TPrinter) ClassName() string {
	return Printer_ClassName(p._instance())
}

// InstanceSize
//
// 获取当前对象实例大小。
//
// Get the current object instance size.
func (p *TPrinter) InstanceSize() int32 {
	return Printer_InstanceSize(p._instance())
}

// InheritsFrom
//
// 判断当前类是否继承自指定类。
//
// Determine whether the current class inherits from the specified class.
func (p *TPrinter) InheritsFrom(AClass TClass) bool {
	return Printer_InheritsFrom(p._instance(), AClass)
}

// Equals
//
// 与一个对象进行比较。
//
// Compare with an object.
func (p *TPrinter) Equals(Obj IObject) bool {
	return Printer_Equals(p._instance(), CheckPtr(Obj))
}

// GetHashCode
//
// 获取类的哈希值。
//
// Get the hash value of the class.
func (p *TPrinter) GetHashCode() int32 {
	return Printer_GetHashCode(p._instance())
}

// ToString
//
// 文本类信息。
//
// Text information.
func (p *TPrinter) ToString() string {
	return Printer_ToString(p._instance())
}

func (p *TPrinter) Aborted() bool {
	return Printer_GetAborted(p._instance())
}

// Canvas
//
// 获取画布。
func (p *TPrinter) Canvas() *TCanvas {
	return AsCanvas(Printer_GetCanvas(p._instance()))
}

func (p *TPrinter) Copies() int32 {
	return Printer_GetCopies(p._instance())
}

func (p *TPrinter) SetCopies(value int32) {
	Printer_SetCopies(p._instance(), value)
}

func (p *TPrinter) Fonts() *TStrings {
	return AsStrings(Printer_GetFonts(p._instance()))
}

func (p *TPrinter) Orientation() TPrinterOrientation {
	return Printer_GetOrientation(p._instance())
}

func (p *TPrinter) SetOrientation(value TPrinterOrientation) {
	Printer_SetOrientation(p._instance(), value)
}

func (p *TPrinter) PageHeight() int32 {
	return Printer_GetPageHeight(p._instance())
}

func (p *TPrinter) PageWidth() int32 {
	return Printer_GetPageWidth(p._instance())
}

func (p *TPrinter) PageNumber() int32 {
	return Printer_GetPageNumber(p._instance())
}

func (p *TPrinter) PrinterIndex() int32 {
	return Printer_GetPrinterIndex(p._instance())
}

func (p *TPrinter) SetPrinterIndex(value int32) {
	Printer_SetPrinterIndex(p._instance(), value)
}

func (p *TPrinter) Printing() bool {
	return Printer_GetPrinting(p._instance())
}

func (p *TPrinter) Printers() *TStrings {
	return AsStrings(Printer_GetPrinters(p._instance()))
}

func (p *TPrinter) Title() string {
	return Printer_GetTitle(p._instance())
}

func (p *TPrinter) SetTitle(value string) {
	Printer_SetTitle(p._instance(), value)
}
