 package DBInsert

import (
    "database/sql"
    "fmt"
    "log"
    "strings"
    _ "github.com/go-sql-driver/mysql"
    _"github.com/mattn/go-sqlite3"
    
)

type SQLInsert struct {
    serverHostName, field, information string
    Database *sql.DB
}

func UpdateProcessor(serverHostName string, Database *sql.DB, information string){
	IncertAttribute(serverHostName, Database, "processor", information)
}

/***********************************************************************
* this incerts data into the database
* Returns: string of data in the 
* 
***********************************************************************/
func UpdateIP(serverHostName string, Database *sql.DB, information string){
	IncertAttribute(serverHostName, Database, "IP", information)
}

/***********************************************************************
* this incerts data into the database
* Returns: string of data in the 
* 
***********************************************************************/
func IncertAttribute(serverHostName string, Database *sql.DB, field string, information string)  {   
	
	
	fmt.Println("make change")
	SQLstring := strings.Replace("UPDATE serverInformation SET field=? WHERE Hostname='#'","#" ,serverHostName, 1)
	SQLstring = strings.Replace(SQLstring,"field", field, 1 )
	fmt.Println(SQLstring)
	stmt, err := Database.Prepare(SQLstring)
	if err != nil {
		fmt.Println("incert error 0")
		log.Fatal(err)
		
	}
	res, err := stmt.Exec(information)
	if err != nil {
		fmt.Println("incert error 1")
		log.Fatal(err)

	}
	lastId, err := res.LastInsertId()
	if err != nil {
		fmt.Println("incert error 2")
		log.Fatal(err)
		
	}
	rowCnt, err := res.RowsAffected()
	if err != nil {
		fmt.Println("incert error 3")
		log.Fatal(err)	
	}
	log.Printf("ID = %d, affected = %d\n", lastId, rowCnt)
	
}
