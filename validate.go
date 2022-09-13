package sudoku

import (
	"errors"
	"fmt"
	"strings"

	"github.com/maxclav/sudoku/pkg/set"
)

// IsClassicValid returns whether the sudoku board
// is a valid (true) or not (false) classic board.
// See: leetcode.com/problems/valid-sudoku/discuss/2571745.
func IsClassicValid(board *Board) (bool, error) {
	if board == nil {
		return false, errors.New("board is nil")
	}

	b := [][]byte(*board)
	if len(b) != NumberRowsClassic {
		return false, errors.New("invalid number of rows")
	}
	if len(b[0]) != NumberColumnsClassic {
		return false, errors.New("invalid number of columns")
	}

	s := set.NewString()
	for i := 0; i < NumberRowsClassic; i++ {
		for j := 0; j < NumberColumnsClassic; j++ {
			if !validate(s, b[i][j], j, i) {
				return false, nil
			}
		}
	}
	return true, nil
}

func validate(s *set.String, b byte, i, j int) bool {
	if str := string(b); !strings.EqualFold(str, ".") {
		keyRow := func(s string, row int) string {
			return fmt.Sprintf("%v in row %v", s, row)
		}

		keyCol := func(s string, row int) string {
			return fmt.Sprintf("%v in col %v", s, row)
		}

		keyBlock := func(s string, row, col int) string {
			return fmt.Sprintf("%v in block %v-%v", s, row/3, col/3)
		}

		keys := func(s string, row, col int) (string, string, string) {
			return keyRow(s, row), keyCol(s, col), keyBlock(s, row, col)
		}

		row, col, block := keys(str, i, j)
		if !s.Add(row) || !s.Add(col) || !s.Add(block) {
			return false
		}
	}
	return true
}
