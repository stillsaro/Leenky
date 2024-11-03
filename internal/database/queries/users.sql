-- name: GetUserByEmail :one
SELECT * 
FROM users 
WHERE email = $1;

-- name: GetUserByUsername :one
SELECT * 
FROM users 
WHERE username = $1;

-- name: InsertUser :exec
INSERT INTO users (
    username, email, password_hash
) VALUES ($1, $2, $3);
