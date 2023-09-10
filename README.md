## Description
XM - An application to handle companies. It should provide the following operations:
    • Create
    • Patch
    • Delete
    • Get (one)

## Installation
Install Golang => https://golang.org/doc/install
Install Docker => sudo snap install docker

### Local Setup
- Clone the Repo -git clone https://github.com/kannanmohan1994/xm-demo.git
- Setup postgres, pgadmin, kafka: Enter repo, RUN docker-compose up 
- Run application in local => make tidy, make run 

### Useful Commands
- Download and Install all the dependent packages
```bash
    make tidy
```
- To Run the Server:
```bash
    make run
```
- To Run Tests:
```bash
    make test
```
- To Run Lint:
```bash
    make lint
```
- To Setup Swagger Docs:
```bash
    make swagger
```

#### Environment Variables

To run the service, you will need to populate environment variables to  .env file

#### Run migrations
    Up migrations => make migrate-up
    Down migrations => make migrate-down

## Tech Stack

*Database:* Postgres
*Server:* Golang
