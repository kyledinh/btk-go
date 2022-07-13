package sandbox_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Square struct {
	Edges   [][]int // 4 for a square
	Matches []int   // matches of varying lengths
}

func (s *Square) isValidSquare() (bool, error) {

	var totalMatchesLength int
	var edgeLen int
	var err error
	_ = edgeLen // arggghhh this linter!!!

	// FIND THE EDGE LENGTHS
	for _, val := range s.Matches {
		totalMatchesLength += val
	}
	edgeLen = totalMatchesLength / 4

	remainder := totalMatchesLength % 4
	if remainder > 0 {
		return false, fmt.Errorf("The edges of %f are not integers", float32(totalMatchesLength)/4.0)
	}

	// TRY TO MAKE EACH EDGE WITH MATCHES PROVIDED
	availableMatches := s.Matches
	for i := 0; i < 4; i++ {
		var unusedMatches []int

		for _, match := range availableMatches {
			_ = match

			//IF EDGE IS EMPTY, ADD A MATCH
			if len(s.Edges[i]) == 0 {
				s.Edges[i] = append(s.Edges[i], match)
				continue
			}

			// IF THE EDGE AND MATCH CAN FIT, THEN ADD TO EDGE AND REMOVE FROM AVAILABLE
			if (sum(s.Edges[i]) + match) <= edgeLen {
				s.Edges[i] = append(s.Edges[i], match)
				continue
			}
			unusedMatches = append(unusedMatches, match)
		}

		if sum(s.Edges[i]) != edgeLen {
			return false, fmt.Errorf("Could not create edge with correct length of %d", edgeLen)
		}
		// RESET AVAILABLE TO WHAT IS LEFT
		availableMatches = unusedMatches

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
func Test_Matches(t *testing.T) {
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
				Matches: []int{4, 4, 4, 4},
			},
			wantValid: true,
			wantError: nil,
		},
		{
			name: "Test should pass ...",
			square: Square{
				Matches: []int{4, 4, 4, 2, 2},
			},
			wantValid: true,
			wantError: nil,
		},
		{
			name: "Test should pass ...",
			square: Square{
				Matches: []int{9, 6, 3, 9, 3, 3, 3},
			},
			wantValid: true,
			wantError: nil,
		},
		{
			name: "Test should pass ...",
			square: Square{
				Matches: []int{9, 1, 8, 3, 3, 6, 1, 2, 2, 1},
			},
			wantValid: true,
			wantError: nil,
		},
		{
			name: "Should also fail",
			square: Square{
				Matches: []int{1, 2, 3, 4, 5, 6, 7, 8, 8},
			},
			wantValid: false,
		},
		{
			name: "Total of matches does not divide evenly",
			square: Square{
				Matches: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
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
			fmt.Printf("Input %v \n", tt.square.Matches)
			fmt.Printf("Edges: %v \n\n", tt.square.Edges)
			// assert.Equal(t, nil, err)
			assert.Equal(t, tt.wantValid, result)
		})
	}
}
