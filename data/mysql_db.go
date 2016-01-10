package data

type MysqlDb struct {
}

func (mdb *MysqlDb) GetAllUsers() ([]*User, error) {
	const MYSQL_GET_ALL_USERS = `SELECT * FROM orek_user;`
	return nil, nil
}

func (mdb *MysqlDb) GetUser(userName string) (*User, error) {
	const MYSQL_GET_USER = `SELECT * FROM orek_user WHERE user_name = ?;`
	return nil, nil
}

func (mdb *MysqlDb) GetUserWithEmail(email string) (*User, error) {
	const MYSQL_GET_USER_WITH_EMAIL = `SELECT * FROM orek_user WHERE email = "?";`
	return nil, nil
}

func (mdb *MysqlDb) CreateOrUpdateUser(user *User) error {
	const MYSQL_CREATE_OR_UPDATE_USER = `
        INSERT INTO orek_user( user_name,
 	    				   first_name,
 	    				   last_name,
 	    				   email )
 	    VALUES( ?, ?, ?, ? ) ON DUPLICATE KEY UPDATE
            first_name = VALUES( first_name ),
            last_name = VALUES( last_name );
            email = VALUES( email );`
	return nil
}

func (mdb *MysqlDb) DeleteUser(userId string) error {
	const MYSQL_DELETE_USER = `DELETE FROM orek_user WHERE user_name = ?;`
	return nil
}

func (mdb *MysqlDb) GetAllSources() ([]*Source, error) {
	const MYSQL_GET_ALL_SOURCES = `SELECT * FROM orek_source;`
	return nil, nil
}

func (mdb *MysqlDb) GetSource(sourceName, owner string) (*Source, error) {
	const MYSQL_GET_SOURCE = `SELECT * FROM orek_source
	                            WHERE name = ? AND owner = ?;`
	return nil, nil
}

func (mdb *MysqlDb) GetSourceWithId(sourceId string) (*Source, error) {
	const MYSQL_GET_SOURCE_WITH_ID = `SELECT * FROM orek_source WHERE source_id = ?;`
	return nil, nil
}

func (mdb *MysqlDb) CreateOrUpdateSource(source *Source) error {
    const MYSQL_CREATE_OR_UPDATE_SOURCE =
        `INSERT INTO orek_source( source_id,
                             name,
                             owner,
                             description,
                             location,
                             access )
        VALUES( ?, ?, ?, ?, ?, ? )
        ON DUPLICATE KEY UPDATE
            name = VALUES( name ),
            owner = VALUES( owner ),
            description = VALUES( description ),
            location = VALUES( location ),
            access = VALUES( access );`
	return nil
}

func (mdb *MysqlDb) DeleteSource(sourceId *Source) error {
    const MYSQL_DELETE_SOURCE = `DELETE FROM orek_source WHERE source_id = ?;`
	return nil
}

func (mdb *MysqlDb) GetAllVariables() ([]*Variable, error) {
    const MYSQL_GET_ALL_VARIABLE = `SELECT * FROM orek_variable;`
    return nil, nil
}

func (mdb *MysqlDb) GetVariable(sourceId, name string) (*Variable, error) {
    const MYSQL_GET_VARIABLE = `SELECT * FROM orek_variable WHERE name = ?
                                    AND source_id = ?;`
	return nil, nil
}

func (mdb *MysqlDb) GetVariableWithId(variableId string) (*Variable, error) {
    const MYSQL_GET_VARIABLE_WITH_ID =
        `SELECT * FROM orek_variable WHERE variable_id = ?;`
	return nil, nil
}

func (mdb *MysqlDb) CreateOrUpdateVariable(variable *Variable) error {
    const MYSQL_CREATE_OR_UPDATE_VARIABLE =
        `INSERT INTO orek_variable( variable_id,
                               name,
                               source_id,
                               description,
                               unit )
            VALUES( ?, ?, ?, ?, ? )
            ON DUPLICATE KEY UPDATE
                name = VALUES( name ),
                source_id = VALUES( source_id ),
                description = VALUES( description ),
                unit = VALUES( unit );`
	return nil
}

func (mdb *MysqlDb) DeleteVariable(variableId string) error {
    const MYSQL_DELETE_VARIABLE =
        `DELETE FROM orek_variable WHERE variable_id = ?;`
	return nil
}

func (mdb *MysqlDb) GetAllUserGroups() ([]*UserGroup, error) {
    const MYSQL_GET_ALL_USER_GROUPS = `SELECT * FROM orek_user_group;`
    return nil, nil
}

func (mdb *MysqlDb) GetUserGroup(userGroupName string) (*UserGroup, error) {
    const MYSQL_GET_USER_GROUP = `SELECT * FROM orek_user_group WHERE name = ?;`
	return nil, nil
}

func (mdb *MysqlDb) CreateOrUpdateUserGroup(userGroup *UserGroup) error {
    const MYSQL_CREATE_OR_UPDATE_USER_GROUP  =
    `INSERT INTO orek_user_group( name,
                             owner,
                             description )
        VALUES( ?, ?, ? )
        ON DUPLICATE KEY UPDATE
            owner = VALUES( owner ),
            description = VALUES( description );`
	return nil, nil
}

func (mdb *MysqlDb) DeleteUserGroup(userGroupName string) error {
    const MYSQL_DELETE_USER_GROUP =
        `DELETE FROM orek_user_group WHERE name = ?;`
	return nil
}

func (mdb *MysqlDb) GetAllVariableGroups() ([]*VariableGroup, error) {
    const MYSQL_GET_ALL_VARIABLE_GROUPS = `SELECT * FROM orek_variable_group;`
    return  nil, nil
}

func (mdb *MysqlDb) GetVariableGroup(varGroupName, owner string) (*VariableGroup, error) {
    const MYSQL_GET_VARIABLE_GROUP =
        `SELECT * FROM orek_variable_group WHERE name = ? AND owner = ?;`
	return nil, nil
}

func (mdb *MysqlDb) GetVariableGroupWithId(varGroupId string) (*VariableGroup, error) {
    const MYSQL_GET_VARIABLE_WITH_GROUP =
        `SELECT * FROM orek_variable_group WHERE group_id = ?;`
	return nil, nil
}

func (mdb *MysqlDb) CreateOrUpdateVariableGroup(variableGroup *VariableGroup) error {
    const CREATE_OR_UPDATE_VARIABLE_GROUP =
    `INSERT INTO orek_variable_group( group_id,
                                 name,
                                 owner,
                                 description,
                                 access )
    VALUES( ?, ?, ?, ?, ? )
    ON DUPLICATE KEY UPDATE
        name = VALUES( name ),
        owner = VALUES( owner ),
        description = VALUES( description ),
        access = VALUES( access );`
	return nil
}

func (mdb *MysqlDb) DeleteVariableGroup(varGroupId string) error {
    const MYSQL_REMOVE_VARIABLE_GROUP =
        `DELETE FROM orek_variable_group WHERE group_id = ?;`
	return nil
}

func (mdb *MysqlDb) AddUserToGroup(userName string, groupName string) error {
    const MYSQL_ADD_USER_TO_GROUP =
        `INSERT IGNORE INTO orek_user_to_group( group_name, user_name )
            VALUES( ?, ? );`
	return nil
}

func (mdb *MysqlDb) RemoveUserFromGroup(userName string, groupName string) error {
    const MYSQL_REMOVE_USER_FROM_GROUP =
        `DELETE FROM orek_user_to_group WHERE user_name = ? AND group_name = ?;`
	return nil
}

func (mdb *MysqlDb) GetUsersInGroup(groupName string) ([]*User, error) {
    const MYSQL_GET_USERS_IN_GROUP =
        `SELECT user_name FROM orek_user_to_group WHERE group_name = ?;`
	return nil, nil
}

func (mdb *MysqlDb) GetGroupsForUser(userName string) ([]*UserGroup, error) {
    const MYSQL_GET_GROUPS_FOR_USER =
         `SELECT group_name FROM orek_user_to_group WHERE user_name = ?;`
	return nil, nil
}

func (mdb *MysqlDb) AddVariableToGroup(variableId string, varGroupId string) error {
    const MYSQL_ADD_VARIABLE_TO_GROUP =
        `INSERT IGNORE INTO orek_variable_to_group(
            var_group_id, variable_name ) VALUES( ? ? );`
	return nil
}

func (mdb *MysqlDb) RemoveVariableFromGroup(variableId string, varGroupId string) error {
    const MYSQL_REMOVE_VARIABLE_FROM_GROUP =
        `DELETE FROM orek_variable_to_group WHERE var_group_id = ?
            AND variable_id = ?;`
	return nil
}

func (mdb *MysqlDb) GetVariablesInGroup(groupName string) ([]*Variable, error) {
    const MYSQL_GET_VARIABLE_IN_GROUP =
        `SELECT variable_id FROM orek_variable_to_group WHERE var_group_id = ?;`
	return nil, nil
}

func (mdb *MysqlDb) GetGroupsForVariable(variableId string) ([]*VariableGroup, error) {
    const MYSQL_GET_GROUPS_FOR_VARIABLE =
        `SELECT var_group_id FROM orek_variable_to_group WHERE variable_id = ?;`
	return nil, nil
}

func (mdb *MysqlDb) AddVariableValue(variableId, value string) error {
    const MYSQL_ADD_VARIABLE_VALUE =
        `INSERT INTO orek_variable_value( variable_id, value ) VALUES( ?, ? );`
    return nil
}

func (mdb *MysqlDb) ClearValuesForVariable(variableId string) error {
    const MYSQL_CLEAR_VALUES_FOR_VARIABLE =
        `DELETE FROM orek_variable_value WHERE variable_id = ?;`
    return nil
}

func (mdb *MysqlDb) GetValuesForVariable(variableId string) error {
    const MYSQL_GET_VALUES_FOR_VARIABLE =
        `DELETE FROM orek_variable_value WHERE variable_id = ?;`
    return nil
}
