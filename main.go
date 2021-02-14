package main

import (
	"awesomeProject1/blkreader"
	"fmt"
	"io"
)

func main() {
	reader := blkreader.NewBlockFileReader("./", 2112, blkreader.MainNetMagic)
	defer reader.Close()
	for {
		block, err := reader.NextBlock()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		fmt.Println(block.BlockHash().String())
	}
}
