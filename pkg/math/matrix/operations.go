package matrix

import (
	"errors"
	"fmt"
)

func (m Matrix[T]) Get(row int, col int) (*T, error) {
	if m.transposed {
		row, col = col, row
	}

	if row < 0 || row >= m.row ||
		col < 0 || col >= m.col {
		//charm.Debug("", "r", row, "c", col, "x", m.row, "y", m.col)
		return nil, errors.New("index out of bounds")
	}

	return &m.mat[row][col], nil
}

type MapFunc[A Numeric, B Numeric] func(in A) (*B, error)

func MapElementWise[A Numeric, B Numeric](m *Matrix[A], f MapFunc[A, B]) (
	*Matrix[B], error) {

	mNew, err := New[B](m.row, m.col)
	if err != nil {
		return nil, err
	}
	mNew.transposed = m.transposed

	for i := range m.row {
		for j := range m.col {
			v, err := f(m.mat[i][j])
			if err != nil {
				return nil, err
			}

			mNew.mat[i][j] = *v
		}
	}

	return mNew, nil
}

func Add[A Numeric, B Numeric, C Numeric](m1 *Matrix[A], m2 *Matrix[B]) (
	*Matrix[C], error) {
	row1, col1 := m1.GetDimensions()
	row2, col2 := m2.GetDimensions()

	if row1 != row2 || col1 != col2 {
		return nil, fmt.Errorf("invalid dimension for matrix addition, m1: { r: %d, c: %d }, m2: { r: %d, c: %d }",
			row1, col1, row2, col2)
	}

	mNew, err := New[C](row1, col1)
	if err != nil {
		return nil, err
	}

	return mNew, nil
}
