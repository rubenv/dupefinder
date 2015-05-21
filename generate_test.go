package dupefinder

import "testing"

func TestGenerateFolders(t *testing.T) {
	err := Generate("catalog.txt", "invalid")
	if err == nil {
		t.Errorf("Expected an error: invalid is not an existing folder")
	}
}

func TestGenerateFinds(t *testing.T) {
	err := Generate("catalog.txt", "fixtures/a")
	if err != nil {
		t.Error(err)
	}

	// TODO: Read catalog, see if files are there and check checksums
}
