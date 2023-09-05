 # Calorie Tracker API 

This is a simple calorie tracker API built using Gin framework in Go. The API allows users to track their daily calorie intake and provides a simple interface for users to interact with the data.

## Prerequisites

- Go 1.18 or later
- PostgreSQL database
- A text editor
- Git

## Setup

1. Clone the repository

```sh
git clone https://github.com/Sadeedpv/go-calorie-tracker.git
cd go-calorie-tracker
```

2. Install dependencies

```sh
go mod tidy
```

3. Create a `.env` file in the project directory and add the environment variables


4. Create a PostgreSQL database and table

```sql
CREATE DATABASE calorie_tracker;

\c calorie_tracker

CREATE TABLE calories (
  id SERIAL PRIMARY KEY,
  food TEXT NOT NULL,
  calorie INTEGER NOT NULL
);
```

5. Run the application

```sh
go run main.go
```

## Usage

The API provides the following endpoints:

- `/v1/calories`: Get all calories
- `/v1/calories`: Add a calorie
- `/v1/calories/:id`: Get a calorie by ID
- `/v1/calories/:id`: Update a calorie by ID
- `/v1/calories/:id`: Delete a calorie by ID
- `/v1/totalcalories`: Get total calories consumed

## Contributing

Contributions are welcome! Please feel free to fork the repository and submit a pull request.