CREATE TABLE IF NOT EXISTS users(
    id uuid PRIMARY KEY,
    user_name VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    email_address TEXT NOT NULL,
    password
);