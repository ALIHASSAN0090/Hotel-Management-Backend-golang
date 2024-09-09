-- Insert data into menus
INSERT INTO menus (id, name, category) VALUES
(1, 'Breakfast Menu', 'Breakfast'),
(2, 'Lunch Menu', 'Lunch'),
(3, 'Dinner Menu', 'Dinner');

-- Insert data into restaurant_tables
INSERT INTO restaurant_tables (id, menu_id) VALUES
(1, 1),
(2, 2),
(3, 3),
(4, 2);

-- Insert data into food_items
INSERT INTO food_items (id, name, price, menu_id, table_id) VALUES
(1, 'Pancakes', 5.99, 1, 1),
(2, 'Waffles', 6.99, 1, 1),
(3, 'Burger', 8.99, 2, 2),
(4, 'Caesar Salad', 7.49, 2, 2),
(5, 'Steak', 15.99, 3, 3),
(6, 'Spaghetti', 12.99, 3, 3),
(7, 'Fish and Chips', 10.99, 2, 4),
(8, 'Chicken Wings', 9.49, 2, 4);

-- Insert data into orders
INSERT INTO orders (id, total_price) VALUES
(1, 21.97),
(2, 15.49),
(3, 28.98);

-- Insert data into invoices
INSERT INTO invoices (id, order_id, payment_method, payment_status) VALUES
(1, 1, 'Credit Card', 'Paid'),
(2, 2, 'Cash', 'Pending'),
(3, 3, 'Credit Card', 'Paid');

-- Insert data into order_food_items
INSERT INTO order_food_items (order_id, food_item_id) VALUES
(1, 1), -- Pancakes
(1, 2), -- Waffles
(2, 3), -- Burger
(2, 4), -- Caesar Salad
(3, 5), -- Steak
(3, 6), -- Spaghetti
(3, 7), -- Fish and Chips
(3, 8); -- Chicken Wings
