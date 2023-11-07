package ex1347_test

import (
	"testing"

	ex1347 "github.com/hulla-hoop/leetcodeex/1347"
)

func TestEx(t *testing.T) {

	data := []struct {
		name string
		s    string
		t    string
		want int
	}{
		{
			name: "1",
			s:    "bab",
			t:    "aba",
			want: 1,
		},
		{
			name: "2",
			s:    "leetcode",
			t:    "practice",
			want: 5,
		}, {
			name: "3",
			s:    "anagram",
			t:    "mangaar",
			want: 0,
		},
	}

	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			result := ex1347.MinSteps(d.s, d.t)
			if result != d.want {
				t.Errorf("Expected %v , got %v ", d.want, result)
			}
		})
	}

}
