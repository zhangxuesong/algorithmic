/**
 * @Author: zhangxuesong(joseph2018@aliyun.com)
 * @Date: 2022/6/22
 * @Description:
 */
package RemoveElement

func RemoveElement(nums []int, val int) int {
	//slow := 0
	//for _, v := range nums {
	//	if v != val {
	//		nums[slow] = v
	//		slow++
	//	}
	//}
	//return slow
	left := 0
	right := len(nums)
	for left < right {
		if nums[left] == val {
			nums[left] = nums[right-1]
			right--
		} else {
			left++
		}
	}
	return left
}
