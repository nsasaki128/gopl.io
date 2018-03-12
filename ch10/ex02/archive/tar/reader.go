package tar

import (
	"os"

	"gopl.io/ch10/ex02/archive"
)

//reference
//https://github.com/BcRikko/learning-go/tree/master/tar-gzip
//http://www.redout.net/data/tar.html

//TODO: FIX this function
func read(f *os.File) (archive.Reader, error) {
	return nil, nil
}
func init() {
	archive.RegisterFormat("tar", "\x75\x73\x74\x61\x72", 0x101, read)
}
