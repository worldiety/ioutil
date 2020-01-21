package ioutil

import (
	"encoding/binary"
	"io"
)

var _ DataInput = (*dataInputImpl)(nil)

// A DataInput provides helpers for reading bytes from a binary stream and interprets the data for any primitive Go
// type. A DataInput is always tied to a specific endianness. A DataInput should not be considered thread safe.
// As soon as any error occurred, any call is a no-op and will result in the same error state.
type DataInput interface {

	// ReadBlob reads a prefixed byte slice
	ReadBlob(p IntSize) []byte

	// ReadUTF8 reads a prefixed unmodified utf8 string sequence
	ReadUTF8(p IntSize) string

	// Reads one byte and returns 0 if the byte is zero, otherwise true
	ReadBool() bool

	// Reads one byte
	ReadUint8() uint8

	// Reads one byte
	ReadInt8() int8

	//
	ReadUInt16() uint16

	ReadUInt24() uint32

	ReadUInt32() uint32

	ReadUInt64() uint64

	ReadInt16() int16

	ReadInt24() int32

	ReadInt32() int32

	ReadInt40() int64

	ReadInt64() int64

	// ReadUvarint reads a variable length integer, up to 9 bytes using zig-zag protobuf encoding.
	ReadUvarint() uint64

	// ReadVarint reads a variable length and signed integer, up to 9 bytes using zig-zag protobuf encoding.
	ReadVarint() int64

	ReadFloat32() float32

	ReadFloat64() float64

	ReadComplex64() complex64

	ReadComplex128() complex128

	// ReadFull reads exactly len(b) bytes. If an error occurs returns the number of read bytes.
	ReadFull(b []byte) int

	// Error returns the first occurred error. Each call to any Read* method may cause an error.
	Error() error

	io.Reader
	io.ByteReader
}

func NewDataInput(order binary.ByteOrder, reader io.Reader) DataInput {
	return dataInputImpl{decoder: NewDecoder(reader, true), order: order}
}

type dataInputImpl struct {
	order   binary.ByteOrder
	decoder *Decoder
}

func (d dataInputImpl) ReadComplex64() complex64 {
	return d.decoder.ReadComplex64(d.order)
}

func (d dataInputImpl) ReadComplex128() complex128 {
	return d.decoder.ReadComplex128(d.order)
}

func (d dataInputImpl) ReadInt40() int64 {
	return d.ReadInt40()
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
	return d.decoder.ReadUInt16(d.order)
}

func (d dataInputImpl) ReadUInt24() uint32 {
	return d.decoder.ReadUInt24(d.order)
}

func (d dataInputImpl) ReadUInt32() uint32 {
	return d.decoder.ReadUInt32(d.order)
}

func (d dataInputImpl) ReadUInt64() uint64 {
	return d.decoder.ReadUInt64(d.order)
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
