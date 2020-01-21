package ioutil

// IntSize defines how a storage class is encoded, using 1,2,3,4,8 oder a variable encoding.
type IntSize int

const (
	// An I8 storage class is 1 byte/8bit and max length is 255 bytes
	I8 IntSize = 1

	// An I16 storage class is 2 byte/16bit and max length is 65.535 bytes (64kb)
	I16 IntSize = 2

	// An I24 storage class is 3 byte/24bit and max length is 16.777.215 bytes (16mb)
	I24 IntSize = 3

	// An I32 storage class is 4 byte/32bit and max length is 4.294.967.295 bytes (4gb)
	I32 IntSize = 4

	// An I40 storage class is 5 byte/40bit and max length is 1.099.511.627.776 (1tb)
	I40 IntSize = 5

	// An I64 storage class is 8 byte/64bit and max length is 9.223.372.036.854.775.806 bytes (8.388.608tb)
	I64 IntSize = 8

	// An IVar storage class uses a varint encoding from 1-9 byte (8.388.608tb)
	IVar IntSize = 0
)
