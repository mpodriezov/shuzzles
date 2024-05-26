-- Create a table for item types
CREATE TABLE
    IF NOT EXISTS sessions (
        session_id TEXT PRIMARY KEY NOT NULL,
        user_id INTEGER NOT NULL,
        created_on TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        modified_on TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        expires_on TIMESTAMP
    );