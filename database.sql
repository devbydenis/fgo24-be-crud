CREATE TABLE users (
  id SERIAL PRIMARY KEY,
  name VARCHAR(100),
  email VARCHAR(100),
  password VARCHAR(100)
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
