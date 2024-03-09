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
	ncovers [SUDOKU_FIELDS][SUDOKU_SIZE]int
}

type SudokuMove struct {
	// 1, 1 are the coordinates of the top left corner
	Row int // row index starts at 1
	Col int // column index starts at 1
	Val int
}

func NewEmptySudoku() Sudoku {
	s := Sudoku{}
	for i := range s.field {
		s.field[i] = -1
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

func (su *Sudoku) Allowed(mv SudokuMove) bool {
	return su.ncovers[sudokuField(mv.Row, mv.Col)][mv.Val-1] == 0
}

func (su *Sudoku) makeMove(mv SudokuMove) {
	su.field[sudokuField(mv.Row, mv.Col)] = mv.Val
	row := mv.Row - 1
	col := mv.Col - 1
	v := mv.Val - 1
	// update row
	for j := 0; j < SUDOKU_SIZE; j++ {
		su.ncovers[row*SUDOKU_SIZE+j][v]++
	}
	// update column
	for j := 0; j < SUDOKU_SIZE; j++ {
		su.ncovers[j*SUDOKU_SIZE+col][v]++
	}
	// update block
	blockRow := row / SUDOKU_BLOCK_SIZE
	blockCol := col / SUDOKU_BLOCK_SIZE
	for j := 0; j < SUDOKU_BLOCK_SIZE; j++ {
		for k := 0; k < SUDOKU_BLOCK_SIZE; k++ {
			r := blockRow*SUDOKU_BLOCK_SIZE + j
			c := blockCol*SUDOKU_BLOCK_SIZE + k
			su.ncovers[r*SUDOKU_SIZE+c][v]++
		}
	}
}

func (su *Sudoku) MakeMove(mv SudokuMove) error {
	if mv.Val < 1 || mv.Val > 9 {
		return fmt.Errorf("invalid value %v for square at (%v, %v) provided (should be between 1 and 9)", mv.Val, mv.Row, mv.Col)
	}
	if mv.Row < 1 || mv.Row > 9 {
		return fmt.Errorf("invalid row index %v (should be between 1 and 9)", mv.Row)
	}
	if mv.Col < 1 || mv.Col > 9 {
		return fmt.Errorf("invalid column index %v (should be between 1 and 9)", mv.Col)
	}
	if !su.Allowed(mv) {
		return fmt.Errorf("value %v is not allowed at (%v, %v)", mv.Val, mv.Row, mv.Col)
	}
	su.makeMove(mv)
	return nil
}

func NewSudoku(field [SUDOKU_FIELDS]int) Sudoku {
	s := NewEmptySudoku()
	for i, val := range field {
		if val == -1 {
			continue
		}
		row, col := sudokuRowCol(i)
		s.MakeMove(SudokuMove{row, col, val})
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
