// 由res2go自动生成，不要编辑。
package main

import (
	"github.com/energye/golcl/lcl"
)

type TAbout struct {
	*lcl.TForm

	//::private::
	TAboutFields
}

var About *TAbout

// 以字节形式加载
// lcl.Application.CreateForm(&About)

func NewAbout(owner lcl.IComponent) (root *TAbout) {
	lcl.CreateResForm(owner, &root)
	return
}

var aboutBytes = []byte("\x54\x50\x46\x30\x06\x54\x41\x62\x6F\x75\x74\x05\x41\x62\x6F\x75\x74\x04\x4C\x65\x66\x74\x03\xE9\x02\x06\x48\x65\x69\x67\x68\x74\x03\x2C\x01\x03\x54\x6F\x70\x03\x98\x01\x05\x57\x69\x64\x74\x68\x03\x90\x01\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x05\x41\x62\x6F\x75\x74\x0D\x44\x65\x73\x69\x67\x6E\x54\x69\x6D\x65\x50\x50\x49\x02\x78\x08\x4F\x6E\x43\x72\x65\x61\x74\x65\x07\x0A\x46\x6F\x72\x6D\x43\x72\x65\x61\x74\x65\x0A\x4C\x43\x4C\x56\x65\x72\x73\x69\x6F\x6E\x06\x07\x31\x2E\x38\x2E\x30\x2E\x36\x00\x00")

// 注册Form资源
var _ = lcl.RegisterFormResource(About, &aboutBytes)