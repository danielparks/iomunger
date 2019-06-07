package iomunger

import (
	"io"
)

func Unix2DosReader(r io.Reader) Reader {
	return NewReader(r, byte('\n'), []byte("\r\n"))
}

func Unix2DosWriter(r io.Writer) Writer {
	return NewWriter(r, byte('\n'), []byte("\r\n"))
}
