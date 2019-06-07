package iomunger

import (
	"bytes"
	"testing"
)

func TestUnix2DosReader(t *testing.T) {
	r := Unix2DosReader(byteReader("test\nfoo\n"))
	testReader(t, r, "test\r\nfoo\r\n", 30)
}

func TestUnix2DosWriter(t *testing.T) {
	var b bytes.Buffer
	w := Unix2DosWriter(&b)
	testWriterTwice(t, &b, w, "test\nfoo\n", "test\r\nfoo\r\n")
}
