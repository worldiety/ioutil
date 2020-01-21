package ioutil

import (
	"encoding/binary"
	"fmt"
	"io"
	"math"
	"strconv"
	"unsafe"
)

// A Decoders implements various decoding helpers for Little Endian and Big Endian. It may optimize some
// paths in the future, so that the generic call with byte order may be slower than the direct invocation.
// The implementation reuses an internal buffer and is not thread safe.
type Decoder struct {
	buf8     []byte
	in       io.Reader
	firstErr error
}

func NewDecoder(in io.Reader) *Decoder {
	return &Decoder{
		buf8: make([]byte, 8),
		in:   in,
	}
}

func (r *Decoder) ReadBlob(order binary.ByteOrder, storageClass PrefixType) []byte {
	var bytesToRead uint64
	switch storageClass {
	case Tiny:
		bytesToRead = uint64(r.ReadUint8())
	case Small:
		bytesToRead = uint64(r.ReadUInt16(order))
	case Medium:
		bytesToRead = uint64(r.ReadUInt24(order))
	case Long:
		bytesToRead = uint64(r.ReadUInt32(order))
	case Large:
		bytesToRead = r.ReadUInt64(order)
	case Var:
		t, err := binary.ReadUvarint(r)
		if r.noteErr(err) {
			return nil
		}
		bytesToRead = t
	default:
		panic("invalid PrefixType " + strconv.Itoa(int(storageClass)))
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

func (r *Decoder) ReadUvarint() uint64 {
	t, err := binary.ReadUvarint(r)
	if r.noteErr(err) {
		return 0
	}
	return t
}

func (r *Decoder) ReadVarint() int64 {
	t, err := binary.ReadVarint(r)
	if r.noteErr(err) {
		return 0
	}
	return t
}

func (r *Decoder) ReadBlobLE(p PrefixType) []byte {
	return r.ReadBlob(binary.LittleEndian, p)
}

func (r *Decoder) ReadBlobBE(p PrefixType) []byte {
	return r.ReadBlob(binary.BigEndian, p)
}

// ReadUTF8 provides a type safe conversion to avoid another heap allocation for the
// returned string.
func (r *Decoder) ReadUTF8(order binary.ByteOrder, p PrefixType) string {
	tmp := r.ReadBlob(order, p) // do not change tmp anymore
	// this hack avoids another allocation for the string, see https://github.com/golang/go/issues/25484
	return *(*string)(unsafe.Pointer(&tmp))
}

func (r *Decoder) ReadBool() bool {
	return r.ReadUint8() != 0
}

func (r *Decoder) ReadInt16(order binary.ByteOrder) int16 {
	return int16(r.ReadUInt16(order))
}

func (r *Decoder) ReadInt16LE() int16 {
	return int16(r.ReadUInt16(binary.LittleEndian))
}

func (r *Decoder) ReadInt16BE() int16 {
	return int16(r.ReadUInt16(binary.BigEndian))
}

func (r *Decoder) ReadInt32(order binary.ByteOrder) int32 {
	return int32(r.ReadUInt32(order))
}

func (r *Decoder) ReadInt32LE() int32 {
	return int32(r.ReadUInt32(binary.LittleEndian))
}

func (r *Decoder) ReadInt32BE() int32 {
	return int32(r.ReadUInt32(binary.BigEndian))
}

func (r *Decoder) ReadInt24(order binary.ByteOrder) int32 {
	return int32(r.ReadUInt24(order))
}

func (r *Decoder) ReadInt24LE() int32 {
	return int32(r.ReadUInt24(binary.LittleEndian))
}

func (r *Decoder) ReadInt24BE() int32 {
	return int32(r.ReadUInt24(binary.BigEndian))
}

func (r *Decoder) ReadInt64(order binary.ByteOrder) int64 {
	return int64(r.ReadUInt64(order))
}

func (r *Decoder) ReadInt64LE() int64 {
	return int64(r.ReadUInt64(binary.LittleEndian))
}

func (r *Decoder) ReadInt64BE() int64 {
	return int64(r.ReadUInt64(binary.BigEndian))
}

func (r *Decoder) ReadUInt16(order binary.ByteOrder) uint16 {
	tmp := r.buf8[:2]
	_, err := io.ReadFull(r.in, tmp)
	if r.noteErr(err) {
		return 0
	}
	return order.Uint16(tmp)
}

func (r *Decoder) ReadUInt16LE() uint16 {
	return r.ReadUInt16(binary.LittleEndian)
}

func (r *Decoder) ReadUInt16BE() uint16 {
	return r.ReadUInt16(binary.BigEndian)
}

func (r *Decoder) ReadUInt24(order binary.ByteOrder) uint32 {
	tmp := r.buf8[:3]
	_, err := io.ReadFull(r.in, tmp)
	if r.noteErr(err) {
		return 0
	}
	// this is to slow
	switch order.String() {
	case binary.BigEndian.String():
		return uint24BE(tmp)
	case binary.LittleEndian.String():
		return uint24LE(tmp)
	default:
		panic("unsupported byte order: " + order.String())
	}
}

func (r *Decoder) ReadUInt24LE() uint32 {
	return r.ReadUInt24(binary.LittleEndian)
}

func (r *Decoder) ReadUInt24BE() uint32 {
	return r.ReadUInt24(binary.BigEndian)
}

func (r *Decoder) ReadUInt32(order binary.ByteOrder) uint32 {
	tmp := r.buf8[:4]
	_, err := io.ReadFull(r.in, tmp)
	if r.noteErr(err) {
		return 0
	}
	return order.Uint32(tmp)
}

func (r *Decoder) ReadUInt32LE() uint32 {
	return r.ReadUInt32(binary.LittleEndian)
}

func (r *Decoder) ReadUInt32BE() uint32 {
	return r.ReadUInt32(binary.BigEndian)
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

func (r *Decoder) ReadUInt64(order binary.ByteOrder) uint64 {
	_, err := io.ReadFull(r.in, r.buf8)
	if r.noteErr(err) {
		return 0
	}
	return order.Uint64(r.buf8)
}

func (r *Decoder) ReadUInt64LE() uint64 {
	return r.ReadUInt64(binary.LittleEndian)
}

func (r *Decoder) ReadUInt64BE() uint64 {
	return r.ReadUInt64(binary.BigEndian)
}

func (r *Decoder) ReadFull(b []byte) int {
	n, err := io.ReadFull(r.in, b)
	if r.noteErr(err) {
		return n
	}
	return n
}

func (r *Decoder) ReadUint8() byte {
	tmp := r.buf8[:1]
	_, err := io.ReadFull(r.in, tmp)
	if r.noteErr(err) {
		return 0
	}
	return tmp[0]
}

func (r *Decoder) ReadByte() (byte, error) {
	tmp := r.buf8[:1]
	_, err := io.ReadFull(r.in, tmp)
	if r.noteErr(err) {
		return 0, err
	}
	return tmp[0], nil
}

// Directly delegates the read
func (r *Decoder) Read(buf []byte) (int, error) {
	n, err := r.in.Read(buf)
	r.noteErr(err)
	return n, err
}

func (r *Decoder) Error() error {
	return r.firstErr
}

func (r *Decoder) ReadFloat64(order binary.ByteOrder) float64 {
	bits := r.ReadUInt64(order)
	return math.Float64frombits(bits)
}

func (r *Decoder) ReadFloat64BE() float64 {
	return r.ReadFloat64(binary.BigEndian)
}

func (r *Decoder) ReadFloat64LE() float64 {
	return r.ReadFloat64(binary.LittleEndian)
}

func (r *Decoder) ReadFloat32(order binary.ByteOrder) float32 {
	bits := r.ReadUInt32(order)
	return math.Float32frombits(bits)
}

func (r *Decoder) ReadFloat32BE() float32 {
	return r.ReadFloat32(binary.BigEndian)
}

func (r *Decoder) ReadFloat32LE() float32 {
	return r.ReadFloat32(binary.LittleEndian)
}
