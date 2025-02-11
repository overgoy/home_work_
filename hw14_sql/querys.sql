CREATE DATABASE shop_db;


CREATE USER shop_user WITH ENCRYPTED PASSWORD '12345';
GRANT ALL PRIVILEGES ON DATABASE shop_db TO shop_user;


CREATE TABLE Users (
                       id SERIAL PRIMARY KEY,
                       name VARCHAR(100) NOT NULL,
                       email VARCHAR(100) UNIQUE NOT NULL,
                       password TEXT NOT NULL
);

CREATE TABLE Orders (
                        id SERIAL PRIMARY KEY,
                        user_id INT NOT NULL,
                        order_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                        total_amount DECIMAL(10,2) NOT NULL,
                        FOREIGN KEY (user_id) REFERENCES Users(id) ON DELETE CASCADE
);

CREATE TABLE Products (
                          id SERIAL PRIMARY KEY,
                          name VARCHAR(255) NOT NULL,
                          price DECIMAL(10,2) NOT NULL
);

CREATE TABLE OrderProducts (
                               order_id INT NOT NULL,
                               product_id INT NOT NULL,
                               quantity INT NOT NULL CHECK (quantity > 0),
                               PRIMARY KEY (order_id, product_id),
                               FOREIGN KEY (order_id) REFERENCES Orders(id) ON DELETE CASCADE,
                               FOREIGN KEY (product_id) REFERENCES Products(id) ON DELETE CASCADE
);

INSERT INTO Users (name, email, password) VALUES ('Иван Иванов', 'ivan@example.com', 'pass');
INSERT INTO Products (name, price) VALUES ('Ноутбук', 50000.00);

UPDATE Users SET name = 'Иван Петров' WHERE id = 1;
UPDATE Products SET price = 45000.00 WHERE id = 1;

DELETE FROM Users WHERE id = 2;
DELETE FROM Products WHERE id = 1;

INSERT INTO Orders (user_id, total_amount) VALUES (1, 50000.00);
INSERT INTO OrderProducts (order_id, product_id, quantity) VALUES (1, 1, 1);

DELETE FROM Orders WHERE id = 1;

SELECT * FROM Users;
SELECT * FROM Products;
SELECT * FROM Orders WHERE user_id = 1;


SELECT user_id, SUM(total_amount) AS total_spent,
       AVG(p.price) AS avg_product_price
FROM Orders o
         JOIN OrderProducts op ON o.id = op.order_id
         JOIN Products p ON op.product_id = p.id
GROUP BY user_id;

CREATE INDEX idx_users_email ON Users(email);
CREATE INDEX idx_orders_user_id ON Orders(user_id);
CREATE INDEX idx_orderproducts_order_id ON OrderProducts(order_id);
CREATE INDEX idx_orderproducts_product_id ON OrderProducts(product_id);
