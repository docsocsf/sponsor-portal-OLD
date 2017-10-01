PREPARE make_sponsor (varchar, varchar, varchar(255)) AS
    INSERT INTO users (name, email, hashed_password) VALUES ($1, $2, $3);

-- Password: 'password'
EXECUTE make_sponsor('Ada Lovelace', 'ada@sponsor.com', '$2a$10$TEKE9mkzJQG5nlp3lPlzj.FAg4cxF8Fw0E.Aht2hHjd4kH.vJInCC');
