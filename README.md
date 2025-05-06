# blog-aggregator
A blog aggregator in go that uses a PostgreSQL database.

# Installation

## Setup
- go version 1.23.2 or higher
- PostgreSQL
- install blog-aggregator

### PostgreSQL Setup
<details>
<summary>Installation process for PostgreSQL</summary>
1. Install PostgresSQL v15 or later
Mac:
`brew install postgresql@15`

Linux: 
```bash
sudo apt update
sudo apt install postgresql postgresql-contrib
```

2. Ensure the installation worked:
```bash
psql --version
```
</details>

3. Start the Postgres server in the background:
- Mac: `brew services start postgresql@15`
- Linux: `sudo service postgresql start`

4. Connect to postgres (e.g. using psql) and create a new database:
```SQL
CREATE DATABASE gator;
```

5. Connect to the database:
```
\c gator
```

6. Set the user password (Linux only):
```SQL
ALTER USER postgres PASSWORD 'postgres';
```

7. Exit psql using the `exit` command in psql

### Install blog-aggregator
Run the following command in your terminal:
```bash
go install github.com/Snansidansi/blog-aggregator@latest
```

# Usage
Coming soon.

# Development Setup (Linux)
- install goose:
```bash
go install github.com/pressly/goose/v3/cmd/goose@latest
```
