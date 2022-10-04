// 在这里写你的事件

package src

import (
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/rtl"
	"github.com/energye/golcl/lcl/types"
)

//::private::
type TForm1Fields struct {
}

func (f *TForm1) OnFormCreate(sender lcl.IObject) {
	f.ScreenCenter()

}

func (f *TForm1) OnFormDestroy(sender lcl.IObject) {

}

func (f *TForm1) OnButton1Click(sender lcl.IObject) {
	// 屏幕截图
	dc := rtl.GetDC(0)         // windows桌面
	defer rtl.ReleaseDC(0, dc) // 一定要释放
	bmp := lcl.NewBitmap()
	defer bmp.Free()
	bmp.LoadFromDevice(dc)
	bmp.Canvas()
	f.Image1.Picture().Assign(bmp)
}

func (f *TForm1) OnPanel1AlignPosition(sender *lcl.TWinControl, control *lcl.TControl, newLeft, newTop, newWidth, newHeight *int32, alignRect *types.TRect, alignInfo types.TAlignInfo) {
	*newLeft = (alignRect.Width() - *newWidth) / 2
	*newTop = (alignRect.Height() - *newHeight) / 2
}

/*
原windows下是这么干的，但lazarus这个跨平台的改变了些东西，直接一个函数就行，反之还只能在Windows下用。
	dc := rtl.GetDC(0)         // windows桌面
	defer rtl.ReleaseDC(0, dc) // 一定要释放

	canvas
	canvas := lcl.NewCanvas()
	defer canvas.Free()

	canvas.SetHandle(dc)

	w := lcl.Screen.Width()
	h := lcl.Screen.Height()

	bmp := lcl.NewBitmap()
	defer bmp.Free()
	bmp.SetSize(w, h)
	bmp.SetPixelFormat(types.Pf32bit)

	bmp.Canvas().CopyRect(types.Rect(0, 0, w, h), canvas, types.Rect(0, 0, w, h))

	f.Image1.Picture().Assign(bmp)
*/
