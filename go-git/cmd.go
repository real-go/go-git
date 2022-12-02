package main

import (
	"flag"
	gogit "go-git/src"
	"os"
)

var ()

func main() {
	switch os.Args[1] {
	case "init":
		gogit.InitRepo()
	default:
		flag.PrintDefaults()
	}
}
