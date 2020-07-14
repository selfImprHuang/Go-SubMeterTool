/*
 *  @Author : huangzj
 *  @Time : 2020/7/13 16:29
 *  @Description：sql拼接工具
 */

package service

import (
	"Go-SubmeterTool/service/extra"
	"Go-SubmeterTool/service/tool"
	"fmt"
	"strings"
)

func (ts *SubMeterTool) updateByKeySql(val interface{}, value string) string {
	fnv := tool.GetPtrFieldTypeMap(val)
	s := fmt.Sprintf("update %s set ", ts.subMeterTable.SubTableName)
	var kvString string
	i := 1
	for k, v := range fnv {
		//primary key不会被修改
		if k != ts.subMeterTable.Pk {
			kvString = fmt.Sprint(kvString, " ", k, " = ", "'", extra.ToStr(v), "'")
			if i != len(fnv)-1 {
				kvString = fmt.Sprint(kvString, ",")
			}
			i++
		}
	}

	result := fmt.Sprint(s, kvString, " where ", ts.subMeterTable.Pk, " = ", "'", value, "'")
	return result
}

func (ts *SubMeterTool) selectByKeySql(value string) string {
	return fmt.Sprintf("select * from %s where %s = %s%s%s", ts.subMeterTable.SubTableName, ts.subMeterTable.Pk, "'", value, "'")
}

func (ts *SubMeterTool) deleteByKeySql(value string) string {
	return fmt.Sprint("delete from ", ts.subMeterTable.SubTableName, " where ", ts.subMeterTable.Pk, " = '", value, "'")
}

func (ts *SubMeterTool) insertSql(fn []string, val []interface{}) string {
	var fs string
	var vs string

	is := fmt.Sprintf("insert into %s (", ts.subMeterTable.SubTableName)
	//拼接key
	for i, v := range fn {
		fs = fmt.Sprint(fs, v)

		if i != len(fn)-1 {
			fs = fmt.Sprint(fs, ",")
		}
	}
	//拼接value
	for i, v := range val {
		vs = fmt.Sprint(vs, "'", extra.ToStr(v), "'")
		if i != len(val)-1 {
			vs = fmt.Sprint(vs, ",")
		}
	}
	return fmt.Sprint(is, fs, " )", " values( ", vs, " )")
}

func (ts *SubMeterTool) selectByKeySqlWithOld(value string) string {
	return fmt.Sprintf("select * from %s where %s = %s%s%s ", ts.subMeterTable.Table, ts.subMeterTable.Pk, "'", value, "'")
}

func (ts *SubMeterTool) selectSqlByField(value string, fieldName string, table []string) interface{} {
	var selectSql string //默认值是""

	for Index, val := range table {
		s := fmt.Sprintf("select * from %s where %s = %s%s%s ", val, fieldName, "'", value, "'")

		if Index != 0 {
			selectSql = fmt.Sprint(selectSql, "union all ", s)
		} else {
			selectSql = fmt.Sprint(selectSql, s)
		}
	}

	return selectSql
}

func (ts *SubMeterTool) selectInKeysSql(tableMap map[string][]string) string {
	var inSql string
	Index := 0

	for key, value := range tableMap {
		selectSql := fmt.Sprintf("select * from  %s where %s in( ", key, ts.subMeterTable.Pk)
		var fs string
		for i, s := range value {
			fs = fmt.Sprint(fs, "'", s, "'")
			if i != len(value)-1 {
				fs = fmt.Sprint(fs, ",")
			}
		}
		selectSql = fmt.Sprint(selectSql, fs, " )")
		if Index == 0 {
			inSql = fmt.Sprint(inSql, selectSql)
		} else if Index < len(tableMap)-1 {
			inSql = fmt.Sprint(inSql, "union all ", selectSql)
		}
		Index++
	}

	return inSql
}

func (ts *SubMeterTool) selectInKeysSqlWithOld(value []string) string {
	selectSql := fmt.Sprintf("select * from %s where %s in( ", ts.subMeterTable.Table, ts.subMeterTable.Pk)
	var fs string
	for i, s := range value {
		fs = fmt.Sprint(fs, "'", s, "'")
		if i != len(value)-1 {
			fs = fmt.Sprint(fs, ",")
		}
	}
	return strings.Join([]string{selectSql, fs, ")"}, "")
}

func (ts *SubMeterTool) deleteByKeySqlWithOld(value string) string {
	return fmt.Sprint("delete from ", ts.subMeterTable.Table, " where ", ts.subMeterTable.Pk, " = '", value, "'")
}

func (ts *SubMeterTool) deleteInKeys(tableValueMap map[string][]string, name string) string {
	var deleteSql string

	for tableName, values := range tableValueMap {
		if len(values) == 0 {
			continue
		}
		var list string
		for i, r := range values {
			list = fmt.Sprintf(" ' %s '", r)
			if i != len(values)-1 {
				list = fmt.Sprint(list, ",")
			}
		}
		deleteSql = fmt.Sprintf(deleteSql, "delete from %s where %s in ( %s );", tableName, name, list)
	}

	return deleteSql
}
