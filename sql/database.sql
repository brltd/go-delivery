CREATE EXTENSION IF NOT EXISTS "uuid-ossp";


CREATE TABLE "users" (
    "id" uuid DEFAULT uuid_generate_v4(),
    "name" VARCHAR(50) NOT NULL,
    "email" VARCHAR(50) NOT NULL,
    "hash" TEXT NOT NULL,
    "dateCreated" TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "dateUpdated" TIMESTAMP WITH TIME ZONE NULL,
    CONSTRAINT "user_pk" PRIMARY KEY ("id")
);

CREATE INDEX id_user_index ON "users" ("id");

CREATE UNIQUE INDEX "users_email_index" ON "users"("email");

CREATE TABLE "stores"(
    "id" uuid DEFAULT uuid_generate_v4(),
    "name" VARCHAR(50) NOT NULL,
    "imageURL" TEXT NOT NULL,
    "icon" VARCHAR(50) NULL,
    "statusId" uuid NOT NULL,
    "dateCreated" TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "dateUpdated" TIMESTAMP WITH TIME ZONE NULL,
    CONSTRAINT "stores_pk" PRIMARY KEY ("id")
);

CREATE INDEX id_stores_index ON "stores" ("id");

CREATE TABLE "storeStatus" (
    "id" uuid DEFAULT uuid_generate_v4(),
    "name" VARCHAR(50) NOT NULL,
    CONSTRAINT "storeStatus_pk" PRIMARY KEY ("id")
);

CREATE INDEX id_storesStatus_index ON "storeStatus" ("id");

ALTER TABLE
    "stores"
ADD
    CONSTRAINT "stores_statusId_fkey" FOREIGN KEY ("statusId") REFERENCES "storeStatus"("id") ON DELETE RESTRICT ON UPDATE CASCADE;

CREATE TABLE "categories" (
    "id" uuid DEFAULT uuid_generate_v4(),
    "name" VARCHAR(50) NOT NULL,
    "storeId" uuid NOT NULL,
    "dateCreated" TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "dateUpdated" TIMESTAMP WITH TIME ZONE NULL,
    CONSTRAINT "categories_pk" PRIMARY KEY ("id")
);

CREATE INDEX id_categories_index ON "categories" ("id");

CREATE TABLE "products" (
    "id" uuid DEFAULT uuid_generate_v4(),
    "name" VARCHAR(50) NOT NULL,
    "imageURL" TEXT NOT NULL,
    "price" DECIMAL(65, 30) NOT NULL,
    "promotionPrice" DECIMAL(65, 30) NULL,
    "categoryId" uuid NOT NULL,
    "storeId" uuid NOT NULL,
    "dateCreated" TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "dateUpdated" TIMESTAMP WITH TIME ZONE NULL,
    CONSTRAINT "products_pk" PRIMARY KEY ("id")
);

CREATE INDEX id_products_index ON "products" ("id");

ALTER TABLE
    "products"
ADD
    CONSTRAINT "stores_statusId_fk" FOREIGN KEY ("storeId") REFERENCES "stores"("id") ON DELETE RESTRICT ON UPDATE CASCADE;

ALTER TABLE
    "products"
ADD
    CONSTRAINT "products_categoryId_fk" FOREIGN KEY ("categoryId") REFERENCES "categories"("id") ON DELETE RESTRICT ON UPDATE CASCADE;