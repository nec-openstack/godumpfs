package godumpfs

import (
	"testing"
	"time"
  "os"
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
  dest := wd+"/fixtures/"

  actual, _ := g.latestSnapshot(time.Date(2017, 7, 20, 23, 59, 59, 0, time.UTC), "", dest, "base")
  expected := wd+"/fixtures/2017/06/21/base"

  if actual != expected {
		t.Errorf("latestSnapshot doesn't detects invalid state.")
  }

}

func TestStart(t *testing.T) {
  g := &GOdumpfs{}
	g.Start("", "", "")
}
