-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE kusokoras (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    pictureUrl TEXT UNIQUE
);

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE kusokoras;
