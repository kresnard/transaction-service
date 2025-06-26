CREATE TABLE IF NOT EXISTS promotions (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    type ENUM('freebie', 'bundle', 'discount') NOT NULL,
    target_sku VARCHAR(50) NOT NULL,
    condition_quantity INT NOT NULL,
    discount_percent DECIMAL(5,2),
    free_sku VARCHAR(50),
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP
);


CREATE TABLE IF NOT EXISTS products (
    sku VARCHAR(50) PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    price DECIMAL(10,2) NOT NULL,
    inventory_qty INT NOT NULL
);

CREATE TABLE IF NOT EXISTS orders (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    total_price DECIMAL(10,2) NOT NULL
);

CREATE TABLE IF NOT EXISTS order_items (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    order_id BIGINT NOT NULL,
    sku VARCHAR(50) NOT NULL,
    quantity INT NOT NULL,
    unit_price DECIMAL(10,2) NOT NULL,
    subtotal DECIMAL(10,2) NOT NULL,
    FOREIGN KEY (order_id) REFERENCES orders(id),
    FOREIGN KEY (sku) REFERENCES products(sku)
);

INSERT INTO products (sku, name, price, inventory_qty) VALUES
('43N23P', 'MacBook Pro', 5399.99, 5),
('234234', 'Raspberry Pi B', 30.00, 10),
('120P90', 'Google Home', 49.99, 10),
('A304SD', 'Alexa Speaker', 109.50, 10);


INSERT INTO promotions (name, type, target_sku, condition_quantity, free_sku)
VALUES ('Free Raspberry Pi with MacBook Pro', 'freebie', '43N23P', 1, '234234');

INSERT INTO promotions (name, type, target_sku, condition_quantity)
VALUES ('Buy 3 Google Homes pay 2', 'bundle', '120P90', 3);

INSERT INTO promotions (name, type, target_sku, condition_quantity, discount_percent)
VALUES ('10% off Alexa if buy > 3', 'discount', 'A304SD', 4, 10.00);

