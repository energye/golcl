package main

import (
	"fmt"
	"github.com/energye/golcl/energy/inits"
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/types"
	_ "github.com/energye/golcl/pkgs/winappres"
)

type TMainForm struct {
	*lcl.TForm
	Button1 *lcl.TButton
}

type TForm1 struct {
	*lcl.TForm
	Button1 *lcl.TButton
}

var (
	mainForm *TMainForm
	form1    *TForm1
)

func main() {
	lcl.DEBUG = true
	inits.Init(nil, nil)
	lcl.RunApp(&mainForm, &form1)
}

// --------------MainForm -----------------
func (f *TMainForm) OnFormCreate(sender lcl.IObject) {
	f.SetCaption("Hello")
	f.EnabledMaximize(false)
	f.SetWidth(600)
	f.SetHeight(600)
	f.ScreenCenter()

	f.Button1 = lcl.NewButton(f)
	f.Button1.SetParent(f)
	f.Button1.SetCaption("窗口1")
	f.Button1.SetLeft(50)
	f.Button1.SetTop(50)
	f.Button1.SetOnClick(f.OnButton1Click)
	//	f.Button1.Font().SetStyle(types.NewSet(types.FsBold)) //f.Button1.Font().Style().Include(types.FsBold))
	//f.Button1.Hide()

	//cbb := lcl.NewComboBox(f)
	//cbb.SetParent(f)
	//cbb.SetLeft(100)
	//cbb.SetTop(100)
	//cbb.SetStyle(types.CsOwnerDrawVariable)
	//cbb.Items().Add("1111")
	//cbb.Items().Add("2222")
	//cbb.Items().Add("3333")
	//cbb.SetOnDrawItem(func(control lcl.IWinControl, index int32, aRect types.TRect, state types.TOwnerDrawState) {
	//	switch index {
	//	case 0:
	//		cbb.Canvas().Font().SetColor(colors.ClRed)
	//	case 1:
	//		cbb.Canvas().Font().SetColor(colors.ClGreen)
	//	case 2:
	//		cbb.Canvas().Font().SetColor(colors.ClBlue)
	//	}
	//	cbb.Canvas().TextOut(aRect.Left, aRect.Top, cbb.Items().Strings(index))
	//})
	//
	//f.SetOnUTF8KeyPress(func(sender lcl.IObject, utf8key *types.TUTF8Char) {
	//	fmt.Println("打印：1111", utf8key.ToString(), utf8key)
	//	utf8key.SetString("这") //每次只一个字符
	//	fmt.Println("打印：2222", utf8key.ToString(), utf8key)
	//})
	//xx := lcl.NewCheckComboBox(f)
	//xx.SetParent(f)
	//xx.Items().Add("fff")
	//xx.Items().Add("bbb")
	//xx.Items().Add("ccc")
	//xx.SetOnItemChange(func(sender lcl.IObject, index int32) {
	//	fmt.Println("checked: ", xx.Checked(index))
	//})
	//x := lcl.NewFloatSpinEdit(f)
	//x.SetParent(f)
	//x.SetLeft(100)
	//x.SetTop(100)
	//x.SetMaxValue(10.0)
	//x.SetIncrement(0.2)
	//x.SetMinValue(0.1)
	//x.SetValue(3.0)
	//x.SetWidth(100)
	//
	//d := lcl.NewDirectoryEdit(f)
	//d.SetParent(f)
	//d.SetLeft(100)
	//d.SetTop(150)
	//d.SetWidth(200)
	//d.SetDirectory("C:\\xxxxx")

	//c := lcl.NewColorButton(f)
	//c.SetParent(f)
	//c.SetOnColorChanged(func(sender lcl.IObject) {
	//	fmt.Println(c.ButtonColor())
	//})

	//pnl := lcl.NewPanel(f)
	//pnl.SetParent(f)
	//pnl.SetBounds(20, 150, 200, 300)
	//pnl.SetOnMouseMove(func(sender lcl.IObject, shift types.TShiftState, x, y int32) {
	//	// x < 5时显示左右调节指针样式
	//	if x < 5 {
	//		pnl.SetCursor(types.CrSizeE)
	//		// y<5时显示上下调节指针样式
	//	} else if y < 5 {
	//		pnl.SetCursor(types.CrSizeN)
	//	} else {
	//		pnl.SetCursor(types.CrDefault)
	//	}
	//})

}

func (f *TMainForm) OnFormCloseQuery(Sender lcl.IObject, CanClose *bool) {
	*CanClose = lcl.MessageDlg("是否退出？", types.MtConfirmation, types.MbYes, types.MbNo) == types.IdYes
}

func (f *TMainForm) OnButton1Click(object lcl.IObject) {
	//_SetParent.Call(0x50C40, uintptr(f.Handle()))
	//	form1.Show()
	fmt.Println("清除事件")
	//f.Button1.SetOnClick(nil)
	f.Button1.SetOnClick(f.OnButton1Click)
	fmt.Println("更换事件")
	f.Button1.SetOnClick(f.OnButton2Click)
}

func (f *TMainForm) OnButton2Click(object lcl.IObject) {
	fmt.Println("换成button2click事件了啊")
}

// ---------- Form1 ----------------

func (f *TForm1) OnFormCreate(sender lcl.IObject) {
	fmt.Println("onCreate")
	f.Button1 = lcl.NewButton(f)
	fmt.Println("f.Button1:", f.Button1.Instance())
	f.Button1.SetParent(f)
	//f.Button1.SetName("Button1")
	f.Button1.SetCaption("我是按钮")
	f.Button1.SetOnClick(f.OnButton1Click)
}

func (f *TForm1) OnButton1Click(object lcl.IObject) {
	lcl.ShowMessage("Click")
}
