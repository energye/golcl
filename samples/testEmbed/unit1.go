// Automatically generated by the res2go IDE plug-in, do not edit.
package main

import (
	_ "embed"
	"github.com/energye/golcl/lcl"
)

type TForm1 struct {
	*lcl.TForm
	Button1 *lcl.TButton

	//::private::
	TForm1Fields
}

var Form1 *TForm1

// lcl.Application.CreateForm(&Form1)

func NewForm1(owner lcl.IComponent) (root *TForm1) {
	lcl.CreateResForm(owner, &root)
	return
}

//go:embed resources/unit1.gfm
var form1Bytes []byte

// Register Form Resources
var _ = lcl.RegisterFormResource(Form1, &form1Bytes)