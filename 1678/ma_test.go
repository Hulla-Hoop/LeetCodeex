package ex1678_test

import (
	"testing"

	ex1678 "github.com/hulla-hoop/leetcodeex/1678"
)

func TestFinal(t *testing.T) {

	data := []struct {
		name    string
		command string
		want    string
	}{
		{
			name:    "1",
			command: "G()(al)",
			want:    "Goal",
		},
		{
			name:    "2",
			command: "G()()()()(al)",
			want:    "Gooooal",
		}, {
			name:    "3",
			command: "(al)G(al)()()G",
			want:    "alGalooG",
		},
	}

	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			result := ex1678.Interpret(d.command)
			if result != d.want {
				t.Errorf("Expected %v , got %v ", d.want, result)
			}
		})
	}

}
