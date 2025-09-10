CREATE TABLE IF NOT EXISTS ordens (
    id varchar(255) NOT NULL,
    preco float NOT NULL,
    taxa float NOT NULL,
    valor float NOT NULL,
    PRIMARY KEY (id)
);