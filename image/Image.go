package image

import (
	"image"
	"image/color"
)

// イメージ生成
func NewImage(w, h int) *image.RGBA {
	m := image.NewRGBA(image.Rect(0, 0, w, h))
	white := color.RGBA{
		255,
		255,
		255,
		255,
	}
	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			m.Set(x, y, white)
		}
	}
	return m
}

func AbsInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// 線分を描く(ブレゼンハムのアルゴリズム)
func DrawLine(img *image.RGBA, x1 int, y1 int,
	x2 int, y2 int, lineColor color.RGBA) {
	var step int = 0
	var dx int = AbsInt(x2 - x1)
	var dy int = AbsInt(y2 - y1)

	if dx > dy {
		if x1 > x2 {
			step = 0
			if y1 > y2 {
				step = 1
			} else {
				step = -1
			}
			x1, x2 = x2, x1 // swap
			y1 = y2
		} else {
			if y1 < y2 {
				step = 1
			} else {
				step = -1
			}
		}
		img.Set(x1, y1, lineColor)
		var s int = dx >> 1
		x1++
		for x1 <= x2 {
			s -= dy
			if s < 0 {
				s += dx
				y1 += step
			}
			img.Set(x1, y1, lineColor)
			x1++
		}
	} else {
		if y1 > y2 {
			if x1 > x2 {
				step = 1
			} else {
				step = -1
			}
			y1, y2 = y2, y1 // swap
			x1 = x2
		} else {
			if x1 < x2 {
				step = 1
			} else {
				step = -1
			}
		}
		img.Set(x1, y1, lineColor)
		var s int = dy >> 1
		y1++
		for y1 <= y2 {
			s -= dx
			if s < 0 {
				s += dy
				x1 += step
			}
			img.Set(x1, y1, lineColor)
			y1++
		}
	}
}
