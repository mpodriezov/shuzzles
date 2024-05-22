# SHUZZLES

A simple pet project website, where puzzle or board games lovers can share their sets.

##Setup

1. Create .env file with the following keys

- PORT=
- DATABASE_URL= <torso sql connection string>, documentation https://turso.tech/

2. Migrations, documentation - https://pkg.go.dev/github.com/gavsidhu/miflo#section-readme,
   To create new migration use: `miflo create [migration_file_name]`

3. To run the web server just do `make run`
