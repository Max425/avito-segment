CREATE TABLE segments (
    id SERIAL PRIMARY KEY,
    slug VARCHAR(255) UNIQUE NOT NULL
);

CREATE TABLE users (
    id SERIAL PRIMARY KEY
);

CREATE TABLE users_segments (
    id SERIAL PRIMARY KEY,
    user_id INT REFERENCES users(id) ON DELETE CASCADE,
    segment_id INT REFERENCES segments(id) ON DELETE CASCADE
);

