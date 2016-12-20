package main

import (
	"net/http"
	"html/template"
)

// リンク：http://qiita.com/taizo/items/bf1ec35a65ad5f608d45


//func handler(w http.ResponseWriter, r *http.Request) {
//	fmt.Fprint(w, "Hello World")
//}
//
//func main() {
//	http.HandleFunc("/", handler)
//	http.ListenAndServe(":8080", nil)
//}


//type String string
//
//func (s String) ServeHTTP(w http.ResponseWriter, r *http.Request) {
//	fmt.Fprint(w, s)
//}
//
//func main() {
//	http.Handle("/", String("Hello World2."))
//	http.ListenAndServe(":8080", nil)
//}


type Page struct {
	Title string
	Count int
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	page := Page{"Hello World3.", 1} // テンプレートデータ
	//tmpl, err := template.New("new").Parse("{{.Title}} {{.Count}} count") // テンプレート文字列
	tmpl, err := template.ParseFiles("layout.html")
	if err != nil {
		panic(err)
	}

	err = tmpl.Execute(w, page)
	if err != nil {
		panic(err)
	}
}

func main() {
	http.HandleFunc("/", viewHandler);
	http.ListenAndServe(":8080", nil)
}
































