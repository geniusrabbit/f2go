package f2go

import (
	"bytes"
	"strings"
	"unicode"
)

// VariableNamePrepare from some string
func VariableNamePrepare(name string) string {
	var (
		names = strings.FieldsFunc(strings.ToLower(name), func(r rune) bool {
			return !unicode.IsNumber(r) && !unicode.IsLetter(r)
		})
		buff bytes.Buffer
	)
	for _, name := range names {
		switch name {
		case "id", "http", "sql", "xml", "json", "yaml", "tcp", "udp":
			buff.WriteString(strings.ToUpper(name))
		default:
			buff.WriteString(strings.Title(name))
		}
	}
	return buff.String()
}
