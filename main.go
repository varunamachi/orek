package main

import (
	"github.com/varunamachi/orek/data"
	"fmt"
)

func main() {
	user := data.User{"the_user", "the", "user", "user@mail.com"}
	fmt.Println(&user)
}
