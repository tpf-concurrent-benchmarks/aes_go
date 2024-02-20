package matrix

import (
	"testing"
)

var testData [4][4]byte = [4][4]byte{
	{0, 1, 2, 3},
	{4, 5, 6, 7},
	{8, 9, 10, 11},
	{12, 13, 14, 15},
}

func TestNew(t *testing.T) {
	matrix := New()
	if matrix == nil {
		t.Error("Expected matrix to be created")
	}
}

func TestGetDimensions(t *testing.T) {
	matrix := New()
	rows := matrix.GetRowsAmount()
	cols := matrix.GetColsAmount()
	if rows != 4 || cols != 4 {
		t.Error("Expected 4x4 matrix")
	}
}

func TestNewValues(t *testing.T) {
	matrix := New()
	for i := 0; i < matrix.GetRowsAmount(); i++ {
		for j := 0; j < matrix.GetColsAmount(); j++ {
			if matrix.Get(i, j) != 0 {
				t.Error("Expected 0 value")
			}
		}
	}
}

func TestGetValue(t *testing.T) {
	matrix := FromData(testData)
	val := matrix.Get(0,0)
	if val != 0 {
		t.Error("Expected 0 value")
	}

	val = matrix.Get(1,1)
	if val != 5 {
		t.Error("Expected 10 value")
	}

	val = matrix.Get(2,2)
	if val != 10 {
		t.Error("Expected 10 value")
	}

	val = matrix.Get(3,3)
	if val != 15 {
		t.Error("Expected 15 value")
	}
}

func TestGetCol(t *testing.T) {
	matrix := FromData(testData)
	col := matrix.GetCol(0)
	if col[0] != 0 || col[1] != 4 || col[2] != 8 || col[3] != 12 {
		t.Error("Expected 0, 4, 8, 12")
	}

	col = matrix.GetCol(1)
	if col[0] != 1 || col[1] != 5 || col[2] != 9 || col[3] != 13 {
		t.Error("Expected 1, 5, 9, 13")
	}

	col = matrix.GetCol(2)
	if col[0] != 2 || col[1] != 6 || col[2] != 10 || col[3] != 14 {
		t.Error("Expected 2, 6, 10, 14")
	}

	col = matrix.GetCol(3)
	if col[0] != 3 || col[1] != 7 || col[2] != 11 || col[3] != 15 {
		t.Error("Expected 3, 7, 11, 15")
	}
}

func TestSetCol(t *testing.T) {
	matrix := New()
	col := [4]byte{0, 4, 8, 12}
	matrix.SetCol(0, col)

	mCol := matrix.GetCol(0)
	if mCol[0] != 0 || mCol[1] != 4 || mCol[2] != 8 || mCol[3] != 12 {
		t.Error("Expected 0, 4, 8, 12")
	}
}

func TestGetCols(t *testing.T) {
	matrix := FromData(testData)
	cols := matrix.GetCols()
	if cols[0][0] != 0 || cols[0][1] != 4 || cols[0][2] != 8 || cols[0][3] != 12 {
		t.Error("Expected 0, 4, 8, 12")
	}

	if cols[1][0] != 1 || cols[1][1] != 5 || cols[1][2] != 9 || cols[1][3] != 13 {
		t.Error("Expected 1, 5, 9, 13")
	}

	if cols[2][0] != 2 || cols[2][1] != 6 || cols[2][2] != 10 || cols[2][3] != 14 {
		t.Error("Expected 2, 6, 10, 14")
	}

	if cols[3][0] != 3 || cols[3][1] != 7 || cols[3][2] != 11 || cols[3][3] != 15 {
		t.Error("Expected 3, 7, 11, 15")
	}
}

func TestShiftLeft(t *testing.T){
	matrix := FromData(testData)
	matrix.ShiftRowLeft(0, 1)
	row := matrix.GetRow(0)
	if row[0] != 1 || row[1] != 2 || row[2] != 3 || row[3] != 0 {
		t.Error("Expected 1, 2, 3, 0")
	}

	matrix.ShiftRowLeft(1, 2)
	row = matrix.GetRow(1)
	if row[0] != 6 || row[1] != 7 || row[2] != 4 || row[3] != 5 {
		t.Error("Expected 6, 7, 4, 5")
	}

	matrix.ShiftRowLeft(2, 3)
	row = matrix.GetRow(2)
	if row[0] != 11 || row[1] != 8 || row[2] != 9 || row[3] != 10 {
		t.Error("Expected 11, 8, 9, 10")
	}
}

func TestShiftRight(t *testing.T){
	matrix := FromData(testData)
	matrix.ShiftRowRight(0, 1)
	row := matrix.GetRow(0)
	if row[0] != 3 || row[1] != 0 || row[2] != 1 || row[3] != 2 {
		t.Error("Expected 3, 0, 1, 2")
	}

	matrix.ShiftRowRight(1, 2)
	row = matrix.GetRow(1)
	if row[0] != 6 || row[1] != 7 || row[2] != 4 || row[3] != 5 {
		t.Error("Expected 6, 7, 4, 5")
	}

	matrix.ShiftRowRight(2, 3)
	row = matrix.GetRow(2)
	if row[0] != 9 || row[1] != 10 || row[2] != 11 || row[3] != 8 {
		t.Error("Expected 9, 10, 11, 8")
	}
}