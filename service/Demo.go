/*
 *  @Author : huangzj
 *  @Time : 2020/4/24 16:32
 *  @Description：
 */
package service

import (
	"Go-SubMeterTool/service/extra"
	"encoding/json"
	"github.com/go-xorm/xorm"
)

func doDemo() {

}

type Example1 struct {
	subMeterTool *SubMeterTool //分表工具对象
}

type SubMeterExample struct {
	UserId   string `xorm:"pk not null varchar(32)"` //
	UserName string `xorm:"varchar(32)"`             //
	Comment  string `xorm:"varchar(255)"`            //
	Age      int    `xorm:"int(11)"`
}

/*
 * 1.创建一个分表对象
 */
func NewExample1(sess *xorm.Session) *Example1 {
	//创建分表工具对象
	tool := NewSubMeterTable(sess)
	//创建分表通用参数
	tool.CreateSubMeterTable("SubMeterExample", "UserId", 3, &SubMeterExample{})

	//直接返回一个携戴分表工具对象的原有对象数据,这样就不用改变原来的对象
	srv := &Example1{
		tool,
	}

	return srv
}

func (srv *Example1) Commit() {
	srv.subMeterTool.Commit()
}

/**
 * 通过主键获取一条分表数据
 * 这边通过json的序列化方法进行初始化成相应对象，这个后期还是可以优化一下的
 */
func (srv *Example1) GetByPrimaryKey(primaryKey string) (*SubMeterExample, bool) {
	var result SubMeterExample

	record, err, has := srv.subMeterTool.SelectByKey(primaryKey)
	extra.CheckErr(err)
	r, err := json.Marshal(record)
	uErr := json.Unmarshal(r, &result)
	if err != nil || uErr != nil {
		return nil, false
	}

	return &result, has

}

/**
 * 新增对象，只会在新的分表进行新增，不会再处理原来的旧表(单表)
 */
func (srv *Example1) Insert(entity *SubMeterExample) {
	err := srv.subMeterTool.InsertOne(entity, entity.UserId)
	extra.CheckErr(err)
}

/*
 * 更新数据，如果新的分表有数据进行数据更新，否则插入这条数据
 */
func (srv *Example1) Update(record *SubMeterExample) {
	srv.subMeterTool.UpdateByKey(record, record.UserId)
}

/*
 * 通过一组主键值查询对应的对象集合，这边的例子如果查询分表找不到会去旧表进行查询
 * 如果没有旧表直接把SelectInKeysWithOld 替换成SelectInKeys
 */
func (srv *Example1) FindMapByPrimaryKeys(rids []string) map[string]*SubMeterExample {

	val, err, _ := srv.subMeterTool.SelectInKeysWithOld(rids)
	extra.CheckErr(err)

	var retMap = make(map[string]*SubMeterExample, 0)
	for _, rec := range val {
		var result SubMeterExample
		r, err := json.Marshal(rec)
		uErr := json.Unmarshal(r, &result)
		if err == nil && uErr == nil {
			retMap[result.UserId] = &result
		}
	}

	return retMap
}

/**
 * 通过非主键的数据库字段值进行查找
 * 这边需要传入对应的字段名称。
 */
func (srv *Example1) FindByCommonField(commonField string) []*SubMeterExample {

	var records []*SubMeterExample

	val, err, _ := srv.subMeterTool.SelectWithCommonField(commonField, "UserName")
	extra.CheckErr(err)
	for _, rec := range val {
		var result SubMeterExample
		r, err := json.Marshal(rec)
		uErr := json.Unmarshal(r, &result)
		if err == nil && uErr == nil {
			records = append(records, &result)
		}
	}
	return records
}

func (srv *Example1) DeleteByPrimaryKey(key string) {
	srv.subMeterTool.DeleteByKey(key)
}

func (srv *Example1) DeleteByKeys(keys []string) {
	srv.subMeterTool.DeleteByKeys(keys, "UserId")
}
