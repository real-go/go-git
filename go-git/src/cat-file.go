package gogit

import (
	"bytes"
	"compress/zlib"
	"fmt"
	"io"
	"log"
	"path"
)

func ReadHashFile(hash string) ([]byte, error) {
	f, err := OpenFileDefault(path.Join(ConstantObjectsPath, hash[:2], hash[2:]))
	if err != nil {
		return nil, err
	}
	defer f.Close()
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
	return b.Bytes(), nil
}

func ParseData(data []byte) []byte {
	var tp string
	for i := 0; data[i] != ' '; i++ {
		tp += string(data[i])
	}
	switch tp {
	case "blob":
		i, length := 0, 0
		data = data[5:] // skip "blob "
		for ; data[i] != 0; i++ {
			length = length*10 + int(data[i]-'0')
		}
		return data[i+1 : i+1+length]
	case "tree":
		// TODO
		return nil
	case "commit":
		// TODO
		return nil
	default:
		return nil
	}
}

func CatFile(hash string) ([]byte, error) {
	data, err := ReadHashFile(hash)
	if err != nil {
		log.Fatalf("read hash file error: %v\n", err)
	}
	return ParseData(data), nil
}

func CatFileCMD(hash string) {
	out, err := CatFile(hash)
	if err != nil {
		log.Fatalf("cat file error: %v\n", err)
	}
	fmt.Println(string(out))
}
