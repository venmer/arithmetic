package arithmetic

func Add(numbers ...int) int {
	sum := 0

	for _, n := range numbers {
		sum += n
	}
	return sum

}
