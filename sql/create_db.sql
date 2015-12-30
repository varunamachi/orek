CREATE DATABASE orek;

USE orek;

CREATE TABLE orek_user(
    user_name     VARCHAR( 255 ) NOT NULL,
    first_name    VARCHAR( 255 ),
    last_name     VARCHAR( 255 ),
    emai          VARCHAR( 255 ) NOT NULL,
    UNIQUE( email )
    PRIMARY KEY( user_name ),
);

CREATE TABLE orek_user_identity(
    user_name   VARCHAR( 255 ) NOT NULL,
    email       VARCHAR( 255 ) NOT NULL,
    digest      VARCHAR( 255 ) NOT NULL,
    PRIMARY KEY( user_name ),
    UNIQUE( email ),
    FOREIGN KEY user_name REFERENCES orek_user( user_name ),
    FOREIGN KEY email REFERENCES orek_user( email )
)


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
    name         VARCHAR( 255 ) NOT NULL,
    source_id    CHAR( 36 )     NOT NULL,
    description  TEXT           NOT NULL,
    unit         CHAR( 30 )     NOT NULL,
    PRIMARY KEY( source_id, name ),
    FOREIGN KEY( source_id ) REFERENCES orek_source( source_id )
    
);


CREATE TABLE orek_user_group(
    name        VARCHAR( 256 ) NOT NULL,
    owner       VARCHAR( 256 ) NOT NULL,
    description TEXT NOT NULL,
    PRIMARY KEY( name ),
    FOREIGN KEY owner REFERENCES orek_user( user_name )
);

CREATE TABLE orek_user_to_group(
    group_name  VARCHAR( 256 ) NOT NULL,
    user_name   VARCHAR( 256 ) NOT NULL,
    FOREIGN KEY group_name REFERENCES orek_user_group( name ),
    FOREIGN KEY user_name REFERENCES orek_user( user_name ),
)


