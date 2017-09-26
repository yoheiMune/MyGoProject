package main

import (
	"github.com/pkg/errors"
	"io/ioutil"
	"fmt"
)

func main() {

	// TODO 途中.

	/*

		pkg/errors のサンプル実装

		# 参照
			https://github.com/pkg/errors

		# インストール
			go get -u github.com/pkg/errors
	*/

	err := a()
	if err != nil {
		fmt.Printf("エラーだよ. %v\n", err)
		fmt.Println("------------")
	} else {
		fmt.Println("無事に処理できました！")
	}

}

func a() error {
	err := b()
	if err != nil {
		return errors.Wrap(err, "〇〇しようと思ったけどエラー.")
	}
	return nil
}

func b() error {
	fpath := "notfound.txt"
	_, err := ioutil.ReadFile(fpath)
	if err != nil {
		return errors.Wrapf(err, "「%s」を読み込めませんでした", fpath)
	}
	return nil
}