/**
 * @Author: zhangxuesong(joseph2018@aliyun.com)
 * @Date: 2022/6/21
 * @Description:
 */
package BinarySearch

func Search(nums []int, target int) int {
	length := len(nums)
	if length < 1 || length > 10*10*10*10 {
		return -1
	}
	//left := 0
	//right := length - 1
	//for left <= right {
	//	middle := left + (right-left)/2
	//	switch {
	//	case nums[middle] == target:
	//		return middle
	//	case nums[middle] < target:
	//		left = middle + 1
	//	case nums[middle] > target:
	//		right = middle - 1
	//	}
	//}
	left := 0
	right := length
	for left < right {
		middle := left + (right-left)/2
		switch {
		case nums[middle] == target:
			return middle
		case nums[middle] < target:
			left = middle + 1
		case nums[middle] > target:
			right = middle - 1
		}
	}
	return -1
}
