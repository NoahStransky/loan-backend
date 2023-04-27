## loan application backend project

this project should be placed in `$GOPATH/src/bernie.top/demyst` directory.

## Prerequisites

---

- [Golang](https://go.dev/) (min version: 1.19)
- [Docker](https://www.docker.com/get-started) (min allocated memory: 6 GB)
- [Docker Compose](https://docs.docker.com/compose/install) (min version: 1.29.2)

## prepare environment

```bash

go mod init
go mod tidy

```

## development

### run command

```bash
docker-compose up
```

### rerun backend project

```bash
docker-compose up --build backend
```

### init database

```bash
./migration.sh
```

### run test

```bash
go test ./... -coverprofile .testCoverage.txt

```
