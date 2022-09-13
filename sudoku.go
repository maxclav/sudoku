package sudoku

// Board is a sudoku board.
type Board [][]byte

const (
	// NumberRowsClassic is the number of row
	// in a classic sudoku board.
	NumberRowsClassic int = 9
	// NumberColumnsClassic is the number of columns
	// in a classic sudoku board.
	NumberColumnsClassic int = 9
)

// NewClassicRandom returns a new valid Baord.
// TODO: take a look at https://github.com/eliben/go-sudoku
func NewClassicRandom() *Board {
	return nil
}
