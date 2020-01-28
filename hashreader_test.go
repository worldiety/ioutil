package ioutil

import (
	"bytes"
	"crypto/md5" //nolint
	"encoding/hex"
	"reflect"
	"testing"
)

func TestHashingReader(t *testing.T) {
	src := []byte{'a', 'b', 'c'}
	reader := NewHashReader(md5.New(), bytes.NewBuffer(src)) //nolint
	tmp := make([]byte, 6)
	n, err := reader.Read(tmp)

	if err != nil {
		t.Fatal(err)
	}

	if n != len(src) {
		t.Fatalf("expected %d but got %d", len(src), n)
	}

	if !reflect.DeepEqual(src, tmp[0:n]) {
		t.Fatalf("expected \n%v\n but got \n%v", src, tmp)
	}

	readHash := reader.Sum()
	expectedHash := md5.Sum(src) //nolint

	if !bytes.Equal(readHash, expectedHash[:]) {
		t.Fatalf("expected \n%x\n but got \n%x", expectedHash, readHash)
	}

	if hex.EncodeToString(readHash) != "900150983cd24fb0d6963f7d28e17f72" {
		t.Fatalf("invalid sum")
	}
}
