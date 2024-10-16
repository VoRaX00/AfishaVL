-- +goose Up
-- +goose StatementBegin
CREATE TABLE users (
    id INT PRIMARY KEY,
    phone TEXT NOT NULL,
    password_hash TEXT NOT NULL
);

CREATE TABLE refresh_tokens (
    id SERIAL PRIMARY KEY,
    user_id UUID REFERENCES users(id) ON DELETE CASCADE NOT NULL,
    refresh_token_hash TEXT UNIQUE NOT NULL,
    ip VARCHAR(45) NOT NULL,
    expires_at TIMESTAMPTZ NOT NULL,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE events (
    id INT PRIMARY KEY,
    user_id INT REFERENCES users(id) ON DELETE CASCADE NOT NULL,
    name TEXT NOT NULL,
    description TEXT,
    city_id INT REFERENCES cities (id) ON DELETE CASCADE NOT NULL,
    category_id INT REFERENCES categories (id) ON DELETE CASCADE NOT NULL,
    image bytea NULL,
    price_id INT REFERENCES prices(id) ON DELETE CASCADE NOT NULL,
    age_limit INT NOT NULL,
    schedule_id INT REFERENCES schedules(id) ON DELETE CASCADE NOT NULL,
    registration BOOLEAN NOT NULL DEFAULT(false),
    contact_id INT REFERENCES contacts(id) ON DELETE CASCADE NOT NULL
);

CREATE TABLE cities (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL
);

CREATE TABLE categories (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL
);

CREATE TABLE prices (
    id SERIAL PRIMARY KEY,
    price_from INT NULL,
    price_to INT NULL,
    comment TEXT NULL,
    free boolean NOT NULL
);

CREATE TABLE schedules (
    id SERIAL PRIMARY KEY,
    place TEXT NOT NULL,
    address TEXT NOT NULL,
    dates TIMESTAMP[] NOT NULL
);

CREATE TABLE contacts (
    id SERIAL PRIMARY KEY,
    email TEXT NOT NULL,
    extra_contacts TEXT NOT NULL
);

INSERT INTO cities (name) VALUES
    ('Владивосток'), ('Спасск-Дальний'), ('Артём'), ('Сибирцево'), ('Приморский край, с. Анисимовка'),
    ('Арсеньев'), ('Уссурийск'), ('Ливадия'), ('Приморский край, с. Утёсное'), ('Покровка'),
    ('Приморский край, с. Новокачалинск'), ('Михайловка'), ('п. Веневитиново'), ('Вольно-Надеждинское'),
    ('Тавричанка'), ('о. Русский'), ('Дальнегорск'), ('Пограничный'), ('Находка'), ('Врангель'),
    ('Штыково'), ('Приморский край, с.Молчановка, ул. с. Молчановка'), ('Приморский край, с. Таёжка'), ('Западный'),
    ('Андреевка'), ('Черниговка'), ('Дальнереченск'), ('Золотая Долина'), ('Владимиро-Александровское'),
    ('Партизанск');

INSERT INTO categories (name) VALUES
    ('Концерты'), ('Театры'), ('Вечеринки'), ('Спорт'), ('Детям'),
    ('Музеи и галереи'),('Цирк'), ('Городское событие'), ('Обучение'), ('Впечатления'),
    ('Экскурсие и туры'), ('Мастер-классы'), ('Стендап');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE cities;
DROP TABLE categories;
DROP TABLE prices;
DROP TABLE schedules;
DROP TABLE contacts;
DROP TABLE users;
DROP TABLE events
-- +goose StatementEnd
