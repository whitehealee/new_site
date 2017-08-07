package main
import (
	"fmt"
	"net/http"
	"html/template"
	"strconv"
	_ "github.com/denisenkom/go-mssqldb"
	"database/sql"
	"crypto/rand"

	//"time"

)

func GenerateId() string{
	b := make([]byte, 10)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}
func indexHendler (w http.ResponseWriter, r *http.Request){
	t, err := template.ParseFiles("html/index.html", "html/futer.html", "html/header.html")
	if err != nil {
		fmt.Fprint(w, err.Error())
	}
	t.ExecuteTemplate(w, "index", nil)
}
func indexTovar (w http.ResponseWriter, r *http.Request){
	t, err := template.ParseFiles("html/tovar.html", "html/futer.html", "html/header.html")
	if err != nil {
		fmt.Fprint(w, err.Error())
	}
	t.ExecuteTemplate(w, "tovar", nil)
}
func indexTovar2 (w http.ResponseWriter, r *http.Request){
	t, err := template.ParseFiles("html/tovar2.html", "html/futer.html", "html/header.html")
	if err != nil {
		fmt.Fprint(w, err.Error())
	}
	t.ExecuteTemplate(w, "tovar2", nil)
}

func basket (w http.ResponseWriter, r *http.Request){
	t, err := template.ParseFiles("html/basket.html", "html/futer.html", "html/header.html")
	if err != nil {
		fmt.Fprint(w, err.Error())
	}
	t.ExecuteTemplate(w, "basket", nil)
}
func thank (w http.ResponseWriter, r *http.Request){
	t, err := template.ParseFiles("html/thank.html", "html/futer.html", "html/header.html")
	if err != nil {
		fmt.Fprint(w, err.Error())
	}
	t.ExecuteTemplate(w, "thank", nil)
}
func rozdel (w http.ResponseWriter, r *http.Request){
	t, err := template.ParseFiles("html/rozdel.html", "html/futer.html", "html/header.html")
	if err != nil {
		fmt.Fprint(w, err.Error())
	}
	t.ExecuteTemplate(w, "rozdel", nil)
}
func SaveOrder (w http.ResponseWriter, r *http.Request) {

	//TovarID :=r.FormValue("itemID")
	//TovarKl :=r.FormValue("kl_tovara")
	FName :=r.FormValue("Lname")
	name :=r.FormValue("Fname")
	Oname :=r.FormValue("Oname")
	Stret :=r.FormValue("Street")
	Home :=r.FormValue("Home")
	Flat :=r.FormValue("Flat")
	Index :=r.FormValue("index")
	Phone :=r.FormValue("Phone")
	Login :="null";
	Password := "null";
	CustumerID, _ := addCustomer(FName, name, Oname, Stret, Home, Flat, Index, Phone, Login, Password)
	Orders_ID, _ := addOrders(CustumerID, "0")

	cou :=  r.FormValue("count")
	count, _ :=strconv.Atoi(cou)
	for  i:=0; i<=count; i++{
		itemID := r.FormValue("itemID"+strconv.Itoa(i))
		TovarKl :=r.FormValue("kl_tovara"+strconv.Itoa(i))
		addBasket(Orders_ID, itemID, TovarKl)
	}
		http.Redirect(w, r, "/thank", 301)
}
func addBasket (Orders_ID, Kod_tovara, quantyty string){
	connectString := "sqlserver://sa:123Oktell321@192.168.10.6?database=shop&connection+timeout=30"
	println("Connection string=" , connectString )

	println("open connection")
	db, err := sql.Open("mssql", connectString)
	defer db.Close()
	println ("Open Error:" , err)
	if err != nil {
		fmt.Println(err)
	}
	Basket_ID := GenerateId()
	rows, err2 := db.Query("SELECT [ID], [Price] from [shop].[dbo].[Tovar] where [Kod_tovara]='"+Kod_tovara+"'")
	if err2 != nil {
		fmt.Println(err2)
	}
	rows.Next()
	var Tovar_ID, price string
	err4 := rows.Scan(&Tovar_ID, &price)
	if err4 != nil {
		fmt.Printf(err.Error())
	}
	query := "INSERT INTO [shop].[dbo].[Basket] (ID, Tovar_ID, Order_ID, quantyty, Price)"+
		"VALUES ('"+Basket_ID+"', '"+Tovar_ID+"', '"+Orders_ID+"', '"+quantyty+"', '"+price+"');"
	_, err1 := db.Query(query)

	if err1 != nil {
		fmt.Println(err1)
	}
}
func addOrders(Custumer_ID, Coment string) (string, error){
	connectString := "sqlserver://sa:123Oktell321@192.168.10.6?database=shop&connection+timeout=30"
	println("Connection string=" , connectString )

	println("open connection")
	db, err := sql.Open("mssql", connectString)
	defer db.Close()
	println ("Open Error:" , err)
	if err != nil {
		fmt.Println(err)
	}
	Orders_ID := GenerateId()

	//current_time := time.Now().Local()
	query := "INSERT INTO [shop].[dbo].[orders] (ID, Summa, Status, Kod_zakaza, Coment, Custumer_ID)"+
		"VALUES ('"+Orders_ID+"','0','0','0','"+Coment+"','"+Custumer_ID+"');"


	_, err1 := db.Query(query)

	if err1 != nil {
		fmt.Println(err1)
	}
	return Orders_ID, err1
}
func addCustomer(FName, name, Oname, Stret, Home, Flat, Index, Phone, Login, Password string) (string, error) {
	connectString := "sqlserver://sa:123Oktell321@192.168.10.6?database=shop&connection+timeout=30"
	println("Connection string=" , connectString )

	println("open connection")
	db, err := sql.Open("mssql", connectString)
	defer db.Close()
	println ("Open Error:" , err)
	if err != nil {
		fmt.Println(err)
	}

	Custumer_ID := GenerateId()


	query := "INSERT INTO [shop].[dbo].[Custumer] (ID, FName, name, Oname, Stret, Home, Flat, [Index], Phone, Login, Password)" +
		" VALUES ('"+Custumer_ID+"','"+FName+"','"+name+"','"+Oname+"','"+Stret+"','"+Home+"','"+Flat+"','"+Index+"','"+Phone+"','"+Login+"','"+Password+"');"


	fmt.Println(query)

	_, err1 := db.Query(query)

	if err1 != nil {
		fmt.Println(err1)
	}
	return Custumer_ID, err1
}
func main()  {

	fmt.Println("Listening on port :3000")
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("./css/"))))
	http.Handle("/png/", http.StripPrefix("/png/", http.FileServer(http.Dir("./png/"))))
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("./js/"))))

	//http.Handle("/css/",http.StripPrefix("/css/", http.FileServer(http.Dir("/css/"))))
	//http.HandleFunc("/", indexHendler)
	http.HandleFunc("/", indexHendler)
	http.HandleFunc("/tovar", indexTovar)
	http.HandleFunc("/1", rozdel)
	http.HandleFunc("/7", indexTovar)
	http.HandleFunc("/SaveOrder", SaveOrder)
	http.HandleFunc("/basket", basket)
	http.HandleFunc("/thank", thank)

	http.ListenAndServe(":3000", nil)
}