-- name: CreateUser :exec
INSERT INTO users (name, email, password) VALUES ($1, $2, $3);

-- name: GetUsers :many
SELECT id, name, email FROM users;

-- name: CreateProduct :exec
INSERT INTO products (name, price) VALUES ($1, $2);

-- name: GetProducts :many
SELECT id, name, price FROM products;
