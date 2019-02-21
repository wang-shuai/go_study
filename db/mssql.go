package main

import (
	_ "github.com/denisenkom/go-mssqldb"
	"github.com/go-xorm/xorm"
	"fmt"
	json "github.com/pquerna/ffjson/ffjson"
	"github.com/go-xorm/core"
	"io/ioutil"
	"os"
	"bytes"
)

var (
	ucar_engine *xorm.Engine // ucar数据库
	keys        = []int{1, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 25, 27, 28, 29, 30}
)

type item struct {
	Name  string `json:"name"`
	Id    int    `json:"id"`
	Spell string `json:"spell"`
}

type area struct {
	CityId int
	PinYin string
}

func main() {

	dic := make(map[string][]item)
	//items := make([]item, 0)
	//i1 := item{"北京", 201, "beijing"}
	//i2 := item{"上海", 2401, "shanghai"}
	//items = append(items, i1, i2)
	//dic["1"] = items //dt{items}
	//dic["2"] = items //dt{items}
	//fmt.Println(dic)
	//
	//byt, _ := json.Marshal(dic)
	//fmt.Println(string(byt))
	//
	//dic1 := make(map[string][]item)
	//json.Unmarshal(byt,&dic1)
	//fmt.Println(dic1)

	if Eg, err := xorm.NewEngine("mssql", `odbc:server=192.168.156.225;user id=zhaogx;password={!QAZxsw2};database=ucar;connection timeout=180`); err != nil {
		fmt.Println(err)
		panic("ucar数据库链接失败")
	} else {
		ucar_engine = Eg
	}
	ucar_engine.SetMapper(core.SameMapper{}) //与字段、表名一致  不区分大小写
	ucar_engine.ShowSQL(true)                //展示每次执行的sql

	dir ,err:= os.Getwd()
	file, err := ioutil.ReadFile(dir+"/src/db/areadata.json")

	if err != nil {
		fmt.Println("打开文件失败", err)
	}
	file = bytes.TrimPrefix(file, []byte("\xef\xbb\xbf")) // Or []byte{239, 187, 191}
	err = json.Unmarshal(file, &dic)
	if err !=nil{
		fmt.Println(err)
	}
	fmt.Println(dic)

	allarea := make([]area,0)
	ucar_engine.SQL("select * from AreaCity where ParentId > 0 ").Find(&allarea)

	for k,_ := range dic{
		itemlist := dic[k]
		for i,_ := range itemlist{
			for _,v := range allarea{
				if v.CityId==itemlist[i].Id{
					itemlist[i].Spell = v.PinYin
				}
			}
		}
	}
	b,_:=json.Marshal(dic)
	fmt.Println(string(b))
}
