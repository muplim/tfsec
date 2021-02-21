package parser

import (
	"path/filepath"
	"testing"
)

func Test_LoadNestedModules(t *testing.T) {

	parser := New(filepath.Dir("../../../../tfsec-error/foo/bar/entrypoint/"), "")
	_, err := parser.ParseDirectory()
	if err != nil {
		t.Fatal(err)
	}

}
