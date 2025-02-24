## Postgres

## Golang Migrate

Need to install Golang Migrate CLI tool for CLI usage if not already installed: `brew install golang-migrate`

To create a migration file, run `migrate create -ext sql -dir migrations <migration_name>` from within the FreeCreate directory.

`-ext` flag is "extension": `sql` is passed as an argument to give the migration files a `.sql` extension.

`-dir` flag stands for director: argument `migrations` is passed in to place migration files in migration directory.

For a failed migration, the database will be in a `dirty` state. Check the database configuration to see if there are not errors. If there are errors, fix them. Once fixed, or if there are no errors, update the schema_migrations table for the migration in question to `f` in the `dirty` column, indicating that the migration is not dirty. Then, rollback the migration in question, fix whatever caused the problem in the up migration, and re-run the migration.
