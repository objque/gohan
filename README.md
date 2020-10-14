# Intro

Recently I've created a few of public/private projects with same dir structure based on copy and paste from other older projects. This took a time.

But it's might be simplified, if i will have prepared public project with common needed packages.

# Aims

## database

- App should be prepared for work with [PostgreSQL](https://www.postgresql.org/docs/12/index.html)
- Driver should be [jackc/pgx](https://github.com/jackc/pgx) or [lib/pq](https://github.com/lib/pq) _(but package is effectively in maintenance mode and is not actively developed)_
- To exec raw queries may be used [sqlx](https://github.com/jmoiron/sqlx)
- Sometimes may be required to build query with dynamic conditions. Silver bullet is [Masterminds/squirrel](https://github.com/Masterminds/squirrel) that compatible with [sqlx](https://github.com/jmoiron/sqlx)
- It is a good practice to keep raw sql migrations for the db near code. There are many libraries, but most popular are [rubenv/sql-migrate](https://github.com/rubenv/sql-migrate) and [golang-migrate](https://github.com/golang-migrate/migrate)
- Also do not forget about tests for models: tests should be run on real postgres _(e.g in docker)_
- To simplify process need to use [testify/assert](https://github.com/stretchr/testify/assert)/[testify/require](https://github.com/stretchr/testify/require). Before start you definitely must check great [article](https://dev.to/techschoolguru/write-go-unit-tests-for-db-crud-with-random-data-53no)
- Another good practices: use context in db methods; use mgr that compatible with query and tx _(like in the gorm)_
- Optional: it's ok to use [kyleconroy/sqlc](https://github.com/kyleconroy/sqlc) for generate type safe Go from SQL