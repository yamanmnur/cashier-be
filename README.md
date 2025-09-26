# Super Bank Assesment

Simple Dashboard using Go & Next Js for Super Bank Assesment

## Specification

### Backend

- **Go** (version 1.24.1 or higher)
- **Go Echo**
- **PostgreSQL** (version 15 or latest)
- **Gorm** (version 1.25 or latest)
- **Docker** and **Docker Compose**

## Install & Running - Using Docker Compose

```bash
  docker-compose up --build -d
```

### Local address Backend

```bash
  http://localhost:3001
```

### Credential Postgre SQL

use this credential if use Postgre from docker-compose

```bash
  host     : localhost
  port     : 5432
  username : postgres
  password : password
  database : cashierdb
```

## Installation - backend

1. Update the .env
2. Run Docker Compose

or

1. Update the .env
2. Run DB Host in Docker Compose or another
3. go run server.go

## Postman Collection

```
cashier_be.postman_collection.json
```
