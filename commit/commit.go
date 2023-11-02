package gomit

import (
	"fmt"
	"regexp"
)

var (
	o = getOption()
)
func Check_commit_size(commit string) bool {
	if len([]rune(commit)) > 10 {
		return true
	}
	fmt.Printf("Your commit message need to contain at least %d characters \n",o.MinimumSize)
	return false
} 

func Check_commit_lint(commit string) bool {
	for _, prefix := range o.Type {
		match, err := regexp.MatchString("^"+prefix+"(: |\\(.*\\): |!: )", commit)
		if match && err == nil {
			return true
		}
	}
	fmt.Println(string(colorRed)+"Your commit message doenst correspond to requirements"+string(colorReset))
	show_req(commit)
	return false
} 

 