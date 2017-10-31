package raindrops

import "strconv"
const testVersion = 3

func Convert(num int) string{
	result := ""
	if num%3 == 0 { result += "Pling" }
	if num%5 == 0 { result += "Plang" }
	if num%7 == 0 { result += "Plong" }
        if result == "" { result = strconv.Itoa(num) }
	return result
}
// Don't forget the test program has a benchmark too.
// How fast does your Convert convert?
