package main

import (
	"flag"
	"fmt"
	"github.com/varunamachi/orek/data"
    "golang.org/x/crypto/ssh/terminal"
	"log"
    "syscall"
)

func main() {
	var userName string
	var host string
	var port string
	var dbName string
	var password string

	flag.StringVar(&userName, "user-name", "", "Database user name")
	flag.StringVar(&host, "host", "localhost", "Database host")
	flag.StringVar(&dbName, "db-name", "orek", "Name of the database")
	flag.StringVar(&port, "port", "3306", "Database port")
	flag.StringVar(&password, "password", "",
		"Password, will be asked later if not provided")
    flag.Parse()
	if len(password) == 0 {
        fmt.Printf("Password for %s: ", userName)
		pbyte, err := terminal.ReadPassword(int(syscall.Stdin))
        if err != nil {
            fmt.Print(err)
        } else {
            password = string(pbyte)
        }
	}
	connStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		userName,
		password,
		host,
		port,
		dbName)
    fmt.Println(connStr)
	db, err := data.MySqlInit(connStr)
	if err == nil {
		data.SetDb(db)
	} else {
		log.Fatal("Failed to initialize database")
	}

}
