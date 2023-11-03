package ex1769_test

import (
	"reflect"
	"testing"

	ex1769 "github.com/hulla-hoop/leetcodeex/1769"
)

func TestFinal(t *testing.T) {

	data := []struct {
		name  string
		boxes string
		want  []int
	}{
		{
			name:  "1",
			boxes: "110",
			want:  []int{1, 1, 3},
		},
		{
			name:  "2",
			boxes: "001011",
			want:  []int{11, 8, 5, 4, 3, 4},
		},
	}

	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			result := ex1769.MinOperations(d.boxes)
			if reflect.DeepEqual(result, d.want) == false {
				t.Errorf("Expected %v , got %v ", d.want, result)
			}
		})
	}

}
