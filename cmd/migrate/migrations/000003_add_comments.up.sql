CREATE TABLE IF NOT EXISTS "comments"(
    "id" uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    "post_id" uuid NOT NULL,
    "user_id" uuid NOT NULL,
    "content" text NOT NULL,
    "created_at" TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    "updated_at" TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    CONSTRAINT fk_post_id
        FOREIGN KEY ("post_id") REFERENCES posts("id"),

    CONSTRAINT fk_user_id
        FOREIGN KEY ("user_id") REFERENCES users("id")
);

CREATE INDEX idx_comment_post_id ON posts(id);
CREATE INDEX idx_comment_user_id ON users(id);