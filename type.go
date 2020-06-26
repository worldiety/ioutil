package ioutil

import "strconv"

// Type enumerates a bunch of data types.
type Type byte

const (
	TUint8      Type = 1
	TUint16     Type = 2
	TUint24     Type = 3
	TUint32     Type = 4
	TUint40     Type = 5
	TUint48     Type = 6
	TUint56     Type = 7
	TUint64     Type = 8
	TInt8       Type = 9
	TInt16      Type = 10
	TInt24      Type = 11
	TInt32      Type = 12
	TInt40      Type = 13
	TInt48      Type = 14
	TInt56      Type = 15
	TInt64      Type = 16
	TBlob8      Type = 17
	TBlob16     Type = 18
	TBlob24     Type = 19
	TBlob32     Type = 20
	TString8    Type = 21
	TString16   Type = 22
	TString24   Type = 23
	TString32   Type = 24
	TFloat32    Type = 25
	TFloat64    Type = 26
	TComplex64  Type = 27
	TComplex128 Type = 28

	minTValid = TUint8
	maxTValid = TFloat64
)

func (d Type) IsValid() bool {
	return d >= minTValid && d <= maxTValid
}

func (d Type) IsNumber() bool {
	switch d {
	case TUint8:
		fallthrough
	case TUint16:
		fallthrough
	case TUint24:
		fallthrough
	case TUint32:
		fallthrough
	case TUint40:
		fallthrough
	case TUint48:
		fallthrough
	case TUint56:
		fallthrough
	case TUint64:
		fallthrough
	case TInt8:
		fallthrough
	case TInt16:
		fallthrough
	case TInt24:
		fallthrough
	case TInt32:
		fallthrough
	case TInt40:
		fallthrough
	case TInt48:
		fallthrough
	case TInt56:
		fallthrough
	case TInt64:
		fallthrough
	case TFloat32:
		fallthrough
	case TFloat64:
		return true
	default:
		return false
	}
}

func (d Type) String() string {
	switch d {
	case TUint8:
		return "uint8"
	case TUint16:
		return "uint16"
	case TUint24:
		return "uint24"
	case TUint32:
		return "uint32"
	case TUint40:
		return "uint40"
	case TUint48:
		return "uint48"
	case TUint56:
		return "uint56"
	case TUint64:
		return "uint64"
	case TInt8:
		return "int8"
	case TInt16:
		return "int16"
	case TInt24:
		return "int24"
	case TInt32:
		return "int32"
	case TInt40:
		return "int40"
	case TInt48:
		return "int48"
	case TInt56:
		return "int56"
	case TInt64:
		return "int64"
	case TBlob8:
		return "blob8"
	case TBlob16:
		return "blob16"
	case TBlob24:
		return "blob24"
	case TBlob32:
		return "blob32"
	case TString8:
		return "string8"
	case TString16:
		return "string16"
	case TString24:
		return "string24"
	case TString32:
		return "string32"
	case TFloat32:
		return "float32"
	case TFloat64:
		return "float64"
	case TComplex64:
		return "complex64"
	case TComplex128:
		return "complex128"
	default:
		return "unspecified " + strconv.Itoa(int(d))
	}
}
