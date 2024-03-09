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
	allowed [SUDOKU_FIELDS][SUDOKU_SIZE]bool
}

func NewEmptySudoku() Sudoku {
	s := Sudoku{}
	for i := 0; i < SUDOKU_FIELDS; i++ {
		s.field[i] = -1
		for j := 0; j < SUDOKU_SIZE; j++ {
			s.allowed[i][j] = true
		}
	}
	return s
}

func sudokuField(row, col int) int {
	return (row-1)*SUDOKU_SIZE + col - 1
}

func sudokuRowCol(i int) (int, int) {
	row := i / SUDOKU_SIZE
	col := i - row*SUDOKU_SIZE
	return row + 1, col + 1
}

func (su *Sudoku) setField(i, row, col, v int) {
	su.field[i] = v
	row--
	col--
	v--
	// update row
	for j := 0; j < SUDOKU_SIZE; j++ {
		su.allowed[row*SUDOKU_SIZE+j][v] = false
	}
	// update column
	for j := 0; j < SUDOKU_SIZE; j++ {
		su.allowed[j*SUDOKU_SIZE+col][v] = false
	}
	// update block
	blockRow := row / SUDOKU_BLOCK_SIZE
	blockCol := col / SUDOKU_BLOCK_SIZE
	for j := 0; j < SUDOKU_BLOCK_SIZE; j++ {
		for k := 0; k < SUDOKU_BLOCK_SIZE; k++ {
			r := blockRow*SUDOKU_BLOCK_SIZE + j
			c := blockCol*SUDOKU_BLOCK_SIZE + k
			su.allowed[r*SUDOKU_SIZE+c][v] = false
		}
	}
}

func (su *Sudoku) FillSquare(sq SudokuSquare) error {
	if sq.Val < 1 || sq.Val > 9 {
		return fmt.Errorf("invalid value %v for square at (%v, %v) provided (should be between 1 and 9)", sq.Val, sq.Row, sq.Col)
	}
	if sq.Row < 1 || sq.Row > 9 {
		return fmt.Errorf("invalid row index %v (should be between 1 and 9)", sq.Row)
	}
	if sq.Col < 1 || sq.Col > 9 {
		return fmt.Errorf("invalid column index %v (should be between 1 and 9)", sq.Col)
	}
	i := sudokuField(sq.Row, sq.Col)
	if !su.allowed[i][sq.Val-1] {
		return fmt.Errorf("value %v is not allowed at (%v, %v)", sq.Val, sq.Row, sq.Col)
	}
	su.setField(i, sq.Row, sq.Col, sq.Val)
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
		row, col := sudokuRowCol(i)
		s.FillSquare(SudokuSquare{row, col, val})
	}
	return s
}

func (s Sudoku) ValueAt(row, col int) int {
	i := sudokuField(row, col)
	if i < 0 || i >= SUDOKU_FIELDS {
		return -1
	}
	return s.field[i]
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
