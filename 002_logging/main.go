package main

import (
	"log"
	"fmt"
	"runtime"
)

func main() {

	/*
		1.
		標準のlogモジュールでファイル名とかを取得する方法
	 */
	// https://github.com/golang/go/blob/d1731f8cbc1adde72a2b2e7513b6ca73476c7f0d/src/log/log.go#L40-L41

	// full file name and line number: /a/b/c/d.go:23
	log.SetFlags(log.Llongfile)
	log.Println("ログ1")
	//=> /Users/munesadayohei/go/src/github.com/yoheiMune/MyGoProject/002_logging/main.go:16: 標準モジュールで表示したログです.

	// final file name element and line number: d.go:23. overrides Llongfile
	log.SetFlags(log.Lshortfile)
	log.Println("ログ2")
	//=> main.go:21: 標準モジュールで表示したログです.



	/*
		自分でも取得してみる.
	 */
	pt, file, line, ok := runtime.Caller(0)
	if !ok {
		fmt.Println("スタックトレースの取得失敗")
		return
	}
	// (参考)関数名を取得する
	// http://kwmt27.net/index.php/2013/10/05/golang-how-to-get-function-name-with-reflection/
	funcName := runtime.FuncForPC(pt).Name()
	fmt.Printf("file=%s, line=%d, func=%v\n", file, line, funcName)
}
