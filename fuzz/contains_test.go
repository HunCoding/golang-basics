package fuzz

import (
	"fmt"
	"strings"
	"testing"
)

func TestMyIndexAny(t *testing.T) {
	tests := []struct {
		s, chars string
		want     int
	}{
		{"", "", -1},
		{"", "a", -1},
		{"", "abc", -1},
		{"a", "", -1},
		{"a", "a", 0},
		{"abc", "xyz", -1},
		{"abc", "xcz", 2},
		{"abøc", "xøyz", 2},
		{"a\x20c", "xyz\x20", 1},
		{"\x00\x01\x02", "\x00", 0},
		{"\x70\x71\x72", "\x73", -1},
		{"sRegExp*", ".(|)*+?^$[]", 7},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("(%q, %q)", test.s, test.chars), func(t *testing.T) {
			if got := MyIndexAny(test.s, test.chars); got != test.want {
				t.Errorf("%d, want %d", got, test.want)
			}
		})
	}
}

func FuzzMyIndexAny(f *testing.F) {
	tests := []struct {
		s, chars string
		want     int
	}{
		{"", "", -1},
		{"", "a", -1},
		{"", "abc", -1},
		{"a", "", -1},
		{"a", "a", 0},
		{"abc", "xyz", -1},
		{"abc", "xcz", 2},
		{"abøc", "xøyz", 2},
		{"a\x20c", "xyz\x20", 1},
		{"\x00\x01\x02", "\x00", 0},
		{"\x70\x71\x72", "\x73", -1},
		{"sRegExp*", ".(|)*+?^$[]", 7},
		{"sRegExp*", ".(|)*+?^$[]", 7},
	}
	for _, test := range tests {
		f.Add(test.s, test.chars)
	}

	f.Fuzz(func(t *testing.T, s, chars string) {
		if want, got := MyIndexAny(s, chars), strings.IndexAny(s, chars); got != want {
			t.Errorf("(%q, %q), MyIndex=%d, strings.IndexAny=%d",
				s, chars, want, got)
		}
	})
}
