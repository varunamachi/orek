package web

import (
	"encoding/json"
	"log"
	"time"

	"github.com/gocraft/web"
	"github.com/varunamachi/orek/data"
)

type Context struct {
	SessionUser  data.User
	CreationTime time.Duration
}

func checkSession(context *Context) bool {
	return true
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
	//	userName := req.FormValue("userName")
	//	password := req.FormValue("password")
	//	sessionId := req.FormValue("sessionId")
	//	data.DataSource().GetUser("varun").FirstName

	//	data.DataSource().GetUser(
}

func (c *Context) Logout(resp web.ResponseWriter, req *web.Request) {
	//resp.Write
}

func (c *Context) GetAllUsers(resp web.ResponseWriter, req *web.Request) {
	encoder := json.NewEncoder(resp)
	users, err := data.DataSource().GetAllUsers()
	if err == nil {
		err = encoder.Encode(users)
		if err != nil {
			err = encoder.Encode(OrekError{
				"MarshalError",
				"Failed to marshal user list"})
		}
	} else {
		err = encoder.Encode(OrekError{
			"DataSourceError",
			"Failed retrieve user list from datasource"})
	}
	if err != nil {
		log.Print(err)
	}
}

func (c *Context) GetUser(resp web.ResponseWriter, req *web.Request) {
	userName := req.PathParams["userName"]
	user, err := data.DataSource().GetUser(userName)
	encoder := json.NewEncoder(resp)
	if err == nil {
		err = encoder.Encode(user)
		if err != nil {
			err = encoder.Encode(OrekError{
				"MarshalError",
				"Failed to marshal user details"})
		}
	} else {
		err = encoder.Encode(OrekError{
			"DataSourceError",
			"Failed to fetch user detail from data source"})
	}
	if err != nil {
		log.Print(err)
	}
}

func (c *Context) CreateUser(resp web.ResponseWriter, req *web.Request) {
	encoder := json.NewEncoder(resp)
	decoder := json.NewDecoder(req.Body)
	var user data.User
	var err error = nil
	if err = decoder.Decode(&user); err == nil {
		err = data.DataSource().CreateOrUpdateUser(&user)
		if err != nil {
			err = encoder.Encode(OrekError{
				"MarshalError",
				"Failed to marshal user creation result"})
		}
	} else {
		err = encoder.Encode(OrekError{
			"UnmarshalError",
			"Failed to create user from given information"})
	}
	if err != nil {
		log.Print(err)
	}
}

func (c *Context) DeleteUser(resp web.ResponseWriter, req *web.Request) {
	err := req.ParseForm()
	userName := req.Form.Get("userName")
	encoder := json.NewEncoder(resp)
	if err == nil && len(userName) > 0 {
		err = data.DataSource().DeleteUser(userName)
		if err != nil {
			err = encoder.Encode(OrekError{
				"DataSourceError",
				"Failed to delete user"})
		}

	} else {
		err = encoder.Encode(OrekError{
			"ArgumentError",
			"Invalid user name given for deletion"})
	}
	if err != nil {
		log.Print(err)
	}
}

func (c *Context) GetAllSources(resp web.ResponseWriter, req *web.Request) {
	sources, err := data.DataSource().GetAllSources()
	encoder := json.NewEncoder(resp)
	if err == nil {
		err = encoder.Encode(sources)
		if err != nil {
			encoder.Encode(OrekError{
				"MarshalError",
				"Failed to marshal source list"})
		}
	} else {
		encoder.Encode(OrekError{
			"DataSourceError",
			"Failed to fetch list of source from database"})

	}
	if err == nil {
		log.Print(err)
	}
}

func (c *Context) GetSource(resp web.ResponseWriter, req *web.Request) {
	sourceName := req.PathParams["sourceName"]
	sourceOwner := req.PathParams["sourceOwner"]
	encoder := json.NewEncoder(resp)
	source, err := data.DataSource().GetSource(sourceName, sourceOwner)
	if err == nil {
		err = encoder.Encode(source)
		if err != nil {
			encoder.Encode(OrekError{
				"MarshalError",
				"Failed to marshal source information"})
		}
	} else {
		encoder.Encode(OrekError{
			"DataSourceError",
			"Failed to fetch source information from database"})
	}
	if err != nil {
		log.Print(err)
	}
}

func (c *Context) GetSourceWithId(resp web.ResponseWriter, req *web.Request) {
	sourceId := req.PathParams["sourceId"]
	encoder := json.NewEncoder(resp)
	source, err := data.DataSource().GetSourceWithId(sourceId)
	if err == nil {
		err = encoder.Encode(source)
		if err != nil {
			encoder.Encode(OrekError{
				"MarshalError",
				"Failed to marshal source information"})
		}
	} else {
		encoder.Encode(OrekError{
			"DataSourceError",
			"Failed to fetch source information from database"})
	}
	if err != nil {
		log.Print(err)
	}
}

func (c *Context) CreateOrUpdateSource(resp web.ResponseWriter, req *web.Request) {
	decoder := json.NewDecoder(req.Body)
	encoder := json.NewEncoder(resp)
	var source data.Source
	err := decoder.Decode(&source)
	if err == nil {
		err = data.DataSource().CreateOrUpdateSource(&source)
		if err != nil {
			encoder.Encode(OrekError{
				"DataSourceError",
				"Failed to add source to database"})
		}
	} else {
		encoder.Encode(OrekError{
			"UnamarshalError",
			"Could not decode the given source information"})
	}
	if err != nil {
		log.Print(err)
	}

}

func (c *Context) DeleteSource(resp web.ResponseWriter, req *web.Request) {
	//sourceId := req.PathParams["sourceId"]
	sourceId := req.Form.Get("sourceId")
	err := data.DataSource().DeleteSource(sourceId)
	encoder := json.NewEncoder(resp)
	if err != nil {
		encoder.Encode(OrekError{
			"DataSourceError",
			"Could not delete the source with given id"})
		log.Print(err)
	}
}

func (c *Context) GetAllVariables(resp web.ResponseWriter, req *web.Request) {
	vars, err := data.DataSource().GetAllVariables()
	encoder := json.NewEncoder(resp)
	if err == nil {
		err = encoder.Encode(vars)
		if err != nil {
			encoder.Encode(OrekError{
				"MarshalError",
				"Could not encode the list of variables"})
		}
	} else {
		encoder.Encode(OrekError{
			"DataSourceError",
			"Could not get list of variables from DB"})
	}
	if err != nil {
		log.Print(err)
	}
}

func (c *Context) GetVariable(resp web.ResponseWriter, req *web.Request) {
	variableName := req.PathParams["variableName"]
	sourceId := req.PathParams["sourceId"]
	variable, err := data.DataSource().GetVariable(sourceId, variableName)
	encoder := json.NewEncoder(resp)
	if err == nil {
		err = encoder.Encode(variable)
		if err != nil {
			encoder.Encode(OrekError{
				"MarshalError",
				"Failed to encode variable information"})
		}
	} else {
		encoder.Encode(OrekError{
			"DataSourceError",
			"Failed to fetch information about the given variable"})

	}
	if err != nil {
		log.Print(err)
	}
}

func (c *Context) GetVariableWithId(resp web.ResponseWriter, req *web.Request) {
	variableId := req.PathParams["variableId"]
	variable, err := data.DataSource().GetVariableWithId(variableId)
	encoder := json.NewEncoder(resp)
	if err == nil {
		err = encoder.Encode(variable)
		if err != nil {
			encoder.Encode(OrekError{
				"MarshalError",
				"Failed to encode variable information"})
		}
	} else {
		encoder.Encode(OrekError{
			"DataSourceError",
			"Failed to fetch information about the given variable"})

	}
	if err != nil {
		log.Print(err)
	}

}

func (c *Context) CreateOrUpdateVariable(resp web.ResponseWriter, req *web.Request) {
	decoder := json.NewDecoder(req.Body)
	encoder := json.NewEncoder(resp)
	var variable data.Variable
	err := decoder.Decode(&variable)
	if err == nil {
		err := data.DataSource().CreateOrUpdateVariable(&variable)
		if err != nil {
			encoder.Encode(OrekError{
				"DataSourceError",
				"Failed to create/update given variable"})
		}
	} else {
		encoder.Encode(OrekError{
			"UnmarshalError",
			"Failed to unmarshal variable information"})
	}
	if err != nil {
		log.Print(err)
	}
}

func (c *Context) DeleteVariable(resp web.ResponseWriter, req *web.Request) {
	//variableId := req.PathParams["variableId"]
	variableId := req.Form.Get("variableId")
	encoder := json.NewEncoder(resp)
	err := data.DataSource().DeleteVariable(variableId)
	if err != nil {
		encoder.Encode(OrekError{
			"DataSourceError",
			"Failed to delete give variable"})
	}
	if err != nil {
		log.Print(err)
	}
}

func (c *Context) GetAllUserGroups(resp web.ResponseWriter, req *web.Request) {
	encoder := json.NewEncoder(resp)
	groups, err := data.DataSource().GetAllUserGroups()
	if err == nil {
		err = encoder.Encode(groups)
		if err != nil {
			encoder.Encode(OrekError{
				"MarshalError",
				"Failed marshal user group list"})
		}
	} else {
		encoder.Encode(OrekError{
			"DataSourceError",
			"Could not fetch list of user groups from database"})
	}
}

func (c *Context) GetUserGroup(resp web.ResponseWriter, req *web.Request) {
	groupName := req.PathParams["groupName"]
	encoder := json.NewEncoder(resp)
	group, err := data.DataSource().GetUserGroup(groupName)
	if err == nil {
		err = encoder.Encode(group)
		if err != nil {
			encoder.Encode(OrekError{
				"MarshalError",
				"Failed to encode user information"})
		}
	} else {
		encoder.Encode(OrekError{
			"DataSourceError",
			"Failed to fetch user group information"})
	}
	if err != nil {
		log.Print(err)
	}
}

func (c *Context) CreateOrUpdateUserGroup(resp web.ResponseWriter,
	req *web.Request) {
	decoder := json.NewDecoder(req.Body)
	encoder := json.NewEncoder(resp)
	var group data.UserGroup
	err := decoder.Decode(&group)
	if err == nil {
		err = data.DataSource().CreateOrUpdateUserGroup(&group)
		if err != nil {
			encoder.Encode(OrekError{
				"DataSourceError",
				"Failed to create/update user group info."})
		}
	} else {
		encoder.Encode(OrekError{
			"UnamrshalError",
			"Failed to decode user group info for create/update"})
	}
	if err != nil {
		log.Print(err)
	}
}

func (c *Context) DeleteUserGroup(resp web.ResponseWriter, req *web.Request) {
	encoder := json.NewEncoder(resp)
	groupName := req.Form.Get("groupName")
	err := data.DataSource().DeleteUserGroup(groupName)
	if err != nil {
		encoder.Encode(OrekError{
			"DataSourceError",
			"Failed to delete given user"})
		log.Print(err)
	}
}

func (c *Context) GetAllVariableGroups(resp web.ResponseWriter, req *web.Request) {
	encoder := json.NewEncoder(resp)
	varGroups, err := data.DataSource().GetAllVariables()
	if err == nil {
		err = encoder.Encode(varGroups)
		if err != nil {
			encoder.Encode(OrekError{
				"MarshalError",
				"Failed to encode variable group list"})
		}
	} else {
		encoder.Encode(OrekError{
			"DataSourceError",
			"Failed to fetch all variable groups"})
	}
	if err != nil {
		log.Print(err)
	}
}

func (c *Context) GetVariableGroup(resp web.ResponseWriter, req *web.Request) {
	varGroupName := req.PathParams["varGroupName"]
	owner := req.PathParams["owner"]
	encoder := json.NewEncoder(resp)
	varGroup, err := data.DataSource().GetVariableGroup(varGroupName, owner)
	if err == nil {
		err = encoder.Encode(varGroup)
		if err != nil {
			encoder.Encode(OrekError{
				"MarshalError",
				"Failed to encode the var group list"})
		}
	} else {
		encoder.Encode(OrekError{
			"DataSourceError",
			"Failed to fetch variable group with give info"})
	}
	if err != nil {
		log.Print(err)
	}
}

func (c *Context) GetVariableGroupWithId(resp web.ResponseWriter, req *web.Request) {
	varGroupId := req.PathParams["varGroupId"]
	encoder := json.NewEncoder(resp)
	varGroup, err := data.DataSource().GetVariableGroupWithId(varGroupId)
	if err == nil {
		err = encoder.Encode(varGroup)
		if err != nil {
			encoder.Encode(OrekError{
				"MarshalError",
				"Failed to encode the var group list"})
		}
	} else {
		encoder.Encode(OrekError{
			"DataSourceError",
			"Failed to fetch variable group with give info"})
	}
	if err != nil {
		log.Print(err)
	}
}

func (c *Context) CreateOrUpdateVariableGroup(resp web.ResponseWriter, req *web.Request) {
	decoder := json.NewDecoder(req.Body)
	encoder := json.NewEncoder(resp)
	var varGroup data.VariableGroup
	err := decoder.Decode(&varGroup)
	if err == nil {
		err = data.DataSource().CreateOrUpdateVariableGroup(&varGroup)
		if err != nil {
			encoder.Encode(OrekError{
				"DataSourceError",
				"Failed create/update variable group"})
		}
	} else {
		encoder.Encode(OrekError{
			"UnmarshalError",
			"Failed decode variable group info"})
	}
	if err != nil {
		log.Print(err)
	}
}

func (c *Context) DeleteVariableGroup(resp web.ResponseWriter, req *web.Request) {
	encoder := json.NewEncoder(resp)
	varGroupId := req.Form.Get("varGroupId")
	err := data.DataSource().DeleteVariableGroup(varGroupId)
	if err != nil {
		encoder.Encode(OrekError{
			"DataSourceError",
			"Failed to delete variable group"})
		log.Print(err)
	}
}

func (c *Context) AddUserToGroup(resp web.ResponseWriter, req *web.Request) {
	userId := req.Form.Get("userName")
	groupId := req.Form.Get("groupId")
	err := data.DataSource().AddUserToGroup(userId, groupId)
	if err != nil {
		encoder := json.NewEncoder(resp)
		encoder.Encode(OrekError{
			"DataSourceError",
			"Failed to add user to group"})
		log.Print(err)
	}
}

func (c *Context) RemoveUserFromGroup(resp web.ResponseWriter, req *web.Request) {
	userName := req.Form.Get("userName")
	groupId := req.Form.Get("groupId")
	err := data.DataSource().RemoveUserFromGroup(userName, groupId)
	if err != nil {
		encoder := json.NewEncoder(resp)
		encoder.Encode(OrekError{
			"DataSourceError",
			"Failed to remove user from group"})
		log.Print(err)
	}
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
