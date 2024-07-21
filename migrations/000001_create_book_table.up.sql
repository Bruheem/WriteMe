CREATE TABLE book (
    id SERIAL PRIMARY KEY,
    title TEXT NOT NULL,
    author TEXT NOT NULL
);

CREATE TABLE chapter (
    id SERIAL PRIMARY KEY,
    book_id BIGINT REFERENCES book(id),
    title TEXT NOT NULL,
    content TEXT NOT NULL
);