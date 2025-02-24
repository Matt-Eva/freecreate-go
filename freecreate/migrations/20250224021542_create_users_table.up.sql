CREATE TABLE users (
    id bigserial PRIMARY KEY,
    uuid UUID NOT NULL,
    email VARCHAR(500) UNIQUE NOT NULL,
    username VARCHAR(50),
    password VARCHAR(50)
);