package archive

import (
	"bufio"
	"errors"
	"io"
	"os"
)

// This construction is originally based on image/format.go

// ErrFormat indicates that decoding encountered an unknown format.
var ErrFormat = errors.New("archive: unknown format")

// Just copied from image/format.go
// A reader is an io.Reader that can also peek ahead.
type Reader interface {
	io.Reader
	io.ByteReader
}

// A format holds an archive format's name, magic header and how to read it.
type format struct {
	name, magic string
	offset      int
	read        func(file *os.File) (Reader, error)
}

var formats []format

func RegisterFormat(name, magic string, offset int, read func(*os.File) (Reader, error)) {
	formats = append(formats, format{name, magic, offset, read})
}

// Just copied from image/format.go
// Match reports whether magic matches b. Magic may contain "?" wildcards.
func match(magic string, b []byte) bool {
	if len(magic) != len(b) {
		return false
	}
	for i, c := range b {
		if magic[i] != c && magic[i] != '?' {
			return false
		}
	}
	return true
}

// Sniff determines the format of file data.
func sniff(file *os.File) format {
	r := bufio.NewReader(file)
	//Below is same as original image/format.go
	for _, f := range formats {
		b, err := r.Peek(len(f.magic))
		if err == nil && match(f.magic, b) {
			return f
		}
	}
	return format{}
}
