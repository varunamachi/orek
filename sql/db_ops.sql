SELECT * FROM orek_user;

SELECT * FROM orek_user WHERE user_name = ?;

SELECT * FROM orek_user WHERE email = "?";

INSERT INTO orek_user( user_name,
 					   first_name,
 					   last_name,
 					   email )
 	VALUES( ?, ?, ?, ? ) ON DUPLICATE KEY UPDATE
        first_name = VALUES( first_name ),
        last_name = VALUES( last_name );
        email = VALUES( email );


DELETE FROM orek_user WHERE user_name = ?;


SELECT * FROM orek_source;

SELECT * FROM orek_source WHERE name = ? AND owner = ?;

SELECT * FROM orek_source WHERE source_id = ?;

INSERT INTO orek_source( source_id,
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
        access = VALUES( access );

DELETE FROM orek_source WHERE source_id = ?;


SELECT * FROM orek_variable;

SELECT * FROM orek_variable WHERE name = ? AND source_id = ?;

SELECT * FROM orek_variable WHERE variable_id = ?;

INSERT INTO orek_variable( variable_id,
                           name,
                           source_id,
                           description,
                           unit )
    VALUES( ?, ?, ?, ?, ? )
    ON DUPLICATE KEY UPDATE
        name = VALUES( name ),
        source_id = VALUES( source_id ),
        description = VALUES( description ),
        unit = VALUES( unit );

DELETE FROM orek_variable WHERE variable_id = ?;


SELECT * FROM orek_user_group;

SELECT * FROM orek_user_group WHERE name = ?;

INSERT INTO orek_user_group( name,
                             owner,
                             description )
    VALUES( ?, ?, ? )
    ON DUPLICATE KEY UPDATE
        owner = VALUES( owner ),
        description = VALUES( description );

DELETE FROM orek_user_group WHERE name = ?;


SELECT * FROM orek_variable_group;

SELECT * FROM orek_variable_group WHERE name = ? AND owner = ?;

SELECT * FROM orek_variable_group WHERE group_id = ?;

INSERT INTO orek_variable_group( group_id,
                                 name,
                                 owner,
                                 description,
                                 access )
    VALUES( ?, ?, ?, ?, ? )
    ON DUPLICATE KEY UPDATE
        name = VALUES( name ),
        owner = VALUES( owner ),
        description = VALUES( description ),
        access = VALUES( access );

DELETE FROM orek_variable_group WHERE group_id = ?;


INSERT IGNORE INTO orek_user_to_group( group_name, user_name )
    VALUES( ?, ? );

DELETE FROM orek_user_to_group WHERE user_name = ? AND group_name = ?;

SELECT user_name FROM orek_user_to_group WHERE group_name = ?;

SELECT group_name FROM orek_user_to_group WHERE user_name = ?;


INSERT IGNORE INTO orek_variable_to_group( var_group_id, variable_name ) VALUES( ? ? );

DELETE FROM orek_variable_to_group WHERE var_group_id = ? AND variable_id = ?;

SELECT variable_id FROM orek_variable_to_group WHERE var_group_id = ?;

SELECT var_group_id FROM orek_variable_to_group WHERE variable_id = ?;


INSERT INTO orek_variable_value( variable_id, value ) VALUES( ?, ? );

DELETE FROM orek_variable_value WHERE variable_id = ?;

SELECT * FROM orek_variable_value WHERE variable_id = ?;

