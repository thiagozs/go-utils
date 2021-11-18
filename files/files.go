package files

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func FileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func DeleteFile(str string) error {
	if err := os.Remove(str); err != nil {
		return err
	}
	return nil
}

func WriteFile(str string, bytes []byte) error {
	return ioutil.WriteFile(str, bytes, 0644)
}

func MkdirAll(path string) error {
	if err := os.MkdirAll(path, os.ModeSticky|os.ModePerm); err != nil {
		return err
	}
	return nil
}

func IsDir(pathFile string) bool {
	if pathAbs, err := filepath.Abs(pathFile); err != nil {
		return false
	} else if fileInfo, err := os.Stat(pathAbs); os.IsNotExist(err) || !fileInfo.IsDir() {
		return false
	}

	return true
}

func RmdirAll(path string) error {
	d, err := os.Open(path)
	if err != nil {
		return err
	}
	defer d.Close()
	names, err := d.Readdirnames(-1)
	if err != nil {
		return err
	}
	for _, name := range names {
		err = os.RemoveAll(filepath.Join(path, name))
		if err != nil {
			return err
		}
	}
	rootFolder := strings.Split(path, "/")
	if len(rootFolder) > 1 {
		err = os.RemoveAll(rootFolder[0])
		if err != nil {
			return err
		}
	}
	return nil
}
