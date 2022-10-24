package main

import (
	"bytes"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		return
	}
	for _, filename := range os.Args[1:] {
		data, err := os.ReadFile(filename)
		if err != nil {
			panic(err)
		}
		data2 := bytes.ReplaceAll(data, []byte("\r\n"), []byte("\n"))
		err = os.WriteFile(filename, data2, 0666)
		if err != nil {
			panic(err)
		}
	}
}
