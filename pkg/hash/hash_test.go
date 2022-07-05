package hash

import (
	"crypto/md5"
	"crypto/sha256"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func HashMd5(s string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(s)))
}

func HashSha256(s string) string {
	return fmt.Sprintf("%x", sha256.Sum256([]byte(s)))
}

func TestMd5(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		alpha    string
		expected string
	}{
		{
			name:     "Foo as MD5 Hash",
			alpha:    "Foo",
			expected: "1356c67d7ad1638d816bfb822dd2c25d",
		},
		{
			name:     "Foo as MD5 Hash",
			alpha:    "Foo",
			expected: "1356c67d7ad1638d816bfb822dd2c25d",
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d: %s", i, tt.name), func(t *testing.T) {
			result := HashMd5(tt.alpha)
			assert.Equal(t, tt.expected, result, "Failed to MD5Hash!")
		})
	}
}

func TestSha256(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		alpha    string
		expected string
	}{
		{
			name:     "Foo as Sha256 Hash",
			alpha:    "Foo",
			expected: "1cbec737f863e4922cee63cc2ebbfaafcd1cff8b790d8cfd2e6a5d550b648afa",
		},
		{
			name:     "Foo as Sha256 Hash",
			alpha:    "Foo",
			expected: "1cbec737f863e4922cee63cc2ebbfaafcd1cff8b790d8cfd2e6a5d550b648afa",
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d: %s", i, tt.name), func(t *testing.T) {
			result := HashSha256(tt.alpha)
			assert.Equal(t, tt.expected, result, "Failed to Sha256Hash!")
		})
	}
}
