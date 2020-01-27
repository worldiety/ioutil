/*
 * Copyright 2020 Torben Schinke
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package ioutil

import "io"

// A ByteReader implements the io.ByteReader and delegates to an io.Reader. It is not thread safe, because
// it reuses an internal buffer. The reason is, that for calling a virtual method (like Read), Go must allocate
// on the heap, because the array may escape which cannot be proven at compile time in general.
type ByteReader struct {
	buf    [1]byte
	reader io.Reader
}

// ReadByte reads a byte or an error.
func (b *ByteReader) ReadByte() (byte, error) {
	n, err := b.reader.Read(b.buf[:])
	if err != nil {
		return 0, err
	}

	if n != 1 { //nolint:gomnd
		return 0, io.EOF
	}

	return b.buf[0], nil
}

// NewByteReader wraps another io.Reader and allows
func NewByteReader(r io.Reader) *ByteReader {
	return &ByteReader{reader: r}
}
