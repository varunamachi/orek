package data

var db OrekDb = nil

func SetDb(dbInst *OrekDb) error {
	db = dbInst
}

type OrekDb interface {
	GetAllUsers() ([]*User, error)
	GetUser(userName string) (*User, error)
	GetUserWithEmail(email string) (*User, error)
	CreateOrUpdateUser(user *User) error
	DeleteUser(userName string) error

	GetAllSources() ([]*User, error)
	GetSource(sourceName, owner string) (*Source, error)
	GetSourceWithId(sourceId string) (*Source, error)
	CreateOrUpdateSource(source *Source) error
	DeleteSource(sourceId string) error

	GetAllVariables() ([]*Variable, error)
	GetVariable(sourceId, name string) (*Variable, error)
	GetVariableWithId(variableId string) (*Variable, error)
	CreateOrUpdateVariable(variable *Variable) error
	DeleteVariable(variableId string) error

	GetAllUserGroups() ([]*UserGroup, error)
	GetUserGroup(userGroupName string) (*UserGroup, error)
	CreateOrUpdateUserGroup(userGroup *UserGroup) error
	DeleteUserGroup(userGroupName string) error

	GetAllVariableGroups([]*VariableGroup, error)
	GetVariableGroup(varGroupName, owner string) (*VariableGroup, error)
	GetVariableGroupWithId(varGroupId string) (*VariableGroup, error)
	CreateOrUpdateVariableGroup(variableGroup *VariableGroup) error
	DeleteVariableGroup(varGroupId string) error

	AddUserToGroup(userName, groupName string) error
	RemoveUserFromGroup(userName, groupName string) error
	GetUsersInGroup(groupName string) ([]*User, error)
	GetGroupsForUser(userName string) ([]*UserGroup, error)

	AddVariableToGroup(variableId, varGroupId string) error
	RemoveVariableFromGroup(variableId, varGroupId string) error
	GetVariablesInGroup(groupName string) ([]*Variable, error)
	GetGroupsForVariable(variableId string) ([]*VariableGroup, error)

	AddVariableValue(variableId, value string) error
	ClearValuesForVariable(variableId string) error
	GetValuesForVariable(variableId string) error
}
