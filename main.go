package main

import (
	"fmt"
	"gommit.go/commit"
	"os"
	"io/ioutil"
	"strconv"
)
var (
	debugMode string
	isDebugMode = false 
)  

func main() {
	isDebugMode, err := strconv.ParseBool(debugMode)
    if err != nil {
        fmt.Println("Error parsing debugMode:", err)
        return
    }

	isValid := false 
	isValid2 := false
	commit := ""
	if !isDebugMode {
		commitFile := os.Args[1] //GIVE THE PATH OF THE COMMIT FILE
		dat, _ := ioutil.ReadFile(commitFile) //READ THE COMMIT FILE
		commit = string(dat)
	}else{
		dat := os.Args[1]
		commit = string(dat)
		
	}
	isValid = gommit.CheckCommitSize(commit) //CHECK THE SIZE OF THE COMMIT
	isValid2 = gommit.CheckCommitLint(commit) //LINT THE COMMIT
	if isValid && isValid2 {
		if isDebugMode {
			fmt.Println("Well executed")
			os.Exit(1)
		}else {
			os.Exit(0)
		}
	} else {
		os.Exit(1)
	}

}  
  