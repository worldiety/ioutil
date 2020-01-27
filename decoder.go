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
	"strconv"
	"unsafe"
)

// A Decoder implements various decoding helpers for Little Endian and Big Endian. It may optimize some
// paths in the future, so that the generic call with byte order may be slower than the direct invocation.
// The implementation reuses an internal buffer to avoid heap allocations and is therefore not thread safe.
type Decoder struct {
	buf8        []byte
	in          io.Reader
	firstErr    error
	failOnError bool
}

// NewDecoder wraps a reader to provide the decoder functions. If failOnError is true, any subsequent call
// after an error occurred  will result in a no-op call, so that no more reads will be
// issued to the wrapped reader.
func NewDecoder(in io.Reader, failOnError bool) *Decoder {
	return &Decoder{
		buf8:        make([]byte, 8),
		in:          in,
		failOnError: failOnError,
	}
}

// Reset removes any error state.
func (r *Decoder) Reset() {
	r.firstErr = nil
}

func (r *Decoder) quickFail() bool {
	return r.failOnError && r.firstErr != nil
}

// ReadBlob reads a prefixed byte slice
func (r *Decoder) ReadBlob(order ByteOrder, storageClass IntSize) []byte {
	if r.quickFail() {
		return nil
	}

	var bytesToRead uint64

	switch storageClass {
	case I8:
		bytesToRead = uint64(r.ReadUint8())
	case I16:
		bytesToRead = uint64(r.ReadUint16(order))
	case I24:
		bytesToRead = uint64(r.ReadUint24(order))
	case I32:
		bytesToRead = uint64(r.ReadUint32(order))
	case I40:
		bytesToRead = r.ReadUint40(order)
	case I64:
		bytesToRead = r.ReadUint64(order)
	case IVar:
		t, err := binary.ReadUvarint(r)
		if r.noteErr(err) {
			return nil
		}

		bytesToRead = t
	default:
		panic("invalid IntSize " + strconv.Itoa(int(storageClass)))
	}

	if bytesToRead > MaxInt {
		err := fmt.Errorf("decoded length %d is larger than allowed (%d)", bytesToRead, MaxInt)
		if r.noteErr(err) {
			return nil
		}
	}

	buf := make([]byte, int(bytesToRead))
	r.ReadFull(buf)

	return buf
}

// ReadBytes just reads a bunch of bytes into a newly allocated buffer
func (r *Decoder) ReadBytes(len int) []byte {
	if r.quickFail() {
		return nil
	}

	buf := make([]byte, len)
	n := r.ReadFull(buf)

	return buf[0:n]
}

// ReadUvarint reads a variable length integer, up to 10 bytes using zig-zag protobuf encoding.
func (r *Decoder) ReadUvarint() uint64 {
	if r.quickFail() {
		return 0
	}

	t, err := binary.ReadUvarint(r)

	if r.noteErr(err) {
		return 0
	}

	return t
}

// ReadVarint reads a variable length and signed integer, up to 10 bytes using zig-zag protobuf encoding.
func (r *Decoder) ReadVarint() int64 {
	if r.quickFail() {
		return 0
	}

	t, err := binary.ReadVarint(r)
	if r.noteErr(err) {
		return 0
	}

	return t
}

// ReadUTF8 provides a type safe conversion to avoid another heap allocation for the
// returned string.
func (r *Decoder) ReadUTF8(order ByteOrder, p IntSize) string {
	tmp := r.ReadBlob(order, p) // do not change tmp anymore
	// this hack avoids another allocation for the string, see https://github.com/golang/go/issues/25484
	return *(*string)(unsafe.Pointer(&tmp))
}

// ReadBool reads one byte and returns 0 if the byte is zero, otherwise true
func (r *Decoder) ReadBool() bool {
	return r.ReadUint8() != 0
}

// ReadInt16 reads 2 bytes and interprets them as signed
func (r *Decoder) ReadInt16(order ByteOrder) int16 {
	return int16(r.ReadUint16(order))
}

// ReadInt24 reads 3 bytes and interprets them as signed
func (r *Decoder) ReadInt24(order ByteOrder) int32 {
	return int32(r.ReadUint24(order))
}

// ReadInt32 reads 4 bytes and interprets them as signed
func (r *Decoder) ReadInt32(order ByteOrder) int32 {
	return int32(r.ReadUint32(order))
}

// ReadInt40 reads 5 bytes and interprets them as signed
func (r *Decoder) ReadInt40(order ByteOrder) int64 {
	return int64(r.ReadUint32(order))
}

// ReadInt48 reads 6 bytes and interprets them as signed
func (r *Decoder) ReadInt48(order ByteOrder) int64 {
	return int64(r.ReadUint32(order))
}

// ReadInt56 reads 7 bytes and interprets them as signed
func (r *Decoder) ReadInt56(order ByteOrder) int64 {
	return int64(r.ReadUint32(order))
}

// ReadInt64 reads 7 bytes and interprets them as signed
func (r *Decoder) ReadInt64(order ByteOrder) int64 {
	return int64(r.ReadUint64(order))
}

// ReadUint16 reads 2 bytes and interprets them as unsigned
func (r *Decoder) ReadUint16(order ByteOrder) uint16 {
	if r.quickFail() {
		return 0
	}

	tmp := r.buf8[:2]
	_, err := io.ReadFull(r.in, tmp)

	if r.noteErr(err) {
		return 0
	}

	return order.Uint16(tmp)
}

// ReadUint24 reads 3 bytes and interprets them as unsigned
func (r *Decoder) ReadUint24(order ByteOrder) uint32 {
	if r.quickFail() {
		return 0
	}

	tmp := r.buf8[:3]
	_, err := io.ReadFull(r.in, tmp)

	if r.noteErr(err) {
		return 0
	}

	return order.Uint24(tmp)
}

// ReadUint32 reads 4 bytes and interprets them as unsigned
func (r *Decoder) ReadUint32(order ByteOrder) uint32 {
	if r.quickFail() {
		return 0
	}

	tmp := r.buf8[:4]
	_, err := io.ReadFull(r.in, tmp)

	if r.noteErr(err) {
		return 0
	}

	return order.Uint32(tmp)
}

// ReadUint40 reads 5 bytes and interprets them as unsigned
func (r *Decoder) ReadUint40(order ByteOrder) uint64 {
	if r.quickFail() {
		return 0
	}

	tmp := r.buf8[:5]
	_, err := io.ReadFull(r.in, tmp)

	if r.noteErr(err) {
		return 0
	}

	return order.Uint40(tmp)
}

// ReadUint48 reads 6 bytes and interprets them as unsigned
func (r *Decoder) ReadUint48(order ByteOrder) uint64 {
	if r.quickFail() {
		return 0
	}

	tmp := r.buf8[:6]
	_, err := io.ReadFull(r.in, tmp)

	if r.noteErr(err) {
		return 0
	}

	return order.Uint48(tmp)
}

// ReadUint56 reads 7 bytes and interprets them as unsigned
func (r *Decoder) ReadUint56(order ByteOrder) uint64 {
	if r.quickFail() {
		return 0
	}

	tmp := r.buf8[:7]
	_, err := io.ReadFull(r.in, tmp)

	if r.noteErr(err) {
		return 0
	}

	return order.Uint56(tmp)
}

// ReadUint64 reads 8 bytes and interprets them as unsigned
func (r *Decoder) ReadUint64(order ByteOrder) uint64 {
	if r.quickFail() {
		return 0
	}

	_, err := io.ReadFull(r.in, r.buf8)
	if r.noteErr(err) {
		return 0
	}

	return order.Uint64(r.buf8)
}

// ReadFull reads exactly len(b) bytes. If an error occurs returns the number of read bytes.
func (r *Decoder) ReadFull(b []byte) int {
	n, err := io.ReadFull(r.in, b)
	if r.noteErr(err) {
		return n
	}

	return n
}

// ReadUint8 reads one byte
func (r *Decoder) ReadUint8() uint8 {
	if r.quickFail() {
		return 0
	}

	tmp := r.buf8[:1]
	_, err := io.ReadFull(r.in, tmp)

	if r.noteErr(err) {
		return 0
	}

	return tmp[0]
}

// ReadInt8 reads one byte
func (r *Decoder) ReadInt8() int8 {
	return int8(r.ReadUint8())
}

// ReadByte reads one byte
func (r *Decoder) ReadByte() (byte, error) {
	if r.quickFail() {
		return 0, r.firstErr
	}

	tmp := r.buf8[:1]
	_, err := io.ReadFull(r.in, tmp)

	if r.noteErr(err) {
		return 0, err
	}

	return tmp[0], nil
}

// Directly delegates the read
func (r *Decoder) Read(buf []byte) (int, error) {
	if r.quickFail() {
		return 0, r.firstErr
	}

	n, err := r.in.Read(buf)
	r.noteErr(err)

	return n, err
}

// ReadFloat64 reads 8 bytes and interprets them as a float64 IEEE 754 4 byte bit sequence.
func (r *Decoder) ReadFloat64(order ByteOrder) float64 {
	bits := r.ReadUint64(order)
	return math.Float64frombits(bits)
}

// ReadFloat32 reads 4 bytes and interprets them as a float32 IEEE 754 4 byte bit sequence.
func (r *Decoder) ReadFloat32(order ByteOrder) float32 {
	bits := r.ReadUint32(order)
	return math.Float32frombits(bits)
}

// ReadComplex64 reads two float32 IEEE 754 4 byte bit sequences for the real and imaginary parts.
func (r *Decoder) ReadComplex64(order ByteOrder) complex64 {
	rnum := r.ReadFloat32(order)
	inum := r.ReadFloat32(order)

	return complex(rnum, inum)
}

// ReadComplex128 reads two float64 IEEE 754 8 byte bit sequences for the real and imaginary parts.
func (r *Decoder) ReadComplex128(order ByteOrder) complex128 {
	rnum := r.ReadFloat64(order)
	inum := r.ReadFloat64(order)

	return complex(rnum, inum)
}

func (r *Decoder) noteErr(err error) bool {
	if err != nil && r.firstErr == nil {
		r.firstErr = err
	}

	if r.firstErr != nil {
		return true
	}

	return false
}

// Error returns the first occurred error. Each call to any Read* method may cause an error.
func (r *Decoder) Error() error {
	return r.firstErr
}
