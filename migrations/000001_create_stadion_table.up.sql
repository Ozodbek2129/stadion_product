CREATE EXTENSION IF NOT EXISTS postgis;

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

CREATE TABLE IF NOT EXISTS stadium_images(
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    stadium_id UUID NOT NULL references stadium(id),
    image_url TEXT NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at BIGINT DEFAULT 0
);

CREATE TABLE IF NOT EXISTS order_stadium(
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL,
    stadium_id UUID NOT NULL references stadium(id),
    date DATE NOT NULL,
    start_time TIME NOT NULL,
    end_time TIME NOT NULL,
    price FLOAT NOT NULL,
    status VARCHAR(50) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at BIGINT DEFAULT 0
);

CREATE OR REPLACE FUNCTION update_deleted_at()
RETURNS TRIGGER AS $$
BEGIN
    -- Agar hozirgi vaqt end_time dan katta bo'lsa
    IF CURRENT_TIMESTAMP >= (NEW.date + NEW.end_time) THEN
        NEW.deleted_at = EXTRACT(EPOCH FROM CURRENT_TIMESTAMP); -- deleted_at maydoniga vaqtni yozish
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Trigger yaratish
CREATE TRIGGER update_deleted_at_trigger
BEFORE INSERT OR UPDATE ON order_stadium
FOR EACH ROW
EXECUTE FUNCTION update_deleted_at();