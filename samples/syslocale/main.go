package main

import (
	"fmt"
	"github.com/energye/golcl/energy/inits"

	"github.com/energye/golcl/lcl/rtl"
	_ "github.com/energye/golcl/pkgs/winappres"
)

func main() {
	inits.Init(nil, nil)
	fmt.Println(rtl.SysLocale)
}
