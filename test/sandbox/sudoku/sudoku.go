package sudoku_test

import (
	"encoding/json"
	"fmt"
	"math/bits"
	"strconv"
)

// DATA STRUCTURE

type SudokuBoard struct {
	Name string     `json:"name"`
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
	Verbose      bool
	Summary      struct {
		Steps       int `json:"steps"`
		SolvedStart int `json:"solved_start"`
		SolvedEnd   int `json:"solved_end"`
	}
}

// SUDOKUBOARD METHODS

func (sb *SudokuBoard) Solve() (final SudokuBoard) {

	history := make([]SudokuBoard, 0)
	sb.Summary.SolvedStart = sb.SolvedCnt()
	history = append(history, *sb)

	firstBoard := sb.SweepMark()
	history = append(history, firstBoard)

	for i := 2; history[i-1].SolvedCnt() < 81; i++ {
		before := history[i-2]
		previous := history[i-1]
		// Check if there was progress from before to previous
		if previous.SolvedCnt() == before.SolvedCnt() {
			break
		}
		next := previous.SweepMark()
		history = append(history, next)
	}

	finalIndex := len(history)
	final = history[finalIndex-1]
	final.Summary.SolvedStart = sb.SolvedCnt()
	final.Summary.SolvedEnd = final.SolvedCnt()
	final.Summary.Steps = finalIndex

	return final
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

func (sb *SudokuBoard) SweepMark() SudokuBoard {
	var newBoard SudokuBoard
	for i, square := range sb.Grid { // check if I need to copy by value or another kind of copy
		newBoard.Grid[i] = square
	}
	newBoard.Neighborhood = sb.Neighborhood
	for i := 0; i < 81; i++ {
		bitmap := sb.Grid[i] // uint16 representatin of a bitmap: "001000101"

		if bits.OnesCount16(bitmap) == 1 { // Square is SOLVED with only 1 possility, 1 bit flipped.
			VerbosePrint(sb.Verbose, fmt.Sprintf("Solved for %v \n", i))
		} else { // Square has 2 or more possibilities
			VerbosePrint(sb.Verbose, fmt.Sprintf("=== NEIGHBORHOOD %v \n", sb.Neighborhood[i]))
			for _, square := range sb.Neighborhood[i] {
				if bits.OnesCount16(sb.Grid[square]) == 1 { // The Neighbor is SOLVED, use it to calculate
					VerbosePrint(sb.Verbose, fmt.Sprintf("=== INDEX %v == BITUINT16 %v == NEIGHBOR %v \n", i, sb.Grid[square], IntToBinaryString(sb.Grid[square])))
					bitmap = bitmap &^ sb.Grid[square] // bitclear &^ : if matches a 1 in the square, flip to 0 in the bitmap
				}
			}
			newBoard.Grid[i] = bitmap
		}
		VerbosePrint(sb.Verbose, fmt.Sprintf("========== AFTER new value %v \n\n", IntToBinaryString(newBoard.Grid[i])))
	}
	newBoard.Summary.SolvedStart = sb.SolvedCnt()
	newBoard.Summary.SolvedEnd = newBoard.SolvedCnt()
	return newBoard
}

// Displays the grid with uint16 values
func (sb *SudokuBoard) PrintBoard() {
	fmt.Printf("%v \n", sb.Grid[:9])
	fmt.Printf("%v \n", sb.Grid[9:18])
	fmt.Printf("%v \n", sb.Grid[18:27])
	fmt.Printf("%v \n", sb.Grid[27:36])
	fmt.Printf("%v \n", sb.Grid[36:45])
	fmt.Printf("%v \n", sb.Grid[45:54])
	fmt.Printf("%v \n", sb.Grid[54:63])
	fmt.Printf("%v \n", sb.Grid[63:72])
	fmt.Printf("%v \n", sb.Grid[72:])
}

// Displays the roman numerals 1-9, or 0
func (sb *SudokuBoard) PrintDecimal() {
	decimal := make([]int, 81)
	for i, val := range sb.Grid {
		decimal[i] = BitmapToDecimal(val)
	}
	fmt.Printf("%v \n", decimal[:9])
	fmt.Printf("%v \n", decimal[9:18])
	fmt.Printf("%v \n", decimal[18:27])
	fmt.Printf("%v \n", decimal[27:36])
	fmt.Printf("%v \n", decimal[36:45])
	fmt.Printf("%v \n", decimal[45:54])
	fmt.Printf("%v \n", decimal[54:63])
	fmt.Printf("%v \n", decimal[63:72])
	fmt.Printf("%v \n", decimal[72:])
}

func (sb *SudokuBoard) SolvedCnt() (cnt int) {
	for _, bitmap := range sb.Grid {
		if bits.OnesCount16(bitmap) == 1 {
			cnt++
		}
	}
	return
}

// PUBLIC FUNCTIONS

func CreateSudokuBoard(ba []byte) (SudokuBoard, error) {
	var sudoku SudokuBoard
	err := json.Unmarshal(ba, &sudoku)
	if err != nil {
		return sudoku, err
	}
	// Converts the JSON decimal number input into uint16 bitmap,
	// ie 3 ->  4 | "0000000100" 3rd binary position
	// ie 5 -> 16 | "0000010000" 5th binary position
	for i := 0; i < 81; i++ {
		sudoku.Grid[i] = DecimalToBitmap(sudoku.Grid[i])
	}
	sudoku.PoplateNeighborhood()

	return sudoku, err
}

func BinaryLookup1to9(i uint16) (uint16, error) {
	binaries := []string{
		"111111111",
		"000000001",
		"000000010",
		"000000100",
		"000001000",
		"000010000",
		"000100000",
		"001000000",
		"010000000",
		"100000000",
	}
	numInt, err := strconv.ParseInt(binaries[i], 2, 64)
	return uint16(numInt), err
}

func IntToBinaryString(i uint16) string {
	binstr := strconv.FormatInt(int64(i), 2)
	return binstr
}

func DecimalToBitmap(i uint16) uint16 {
	bitmap, err := BinaryLookup1to9(i)
	if err != nil {
		// TODO: handke error instead of returning 0
		return uint16(511) // binary "111111111"
	}
	return bitmap
}

func BitmapToDecimal(bm uint16) int {
	// bin := strconv.FormatInt(int64(bm), 2)
	// decimal, err := strconv.ParseInt(bin, 2, 64)
	// TODO: find better way to find exp of 2
	if bm == 1 {
		return 1
	}
	if bm == 2 {
		return 2
	}
	if bm == 4 {
		return 3
	}
	if bm == 8 {
		return 4
	}
	if bm == 16 {
		return 5
	}
	if bm == 32 {
		return 6
	}
	if bm == 64 {
		return 7
	}
	if bm == 128 {
		return 8
	}
	if bm == 256 {
		return 9
	}
	return 0
}

// PRIVATE HELPER FUNCTIONS

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

func strIsNumeral(str string) bool {
	s := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "0"}
	for _, v := range s {
		if v == str {
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

func VerbosePrint(verbose bool, message string) {
	if verbose {
		fmt.Print(message)
	}
}

// LEETCODE WRAPPER

func (sb *SudokuBoard) LeetOut() [][]string {
	var gridStr [81]string
	out := make([][]string, 9)
	for i, bitmap := range sb.Grid {
		gridStr[i] = fmt.Sprint(BitmapToDecimal(bitmap))
	}
	out[0] = gridStr[:9]
	out[1] = gridStr[9:18]
	out[2] = gridStr[18:27]
	out[3] = gridStr[27:36]
	out[4] = gridStr[36:45]
	out[5] = gridStr[45:54]
	out[6] = gridStr[54:63]
	out[7] = gridStr[63:72]
	out[8] = gridStr[72:]
	return out
}

func leet2bitmap(str string) (bitmap uint16, err error) {
	if strIsNumeral(str) {
		intNum, err := strconv.Atoi(str)
		return DecimalToBitmap(uint16(intNum)), err
	}
	return DecimalToBitmap(uint16(0)), err
}

func solveSudoku(board [][]string) (solution [][]string) {
	oneline := make([]string, 0)
	for _, sa := range board {
		oneline = append(oneline, sa...)
	}

	var grid [81]uint16
	for i, str := range oneline {
		square, err := leet2bitmap(str)
		grid[i] = square
		if err != nil {
			panic(err)
		}
	}

	sudokuBoard := SudokuBoard{
		Name: "Leet",
		Grid: grid,
	}
	sudokuBoard.PoplateNeighborhood()
	final := sudokuBoard.Solve()

	return final.LeetOut()
}
