/*
 *  @Author : huangzj
 *  @Time : 2020/7/14 10:05
 *  @Description：拓展工具
 */

package service

import (
	"fmt"
	"strconv"
)

//获取 表名 -> [值1,值2，值3] 的数据结构，用于后面拼接in的操作
func (ts *SubMeterTool) getTableValueMap(keys []string) map[string][]string {
	mapList := make(map[string][]string)
	for _, s := range keys {
		index := ts.tableIndexFunc(s, ts.subMeterTable.Mod)
		tableName := ts.tableNameFunc(ts.subMeterTable.Table, index)
		ts.assertTableExist(tableName)
		v, ok := mapList[tableName]

		if ok {
			v = append(v, s) //把对应的key的值添加到集合
			mapList[tableName] = v
		} else {
			mapList[tableName] = []string{s}
		}
	}

	return mapList

}

//	根据Mod返回所有分表的名称数组
func (ts *SubMeterTool) getAllSubMeterTable() []string {
	ts.subMeterParamValid()
	val := 1
	tableNames := make([]string, 0)
	for ; val <= ts.subMeterTable.Mod; val++ {
		name := fmt.Sprint(ts.subMeterTable.Table, "_", strconv.Itoa(val))
		tableNames = append(tableNames, name)
		ts.assertTableExist(name) //校验表是否存在
	}

	return tableNames
}
