package main

import (
	"flag"
	"fmt"
	"log"
	"syscall"

	"github.com/varunamachi/orek/data"
	orekweb "github.com/varunamachi/orek/web"
	"golang.org/x/crypto/ssh/terminal"
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
	if len(password) == 0 && len(userName) != 0 {
		fmt.Printf("Password for %s: ", userName)
		pbyte, err := terminal.ReadPassword(int(syscall.Stdin))
		if err != nil {
			fmt.Print(err)
		} else {
			password = string(pbyte)
		}
	} else {
		log.Fatal("Insufficient parameters for Orek to run!")
	}
	options := &data.MysqlOptions{
		UserName: userName,
		Password: password,
		Host:     host,
		Port:     port,
		DbName:   dbName}
	// fmt.Println(options)
	db, err := data.MysqlInit(options)
	if err == nil {
		data.SetDataSource(db)
		orekweb.Serve()
	} else {
		log.Fatal("Failed to initialize database")
	}

}
