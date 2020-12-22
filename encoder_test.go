package f2go

import (
	"bytes"
	"io"
	"strings"
	"testing"
)

func Test_byteEncoder(t *testing.T) {
	var tests = []struct {
		r      io.Reader
		result string
	}{
		{
			r:      strings.NewReader("test"),
			result: `'\x74','\x65','\x73','\x74'`,
		},
		{
			r:      bytes.NewBuffer([]byte{'\x74', '\x65', '\x73', '\x74'}),
			result: `'\x74','\x65','\x73','\x74'`,
		},
	}

	var buff bytes.Buffer
	for _, test := range tests {
		buff.Reset()
		if err := byteEncoder(test.r, &buff); err != nil {
			t.Error(err)
		}
		if buff.String() != test.result {
			t.Errorf("invalid byte response `%s` must be `%s`", buff.String(), test.result)
		}
	}
}

func Test_stringEncoder(t *testing.T) {
	var tests = []struct {
		r      io.Reader
		result string
	}{
		{
			r:      strings.NewReader("test"),
			result: `"test"`,
		},
		{
			r:      bytes.NewBuffer([]byte{'\x74', '\x65', '\x73', '\x74'}),
			result: `"test"`,
		},
		{
			r:      strings.NewReader("the japan news【モスクワ＝畑武尊】!"),
			result: `"the japan news\xe3\x80\x90\xe3\x83\xa2\xe3\x82\xb9\xe3\x82\xaf\xe3\x83\xaf\xef\xbc\x9d\xe7\x95\x91\xe6\xad\xa6\xe5\xb0\x8a\xe3\x80\x91\x21"`,
		},
		{
			r:      bytes.NewBuffer([]byte("New\n\tline")),
			result: `"New\n\tline"`,
		},
		{
			r:      bytes.NewBuffer([]byte("{\"text\":\"test text\\nnewline\"}")),
			result: `"{\"text\":\"test text\\nnewline\"}"`,
		},
	}

	var buff bytes.Buffer
	for _, test := range tests {
		buff.Reset()
		if err := stringEncoder(test.r, &buff); err != nil {
			t.Error(err)
		}
		if buff.String() != test.result {
			t.Errorf("invalid string response `%s` must be `%s`", buff.String(), test.result)
		}
	}
}
