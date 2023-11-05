package ex2391_test

import (
	"testing"

	ex2391 "github.com/hulla-hoop/leetcodeex/2391"
)

func TestFinal(t *testing.T) {

	data := []struct {
		name    string
		garbage []string
		travel  []int
		want    int
	}{
		{
			name:    "1",
			garbage: []string{"G", "P", "GP", "GG"},
			travel:  []int{2, 4, 3},
			want:    21,
		},
		{
			name:    "1",
			garbage: []string{"MMM", "PGM", "GP"},
			travel:  []int{3, 10},
			want:    37,
		},
	}

	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			result := ex2391.GarbageCollection(d.garbage, d.travel)
			if result != d.want {
				t.Errorf("Expected %v , got %v ", d.want, result)
			}
		})
	}

}
