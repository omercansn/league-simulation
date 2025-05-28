package repository

import (
    "database/sql"
    _ "github.com/lib/pq"
    "log"
)

var DB *sql.DB

func InitDB(dataSourceName string) {
    var err error
    DB, err = sql.Open("postgres", dataSourceName)
    if err != nil {
        log.Fatal("failed to open db:", err)
    }
    if err = DB.Ping(); err != nil {
        log.Fatal("failed to connect to db:", err)
    }
}
