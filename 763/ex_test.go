package ex763_test

import (
	"reflect"
	"testing"

	ex763 "github.com/hulla-hoop/leetcodeex/763"
)

func TestFinal(t *testing.T) {

	data := []struct {
		name string
		s    string
		want []int
	}{
		{
			name: "1",
			s:    "ababcbacadefegdehijhklij",
			want: []int{9, 7, 8},
		},
		{
			name: "2",
			s:    "eccbbbbdec",
			want: []int{10},
		},
	}

	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			result := ex763.PartitionLabels(d.s)
			if reflect.DeepEqual(result, d.want) == false {
				t.Errorf("Expected %v , got %v ", d.want, result)
			}
		})
	}

}
