package components

import (
	aes "aes_go/aes"
	"bufio"
	"os"
)

func padBatch(buff []byte, n int) {
	for i := n; i < int(BatchSize)*int(aes.BlockSize); i++ {
		buff[i] = 0
	}
}

func SendMessages(inputFile string, inputChan chan Message) {
	buff := make([]byte, int(BatchSize)*int(aes.BlockSize))
	f, err := os.Open(inputFile)
	Check(err)
	defer f.Close()
	reader := bufio.NewReader(f)

	blockNum := uint32(0)
	for n, err := reader.Read(buff); n > 0; n, err = reader.Read(buff) {
		Check(err)

		blockAmount := BatchSize
		if n < int(BatchSize)*int(aes.BlockSize) {
			padBatch(buff, n)
			blockAmount = uint8(n)/aes.BlockSize + 1
		}

		blocks := [BatchSize]aes.Block{}

		for i := 0; i < n; i += int(aes.BlockSize) {
			block := aes.Block{}
			copy(block[:], buff[i:i+int(aes.BlockSize)])
			blocks[i/int(aes.BlockSize)] = block
		}

		inputChan <- Message{Num: blockNum, Batch: blocks, Blocks: blockAmount}
		blockNum++
	}
}