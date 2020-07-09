package ioutil

import "testing"

func BenchmarkLittleEndianBuffer_ReadUint16(b *testing.B) {
	le := LittleEndianBuffer{
		Bytes: []byte{1,2,3,4,5,6,7,8,9,10},
		Pos:   0,
	}
	for n := 0; n < b.N; n++ {
		le.ReadUint16()
		le.ReadUint16()
		le.ReadUint16()
		le.ReadUint16()
		le.ReadUint16()
		le.Pos = 0
	}
}

func BenchmarkLittleEndianBuffer_ReadUint32(b *testing.B) {
	le := LittleEndianBuffer{
		Bytes: []byte{1,2,3,4,5,6,7,8,9,10,1,2,3,4,5,6,7,8,9,10,1,2,3,4,5,6,7,8,9,10,1,2,3,4,5,6,7,8,9,10,1,2,3,4,5,6,7,8,9,10},
		Pos:   0,
	}
	for n := 0; n < b.N; n++ {
		le.ReadUint32()
		le.ReadUint32()
		le.ReadUint32()
		le.ReadUint32()
		le.ReadUint32()
		le.Pos = 0
	}
}