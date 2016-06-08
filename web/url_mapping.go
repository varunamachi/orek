package web

import(
    gcweb "github.com/gocraft/web"
    orekweb "github.com/varunamachi/orek/web"
    "fmt"
    "net/http"
    "strings"
)

func MapUrls() {
    router := gcweb.New(orekweb.Context{}).
        Middleware(gcweb.LoggerMiddleware).
        Get("/", (*orekweb.Context).Default)
    http.ListenAndServe("localhost:3000", router)
}


