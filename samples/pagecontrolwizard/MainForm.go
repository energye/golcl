// Automatically generated by the res2go, do not edit.

package main

import (
	"github.com/energye/golcl/lcl"
)

type TMainForm struct {
	*lcl.TForm
	Panel1       *lcl.TPanel
	PageControl1 *lcl.TPageControl
	TabSheet1    *lcl.TTabSheet
	Button1      *lcl.TButton
	TabSheet2    *lcl.TTabSheet
	Button2      *lcl.TButton
	TabSheet3    *lcl.TTabSheet
	Button3      *lcl.TButton
	TabSheet4    *lcl.TTabSheet
	Button4      *lcl.TButton
	TabSheet5    *lcl.TTabSheet
	Button5      *lcl.TButton
	TabSheet6    *lcl.TTabSheet
	Button6      *lcl.TButton
	TabSheet7    *lcl.TTabSheet
	Button7      *lcl.TButton
	Panel2       *lcl.TPanel
	Bevel1       *lcl.TBevel
	Button8      *lcl.TButton
	Button9      *lcl.TButton
	ActionList1  *lcl.TActionList
	ActPagePrev  *lcl.TAction
	ActPageNext  *lcl.TAction

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

var MainFormBytes = []byte("\x54\x50\x46\x30\x0B\x54\x44\x65\x73\x69\x67\x6E\x46\x6F\x72\x6D\x08\x4D\x61\x69\x6E\x46\x6F\x72\x6D\x04\x4C\x65\x66\x74\x02\x08\x06\x48\x65\x69\x67\x68\x74\x03\xA8\x01\x03\x54\x6F\x70\x02\x08\x05\x57\x69\x64\x74\x68\x03\xE2\x02\x0B\x42\x6F\x72\x64\x65\x72\x49\x63\x6F\x6E\x73\x0B\x0C\x62\x69\x53\x79\x73\x74\x65\x6D\x4D\x65\x6E\x75\x0A\x62\x69\x4D\x69\x6E\x69\x6D\x69\x7A\x65\x00\x0B\x42\x6F\x72\x64\x65\x72\x53\x74\x79\x6C\x65\x07\x08\x62\x73\x53\x69\x6E\x67\x6C\x65\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x1D\x50\x61\x67\x65\x43\x6F\x6E\x74\x72\x6F\x6C\xE5\x90\x91\xE5\xAF\xBC\xE7\xA8\x8B\xE5\xBA\x8F\xE6\xBC\x94\xE7\xA4\xBA\x0C\x43\x6C\x69\x65\x6E\x74\x48\x65\x69\x67\x68\x74\x03\xA8\x01\x0B\x43\x6C\x69\x65\x6E\x74\x57\x69\x64\x74\x68\x03\xE2\x02\x05\x43\x6F\x6C\x6F\x72\x07\x09\x63\x6C\x42\x74\x6E\x46\x61\x63\x65\x0E\x44\x6F\x75\x62\x6C\x65\x42\x75\x66\x66\x65\x72\x65\x64\x09\x0A\x46\x6F\x6E\x74\x2E\x43\x6F\x6C\x6F\x72\x07\x0C\x63\x6C\x57\x69\x6E\x64\x6F\x77\x54\x65\x78\x74\x0B\x46\x6F\x6E\x74\x2E\x48\x65\x69\x67\x68\x74\x02\xF3\x09\x46\x6F\x6E\x74\x2E\x4E\x61\x6D\x65\x06\x06\x54\x61\x68\x6F\x6D\x61\x14\x50\x61\x72\x65\x6E\x74\x44\x6F\x75\x62\x6C\x65\x42\x75\x66\x66\x65\x72\x65\x64\x08\x08\x50\x6F\x73\x69\x74\x69\x6F\x6E\x07\x0E\x70\x6F\x53\x63\x72\x65\x65\x6E\x43\x65\x6E\x74\x65\x72\x00\x06\x54\x50\x61\x6E\x65\x6C\x06\x50\x61\x6E\x65\x6C\x31\x04\x4C\x65\x66\x74\x02\x00\x06\x48\x65\x69\x67\x68\x74\x03\x6A\x01\x03\x54\x6F\x70\x02\x00\x05\x57\x69\x64\x74\x68\x03\xE2\x02\x05\x41\x6C\x69\x67\x6E\x07\x08\x61\x6C\x43\x6C\x69\x65\x6E\x74\x0A\x42\x65\x76\x65\x6C\x4F\x75\x74\x65\x72\x07\x06\x62\x76\x4E\x6F\x6E\x65\x0C\x43\x6C\x69\x65\x6E\x74\x48\x65\x69\x67\x68\x74\x03\x6A\x01\x0B\x43\x6C\x69\x65\x6E\x74\x57\x69\x64\x74\x68\x03\xE2\x02\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x00\x00\x0C\x54\x50\x61\x67\x65\x43\x6F\x6E\x74\x72\x6F\x6C\x0C\x50\x61\x67\x65\x43\x6F\x6E\x74\x72\x6F\x6C\x31\x04\x4C\x65\x66\x74\x02\x00\x06\x48\x65\x69\x67\x68\x74\x03\x6A\x01\x03\x54\x6F\x70\x02\x00\x05\x57\x69\x64\x74\x68\x03\xE2\x02\x0A\x41\x63\x74\x69\x76\x65\x50\x61\x67\x65\x07\x09\x54\x61\x62\x53\x68\x65\x65\x74\x37\x05\x41\x6C\x69\x67\x6E\x07\x08\x61\x6C\x43\x6C\x69\x65\x6E\x74\x08\x54\x61\x62\x49\x6E\x64\x65\x78\x02\x06\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x00\x0B\x54\x61\x62\x50\x6F\x73\x69\x74\x69\x6F\x6E\x07\x08\x74\x70\x42\x6F\x74\x74\x6F\x6D\x00\x09\x54\x54\x61\x62\x53\x68\x65\x65\x74\x09\x54\x61\x62\x53\x68\x65\x65\x74\x31\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x09\x54\x61\x62\x53\x68\x65\x65\x74\x31\x0C\x43\x6C\x69\x65\x6E\x74\x48\x65\x69\x67\x68\x74\x02\x00\x0B\x43\x6C\x69\x65\x6E\x74\x57\x69\x64\x74\x68\x02\x00\x00\x07\x54\x42\x75\x74\x74\x6F\x6E\x07\x42\x75\x74\x74\x6F\x6E\x31\x04\x4C\x65\x66\x74\x03\x20\x01\x06\x48\x65\x69\x67\x68\x74\x02\x19\x03\x54\x6F\x70\x03\x90\x00\x05\x57\x69\x64\x74\x68\x02\x4B\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x09\xE7\xAC\xAC\xE4\xB8\x80\xE9\xA1\xB5\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x00\x00\x00\x00\x09\x54\x54\x61\x62\x53\x68\x65\x65\x74\x09\x54\x61\x62\x53\x68\x65\x65\x74\x32\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x09\x54\x61\x62\x53\x68\x65\x65\x74\x32\x0C\x43\x6C\x69\x65\x6E\x74\x48\x65\x69\x67\x68\x74\x02\x00\x0B\x43\x6C\x69\x65\x6E\x74\x57\x69\x64\x74\x68\x02\x00\x00\x07\x54\x42\x75\x74\x74\x6F\x6E\x07\x42\x75\x74\x74\x6F\x6E\x32\x04\x4C\x65\x66\x74\x03\x2A\x01\x06\x48\x65\x69\x67\x68\x74\x02\x19\x03\x54\x6F\x70\x03\x9A\x00\x05\x57\x69\x64\x74\x68\x02\x4B\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x09\xE7\xAC\xAC\xE4\xBA\x8C\xE9\xA1\xB5\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x00\x00\x00\x00\x09\x54\x54\x61\x62\x53\x68\x65\x65\x74\x09\x54\x61\x62\x53\x68\x65\x65\x74\x33\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x09\x54\x61\x62\x53\x68\x65\x65\x74\x33\x0C\x43\x6C\x69\x65\x6E\x74\x48\x65\x69\x67\x68\x74\x02\x00\x0B\x43\x6C\x69\x65\x6E\x74\x57\x69\x64\x74\x68\x02\x00\x00\x07\x54\x42\x75\x74\x74\x6F\x6E\x07\x42\x75\x74\x74\x6F\x6E\x33\x04\x4C\x65\x66\x74\x03\x2A\x01\x06\x48\x65\x69\x67\x68\x74\x02\x19\x03\x54\x6F\x70\x03\x9A\x00\x05\x57\x69\x64\x74\x68\x02\x4B\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x09\xE7\xAC\xAC\xE4\xB8\x89\xE9\xA1\xB5\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x00\x00\x00\x00\x09\x54\x54\x61\x62\x53\x68\x65\x65\x74\x09\x54\x61\x62\x53\x68\x65\x65\x74\x34\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x09\x54\x61\x62\x53\x68\x65\x65\x74\x34\x0C\x43\x6C\x69\x65\x6E\x74\x48\x65\x69\x67\x68\x74\x02\x00\x0B\x43\x6C\x69\x65\x6E\x74\x57\x69\x64\x74\x68\x02\x00\x00\x07\x54\x42\x75\x74\x74\x6F\x6E\x07\x42\x75\x74\x74\x6F\x6E\x34\x04\x4C\x65\x66\x74\x03\x2A\x01\x06\x48\x65\x69\x67\x68\x74\x02\x19\x03\x54\x6F\x70\x03\x9A\x00\x05\x57\x69\x64\x74\x68\x02\x4B\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x09\xE7\xAC\xAC\xE5\x9B\x9B\xE9\xA1\xB5\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x00\x00\x00\x00\x09\x54\x54\x61\x62\x53\x68\x65\x65\x74\x09\x54\x61\x62\x53\x68\x65\x65\x74\x35\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x09\x54\x61\x62\x53\x68\x65\x65\x74\x35\x0C\x43\x6C\x69\x65\x6E\x74\x48\x65\x69\x67\x68\x74\x02\x00\x0B\x43\x6C\x69\x65\x6E\x74\x57\x69\x64\x74\x68\x02\x00\x00\x07\x54\x42\x75\x74\x74\x6F\x6E\x07\x42\x75\x74\x74\x6F\x6E\x35\x04\x4C\x65\x66\x74\x03\x2A\x01\x06\x48\x65\x69\x67\x68\x74\x02\x19\x03\x54\x6F\x70\x03\x9A\x00\x05\x57\x69\x64\x74\x68\x02\x4B\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x09\xE7\xAC\xAC\xE4\xBA\x94\xE9\xA1\xB5\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x00\x00\x00\x00\x09\x54\x54\x61\x62\x53\x68\x65\x65\x74\x09\x54\x61\x62\x53\x68\x65\x65\x74\x36\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x09\x54\x61\x62\x53\x68\x65\x65\x74\x36\x0C\x43\x6C\x69\x65\x6E\x74\x48\x65\x69\x67\x68\x74\x02\x00\x0B\x43\x6C\x69\x65\x6E\x74\x57\x69\x64\x74\x68\x02\x00\x00\x07\x54\x42\x75\x74\x74\x6F\x6E\x07\x42\x75\x74\x74\x6F\x6E\x36\x04\x4C\x65\x66\x74\x03\x2A\x01\x06\x48\x65\x69\x67\x68\x74\x02\x19\x03\x54\x6F\x70\x03\x9A\x00\x05\x57\x69\x64\x74\x68\x02\x4B\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x09\xE7\xAC\xAC\xE5\x85\xAD\xE9\xA1\xB5\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x00\x00\x00\x00\x09\x54\x54\x61\x62\x53\x68\x65\x65\x74\x09\x54\x61\x62\x53\x68\x65\x65\x74\x37\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x09\x54\x61\x62\x53\x68\x65\x65\x74\x37\x0C\x43\x6C\x69\x65\x6E\x74\x48\x65\x69\x67\x68\x74\x03\x4C\x01\x0B\x43\x6C\x69\x65\x6E\x74\x57\x69\x64\x74\x68\x03\xDA\x02\x00\x07\x54\x42\x75\x74\x74\x6F\x6E\x07\x42\x75\x74\x74\x6F\x6E\x37\x04\x4C\x65\x66\x74\x03\x2A\x01\x06\x48\x65\x69\x67\x68\x74\x02\x19\x03\x54\x6F\x70\x03\x9A\x00\x05\x57\x69\x64\x74\x68\x02\x4B\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x09\xE7\xAC\xAC\xE4\xB8\x83\xE9\xA1\xB5\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x00\x00\x00\x00\x00\x00\x06\x54\x50\x61\x6E\x65\x6C\x06\x50\x61\x6E\x65\x6C\x32\x04\x4C\x65\x66\x74\x02\x00\x06\x48\x65\x69\x67\x68\x74\x02\x3E\x03\x54\x6F\x70\x03\x6A\x01\x05\x57\x69\x64\x74\x68\x03\xE2\x02\x05\x41\x6C\x69\x67\x6E\x07\x08\x61\x6C\x42\x6F\x74\x74\x6F\x6D\x0A\x42\x65\x76\x65\x6C\x4F\x75\x74\x65\x72\x07\x06\x62\x76\x4E\x6F\x6E\x65\x0C\x43\x6C\x69\x65\x6E\x74\x48\x65\x69\x67\x68\x74\x02\x3E\x0B\x43\x6C\x69\x65\x6E\x74\x57\x69\x64\x74\x68\x03\xE2\x02\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x01\x00\x06\x54\x42\x65\x76\x65\x6C\x06\x42\x65\x76\x65\x6C\x31\x04\x4C\x65\x66\x74\x02\x05\x06\x48\x65\x69\x67\x68\x74\x02\x07\x03\x54\x6F\x70\x02\x04\x05\x57\x69\x64\x74\x68\x03\xD8\x02\x05\x53\x68\x61\x70\x65\x07\x09\x62\x73\x54\x6F\x70\x4C\x69\x6E\x65\x00\x00\x07\x54\x42\x75\x74\x74\x6F\x6E\x07\x42\x75\x74\x74\x6F\x6E\x38\x04\x4C\x65\x66\x74\x03\x25\x02\x06\x48\x65\x69\x67\x68\x74\x02\x19\x03\x54\x6F\x70\x02\x13\x05\x57\x69\x64\x74\x68\x02\x4B\x06\x41\x63\x74\x69\x6F\x6E\x07\x0B\x41\x63\x74\x50\x61\x67\x65\x50\x72\x65\x76\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x00\x00\x00\x07\x54\x42\x75\x74\x74\x6F\x6E\x07\x42\x75\x74\x74\x6F\x6E\x39\x04\x4C\x65\x66\x74\x03\x80\x02\x06\x48\x65\x69\x67\x68\x74\x02\x19\x03\x54\x6F\x70\x02\x13\x05\x57\x69\x64\x74\x68\x02\x4B\x06\x41\x63\x74\x69\x6F\x6E\x07\x0B\x41\x63\x74\x50\x61\x67\x65\x4E\x65\x78\x74\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x01\x00\x00\x00\x0B\x54\x41\x63\x74\x69\x6F\x6E\x4C\x69\x73\x74\x0B\x41\x63\x74\x69\x6F\x6E\x4C\x69\x73\x74\x31\x04\x6C\x65\x66\x74\x03\x9E\x01\x03\x74\x6F\x70\x02\x76\x00\x07\x54\x41\x63\x74\x69\x6F\x6E\x0B\x41\x63\x74\x50\x61\x67\x65\x50\x72\x65\x76\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x09\xE4\xB8\x8A\xE4\xB8\x80\xE9\xA1\xB5\x00\x00\x07\x54\x41\x63\x74\x69\x6F\x6E\x0B\x41\x63\x74\x50\x61\x67\x65\x4E\x65\x78\x74\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x09\xE4\xB8\x8B\xE4\xB8\x80\xE9\xA1\xB5\x00\x00\x00\x00")

// Register Form Resources
var _ = lcl.RegisterFormResource(MainForm, &MainFormBytes)
