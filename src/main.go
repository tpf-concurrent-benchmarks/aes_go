package main

import (
	aes "aes_go/aes"
	"time"

	"github.com/cactus/go-statsd-client/v5/statsd"
)

func main() {

	startTime := time.Now()

	plainText := "The quick brown fox jumps over the lazy dog. Sphinx of black quartz, judge my vow. A wizard's job is to vex chumps quickly in fog"
	key := "0123456789abcdef"

	cipher, _ := aes.FromString(key)

	cipherText := cipher.CipherBlock(aes.StringToBytes(plainText))

	decipheredText := cipher.InvCipherBlock(cipherText)

	println("Plain text: ", plainText)
	println("Deciphered text: ", string(decipheredText[:]))
	println("Cipher text: ", string(cipherText[:]))

	endTime := time.Now()
	elapsedTime := endTime.Sub(startTime)
	sendTime(elapsedTime)

}		

func sendTime(time time.Duration) {
	println("Time: ", time)

	statsdDir := "graphite:8125"

	statsdClient, err := statsd.NewClient(statsdDir, "aes_go")

	if err != nil {
		println("Error: ", err)
		return
	}

	statsdClient.Inc("aes_go_metric", 1, 1)
	
}
