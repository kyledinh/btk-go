package some_test

import (
	"reflect"
	"strings"
	"testing"
)

func Split(s, sep string) []string {
	var result []string
	i := strings.Index(s, sep)
	for i > -1 {
		result = append(result, s[:i])
		s = s[i+len(sep):]
		i = strings.Index(s, sep)
	}
	return append(result, s)
}

func TestSplit(t *testing.T) {
	tests := []struct {
		input string
		sep   string
		want  []string
	}{
		{input: "a/b/c", sep: "/", want: []string{"a", "b", "c"}},
		{input: "a/b/c", sep: ",", want: []string{"a/b/c"}},
		{input: "abc", sep: "/", want: []string{"abc"}},
	}

	for _, tc := range tests {
		got := Split(tc.input, tc.sep)
		if !reflect.DeepEqual(tc.want, got) {
			t.Fatalf("expected: %v, got: %v", tc.want, got)
		}
	}
}
