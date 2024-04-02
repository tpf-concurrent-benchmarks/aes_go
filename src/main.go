package main

import (
	aes "aes_go/aes"
	. "aes_go/components"
	"bufio"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"

	godotenv "github.com/joho/godotenv"
)

func padBatch(buff []byte, n int) {
	for i := n; i < int(BatchSize)*int(aes.BlockSize); i++ {
		buff[i] = 0
	}
}

func sendMessages(inputFile string, inputChan chan Message) {
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
			blockAmount = uint8(n) / aes.BlockSize+1
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

func processFile(inputFile string, outputFile string, makeWorkers WorkerBuilder, key string, numWorkers int) {
	
	workers_wg := sync.WaitGroup{}
	inputChan := make(chan Message, numWorkers*2)
	sink_wg := sync.WaitGroup{}
	outputChan := make(chan Message, numWorkers*10)
	
	makeWorkers(numWorkers, &workers_wg, inputChan, outputChan, key)
	MakeSink(&sink_wg, outputChan, outputFile)

	sendMessages(inputFile, inputChan)

	close(inputChan)
	workers_wg.Wait()
	close(outputChan)
	sink_wg.Wait()
}


func _main() {
	cipherKey := "0123456789abcdef"

	numWorkers, err := strconv.Atoi(os.Getenv("CORES"))
	Check(err)

	plainText := os.Getenv("PLAIN_TEXT")
	encryptedText := os.Getenv("ENCRYPTED_TEXT")
	decryptedText := os.Getenv("DECRYPTED_TEXT")

	if plainText != "" && encryptedText != "" {
		log.Println("  > Encrypting", plainText, "to", encryptedText)
		processFile(plainText, encryptedText, MakeCipherWorkers, cipherKey, numWorkers)
	}

	if encryptedText != "" && decryptedText != "" {
		log.Println("  > Decrypting", encryptedText, "to", decryptedText)
		processFile(encryptedText, decryptedText, MakeInvCipherWorkers, cipherKey, numWorkers)
		RemoveTrailingNulls(decryptedText)
	}
}

func _loop_main() {
	times, err := strconv.Atoi(os.Getenv("REPEAT"))
	Check(err)
	for i := 0; i < times; i++ {
		log.Println("Iteration", i)
		_main()
	}
}

func loadEnv() {
	err := godotenv.Load(".env")
	Check(err)

	_doLog := os.Getenv("LOG")
	doLog := strings.ToLower(_doLog) == "true"

	if doLog {
		log.SetOutput(os.Stdout)
	} else {
		log.SetOutput(io.Discard)
	}
}


func main() {
	println("Starting...")
	loadEnv()
	RunAndMeasure(_loop_main)
	println("Done!")
}
