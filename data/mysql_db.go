package data

type MysqlDb struct {
}

const SQL_GET_ALL_USERS = `
	SELECT * FROM orek_user;
`

const SQL_GET_USER_WITH_EMAIL = `
	SELECT * FROM orek_user WHERE email = "?";
`

const SQL_GET_USER = `
	SELECT * FROM orek_user WHERE user_name = ?;
`

const SQL_CREATE_OR_UPDATE_USER = `
	INSERT INTO orek_user( user_name,
	 					   first_name,
	 					   last_name,
	 					   email )
	VALUES( ?, ?, ?, ? );
`

const SQL_UPDATE_USER_FILED = `
	???
`

const SQL_DELETE_USER = `
	DELETE FROM orek_user WHERE user_name = ?;
`


const SQL_GET_ALL_SOURCE = `
	SELECT * FROM orek_source;
`


func (mdb *MysqlDb) GetUser(userName string) (*User, error) {
	return nil, nil
}

func (mdb *MysqlDb) GetUserWithEmail(email string) (*User, error) {
	return nil, nil
}

func (mdb *MysqlDb) CreateOrUpdateUser(user *User) error {
	return nil
}

func (mdb *MysqlDb) UpdateUserField(userName string, fields map[string]string) error {
	return nil
}

func (mdb *MysqlDb) DeleteUser(userId string) error {
	return nil
}

func (mdb *MysqlDb) GetSource(sourceName string, owner string) (*Source, error) {
	return nil, nil
}

func (mdb *MysqlDb) GetSourceWithId(sourceId string) (*Source, error) {
	return nil, nil
}

func (mdb *MysqlDb) CreateOrUpdateSource(source *Source) error {
	return nil
}

func (mdb *MysqlDb) UpdateSourceField(sourceId string, fields map[string]string) error {
	return nil
}

func (mdb *MysqlDb) DeleteSource(sourceId *Source) error {
	return nil
}

func (mdb *MysqlDb) GetVariable(sourceId string, name string) (*Variable, error) {
	return nil, nil
}

func (mdb *MysqlDb) GetVariableWithId(variableId string) (*Variable, error) {
	return nil, nil
}

func (mdb *MysqlDb) CreateOrUpdateVariable(variable *Variable) error {
	return nil
}

func (mdb *MysqlDb) UpdateVariableField(variableId string, fields map[string]string) error {
	return nil
}

func (mdb *MysqlDb) DeleteVariable(variableId string) error {
	return nil
}

func (mdb *MysqlDb) GetUserGroup(userGroupName string) (*UserGroup, error) {
	return nil, nil
}

func (mdb *MysqlDb) CreateOrUpdateUserGroup(userGroup *UserGroup) error {
	return nil, nil
}

func (mdb *MysqlDb) UpdateUserGroupField(userGroupName, fields map[string]string) error {
	return nil
}

func (mdb *MysqlDb) DeleteUserGroup(userGroupName string) error {
	return nil
}

func (mdb *MysqlDb) GetVariableGroup(varGroupName string, owner string) (*VariableGroup, error) {
	return nil, nil
}

func (mdb *MysqlDb) GetVariableGroupWithId(varGroupId string) (*VariableGroup, error) {
	return nil, nil
}

func (mdb *MysqlDb) CreateOrUpdateVariableGroup(variableGroup *VariableGroup) error {
	return nil
}

func (mdb *MysqlDb) UpdateVariableGroupField(varGroupId string, fields map[string]string) error {
	return nil
}

func (mdb *MysqlDb) DeleteVariableGroup(varGroupId string) error {
	return nil
}

func (mdb *MysqlDb) AddUserToGroup(userName string, groupName string) error {
	return nil
}

func (mdb *MysqlDb) RemoveUserFromGroup(userName string, groupName string) error {
	return nil
}

func (mdb *MysqlDb) GetUsersInGroup(groupName string) ([]*User, error) {
	return nil, nil
}

func (mdb *MysqlDb) GetGroupsForUser(userName string) ([]*UserGroup, error) {
	return nil, nil
}

func (mdb *MysqlDb) AddVariableToGroup(variableId string, varGroupId string) error {
	return nil
}

func (mdb *MysqlDb) RemoveVariableFromGroup(variableId string, varGroupId string) error {
	return nil
}

func (mdb *MysqlDb) GetVariablesInGroup(groupName string) ([]*Variable, error) {
	return nil, nil
}

func (mdb *MysqlDb) GetGroupsForVariable(variableId string) ([]*VariableGroup, error) {
	return nil, nil
}
