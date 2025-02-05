-- Drop tables if they exist (in correct order due to dependencies)
DROP TABLE IF EXISTS user_interactions;
DROP TABLE IF EXISTS daily_posts;
DROP TABLE IF EXISTS messages;
DROP TABLE IF EXISTS users;

-- Create users table
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(50) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Create messages table
CREATE TABLE messages (
    id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES users(id) ON DELETE CASCADE,
    content VARCHAR(280) NOT NULL,
    volume INTEGER DEFAULT 0,
    echo_count INTEGER DEFAULT 0,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Create daily posts tracking
CREATE TABLE daily_posts (
    user_id INTEGER REFERENCES users(id) ON DELETE CASCADE,
    post_date DATE NOT NULL,
    post_count INTEGER DEFAULT 1,
    PRIMARY KEY (user_id, post_date)
);

-- Create user interactions table
CREATE TABLE user_interactions (
    id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES users(id) ON DELETE CASCADE,
    message_id INTEGER REFERENCES messages(id) ON DELETE CASCADE,
    interaction_type VARCHAR(10) NOT NULL CHECK (interaction_type IN ('volume', 'echo')),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    UNIQUE(user_id, message_id, interaction_type)
);

-- Create indexes for better performance
CREATE INDEX idx_messages_user_id ON messages(user_id);
CREATE INDEX idx_messages_created_at ON messages(created_at);
CREATE INDEX idx_user_interactions_message_id ON user_interactions(message_id);
CREATE INDEX idx_user_interactions_user_id ON user_interactions(user_id); 