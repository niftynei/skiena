package three

import (
)

func XORProduct(arr []int) ([]int64) {
	back := make([]int64, len(arr))
	ans := make([]int64, len(arr))

	if len(arr) < 2 {
		for i, val := range arr {
			ans[i] = int64(val)
		}
		return ans
	}

	for i := len(arr)-1; i > -1; i-- {
		if i == len(arr) -1 {
			back[i] = int64(arr[i])
		} else {
			back[i] = back[i+1] * int64(arr[i])
		}
	}
	var acc int64
	for i, val := range arr {
		if i == 0 {
			acc = int64(val)
			ans[i] = back[i+1]
		} else if i == len(arr) - 1 {
			ans[i] = acc
		} else {
			ans[i] = acc * back[i+1]
			acc = acc * int64(val)
		}
	}

	return ans
}
