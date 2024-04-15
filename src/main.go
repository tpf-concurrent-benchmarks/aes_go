package main

import (
	. "aes_go/components"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"

	godotenv "github.com/joho/godotenv"
)

func processFile(inputFile string, outputFile string, makeWorkers WorkerBuilder, key string, numWorkers int) {
	
	workers_wg := sync.WaitGroup{}
	inputChan := make(chan Message, numWorkers*2)
	sink_wg := sync.WaitGroup{}
	outputChan := make(chan Message, numWorkers*10)
	
	makeWorkers(numWorkers, &workers_wg, inputChan, outputChan, key)
	MakeSink(&sink_wg, outputChan, outputFile)

	SendMessages(inputFile, inputChan)

	close(inputChan)
	workers_wg.Wait()
	close(outputChan)
	sink_wg.Wait()
}


func _main() {
	cipherKey := "0123456789abcdef"

	numWorkers, err := strconv.Atoi(os.Getenv("CORES"))
	println("Using", numWorkers, "goroutines")
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
	println("Repeating", times, "times")
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
	println("Started")
	loadEnv()
	RunAndMeasure(_loop_main)
	println("Done!")
}
