/*
 * Copyright 2020 Torben Schinke
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package ioutil

import (
	"encoding/binary"
	"io"
	"math"
	"reflect"
	"strconv"
	"unsafe"
)

// An Encoder implements various encoding helpers for Little Endian and Big Endian. It may optimize some
// paths in the future, so that a generic call with byte order may be slower than the direct invocation.
// The implementation reuses an internal buffer to avoid heap allocations and is therefore not thread safe.
type Encoder struct {
	buf8        []byte
	out         io.Writer
	firstErr    error
	failOnError bool
}

// Resets removes any error state.
func (e *Encoder) Reset() {
	e.firstErr = nil
}

// quickFail returns true, if an error is already pending and we should not bother the writer again.
func (e *Encoder) quickFail() bool {
	return e.failOnError && e.firstErr != nil
}

// WriteArray just writes the slice out, without any prefix for the length.
// If an error occurs returns the number of written bytes.
func (e *Encoder) WriteArray(v []byte) int {
	if e.quickFail() {
		return 0
	}
	n, err := e.out.Write(v)
	e.noteErr(err)
	return n
}

// WriteBlob writes a prefixed byte slice of variable length.
func (e *Encoder) WriteBlob(o binary.ByteOrder, p IntSize, v []byte) {
	if e.quickFail() {
		return
	}

	switch p {
	case I8:
		if len(v) > math.MaxUint8 {
			e.noteErr(IntegerOverflow{Val: len(v), Max: math.MaxUint8})
			return
		}
		e.WriteUint8(uint8(len(v)))
	case I16:
		if len(v) > math.MaxUint16 {
			e.noteErr(IntegerOverflow{Val: len(v), Max: math.MaxUint16})
			return
		}
		e.WriteUint16(o, uint16(len(v)))
	case I24:
		if len(v) > MaxUint24 {
			e.noteErr(IntegerOverflow{Val: len(v), Max: MaxUint24})
			return
		}
		e.WriteUint24(o, uint32(len(v)))
	case I32:
		if len(v) > math.MaxUint32 {
			e.noteErr(IntegerOverflow{Val: len(v), Max: math.MaxUint32})
			return
		}
		e.WriteUint32(o, uint32(len(v)))
	case I40:
		if len(v) > MaxUint40 {
			e.noteErr(IntegerOverflow{Val: len(v), Max: MaxUint40})
			return
		}
		e.WriteUint40(o, uint64(len(v)))
	case I64:
		// overflow cannot happen, len is at most positive signed 64 bit value
		e.WriteUint64(o, uint64(len(v)))
	case IVar:
		// overflow cannot happen, len is at most positive signed 64 bit value
		e.WriteUvarint(uint64(len(v)))
	default:
		panic("unknown IntSize: " + strconv.Itoa(int(p)))
	}
	e.WriteArray(v)
}

// WriteUTF8 writes a prefixed unmodified utf8 string sequence of variable length.
func (e *Encoder) WriteUTF8(o binary.ByteOrder, p IntSize, v string) {
	if e.quickFail() {
		return
	}

	str := *(*reflect.StringHeader)(unsafe.Pointer(&v))
	// do not modify the slice, because this is a hack to avoid an unnecessary copy and heap allocation
	slice := *(*[]byte)(unsafe.Pointer(&reflect.SliceHeader{
		Data: str.Data,
		Len:  str.Len,
		Cap:  str.Len,
	}))

	e.WriteBlob(o, p, slice)
}

// WriteBool writes one byte.
func (e *Encoder) WriteBool(v bool) {

}

// WriteUint8 writes an unsigned byte
func (e *Encoder) WriteUint8(v uint8) {

}

// WriteInt8 writes a signed byte
func (e *Encoder) WriteInt8(v int8) {

}

// WriteUInt16 writes an unsigned 2 byte integer.
func (e *Encoder) WriteUint16(o binary.ByteOrder, v uint16) {

}

// WriteInt16 writes a signed 2 byte integer.
func (e *Encoder) WriteInt16(o binary.ByteOrder, v int16) {

}

// WriteUint24 writes an unsigned 3 byte integer.
func (e *Encoder) WriteUint24(o binary.ByteOrder, v uint32) {

}

// WriteInt24 writes a signed 3 byte integer.
func (e *Encoder) WriteInt24(o binary.ByteOrder, v int32) {

}

// WriteUint32 writes an unsigned 4 byte integer.
func (e *Encoder) WriteUint32(o binary.ByteOrder, v uint32) {

}

// WriteInt32 writes a signed 4 byte integer.
func (e *Encoder) WriteInt32(o binary.ByteOrder, v int32) {

}

// WriteInt40 writes a signed 5 byte integer.
func (e *Encoder) WriteInt40(o binary.ByteOrder, v int64) {

}

// WriteUint40 writes an unsigned 5 byte integer.
func (e *Encoder) WriteUint40(o binary.ByteOrder, v uint64) {

}

// WriteUint64 writes an unsigned 8 byte integer.
func (e *Encoder) WriteUint64(o binary.ByteOrder, v uint64) {

}

// WriteInt64 writes a signed 8 byte integer.
func (e *Encoder) WriteInt64(o binary.ByteOrder, v int64) {

}

// WriteUvarint writes a variable length integer, up to 9 bytes using zig-zag protobuf encoding.
func (e *Encoder) WriteUvarint(v uint64) {

}

// WriteVarint writes a variable length and signed integer, up to 9 bytes using zig-zag protobuf encoding.
func (e *Encoder) WriteVarint(v int64) {

}

// WriteFloat32 writes a float32 IEEE 754 4 byte bit sequence.
func (e *Encoder) WriteFloat32(o binary.ByteOrder, v float32) {

}

// WriteFloat64 writes a float64 IEEE 754 8 byte bit sequence.
func (e *Encoder) WriteFloat64(o binary.ByteOrder, v float64) {

}

// WriteComplex64 writes two float32 IEEE 754 4 byte bit sequences.
func (e *Encoder) WriteComplex64(o binary.ByteOrder, v complex64) {

}

// WriteComplex128 writes two float32 IEEE 754 4 byte bit sequences.
func (e *Encoder) WriteComplex128(o binary.ByteOrder, v complex128) {

}

// Write follows the io.Writer contract.
func (e *Encoder) Write(p []byte) (int, error) {
	if e.quickFail() {
		return 0, e.firstErr
	}
	n, err := e.out.Write(p)
	e.noteErr(err)
	return n, err
}

// WriteByte follows the io.ByteWriter contract.
func (e *Encoder) WriteByte(c byte) error {
	if e.quickFail() {
		return e.firstErr
	}
	tmp := e.buf8[:1]
	_, err := e.out.Write(tmp)
	e.noteErr(err)
	return err
}

func (e *Encoder) noteErr(err error) bool {
	if err != nil && e.firstErr == nil {
		e.firstErr = err
	}
	if e.firstErr != nil {
		return true
	}
	return false
}

// Error returns the first occurred error. Each call to any Write* method may cause an error. Per definition,
// any other call after the first error is a no-op.
func (e *Encoder) Error() error {
	return e.firstErr
}
