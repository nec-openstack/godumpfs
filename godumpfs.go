package godumpfs

import (
	"fmt"
	"github.com/nec-openstack/godumpfs/pkg/file"
	"path"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"time"
)

type GOdumpfs struct {
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

type SnapshotNotFound struct {
}

func (e *SnapshotNotFound) Error() string {
	return ""
}

type Dirs []string

func (d Dirs) Len() int { return len(d) }

func (d Dirs) Swap(i, j int) {
	d[i], d[j] = d[j], d[i]
}

func (d Dirs) Less(i, j int) bool { return strings.Compare(d[i], d[j]) < 0 }

func (g *GOdumpfs) latestSnapshot(startTime time.Time, src string, dest string, base string) (string, error) {
	dd := "[0-9][0-9]"
	dddd := dd + dd
	globPath := path.Join(dest, dddd, dd, dd)
	dirs, err := filepath.Glob(globPath)
	if err != nil {
		return "", err
	}
	if dirs == nil {
		return "", &SnapshotNotFound{}
	}
	sort.Sort(sort.Reverse(Dirs(dirs)))

	for _, dir := range dirs {
		p := path.Join(dir, base)
		ps := strings.Split(dir, string(filepath.Separator))
		l := len(ps)
		y, _ := strconv.Atoi(ps[l-3])
		m, _ := strconv.Atoi(ps[l-2])
		d, _ := strconv.Atoi(ps[l-1])
		t := time.Date(y, time.Month(m), d, 0, 0, 0, 0, time.Local)
		if y != t.Year() || time.Month(m) != t.Month() || d != t.Day() {
			// 日付じゃない
			continue
		}
		if isDir, err := file.IsDir(p); err != nil || !isDir {
			// ディレクトリじゃない
			continue
		}
		// TODO 未来か?
		if t.After(startTime) {
			continue
		}
		return p, nil
	}

	return "", &SnapshotNotFound{}
}

func (g *GOdumpfs) Start(src string, dest string, base string) error {
	if len(base) == 0 {
		base = path.Base(src)
	}

	return nil
}
