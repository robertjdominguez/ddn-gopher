CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    name VARCHAR(255) NOT NULL
);

INSERT INTO users (username, password, name) VALUES ('exampleuser', 'examplepassword', 'Hasura');
INSERT INTO users (username, password, name) VALUES ('exampleuser1', 'examplepassword1', 'Gopher');
