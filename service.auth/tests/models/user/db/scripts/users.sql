CREATE TABLE IF NOT EXISTS users (
    id            INT PRIMARY KEY AUTO_INCREMENT,
    username      TEXT,
    password      TEXT,
    role          TEXT
);