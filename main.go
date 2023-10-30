package main

import (
	"fmt"
	"gommit.go/commit"
	"os"
	"io/ioutil"
)
var (
	debugMode = false 
)  

func main() {

	commitFile := os.Args[1] //GIVE THE PATH OF THE COMMIT FILE
	dat, _ := ioutil.ReadFile(commitFile) //READ THE COMMIT FILE
	commit := string(dat)
	isValid := gomit.Check_commit_size(commit) //CHECK THE SIZE OF THE COMMIT
	isValid2 := gomit.Check_commit_lint(commit) //LINT THE COMMIT

	if isValid && isValid2 {
		if debugMode {
			fmt.Println("Well executed")
			os.Exit(1)
		}else {
			os.Exit(0)
		}
	} else {
		os.Exit(1)
	}

}  
  