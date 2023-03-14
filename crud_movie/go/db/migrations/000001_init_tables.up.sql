CREATE TABLE "Movie" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL,
  "year" int NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "Filmmaker" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL,
  "dob" timestamptz NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "Role" (
  "id" bigserial PRIMARY KEY,
  "name" varchar UNIQUE NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "MovieCredits" (
  "id" bigserial PRIMARY KEY,
  "filmmaker_id" bigint NOT NULL,
  "movie_id" bigint NOT NULL,
  "role_id" bigint NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE INDEX ON "Movie" ("name");

CREATE INDEX ON "Filmmaker" ("name");

CREATE INDEX ON "Role" ("name");

CREATE INDEX ON "MovieCredits" ("filmmaker_id");

CREATE INDEX ON "MovieCredits" ("movie_id");

CREATE INDEX ON "MovieCredits" ("role_id");

CREATE INDEX ON "MovieCredits" ("filmmaker_id", "movie_id");

CREATE UNIQUE INDEX ON "MovieCredits" ("filmmaker_id", "movie_id", "role_id");

COMMENT ON COLUMN "Role"."name" IS 'director, actor, etc';

ALTER TABLE "MovieCredits" ADD FOREIGN KEY ("filmmaker_id") REFERENCES "Filmmaker" ("id");

ALTER TABLE "MovieCredits" ADD FOREIGN KEY ("movie_id") REFERENCES "Movie" ("id");

ALTER TABLE "MovieCredits" ADD FOREIGN KEY ("role_id") REFERENCES "Role" ("id");
