package calculate

import "strconv"

func string2int(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}

func Add(a, b int) int {
	return a + b
}

func Multi(a, b int) int {
	return a * b
}
