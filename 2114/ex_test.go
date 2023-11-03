package ex2114_test

import (
	"testing"

	ex2114 "github.com/hulla-hoop/leetcodeex/2114"
)

func TestEx(t *testing.T) {

	data := []struct {
		name      string
		sentences []string
		want      int
	}{
		{
			name:      "1",
			sentences: []string{"alice and bob love leetcode", "i think so too", "this is great thanks very much"},
			want:      6,
		},
		{
			name:      "2",
			sentences: []string{"please wait", "continue to fight", "continue to win"},
			want:      3,
		},
	}

	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			result := ex2114.MostWordsFound(d.sentences)
			if result != d.want {
				t.Errorf("Expected %v , got %v ", d.want, result)
			}
		})
	}

}
