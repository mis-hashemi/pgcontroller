-- +migrate Up
CREATE TABLE users (
id uuid DEFAULT uuid_generate_v4() PRIMARY KEY ,
first_name VARCHAR(255),
last_name VARCHAR(255),
password VARCHAR(255) NOT NULL,
phone_number VARCHAR(255) NOT NULL UNIQUE
);

-- +migrate Down
DROP TABLE users;