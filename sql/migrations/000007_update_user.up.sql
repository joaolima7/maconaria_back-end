ALTER TABLE users
ADD COLUMN cim VARCHAR(50) NOT NULL,
ADD COLUMN degree ENUM('apprentice', 'companion', 'master') NOT NULL;