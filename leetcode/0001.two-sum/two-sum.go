package main

import "fmt"

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for k, v := range nums {
		idx, ok := m[target-v]
		if ok {
			return []int{idx, k}
		}
		m[v] = k
	}
	return nil
}
func main() {
	var q = []int{1, 2, 3}
	fmt.Println(twoSum(q, 4))
}
