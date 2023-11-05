package ex2120_test

import (
	"reflect"
	"testing"

	ex2120 "github.com/hulla-hoop/leetcodeex/2120"
)

func TestFinal(t *testing.T) {

	data := []struct {
		name     string
		n        int
		startPos []int
		s        string
		want     []int
	}{
		{
			name:     "1",
			n:        3,
			startPos: []int{0, 1},
			s:        "RRDDLU",
			want:     []int{1, 5, 4, 3, 1, 0},
		},
		{
			name:     "2",
			n:        2,
			startPos: []int{1, 1},
			s:        "LURD",
			want:     []int{4, 1, 0, 0},
		}, {
			name:     "3",
			n:        0,
			startPos: []int{0, 0},
			s:        "LRUD",
			want:     []int{0, 0, 0, 0},
		},
	}

	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			result := ex2120.ExecuteInstructions(d.n, d.startPos, d.s)
			if reflect.DeepEqual(result, d.want) == false {
				t.Errorf("Expected %v , got %v ", d.want, result)
			}
		})
		t.Run(d.name, func(t *testing.T) {
			result := ex2120.ExecuteInstructionsTwo(d.n, d.startPos, d.s)
			if reflect.DeepEqual(result, d.want) == false {
				t.Errorf("Expected %v , got %v ", d.want, result)
			}
		})
	}

}
