package ex_test

import (
	"testing"

	ex "github.com/hulla-hoop/leetcodeex/2011"
)

func TestFinal(t *testing.T) {

	data := []struct {
		name       string
		operations []string
		want       int
	}{
		{
			name:       "1",
			operations: []string{"--X", "X++", "X++"},
			want:       1,
		},
		{
			name:       "2",
			operations: []string{"--X", "X++", "X++", "--X", "X++", "X++"},
			want:       2,
		},
	}

	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			result := ex.FinalValueAfterOperations(d.operations)
			if result != d.want {
				t.Errorf("Expected %v , got %v ", d.want, result)
			}
		})
	}

}
