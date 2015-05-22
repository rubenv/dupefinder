package dupefinder

import (
	"strings"
	"testing"
)

func TestCatalogParse(t *testing.T) {
	in := `
# A comment and an empty line:

00e3261a6e0d79c329445acd540fb2b07187a0dcf6017065c8814010283ac67f test
5891b5b522d5df086d0ff0b110fbd9d21bb4fc7163af34d08286a2e846f6be03 a/b.txt
`

	reader := strings.NewReader(in)

	entries, err := ParseCatalogReader(reader)
	if err != nil {
		t.Error(err)
	}
	if len(entries) != 2 {
		t.Errorf("Unexpected number of entries: %d", len(entries))
	}
	if entries["00e3261a6e0d79c329445acd540fb2b07187a0dcf6017065c8814010283ac67f"] != "test" {
		t.Error("Bad entry")
	}
	if entries["5891b5b522d5df086d0ff0b110fbd9d21bb4fc7163af34d08286a2e846f6be03"] != "a/b.txt" {
		t.Error("Bad entry")
	}
}
