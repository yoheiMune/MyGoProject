package main

import "fmt"

func main() {

	// 配列.
	{
		var a [3]int
		println(a[0])
		println(a[len(a) - 1])

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

		c := [...]int{1, 2, 3}
		// println(a == c)  // コンパイルエラー（型が違うから）
	}







}
