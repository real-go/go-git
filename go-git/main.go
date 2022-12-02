package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := OpenFile("./", "cat-file.go", os.O_RDWR|os.O_CREATE)
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
