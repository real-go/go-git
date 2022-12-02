package main

import (
	"bytes"
	"compress/zlib"
	"io"
	"os"
)

func DeCompress(f *os.File) ([]byte, error) {
	z, err := zlib.NewReader(f)
	if err != nil {
		return nil, err
	}
	defer z.Close()
	var b bytes.Buffer
	io.Copy(&b, z)
	return b.Bytes(), nil
}

func Compress(f *os.File, b []byte) error {
	w := zlib.NewWriter(f)
	_, err := w.Write(b)
	if err != nil {
		return err
	}
	return w.Close()
}
