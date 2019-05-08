package iomunger

import (
	"bytes"
	"reflect"
	"testing"
)

func TestReaderPartial(t *testing.T) {
	testStr := []byte("abcabcabcabc")
	resultStr := []byte("ZbcZbc")
	src := bytes.NewReader(testStr)

	r := NewReader(src, byte('a'), byte('Z'))

	b := make([]byte, 6)
	read, err := r.Read(b)
	if err != nil {
		t.Errorf("Could not read: %v", err)
		return
	} else if read != 6 {
		t.Errorf("Tried to read 6 bytes, but only got %d", read)
		return
	}

	if !reflect.DeepEqual(b, resultStr) {
		t.Errorf("Expected %q, got %q", resultStr, b)
	}
}

func TestReaderShort(t *testing.T) {
	testStr := []byte("abcabc")
	resultStr := []byte("ZbcZbc")
	src := bytes.NewReader(testStr)

	r := NewReader(src, byte('a'), byte('Z'))

	b := make([]byte, 30)
	read, err := r.Read(b)
	if err != nil {
		t.Errorf("Could not read: %v", err)
		return
	} else if read != 6 {
		t.Errorf("Tried to read 6 bytes, but got %d", read)
		return
	}

	if !reflect.DeepEqual(b[0:read], resultStr) {
		t.Errorf("Expected %q, got %q", resultStr, b[0:read])
	}
}


/// FIXME TestReaderFailure
