package utils

import "reflect"

func Remove[T any](slice []T, s T) []T {
	n := 0
	for _, x := range slice {
		if !reflect.DeepEqual(x, s) {
			slice[n] = x
			n++
		}
	}
	return slice[:n]
}
