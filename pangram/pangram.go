package pangram

import (
	"strings"
)

const testVersion = 1

func IsPangram(s string) bool {
    l := make([]bool, 26)
    for _, r := range strings.ToLower(s) {
        c := string(r)
        if(IsLetter(c)){
            l[r-97] = true
        }
    }
    for _, b := range l {
	if b == false {
	    return false
	}
    }
    return true
}

func IsLetter(s string) bool {
    for _, r := range s {
        if (r < 'a' || r > 'z') && (r < 'A' || r > 'Z') {
            return false
        }
    }
    return true
}
func Reverse(s string) string {
	var rev string
	for i := len(s)-1; i >= 0; i-- {
		rev += string(s[i])
	}
	return rev
}

