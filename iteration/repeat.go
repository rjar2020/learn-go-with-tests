package iteration

//Repeat prints a character multiple times
func Repeat(character string, times int) string {
	repeated := character
	if times > 0 {
		for i := 1; i < times; i++ {
			repeated += character
		}
	}
	return repeated
}
