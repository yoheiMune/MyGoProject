package main

import (

	"os"
	"math"
	"bufio"

	"image/color"
	"math/rand"
	"time"
	"image/gif"
	"image"
	"fmt"
	"io/ioutil"
)

func main() {

	// とりあえず動かして見た.
	{
		fmt.Println("Hello Go-Lang")
	}

	// コマンドライン引数を扱う例.
	{
		var s, sep string
		for i := 1; i < len(os.Args); i++ {
			s += sep + os.Args[i]
			sep = " "
		}
		fmt.Println(s)
	}

	// for文（従来のwhile文）.
	{
		i := 0
		for i <= 10 {
			i++
		}
		fmt.Println(i)
	}

	// 無限ループはこちら.
	{
		i := 0
		for {
			if i > 20 {
				break
			}
			i++
		}
		fmt.Println(i)
	}

	// rangeを使った例.
	{
		s, sep := "", ""
		for _, arg := range os.Args[1:] {
			s += sep + arg
			sep = " "
		}
		fmt.Println(s)
		fmt.Println(os.Args)
		fmt.Println(os.Args[1:])
	}

	// Dup1は標準入力から2回以上現れる行を出現回数と一緒に表示します.
	//{
	//	counts := make(map[string]int)
	//	input := bufio.NewScanner(os.Stdin)
	//	for input.Scan() {
	//		text := input.Text()
	//		if text == "q" {
	//			break
	//		}
	//		counts[text]++
	//	}
	//	// 注意：input.Err()からのエラーの可能性は無視している.
	//	for line, n := range counts {
	//		fmt.Printf("%d\t%s\n", n, line)
	//	}
	//}

	// フォーマット文字列.
	{
		fmt.Printf("%d\n", 12)  // %d：整数（10進数）
		fmt.Printf("%x\n", 12)  // %x：整数（16進数）
		fmt.Printf("%o\n", 12)  // %o：整数（8進数）
		fmt.Printf("%b\n", 12)  // %b：整数（2進数）
		fmt.Printf("%f\n", math.Pi)  // %f：浮動点少数
		fmt.Printf("%e\n", math.Pi)  // %e：浮動点少数（e記法）
		fmt.Printf("%t\n", true)  // %t：ブーリアン
		fmt.Printf("%c\n", "a")  // %c：ルーン →ナニコレ？
		fmt.Printf("%s\n", "abc")  // %s：文字列
		fmt.Printf("%v\n", 12)  // %v：任意の値に対する自然なフォーマット
		fmt.Printf("%T\n", 12)  // %T：任意の値の型
	}

	// ファイルの読み書きと、関数の定義.
	{
		f, err := os.Open("./sample.txt")
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
			os.Exit(1)
		}
		counts := make(map[string]int)
		countLines(f, counts)
		for line, n := range counts {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}

	// ファイルの読み書き.
	{
		data, err := ioutil.ReadFile("./sample.txt")
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			os.Exit(1)
		}
		fmt.Printf("%s\n", data)
	}


	// リサージュ図形
	{
		var palette = []color.Color{color.White, color.Black}

		const (
			whiteIndex = 0 // パレットの最初の色
			blackIndex = 1 // パレットの次の色
		)

		rand.Seed(time.Now().UTC().UnixNano())

		const (
			cycles  = 5     // 発振器xが完了する周回の回数
			res     = 0.001 // 回転の分解能
			size    = 100   // 画像キャンパスは [-size..+size] の範囲を扱う
			nframes = 64    // アニメーションフレーム数
			delay   = 8     // 10ms 単位でのフレーム間の遅延
		)

		freq := rand.Float64() * 3.0 // 発振器 y の相対周波数
		anim := gif.GIF{LoopCount: nframes}
		phase := 0.0 // 位相差
		for i := 0; i < nframes; i++ {
			rect := image.Rect(0, 0, 2*size+1, 2*size+1)
			img := image.NewPaletted(rect, palette)
			for t := 0.0; t < cycles*2*math.Pi; t += res {
				x := math.Sin(t)
				y := math.Sin(t*freq + phase)
				img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), blackIndex)
			}
			phase += 0.1
			anim.Delay = append(anim.Delay, delay)
			anim.Image = append(anim.Image, img)
		}
		//gif.EncodeAll(os.Stdout, &anim)
	}






























}

func countLines(f *os.File, counts map[string]int) {

	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
	// 注意：input.Err()からのエラーの可能性は無視している.

}
