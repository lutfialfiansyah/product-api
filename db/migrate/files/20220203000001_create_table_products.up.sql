-- Table Definition ----------------------------------------------

CREATE TABLE products
(
    id          SERIAL PRIMARY KEY,
    name        varchar(50)        DEFAULT NULL,
    price       double precision   DEFAULT 0,
    description text               DEFAULT NULL,
    quantity    double precision   DEFAULT 0,
    created_at  timestamp NOT NULL DEFAULT now(),
    updated_at  timestamp          DEFAULT NULL
);

-- Indices -------------------------------------------------------
