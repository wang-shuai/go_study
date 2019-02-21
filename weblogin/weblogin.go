package main
import(
	"fmt"
	//"strings"
	"net/http"
	"log"
	"html/template"
)

func sayHelloName(w http.ResponseWriter,r *http.Request){
	r.ParseForm()
	for k,_ := range r.Form{
		fmt.Println("key:",k)
		//fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprint(w,"hello login")
}

func login(w http.ResponseWriter,r *http.Request){

	//t, _ := template.New("foo").Parse(`{{define "T"}}Hello, {{.}}!{{end}}`)
	//t.ExecuteTemplate(w, "T", "<script>alert('you have been pwned')</script>")
	//
	//t2,_ := template.New("foo2").Parse(`{{define "T"}}Hello, {{.}}!{{end}}`)
	//t2.ExecuteTemplate(w, "T", template.HTML("<script>alert('you have been pwned')</script>"))

	fmt.Println("method:"+r.Method)
	r.ParseForm()
	if(r.Method=="GET"){
		t,err := template.ParseFiles("login.gtpl")
		if err!=nil{
			fmt.Println(err)
		} else{
			t.Execute(w,nil)
		}
	}else{
		fmt.Println("post")

		fmt.Println("username:",r.Form["username"])
		fmt.Println("password:",template.HTMLEscapeString(r.Form.Get("password")))
		b := []byte(r.Form.Get("username"))
		template.HTMLEscape(w,b)
	}
}

func main(){
	http.HandleFunc("/",sayHelloName)
	http.HandleFunc("/login",login)
	err := http.ListenAndServe(":9900",nil)
	if err != nil{
		log.Fatal("listenAndServe:",err)
	}
}