package ioutil

// PrefixType defines how a storage class is encoded, using 1,2,3,4,8 oder a variable encoding.
type PrefixType int

const (
	// A Tiny storage class is 1 byte/8bit and max length is 255 bytes
	Tiny PrefixType = 1

	// A Small storage class is 2 byte/16bit and max length is 65.535 bytes (64kb)
	Small PrefixType = 2

	// A Medium storage class is 3 byte/24bit and max length is 16.777.215 bytes (16mb)
	Medium PrefixType = 3

	// A Long storage class is 4 byte/32bit and max length is 4.294.967.295 bytes (4gb)
	Long PrefixType = 4

	// A Large storage class is 8 byte/64bit and max length is 9.223.372.036.854.775.806 bytes (8.388.608tb)
	Large PrefixType = 8

	// A Var storage class uses a varint encoding from 1-9 byte (8.388.608tb)
	Var PrefixType = 0
)
