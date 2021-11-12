CREATE TABLE "users" (
  "id" varchar PRIMARY KEY,
  "password" varchar NOT NULL,
  "firstname" varchar NOT NULL,
  "lastname" varchar NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "status" tinyint DEFAULT 0,
  "password_changed_at" timestamptz NOT NULL DEFAULT('0001-01-01 00:00:00Z'),  
  "created_at" timestamptz NOT NULL DEFAULT (now())
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);