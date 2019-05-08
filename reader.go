package iomunger

import (
	"fmt"
	"io"
	"bytes"
)

type Reader struct {
	raw io.Reader
	old byte
	new byte
}

func NewReader(r io.Reader, old, new byte) Reader {
	return Reader{r, old, new}
}

func (r Reader) Read(s []byte) (int, error) {
	read, err := r.raw.Read(s)

	copied := copy(s, bytes.ReplaceAll(s[0:read], []byte{r.old}, []byte{r.new}))
	if err == nil && copied != read {
		// This should never happen. Only here out of paranoia.
		return copied, fmt.Errorf("Could only copy %d out of %d bytes into buffer", copied, read)
	}

	return read, err
}
