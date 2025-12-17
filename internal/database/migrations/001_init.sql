BEGIN;
CREATE TYPE user_role AS ENUM ('user', 'admin', 'director');
-- USERS
CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    email TEXT NOT NULL UNIQUE,
    password TEXT NOT NULL,
    role TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);


-- PROFILES

CREATE TABLE IF NOT EXISTS profiles (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    name TEXT NOT NULL,
    age INT NOT NULL,
    preferences JSONB,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),

    CONSTRAINT fk_profiles_user
        FOREIGN KEY (user_id)
        REFERENCES users(id)
        ON DELETE CASCADE
);


-- CONTENT

CREATE TABLE IF NOT EXISTS contents (
    id SERIAL PRIMARY KEY,
    title TEXT NOT NULL,
    category TEXT NOT NULL,
    genre TEXT NOT NULL,
    duration INT NOT NULL,
    year INT NOT NULL,
    artist TEXT,
    rating TEXT,
    min_age INT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);


-- SUBSCRIPTION PLANS

CREATE TABLE IF NOT EXISTS subscription_plans (
    id SERIAL PRIMARY KEY,
    plan TEXT NOT NULL UNIQUE,
    price NUMERIC(10,2) NOT NULL,
    max_profiles INT NOT NULL,
    features JSONB,
    is_active BOOLEAN NOT NULL DEFAULT TRUE
);


-- USER SUBSCRIPTIONS

CREATE TABLE IF NOT EXISTS user_subscriptions (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    plan TEXT NOT NULL,
    active BOOLEAN NOT NULL DEFAULT TRUE,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),

    CONSTRAINT fk_user_subscriptions_user
        FOREIGN KEY (user_id)
        REFERENCES users(id)
        ON DELETE CASCADE
);


-- PAYMENTS

CREATE TABLE IF NOT EXISTS payments (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    plan_id INT NOT NULL,
    amount NUMERIC(10,2) NOT NULL,
    payment_date TIMESTAMP NOT NULL DEFAULT NOW(),
    method TEXT NOT NULL,

    CONSTRAINT fk_payments_user
        FOREIGN KEY (user_id)
        REFERENCES users(id)
        ON DELETE CASCADE,

    CONSTRAINT fk_payments_plan
        FOREIGN KEY (plan_id)
        REFERENCES subscription_plans(id)
);


-- HISTORY

CREATE TABLE IF NOT EXISTS history (
    id SERIAL PRIMARY KEY,
    profile_id INT NOT NULL,
    content_id INT NOT NULL,
    progress INT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),

    CONSTRAINT fk_history_profile
        FOREIGN KEY (profile_id)
        REFERENCES profiles(id)
        ON DELETE CASCADE,

    CONSTRAINT fk_history_content
        FOREIGN KEY (content_id)
        REFERENCES contents(id)
        ON DELETE CASCADE
);


-- FAVORITES

CREATE TABLE IF NOT EXISTS favorites (
    id SERIAL PRIMARY KEY,
    profile_id INT NOT NULL,
    content_id INT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),

    CONSTRAINT fk_favorites_profile
        FOREIGN KEY (profile_id)
        REFERENCES profiles(id)
        ON DELETE CASCADE,

    CONSTRAINT fk_favorites_content
        FOREIGN KEY (content_id)
        REFERENCES contents(id)
        ON DELETE CASCADE,

    UNIQUE (profile_id, content_id)
);


-- RATINGS

CREATE TABLE IF NOT EXISTS ratings (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    content_id INT NOT NULL,
    value NUMERIC(2,1) NOT NULL CHECK (value >= 0 AND value <= 5),
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),

    CONSTRAINT fk_ratings_user
        FOREIGN KEY (user_id)
        REFERENCES users(id)
        ON DELETE CASCADE,

    CONSTRAINT fk_ratings_content
        FOREIGN KEY (content_id)
        REFERENCES contents(id)
        ON DELETE CASCADE,

    UNIQUE (user_id, content_id)
);


-- INDEXES (RENDIMIENTO)

CREATE INDEX IF NOT EXISTS idx_profiles_user_id ON profiles(user_id);
CREATE INDEX IF NOT EXISTS idx_contents_genre ON contents(genre);
CREATE INDEX IF NOT EXISTS idx_history_profile_id ON history(profile_id);
CREATE INDEX IF NOT EXISTS idx_favorites_profile_id ON favorites(profile_id);
CREATE INDEX IF NOT EXISTS idx_ratings_content_id ON ratings(content_id);

COMMIT;