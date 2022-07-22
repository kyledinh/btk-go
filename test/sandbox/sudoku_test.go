package sandbox_test

import (
	"encoding/json"
	"fmt"
	"math/bits"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Board struct {
	Squares [][]uint16 `json:"squares"`
	// [0,0][0,1][0,2]
	// [1,0][1,1][1,2]
	// [2,0][2,1][2,2]
	// Second index is the bitmap that holds the val or possibilities
	// Only one bit set = value, more than one is possbilities
}

func SquarePrint(bitmap uint16) string {
	// 	if bits.OnesCount16(bitmap) == 1 {
	// 		return strconv.Itoa(int(bitmap))
	// 	}
	// 	return ""
	return strconv.Itoa(int(bitmap))
}

func SquareValue(bitmap uint16) (int, []int) {
	switch bits.OnesCount16(bitmap) {
	case 1: // square is solved
		return int(bitmap), nil
	case 0: // error
		return 0, nil
	}
	// return all possbilities
	return 0, nil
}

func (b *Board) PrintRow(row int) []string {
	out := make([]string, 0)
	out = append(out, SquarePrint(b.Squares[row][0]))
	out = append(out, SquarePrint(b.Squares[row][1]))
	out = append(out, SquarePrint(b.Squares[row][2]))
	return out
}

type Sudoku struct {
	// [0][1][2]
	// [3][4][5]
	// [6][7][8]
	Boards [9]Board `json:"boards"`
}

func (s *Sudoku) GridPrint() [9]string {
	var rows = [9]string{}
	rows[0] = strings.Join(append(s.Boards[0].PrintRow(0), append(s.Boards[1].PrintRow(0), s.Boards[2].PrintRow(0)...)...), ",")
	rows[1] = strings.Join(append(s.Boards[0].PrintRow(1), append(s.Boards[1].PrintRow(1), s.Boards[2].PrintRow(1)...)...), ",")
	rows[2] = strings.Join(append(s.Boards[0].PrintRow(2), append(s.Boards[1].PrintRow(2), s.Boards[2].PrintRow(2)...)...), ",")
	rows[3] = strings.Join(append(s.Boards[3].PrintRow(0), append(s.Boards[4].PrintRow(0), s.Boards[5].PrintRow(0)...)...), ",")
	rows[4] = strings.Join(append(s.Boards[3].PrintRow(1), append(s.Boards[4].PrintRow(1), s.Boards[5].PrintRow(1)...)...), ",")
	rows[5] = strings.Join(append(s.Boards[3].PrintRow(2), append(s.Boards[4].PrintRow(2), s.Boards[5].PrintRow(2)...)...), ",")
	rows[6] = strings.Join(append(s.Boards[6].PrintRow(0), append(s.Boards[7].PrintRow(0), s.Boards[8].PrintRow(0)...)...), ",")
	rows[7] = strings.Join(append(s.Boards[6].PrintRow(1), append(s.Boards[7].PrintRow(1), s.Boards[8].PrintRow(1)...)...), ",")
	rows[8] = strings.Join(append(s.Boards[6].PrintRow(2), append(s.Boards[7].PrintRow(2), s.Boards[8].PrintRow(2)...)...), ",")
	return rows
}

func CreateSudoku(ba []byte) (Sudoku, error) {
	var sudoku Sudoku
	err := json.Unmarshal(ba, &sudoku)
	return sudoku, err
}

func Test_Sudoku(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		payload []byte
		want    string
	}{
		{
			name:    "SDK 1",
			payload: datum_1,
			want:    "0,8,4",
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d: %s", i, tt.name), func(t *testing.T) {

			sudoku, err := CreateSudoku(tt.payload)
			assert.Equal(t, nil, err)
			block0_row0 := strings.Join(sudoku.Boards[0].PrintRow(0), ",")
			assert.Equal(t, tt.want, block0_row0)
		})
	}
}

var datum_1 = []byte(`
	{
		"boards": [
			{ "squares": [[0,8,4],[2,0,7],[6,0,0]] },
			{ "squares": [[0,7,2],[8,3,0],[5,0,9]] },
			{ "squares": [[1,0,5],[9,0,0],[0,0,8]] },
			{ "squares": [[0,6,0],[0,7,0],[0,2,0]] },
			{ "squares": [[9,2,8],[0,0,0],[0,0,0]] },
			{ "squares": [[4,0,0],[0,6,9],[0,8,1]] },
			{ "squares": [[0,3,2],[7,0,0],[1,0,0]] },
			{ "squares": [[0,5,0],[0,0,0],[2,0,4]] },
			{ "squares": [[6,9,4],[0,0,2],[0,0,7]] }
		]
	}
`)
