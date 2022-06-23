/**
 * @Author: zhangxuesong(joseph2018@aliyun.com)
 * @Date: 2022/6/23
 * @Description:
 */
package MoveZeroes

func MoveZeroes(nums []int) {
	slow := 0
	for fast := 0; fast < len(nums); fast++ {
		if nums[fast] != 0 {
			nums[slow], nums[fast] = nums[fast], nums[slow]
			slow++
		}
	}
}
