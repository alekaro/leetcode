// An array is monotonic if it is either monotone increasing or monotone decreasing.
// An array nums is monotone increasing if for all i <= j, nums[i] <= nums[j]. An array nums is monotone decreasing if for all i <= j, nums[i] >= nums[j].
// Given an integer array nums, return true if the given array is monotonic, or false otherwise.

// Example 1:

// Input: nums = [1,2,2,3]
// Output: true
// Example 2:

// Input: nums = [6,5,4,4]
// Output: true
// Example 3:

// Input: nums = [1,3,2]
// Output: false

// Constraints:

// 1 <= nums.length <= 105
// -105 <= nums[i] <= 105

package main

import "fmt"

func determine(num1 *int, num2 *int) int {
	if *num1 < *num2 {
		return 1
	} else if *num1 == *num2 {
		return 0
	}
	return -1
}

func isMonotonic(nums []int) bool {
	if len(nums) <= 2 {
		return true
	}
	monotony := -2
	var temp int
	for index := 0; index < len(nums)-1; index++ {
		temp = determine(&nums[index], &nums[index+1])
		if temp == 0 {
			continue
		} else if monotony == -2 {
			monotony = temp
		} else if temp != monotony {
			return false
		}
	}
	return true
}

func main() {
	nums := []int{1, 3, 2}
	// var temp int
	fmt.Println(isMonotonic(nums))
	// fmt.Println(temp)
}
