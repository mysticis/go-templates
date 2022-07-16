CREATE TYPE "taskstatus" AS ENUM (
  'COMPLETED',
  'PENDING',
  'DELETED'
);

CREATE TABLE "users" (
  "id" bigserial PRIMARY KEY,
  "username" varchar UNIQUE NOT NULL,
  "password" varchar NOT NULL,
  "email" varchar UNIQUE NOT NULL
);

CREATE TABLE "category" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL,
  "user_id" bigint NOT NULL
);

CREATE TABLE "status" (
  "id" bigserial PRIMARY KEY,
  "task_status" taskstatus NOT NULL
);

CREATE TABLE "task" (
  "id" bigserial PRIMARY KEY,
  "title" varchar NOT NULL,
  "content" text NOT NULL,
  "created_date" timestamp NOT NULL DEFAULT (now()),
  "last_modified_at" timestamp NOT NULL DEFAULT (now()),
  "finish_date" timestamp NOT NULL DEFAULT (now()),
  "priority" integer,
  "category_id" bigint,
  "task_status_id" integer,
  "due_date" timestamp NOT NULL DEFAULT (now()),
  "user_id" bigint,
  "hide" int
);

CREATE TABLE "comments" (
  "id" bigserial PRIMARY KEY,
  "content" text NOT NULL,
  "taskID" bigint NOT NULL,
  "created" timestamp NOT NULL,
  "user_id" integer
);

CREATE TABLE "files" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL,
  "autoName" varchar NOT NULL,
  "user_id" bigint,
  "created_date" timestamp NOT NULL DEFAULT (now())
);

CREATE INDEX ON "users" ("username");

CREATE INDEX ON "category" ("name");

CREATE INDEX ON "task" ("title");

CREATE INDEX ON "task" ("user_id");

CREATE INDEX ON "comments" ("taskID");

CREATE INDEX ON "comments" ("user_id");

CREATE INDEX ON "files" ("name");

ALTER TABLE "category" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "task" ADD FOREIGN KEY ("category_id") REFERENCES "category" ("id");

ALTER TABLE "task" ADD FOREIGN KEY ("task_status_id") REFERENCES "status" ("id");

ALTER TABLE "task" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "comments" ADD FOREIGN KEY ("taskID") REFERENCES "task" ("id");

ALTER TABLE "comments" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "files" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");
