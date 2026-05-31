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