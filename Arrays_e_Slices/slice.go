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
		var rest []int
		for i, j := range v {
			if i != 0 {
				rest = append(rest, j)
			}
		}
		soma = append(soma, Soma(rest))
	}
	return soma
}
