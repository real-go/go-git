package gogit_test

import (
	gogit "go-git/src"
	"os"
	"path"
	"testing"
)

func TestInitRepo(t *testing.T) {
	err := gogit.InitRepo("/tmp/go-git")
	if err != nil {
		t.Error(err)
	}
	if _, err = os.Stat(path.Join("/tmp/go-git", gogit.ConstantObjectsPath)); os.IsNotExist(err) {
		if err != nil {
			t.Error(err)
		}
	}
}
