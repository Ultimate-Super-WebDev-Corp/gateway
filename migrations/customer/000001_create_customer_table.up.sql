create table customer (
    id serial PRIMARY KEY,
    email VARCHAR(100) NOT NULL,
    password VARCHAR(50) NOT NULL,
    name VARCHAR(100) NOT NULL,
    UNIQUE(email)
);