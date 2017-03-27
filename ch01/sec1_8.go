package main

func main() {

	// switch文
	target := "header"
	switch target {
	case "header":
		println("header!!!")
	case "tail":
		println("tail!!!")
		default:
		println("default!!!")
	}

	// switch文
	num := -10
	switch  {
	case num >= 0:
		println("unsigned")
	default:
		println("signed")
	}

	// 名前付きの型
	type Point struct {
		X, Y int
	}
	//var point Point


}
