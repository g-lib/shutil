package shutil

import "fmt"

// ArchiveFormat obj
type ArchiveFormat struct {
	Name  string
	Sufix []string
	Desc  string
}

func (af ArchiveFormat) String() string {
	return fmt.Sprintf("\nName:%s\nSufix:%v\nDesc:%v\n",
		af.Name, af.Sufix, af.Desc)
}

// RegisterArchiveFormat Register an archiver for the format name.
func RegisterArchiveFormat(name string, f func()) (err error) {
	return
}

// UnRegisterArchiveFormat Remove the archive format name from the list of supported formats.
func UnRegisterArchiveFormat(name string) (err error) {
	return
}

// GetArchiveFormats Return a list of supported formats for archiving
func GetArchiveFormats() (afs []*ArchiveFormat) {
	afs = []*ArchiveFormat{
		{"bztar", []string{".tar.bz2", ".tbz2"}, "bzip2'ed tar-file"},
		{"gztar", []string{".tar.gz", ".tgz"}, "gzip'ed tar-file"},
		{"tar", []string{".tar"}, "uncompressed tar file"},
		{"xztar", []string{".tar.xz", ".txz"}, "xz'ed tar-file"},
		{"zip", []string{".zip"}, "ZIP file"},
	}
	return
}

// MakeOpt obj
type MakeOpt struct {
	Name    string
	Format  string
	RootDir string
	BaseDir string
	Verbose bool
	DryRun  bool
	Owner   int64
	Group   int64
	Logger  string
}

// MakeArchive Create an archive file (such as zip or tar) and return its name.
func MakeArchive(opt *MakeOpt) (err error) {
	return
}

// RegisterUnpackFormat Registers an unpack format.
func RegisterUnpackFormat(name string, f func()) (err error) {
	return
}

// UnRegisterUnpackFormat Unregister an unpack format. name is the name of the format.
func UnRegisterUnpackFormat(name string) (err error) {
	return
}

// GetUnpackFormats Return a list of all registered formats for unpacking.
func GetUnpackFormats() (afs []*ArchiveFormat) {
	afs = []*ArchiveFormat{
		{"bztar", []string{".tar.bz2", ".tbz2"}, "bzip2'ed tar-file"},
		{"gztar", []string{".tar.gz", ".tgz"}, "gzip'ed tar-file"},
		{"tar", []string{".tar"}, "uncompressed tar file"},
		{"xztar", []string{".tar.xz", ".txz"}, "xz'ed tar-file"},
		{"zip", []string{".zip"}, "ZIP file"},
	}
	return
}

// UnpackOpt obj
type UnpackOpt struct {
	Filename   string
	ExtractDir string
	Format     string
}

// UnpackArchive Unpack an archive. filename is the full path of the archive.
func UnpackArchive(opt *UnpackOpt) {

}

/*
make_archive
shutil.register_archive_format


*/
