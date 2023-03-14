package gogit

import (
	"bytes"
	"compress/zlib"
	"encoding/json"
	"io"
	"log"
)

type IndexItem struct {
	Hash string
	Mode string
	Name string
	Type string
}

func UpdateIndexCMD(f, hash string, isCacheInfo bool) {
	UpdateIndex(f, hash, isCacheInfo)
}

func UpdateIndex(f, hash string, isCacheInfo bool) {
	if !isCacheInfo {
		var err error
		file, err := OpenFileDefault(f)
		if err != nil {
			log.Fatalf("open file error: %v", err)
		}
		defer file.Close()
		body, _ := io.ReadAll(file)
		hash, err = HashObject(body, BlobType, true)
		if err != nil {
			log.Fatalf("hash object error: %v", err)
		}
	}
	// read index file
	indexFile, err := OpenFileDefault(ConstantIndexFile)
	if err != nil {
		log.Fatalf("open file error: %v", err)
	}
	defer indexFile.Close()
	// read all index file
	var indexItems []IndexItem
	var b bytes.Buffer
	_, _ = io.Copy(&b, indexFile)
	if b.Len() > 0 {
		z, err := zlib.NewReader(&b)
		if err != nil {
			log.Fatalf("decompress index file error: %v", err)
		}
		defer z.Close()
		var bb bytes.Buffer
		_, _ = io.Copy(&bb, z)
		if err != nil {
			log.Fatalf("read index file error: %v", err)
		}
		// parse index file
		_ = json.Unmarshal(bb.Bytes(), &indexItems)
	}
	// update index file
	var isExist bool
	for i, item := range indexItems {
		if item.Name == f {
			indexItems[i].Hash = hash
			isExist = true
			break
		}
	}
	if !isExist {
		indexItems = append(indexItems, IndexItem{
			Hash: hash,
			Mode: "100644",
			Name: f,
			Type: "blob",
		})
	}
	// marshal index file
	indexBytes, err := json.Marshal(indexItems)
	if err != nil {
		log.Fatalf("marshal index file error: %v", err)
	}
	// compress index file
	var buf bytes.Buffer
	w := zlib.NewWriter(&buf)
	_, _ = w.Write(indexBytes)
	_ = w.Close()
	// rewrite index file, use old fd
	_, _ = indexFile.Seek(0, 0)
	_, _ = indexFile.Write(buf.Bytes())
	_ = indexFile.Truncate(int64(buf.Len()))
}
