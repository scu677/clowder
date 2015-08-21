package praticeReadDB

import (
    _ "driver"
    "database/sql"
    "github.com/jmoiron/sqlx"
)
type User struct {
    Id int
    Name string
}
func main() {
    db := sqlx.Connect(...)
    user := User{}
    rows, _ := db.Queryx("SELECT id, name FROM users;")
    for rows.Next() {
        rows.StructScan(&user)
    }
}
