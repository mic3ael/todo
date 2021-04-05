CREATE TABLE users(
    id SERIAL NOT null UNIQUE,
    name VARCHAR(255) NOT null,
    username VARCHAR(255) NOT null UNIQUE,
    password_hash VARCHAR(255) NOT null
);

CREATE TABLE todo_lists(
    id SERIAL NOT null UNIQUE,
    title VARCHAR(255) NOT null,
    description VARCHAR(255)
);

CREATE TABLE users_lists(
    id SERIAL NOT null UNIQUE,
    user_id INT REFERENCES users(id) ON DELETE CASCADE NOT null,
    list_id INT REFERENCES todo_lists(id) ON DELETE CASCADE NOT null
);

CREATE TABLE todo_items(
    id SERIAL NOT null UNIQUE,
    title VARCHAR(255) NOT null,
    description VARCHAR(255),
    done BOOLEAN NOT null DEFAULT false
);

CREATE TABLE lists_items(
    id SERIAL NOT null UNIQUE,
    item_id INT REFERENCES todo_items(id) ON DELETE CASCADE NOT null,
    list_id INT REFERENCES todo_lists(id) ON DELETE CASCADE NOT null
);