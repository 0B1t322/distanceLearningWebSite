CREATE DATABASE IF NOT EXISTS auth;
USE auth;
CREATE TABLE IF NOT EXISTS tokens (
    user_id       INT,
    token         TEXT,
    duration      TIME
);
CREATE TABLE IF NOT EXISTS users (
    id            INT PRIMARY KEY AUTO_INCREMENT,
    username      TEXT,
    password      TEXT,
    role          TEXT
);