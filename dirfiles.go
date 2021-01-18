package shutil

import "os"

// CopyFileObj Copy the contents of the file-like object fsrc to the file-like object fdst.
func CopyFileObj(fsrc, fdst os.File, length ...int64) (err error) {
	return
}

// CopyFile Copy the contents (no metadata) of the file named src to a file named dst
// and return dst in the most efficient way possible.
// src and dst are path-like objects or path names given as strings.
func CopyFile(src, dst string, followSymlinks bool) (err error) {
	return
}

// CopyMode Copy the permission bits from src to dst.
func CopyMode(src, dst string, followSymlinks bool) (err error) {
	return

}

// CopyStat Copy the permission bits, last access time, last modification time,
// and flags from src to dst.
func CopyStat(src, dst string, followSymlinks bool) (err error) {
	return
}

// Copy Copies the file src to the file or directory dst.
func Copy(src, dst string, followSymlinks bool) (err error) {
	return
}

// Copy2 Copies the file src to the file or directory dst.
// also attempts to preserve file metadata
func Copy2(src, dst string, followSymlinks bool) (err error) {
	return
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
func RmTree() (err error) {
	return
}

// Move Recursively move a file or directory (src) to another location (dst) and return the destination.
func Move() (err error) {
	return
}

// DiskUsage Return disk usage statistics about the given path
func DiskUsage(path string) (err error) {
	return
}

// Chown Change owner user and/or group of the given path.
func Chown(path string, user, group uint) (err error) {
	return
}

// Which Return the path to an executable which would be run if the given cmd was called.
func Which(cmd string, mode uint, path ...string) (p string, err error) {
	return
}
