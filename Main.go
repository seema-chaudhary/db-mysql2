package main

import (
    "database/sql"
    "fmt"
    _ "github.com/go-sql-driver/mysql"
)

func main() {
    //open connection
    conn, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/mydb")
    if err != nil {
        panic(err.Error())
    }
    defer conn.Close()

    //prepare statement to get particular detail
    stmt, err := conn.Prepare("select * from employee where empId=?")
    if err != nil {
        panic(err.Error()) // log.Fatel(err)
    }
    defer stmt.Close()
    
    //execute query/statement
    var empId int
    var empName, email string

    err = stmt.QueryRow(2).Scan(&empId, &empName, &email)
    if err != nil {
        panic(err.Error())
    }
    fmt.Printf("EmpId : %d, EmpName : %s, Email : %s\n", empId, empName, email)

    //close connection
    conn.Close()

}
