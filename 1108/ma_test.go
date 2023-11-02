package ex1108_test

import (
	"testing"

	ex1108 "github.com/hulla-hoop/leetcodeex/1108"
)

func TestFinal(t *testing.T) {

	data := []struct {
		name    string
		address string
		want    string
	}{
		{
			name:    "1",
			address: "1.1.1.1",
			want:    "1[.]1[.]1[.]1",
		},
		{
			name:    "2",
			address: "21.31.251.1",
			want:    "21[.]31[.]251[.]1",
		},
	}

	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			result := ex1108.DefangIPaddr(d.address)
			if result != d.want {
				t.Errorf("Expected %v , got %v ", d.want, result)
			}
		})
	}

}
