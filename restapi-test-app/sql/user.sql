CREATE TABLE IF NOT EXISTS users (
    id serial PRIMARY KEY,
    email VARCHAR(128) NOT NULL UNIQUE,
    password VARCHAR(256) NOT NULL,
    fullname VARCHAR(256) NULL
);