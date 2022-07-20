package benchmark

import (
	"strings"
	"testing"
)

func join(strSlice []string) string {
	s := ""
	for _, str := range strSlice {
		s += str
	}
	return s
}

func BenchmarkJoin(b *testing.B) {
	join([]string{"aaa", "bbb", "ccc"})
}

func BenchmarkStringsJoin(b *testing.B) {
	strings.Join([]string{"aaa", "bbb", "ccc"}, "")
}
