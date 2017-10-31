package main

import (
	"runtime"
	"fmt"
)

func main() {

	/**
		runtime.Callerでスタックトレースを取得するサンプル.
		この機能を用いて、logでファイル名や行数を表示している.

		https://github.com/golang/go/blob/d1731f8cbc1adde72a2b2e7513b6ca73476c7f0d/src/log/log.go#L159
	 */
	a()
	//=> file=/Users/munesadayohei/go/src/github.com/yoheiMune/MyGoProject/002_logging/main.go, line=30, func=main.b
	//=> file=/Users/munesadayohei/go/src/github.com/yoheiMune/MyGoProject/002_logging/main.go, line=26, func=main.a
	//=> file=/Users/munesadayohei/go/src/github.com/yoheiMune/MyGoProject/002_logging/main.go, line=16, func=main.main
	//=> file=/usr/local/go/src/runtime/proc.go, line=185, func=runtime.main
	//=> file=/usr/local/go/src/runtime/asm_amd64.s, line=2197, func=runtime.goexit

}

func a() {
	b()
}

func b() {
	c()
}

func c() {
	i := 0
	for {
		pt, file, line, ok := runtime.Caller(i)
		if !ok {
			break
		}
		// (参考)関数名を取得する
		// http://kwmt27.net/index.php/2013/10/05/golang-how-to-get-function-name-with-reflection/
		funcName := runtime.FuncForPC(pt).Name()
		fmt.Printf("file=%s, line=%d, func=%v\n", file, line, funcName)
		i += 1
	}
}
