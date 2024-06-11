-- +goose Up
CREATE TABLE tokens (
  value VARCHAR(64) UNIQUE NOT NULL PRIMARY KEY,
  created_at TIMESTAMP NOT NULL,
  expired_at TIMESTAMP NOT NULL,
  is_revoked BOOL NOT NULL DEFAULT false,
  user_id UUID NOT NULL REFERENCES users (id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE tokens;
