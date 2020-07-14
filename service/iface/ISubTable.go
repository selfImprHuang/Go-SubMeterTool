/*
 *  @Author : huangzj
 *  @Time : 2020/4/23 17:25
 *  @Description：
 */

package subTable

/*
 * 一开始就分表没有旧表的情况
 */
type SubOnlyNew interface {
	UpdateByKey(val interface{}, value string) interface{}

	DeleteByKey(value string)

	SelectByKey(value string) (interface{}, error, bool)

	SelectWithCommonField(value string, keyName string) ([]interface{}, error, bool)

	SelectInKeys(value []string) ([]interface{}, error, bool)
}

/*
 * 兼容旧表的分表方案
 */
type SubWithOld interface {
	DeleteByKeyWithOld(value string)

	SelectByKeyWithOld(value string) (interface{}, error, bool)

	SelectInKeysWithOld(value []string) ([]interface{}, error, bool)
}
