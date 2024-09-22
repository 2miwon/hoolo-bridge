-- Drop existing tables if they exist
DROP TABLE IF EXISTS public.travel_schedule;
DROP TABLE IF EXISTS public.narration_style;
DROP TABLE IF EXISTS public.narration;
DROP TABLE IF EXISTS public.narration_series;
DROP TABLE IF EXISTS public.announce_post;
DROP TABLE IF EXISTS public.session;
DROP TABLE IF EXISTS public.holog;
DROP TABLE IF EXISTS public.region;
DROP TABLE IF EXISTS public.users;

CREATE SCHEMA IF NOT EXISTS public;

CREATE TABLE IF NOT EXISTS public.users (
    id TEXT PRIMARY KEY NOT NULL,
    password TEXT NOT NULL,
    username VARCHAR NOT NULL,
    profile_image_url TEXT,
    created_at DATE DEFAULT CURRENT_DATE,
    deleted_at TIMESTAMPTZ DEFAULT NULL
);

CREATE TABLE IF NOT EXISTS public.schedule (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id TEXT NOT NULL,
    start_date TIMESTAMP NOT NULL,
    end_date TIMESTAMP NOT NULL,
    created_at DATE DEFAULT CURRENT_DATE,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,

    FOREIGN KEY (user_id) REFERENCES public.users(id)
);

CREATE TABLE IF NOT EXISTS public.schedule_detail (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    schedule_id UUID NOT NULL,
    place_id TEXT NOT NULL,
    created_at DATE DEFAULT CURRENT_DATE,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,

    FOREIGN KEY (schedule_id) REFERENCES public.schedule(id)
);

CREATE TABLE IF NOT EXISTS public.holog (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    place_id TEXT NOT NULL,
    creator_id TEXT NOT NULL,
    schedule_id UUID,
    type TEXT NOT NULL DEFAULT 'holog', -- tistory / naver
    title TEXT NOT NULL,
    content TEXT NOT NULL,
    thumbnail_url TEXT,
    image_url TEXT,
    external_url TEXT,
    created_at DATE DEFAULT CURRENT_DATE,
    deleted_at TIMESTAMPTZ,

    FOREIGN KEY (creator_id) REFERENCES public.users(id)
);

CREATE TABLE IF NOT EXISTS public.bookmark (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id TEXT NOT NULL,
    holog_id INTEGER NOT NULL,

    FOREIGN KEY (user_id) REFERENCES public.users(id),
    FOREIGN KEY (place_id) REFERENCES public.holog(id)
);

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

CREATE TABLE IF NOT EXISTS public.announce (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    title TEXT NOT NULL,
    content TEXT NOT NULL,
    created_at DATE DEFAULT CURRENT_DATE,
    deleted_at TIMESTAMP
);

-- CREATE TABLE IF NOT EXISTS public.session (
--     id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
--     user_id UUID NOT NULL,
--     access_token TEXT NOT NULL,
--     created_at DATE DEFAULT CURRENT_DATE,
--     updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
--     deleted_at TIMESTAMP,
--     FOREIGN KEY (user_id) REFERENCES public.users(id)
-- );
