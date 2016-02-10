package web

import "github.com/varunamachi/orek/data"

type Session struct {
    Result      bool        `json:"result"`
    SessionId   string      `json:"sessionId"`
    User        data.User   `json:"user"`
}
