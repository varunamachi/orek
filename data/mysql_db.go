package data

import (
	"database/sql"
	"errors"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type MysqlDb struct {
	*sql.DB
}

func MySqlInit(options string) (*MysqlDb, error) {
	mdb, err := sql.Open("mysql", options)
	if err != nil {
		log.Print("Could not connect to mysql database")
	} else if err = mdb.Ping(); err != nil {
		log.Print("Could not connect to mysql database")
	} else {
		log.Print("Database opened successfuly")
	}
	return &MysqlDb{mdb}, err
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
		row := stmt.QueryRow(userName)
		err = row.Scan(&user.Name,
			&user.FirstName,
			&user.SecondName,
			&user.Email)
		if err != nil {
			log.Printf(`Error occured while reading info about user with
                        name %s`, userName)
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
		row := stmt.QueryRow(email)
		err = row.Scan(&user.Name,
			&user.FirstName,
			&user.SecondName,
			&user.Email)
		if err != nil {
			log.Printf(`Error occured while reading info about user with
                        email %s`, email)
		}
		log.Printf(`Error occured while querying info about user with email
                         %s from database`, email)
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
			log.Printf(`Error occured while creating/updating the user with
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
		var rows *sql.Rows
		rows, err = stmt.Query()
		if err == nil {
			defer rows.Close()
			for rows.Next() {
				source := &Source{}
				err = rows.Scan(&source.SourceId,
					&source.Name,
					&source.Owner,
					&source.Description,
					&source.Location,
					&source.Access)
				if err == nil {
					sources = append(sources, source)
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
	return sources, err
}

func (mdb *MysqlDb) GetSource(sourceName, owner string) (*Source, error) {
	const MYSQL_GET_SOURCE = `SELECT * FROM orek_source
	                            WHERE name = ? AND owner = ?;`
	source := &Source{}
	stmt, err := mdb.Prepare(MYSQL_GET_SOURCE)
	if err == nil {
		defer stmt.Close()
		row := stmt.QueryRow(sourceName, owner)
		err = row.Scan(&source.SourceId,
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
		log.Printf(`Error occured while preparing query to get info about source
                    with name '%s' owned by '%s'`, sourceName, owner)
	}
	return source, err
}

func (mdb *MysqlDb) GetSourceWithId(sourceId string) (*Source, error) {
	const MYSQL_GET_SOURCE_WITH_ID = `SELECT * FROM orek_source WHERE
	                                  source_id = ?;`

	source := &Source{}
	stmt, err := mdb.Prepare(MYSQL_GET_SOURCE_WITH_ID)
	if err == nil {
		defer stmt.Close()
		row := stmt.QueryRow(sourceId)
		err = row.Scan(&source.SourceId,
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
		log.Printf(`Error occured while preparing query to get info about source
                    with id '%s'`, sourceId)
	}
	return source, err
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
			log.Printf(`Error occured while creating/updating the source with
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
	variables := make([]*Variable, 10)
	const MYSQL_GET_ALL_VARIABLE = `SELECT * FROM orek_variable;`
	stmt, err := mdb.Prepare(MYSQL_GET_ALL_VARIABLE)
	if err == nil {
		defer stmt.Close()
		var rows *sql.Rows
		rows, err = stmt.Query()
		if err == nil {
			defer rows.Close()
			for rows.Next() {
				variable := &Variable{}
				err = rows.Scan(&variable.VariableId,
					&variable.Name,
					&variable.SourceId,
					&variable.Description,
					&variable.Unit)
				if err == nil {
					variables = append(variables, variable)
				} else {
					log.Printf(`Error occured while retrieving list of variables
					             from the database`)
					break
				}
			}
			if err = rows.Err(); err != nil {
				log.Printf(`Error occured while retrieving list of variables
				            from the database`)
			}
		} else {
			log.Printf(`Error while querying list of variables from the database`)
		}
	} else {
		log.Printf(`Failed to get the list of variables`)
	}
	return variables, err
}

func (mdb *MysqlDb) GetVariable(sourceId, name string) (*Variable, error) {
	const MYSQL_GET_VARIABLE = `SELECT * FROM orek_variable WHERE name = ?
                                    AND source_id = ?;`
	stmt, err := mdb.Prepare(MYSQL_GET_VARIABLE)
	variable := &Variable{}
	if err == nil {
		defer stmt.Close()
		row := stmt.QueryRow(sourceId, name)
		err = row.Scan(&variable.VariableId,
			&variable.Name,
			&variable.SourceId,
			&variable.Description,
			&variable.Unit)
		if err != nil {
			log.Printf(`Error occured while retrieving variable %s:%s`,
				sourceId, name)
		}

	} else {
		log.Printf(`Failed to get variable %s:%s from database`, sourceId, name)
	}
	return variable, err
}

func (mdb *MysqlDb) GetVariableWithId(variableId string) (*Variable, error) {
	const MYSQL_GET_VARIABLE_WITH_ID = `SELECT * FROM orek_variable WHERE
	                                    variable_id = ?;`
	stmt, err := mdb.Prepare(MYSQL_GET_VARIABLE_WITH_ID)
	variable := &Variable{}
	if err == nil {
		defer stmt.Close()
		row := stmt.QueryRow(variableId)
		err = row.Scan(&variable.VariableId,
			&variable.Name,
			&variable.SourceId,
			&variable.Description,
			&variable.Unit)
		if err != nil {
			log.Printf(`Error occured while retrieving variable %s`,
				variableId)
		}

	} else {
		log.Printf(`Failed to get variable %s from database`, variableId)
	}
	return variable, err
}

func (mdb *MysqlDb) CreateOrUpdateVariable(variable *Variable) error {
	const MYSQL_CREATE_OR_UPDATE_VARIABLE = `INSERT INTO orek_variable(
	                           variable_id,
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
	if variable == nil {
		return errors.New("Create/Update Variable: Invalid object given")
	}
	stmt, err := mdb.Prepare(MYSQL_CREATE_OR_UPDATE_VARIABLE)
	if err == nil {
		defer stmt.Close()
		_, err = stmt.Exec(variable.VariableId,
			variable.Name,
			variable.SourceId,
			variable.Description,
			variable.Unit)
		if err != nil {
			log.Printf(`Error occured while creating/updating variable
			            with id %s`, variable.VariableId)
		}
	} else {
		log.Printf(`Error occured while preparing query for create/update
		            variable %s`, variable.VariableId)
	}
	return err
}

func (mdb *MysqlDb) DeleteVariable(variableId string) error {
	const MYSQL_DELETE_VARIABLE = `DELETE FROM orek_variable WHERE variable_id = ?;`
	stmt, err := mdb.Prepare(MYSQL_DELETE_VARIABLE)
	if err == nil {
		defer stmt.Close()
		_, err = stmt.Exec(variableId)
		if err != nil {
			log.Printf(`Failed to delete variable with id %s`, variableId)
		}
	} else {
		log.Printf(`Error occured while preparing query to delete variable with id %s`, variableId)
	}
	return err
}

func (mdb *MysqlDb) GetAllUserGroups() ([]*UserGroup, error) {
	const MYSQL_GET_ALL_USER_GROUPS = `SELECT * FROM orek_user_group;`
	groups := make([]*UserGroup, 10)
	stmt, err := mdb.Prepare(MYSQL_GET_ALL_USER_GROUPS)
	if err == nil {
		defer stmt.Close()
		var rows *sql.Rows
		rows, err = stmt.Query()
		if err == nil {
			defer rows.Close()
			for rows.Next() {
				ugrp := &UserGroup{}
				err = rows.Scan(&ugrp.Name, &ugrp.Owner, ugrp.Description)
				if err == nil {
					groups = append(groups, ugrp)
				} else {
					log.Printf(`Error occured while reading user group list from database`)
					break
				}
			}
			if err = rows.Err(); err != nil {
				log.Printf(`Error occured while reading list of user group from database`)
			}
		} else {
			log.Printf(`Failed to fetch user group list from DB `)
		}
	} else {
		log.Printf(`Failed to prepare query for getting the list of user groups`)
	}
	return groups, err
}

func (mdb *MysqlDb) GetUserGroup(userGroupName string) (*UserGroup, error) {
	const MYSQL_GET_USER_GROUP = `SELECT * FROM orek_user_group WHERE name = ?;`
	ugrp := &UserGroup{}
	stmt, err := mdb.Prepare(MYSQL_GET_USER_GROUP)
	if err == nil {
		defer stmt.Close()
		row := stmt.QueryRow()
		err = row.Scan(&ugrp.Name, &ugrp.Owner, ugrp.Description)
		if err != nil {
			log.Printf(`Error occured while reading info about user
				            group %s`, userGroupName)
		}
	} else {
		log.Printf(`Failed to prepare query for getting the info about user
		            group %s`, userGroupName)
	}
	return ugrp, err
}

func (mdb *MysqlDb) CreateOrUpdateUserGroup(userGroup *UserGroup) error {
	const MYSQL_CREATE_OR_UPDATE_USER_GROUP = `INSERT INTO orek_user_group(
	                        name,
                            owner,
                            description )
        VALUES( ?, ?, ? )
        ON DUPLICATE KEY UPDATE
            owner = VALUES( owner ),
            description = VALUES( description );`
	if userGroup == nil {
		return errors.New("Create/Update User Group: Invalid object given")
	}
	stmt, err := mdb.Prepare(MYSQL_CREATE_OR_UPDATE_USER_GROUP)
	if err == nil {
		defer stmt.Close()
		_, err = stmt.Exec(userGroup.Name,
			userGroup.Owner,
			userGroup.Description)
		if err != nil {
			log.Printf(`Error occured while creating/updating user group '%s'`,
				userGroup)
		}
	} else {
		log.Printf(`Error while preparing query for creating/updating user
                    group with name %s`, userGroup.Name)
	}
	return err
}

func (mdb *MysqlDb) DeleteUserGroup(userGroupName string) error {
	const MYSQL_DELETE_USER_GROUP = `DELETE FROM orek_user_group WHERE
	                                 name = ?;`
	stmt, err := mdb.Prepare(MYSQL_DELETE_USER_GROUP)
	if err == nil {
		defer stmt.Close()
		_, err = stmt.Exec(userGroupName)
		if err != nil {
			log.Printf(`Failed to delete user group '%s'`, userGroupName)
		}
	} else {
		log.Printf(`Failed to prepare query to delete user group '%s'`,
			userGroupName)
	}
	return err
}

func (mdb *MysqlDb) GetAllVariableGroups() ([]*VariableGroup, error) {
	const MYSQL_GET_ALL_VARIABLE_GROUPS = `SELECT * FROM orek_variable_group;`
	stmt, err := mdb.Prepare(MYSQL_GET_ALL_VARIABLE_GROUPS)
	varGroups := make([]*VariableGroup, 10)
	if err != nil {
		defer stmt.Close()
		var rows *sql.Rows
		rows, err = stmt.Query()
		if err == nil {
			defer rows.Close()
			for rows.Next() {
				group := &VariableGroup{}
				err = rows.Scan(&group.VarGroupId,
					&group.Name,
					&group.Owner,
					&group.Description,
					&group.Access)
				if err == nil {
					varGroups = append(varGroups, group)
				} else {
					log.Printf(`Error occured while reading list of variable groups`)
					break
				}
			}
			if err = rows.Err(); err != nil {
				log.Printf(`Error occured while reading list of variable groups`)
			}
		} else {
			log.Printf(`Failed to fetch list variable groups`)
		}

	} else {
		log.Printf(`Error while preparing query to fetch variable group list`)
	}
	return varGroups, err
}

func (mdb *MysqlDb) GetVariableGroup(varGroupName, owner string) (*VariableGroup, error) {
	const MYSQL_GET_VARIABLE_GROUP = `SELECT * FROM orek_variable_group WHERE name = ? AND owner = ?;`
	var varGroup *VariableGroup
	stmt, err := mdb.Prepare(MYSQL_GET_VARIABLE_GROUP)
	if err != nil {
		defer stmt.Close()
		row := stmt.QueryRow(varGroupName, owner)
		err = row.Scan(&varGroup.VarGroupId,
			&varGroup.Name,
			&varGroup.Owner,
			&varGroup.Description,
			&varGroup.Access)
		if err != nil {
			log.Printf(`Error occured while reading variable group '%s' owned by %s`,
				varGroupName, owner)
		}
	} else {
		log.Printf(`Error while preparing query to fetch variable group '%s' owned by %s`,
			varGroupName, owner)
	}
	return varGroup, err
}

func (mdb *MysqlDb) GetVariableGroupWithId(varGroupId string) (*VariableGroup, error) {
	const MYSQL_GET_VARIABLE_WITH_GROUP = `SELECT * FROM orek_variable_group WHERE group_id = ?;`
	var varGroup *VariableGroup
	stmt, err := mdb.Prepare(MYSQL_GET_VARIABLE_WITH_GROUP)
	if err != nil {
		defer stmt.Close()
		row := stmt.QueryRow(varGroupId)
		err = row.Scan(&varGroup.VarGroupId,
			&varGroup.Name,
			&varGroup.Owner,
			&varGroup.Description,
			&varGroup.Access)
		if err != nil {
			log.Printf(`Error occured while reading variable group '%s'`, varGroupId)
		}
	} else {
		log.Printf(`Error while preparing query to fetch variable group '%s'`, varGroupId)
	}
	return varGroup, err
}

func (mdb *MysqlDb) CreateOrUpdateVariableGroup(
	variableGroup *VariableGroup) error {
	if variableGroup == nil {
		return errors.New(`Failed to create/update variable group:
		    Invalid object give`)
	}
	const CREATE_OR_UPDATE_VARIABLE_GROUP = `INSERT INTO orek_variable_group(
	                             group_id,
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
	stmt, err := mdb.Prepare(CREATE_OR_UPDATE_VARIABLE_GROUP)
	if err == nil {
		defer stmt.Close()
		_, err = stmt.Exec(variableGroup.VarGroupId,
			variableGroup.Name,
			variableGroup.Owner,
			variableGroup.Description,
			variableGroup.Access)
		if err != nil {
			log.Printf(`Failed to create/update variable group '%s' in DB`,
				variableGroup.VarGroupId)
		}
	} else {
		log.Printf(`Error occured while preparing query for creating/updating
		    user group %`, variableGroup.VarGroupId)
	}
	return err
}

func (mdb *MysqlDb) DeleteVariableGroup(varGroupId string) error {
	const MYSQL_REMOVE_VARIABLE_GROUP = `DELETE FROM orek_variable_group
	    WHERE group_id = ?;`
	stmt, err := mdb.Prepare(MYSQL_REMOVE_VARIABLE_GROUP)
	if err == nil {
		defer stmt.Close()
		_, err = stmt.Exec(varGroupId)
		if err != nil {
			log.Printf(`Error occured while deleting variable group '%s`,
				varGroupId)
		}
	} else {
		log.Printf(`Error occured while preparing query for deleing a
		    variable group %s`, varGroupId)
	}
	return err
}

func (mdb *MysqlDb) AddUserToGroup(userName, groupName string) error {
	const MYSQL_ADD_USER_TO_GROUP = `INSERT IGNORE INTO orek_user_to_group(
									group_name,
									user_name ) VALUES( ?, ? );`
	stmt, err := mdb.Prepare(MYSQL_ADD_USER_TO_GROUP)
	if err == nil {
		defer stmt.Close()
		_, err = mdb.Exec(groupName, userName)
		if err != nil {
			log.Printf(`Error occured while adding user '%s' to group '%s'`,
				userName, groupName)
		}
	} else {
		log.Printf(`Error occured while preparing query to add user '%s'
                    to group '%s'`, userName, groupName)
	}
	return err
}

func (mdb *MysqlDb) RemoveUserFromGroup(userName string,
	groupName string) error {
	const MYSQL_REMOVE_USER_FROM_GROUP = `DELETE FROM orek_user_to_group
	                                WHERE user_name = ? AND group_name = ?;`
	stmt, err := mdb.Prepare(MYSQL_REMOVE_USER_FROM_GROUP)
	if err == nil {
		defer stmt.Close()
		_, err = mdb.Exec(userName, groupName)
		if err != nil {
			log.Printf(`Error occured while removing user '%s' from group '%s'`,
				userName, groupName)
		}
	} else {
		log.Printf(`Error occured while preparing query to remove user '%s'
                    from group '%s'`, userName, groupName)
	}
	return err
}

func (mdb *MysqlDb) GetUsersInGroup(groupName string) ([]*User, error) {
	const MYSQL_GET_USERS_IN_GROUP = `SELECT user_name FROM orek_user_to_group
	                                    WHERE group_name = ?;`
	stmt, err := mdb.Prepare(MYSQL_GET_USERS_IN_GROUP)
	users := make([]*User, 10)
	if err == nil {
		defer stmt.Close()
		var rows *sql.Rows
		rows, err = stmt.Query(groupName)
		if err == nil {
			defer rows.Close()
			for rows.Next() {
				var userName string
				err = rows.Scan(&userName)
				if err == nil {
					user, err := mdb.GetUser(userName)
					if err == nil {
						users = append(users, user)
					} else {
						break
					}
				} else {
					log.Printf(`Error occured while reading list of user names
                        in group '%s'`, groupName)
					break
				}
			}
			err = rows.Err()
		} else {
			log.Printf(`Failed to to get users in group '%s`, groupName)
		}
	} else {
		log.Printf(`Failed to prepare query to get users in group '%s`,
			groupName)
	}
	return users, err
}

func (mdb *MysqlDb) GetGroupsForUser(userName string) ([]*UserGroup, error) {
	const MYSQL_GET_GROUPS_FOR_USER = `SELECT group_name FROM
	                        orek_user_to_group WHERE user_name = ?;`
	stmt, err := mdb.Prepare(MYSQL_GET_GROUPS_FOR_USER)
	groups := make([]*UserGroup, 10)
	if err == nil {
		defer stmt.Close()
		var rows *sql.Rows
		rows, err = mdb.Query(userName)
		if err == nil {
			defer rows.Close()
			for rows.Next() {
				var groupName string
				err = rows.Scan(&groupName)
				if err == nil {
					group, err := mdb.GetUserGroup(groupName)
					if err == nil {
						groups = append(groups, group)
					} else {
						break
					}
				} else {
					log.Printf(`Failed to fetch list of groups of user '%s'`,
						userName)
					break
				}
			}
			err = rows.Err()
		} else {
			log.Printf(`Failed to prepare query to get groups of user %s`,
				userName)
		}
	}
	return groups, err
}

func (mdb *MysqlDb) AddVariableToGroup(variableId, varGroupId string) error {
	const MYSQL_ADD_VARIABLE_TO_GROUP = `INSERT IGNORE INTO
	        orek_variable_to_group(var_group_id, variable_name ) VALUES( ? ? );`
	stmt, err := mdb.Prepare(MYSQL_ADD_VARIABLE_TO_GROUP)
	if err == nil {
		defer stmt.Close()
		_, err = stmt.Exec(varGroupId, variableId)
		if err != nil {
			log.Printf(`Failed to add variable '%s' to group '%s'`,
				variableId, varGroupId)
		}
	} else {
		log.Printf(`Failed to prepare query to add variable '%s' to group '%s'`,
			variableId, varGroupId)
	}
	return err
}

func (mdb *MysqlDb) RemoveVariableFromGroup(
	variableId, varGroupId string) error {
	const MYSQL_REMOVE_VARIABLE_FROM_GROUP = `DELETE FROM orek_variable_to_group
	                        WHERE var_group_id = ? AND variable_id = ?;`
	stmt, err := mdb.Prepare(MYSQL_REMOVE_VARIABLE_FROM_GROUP)
	if err == nil {
		defer stmt.Close()
		_, err = stmt.Exec(varGroupId, variableId)
		if err != nil {
			log.Printf(`Failed to remove variable '%s' from group '%s'`,
				variableId, varGroupId)
		}
	} else {
		log.Printf(`Failed to prepare query to remove variable '%s'
		            from group '%s'`)
	}
	return err
}

func (mdb *MysqlDb) GetVariablesInGroup(
	varGroupId string) ([]*Variable, error) {
	const MYSQL_GET_VARIABLE_IN_GROUP = `SELECT variable_id FROM
	    orek_variable_to_group WHERE var_group_id = ?;`
	stmt, err := mdb.Prepare(MYSQL_GET_VARIABLE_IN_GROUP)
	variables := make([]*Variable, 10)
	if err == nil {
		defer stmt.Close()
		var rows *sql.Rows
		rows, err = mdb.Query(varGroupId)
		if err == nil {
			defer rows.Close()
			for rows.Next() {
				var variableId string
				err = rows.Scan(&variableId)
				if err == nil {
					var variable *Variable
					variable, err = mdb.GetVariableWithId(variableId)
					if err == nil {
						variables = append(variables, variable)
					} else {
						break
					}
				} else {
					log.Printf(`Could not find a variable with id '%s'`,
						varGroupId)
				}
			}
			err = rows.Err()
		} else {
			log.Printf(`Failed to list variables in group '%s'`, varGroupId)
		}
	} else {
		log.Printf(`Failed to prepare query to list variables in group '%s'`,
			varGroupId)
	}
	return variables, err
}

func (mdb *MysqlDb) GetGroupsForVariable(
	variableId string) ([]*VariableGroup, error) {
	const MYSQL_GET_GROUPS_FOR_VARIABLE = `SELECT var_group_id FROM
	    orek_variable_to_group WHERE variable_id = ?;`
	stmt, err := mdb.Prepare(MYSQL_GET_GROUPS_FOR_VARIABLE)
	variableGroups := make([]*VariableGroup, 10)
	if err == nil {
		defer stmt.Close()
		var rows *sql.Rows
		rows, err = mdb.Query(variableId)
		if err == nil {
			defer rows.Close()
			for rows.Next() {
				var groupId string
				err = rows.Scan(&groupId)
				if err == nil {
					var group *VariableGroup
					group, err = mdb.GetVariableGroupWithId(groupId)
					if err == nil {
						variableGroups = append(variableGroups, group)
					} else {
						break
					}
				} else {
					log.Printf(`Failed to fetch group id from database for
					            variable '%s'`, variableId)
					break
				}
			}
		} else {
			log.Printf(`Failed to fetch groups with which variable '%s' is
			            associated`, variableId)
		}
	} else {
		log.Printf(`Failed to prepare query to fetch groups the variable '%s'
		            associated with`, variableId)
	}
	return variableGroups, err
}

func (mdb *MysqlDb) AddVariableValue(variableId, value string) error {
	const MYSQL_ADD_VARIABLE_VALUE = `INSERT INTO orek_variable_value(
	    variable_id, value ) VALUES( ?, ? );`
	stmt, err := mdb.Prepare(MYSQL_ADD_VARIABLE_VALUE)
	if err == nil {
		defer stmt.Close()
		_, err = stmt.Exec(variableId, value)
		if err != nil {
			log.Printf(`Failed to add value '%s' of variable '%s'`,
				value, variableId)
		}
	} else {
		log.Printf(`Failed prepare query to add value '%s' of variable '%s'`,
			value, variableId)
	}
	return err
}

func (mdb *MysqlDb) ClearValuesForVariable(variableId string) error {
	const MYSQL_CLEAR_VALUES_FOR_VARIABLE = `DELETE FROM orek_variable_value
	    WHERE variable_id = ?;`
	stmt, err := mdb.Prepare(MYSQL_CLEAR_VALUES_FOR_VARIABLE)
	if err == nil {
		defer stmt.Close()
		_, err = stmt.Exec(variableId)
		if err != nil {
			log.Printf(`Failed to clear values of variable '%s'`, variableId)
		}
	} else {
		log.Printf(`Failed to prepare query to clear value for variable '%s'`,
			variableId)
	}
	return nil
}

func (mdb *MysqlDb) GetValuesForVariable(variableId string) ([]*string, error) {
	const MYSQL_GET_VALUES_FOR_VARIABLE = `DELETE FROM orek_variable_value WHERE
											variable_id = ?;`
	stmt, err := mdb.Prepare(MYSQL_GET_VALUES_FOR_VARIABLE)
	values := make([]*string, 10)
	if err == nil {
		defer stmt.Close()
		var rows *sql.Rows
		rows, err = mdb.Query(variableId)
		if err == nil {
			defer rows.Close()
			for rows.Next() {
				var value string
				err = rows.Scan(&value)
				if err == nil {
					values = append(values, &value)
				} else {
					log.Printf(`Error occured while retrieving a value for
					    variable '%s'`, variableId)
				}
			}
			err = rows.Err()
		} else {
			log.Printf(`Failed to fetch values of variable '%s'`, variableId)
		}
	} else {
		log.Printf(`Failed to prepare query to fetch values of variable '%s'`,
			variableId)
	}
	return values, err
}
