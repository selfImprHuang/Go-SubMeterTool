/*
 *  @Author : huangzj
 *  @Time : 2020/7/14 9:55
 *  @Description：分表查询，如果没有会查询旧表
 */

package service

import "Go-SubmeterTool/service/extra"

/*
 * @description 删除新表并且把旧表数据删除
 */
func (ts *SubMeterTool) DeleteByKeyWithOld(value string) {
	ts.DeleteByKey(value)
	ts.assertTableExist(ts.subMeterTable.Table)
	deleteSql := ts.deleteByKeySqlWithOld(value)
	_, err := ts.engine.Query(deleteSql)
	extra.CheckErr(err)
}

// 新表查不到数据再去旧表查,新表能查到数据就直接返回查询结果
func (ts *SubMeterTool) SelectByKeyWithOld(value string) (interface{}, error, bool) {
	v, err, has := ts.SelectByKey(value) //查询新表数据
	if has {
		return v, err, has
	}

	ts.assertTableExist(ts.subMeterTable.Table) //查询旧表数据
	return ts.selectFuc(value, ts.selectByKeySqlWithOld)
}

// 如果新分表查不到数据就从旧表中去获取数据
func (ts *SubMeterTool) SelectInKeysWithOld(value []string) ([]interface{}, error, bool) {
	v, err, has := ts.SelectInKeys(value)
	//如果全部数据都找到了就直接返回
	if len(v) == len(value) {
		return v, err, has
	}

	keyList := make([]string, 0)
	for _, row := range v {
		r := row.(map[string]interface{})
		v1, _ := r[ts.subMeterTable.Pk]
		keyList = append(keyList, v1.(string))
	}

	notFoundKeyList := make([]string, 0)
	for _, row := range value {
		if !extra.Contains(keyList, row) {
			notFoundKeyList = append(notFoundKeyList, row)
		}
	}

	vOld, errOld, hasOld := ts.selectInKeys(ts.selectInKeysSqlWithOld(notFoundKeyList))
	if hasOld || has {
		list := extra.Combine(vOld, v)
		return list, nil, true
	}

	if err != nil {
		return nil, err, false
	}
	return nil, errOld, false
}

//从新的分表和旧表中查询对应的数据
func (ts *SubMeterTool) SelectWithCommonFieldWithOld() {

}
