package main

import (
	"flag"
	"fmt"
	gogit "go-git/src"
	"log"
	"os"
)

func main() {
	switch os.Args[1] {
	case "help": // help
		fmt.Println(
			"Usage: go-git <command> [<args>]\n",
			"  init: create an empty Git repository\n",
			"  hash-object: compute object ID and optionally creates a blob from a file\n",
			"  cat-file: provide content or type and size information for repository objects\n",
			"  update-index: add file contents to the index\n",
			"  write-tree: build a tree object from the current index\n",
			"  ls-tree: list the contents of a tree object",
		)
	case "init":
		gogit.InitRepo(".")
	case "hash-object":
		if len(os.Args) < 3 {
			log.Fatalln("usage: hash-object <file>")
		}
		if len(os.Args) == 4 && os.Args[3] == "-w" {
			gogit.HashObjectCMD(os.Args[2], true)
			return
		}
		gogit.HashObjectCMD(os.Args[2], false)

	case "cat-file":
		if len(os.Args) < 3 {
			log.Fatalln("usage: cat-file <hash>")
		}
		gogit.CatFileCMD(os.Args[2])
	case "update-index":
		if len(os.Args) == 4 {
			gogit.UpdateIndexCMD(os.Args[3], "", false)
			return
		}
		gogit.UpdateIndexCMD(os.Args[3], os.Args[4], true)
	case "write-tree":
		gogit.WriteTreeCMD()
	case "ls-files":
		if len(os.Args) == 3 && os.Args[2] == "-s" {
			gogit.LsStageFiles()
		}
	default:
		flag.PrintDefaults()
	}
}
