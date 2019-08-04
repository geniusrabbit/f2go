package f2go

import (
	"fmt"
	"io"
	"strings"
)

// Converter wrapper
type Converter struct {
	encoder Encoder
	reader  io.Reader
	writer  io.Writer
}

// NewConverter object
func NewConverter(r io.Reader, w io.Writer, enc Encoder) *Converter {
	return &Converter{
		encoder: enc,
		reader:  r,
		writer:  w,
	}
}

// Render template object
func (c *Converter) Render(template string, ctx *RenderContext) error {
	tpl, err := c.template("data.go")
	if err != nil {
		return err
	}
	return tpl.Render(c.writer, ctx)
}

func (c *Converter) template(name string) (_ *templateWrapper, err error) {
	switch name {
	case "data.go":
		tpl := &templateWrapper{dataEncoder: c.dataWriter}
		if c.isByteEncoder() {
			err = tpl.registerBeginTemplate(
				strings.Replace(dataGoTemplate, "const", "var", -1) + "[]byte{\n",
			)
			if err != nil {
				return nil, err
			}
			tpl.end = []byte("}\n")
		} else {
			if err = tpl.registerBeginTemplate(dataGoTemplate); err != nil {
				return nil, err
			}
		}
		return tpl, nil
	}
	return nil, fmt.Errorf(`'%s' undefined template`, name)
}

func (c *Converter) isByteEncoder() bool {
	return c.encoder == ByteEncoder
}

func (c *Converter) dataWriter(w io.Writer) error {
	return c.encoder.encoder()(c.reader, c.writer)
}

const dataGoTemplate = "package {{.PackageName}}\n\nconst {{.DataName}} = "
