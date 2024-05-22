-- Create a table for item types
CREATE TABLE
    item_types (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        name TEXT,
        created_at TIMESTAMP,
        updated_at TIMESTAMP
    );

-- Create a table for items
CREATE TABLE
    items (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        user_id INTEGER,
        title TEXT,
        type_id INTEGER,
        description TEXT,
        created_at TIMESTAMP,
        updated_at TIMESTAMP
    );

-- Create a table for media
CREATE TABLE
    media (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        item_id INTEGER,
        mime TEXT,
        url TEXT,
        size INTEGER,
        created_at TIMESTAMP,
        updated_at TIMESTAMP,
        FOREIGN KEY (item_id) REFERENCES items (id)
    );

-- Create a table for listings
CREATE TABLE
    listings (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        user_id INTEGER,
        title TEXT,
        description TEXT,
        altitude REAL,
        longitude REAL,
        created_at TIMESTAMP,
        updated_at TIMESTAMP
    );

-- Create a table for items in listings
CREATE TABLE
    items_in_listings (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        listing_id INTEGER,
        item_id INTEGER,
        FOREIGN KEY (listing_id) REFERENCES listings (id),
        FOREIGN KEY (item_id) REFERENCES items (id)
    );

-- Create a table for users
CREATE TABLE
    users (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        username TEXT,
        email TEXT,
        password_hash TEXT,
        bio TEXT,
        altitude REAL,
        longitude REAL,
        created_at TIMESTAMP,
        updated_at TIMESTAMP
    );
