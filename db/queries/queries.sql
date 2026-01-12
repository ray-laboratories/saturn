-- name: GetUser :one
SELECT * FROM users WHERE username = ?;

-- name: InsertUser :execresult
INSERT INTO users (username, hashed_password) VALUES (?, ?);

-- name: UpdateUser :execresult
UPDATE users SET username = ?, hashed_password = ? WHERE username = ?;