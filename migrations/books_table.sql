
CREATE TABLE IF NOT EXISTS books(
    id serial NOT NULL CONSTRAINT books_pk PRIMARY KEY,
    name VARCHAR (100) NOT NULL
);

CREATE UNIQUE INDEX IF NOT EXISTS books_id_uindex on books (id);