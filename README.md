 # Go Calorie Tracker

This is a simple Go application that allows users to track their daily calorie intake. The application uses a PostgreSQL database to store the data. The application is built using the Gin framework.

## Prerequisites

* PostgreSQL database
* Go 1.21 or later
* Docker

## Installation

1. Clone the repository

```
git clone https://github.com/Sadeedpv/go-calorie-tracker.git
```

2. Install dependencies

```
go mod download
```

3. Create a .env file and add your database credentials

```
DB_URL=postgres://user:password@localhost:5432/calorie_tracker
```

4. Run the application

```
go run main.go
```

## Usage

The application exposes the following endpoints:

* `/v1/calories`: Get all calories
* `/v1/calories`: Add a calorie
* `/v1/calories`: Delete all calories
* `/v1/calories/:id`: Get a calorie by ID
* `/v1/calories/:id`: Update a calorie by ID
* `/v1/totalcalories`: Get the total number of calories

## Docker

You can also run the application using Docker.

1. Build the image

```
docker build -t go-calorie-tracker .
```

2. Run the container

```
docker run -p 8000:8000 go-calorie-tracker
```

## Contributing

Contributions are welcome! Please open a pull request if you have any suggestions or improvements.