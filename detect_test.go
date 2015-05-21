package dupefinder

import "testing"

func TestDetectFolders(t *testing.T) {
	err := Detect("catalog.txt", true, false, "invalid")
	if err == nil {
		t.Errorf("Expected an error: invalid is not an existing folder")
	}
}
