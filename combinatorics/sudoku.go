package combinatorics

import (
	"fmt"
	"strings"
)

const SUDOKU_SIZE = 9
const SUDOKU_BLOCK_SIZE = SUDOKU_SIZE / 3
const SUDOKU_FIELDS = SUDOKU_SIZE * SUDOKU_SIZE

type Sudoku struct {
	field       []int // row major
	ncovers     [][]int
	NFreeFields int
}

type SudokuSquare struct {
	// 1, 1 are the coordinates of the top left corner
	Row int // row index starts at 1
	Col int // column index starts at 1
}

func (sq SudokuSquare) fieldIndex() int {
	return (sq.Row-1)*SUDOKU_SIZE + sq.Col - 1
}

func newSudokuSquare(i int) SudokuSquare {
	row := i / SUDOKU_SIZE
	col := i - row*SUDOKU_SIZE
	return SudokuSquare{row + 1, col + 1}

}

func NewEmptySudoku() Sudoku {
	field := make([]int, SUDOKU_FIELDS)
	for i := range field {
		field[i] = -1
	}
	ncovers := make([][]int, SUDOKU_FIELDS)
	for i := range ncovers {
		ncovers[i] = make([]int, SUDOKU_SIZE)
	}
	s := Sudoku{field, ncovers, SUDOKU_FIELDS}
	return s
}

func (su Sudoku) Allowed(sq SudokuSquare, val int) bool {
	return su.ncovers[sq.fieldIndex()][val-1] == 0
}

func (su *Sudoku) updateNcovers(sq SudokuSquare, val, delta int) {
	row := sq.Row - 1
	col := sq.Col - 1
	v := val - 1
	// update row
	for j := 0; j < SUDOKU_SIZE; j++ {
		su.ncovers[row*SUDOKU_SIZE+j][v] += delta
	}
	// update column
	for j := 0; j < SUDOKU_SIZE; j++ {
		su.ncovers[j*SUDOKU_SIZE+col][v] += delta
	}
	// update block
	blockRow := row / SUDOKU_BLOCK_SIZE
	blockCol := col / SUDOKU_BLOCK_SIZE
	for j := 0; j < SUDOKU_BLOCK_SIZE; j++ {
		for k := 0; k < SUDOKU_BLOCK_SIZE; k++ {
			r := blockRow*SUDOKU_BLOCK_SIZE + j
			c := blockCol*SUDOKU_BLOCK_SIZE + k
			su.ncovers[r*SUDOKU_SIZE+c][v] += delta
		}
	}
}

func (su *Sudoku) makeMove(sq SudokuSquare, val int) {
	su.field[sq.fieldIndex()] = val
	su.updateNcovers(sq, val, 1)
	su.NFreeFields--
}

func (su *Sudoku) unmakeMove(sq SudokuSquare, oldVal int) {
	su.field[sq.fieldIndex()] = -1
	su.updateNcovers(sq, oldVal, -1)
	su.NFreeFields++
}

func checkSquare(sq SudokuSquare) error {
	if sq.Row < 1 || sq.Row > 9 {
		return fmt.Errorf("invalid row %v (should be between 1 and 9)", sq.Row)
	}
	if sq.Col < 1 || sq.Col > 9 {
		return fmt.Errorf("invalid column %v (should be between 1 and 9)", sq.Col)
	}
	return nil
}

func (su *Sudoku) MakeMove(sq SudokuSquare, val int) error {
	if val < 1 || val > 9 {
		return fmt.Errorf("invalid value %v (should be between 1 and 9)", val)
	}
	e := checkSquare(sq)
	if e != nil {
		return e
	}
	if !su.Allowed(sq, val) {
		return fmt.Errorf("value %v is not allowed at (%v, %v)", val, sq.Row, sq.Col)
	}
	su.makeMove(sq, val)
	return nil
}

func (su *Sudoku) UnmakeMove(sq SudokuSquare) error {
	e := checkSquare(sq)
	if e != nil {
		return e
	}
	oldVal := su.ValueAt(sq)
	if oldVal == -1 {
		return nil
	}
	su.unmakeMove(sq, oldVal)
	return nil
}

func NewSudoku(field [SUDOKU_FIELDS]int) Sudoku {
	su := NewEmptySudoku()
	for i, val := range field {
		if val == -1 {
			continue
		}
		sq := newSudokuSquare(i)
		su.MakeMove(sq, val)
	}
	return su
}

func (su Sudoku) ValueAt(sq SudokuSquare) int {
	i := sq.fieldIndex()
	if i < 0 || i >= SUDOKU_FIELDS {
		return -1
	}
	return su.field[i]
}

func (su Sudoku) String() string {
	s := ""
	rowStrings := make([]string, SUDOKU_SIZE)
	for row := 1; row < 10; row++ {
		for col := 1; col < 10; col++ {
			rowStrings[col-1] = fmt.Sprintf("%2v", su.ValueAt(SudokuSquare{row, col}))
		}
		s = s + strings.Join(rowStrings[:], " ")
		if row != 9 {
			s = s + "\n"
		}
	}
	return s
}

func nextSquare(su Sudoku) SudokuSquare {
	// we assume that su is not fully filled
	nposs := make([]int, SUDOKU_FIELDS)
	for i := range nposs {
		if su.field[i] != -1 {
			continue
		}
		for _, n := range su.ncovers[i] {
			if n == 0 {
				nposs[i]++
			}
		}
	}

	var min int
	for i, n := range nposs {
		if n > 0 {
			min = i
			break
		}
	}
	for i := min + 1; i < len(nposs); i++ {
		if nposs[i] == 0 {
			continue
		}
		if nposs[i] < nposs[min] {
			min = i
		}
	}
	return newSudokuSquare(min)
}

func backtrackSudoku(moves []SudokuSquare, values []int, k int, su *Sudoku) {
	if su.NFreeFields == 0 {
		return
	}

	sq := nextSquare(*su)
	fmt.Println(sq)
}

func (su *Sudoku) Solve() {
	moves := make([]SudokuSquare, su.NFreeFields)
	values := make([]int, SUDOKU_SIZE)
	backtrackSudoku(moves, values, 0, su)
}
