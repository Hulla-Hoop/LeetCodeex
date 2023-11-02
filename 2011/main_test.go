package main_test

import "testing"

func TestMain(t *testing.T) {

	data := []struct {
		name       string
		operations []string
		want       int
	}{}

	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			result := finalValueAfterOperations(d.user)
			if ist != d.wantBool {
				t.Errorf("Expected %t | %v , got %t  | %v", d.wantBool, d.want, ist, result)
			}
		})
	}

}
