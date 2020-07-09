package ioutil

import (
	"math"
)

// TypedLittleEndianBuffer is a light weight helper to modify bytes within a buffer in little endian format and
// each written type has a type prefix.
type TypedLittleEndianBuffer LittleEndianBuffer

// WriteFloat inspects the value and chooses automatically between int 1/2/3/4/5/6/7/8 byte signed or signed
// integers or float32/float64. The concrete value is prefixed with a type, so the written length is 2-9 byte.
// Floats with a fraction upto 1/1000 and smaller than 16777215 are encoded as float32.
func (t *TypedLittleEndianBuffer) WriteFloat(v float64) {
	f := (*LittleEndianBuffer)(t)
	const epsilon = 1e-9

	// looks like an int?
	if _, frac := math.Modf(math.Abs(v)); frac < epsilon || frac > 1.0-epsilon {
		t.WriteInt(int64(v))
		return
	}

	// looks like it fits into float32?
	tmp := v * 1000
	if _, frac := math.Modf(math.Abs(tmp)); frac < epsilon || frac > 1.0-epsilon && v <= 16777215 {
		f.WriteType(TFloat32)
		f.WriteFloat32(float32(v))
		return
	}

	f.WriteType(TFloat64)
	f.WriteFloat64(v)
}

// ReadFloat reads any number into a float
func (t *TypedLittleEndianBuffer) ReadFloat() float64 {
	f := (*LittleEndianBuffer)(t)
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
func (t *TypedLittleEndianBuffer) ReadInt() int64 {
	f := (*LittleEndianBuffer)(t)
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
func (t *TypedLittleEndianBuffer) WriteInt(v int64) {
	f := (*LittleEndianBuffer)(t)

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
func (t *TypedLittleEndianBuffer) WriteString(str string) {
	f := (*LittleEndianBuffer)(t)

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
func (t *TypedLittleEndianBuffer) ReadString(strBuffer []byte) string {
	if strBuffer == nil {
		//ups, need to work around our mutable owned string approach
		// otherwise we will get weired sigsegv somewhere later
		tmp := make([]byte, 1024*64) // 64k
		strBuffer = tmp
	}
	f := (*LittleEndianBuffer)(t)

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
func (t *TypedLittleEndianBuffer) WriteBlob(b []byte) {
	f := (*LittleEndianBuffer)(t)

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
func (t *TypedLittleEndianBuffer) ReadBlob(b []byte) int {
	f := (*LittleEndianBuffer)(t)

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

func (t *TypedLittleEndianBuffer) WriteUint8(v uint8) {
	f := (*LittleEndianBuffer)(t)
	f.WriteType(TUint8)
	f.WriteUint8(v)
}

func (t *TypedLittleEndianBuffer) WriteInt8(v int8) {
	f := (*LittleEndianBuffer)(t)
	f.WriteType(TInt8)
	f.WriteUint8(uint8(v))
}

func (t *TypedLittleEndianBuffer) WriteUint16(v uint16) {
	f := (*LittleEndianBuffer)(t)
	f.WriteType(TUint16)
	f.WriteUint16(v)
}

func (t *TypedLittleEndianBuffer) WriteInt16(v int16) {
	f := (*LittleEndianBuffer)(t)
	f.WriteType(TInt16)
	f.WriteUint16(uint16(v))
}

func (t *TypedLittleEndianBuffer) WriteUint24(v uint32) {
	f := (*LittleEndianBuffer)(t)
	f.WriteType(TUint24)
	f.WriteUint24(v)
}

func (t *TypedLittleEndianBuffer) WriteInt24(v int32) {
	f := (*LittleEndianBuffer)(t)
	f.WriteType(TInt24)
	f.WriteUint24(uint32(v))
}

func (t *TypedLittleEndianBuffer) WriteUint32(v uint32) {
	f := (*LittleEndianBuffer)(t)
	f.WriteType(TUint32)
	f.WriteUint32(v)
}

func (t *TypedLittleEndianBuffer) WriteInt32(v int32) {
	f := (*LittleEndianBuffer)(t)
	f.WriteType(TInt32)
	f.WriteUint32(uint32(v))
}

func (t *TypedLittleEndianBuffer) WriteUint40(v uint64) {
	f := (*LittleEndianBuffer)(t)
	f.WriteType(TUint40)
	f.WriteUint40(v)
}

func (t *TypedLittleEndianBuffer) WriteInt40(v int64) {
	f := (*LittleEndianBuffer)(t)
	f.WriteType(TInt40)
	f.WriteUint40(uint64(v))
}

func (t *TypedLittleEndianBuffer) WriteUint48(v uint64) {
	f := (*LittleEndianBuffer)(t)
	f.WriteType(TUint48)
	f.WriteUint48(v)
}

func (t *TypedLittleEndianBuffer) WriteInt48(v int64) {
	f := (*LittleEndianBuffer)(t)
	f.WriteType(TInt48)
	f.WriteUint48(uint64(v))
}

func (t *TypedLittleEndianBuffer) WriteUint56(v uint64) {
	f := (*LittleEndianBuffer)(t)
	f.WriteType(TUint56)
	f.WriteUint56(v)
}

func (t *TypedLittleEndianBuffer) WriteInt56(v int64) {
	f := (*LittleEndianBuffer)(t)
	f.WriteType(TInt56)
	f.WriteUint56(uint64(v))
}

func (t *TypedLittleEndianBuffer) WriteUint64(v uint64) {
	f := (*LittleEndianBuffer)(t)
	f.WriteType(TUint64)
	f.WriteUint64(v)
}

func (t *TypedLittleEndianBuffer) WriteInt64(v int64) {
	f := (*LittleEndianBuffer)(t)
	f.WriteType(TInt64)
	f.WriteUint64(uint64(v))
}

func (t *TypedLittleEndianBuffer) WriteFloat32(v float32) {
	f := (*LittleEndianBuffer)(t)
	f.WriteType(TFloat32)
	f.WriteFloat32(v)
}

func (t *TypedLittleEndianBuffer) WriteFloat64(v float64) {
	f := (*LittleEndianBuffer)(t)
	f.WriteType(TFloat64)
	f.WriteFloat64(v)
}

func (t *TypedLittleEndianBuffer) WriteBlob8(v []byte) {
	f := (*LittleEndianBuffer)(t)
	f.WriteType(TBlob8)
	f.WriteBlob8(v)
}

func (t *TypedLittleEndianBuffer) WriteBlob16(v []byte) {
	f := (*LittleEndianBuffer)(t)
	f.WriteType(TBlob16)
	f.WriteBlob16(v)
}

func (t *TypedLittleEndianBuffer) WriteBlob24(v []byte) {
	f := (*LittleEndianBuffer)(t)
	f.WriteType(TBlob24)
	f.WriteBlob24(v)
}

func (t *TypedLittleEndianBuffer) WriteBlob32(v []byte) {
	f := (*LittleEndianBuffer)(t)
	f.WriteType(TBlob32)
	f.WriteBlob32(v)
}

func (t *TypedLittleEndianBuffer) ReadUint8() uint8 {
	f := (*LittleEndianBuffer)(t)
	t.assertType(TUint8)
	return f.ReadUint8()
}

func (t *TypedLittleEndianBuffer) ReadInt8() int8 {
	f := (*LittleEndianBuffer)(t)
	t.assertType(TInt8)
	return int8(f.ReadUint8())
}

func (t *TypedLittleEndianBuffer) ReadUint16() uint16 {
	f := (*LittleEndianBuffer)(t)
	t.assertType(TUint16)
	return f.ReadUint16()
}

func (t *TypedLittleEndianBuffer) ReadInt16() int16 {
	f := (*LittleEndianBuffer)(t)
	t.assertType(TInt16)
	return int16(f.ReadUint16())
}

func (t *TypedLittleEndianBuffer) ReadUint24() uint32 {
	f := (*LittleEndianBuffer)(t)
	t.assertType(TUint24)
	return f.ReadUint24()
}

func (t *TypedLittleEndianBuffer) ReadInt24() int32 {
	f := (*LittleEndianBuffer)(t)
	t.assertType(TInt24)
	return int32(f.ReadUint24())
}

func (t *TypedLittleEndianBuffer) ReadUint32() uint32 {
	f := (*LittleEndianBuffer)(t)
	t.assertType(TUint32)
	return f.ReadUint32()
}

func (t *TypedLittleEndianBuffer) ReadInt32() int32 {
	f := (*LittleEndianBuffer)(t)
	t.assertType(TInt32)
	return int32(f.ReadUint32())
}

func (t *TypedLittleEndianBuffer) ReadUint64() uint64 {
	f := (*LittleEndianBuffer)(t)
	t.assertType(TUint64)
	return f.ReadUint64()
}

func (t *TypedLittleEndianBuffer) ReadInt64() int64 {
	f := (*LittleEndianBuffer)(t)
	t.assertType(TInt64)
	return int64(f.ReadUint64())
}

func (t *TypedLittleEndianBuffer) ReadBlob8(dst []byte) int {
	f := (*LittleEndianBuffer)(t)
	t.assertType(TBlob8)
	return f.ReadBlob8(dst)
}

func (t *TypedLittleEndianBuffer) ReadBlob16(dst []byte) int {
	f := (*LittleEndianBuffer)(t)
	t.assertType(TBlob16)
	return f.ReadBlob16(dst)
}

func (t *TypedLittleEndianBuffer) ReadBlob24(dst []byte) int {
	f := (*LittleEndianBuffer)(t)
	t.assertType(TBlob24)
	return f.ReadBlob24(dst)
}

func (t *TypedLittleEndianBuffer) ReadBlob32(dst []byte) int {
	f := (*LittleEndianBuffer)(t)
	t.assertType(TBlob32)
	return f.ReadBlob32(dst)
}

func (t *TypedLittleEndianBuffer) ReadFloat32() float32 {
	f := (*LittleEndianBuffer)(t)
	t.assertType(TFloat32)
	return f.ReadFloat32()
}

func (t *TypedLittleEndianBuffer) ReadFloat64() float64 {
	f := (*LittleEndianBuffer)(t)
	t.assertType(TFloat64)
	return f.ReadFloat64()
}

func (t *TypedLittleEndianBuffer) assertType(kind Type) {
	if debug{
		f := (*LittleEndianBuffer)(t)
		x := f.ReadType()
		if x != kind {
			panic("expected " + kind.String() + " but got " + x.String()) // this is not inlineable
		}
	}
}
