/*
 *  @Author : huangzj
 *  @Time : 2020/3/25 12:01
 *  @Description：
 */

package extra

import (
	"encoding/json"
	"reflect"
	"strconv"
	"time"
)

/*
 * @param value 入参
 * @return string 转换成字符串的结果
 * @description 这个方法是网上找的，大概看了一下，对应到的类型基本都是我们需要的
 */
func ToStr(value interface{}) string {
	var key string
	if value == nil {
		return key
	}

	switch value.(type) {
	case float64:
		ft := value.(float64)
		key = strconv.FormatFloat(ft, 'f', -1, 64)
	case float32:
		ft := value.(float32)
		key = strconv.FormatFloat(float64(ft), 'f', -1, 64)
	case int:
		it := value.(int)
		key = strconv.Itoa(it)
	case uint:
		it := value.(uint)
		key = strconv.Itoa(int(it))
	case int8:
		it := value.(int8)
		key = strconv.Itoa(int(it))
	case uint8:
		it := value.(uint8)
		key = strconv.Itoa(int(it))
	case int16:
		it := value.(int16)
		key = strconv.Itoa(int(it))
	case uint16:
		it := value.(uint16)
		key = strconv.Itoa(int(it))
	case int32:
		it := value.(int32)
		key = strconv.Itoa(int(it))
	case uint32:
		it := value.(uint32)
		key = strconv.Itoa(int(it))
	case int64:
		it := value.(int64)
		key = strconv.FormatInt(it, 10)
	case uint64:
		it := value.(uint64)
		key = strconv.FormatUint(it, 10)
	case string:
		key = value.(string)
	case []byte:
		key = string(value.([]byte))
	case time.Time:
		//时间这个应该设置成可配的。
		it := value.(time.Time)
		it.Format(time.RFC3339)
	default:
		newValue, _ := json.Marshal(value)
		key = string(newValue)
	}

	return key
}

func ToVal(s string, value reflect.Kind) interface{} {
	if value == reflect.Invalid {
		return s
	}

	var ft interface{}
	var err error

	switch value {
	case reflect.Float64:
		ft, err = strconv.ParseFloat(s, 64)
	case reflect.Float32:
		ft, err = strconv.ParseFloat(s, 32)
	case reflect.Int:
		ft, err = strconv.Atoi(s)
	case reflect.Int8:
		ft, err = strconv.ParseInt(s, 10, 8)
	case reflect.Int16:
		ft, err = strconv.ParseInt(s, 10, 16)
	case reflect.Int32:
		ft, err = strconv.ParseInt(s, 10, 32)
	case reflect.Int64:
		ft, err = strconv.ParseInt(s, 10, 64)
	case reflect.String:
		ft = s
	case reflect.Array:
		ft = []byte(s)
	default:
		newValue, _ := json.Marshal(value)
		ft = string(newValue)
	}
	if err != nil {
		panic("转换失败" + err.Error())
	}

	return ft
}
