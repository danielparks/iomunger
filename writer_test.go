package iomunger

import (
	"bytes"
	"testing"
)

func TestWriterSuccess(t *testing.T) {
	var b bytes.Buffer
	w := NewWriter(&b, byte('a'), []byte("Z"))
	testWriterTwice(t, &b, w, "abcabc", "ZbcZbc")
}

func TestWriterMultibyte(t *testing.T) {
	var b bytes.Buffer
	w := NewWriter(&b, byte('a'), []byte("ZZ"))
	testWriterTwice(t, &b, w, "abcabc", "ZZbcZZbc")
}

/// FIXME TestWriterFailure
