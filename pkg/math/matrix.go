package math

type Matrix struct {
	M    []int
	Rows int
	Cols int
}

// NewMatrix returns a blank matrix with the given size.
func NewMatrix(rows, cols int) *Matrix {
	return &Matrix{
		M:    make([]int, rows*cols),
		Rows: rows,
		Cols: cols,
	}
}

// loc gets the location of row, col in the 1D array.
func (m *Matrix) loc(row, col int) int {
	return row*m.Cols + col
}

// Get gets the value of the Grid at the specified location.
func (m *Matrix) Get(row, col int) int {
	return m.M[m.loc(row, col)]
}

// Set sets the value of the Grid at the specified location to the given int.
func (m *Matrix) Set(row, col, val int) {
	m.M[m.loc(row, col)] = val
}

// GetCol gets the given column.
func (m *Matrix) GetCol(col int) []int {
	vals := make([]int, m.Rows)
	for row := range vals {
		vals[i] = m.Get(row, col)
	}
	return vals
}

// SetCol sets the given column to the given values.
func (m *Matrix) SetCol(col int, vals []int) {
	for row, val := range vals {
		m.Set(row, col, val)
	}
}

// GetRow gets the given row.
func (m *Matrix) GetRow(row int) []int {
	vals := make([]int, m.Cols)
	for col := range vals {
		vals[i] = m.Get(row, col)
	}
	return vals
}

// SetRow sets the given row to the given values.
func (m *Matrix) SetRow(row int, vals []int) {
	for col, val := range vals {
		m.Set(row, col, val)
	}
}

// SwapRows swaps the two given rows.
func (m *Matrix) SwapRows(row1, row2 int) {
	from := m.GetRow(row1)
	m.SetRow(row1, m.GetRow(row2))
	m.SetRow(row2, from)
}

// ScaleRow scales the given row by the given scale factor.
func (m *Matrix) ScaleRow(row, scale int) {
	curr := m.GetRow(row)
	for col, val := range curr {
		m.Set(row, col, val*scale)
	}
}

// AddRow adds the scaled value of the `from` row to the `to` row.
func (m *Matrix) AddRow(from, to, scale int) {
	start := m.GetRow(from)
	for col, val := range start {
		m.Set(to, col, m.Get(to, col)+val*scale)
	}
}

// Solve reduces the matrix to row echelon form, transforming the right hand side and returning it.
func (m *Matrix) RowEchelon(rhs []int) []int {
	res := make([]int, len(rhs))
	copy(res, rhs)
	for i := 0; i < m.Rows; i++ {
		if m.Get(i, i) == 0 {
			for swapRow := i + 1; m.Get(i, i) == 0 && swapRow < m.Rows; swapRow++ {
				m.SwapRows(i, swapRow)
				res[i], res[swapRow] = res[swapRow], res[i]
			}
		}
		for j := 0; j < m.Rows; j++ {
			if i != j {
				m.ScaleRow(j, m.Get(i, i))
				res[j] *= m.Get(i, i)
				m.AddRow(i, j, -m.Get(j, i))
				res[j] -= res[i] * m.Get(j, i)
			}
		}
	}
	return res
}
