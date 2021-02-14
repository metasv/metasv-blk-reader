# MetaSV blk Reader

MetaSV golang tool to read blk.data file , and parse it into blocks

## Usage

```go

func main()  {
	// create a new reader, replace the blk folder path and file number
	reader := blkreader.NewBlockFileReader("./", 2112, blkreader.MainNetMagic)
	defer reader.Close()
	for  {
		// read next block from blk file
		block, err := reader.NextBlock()
		// break if EOF
		if err == io.EOF{
			break
		}
		if err != nil {
			panic(err)
		}
		fmt.Println(block.BlockHash().String())
	}
}

```