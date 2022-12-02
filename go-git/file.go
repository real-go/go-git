package main

import (
	"log"
	"os"
	"path"
)

func OpenFile(p, f string, flag int) (*os.File, error) {
	if _, err := os.Stat(p); os.IsNotExist(err) {
		err := os.MkdirAll(p, 0755)
		if err != nil {
			log.Printf("create path error, %s", err.Error())
			return nil, err
		}
	}
	file, err := os.OpenFile(path.Join(p+f), flag, 0755)
	if err != nil {
		log.Printf("open file error, %s", err.Error())
	}
	return file, nil
}
