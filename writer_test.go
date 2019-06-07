package iomunger

import (
	"bytes"
	"reflect"
	"testing"
)

func TestWriterSuccess(t *testing.T) {
	var b bytes.Buffer
	w := NewWriter(&b, byte('a'), []byte("Z"))
	testStr := []byte("abcabc")
	resultStr := []byte("ZbcZbcZbcZbc")
	expectedWrite := len(resultStr) / 2

	// Write twice
	for _ = range []int{0, 0} {
		l, err := w.Write(testStr)
		if err != nil {
			t.Errorf("Could not write %q: %v", testStr, err)
			return
		} else if l != expectedWrite {
			t.Errorf("Tried to write %d bytes, but only wrote %d", expectedWrite, l)
			return
		}
	}

	if !reflect.DeepEqual(b.Bytes(), resultStr) {
		t.Errorf("Expected %q, got %q", resultStr, b.Bytes())
	}
}

func TestWriterMultibyte(t *testing.T) {
	var b bytes.Buffer
	w := NewWriter(&b, byte('a'), []byte("ZZ"))
	testStr := []byte("abcabc")
	resultStr := []byte("ZZbcZZbcZZbcZZbc")
	expectedWrite := len(resultStr) / 2

	// Write twice
	for _ = range []int{0, 0} {
		l, err := w.Write(testStr)
		if err != nil {
			t.Errorf("Could not write %q: %v", testStr, err)
			return
		} else if l != expectedWrite {
			t.Errorf("Tried to write %d bytes, but only wrote %d", expectedWrite, l)
			return
		}
	}

	if !reflect.DeepEqual(b.Bytes(), resultStr) {
		t.Errorf("Expected %q, got %q", resultStr, b.Bytes())
	}
}

/// FIXME TestWriterFailure
