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

func (f *LittleEndianBuffer) ReadUint16() uint16 {
	b := f.Bytes[f.Pos:]
	f.Pos += 2
	_ = b[1] // bounds check hint to compiler; see golang.org/issue/14808
	return uint16(b[0]) | uint16(b[1])<<8
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
	return *(*string)(unsafe.Pointer(&vLen))
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
	return *(*string)(unsafe.Pointer(&vLen))
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
	return *(*string)(unsafe.Pointer(&vLen))
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
	return *(*string)(unsafe.Pointer(&vLen))
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

// WriteFloat inspects the value and chooses automatically between int 1/2/3/4/5/6/7/8 byte signed or signed
// integers or float32/float64. The concrete value is prefixed with a type, so the written length is 2-9 byte.
// Floats with a fraction upto 1/1000 are encoded as float32.
func (f *LittleEndianBuffer) WriteFloat(v float64) {
	const epsilon = 1e-9

	// looks like an int?
	if _, frac := math.Modf(math.Abs(v)); frac < epsilon || frac > 1.0-epsilon {
		f.WriteInt(int64(v))
		return
	}

	// looks like it fits into float32?
	tmp := v * 1000
	if _, frac := math.Modf(math.Abs(tmp)); frac < epsilon || frac > 1.0-epsilon {
		f.WriteType(TFloat32)
		f.WriteFloat32(float32(v))
		return
	}

	f.WriteType(TFloat64)
	f.WriteFloat64(v)
}

// ReadFloat reads any number into a float
func (f *LittleEndianBuffer) ReadFloat() float64 {
	typ := f.ReadType()
	switch typ {
	case TUint8:
		return float64(f.ReadUint8())
	case TInt8:
		return float64(int8(f.ReadUint8()))

	case TUint16:
		return float64(f.ReadUint16())
	case TInt16:
		return float64(int16(f.ReadUint16()))

	case TUint24:
		return float64(f.ReadUint24())
	case TInt24:
		return float64(int32(f.ReadUint24()))

	case TUint32:
		return float64(f.ReadUint32())
	case TInt32:
		return float64(int32(f.ReadUint32()))

	case TUint40:
		return float64(f.ReadUint40())
	case TInt40:
		return float64(int64(f.ReadUint40()))

	case TUint48:
		return float64(f.ReadUint48())
	case TInt48:
		return float64(int64(f.ReadUint48()))

	case TUint56:
		return float64(f.ReadUint56())
	case TInt56:
		return float64(int64(f.ReadUint56()))

	case TUint64:
		return float64(f.ReadUint64())
	case TInt64:
		return float64(int64(f.ReadUint64()))

	case TFloat32:
		return float64(f.ReadFloat32())
	case TFloat64:
		return f.ReadFloat64()
	default:
		panic("unsupported type " + typ.String())

	}
}

// ReadInt reads any number into an integer
func (f *LittleEndianBuffer) ReadInt() int64 {
	typ := f.ReadType()
	switch typ {
	case TUint8:
		return int64(f.ReadUint8())
	case TInt8:
		return int64(int8(f.ReadUint8()))

	case TUint16:
		return int64(f.ReadUint16())
	case TInt16:
		return int64(int16(f.ReadUint16()))

	case TUint24:
		return int64(f.ReadUint24())
	case TInt24:
		return int64(int32(f.ReadUint24()))

	case TUint32:
		return int64(f.ReadUint32())
	case TInt32:
		return int64(int32(f.ReadUint32()))

	case TUint40:
		return int64(f.ReadUint40())
	case TInt40:
		return int64(int64(f.ReadUint40()))

	case TUint48:
		return int64(f.ReadUint48())
	case TInt48:
		return int64(int64(f.ReadUint48()))

	case TUint56:
		return int64(f.ReadUint56())
	case TInt56:
		return int64(int64(f.ReadUint56()))

	case TUint64:
		return int64(f.ReadUint64())
	case TInt64:
		return int64(int64(f.ReadUint64()))

	case TFloat32:
		return int64(f.ReadFloat32())
	case TFloat64:
		return int64(f.ReadFloat64())
	default:
		panic("unsupported type " + typ.String())

	}
}

// WriteInt inspects the value and chooses automatically between 1/2/3/4/5/6/7/8 byte representation
// of signed or unsigned values. The resulting size is 2-9 byte. This contains a lot of branches.
// The concrete value is prefixed with a type, so the written length is 2-9 byte.
func (f *LittleEndianBuffer) WriteInt(v int64) {
	// The structure increases the minimum amount of checks but the worst case amount of checks
	// is halfed.

	if v >= int64(MinInt8) && v <= int64(MaxUint8) {
		if v > int64(MaxInt8) {
			f.WriteType(TUint8)
		} else {
			f.WriteType(TInt8)
		}
		f.WriteUint8(uint8(v))
		return
	}

	if v >= int64(MinInt16) && v <= int64(MaxUint16) {
		if v > int64(MaxInt16) {
			f.WriteType(TUint16)
		} else {
			f.WriteType(TInt16)
		}
		f.WriteUint16(uint16(v))
		return
	}

	if v >= int64(MinInt24) && v <= int64(MaxUint24) {
		if v > int64(MaxInt24) {
			f.WriteType(TUint24)
		} else {
			f.WriteType(TInt24)
		}
		f.WriteUint24(uint32(v))
		return
	}

	if v >= int64(MinInt32) && v <= int64(MaxUint32) {
		if v > int64(MaxInt32) {
			f.WriteType(TUint32)
		} else {
			f.WriteType(TInt32)
		}
		f.WriteUint32(uint32(v))
		return
	}

	if v >= int64(MinInt40) && v <= int64(MaxUint40) {
		if v > int64(MaxInt40) {
			f.WriteType(TUint40)
		} else {
			f.WriteType(TInt40)
		}
		f.WriteUint40(uint64(v))
		return
	}

	if v >= int64(MinInt48) && v <= int64(MaxUint48) {
		if v > int64(MaxInt48) {
			f.WriteType(TUint48)
		} else {
			f.WriteType(TInt48)
		}
		f.WriteUint48(uint64(v))
		return
	}

	if v >= int64(MinInt56) && v <= int64(MaxUint56) {
		if v > int64(MaxInt56) {
			f.WriteType(TUint56)
		} else {
			f.WriteType(TInt56)
		}
		f.WriteUint56(uint64(v))
		return
	}

	f.WriteType(TInt64)
	f.WriteUint64(uint64(v))
}

// WriteString determines how many bytes the string has and chooses between an 1,2,3 or 4 byte length prefix. It
// is prefixed with a type, indicating the max size and followed by the actual length prefix and string bytes.
func (f *LittleEndianBuffer) WriteString(str string) {
	if len(str) <= int(MaxUint8) {
		f.WriteType(TString8)
		f.WriteString8(str)
		return
	}

	if len(str) <= int(MaxUint16) {
		f.WriteType(TString16)
		f.WriteString16(str)
		return
	}

	if len(str) <= int(MaxUint24) {
		f.WriteType(TString24)
		f.WriteString24(str)
		return
	}

	f.WriteType(TString32)
	f.WriteString32(str)
	return
}

// ReadString reads a string8/16/24 or 32 string into the strBuffer and returns a mutable string from it.
func (f *LittleEndianBuffer) ReadString(strBuffer []byte) string {
	typ := f.ReadType()
	switch typ {
	case TString8:
		return f.ReadString8(strBuffer)
	case TString16:
		return f.ReadString16(strBuffer)
	case TString24:
		return f.ReadString24(strBuffer)
	case TString32:
		return f.ReadString32(strBuffer)
	default:
		panic("unsupported type " + typ.String())
	}
}

// WriteBlob determines how many bytes the buffer has and chooses between an 1,2,3 or 4 byte length prefix. It
// is prefixed with a blob type, indicating the max size and followed by the actual length prefix and string bytes.
func (f *LittleEndianBuffer) WriteBlob(b []byte) {
	if len(b) <= int(MaxUint8) {
		f.WriteType(TBlob8)
		f.WriteBlob8(b)
		return
	}

	if len(b) <= int(MaxUint16) {
		f.WriteType(TBlob16)
		f.WriteBlob16(b)
		return
	}

	if len(b) <= int(MaxUint24) {
		f.WriteType(TBlob24)
		f.WriteBlob24(b)
		return
	}

	f.WriteType(TBlob32)
	f.WriteBlob32(b)
	return
}

// ReadBlob reads a blob8/16/24 or 32 into the buffer.
func (f *LittleEndianBuffer) ReadBlob(b []byte) int {
	typ := f.ReadType()
	switch typ {
	case TBlob8:
		return f.ReadBlob8(b)
	case TBlob16:
		return f.ReadBlob16(b)
	case TBlob24:
		return f.ReadBlob24(b)
	case TBlob32:
		return f.ReadBlob32(b)
	default:
		panic("unsupported type " + typ.String())
	}
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
