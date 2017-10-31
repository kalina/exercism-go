package accumulate

const testVersion = 1

func Accumulate(array []string, fun func(string) string) []string {
	for i := 0; i < len(array); i ++ {
		array[i] = fun(array[i])
	}
	return array
}
