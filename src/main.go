package main

import (
	aes "aes_go/aes"
	. "aes_go/components"
	"os"
	"sync"
)

func padBlock(buff []byte, n int) {
	//fill the rest of the block with 0 (NULL)
	for i := uint8(n); i < aes.BlockSize; i++ {
		buff[i] = 0
	}
}

func processFile(inputFile string, outputFile string, makeWorkers WorkerBuilder, key string, numWorkers int) {

	buff := make([]byte, aes.BlockSize)

	f, err := os.Open(inputFile)
	Check(err)
	defer f.Close()

	workers_wg := sync.WaitGroup{}
	inputChan := make(chan Message, numWorkers*2)
	sink_wg := sync.WaitGroup{}
	outputChan := make(chan Message, numWorkers*2)

	makeWorkers(numWorkers, &workers_wg, inputChan, outputChan, key)
	MakeSink(&sink_wg, outputChan, outputFile)

	blockNum := uint32(0)
	for n, err := f.Read(buff); n > 0; n, err = f.Read(buff) {
		Check(err)
		if uint8(n) < aes.BlockSize {
			padBlock(buff, n)
		}
		plainText := aes.Block(buff)
		inputChan <- Message{BlockNum: blockNum, Block: plainText}
		blockNum++
	}

	close(inputChan)
	workers_wg.Wait()
	close(outputChan)
	sink_wg.Wait()
}

func _main() {
	cipherKey := "0123456789abcdef"
	numWorkers := 10

	processFile("input.txt", "ciphered.txt", MakeCipherWorkers, cipherKey, numWorkers)
	processFile("ciphered.txt", "deciphered.txt", MakeInvCipherWorkers, cipherKey, numWorkers)
	RemoveTrailingNulls("deciphered.txt")
}

func main() {
	RunAndMeasure(_main)
}
