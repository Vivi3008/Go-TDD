package arrayseslices

func Soma(numbers []int) int {
	var total int
	for _, v := range numbers {
		total += v
	}
	return total
}

func SomaTodoResto(numeros ...[]int) []int {
	var soma []int
	for _, v := range numeros {
		if len(v) == 0 {
			soma = append(soma, 0)
		} else {
			final := v[1:]
			soma = append(soma, Soma(final))
		}
	}
	return soma
}
