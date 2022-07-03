package main

import "fmt"

func main() {
	fmt.Println("hello")
	fmt.Println("hello")
	fmt.Println("hello")
	fmt.Println("hello")

	nums := []int{2, 3, 4, 5, 1}
	fmt.Println(twoSum(nums, 7))
}

func twoSum(nums []int, target int) []int {
	dic := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		another := target - nums[i]
		if _, ok := dic[another]; ok {
			return []int{dic[another], i}
		}
		dic[nums[i]] = i
	}

	return nil
}
