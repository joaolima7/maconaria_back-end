CREATE TABLE IF NOT EXISTS acacias (
    id VARCHAR(36) PRIMARY KEY DEFAULT (UUID()),
    name VARCHAR(255) NOT NULL UNIQUE,
    terms JSON,
    is_president BOOLEAN NOT NULL DEFAULT false,
    deceased BOOLEAN NOT NULL DEFAULT false,
    image_url TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    INDEX idx_acacias_name (name),
    INDEX idx_acacias_president (is_president),
    INDEX idx_acacias_deceased (deceased)
);