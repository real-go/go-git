package gogit

import (
	"fmt"
	"log"
	"os"
	"path"
)

func InitRepo(p string) error {
	pa := path.Join(p, ConstantObjectsPath)
	_, err := os.Stat(pa)
	if err == nil {
		return fmt.Errorf("fatal: destination path '%s' already exists and is not an empty directory", pa)
	}
	if os.IsNotExist(err) {
		err = os.MkdirAll(pa, 0755)
		if err != nil {
			log.Printf("create path error, %s", err.Error())
			return err
		}
	}
	return nil
}
