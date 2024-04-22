CREATE TABLE IF NOT EXISTS producer(
    id SERIAL PRIMARY KEY AUTOINCREMENT,
    name varchar(100) not null,
    openDate date not null,
    PRIMARY KEY(id)
);

CREATE TABLE IF NOT EXISTS console(
    id SERIAL PRIMARY KEY AUTOINCREMENT,
    name varchar(100) not null,
    producerId bigint not null,
    releaseDate date,
    buyDate date,
    own tinyint,
    PRIMARY KEY(id),
    FOREIGN KEY(producerId) REFERENCES producer(id)
);

CREATE TABLE IF NOT EXISTS developer(
    id SERIAL PRIMARY KEY AUTOINCREMENT,
    name varchar(100) not null,
    PRIMARY KEY(id)
);

CREATE TABLE IF NOT EXISTS game(
    id SERIAL PRIMARY KEY AUTOINCREMENT,
    name varchar(100) not null,
    consoleId bigint not null,
    developerId bigint not null,
    releaseDate date,
    buyDate date,
    PRIMARY KEY(id),
    FOREIGN KEY(consoleId) REFERENCES console(id),
    FOREIGN KEY(developerId) REFERENCES developer(id)
);