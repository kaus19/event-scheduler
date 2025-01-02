# event-scheduler

A RESTful API in Golang for suggesting time slots for events. The API allows users to 

## Table of Contents
- [Features](#features)
- [Installation](#installation)
- [Usage](#usage)
- [Running Tests](#running-tests)
- [Technologies Used](#technologies-used)

## Features
- Supports creating, reading, updating, and deleting events
- Supports creating, reading, updating, and deleting preferred time slots by each user.
- Endpoint that shows the possible time slots for the event.

## Installation

### 1. Clone the repository:

```sh
git clone https://github.com/laus19/event-scheduler.git

cd event-scheduler
```

### 2. Start the server:

```sh
docker compose up --build
```

### 3. Shut down the server:

```sh
docker compose down
```

## Usage

After starting the server, access the API at `http://localhost:8080`.

## Running tests

Tests are automatically run using github actions on pushed directly to main branch pull request is created on main branch

### Running tests in local

```sh
# Run test
make test
```

## Technologies Used

- go
- postgreSQL: SQL engine
- gin framework: HTTP server
- sqlc: Generating go code for the sql queries
- make: Important commands documentation
- migrate: Database setup/migration utility
- docker: Containerization
- openAPI: API documentation