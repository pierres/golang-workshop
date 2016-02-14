package main

import (
	"bytes"
	"reflect"
	"strings"
	"testing"
)

func Test_String(t *testing.T) {
	cases := []struct {
		in   Data
		want string
	}{
		{Data{}, ""},
		{Data{"a": "b"}, "a = b\n"},
		{Data{"a": "b", "c": "d"}, "a = b\nc = d\n"},
	}
	for _, c := range cases {
		s := NewStore(c.in)
		if got := s.String(); got != c.want {
			t.Errorf("Got %q but wanted %q", got, c.want)
		}
	}
}

func Test_Merge(t *testing.T) {
	cases := []struct {
		base Data
		in   Data
		want Data
	}{
		{Data{}, Data{}, Data{}},
		{Data{}, Data{"a": "b"}, Data{"a": "b"}},
		{Data{"a": "b"}, Data{}, Data{"a": "b"}},
		{Data{"a": "b"}, Data{"c": "d"}, Data{"a": "b", "c": "d"}},
		{Data{"a": "b"}, Data{"a": "c"}, Data{"a": "c"}},
	}
	for _, c := range cases {
		s := NewStore(c.base)
		s.Merge(c.in)
		if !reflect.DeepEqual(s, NewStore(c.want)) {
			t.Errorf("Got %q but wanted %q", s, c.want)
		}
	}
}

func Test_Read(t *testing.T) {
	cases := []struct {
		in   string
		want Data
	}{
		{"{}", Data{}},
		{"{\"a\":\"b\"}", Data{"a": "b"}},
		{"{\"a\":\"b\",\"c\":\"d\"}", Data{"a": "b", "c": "d"}},
	}
	s := NewStore(Data{})
	for _, c := range cases {
		s.Read(strings.NewReader(c.in))
		if !reflect.DeepEqual(s, NewStore(c.want)) {
			t.Errorf("Got %q but wanted %q", s, c.want)
		}
	}
}

func Test_Write(t *testing.T) {
	cases := []struct {
		in   Data
		want string
	}{
		{Data{}, "{}"},
		{Data{"a": "b"}, "{\"a\":\"b\"}"},
		{Data{"a": "b", "c": "d"}, "{\"a\":\"b\",\"c\":\"d\"}"},
	}
	for _, c := range cases {
		s := NewStore(c.in)
		out := new(bytes.Buffer)
		s.Write(out)
		if out.String() != c.want {
			t.Errorf("Got %q but wanted %q", out, c.want)
		}
	}
}

func Test_Filter(t *testing.T) {
	cases := []struct {
		in     Data
		filter []string
		want   Data
	}{
		{Data{}, []string{}, Data{}},
		{Data{"a": "b"}, []string{"a"}, Data{"a": "b"}},
		{Data{"a": "b", "c": "d"}, []string{"a"}, Data{"a": "b"}},
		{Data{"a": "b", "c": "d"}, []string{"c"}, Data{"c": "d"}},
		{Data{"a": "b", "c": "d"}, []string{"d"}, Data{}},
	}
	for _, c := range cases {
		s := NewStore(c.in)
		got := s.Filter(c.filter)
		if !reflect.DeepEqual(got, NewStore(c.want)) {
			t.Errorf("Got %q but wanted %q", got, c.want)
		}
	}
}

func Benchmark_ConcurrentWrites(b *testing.B) {
	doneChannel := make(chan bool)
	s := NewStore(Data{})
	out := new(bytes.Buffer)

	for i := 0; i < b.N; i++ {
		go func() {
			s.Write(out)
			doneChannel <- true
		}()
	}

	for i := 0; i < b.N; i++ {
		<-doneChannel
	}
}

func Benchmark_ConcurrentReads(b *testing.B) {
	doneChannel := make(chan bool)
	s := NewStore(Data{})

	for i := 0; i < b.N; i++ {
		go func() {
			s.Read(strings.NewReader(""))
			doneChannel <- true
		}()
	}

	for i := 0; i < b.N; i++ {
		<-doneChannel
	}
}
