package image

import (
	"image"
	"image/color"
	"math"
)

// 円群クラス
type Circle struct {
	center    Point2d
	radius    int
	lineColor color.RGBA
}

// 座標クラスコンストラクタ
func (ci *Circle) Initialize(x, y, r int, color color.RGBA) {
	ci.center.X = x
	ci.center.Y = y
	ci.radius = r
	ci.lineColor = color
}

// 円描画
func (ci *Circle) DrawArc(img *image.RGBA) {
	var bx, by int
	var firstX, firstY int
	const NUM = 400

	for i := 0; i < NUM; i++ {
		angle := float64(i) * math.Pi * 2.0 / float64(NUM)
		x := int(float64(ci.radius)*math.Cos(angle) + float64(ci.center.X))
		y := int(float64(ci.radius)*math.Sin(angle) + float64(ci.center.Y))
		if i == 0 {
			firstX, firstY = x, y
		} else {
			// 円弧(線分)を描く
			DrawLine(img, x, y, bx, by, ci.lineColor)
		}
		bx, by = x, y
	}

	// 円弧(線分)を描く
	DrawLine(img, firstX, firstY, bx, by, ci.lineColor)
}
