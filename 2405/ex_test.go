package ex2405_test

import (
	"testing"

	ex2405 "github.com/hulla-hoop/leetcodeex/2405"
)

func TestFinal(t *testing.T) {

	data := []struct {
		name string
		s    string
		want int
	}{
		{
			name: "1",
			s:    "abacaba",
			want: 4,
		},
		{
			name: "2",
			s:    "ssssss",
			want: 6,
		}, {
			name: "3",
			s:    "yzubfsiypfrepcfftiov",
			want: 4,
		},
	}

	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			result := ex2405.PartitionString(d.s)
			if result != d.want {
				t.Errorf("Expected %v , got %v ", d.want, result)
			}
		})
	}

}
