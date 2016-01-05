package data

var db OrekDb = nil

type OrekDb interface {
	GetUser(userName string) (*User, error)
	GetUserWithEmail(email string) (*User, error)
	CreateOrUpdateUser(user *User) error
	UpdateUserField(userName string, fields map[string]string) error
	DeleteUser(user *User) error
	DeleteUserWithId(userId string) error

	GetSource(sourceName string, owner string) (*Source, error)
	GetSourceWithId(sourceId string) (*Source, error)
	CreateOrUpdateSource(source *Source) error
	UpdateSourceField(sourceId string, fields map[string]string) error
	DeleteSource(source *Source) error
	DeleteSourceWithId(sourceId *Source) error

	GetVariable(sourceId string, name string) (*Variable, error)
	GetVariableWithId(variableId string) (*Variable, error)
	CreateOrUpdateVariable(variable *Variable) error
	UpdateVariableField(variableId string, fields map[string]string) error
	DeleteVariable(variable *Variable) error
	DeleteVariableWithId(variableId string) error

	GetUserGroup(userGroupName string) (*UserGroup, error)
	CreateOrUpdateUserGroup(userGroup *UserGroup) error
	UpdateUserGroupField(userGroupName, fields map[string]string) error
	DeleteUserGroup(userGroup *UserGroup) error
	DeleteUserGroupWithName(userGroupName string) error

	GetVariableGroup(varGroupName string, owner string) (*VariableGroup, error)
	GetVariableGroupWithId(varGroupId string) (*VariableGroup, error)
	CreateOrUpdateVariableGroup(variableGroup *VariableGroup) error
	UpdateVariableGroupField(varGroupId string, fields map[string]string) error
	DeleteVariableGroup(varGroup *VariableGroup) error
	DeleteVariableGroupWithId(varGroupId string) error

	AddUserToGroup(userName string, groupName string) error
	RemoveUserFromGroup(userName string, groupName string) error
	GetUsersInGroup(groupName string) ([]*User, error )
	GetGroupsForUser(userName string) ([]*UserGroup, error)

	AddVariableToGroup(variableId string, varGroupId string) error
	RemoveVariableFromGroup(variableId string, varGroupId string) error
	GetVariablesInGroup(groupName string) ([]*Variable, error)
	GetGroupsForVariable(variableId string) ([]*VariableGroup, error)
}
