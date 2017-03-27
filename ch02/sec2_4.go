package main

import "fmt"

func main() {

	// 2.4.1 タプル代入
	{
		var a, b, c int
		a, b, c = 1, 2, 3
		println(a, b, c)
	}

	// 2.5 型宣言
	{
		type Celsius float64
		type Fahrenheit float64
		const (
			AbsoluteZeroC Celsius = -273.15
		)
		println(AbsoluteZeroC)

		var c Celsius
		var f Fahrenheit
		println(c == 0)
		println(f >= 0)
		//print(c == f) // コンパイルエラー（型が違う）
		println(c == Celsius((f)))

		// 型に振る舞いを追加
		var v MyFloat = 1.1
		println(v.String())
	}
}

type MyFloat float64
func (m MyFloat) String() string {
	return fmt.Sprintf("Myfloat_%g", m)
}