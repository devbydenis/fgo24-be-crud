CREATE TABLE users (
  id SERIAL PRIMARY KEY,
  name VARCHAR(100),
  email VARCHAR(100),
  password VARCHAR(100),
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP
);

INSERT INTO users (name, email, password) VALUES 
('John', '2M5lM@example.com', 'password123'),
('Jane', 'QH6kW@example.com', 'password456'),
('Bob Smith', 'QH6kW@example.com', 'password789'),
('Alice Smith', 'QH6kW@example.com', 'password123'),
('Bahlil', 'QH6kW@example.com', 'password456'),
('Joekowie', 'QH6kW@example.com', 'password789'),
('Gibrun', 'QH6kW@example.com', 'password123'),
('prabroro', 'QH6kW@example.com', 'password456'),
('Johnson', 'QH6kW@example.com', 'password789'),
('Jones', 'QH6kW@example.com', 'password123');


INSERT INTO users (name, email, password) VALUES
('Rojali', 'rojali@example.com', 'password456');
SELECT * FROM users