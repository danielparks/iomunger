package iomunger

import (
	"bytes"
	"io"
)

type Writer struct {
	raw io.Writer
	old byte
	new byte
}

func NewWriter(w io.Writer, old, new byte) Writer {
	return Writer{w, old, new}
}

func (w Writer) Write(s []byte) (int, error) {
	return w.raw.Write(bytes.ReplaceAll(s, []byte{w.old}, []byte{w.new}))
}
