package main

func main() {
	for {
		select {
		case data := <-reader:
			// dataの利用
		case <-exit:
			// ループを抜ける
			break
		default:
			// まだデータがきてない
			break
		}
	}
}
