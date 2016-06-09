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
		Delete("/orek/v0/users/:userName", (*Context).DeleteUser).
		Get("/orek/v0/users/:userName/groups", (*Context).GetGroupsForUser)
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
		Delete("/orek/v0/variables/:variableId", (*Context).DeleteVariable).
		Get("/orek/variables/:variableId/groups", (*Context).GetGroupsForVariable).
		Put("/orek/variables/:variableId/values", (*Context).AddVariableValue).
		Get("/orek/variables/:variableId/values", (*Context).GetValuesForVariable).
		Delete("/orek/variables/:variableId/values", (*Context).ClearValuesForVariable)
}

func mapUserGroupUrls(router *web.Router) {
	router.
		Get("/orek/v0/userGroups", (*Context).GetAllUserGroups).
		Get("/orek/v0/userGroups/:groupName", (*Context).GetUserGroup).
		Post("/orek/v0/userGroups", (*Context).CreateOrUpdateUserGroup).
		Delete("/orek/v0/userGroups/:groupName", (*Context).DeleteUserGroup).
		Get("/orek/v0/userGroups/users", (*Context).GetUsersInGroup).
		Put("/orek/v0/userGroups/:groupName/:userName", (*Context).AddUserToGroup).
		Delete("/orek/v0/userGroups/:groupName/:userName", (*Context).RemoveUserFromGroup)
}

func mapVariableGroupUrls(router *web.Router) {
	router.
		Get("/orek/v0/variableGroups", (*Context).GetAllVariableGroups).
		Get("/orek/v0/variableGroups/:owner/:varGroupName", (*Context).GetVariableGroup).
		Get("/orek/v0/variableGroups/:varGroupId", (*Context).GetVariableGroupWithId).
		Post("/orek/v0/variableGroups", (*Context).CreateOrUpdateVariableGroup).
		Delete("/orek/v0/variableGroups/:varGroupId", (*Context).DeleteVariableGroup).
		Put("/orek/v0/variableGroups/:varGroupId/:variableId", (*Context).AddVariableToGroup).
		Get("/orek/v0/variableGroups/:varGroupId", (*Context).GetVariablesInGroup)
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
