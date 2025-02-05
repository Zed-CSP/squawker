#!/bin/bash

# Exit on error
set -e

echo "🗑️  Dropping existing database..."
dropdb --if-exists squawker

echo "🛠️  Creating fresh database..."
createdb squawker

echo "📊 Creating tables..."
psql squawker << 'EOSQL'
-- Users table
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(50) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Messages table (squawks)
CREATE TABLE messages (
    id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES users(id) ON DELETE CASCADE,
    content VARCHAR(280) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Daily post tracking
CREATE TABLE daily_posts (
    user_id INTEGER REFERENCES users(id) ON DELETE CASCADE,
    post_date DATE NOT NULL,
    post_count INTEGER DEFAULT 1,
    PRIMARY KEY (user_id, post_date)
);

-- Follows table
CREATE TABLE follows (
    follower_id INTEGER REFERENCES users(id) ON DELETE CASCADE,
    following_id INTEGER REFERENCES users(id) ON DELETE CASCADE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (follower_id, following_id)
);

-- Indexes
CREATE INDEX messages_user_id_idx ON messages(user_id);
CREATE INDEX messages_created_at_idx ON messages(created_at);
CREATE INDEX daily_posts_date_idx ON daily_posts(post_date);
EOSQL

echo "🌱 Seeding database..."
psql squawker < "$(dirname "$0")/seed.sql"

echo "✅ Database reset and seeded successfully!" 