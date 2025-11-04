CREATE TABLE IF NOT EXISTS posts (
    id VARCHAR(36) PRIMARY KEY DEFAULT (UUID()),
    title TEXT NOT NULL,
    category VARCHAR(100) NOT NULL,
    small_description TEXT NOT NULL,
    complete_description TEXT NOT NULL,
    date VARCHAR(50),
    time VARCHAR(50),
    location TEXT,
    is_featured BOOLEAN NOT NULL DEFAULT false,
    post_type ENUM('event', 'commemoration', 'article', 'news') NOT NULL DEFAULT 'event',
    user_id VARCHAR(36) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    INDEX idx_posts_user_id (user_id),
    INDEX idx_posts_type (post_type),
    INDEX idx_posts_featured (is_featured)
);