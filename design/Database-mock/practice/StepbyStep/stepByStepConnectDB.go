/******************************************
http://golang-basic.blogspot.ca/2014/06/golang-database-step-by-step-guide-on.html
******************************************/

package DB_DRIVERconnection

import (
"database/sql"
"database/sql/driver"
"fmt"
"log"
"github.com/mattn/go-sqlite3"
//"code.google.com/p/go-sqlite/go1/sqlite3"
//"sqlite3"
)



func main(){
	var DB_DRIVER string
	sql.Register(DB_DRIVER, &sqlite3.SQLiteDriver{Extensions: []string{
			"sqlite3_mod_regexp",
		},
		}) //Problem here 

	
	
	
	/*	database is the pointer to the database
		DB_DRIVER is the name of Database Driver that we registered in the previous step.
		mysqlite_3 is the name of the datasource that gets created as a successful result of Open operation
	*/ 
	database, err := sql.Open(DB_DRIVER, "mysqlite_3")
	if err != nil {
	 fmt.Println("Failed to create the handle")
	}
	
	makeDB()
	
	if err2 := database.Ping(); err2 != nil {
		fmt.Println("Failed to keep connection alive")
	}
	
	
	
	tx, err := database.Begin()
		if err != nil {
		 log.Fatal(err)
	}
	
	
	
	
	// Following is the create statement for Persons struct

		result, err := database.Exec(
		 "CREATE TABLE IF NOT EXISTS Persons ( id integer PRIMARY KEY, LastName varchar(255) NOT NULL, FirstName varchar(255), Address varchar(255), City varchar(255), CONSTRAINT uc_PersonID UNIQUE (id,LastName))",)

		if err != nil {
		 log.Fatal(err)
		}


	// Following is the create statement exec command for Employee struct
	// Here we are using the person_id as the foreign key

	result, err = database.Exec(
	 "create table IF NOT EXISTS employee (employeeID integer PRIMARY KEY,name varchar(255) NOT null,age int, person_id int, FOREIGN KEY (person_id) REFERENCES persons(id), CONSTRAINT uc_empID UNIQUE (employeeID, person_id, name))",)
	 
	if err != nil {
	 log.Fatal(err)
	}






	// Lets first insert values in person table as Employee table has some dependency on the Person's table.

	result, err = database.Exec("Insert into Persons (id, LastName, FirstName, Address, City) values (?, ?, ?, ?, ?)", nil, "soni", "swati", "110 Eastern drive", "Mountain view, CA")

	if err != nil {
	 log.Fatal(err)
	}


	// Next, we will insert values in the Employee table. 
	// Here the value of employeeID is nil as it should be auto-incremented by SQLITE-3

	result, err = database.Exec(
	  "INSERT INTO employee (employeeID, name, age, person_id) VALUES (?, ?, ?, ?)", nil,
	  "Swati Soni",
	  24, 1,)

	if err != nil {
	 log.Fatal(err)
	}



	tx.Commit()


	   rows, err := database.Query("SELECT * FROM employee")
		if err != nil {
			log.Fatal(err)
		}

    id := 123
    var firstName string
    row, err := db.QueryRow(
    "SELECT Firstname FROM Persons WHERE id=?", id)
    
    
    
    
    // In the previous step , we used Query() function to
	// retrieve all rows of the employee table into rows* of type Rows.
	// Now we will iterate through each one of them to get the values for each row.

	for rows.Next() {

	 var empID sql.NullInt64
	 var empName sql.NullString
	 var empAge sql.NullInt64
	 var empPersonId sql.NullInt64

	 if err := rows.Scan(&empID, &empName, &empAge, 
							   &empPersonId); err != nil {
			  log.Fatal(err)
	 }

	 fmt.Printf("ID %d with personID:%d & name %s is age %d\n",       
					   empID.Int64, empPersonId.Int64, empName.String, empAge.Int64)
	}



	// In the previous step, we used QueryRow() to 
	// retrieve exactly one row from the Person table.
	// The record which has an ID of 123.
	// Since ID is the PRIMARY_KEY on Person table, 
	// there should be either none or exactly one returned row from the Query.

	// We can use the Scan function on Row to retrieve values from row :

	var firstName2 string
	row.Scan(& firstName2)
		
	switch {
		case err == sql.ErrNoRows:
				log.Printf("No Person with that ID.")
		case err != nil:
				log.Fatal(err)
		default:
				fmt.Printf("Person is %s\n", firstName2)
	}
    
}
	// Syntax of Register function inside or database/sql package
	//func Register(name string, driver driver.Driver)



func makeDB(){
		 type Person struct{
		 key       sql.NullInt64
		 firstName sql.NullString
		 lastName  sql.NullString
		 address   sql.NullString
		 city      sql.NullString
		}

		type employee struct{
		 empID       sql.NullInt64
		 empName     sql.NullString
		 empAge      sql.NullInt64
		 empPersonId sql.NullInt64
		}
						
	}
