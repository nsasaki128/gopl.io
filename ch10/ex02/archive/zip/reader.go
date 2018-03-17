package zip

import (
	"os"

	"archive/zip"

	"errors"
	"io"

	"gopl.io/ch10/ex02/archive"
)

//reference
//http://text.baldanders.info/golang/compress-data/
//https://blog.freedom-man.com/zip-structure-golang/
//https://imgur.com/BXuOFqT

type reader struct {
	r *zip.ReadCloser
	i int
}

func (r *reader) Next() (*archive.Header, error) {
	r.i++
	if len(r.r.File) <= r.i {
		r.r.Close()
		return nil, io.EOF
	}
	return &archive.Header{Name: r.r.File[r.i].Name, FileInfo: r.r.File[r.i].FileInfo()}, nil
}

func (r *reader) Read(b []byte) (n int, err error) {
	if r.r == nil {
		return 0, errors.New("invalid use of zip reader")
	}
	f := r.r.File[r.i]
	reader, err := f.Open()
	if err != nil {
		return 0, err
	}
	return reader.Read(b)
}

func read(f *os.File) (archive.Reader, error) {
	_, err := f.Stat()
	if err != nil {
		return nil, err
	}
	r, err := zip.OpenReader(f.Name())
	if err != nil {
		return nil, err
	}
	return &reader{r: r, i: -1}, nil
}

func init() {
	//PK\003\004 = \x50\x4b
	archive.RegisterFormat("zip", "\x50\x4b", 0, read)
}
