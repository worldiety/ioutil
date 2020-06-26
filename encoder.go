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
	"fmt"
	"io"
	"math"
	"reflect"
	"strconv"
	"unsafe"
)

// An Encoder implements various encoding helpers for Little Endian and Big Endian.
// The implementation reuses an internal buffer to avoid heap allocations and is therefore not thread safe.
type Encoder struct {
	buf10       []byte
	out         io.Writer
	firstErr    error
	failOnError bool
}

// NewEncoder allocates a new encoder instance with a shared buffer
func NewEncoder(out io.Writer, failOnError bool) *Encoder {
	return &Encoder{
		buf10:       make([]byte, binary.MaxVarintLen64),
		out:         out,
		failOnError: failOnError}
}

// Reset removes any error state.
func (e *Encoder) Reset() {
	e.firstErr = nil
}

// quickFail returns true, if an error is already pending and we should not bother the writer again.
func (e *Encoder) quickFail() bool {
	return e.failOnError && e.firstErr != nil
}

// WriteBytes just writes the slice out, without any prefix for the length.
// If an error occurs returns the number of written bytes.
func (e *Encoder) WriteBytes(v ...byte) int {
	return e.WriteSlice(v)
}

// WriteSlice just writes the slice out, without any prefix for the length.
// If an error occurs returns the number of written bytes.
func (e *Encoder) WriteSlice(v []byte) int {
	if e.quickFail() {
		return 0
	}

	n, err := e.out.Write(v)
	if e.noteErr(err) || n != len(v) {
		e.noteErr(fmt.Errorf("writer buffer underrun"))
	}

	return n
}

// WriteBlob writes a prefixed byte slice of variable length.
func (e *Encoder) WriteBlob(o ByteOrder, p IntSize, v []byte) {
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
		if uint32(len(v)) > MaxUint24 {
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
		if uint64(len(v)) > MaxUint40 {
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

	e.WriteSlice(v)
}

// WriteUTF8 writes a prefixed unmodified utf8 string sequence of variable length.
func (e *Encoder) WriteUTF8(o ByteOrder, p IntSize, v string) {
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
	if v {
		e.WriteUint8(1) //nolint:gomnd
	} else {
		e.WriteUint8(0) //nolint:gomnd
	}
}

// WriteUint8 writes an unsigned byte
func (e *Encoder) WriteUint8(v uint8) {
	_ = e.WriteByte(v)
}

// WriteInt8 writes a signed byte
func (e *Encoder) WriteInt8(v int8) {
	e.WriteUint8(uint8(v))
}

// WriteUint16 writes an unsigned 2 byte integer.
func (e *Encoder) WriteUint16(o ByteOrder, v uint16) {
	tmp := e.buf10[:2]
	o.PutUint16(tmp, v)
	e.WriteSlice(tmp)
}

// WriteInt16 writes a signed 2 byte integer.
func (e *Encoder) WriteInt16(o ByteOrder, v int16) {
	e.WriteUint16(o, uint16(v))
}

// WriteUint24 writes an unsigned 3 byte integer.
func (e *Encoder) WriteUint24(o ByteOrder, v uint32) {
	tmp := e.buf10[:3]
	o.PutUint24(tmp, v)
	e.WriteSlice(tmp)
}

// WriteInt24 writes a signed 3 byte integer.
func (e *Encoder) WriteInt24(o ByteOrder, v int32) {
	e.WriteUint24(o, uint32(v))
}

// WriteUint32 writes an unsigned 4 byte integer.
func (e *Encoder) WriteUint32(o ByteOrder, v uint32) {
	tmp := e.buf10[:4]
	o.PutUint32(tmp, v)
	e.WriteSlice(tmp)
}

// WriteInt32 writes a signed 4 byte integer.
func (e *Encoder) WriteInt32(o ByteOrder, v int32) {
	e.WriteUint32(o, uint32(v))
}

// WriteInt40 writes a signed 5 byte integer.
func (e *Encoder) WriteInt40(o ByteOrder, v int64) {
	e.WriteUint40(o, uint64(v))
}

// WriteUint40 writes an unsigned 5 byte integer.
func (e *Encoder) WriteUint40(o ByteOrder, v uint64) {
	tmp := e.buf10[:5]
	o.PutUint40(tmp, v)
	e.WriteSlice(tmp)
}

// WriteInt48 writes a signed 6 byte integer.
func (e *Encoder) WriteInt48(o ByteOrder, v int64) {
	e.WriteUint48(o, uint64(v))
}

// WriteUint48 writes an unsigned 6 byte integer.
func (e *Encoder) WriteUint48(o ByteOrder, v uint64) {
	tmp := e.buf10[:6]
	o.PutUint48(tmp, v)
	e.WriteSlice(tmp)
}

// WriteInt56 writes a signed 7 byte integer.
func (e *Encoder) WriteInt56(o ByteOrder, v int64) {
	e.WriteUint56(o, uint64(v))
}

// WriteUint56 writes an unsigned 7 byte integer.
func (e *Encoder) WriteUint56(o ByteOrder, v uint64) {
	tmp := e.buf10[:7]
	o.PutUint56(tmp, v)
	e.WriteSlice(tmp)
}

// WriteUint64 writes an unsigned 8 byte integer.
func (e *Encoder) WriteUint64(o ByteOrder, v uint64) {
	tmp := e.buf10[:8]
	o.PutUint64(tmp, v)
	e.WriteSlice(tmp)
}

// WriteInt64 writes a signed 8 byte integer.
func (e *Encoder) WriteInt64(o ByteOrder, v int64) {
	e.WriteUint64(o, uint64(v))
}

// WriteUvarint writes a variable length integer, up to 10 bytes using zig-zag protobuf encoding.
func (e *Encoder) WriteUvarint(v uint64) {
	n := binary.PutUvarint(e.buf10, v)
	e.WriteBytes(e.buf10[:n]...)
}

// WriteVarint writes a variable length and signed integer, up to 10 bytes using zig-zag protobuf encoding.
func (e *Encoder) WriteVarint(v int64) {
	n := binary.PutVarint(e.buf10, v)
	e.WriteBytes(e.buf10[:n]...)
}

// WriteFloat32 writes a float32 IEEE 754 4 byte bit sequence.
func (e *Encoder) WriteFloat32(o ByteOrder, v float32) {
	bits := math.Float32bits(v)
	e.WriteUint32(o, bits)
}

// WriteFloat64 writes a float64 IEEE 754 8 byte bit sequence.
func (e *Encoder) WriteFloat64(o ByteOrder, v float64) {
	bits := math.Float64bits(v)
	e.WriteUint64(o, bits)
}

// WriteComplex64 writes two float32 IEEE 754 4 byte bit sequences.
func (e *Encoder) WriteComplex64(o ByteOrder, v complex64) {
	e.WriteFloat32(o, real(v))
	e.WriteFloat32(o, imag(v))
}

// WriteComplex128 writes two float32 IEEE 754 4 byte bit sequences.
func (e *Encoder) WriteComplex128(o ByteOrder, v complex128) {
	e.WriteFloat64(o, real(v))
	e.WriteFloat64(o, imag(v))
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

	tmp := e.buf10[:1]
	tmp[0] = c
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
