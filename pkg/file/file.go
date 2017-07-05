package file

import (
  "os"
)

func IsRealFile(path string) (bool, error) {
  fInfo, err := os.Lstat(path)
  if err != nil {
    isExist := os.IsExist(err)
    if isExist != true {
      return false, nil
    }
    return false, err
  }
  fMode := fInfo.Mode()
  return fMode.IsRegular(), nil
}
