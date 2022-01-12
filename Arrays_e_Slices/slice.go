package arrayseslices

func Soma(numbers []int) int {
	var total int
	for _, v := range numbers {
		total += v
	}
	return total
}

func SomaTudo(numeros ...[]int) []int {
	var soma []int
	for _, v := range numeros {
		soma = append(soma, Soma(v))
	}
	return soma
}
