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

import (
	"fmt"
	"hash"
	"io"
)

// A HashReader calculates for every transferred byte the hash until Sum() is called.
// This is useful to create a middleware component which just calculates a hash of a processed byte stream.
type HashReader struct {
	hasher hash.Hash
	reader io.Reader
	count  uint64
}

// NewHashReader creates a new instance. The given hash instance is unchanged until the first read.
func NewHashReader(h hash.Hash, reader io.Reader) *HashReader {
	hr := &HashReader{hasher: h, reader: reader}
	return hr
}

func (h *HashReader) Read(p []byte) (n int, err error) {
	n, err = h.reader.Read(p)
	n2, err2 := h.hasher.Write(p[0:n])

	if err != nil && err2 != nil {
		return n, fmt.Errorf("failed to hash: %w", fmt.Errorf("failed to read: %w", err))
	}

	if err != nil {
		return n, err
	}

	if err2 != nil {
		return n2, err2
	}

	if n != n2 {
		return n, fmt.Errorf("unable to hash the buffer properly")
	}

	return n, nil
}

// Sum returns the resulting slice.
// It does not change the underlying hash state.
func (h *HashReader) Sum() []byte {
	return h.hasher.Sum(nil)
}

// Hash returns the wrapped hasher
func (h *HashReader) Hash() hash.Hash {
	return h.hasher
}

// Count returns the total amount of read bytes so far.
func (h *HashReader) Count() uint64 {
	return h.count
}

// Reset sets the internal byte count to 0 and resets the hash
func (h *HashReader) Reset() {
	h.count = 0
	h.hasher.Reset()
}
