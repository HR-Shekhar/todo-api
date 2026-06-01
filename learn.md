`migrate create -ext sql -dir migrations -seq create_users_table`

migrations/
├── 000001_create_users_table.up.sql
└── 000001_create_users_table.down.sql

create           -> create migration files
-ext sql         -> SQL migration
-dir migrations  -> put files in migrations/
-seq             -> 000001, 000002, ...


### Run migrations for docker container(postgres)

`migrate -path migrations -database "postgres://postgres:secretkey@localhost:5431/todo_api?sslmode=disable" up`

## Verify

`docker exec -it todo_postgres psql -U postgres -d todo_api`


### Postgres, writing query in repo
1. Exec()

Use when you don't expect rows back.

Example:-

`DELETE FROM users WHERE id = ...`

or

`UPDATE users SET ...`

You only care:

Did it succeed?
How many rows affected?

No row data returned.

2. QueryRow()

Use when you expect exactly:

one row

back.

Example:

`SELECT * FROM users WHERE email = ...`

One user.

Or:

`INSERT ... RETURNING id`

One inserted row.

3. Query()

Use when you expect:

many rows

Example:

`SELECT * FROM todos`

which may return:

0 rows
1 row
100 rows

You then iterate through them.