package main

import (
	"gommit.go/commit"
	"os"
	//"io/ioutil"
)

func main() {

	commitFile := os.Args[1]
	// dat, _ := ioutil.ReadFile(commitFile)
	// commit := string(dat)
	commit := commitFile
	isValid := gomit.Check_commit_size(commit)
	gomit.Check_commit_lint(commit)
	if isValid {
		os.Exit(0)
	} else {
		os.Exit(1)
	}

}  
 