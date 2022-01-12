package arrayseslices

func Soma(numbers []int) int {
	var total int
	for _, v := range numbers {
		total += v
	}
	return total
}
