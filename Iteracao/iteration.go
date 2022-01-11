package iteracao

// Receive two parameters, first the string that will be repeat.
// Second a int wich represents the number of times that the string be repeat
func Repeat(txt string, rep int) string {
	var res string
	for i := 0; i < rep; i++ {
		res += txt
	}
	return res
}
