package main

import (
	"fmt"
	"github.com/varunamachi/orek/data"
)

func main() {
	user := data.User{"the_user", "the", "user", "user@mail.com"}
	fmt.Println(&user)
}
