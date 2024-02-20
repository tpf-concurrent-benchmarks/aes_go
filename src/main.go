package main

import (
	"aes_go/aes"
)

func main() {

	plainText := "The quick brown fox jumps over the lazy dog. Sphinx of black quartz, judge my vow. A wizard's job is to vex chumps quickly in fog"
	key := "0123456789abcdef"

	cipher := aes.FromString(key)

	cipherText := cipher.EncryptString(plainText)

	decipheredText := cipher.Decrypt(cipherText)

	println("Plain text: ", plainText)
	println("Deciphered text: ", string(decipheredText[:]))
	println("Cipher text: ", string(cipherText[:]))

}		
