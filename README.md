# CMS

* Manage the database schema and apply migrations in `pkg/storage/db/migrations`.
* Manage the queries in `pkg/storage/db/sqlc`.
* Generated query directory `pkg/storage/db/dbx`.
* Manage the templates in `pkg/ui`.

## Go Tasks
install go [task](https://taskfile.dev) package to run commands defined in [Taskfile.yml](Taskfile.yml):
```bash
go install github.com/go-task/task/v3/cmd/task@latest
```
then install all dev tooling:
```bash
task setup:tooling
```
list all tasks:
```bash
task --list-all
```

## Config
For configuration see the `config.toml` passed in as the `--config` flag to app.

create your dev config:
```bash
task setup:config
```

## SQL
make sure to use the correct db dsn in `sqlc.yml` and that the db is fully migrated.

generate go code from sql:
```bash
task gen:sqlc
```

## Migrations

new:
```bash
task migrate:add SEQ="migration_name"
```

up:
```bash
task migrate:up
```

down:
```bash
task migrate:down
```

## Templates

Generate template code with [templ.guide](https://templ.guide)
```bash
task gen:templ
```

## Usage

use air to generate the templates and run the server:
```bash
task air
```

list app commands:
```bash
go run . help
```

run the server:
```bash
go run . server --config config.dev.toml
```

## Tailwind

For simplicity we are using the [standalone cli](https://tailwindcss.com/blog/standalone-cli).

download the cli to the tmp folder (created once air is run):
```bash 
curl -sLO https://github.com/tailwindlabs/tailwindcss/releases/latest/download/tailwindcss-macos-arm64
chmod +x tailwindcss-macos-arm64
mv tailwindcss-macos-arm64 tmp/tailwindcss
```

build and watch:
```bash
task gen:tailwind
```

When content is in the database you will need a safelist of css classes.
To do this there is a query in `pkg/storage/db/sqlc/tailwind.sql`, this returns a unique list
of css classes found in the `page_html.html` field. You can expand this to fit further requirements.

running the following task will rebuild the `tailwind.safelist.txt` file with an updated list of css classes.
```bash
task gen:tailwind-safelist
```

to expand the config make sure you update the `tailwind.config.js.tmpl` file to include new settings.
With the config rebuilt, re-run the main `gen:tailwind` task.

## Docker
This can be used to build the app as a binary and run it in a container.

build:
```bash
docker build --build-arg EXPOSE_PORT=80 -t cms:latest .
```

run:
```bash
docker run \
--rm \
--name cms \
--publish "80:80" \
--env "DATABASE_HOST=host.docker.internal" \
--env "SERVER_DEBUG=false" \
--env "SERVER_PORT=80" \
cms:latest \
server --config config.toml
```

if you need it create a postgres container for the database:
```bash
docker run \
--detach \
--name "cms-postgres" \
--mount type=tmpfs,destination=/var/lib/postgresql/data \
--publish "5432:5432" \
--env POSTGRES_USER=postgres \
--env POSTGRES_PASSWORD=password \
--env POSTGRES_DB=cms \
postgres:latest
```

## Links

* [Boilerplate](https://github.com/stuartaccent/echo-boilerplate)
* [Golang](https://go.dev)
* [Task](https://taskfile.dev)
* [Templ](https://templ.guide)
* [Air](https://github.com/air-verse/air)
* [Icônes](https://icones.js.org/collection/lucide)
* [Tailwind](https://tailwindcss.com)
* [Owl](https://github.com/AccentDesign/owl)
* [Shadcn](https://ui.shadcn.com/docs)
* [Htmx](https://htmx.org)
* [Writing secure Go code](https://jarosz.dev/article/writing-secure-go-code)