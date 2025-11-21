CREATE TABLE IF NOT EXISTS libraries (
    id VARCHAR(36) PRIMARY KEY DEFAULT (UUID()),
    title VARCHAR(255) NOT NULL UNIQUE,
    small_description TEXT NOT NULL,
    degree ENUM('apprentice', 'companion', 'master') NOT NULL,
    file_url TEXT,
    cover_url TEXT,
    link VARCHAR(500),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    INDEX idx_libraries_title (title),
    INDEX idx_libraries_degree (degree)
);