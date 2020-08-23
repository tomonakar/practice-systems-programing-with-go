package main

import (
	"compress/gzip"
	"io"
	"os"
)

func main() {
	// ファイル作成
	file, err := os.Create("test.txt.gz")
	if err != nil {
		panic(err)
	}
	// gzipファイルのwriterを作成
	writer := gzip.NewWriter(file)
	writer.Header.Name = "test.txt"
	// ファイルに書き込み
	io.WriteString(writer, "gzip.Writer example\n")
	// ファイルを閉じる
	writer.Close()
}
