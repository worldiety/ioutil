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

// A DataOutput provides helpers for writing bytes into a binary stream and interprets the data for any primitive Go
// type. A DataOutput is always tied to a specific endianness. A DataOutput should not be considered thread safe.
// As soon as any error occurred, any call is a no-op and will result in the same error state.
//
// Example:
//
//   writer, err := os.Open("file")
//   if err != nil {
//      return err
//   }
//   defer writer.close()
//   dout := ioutil.NewDataOutput(binary.LittleEndian, writer)
//   dout.WriteArray([]byte{'h', 'e', 'l', 'l', 'o'})
//   dout.WriteInt32(1234)
//   dout.WriteUTF8(ioutil.I8, "hello world")
//   dout.WriteBool(true)
//	 if dout.Error() != nil{
//      return dout.Error()
//	 }
//
type DataOutput interface {
	// WriteArray just writes the slice out, without any prefix for the length.
	// If an error occurs returns the number of written bytes.
	WriteArray(v []byte) int

	// WriteBlob writes a prefixed byte slice of variable length.
	WriteBlob(p IntSize, v []byte)

	// WriteUTF8 writes a prefixed unmodified utf8 string sequence of variable length.
	WriteUTF8(p IntSize, v string)

	// WriteBool writes one byte.
	WriteBool(v bool)

	// WriteUint8 writes an unsigned byte
	WriteUint8(v uint8)

	// WriteInt8 writes a signed byte
	WriteInt8(v int8)

	// WriteUint16 writes an unsigned 2 byte integer.
	WriteUint16(v uint16)

	// WriteInt16 writes a signed 2 byte integer.
	WriteInt16(v int16)

	// WriteUint24 writes an unsigned 3 byte integer.
	WriteUint24(v uint32)

	// WriteInt24 writes a signed 3 byte integer.
	WriteInt24(v int32)

	// WriteUint32 writes an unsigned 4 byte integer.
	WriteUint32(v uint32)

	// WriteInt32 writes a signed 4 byte integer.
	WriteInt32(v int32)

	// WriteInt40 writes a signed 5 byte integer.
	WriteInt40(v int64)

	// WriteUint40 writes an unsigned 5 byte integer.
	WriteUint40(v uint64)

	// WriteUint64 writes an unsigned 8 byte integer.
	WriteUint64(v uint64)

	// WriteInt64 writes a signed 8 byte integer.
	WriteInt64(v int64)

	// WriteUvarint writes a variable length integer, up to 9 bytes using zig-zag protobuf encoding.
	WriteUvarint(v uint64)

	// WriteVarint writes a variable length and signed integer, up to 9 bytes using zig-zag protobuf encoding.
	WriteVarint(v int64)

	// WriteFloat32 writes a float32 IEEE 754 4 byte bit sequence.
	WriteFloat32(v float32)

	// WriteFloat64 writes a float64 IEEE 754 8 byte bit sequence.
	WriteFloat64(v float64)

	// WriteComplex64 writes two float32 IEEE 754 4 byte bit sequences.
	WriteComplex64(v complex64)

	// WriteComplex128 writes two float32 IEEE 754 4 byte bit sequences.
	WriteComplex128(v complex128)

	// Error returns the first occurred error. Each call to any Write* method may cause an error. Per definition,
	// any other call after the first error is a no-op.
	Error() error

	io.Writer
	io.ByteWriter
}
