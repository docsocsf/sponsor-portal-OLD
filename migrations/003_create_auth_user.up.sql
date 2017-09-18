ALTER TABLE users
  ADD COLUMN hashed_password VARCHAR(255),
  ADD COLUMN role int;
