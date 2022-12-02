package main

import (
	"fmt"
	"os"
	"testing"
)

func TestBlobBase(t *testing.T) {
	f, err := OpenFile("./", "LICENSE", os.O_RDWR|os.O_CREATE)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	hash, err := HashObject(f)
	if err != nil {
		panic(err)
	}
	b, err := ReadObject(hash)
	if err != nil {
		panic(err)
	}
	fmt.Println("hash: ", hash, "\nbody: \n", b.Body)
}

func TestTreeUpdateIndexWithTemp(t *testing.T) {
	Write("", "72d19d4a770684aa2355d212e4b0c029bbf70ae8", "cat-file.go", BlobType)
}

func TestTreeLsTreeWithTemp(t *testing.T) {
	tt, _ := LsTree("")
	fmt.Println(tt)
	Write("", "caa7f778a50bb8f5ba44e025536c9308739dc1a0", "LICENSE", BlobType)
	tt, _ = LsTree("")
	fmt.Println(tt)
}
