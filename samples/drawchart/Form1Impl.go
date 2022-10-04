// 在这里写你的事件

package main

import (
	"fmt"
	"image"
	"runtime"

	"github.com/energye/golcl/lcl/bitmap"

	"github.com/energye/golcl/lcl"
)

//::private::
type TForm1Fields struct {
	drawFuncs map[string]func()
	buffBmp   *lcl.TBitmap
}

// 使用chart来生成图表
// chart部分的代码来自chart的例子
// https://github.com/vdobler/chart

func (f *TForm1) OnFormCreate(sender lcl.IObject) {
	f.ScreenCenter()
	f.buffBmp = lcl.NewBitmap()
	f.drawFuncs = map[string]func(){
		"stripChart":          f.stripChart,
		"keyStyles":           f.keyStyles,
		"scatterTics":         f.scatterTics,
		"scatterChart":        f.scatterChart,
		"functionPlots":       f.functionPlots,
		"autoscale":           f.autoscale,
		"boxChart":            f.boxChart,
		"kernels":             f.kernels,
		"histCharthistFalse":  f.histCharthistFalse,
		"histCharthistTrue":   f.histCharthistTrue,
		"histChartshistFalse": f.histChartshistFalse,
		"histChartshistTrue":  f.histChartshistTrue,
		"histChartohistFalse": f.histChartohistFalse,
		"histChartohistTrue":  f.histChartohistTrue,
		"barChart":            f.barChart,
		"categoricalBarChart": f.categoricalBarChart,
		"logAxis":             f.logAxis,
		"pieChart":            f.pieChart,
		"testGraphics":        f.testGraphics,
		"bestOf":              f.bestOf,
		"mietenChart":         f.mietenChart,
	}
}

func (f *TForm1) OnFormDestroy(sender lcl.IObject) {
	if f.buffBmp != nil && f.buffBmp.IsValid() {
		f.buffBmp.Free()
	}
}

type TDrawFunc = func(img image.Image)

func (f *TForm1) drawImage(img image.Image) {
	// linux跟macOS下有些效果不太好，所以用png，具体原因还待分析。
	if runtime.GOOS != "windows" {
		pngObj, err := bitmap.ToPngImage(img)
		if err == nil && pngObj != nil {
			defer pngObj.Free()
			f.Canvas().Draw(0, 0, pngObj)
		}
	} else {
		if f.buffBmp != nil {
			if bitmap.ToBitmap2(img, f.buffBmp) == nil {
				f.Canvas().Draw(0, 0, f.buffBmp)
			}
		}
	}

}

func (f *TForm1) OnFormPaint(sender lcl.IObject) {
	if f.ListBox1.ItemIndex() == -1 {
		return
	}

	fn, ok := f.drawFuncs[f.ListBox1.Items().Strings(f.ListBox1.ItemIndex())]
	if !ok {
		fmt.Println("没有找到指定的方法。")
		return
	}
	fn()
}

func (f *TForm1) OnListBox1Click(sender lcl.IObject) {
	f.Repaint()
}
