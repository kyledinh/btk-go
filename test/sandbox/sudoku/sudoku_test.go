package sudoku_test

import (
	"fmt"
	"math/bits"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_Sudoku_2_BIGTEST(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		payload []byte
	}{
		{
			name:    "SDK 2",
			payload: datum_3,
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d: %s", i, tt.name), func(t *testing.T) {
			sudoku, err := CreateSudokuBoard(tt.payload)
			assert.Equal(t, nil, err)
			sudoku.PoplateNeighborhood()
			assert.Equal(t, []int{1, 2, 9, 10, 11, 18, 19, 20, 3, 4, 5, 6, 7, 8, 27, 36, 45, 54, 63, 72}, sudoku.Neighborhood[0])
			assert.Equal(t, []int{33, 34, 35, 42, 43, 44, 51, 52, 45, 46, 47, 48, 49, 50, 8, 17, 26, 62, 71, 80}, sudoku.Neighborhood[53])
			assert.Equal(t, []int{57, 58, 59, 67, 68, 75, 76, 77, 63, 64, 65, 69, 70, 71, 3, 12, 21, 30, 39, 48}, sudoku.Neighborhood[66])

			secondBoard := sudoku.SweepMark()
			time.Sleep(3 * time.Second)
			thirdBoard := secondBoard.SweepMark()
			fourthBoard := thirdBoard.SweepMark()

			time.Sleep(time.Second)
			println("======" + sudoku.Name + "========")
			sudoku.PrintRoman()

			time.Sleep(time.Second)
			println("-----------")
			secondBoard.PrintRoman()

			time.Sleep(time.Second)
			println("-----------")
			thirdBoard.PrintRoman()

			time.Sleep(time.Second)
			println("-----------")
			fourthBoard.PrintRoman()

		})
	}
}

func Test_MarkSweep(t *testing.T) {
	t.Run(fmt.Sprintf("Test: %s", "Mark dna Sweep"), func(t *testing.T) {

		sb, err := CreateSudokuBoard(datum_2)
		assert.Equal(t, nil, err)
		sb.PoplateNeighborhood()
		newBoard := sb

		for i := 0; i < 81; i++ {
			bitmap := sb.Grid[i]
			bitCount := bits.OnesCount16(bitmap)
			fmt.Printf("for %d count is %d \n", i, bitCount)
			if bitCount == 1 {
				fmt.Printf("Solved for %v \n", i)
			} else {
				fmt.Printf("doing(%v) with: ", i)
				fmt.Printf("========== BEFORE %v \n", IntToBinaryString(newBoard.Grid[i]))
				for _, square := range sb.Neighborhood[i] {
					if bits.OnesCount16(sb.Grid[square]) == 1 { // Square is SOLVED with only 1 possility
						fmt.Printf("========== NEIGHBOR %v \n", IntToBinaryString(sb.Grid[square]))
						bitmap = bitmap &^ sb.Grid[square]
					}

				}
				newBoard.Grid[i] = bitmap
			}
			fmt.Printf("========== AFTER new value %v \n\n", IntToBinaryString(newBoard.Grid[i]))
		}

	})
}

func Test_findMultiplier(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		index   int
		wantRow int
		wantCol int
	}{
		{
			name:    "mulit 0",
			index:   2,
			wantRow: 0,
			wantCol: 2,
		},
		{
			name:    "multi 5",
			index:   45,
			wantRow: 5,
			wantCol: 0,
		},
		{
			name:    "mulit 8",
			index:   78,
			wantRow: 8,
			wantCol: 6,
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d: %s", i, tt.name), func(t *testing.T) {
			row, col := findCurrentRowAndColumn(tt.index)
			assert.Equal(t, tt.wantRow, row)
			assert.Equal(t, tt.wantCol, col)
		})
	}
}
