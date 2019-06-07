package iomunger

import (
	"testing"
)

func TestReaderSingleByte(t *testing.T) {
	r := NewReader(byteReader("abcabc"), byte('a'), []byte("Z"))
	testReader(t, r, "ZbcZbc", 30)
}

func TestReaderPartial(t *testing.T) {
	r := NewReader(byteReader("abcabcabcabc"), byte('a'), []byte("Z"))
	testReader(t, r, "ZbcZbc", 6)
}

func TestReaderMultibyte(t *testing.T) {
	r := NewReader(byteReader("abcabc"), byte('a'), []byte("ZZ"))
	testReader(t, r, "ZZbcZZbc", 30)
}

/// FIXME TestReaderFailure
