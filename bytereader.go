package ioutil

import "io"

type ByteReader struct {
	buf    [1]byte
	reader io.Reader
}

func (b ByteReader) ReadByte() (byte, error) {
	n, err := b.reader.Read(b.buf[:])
	if err != nil {
		return 0, err
	}
	if n != 1 {
		return 0, io.EOF
	}
	return b.buf[0], nil
}

// NewByteReader wraps another io.Reader and allows
func NewByteReader(r io.Reader) ByteReader {
	return ByteReader{reader: r}
}
