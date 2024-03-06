package components

import (
	aes "aes_go/aes"
	"sync"
)

func cipherBytes(wg *sync.WaitGroup, plainChan chan Message, cipherChan chan Message, key string) {	
	cipher, _ := aes.FromString(key)

	for message := range plainChan {
		message.Block = cipher.CipherBlock(message.Block)
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
		message.Block = cipher.InvCipherBlock(message.Block)
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
