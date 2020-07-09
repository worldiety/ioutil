package ioutil

import "testing"

func TestTypedLittleEndianBuffer_assertType(t *testing.T) {
	le := &TypedLittleEndianBuffer{
		Bytes: nil,
		Pos:   0,
	}
	//le.assertType(0)
}