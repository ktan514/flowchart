package image

import (
	"image"
	"image/color"
	"math"
)

// 円群クラス
type Ellipse struct {
	left      Point2d
	right     Point2d
	length    int
	radius    int
	lineColor color.RGBA
}

// 座標クラスコンストラクタ
func (el *Ellipse) Initialize(x1, y1, x2, y2, r, l int, color color.RGBA) {
	el.left.X = x1
	el.left.Y = y1
	el.right.X = x2
	el.right.Y = y2
	el.radius = r
	el.length = l
	el.lineColor = color
}

// 円描画
func (el *Ellipse) DrawArc(img *image.RGBA) {
	var bx, by int
	var firstX, firstY int
	const NUM = 600

	for i := 0; i < NUM/2; i++ {
		angle := float64(i)*math.Pi*2.0/float64(NUM) + (math.Pi / 2.0)
		x := int(float64(el.radius)*math.Cos(angle) + float64(el.left.X))
		y := int(float64(el.radius)*math.Sin(angle) + float64(el.left.Y))
		if i == 0 {
			firstX, firstY = x, y
		} else {
			// 円弧(線分)を描く
			DrawLine(img, x, y, bx, by, el.lineColor)
		}
		bx, by = x, y
	}
	{
		// 円弧(線分)を描く
		x := int(float64(el.radius)*math.Cos(math.Pi/2.0) + float64(el.right.X))
		y := int(float64(el.radius)*math.Sin(math.Pi/2.0) + float64(el.right.Y))
		DrawLine(img, firstX, firstY, x, y, el.lineColor)
	}

	for i := 0; i < NUM/2; i++ {
		angle := float64(i)*math.Pi*2.0/float64(NUM) - (math.Pi / 2.0)
		x := int(float64(el.radius)*math.Cos(angle) + float64(el.right.X))
		y := int(float64(el.radius)*math.Sin(angle) + float64(el.right.Y))
		if i == 0 {
			firstX, firstY = x, y
		} else {
			// 円弧(線分)を描く
			DrawLine(img, x, y, bx, by, el.lineColor)
		}
		bx, by = x, y
	}
	{
		// 円弧(線分)を描く
		x := int(float64(el.radius)*math.Cos(-math.Pi/2.0) + float64(el.left.X))
		y := int(float64(el.radius)*math.Sin(-math.Pi/2.0) + float64(el.left.Y))
		DrawLine(img, firstX, firstY, x, y, el.lineColor)
	}
}
