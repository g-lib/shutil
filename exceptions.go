package shutil

// Error shutil base error
type Error struct {
	Err error
}

// SameFileError Raised when source and destination are the same file.
type SameFileError struct {
	Error
}

// SpecialFileError Raised when trying to do a kind of operation
// (e.g. copying) which is not supported on a special file
// (e.g. a named pipe)
type SpecialFileError struct {
	Error
}

// ExecError Raised when a command could not be executed
type ExecError struct {
	Error
}

// ReadError Raised when an archive cannot be read
type ReadError struct {
	Error
}

// RegistryError Raised when a registry operation with
// the archiving and unpacking registries fails
type RegistryError struct {
	Error
}

// giveUpOnFastCopy Raised as a signal to fallback on using
// raw read()/write() file copy when fast-copy functions fail to do so.
type giveUpOnFastCopy struct {
	Error
}

/* os error
var (
	// ErrInvalid indicates an invalid argument.
	// Methods on File will return this error when the receiver is nil.
	ErrInvalid = errInvalid() // "invalid argument"

	ErrPermission       = errPermission()       // "permission denied"
	ErrExist            = errExist()            // "file already exists"
	ErrNotExist         = errNotExist()         // "file does not exist"
	ErrClosed           = errClosed()           // "file already closed"
	ErrNoDeadline       = errNoDeadline()       // "file type does not support deadline"
	ErrDeadlineExceeded = errDeadlineExceeded() // "i/o timeout"
)
*/
