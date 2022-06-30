/**
 * @Author: zhangxuesong(joseph2018@aliyun.com)
 * @Date: 2022/6/28
 * @Description:
 */
package BackspaceCompare

func BackspaceCompare(s, t string) bool {
	ns, nt := 0, 0
	ls, lt := len(s)-1, len(t)-1
	for ls >= 0 || lt >= 0 {
		for ls >= 0 {
			if s[ls] == '#' {
				ls--
				ns++
			} else if ns > 0 {
				ls--
				ns--
			} else {
				break
			}
		}
		for lt >= 0 {
			if t[lt] == '#' {
				lt--
				nt++
			} else if nt > 0 {
				lt--
				nt--
			} else {
				break
			}
		}
		if ls >= 0 && lt >= 0 {
			if s[ls] != t[lt] {
				return false
			}
		} else if ls >= 0 || lt >= 0 {
			return false
		}

		ls--
		lt--
	}
	return true
}
