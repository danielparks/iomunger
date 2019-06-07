package iomunger

// Utility functions for testing

import (
	"bytes"
	"reflect"
	"testing"
)

func byteReader(src string) *bytes.Reader {
	return bytes.NewReader([]byte(src))
}

func testReader(t *testing.T, r Reader, expected string, bufSize int) {
	b := make([]byte, bufSize)
	read, err := r.Read(b)
	if err != nil {
		t.Errorf("Could not read: %v", err)
		return
	} else if read != len(expected) {
		t.Errorf("Expected to read %d bytes, but only got %d", len(expected), read)
		return
	}

	if !reflect.DeepEqual(b[0:read], []byte(expected)) {
		t.Errorf("Expected %q, got %q", expected, b[0:read])
	}
}

func testWriterTwice(t *testing.T, b *bytes.Buffer, w Writer, src, expected string) {
	fullExpected := expected + expected

	// Write twice
	for _ = range []int{0, 0} {
		l, err := w.Write([]byte(src))
		if err != nil {
			t.Errorf("Could not write %q: %v", src, err)
			return
		} else if l != len(expected) {
			t.Errorf("Expected to write %d bytes, but only wrote %d", len(expected), l)
			return
		}
	}

	if !reflect.DeepEqual(b.Bytes(), []byte(fullExpected)) {
		t.Errorf("Expected %q, got %q", fullExpected, b.Bytes())
	}
}
