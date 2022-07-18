package sandbox_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Square struct {
	Edges       [][]int `json:"edges"`       // 4 for a square
	Matchsticks []int   `json:"matchsticks"` // matchsticks of varying lengths
}

func (s *Square) isValidSquare() (bool, error) {

	var (
		totalMatchestickLen int
		edgeLen             int
		err                 error
	)

	// FIND THE EDGE LENGTHS
	for _, val := range s.Matchsticks {
		totalMatchestickLen += val
	}
	edgeLen = totalMatchestickLen / 4
	remainder := totalMatchestickLen % 4
	if remainder > 0 {
		return false, fmt.Errorf("The edges of %f are not integers", float32(totalMatchestickLen)/4.0)
	}

	// TRY TO MAKE EACH EDGE WITH MATCHES PROVIDED
	availableSticks := s.Matchsticks
	for i, _ := range s.Edges {
		var unusedSticks []int

		for _, matchstick := range availableSticks {
			_ = matchstick

			//IF EDGE IS EMPTY, ADD A MATCH
			if len(s.Edges[i]) == 0 {
				s.Edges[i] = append(s.Edges[i], matchstick)
				continue
			}

			// IF THE EDGE AND MATCH CAN FIT, THEN ADD TO EDGE AND REMOVE FROM AVAILABLE
			if (sum(s.Edges[i]) + matchstick) <= edgeLen {
				s.Edges[i] = append(s.Edges[i], matchstick)
				continue
			}
			unusedSticks = append(unusedSticks, matchstick)
		}

		if sum(s.Edges[i]) != edgeLen {
			return false, fmt.Errorf("Could not create edge with correct length of %d", edgeLen)
		}
		// RESET AVAILABLE TO WHAT IS LEFT/UNSUED
		availableSticks = unusedSticks
	}
	return true, err
}

func sum(array []int) int {
	result := 0
	for _, v := range array {
		result += v
	}
	return result
}

// WILL WORK FOR SIMPLE CASES, THAT REQUIRE ONE PASS, WILL NOT SURE MORE COMPLEX INPUTS OF MATCHES??
func Test_Matchesticks(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		square    Square
		wantValid bool
		wantError error
	}{
		{
			name: "Test should pass with all 4 edges as 4",
			square: Square{
				Matchsticks: []int{4, 4, 4, 4},
			},
			wantValid: true,
			wantError: nil,
		},
		{
			name: "Test should pass ...",
			square: Square{
				Matchsticks: []int{4, 4, 4, 2, 2},
			},
			wantValid: true,
			wantError: nil,
		},
		{
			name: "Test should pass ...",
			square: Square{
				Matchsticks: []int{9, 6, 3, 9, 3, 3, 3},
			},
			wantValid: true,
			wantError: nil,
		},
		{
			name: "Test should pass ...",
			square: Square{
				Matchsticks: []int{9, 1, 8, 3, 3, 6, 1, 2, 2, 1},
			},
			wantValid: true,
			wantError: nil,
		},
		{
			name: "Should also fail",
			square: Square{
				Matchsticks: []int{1, 2, 3, 4, 5, 6, 7, 8, 8},
			},
			wantValid: false,
		},
		{
			name: "Total of matches does not divide evenly",
			square: Square{
				Matchsticks: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			},
			wantValid: false,
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d: %s", i, tt.name), func(t *testing.T) {
			tt.square.Edges = [][]int{{}, {}, {}, {}}
			result, err := tt.square.isValidSquare()
			_ = err
			fmt.Printf("======= %s \n", tt.name)
			fmt.Printf("Input %v \n", tt.square.Matchsticks)
			fmt.Printf("Edges: %v \n\n", tt.square.Edges)
			// assert.Equal(t, nil, err)
			assert.Equal(t, tt.wantValid, result)
		})
	}
}

// VERSION 2 OF TESTING SUITE/WITH JSON INPUTS
type TestableSquare struct {
	Square
	Title     string `json:"title"`
	WantValid bool   `json:"want_valid"`
	WantError string `json:"want_error"`
}

// WILL WORK FOR SIMPLE CASES, THAT REQUIRE ONE PASS, WILL NOT SURE MORE COMPLEX INPUTS OF MATCHES??
func Test_Able_Matchesticks(t *testing.T) {
	t.Parallel()

	var tests []TestableSquare
	json.Unmarshal(bufTestableSuite, &tests)

	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d: %s", i, tt.Title), func(t *testing.T) {
			tt.Square.Edges = [][]int{{}, {}, {}, {}}
			result, err := tt.Square.isValidSquare()
			_ = err
			fmt.Printf("======= %s \n", tt.Title)
			fmt.Printf("Input %v \n", tt.Square.Matchsticks)
			fmt.Printf("Edges: %v \n\n", tt.Square.Edges)
			// assert.Equal(t, nil, err)
			assert.Equal(t, tt.WantValid, result)
		})
	}
}

var bufTestableSuite = []byte(`
[
 	{
 		"title": "Test should pass with all 4 edges as 4",
	 	"want_valid": true,
	 	"want_error": "",
 		"square": { "matchsticks": {4, 4, 4, 4} }
	},
 	{
 		"title": "Test total of 9 per egde",
	 	"want_valid": true,
	 	"want_error": "",
 		"square": { "matchsticks": {9, 1, 3, 7, 4, 1, 6, 3, 2} }
	},
 	{
 		"title": "Test should pass with all 4 edges as 4",
	 	"want_valid": false,
	 	"want_error": "",
 		"square": { "matchsticks": {4, 1, 3, 4, 5} }
	}
]
`)
