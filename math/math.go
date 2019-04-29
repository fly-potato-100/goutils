package math

func Sum(sli []int) int {
	sum := 0
	for _, v := range sli {
		sum += v
	}
	return sum
}
