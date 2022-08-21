package shutil

import (
	"fmt"
	"io"
	"os"
	"path"
	"path/filepath"
	"strings"
	"syscall"
	"time"
)

// CopyFileObj Copy the contents of the file-like object fsrc to the file-like object fdst.
func CopyFileObj(fsrc, fdst *os.File, length ...int64) (err error) {
	defer fsrc.Close()
	defer fdst.Close()
	_, err = io.Copy(fdst, fsrc)
	return err
}

// CopyFile Copy the contents (no metadata) of the file named src to a file named dst
// and return dst in the most efficient way possible.
func CopyFile(src, dst string, followSymlinks ...bool) (err error) {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, in)
	if err != nil {
		return err
	}
	return out.Close()
}

// CopyMode Copy the permission bits from src to dst.
func CopyMode(src, dst string, followSymlinks ...bool) (err error) {
	fileInfo, err := os.Stat(src)
	if err != nil {
		return err
	}
	return os.Chmod(dst, fileInfo.Mode())

}

// CopyStat Copy the permission bits, last access time, last modification time,
// and flags from src to dst.
func CopyStat(src, dst string, followSymlinks ...bool) (err error) {
	fileInfo, err := os.Stat(src)
	if err != nil {
		return err
	}
	err = os.Chmod(dst, fileInfo.Mode())
	if err != nil {
		return err
	}
	atime := fileInfo.Sys().(*syscall.Stat_t).Atim
	return os.Chtimes(dst, time.Unix(atime.Sec, atime.Nsec), fileInfo.ModTime())

}

// Copy Copies the file src to the file or directory dst.
func Copy(src, dst string, followSymlinks ...bool) (err error) {
	s, err := os.Stat(dst)
	if err != nil {
		return err
	}
	if !s.IsDir() {
		return CopyFile(src, dst)
	}
	return CopyFile(src, path.Join(dst, path.Base(src)))
}

// Copy2 Copies the file src to the file or directory dst.
// also attempts to preserve file metadata
func Copy2(src, dst string, followSymlinks ...bool) (err error) {
	if err := Copy(src, dst, followSymlinks...); err != nil {
		return err
	}
	if i, _ := os.Stat(dst); i.IsDir() {
		dst = path.Join(dst, path.Base(src))
	}
	return CopyStat(src, dst, followSymlinks...)
}

// IgnorePatterns A factory function creates a function that
// can be used as a callable for copytree()’s ignore argument
func IgnorePatterns(patterns ...string) {

}

// CopyTree Recursively copy an entire directory tree rooted at src to
// a directory named dst and return the destination directory.
func CopyTree() (err error) {
	return
}

// CopyTreeWithOpt same to copytree,but can add extra opt-config
func CopyTreeWithOpt() (err error) {
	return
}

// RmTree Delete an entire directory tree;
// path must point to a directory (but not a symbolic link to a directory).
func RmTree(dir string) (err error) {
	return os.RemoveAll(dir)
}

func RemoveDirContents(dir string, p ...string) (err error) {
	fp := "*"
	if len(p) == 0 {
		fp = p[0]
	}
	files, err := filepath.Glob(filepath.Join(dir, fp))
	if err != nil {
		return err
	}
	for _, file := range files {
		err = os.RemoveAll(file)
		if err != nil {
			return err
		}
	}
	return nil

}

// Move Recursively move a file or directory (src) to another location (dst) and return the destination.
func Move(srt, dst string) (err error) {
	return os.Rename(srt, dst)
}

type Diskusage struct {
	Total     uint64  //  total size of the file system
	Used      uint64  //  total bytes used in file system
	Free      uint64  // total free bytes on file system
	Available uint64  // total available bytes on file system to an unprivileged user
	Usage     float32 //percentage of use on the file system
}

type DiskusageHuman struct {
	Total     string //  total size of the file system
	Used      string //  total bytes used in file system
	Free      string // total free bytes on file system
	Available string // total available bytes on file system to an unprivileged user
	Usage     string //percentage of use on the file system
}

func (du Diskusage) ToHuman(unit ...string) DiskusageHuman {
	m := map[string]string{
		"b": "Bytes", "kb": "Kilobytes",
		"mb": "MMegabytes",
		"gb": "Gigabytes",
		"tb": "Terrabytes",
		"pb": "PetaBytes"}

	u := "mb"
	if len(unit) > 0 {
		u = strings.ToLower(unit[0])
	}
	var d uint64 = 1
	switch u {
	case "kb":
		d = 1024
	case "mb":
		d = 1024 * 1024
	case "gb":
		d = 1024 * 1024 * 1024
	case "tb":
		d = 1024 * 1024 * 1024 * 1024 * 1024
	case "pb":
		d = 1024 * 1024 * 1024 * 1024 * 1024 * 1024

	default:
		d = 1024 * 1024
	}

	duh := DiskusageHuman{}
	duh.Available = fmt.Sprintf("%v %s", du.Available/d, m[u])
	duh.Total = fmt.Sprintf("%v %s", du.Total/d, m[u])
	duh.Free = fmt.Sprintf("%v %s", du.Free/d, m[u])
	duh.Used = fmt.Sprintf("%v %s", du.Used/d, m[u])
	duh.Usage = fmt.Sprintf("%2.f %%", du.Usage*100)

	return duh
}

// DiskUsage Return disk usage statistics about the given path
func DiskUsage(path string) (usage Diskusage, err error) {
	diskUsage := Diskusage{}
	var stat syscall.Statfs_t
	if err := syscall.Statfs(path, &stat); err != nil {
		return diskUsage, err
	}
	diskUsage.Free = stat.Bfree * uint64(stat.Bsize)
	diskUsage.Total = uint64(stat.Blocks) * uint64(stat.Bsize)
	diskUsage.Available = stat.Bavail * uint64(stat.Bsize)
	diskUsage.Used = diskUsage.Total - diskUsage.Free
	diskUsage.Usage = float32(diskUsage.Used) / float32(diskUsage.Total)
	return diskUsage, nil
}

// Chown Change owner user and/or group of the given path.
func Chown(path string, user, group uint) (err error) {
	return
}

// Which Return the path to an executable which would be run if the given cmd was called.
func Which(cmd string, mode uint, path ...string) (p string, err error) {
	if access_check(cmd) {
		return cmd, nil
	}
	paths := ""
	if len(path) == 0 {
		paths = os.Getenv("PATH")
	}
	if paths == "" {
		paths = "/bin:/usr/bin"
	}
	pathS := strings.Split(paths, ":")
	for _, p := range pathS {
		name := filepath.Join(p, cmd)
		if access_check(name) {
			return name, nil
		}

	}
	return
}

func access_check(fn string, mode ...os.FileMode) bool {
	f, err := os.Stat(fn)
	if err != nil && !os.IsExist(err) {
		return false
	}
	if f.IsDir() {
		return false
	}
	// TODO:check access

	return true
}
