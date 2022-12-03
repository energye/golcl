package main

import (
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/types"
	"github.com/energye/golcl/lcl/types/colors"
	_ "github.com/energye/golcl/pkgs/winappres"
)

var (
	mainForm *lcl.TForm
	richEdit *lcl.TRichEdit
)

func main() {

	lcl.Application.Initialize()
	lcl.Application.SetMainFormOnTaskBar(true)

	initMainForm()
	initMainMenu()
	tlbar := lcl.NewToolBar(mainForm)
	tlbar.SetParent(mainForm)

	richEdit = lcl.NewRichEdit(mainForm)
	richEdit.SetParent(mainForm)
	richEdit.SetAlign(types.AlClient)
	richEdit.Lines().Add("这是一段文字红色，粗体，斜體")
	richEdit.SetSelStart(6)
	richEdit.SetSelLength(2)
	richEdit.SelAttributes().SetColor(colors.ClRed)

	richEdit.SetSelStart(9)
	richEdit.SetSelLength(2)

	richEdit.SelAttributes().SetStyle(types.NewSet(types.FsBold))

	richEdit.SetSelStart(12)
	richEdit.SetSelLength(2)

	richEdit.SelAttributes().SetStyle(types.NewSet(types.FsItalic))

	richEdit.SetSelStart(15)

	initRichEditPopupMenu()

	stabar := lcl.NewStatusBar(mainForm)
	stabar.SetParent(mainForm)

	lcl.Application.Run()
}

func initMainForm() {
	mainForm = lcl.Application.CreateForm()
	mainForm.SetCaption("Hello")
	mainForm.SetPosition(types.PoScreenCenter)
	mainForm.EnabledMaximize(false)
	mainForm.SetWidth(600)
	mainForm.SetHeight(400)
}

func initMainMenu() {
	mainMenu := lcl.NewMainMenu(mainForm)

	item := lcl.NewMenuItem(mainForm)
	item.SetCaption("&File")
	mainMenu.Items().Add(item)

	item = lcl.NewMenuItem(mainForm)
	item.SetCaption("&Help")
	mainMenu.Items().Add(item)
}

func initRichEditPopupMenu() {
	pm := lcl.NewPopupMenu(mainForm)
	item := lcl.NewMenuItem(mainForm)
	item.SetCaption("&Clear")
	pm.Items().Add(item)

	richEdit.SetPopupMenu(pm)
}
