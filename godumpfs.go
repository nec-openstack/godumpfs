package godumpfs

import (
	"fmt"
	"github.com/nec-openstack/godumpfs/pkg/file"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"syscall"
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

func (g *GOdumpfs) sameDirectory(src, dest string) (bool, error) {
	src, err := filepath.Abs(src)
	if err != nil {
		// error message
		return false, err
	}
	dest, err = filepath.Abs(dest)
	if err != nil {
		// error message
		return false, err
	}
	return src == dest, nil
}

func (g *GOdumpfs) subDirectory(src, dest string) (bool, error) {
	src, err := filepath.Abs(src)
	if err != nil {
		// error message
		return false, err
	}
	dest, err = filepath.Abs(dest)
	if err != nil {
		// error message
		return false, err
	}

	match, err := regexp.MatchString(string(filepath.Separator)+"$", src)
	if err != nil {
		// error message
		return false, err
	}
	if !match {
		src += string(filepath.Separator)
	}

	return regexp.MatchString("^"+regexp.QuoteMeta(src), dest)
}

func (g *GOdumpfs) datedir(date *time.Time) string {
	s := string(filepath.Separator)
	return fmt.Sprintf("%d%s%02d%s%02d", date.Year, s, date.Month, s, date.Day)
}

func (g *GOdumpfs) Start(src string, dest string, base string) error {
	startTime := time.Now()
	if same, err := g.sameDirectory(src, dest); same || err != nil {
		return fmt.Errorf("cannot copy a directory, `%s', into itself, `%s'", src, dest)
	}
	if sub, err := g.subDirectory(src, dest); sub || err != nil {
		return fmt.Errorf("cannot copy a directory, `%s', into itself, `%s'", src, dest)
	}

	if src != "/" {
		re := regexp.MustCompile("/+$")
		src = re.ReplaceAllString(src, "")
	}
	if len(base) == 0 {
		base = path.Base(src)
	}

	latest, err := g.latestSnapshot(startTime, src, dest, base)
	if err != nil {
		if _, ok := err.(*SnapshotNotFound); !ok { // SnapShot not foundじゃないエラーの場合の意
			return err
		}
	}
	today := path.Join(dest, g.datedir(&startTime), base)

	syscall.Umask(0077)
	// TODO dry run
	os.MkdirAll(today, 0700)
	if err == nil {
		// SnapShot 居る
		// TODO g.UpdateSnapshot(src, latest, today)
	} else {
		// SnapShot 居いない
		// TODO g.RecursiveCopy(src, today)
	}
	fmt.Println(latest) // TODO Delete

	return nil
}
