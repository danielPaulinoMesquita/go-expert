CREATE TABLE categories (
                            id VARCHAR(36) NOT NULL PRIMARY KEY,
                            name text not null ,
                            description TEXT
);

CREATE TABLE courses (
                         id VARCHAR(36) NOT NULL PRIMARY KEY,
                         category_id VARCHAR(36) NOT NULL,
                         name text,
                         description TEXT,
                         price DECIMAL(10, 2) not null ,
                         FOREIGN KEY (category_id) REFERENCES categories(id)
);