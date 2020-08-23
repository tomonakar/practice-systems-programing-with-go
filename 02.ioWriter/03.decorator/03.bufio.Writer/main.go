package main

import (
	"bufio"
	"os"
)

func main() {
	// bufio.Writerは出力結果を一時的にためておいて、ある程度の分量ごとにまとめて書き出す
	buffer := bufio.NewWriter(os.Stdout)
	buffer.WriteString("bufio.Writer")
	// Flushメソッドを呼ぶと、後続のio.Writerに書き出す
	// 呼ばないと、書き込まれたデータを抱えたまま消滅する
	// Flushを自動で呼ぶ場合は、バッファサイズ指定のbufio.NewWriterSize(os.Stdout, size)関数でbufio.Writerを作ると良い
	// sizeのデフォルトは4096バイト
	buffer.Flush()
	buffer.WriteString("example\n")
	buffer.Flush()
}
