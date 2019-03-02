package main

import (
	"errors"
	"fmt"
	"github.com/jason-wj/bitesla/common/util/mysql/converter"
	"os"
)

//将表文件映射为通用结构体入口
func main() {
	rootPath := "/Users/su/Documents/project/go/src/github.com/jason-wj/bitesla/service/"
	if rootPath == "" {
		if len(os.Args) < 2 {
			panic(errors.New("请传入根路径！"))
		} else if len(os.Args) > 2 {
			panic(errors.New("传入参数数量过多！"))
		}
		rootPath = os.Args[1]
	}

	//contract_conf结构体生成
	generateStruct("user", rootPath+"/service-user/orm/user.go")

}

//generateStruct 生成结构体
// tableName：去掉前缀的表名，切记不要带上前缀
// path:生成结构体后要保存的路径，注意，不同结构体要起成不同的文件名，比如：/xx/xx/xx/struct1.go，否则会被覆盖
func generateStruct(tableName, path string) {
	//初始化
	t2t := converter.NewTable2Struct()
	t2t.Config(&converter.T2tConfig{
		// 如果字段首字母本来就是大写, 就不添加tag, 默认false添加, true不添加
		RmTagIfUcFirsted: false,
		// tag的字段名字是否转换为小写, 如果本身有大写字母的话, 默认false不转
		TagToLower: false,
		// 字段首字母大写的同时, 是否要把其他字母转换为小写,默认false不转换
		UcFirstOnly: false,
		//// 每个struct放入单独的文件,默认false,放入同一个文件(暂未提供)
		//SeperatFile: false,
	})

	// 开始迁移转换
	err := t2t.
		// 指定某个表,如果不指定,则默认全部表都迁移
		// 不要加前缀
		Table(tableName).
		//Table("contract_conf").
		// 表前缀
		Prefix("t_").
		// 是否添加json tag
		EnableJsonTag(true).
		// 生成struct的包名(默认为空的话, 则取名为: package model)
		PackageName("orm").
		// tag字段的key值,默认是orm
		TagKey("orm").
		// 是否添加结构体方法获取表名
		RealNameMethod("TableName").
		// 生成的结构体保存路径
		SavePath(path).
		// 数据库dsn,这里可以使用 t2t.DB() 代替,参数为 *sql.DB 对象
		Dsn("root:wuj123@tcp(localhost:3306)/db_bitesla?charset=utf8").
		// 执行
		Run()
	fmt.Println(err)
}
