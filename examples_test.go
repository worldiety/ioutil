package ioutil

import (
	"os"
)

//nolint:gomnd
func ExampleDataOutput() {
	writer, err := os.Open("file")
	if err != nil {
		panic(err)
	}
	defer writer.Close()
	dout := NewDataOutput(LittleEndian, writer)
	dout.WriteBytes('h', 'e', 'l', 'l', 'o')
	dout.WriteInt32(1234)
	dout.WriteUTF8(I8, "hello world")
	dout.WriteBool(true)

	if dout.Error() != nil {
		panic(dout.Error())
	}
}

/*
func ExampleFS() {
	file := NewOpenOptions().
		Read(true).
		Write(true).
		Create(true).
		Open("foo.txt")
	defer file.Close()

	NewQuery().
		Select().Files().From("/foo/bar").Recursively().
		Where().Attribute().NotHidden().And().Name().EndsWidth(".jpg").
		Execute()
}
*/
