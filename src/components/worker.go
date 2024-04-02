package components

import (
	aes "aes_go/aes"
	"sync"
)

func cipherBytes(wg *sync.WaitGroup, plainChan chan Message, cipherChan chan Message, key string) {	
	cipher, _ := aes.FromString(key)

	for message := range plainChan {
		for i, block := range message.Batch {
			message.Batch[i] = cipher.CipherBlock(block)
		}
		cipherChan <- message
	}

	wg.Done()
}

func MakeCipherWorkers(num int, wg *sync.WaitGroup, plainChan chan Message, cipherChan chan Message, key string) {
	for i := 0; i < num; i++ {
		wg.Add(1)
		go cipherBytes(wg, plainChan, cipherChan,key)
	}
}

func invCipherBytes(wg *sync.WaitGroup, cipherChan chan Message, plainChan chan Message, key string) {
	cipher, _ := aes.FromString(key)

	for message := range cipherChan {
		for i, block := range message.Batch {
			message.Batch[i] = cipher.InvCipherBlock(block)
		}
		plainChan <- message
	}

	wg.Done()
}

func MakeInvCipherWorkers(num int, wg *sync.WaitGroup, cipherChan chan Message, plainChan chan Message, key string) {
	for i := 0; i < num; i++ {
		wg.Add(1)
		go invCipherBytes(wg, cipherChan, plainChan,key)
	}
}


type WorkerBuilder func (int, *sync.WaitGroup, chan Message, chan Message, string)
