package main

import (
	"github.com/varunamachi/orek/data"
    "log"
)

func main() {
    db, err := data.MySqlInit("hello")
    if err == nil {
        data.SetDb(db)
    } else {
        log.Fatal("Failed to initialize database")
    }

}
