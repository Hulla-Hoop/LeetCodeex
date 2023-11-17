package ex1061_test

import (
	"testing"

	ex1061 "github.com/hulla-hoop/leetcodeex/1061"
)

func TestFinal(t *testing.T) {

	data := []struct {
		name    string
		s1      string
		s2      string
		baseStr string
		want    string
	}{
		/* {
			name:    "1",
			s1:      "parker",
			s2:      "morris",
			baseStr: "parser",
			want:    "makkek",
		},
		{
			name:    "2",
			s1:      "hello",
			s2:      "world",
			baseStr: "hold",
			want:    "hdld",
		}, */{
			name:    "3",
			s1:      "leetcode",
			s2:      "programs",
			baseStr: "sourcecode",
			want:    "aauaaaaada",
		}, /* {
			name:    "4",
			s1:      "abc",
			s2:      "cde",
			baseStr: "eed",
			want:    "aab",
		}, */
	}

	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			result := ex1061.Smollest(d.s1, d.s2, d.baseStr)
			if result != d.want {
				t.Errorf(" Base %v, Expected %v , got %v ", d.baseStr, d.want, result)
			}
		})
	}

}
