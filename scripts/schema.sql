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
    dateonly boolean NOT NULL,
    tags integer[] element references tag(id)
    place integer NOT NULL references place(id)
);

CREATE TABLE place (
    id serial PRIMARY KEY,
    name varchar(100) NOT NULL,
    gps point NOT NULL,
    street varchar(100),
    city varchar(50),
    zip varchar(10),
    tags integer[] element references tag(id)
);

CREATE TABLE tag (
    id serial PRIMARY KEY,
    name varchar(25) NOT NULL,
    label varchar(50)
);

CREATE TABLE source (
    id serial PRIMARY KEY,
    name varchar(100),
    link varchar(256) NOT NULL,
    visited timestamp,
    place integer NOT NULL references place(id)

);