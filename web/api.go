package web

import (
	"github.com/gocraft/web"
	"github.com/varunamachi/orek/data"
	"net/http"
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
	next(req, req)
}

func (c *Context) Login(resp http.ResponseWriter, req *http.Request) {
	//Get the user name and password from the request
	//Authenticate the user and create the user object, put the user into
	// context
	//Check if the current session already registered to some other login
	//if so expire this session
	//Create a new session in the session table
}

func (c *Context) Logout(resp http.ResponseWriter, req *http.Request) {

}

func (c *Context) GetAllUsers(resp http.ResponseWriter, req *http.Request) {

}

func (c *Context) GetUser(resp http.ResponseWriter, req *http.Request) {

}

func (c *Context) CreateUser(resp http.ResponseWriter, req *http.Request) {

}

func (c *Context) DeleteUser(resp http.ResponseWriter, req *http.Request) {

}

func (c *Context) GetAllSources(resp http.ResponseWriter, req *http.Request) {

}

func (c *Context) GetSource(resp http.ResponseWriter, req *http.Request) {

}

func (c *Context) GetSourceWithId(resp http.ResponseWriter, req *http.Request) {

}

func (c *Context) CreateOrUpdateSource(resp http.ResponseWriter,
	req *http.Request) {

}

func (c *Context) DeleteSource(resp http.ResponseWriter, req *http.Request) {

}

func (c *Context) GetAllVariables(resp http.ResponseWriter, req *http.Request) {

}

func (c *Context) GetVariable(resp http.ResponseWriter, req *http.Request) {

}

func (c *Context) GetVariableWithId(resp http.ResponseWriter,
	req *http.Request) {

}

func (c *Context) CreateOrUpdateVariable(resp http.ResponseWriter,
	req *http.Request) {

}

func (c *Context) DeleteVariable(resp http.ResponseWriter, req *http.Request) {

}

func (c *Context) GetAllUserGroups(resp http.ResponseWriter,
	req *http.Request) {

}

func (c *Context) GetUserGroup(resp http.ResponseWriter, req *http.Request) {

}

func (c *Context) CreateOrUpdateUserGroup(resp http.ResponseWriter,
	req *http.Request) {

}

func (c *Context) DeleteUserGroup(resp http.ResponseWriter, req *http.Request) {

}

func (c *Context) GetAllVariableGroups(resp http.ResponseWriter,
	req *http.Request) {

}

func (c *Context) GetVariableGroup(resp http.ResponseWriter,
	req *http.Request) {

}

func (c *Context) GetVariableGroupWithId(resp http.ResponseWriter,
	req *http.Request) {

}

func (c *Context) CreateOrUpdateVariableGroup(resp http.ResponseWriter,
	req *http.Request) {

}

func (c *Context) DeleteVariableGroup(resp http.ResponseWriter,
	req *http.Request) {

}

func (c *Context) AddUserToGroup(resp http.ResponseWriter, req *http.Request) {

}

func (c *Context) RemoveUserFromGroup(resp http.ResponseWriter,
	req *http.Request) {

}

func (c *Context) GetUsersInGroup(resp http.ResponseWriter, req *http.Request) {

}

func (c *Context) GetGroupsForUser(resp http.ResponseWriter,
	req *http.Request) {

}

func (c *Context) AddVariableToGroup(resp http.ResponseWriter,
	req *http.Request) {

}

func (c *Context) RemoveVariableFromGroup(resp http.ResponseWriter,
	req *http.Request) {

}

func (c *Context) GetVariablesInGroup(resp http.ResponseWriter,
	req *http.Request) {

}

func (c *Context) GetGroupsForVariable(resp http.ResponseWriter,
    req *http.Request) {

}

func (c *Context) AddVariableValue(resp http.ResponseWriter,
    req *http.Request) {

}

func (c *Context) ClearValuesForVariable(resp http.ResponseWriter,
    req *http.Request) {

}

func (c *Context) GetValuesForVariable(resp http.ResponseWriter,
    req *http.Request) {

}

func Setup() error {

	return nil
}
