package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"time"
)

// refer url
// https://github.com/go-xorm/xorm/blob/master/README_CN.md
/***
名称映射规则
名称映射规则主要负责结构体名称到表名和结构体field到表字段的名称映射。
由core.IMapper接口的实现者来管理，xorm内置了三种IMapper实现：
core.SnakeMapper ， core.SameMapper和core.GonicMapper。
* SnakeMapper 支持struct为驼峰式命名，表结构为下划线命名之间的转换，这个是默认的Maper；
* SameMapper 支持结构体名称和对应的表名称以及结构体field名称与对应的表字段名称相同的命名；
* GonicMapper 和SnakeMapper很类似，但是对于特定词支持更好，比如ID会翻译成id而不是i_d。

当前SnakeMapper为默认值，如果需要改变时，在engine创建完成后使用

engine.SetMapper(core.SameMapper{})
同时需要注意的是：

如果你使用了别的命名规则映射方案，也可以自己实现一个IMapper。
表名称和字段名称的映射规则默认是相同的，当然也可以设置为不同，如：
engine.SetTableMapper(core.SameMapper{})
engine.SetColumnMapper(core.SnakeMapper{})
***/
type fleet struct {
	Id int

	Name string

	NameSpell string //`xorm:"name_spell"`

	Logo string

	Modifier int

	ModifyTime time.Time //`xorm:"modify_time"`

	IsDelete int //`xorm:"is_delete"`

	Description string //`xorg:"description"`
}

var engine *xorm.Engine

func main() {
	var err error
	engine, err = xorm.NewEngine("mysql", "root:nopass.123@tcp(192.168.145.5:3306)/crowdfunding")
	checkErr(err)

	// results, err := engine.Query("select * from fleet limit 5")
	// checkErr(err)
	// fmt.Println(results) // []map[string][]byte

	results1, err := engine.QueryString("select * from fleet limit 5")
	checkErr(err)
	fmt.Println(results1) // []map[string]string

	// results2, err := engine.QueryInterface("select * from fleet limit 5")
	// checkErr(err)
	// fmt.Println(results2) // []map[string]interface{}

	var ft fleet
	has, err := engine.Get(&ft)
	if has {
		fmt.Println(ft)
	}

	var ft1 fleet
	has1, _ := engine.Desc("id").Get(&ft1)
	if has1 {
		fmt.Println(ft1.Name, ft1.NameSpell)
	}
	// fleet := new(Fleet)
	// 		engine.Query("select * from fleet where id = 1")
}
func checkErr(err error) {
	if err != nil {
		fmt.Println(err.Error())
	}
}
