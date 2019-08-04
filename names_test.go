package f2go

import "testing"

func Test_VariableNamePrepare(t *testing.T) {
	var tests = [][]string{
		[]string{"new_test_name", "NewTestName"},
		[]string{"http/server", "HTTPServer"},
		[]string{"Xml_Encoder Server", "XMLEncoderServer"},
		[]string{"TREE-BIN-CON", "TreeBinCon"},
	}

	for _, test := range tests {
		if name := VariableNamePrepare(test[0]); name != test[1] {
			t.Errorf("invalid name convesion: `%s` to `%s`", test[0], name)
		}
	}
}
