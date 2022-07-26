package sudoku_test

import (
	"fmt"
	"math/bits"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_Sudoku_Leet_Format(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		payload [][]string
		want    [][]string
	}{
		{
			name: "Leet format",
			payload: [][]string{
				{"5", "3", ".", ".", "7", ".", ".", ".", "."},
				{"6", ".", ".", "1", "9", "5", ".", ".", "."},
				{".", "9", "8", ".", ".", ".", ".", "6", "."},
				{"8", ".", ".", ".", "6", ".", ".", ".", "3"},
				{"4", ".", ".", "8", ".", "3", ".", ".", "1"},
				{"7", ".", ".", ".", "2", ".", ".", ".", "6"},
				{".", "6", ".", ".", ".", ".", "2", "8", "."},
				{".", ".", ".", "4", "1", "9", ".", ".", "5"},
				{".", ".", ".", ".", "8", ".", ".", "7", "9"},
			},
			want: [][]string{
				{"5", "3", "4", "6", "7", "8", "9", "1", "2"},
				{"6", "7", "2", "1", "9", "5", "3", "4", "8"},
				{"1", "9", "8", "3", "4", "2", "5", "6", "7"},
				{"8", "5", "9", "7", "6", "1", "4", "2", "3"},
				{"4", "2", "6", "8", "5", "3", "7", "9", "1"},
				{"7", "1", "3", "9", "2", "4", "8", "5", "6"},
				{"9", "6", "1", "5", "3", "7", "2", "8", "4"},
				{"2", "8", "7", "4", "1", "9", "6", "3", "5"},
				{"3", "4", "5", "2", "8", "6", "1", "7", "9"},
			},
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d: %s", i, tt.name), func(t *testing.T) {
			result := solveSudoku(tt.payload)
			assert.Equal(t, tt.want, result)
		})
	}
}
func Test_Sudoku_MainTest_Solve(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		payload []byte
	}{
		{
			name:    "SDK 2",
			payload: datum_2,
		},
		{
			name:    "SDK 3",
			payload: datum_3,
		},
		{
			name:    "SDK 4",
			payload: datum_4,
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d: %s", i, tt.name), func(t *testing.T) {
			sudoku, err := CreateSudokuBoard(tt.payload)
			assert.Equal(t, nil, err)
			final := sudoku.Solve()
			assert.Equal(t, 81, final.SolvedCnt())
		})
	}
}

func Test_Sudoku_2_BIGTEST(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		payload []byte
	}{
		{
			name:    "SDK 2",
			payload: datum_4,
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d: %s", i, tt.name), func(t *testing.T) {
			sudoku, err := CreateSudokuBoard(tt.payload)
			assert.Equal(t, nil, err)
			assert.Equal(t, []int{1, 2, 9, 10, 11, 18, 19, 20, 3, 4, 5, 6, 7, 8, 27, 36, 45, 54, 63, 72}, sudoku.Neighborhood[0])
			assert.Equal(t, []int{33, 34, 35, 42, 43, 44, 51, 52, 45, 46, 47, 48, 49, 50, 8, 17, 26, 62, 71, 80}, sudoku.Neighborhood[53])
			assert.Equal(t, []int{57, 58, 59, 67, 68, 75, 76, 77, 63, 64, 65, 69, 70, 71, 3, 12, 21, 30, 39, 48}, sudoku.Neighborhood[66])

			secondBoard := sudoku.SweepMark()
			thirdBoard := secondBoard.SweepMark()
			fourthBoard := thirdBoard.SweepMark()

			time.Sleep(time.Second)
			println("======" + sudoku.Name + "==[" + strconv.Itoa(sudoku.SolvedCnt()) + "]=======")
			sudoku.PrintDecimal()

			time.Sleep(time.Second)
			println("---- " + strconv.Itoa(secondBoard.SolvedCnt()) + " -------")
			secondBoard.PrintDecimal()

			time.Sleep(time.Second)
			println("---- " + strconv.Itoa(thirdBoard.SolvedCnt()) + " -------")
			thirdBoard.PrintDecimal()

			time.Sleep(time.Second)
			println("---- " + strconv.Itoa(fourthBoard.SolvedCnt()) + " -------")
			fourthBoard.PrintDecimal()

		})
	}
}

func Test_MarkSweep(t *testing.T) {
	t.Run(fmt.Sprintf("Test: %s", "Mark dna Sweep"), func(t *testing.T) {

		sb, err := CreateSudokuBoard(datum_2)
		assert.Equal(t, nil, err)
		// sb.PoplateNeighborhood()
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
