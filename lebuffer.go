package ioutil

import (
	"math"
	"reflect"
	"strconv"
	"unsafe"
)

// LittleEndianBuffer is a light weight helper to modify bytes within a buffer in little endian format.
type LittleEndianBuffer struct {
	Bytes []byte
	Pos   int
}

func (f *LittleEndianBuffer) ReadUint8() uint8 {
	b := f.Bytes[f.Pos]
	f.Pos++
	return b
}

func (f *LittleEndianBuffer) WriteUint8(v uint8) {
	f.Bytes[f.Pos] = v
	f.Pos++
}

func (f *LittleEndianBuffer) postInc()int{
	i := f.Pos
	f.Pos++
	return i
}

func (f *LittleEndianBuffer) ReadUint16() uint16 {
/*	b := f.Bytes[f.Pos:]
	f.Pos += 2
	_ = b[1] // bounds check hint to compiler; see golang.org/issue/14808
	return uint16(b[0]) | uint16(b[1])<<8
*/

	// equal or slower
	//return uint16(f.Bytes[f.postInc()]) | uint16(f.Bytes[f.postInc()])<<8

	// the following is 40% faster in microbenchmark than all above
	_ = f.Bytes[1+f.Pos]
	i := uint16(f.Bytes[f.Pos]) | uint16(f.Bytes[f.Pos+1])<<8
	f.Pos+=2
	return i
}

func (f *LittleEndianBuffer) WriteUint16(v uint16) {
	b := f.Bytes[f.Pos:]
	f.Pos += 2
	_ = b[1] // bounds check hint to compiler; see golang.org/issue/14808
	b[0] = byte(v)
	b[1] = byte(v >> 8)
}

func (f *LittleEndianBuffer) ReadUint24() uint32 {
	b := f.Bytes[f.Pos:]
	f.Pos += 3

	_ = b[2] // bounds check hint to compiler; see golang.org/issue/14808
	return uint32(b[0]) | uint32(b[1])<<8 | uint32(b[2])<<16
}

func (f *LittleEndianBuffer) WriteUint24(v uint32) {
	b := f.Bytes[f.Pos:]
	f.Pos += 3

	_ = b[2] // early bounds check to guarantee safety of writes below
	b[0] = byte(v)
	b[1] = byte(v >> 8)  //nolint:gomnd
	b[2] = byte(v >> 16) //nolint:gomnd
}

func (f *LittleEndianBuffer) ReadUint32() uint32 {

	b := f.Bytes[f.Pos:]
	f.Pos += 4
	_ = b[3] // bounds check hint to compiler; see golang.org/issue/14808
	return uint32(b[0]) | uint32(b[1])<<8 | uint32(b[2])<<16 | uint32(b[3])<<24

	// the following is 50% slower
	/*
	_ = f.Bytes[3+f.Pos]
	i := uint32(f.Bytes[f.Pos]) | uint32(f.Bytes[f.Pos+1])<<8 | uint32(f.Bytes[f.Pos+2])<<16 | uint32(f.Bytes[f.Pos+3])<<24
	f.Pos+=4
	return i*/
}

func (f *LittleEndianBuffer) WriteUint32(v uint32) {
	b := f.Bytes[f.Pos:]
	f.Pos += 4

	_ = b[3] // bounds check hint to compiler; see golang.org/issue/14808
	b[0] = byte(v)
	b[1] = byte(v >> 8)
	b[2] = byte(v >> 16)
	b[3] = byte(v >> 24)
}

func (f *LittleEndianBuffer) ReadUint40() uint64 {
	b := f.Bytes[f.Pos:]
	f.Pos += 5

	_ = b[4] // bounds check hint to compiler; see golang.org/issue/14808
	return uint64(b[0]) | uint64(b[1])<<8 | uint64(b[2])<<16 | uint64(b[3])<<24 | uint64(b[4])<<32
}

func (f *LittleEndianBuffer) WriteUint40(v uint64) {
	b := f.Bytes[f.Pos:]
	f.Pos += 5

	_ = b[4] // bounds check hint to compiler; see golang.org/issue/14808
	b[0] = byte(v)
	b[1] = byte(v >> 8)
	b[2] = byte(v >> 16)
	b[3] = byte(v >> 24)
	b[4] = byte(v >> 32)
}

func (f *LittleEndianBuffer) ReadUint48() uint64 {
	b := f.Bytes[f.Pos:]
	f.Pos += 6

	_ = b[5] // bounds check hint to compiler; see golang.org/issue/14808
	return uint64(b[0]) | uint64(b[1])<<8 | uint64(b[2])<<16 | uint64(b[3])<<24 |
		uint64(b[4])<<32 | uint64(b[5])<<40
}

func (f *LittleEndianBuffer) WriteUint48(v uint64) {
	b := f.Bytes[f.Pos:]
	f.Pos += 6

	_ = b[5] // bounds check hint to compiler; see golang.org/issue/14808
	b[0] = byte(v)
	b[1] = byte(v >> 8)
	b[2] = byte(v >> 16)
	b[3] = byte(v >> 24)
	b[4] = byte(v >> 32)
	b[5] = byte(v >> 40)
}

func (f *LittleEndianBuffer) ReadUint56() uint64 {
	b := f.Bytes[f.Pos:]
	f.Pos += 7

	_ = b[6] // bounds check hint to compiler; see golang.org/issue/14808
	return uint64(b[0]) | uint64(b[1])<<8 | uint64(b[2])<<16 | uint64(b[3])<<24 |
		uint64(b[4])<<32 | uint64(b[5])<<40 | uint64(b[6])<<48
}

func (f *LittleEndianBuffer) WriteUint56(v uint64) {
	b := f.Bytes[f.Pos:]
	f.Pos += 7

	_ = b[6] // bounds check hint to compiler; see golang.org/issue/14808
	b[0] = byte(v)
	b[1] = byte(v >> 8)
	b[2] = byte(v >> 16)
	b[3] = byte(v >> 24)
	b[4] = byte(v >> 32)
	b[5] = byte(v >> 40)
	b[6] = byte(v >> 48)
}

func (f *LittleEndianBuffer) ReadUint64() uint64 {
	b := f.Bytes[f.Pos:]
	f.Pos += 8

	_ = b[7] // bounds check hint to compiler; see golang.org/issue/14808
	return uint64(b[0]) | uint64(b[1])<<8 | uint64(b[2])<<16 | uint64(b[3])<<24 |
		uint64(b[4])<<32 | uint64(b[5])<<40 | uint64(b[6])<<48 | uint64(b[7])<<56
}

func (f *LittleEndianBuffer) WriteUint64(v uint64) {
	b := f.Bytes[f.Pos:]
	f.Pos += 8

	_ = b[7] // bounds check hint to compiler; see golang.org/issue/14808
	b[0] = byte(v)
	b[1] = byte(v >> 8)
	b[2] = byte(v >> 16)
	b[3] = byte(v >> 24)
	b[4] = byte(v >> 32)
	b[5] = byte(v >> 40)
	b[6] = byte(v >> 48)
	b[7] = byte(v >> 56)
}

// WriteSlice copies the content of the given buffer into the destination
func (f *LittleEndianBuffer) WriteSlice(v []byte) {
	b := f.Bytes[f.Pos : f.Pos+len(v)]
	copy(b, v)
	f.Pos += len(v)
}

// ReadSlice reads fully into the given buffer
func (f *LittleEndianBuffer) ReadSlice(v []byte) {
	b := f.Bytes[f.Pos : f.Pos+len(v)]
	copy(v, b)
	f.Pos += len(v)
}

// ReadBlob8 reads up to 255 bytes. The blob is truncated.
func (f *LittleEndianBuffer) ReadBlob8(v []byte) int {
	vLen := f.ReadUint8()
	vBuf := v[0:vLen]

	f.ReadSlice(vBuf)
	return int(vLen)
}

// WriteBlob8 writes up to 255 bytes. The blob is truncated.
func (f *LittleEndianBuffer) WriteBlob8(v []byte) {
	vLen := len(v)
	if vLen > int(MaxUint8) {
		vLen = int(MaxUint8)
	}

	f.WriteUint8(uint8(vLen))
	f.WriteSlice(v[:vLen])
}

// ReadBlob16 reads up to 65535 bytes. The blob is truncated.
func (f *LittleEndianBuffer) ReadBlob16(v []byte) int {
	vLen := f.ReadUint16()
	vBuf := v[0:vLen]

	f.ReadSlice(vBuf)
	return int(vLen)
}

// WriteBlob16 writes up to 65535 bytes. The blob is truncated.
func (f *LittleEndianBuffer) WriteBlob16(v []byte) {
	vLen := len(v)
	if vLen > int(MaxUint16) {
		vLen = int(MaxUint16)
	}

	f.WriteUint16(uint16(vLen))
	f.WriteSlice(v[:vLen])
}

// ReadBlob16 reads up to 16777215 bytes. The blob is truncated.
func (f *LittleEndianBuffer) ReadBlob24(v []byte) int {
	vLen := f.ReadUint24()
	vBuf := v[0:vLen]

	f.ReadSlice(vBuf)
	return int(vLen)
}

// WriteBlob16 writes up to 16777215 bytes. The blob is truncated.
func (f *LittleEndianBuffer) WriteBlob24(v []byte) {
	vLen := len(v)
	if vLen > int(MaxUint24) {
		vLen = int(MaxUint24)
	}

	f.WriteUint24(uint32(vLen))
	f.WriteSlice(v[:vLen])
}

// ReadBlob32 reads up to 4294967295 bytes. The blob is truncated.
func (f *LittleEndianBuffer) ReadBlob32(v []byte) int {
	vLen := f.ReadUint32()
	vBuf := v[0:vLen]

	f.ReadSlice(vBuf)
	return int(vLen)
}

// WriteBlob32 writes up to 4294967295 bytes. The blob is truncated.
func (f *LittleEndianBuffer) WriteBlob32(v []byte) {
	vLen := len(v)
	if vLen > int(MaxUint32) {
		vLen = int(MaxUint32)
	}

	f.WriteUint32(uint32(vLen))
	f.WriteSlice(v[:vLen])
}

// WriteString8 writes the string into a blob, avoiding another allocation.
func (f *LittleEndianBuffer) WriteString8(v string) {
	str := *(*reflect.StringHeader)(unsafe.Pointer(&v))
	// do not modify the slice, because this is a hack to avoid an unnecessary copy and heap allocation
	slice := *(*[]byte)(unsafe.Pointer(&reflect.SliceHeader{
		Data: str.Data,
		Len:  str.Len,
		Cap:  str.Len,
	}))

	f.WriteBlob8(slice)
}

// ReadString8 creates a (mutable) string, by using the strBuffer.
func (f *LittleEndianBuffer) ReadString8(strBuffer []byte) string {
	vLen := f.ReadBlob8(strBuffer)
	strBuffer = strBuffer[:vLen]
	// this hack avoids another allocation for the string, see https://github.com/golang/go/issues/25484
	return *(*string)(unsafe.Pointer(&strBuffer))
}

// WriteString16 writes the string into a blob, avoiding another allocation.
func (f *LittleEndianBuffer) WriteString16(v string) {
	str := *(*reflect.StringHeader)(unsafe.Pointer(&v))
	// do not modify the slice, because this is a hack to avoid an unnecessary copy and heap allocation
	slice := *(*[]byte)(unsafe.Pointer(&reflect.SliceHeader{
		Data: str.Data,
		Len:  str.Len,
		Cap:  str.Len,
	}))

	f.WriteBlob16(slice)
}

// ReadString16 creates a (mutable) string, by using the strBuffer.
func (f *LittleEndianBuffer) ReadString16(strBuffer []byte) string {
	vLen := f.ReadBlob16(strBuffer)
	strBuffer = strBuffer[:vLen]
	// this hack avoids another allocation for the string, see https://github.com/golang/go/issues/25484
	return *(*string)(unsafe.Pointer(&strBuffer))
}

// WriteString24 writes the string into a blob, avoiding another allocation.
func (f *LittleEndianBuffer) WriteString24(v string) {
	str := *(*reflect.StringHeader)(unsafe.Pointer(&v))
	// do not modify the slice, because this is a hack to avoid an unnecessary copy and heap allocation
	slice := *(*[]byte)(unsafe.Pointer(&reflect.SliceHeader{
		Data: str.Data,
		Len:  str.Len,
		Cap:  str.Len,
	}))

	f.WriteBlob24(slice)
}

// ReadString24 creates a (mutable) string, by using the strBuffer.
func (f *LittleEndianBuffer) ReadString24(strBuffer []byte) string {
	vLen := f.ReadBlob24(strBuffer)
	strBuffer = strBuffer[:vLen]
	// this hack avoids another allocation for the string, see https://github.com/golang/go/issues/25484
	return *(*string)(unsafe.Pointer(&strBuffer))
}

// WriteString32 writes the string into a blob, avoiding another allocation.
func (f *LittleEndianBuffer) WriteString32(v string) {
	str := *(*reflect.StringHeader)(unsafe.Pointer(&v))
	// do not modify the slice, because this is a hack to avoid an unnecessary copy and heap allocation
	slice := *(*[]byte)(unsafe.Pointer(&reflect.SliceHeader{
		Data: str.Data,
		Len:  str.Len,
		Cap:  str.Len,
	}))

	f.WriteBlob32(slice)
}

// ReadString32 creates a (mutable) string, by using the strBuffer.
func (f *LittleEndianBuffer) ReadString32(strBuffer []byte) string {
	vLen := f.ReadBlob32(strBuffer)
	strBuffer = strBuffer[:vLen]
	// this hack avoids another allocation for the string, see https://github.com/golang/go/issues/25484
	return *(*string)(unsafe.Pointer(&strBuffer))
}

// ReadFloat64 reads 8 bytes and interprets them as a float64 IEEE 754 4 byte bit sequence.
func (f *LittleEndianBuffer) ReadFloat64() float64 {
	bits := f.ReadUint64()
	return math.Float64frombits(bits)
}

// ReadFloat32 reads 4 bytes and interprets them as a float32 IEEE 754 4 byte bit sequence.
func (f *LittleEndianBuffer) ReadFloat32() float32 {
	bits := f.ReadUint32()
	return math.Float32frombits(bits)
}

// WriteFloat32 writes a float32 IEEE 754 4 byte bit sequence.
func (f *LittleEndianBuffer) WriteFloat32(v float32) {
	bits := math.Float32bits(v)
	f.WriteUint32(bits)
}

// WriteFloat64 writes a float64 IEEE 754 8 byte bit sequence.
func (f *LittleEndianBuffer) WriteFloat64(v float64) {
	bits := math.Float64bits(v)
	f.WriteUint64(bits)
}

// WriteType writes the type as uint8
func (f *LittleEndianBuffer) WriteType(typ Type) {
	f.WriteUint8(uint8(typ))
}

func (f *LittleEndianBuffer) ReadType() Type {
	return Type(f.ReadUint8())
}

var drainJumpTable = [29]int{
	0, // undefined
	1, // TUint8      Type = 1
	2, // TUint16     Type = 2
	3, // TUint24     Type = 3
	4, // TUint32     Type = 4
	5, // TUint40     Type = 5
	6, // TUint48     Type = 6
	7, // TUint56     Type = 7
	8, // TUint64     Type = 8

	1, // TInt8       Type = 9
	2, // TInt16      Type = 10
	3, // TInt24      Type = 11
	4, // TInt32      Type = 12
	5, // TInt40      Type = 13
	6, // TInt48      Type = 14
	7, // TInt56      Type = 15
	8, // TInt64      Type = 16

	0, // TBlob8      Type = 17
	0, // TBlob16     Type = 18
	0, // TBlob24     Type = 19
	0, // TBlob32     Type = 20
	0, // TString8    Type = 21
	0, // TString16   Type = 22
	0, // TString24   Type = 23
	0, // TString32   Type = 24

	4, // TFloat32    Type = 25
	8, // TFloat64    Type = 26

	8,  // TComplex64  Type = 27
	16, // TComplex128 Type = 28

}

// DrainFast uses an inlineable jump table for fixed types and returns -1 for unsupported types. In that case, you
// have to fallback into the slow Drain. See also https://github.com/golang/go/issues/17566
func (f *LittleEndianBuffer) DrainFast(t Type) int {
	x := drainJumpTable[t]
	if x != 0 {
		f.Pos += x
		return x
	}

	return -1
}

// Drain moves the buffer position the right amount of bytes without actually parsing it
func (f *LittleEndianBuffer) Drain(t Type) int {
	oldPos := f.Pos
	switch t {
	case TInt8:
		fallthrough
	case TUint8:
		f.Pos++
	case TInt16:
		fallthrough
	case TUint16:
		f.Pos += 2
	case TInt24:
		fallthrough
	case TUint24:
		f.Pos += 3
	case TInt32:
		fallthrough
	case TUint32:
		f.Pos += 4
	case TInt64:
		fallthrough
	case TUint64:
		f.Pos += 8
	case TString8:
		fallthrough
	case TBlob8:
		vLen := int(f.ReadUint8())
		f.Pos += vLen
	case TString16:
		fallthrough
	case TBlob16:
		vLen := int(f.ReadUint16())
		f.Pos += vLen
	case TString24:
		fallthrough
	case TBlob24:
		vLen := int(f.ReadUint24())
		f.Pos += vLen
	case TString32:
		fallthrough
	case TBlob32:
		vLen := int(f.ReadUint32())
		f.Pos += vLen
	case TFloat32:
		f.Pos += 4
	case TFloat64:
		f.Pos += 8
	default:
		panic("not implemented " + strconv.Itoa(int(t)))
	}
	return f.Pos - oldPos
}
