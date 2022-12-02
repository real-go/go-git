package gogit

import (
	"bytes"
	"compress/zlib"
	"fmt"
	"io"
	"os"
	"path"
)

func ReadBlobObj(hash string) (*Blob, error) {
	f, err := OpenFileDefault(path.Join(ConstantObjectsPath, hash[:2], hash[2:]))
	if err != nil {
		return nil, err
	}
	defer func(f *os.File) {
		if err := f.Close(); err != nil {
			panic(err)
		}
	}(f)
	z, err := zlib.NewReader(f)
	if err != nil {
		return nil, err
	}
	defer z.Close()
	var b bytes.Buffer
	_, _ = io.Copy(&b, z)
	if err != nil {
		return nil, err
	}
	return &Blob{Hash: hash, Body: parseBlob(b.Bytes())}, nil
}

func parseBlob(data []byte) []byte {
	i, length := 0, 0
	data = data[5:] // skip "blob "
	for ; data[i] != 0; i++ {
		length = length*10 + int(data[i]-'0')
	}
	fmt.Println(i, length)
	return data[i+1 : i+1+length]
}
