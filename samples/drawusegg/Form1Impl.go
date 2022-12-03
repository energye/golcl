// 在这里写你的事件

package main

import (
	"fmt"
	"runtime"

	"github.com/energye/golcl/lcl/bitmap"

	"github.com/fogleman/gg"

	"github.com/energye/golcl/lcl"
)

//::private::
type TForm1Fields struct {
	drawFunds map[string]func()
	bufBmp    *lcl.TBitmap
}

// 使用gg 2d图形库来绘制，最终显示到govcl控件上。
// gg部分的代码来自gg的例子
// https://github.com/fogleman/gg

func (f *TForm1) OnFormCreate(sender lcl.IObject) {
	f.ScreenCenter()
	f.bufBmp = lcl.NewBitmap()
	f.drawFunds = map[string]func(){
		"ggDrawLines":          f.ggDrawLines,
		"ggDrawEllipse":        f.ggDrawEllipse,
		"ggDrawGradientLinear": f.ggDrawGradientLinear,
		"ggDrawGradientRadial": f.ggDrawGradientRadial,
		"ggDrawBeziers":        f.ggDrawBeziers,
		"ggDrawCircle":         f.ggDrawCircle,
		"ggDrawCrisp":          f.ggDrawCrisp,
		"ggDrawCubic":          f.ggDrawCubic,
		"ggDrawFont":           f.ggDrawFont,
		"ggDrawGradientText":   f.ggDrawGradientText,
		"ggDrawInvertMask":     f.ggDrawInvertMask,
		"ggDrawLineWidth":      f.ggDrawLineWidth,
		"ggDrawLorem":          f.ggDrawLorem,
		"ggDrawRotatedText":    f.ggDrawRotatedText,
		"ggDrawRotatedImage":   f.ggDrawRotatedImage,
		"ggDrawWarpText":       f.ggDrawWarpText,
		"ggDrawStar":           f.ggDrawStar,
		"ggDrawStars":          f.ggDrawStars,
	}

}

func (f *TForm1) OnFormDestroy(sender lcl.IObject) {
	if f.bufBmp != nil && f.bufBmp.IsValid() {
		f.bufBmp.Free()
	}
}

func (f *TForm1) ggImageToBitmap(dc *gg.Context) error {
	err := bitmap.ToBitmap2(dc.Image(), f.bufBmp)
	if err != nil {
		return err
	}
	return nil
}

func (f *TForm1) ggImageToPng(dc *gg.Context) *lcl.TPngImage {
	obj, err := bitmap.ToPngImage(dc.Image())
	if err != nil {
		return nil
	}
	return obj
}

func (f *TForm1) ggDrawImage(dc *gg.Context) {
	// linux跟macOS下有些效果不太好，所以用png，具体原因还待分析。
	if runtime.GOOS != "windows" {
		pngObj := f.ggImageToPng(dc)
		if pngObj != nil {
			defer pngObj.Free()
			f.Canvas().Draw(0, 0, pngObj)
		}
	} else {
		if f.ggImageToBitmap(dc) == nil {
			f.Canvas().Draw(0, 0, f.bufBmp)
		}
	}

}

func (f *TForm1) OnFormPaint(sender lcl.IObject) {
	if f.ListBox1.ItemIndex() == -1 {
		return
	}
	fn, ok := f.drawFunds[f.ListBox1.Items().Strings(f.ListBox1.ItemIndex())]
	if !ok {
		fmt.Println("没有找到指定的滤镜。")
		return
	}
	fn()
}

func (f *TForm1) OnListBox1Click(sender lcl.IObject) {
	f.Repaint()
}
