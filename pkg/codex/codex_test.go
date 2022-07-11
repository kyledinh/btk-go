package codex_test

import (
	"fmt"
	"testing"

	"github.com/kyledinh/btk-go/pkg/codex"
	"github.com/stretchr/testify/assert"
)

func Test_KeywordFromFilename(t *testing.T) {
	t.Parallel()

	type Want struct {
		keyword string
	}

	tests := []struct {
		name     string
		filename string
		want     Want
	}{
		{
			name:     "Limit 1",
			filename: "diretory/somefile.txt",
			want: Want{
				keyword: "somefile.txt",
			},
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d: %s", i, tt.name), func(t *testing.T) {
			assert.Equal(t, tt.want.keyword, codex.KeywordFromFilename(tt.filename))
		})
	}
}
