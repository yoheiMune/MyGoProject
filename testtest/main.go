package main

import (
	"math"
	"fmt"
)

// 値を2乗する
func Square(x int) int {
	// バグを埋め込んでみる
	//if x == -2 {
	//	return -4
	//}
	result := math.Pow(float64(x), 2.0)
	return int(result)
}

func main() {
	result := Square(10)
	fmt.Println(result)
}
