package web

import (
	"github.com/gocraft/web"
	"github.com/varunamachi/orek/data"
	// "net/http"
	"time"
)

type Context struct {
	SessionUser  data.User
	CreationTime time.Duration
}

func (c *Context) SessionChecker(rw web.ResponseWriter,
	req *web.Request,
	next web.NextMiddlewareFunc) {
	//get the session ID
	//see if it is still valid
	//put the user and session objects into the context object
	next(rw, req)
}

func (c *Context) Login(resp web.ResponseWriter, req *web.Request) {
	//Get the user name and password from the request
	//Authenticate the user and create the user object, put the user into
	// context
	//Check if the current session already registered to some other login
	//if so expire this session
	//Create a new session in the session table
    userName := req.FormValue("userName")
    password := req.FormValue("password")
    sessionId := req.FormValue("sessionId")
//    data.OrekDb().
}

func (c *Context) Logout(resp web.ResponseWriter, req *web.Request) {
    
}

func (c *Context) GetAllUsers(resp web.ResponseWriter, req *web.Request) {

}

func (c *Context) GetUser(resp web.ResponseWriter, req *web.Request) {

}

func (c *Context) CreateUser(resp web.ResponseWriter, req *web.Request) {

}

func (c *Context) DeleteUser(resp web.ResponseWriter, req *web.Request) {

}

func (c *Context) GetAllSources(resp web.ResponseWriter, req *web.Request) {

}

func (c *Context) GetSource(resp web.ResponseWriter, req *web.Request) {

}

func (c *Context) GetSourceWithId(resp web.ResponseWriter, req *web.Request) {

}

func (c *Context) CreateOrUpdateSource(resp web.ResponseWriter,
	req *web.Request) {

}

func (c *Context) DeleteSource(resp web.ResponseWriter, req *web.Request) {

}

func (c *Context) GetAllVariables(resp web.ResponseWriter, req *web.Request) {

}

func (c *Context) GetVariable(resp web.ResponseWriter, req *web.Request) {

}

func (c *Context) GetVariableWithId(resp web.ResponseWriter,
	req *web.Request) {

}

func (c *Context) CreateOrUpdateVariable(resp web.ResponseWriter,
	req *web.Request) {

}

func (c *Context) DeleteVariable(resp web.ResponseWriter, req *web.Request) {

}

func (c *Context) GetAllUserGroups(resp web.ResponseWriter,
	req *web.Request) {

}

func (c *Context) GetUserGroup(resp web.ResponseWriter, req *web.Request) {

}

func (c *Context) CreateOrUpdateUserGroup(resp web.ResponseWriter,
	req *web.Request) {

}

func (c *Context) DeleteUserGroup(resp web.ResponseWriter, req *web.Request) {

}

func (c *Context) GetAllVariableGroups(resp web.ResponseWriter,
	req *web.Request) {

}

func (c *Context) GetVariableGroup(resp web.ResponseWriter,
	req *web.Request) {

}

func (c *Context) GetVariableGroupWithId(resp web.ResponseWriter,
	req *web.Request) {

}

func (c *Context) CreateOrUpdateVariableGroup(resp web.ResponseWriter,
	req *web.Request) {

}

func (c *Context) DeleteVariableGroup(resp web.ResponseWriter,
	req *web.Request) {

}

func (c *Context) AddUserToGroup(resp web.ResponseWriter, req *web.Request) {

}

func (c *Context) RemoveUserFromGroup(resp web.ResponseWriter,
	req *web.Request) {

}

func (c *Context) GetUsersInGroup(resp web.ResponseWriter, req *web.Request) {

}

func (c *Context) GetGroupsForUser(resp web.ResponseWriter,
	req *web.Request) {

}

func (c *Context) AddVariableToGroup(resp web.ResponseWriter,
	req *web.Request) {

}

func (c *Context) RemoveVariableFromGroup(resp web.ResponseWriter,
	req *web.Request) {

}

func (c *Context) GetVariablesInGroup(resp web.ResponseWriter,
	req *web.Request) {

}

func (c *Context) GetGroupsForVariable(resp web.ResponseWriter,
    req *web.Request) {

}

func (c *Context) AddVariableValue(resp web.ResponseWriter,
    req *web.Request) {

}

func (c *Context) ClearValuesForVariable(resp web.ResponseWriter,
    req *web.Request) {

}

func (c *Context) GetValuesForVariable(resp web.ResponseWriter,
    req *web.Request) {

}

func Setup() error {

	return nil
}
