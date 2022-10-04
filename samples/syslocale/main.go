package main

import (
	"fmt"

	"github.com/energye/golcl/lcl/rtl"
	_ "github.com/energye/golcl/pkgs/winappres"
)

func main() {
	fmt.Println(rtl.SysLocale)
}
