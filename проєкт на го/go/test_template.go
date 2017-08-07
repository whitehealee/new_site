package main

import (
	"html/template"
	"fmt"
	"net/http"
)
var v = make(map[int]*Post)
type Post struct{
	Id	string
	Price	string
	Name	string
}

func addPost(id, price, name string) *Post  {
	return &Post{id, price, name}
}
func indexHendler (w http.ResponseWriter, r *http.Request){

	t, err := template.ParseFiles("проєкт на го/go/index.html")
	if err != nil {
		fmt.Fprint(w, err.Error())
	}
	fmt.Println(v[1].Name)
	for i, g:= range v {
		fmt.Println("v[",i,"]=", g.Id)

	}
	t.ExecuteTemplate(w, "index", v)
}
func main()  {
	//v := make(map[int]*Post)
	k:=addPost("1","name_v1","3")
	v[1]=k
	k=addPost("2","name_v2","2")
	v[2]=k

	fmt.Println("Listening on port :3006")

	//http.Handle("/css/",http.StripPrefix("/css/", http.FileServer(http.Dir("/css/"))))
	//http.HandleFunc("/", indexHendler)
	http.HandleFunc("/", indexHendler)

	http.ListenAndServe(":3006", nil)
}