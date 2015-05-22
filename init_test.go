package dupefinder

import (
	"io/ioutil"
	"testing"
)

func tempFilename(t *testing.T) string {
	f, err := ioutil.TempFile("", "dupefinder")
	if err != nil {
		t.Fatal(err)
	}
	f.Close()

	return f.Name()
}
