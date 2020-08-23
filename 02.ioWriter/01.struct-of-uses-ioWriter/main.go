package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func main() {
	// io.Writerインターフェースを使うことで、どれも同じWrite()メソッドを使って書き出すことができる

	// ファイル出力
	file, err := os.Create("test.test")
	if err != nil {
		panic(err)
	}
	file.Write([]byte("os.File example\n"))
	file.Close()

	// 画面出力
	os.Stdout.Write([]byte("os.Stdout example\n"))

	// 書かれた内容を記憶しておくバッファ
	var buffer bytes.Buffer
	buffer.Write([]byte("bytes.Buffer example\n"))
	fmt.Println(buffer.String())

	// bytes.BufferのWriteStringメソッドは文字列を受け取れる
	// しかし、io.Writerのメソッドではないので、他の構造体では使えない
	buffer.WriteString("bytes.Buffer example2\n")
	fmt.Println(buffer.String())

	// 代わりにio.WriteString() を使えばキャストは不要になる
	var buffer2 bytes.Buffer
	io.WriteString(&buffer2, "bytes.Buffer example3\n")
	fmt.Println(buffer2.String())
}
