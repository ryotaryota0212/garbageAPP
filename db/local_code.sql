CREATE TABLE local_code (
    id INT AUTO_INCREMENT NOT NULL PRIMARY KEY,
    global_code VARCHAR(10),
    local_code VARCHAR(10),
    name VARCHAR(50),
    furigana VARCHAR(255)
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);