/*
 *  @Author : huangzj
 *  @Time : 2020/7/14 9:57
 *  @Description：
 */

package service

import (
	"Go-SubmeterTool/service/extra"
	"github.com/go-xorm/xorm"
	"reflect"
)

//分表工具对象
type SubMeterTool struct {
	sess           xorm.Session //数据库session管理
	engine         *xorm.Engine
	tableIndexFunc func(string, int) string    //支持Index的计算可配
	tableNameFunc  func(string, string) string //table名称的设置方法
	subMeterTable  *extra.SubMeterTable        //分表信息对象
}

//------------------------------------初始化和规则定义--------------------------------------

//初始给定的subMeterTable是nil
//程序会判断nil不进行处理，可清空后重新设置，所以其实一个SubMeterTool是支持复用的
func NewSubMeterTable(engine *xorm.Engine) *SubMeterTool {
	return &SubMeterTool{
		engine: engine,
		tableIndexFunc: func(s string, i int) string {
			return extra.GetIndex(s, i)
		},
		tableNameFunc: func(s, s1 string) string {
			return extra.MakeTable(s, s1)
		},
	}
}

//重置分表工具对象化(但是不重置Index和table名的方法)
func (ts *SubMeterTool) ClearSubMeterTable() {
	ts.subMeterTable = nil
}

//obj必须是一个ptr类型，也就是指针,否则报错(这样做主要是为了代码的统一)
func (ts *SubMeterTool) CreateSubMeterTable(tableName, pk string, Mod int, obj interface{}) {
	ts.assertPtr(obj)
	ts.subMeterTable = &extra.SubMeterTable{
		Mod:       Mod,
		Table:     tableName,
		Pk:        pk,
		StructObj: reflect.TypeOf(obj),
	}
}

// @param tableNameFunc 分表名合成规则
// @description 设置分表名合成规则
func (ts *SubMeterTool) SetMyTableNameRule(tableNameFunc func(tableName, Index string) string) {
	ts.tableNameFunc = tableNameFunc
}

func (ts *SubMeterTool) SetMyTableIndexRule(tableIndexFunc func(tableName string, Index int) string) {
	ts.tableIndexFunc = tableIndexFunc
}
