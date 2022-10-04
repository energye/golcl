package main

import (
	"fmt"
	"github.com/energye/golcl/inits"
	"io/ioutil"
	"net/http"

	"github.com/energye/golcl/lcl/rtl"
	//	"syscall"
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/types"
	_ "github.com/energye/golcl/pkgs/winappres"
)

func main() {
	inits.Init(nil, nil)
	fmt.Println("main:currentThreadId:", GetCurrentThreadId())

	lcl.Application.Initialize()
	lcl.Application.SetMainFormOnTaskBar(true)

	mainForm := lcl.Application.CreateForm()
	mainForm.SetCaption("Hello")
	mainForm.SetPosition(types.PoScreenCenter)
	mainForm.EnabledMaximize(false)
	mainForm.SetWidth(500)
	mainForm.SetHeight(600)

	img := lcl.NewImage(mainForm)
	img.SetParent(mainForm)
	// 本地加载
	jpgFileName := "./1.jpg"
	if rtl.FileExists(jpgFileName) {
		//mem := lcl.NewMemoryStream()
		//mem.LoadFromFile(jpgFileName)
		//mem.SetPosition(0)
		//img.Picture().LoadFromStream(mem)
		img.Picture().LoadFromFile(jpgFileName)
		//mem.Free()
	}

	// 网络图片加载
	img2 := lcl.NewImage(mainForm)
	img2.SetParent(mainForm)
	img2.SetTop(img.Height() + 10)
	img2.SetAutoSize(true)

	// 异步加载，一般来说不要在非主线程中访问UI组件,需要在线程中访问ui组件请使用 lcl.ThreadSync
	go func() {
		fmt.Println("main:currentThreadId2:", GetCurrentThreadId())
		resp, err := http.Get("http://ww2.sinaimg.cn/large/df780e95jw1egxm06uxerj20cs05hjs8.jpg")
		if err == nil {
			defer resp.Body.Close()
			bs, err := ioutil.ReadAll(resp.Body)
			if err == nil {
				//mem := lcl.NewMemoryStream()
				//defer mem.Free()
				//mem.Write(bs)
				//mem.SetPosition(0)
				// 让以下代码运行在主线程中
				lcl.ThreadSync(func() {
					fmt.Println("main:currentThreadId3:", GetCurrentThreadId())
					//img2.Picture().LoadFromStream(mem)
					img2.Picture().LoadFromBytes(bs)
				})
				fmt.Println("测试运行到此。")
			} else {
				fmt.Println(err)
			}
		} else {
			fmt.Println(err)
		}
	}()
	lcl.Application.Run()
}
