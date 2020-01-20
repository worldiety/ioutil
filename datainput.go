package ioutil

import (
	"encoding/binary"
	"io"
)

var _ DataInput = (*DataInputReader)(nil)

// A DataInput provides helpers for reading bytes from a binary stream and interprets the data for any primitive Go
// type.
type DataInput interface {

	// ReadUTF8 reads a 2 byte prefixed
	ReadUTF8() string

	ReadBool() bool

	ReadUInt16(order binary.ByteOrder) uint16
	ReadUInt16LE() uint16
	ReadUInt16BE() uint16

	ReadUInt32(order binary.ByteOrder) uint32
	ReadUInt32LE() uint32
	ReadUInt32BE() uint32

	ReadUInt64(order binary.ByteOrder) uint64
	ReadUInt64LE() uint64
	ReadUInt64BE() uint64

	ReadInt16(order binary.ByteOrder) int16
	ReadInt16LE() int16
	ReadInt16BE() int16

	ReadInt32(order binary.ByteOrder) int32
	ReadInt32LE() int32
	ReadInt32BE() int32

	ReadInt64(order binary.ByteOrder) int64
	ReadInt64LE() int64
	ReadInt64BE() int64

	ReadByte() byte
	Read(buf []byte) (int, error)

	// ReadFull reads exactly len(b) bytes
	ReadFull(b []byte) int

	// Error returns the first occurred error. Each call to any Read* method may cause an error.
	Error() error
}

type DataInputReader struct {
	buf8     []byte
	in       io.Reader
	firstErr error
}

func NewDataInputReader(in io.Reader) *DataInputReader {
	return &DataInputReader{
		buf8: make([]byte, 8),
		in:   in,
	}
}

func (r *DataInputReader) ReadBool() bool {
	return r.ReadByte() != 0
}

func (r *DataInputReader) ReadInt16(order binary.ByteOrder) int16 {
	return int16(r.ReadUInt16(order))
}

func (r *DataInputReader) ReadInt16LE() int16 {
	return int16(r.ReadUInt16(binary.LittleEndian))
}

func (r *DataInputReader) ReadInt16BE() int16 {
	return int16(r.ReadUInt16(binary.BigEndian))
}

func (r *DataInputReader) ReadInt32(order binary.ByteOrder) int32 {
	return int32(r.ReadUInt32(order))
}

func (r *DataInputReader) ReadInt32LE() int32 {
	return int32(r.ReadUInt32(binary.LittleEndian))
}

func (r *DataInputReader) ReadInt32BE() int32 {
	return int32(r.ReadUInt32(binary.BigEndian))
}

func (r *DataInputReader) ReadInt64(order binary.ByteOrder) int64 {
	return int64(r.ReadUInt64(order))
}

func (r *DataInputReader) ReadInt64LE() int64 {
	return int64(r.ReadUInt64(binary.LittleEndian))
}

func (r *DataInputReader) ReadInt64BE() int64 {
	return int64(r.ReadUInt64(binary.BigEndian))
}

func (r *DataInputReader) ReadUInt16(order binary.ByteOrder) uint16 {
	tmp := r.buf8[:2]
	_, err := io.ReadFull(r.in, tmp)
	if r.noteErr(err) {
		return 0
	}
	return order.Uint16(tmp)
}

func (r *DataInputReader) ReadUInt16LE() uint16 {
	return r.ReadUInt16(binary.LittleEndian)
}

func (r *DataInputReader) ReadUInt16BE() uint16 {
	return r.ReadUInt16(binary.BigEndian)
}

func (r *DataInputReader) ReadUInt32(order binary.ByteOrder) uint32 {
	tmp := r.buf8[:4]
	_, err := io.ReadFull(r.in, tmp)
	if r.noteErr(err) {
		return 0
	}
	return order.Uint32(tmp)
}

func (r *DataInputReader) ReadUInt32LE() uint32 {
	return r.ReadUInt32(binary.LittleEndian)
}

func (r *DataInputReader) ReadUInt32BE() uint32 {
	return r.ReadUInt32(binary.BigEndian)
}

func (r *DataInputReader) noteErr(err error) bool {
	if err != nil && r.firstErr == nil {
		r.firstErr = err
	}
	if r.firstErr != nil {
		return true
	}
	return false
}

func (r *DataInputReader) ReadUInt64(order binary.ByteOrder) uint64 {
	_, err := io.ReadFull(r.in, r.buf8)
	if r.noteErr(err) {
		return 0
	}
	return order.Uint64(r.buf8)
}

func (r *DataInputReader) ReadUInt64LE() uint64 {
	return r.ReadUInt64(binary.LittleEndian)
}

func (r *DataInputReader) ReadUInt64BE() uint64 {
	return r.ReadUInt64(binary.BigEndian)
}

func (r *DataInputReader) ReadFull(b []byte) int {
	n, err := io.ReadFull(r.in, b)
	if r.noteErr(err) {
		return n
	}
	return n
}

func (r *DataInputReader) ReadByte() byte {
	tmp := r.buf8[:1]
	_, err := io.ReadFull(r.in, tmp)
	if r.noteErr(err) {
		return 0
	}
	return tmp[0]
}

// Directly delegates the read
func (r *DataInputReader) Read(buf []byte) (int, error) {
	n, err := r.in.Read(buf)
	r.noteErr(err)
	return n, err
}

func (r *DataInputReader) Error() error {
	return r.firstErr
}
