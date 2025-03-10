CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(50) NOT NULL,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    status VARCHAR(20) NOT NULL
);

CREATE TABLE IF NOT EXISTS 'authentication' (
    id SERIAL PRIMARY KEY,
    hash VARCHAR(255) NOT NULL,
    iduser INTEGER REFERENCES users(id),
    status VARCHAR(20) NOT NULL
);
