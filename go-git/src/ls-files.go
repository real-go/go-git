package gogit

import (
	"bytes"
	"compress/zlib"
	"encoding/json"
	"fmt"
	"io"
	"log"
)

func LsFilesCMD(hash string, isStage bool) {

}

func LsStageFiles() {
	// read index file
	file, err := OpenFileDefault(ConstantIndexFile)
	if err != nil {
		log.Fatalf("open file error: %v", err)
	}
	defer file.Close()
	z, err := zlib.NewReader(file)
	if err != nil {
		log.Fatalf("decompress index file error: %v", err)
	}
	defer z.Close()
	var b bytes.Buffer
	_, _ = io.Copy(&b, z)
	if err != nil {
		log.Fatalf("read index file error: %v", err)
	}

	// parse index file
	var indexItems []IndexItem
	_ = json.Unmarshal(b.Bytes(), &indexItems)
	for _, item := range indexItems {
		fmt.Printf("%s %s %s %s\n", item.Mode, item.Type, item.Hash, item.Name)
	}
}
