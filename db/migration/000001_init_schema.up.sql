CREATE TABLE users (
    "user_id" SERIAL PRIMARY KEY,
    "name" VARCHAR(100) NOT NULL,
    "created_at" TIMESTAMPTZ NOT NULL DEFAULT (now())
);

CREATE TABLE events (
    "event_id" SERIAL PRIMARY KEY,
    "organizer_id" INT NOT NULL REFERENCES users(user_id),
    "event_name" VARCHAR(200) NOT NULL,
    "event_description" TEXT NOT NULL,
    "duration" INT NOT NULL, -- In hours
    "created_at" TIMESTAMPTZ NOT NULL DEFAULT (now()),
    "updated_at" TIMESTAMPTZ NOT NULL DEFAULT (now())
);

CREATE TABLE time_slots_event (
    "id" SERIAL PRIMARY KEY,
    "event_id" INT NOT NULL REFERENCES events(event_id),
    "start_time" TIMESTAMPTZ NOT NULL,
    "end_time" TIMESTAMPTZ NOT NULL
);

CREATE TABLE time_slots_user (
    "id" SERIAL PRIMARY KEY,
    "user_id" INT NOT NULL REFERENCES users(user_id),
    "start_time" TIMESTAMPTZ NOT NULL,
    "end_time" TIMESTAMPTZ NOT NULL
);