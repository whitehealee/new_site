package main
import (
"fmt"
"net/http"
"html/template"
)
func indexHendler (w http.ResponseWriter, r *http.Request){
	t, err := template.ParseFiles("tets_post_with_go/index.html")
	if err != nil {
		fmt.Fprint(w, err.Error())
	}
	t.ExecuteTemplate(w, "index", nil)
}
func thanks(w http.ResponseWriter, r *http.Request){
	t, err := template.ParseFiles("tets_post_with_go/thanks.html")
	if err != nil {
		fmt.Fprint(w, err.Error())
	}
	t.ExecuteTemplate(w, "tkanks", nil)
}
func SaveOrder (w http.ResponseWriter, r *http.Request){
	b := r.FormValue("b")
	a := r.FormValue("answer1")
	fmt.Println(b, a)
	http.Redirect(w, r, "/", 301)
}

func main()  {
	fmt.Println("Listening on port :3003")

	//http.Handle("/css/",http.StripPrefix("/css/", http.FileServer(http.Dir("/css/"))))
	http.HandleFunc("/", indexHendler)
	http.HandleFunc("/SaveOrder", SaveOrder)
	http.HandleFunc("/thanks", thanks)
	http.ListenAndServe(":3003", nil)
}