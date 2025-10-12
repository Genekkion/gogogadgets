package matrix

import (
	"errors"
	"fmt"
	"strings"

	"golang.org/x/exp/constraints"
)

type Numeric interface {
	constraints.Integer | constraints.Float
}

type Matrix[T Numeric] struct {
	transposed bool
	mat        [][]T
	row        int
	col        int
}

func New[T Numeric](row int, col int) (*Matrix[T], error) {
	if row == 0 || col == 0 {
		return nil, fmt.Errorf("invalid matrix dimensions, r: %d, c: %d", row, col)
	}

	mat := make([][]T, row)
	for i := range row {
		mat[i] = make([]T, col)
	}

	return &Matrix[T]{
		transposed: false,
		mat:        mat,
		row:        row,
		col:        col,
	}, nil
}

func NewFromSlices[T Numeric](original [][]T) (*Matrix[T], error) {
	if len(original) == 0 {
		return nil, errors.New("invalid matrix dimensions")
	}

	m, err := New[T](len(original), len(original[0]))
	if err != nil {
		return nil, err
	}

	for i, row := range original {
		col := len(row)
		if col != m.col {
			return nil, fmt.Errorf("inconsistent column count, expected: %d, got: %d",
				m.col, col)
		}

		copy(m.mat[i], row)
	}

	return m, nil
}

func (m *Matrix[T]) Transpose() bool {
	m.transposed = !m.transposed
	return m.transposed
}

func (m Matrix[T]) String() string {
	var row int
	var col int
	if !m.transposed {
		row = m.row
		col = m.col
	} else {
		row = m.col
		col = m.row
	}

	return fmt.Sprintf("Matrix: { row: %d, col: %d }", row, col)
}

func (m Matrix[T]) Display() string {
	b := strings.Builder{}
	b.WriteString(m.String())
	b.WriteRune('\n')

	if !m.transposed {
		for i := range m.row {
			if i > 0 {
				b.WriteRune('\n')
			}
			b.WriteRune('[')
			for j := range m.col {
				if j > 0 {
					b.WriteString(", ")
				}
				b.WriteString(fmt.Sprintf("%v", m.mat[i][j]))
			}
			b.WriteRune(']')
		}
	} else {
		for i := range m.col {
			if i > 0 {
				b.WriteRune('\n')
			}
			b.WriteRune('[')
			for j := range m.row {
				if j > 0 {
					b.WriteString(", ")
				}
				b.WriteString(fmt.Sprintf("%v", m.mat[j][i]))
			}
			b.WriteRune(']')
		}
	}

	return b.String()
}

func (m Matrix[T]) GetDimensions() (int, int) {
	if !m.transposed {
		return m.row, m.col
	} else {
		return m.col, m.row
	}
}
