CREATE TABLE users (
    id SERIAL PRIMARY KEY,

    name TEXT,
    email TEXT UNIQUE
)
