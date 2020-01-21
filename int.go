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
