package components

import (
	aes "aes_go/aes"
	"os"
)

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

func _removeTrailingNulls(f *os.File) {
	// Read last block,

	lastOffest, err := f.Seek(-int64(aes.BlockSize), os.SEEK_END)
	Check(err)

	buff := make([]byte, aes.BlockSize)

	_, err = f.Read(buff)
	Check(err)

	// Calc valid chars

	validChars := 0
	for i := 0; i < int(aes.BlockSize); i++ {
		if buff[i] != 0 {
			validChars = i + 1
		}
	}

	// Truncate file

	f.Truncate(lastOffest + int64(validChars))

}

func RemoveTrailingNulls(filename string) {
	f, err := os.OpenFile(filename, os.O_RDWR, 0644)
	Check(err)
	defer f.Close()

	_removeTrailingNulls(f)
}