package main

import (
	"testing"
)

func Test_String(t *testing.T) {
	cases := []struct {
		in   Store
		want string
	}{
		{Store{}, ""},
		{Store{"a": "b"}, "a = b\n"},
		{Store{"a": "b", "c": "d"}, "a = b\nc = d\n"},
	}
	for _, c := range cases {
		s := Store(c.in)
		if got := s.String(); got != c.want {
			t.Errorf("Got %q but wanted %q", got, c.want)
		}
	}
}
