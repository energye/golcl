// Automatically generated by the res2go, do not edit.

package main

import (
	"github.com/energye/golcl/lcl"
)

type TMainForm struct {
	*lcl.TForm
	Panel1                               *lcl.TPanel
	BtnPaste                             *lcl.TButton
	BtnCopy                              *lcl.TButton
	ChkTryConvInt                        *lcl.TCheckBox
	ChkRemoveUnderlineAndTryUseCamelCase *lcl.TCheckBox
	Panel2                               *lcl.TPanel
	Splitter1                            *lcl.TSplitter
	MmoInput                             *lcl.TMemo
	MmoOutput                            *lcl.TMemo

	//::private::
	TMainFormFields
}

var MainForm *TMainForm

// Loaded in bytes.
// lcl.Application.CreateForm(&MainForm)

func NewMainForm(owner lcl.IComponent) (root *TMainForm) {
	lcl.CreateResForm(owner, &root)
	return
}

var MainFormBytes = []byte("\x54\x50\x46\x30\x0B\x54\x44\x65\x73\x69\x67\x6E\x46\x6F\x72\x6D\x08\x4D\x61\x69\x6E\x46\x6F\x72\x6D\x04\x4C\x65\x66\x74\x02\x08\x06\x48\x65\x69\x67\x68\x74\x03\xBE\x02\x03\x54\x6F\x70\x02\x08\x05\x57\x69\x64\x74\x68\x03\x44\x04\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x08\x6A\x73\x6F\x6E\x54\x6F\x47\x6F\x0C\x43\x6C\x69\x65\x6E\x74\x48\x65\x69\x67\x68\x74\x03\xBE\x02\x0B\x43\x6C\x69\x65\x6E\x74\x57\x69\x64\x74\x68\x03\x44\x04\x05\x43\x6F\x6C\x6F\x72\x07\x09\x63\x6C\x42\x74\x6E\x46\x61\x63\x65\x0A\x46\x6F\x6E\x74\x2E\x43\x6F\x6C\x6F\x72\x07\x0C\x63\x6C\x57\x69\x6E\x64\x6F\x77\x54\x65\x78\x74\x0B\x46\x6F\x6E\x74\x2E\x48\x65\x69\x67\x68\x74\x02\xF0\x09\x46\x6F\x6E\x74\x2E\x4E\x61\x6D\x65\x06\x06\x54\x61\x68\x6F\x6D\x61\x08\x50\x6F\x73\x69\x74\x69\x6F\x6E\x07\x0E\x70\x6F\x53\x63\x72\x65\x65\x6E\x43\x65\x6E\x74\x65\x72\x06\x53\x63\x61\x6C\x65\x64\x08\x00\x06\x54\x50\x61\x6E\x65\x6C\x06\x50\x61\x6E\x65\x6C\x31\x04\x4C\x65\x66\x74\x02\x00\x06\x48\x65\x69\x67\x68\x74\x02\x4A\x03\x54\x6F\x70\x02\x00\x05\x57\x69\x64\x74\x68\x03\x44\x04\x05\x41\x6C\x69\x67\x6E\x07\x05\x61\x6C\x54\x6F\x70\x0C\x43\x6C\x69\x65\x6E\x74\x48\x65\x69\x67\x68\x74\x02\x4A\x0B\x43\x6C\x69\x65\x6E\x74\x57\x69\x64\x74\x68\x03\x44\x04\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x00\x00\x07\x54\x42\x75\x74\x74\x6F\x6E\x08\x42\x74\x6E\x50\x61\x73\x74\x65\x04\x4C\x65\x66\x74\x03\xCD\x00\x06\x48\x65\x69\x67\x68\x74\x02\x19\x03\x54\x6F\x70\x02\x1A\x05\x57\x69\x64\x74\x68\x02\x4B\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x06\xE7\xB2\x98\xE8\xB4\xB4\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x00\x00\x00\x07\x54\x42\x75\x74\x74\x6F\x6E\x07\x42\x74\x6E\x43\x6F\x70\x79\x04\x4C\x65\x66\x74\x03\x1E\x03\x06\x48\x65\x69\x67\x68\x74\x02\x19\x03\x54\x6F\x70\x02\x1A\x05\x57\x69\x64\x74\x68\x02\x4B\x07\x41\x6E\x63\x68\x6F\x72\x73\x0B\x05\x61\x6B\x54\x6F\x70\x07\x61\x6B\x52\x69\x67\x68\x74\x00\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x06\xE5\xA4\x8D\xE5\x88\xB6\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x01\x00\x00\x09\x54\x43\x68\x65\x63\x6B\x42\x6F\x78\x0D\x43\x68\x6B\x54\x72\x79\x43\x6F\x6E\x76\x49\x6E\x74\x04\x4C\x65\x66\x74\x03\x92\x01\x06\x48\x65\x69\x67\x68\x74\x02\x17\x03\x54\x6F\x70\x02\x10\x05\x57\x69\x64\x74\x68\x03\xA8\x00\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x1B\xE5\xB0\x9D\xE8\xAF\x95\xE5\xB0\x86\xE5\xAD\x97\xE7\xAC\xA6\xE4\xB8\xB2\xE8\xBD\xAC\xE4\xB8\xBA\x69\x6E\x74\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x02\x00\x00\x09\x54\x43\x68\x65\x63\x6B\x42\x6F\x78\x24\x43\x68\x6B\x52\x65\x6D\x6F\x76\x65\x55\x6E\x64\x65\x72\x6C\x69\x6E\x65\x41\x6E\x64\x54\x72\x79\x55\x73\x65\x43\x61\x6D\x65\x6C\x43\x61\x73\x65\x04\x4C\x65\x66\x74\x03\x92\x01\x06\x48\x65\x69\x67\x68\x74\x02\x17\x03\x54\x6F\x70\x02\x29\x05\x57\x69\x64\x74\x68\x03\xF6\x00\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x2A\xE7\xA7\xBB\xE9\x99\xA4\xE4\xB8\x8B\xE5\x88\x92\xE7\xBA\xBF\xEF\xBC\x8C\xE5\xB0\x9D\xE8\xAF\x95\xE4\xBD\xBF\xE7\x94\xA8\xE5\xA4\xA7\xE9\xA9\xBC\xE5\xB3\xB0\xE6\xB3\x95\x07\x43\x68\x65\x63\x6B\x65\x64\x09\x05\x53\x74\x61\x74\x65\x07\x09\x63\x62\x43\x68\x65\x63\x6B\x65\x64\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x03\x00\x00\x00\x06\x54\x50\x61\x6E\x65\x6C\x06\x50\x61\x6E\x65\x6C\x32\x04\x4C\x65\x66\x74\x02\x00\x06\x48\x65\x69\x67\x68\x74\x03\x74\x02\x03\x54\x6F\x70\x02\x4A\x05\x57\x69\x64\x74\x68\x03\x44\x04\x05\x41\x6C\x69\x67\x6E\x07\x08\x61\x6C\x43\x6C\x69\x65\x6E\x74\x0C\x43\x6C\x69\x65\x6E\x74\x48\x65\x69\x67\x68\x74\x03\x74\x02\x0B\x43\x6C\x69\x65\x6E\x74\x57\x69\x64\x74\x68\x03\x44\x04\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x01\x00\x09\x54\x53\x70\x6C\x69\x74\x74\x65\x72\x09\x53\x70\x6C\x69\x74\x74\x65\x72\x31\x04\x4C\x65\x66\x74\x03\x15\x02\x06\x48\x65\x69\x67\x68\x74\x03\x72\x02\x03\x54\x6F\x70\x02\x01\x05\x57\x69\x64\x74\x68\x02\x05\x00\x00\x05\x54\x4D\x65\x6D\x6F\x08\x4D\x6D\x6F\x49\x6E\x70\x75\x74\x04\x4C\x65\x66\x74\x02\x01\x06\x48\x65\x69\x67\x68\x74\x03\x72\x02\x03\x54\x6F\x70\x02\x01\x05\x57\x69\x64\x74\x68\x03\x14\x02\x05\x41\x6C\x69\x67\x6E\x07\x06\x61\x6C\x4C\x65\x66\x74\x0D\x4C\x69\x6E\x65\x73\x2E\x53\x74\x72\x69\x6E\x67\x73\x01\x06\x01\x7B\x06\x0B\x20\x20\x20\x22\x6B\x65\x79\x22\x3A\x31\x2C\x06\x0B\x20\x20\x20\x22\x6B\x65\x79\x32\x22\x3A\x5B\x06\x20\x20\x20\x20\x7B\x22\x61\x61\x61\x22\x3A\x22\x61\x61\x61\x61\x22\x2C\x20\x22\x61\x31\x32\x32\x22\x3A\x22\x61\x31\x22\x7D\x2C\x20\x06\x11\x20\x20\x20\x7B\x22\x62\x62\x62\x22\x3A\x20\x22\x62\x62\x62\x22\x7D\x06\x04\x20\x20\x5D\x2C\x06\x0B\x20\x20\x22\x6B\x65\x79\x33\x22\x3A\x20\x5B\x06\x0B\x20\x20\x20\x20\x20\x31\x2C\x20\x32\x2C\x33\x06\x05\x20\x20\x20\x5D\x2C\x06\x0E\x20\x20\x20\x20\x20\x22\x6B\x65\x79\x34\x22\x3A\x20\x5B\x06\x13\x20\x20\x20\x20\x20\x22\x73\x74\x72\x31\x22\x2C\x20\x22\x73\x74\x72\x32\x22\x06\x04\x20\x20\x20\x5D\x06\x01\x7D\x00\x0A\x53\x63\x72\x6F\x6C\x6C\x42\x61\x72\x73\x07\x06\x73\x73\x42\x6F\x74\x68\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x00\x00\x00\x05\x54\x4D\x65\x6D\x6F\x09\x4D\x6D\x6F\x4F\x75\x74\x70\x75\x74\x04\x4C\x65\x66\x74\x03\x1A\x02\x06\x48\x65\x69\x67\x68\x74\x03\x72\x02\x03\x54\x6F\x70\x02\x01\x05\x57\x69\x64\x74\x68\x03\x29\x02\x05\x41\x6C\x69\x67\x6E\x07\x08\x61\x6C\x43\x6C\x69\x65\x6E\x74\x08\x52\x65\x61\x64\x4F\x6E\x6C\x79\x09\x0A\x53\x63\x72\x6F\x6C\x6C\x42\x61\x72\x73\x07\x06\x73\x73\x42\x6F\x74\x68\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x01\x00\x00\x00\x00")

// Register Form Resources
var _ = lcl.RegisterFormResource(MainForm, &MainFormBytes)
