package dialect

import "reflect"

var dialectsMap = map[string]Dialect{}

//这一层其实是为了屏蔽 不同数据库的差异
type Dialect interface {
	//DataTypeOf 用于将 Go 语言的类型转换为该数据库的数据类型。
	DataTypeOf(typ reflect.Value) string
	//TableExistSQL 返回某个表是否存在的 SQL 语句，参数是表名(table)。
	TableExistSQL(tableName string) (string, []interface{})
}

//声明了 RegisterDialect 和 GetDialect 两个方法用于注册和获取 dialect 实例。
//如果新增加对某个数据库的支持，
//那么调用 RegisterDialect 即可注册到全局。
func RegisterDialect(name string, dialect Dialect) {
	dialectsMap[name] = dialect
}

func GetDialect(name string) (dialect Dialect, ok bool) {
	dialect, ok = dialectsMap[name]
	return
}
