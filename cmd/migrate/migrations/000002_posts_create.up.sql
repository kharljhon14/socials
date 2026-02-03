CREATE TABLE IF NOT EXISTS "posts"(
    "id" uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    "title" text NOT NULL,
    "user_id" uuid NOT NULL,
    "content" text NOT NULL,
    "created_at" TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    "updated_at" TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    CONSTRAINT fk_user_id
        FOREIGN KEY ("user_id") REFERENCES "users"(id)
);

CREATE INDEX idx_posts_user_id ON users(id);