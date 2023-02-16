package snippets

import (
	"testing"

	"github.com/kyledinh/btk-go/pkg/moxutil"
)

func BenchmarkGenerateAKey_5(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = moxutil.GetRandomHyphenedKeyByLimit(5)
	}
}
