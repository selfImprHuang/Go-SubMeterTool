/*
 *  @Author : huangzj
 *  @Time : 2020/7/14 10:27
 *  @Description：
 */

package main

import (
	"Go-SubmeterTool/service"
	"Go-SubmeterTool/service/extra"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

const (
	dnDriver = "mysql"
	dbName   = "root:数据库连接/表名?charset=utf8mb4&parseTime=True&loc=Local"
)

func main() {
	engine, err := xorm.NewEngine(dnDriver, dbName)
	extra.CheckErr(err)
	example1 := service.NewExample1(engine)

	insertSql(example1) //插入数据库
	updateSql(example1) //更新数据
	findSql(example1)   //查询数据
}

func findSql(example1 *service.Example1) {
	fmt.Println("根据一组Id来查询玩家数据")
	result := example1.FindMapByPrimaryKeys([]string{"abvcasdas", "3213213", "4324韦尔奇若"})
	for k, v := range result {
		fmt.Println(fmt.Sprintf("主键为:%s ,用户名为： %s,用户备注为: %s", k, v.UserName, v.Comment))
		fmt.Println()
	}

	fmt.Println("根据非主键字段进行查询")
	list := example1.FindByCommonField("xyzze3432")
	for _, l := range list {
		fmt.Println(fmt.Sprintf("主键为:%s ,用户名为：%s,用户备注为: %s", l.UserId, l.UserName, l.Comment))
		fmt.Println()
	}
}

func updateSql(example1 *service.Example1) {
	example1.Update(&service.SubMeterExample{
		UserId:   "abvcasdas",
		UserName: "我是被修改后的名字",
		Comment:  "我是被修改后的备注",
		Age:      12,
	})

	result, has := example1.GetByPrimaryKey("abvcasdas")
	if !has {
		panic("查询不到对应玩家数据")
	}
	fmt.Println("更新玩家数据")
	fmt.Print(fmt.Sprintf("玩家Id:%s,玩家名称：%s，玩家备注：%s", result.UserId, result.UserName, result.Comment))
}

func insertSql(example1 *service.Example1) {
	//先删除数据才能进行新增
	example1.DeleteByPrimaryKey("abvcasdas")
	example1.DeleteByPrimaryKey("3213213")
	example1.DeleteByPrimaryKey("4324韦尔奇若")
	example1.DeleteByPrimaryKey("热舞区4123")
	example1.DeleteByPrimaryKey("rwqereqw4132")

	example1.Insert(&service.SubMeterExample{
		UserId:   "abvcasdas",
		UserName: "3241324231",
		Comment:  "发大水发斯蒂芬暗示法士大夫沙发上",
		Age:      30,
	})

	example1.Insert(&service.SubMeterExample{
		UserId:   "3213213",
		UserName: "xyzze3432",
		Comment:  "发大水发斯蒂芬暗示法士大夫沙发上",
		Age:      50,
	})

	example1.Insert(&service.SubMeterExample{
		UserId:   "4324韦尔奇若",
		UserName: "xyzze3432",
		Comment:  "发大水发斯蒂芬暗示法士大夫沙发上",
		Age:      13,
	})
	example1.Insert(&service.SubMeterExample{
		UserId:   "热舞区4123",
		UserName: "xyzze3432",
		Comment:  "发大水发斯蒂芬暗示法士大夫沙发上",
		Age:      27,
	})
	example1.Insert(&service.SubMeterExample{
		UserId:   "rwqereqw4132",
		UserName: "3241324231",
		Comment:  "发大水发斯蒂芬暗示法士大夫沙发上",
		Age:      99,
	})
}
