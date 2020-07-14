/*
 *  @Author : huangzj
 *  @Time : 2020/7/13 17:44
 *  @Description：检验工具
 */

package service

import (
	"fmt"
	"reflect"
)

func (ts *SubMeterTool) assertTableExist(table string) {
	for _, t := range tableNameList {
		if t == table {
			return //找到对应的表名直接返回
		}
	}
	//db查询判断表是否存在
	dec, notExist := ts.engine.Query(fmt.Sprint("show tables like ", "'", table, "'"))
	if notExist != nil || dec == nil {
		panic(fmt.Sprintf("%s 表不存在，报错信息%v", table, notExist))
	}

	tableNameList = append(tableNameList, table)
}

//分表对象的属性校验，不存在报错
func (ts *SubMeterTool) subMeterParamValid() {
	if ts.subMeterTable == nil {
		panic("请先设置分表属性")
	}
}

func (ts *SubMeterTool) assertPtr(obj interface{}) {
	if reflect.TypeOf(obj).Kind() != reflect.Ptr {
		panic("请传入相应指针类型")
	}
}
