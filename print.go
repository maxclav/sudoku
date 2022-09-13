package sudoku

import (
	"errors"
	"fmt"
)

// Print prints the Board.
func Print(board *Board) error {
	if board == nil {
		return errors.New("board is nil")
	}

	b := [][]byte(*board)
	for i, row := range b {
		for j, val := range row[:len(row)-1] {
			fmt.Printf("%v ", val)
			if (j+1)%3 == 0 {
				fmt.Printf("| ")
			}
		}
		fmt.Printf("%v\n", row[len(row)-1])
		if (i+1)%3 == 0 {
			for j := 0; j < len(row); j++ {
				fmt.Print("-")
			}
			fmt.Println()
		}
	}
	fmt.Println()
	return nil
}
