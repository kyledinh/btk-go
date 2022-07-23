package sandbox_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type SudokuBoard struct {
	Grid [81]uint16 `json:"grid"`
	//  0, 1, 2  3, 4, 5  6, 7, 8
	//  9,10,11 12,13,14 15,16,17
	// 18,19,20 21,22,23 24,25,26

	// 27,28,29 30,31,32 33,34,35
	// 36,37,38 39,40,41 42,43,44
	// 45,46,47 48,49,50 51,52,53

	// 54,55,56 57,58,59 60,61,62
	// 63,64,65 66,67,68 69,70,71
	// 72,73,74 75,76,77 78,79,80
	Neighborhood [][]int
}

func findCurrentRowAndColumn(num int) (row int, col int) {
	for multi := 0; multi < 9; multi++ {
		lowerLimit := (9 * multi) - 1
		upperLimit := (9 * multi) + 9
		if num > lowerLimit && num < upperLimit {
			row = multi
		}
	}
	col = num % 9
	return
}

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func RemoveIntElement(s []int, element int) []int {
	newArray := make([]int, 0)
	for _, val := range s {
		if val != element {
			newArray = append(newArray, val)
		}
	}
	return newArray
}

func (sb *SudokuBoard) PoplateNeighborhood() {
	sb.Neighborhood = make([][]int, 81)
	for i := 0; i <= 80; i++ {
		row, col := findCurrentRowAndColumn(i)
		// Populate Grid Neighborhood
		if row < 3 && col < 3 {
			sb.Neighborhood[i] = RemoveIntElement([]int{0, 1, 2, 9, 10, 11, 18, 19, 20}, i)
		}
		if row >= 3 && row < 6 && col < 3 {
			sb.Neighborhood[i] = RemoveIntElement([]int{27, 28, 29, 36, 37, 38, 45, 46, 47}, i)
		}
		if row >= 6 && row < 9 && col < 3 {
			sb.Neighborhood[i] = RemoveIntElement([]int{54, 55, 56, 63, 64, 65, 72, 73, 74}, i)
		}
		// Grids: 1,4,7
		if row < 3 && col >= 3 && col < 6 {
			sb.Neighborhood[i] = RemoveIntElement([]int{3, 4, 5, 12, 13, 14, 21, 22, 23}, i)
		}
		if row >= 3 && row < 6 && col >= 3 && col < 6 {
			sb.Neighborhood[i] = RemoveIntElement([]int{30, 31, 32, 39, 40, 41, 48, 49, 50}, i)
		}
		if row >= 6 && row < 9 && col >= 3 && col < 6 {
			sb.Neighborhood[i] = RemoveIntElement([]int{57, 58, 59, 66, 67, 68, 75, 76, 77}, i)
		}
		// Grids: 2,5,8
		if row < 3 && col >= 6 && col < 9 {
			sb.Neighborhood[i] = RemoveIntElement([]int{6, 7, 8, 15, 16, 17, 24, 25, 26}, i)
		}
		if row >= 3 && row < 6 && col >= 6 && col < 9 {
			sb.Neighborhood[i] = RemoveIntElement([]int{33, 34, 35, 42, 43, 44, 51, 52, 53}, i)
		}
		if row >= 6 && row < 9 && col >= 6 && col < 9 {
			sb.Neighborhood[i] = RemoveIntElement([]int{60, 61, 62, 69, 70, 71, 78, 79, 80}, i)
		}
		// Check Row Neighbors
		lowerBound := (row * 9)
		upperBound := (row * 9) + 9 // lowerBound +9
		for j := lowerBound; j >= lowerBound && j < upperBound; j++ {
			if !contains(sb.Neighborhood[i], j) && i != j {
				sb.Neighborhood[i] = append(sb.Neighborhood[i], j)
			}
		}
		// interate the rows to get the column neigbors
		for r := 0; r < 9; r++ {
			neighbor := (r * 9) + col
			if !contains(sb.Neighborhood[i], neighbor) && i != neighbor {
				sb.Neighborhood[i] = append(sb.Neighborhood[i], neighbor)
			}
		}
	}
}

func CreateSudokuBoard(ba []byte) (SudokuBoard, error) {
	var sudoku SudokuBoard
	err := json.Unmarshal(ba, &sudoku)
	return sudoku, err
}

func Test_Sudoku_2_Neighborhoods(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		payload []byte
		want    string
	}{
		{
			name:    "SDK 1",
			payload: datum_2,
			want:    "",
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
		})
	}
}

var datum_2 = []byte(`
	{
		"grid": [
			3,0,0,1,0,8,0,6,0,
			0,2,6,0,4,0,8,0,0,
			0,0,1,0,0,0,4,0,0,
			5,0,8,0,0,7,0,0,1,
			0,0,0,0,9,5,7,0,0,
			0,7,9,2,3,0,0,0,0,
			9,0,0,0,0,0,5,0,6,
			2,0,4,0,0,0,1,0,0,
			6,0,0,5,0,0,0,2,0
		]
	}
`)

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
