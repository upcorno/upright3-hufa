package utils

import (
	"math/rand"
)

func RandSlice(num int, exceptArr []int) []int {
	res := rand.Perm(10)
	for k := range res {
		res[k]++
	}
	for _, v := range exceptArr {
		if v >= 1 && v <= 10 {
			res = append(res[:v-1], res[v:]...)
		}
	}
	return res[0:num]
}
