package main

import (
	"fmt"

	"github.com/energye/golcl/lcl"
	_ "github.com/energye/golcl/pkgs/winappres"
)

func main() {

	lcl.Application.Initialize()
	lcl.Application.SetMainFormOnTaskBar(true)

	lcl.Application.CreateForm(&FormListViewDraw)

	// 加载信息
	trainData, err := parseFromFile("testtraindata.json")
	if err == nil {
		// 设计器设计的ListView
		fullResFormListViewInstance(trainData)
		// 代码创建的ListView
		codeBuildistViewInstance(trainData)
	} else {
		fmt.Println("err:", err)
	}

	lcl.Application.Run()
}
