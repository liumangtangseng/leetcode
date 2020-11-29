package main

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9

	for _, v := range twoSum(nums, target) {
		println(v)
	}
}
func twoSum(nums []int, target int) []int {
	for key, value := range nums {
		left := target - value
		for k, v := range nums {
			if k == key {
				continue
			}
			if v == left {
				return []int{key, k}
			}
		}
	}
	return nil
}
