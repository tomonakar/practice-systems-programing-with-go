# 第３章
## 低レベルアクセスへの入り口: io.Reader

- io.Readerとその仲間たち
- 少ないコード量でio.Readerからデータを効率よく読み込むための補助的な関数群
- io.Readerとio.Writer以外の入出力インターフェース
- io.Readerを満たす構造体で特に品案に使われるもの（標準入力、ファイル、ソケット、メモリのバッファ）
- バイナリ解析に便利な機能群
- テキスト解析に便利な機能群
- ちょっと抽象的なio.Readerの構造体

## io.Readerの概要
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

## io.Readerの補助関数
### 読み込みの補助関数

`ioutil.ReadAll()` 終端記号にあたるまで全てのデータを読み込む。おそらくもっとも利用する。
```go
// 全て読み込む
buffer, err := ioutil.ReadAll(reader)
```

`io.ReadFull()` 指定したバッファサイズ分まで読み込む。読み込めない場合はエラーを返す。
```go
// 4バイト読み込めないとエラー
buffer := make([]byte, 4)
size, err := io.ReadFull(reader, buffer)
```

### コピーの補助関数
io.Readerからio.Writerにそのままデータを渡すときにはコピーを使う。もっともよく使うのが、全てを読み尽くして書き込む `io.Copy()`. コピーするバイト数を指定する場合は、 `io.CopyN()`を使う

```go
// 全てコピー
writeSize, err := io.Copy(writer, reader)

// 指定したサイズ分コピー
writeSize, err := io.CopyN(writer, reader, size)
```

あらかじめコピー量が決まっていて無駄にバッファを使いたくないとき、なんどもコピーするからバッファを取り回したい時などに使うのが `io.CopyBuffer()`。これを使うと自分が作った作業バッファを渡すことができる。デフォは32KBを内部で確保して使っている。
```go
// 8KBのバッファを使う
buffer := make([]byte, 8 * 1024)
io.CopyBuffer(writer, reader, buffer)
```

## 入出力に関するio.Weiterとio.Reader以外のインターフェース
入出力ではクローズ処理など、Read・Write以外にも様々な処理が必要。よく使うインターフェースを記載する。

- io.Closerインターフェース
  - `func Close() error` メソッドを持つ
  - 使い終わったファイルを閉じる

- io.Seekerインターフェース
  - `func Seek(offset int64, whence int) (int64, error)` メソッドを持つ
  - 読み書き位置を移動する

- io.ReaderAtインターフェース
  - `io.ReadAt(p []byte, off int64)（n int, err error)` メソッドを持つ
  - 対象となるオブジェクトがランダムアクセスを行える場合に、好きな位置を自由にアクセスする時に使う

### 入出力関連インターフェースのキャスト

引数に`io.ReadCloser`が要求されているが、今あるオブジェクトはio.Readerしか満たしてないという場合がある。
例えば、ソケット読み込み用の関数を作成していて、その関数の引数はio.ReadCloserだが、
ユニットテストにはio.Readerインターフェースを満たすString.Readerや
bytes.Bufferを使いたいというケースが考えられる。

その場合は、ioutil.NopCloser()関数を使うと、ダミーのCloser()メソッドを持って、
io.ReadCloserのふりをするラッパーオブジェクトを得られる。

```go
import (
  "io"
  "io/ioutil"
  "strings"
)

var reader io.Reader = strings.NewReader("テストデータ")
var readCloser io.ReadCloser = ioutil.NopCloser(reader)
```

## io.Readerの構造体を満たしよく使うもの
### 標準入力: os.Stdin

標準入力の構造体はos.Stdin.
以下のプログラムを実行すると入力待ちになり、以降はEnterが押されるたびに結果が帰ってくる。
Ctrl+Dで終了する。

```go
package main

import (
  "fmt"
  "io"
  "os"
)

func main() {
  for {
    buffer := make([]bytes, 5)
    size, err := os.Stdin.Read(buffer)
    if err := io.EOF {
      fmt.Println("EOF")
      break
    }
    fmt.Printf("size=%d input='%s'\n", size, string(buffer))
  }
}
```

以下のようにすると、自身のプログラムを指定したバッファサイズ（ここでは５）ごとに区切って表示する

```sh
$ go run stdin.go < stdin.go
```

### ファイル入力: os.File
ファイル入力はos.File構造体を使う.
ファイルの新規作成はos.Create()関数で実行する。
os.Open()を使うと既存ファイルを開くことができる。

```go
func Open(name string) (*File, error) {
  return OpenFile(name, O_RDONLY, O)
}
func Create(name string) (*File, error) {
  return OpenFIle(name, O_RDWR|O_CREATE|O_TRUNC, 06666)
}
```

ファイル読み込みの例。
io.Copy()を使い、標準出力にファイルの内容を書き出している。
```go
package main
import (
  "io"
  "os"
)
func main() {
  file, err := os.Open("file.go")
  if err != nil {
    panic(err)
  }
  defer file.Close()
  io.Copy(os.Stdout, file)
}
```

### ネットワーク通信の読み込み

net.Dial()で返えされるconnがnet.Conn型で、これをio.Copyを使って標準出力にコピーすることでデータを一括して読み込んでいる。
以下のコードではコピーする内容は生のHTTP通信内容そのもの。

```go
package main
import (
  "io"
  "os"
  "net"
)

func main() {
  conn, err := net.Dial("tcp", "ascii.jp:80")
  if err != nil {
    panic(err)
  }
  conn.Write([]byte("GET / HTTP/1.0/\r\nHOST: ascii.jp\r\n\r\n"))
  io.Copy(os.Stdout, conn)
}
```

上記だと、HTTPを読み込む際に、RFCにしたがって通信内容をパースする必要が生じる。
`http.ReadResponse()`を使うことでHTTP通信をパースして使いやすくしてくれる。


```go
package main
import (
  "io"
  "os"
  "net"
)

func main() {
  conn, err := net.Dial("tcp", "ascii.jp:80")
  if err != nil {
    panic(err)
  }
  conn.Write([]byte("GET / HTTP/1.0/\r\nHOST: ascii.jp\r\n\r\n"))
  res, err := http.ReadResponse(bufio.NewReader(conn), nil)

  // ヘッダーを表示
  fmt.Printl(res.Header)
  //ボディを表示.ボディ表示後はClose()する必要がある
  defer res.Body.Close()
  io.Copy(os.Stdout, res.Body)
}
```