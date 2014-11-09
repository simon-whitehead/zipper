package zipper

import (
	"archive/zip"
	"io"
	"os"
	"path/filepath"
)

type zipper struct {
	p  string
	zw *zip.Writer
}

func NewZipper(p string) *zipper {
	return &zipper{p: p}
}

func (z *zipper) Zip(w io.Writer) error {
	z.zw = zip.NewWriter(w)
	filepath.Walk(z.p, z.walk)

	return z.zw.Close()
}

func (z *zipper) walk(p string, fi os.FileInfo, e error) error {
	fn := p
	if fi.IsDir() {
		fn += string(os.PathSeparator)
	}

	f, e := z.zw.Create(fn)
	if e != nil {
		return e
	}

	o, e := os.Open(fn)
	if e != nil {
		return e
	}

	defer o.Close()

	io.Copy(f, o)

	return nil
}
