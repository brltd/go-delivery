-- Create the 'users' table
CREATE TABLE "users" (
    "id" TEXT DEFAULT (lower(hex(randomblob(4)) || '-' || hex(randomblob(2)) || '-4' || substr(hex(randomblob(2)), 2) || '-' || substr('AB89', 1 + (abs(random()) % 4), 1) || substr(hex(randomblob(2)), 2) || '-' || hex(randomblob(6)))),
    "name" VARCHAR(50) NOT NULL,
    "email" VARCHAR(50) NOT NULL,
    "hash" TEXT NOT NULL,
    "dateCreated" TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "dateUpdated" TEXT NULL,
    PRIMARY KEY ("id")
);

CREATE INDEX id_user_index ON "users" ("id");

CREATE UNIQUE INDEX "users_email_index" ON "users"("email");

-- Create the 'storeStatus' table
CREATE TABLE "storeStatus" (
    "id" TEXT DEFAULT (lower(hex(randomblob(4)) || '-' || hex(randomblob(2)) || '-4' || substr(hex(randomblob(2)), 2) || '-' || substr('AB89', 1 + (abs(random()) % 4), 1) || substr(hex(randomblob(2)), 2) || '-' || hex(randomblob(6)))),
    "name" VARCHAR(50) NOT NULL,
    PRIMARY KEY ("id")
);

CREATE INDEX id_storesStatus_index ON "storeStatus" ("id");

-- Create the 'stores' table with foreign key to 'storeStatus'
CREATE TABLE "stores"(
    "id" TEXT DEFAULT (lower(hex(randomblob(4)) || '-' || hex(randomblob(2)) || '-4' || substr(hex(randomblob(2)), 2) || '-' || substr('AB89', 1 + (abs(random()) % 4), 1) || substr(hex(randomblob(2)), 2) || '-' || hex(randomblob(6)))),
    "name" VARCHAR(50) NOT NULL,
    "imageURL" TEXT NOT NULL,
    "icon" VARCHAR(50) NULL,
    "statusId" TEXT NOT NULL,
    "dateCreated" TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "dateUpdated" TEXT NULL,
    PRIMARY KEY ("id"),
    FOREIGN KEY ("statusId") REFERENCES "storeStatus" ("id") ON DELETE RESTRICT ON UPDATE CASCADE
);

CREATE INDEX id_stores_index ON "stores" ("id");

-- Create the 'categories' table with foreign key to 'stores'
CREATE TABLE "categories" (
    "id" TEXT DEFAULT (lower(hex(randomblob(4)) || '-' || hex(randomblob(2)) || '-4' || substr(hex(randomblob(2)), 2) || '-' || substr('AB89', 1 + (abs(random()) % 4), 1) || substr(hex(randomblob(2)), 2) || '-' || hex(randomblob(6)))),
    "name" VARCHAR(50) NOT NULL,
    "storeId" TEXT NOT NULL,
    "dateCreated" TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "dateUpdated" TEXT NULL,
    PRIMARY KEY ("id"),
    FOREIGN KEY ("storeId") REFERENCES "stores" ("id") ON DELETE RESTRICT ON UPDATE CASCADE
);

CREATE INDEX id_categories_index ON "categories" ("id");

-- Create the 'products' table with foreign keys to 'stores' and 'categories'
CREATE TABLE "products" (
    "id" TEXT DEFAULT (lower(hex(randomblob(4)) || '-' || hex(randomblob(2)) || '-4' || substr(hex(randomblob(2)), 2) || '-' || substr('AB89', 1 + (abs(random()) % 4), 1) || substr(hex(randomblob(2)), 2) || '-' || hex(randomblob(6)))),
    "name" VARCHAR(50) NOT NULL,
    "imageURL" TEXT NOT NULL,
    "price" DECIMAL(65, 30) NOT NULL,
    "promotionPrice" DECIMAL(65, 30) NULL,
    "categoryId" TEXT NOT NULL,
    "storeId" TEXT NOT NULL,
    "dateCreated" TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "dateUpdated" TEXT NULL,
    PRIMARY KEY ("id"),
    FOREIGN KEY ("categoryId") REFERENCES "categories" ("id") ON DELETE RESTRICT ON UPDATE CASCADE,
    FOREIGN KEY ("storeId") REFERENCES "stores" ("id") ON DELETE RESTRICT ON UPDATE CASCADE
);

CREATE INDEX id_products_index ON "products" ("id");
