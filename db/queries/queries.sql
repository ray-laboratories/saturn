-- name: GetAccessor :one
SELECT * FROM accessors WHERE username = ?;

-- name: InsertAccessor :execresult
INSERT INTO accessors (username, hashed_password) VALUES (?, ?);

-- name: UpdateAccessor :execresult
UPDATE accessors SET username = ?, hashed_password = ? WHERE username = ?;