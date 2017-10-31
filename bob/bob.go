package bob

import (
	"strings"
)
const testVersion = 3

const alpha = "abcdefghijklmnopqrstuvwxyz"

func ContainsAlpha(s string) bool {
   for _, char := range s {
      if strings.Contains(alpha, strings.ToLower(string(char))) {
         return true
      }
   }
   return false
}

func Hey(s string) string {

     //get rid of whitespace
     s = strings.TrimSpace(s)

     //empty string
     if(len(s) == 0) {
	 return "Fine. Be that way!"
     }
     //shouting
     if ContainsAlpha(s) && strings.ToUpper(s) == s {
         return "Whoa, chill out!"
     }
     //question
     if strings.HasSuffix(s, "?") { //strings.Contains(s, "?"){
         return "Sure."
     }
     //default
     return "Whatever."
}
