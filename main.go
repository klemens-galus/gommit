package main

import (
	"gommit.go/commit"
	"os"
	"io/ioutil"
)

func main() {

	commitFile := os.Args[1]
	dat, _ := ioutil.ReadFile(commitFile)
	commit := string(dat)
	isValid := gomit.Check_commit(commit)
	if isValid {
		os.Exit(0)
	} else {
		os.Exit(1)
	}

}  
 