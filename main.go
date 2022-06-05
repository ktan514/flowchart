package main

import (
	"fmt"
	"image/color"
	"image/png"
	"os"

	"github.com/ktan514/flowchart/image"
)

const DIV_NUM = 2

func main() {
	// 保存するファイル名
	imgFile, err := os.Create("line_img.png")
	if err != nil {
		fmt.Println("画像ファイルが作成できませんでした。")
		os.Exit(1)
	}
	defer imgFile.Close()

	img := image.NewImage(1800, 1200)

	lineColor := color.RGBA{
		0,
		255,
		255,
		255,
	}

	var radius = 36

	// 円群クラス
	el := &image.Ellipse{}
	el.Initialize(100, 100, 300, 100, radius, 200, lineColor)

	// 円群を描画
	el.DrawArc(img)

	// PNG形式で保存する
	png.Encode(imgFile, img)
}
