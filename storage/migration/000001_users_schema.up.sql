CREATE TABLE IF NOT EXISTS "users" (
    "id" CHAR(36) PRIMARY KEY NOT NULL,
    "name" CHAR(50) NOT NULL,
    "email" CHAR(100) NOT NULL,
    "password" CHAR(64) NOT NULL
);