package blkreader

import (
	"bufio"
	"fmt"
	"github.com/metasv/bsvd/wire"
	"log"
	"os"
	"path/filepath"
)

// blk reader read
type BlockFileReader struct {
	// dir for blk files
	Dir string
	// number to read
	FileNum int
	// magic number for different network
	Magic uint32

	currentPos int
	// file reader indicating current file
	reader *bufio.Reader
	file   *os.File
}

//ConcurrentLimit 增加并发数限制
const ConcurrentLimit = 5

func NewBlockFileReader(dirPath string, fileNumber int, magic uint32) *BlockFileReader {
	path := filepath.Join(dirPath, fmt.Sprintf("%s%05d.dat", "blk", fileNumber))
	log.Printf("Scanning file: %v", path)
	fileReader, err := os.Open(path)
	if err != nil {
		log.Println("error while reading blk file", err)
		panic(err)
	}
	newReader := bufio.NewReaderSize(fileReader, 64*1024)

	return &BlockFileReader{
		Dir:        dirPath,
		FileNum:    fileNumber,
		Magic:      magic,
		currentPos: 0,
		reader:     newReader,
		file:       fileReader,
	}
}

// must close file after reading one
func (b *BlockFileReader) Close() {
	_ = b.file.Close()
}

// read next msg block from blk file
func (b *BlockFileReader) NextBlock() (*wire.MsgBlock, error) {
	// 首先使用内置reader读出magic，让
	m, err := readMagic(b.reader)
	if err != nil {
		return nil, err
	}
	// magic 不正确
	if b.Magic > 0 && b.Magic != m {
		return nil, fmt.Errorf("bad magic: %d \n", m)
	}
	// 读取block size
	var size uint32
	err = BinRead(&size, b.reader)
	if err != nil {
		return nil, err
	}
	// 使用wire读取block
	block := wire.NewMsgBlock(&wire.BlockHeader{})
	err = block.Deserialize(b.reader)
	return block, err
}
