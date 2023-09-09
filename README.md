## Description
xm - An application to handle companies. It should provide the following operations:
    • Create
    • Patch
    • Delete
    • Get (one)

## Installation

### Local Setup
- Clone the Repo -git clone git@bitbucket.org:keyvaluesoftwaresystems/go-gin-boilerplate.git
- Install Golang: https://golang.org/doc/install
- Install precommit hooks for github. It is required for prehook commit checks like lint, unit testing, formatting and so on.

```bash
    make prehook-install
```
- Install Swagger required for API Documentation: 
```bash 
    make swagger-install 
```
- Generate Swagger API Specs:
```bash
    make gen-swagger
```
- Download and Install all the dependent packages
```bash
    go mod download
```
- To Build the Go Binary:
```bash
    make build
```
- To Run the Server:
```bash
    make run
```
- To Run Tests:
```bash
    make test
```

#### Environment Variables

To run the service, you will need to populate environment variables to  .env file

#### To create the migration File

Install goose binary from https://github.com/pressly/goose and move it under one of your PATH folder
Or you can download using go

`go install github.com/pressly/goose/v3/internal/goose@latest`

To Create a Migration File

`goose -dir <migration_folder> create <name> sql`

To Execute a Migration

`goose <driver> "postgresql://username:password@host:port/database?sslmode=disable" up`

  
## Tech Stack

*Database:* Postgres
*Server:* Golang
