package acronym

import (
	"strings"
	"regexp"
)

const testVersion = 3

func Abbreviate(str string) string {
	acronym := ""
	re := regexp.MustCompile("[- ]")
	var str_arr = re.Split(str, -1)

	for i := 0; i < len(str_arr); i ++ {
		acronym += strings.ToUpper(str_arr[i][:1])
	}
	return acronym
}
