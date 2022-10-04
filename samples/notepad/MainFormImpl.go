package main

import (
	"time"

	"fmt"

	"github.com/energye/golcl/lcl"
)

//::private::
type TMainFormFields struct {
}

// 此方法会在将资源窗口处理完后调用，命名规则为  On+窗口Name+Create，除TForm外
func (m *TMainForm) OnFormCreate(sender lcl.IObject) {
	// 这里可以初始化些东西

}

// File

func (m *TMainForm) OnActEditCopyUpdate(sender lcl.IObject) {
	lcl.AsAction(sender).SetEnabled(MainForm.MemoBody.SelLength() > 0)
}

func (m *TMainForm) OnActFileNewExecute(sender lcl.IObject) {
	m.MemoBody.Clear()
}

func (m *TMainForm) OnActFileOpenExecute(sender lcl.IObject) {
	if m.DlgOpen.Execute() {
		m.MemoBody.Lines().LoadFromFile(m.DlgOpen.FileName())
	}
}

func (m *TMainForm) OnActFileSaveExecute(sender lcl.IObject) {
	if m.DlgSave.Execute() {
		m.MemoBody.Lines().SaveToFile(m.DlgSave.FileName())
	}
}

func (m *TMainForm) OnActFileSaveAsExecute(sender lcl.IObject) {
	if m.DlgSave.Execute() {
		m.MemoBody.Lines().SaveToFile(m.DlgSave.FileName())
	}
}

func (m *TMainForm) OnActFilePageSettingsExecute(sender lcl.IObject) {
	m.PageSetupDialog1.Execute()
}

func (m *TMainForm) OnActFilePrintExecute(sender lcl.IObject) {
	m.PrinterSetupDialog1.Execute()
}

func (m *TMainForm) OnActFileExitExecute(sender lcl.IObject) {
	lcl.Application.Terminate()
}

// Edit
func (m *TMainForm) OnActEditUndoExecute(sender lcl.IObject) {
	m.MemoBody.Undo()
}

func (m *TMainForm) OnActEditUndoUpdate(sender lcl.IObject) {
	lcl.AsAction(sender).SetEnabled(m.MemoBody.CanUndo())
}

func (m *TMainForm) OnActEditCopyExecute(sender lcl.IObject) {
	m.MemoBody.CopyToClipboard()
}

func (m *TMainForm) OnActEditCutExecute(sender lcl.IObject) {
	m.MemoBody.CutToClipboard()
}

func (m *TMainForm) OnActEditPasteExecute(sender lcl.IObject) {
	m.MemoBody.PasteFromClipboard()
}

func (m *TMainForm) OnActEditFindExecute(sender lcl.IObject) {
	m.FindDialog1.Execute()
}

func (m *TMainForm) OnFindDialog1Find(sender lcl.IObject) {
	fmt.Println("查找 查找")
}

func (m *TMainForm) OnActEditReplaceExecute(sender lcl.IObject) {
	m.ReplaceDialog1.Execute()
}

func (m *TMainForm) OnReplaceDialog1Find(sender lcl.IObject) {
	fmt.Println("替换 查找")
}

func (m *TMainForm) OnReplaceDialog1Replace(sender lcl.IObject) {
	fmt.Println("替换 替换")
}

func (m *TMainForm) OnActEditGoLineExecute(sender lcl.IObject) {
	s := lcl.InputBox("跳转行数", "行数：", "")
	if s != "" {

	}
}

func (m *TMainForm) OnActEditDeleteExecute(sender lcl.IObject) {
	m.MemoBody.SetSelText("")
}

func (m *TMainForm) OnActEditSelectAllExecute(sender lcl.IObject) {
	m.MemoBody.SelectAll()
}

func (m *TMainForm) OnActEditSelectAllUpdate(sender lcl.IObject) {
	lcl.AsAction(sender).SetEnabled(m.MemoBody.GetTextLen() > 0)
}

func (m *TMainForm) OnActEditTimeOrDateExecute(sender lcl.IObject) {
	m.MemoBody.Lines().Add(time.Now().Format("2006-01-02 15:04:05"))
}

// Format

func (m *TMainForm) OnActFormatWordWapExecute(sender lcl.IObject) {
	val := !lcl.AsAction(sender).Checked()
	lcl.AsAction(sender).SetChecked(val)
	m.StatusBar1.SetVisible(!val && m.ActViewStatusBar.Checked())
	m.MemoBody.SetWordWrap(val)
	fmt.Println(val)
}

func (m *TMainForm) OnActFormatFontExecute(sender lcl.IObject) {
	if m.FontDialog1.Execute() {
		m.MemoBody.SetFont(MainForm.FontDialog1.Font())
	}
}

func (m *TMainForm) OnActViewStatusBarExecute(sender lcl.IObject) {
	val := !lcl.AsAction(sender).Checked()
	lcl.AsAction(sender).SetChecked(val)
	m.StatusBar1.SetVisible(val)
}

func (m *TMainForm) OnActViewStatusBarUpdate(sender lcl.IObject) {
	lcl.AsAction(sender).SetEnabled(!m.ActFormatWordWap.Checked())
}
