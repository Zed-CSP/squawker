-- Clear existing data
TRUNCATE users, messages, daily_posts, follows CASCADE;

-- Reset sequences
ALTER SEQUENCE users_id_seq RESTART WITH 1;
ALTER SEQUENCE messages_id_seq RESTART WITH 1;

-- Insert test users (passwords are 'password123' - we'll implement proper hashing later)
INSERT INTO users (username, password_hash, created_at) VALUES
    ('alice', 'temporary_hash', NOW() - INTERVAL '7 days'),
    ('bob', 'temporary_hash', NOW() - INTERVAL '6 days'),
    ('charlie', 'temporary_hash', NOW() - INTERVAL '5 days'),
    ('diana', 'temporary_hash', NOW() - INTERVAL '4 days'),
    ('evan', 'temporary_hash', NOW() - INTERVAL '3 days');

-- Insert test messages
INSERT INTO messages (user_id, content, created_at) VALUES
    (1, 'Hello world! This is my first squawk!', NOW() - INTERVAL '7 days'),
    (2, 'Just joined Squawker! Excited to be here!', NOW() - INTERVAL '6 days'),
    (1, 'Beautiful day for coding! #programming', NOW() - INTERVAL '5 days'),
    (3, 'Anyone want to collaborate on a project?', NOW() - INTERVAL '4 days'),
    (4, 'Just deployed my first app! #milestone', NOW() - INTERVAL '3 days'),
    (5, 'Learning Go and React - amazing combo!', NOW() - INTERVAL '2 days'),
    (2, 'Who else is loving Squawker?', NOW() - INTERVAL '1 day'),
    (3, 'Working on some exciting features!', NOW() - INTERVAL '12 hours'),
    (1, 'Remember: one squawk per day keeps the doctor away 😄', NOW() - INTERVAL '6 hours'),
    (4, 'Just found a great new coffee shop! ☕', NOW() - INTERVAL '1 hour');

-- Insert test follows
INSERT INTO follows (follower_id, following_id, created_at) VALUES
    (1, 2, NOW() - INTERVAL '6 days'),  -- Alice follows Bob
    (1, 3, NOW() - INTERVAL '5 days'),  -- Alice follows Charlie
    (2, 1, NOW() - INTERVAL '5 days'),  -- Bob follows Alice
    (3, 1, NOW() - INTERVAL '4 days'),  -- Charlie follows Alice
    (4, 1, NOW() - INTERVAL '3 days'),  -- Diana follows Alice
    (5, 1, NOW() - INTERVAL '2 days'),  -- Evan follows Alice
    (2, 3, NOW() - INTERVAL '2 days'),  -- Bob follows Charlie
    (3, 2, NOW() - INTERVAL '1 day');   -- Charlie follows Bob

-- Insert daily post tracking for today
INSERT INTO daily_posts (user_id, post_date, post_count) VALUES
    (1, CURRENT_DATE, 1),
    (2, CURRENT_DATE, 1),
    (4, CURRENT_DATE, 1);

-- Count of users
SELECT COUNT(*) FROM users;

-- Count of messages per user
SELECT u.username, COUNT(m.id) as message_count 
FROM users u 
LEFT JOIN messages m ON u.id = m.user_id 
GROUP BY u.username;

-- Count of followers per user
SELECT u.username, COUNT(f.follower_id) as follower_count 
FROM users u 
LEFT JOIN follows f ON u.id = f.following_id 
GROUP BY u.username; 