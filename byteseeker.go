package ioutil

import (
	"io"
	"math"
)

// ByteSeeker is an implementation for an in-memory io.WriteSeeker and io.ReadSeeker
type ByteSeeker struct {
	buf []byte
	pos int
}

// Read returns EOF if no bytes can be read anymore.
func (b *ByteSeeker) Read(p []byte) (n int, err error) {
	if b.pos == len(b.buf)-1 && len(p) > 0 {
		return 0, io.EOF
	}

	var atMost int

	if b.pos+len(p) > len(b.buf) {
		atMost = len(b.buf) - b.pos
	} else {
		atMost = len(p)
	}

	copy(p[:atMost], b.buf[b.pos:b.pos+atMost])

	return atMost, nil
}

func (b *ByteSeeker) Write(p []byte) (n int, err error) {
	size := b.pos + len(p)
	b.ensureBuffer(size)
	copy(b.buf[b.pos:], p)
	b.pos += len(p)

	return len(p), nil
}

// ensureBuffer ensures the required size. New capacity either doubles or uses the exact size, whatever is larger.
// This will result in a nice adaptive behavior, where an initial write buffers
// The exact size and does not cause any unused over provisioning
func (b *ByteSeeker) ensureBuffer(size int) {
	if size > cap(b.buf) {
		newCap := int(math.Max(float64(size), float64(len(b.buf))))
		tmp := make([]byte, len(b.buf), newCap)
		copy(tmp, b.buf)
		b.buf = tmp
	}

	if size > len(b.buf) {
		b.buf = b.buf[:size]
	}
}

// Seek returns EOF if seeking before the beginning and enlarges the buffer, if required, seeks and allocates
// beyond the buffer
func (b *ByteSeeker) Seek(offset int64, whence int) (int64, error) {
	newPos, offs := 0, int(offset)

	switch whence {
	case io.SeekStart:
		newPos = offs
	case io.SeekCurrent:
		newPos = b.pos + offs
	case io.SeekEnd:
		newPos = len(b.buf) + offs
	}

	if newPos < 0 {
		b.pos = 0
		return 0, io.EOF
	}

	b.ensureBuffer(newPos)
	b.pos = newPos

	return int64(b.pos), nil
}

// Close is a no-op
func (b *ByteSeeker) Close() error {
	return nil
}

// Bytes returns the backing buffer.
func (b *ByteSeeker) Bytes() []byte {
	return b.buf
}

// Pos returns the current position
func (b *ByteSeeker) Pos() int {
	return b.pos
}
