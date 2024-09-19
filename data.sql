// ... existing code ...

INSERT INTO users (username, email, password_hash, first_name, last_name, role) VALUES
('user1', 'user1@example.com', 'hash1', 'John', 'Doe', 'user'),
('user2', 'user2@example.com', 'hash2', 'Jane', 'Doe', 'user'),
('user3', 'user3@example.com', 'hash3', 'Jim', 'Beam', 'user'),
('user4', 'user4@example.com', 'hash4', 'Jack', 'Daniels', 'user'),
('user5', 'user5@example.com', 'hash5', 'Johnny', 'Walker', 'user'),
('user6', 'user6@example.com', 'hash6', 'Jameson', 'Irish', 'user'),
('user7', 'user7@example.com', 'hash7', 'Jose', 'Cuervo', 'user'),
('user8', 'user8@example.com', 'hash8', 'Jager', 'Meister', 'user'),
('user9', 'user9@example.com', 'hash9', 'Captain', 'Morgan', 'user'),
('user10', 'user10@example.com', 'hash10', 'Don', 'Julio', 'user');

INSERT INTO menus (name, category) VALUES
('Breakfast', 'Morning'),
('Lunch', 'Afternoon'),
('Dinner', 'Evening'),
('Desserts', 'Anytime'),
('Beverages', 'Anytime'),
('Snacks', 'Anytime'),
('Specials', 'Anytime'),
('Kids Menu', 'Anytime'),
('Vegan', 'Anytime'),
('Gluten-Free', 'Anytime');

INSERT INTO food_items (name, price, menu_id) VALUES
('Pancakes', 5.99, 1),
('Burger', 8.99, 2),
('Steak', 15.99, 3),
('Ice Cream', 3.99, 4),
('Coffee', 2.99, 5),
('Chips', 1.99, 6),
('Daily Special', 12.99, 7),
('Kids Burger', 4.99, 8),
('Vegan Salad', 7.99, 9),
('Gluten-Free Bread', 3.99, 10);

INSERT INTO orders (total_price, user_id) VALUES
(25.99, 1),
(15.99, 2),
(30.99, 3),
(10.99, 4),
(50.99, 5),
(20.99, 6),
(35.99, 7),
(40.99, 8),
(45.99, 9),
(55.99, 10);

INSERT INTO invoices (order_id, payment_method, payment_status) VALUES
(1, 'Credit Card', 'Paid'),
(2, 'Credit Card', 'Paid'),
(3, 'Cash', 'Paid'),
(4, 'Credit Card', 'Paid'),
(5, 'Cash', 'Paid'),
(6, 'Credit Card', 'Paid'),
(7, 'Credit Card', 'Paid'),
(8, 'Cash', 'Paid'),
(9, 'Credit Card', 'Paid'),
(10, 'Cash', 'Paid');

INSERT INTO order_food_items (order_id, food_item_id) VALUES
(1, 1),
(1, 2),
(2, 3),
(2, 4),
(3, 5),
(3, 6),
(4, 7),
(4, 8),
(5, 9),
(5, 10),
(6, 1),
(6, 2),
(7, 3),
(7, 4),
(8, 5),
(8, 6),
(9, 7),
(9, 8),
(10, 9),
(10, 10);