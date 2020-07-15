/*
 *  @Author : huangzj
 *  @Time : 2020/3/24 16:23
 *  @Description：该方法只查询新表
 */

package service

import (
	"Go-SubMeterTool/service/extra"
	"Go-SubMeterTool/service/tool"
	"errors"
	"fmt"
	"reflect"
)

var tableNameList []string //存储表名的数组，可以存取一步查询

// 需要更新的对象的值，可以传入对应设置的结构体的对象，正常情况下也应该支持传入map才对，但是现在还不知道怎么处理，妈耶
// 主键的值.
func (ts *SubMeterTool) UpdateByKey(val interface{}, value string) interface{} {
	ts.assertPtr(val)
	//检查subMeterTable对象是否存在，因为这个对象可以清空然后重新设置，所以为了防止出错这边需要检查一下
	ts.subMeterParamValid()
	ts.subMeterTable.Index = ts.tableIndexFunc(value, ts.subMeterTable.Mod)
	ts.subMeterTable.SubTableName = ts.tableNameFunc(ts.subMeterTable.Table, ts.subMeterTable.Index)
	ts.assertTableExist(ts.subMeterTable.SubTableName)

	//判断传进来的map还是对应的结构体数据，如果都不是需要报错
	if reflect.TypeOf(val) == ts.subMeterTable.StructObj {
		//如果当前表有这个数据就进行更新
		if _, _, has := ts.SelectByKey(value); has {
			updateByKeySql := ts.updateByKeySql(val, value) //更新语句sql拼接
			_, err := ts.Sess.Query(updateByKeySql)
			extra.CheckErr(err)
			return val
		}

		err := ts.InsertOne(val, value) //如果当前表没有数据就进行插入
		extra.CheckErr(err)
		return val
	}

	panic("传入对象类型对不上:" + reflect.TypeOf(val).String())
}

// 删除新分表的数据，不删除原来旧表的数据
func (ts *SubMeterTool) DeleteByKey(value string) {
	ts.subMeterParamValid()
	ts.subMeterTable.Index = ts.tableIndexFunc(value, ts.subMeterTable.Mod)
	ts.subMeterTable.SubTableName = ts.tableNameFunc(ts.subMeterTable.Table, ts.subMeterTable.Index)
	deleteSql := ts.deleteByKeySql(value)
	_, err := ts.Sess.Query(deleteSql)
	extra.CheckErr(err)
}

/*
 * @param value 主键值
 * @return 数据，错误，是否存在
 * @description 通过主键获取到对应的属性值
 */
func (ts *SubMeterTool) SelectByKey(value string) (interface{}, error, bool) {
	ts.subMeterParamValid()
	ts.subMeterTable.Index = ts.tableIndexFunc(value, ts.subMeterTable.Mod)
	ts.subMeterTable.SubTableName = ts.tableNameFunc(ts.subMeterTable.Table, ts.subMeterTable.Index)
	ts.assertTableExist(ts.subMeterTable.SubTableName) //校验表是否存在
	return ts.selectFuc(value, ts.selectByKeySql)
}

// 根据不是主键的字段进行查询，不会有多个数据的情况下，需要组合所有的查询结果
func (ts *SubMeterTool) SelectWithCommonField(value string, keyName string) ([]interface{}, error, bool) {
	ts.subMeterParamValid()
	selectResultMap, err := ts.Sess.QueryInterface(ts.selectSqlByField(value, keyName, ts.getAllSubMeterTable()))
	if err != nil || selectResultMap == nil {
		return nil, err, false
	}

	var list []interface{}

	for key, value := range selectResultMap {
		fnv := tool.GetFieldKindMap(ts.subMeterTable.StructObj) //因为这边得到的是byte数组，所以需要进行一次转换
		for k, v := range value {
			s := extra.ToStr(v) //先把字节数组转换成字符串
			t, ok := fnv[k]     //再通过参数的类型，讲转换出来的字符串进行对应类型的转换
			if !ok {
				continue
			}
			result := extra.ToVal(s, t)
			selectResultMap[key][k] = result
		}
		list = append(list, value)
	}

	return list, nil, true
}

//根据一组主键查询结果.
func (ts *SubMeterTool) SelectInKeys(value []string) ([]interface{}, error, bool) {
	ts.subMeterParamValid()
	tableMap := ts.getTableValueMap(value) //通过主键拼接 表名-主键数组的map数据
	return ts.selectInKeys(ts.selectInKeysSql(tableMap))
}

// @param i 结构体对象
// @param value 传入主键的值，这边不计算，感觉传入更好
// @description 分表的新增方法
func (ts *SubMeterTool) InsertOne(structObj interface{}, value string) error {
	ts.assertPtr(structObj)
	ts.subMeterParamValid()
	ts.subMeterTable.Index = ts.tableIndexFunc(value, ts.subMeterTable.Mod)
	ts.subMeterTable.SubTableName = ts.tableNameFunc(ts.subMeterTable.Table, ts.subMeterTable.Index)
	if reflect.TypeOf(structObj) == ts.subMeterTable.StructObj {
		fn, val := tool.GetFieldNameAndValue(structObj) //获取结构体的属性

		//这个地方如果用表注册的方式可以省去一次查询数据库的消耗,这边还有一步就是给table属性赋值
		ts.assertTableExist(ts.subMeterTable.SubTableName)
		insertSql := ts.insertSql(fn, val)
		_, err := ts.Sess.Query(insertSql)
		extra.CheckErr(err)
		return nil
	}
	return errors.New(fmt.Sprintf("传入的对象类型和设置的类型不一致：%s", reflect.TypeOf(structObj).String()))
}

//根据一组key进行删除
//@param keys 一组key的值
//@param keyName primary key的字段名
func (ts *SubMeterTool) DeleteByKeys(keys []string, keyName string) {
	ts.subMeterParamValid()
	tableValueMap := ts.getTableValueMap(keys)
	sqlList := ts.deleteInKeys(tableValueMap, keyName)
	for _, sql := range sqlList {
		_, err := ts.Sess.Query(sql)
		extra.CheckErr(err)
	}
}
