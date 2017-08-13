CREATE TABLE user_cv (
    user_id SERIAL UNIQUE,
    name TEXT,
    file TEXT,
    FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE ON UPDATE CASCADE
)
