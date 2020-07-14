/*
 *  @Author : huangzj
 *  @Time : 2020/7/13 16:17
 *  @Description：
 */

package extra

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}
