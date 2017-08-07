package main
import (
	// Import go-mssqldb strictly for side-effects
	_ "github.com/denisenkom/go-mssqldb"
	"database/sql"
	"log"
	"fmt"
)


func main() {


	//var n_tables string

	println (sql.Drivers())

	// URL connection string formats
	//    sqlserver://sa:mypass@localhost?database=master&connection+timeout=30         // username=sa, password=mypass.
	//    sqlserver://sa:my%7Bpass@somehost?connection+timeout=30                       // password is "my{pass"
	// note: pwd is "myP@55w0rd"
	connectString := "sqlserver://sa:123Oktell321@192.168.10.6?database=test_upk_date&connection+timeout=30"
	println("Connection string=" , connectString )

	println("open connection")
	db, err := sql.Open("mssql", connectString)
	defer db.Close()
	println ("Open Error:" , err)
	if err != nil {
		log.Fatal(err)
	}

	println("count records in TS_TABLES & scan")


	rows, err := db.Query("SELECT TOP 1000 [_LineNo271],[_Fld272] FROM [test_upk_date].[dbo].[_Reference16_VT270]")
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		var source, content string
		err := rows.Scan(&source, &content)
		if err != nil {
			fmt.Printf(err.Error())
		}
		fmt.Println(source, "   ", content)
	}
	println("closing connection")
	db.Close()
}