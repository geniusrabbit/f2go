package f2go

import (
	"fmt"
	"io"
)

type Encoder int

const (
	ByteEncoder Encoder = iota
	StringEncoder
)

func (enc Encoder) encoder() encoder {
	if enc == StringEncoder {
		return stringEncoder
	}
	return byteEncoder
}

type encoder func(io.Reader, io.Writer) error
type byteFmt func(v byte) []byte

// ByteEncoder converts file to the byte style array
// example: '\x74', '\x65', '\x73', '\x74'
func byteEncoder(r io.Reader, w io.Writer) error {
	return _encoder(r, w, ',', byteHexFmt)
}

// StringEncoder converts file to the string style encoding
func stringEncoder(r io.Reader, w io.Writer) (err error) {
	if _, err = w.Write([]byte(`"`)); err != nil {
		return err
	}
	if err = _encoder(r, w, '\x00', byteStrFmt); err != nil {
		return err
	}
	_, err = w.Write([]byte(`"`))
	return err
}

func _encoder(r io.Reader, w io.Writer, splitSimbol byte, f byteFmt) error {
	buff := make([]byte, 128)
	started := false
	for {
		n, err := r.Read(buff)
		if n > 0 {
			for i := 0; i < n; i++ {
				if started && splitSimbol != '\x00' {
					_, err = w.Write([]byte{splitSimbol})
				}
				if err == nil {
					_, err = w.Write(f(buff[i]))
				}
				if err != nil {
					return err
				}
				started = true
			}
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
	}
	return nil
}

func byteHexFmt(v byte) []byte {
	return []byte(fmt.Sprintf("'\\x%02x'", v))
}

// typed ASCII symbols prints without changes
func byteStrFmt(v byte) []byte {
	switch {
	case v == '\n':
		return []byte{'\\', 'n'}
	case v == '\t':
		return []byte{'\\', 't'}
	case v == ' ':
		return []byte{' '}
	case v == '"':
		return []byte{'\\', '"'}
	case v > 0x32 && v < 0x7f:
		return []byte{v}
	}
	return []byte(fmt.Sprintf("\\x%02x", v))
}
