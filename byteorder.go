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

// A ByteOrder specifies how to convert byte sequences into
// 16-, 24-, 32-, 40-, 48-, 56- or 64-bit unsigned integers.
// It is compatible with the standard library ByteOrder but contains
// more integer types.
type ByteOrder interface {
	// Uint16 reads 2 bytes
	Uint16([]byte) uint16
	// Uint24 reads 3 bytes
	Uint24([]byte) uint32
	// Uint32 reads 4 bytes
	Uint32([]byte) uint32
	// Uint40 reads 5 bytes
	Uint40([]byte) uint64
	// Uint48 reads 6 bytes
	Uint48([]byte) uint64
	// Uint56 reads 7 bytes
	Uint56([]byte) uint64
	// Uint64 reads 8 bytes
	Uint64([]byte) uint64
	// PutUint16 writes 2 bytes
	PutUint16([]byte, uint16)
	// PutUint24 writes 3 bytes
	PutUint24([]byte, uint32)
	// PutUint32 writes 4 bytes
	PutUint32([]byte, uint32)
	// PutUint40 writes 5 bytes
	PutUint40([]byte, uint64)
	// PutUint48 writes 6 bytes
	PutUint48([]byte, uint64)
	// PutUint56 writes 7 bytes
	PutUint56([]byte, uint64)
	// PutUint64 writes 8 bytes
	PutUint64([]byte, uint64)
	// String returns the endianness name
	String() string
}

// LittleEndian instance of ByteOrder.
var LittleEndian littleEndian //nolint:gochecknoglobals

// BigEndian instance of ByteOrder.
var BigEndian bigEndian //nolint:gochecknoglobals

type littleEndian struct {
}

func (littleEndian) Uint16(b []byte) uint16 {
	_ = b[1] // bounds check hint to compiler; see golang.org/issue/14808
	return uint16(b[0]) | uint16(b[1])<<8
}

func (littleEndian) Uint24(b []byte) uint32 {
	_ = b[2] // bounds check hint to compiler; see golang.org/issue/14808
	return uint32(b[0]) | uint32(b[1])<<8 | uint32(b[2])<<16
}

func (littleEndian) Uint32(b []byte) uint32 {
	_ = b[3] // bounds check hint to compiler; see golang.org/issue/14808
	return uint32(b[0]) | uint32(b[1])<<8 | uint32(b[2])<<16 | uint32(b[3])<<24
}

func (littleEndian) Uint40(b []byte) uint64 {
	_ = b[4] // bounds check hint to compiler; see golang.org/issue/14808
	return uint64(b[0]) | uint64(b[1])<<8 | uint64(b[2])<<16 | uint64(b[3])<<24 |
		uint64(b[4])<<32
}

func (littleEndian) Uint48(b []byte) uint64 {
	_ = b[5] // bounds check hint to compiler; see golang.org/issue/14808
	return uint64(b[0]) | uint64(b[1])<<8 | uint64(b[2])<<16 | uint64(b[3])<<24 |
		uint64(b[4])<<32 | uint64(b[5])<<40
}

func (littleEndian) Uint56(b []byte) uint64 {
	_ = b[6] // bounds check hint to compiler; see golang.org/issue/14808
	return uint64(b[0]) | uint64(b[1])<<8 | uint64(b[2])<<16 | uint64(b[3])<<24 |
		uint64(b[4])<<32 | uint64(b[5])<<40 | uint64(b[6])<<48
}

func (littleEndian) Uint64(b []byte) uint64 {
	_ = b[7] // bounds check hint to compiler; see golang.org/issue/14808
	return uint64(b[0]) | uint64(b[1])<<8 | uint64(b[2])<<16 | uint64(b[3])<<24 |
		uint64(b[4])<<32 | uint64(b[5])<<40 | uint64(b[6])<<48 | uint64(b[7])<<56
}

func (littleEndian) PutUint16(b []byte, v uint16) {
	_ = b[1] // early bounds check to guarantee safety of writes below
	b[0] = byte(v)
	b[1] = byte(v >> 8) //nolint:gomnd
}

func (littleEndian) PutUint24(b []byte, v uint32) {
	_ = b[2] // early bounds check to guarantee safety of writes below
	b[0] = byte(v)
	b[1] = byte(v >> 8)  //nolint:gomnd
	b[2] = byte(v >> 16) //nolint:gomnd
}

func (littleEndian) PutUint32(b []byte, v uint32) {
	_ = b[3] // early bounds check to guarantee safety of writes below
	b[0] = byte(v)
	b[1] = byte(v >> 8)  //nolint:gomnd
	b[2] = byte(v >> 16) //nolint:gomnd
	b[3] = byte(v >> 24) //nolint:gomnd
}

func (littleEndian) PutUint40(b []byte, v uint64) {
	_ = b[4] // early bounds check to guarantee safety of writes below
	b[0] = byte(v)
	b[1] = byte(v >> 8)  //nolint:gomnd
	b[2] = byte(v >> 16) //nolint:gomnd
	b[3] = byte(v >> 24) //nolint:gomnd
	b[4] = byte(v >> 32) //nolint:gomnd
}

func (littleEndian) PutUint48(b []byte, v uint64) {
	_ = b[5] // early bounds check to guarantee safety of writes below
	b[0] = byte(v)
	b[1] = byte(v >> 8)  //nolint:gomnd
	b[2] = byte(v >> 16) //nolint:gomnd
	b[3] = byte(v >> 24) //nolint:gomnd
	b[4] = byte(v >> 32) //nolint:gomnd
	b[5] = byte(v >> 40) //nolint:gomnd
}

func (littleEndian) PutUint56(b []byte, v uint64) {
	_ = b[6] // early bounds check to guarantee safety of writes below
	b[0] = byte(v)
	b[1] = byte(v >> 8)  //nolint:gomnd
	b[2] = byte(v >> 16) //nolint:gomnd
	b[3] = byte(v >> 24) //nolint:gomnd
	b[4] = byte(v >> 32) //nolint:gomnd
	b[5] = byte(v >> 40) //nolint:gomnd
	b[6] = byte(v >> 48) //nolint:gomnd
}

func (littleEndian) PutUint64(b []byte, v uint64) {
	_ = b[7] // early bounds check to guarantee safety of writes below
	b[0] = byte(v)
	b[1] = byte(v >> 8)  //nolint:gomnd
	b[2] = byte(v >> 16) //nolint:gomnd
	b[3] = byte(v >> 24) //nolint:gomnd
	b[4] = byte(v >> 32) //nolint:gomnd
	b[5] = byte(v >> 40) //nolint:gomnd
	b[6] = byte(v >> 48) //nolint:gomnd
	b[7] = byte(v >> 56) //nolint:gomnd
}

func (littleEndian) String() string { return "LittleEndian" }

type bigEndian struct{}

func (bigEndian) Uint16(b []byte) uint16 {
	_ = b[1] // bounds check hint to compiler; see golang.org/issue/14808
	return uint16(b[1]) | uint16(b[0])<<8
}

func (bigEndian) Uint24(b []byte) uint32 {
	_ = b[2] // bounds check hint to compiler; see golang.org/issue/14808
	return uint32(b[2]) | uint32(b[1])<<8 | uint32(b[0])<<16
}

func (bigEndian) Uint32(b []byte) uint32 {
	_ = b[3] // bounds check hint to compiler; see golang.org/issue/14808
	return uint32(b[3]) | uint32(b[2])<<8 | uint32(b[1])<<16 | uint32(b[0])<<24
}

func (bigEndian) Uint40(b []byte) uint64 {
	_ = b[4] // bounds check hint to compiler; see golang.org/issue/14808
	return uint64(b[4]) | uint64(b[3])<<8 | uint64(b[2])<<16 | uint64(b[1])<<24 |
		uint64(b[0])<<32
}

func (bigEndian) Uint48(b []byte) uint64 {
	_ = b[5] // bounds check hint to compiler; see golang.org/issue/14808
	return uint64(b[5]) | uint64(b[4])<<8 | uint64(b[3])<<16 | uint64(b[2])<<24 |
		uint64(b[1])<<32 | uint64(b[0])<<40
}

func (bigEndian) Uint56(b []byte) uint64 {
	_ = b[6] // bounds check hint to compiler; see golang.org/issue/14808
	return uint64(b[6]) | uint64(b[5])<<8 | uint64(b[4])<<16 | uint64(b[3])<<24 |
		uint64(b[2])<<32 | uint64(b[1])<<40 | uint64(b[0])<<48
}

func (bigEndian) Uint64(b []byte) uint64 {
	_ = b[7] // bounds check hint to compiler; see golang.org/issue/14808
	return uint64(b[7]) | uint64(b[6])<<8 | uint64(b[5])<<16 | uint64(b[4])<<24 |
		uint64(b[3])<<32 | uint64(b[2])<<40 | uint64(b[1])<<48 | uint64(b[0])<<56
}

func (bigEndian) PutUint16(b []byte, v uint16) {
	_ = b[1]            // early bounds check to guarantee safety of writes below
	b[0] = byte(v >> 8) //nolint:gomnd
	b[1] = byte(v)
}

func (bigEndian) PutUint24(b []byte, v uint32) {
	_ = b[2]             // early bounds check to guarantee safety of writes below
	b[0] = byte(v >> 16) //nolint:gomnd
	b[1] = byte(v >> 8)  //nolint:gomnd
	b[2] = byte(v)
}

func (bigEndian) PutUint32(b []byte, v uint32) {
	_ = b[3]             // early bounds check to guarantee safety of writes below
	b[0] = byte(v >> 24) //nolint:gomnd
	b[1] = byte(v >> 16) //nolint:gomnd
	b[2] = byte(v >> 8)  //nolint:gomnd
	b[3] = byte(v)
}

func (bigEndian) PutUint40(b []byte, v uint64) {
	_ = b[4]             // early bounds check to guarantee safety of writes below
	b[0] = byte(v >> 32) //nolint:gomnd
	b[1] = byte(v >> 24) //nolint:gomnd
	b[2] = byte(v >> 16) //nolint:gomnd
	b[3] = byte(v >> 8)  //nolint:gomnd
	b[4] = byte(v)
}

func (bigEndian) PutUint48(b []byte, v uint64) {
	_ = b[5]             // early bounds check to guarantee safety of writes below
	b[0] = byte(v >> 40) //nolint:gomnd
	b[1] = byte(v >> 32) //nolint:gomnd
	b[2] = byte(v >> 24) //nolint:gomnd
	b[3] = byte(v >> 16) //nolint:gomnd
	b[4] = byte(v >> 8)  //nolint:gomnd
	b[5] = byte(v)
}

func (bigEndian) PutUint56(b []byte, v uint64) {
	_ = b[6]             // early bounds check to guarantee safety of writes below
	b[0] = byte(v >> 48) //nolint:gomnd
	b[1] = byte(v >> 40) //nolint:gomnd
	b[2] = byte(v >> 32) //nolint:gomnd
	b[3] = byte(v >> 24) //nolint:gomnd
	b[4] = byte(v >> 16) //nolint:gomnd
	b[5] = byte(v >> 8)  //nolint:gomnd
	b[6] = byte(v)
}

func (bigEndian) PutUint64(b []byte, v uint64) {
	_ = b[7]             // early bounds check to guarantee safety of writes below
	b[0] = byte(v >> 56) //nolint:gomnd
	b[1] = byte(v >> 48) //nolint:gomnd
	b[2] = byte(v >> 40) //nolint:gomnd
	b[3] = byte(v >> 32) //nolint:gomnd
	b[4] = byte(v >> 24) //nolint:gomnd
	b[5] = byte(v >> 16) //nolint:gomnd
	b[6] = byte(v >> 8)  //nolint:gomnd
	b[7] = byte(v)
}

func (bigEndian) String() string { return "BigEndian" }
