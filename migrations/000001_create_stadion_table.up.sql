CREATE TABLE IF NOT EXISTS stadium(
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL,
    name VARCHAR(50) NOT NULL,
    location GEOGRAPHY(Point, 4326) NOT NULL,
    address TEXT NOT NULL,
    phonenummer VARCHAR(50) NOT NULL,
    price FLOAT NOT NULL,
    length FLOAT NOT NULL,
    width FLOAT NOT NULL,
    situation VARCHAR(255) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at BIGINT DEFAULT 0
);

CREATE TABLE IF NOT EXISTS stadium_image(
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    stadium_id UUID NOT NULL references stadium(id),
    image_url TEXT NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at BIGINT DEFAULT 0
);