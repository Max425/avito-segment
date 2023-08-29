CREATE TABLE segments (
    slug VARCHAR(255) PRIMARY KEY
);

CREATE TABLE users (
    id SERIAL PRIMARY KEY
);

CREATE TABLE users_segments (
    id SERIAL PRIMARY KEY,
    user_id INT REFERENCES users(id) ON DELETE CASCADE,
    segment_slug VARCHAR(255) REFERENCES segments(slug) ON DELETE CASCADE,
    added_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    expires_at TIMESTAMP,
    CONSTRAINT unique_user_segment UNIQUE (user_id, segment_slug)
);

-- Создаем функцию, которая удаляет просроченные записи
CREATE OR REPLACE FUNCTION delete_expired_user_segments()
RETURNS TRIGGER AS $$
BEGIN
DELETE FROM users_segments WHERE expires_at < NOW();
RETURN NULL;
END;
$$ LANGUAGE plpgsql;

-- Создаем триггер, который будет вызывать функцию при каждой операции вставки в таблицу users_segments
CREATE TRIGGER delete_expired_segments_trigger
    AFTER INSERT ON users_segments
    EXECUTE FUNCTION delete_expired_user_segments();
