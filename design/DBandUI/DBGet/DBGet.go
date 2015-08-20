 package DBGet

import (
    "database/sql"
    "fmt"
    "log"
    "strings"
    _ "github.com/go-sql-driver/mysql"
    _"github.com/mattn/go-sqlite3"  
    
)
type SQLQuery struct {
    serverHostName, field string
    Database *sql.DB
}

type Row struct {
	servername string 
	availability string
	BookedUntill, user, MacAddress, IP, processor, cores, RAM, DiskSpace,NetCard,Comments string
}

/***********************************************************************
* this gets the Availblity of the server from the database
* Returns: bool true if it is in use
* 
***********************************************************************/

func GetAvailblity(serverHostName string, Database *sql.DB) string { 
	isAvailable := GetAttribute(serverHostName, Database, "Availability")
	if isAvailable == "true" || isAvailable == "True"{
		return "Not Booked"
	}else{
		return "Booked"
	}
}


/***********************************************************************
* this gets the current user of the server from the database
* Returns: string of the user. 
* 
***********************************************************************/
func GetUser(serverHostName string, Database *sql.DB) string {   
	
	return GetAttribute(serverHostName, Database, "User")
}

/***********************************************************************
* this quearies the databace based on what attrbuite it was passed in
* Returns: string of data in the 
* 
***********************************************************************/
func GetAttribute(serverHostName string, Database *sql.DB, attribute string) string {   
	
    rows, err := Database.Query(strings.Replace("SELECT ? FROM serverInformation WHERE Hostname=?", "?", attribute, 1), serverHostName)
    if err != nil {
            log.Fatal(err)
    }
    var user string
    defer rows.Close()
    for rows.Next() {
            if err := rows.Scan(&user); err != nil {
                    log.Fatal(err)
            }
            //fmt.Printf("curent User is %s for %s\n", user, serverHostName)
    }
    if err := rows.Err(); err != nil {
            log.Fatal(err)
    }  
    return user
}



/***********************************************************************
* this quearies the databace to get one full row based on the server
* hostname
* Returns: string  
* 
***********************************************************************/
/*func GetRow(serverHostName string, Database *sql.DB)Row{
	row, err := Database.Query("SELECT * FROM serverInformation WHERE Hostname=?", serverHostName)
	if err != nil {
		fmt.Println("error triped by queary")
		log.Fatal(err)
	}
	var aRow, b, c, d, e, f, g, h, i, j, k, l string
	
    defer row.Close()
    for row.Next() {
            if err := row.Scan(&aRow, &b, &c, &d, &e, &f, &g, &h, &i, &j, &k, &l); err != nil {
           // if err := row.Scan(&fullRow); err != nil {
                    log.Fatal(err)
            }
            ///fmt.Printf("Number of entries are %i \n", aRows)
    }
    if err := row.Err(); err != nil {
            log.Fatal(err)
    }  
    var bb bool
    //fullRow = aRow, time, c, d, e, f, g, h, i, j, k, l
    if b == "true" || b == "True"{
		 bb = true
	}else{
		bb = false
	}
    var fullRow = Row{aRow, bb, c, d, e, f, g, h, i, j, k, l}
	//fmt.Println(fullRow)
	
	return fullRow	
}*/
func GetRow(serverHostName string, Database *sql.DB) *sql.Row{
	row := Database.QueryRow("SELECT * FROM serverInformation WHERE Hostname=?", serverHostName)	
	return row	
}

/***********************************************************************
* this quearies the databace for the full table values
* Returns: string  
* 
* 
* not tested
***********************************************************************/
func GetAllRows(serverHostName string, Database *sql.DB) *sql.Rows {
	allrows, err := Database.Query("SELECT * FROM serverInformation")
	if err != nil {
		log.Fatal(err)
	}
	var aRow, b, c, d, e, f, g, h, i, j, k, l string
	z := 0
    defer allrows.Close()
    for allrows.Next() {
		z++
		fmt.Println(z)
		if err := allrows.Scan(&aRow, &b, &c, &d, &e, &f, &g, &h, &i, &j, &k, &l); err != nil {
			// if err := row.Scan(&fullRow); err != nil {
			log.Fatal(err)
		}
		fmt.Println(aRow)
		///fmt.Printf("Number of entries are %i \n", aRows)
    }
    if err := allrows.Err(); err != nil {
            log.Fatal(err)
    }/* 
    var bb bool 
	if b == "true" || b == "True"{
		 bb = true
	}else{
		bb = false
	}*/
    //var allTheRows = Row{aRow, bb, c, d, e, f, g, h, i, j, k, l}
	//fmt.Println(allTheRows)
	
	//return allTheRows		
	
	return allrows 
}



/***********************************************************************
* this quearies the databace based on what attrbuite it was passed in
* Returns: string of data in the 
* 
* 
* not tested
***********************************************************************/
func GetEntryCount(Database *sql.DB) sql.Result{
	/*NumberOfRows, err := Database.Query("SELECT COUNT(*) FROM ?")
	if err != nil {
		log.Fatal(err)
	}
	var count string
    defer NumberOfRows.Close()
    for NumberOfRows.Next() {
            if err := NumberOfRows.Scan(&count); err != nil {
                    log.Fatal(err)
            }
            //fmt.Printf("Number of entries are %i \n", count)
    }
    if err := NumberOfRows.Err(); err != nil {
            log.Fatal(err)
    }  
	fmt.Printf(count)
	return count
	*/
	NumberOfRows, err := Database.Exec("SELECT COUNT(*) FROM ?")
	if err != nil {
		log.Fatal(err)
	}
	//var count string
    //defer NumberOfRows.Close()
    /*for NumberOfRows.Next() {
            if err := NumberOfRows.Scan(&count); err != nil {
                    log.Fatal(err)
            }
            //fmt.Printf("Number of entries are %i \n", count)
    }
    if err := NumberOfRows.Err(); err != nil {
            log.Fatal(err)
    }  */
	//fmt.Printf(NumberOfRows.LastInsertId)
	return NumberOfRows
	


}
