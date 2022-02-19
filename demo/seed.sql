CREATE TABLE "users" (
    "id" TEXT PRIMARY KEY,
    "created_at" TIMESTAMPTZ NOT NULL DEFAULT now(),
    "username" TEXT NOT NULL
);

CREATE TABLE "articles" (
    "id" TEXT PRIMARY KEY,
    "created_at" TIMESTAMPTZ NOT NULL DEFAULT now(),
    "title" TEXT NOT NULL,
    "content" TEXT NOT NULL,
    "author" TEXT NOT NULL,
    FOREIGN KEY ("author") REFERENCES "users" ("id")
);

CREATE TABLE "comments" (
    "id" TEXT PRIMARY KEY,
    "created_at" TIMESTAMPTZ NOT NULL DEFAULT now(),
    "content" TEXT NOT NULL,
    "article" TEXT NOT NULL,
    "author" TEXT NOT NULL,
    FOREIGN KEY ("article") REFERENCES "articles" ("id"),
    FOREIGN KEY ("author") REFERENCES "users" ("id")
);

CREATE TABLE "articles_tags" (
    "label" TEXT NOT NULL,
    "article" TEXT NOT NULL,
    PRIMARY KEY ("label", "article"),
    FOREIGN KEY ("article") REFERENCES "articles" ("id")
);

CREATE TABLE "likes" (
    "user" TEXT NOT NULL,
    "article" TEXT NOT NULL,
    PRIMARY KEY ("user", "article"),
    FOREIGN KEY ("user") REFERENCES "users" ("id"),
    FOREIGN KEY ("article") REFERENCES "articles" ("id")
);

-- Users
INSERT INTO "users" ("id", "created_at", "username")
    VALUES ('u1', NOW(), 'user1');
INSERT INTO "users" ("id", "created_at", "username")
    VALUES ('u2', NOW(), 'user2');
INSERT INTO "users" ("id", "created_at", "username")
    VALUES ('u3', NOW(), 'user3');

-- Articles
INSERT INTO "articles" ("id", "created_at", "title", "content", "author")
    VALUES ('a1', NOW(), 'Article 1', 'This is the content of article 1.', 'u1');
INSERT INTO "articles" ("id", "created_at", "title", "content", "author")
    VALUES ('a2', NOW(), 'Article 2', 'This is the content of article 2.', 'u1');
INSERT INTO "articles" ("id", "created_at", "title", "content", "author")
    VALUES ('a3', NOW(), 'Article 3', 'This is the content of article 3.', 'u2');
INSERT INTO "articles" ("id", "created_at", "title", "content", "author")
    VALUES ('a4', NOW(), 'Article 4', 'This is the content of article 4.', 'u2');
INSERT INTO "articles" ("id", "created_at", "title", "content", "author")
    VALUES ('a5', NOW(), 'Article 5', 'This is the content of article 5.', 'u3');

-- Comments
INSERT INTO "comments" ("id", "created_at", "content", "article", "author")
    VALUES ('c1', NOW(), 'This is comment 1.', 'a1', 'u1');
INSERT INTO "comments" ("id", "created_at", "content", "article", "author")
    VALUES ('c2', NOW(), 'This is comment 2.', 'a1', 'u1');
INSERT INTO "comments" ("id", "created_at", "content", "article", "author")
    VALUES ('c3', NOW(), 'This is comment 3.', 'a2', 'u2');
INSERT INTO "comments" ("id", "created_at", "content", "article", "author")
    VALUES ('c4', NOW(), 'This is comment 4.', 'a2', 'u2');
INSERT INTO "comments" ("id", "created_at", "content", "article", "author")
    VALUES ('c5', NOW(), 'This is comment 5.', 'a3', 'u1');
INSERT INTO "comments" ("id", "created_at", "content", "article", "author")
    VALUES ('c6', NOW(), 'This is comment 5.', 'a4', 'u3');
INSERT INTO "comments" ("id", "created_at", "content", "article", "author")
    VALUES ('c7', NOW(), 'This is comment 5.', 'a4', 'u2');
INSERT INTO "comments" ("id", "created_at", "content", "article", "author")
    VALUES ('c8', NOW(), 'This is comment 5.', 'a4', 'u1');

-- Tags
INSERT INTO "articles_tags" ("label", "article")
    VALUES ('tech', 'a1');
INSERT INTO "articles_tags" ("label", "article")
    VALUES ('tech', 'a2');
INSERT INTO "articles_tags" ("label", "article")
    VALUES ('sports', 'a1');
INSERT INTO "articles_tags" ("label", "article")
    VALUES ('sports', 'a3');
INSERT INTO "articles_tags" ("label", "article")
    VALUES ('tech', 'a4');
INSERT INTO "articles_tags" ("label", "article")
    VALUES ('health', 'a4');
INSERT INTO "articles_tags" ("label", "article")
    VALUES ('politics', 'a5');
INSERT INTO "articles_tags" ("label", "article")
    VALUES ('health', 'a5');

-- Likes
INSERT INTO "likes" ("user", "article")
    VALUES ('u1', 'a1');
INSERT INTO "likes" ("user", "article")
    VALUES ('u2', 'a2');
INSERT INTO "likes" ("user", "article")
    VALUES ('u1', 'a4');
INSERT INTO "likes" ("user", "article")
    VALUES ('u2', 'a1');
INSERT INTO "likes" ("user", "article")
    VALUES ('u3', 'a4');
INSERT INTO "likes" ("user", "article")
    VALUES ('u3', 'a5');
INSERT INTO "likes" ("user", "article")
    VALUES ('u2', 'a3');
INSERT INTO "likes" ("user", "article")
    VALUES ('u1', 'a5');
