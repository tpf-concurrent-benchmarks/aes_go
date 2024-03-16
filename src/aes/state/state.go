package state

import (
	. "aes_go/aes/constants"
	"aes_go/matrix"
)

type State struct {
	data matrix.Matrix
}

func New() (state *State) {
	return &State{data: *matrix.New()}
}

func FromMatrix(data matrix.Matrix) (state *State) {
	return &State{data: data}
}

func FromData(data [4][4]byte) (state *State) {
	return &State{data: *matrix.FromData(data)}
}

func FromList(dataIn [4 * N_B]byte) (state *State) { // set_data_in
	state = New()
	for i := 0; i < int(N_B); i++ {
		col := [4]byte{dataIn[4*i], dataIn[4*i+1], dataIn[4*i+2], dataIn[4*i+3]}
		state.data.SetCol(i, col)
	}
	return state
}

func FromWords(words [N_B]Word) (state *State) {
	state = New()
	for i := 0; i < int(N_B); i++ {
		word := words[i]
		wordBytes := word.ToBytes()
		col := [4]byte{wordBytes[0], wordBytes[1], wordBytes[2], wordBytes[3]}
		state.data.SetCol(i, col)
	}
	return state
}

func (state *State) ToList(dataOut *[4 * N_B]byte) { // new_from_data_in
	for i := 0; i < int(N_B); i++ {
		col := state.data.GetCol(i)
		dataOut[4*i] = col[0]
		dataOut[4*i+1] = col[1]
		dataOut[4*i+2] = col[2]
		dataOut[4*i+3] = col[3]
	}
}

func (state *State) SubBytes() {
	state.applySubstitution(S_BOX)
}

func (state *State) InvSubBytes() {
	state.applySubstitution(INV_S_BOX)
}

func (state *State) applySubstitution(subBox [256]byte) {
	for row := 0; row < int(state.data.GetRowsAmount()); row++ {
		for col := 0; col < int(state.data.GetColsAmount()); col++ {
			value := state.data.Get(row, col)
			state.data.Set(row, col, subBox[value])
		}
	}
}

func (state *State) ShiftRows() {
	for i := 1; i < int(state.data.GetRowsAmount()); i++ {
		state.data.ShiftRowLeft(i, i)
	}
}

func (state *State) InvShiftRows() {
	for i := 1; i < int(state.data.GetRowsAmount()); i++ {
		state.data.ShiftRowRight(i, i)
	}
}

func mixColumn(col [4]byte) [4]byte {
	a := col[0]
	b := col[1]
	c := col[2]
	d := col[3]

	newCol := [4]byte{
		galoisDouble(a^b) ^ b ^ c ^ d,
		galoisDouble(b^c) ^ c ^ d ^ a,
		galoisDouble(c^d) ^ d ^ a ^ b,
		galoisDouble(d^a) ^ a ^ b ^ c,
	}
	return newCol
}

func (state *State) MixColumns() {
	for i := 0; i < int(N_B); i++ {
		col := state.data.GetCol(i)
		newCol := mixColumn(col)
		state.data.SetCol(i, newCol)
	}
}

func invMixColumn(col [4]byte) [4]byte {
	a := col[0]
	b := col[1]
	c := col[2]
	d := col[3]

	x := galoisDouble(a^b^c^d)
	y := galoisDouble(x^a^c)
	z := galoisDouble(x^b^d)

	newCol := [4]byte{
		galoisDouble(y^a^b) ^ b ^ c ^ d,
		galoisDouble(z^b^c) ^ c ^ d ^ a,
		galoisDouble(y^c^d) ^ d ^ a ^ b,
		galoisDouble(z^d^a) ^ a ^ b ^ c,
	}
	return newCol
}

func (state *State) InvMixColumns() {
	for i := 0; i < int(N_B); i++ {
		col := state.data.GetCol(i)
		newCol := invMixColumn(col)
		state.data.SetCol(i, newCol)
	}
}

func (state *State) AddRoundKey(roundKey [N_B]Word) {
	for i := 0; i < int(N_B); i++ {
		col := state.data.GetCol(i)
		word := roundKey[i]
		wordBytes := word.ToBytes()
		newCol := [4]byte{
			col[0] ^ wordBytes[0],
			col[1] ^ wordBytes[1],
			col[2] ^ wordBytes[2],
			col[3] ^ wordBytes[3],
		}
		state.data.SetCol(i, newCol)
	}
}

// Deprecated - all galoisMul calls were replaced optimized implementations using galoisDouble
// func galoisMul(a byte, b byte) byte {
// 	var result byte = 0
// 	for b != 0 {
// 		if b&1 != 0 {
// 			result ^= a
// 		}
// 		if a&0x80 != 0 {
// 			a = (a << 1) ^ 0x1b
// 		} else {
// 			a <<= 1
// 		}
// 		b >>= 1
// 	}
// 	return result
// }

func galoisDouble(a byte) byte {
	if a&0x80 != 0 {
		return (a << 1) ^ 0x1b
	}
	return a << 1
}

func (state *State) GetRow(row int) [4]byte {
	return state.data.GetRow(row)
}

func (state *State) GetCol(col int) [4]byte {
	return state.data.GetCol(col)
}
