package directory

import (
	"errors"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
)

var (
	ErrDirEmpty    = errors.New("directory is empty")
	ErrDirNotFound = errors.New("directory not found")
)

func DirContent(path string) ([]string, error) {
	if dirNotExists(path) {
		return nil, ErrDirNotFound
	}
	empty, err := dirEmpty(path)
	if err != nil {
		return nil, err
	}
	if empty {
		return nil, ErrDirEmpty
	}

	return walkDir(path)
}

func dirEmpty(path string) (bool, error) {
	f, err := os.Open(path)
	if err != nil {
		return false, err
	}
	_, err = f.Readdirnames(3)
	if err == io.EOF {
		return true, nil
	}
	return false, err
}

func dirNotExists(path string) bool {
	_, err := os.Stat(path)
	return os.IsNotExist(err)
}

func walkDir(path string) ([]string, error) {
	var content []string
	picRegex := regexp.MustCompile(`^.+\.(jpe?g|png)`)

	err := filepath.WalkDir(path, func(path string, info fs.DirEntry, err error) error {
		if !info.IsDir() && picRegex.MatchString(path) {
			content = append(content, path)
		}
		return nil
	})

	return content, err
}
