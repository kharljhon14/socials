CREATE EXTENSION IF NOT EXISTS pgcrypto;

CREATE EXTENSION IF NOT EXISTS citext;

CREATE TABLE IF NOT EXISTS "users"(
    "id" uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    "username" varchar(255) UNIQUE NOT NULL,
    "email" citext UNIQUE NOT NULL,
    "password" bytea NOT NULL,
    "created_at" TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    "updated_at" TIMESTAMPTZ NOT NULL DEFAULT NOW()
);