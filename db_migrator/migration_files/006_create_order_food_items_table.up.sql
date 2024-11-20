CREATE TABLE IF NOT EXISTS order_food_items (
    order_id BIGINT NOT NULL,
    food_item_id BIGINT NOT NULL,
    PRIMARY KEY (order_id, food_item_id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (order_id) REFERENCES orders(id) ON DELETE CASCADE,
    FOREIGN KEY (food_item_id) REFERENCES food_items(id) ON DELETE CASCADE
);