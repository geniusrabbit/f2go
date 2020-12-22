package f2go

import (
	"bytes"
	"html/template"
	"io"
)

type templateEncoder func(w io.Writer) error

// RenderContext data
type RenderContext struct {
	Filename    string
	Filepath    string
	PackageName string
	DataName    string
}

type templateWrapper struct {
	begin       []byte
	beginTmp    *template.Template
	end         []byte
	dataEncoder templateEncoder
}

func (t *templateWrapper) registerBeginTemplate(tpl string) (err error) {
	t.beginTmp, err = template.New("tpl").Parse(tpl)
	return err
}

func (t *templateWrapper) Render(w io.Writer, ctx *RenderContext) (err error) {
	if len(t.begin) > 0 {
		err = t.prepareDataAdWrite(w, t.begin, ctx)
		if err != nil {
			return err
		}
	}
	if t.beginTmp != nil {
		if err = t.beginTmp.Execute(w, ctx); err != nil {
			return err
		}
	}
	if err = t.dataEncoder(w); err != nil {
		return err
	}
	if len(t.end) > 0 {
		err = t.prepareDataAdWrite(w, t.end, ctx)
	}
	return err
}

func (t *templateWrapper) prepareDataAdWrite(w io.Writer, data []byte, ctx *RenderContext) (err error) {
	data = t.prepareData(data, ctx)
	_, err = w.Write(data)
	return
}

func (t *templateWrapper) prepareData(data []byte, ctx *RenderContext) []byte {
	return t.replaceAll(data,
		[]byte(`{{.PackageName}}`), []byte(ctx.PackageName),
		[]byte(`{{.DataName}}`), []byte(ctx.DataName),
		[]byte(`{{.Filepath}}`), []byte(ctx.Filepath),
		[]byte(`{{.Filename}}`), []byte(ctx.Filename),
	)
}

func (t *templateWrapper) replaceAll(data []byte, items ...[]byte) []byte {
	for i := 0; i < len(items); i += 2 {
		data = bytes.ReplaceAll(data, items[i], items[i+1])
	}
	return data
}
