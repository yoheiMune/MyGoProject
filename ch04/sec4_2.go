package main

import "fmt"

func main() {

	// 4.2. スライス
	{
		members := [...]string{"Yohei", "John", "Kento", "Taiga", "Shingo"}

		m1 := members[0:2]  // スライス.
		fmt.Printf("%T\n", m1)
		println(cap(m1), len(m1))

		m2 := members[3:]
		println(cap(m2), len(m2))

		fmt.Println(m2)

		//_ = m2[:20] // パニック：範囲外
	}

	// 4.2.1. append関数
	{
		var runes []rune
		for _, r := range "Hello,世界" {
			runes = append(runes, r)
		}
		fmt.Printf("%q\n", runes) // ['H' 'e' 'l' 'l' 'o' ',' '世' '界']
	}
}
