package gomit

import (
	"fmt"
	"regexp"
)

var (
	o = getOption()
)
func CheckCommitSize(commit string) bool {
	if len([]rune(commit)) > 10 {
		return true
	}
	fmt.Printf("Your commit message need to contain at least %d characters \n",o.MinimumSize)
	return false
} 

func CheckCommitLint(commit string) bool {
	for _, prefix := range o.Type {
		match, err := regexp.MatchString("^"+prefix+"(: |\\(.*\\): |!: )", commit)
		if match && err == nil {
			return true
		}
	}
	fmt.Println(string(colorRed)+"Your commit message doesn't correspond to requirements"+string(colorReset))
	ShowReq(commit)
	return false
} 

 