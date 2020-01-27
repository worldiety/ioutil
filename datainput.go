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
	"io"
)

// A DataInput provides helpers for reading bytes from a binary stream and interprets the data for any primitive Go
// type. A DataInput is always tied to a specific endianness. A DataInput should not be considered thread safe.
// As soon as any error occurred, any call is a no-op and will result in the same error state.
type DataInput interface {

	// ReadBlob reads a prefixed byte slice
	ReadBlob(p IntSize) []byte

	// ReadUTF8 reads a prefixed unmodified utf8 string sequence
	ReadUTF8(p IntSize) string

	// ReadBool reads one byte and returns 0 if the byte is zero, otherwise true
	ReadBool() bool

	// ReadUint8 reads one byte
	ReadUint8() uint8

	// ReadBytes just reads a bunch of bytes into a newly allocated buffer
	ReadBytes(len int) []byte

	// ReadUint16 reads 2 bytes and interprets them as unsigned
	ReadUint16() uint16

	// ReadUint24 reads 3 bytes and interprets them as unsigned
	ReadUint24() uint32

	// ReadUint32 reads 4 bytes and interprets them as unsigned
	ReadUint32() uint32

	// ReadUint40 reads 5 bytes and interprets them as unsigned
	ReadUint40() uint64

	// ReadUint48 reads 6 bytes and interprets them as unsigned
	ReadUint48() uint64

	// ReadUint56 reads 7 bytes and interprets them as unsigned
	ReadUint56() uint64

	// ReadUint64 reads 8 bytes and interprets them as unsigned
	ReadUint64() uint64

	// ReadInt8 reads one byte
	ReadInt8() int8

	// ReadUint16 reads 2 bytes and interprets them as signed
	ReadInt16() int16

	// ReadInt24 reads 3 bytes and interprets them as signed
	ReadInt24() int32

	// ReadInt32 reads 4 bytes and interprets them as signed
	ReadInt32() int32

	// ReadInt40 reads 5 bytes and interprets them as signed
	ReadInt40() int64

	// ReadInt48 reads 6 bytes and interprets them as signed
	ReadInt48() int64

	// ReadInt56 reads 7 bytes and interprets them as signed
	ReadInt56() int64

	// ReadInt64 reads 8 bytes and interprets them as signed
	ReadInt64() int64

	// ReadUvarint reads a variable length integer, up to 10 bytes using zig-zag protobuf encoding.
	ReadUvarint() uint64

	// ReadVarint reads a variable length and signed integer, up to 10 bytes using zig-zag protobuf encoding.
	ReadVarint() int64

	// ReadFloat32 reads 4 bytes and interprets them as a float32 IEEE 754 4 byte bit sequence.
	ReadFloat32() float32

	// ReadFloat64 reads 8 bytes and interprets them as a float64 IEEE 754 4 byte bit sequence.
	ReadFloat64() float64

	// ReadComplex64 reads two float32 IEEE 754 4 byte bit sequences for the real and imaginary parts.
	ReadComplex64() complex64

	// ReadComplex128 reads two float64 IEEE 754 8 byte bit sequences for the real and imaginary parts.
	ReadComplex128() complex128

	// ReadFull reads exactly len(b) bytes. If an error occurs returns the number of read bytes.
	ReadFull(b []byte) int

	// Error returns the first occurred error. Each call to any Read* method may cause an error.
	Error() error

	io.Reader
	io.ByteReader
}

// NewDataInput creates a new DataInput instance according to the given byte order
func NewDataInput(order ByteOrder, reader io.Reader) DataInput {
	return dataInputImpl{decoder: NewDecoder(reader, true), order: order}
}

var _ DataInput = (*dataInputImpl)(nil)

type dataInputImpl struct {
	order   ByteOrder
	decoder *Decoder
}

func (d dataInputImpl) ReadBytes(len int) []byte {
	return d.decoder.ReadBytes(len)
}

func (d dataInputImpl) ReadUint16() uint16 {
	return d.decoder.ReadUint16(d.order)
}

func (d dataInputImpl) ReadUint24() uint32 {
	return d.decoder.ReadUint24(d.order)
}

func (d dataInputImpl) ReadUint32() uint32 {
	return d.decoder.ReadUint32(d.order)
}

func (d dataInputImpl) ReadUint40() uint64 {
	return d.decoder.ReadUint40(d.order)
}

func (d dataInputImpl) ReadUint48() uint64 {
	return d.decoder.ReadUint48(d.order)
}

func (d dataInputImpl) ReadUint56() uint64 {
	return d.decoder.ReadUint56(d.order)
}

func (d dataInputImpl) ReadUint64() uint64 {
	return d.decoder.ReadUint64(d.order)
}

func (d dataInputImpl) ReadComplex64() complex64 {
	return d.decoder.ReadComplex64(d.order)
}

func (d dataInputImpl) ReadComplex128() complex128 {
	return d.decoder.ReadComplex128(d.order)
}

func (d dataInputImpl) ReadInt40() int64 {
	return d.decoder.ReadInt40(d.order)
}

func (d dataInputImpl) ReadFloat32() float32 {
	return d.decoder.ReadFloat32(d.order)
}

func (d dataInputImpl) ReadFloat64() float64 {
	return d.decoder.ReadFloat64(d.order)
}

func (d dataInputImpl) ReadUint8() uint8 {
	return d.decoder.ReadUint8()
}

func (d dataInputImpl) ReadByte() (byte, error) {
	return d.decoder.ReadByte()
}

func (d dataInputImpl) ReadInt8() int8 {
	return d.decoder.ReadInt8()
}

func (d dataInputImpl) ReadBlob(p IntSize) []byte {
	return d.decoder.ReadBlob(d.order, p)
}

func (d dataInputImpl) ReadUTF8(p IntSize) string {
	return d.decoder.ReadUTF8(d.order, p)
}

func (d dataInputImpl) ReadBool() bool {
	return d.decoder.ReadBool()
}

func (d dataInputImpl) ReadUInt16() uint16 {
	return d.decoder.ReadUint16(d.order)
}

func (d dataInputImpl) ReadUInt24() uint32 {
	return d.decoder.ReadUint24(d.order)
}

func (d dataInputImpl) ReadUInt32() uint32 {
	return d.decoder.ReadUint32(d.order)
}

func (d dataInputImpl) ReadUInt64() uint64 {
	return d.decoder.ReadUint64(d.order)
}

func (d dataInputImpl) ReadInt16() int16 {
	return d.decoder.ReadInt16(d.order)
}

func (d dataInputImpl) ReadInt24() int32 {
	return d.decoder.ReadInt24(d.order)
}

func (d dataInputImpl) ReadInt32() int32 {
	return d.decoder.ReadInt32(d.order)
}

func (d dataInputImpl) ReadInt48() int64 {
	return d.decoder.ReadInt48(d.order)
}

func (d dataInputImpl) ReadInt56() int64 {
	return d.decoder.ReadInt56(d.order)
}

func (d dataInputImpl) ReadInt64() int64 {
	return d.decoder.ReadInt64(d.order)
}

func (d dataInputImpl) ReadUvarint() uint64 {
	return d.decoder.ReadUvarint()
}

func (d dataInputImpl) ReadVarint() int64 {
	return d.decoder.ReadVarint()
}

func (d dataInputImpl) Read(buf []byte) (int, error) {
	return d.decoder.Read(buf)
}

func (d dataInputImpl) ReadFull(b []byte) int {
	return d.decoder.ReadFull(b)
}

func (d dataInputImpl) Error() error {
	return d.decoder.Error()
}
