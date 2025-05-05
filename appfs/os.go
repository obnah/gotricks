package appfs

import "os"

//Functions from os package

func Chdir(dir string) error {
	return os.Chdir(OsPath(dir))
}

func Chmod(name string, mode os.FileMode) error {
	return os.Chmod(OsPath(name), mode)
}

func Chown(name string, uid, gid int) error {
	return os.Chown(OsPath(name), uid, gid)
}

func Mkdir(name string, perm os.FileMode) error {
	return os.Mkdir(OsPath(name), perm)
}

func MkdirAll(path string, perm os.FileMode) error {
	return os.MkdirAll(OsPath(path), perm)
}

func ReadFile(name string) ([]byte, error) {
	return os.ReadFile(OsPath(name))
}

func Readlink(name string) (string, error) {
	return os.Readlink(OsPath(name))
}

func Remove(name string) error {
	return os.Remove(OsPath(name))
}

func RemoveAll(path string) error {
	return os.RemoveAll(OsPath(path))
}

func Rename(oldpath, newpath string) error {
	return os.Rename(OsPath(oldpath), OsPath(newpath))
}

func Symlink(oldname, newname string) error {
	return os.Symlink(OsPath(oldname), OsPath(newname))
}

func WriteFile(name string, data []byte, perm os.FileMode) error {
	return os.WriteFile(OsPath(name), data, perm)
}

// Type DirEntry

func ReadDir(name string) ([]os.DirEntry, error) {
	return os.ReadDir(OsPath(name))
}

// Type File

func Create(name string) (*os.File, error) {
	return os.Create(OsPath(name))
}

func Open(name string) (*os.File, error) {
	return os.Open(OsPath(name))
}

func OpenFile(name string, flag int, perm os.FileMode) (*os.File, error) {
	return os.OpenFile(OsPath(name), flag, perm)
}

// Type FileInfo

func Lstat(name string) (os.FileInfo, error) {
	return os.Lstat(OsPath(name))
}

func Stat(name string) (os.FileInfo, error) {
	return os.Stat(OsPath(name))
}
