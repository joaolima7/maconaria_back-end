CREATE TABLE IF NOT EXISTS wordkeys (
    id VARCHAR(36) PRIMARY KEY DEFAULT (UUID()),
    wordkey VARCHAR(255) NOT NULL UNIQUE,
    active BOOLEAN NOT NULL DEFAULT false,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    INDEX idx_wordkeys_active (active),
    INDEX idx_wordkeys_wordkey (wordkey)
);
