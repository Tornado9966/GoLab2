CREATE DATABASE "restaurant_db";
DROP TABLE IF EXISTS "menu_table";
CREATE TABLE "menu_table"
(
    "id"   SERIAL PRIMARY KEY,
    "name" VARCHAR(50) NOT NULL UNIQUE,
	"price" FLOAT NOT NULL
);

-- Insert demo data.
INSERT INTO "menu_table" (name, price) VALUES ('chicken', 123);
INSERT INTO "menu_table" (name, price) VALUES ('beef', 305);
INSERT INTO "menu_table" (name, price) VALUES ('crab', 442);
INSERT INTO "menu_table" (name, price) VALUES ('salad', 95);
INSERT INTO "menu_table" (name, price) VALUES ('lobster', 500);
INSERT INTO "menu_table" (name, price) VALUES ('octopus', 400);
INSERT INTO "menu_table" (name, price) VALUES ('water', 20);
INSERT INTO "menu_table" (name, price) VALUES ('tea', 30);
INSERT INTO "menu_table" (name, price) VALUES ('coffe', 45);
INSERT INTO "menu_table" (name, price) VALUES ('wine', 345);
