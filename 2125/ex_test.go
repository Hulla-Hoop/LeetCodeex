package ex2125_test

import (
	"testing"

	ex2125 "github.com/hulla-hoop/leetcodeex/2125"
)

func TestFinal(t *testing.T) {

	data := []struct {
		name string
		bank []string
		want int
	}{
		{
			name: "1",
			bank: []string{"011001", "000000", "010100", "001000"},
			want: 8,
		},
		{
			name: "1",
			bank: []string{"000", "111", "000"},
			want: 0,
		},
	}

	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			result := ex2125.NumberOfBeams(d.bank)
			if result != d.want {
				t.Errorf("Expected %v , got %v ", d.want, result)
			}
		})
	}

}
