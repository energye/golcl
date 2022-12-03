//go:build darwin
// +build darwin

package main

import (
	"github.com/energye/golcl/lcl"
	_ "github.com/energye/golcl/pkgs/winappres"
)

func main() {
	lcl.RunApp(&Form1)
}
