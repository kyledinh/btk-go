package gencode_test

import (
	"fmt"
	"testing"

	"github.com/kyledinh/btk-go/pkg/gencode"
	"github.com/stretchr/testify/assert"
)

func Test_PascalFrom_snake_case(t *testing.T) {
	t.Parallel()

	type Want struct {
		Pascal string
	}

	tests := []struct {
		name string
		want Want
	}{
		{
			name: "this_is_a_snake_case",
			want: Want{
				Pascal: "ThisIsASnakeCase",
			},
		},
		{
			name: "snake",
			want: Want{
				Pascal: "Snake",
			},
		},
		{
			name: "Pascal",
			want: Want{
				Pascal: "Pascal",
			},
		},
		{
			name: "_is_a_snake_case",
			want: Want{
				Pascal: "IsASnakeCase",
			},
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d: %s", i, tt.name), func(t *testing.T) {
			result := gencode.PascalFrom_snake_case(tt.name)
			assert.Equal(t, tt.want.Pascal, result, "Did not format correcty")
		})
	}
}
