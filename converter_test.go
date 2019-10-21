package f2go

import (
	"bytes"
	"strings"
	"testing"
)

func Test_ConverterRender(t *testing.T) {
	var (
		buff bytes.Buffer
		rctx = RenderContext{PackageName: "test", DataName: "data"}
		conv = NewConverter(strings.NewReader("data"), &buff, ByteEncoder)
	)
	if err := conv.Render("data.go", &rctx); err != nil {
		t.Error(err)
	}
	if buff.Len() < 1 {
		t.Errorf("invalid response")
	}
}
