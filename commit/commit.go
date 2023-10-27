package gomit

import (

)


func Check_commit(commit string) bool {
	if len([]rune(commit)) > 10 {
		return true
	} else { 
		return false
	}
} 