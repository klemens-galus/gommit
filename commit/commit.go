package gomit

import (
	"fmt"
)

func Check_commit_size(commit string) bool {
	if len([]rune(commit)) > 10 {
		return true
	} else { 
		fmt.Printf("Your commit message need to contain at least %d characters \n",minCommitSize)
		return false
	}
} 

