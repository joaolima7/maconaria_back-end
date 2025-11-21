CREATE TABLE IF NOT EXISTS timelines (
    id VARCHAR(36) PRIMARY KEY DEFAULT (UUID()),
    period VARCHAR(50) NOT NULL UNIQUE,
    pdf_url TEXT NOT NULL,
    is_highlight BOOLEAN NOT NULL DEFAULT false,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    INDEX idx_timelines_period (period),
    INDEX idx_timelines_highlight (is_highlight)
);