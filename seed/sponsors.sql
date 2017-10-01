PREPARE make_sponsor (varchar, varchar, varchar(255)) AS
    INSERT INTO users (name, email, hashed_password) VALUES ($1, $2, $3);

-- Password: 'password'
EXECUTE make_sponsor('Ada Lovelace', 'ada@sponsor.com', '$2a$10$lCnz/7edCULg31XqU4.KOOJuyn4U6VKwkiYOwweRcdqykVpj9rT7W');
