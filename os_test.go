package attest

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestFileState(t *testing.T) {
	f, err := ioutil.TempFile(os.TempDir(), "")
	if err != nil {
		t.Fatal(err)
	}
	if err := FileState(f.Name(), true); err != nil {
		t.Fatal("unexpected error")
	}
	if err := FileState(f.Name(), false); err == nil {
		t.Fatal("error expected")
	}
	if err := os.Remove(f.Name()); err != nil {
		t.Fatal(err)
	}
	if err := FileState(f.Name(), false); err != nil {
		t.Fatal("unexpected error")
	}
	if err := FileState(f.Name(), true); err == nil {
		t.Fatal("error expected")
	}
}
