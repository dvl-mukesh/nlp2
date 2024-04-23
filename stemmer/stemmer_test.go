package stemmer

import (
	"fmt"
	"testing"
	// "github.com/stretchr/testify/require"
)

var testCases = []struct {
	word string
	stem string
}{
	{"working", "work"},
	{"works", "work"},
	{"worked", "work"},
	{"work", "work"},
}

func TestStem(t *testing.T) {
	for _, tc := range testCases {
		name := fmt.Sprintf("%s:%s", tc.word, tc.stem)
		t.Run(name, func(t *testing.T) {
			stem := Stem(tc.word)
			if stem != tc.stem {
				t.Fatalf("extected %s, actual %s", tc.stem, stem)
			}
		})
	}
}
