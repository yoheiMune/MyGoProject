package main

import (
	"log"
)

func main() {

	/*
		1.
		標準のlogモジュールでファイル名とかを取得する方法
	 */
	// https://github.com/golang/go/blob/d1731f8cbc1adde72a2b2e7513b6ca73476c7f0d/src/log/log.go#L40-L41

	// full file name and line number: /a/b/c/d.go:23
	log.SetFlags(log.Llongfile)
	log.Println("標準モジュールで表示したログです.")
	//=> /Users/munesadayohei/go/src/github.com/yoheiMune/MyGoProject/002_logging/main.go:16: 標準モジュールで表示したログです.

	// final file name element and line number: d.go:23. overrides Llongfile
	log.SetFlags(log.Lshortfile)
	log.Println("標準モジュールで表示したログです.")
	//=> main.go:21: 標準モジュールで表示したログです.

}
