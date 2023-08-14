package main

import (
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/types"
)

//
type TImageTrackBar struct {
	*lcl.TPanel
	thumb      *lcl.TImage
	background *lcl.TImage
	fore       *lcl.TImage
	//max        int
	//min        int
	mouseDown bool

	position float32 // 最大100，百分比

	OnTrackChange lcl.TNotifyEvent
}

func NewImageTrackBar(owner lcl.IComponent) *TImageTrackBar {
	t := new(TImageTrackBar)
	t.TPanel = lcl.NewPanel(owner)
	t.TPanel.SetBevelOuter(types.BvNone)
	t.TPanel.SetHeight(8)
	t.TPanel.SetParentColor(true)
	t.TPanel.SetDoubleBuffered(true)

	t.background = lcl.NewImage(owner)
	t.fore = lcl.NewImage(owner)
	t.thumb = lcl.NewImage(owner)

	t.background.SetParent(t.TPanel)
	t.fore.SetParent(t.TPanel)
	t.thumb.SetParent(t.TPanel)

	t.background.SetAlign(types.AlClient)
	t.background.SetStretch(true)

	t.fore.SetWidth(100)
	t.fore.SetStretch(true)

	//t.thumb.SetAutoSize(true)
	t.thumb.SetCursor(types.CrHandPoint)

	//fmt.Println(t.thumb.)

	t.thumb.SetOnMouseDown(t.onMouseDown)
	t.thumb.SetOnMouseUp(t.onMouseUp)
	t.thumb.SetOnMouseMove(t.onMouseMove)
	return t
}

func (t *TImageTrackBar) SetImages(thumb, background, fore *lcl.TImage) {
	t.thumb.Picture().Assign(thumb.Picture())
	t.background.Picture().Assign(background.Picture())
	t.fore.Picture().Assign(fore.Picture())
	t.thumb.SetWidth(thumb.Picture().Width())
	t.thumb.SetHeight(thumb.Picture().Height())
}

func (t *TImageTrackBar) Free() {
	t.thumb.Free()
	t.background.Free()
	t.fore.Free()
	t.TPanel.Free()
}

func (t *TImageTrackBar) adjust() {
	r := t.ClientRect()
	//t.background.SetBounds(0, 0, r.Right, r.Bottom)
	posi := int32(t.position / 100.0 * float32(r.Right))
	t.fore.SetBounds(0, 0, posi, r.Bottom)
	thumbPosi := posi
	w := t.thumb.Width()
	if posi <= 0 {
		thumbPosi = 0
	} else if posi+w >= r.Right {
		thumbPosi = r.Right - w
	}
	t.thumb.SetLeft(thumbPosi)
}

func (t *TImageTrackBar) SetPosition(value int) {
	t.setPosition(float32(value))
}

func (t *TImageTrackBar) Position() int {
	return int(t.position)
}

func (t *TImageTrackBar) setPosition(value float32) {
	t.position = value
	if t.position < 0 {
		t.position = 0
	}
	if t.position > 100 {
		t.position = 100
	}
	t.adjust()
	if t.OnTrackChange != nil {
		t.OnTrackChange(t)
	}
}

func (t *TImageTrackBar) onMouseDown(sender lcl.IObject, button types.TMouseButton, shift types.TShiftState, x, y int32) {
	if button == types.MbLeft {
		t.mouseDown = true
	}
}

func (t *TImageTrackBar) onMouseUp(sender lcl.IObject, button types.TMouseButton, shift types.TShiftState, x, y int32) {
	t.mouseDown = false
}

func (t *TImageTrackBar) onMouseMove(sender lcl.IObject, shift types.TShiftState, x, y int32) {
	if t.mouseDown {
		pt := t.thumb.ClientToParent(types.TPoint{x, y}, t)
		val := (float32(pt.X) / float32(t.Width())) * 100
		t.setPosition(val)
	}
}
