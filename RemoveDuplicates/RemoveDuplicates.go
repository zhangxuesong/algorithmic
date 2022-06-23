/**
 * @Author: zhangxuesong(joseph2018@aliyun.com)
 * @Date: 2022/6/23
 * @Description:
 */
package RemoveDuplicates

func RemoveDuplicates(nums []int) int {
	slow := 1
	for fast := 1; fast < len(nums); fast++ {
		if nums[fast] != nums[fast-1] {
			nums[slow] = nums[fast]
			slow++
		}
	}
	return slow
}
