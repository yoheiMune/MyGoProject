package main

import (
	log "github.com/cihub/seelog"
	"fmt"
	"github.com/yoheiMune/MyGoProject/002_logging/sub"
)

func main() {

	/**
		SeeLogの調査.

		# 参照
			https://github.com/cihub/seelog
		# インストール
			go get -u github.com/cihub/seelog
	 */

	defer log.Flush()

	// デフォルトで出力.
	log.Info("Hello from SeeLog!")


	// フォーマット指定
	// フォーマット一覧：https://github.com/cihub/seelog/wiki/Format-reference
	config := `
		<seelog type="sync">
			<outputs formatid="main">
				<console/>
			</outputs>
			<formats>
				<format id="main" format="%Date(2006-01-02T15:04:05Z07:00) %LEVEL %RelFile(%Line)(%FuncShort) %Msg%n"/>
			</formats>
		</seelog>
	`
	logger, err := log.LoggerFromConfigAsBytes([]byte(config))
	if err != nil {
		fmt.Printf("エラーだよ：%v", err)
		panic("ロガーの初期化エラー")
	}
	log.ReplaceLogger(logger)

	// 試しに表示.
	log.Debug("デバッグログ")

	// サブパッケージでの呼び出しもしてみる
	sub.A()
}
