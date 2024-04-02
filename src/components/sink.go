package components

import (
	"bufio"
	"os"
	"sync"
)


func handleBatch(message Message, writer *bufio.Writer) {
	for i:=0; i<int(message.Blocks); i++ {
		writer.Write(message.Batch[i][:])
	}
}

func sink(wg *sync.WaitGroup, cipherChan chan Message, outputFile string) {

	f, err := os.Create(outputFile)
	Check(err)
	defer f.Close()
	writer := bufio.NewWriter(f)


	pending := MessageHeap{}
	nextMessage := uint32(0)

	for message := range cipherChan {

		if message.Num > nextMessage {
			pending.Push(message)
			continue
		}
		
		if message.Num < nextMessage {
			panic("Out of order block")
		}

		handleBatch(message, writer)
		nextMessage++

		for pending.Len() > 0 && pending.Peek().Num == nextMessage {
			message := pending.Pop()
			handleBatch(message, writer)
			nextMessage++
		}

	}

	writer.Flush()

	wg.Done()
}

func MakeSink(wg *sync.WaitGroup, cipherChan chan Message, outputFile string) {
	// There should only be one sink!
	wg.Add(1)
	go sink(wg, cipherChan, outputFile)
}