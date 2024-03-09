package combinatorics

import (
	"fmt"
	"strings"
)

const SUDOKU_SIZE = 9
const SUDOKU_BLOCK_SIZE = SUDOKU_SIZE / 3
const SUDOKU_FIELDS = SUDOKU_SIZE * SUDOKU_SIZE

type Sudoku struct {
	field   [SUDOKU_FIELDS]int // row major
	rows    [SUDOKU_SIZE][SUDOKU_SIZE]bool
	columns [SUDOKU_SIZE][SUDOKU_SIZE]bool
	blocks  [SUDOKU_SIZE][SUDOKU_SIZE]bool
}

func NewEmptySudoku() Sudoku {
	s := Sudoku{}
	for i := range s.field {
		s.field[i] = -1
	}
	return s
}

func getSudokuFieldIndex(row, col int) (int, error) {
	if row < 1 || row > SUDOKU_SIZE {
		return -1, fmt.Errorf("sudoku square has incorrect row %v (should be between 1 and 9)", row)
	}
	if col < 1 || col > SUDOKU_SIZE {
		return -1, fmt.Errorf("sudoku square has incorrect column %v (should be between 1 and 9)", col)
	}
	return (row-1)*SUDOKU_SIZE + col - 1, nil
}

func (su *Sudoku) FillSquare(sq SudokuSquare) error {
	if sq.Val < 1 || sq.Val > 9 {
		return fmt.Errorf("invalid value %v for square at (%v, %v) provided (should be between 1 and 9)", sq.Val, sq.Row, sq.Col)
	}
	i, e := getSudokuFieldIndex(sq.Row, sq.Col)
	if e != nil {
		return e
	}
	if su.field[i] != -1 {
		return fmt.Errorf("square at (%v, %v) is already filled with %v, cannot fill it with %v a second time", sq.Col, sq.Row, su.field[i], sq.Val)
	}
	su.field[i] = sq.Val
	r, c, v := sq.Row-1, sq.Col-1, sq.Val-1
	su.rows[r][v] = true
	su.columns[c][v] = true

	b := (r/SUDOKU_BLOCK_SIZE)*SUDOKU_BLOCK_SIZE + c/SUDOKU_BLOCK_SIZE
	su.blocks[b][v] = true
	return nil
}

type SudokuSquare struct {
	// 1, 1 are the coordinates of the top left corner
	Row int // row index starts at 1
	Col int // column index starts at 1
	Val int
}

func NewSudoku(field [SUDOKU_FIELDS]int) Sudoku {
	s := NewEmptySudoku()
	for i, val := range field {
		if val == -1 {
			continue
		}
		row := i / SUDOKU_SIZE
		col := i - row*SUDOKU_SIZE
		s.FillSquare(SudokuSquare{Row: row + 1, Col: col + 1, Val: val})
	}
	return s
}

func (s Sudoku) ValueAt(row, col int) int {
	return s.field[(row-1)*SUDOKU_SIZE+col-1]
}

func (su Sudoku) String() string {
	s := ""
	var rowStrings [SUDOKU_SIZE]string
	for row := 1; row < 10; row++ {
		for col := 1; col < 10; col++ {
			rowStrings[col-1] = fmt.Sprintf("%2v", su.ValueAt(row, col))
		}
		s = s + strings.Join(rowStrings[:], " ") + "\n"
	}
	return s
}
