package file

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestIsRealFileTrue(t *testing.T) {
	f, _ := ioutil.TempFile("/tmp", "godumpfs")
	defer os.Remove(f.Name())
	r, _ := IsRealFile(f.Name())
	if r != true {
		t.Errorf("%v is not True", r)
	}
}

func TestIsRealFileFalse(t *testing.T) {
	r, _ := IsRealFile("/tmp/pokemon/pikachu")
	if r != false {
		t.Errorf("%v is not False", r)
	}
}

func TestIsRealFileCheckDirectory(t *testing.T) {
	d, _ := ioutil.TempDir("/tmp", "godumpfs")
	defer os.Remove(d)
	r, _ := IsRealFile(d)
	if r != false {
		t.Errorf("%v is not False", r)
	}
}

func TestIsRealFileCheckSymLink(t *testing.T) {
	f, _ := ioutil.TempFile("/tmp", "godumpfs")
	defer os.Remove(f.Name())
	fSymLink := f.Name() + "-symlink"
	os.Symlink(f.Name(), fSymLink)
	defer os.Remove(fSymLink)
	r, _ := IsRealFile(fSymLink)
	if r != false {
		t.Errorf("%v is not False", r)
	}
}
