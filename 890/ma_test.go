package ex890_test

import (
	"reflect"
	"testing"

	ex890 "github.com/hulla-hoop/leetcodeex/890"
)

func TestFinal(t *testing.T) {

	data := []struct {
		name    string
		words   []string
		pattern string
		want    []string
	}{
		{
			name:    "1",
			words:   []string{"abc", "deq", "mee", "aqq", "dkd", "ccc"},
			pattern: "abb",
			want:    []string{"mee", "aqq"},
		},
		{
			name:    "2",
			words:   []string{"a", "b", "c"},
			pattern: "a",
			want:    []string{"a", "b", "c"},
		},
		{
			name:    "3",
			words:   []string{"ktittgzawn", "dgphvfjniv", "gceqobzmis", "alrztxdlah", "jijuevoioe", "mawiizpkub", "onwpmnujos", "zszkptjgzj", "zwfvzhrucv", "isyaphcszn"},
			pattern: "zdqmjnczma",
			want:    nil,
		}, {
			name:    "4",
			words:   []string{"ttkxo", "tnron"},
			pattern: "drpsr",
			want:    []string{"tnron"},
		},
	}

	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			result := ex890.FindAndReplacePattern(d.words, d.pattern)
			if reflect.DeepEqual(result, d.want) == false {
				t.Errorf("Expected %v , got %v ", d.want, result)
			}
		})
	}

}
