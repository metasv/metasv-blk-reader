package main

import (
	"fmt"
	"github.com/metasv/metasv-blk-reader/blkreader"
	"io"
)

func main() {
	reader := blkreader.NewBlockFileReader("./", 2005, blkreader.MainNetMagic)
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
