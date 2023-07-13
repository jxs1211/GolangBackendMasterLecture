- P11- Design DB schema and generate SQL code with dbdiagram.io

  postgres sql statement

  ```sql
  CREATE TABLE "accounts" (
    "id" bigserial PRIMARY KEY,
    "owner" varchar NOT NULL,
    "balance" bigint NOT NULL,
    "currency" varchar NOT NULL,
    "created_at" timestamptz NOT NULL DEFAULT (now())
  );
  
  CREATE TABLE "entries" (
    "id" bigserial PRIMARY KEY,
    "account_id" bigint,
    "amount" bigint NOT NULL,
    "created_at" timestamptz NOT NULL DEFAULT (now())
  );
  
  CREATE TABLE "transfers" (
    "id" bigserial PRIMARY KEY,
    "from_account_id" bigint,
    "to_account_id" bigint,
    "amount" bigint NOT NULL,
    "created_at" timestamptz NOT NULL DEFAULT (now())
  );
  
  CREATE INDEX ON "accounts" ("owner");
  
  CREATE INDEX ON "entries" ("account_id");
  
  CREATE INDEX ON "transfers" ("from_account_id");
  
  CREATE INDEX ON "transfers" ("to_account_id");
  
  CREATE INDEX ON "transfers" ("from_account_id", "to_account_id");
  
  COMMENT ON COLUMN "entries"."amount" IS 'can be negative';
  
  COMMENT ON COLUMN "transfers"."amount" IS 'must be postive';
  
  ALTER TABLE "entries" ADD FOREIGN KEY ("account_id") REFERENCES "accounts" ("id");
  
  ALTER TABLE "transfers" ADD FOREIGN KEY ("from_account_id") REFERENCES "accounts" ("id");
  
  ALTER TABLE "transfers" ADD FOREIGN KEY ("to_account_id") REFERENCES "accounts" ("id");
  
  ```

  mysql statement

  ```sql
  CREATE TABLE `accounts` (
    `id` BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    `owner` varchar(255) NOT NULL,
    `balance` bigint NOT NULL,
    `currency` varchar(255) NOT NULL,
    `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
  );
  
  CREATE TABLE `entries` (
    `id` BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    `account_id` bigint,
    `amount` bigint NOT NULL COMMENT 'can be negative',
    `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
  );
  
  CREATE TABLE `transfers` (
    `id` BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    `from_account_id` bigint,
    `to_account_id` bigint,
    `amount` bigint NOT NULL COMMENT 'must be postive',
    `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
  );
  
  CREATE INDEX `accounts_index_0` ON `accounts` (`owner`);
  
  CREATE INDEX `entries_index_1` ON `entries` (`account_id`);
  
  CREATE INDEX `transfers_index_2` ON `transfers` (`from_account_id`);
  
  CREATE INDEX `transfers_index_3` ON `transfers` (`to_account_id`);
  
  CREATE INDEX `transfers_index_4` ON `transfers` (`from_account_id`, `to_account_id`);
  
  ALTER TABLE `entries` ADD CONSTRAINT `fk_entries_account_id` FOREIGN KEY (`account_id`) REFERENCES `accounts` (`id`);
  
  ALTER TABLE `transfers` ADD CONSTRAINT fk_from_account FOREIGN KEY (from_account_id) REFERENCES `accounts` (id);
  
  ALTER TABLE `transfers` ADD CONSTRAINT fk_to_account FOREIGN KEY (to_account_id) REFERENCES `accounts` (id);
  
  ```

  

- P22- Install & use Docker + Postgres + TablePlus to create DB schema

  install migrate tool

  ```
  $ wget https://github.com/golang-migrate/migrate/releases/download/v4.16.2/migrate.linux-amd64.tar.gz
  tar -xvzf migrate.linux-amd64.tar.gz
  mv migrate /usr/local/bin/
  ```

  

- P33- How to write - run database migration in Golang

  create docker compose file

  ```yaml
  version: '3.1'
  
  services:
    postgres:
      networks:
        wave-net:
          aliases:
            - postgres
      deploy:
        resources:
          limits:
            cpus: '0.50'
            memory: 500M
          reservations:
            cpus: '0.25'
            memory: 200M
      image: postgres:12-alpine
      container_name: postgres
      environment:
        POSTGRES_USER: root
        POSTGRES_PASSWORD: root
      ports:
      - 5432:5432
      volumes:
      - /home/going/docker-env/data/postgres/pgdata:/var/lib/postgresql/data
  networks:
    wave-net:
      external: true
  ```

  create migrate scripts using migrate

- ```sh
  [going@dev GolangBackendMasterLecture]$ pwd
  /home/going/workspace/GolangBackendMasterLecture
  [going@dev GolangBackendMasterLecture]$ gs
  On branch master
  #
  mkdir -p db/migration
  migrate create -ext sql -dir db/migration -seq init_schema
  [going@dev GolangBackendMasterLecture]$ tree db/
  db/
  ├── migration
  │   ├── 000001_init_schema.down.sql
  │   └── 000001_init_schema.up.sql
  ```

  copy or write sql statement to migration script

  ```sql
  [going@dev GolangBackendMasterLecture]$ cat db/migration/000001_init_schema.up.sql
  CREATE TABLE "accounts" (
    "id" bigserial PRIMARY KEY,
    "owner" varchar NOT NULL,
    "balance" bigint NOT NULL,
    "currency" varchar NOT NULL,
    "created_at" timestamptz NOT NULL DEFAULT (now())
  );
  
  CREATE TABLE "entries" (
    "id" bigserial PRIMARY KEY,
    "account_id" bigint,
    "amount" bigint NOT NULL,
    "created_at" timestamptz NOT NULL DEFAULT (now())
  );
  
  CREATE TABLE "transfers" (
    "id" bigserial PRIMARY KEY,
    "from_account_id" bigint,
    "to_account_id" bigint,
    "amount" bigint NOT NULL,
    "created_at" timestamptz NOT NULL DEFAULT (now())
  );
  
  ALTER TABLE "entries" ADD FOREIGN KEY ("account_id") REFERENCES "accounts" ("id");
  
  ALTER TABLE "transfers" ADD FOREIGN KEY ("from_account_id") REFERENCES "accounts" ("id");
  
  ALTER TABLE "transfers" ADD FOREIGN KEY ("to_account_id") REFERENCES "accounts" ("id");
  
  CREATE INDEX ON "accounts" ("owner");
  
  CREATE INDEX ON "entries" ("account_id");
  
  CREATE INDEX ON "transfers" ("from_account_id");
  
  CREATE INDEX ON "transfers" ("to_account_id");
  
  CREATE INDEX ON "transfers" ("from_account_id", "to_account_id");
  
  COMMENT ON COLUMN "entries"."amount" IS 'can be negative or postive';
  
  COMMENT ON COLUMN "transfers"."amount" IS 'must be postive';
  [going@dev GolangBackendMasterLecture]$ cat db/migration/000001_init_schema.down.sql
  DROP TABLE IF EXISTS accounts CASCADE;
  DROP TABLE IF EXISTS entries CASCADE;
  DROP TABLE IF EXISTS transfers CASCADE;
  ```

  create makefile to manage the project

  ```makefile
  [going@dev GolangBackendMasterLecture]$ cat makefile
  dbup:
  	docker compose up -d postgres
  dbdown:
  	docker compose down 
  createdb:
  	docker exec -it postgres createdb --username=root --owner=root simple_bank
  dropdb:
  	docker exec -it postgres dropdb simple_bank
  migrateup:
  	migrate -path simplebank/db/migration/ -database postgres://root:root@localhost:5432/simple_bank?sslmode=disable up
  migratedown:
  	migrate -path simplebank/db/migration/ -database postgres://root:root@localhost:5432/simple_bank?sslmode=disable down
  
  .PHONY: dbup dbdown createdb dropdb migrateup migratedown
  ```

  

- P44- Generate CRUD Golang code from SQL Compare dbsql- gorm- sqlx - sqlc

  21:22

- P55- Write Golang unit tests for database CRUD with random data

  20:05

- P6课程特别提醒

  00:30

- P76- A clean way to implement database transaction in Golang

  19:53

- P87- DB transaction lock - How to handle deadlock in Golang

  28:21

- P98- How to avoid deadlock in DB transaction Queries order matters-

  13:04

- P109- Understand isolation levels - read phenomena in MySQL - PostgreSQL via examp

  29:37

- P1110- Setup Github Actions for Golang - Postgres to run automated tests

  19:27

- P1211- Implement RESTful HTTP API in Go using Gin【Building RESTful HTTP JSON API】

  25:22

- P1312- Load config from file - environment variables in Golang with Viper

  09:33

- P1413- Mock DB for testing HTTP API in Go and achieve 100- coverage

  26:55

- P1514- Implement transfer money API with a custom params validator

  14:25

- P1615- Add users table with unique - foreign key constraints in PostgreSQL

  14:09

- P1716- How to handle DB errors in Golang correctly

  11:13

- P1817- How to securely store passwords Hash password in Go with Bcrypt-

  16:54

- P1918- How to write stronger unit tests with a custom gomock matcher

  12:01

- P2019- Why PASETO is better than JWT for token-based authentication

  15:25

- P2120- How to create and verify JWT - PASETO token in Golang

  23:31

- P2221- Implement login user API that returns PASETO or JWT access token in Go

  13:52

- P2322- Implement authentication middleware and authorization rules in Golang using

  29:18

- P2423- Build a minimal Golang Docker image with a multistage Dockerfile

  12:00

- P2524- How to use docker network to connect 2 stand-alone containers

  10:08

- P2625- How to write docker-compose file and control service start-up orders with w

  16:09

- P2726- How to create a free tier AWS account

  06:45

- P2827- Auto build - push docker image to AWS ECR with Github Actions

  20:45

- P2928- How to create a production DB on AWS RDS

  10:09

- P3029- Store - retrieve production secrets with AWS secrets manager

  23:32

- P3130- Kubernetes architecture - How to create an EKS cluster on AWS

  17:28

- P3231- How to use kubectl - k9s to connect to a kubernetes cluster on AWS EKS

  15:04

- P3332- How to deploy a web app to Kubernetes cluster on AWS EKS

  20:54

- P3433- Register a domain - set up A-record using Route53

  10:31

- P3534- How to use Ingress to route traffics to different services in Kubernetes

  09:51

- P3635- Automatic issue TLS certificates in Kubernetes with Let-s Encrypt

  14:26

- P3736- Automatic deploy to Kubernetes with Github Action

  14:39

- P3837- How to manage user session with refresh token - Golang【Advanced Backend Top】

  22:56