CREATE TABLE "boards" (
    id uuid DEFAULT uuid_generate_v4(),
    name VARCHAR,
    gender INTEGER,
    birthday TIMESTAMP,
    PRIMARY KEY (id)
)