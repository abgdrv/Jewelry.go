CREATE TABLE IF NOT EXISTS watches (
    id BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMP(0) with time zone NOT NULL DEFAULT NOW(),
    brand TEXT NULL,
    model TEXT NULL,
    dial_color TEXT NOT NULL,
    strap_type TEXT NOT NULL,
    diameter INT NOT NULL,
    energy TEXT NOT NULL,
    gender TEXT NOT NULL,
    price FLOAT NOT NULL,
    image_url TEXT NOT NULL
);
