package main

import "fmt"

func main() {

	// 配列.
	{
		var a [3]int
		println(a[0])
		println(a[len(a) - 1])

		// 1番目に値を代入
		a[1] = 2

		// 参照もできる
		println(a[0], a[1], a[2])


		for i, v := range a {
			println(i, v)
		}

		for _, v := range a {
			println(v)
		}

		var q [3]int = [3]int{1, 2, 3}
		println(q[1])

		r := [...]int{1, 2, 3, 4, 5}
		fmt.Printf("%T\n", r)

		type Currency int
		const (
			USD Currency = iota
			EUR
			JPN
		)
		symbol := [...]string{USD: "$", EUR: "€", JPN: "¥"}
		println(USD, EUR, JPN)
		println(symbol[EUR])

		p := [...]int{99: -1}
		fmt.Printf("%T\n", p)
		for _, v := range p {
			fmt.Printf("%d ", v)
		}
		println("")
	}

	{
		a := [2]int{1, 2}
		b := [...]int{1, 2}
		println(a == b) // true

		//c := [...]int{1, 2, 3}
		// println(a == c)  // コンパイルエラー（型が違うから）
	}


	{

		println("=================")

		// 初期化（要素数=3、値は未指定）
		var a [3]int
		fmt.Printf("%v\n", a)

		// 初期化（要素数=3、値を指定）
		var b [3]int = [3]int{1, 2, 3}
		fmt.Printf("%v\n", b)
		// または
		c := [3]int{1, 2, 3}
		fmt.Printf("%v\n", c)

		// 初期化（要素数を省略、値を指定）
		d := [...]int{1, 2, 3, 4, 5}
		fmt.Printf("%v\n", d)

		// 初期化（要素番号を指定）
		e := [...]string{0:"a", 2:"c", 3:"d"}
		fmt.Printf("%v\n", e)


		println(c[1])  // 2

		for i, v := range c {
			println(i, v)
		}

		a[1] = 10
		fmt.Printf("%v\n", a)


		f := [2]int{1, 2}
		g := [...]int{1, 2}
		h := [2]int{1, 3}
		println(f == g, f == h)  // true false

		//i := [3]int{1, 2, 3}
		//println(f == i)

	}







}
