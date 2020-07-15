/*
 *  @Author : huangzj
 *  @Time : 2020/7/13 18:02
 *  @Description：默认方法
 */

package extra

import (
	"Go-SubMeterTool/service/tool"
	"fmt"
	"strconv"
)

// @param s 需要进行取模的pk的值，正常这边好像都是字符串，如果有其他的可以补充一下方法付
// @param Mod 模
// @return 取模的结果
func GetIndex(s string, Mod int) string {
	return strconv.Itoa(tool.Crc32Mode(s, Mod))
}

func MakeTable(tableName, Index string) string {
	return fmt.Sprint(tableName, "_", Index)
}
