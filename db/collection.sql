CREATE TABLE collections (
    id INT AUTO_INCREMENT NOT NULL PRIMARY KEY,
    name VARCHAR(50),
    division VARCHAR(50),
    memo VARCHAR(250),
    search_index VARCHAR(50),
    add_search_index VARCHAR(50),
    local_number VARCHAR(6),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);