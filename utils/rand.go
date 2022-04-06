package utils

import "math/rand"

func RandSlice(num int) []int {
	res := rand.Perm(10)
	for k, _ := range res {
       res[k]++
	}
	return res[0: num]
}