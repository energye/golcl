//----------------------------------------
//
// Copyright © ying32. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package lcl

import (
	. "github.com/energye/golcl/lcl/api"
	. "github.com/energye/golcl/lcl/types"
)

// BrushCopy 画刷复制
func (c *TCanvas) BrushCopy(dest TRect, bitmap IObject, source TRect, color TColor) {
	Canvas_BrushCopy(c._instance(), dest, CheckPtr(bitmap), source, color)
}

// CopyRect 制矩形
func (c *TCanvas) CopyRect(dest TRect, canvas IObject, source TRect) {
	Canvas_CopyRect(c._instance(), dest, CheckPtr(canvas), source)
}

// Draw 绘制graphic对象
func (c *TCanvas) Draw(x, y int32, graphic IGraphic) {
	Canvas_Draw1(c._instance(), x, y, CheckPtr(graphic))
}

// DrawFocusRect 画焦点矩形
func (c *TCanvas) DrawFocusRect(aRect TRect) {
	Canvas_DrawFocusRect(c._instance(), aRect)
}

// FillRect 填充矩形
func (c *TCanvas) FillRect(aRect TRect) {
	Canvas_FillRect(c._instance(), aRect)
}

// FrameRect 绘制边框
func (c *TCanvas) FrameRect(aRect TRect) {
	Canvas_FrameRect(c._instance(), aRect)
}

// TextRect 在矩形内绘制文字
func (c *TCanvas) TextRect(aRect TRect, x, y int32, text string) {
	Canvas_TextRect1(c._instance(), aRect, x, y, text)
}

// TextRect2 在矩形内绘制文字
func (c *TCanvas) TextRect2(aRect *TRect, text string, textFormat TTextFormat) {
	Canvas_TextRect2(c._instance(), aRect, text, textFormat)
}

// Polygon 填充多边形
func (c *TCanvas) Polygon(points []TPoint) {
	Canvas_Polygon(c._instance(), points)
}

// Polyline 画多边形，不填充
func (c *TCanvas) Polyline(points []TPoint) {
	Canvas_Polyline(c._instance(), points)
}

// PolyBezier 多边形贝塞尔曲线
func (c *TCanvas) PolyBezier(points []TPoint) {
	Canvas_PolyBezier(c._instance(), points)
}
