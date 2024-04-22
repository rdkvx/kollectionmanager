CREATE TABLE IF NOT EXISTS producer(
    id SERIAL PRIMARY KEY,
    name varchar(100) UNIQUE,
    openDate date UNIQUE
);

CREATE TABLE IF NOT EXISTS console(
    id SERIAL PRIMARY KEY,
    name varchar(100) UNIQUE,
    producerId bigint UNIQUE,
    releaseDate date,
    buyDate date,
    own BOOLEAN,
    FOREIGN KEY(producerId) REFERENCES producer(id)
);

CREATE TABLE IF NOT EXISTS developer(
    id SERIAL PRIMARY KEY,
    name varchar(100) UNIQUE
);

CREATE TABLE IF NOT EXISTS game(
    id SERIAL PRIMARY KEY,
    name varchar(100) UNIQUE,
    consoleId bigint UNIQUE,
    developerId bigint UNIQUE,
    releaseDate date,
    buyDate date,
    FOREIGN KEY(consoleId) REFERENCES console(id),
    FOREIGN KEY(developerId) REFERENCES developer(id)
);