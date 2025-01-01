CREATE TABLE users (
    "user_id" SERIAL PRIMARY KEY,
    "name" VARCHAR(100) NOT NULL,
    "created_at" TIMESTAMPTZ NOT NULL DEFAULT (now())
);

CREATE TABLE events (
    "event_id" SERIAL PRIMARY KEY,
    "organizer_id" INT NOT NULL REFERENCES users(user_id),
    "event_name" VARCHAR(200) NOT NULL,
    "event_description" TEXT,
    "created_at" TIMESTAMPTZ NOT NULL DEFAULT (now())
);

CREATE TABLE time_preferences (
    "id" SERIAL PRIMARY KEY,
    "owner_type" VARCHAR(50) CHECK (owner_type IN ('user', 'event')),
    "owner_id" INT NOT NULL, -- Links to user_id or event_id
    "start_time" TIMESTAMPTZ NOT NULL,
    "end_time" TIMESTAMPTZ NOT NULL,
    UNIQUE(owner_type, owner_id, start_time, end_time)
);

CREATE TABLE event_participants (
    "event_id" INT NOT NULL REFERENCES events(event_id),
    "user_id" INT NOT NULL REFERENCES users(user_id),
    "can_attend" BOOLEAN DEFAULT NULL,
    PRIMARY KEY (event_id, user_id)
);
