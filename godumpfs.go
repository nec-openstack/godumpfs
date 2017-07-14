// @matcher  = (config[:matcher]  or NullMatcher.new)
// @reporter = (config[:reporter] or lambda {|x| puts x })
// @log_file = (config[:log_file] or nil)
// @dry_run  = (config[:dry_run] or false)
// @interval_proc = (config[:interval_proc] or lambda {})
// @written_bytes = 0

package godumpfs

import (
  "fmt"
  "github.com/nec-openstack/godumpfs/pkg/file"
  "path"
  "path/filepath"
  "strings"
  "time"
)

type GOdumpfs struct{ 

}

func (g *GOdumpfs) ValidateDirs(src string, dest string) error {
  src, err := filepath.Abs(src)
  if err != nil {
    // error message
    return err
  }

  dest, err = filepath.Abs(dest)
  if err != nil {
    // error message
    return err
  }

  // is src directory?
  isDir, err := file.IsDir(src)
  if err != nil {
    // error message
    return err
  }
  if isDir == false {
    // error message
    return fmt.Errorf("no such directory %v", src)
  }

  // is dest directory?
  isDir, err = file.IsDir(dest)
  if err != nil {
    // error message
    return err
  }
  if isDir == false {
    // error message
    return fmt.Errorf("no such directory %v", dest)
  }

  // are 
  if src == dest {
    return fmt.Errorf("src and dest are same :%v", src)
  }

  index := strings.Index(dest, src)
  if index == 0 {
    return fmt.Errorf("src(%v) contains dest(%v)", src, dest)
  }

  return nil
}

func (g *GOdumpfs) latestSnapshot(startTime time.Time, src string, dest string, base string) (string, error) {

  return "/home/inou/go/src/github.com/nec-openstack/godumpfs/fixtures/2017/06/21/base", nil
}

func (g *GOdumpfs) Start(src string, dest string, base string) error {
  if len(base) == 0 {
    base = path.Base(src)
  }

  return nil
}




