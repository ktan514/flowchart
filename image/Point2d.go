package image

// 座標クラス
type Point2d struct {
	X int
	Y int
}

// 座標クラスコンストラクタ
func (pt *Point2d) Initialize(x, y int) {
	pt.X = x
	pt.Y = y
}
