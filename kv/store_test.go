package main

import (
	"reflect"
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

func Test_Merge(t *testing.T) {
	cases := []struct {
		base Store
		in   Store
		want Store
	}{
		{Store{}, Store{}, Store{}},
		{Store{}, Store{"a": "b"}, Store{"a": "b"}},
		{Store{"a": "b"}, Store{}, Store{"a": "b"}},
		{Store{"a": "b"}, Store{"c": "d"}, Store{"a": "b", "c": "d"}},
		{Store{"a": "b"}, Store{"a": "c"}, Store{"a": "c"}},
	}
	for _, c := range cases {
		s := Store(c.base)
		s.Merge(c.in)
		if !reflect.DeepEqual(s, c.want) {
			t.Errorf("Got %q but wanted %q", s, c.want)
		}
	}
}
