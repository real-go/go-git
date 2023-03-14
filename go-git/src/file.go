package gogit

import (
	"log"
	"os"
	"path"
)

func openFile(f string, flag int) (*os.File, error) {
	if checkErr := checkOrCreate(path.Dir(f)); checkErr != nil {
		return nil, checkErr
	}
	file, err := os.OpenFile(f, flag, 0644)
	if err != nil {
		log.Printf("open file error, %s", err.Error())
	}
	return file, nil
}

func OpenFileDefault(f string) (*os.File, error) {
	return openFile(f, os.O_RDWR|os.O_CREATE)
}

func checkOrCreate(p string) error {
	if _, err := os.Stat(p); os.IsNotExist(err) {
		err = os.MkdirAll(p, 0755)
		if err != nil {
			log.Printf("create path error, %s", err.Error())
			return err
		}
	}
	return nil
}
