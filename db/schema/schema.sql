CREATE TABLE accessors (
    username TEXT NOT NULL UNIQUE PRIMARY KEY,
    hashed_password TEXT NOT NULL
)