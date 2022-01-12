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
		final := v[1:]
		soma = append(soma, Soma(final))
	}
	return soma
}
