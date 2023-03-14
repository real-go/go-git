package gogit

import (
	"bytes"
	"compress/zlib"
	"crypto/sha1"
	"fmt"
	"io"
	"log"
	"path"
)

type Blob struct {
	Hash string
	Body []byte
}

func (b *Blob) Type() ObjectType {
	return BlobType
}

func (b *Blob) HashObject(isWrite bool) (string, error) {
	header := fmt.Sprintf("%s %d\x00", b.Type().String(), len(b.Body))
	data := header + string(b.Body)
	h := sha1.New()
	_, err := io.Copy(h, bytes.NewReader([]byte(data)))
	if err != nil {
		return "", err
	}
	b.Hash = fmt.Sprintf("%x", h.Sum(nil))
	if isWrite {
		if err = Write(b.Hash, []byte(data)); err != nil {
			return "", err
		}
	}
	return b.Hash, nil
}

func Write(hash string, data []byte) error {
	f, err := OpenFileDefault(path.Join(ConstantObjectsPath, hash[:2], hash[2:]))
	if err != nil {
		return err
	}
	defer f.Close()
	w := zlib.NewWriter(f)
	_, err = w.Write(data)
	if err != nil {
		return err
	}
	_ = w.Close()
	return nil
}

func HashObject(body []byte, tp ObjectType, isWrite bool) (string, error) {
	switch tp {
	case BlobType:
		var blob = Blob{Body: body}
		out, err := blob.HashObject(isWrite)
		if err != nil {
			return out, err
		}
		return out, nil
	default:
		return "", fmt.Errorf("unknown object type")
	}
}

func HashObjectCMD(f string, isWrite bool) {
	file, err := OpenFileDefault(f)
	if err != nil {
		log.Fatalf("open file error: %v", err)
	}
	defer file.Close()
	body, _ := io.ReadAll(file)
	hash, err := HashObject(body, BlobType, isWrite)
	if err != nil {
		log.Fatalf("hash object error: %v", err)
	}
	fmt.Println(hash)
}
