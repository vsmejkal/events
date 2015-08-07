-- Turn on spatial extensions
CREATE EXTENSION cube;
CREATE EXTENSION earthdistance;

CREATE TABLE event (
    id serial PRIMARY KEY,
    name varchar(100) NOT NULL,
    description text,
    link varchar(256) NOT NULL,
    image varchar(256),
    starttime timestamp NOT NULL,
    endtime timestamp NOT NULL,
    dateonly boolean,
    tags varchar(20)[]
);

CREATE TABLE place (
    id serial PRIMARY KEY,
    name varchar(100) NOT NULL,
    gps point NOT NULL,
    street varchar(100),
    city varchar(50),
    zip varchar(10),
    tags varchar(20)[]
);