// 由res2go自动生成，不要编辑。
package main

import (
	"github.com/energye/golcl/lcl"
)

type TFrameRedis struct {
	*lcl.TFrame

	//::private::
	TFrameRedisFields
}

func NewFrameRedis(owner lcl.IComponent) (root *TFrameRedis) {
	lcl.CreateResFrame(owner, &root)
	return
}

var frameRedisBytes = []byte("\x54\x50\x46\x30\x0B\x54\x46\x72\x61\x6D\x65\x52\x65\x64\x69\x73\x0A\x46\x72\x61\x6D\x65\x52\x65\x64\x69\x73\x04\x4C\x65\x66\x74\x02\x00\x06\x48\x65\x69\x67\x68\x74\x03\xF0\x00\x03\x54\x6F\x70\x02\x00\x05\x57\x69\x64\x74\x68\x03\x40\x01\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x00\x0A\x44\x65\x73\x69\x67\x6E\x4C\x65\x66\x74\x03\x9B\x01\x09\x44\x65\x73\x69\x67\x6E\x54\x6F\x70\x02\x1C\x00\x00")

// 注册Form资源
var _ = lcl.RegisterFormResource(TFrameRedis{}, &frameRedisBytes)
