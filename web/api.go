package web

import (
	"github.com/varunamachi/orek/data"
	"time"
	"github.com/gocraft/web"
)

type Context struct{
	SessionUser 	data.User
	CreationTime	time.Duration
}


func (c *Context) Login(rw web.ResponseWriter, req *web.Request, next web.NextMiddlewareFunc) {

}

func Setup() error {


	return  nil
}