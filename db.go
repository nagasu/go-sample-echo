package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

// DB設定値
const (
    HOST     = "localhost"
    DATABASE = "beauty"
    USER     = "postgres"
    PASSWORD = "postgres"
)

type Staff struct {
    ID    int             `json:"id"`
    No    sql.NullInt64   `json:no`
    FName  sql.NullString `json:"f_name" validate:"required"`
    SName sql.NullString `json:"s_name" validate:"required,email"`
}

type Staffs []*Staff

func checkError(err error) {
    if err != nil {
        panic(err)
    }
}

func allUsers() []*Staff {
    connectionString := fmt.Sprintf(
        "host=%s user=%s password=%s dbname=%s sslmode=disable",
        HOST,
        USER,
        PASSWORD,
        DATABASE)

    // Initialize connection object.
    db, err := sql.Open("postgres", connectionString)
    checkError(err)

    err = db.Ping()
    checkError(err)
    fmt.Println("Successfully created connection to database")

    // Read rows from table.
    var id int
    var no sql.NullInt64
    var fName sql.NullString
    var sName sql.NullString
    var ss Staffs

    sqlStatement := "SELECT id, no, f_name, s_name from staff"
    rows, err := db.Query(sqlStatement)
    checkError(err)
    defer rows.Close()

    for rows.Next() {
        switch err := rows.Scan(&id, &no, &fName, &sName); err {
        case sql.ErrNoRows:
            fmt.Println("No rows were returned")
        case nil:
            s := &Staff {
                ID: id,
                No: no,
                FName: fName,
                SName: sName,
            }
            ss = append(ss, s)
            fmt.Printf("Data row = (%d, %d, %s, %s)\n", id, no, fName, sName)
        default:
            checkError(err)
        }
    }

    return ss
}

func create() {
    connectionString := fmt.Sprintf(
        "host=%s user=%s password=%s dbname=%s sslmode=disable",
        HOST,
        USER,
        PASSWORD,
        DATABASE)

    // Initialize connection object.
    db, err := sql.Open("postgres", connectionString)
    checkError(err)

    err = db.Ping()
    checkError(err)
    fmt.Println("Successfully created connection to database")

    // Drop previous table of same name if one exists.
    _, err = db.Exec("DROP TABLE IF EXISTS inventory;")
    checkError(err)
    fmt.Println("Finished dropping table (if existed)")

    // Create table.
    _, err = db.Exec("CREATE TABLE inventory (id serial PRIMARY KEY, name VARCHAR(50), quantity INTEGER);")
    checkError(err)
    fmt.Println("Finished creating table")

    // Insert some data into table.
    sqlStatement := "INSERT INTO inventory (name, quantity) VALUES ($1, $2);"
    _, err = db.Exec(sqlStatement, "banana", 150)
    checkError(err)
    _, err = db.Exec(sqlStatement, "orange", 154)
    checkError(err)
    _, err = db.Exec(sqlStatement, "apple", 100)
    checkError(err)
    fmt.Println("Inserted 3 rows of data")
}
