package main

type Reader interface {
	func Read(p []byte) (n int, err error)
}

func main() {

}