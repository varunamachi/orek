package data

import (
	"database/sql"
	"errors"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type MysqlDb struct {
	sql.DB
}

func (mdb *MysqlDb) init(options string) error {
	var err error = nil
	mdb, err = sql.Open("mysql", options)
	if err != nil {
		log.Print("Could not connect to mysql database")
	} else if err = mdb.Ping(); err != nil {
		log.Print("Could not connect to mysql database")
	} else {
		log.Print("Database opened successfuly")
	}
	return err
}

func (mdb *MysqlDb) GetAllUsers() ([]*User, error) {
	const MYSQL_GET_ALL_USERS = `SELECT * FROM orek_user;`
	rows, err := mdb.Query(MYSQL_GET_ALL_USERS)
	users := make([]*User, 10)
	if err == nil {
		defer rows.Close()
		for rows.Next() {
			user := &User{}
			err = rows.Scan(&user.Name,
				&user.FirstName,
				&user.SecondName,
				&user.Email)
			if err != nil {
				log.Print("Error occured while retrieving list of users")
				break
			} else {
				users = append(users, user)
			}
		}
		if err = rows.Err(); err != nil {
			log.Print("Error occured while iterating over the results")
		}
	} else {
		log.Print("Error occured while retrieving list of users")
	}
	return users, err
}

func (mdb *MysqlDb) GetUser(userName string) (*User, error) {
	const MYSQL_GET_USER = `SELECT * FROM orek_user WHERE user_name = ?;`
	stmt, err := mdb.Prepare(MYSQL_GET_USER)
	var user *User
	if err == nil {
		defer stmt.Close()
		var row sql.Row
		row, err := stmt.QueryRow(userName)
		if err == nil {
			err = row.Scan(&user.Name,
				&user.FirstName,
				&user.SecondName,
				&user.Email)
			if err != nil {
				log.Printf(`Error occured while reading info about user with
                        name %s`, userName)
			}
		} else {
			log.Printf(`Error occured while querying info about user %s from
                        database`, userName)
		}
	} else {
		log.Print(`Error occured while preparing query to get user info`)
	}
	return user, err
}

func (mdb *MysqlDb) GetUserWithEmail(email string) (*User, error) {
	const MYSQL_GET_USER_WITH_EMAIL = `SELECT * FROM orek_user
	                                    WHERE email = "?";`
	stmt, err := mdb.Prepare(MYSQL_GET_USER_WITH_EMAIL)
	var user *User
	if err == nil {
		defer stmt.Close()
		var row sql.Row
		row, err := stmt.QueryRow(email)
		if err == nil {
			err = row.Scan(&user.Name,
				&user.FirstName,
				&user.SecondName,
				&user.Email)
			if err != nil {
				log.Printf(`Error occured while reading info about user with
                        email %s`, email)
			}
		} else {
			log.Printf(`Error occured while querying info about user with email
                         %s from database`, email)
		}
	} else {
		log.Print(`Error occured while preparing query to get user info`)
	}
	return user, err
}

func (mdb *MysqlDb) CreateOrUpdateUser(user *User) error {
	const MYSQL_CREATE_OR_UPDATE_USER = `
        INSERT INTO orek_user( user_name,
 	    				   first_name,
 	    				   second_name,
 	    				   email )
 	    VALUES( ?, ?, ?, ? ) ON DUPLICATE KEY UPDATE
            first_name = VALUES( first_name ),
            last_name = VALUES( last_name );
            email = VALUES( email );`
	if user == nil {
		return errors.New("Create/Update user: Invalid object given")
	}
	stmt, err := mdb.Prepare(MYSQL_CREATE_OR_UPDATE_USER)
	if err == nil {
        defer stmt.Close()
		_, err = stmt.Exec(user.Name,
			user.FirstName,
			user.SecondName,
			user.Email)
        if err != nil {
            log.Printf( `Error occured while creating/updating the user with
                        name %s`, user.Name)
        }
	} else {
		log.Printf(`Error occured while preparing statement to create/update
                    user in %s the database`, user.Name)
	}
	return err
}

func (mdb *MysqlDb) DeleteUser(userId string) error {
	const MYSQL_DELETE_USER = `DELETE FROM orek_user WHERE user_name = ?;`
    stmt, err := mdb.Prepare(MYSQL_DELETE_USER)
    if err == nil {
        defer stmt.Close()
        _, err = stmt.Exec(userId)
        if err != nil {
            log.Printf(`Error occured while deleting information about user
                        with name %s from database`, userId)
        }
    } else {
        log.Printf(`Error occured while preparing query to delete user %s`,
                    userId)
    }
	return err
}

func (mdb *MysqlDb) GetAllSources() ([]*Source, error) {
	const MYSQL_GET_ALL_SOURCES = `SELECT * FROM orek_source;`
    sources := make([]*Source, 10)
    stmt, err := mdb.Prepare(MYSQL_GET_ALL_SOURCES)
    if err == nil {
        defer stmt.Close()
        var rows sql.Rows
        rows, err = stmt.Query()
        if err == nil {
            defer rows.Close()
            for rows.Next() {
                source := &Source{}
                err = rows.Scan( &source.SourceId,
                    &source.Name,
                    &source.Owner,
                    &source.Description,
                    &source.Location,
                    &source.Access)
                if err == nil {
                    source = append(sources, source)
                } else {
                    log.Printf(`Error occured while getting the list of
                                sources from the database`)
                    break
                }
            }
            if err = rows.Err(); err != nil {
                log.Printf(`Error occured while finishing the source list
                            fetch operation`)
            }
        } else {
            log.Printf(`Error occured while querying for the list of
                                sources from the database`)
        }
    } else {
        log.Printf(`Error occured while preparing query to get the list of
                     sources`)
    }
	return err, sources
}

func (mdb *MysqlDb) GetSource(sourceName, owner string) (*Source, error) {
	const MYSQL_GET_SOURCE = `SELECT * FROM orek_source
	                            WHERE name = ? AND owner = ?;`
    source := &Source{}
    stmt, err := mdb.Prepare(MYSQL_GET_SOURCE)
    if err == nil {
        defer stmt.Close()
        var row sql.Row
        row, err = stmt.QueryRow(sourceName, owner)
        if err == nil {
            err = row.Scan( &source.SourceId,
                &source.Name,
                &source.Owner,
                &source.Description,
                &source.Location,
                &source.Access)
            if err != nil {
                log.Printf(`Error occured while getting info about source
                            with name '%s' owned by '%s'`, sourceName, owner)
            }
        } else {
            log.Printf(`Error occured while querying for info about source
                    with name '%s' owned by '%s'`, sourceName, owner)
        }
    } else {
        log.Printf(`Error occured while preparing query to get info about source
                    with name '%s' owned by '%s'`, sourceName, owner)
    }
    return err, source
}

func (mdb *MysqlDb) GetSourceWithId(sourceId string) (*Source, error) {
	const MYSQL_GET_SOURCE_WITH_ID = `SELECT * FROM orek_source WHERE
	                                  source_id = ?;`

    source := &Source{}
    stmt, err := mdb.Prepare(MYSQL_GET_SOURCE_WITH_ID)
    if err == nil {
        defer stmt.Close()
        var row sql.Row
        row, err = stmt.QueryRow(sourceId)
        if err == nil {
            err = row.Scan( &source.SourceId,
                &source.Name,
                &source.Owner,
                &source.Description,
                &source.Location,
                &source.Access)
            if err != nil {
                log.Printf(`Error occured while getting info about source
                            with id '%s'`, sourceId)
            }
        } else {
            log.Printf(`Error occured while querying for info about source
                    with id '%s'`, sourceId)
        }
    } else {
        log.Printf(`Error occured while preparing query to get info about source
                    with id '%s'`, sourceId)
    }
    return err, source
}

func (mdb *MysqlDb) CreateOrUpdateSource(source *Source) error {
	const MYSQL_CREATE_OR_UPDATE_SOURCE = `INSERT INTO orek_source( source_id,
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

    if source == nil {
        return errors.New("Create/Update Source: Invalid object given")
    }
    stmt, err := mdb.Prepare(MYSQL_CREATE_OR_UPDATE_SOURCE)
    if err == nil {
        defer stmt.Close()
        _, err = stmt.Exec(source.SourceId,
            source.Name,
            source.Owner,
            source.Description,
            source.Location,
            source.Access)
        if err != nil {
            log.Printf( `Error occured while creating/updating the source with
                        id %s`, source.SourceId)
        }
    } else {
        log.Printf(`Error occured while preparing statement to create/update
                    source with id %s the database`, source.SourceId)
    }
    return err
}

func (mdb *MysqlDb) DeleteSource(sourceId string) error {
	const MYSQL_DELETE_SOURCE = `DELETE FROM orek_source WHERE source_id = ?;`
    stmt, err := mdb.Prepare(MYSQL_DELETE_SOURCE)
    if err == nil {
        defer stmt.Close()
        _, err = stmt.Exec(sourceId)
        if err != nil {
            log.Printf(`Error occured while deleting information about source
                        with id  %s from database`, sourceId)
        }
    } else {
        log.Printf(`Error occured while preparing query to delete source with
                    id %s`, sourceId)
    }
    return err
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
	const MYSQL_GET_VARIABLE_WITH_ID = `SELECT * FROM orek_variable WHERE variable_id = ?;`
	return nil, nil
}

func (mdb *MysqlDb) CreateOrUpdateVariable(variable *Variable) error {
	const MYSQL_CREATE_OR_UPDATE_VARIABLE = `INSERT INTO orek_variable( variable_id,
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
	const MYSQL_DELETE_VARIABLE = `DELETE FROM orek_variable WHERE variable_id = ?;`
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
	const MYSQL_CREATE_OR_UPDATE_USER_GROUP = `INSERT INTO orek_user_group( name,
                             owner,
                             description )
        VALUES( ?, ?, ? )
        ON DUPLICATE KEY UPDATE
            owner = VALUES( owner ),
            description = VALUES( description );`
	return nil, nil
}

func (mdb *MysqlDb) DeleteUserGroup(userGroupName string) error {
	const MYSQL_DELETE_USER_GROUP = `DELETE FROM orek_user_group WHERE name = ?;`
	return nil
}

func (mdb *MysqlDb) GetAllVariableGroups() ([]*VariableGroup, error) {
	const MYSQL_GET_ALL_VARIABLE_GROUPS = `SELECT * FROM orek_variable_group;`
	return nil, nil
}

func (mdb *MysqlDb) GetVariableGroup(varGroupName, owner string) (*VariableGroup, error) {
	const MYSQL_GET_VARIABLE_GROUP = `SELECT * FROM orek_variable_group WHERE name = ? AND owner = ?;`
	return nil, nil
}

func (mdb *MysqlDb) GetVariableGroupWithId(varGroupId string) (*VariableGroup, error) {
	const MYSQL_GET_VARIABLE_WITH_GROUP = `SELECT * FROM orek_variable_group WHERE group_id = ?;`
	return nil, nil
}

func (mdb *MysqlDb) CreateOrUpdateVariableGroup(variableGroup *VariableGroup) error {
	const CREATE_OR_UPDATE_VARIABLE_GROUP = `INSERT INTO orek_variable_group( group_id,
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
	const MYSQL_REMOVE_VARIABLE_GROUP = `DELETE FROM orek_variable_group WHERE group_id = ?;`
	return nil
}

func (mdb *MysqlDb) AddUserToGroup(userName string, groupName string) error {
	const MYSQL_ADD_USER_TO_GROUP = `INSERT IGNORE INTO orek_user_to_group( group_name, user_name )
            VALUES( ?, ? );`
	return nil
}

func (mdb *MysqlDb) RemoveUserFromGroup(userName string, groupName string) error {
	const MYSQL_REMOVE_USER_FROM_GROUP = `DELETE FROM orek_user_to_group WHERE user_name = ? AND group_name = ?;`
	return nil
}

func (mdb *MysqlDb) GetUsersInGroup(groupName string) ([]*User, error) {
	const MYSQL_GET_USERS_IN_GROUP = `SELECT user_name FROM orek_user_to_group WHERE group_name = ?;`
	return nil, nil
}

func (mdb *MysqlDb) GetGroupsForUser(userName string) ([]*UserGroup, error) {
	const MYSQL_GET_GROUPS_FOR_USER = `SELECT group_name FROM orek_user_to_group WHERE user_name = ?;`
	return nil, nil
}

func (mdb *MysqlDb) AddVariableToGroup(variableId string, varGroupId string) error {
	const MYSQL_ADD_VARIABLE_TO_GROUP = `INSERT IGNORE INTO orek_variable_to_group(
            var_group_id, variable_name ) VALUES( ? ? );`
	return nil
}

func (mdb *MysqlDb) RemoveVariableFromGroup(variableId string, varGroupId string) error {
	const MYSQL_REMOVE_VARIABLE_FROM_GROUP = `DELETE FROM orek_variable_to_group WHERE var_group_id = ?
            AND variable_id = ?;`
	return nil
}

func (mdb *MysqlDb) GetVariablesInGroup(groupName string) ([]*Variable, error) {
	const MYSQL_GET_VARIABLE_IN_GROUP = `SELECT variable_id FROM orek_variable_to_group WHERE var_group_id = ?;`
	return nil, nil
}

func (mdb *MysqlDb) GetGroupsForVariable(variableId string) ([]*VariableGroup, error) {
	const MYSQL_GET_GROUPS_FOR_VARIABLE = `SELECT var_group_id FROM orek_variable_to_group WHERE variable_id = ?;`
	return nil, nil
}

func (mdb *MysqlDb) AddVariableValue(variableId, value string) error {
	const MYSQL_ADD_VARIABLE_VALUE = `INSERT INTO orek_variable_value( variable_id, value ) VALUES( ?, ? );`
	return nil
}

func (mdb *MysqlDb) ClearValuesForVariable(variableId string) error {
	const MYSQL_CLEAR_VALUES_FOR_VARIABLE = `DELETE FROM orek_variable_value WHERE variable_id = ?;`
	return nil
}

func (mdb *MysqlDb) GetValuesForVariable(variableId string) error {
	const MYSQL_GET_VALUES_FOR_VARIABLE = `DELETE FROM orek_variable_value WHERE variable_id = ?;`
	return nil
}
