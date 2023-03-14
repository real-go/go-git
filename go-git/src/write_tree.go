package gogit

import (
	"crypto/sha1"
	"fmt"
	"io"
	"log"
)

func WriteTreeCMD() {
	// TODO:
	// sha1 index file
	f, err := OpenFileDefault(ConstantIndexFile)
	if err != nil {
		log.Fatalf("open file error: %v", err)
	}
	defer f.Close()
	h := sha1.New()
	_, err = io.Copy(h, f)
	if err != nil {
		log.Fatalf("copy file error: %v", err)
	}

	hash := fmt.Sprintf("%x", h.Sum(nil))
	// open
	//ff, err := OpenFileDefault(path.Join(ConstantObjectsPath, hash[:2], hash[2:]))

}
