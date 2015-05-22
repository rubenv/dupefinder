package dupefinder

import (
	"os"
	"testing"
)

func TestGenerateFolders(t *testing.T) {
	err := Generate("catalog.txt", "invalid")
	if err == nil {
		t.Errorf("Expected an error: invalid is not an existing folder")
	}
}

func TestGenerateFinds(t *testing.T) {
	catalog := tempFilename(t)
	defer os.Remove(catalog)

	err := Generate(catalog, "fixtures/a")
	if err != nil {
		t.Error(err)
	}

	entries, err := ParseCatalog(catalog)
	if err != nil {
		t.Error(err)
	}
	if len(entries) != 2 {
		t.Errorf("Unexpected number of entries: %d", len(entries))
	}

	if entries["00e3261a6e0d79c329445acd540fb2b07187a0dcf6017065c8814010283ac67f"] != "fixtures/a/c/bla.txt" {
		t.Error("Bad entry")
	}
}
