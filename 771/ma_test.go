package ex771_test

import (
	"testing"

	ex771 "github.com/hulla-hoop/leetcodeex/771"
)

func TestFinal(t *testing.T) {

	data := []struct {
		name   string
		jevels string
		stones string
		want   int
	}{
		{
			name:   "1",
			jevels: "aA",
			stones: "aAAAAAbbb",
			want:   6,
		},
		{
			name:   "2",
			jevels: "z",
			stones: "ZZ",
			want:   0,
		},
	}

	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			result := ex771.NumJewelsInStones(d.jevels, d.stones)
			if result != d.want {
				t.Errorf("Expected %v , got %v ", d.want, result)
			}
		})
	}

}
