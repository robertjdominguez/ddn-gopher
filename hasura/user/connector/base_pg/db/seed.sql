CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL
);

INSERT INTO users (username, password) VALUES ('exampleuser', 'examplepassword');
INSERT INTO users (username, password) VALUES ('exampleuser1', 'examplepassword1');
