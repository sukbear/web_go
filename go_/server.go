package main

import (
	"net/http"
	"fmt"
	"html/template"
	"log"
)

func sayHelloName(w http.ResponseWriter,r  *http.Request){
	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Println("path",r.URL.Scheme)
	fmt.Println("scheme",r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form{
		fmt.Println("key: ",k)
		fmt.Println("val: " ,v)
	}
	fmt.Fprintf(w,"Hello astaxie!")
}

func login(w http.ResponseWriter, r *http.Request){
	fmt.Println("method: ",r.Method)
	if r.Method=="GET"{
		t, _:= template.ParseFiles("static/login.gtpl")
		logInfo := t.Execute(w,nil)
		println(logInfo)
	}else {
		//请求的是登录数据，那么执行登录的逻辑判断
		r.ParseForm()
		fmt.Println("username:", r.Form["username"])
		fmt.Println("password:", r.Form["password"])
	}
}

func main() {
	http.HandleFunc("/",sayHelloName)
	http.HandleFunc("/login",login)
	err := http.ListenAndServe(":9090",nil)
	if err != nil{
		log.Fatal("ListAndServe", err)
	}
}