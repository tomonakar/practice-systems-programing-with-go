package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

func main() {
	// Fprintf はフォーマット出力のための関数
	// %v はなんでも表示できるフォーマット指定子
	// 以下ではDateを表示している
	fmt.Fprintf(os.Stdout, "Write with os.Stdout at %v\n", time.Now())

	// JSONを整形してio.Writerに書き出す
	encoder := json.NewEncoder(os.Stdout)
	encoder.SetIndent("", "  ")
	encoder.Encode(map[string]string{
		"example": "encoding/json",
		"hello":   "world",
	})

	// net/httpパッケージのRequestは、io.Writerに書き出すメソッドを持つ用途が限定された構造体
	// クライアント側のリクエストを送るときにも使えるし、サーバ側でレスポンスを返すときにクライアントの情報をパースするのにも使える
	request, err := http.NewRequest("GET", "http://ascii.jp", nil)
	if err != nil {
		panic(err)
	}
	request.Header.Set("X-TEST", "ヘッダーも追加できます")
	request.Write(os.Stdout)
}
