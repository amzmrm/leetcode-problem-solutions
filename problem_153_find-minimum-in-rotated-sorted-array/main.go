package main

import (
	"log"
)

func main() {
	type test struct {
		name   string
		nums   []int
		expect int
	}

	tests := make([]*test, 0)
	tests = append(tests, &test{
		name:   "length is 1",
		nums:   []int{1},
		expect: 1,
	})
	tests = append(tests, &test{
		name:   "length is 2",
		nums:   []int{2, 1},
		expect: 1,
	})
	tests = append(tests, &test{
		name:   "0 nums rotated",
		nums:   []int{1, 2, 3, 4, 5},
		expect: 1,
	})
	tests = append(tests, &test{
		name:   "mid location is uncertain case 1",
		nums:   []int{2, 1, 2, 2, 2},
		expect: 1,
	})
	tests = append(tests, &test{
		name:   "mid location is uncertain case 2",
		nums:   []int{2, 1, 2, 2, 2},
		expect: 1,
	})
	tests = append(tests, &test{
		name:   "normal",
		nums:   []int{3, 4, 5, 1, 2},
		expect: 1,
	})

	for _, test := range tests {
		got := FindMin(test.nums)
		if got != test.expect {
			log.Printf("test [%s] failed, expect: %d, got: %d", test.name, test.expect, got)
		} else {
			log.Printf("test [%s] pass", test.name)
		}
	}
}

func FindMin(nums []int) int {
	left, right, mid := 0, len(nums)-1, 0

	for nums[left] >= nums[right] {
		if left+1 == right {
			mid = right
			break
		}

		mid = (left + right) / 2

		// 处理无法判断中间的数字是位于前面的子数组还是后面的子数组的情况
		// 比如：数组{0,1,1,1,1}的两个旋转{1,0,1,1,1}和{1,1,1,0,1}
		if nums[left] == nums[right] && nums[left] == nums[mid] {
			min := nums[left]
			for i := left + 1; i <= right; i++ {
				if nums[i] < min {
					min = nums[i]
				}
			}
			return min
		}

		if nums[mid] >= nums[left] {
			left = mid
		} else if nums[mid] <= nums[right] {
			right = mid
		}
	}
	return nums[mid]
}