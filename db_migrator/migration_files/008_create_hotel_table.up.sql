CREATE TABLE IF NOT EXISTS hotel_info (
    id SERIAL PRIMARY KEY,
    opening_time VARCHAR(50),
    closing_time VARCHAR(50),
    holiday VARCHAR(50),
    home_delivery VARCHAR(50),
    reservations VARCHAR(50)
);