package test

import "fmt"

/**
 * https://leetcode-cn.com/problems/two-sum/description/
 * 两数之和
 */
func main() {
	var data = []int{2, 7, 11, 15}
	data = []int{2, 2, 3, 5}
	var target = 4
	res := twoSum(data, target)
	fmt.Println(res)
}

func twoSum(nums []int, target int) []int {
	var numsMap map[int]int
	numsMap = make(map[int]int)
	for k, v := range nums {
		fmt.Printf("k=%d ;v=%d ", k, v)
		key, ok := numsMap[target-v]
		fmt.Printf("key=%d ;v=%d ", key, ok)
		if ok && k != key {
			return []int{key, k}
		} else {
			numsMap[v] = k
		}
	}

	return []int{}
}
