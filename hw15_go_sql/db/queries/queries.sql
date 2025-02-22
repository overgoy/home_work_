
-- name: CreateUser :exec
INSERT INTO users (name, email, password) VALUES ($1, $2, $3);

-- name: UpdateUser :exec
UPDATE users SET name = $2, email = $3, password = $4 WHERE id = $1;

-- name: DeleteUser :exec
DELETE FROM users WHERE id = $1;

-- name: GetUsers :many
SELECT id, name, email FROM users;

-- name: GetUser :one
SELECT id, name, email FROM users WHERE id = $1;

-- name: CreateProduct :exec
INSERT INTO products (name, price) VALUES ($1, $2);

-- name: UpdateProduct :exec
UPDATE products SET name = $2, price = $3 WHERE id = $1;

-- name: DeleteProduct :exec
DELETE FROM products WHERE id = $1;

-- name: GetProducts :many
SELECT id, name, price FROM products;

-- name: GetProduct :one
SELECT id, name, price FROM products WHERE id = $1;

-- name: CreateOrder :exec
INSERT INTO orders (user_id, product_id, quantity, total_price) VALUES ($1, $2, $3, $4);

-- name: DeleteOrder :exec
DELETE FROM orders WHERE id = $1;

-- name: GetOrders :many
SELECT id, user_id, product_id, quantity, total_price, created_at FROM orders;

-- name: GetOrdersByUser :many
SELECT o.id, o.user_id, o.product_id, p.name AS product_name, o.quantity, o.total_price, o.created_at
FROM orders o
         JOIN products p ON o.product_id = p.id
WHERE o.user_id = $1;

-- name: GetUserStatistics :one
SELECT
    u.id AS user_id,
    u.name AS user_name,
    COALESCE(SUM(o.total_price), 0) AS total_spent,
    COALESCE(AVG(p.price), 0) AS avg_product_price
FROM users u
         LEFT JOIN orders o ON u.id = o.user_id
         LEFT JOIN products p ON o.product_id = p.id
WHERE u.id = $1
GROUP BY u.id, u.name;


