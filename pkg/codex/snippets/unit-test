package moxutil_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/kyledinh/btk-go/pkg/moxutil"
	"github.com/stretchr/testify/assert"
)

func Test_GetRandomHyphenatedKeyByLimit(t *testing.T) {
	t.Parallel()

	type Want struct {
		sliceLen  int
		hyphenCnt int
	}

	tests := []struct {
		name  string
		limit int
		want  Want
	}{
		{
			name:  "Limit 1",
			limit: 1,
			want: Want{
				sliceLen:  1,
				hyphenCnt: 0,
			},
		},
		{
			name:  "Limit 5",
			limit: 5,
			want: Want{
				sliceLen:  5,
				hyphenCnt: 4,
			},
		},
		{
			name:  "Above the limit of 10 with 13",
			limit: 13,
			want: Want{
				sliceLen:  4,
				hyphenCnt: 3,
			},
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d: %s", i, tt.name), func(t *testing.T) {
			result := moxutil.GetRandomHyphenedKeyByLimit(tt.limit)
			wordsArr := strings.Split(result, "-")
			assert.Equal(t, tt.want.sliceLen, len(wordsArr), "Wrong number of keys returned")
			assert.Equal(t, tt.want.hyphenCnt, strings.Count(result, "-"), "Wrong number of hyphens returned")
		})
	}
}
