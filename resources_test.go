package resources

import (
	"strings"
	"testing"

	"github.com/cugu/go-resources/testdata/generated"
)

//go:generate go build -o testdata/resources github.com/cugu/go-resources/cmd/resources
//go:generate testdata/resources -declare -package generated -output testdata/generated/store_prod.go  testdata/*.txt testdata/*.sql

func TestGenerated(t *testing.T) {
	for _, tt := range []struct {
		name    string
		snippet string
	}{
		{name: "test.txt", snippet: "this is test.txt"},
		{name: "patrick.txt", snippet: "no, this is patrick!"},
		{name: "query.sql", snippet: `drop table "files";`},
	} {
		t.Run(tt.name, func(t *testing.T) {
			content, ok := generated.FS.Files["/testdata/"+tt.name]

			if !ok {
				t.Fatalf("expected no error opening file")
			}

			if !strings.Contains(string(content), tt.snippet) {
				t.Errorf("expected to find snippet %q in file", tt.snippet)
			}
		})
	}
}
