package iomunger

import (
	"bytes"
	"io"
)

type Reader struct {
	raw io.Reader
	old byte
	new []byte
}

func NewReader(r io.Reader, old byte, new []byte) Reader {
	return Reader{r, old, new}
}

func (r Reader) Read(s []byte) (int, error) {
	read, err := r.raw.Read(s)
	copied := copy(s, bytes.ReplaceAll(s[0:read], []byte{r.old}, r.new))
	return copied, err
}
