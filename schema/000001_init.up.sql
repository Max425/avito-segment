CREATE TABLE segments (
    slug VARCHAR(255) PRIMARY KEY
);

CREATE TABLE users (
    id SERIAL PRIMARY KEY
);

CREATE TABLE users_segments (
    id SERIAL PRIMARY KEY,
    user_id INT REFERENCES users(id) ON DELETE CASCADE,
    segment_slug VARCHAR(255) REFERENCES segments(slug) ON DELETE CASCADE
);
