CREATE TABLE IF NOT EXISTS chats (
    id SERIAL PRIMARY KEY,
    usernames TEXT[] NOT NULL,
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);

CREATE TABLE chat_messages (
    id SERIAL PRIMARY KEY,
    chat_id INT UNIQUE NOT NULL,
    sender VARCHAR(255) NOT NULL,
    message_text TEXT NOT NULL,
    sent_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (chat_id) REFERENCES chats(id) ON DELETE CASCADE
);