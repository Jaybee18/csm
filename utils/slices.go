package utils

func Sum(slice []int) (res int) {
	for _, num := range slice {
		res += num
	}
	return
}
