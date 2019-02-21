package parsetmpl

import (
	"html/template"
	"os"
)

func checkerr(err error) {
	if err != nil {
		panic(err)
	}
}

type inventory struct {
	Material string
	Count    uint
}

type latlng struct {
	Lat float32
	Lng float32
}

func ParseString() {
	b, err := os.Getwd()
	output, err := os.OpenFile(b+"/test.html", os.O_CREATE|os.O_RDWR|os.O_TRUNC, os.ModePerm)
	checkerr(err)

	name := "shine"
	tmpl, err := template.New("test").Parse("hello,{{.}}")
	checkerr(err)
	//err = tmpl.Execute(os.Stdout,name)
	err = tmpl.Execute(output, name)
	checkerr(err)

	infos := []string{"Han Meimei", "Lilei"}
	t := template.Must(template.New("test").Parse("Students List:" +
		"{{range $index, $_ := .}}" +
		"\n{{$index}}. {{.}}," +
		"{{end}}"))
	t.Execute(os.Stdout, infos)

}

func ParseStruct() {
	s := inventory{"wool", 12}
	muban := "{{.Count}} items are made of {{.Material}} \n"
	tmpl, err := template.New("struct").Parse(muban)
	checkerr(err)
	err = tmpl.Execute(os.Stdout, s)
	checkerr(err)

	info := map[string]bool{
		"Han Meimei": true,
		"LiLei":      false,
	}
	t := template.Must(template.New("test").Parse(`Han Meimei:{{index . "Han Meimei"}}; Li Lei:{{index . "LiLei"}}`))
	t.Execute(os.Stdout, info)
}

func ParseDefineTmpl() {
	data := []template.FuncMap{}
	data = append(data, template.FuncMap{"name": "dotcoo1", "url": "http://www.dotcoo.com/", "latlng": latlng{24.1, 135.1}})
	data = append(data, template.FuncMap{"name": "dotcoo2", "url": "http://www.dotcoo.com/", "latlng": latlng{24.2, 135.2}})
	data = append(data, template.FuncMap{"name": "dotcoo2", "url": "http://www.dotcoo.com/", "latlng": latlng{24.3, 135.3}})

	datatpl := `{{range .}}{{template "user" .}}{{end}}`
	usertpl := `{{define "user"}}name:{{.name}}, url:{{.url}}, latlng:{{.latlng}} lat:{{.latlng.Lat}} lng:{{.latlng.Lng}}
{{end}}`

	tpl, err := template.New("data").Parse(datatpl)
	if err != nil {
		panic(err)
	}
	_, err = tpl.Parse(usertpl)
	if err != nil {
		panic(err)
	}

	err = tpl.Execute(os.Stdout, data)
	if err != nil {
		panic(err)
	}
}

func ParseDefineFiles() {
	data := []template.FuncMap{}
	data = append(data, template.FuncMap{"name": "dotcoo1", "url": "http://www.dotcoo.com/", "latlng": latlng{24.1, 135.1}})
	data = append(data, template.FuncMap{"name": "dotcoo2", "url": "http://www.dotcoo.com/", "latlng": latlng{24.2, 135.2}})
	data = append(data, template.FuncMap{"name": "dotcoo2", "url": "http://www.dotcoo.com/", "latlng": latlng{24.3, 135.3}})

	b, err := os.Getwd()
	output, err := os.OpenFile(b+"/embed.html", os.O_CREATE|os.O_RDWR|os.O_TRUNC, os.ModePerm)
	checkerr(err)

	tmpl, err := template.ParseFiles("layout.html", "user.html")
	checkerr(err)

	err = tmpl.Execute(output, data)
	checkerr(err)

	//// template object named  the first name of the parseFiles
	//tmpl1, err := template.New("layout.html").ParseFiles("layout.html", "user.html")
	//checkerr(err)
	//
	//err = tmpl1.Execute(os.Stdout, data)
	//checkerr(err)
}

func tfunc(i,j int) int {
	r := i + j
	return r
}

func ParseFunc() {
	tmpl := template.New("func").Funcs(template.FuncMap{"tfunc": tfunc})
	s := []string{"aaa", "bbb"}
	tmpl.Parse("list:" +
		"{{range $i,$_ := .}}" +   //循环变量 加$前缀
		"{{$j := tfunc $i $i}}\n" +   //方法调用 直接传递参数  模板内不能对变量做任何操作，只能通过调用方法接收返回值
		"返回值 {{$j}} - {{.}}" +
		"{{end}}")
	tmpl.Execute(os.Stdout, s)
}
