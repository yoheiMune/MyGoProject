package main

import (
	"net/http"
	"fmt"
)

// http://qiita.com/taizo/items/bf1ec35a65ad5f608d45
// https://golang.org/pkg/net/http/

type String string

func (s String) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, s)
}

func main() {

	// handlerを定義する場合.
	http.HandleFunc("/a", handler)

	// クロージャー的なのでもOK.
	http.HandleFunc("/b", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello from Closure")
	})

	// DuckTyping的に、ServeHTTP関数があれば良い.
	http.Handle("/c", String("Duck Typing!!!"))

	// 例えばPOSTのみ.
	http.HandleFunc("/onlyPost", handleOnlyPost)

	// パラメータ受け取り
	http.HandleFunc("/params", handleParams)

	// 静的ファイル配信.
	http.Handle("/contents/", http.FileServer(http.Dir("./")))
	http.Handle("/mysecret/", http.StripPrefix("/mysecret/", http.FileServer(http.Dir("./contents"))))

	// basic認証.
	http.HandleFunc("/basicAuth", handleBasicAuth)

	http.ListenAndServe(":8080", nil)
	//http.ListenAndServe(":8080", String("Other Handler"))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World from Go.")
}

func handleOnlyPost(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed) // 405
		w.Write([]byte("POSTだけだよー"))
		return
	}

	w.Write([]byte("OK"))
}

func handleParams(w http.ResponseWriter, r *http.Request) {

	// https://golang.org/pkg/net/http/#Request
	// https://golang.org/pkg/net/url/#URL

	// クエリパラメータ取得してみる
	fmt.Fprintf(w, "クエリ：%s\n", r.URL.RawQuery)

	// ボディを受け取ってみる
	//if r.Method == http.MethodPost {

		r.ParseForm()
		form := r.PostForm
		fmt.Fprintf(w, "フォーム：\n%v\n", form)

		// または、クエリパラメータも含めて全部.
		params := r.Form
		fmt.Fprintf(w, "フォーム2：\n%v\n", params)
	//}
}

func handleBasicAuth(w http.ResponseWriter, r *http.Request) {

	username, password, ok := r.BasicAuth()
	if ok == false {
		w.Header().Set("WWW-Authenticate", `Basic realm="SECRET AREA"`)
		w.WriteHeader(http.StatusUnauthorized) // 401
		fmt.Fprintf(w, "%d Not authorized.", http.StatusUnauthorized)
		return
	}

	// 認証チェック
	if username != "my" || password != "secret" {
		w.Header().Set("WWW-Authenticate", `Basic realm="SECRET AREA"`)
		w.WriteHeader(http.StatusUnauthorized) // 401
		fmt.Fprintf(w, "%d Not authorized.", http.StatusUnauthorized)
		return
	}

	// OK
	fmt.Fprint(w, "OK")

}