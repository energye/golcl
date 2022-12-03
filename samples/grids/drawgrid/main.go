package main

import (
	"fmt"
	"math/rand"

	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/types"
	_ "github.com/energye/golcl/pkgs/winappres"
)

func main() {
	lcl.Application.Initialize()

	mainForm := lcl.Application.CreateForm()
	mainForm.SetWidth(700)
	mainForm.SetHeight(500)
	mainForm.WorkAreaCenter()
	mainForm.SetCaption("表格自绘")
	mainForm.ScaleSelf()
	grid := NewPlayControl(mainForm)
	grid.SetParent(mainForm)
	grid.SetAlign(types.AlClient)
	for i := 1; i <= 100; i++ {
		grid.Add(TPlayListItem{fmt.Sprintf("标题%d", i), "张三", 100000 + rand.Int31n(100000), "", ""})

	}
	lcl.Application.Run()
}
