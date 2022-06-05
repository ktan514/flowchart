package main

import (
	"bufio"
	"fmt"
	"image/color"
	"image/png"
	"os"
	"regexp"
	"strings"

	"github.com/ktan514/flowchart/image"
)

func IsMethodDefine(line string) bool {
	return true
}

func OpenFile(fileName string) []string {
	fmt.Println("ファイル読み取り処理を開始します")

	// ファイルオープン
	fp, err := os.Open(fileName)
	if err != nil {
		// エラー処理
	}
	defer fp.Close()

	file := []string{}

	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		// ここで一行ずつ処理
		fmt.Println(scanner.Text())
		file = append(file, scanner.Text())
	}

	if err = scanner.Err(); err != nil {
		// エラー処理
	}

	return file
}

func AnalyzerIsFunction(text string) bool {
	return strings.Contains(text, "FUNC")
}

func AnalyzerIsComment(text string) bool {
	match, _ := regexp.MatchString(`^\s*#`, text)
	return match
}

func AnalyzerIsEmptyLine(text string) bool {
	match, _ := regexp.MatchString(`^\s*\r\n`, text)
	return match
}

func AnalyzerCheckIndent(text string) int {
	count := 0
	for _, c := range text {
		if c == '\t' {
			count++
			continue
		}
		break
	}
	return count
}

func Analyzer(text []string, i int, beforeTab int) {

	// ReservedWordList := [...]string{
	// 	"FUNC",
	// 	"WHILE",
	// 	"FOR",
	// 	"IF",
	// 	"ELIF",
	// 	"ELSE",
	// 	"SWITCH",
	// 	"CASE",
	// }

	// FUNC句
	if AnalyzerIsFunction(text[i]) {
		Analyzer(text, i+1, beforeTab+1)
	}

	for j := i; j < len(text); j++ {
		tab := AnalyzerCheckIndent(text[j])

		// タブ数(インデント)が同じなら同じブロック内。
		if beforeTab == tab {
			// コメント行
			if AnalyzerIsComment(text[j]) {
				continue
			}

			// 空白行
			if AnalyzerIsEmptyLine(text[j]) {
				continue
			}

			if strings.Index(text[j], "\t") == 0 {

				fmt.Println(text[j], "       ここは関数ブロック内。")
			}
		}

	}
}

func main() {

	text := OpenFile("./source.txt")
	Analyzer(text, 0, 0)

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
