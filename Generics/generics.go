package generics

func sumIntsOrFloas[V int | float64](m []V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}
