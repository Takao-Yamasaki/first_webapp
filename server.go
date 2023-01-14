package main

import (
	"fmt"
	"net/http"
)

// ハンドラ: HTTPリクエストを受理して、それを処理する
// Requestから情報を取り出して、ResponseWriterで書き込む
func handler(writer http.ResponseWriter, request *http.Request) {
	// レスポンスを生成
	fmt.Fprintf(writer, "Hello World, %s!", request.URL.Path[1:])
}

func main() {
	// ルートが呼び出された時、ハンドラが起動する
	http.HandleFunc("/", handler)
	// ポート8080を監視するようサーバーを起動する
	http.ListenAndServe(":8080", nil)
}
