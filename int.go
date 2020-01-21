package ioutil

// UintSize is either 32 or 64
const UintSize = 32 << (^uint(0) >> 32 & 1)

const (
	// MaxUint is either 1<<31 - 1 or 1<<63 - 1
	MaxInt = 1<<(UintSize-1) - 1

	// MinInt is either -1 << 31 or -1 << 63
	MinInt = -MaxInt - 1

	// MaxUint is either 1<<32 - 1 or 1<<64 - 1
	MaxUint = 1<<UintSize - 1
)

func uint24BE(b []byte) uint32 {
	_ = b[2] // bounds check hint to compiler; see golang.org/issue/14808
	return uint32(b[2])<<0 | uint32(b[1])<<8 | uint32(b[0])<<16
}

func uint24LE(b []byte) uint32 {
	_ = b[2] // bounds check hint to compiler; see golang.org/issue/14808
	return uint32(b[0]) | uint32(b[1])<<8 | uint32(b[2])<<16
}

func uint40BE(b []byte) uint64 {
	_ = b[4] // bounds check hint to compiler; see golang.org/issue/14808
	return uint64(b[4])<<0 | uint64(b[3])<<8 | uint64(b[2])<<16 | uint64(b[1])<<24 | uint64(b[0])<<32
}

func uint40LE(b []byte) uint64 {
	_ = b[4] // bounds check hint to compiler; see golang.org/issue/14808
	return uint64(b[0]) | uint64(b[1])<<8 | uint64(b[2])<<16 | uint64(b[3])<<24 | uint64(b[4])<<32
}
