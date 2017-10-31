package leap

const testVersion = 3

func IsLeapYear(a int) bool {
	// Write some code here to pass the test suite.
	return a%4 ==0  && (!(a%100 ==0) || a%400 == 0)
}
