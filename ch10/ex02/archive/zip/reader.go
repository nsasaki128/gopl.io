package zip

import (
	"os"

	"gopl.io/ch10/ex02/archive"
)

//reference
//http://text.baldanders.info/golang/compress-data/
//https://blog.freedom-man.com/zip-structure-golang/
//https://imgur.com/BXuOFqT

//TODO: FIX this function
func read(f *os.File) (archive.Reader, error) {
	return nil, nil
}

func init() {
	//PK\003\004 = \x50\x4b
	archive.RegisterFormat("zip", "\x50\x4b", 0, read)
}
