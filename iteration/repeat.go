package iteration

func Repeat(item string, times int) string {
	var repeated string
	for i := 0; i < times; i++ {
		repeated += item
	}
	return repeated
}
