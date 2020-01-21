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

// UintSize is either 32 or 64
const UintSize = 32 << (^uint(0) >> 32 & 1)

const (
	// MaxUint is either 1<<31 - 1 or 1<<63 - 1
	MaxInt = 1<<(UintSize-1) - 1

	// MinInt is either -1 << 31 or -1 << 63
	MinInt = -MaxInt - 1

	// MaxUint is either 1<<32 - 1 or 1<<64 - 1
	MaxUint = 1<<UintSize - 1


	MaxInt8   = 1<<7 - 1
	MinInt8   = -1 << 7
	MaxInt16  = 1<<15 - 1
	MinInt16  = -1 << 15
	MaxInt24  = 1<<23 - 1
	MinInt24  = -1 << 23
	MaxInt32  = 1<<31 - 1
	MinInt32  = -1 << 31
	MaxInt40  = 1<<39 - 1
	MinInt40  = -1 << 39
	MaxInt64  = 1<<63 - 1
	MinInt64  = -1 << 63
	MaxUint8  = 1<<8 - 1
	MaxUint16 = 1<<16 - 1
	MaxUint24 = 1<<24 - 1
	MaxUint32 = 1<<32 - 1
	MaxUint40 = 1<<40 - 1
	MaxUint64 = 1<<64 - 1


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
