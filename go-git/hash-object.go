package main

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"
)

func HashObject(f *os.File) (string, error) {
	h := sha1.New()
	b := make([]byte, 1024)
	for {
		n, err := f.Read(b)
		if err != nil {
			if err == io.EOF {
				break
			}
			return "", err
		}
		_, err = h.Write(b[:n])
		if err != nil {
			return "", err
		}
	}
	defer func() {
		if err := f.Close(); err != nil {
			panic(err)
		}
	}()
	// save to the objects file.
	fo, err := OpenFile(ObjectsPath, fmt.Sprintf("%x", h.Sum(nil)), os.O_RDWR|os.O_CREATE)
	if err != nil {
		return "", err
	}
	defer func(f *os.File) {
		if err := fo.Close(); err != nil {
			panic(err)
		}
	}(f)
	var blob = Blob{
		Hash: fmt.Sprintf("%x", h.Sum(nil)),
		Body: string(b),
	}
	err = Compress(fo, []byte(blob.Body))
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", h.Sum(nil)), nil
}

func SHA1Hash(files []string, open func(string) (io.ReadCloser, error)) (string, error) {
	h := sha1.New()
	files = append([]string(nil), files...)
	sort.Strings(files)
	for _, file := range files {
		if strings.Contains(file, "\n") {
			return "", errors.New("dir_hash: filenames with newlines are not supported")
		}
		r, err := open(file)
		if err != nil {
			return "", err
		}
		hf := sha1.New()
		_, err = io.Copy(hf, r)
		_ = r.Close()
		if err != nil {
			return "", err
		}
		_, _ = fmt.Fprintf(h, "%x  %s\n", hf.Sum(nil), file)
	}
	return fmt.Sprintf("%x", h.Sum(nil)), nil
}
