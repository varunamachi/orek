package data

var db OrekDb = nil

type OrekDb interface {
	GetUser(userName string) *User
	GetUserWithEmail(email string) *User
	CreateOrUpdateUser(user *User) error
	UpdateUserField(userName string, fields map[string]string) error
	DeleteUser(user *User) error
	DeleteUserWithId(userId string) error

	GetSource(sourceName string, owner string) *Source
	GetSourceWithId(sourceId string) *Source
	CreateOrUpdateSource(source *Source) error
	UpdateSourceField(sourceId string, fields map[string]string) error
	DeleteSource(source *Source) error
	DeleteSourceWithId(sourceId *Source) error

	GetVariable(sourceId string, name string) *Variable
	GetVariableWithId(variableId string) *Variable
	CreateOrUpdateVariable(variable *Variable) error
	UpdateVariableField(variableId string, fields map[string]string) error
	DeleteVariable(variable *Variable) error
	DeleteVariableWithId(variableId string) error

	GetUserGroup(userGroupName string) UserGroup
	CreateOrUpdateUserGroup(userGroup *UserGroup) error
	UpdateUserGroupField(userGroupName, fields map[string]string) error
	DeleteUserGroup(userGroup *UserGroup) error
	DeleteUserGroupWithName(userGroupName string) error

	GetVariableGroup(varGroupName string, owner string) VariableGroup
	GetVariableGroupWithId(varGroupId string) VariableGroup
	CreateOrUpdateVariableGroup(variableGroup *VariableGroup) error
	UpdateVariableGroupField(varGroupId string, fields map[string]string) error
	DeleteVariableGroup(varGroup *VariableGroup) error
	DeleteVariableGroupWithId(varGroupId string) error
}
