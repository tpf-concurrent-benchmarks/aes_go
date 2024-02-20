package key

import (
	. "aes_go/aes/constants"
	"aes_go/aes/state"
)

type Key struct {
	Data [N_B * (N_R + 1)]Word
}

func NewDirect(cipherKey [4 * N_K]byte) (key *Key) {
	data := expandKey(cipherKey)
	return &Key{Data: data}
}

func NewInverse(cipherKey [4 * N_K]byte) (key *Key) {
	data := invExpandKey(cipherKey)
	return &Key{Data: data}
}

func expandKey(cipherKey [4 * N_K]byte) (data [N_B * (N_R + 1)]Word) {
	var temp Word
	i := 0
	for i < int(N_K) {
		data[i] = Word(cipherKey[4*i])<<24 | Word(cipherKey[4*i+1])<<16 | Word(cipherKey[4*i+2])<<8 | Word(cipherKey[4*i+3])
		i++
	}
	i = int(N_K)
	for i < int(N_B*(N_R+1)) {
		temp = data[i-1]
		if i%int(N_K) == 0 {
			temp = subWord(rotWord(temp)) ^ RCON[i/int(N_K)-1]
		}
		data[i] = data[i-int(N_K)] ^ temp
		i++
	}
	return data
}

func invExpandKey(cipherKey [4 * N_K]byte) (data [N_B * (N_R + 1)]Word) {
	data = expandKey(cipherKey)
	for round := 1; round < int(N_R); round++ {
		newWords := invMixColumnsWords([N_B]Word(data[round*int(N_B) : (round+1)*int(N_B)]))
		for i := 0; i < int(N_B); i++ {
			data[round*int(N_B)+i] = newWords[i]
		}
	}
	return data
}

func subWord(word Word) (result Word) {
	for i := 0; i < 4; i++ {
		byte := byte(word >> (8 * i))
		newByte := applySBox(byte)
		result |= Word(newByte) << (8 * i)
	}
	return result
}

func rotWord(word Word) Word {
	return word<<8 | word>>24
}

func invMixColumnsWords(words [N_B]Word) (newWords [N_B]Word) {
	state := state.FromWords(words)
	state.InvMixColumns()
	for i := 0; i < int(N_B); i++ {
		newWords[i] = WordFromBytes(state.GetCol(i))
	}
	return newWords
}

func applySBox(value byte) byte {
	posX := int(value >> 4)
	posY := int(value & 0x0f)
	return S_BOX[posX*16+posY]
}
