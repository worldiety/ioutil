package ioutil

import (
	"fmt"
	"testing"
)

func Test_Uint16(t *testing.T) {
	tmp := make([]byte, 2)

	for _, bo := range []ByteOrder{LittleEndian, BigEndian} {
		var i uint16
		for i = 0; i < MaxUint16; i++ {
			bo.PutUint16(tmp, i)
			i2 := bo.Uint16(tmp)

			if i != i2 {
				t.Fatalf("expected %d but got %d", i, i2)
			}
		}
	}
}

func Test_Uint24(t *testing.T) {
	tmp := make([]byte, 3)

	for _, bo := range []ByteOrder{LittleEndian, BigEndian} {
		var i uint32
		for i = 0; i < MaxUint24; i++ {
			bo.PutUint24(tmp, i)
			i2 := bo.Uint24(tmp)

			if i != i2 {
				t.Fatalf("expected %d but got %d", i, i2)
			}
		}
	}
}

func Test_Uint32(t *testing.T) {
	tmp := make([]byte, 4)

	for _, bo := range []ByteOrder{LittleEndian, BigEndian} {
		for _, i := range []uint32{0x0, uint32(MaxUint8), uint32(MaxUint16), uint32(MaxUint24), uint32(MaxUint32), 42} {
			bo.PutUint32(tmp, i)
			i2 := bo.Uint32(tmp)

			if i != i2 {
				t.Fatalf("expected %d but got %d", i, i2)
			}
		}
	}
}

func Test_Uint40(t *testing.T) {
	tmp := make([]byte, 5)

	for _, bo := range []ByteOrder{LittleEndian, BigEndian} {
		for _, i := range []uint64{0x0, uint64(MaxUint8), uint64(MaxUint16), uint64(MaxUint24), uint64(MaxUint32), uint64(MaxInt40), 42} {
			bo.PutUint40(tmp, i)
			i2 := bo.Uint40(tmp)

			if i != i2 {
				t.Fatalf("expected %d but got %d", i, i2)
			}
		}
	}
}

func Test_Uint48(t *testing.T) {
	tmp := make([]byte, 6)

	for _, bo := range []ByteOrder{LittleEndian, BigEndian} {
		for _, i := range []uint64{0x0, uint64(MaxUint8), uint64(MaxUint16), uint64(MaxUint24), uint64(MaxUint32), uint64(MaxInt40), uint64(MaxInt48), 42} {
			bo.PutUint48(tmp, i)
			i2 := bo.Uint48(tmp)

			if i != i2 {
				t.Fatalf("expected %d but got %d", i, i2)
			}
		}
	}
}

func Test_Uint56(t *testing.T) {
	tmp := make([]byte, 7)

	for _, bo := range []ByteOrder{LittleEndian, BigEndian} {
		for _, i := range []uint64{0x0, uint64(MaxUint8), uint64(MaxUint16), uint64(MaxUint24), uint64(MaxUint32), uint64(MaxInt40), uint64(MaxInt48), uint64(MaxUint56), 42} {
			bo.PutUint56(tmp, i)
			i2 := bo.Uint56(tmp)

			if i != i2 {
				t.Fatalf("expected %d but got %d", i, i2)
			}
		}
	}
}

func Test_Uint64(t *testing.T) {
	tmp := make([]byte, 8)

	for _, bo := range []ByteOrder{LittleEndian, BigEndian} {
		for _, i := range []uint64{0x0,
			uint64(MaxUint8), uint64(MaxUint16), uint64(MaxUint24), uint64(MaxUint32),
			uint64(MaxInt40), uint64(MaxInt48), uint64(MaxUint56), uint64(MaxUint64), 42} {
			bo.PutUint64(tmp, i)
			i2 := bo.Uint64(tmp)

			if i != i2 {
				t.Fatalf("expected %d but got %d", i, i2)
			}
		}
	}
}

func Test_PrintMaxes(t *testing.T) {
	values := []interface{}{MaxInt8, MinInt8,
		MaxInt16, MinInt16,
		MaxInt24, MinInt24,
		MaxInt32, MinInt32,
		MaxInt40, MinInt40,
		MaxInt48, MinInt48,
		MaxInt56, MinInt56,
		MaxInt64, MinInt64,
		MaxUint8, MaxUint16, MaxUint24, MaxUint32, MaxUint40, MaxUint48, MaxUint56, uint64(MaxUint64)}
	names := []string{"MaxInt8", "MinInt8",
		"MaxInt16", "MinInt16",
		"MaxInt24", "MinInt24",
		"MaxInt32", "MinInt32",
		"MaxInt40", "MinInt40",
		"MaxInt48", "MinInt48",
		"MaxInt56", "MinInt56",
		"MaxInt64", "MinInt64",
		"MaxUint8", "MaxUint16", "MaxUint24", "MaxUint32", "MaxUint40", "MaxUint48", "MaxUint56", "MaxUint64"}

	for i := range values {
		fmt.Printf("// %s is %d\n", names[i], values[i])
	}
}
