/*
 *  @Author : huangzj
 *  @Time : 2020/7/13 16:48
 *  @Description：通过sql查询结果处理工具
 */

package service

import (
	"Go-SubmeterTool/service/extra"
	"Go-SubmeterTool/service/tool"
)

func (ts *SubMeterTool) selectFuc(value string, selectSqlFunc func(string) string) (interface{}, error, bool) {
	selectResultMap, err := ts.engine.QueryInterface(selectSqlFunc(value))
	if err != nil || selectResultMap == nil {
		return nil, err, false
	}

	fnv := tool.GetFieldKindMap(ts.subMeterTable.StructObj) //因为这边得到的是byte数组，所以需要进行一次转换
	for k, v := range selectResultMap[0] {
		s := extra.ToStr(v) //先把字节数组转换成字符串
		t, ok := fnv[k]     //再通过参数的类型，讲转换出来的字符串进行对应类型的转换
		if !ok {
			continue
		}
		result := extra.ToVal(s, t)
		selectResultMap[0][k] = result
	}

	return selectResultMap[0], nil, true
}

func (ts *SubMeterTool) selectInKeys(sql string) ([]interface{}, error, bool) {
	m, err := ts.engine.QueryInterface(sql)
	if err != nil || m == nil {
		return nil, err, false
	}

	var list []interface{}
	for key, value := range m {
		fnv := tool.GetFieldKindMap(ts.subMeterTable.StructObj) //因为这边得到的是byte数组，所以需要进行一次转换
		for k, v := range value {
			s := extra.ToStr(v) //先把字节数组转换成字符串
			t, ok := fnv[k]     //再通过参数的类型，讲转换出来的字符串进行对应类型的转换
			if !ok {
				continue
			}
			result := extra.ToVal(s, t)
			m[key][k] = result
		}
		list = append(list, value)
	}

	return list, nil, true
}
