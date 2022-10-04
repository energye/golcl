package main

import (
	"fmt"

	"strconv"

	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/types"
)

//::private::
type TForm1Fields struct {
}

func (f *TForm1) OnFormCreate(object lcl.IObject) {
	fmt.Println("OnForm1Create被调用了:", object.Instance())
	lcl.Application.SetTitle("我是新标题")

	f.SetKeyPreview(true)
	f.Edit2.SetText("1")
	f.Edit2.SetNumbersOnly(true)

	f.ActFileNew.SetCaption("我是一个共享的动作")

}

func (f *TForm1) OnButton1Click(object lcl.IObject) {
	fmt.Println("f:", f.Button1.Name())
	fmt.Println("窗口1的Button1单击")
	img := lcl.NewJPEGImage()
	defer img.Free()
	img.LoadFromFile("111.jpg")
}

func (f *TForm1) OnButton2Click(object lcl.IObject) {
	fmt.Println("f:", f.Button2.Name(), "xxx:", object.ClassName())
	fmt.Println("窗口1的Button2单击")
	result := Form2.ShowModal()
	if result == types.MrOk {
		lcl.ShowMessage("Form2返回了OK")
	} else if result == types.MrClose || result == types.MrNone {
		lcl.ShowMessage("Form2返回了Close")
	} else if result == types.MrCancel {
		lcl.ShowMessage("Form2返回了Cancel")
	}
}

func (f *TForm1) OnActExitExecute(sender lcl.IObject) {
	lcl.Application.Terminate()
}

func (f *TForm1) OnActFileNewExecute(sender lcl.IObject) {
	lcl.ShowMessage("ActFileNew Execute.")
}

func (f *TForm1) OnCheckBox1Click(sender lcl.IObject) {
	f.Button1.SetEnabled(f.CheckBox1.Checked())
}

func (f *TForm1) OnFormClick(sender lcl.IObject) {
	fmt.Println("OnForm1Click")
}

func (f *TForm1) OnFormDestroy(sender lcl.IObject) {
	fmt.Println("OnForm1Destroy")
}

func (f *TForm1) OnFormKeyDown(sender lcl.IObject, key *types.Char, shift types.TShiftState) {
	fmt.Println("OnForm1KeyDown:", *key, ", ", shift)
}

func (f *TForm1) OnFormMouseDown(sender lcl.IObject, button types.TMouseButton, shift types.TShiftState, x, y int32) {
	fmt.Println("OnForm1MouseDown:", button, ", ", shift, ", ", x, ", ", y)
}

func (f *TForm1) OnFormCloseQuery(sender lcl.IObject, canClose *bool) {
	fmt.Println("OnForm1CloseQuery")
	//*canClose = false
}

func (f *TForm1) OnFormClose(sender lcl.IObject, action *types.TCloseAction) {
	fmt.Println("OnForm1Close")
	//*canClose = false
}

func (f *TForm1) OnApplicationException(sender lcl.IObject, e *lcl.Exception) {
	fmt.Println("OnApplicationException")
	lcl.ShowMessage(e.Message())
}

func (f *TForm1) OnApplicationMinimize(sender lcl.IObject) {
	fmt.Println("OnApplicationMinimize")
}

func (f *TForm1) OnApplicationRestore(sender lcl.IObject) {
	fmt.Println("OnApplicationRestore")
}

func (f *TForm1) OnComboBox1Change(sender lcl.IObject) {
	fmt.Println("OnComboBox1Change: ", f.ComboBox1.ItemIndex())
}

func (f *TForm1) OnImageButton1Click(sender lcl.IObject) {
	lcl.ShowMessage("查杀，查杀，我查。。。")
}

func (f *TForm1) OnEdit2Change(sender lcl.IObject) {
	val, _ := strconv.Atoi(f.Edit2.Text())
	if val >= 0 && val <= 100 {
		f.Gauge1.SetProgress(int32(val))
	}
}

func (f *TForm1) OnMemo1Click(lcl.IObject) {
	fmt.Println("OnMemo1Click")
}

func (f *TForm1) OnListView1Click(lcl.IObject) {
	idx := f.ListView1.ItemIndex()
	if idx != -1 {
		item := f.ListView1.Items().Item(idx)
		fmt.Println(item.Caption())
		fmt.Println("subitem:", item.SubItems().Text())
	}
}

//func (f *TForm1) OnApplicationHint(sender lcl.IObject) {
//	fmt.Println("提示了：", lcl.Application.Hint())
//}
