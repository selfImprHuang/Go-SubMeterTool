/*
 *  @Author : huangzj
 *  @Time : 2020/7/13 16:47
 *  @Description：
 */

package extra

func Contains(value []string, s string) bool {
	for _, v := range value {
		if v == s {
			return true
		}
	}
	return false
}

func Combine(vOld []interface{}, v []interface{}) []interface{} {
	list := make([]interface{}, 0)
	for _, r := range vOld {
		list = append(list, r)
	}
	for _, r := range v {
		list = append(list, r)
	}
	return list
}
