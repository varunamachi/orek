CREATE DATABASE orek;

USE orek;

CREATE TABLE orek_user(
    user_name     VARCHAR( 255 ) NOT NULL,
    first_name    VARCHAR( 255 ),
    last_name     VARCHAR( 255 ),
    email          VARCHAR( 255 ) NOT NULL,
    PRIMARY KEY( user_name ),
    UNIQUE( email )
);

CREATE TABLE orek_user_identity(
    user_name   VARCHAR( 255 ) NOT NULL,
    email       VARCHAR( 255 ) NOT NULL,
    digest      VARCHAR( 255 ) NOT NULL,
    PRIMARY KEY( user_name ),
    UNIQUE( email ),
    FOREIGN KEY( user_name ) REFERENCES orek_user( user_name ),
    FOREIGN KEY( email ) REFERENCES orek_user( email )
);


--'access' is an enum -> private, group, public
CREATE TABLE orek_source(
    source_id     CHAR( 36 )     NOT NULL,
    name          VARCHAR( 255 ) NOT NULL,
    owner         VARCHAR( 255 ) NOT NULL,         
    description   TEXT,
    location      VARCHAR( 255 ) NOT NULL,
    access        CHAR( 20 )     NOT NULL,
    PRIMARY KEY( source_id ),
    FOREIGN KEY( owner ) REFERENCES orek_user( user_name )
);


CREATE TABLE orek_variable(
    variable_id  CHAR( 36 )     NOT NULL,
    name         VARCHAR( 255 ) NOT NULL,
    source_id    CHAR( 36 )     NOT NULL,
    description  TEXT           NOT NULL,
    unit         CHAR( 30 )     NOT NULL,
    PRIMARY KEY( variable_id ),
    UNIQUE( source_id, name ),
    FOREIGN KEY( source_id ) REFERENCES orek_source( source_id )
    
);


CREATE TABLE orek_user_group(
    name        VARCHAR( 256 ) NOT NULL,
    owner       VARCHAR( 256 ) NOT NULL,
    description TEXT NOT NULL,
    PRIMARY KEY( name ),
    FOREIGN KEY( owner ) REFERENCES orek_user( user_name )
);

CREATE TABLE orek_user_to_group(
    group_name  VARCHAR( 256 ) NOT NULL,
    user_name   VARCHAR( 256 ) NOT NULL,
    FOREIGN KEY( group_name ) REFERENCES orek_user_group( name ),
    FOREIGN KEY( user_name ) REFERENCES orek_user( user_name ),
    PRIMARY KEY( group_name, user_name )
);


CREATE TABLE orek_variable_group(
    group_id    CHAR( 36 )     NOT NULL,
    name        VARCHAR( 256 ) NOT NULL,
    owner       VARCHAR( 256 ) NOT NULL,
    description TEXT           NOT NULL,
    access      CHAR( 20 )     NOT NULL DEFAULT "public",
    PRIMARY KEY( group_id ),
    UNIQUE( name, owner ),
    FOREIGN KEY( owner ) REFERENCES orek_user( user_name )
);


CREATE TABLE orek_variable_to_group(
    var_group_id        CHAR( 36 ) NOT NULL,    
    variable_id         CHAR( 36 ) NOT NULL,
    PRIMARY KEY( var_group_id, variable_id ),
    FOREIGN KEY( var_group_id ) REFERENCES orek_variable_group( group_id ),
    FOREIGN KEY( variable_id ) REFERENCES orek_variable( variable_id )
);


CREATE TABLE orek_variable_value(
    variable_id         CHAR( 36 ) NOT NULL,
    value               VARCHAR( 256 ) NOT NULL,
    time                TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_orek_var_value ON orek_variable_value( variable_id );
 
