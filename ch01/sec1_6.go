package main

import (
	"net/http"
	"fmt"
	"os"
	"io/ioutil"
	"time"
	"io"
	"strings"
)

func main() {

	// Ch01_Sec1.6 - URLからの取得.

	// 1つを取得.
	load()

	// ゴルーチンを使った、並行取得
	load_parallel()

}

func load_parallel()  {

	// 対象のURL
	urls := []string{
		"http://www.yoheim.net/image/300.jpg",
		"http://www.yoheim.net/image/301.jpg",
		"http://www.yoheim.net/image/302.jpg",
		"http://www.yoheim.net/image/303.jpg",
		"http://www.yoheim.net/image/304.jpg"}

	// ディレクトリの準備
	if err := os.Mkdir("img", 0777); err != nil && !os.IsExist(err) {
		fmt.Println(err)
		return
	}

	start := time.Now()

	ch := make(chan string)

	for i := 0; i < len(urls); i++ {
		url := urls[i]
		go fetch(url, ch)
	}

	for i := 0; i < len(urls); i++ {
		fmt.Println(<-ch)
	}

	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())

}

func fetch(url string, ch chan<- string)  {

	start := time.Now()

	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}

	frags := strings.Split(url, "/")
	savePath := "img/" + frags[len(frags) - 1]
	f, err := os.Create(savePath)
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	defer f.Close()

	//nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	nbytes, err := io.Copy(f, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}

	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d, %s", secs, nbytes, url)
}



func load()  {

	url := "http://www.yoheim.net"

	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
		os.Exit(1)
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch reading %s: %v\n", url, err)
		os.Exit(1)
	}
	resp.Body.Close()

	fmt.Printf("%s", b)

}




