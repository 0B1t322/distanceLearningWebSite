CREATE DATABASE IF NOT EXISTS courses;
USE courses;

CREATE TABLE IF NOT EXISTS tasks (
    id              INT PRIMARY KEY AUTO_INCREMENT,
    task_header_id  INT,
    name            TEXT,
    imgURL          TEXT,
    contentURL      TEXT
);
CREATE TABLE IF NOT EXISTS task_headers (
    id              INT PRIMARY KEY AUTO_INCREMENT,
    course_id       INT,
    name            TEXT
);
CREATE TABLE IF NOT EXISTS courses (
    id              INT PRIMARY KEY AUTO_INCREMENT,
    name            TEXT,
    imgURL          TEXT
);