-- Adminer 4.8.1 PostgreSQL 16.1 (Debian 16.1-1.pgdg120+1) dump



CREATE DATABASE "blob";

\connect "blob";

DROP TABLE IF EXISTS "buckets";
DROP SEQUENCE IF EXISTS buckets_id_seq;
CREATE SEQUENCE buckets_id_seq INCREMENT 1 MINVALUE 1 MAXVALUE 2147483647 CACHE 1;

CREATE TABLE "public"."buckets" (
    "id" integer DEFAULT nextval('buckets_id_seq') NOT NULL,
    "name" text NOT NULL,
    "type" text NOT NULL,
    CONSTRAINT "buckets_pkey" PRIMARY KEY ("id")
) WITH (oids = false);


DROP TABLE IF EXISTS "files";
DROP SEQUENCE IF EXISTS files_id_seq;
CREATE SEQUENCE files_id_seq INCREMENT 1 MINVALUE 1 MAXVALUE 2147483647 CACHE 1;

CREATE TABLE "public"."files" (
    "id" integer DEFAULT nextval('files_id_seq') NOT NULL,
    "name" text NOT NULL,
    "path" text NOT NULL,
    "mime" text NOT NULL,
    "size" text NOT NULL,
    "uploaded" timestamp DEFAULT CURRENT_TIMESTAMP(0) NOT NULL,
    CONSTRAINT "files_pkey" PRIMARY KEY ("id")
) WITH (oids = false);


DROP TABLE IF EXISTS "filesToBuckets";
DROP SEQUENCE IF EXISTS "filesToBuckets_id_seq";
CREATE SEQUENCE "filesToBuckets_id_seq" INCREMENT 1 MINVALUE 1 MAXVALUE 2147483647 CACHE 1;

CREATE TABLE "public"."filesToBuckets" (
    "id" integer DEFAULT nextval('"filesToBuckets_id_seq"') NOT NULL,
    "bucketID" integer NOT NULL,
    "bucketName" text NOT NULL,
    "fileID" integer NOT NULL,
    "fileName" text NOT NULL,
    CONSTRAINT "filesToBuckets_pkey" PRIMARY KEY ("id")
) WITH (oids = false);


ALTER TABLE ONLY "public"."filesToBuckets" ADD CONSTRAINT "filesToBuckets_bucketID_fkey" FOREIGN KEY ("bucketID") REFERENCES buckets(id) ON UPDATE CASCADE ON DELETE CASCADE NOT DEFERRABLE;
ALTER TABLE ONLY "public"."filesToBuckets" ADD CONSTRAINT "filesToBuckets_fileID_fkey" FOREIGN KEY ("fileID") REFERENCES files(id) ON UPDATE CASCADE ON DELETE CASCADE NOT DEFERRABLE;

-- 2023-12-08 23:14:19.636381+00