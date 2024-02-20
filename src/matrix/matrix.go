package matrix

type Matrix struct {
	data [4][4]byte // Hardcoded 4x4 matrix for AES-128
}

func New() (matrix *Matrix) {
	return &Matrix{}
}

func FromData(data [4][4]byte) (matrix *Matrix) {
	return &Matrix{data: data}
}

func (m *Matrix) Get(row int, col int) byte {
	return m.data[row][col]
}

func (m *Matrix) Set(row int, col int, value byte) {
	m.data[row][col] = value
}

func (m *Matrix) GetRowsAmount() int {
	return 4
}

func (m *Matrix) GetColsAmount() int {
	return 4
}

func (m *Matrix) GetCols() (cols [4][4]byte) {
	for i := 0; i < m.GetColsAmount(); i++ {
		cols[i] = m.GetCol(i)
	}
	return cols
}

func (m *Matrix) GetCol(col int) (colData [4]byte) {
	for i := 0; i < m.GetRowsAmount(); i++ {
		colData[i] = m.data[i][col]
	}
	return colData
}

func (m *Matrix) SetCol(col int, data [4]byte) {
	for i := 0; i < m.GetRowsAmount(); i++ {
		m.data[i][col] = data[i]
	}
}

func (m *Matrix) GetRow(row int) (rowData [4]byte) {
	rowData = m.data[row]
	return rowData
}

func (m *Matrix) ShiftRowLeft(row int, amount int) {
	for i := 0; i < amount; i++ {
		temp := m.data[row][0]
		for j := 0; j < m.GetColsAmount()-1; j++ {
			m.data[row][j] = m.data[row][j+1]
		}
		m.data[row][m.GetColsAmount()-1] = temp
	}
}

func (m *Matrix) ShiftRowRight(row int, amount int) {
	for i := 0; i < amount; i++ {
		temp := m.data[row][m.GetColsAmount()-1]
		for j := m.GetColsAmount() - 1; j > 0; j-- {
			m.data[row][j] = m.data[row][j-1]
		}
		m.data[row][0] = temp
	}
}
