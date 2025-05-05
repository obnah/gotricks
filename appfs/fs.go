package appfs

import (
	"path/filepath"
	"runtime"
	"strings"

	"github.com/spf13/afero"
)

type FS struct {
	afero.Fs
	Root string
}

var (
	AppFs = newFS("/")
)

func newFS(path string) FS {
	var fs afero.Fs
	if path == "/" {
		fs = afero.NewOsFs()
	} else {
		fs = afero.NewBasePathFs(afero.NewOsFs(), path)
	}
	return FS{
		Fs:   fs,
		Root: path,
	}
}

func Fake(path string) {
	if !filepath.IsAbs(path) {
		_, file, _, _ := runtime.Caller(1)
		path = filepath.Join(filepath.Dir(file), path)
	}
	AppFs = newFS(path)
}

func Reset() {
	AppFs = newFS("/")
}

func OsPath(path string) string {
	if filepath.IsAbs(path) {
		path = filepath.Join(AppFs.Root, path)
	}
	return path
}

func appPath(path string) string {
	if AppFs.Root == "/" {
		return path
	}

	if strings.HasPrefix(path, AppFs.Root) {
		return strings.TrimPrefix(path, AppFs.Root)
	}

	root, _ := filepath.EvalSymlinks(AppFs.Root)

	return strings.TrimPrefix(path, root)
}
