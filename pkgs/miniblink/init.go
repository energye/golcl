//go:build windows
// +build windows

package miniblink

import "github.com/energye/golcl/lcl/rtl"

var (
	isLcl bool
)

func init() {
	isLcl = rtl.LcLLoaded()
	//Initialize()
}
