-- Drop existing tables if they exist
DROP TABLE IF EXISTS public.users;
DROP TABLE IF EXISTS public.region;
DROP TABLE IF EXISTS public.place;
DROP TABLE IF EXISTS public.holog;
DROP TABLE IF EXISTS public.travel_schedule;
DROP TABLE IF EXISTS public.narration_style;
DROP TABLE IF EXISTS public.narration;
DROP TABLE IF EXISTS public.narration_series;
DROP TABLE IF EXISTS public.announce_post;
DROP TABLE IF EXISTS public.session;

CREATE SCHEMA IF NOT EXISTS public;

CREATE TABLE IF NOT EXISTS public.users (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    email TEXT NOT NULL UNIQUE,
    password TEXT NOT NULL,
    username VARCHAR NOT NULL,
    profile_image_url TEXT,
    created_at DATE DEFAULT CURRENT_DATE,
    deleted_at TIMESTAMP
);

-- CREATE TABLE IF NOT EXISTS public.region (
--     id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
--     name TEXT NOT NULL,
--     address TEXT NOT NULL
-- );

-- CREATE TABLE IF NOT EXISTS public.place (
--     id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
--     region_id INTEGER NOT NULL,
--     title TEXT NOT NULL,
--     address TEXT NOT NULL,
--     thumbnail TEXT,
--     latitude FLOAT,
--     longitude FLOAT,
--     num TEXT,
--     description TEXT,

--     FOREIGN KEY (region_id) REFERENCES public.region(id)
-- );

CREATE TABLE IF NOT EXISTS public.holog (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    place_id INTEGER NOT NULL,
    creator_id INTEGER NOT NULL,
    type TEXT NOT NULL,
    title TEXT NOT NULL,
    content TEXT NOT NULL,
    thumbnail_url TEXT,
    external_url TEXT,
    created_at DATE DEFAULT CURRENT_DATE,
    deleted_at TIMESTAMP,

    -- FOREIGN KEY (place_id) REFERENCES public.place(id),
    FOREIGN KEY (creator_id) REFERENCES public.users(id)
);

-- CREATE TABLE IF NOT EXISTS public.travel_schedule (
--     id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
--     place_id UUID NOT NULL,
--     owner_id UUID NOT NULL,
--     start_date TIMESTAMP NOT NULL,
--     end_date TIMESTAMP NOT NULL,
--     status TEXT NOT NULL, -- enum 타입을 TEXT로 정의
--     created_at DATE DEFAULT CURRENT_DATE,
--     updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
--     deleted_at TIMESTAMP,
--     FOREIGN KEY (place_id) REFERENCES public.place(id),
--     FOREIGN KEY (owner_id) REFERENCES public.users(id)
-- );

-- CREATE TABLE IF NOT EXISTS public.narration_style (
--     id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
--     name TEXT NOT NULL,
--     created_at DATE DEFAULT CURRENT_DATE,
--     updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
--     deleted_at TIMESTAMP
-- );

-- CREATE TABLE IF NOT EXISTS public.narration (
--     id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
--     style_id UUID NOT NULL,
--     creator_id UUID NOT NULL,
--     place_id UUID NOT NULL,
--     series_id UUID NOT NULL,
--     title TEXT NOT NULL,
--     text_ref TEXT,
--     paid_ver BOOLEAN,
--     created_at DATE DEFAULT CURRENT_DATE,
--     updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
--     deleted_at TIMESTAMP,
--     FOREIGN KEY (style_id) REFERENCES public.narration_style(id),
--     FOREIGN KEY (creator_id) REFERENCES public.users(id),
--     FOREIGN KEY (place_id) REFERENCES public.place(id),
--     FOREIGN KEY (series_id) REFERENCES public.narration_series(id)
-- );

-- CREATE TABLE IF NOT EXISTS public.narration_series (
--     id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
--     creator_id UUID NOT NULL,
--     title TEXT NOT NULL,
--     created_at DATE DEFAULT CURRENT_DATE,
--     updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
--     deleted_at TIMESTAMP,
--     FOREIGN KEY (creator_id) REFERENCES public.users(id)
-- );

-- CREATE TABLE IF NOT EXISTS public.announce_post (
--     id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
--     title TEXT NOT NULL,
--     content TEXT NOT NULL,
--     created_at DATE DEFAULT CURRENT_DATE,
--     updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
--     deleted_at TIMESTAMP
-- );

-- CREATE TABLE IF NOT EXISTS public.session (
--     id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
--     user_id UUID NOT NULL,
--     access_token TEXT NOT NULL,
--     created_at DATE DEFAULT CURRENT_DATE,
--     updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
--     deleted_at TIMESTAMP,
--     FOREIGN KEY (user_id) REFERENCES public.users(id)
-- );
