package main

import "github.com/yoheiMune/MyGoProject/ch02/tempconv"

func main() {

	f1 := tempconv.Fahrenheit(12)
	println(f1)
	println(tempconv.FToC(f1))
	println(tempconv.FreezingC)
}
