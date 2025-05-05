package appfs

import (
	"io/fs"
	"path/filepath"
	"strings"

	"github.com/spf13/afero"
)

func Abs(path string) (string, error) {
	if p, err := filepath.Abs(OsPath(path)); err != nil {
		return p, err
	} else {
		return appPath(p), nil
	}
}

func EvalSymlinks(path string) (string, error) {
	targ, err := filepath.EvalSymlinks(OsPath(path))
	if err != nil {
		return targ, err
	}

	root, err := filepath.EvalSymlinks(AppFs.Root)
	if err != nil {
		return "", err
	}

	if len(root) > 0 && root != "/" {
		targ = strings.TrimPrefix(targ, root)
	}

	return targ, nil
}

func Glob(pattern string) (matches []string, err error) {
	return afero.Glob(AppFs, pattern)
}

func Walk(root string, fn filepath.WalkFunc) error {
	return filepath.Walk(OsPath(root),
		func(path string, info fs.FileInfo, err error) error {
			return fn(appPath(path), info, err)
		})
}

func WalkDir(root string, fn fs.WalkDirFunc) error {
	return filepath.WalkDir(OsPath(root),
		func(path string, d fs.DirEntry, err error) error {
			return fn(appPath(path), d, err)
		})
}
