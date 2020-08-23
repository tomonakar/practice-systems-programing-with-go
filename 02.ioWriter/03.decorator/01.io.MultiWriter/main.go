package main

import (
	"io"
	"os"
)

func main() {
	// io.Writerのデコレータ
	// デコレータ：オブジェクトをラップして追加の機能を実現するという、GoFのデザインパターン用語のデコレータ

	file, err := os.Create("multiwriter.txt")
	if err != nil {
		panic(err)
	}

	// io.MultiWriterは複数のio.Writerを受け取り、それら全てに対して書き込まれた内容を同時に書き込むデコレータ
	writer := io.MultiWriter(file, os.Stdout)
	io.WriteString(writer, "io.Multiwriter example\n")
}
