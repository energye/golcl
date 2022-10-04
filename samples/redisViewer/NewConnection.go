// 由res2go自动生成，不要编辑。
package main

import (
	"github.com/energye/golcl/lcl"
)

type TNewConnection struct {
	*lcl.TForm
	Panel1            *lcl.TPanel
	BtnTestConnection *lcl.TButton
	BtnOk             *lcl.TButton
	BtnCancel         *lcl.TButton
	Panel2            *lcl.TPanel
	Label1            *lcl.TLabel
	EdtConnName       *lcl.TEdit
	Label2            *lcl.TLabel
	EdtHost           *lcl.TEdit
	Label3            *lcl.TLabel
	EdtPort           *lcl.TEdit
	Label5            *lcl.TLabel
	EdtPassword       *lcl.TEdit

	//::private::
	TNewConnectionFields
}

var NewConnection *TNewConnection

// 以字节形式加载
// lcl.Application.CreateForm(&NewConnection)

func NewNewConnection(owner lcl.IComponent) (root *TNewConnection) {
	lcl.CreateResForm(owner, &root)
	return
}

var newConnectionBytes = []byte("\x54\x50\x46\x30\x0E\x54\x4E\x65\x77\x43\x6F\x6E\x6E\x65\x63\x74\x69\x6F\x6E\x0D\x4E\x65\x77\x43\x6F\x6E\x6E\x65\x63\x74\x69\x6F\x6E\x04\x4C\x65\x66\x74\x03\x9B\x02\x06\x48\x65\x69\x67\x68\x74\x03\x8D\x01\x03\x54\x6F\x70\x03\xFA\x00\x05\x57\x69\x64\x74\x68\x03\x60\x01\x0B\x42\x6F\x72\x64\x65\x72\x49\x63\x6F\x6E\x73\x0B\x0C\x62\x69\x53\x79\x73\x74\x65\x6D\x4D\x65\x6E\x75\x00\x0B\x42\x6F\x72\x64\x65\x72\x53\x74\x79\x6C\x65\x07\x08\x62\x73\x53\x69\x6E\x67\x6C\x65\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x06\x4E\x65\x77\x2E\x2E\x2E\x0C\x43\x6C\x69\x65\x6E\x74\x48\x65\x69\x67\x68\x74\x03\x8D\x01\x0B\x43\x6C\x69\x65\x6E\x74\x57\x69\x64\x74\x68\x03\x60\x01\x08\x4F\x6E\x43\x72\x65\x61\x74\x65\x07\x0A\x46\x6F\x72\x6D\x43\x72\x65\x61\x74\x65\x00\x06\x54\x50\x61\x6E\x65\x6C\x06\x50\x61\x6E\x65\x6C\x31\x04\x4C\x65\x66\x74\x02\x00\x06\x48\x65\x69\x67\x68\x74\x02\x32\x03\x54\x6F\x70\x03\x5B\x01\x05\x57\x69\x64\x74\x68\x03\x60\x01\x05\x41\x6C\x69\x67\x6E\x07\x08\x61\x6C\x42\x6F\x74\x74\x6F\x6D\x0A\x42\x65\x76\x65\x6C\x4F\x75\x74\x65\x72\x07\x06\x62\x76\x4E\x6F\x6E\x65\x0C\x43\x6C\x69\x65\x6E\x74\x48\x65\x69\x67\x68\x74\x02\x32\x0B\x43\x6C\x69\x65\x6E\x74\x57\x69\x64\x74\x68\x03\x60\x01\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x00\x00\x07\x54\x42\x75\x74\x74\x6F\x6E\x11\x42\x74\x6E\x54\x65\x73\x74\x43\x6F\x6E\x6E\x65\x63\x74\x69\x6F\x6E\x04\x4C\x65\x66\x74\x02\x08\x06\x48\x65\x69\x67\x68\x74\x02\x19\x03\x54\x6F\x70\x02\x10\x05\x57\x69\x64\x74\x68\x03\x8B\x00\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x0C\xE6\xB5\x8B\xE8\xAF\x95\xE8\xBF\x9E\xE6\x8E\xA5\x07\x4F\x6E\x43\x6C\x69\x63\x6B\x07\x16\x42\x74\x6E\x54\x65\x73\x74\x43\x6F\x6E\x6E\x65\x63\x74\x69\x6F\x6E\x43\x6C\x69\x63\x6B\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x00\x00\x00\x07\x54\x42\x75\x74\x74\x6F\x6E\x05\x42\x74\x6E\x4F\x6B\x04\x4C\x65\x66\x74\x03\xC0\x00\x06\x48\x65\x69\x67\x68\x74\x02\x19\x03\x54\x6F\x70\x02\x10\x05\x57\x69\x64\x74\x68\x02\x4B\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x06\xE7\xA1\xAE\xE5\xAE\x9A\x0B\x4D\x6F\x64\x61\x6C\x52\x65\x73\x75\x6C\x74\x02\x01\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x01\x00\x00\x07\x54\x42\x75\x74\x74\x6F\x6E\x09\x42\x74\x6E\x43\x61\x6E\x63\x65\x6C\x04\x4C\x65\x66\x74\x03\x0D\x01\x06\x48\x65\x69\x67\x68\x74\x02\x19\x03\x54\x6F\x70\x02\x10\x05\x57\x69\x64\x74\x68\x02\x4B\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x06\xE5\x8F\x96\xE6\xB6\x88\x0B\x4D\x6F\x64\x61\x6C\x52\x65\x73\x75\x6C\x74\x02\x02\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x02\x00\x00\x00\x06\x54\x50\x61\x6E\x65\x6C\x06\x50\x61\x6E\x65\x6C\x32\x04\x4C\x65\x66\x74\x02\x00\x06\x48\x65\x69\x67\x68\x74\x03\x5B\x01\x03\x54\x6F\x70\x02\x00\x05\x57\x69\x64\x74\x68\x03\x60\x01\x05\x41\x6C\x69\x67\x6E\x07\x08\x61\x6C\x43\x6C\x69\x65\x6E\x74\x0A\x42\x65\x76\x65\x6C\x4F\x75\x74\x65\x72\x07\x06\x62\x76\x4E\x6F\x6E\x65\x0C\x43\x6C\x69\x65\x6E\x74\x48\x65\x69\x67\x68\x74\x03\x5B\x01\x0B\x43\x6C\x69\x65\x6E\x74\x57\x69\x64\x74\x68\x03\x60\x01\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x01\x00\x06\x54\x4C\x61\x62\x65\x6C\x06\x4C\x61\x62\x65\x6C\x31\x04\x4C\x65\x66\x74\x02\x20\x06\x48\x65\x69\x67\x68\x74\x02\x19\x03\x54\x6F\x70\x02\x20\x05\x57\x69\x64\x74\x68\x02\x60\x08\x41\x75\x74\x6F\x53\x69\x7A\x65\x08\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x0C\xE8\xBF\x9E\xE6\x8E\xA5\xE5\x90\x8D\xEF\xBC\x9A\x0B\x50\x61\x72\x65\x6E\x74\x43\x6F\x6C\x6F\x72\x08\x00\x00\x05\x54\x45\x64\x69\x74\x0B\x45\x64\x74\x43\x6F\x6E\x6E\x4E\x61\x6D\x65\x04\x4C\x65\x66\x74\x03\x80\x00\x06\x48\x65\x69\x67\x68\x74\x02\x19\x03\x54\x6F\x70\x02\x20\x05\x57\x69\x64\x74\x68\x03\xC0\x00\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x00\x00\x00\x06\x54\x4C\x61\x62\x65\x6C\x06\x4C\x61\x62\x65\x6C\x32\x04\x4C\x65\x66\x74\x02\x20\x06\x48\x65\x69\x67\x68\x74\x02\x19\x03\x54\x6F\x70\x02\x58\x05\x57\x69\x64\x74\x68\x02\x60\x08\x41\x75\x74\x6F\x53\x69\x7A\x65\x08\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x12\xE4\xB8\xBB\xE6\x9C\xBA\xE5\x90\x8D\x2F\x49\x50\xE5\x9C\xB0\xEF\xBC\x9A\x0B\x50\x61\x72\x65\x6E\x74\x43\x6F\x6C\x6F\x72\x08\x00\x00\x05\x54\x45\x64\x69\x74\x07\x45\x64\x74\x48\x6F\x73\x74\x04\x4C\x65\x66\x74\x03\x80\x00\x06\x48\x65\x69\x67\x68\x74\x02\x19\x03\x54\x6F\x70\x02\x58\x05\x57\x69\x64\x74\x68\x03\xC0\x00\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x01\x04\x54\x65\x78\x74\x06\x09\x31\x32\x37\x2E\x30\x2E\x30\x2E\x31\x00\x00\x06\x54\x4C\x61\x62\x65\x6C\x06\x4C\x61\x62\x65\x6C\x33\x04\x4C\x65\x66\x74\x02\x20\x06\x48\x65\x69\x67\x68\x74\x02\x19\x03\x54\x6F\x70\x02\x78\x05\x57\x69\x64\x74\x68\x02\x60\x08\x41\x75\x74\x6F\x53\x69\x7A\x65\x08\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x09\xE7\xAB\xAF\xE5\x8F\xA3\xEF\xBC\x9A\x0B\x50\x61\x72\x65\x6E\x74\x43\x6F\x6C\x6F\x72\x08\x00\x00\x05\x54\x45\x64\x69\x74\x07\x45\x64\x74\x50\x6F\x72\x74\x04\x4C\x65\x66\x74\x03\x80\x00\x06\x48\x65\x69\x67\x68\x74\x02\x19\x03\x54\x6F\x70\x02\x78\x05\x57\x69\x64\x74\x68\x02\x38\x0B\x4E\x75\x6D\x62\x65\x72\x73\x4F\x6E\x6C\x79\x09\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x02\x04\x54\x65\x78\x74\x06\x04\x36\x33\x37\x39\x00\x00\x06\x54\x4C\x61\x62\x65\x6C\x06\x4C\x61\x62\x65\x6C\x35\x04\x4C\x65\x66\x74\x02\x20\x06\x48\x65\x69\x67\x68\x74\x02\x19\x03\x54\x6F\x70\x03\x98\x00\x05\x57\x69\x64\x74\x68\x02\x60\x08\x41\x75\x74\x6F\x53\x69\x7A\x65\x08\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x09\xE5\xAF\x86\xE7\xA0\x81\xEF\xBC\x9A\x0B\x50\x61\x72\x65\x6E\x74\x43\x6F\x6C\x6F\x72\x08\x00\x00\x05\x54\x45\x64\x69\x74\x0B\x45\x64\x74\x50\x61\x73\x73\x77\x6F\x72\x64\x04\x4C\x65\x66\x74\x03\x80\x00\x06\x48\x65\x69\x67\x68\x74\x02\x19\x03\x54\x6F\x70\x03\x98\x00\x05\x57\x69\x64\x74\x68\x03\x90\x00\x08\x45\x63\x68\x6F\x4D\x6F\x64\x65\x07\x0A\x65\x6D\x50\x61\x73\x73\x77\x6F\x72\x64\x0C\x50\x61\x73\x73\x77\x6F\x72\x64\x43\x68\x61\x72\x06\x01\x2A\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x03\x00\x00\x00\x00")

// 注册Form资源
var _ = lcl.RegisterFormResource(NewConnection, &newConnectionBytes)
