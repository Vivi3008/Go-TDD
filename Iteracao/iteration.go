package iteracao

func Repeat(txt string) string {
	var res string
	for i := 0; i < 5; i++ {
		res += txt
	}
	return res
}
