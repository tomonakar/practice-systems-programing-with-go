# 第３章
### 低レベルアクセスへの入り口: io.Reader

- io.Readerとその仲間たち
- 少ないコード量でio.Readerからデータを効率よく読み込むための補助的な関数群
- io.Readerとio.Writer以外の入出力インターフェース
- io.Readerを満たす構造体で特に品案に使われるもの（標準入力、ファイル、ソケット、メモリのバッファ）
- バイナリ解析に便利な機能群
- テキスト解析に便利な機能群
- ちょっと抽象的なio.Readerの構造体

### io.Readerの概要
- 前章の復習：インターフェースはメソッド宣言をまとめたもの
- io.ReaderのインターフェースはReadメソッドを持っている

```go
type Reader interface {
	func Read(p []byte) (n int, err error)
}
```

- 以下は、io.Readerインターフェースを満たす何らかの型rを使って、データを読み込む例
- Go言語でメモリを確保するには、make()を使う
- make()を使って1024バイトの入力用バッファを用意し、そこにデータを読み込む

```go
	// 1024バイトのバッファをmakeで作る
	buffer := make([]byte, 1024)

	// sizeは実際に読み込んだ数
	size, err := r.Read(buffer)
```