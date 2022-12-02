package main

import (
	"os"
)

var nilBlob = Blob{}

func ReadObject(hash string) (Blob, error) {
	f, err := OpenFile(ObjectsPath, hash, os.O_RDWR|os.O_CREATE)
	if err != nil {
		return nilBlob, err
	}
	defer func(f *os.File) {
		if err := f.Close(); err != nil {
			panic(err)
		}
	}(f)
	out, err := DeCompress(f)
	if err != nil {
		return nilBlob, err
	}
	var blob = Blob{
		Hash: hash,
		Body: string(out),
	}
	return blob, nil
}
