package cipher

import (
	. "aes_go/aes/constants"
	"aes_go/aes/state"
	"testing"
)

func TestNew(t *testing.T) {
	cipherKey := [4 * N_B]byte{}
	cipher := New(cipherKey)
	if cipher == nil {
		t.Errorf("New() = nil, want *AESCipher")
	}
}

func TestFromString(t *testing.T) {
	cipherKey := "0123456789abcdef"
	cipher, err := FromString(cipherKey)
	if err != nil {
		t.Errorf("FromString() error = %v, want nil", err)
	}
	if cipher == nil {
		t.Errorf("FromString() = nil, want *AESCipher")
	}
}

func TestShiftRows( t*testing.T) {
	initialState := state.FromData([4][4]byte{
		{0xd4, 0xe0, 0xb8, 0x1e},
		{0xbf, 0xb4, 0x41, 0x27},
		{0x5d, 0x52, 0x11, 0x98},
		{0x30, 0xae, 0xf1, 0xe5},
	})

	expectedState := state.FromData([4][4]byte{
		{0xd4, 0xe0, 0xb8, 0x1e},
		{0xb4, 0x41, 0x27, 0xbf},
		{0x11, 0x98, 0x5d, 0x52},
		{0xe5, 0x30, 0xae, 0xf1},
	})

	initialState.ShiftRows()

	for i := 0; i < 4; i++ {
		if initialState.GetRow(i) != expectedState.GetRow(i) {
			t.Errorf("ShiftRows() = %v, want %v", initialState.GetRow(i), expectedState.GetRow(i))
		}
	}
}

func TestInvShiftRows( t*testing.T) {

	initialState := state.FromData([4][4]byte{
		{0xd4, 0xe0, 0xb8, 0x1e},
		{0xb4, 0x41, 0x27, 0xbf},
		{0x11, 0x98, 0x5d, 0x52},
		{0xe5, 0x30, 0xae, 0xf1},
	})

	expectedState := state.FromData([4][4]byte{
		{0xd4, 0xe0, 0xb8, 0x1e},
		{0xbf, 0xb4, 0x41, 0x27},
		{0x5d, 0x52, 0x11, 0x98},
		{0x30, 0xae, 0xf1, 0xe5},
	})

	initialState.InvShiftRows()

	for i := 0; i < 4; i++ {
		if initialState.GetRow(i) != expectedState.GetRow(i) {
			t.Errorf("InvShiftRows() = %v, want %v", initialState.GetRow(i), expectedState.GetRow(i))
		}
	}
}

func TestSubBytes( t*testing.T) {
	initialState := state.FromData([4][4]byte{
		{0x19, 0xa0, 0x9a, 0xe9},
		{0x3d, 0xf4, 0xc6, 0xf8},
		{0xe3, 0xe2, 0x8d, 0x48},
		{0xbe, 0x2b, 0x2a, 0x08},
	})

	expectedState := state.FromData([4][4]byte{
		{0xd4, 0xe0, 0xb8, 0x1e},
		{0x27, 0xbf, 0xb4, 0x41},
		{0x11, 0x98, 0x5d, 0x52},
		{0xae, 0xf1, 0xe5, 0x30},
	})

	initialState.SubBytes()

	for i := 0; i < 4; i++ {
		if initialState.GetRow(i) != expectedState.GetRow(i) {
			t.Errorf("SubBytes() = %v, want %v", initialState.GetRow(i), expectedState.GetRow(i))
		}
	}
}

func TestInvSubBytes( t*testing.T) {
	initialState := state.FromData([4][4]byte{
		{0xd4, 0xe0, 0xb8, 0x1e},
		{0x27, 0xbf, 0xb4, 0x41},
		{0x11, 0x98, 0x5d, 0x52},
		{0xae, 0xf1, 0xe5, 0x30},
	})

	expectedState := state.FromData([4][4]byte{
		{0x19, 0xa0, 0x9a, 0xe9},
		{0x3d, 0xf4, 0xc6, 0xf8},
		{0xe3, 0xe2, 0x8d, 0x48},
		{0xbe, 0x2b, 0x2a, 0x08},
	})

	initialState.InvSubBytes()

	for i := 0; i < 4; i++ {
		if initialState.GetRow(i) != expectedState.GetRow(i) {
			t.Errorf("InvSubBytes() = %v, want %v", initialState.GetRow(i), expectedState.GetRow(i))
		}
	}
}

func TestGetStateFromDataIn( t*testing.T ){ 
	dataIn := [4 * N_B]byte{
		0x32, 0x88, 0x31, 0xe0, 0x43, 0x5a, 0x31, 0x37, 0xf6, 0x30, 0x98, 0x07, 0xa8, 0x8d, 0xa2, 0x34,
	}

	expectedState := state.FromData([4][4]byte{
		{0x32, 0x43, 0xf6, 0xa8},
		{0x88, 0x5a, 0x30, 0x8d},
		{0x31, 0x31, 0x98, 0xa2},
		{0xe0, 0x37, 0x07, 0x34},
	})

	state := state.FromList(dataIn)

	for i := 0; i < 4; i++ {
		if state.GetRow(i) != expectedState.GetRow(i) {
			t.Errorf("GetStateFromDataIn() = %v, want %v", state.GetRow(i), expectedState.GetRow(i))
		}
	}
}

func TestSetDataOutFromState( t*testing.T ){
	dataOut := [4 * N_B]byte{}

	state := state.FromData([4][4]byte{
		{0x39, 0x02, 0xdc, 0x19},
		{0x25, 0xdc, 0x11, 0x6a},
		{0x84, 0x09, 0x85, 0x0b},
		{0x1d, 0xfb, 0x97, 0x32},
	})

	expectedDataOut := [4 * N_B]byte{
		0x39, 0x25, 0x84, 0x1d, 0x02, 0xdc, 0x09, 0xfb, 0xdc, 0x11, 0x85, 0x97, 0x19, 0x6a, 0x0b, 0x32,
	}

	state.ToList(&dataOut)

	for i := 0; i < 4 * int(N_B); i++ {
		if dataOut[i] != expectedDataOut[i] {
			t.Errorf("SetDataOutFromState() = %v, want %v", dataOut[i], expectedDataOut[i])
		}
	}
}

func TestMixColumns( t*testing.T ){
	initialState := state.FromData([4][4]byte{
		{0xdb, 0xf2, 0x01, 0xc6},
		{0x13, 0x0a, 0x01, 0xc6},
		{0x53, 0x22, 0x01, 0xc6},
		{0x45, 0x5c, 0x01, 0xc6},
	})

	expectedState := state.FromData([4][4]byte{
		{0x8e, 0x9f, 0x01, 0xc6},
		{0x4d, 0xdc, 0x01, 0xc6},
		{0xa1, 0x58, 0x01, 0xc6},
		{0xbc, 0x9d, 0x01, 0xc6},
	})

	initialState.MixColumns()

	for i := 0; i < 4; i++ {
		if initialState.GetRow(i) != expectedState.GetRow(i) {
			t.Errorf("MixColumns() = %v, want %v", initialState.GetRow(i), expectedState.GetRow(i))
		}
	}
}

func TestCipher( t*testing.T ){ 
	plainBytes := [4 * N_B]byte{
		0x32, 0x43, 0xf6, 0xa8, 0x88, 0x5a, 0x30, 0x8d, 0x31, 0x31, 0x98, 0xa2, 0xe0, 0x37, 0x07, 0x34,
	}

	cipherKey := [4 * N_K]byte{
		0x2b, 0x7e, 0x15, 0x16, 0x28, 0xae, 0xd2, 0xa6, 0xab, 0xf7, 0x15, 0x88, 0x09, 0xcf, 0x4f, 0x3c,
	}

	expectedCipherBytes := [4 * N_B]byte{
		0x39, 0x25, 0x84, 0x1d, 0x02, 0xdc, 0x09, 0xfb, 0xdc, 0x11, 0x85, 0x97, 0x19, 0x6a, 0x0b, 0x32,
	}

	cipher := New(cipherKey)

	cipherBytes := cipher.CipherBlock(plainBytes)

	for i := 0; i < 4 * int(N_B); i++ {
		if cipherBytes[i] != expectedCipherBytes[i] {
			t.Errorf("Cipher() = %v, want %v", cipherBytes[i], expectedCipherBytes[i])
		}
	}
}

func TestInvMixColumns( t*testing.T ){
	initialState := state.FromData([4][4]byte{
		{0x8e, 0x9f, 0x01, 0xc6},
		{0x4d, 0xdc, 0x01, 0xc6},
		{0xa1, 0x58, 0x01, 0xc6},
		{0xbc, 0x9d, 0x01, 0xc6},
	})

	expectedState := state.FromData([4][4]byte{
		{0xdb, 0xf2, 0x01, 0xc6},
		{0x13, 0x0a, 0x01, 0xc6},
		{0x53, 0x22, 0x01, 0xc6},
		{0x45, 0x5c, 0x01, 0xc6},
	})

	initialState.InvMixColumns()

	for i := 0; i < 4; i++ {
		if initialState.GetRow(i) != expectedState.GetRow(i) {
			t.Errorf("InvMixColumns() = %v, want %v", initialState.GetRow(i), expectedState.GetRow(i))
		}
	}
}

func TestInvCipher( t*testing.T ){
	cipherKey := [4 * N_K]byte{
		0x2b, 0x7e, 0x15, 0x16, 0x28, 0xae, 0xd2, 0xa6, 0xab, 0xf7, 0x15, 0x88, 0x09, 0xcf, 0x4f, 0x3c,
	}

	cipherBytes := [4 * N_B]byte{
		0x39, 0x25, 0x84, 0x1d, 0x02, 0xdc, 0x09, 0xfb, 0xdc, 0x11, 0x85, 0x97, 0x19, 0x6a, 0x0b, 0x32,
	}

	expectedPlainBytes := [4 * N_B]byte{
		0x32, 0x43, 0xf6, 0xa8, 0x88, 0x5a, 0x30, 0x8d, 0x31, 0x31, 0x98, 0xa2, 0xe0, 0x37, 0x07, 0x34,
	}

	cipher := New(cipherKey)

	plainBytes := cipher.InvCipherBlock(cipherBytes)

	for i := 0; i < 4 * int(N_B); i++ {
		if plainBytes[i] != expectedPlainBytes[i] {
			t.Errorf("InvCipher() = %v, want %v", plainBytes[i], expectedPlainBytes[i])
		}
	}
}

// func TestCipherInvCipher( t*testing.T ){ 
// 	plainBytes := [4 * N_B]byte{
// 		0x32, 0x43, 0xf6, 0xa8, 0x88, 0x5a, 0x30, 0x8d, 0x31, 0x31, 0x98, 0xa2, 0xe0, 0x37, 0x07, 0x34,
// 	}

// 	cipherKey := [4 * N_K]byte{
// 		0x2b, 0x7e, 0x15, 0x16, 0x28, 0xae, 0xd2, 0xa6, 0xab, 0xf7, 0x15, 0x88, 0x09, 0xcf, 0x4f, 0x3c,
// 	}

// 	cipher := New(cipherKey)

// 	cipherBytes := cipher.CipherBlock(plainBytes)

// 	decipheredBytes := cipher.InvCipherBlock(cipherBytes)

// 	for i := 0; i < 4 * int(N_B); i++ {
// 		if decipheredBytes[i] != plainBytes[i] {
// 			t.Errorf("CipherInvCipher() = %v, want %v", decipheredBytes[i], plainBytes[i])
// 		}
// 	}
// }


// func TestCipherBlock(t *testing.T) {
// 	cipherKey := "0123456789abcdef"
// 	cipher, err := FromString(cipherKey)
// 	if err != nil {
// 		t.Errorf("FromString() error = %v, want nil", err)
// 	}

// 	dataIn := StringToBytes("0123456789abcdef")
// 	dataOut := cipher.CipherBlock(dataIn)

// 	decipheredData := cipher.InvCipherBlock(dataOut)

// 	// print dataOut
// 	fmt.Println(BytesToString(decipheredData))
// }

// #[test]
// fn test_cipher_using_new_u128() {
//     let plain_bytes: [u8; 4 * N_B] = [
//         0x32, 0x43, 0xf6, 0xa8, 0x88, 0x5a, 0x30, 0x8d, 0x31, 0x31, 0x98, 0xa2, 0xe0, 0x37, 0x07,
//         0x34,
//     ];

//     let cipher_key: u128 = 0x2b7e151628aed2a6abf7158809cf4f3c;

//     let expected_cipher_bytes: [u8; 4 * N_B] = [
//         0x39, 0x25, 0x84, 0x1d, 0x02, 0xdc, 0x09, 0xfb, 0xdc, 0x11, 0x85, 0x97, 0x19, 0x6a, 0x0b,
//         0x32,
//     ];

//     let cipher = AESCipher::new_u128(cipher_key);

//     let block = cipher.cipher_block(plain_bytes);

//     for i in 0..(N_B * 4) {
//         assert_eq!(block[i], expected_cipher_bytes[i]);
//     }
// }