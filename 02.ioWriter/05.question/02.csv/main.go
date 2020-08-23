package main

import (
	"encoding/csv"
	"os"
)

func main() {
	file, err := os.Create("test.csv")
	if err != nil {
		panic(err)
	}
	writer := csv.NewWriter(file)
	defer file.Close()

	writer.Write([]string{"hoge", "fuga", "bar"})
	writer.Flush()
}
