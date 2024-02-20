package cipher

import (
	. "aes_go/aes/constants"
	"aes_go/aes/key"
	"aes_go/aes/state"
	"errors"
)

type AESCipher struct {
	expandedKey key.Key
	invExpandedKey key.Key
}

func New(cipherKey [4 * N_B]byte) (cipher *AESCipher) {
	expandedKey := key.NewDirect(cipherKey)
	invExpandedKey := key.NewInverse(cipherKey)
	return &AESCipher{expandedKey: *expandedKey, invExpandedKey: *invExpandedKey}
}

func FromString(cipherKey string) (cipher *AESCipher, err error) {

	keyBytes := [4*N_B]byte( []byte( cipherKey ) )
	copy(keyBytes[:], []byte(cipherKey))
	if len(cipherKey) != 4 * int(N_B) {
		return New(keyBytes), errors.New("invalid key length")
	}
	return New(keyBytes), nil
}

func StringToBytes(s string) (bytes [4 * N_B]byte) {
	copy(bytes[:], []byte(s))
	return bytes
}

func BytesToString(bytes [4 * N_B]byte) (s string) {
	return string(bytes[:])
}

func (cipher *AESCipher) CipherBlock(dataIn [4 * N_B]byte) (dataOut [4 * N_B]byte) {
	state := state.FromList(dataIn)
	state.AddRoundKey([4]Word(cipher.expandedKey.Data[:N_B]))
	for round := 1; round < int(N_R); round++ {
		state.SubBytes()
		state.ShiftRows()
		state.MixColumns()
		state.AddRoundKey([4]Word(cipher.expandedKey.Data[round*int(N_B):(round+1)*int(N_B)]))
	}
	state.SubBytes()
	state.ShiftRows()
	state.AddRoundKey([4]Word(cipher.expandedKey.Data[int(N_R)*int(N_B):(int(N_R)+1)*int(N_B)]))
	state.ToList(&dataOut)
	return dataOut
}

func (cipher *AESCipher) InvCipherBlock(dataIn [4 * N_B]byte) (dataOut [4 * N_B]byte) {
	state := state.FromList(dataIn)
	state.AddRoundKey([4]Word(cipher.invExpandedKey.Data[int(N_R*N_B):int((N_R+1)*N_B)]))
	
	for round := int(N_R) - 1; round > 0; round-- {
		state.InvSubBytes()
		state.InvShiftRows()
		state.InvMixColumns()
		state.AddRoundKey([4]Word(cipher.invExpandedKey.Data[round*int(N_B):(round+1)*int(N_B)]))
	}

	state.InvSubBytes()
	state.InvShiftRows()
	state.AddRoundKey([4]Word(cipher.invExpandedKey.Data[:N_B]))
	state.ToList(&dataOut)
	return dataOut
}
