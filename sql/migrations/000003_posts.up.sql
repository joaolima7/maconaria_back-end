DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'post_type') THEN
        EXECUTE 'CREATE TYPE post_type AS ENUM (''notice'',''event'')';
    END IF;
END$$;

CREATE TABLE IF NOT EXISTS posts (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    title TEXT NOT NULL,
    description TEXT,
    date TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
    image TEXT,
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    post_type post_type NOT NULL DEFAULT 'notice'
);

CREATE INDEX IF NOT EXISTS idx_posts_user_id ON posts(user_id);