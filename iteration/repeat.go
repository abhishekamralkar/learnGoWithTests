package iteration

func Repeat(x string, n int) string {
	var repeated string
	for i := 0; i < n; i++ {
		repeated += x
	}
	return repeated
}
