package ioutil

import (
	"encoding/binary"
	"os"
)

func ExampleDataOutput() error {
	writer, err := os.Open("file")
	if err != nil {
		return err
	}
	defer writer.Close()
	dout := NewDataOutput(binary.LittleEndian, writer)
	dout.WriteBytes([]byte{'h', 'e', 'l', 'l', 'o'})
	dout.WriteInt32(1234)
	dout.WriteUTF8(I8, "hello world")
	dout.WriteBool(true)
	if dout.Error() != nil {
		return dout.Error()
	}
	return nil
}
