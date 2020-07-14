/*
 *  @Author : huangzj
 *  @Time : 2020/7/13 17:08
 *  @Description：
 */

package tool

import "reflect"

// 方法仅支持【结构体的指针对象】
func GetPtrFieldTypeMap(i interface{}) map[string]interface{} {
	if getValueType(i) != reflect.Ptr {
		panic("该方法只支持结构体指针对象")
	}

	var m map[string]interface{}
	m = make(map[string]interface{})

	ind := reflect.Indirect(reflect.ValueOf(i))
	t := ind.Type()
	for i := 0; i < t.NumField(); i++ {
		fn := t.Field(i).Name
		val := ind.Field(i).Interface()
		m[fn] = val
	}

	return m
}

func getValueType(value interface{}) reflect.Kind {
	return reflect.TypeOf(value).Kind()
}

// @return map{ "属性名"：nil }
// @description 通过传入reflect.Type获取名称-value(nil)的map结构
func GetFieldKindMap(t reflect.Type) map[string]reflect.Kind {
	kindMap := make(map[string]reflect.Kind)
	typ := t.Elem()
	for i := 0; i < typ.NumField(); i++ {
		fn := typ.Field(i).Name
		kindMap[fn] = typ.Field(i).Type.Kind()
	}
	return kindMap
}

// @param   i 结构体对象
// @return	[]string 对应的结构体字段名称
// @return  []interface 对应的结构体字段值
// @description 这个方法只支持结构体对象，但是这边没有校验（我不知道怎么校验😂）
func GetFieldNameAndValue(i interface{}) ([]string, []interface{}) {
	var fieldName []string
	var values []interface{}
	ind := reflect.Indirect(reflect.ValueOf(i))
	t := ind.Type()
	for i := 0; i < t.NumField(); i++ {
		fn := t.Field(i).Name
		val := ind.Field(i).Interface()
		fieldName = append(fieldName, fn)
		values = append(values, val)
	}

	return fieldName, values
}
