package ex1689_test

import (
	"testing"

	ex1689 "github.com/hulla-hoop/leetcodeex/1689"
)

func TestFinal(t *testing.T) {

	data := []struct {
		name string
		n    string
		want int
	}{
		{
			name: "1",
			n:    "32",
			want: 3,
		},
		{
			name: "2",
			n:    "82734",
			want: 8,
		}, {
			name: "3",
			n:    "27346209830709182346",
			want: 9,
		},
	}

	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			result := ex1689.MinPartitions(d.n)
			if result != d.want {
				t.Errorf("Expected %v , got %v ", d.want, result)
			}
		})
	}

}
