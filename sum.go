package gosum

func Sum(a ...int) int {
	s := 0
	for _, v := range a {
		s += v
	}
	return s
}
