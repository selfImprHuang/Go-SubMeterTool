/*
 *  @Author : huangzj
 *  @Time : 2020/7/13 16:05
 *  @Description：分表结构体对象
 */

package extra

import (
	"reflect"
)

//分表结构体
type SubMeterTable struct {
	Mod          int          //分表数量
	Table        string       //统一的分表名(比如说分表规则xx_1\xx_2，这个时候存的就是xx)
	Pk           string       //primary key
	Index        string       //表后缀下标
	SubTableName string       //真正的分表表名
	StructObj    reflect.Type //需要转换的对象类型
}
