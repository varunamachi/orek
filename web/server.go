package web

import (
	"net/http"

	"github.com/gocraft/web"
)

func mapUserUrls(router *web.Router) {
	router.
		Get("/orek/v0/users", (*Context).GetAllUsers).
		Get("/orek/v0/users/:userName", (*Context).GetUser).
		Post("/orek/v0/users", (*Context).CreateUser).
		Delete("/orek/v0/users/:userName", (*Context).DeleteUser)
}

func mapSourceUrls(router *web.Router) {
	router.
		Get("/orek/v0/sources", (*Context).GetAllSources).
		Get("/orek/v0/sources/:sourceOwner/:sourceName", (*Context).GetSource).
		Get("/orek/v0/sources/:sourceId", (*Context).GetSourceWithId).
		Post("/orek/v0/sources", (*Context).CreateOrUpdateSource).
		Delete("/orek/v0/sources/:sourceId", (*Context).DeleteSource)
}

func mapVariableUrls(router *web.Router) {
	router.
		Get("/orek/v0/variables", (*Context).GetAllVariables).
		Get("/orek/v0/variables/:sourceId/:variableName", (*Context).GetVariable).
		Get("/orek/v0/variables/:variableId", (*Context).GetVariableWithId).
		Post("/orek/v0/variables", (*Context).CreateOrUpdateVariable).
		Delete("/orek/v0/variables/:variableId", (*Context).DeleteVariable)
}

func def(router *web.Router) {
	router.Middleware(web.LoggerMiddleware).
		Get("/*", (*Context).Default)
}

func Serve() {
	router := web.New(Context{})
	def(router)
	mapUserUrls(router)
	mapSourceUrls(router)
	mapVariableUrls(router)
	http.ListenAndServe("localhost:3000", router)
}
