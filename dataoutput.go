package ioutil

type DataOutput interface {
	WriteUInt32LE(val uint32)
	WriteUInt32BE(val uint32)
	WriteByte(val byte)
	Write(buf []byte) (int, error)



	Error() error
}
