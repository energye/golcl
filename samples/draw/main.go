package main

import (
	"fmt"
	"github.com/energye/golcl/energy/inits"
	"github.com/energye/golcl/lcl/win"

	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/rtl"
	"github.com/energye/golcl/lcl/types"
	"github.com/energye/golcl/lcl/types/colors"
	_ "github.com/energye/golcl/pkgs/winappres"
)

type TPoint struct {
	X, Y int32
	Down bool
}

var (
	isMouseDown bool
	points      = make([]TPoint, 0)
)

func main() {
	inits.Init(nil, nil)
	lcl.Application.Initialize()
	lcl.Application.SetMainFormOnTaskBar(true)

	jpgFileName := "./imgs/1.jpg"
	canLoad := rtl.FileExists(jpgFileName)
	var jpgimg *lcl.TJPEGImage
	if canLoad {
		jpgimg = lcl.NewJPEGImage()
		jpgimg.LoadFromFile(jpgFileName)
	}

	mainForm := lcl.Application.CreateForm()
	mainForm.SetCaption("Hello")
	mainForm.SetPosition(types.PoScreenCenter)
	//mainForm.EnabledMaximize(false)
	mainForm.SetWidth(600)
	mainForm.SetHeight(600)
	mainForm.SetDoubleBuffered(true)
	mainForm.SetOnDestroy(func(sender lcl.IObject) {
		if jpgimg != nil {
			jpgimg.Free()
		}
	})
	mainForm.SetBorderStyle(types.BsNone)
	mainForm.SetColor(colors.ClNavy)
	Exstyle := win.GetWindowLong(mainForm.Handle(), win.GWL_EXSTYLE)
	Exstyle = Exstyle | win.WS_EX_LAYERED&^win.WS_EX_TRANSPARENT
	win.SetWindowLong(mainForm.Handle(), win.GWL_EXSTYLE, uintptr(Exstyle))
	win.SetLayeredWindowAttributes(mainForm.Handle(), //指定分层窗口句柄
		colors.ClNavy,                  //crKey指定需要透明的背景颜色值，可用RGB()宏  0-255
		255,                            //bAlpha设置透明度，0表示完全透明，255表示不透明
		win.LWA_ALPHA|win.LWA_COLORKEY) //LWA_ALPHA: crKey无效，bAlpha有效；
	//LWA_COLORKEY：窗体中的所有颜色为crKey的地方全透明，bAlpha无效。
	//LWA_ALPHA | LWA_COLORKEY：crKey的地方全透明，其它地方根据bAlpha确定透明度

	mainForm.SetOnPaint(func(lcl.IObject) {

		canvas := mainForm.Canvas()

		canvas.MoveTo(10, 10)
		canvas.LineTo(50, 10)
		s := "这是一段文字"
		canvas.Font().SetColor(colors.ClRed) // red
		canvas.Font().SetSize(20)
		style := canvas.Font().Style()
		canvas.Brush().SetStyle(types.BsClear)
		canvas.Font().SetStyle(style.Include(types.FsBold, types.FsItalic))
		canvas.TextOut(100, 30, s)
		fmt.Println("canvas.Font()", canvas.Font().Height(), canvas.Font().Size())

		r := types.TRect{0, 0, 80, 80}

		// 计算文字
		//fmt.Println("TfSingleLine: ", types.TfSingleLine)
		s = "由于现有第三方的Go UI库不是太庞大就是用的不习惯，或者组件太少。"
		canvas.TextRect2(&r, s, types.NewSet(types.TfCenter, types.TfVerticalCenter, types.TfSingleLine))
		//fmt.Println("r: ", r, ", s: ", s)

		s = "测试输出"
		r = types.TRect{0, 0, 80, 80}
		// brush
		canvas.Brush().SetColor(colors.ClGreen)
		canvas.FillRect(r)

		// font
		canvas.Font().SetStyle(0)
		canvas.Font().SetSize(9)
		canvas.Font().SetColor(colors.ClBlue)

		// pen
		canvas.Pen().SetWidth(2)
		canvas.Pen().SetColor(colors.ClFuchsia)
		canvas.Rectangle(r.Left, r.Top, r.Right, r.Bottom)

		textFmt := types.NewSet(types.TfCenter, types.TfSingleLine, types.TfVerticalCenter)
		//fmt.Println("format: ", textFmt)
		//		canvas.TextRect(r, 0, 0, s)
		canvas.TextRect2(&r, s, textFmt)

		canvas.Draw(0, 80, jpgimg)
		//canvas.Draw2(0, 200, jpgimg, 10)

		// 画多边形

		canvas.Brush().SetColor(colors.ClYellow)
		canvas.Polygon([]types.TPoint{{15, 40}, {43, 123}, {81, 42}, {45, 11}})

		canvas.Polyline([]types.TPoint{{15 + 100, 40}, {43 + 100, 123}, {81 + 100, 42}, {45 + 100, 11}})

	})

	paintbox := lcl.NewPaintBox(mainForm)
	paintbox.SetParent(mainForm)
	paintbox.SetAlign(types.AlBottom)
	paintbox.SetHeight(mainForm.Height() - 280)
	//paintbox.SetColor(colors.ClRed)
	paintbox.SetOnPaint(func(lcl.IObject) {
		canvas := paintbox.Canvas()
		canvas.Pen().SetColor(colors.ClRed)
		canvas.Pen().SetWidth(10)
		r := paintbox.ClientRect()
		canvas.Rectangle(r.Left, r.Top, r.Right, r.Bottom)

		canvas.Font().SetColor(colors.ClSkyblue)
		rect := paintbox.ClientRect()
		s := "在这可以用鼠标绘制"
		textFmt := types.NewSet(types.TfCenter, types.TfSingleLine, types.TfVerticalCenter)
		canvas.TextRect2(&rect, s, textFmt)

		canvas.Pen().SetColor(colors.ClGreen)
		for _, p := range points {
			if p.Down {
				canvas.MoveTo(p.X, p.Y)
			} else {
				canvas.LineTo(p.X, p.Y)
			}
		}

	})

	paintbox.SetOnMouseDown(func(sender lcl.IObject, button types.TMouseButton, shift types.TShiftState, x, y int32) {
		fmt.Println("mouse down")
		if button == types.MbLeft {
			points = append(points, TPoint{X: x, Y: y, Down: true})
			isMouseDown = true
		}
	})

	paintbox.SetOnMouseMove(func(sender lcl.IObject, shift types.TShiftState, x, y int32) {
		if isMouseDown {
			points = append(points, TPoint{X: x, Y: y, Down: false})
			paintbox.Repaint()
		}
	})

	paintbox.SetOnMouseUp(func(sender lcl.IObject, button types.TMouseButton, shift types.TShiftState, x, y int32) {
		fmt.Println("mouse SetOnMouseUp")
		if button == types.MbLeft {
			isMouseDown = false
		}
	})

	btnClear := lcl.NewButton(mainForm)
	btnClear.SetParent(mainForm)
	btnClear.SetCaption("清除绘制")
	btnClear.SetLeft(mainForm.Width() - btnClear.Width() - 20)
	btnClear.SetTop(10)
	btnClear.SetOnClick(func(lcl.IObject) {
		points = make([]TPoint, 0)
		paintbox.Repaint()
	})

	lcl.Application.Run()
}
