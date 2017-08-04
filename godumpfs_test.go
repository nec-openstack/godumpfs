package godumpfs

import (
	"os"
	"testing"
	"time"
)

func TestValidateDirsSrcContainsDest(t *testing.T) {
	g := &GOdumpfs{}
	wd, _ := os.Getwd()
	err := g.ValidateDirs(wd+"/fixtures/src", wd+"/fixtures/src/dest")
	if err == nil {
		t.Errorf("ValidateDirs doesn't detects invalid state.")
	}
}

func TestLatestSnapshot(t *testing.T) {
	g := &GOdumpfs{}
	wd, _ := os.Getwd()
	dest := wd + "/fixtures/"

	actual, _ := g.latestSnapshot(time.Date(2017, 7, 20, 23, 59, 59, 0, time.UTC), "", dest, "base")
	expected := wd + "/fixtures/2017/06/21/base"

	if actual != expected {
		t.Errorf("latestSnapshot doesn't detects invalid state.")
	}

}

func TestStart(t *testing.T) {
	g := &GOdumpfs{}
	err := g.Start("/foo/bar", "/foo/bar/baz", "")
	if err == nil {
		t.Errorf("Start doesn't detects invalid state.")
	}
	err = g.Start("/foo/bar", "/foo/bar", "")
	if err == nil {
		t.Errorf("Start doesn't detects invalid state.")
	}
}

func TestSameDirectory(t *testing.T) {
	g := &GOdumpfs{}
	actual, err := g.sameDirectory("foo/bar/baz", "foo/../foo/bar/baz")
	if err != nil {
		t.Errorf("SameDirectory doesn't detects invalid state.")
	}
	expected := true
	if actual != expected {
		t.Errorf("SameDirectory doesn't detects invalid state.")
	}
	// TODO test motto tuika
}

func TestSubDirectory(t *testing.T) {
	g := &GOdumpfs{}
	actual, err := g.subDirectory("foo/bar/baz", "foo/bar/baz/asdf")
	if err != nil {
		t.Errorf("SubDirectory doesn't detects invalid state.")
	}
	expected := true
	if actual != expected {
		t.Errorf("SubDirectory doesn't detects invalid state.")
	}

	actual, err = g.subDirectory("foo/.bar/baz", "foo/abar/baz/asdf")
	if err != nil {
		t.Errorf("SubDirectory doesn't detects invalid state.")
	}
	expected = false
	if actual != expected {
		t.Errorf("SubDirectory doesn't detects invalid state.")
	}

	// TODO motto motto
}
