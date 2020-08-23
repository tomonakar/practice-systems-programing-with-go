package main

import (
	"io"
	"log"
	"net"
	"net/http"
	"os"
)

func main() {
	// io.Writerでインターネットにアクセスしてみる
	// net.Dial()を使うと、net.Connという通信コネクションを表すインターフェースが返ってくる
	// これは、io.Writerでもio.Readerでも使える

	conn, err := net.Dial("tcp", "ascii.jp:80")
	if err != nil {
		panic(err)
	}

	// サーバからのレスポンスをconnに書き込む
	io.WriteString(conn, "GET / HTTP/1.0/r/nHost: ascii.jp\r\n\r\n")
	// connに書き込まれたデータを標準出力にコピーし、画面に表示
	io.Copy(os.Stdout, conn)

	// httpリクエストを直接書き込む
	req, err := http.NewRequest("GET", "http://ascii.jp", nil)
	req.Write(conn)
	err = conn.Close()
	if err != nil {
		log.Fatal(err)
	}

	// http.ResponseWriterはwebサーバーからブラウザに対してメッセージを書き込む
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

// http.ResponseWriterを使って、web serverからブラウザに対してメッセージを書き込む
func handler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "http.ResponseWriter sample")
}
