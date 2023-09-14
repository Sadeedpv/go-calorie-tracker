-- create_calories_table.sql
CREATE TABLE IF NOT EXISTS calories (
    id SERIAL PRIMARY KEY,
    food VARCHAR(255) NOT NULL,
    calorie INT NOT NULL
);