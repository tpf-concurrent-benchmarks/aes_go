package components

import (
	aes "aes_go/aes"
	"os"
	"sync"
)


func handleBlock(block aes.Block, file *os.File) {
	_, err := file.Write(block[:aes.BlockSize])
	Check(err)
}

func sink(wg *sync.WaitGroup, cipherChan chan Message, outputFile string) {

	f, err := os.Create(outputFile)
	Check(err)
	defer f.Close()


	pending := MessageHeap{}
	nextBlock := uint32(0)

	for message := range cipherChan {

		if message.BlockNum > nextBlock {
			pending.Push(message)
			continue
		}
		
		if message.BlockNum < nextBlock {
			panic("Out of order block")
		}

		handleBlock(message.Block, f)
		println(nextBlock," ", message.BlockNum)
		nextBlock++

		for pending.Len() > 0 && pending.Peek().BlockNum == nextBlock {
			message := pending.Pop()
			handleBlock(message.Block, f)
			println(nextBlock," ", message.BlockNum)
			nextBlock++
		}

	}

	wg.Done()
}

func MakeSink(wg *sync.WaitGroup, cipherChan chan Message, outputFile string) {
	// There should only be one sink!
	wg.Add(1)
	go sink(wg, cipherChan, outputFile)
}