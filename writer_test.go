package iomunger

import (
	"bytes"
	"reflect"
	"testing"
)

func TestWriterSuccess(t *testing.T) {
	var b bytes.Buffer
	w := NewWriter(&b, byte('a'), byte('Z'))
	testStr := []byte("abcabc")
	resultStr := []byte("ZbcZbcZbcZbc")

	// Write twice
	for _ = range []int{0, 0} {
		l, err := w.Write(testStr)
		if err != nil {
			t.Errorf("Could not write %q: %v", testStr, err)
			return
		} else if l != len(testStr) {
			t.Errorf("Tried to write %d bytes, but only wrote %d", len(testStr), l)
			return
		}
	}

	if !reflect.DeepEqual(b.Bytes(), resultStr) {
		t.Errorf("Expected %q, got %q", resultStr, b.Bytes())
	}
}

/// FIXME TestWriterFailure
