-- +goose Up

-- +goose StatementBegin
CREATE TABLE "user"(
    id SERIAL PRIMARY KEY,
    login text NOT NULL,
    password text NOT NULL,
    name text NOT NULL,
    age int NOT NULL,
    created_at timestamp NOT NULL default now(),
    updated_at timestamp,
    salt uuid NOT NULL
);
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TABLE post(
    id SERIAL PRIMARY KEY,
    author_id int NOT NULL,
    created_at timestamp NOT NULL default now(),
    updated_at timestamp,
    allow_comments boolean NOT NULL,
    FOREIGN KEY (author_id) REFERENCES "user" (id)
);

CREATE INDEX posts_created_at_index ON post(created_at);
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TABLE comment(
    id SERIAL PRIMARY KEY,
    post_id int NOT NULL,
    msg text NOT NULL,
    par_id int,
    author_id int NOT NULL,
    created_at timestamp NOT NULL default now(),
    updated_at timestamp,
    FOREIGN KEY (post_id) REFERENCES post (id),
    FOREIGN KEY (author_id) REFERENCES "user" (id),
    FOREIGN KEY (par_id) REFERENCES  comment (id)
);

CREATE INDEX comment_post_index ON comment(post_id,created_at);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE comment;
-- +goose StatementEnd

-- +goose StatementBegin
DROP TABLE post;
-- +goose StatementEnd

-- +goose StatementBegin
DROP TABLE "user"
-- +goose StatementEnd
