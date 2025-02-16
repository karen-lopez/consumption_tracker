# Consumption Tracker

![Go](https://img.shields.io/badge/Go-1.20+-00ADD8?logo=go)
![Gin](https://img.shields.io/badge/Gin-1.9.0-00ADD8?logo=go)
![Swagger](https://img.shields.io/badge/Swagger-2.0-85EA2D?logo=swagger)
![Docker](https://img.shields.io/badge/Docker-24.0+-2496ED?logo=docker)

**Consumption Tracker** is a backend application developed in Go that allows tracking and managing electrical energy consumption. It provides a RESTful API to query energy consumption data, such as active energy, reactive energy, capacitive reactive, and solar energy, for a specific meter within a date range.

---

## Key Features

- **Energy consumption query**: Retrieve energy consumption data for a specific meter within a date range.
- **Period filtering**: Supports daily, weekly, and monthly queries.
- **Swagger integration**: Automatic API documentation using Swagger.
- **Database connection**: Store and retrieve data from a PostgreSQL database.
- **Address client**: Integration with an external service to fetch the address associated with a meter.

---

## Prerequisites

Before running the project, ensure you have the following installed:

- **Go** (version 1.20 or higher)
- **PostgreSQL** (for the database)
- **Git** (for cloning the repository)
- **Docker** and **Docker Compose** (optional, for running the project in containers)

---

## Installation

### Option 1: Run Locally

1. Clone the repository:

   ```bash
   git clone https://github.com/karen-lopez/consumption_tracker.git
   cd consumption_tracker
   ```

2. Set up the database:

    - Create a database in PostgreSQL.
    - Configure environment variables in a `.env` file or directly in your system:

      ```bash
      export DB_HOST=localhost
      export DB_PORT=5432
      export DB_USER=your_user
      export DB_PASSWORD=your_password
      export DB_NAME=consumption_tracker
      ```

3. Install dependencies:

   ```bash
   go mod download
   ```

4. Start the server:

   ```bash
   go run cmd/app/main.go
   ```

### Option 2: Run with Docker Compose

1. Clone the repository:

   ```bash
   git clone https://github.com/karen-lopez/consumption_tracker.git
   cd consumption_tracker
   ```

2. Create a `.env` file in the root directory with the following environment variables:

   ```bash
   POSTGRES_USER=your_user
   POSTGRES_PASSWORD=your_password
   POSTGRES_DB=consumption_tracker
   DB_HOST=db
   DB_PORT=5432
   DB_USER=your_user
   DB_PASSWORD=your_password
   DB_NAME=consumption_tracker
   ```

3. Start the application and database using Docker Compose:

   ```bash
   docker-compose up --build
   ```

   This will:
    - Build the Go application container.
    - Start a PostgreSQL database container.
    - Initialize the database with the schema and data from `init.sql` and `consumptions.csv`.

4. Access the application:

    - The API will be available at `http://localhost:8080`.
    - The Swagger documentation will be available at `http://localhost:8080/swagger/index.html`.

---

## Usage

### API Endpoints

The API provides the following endpoints:

- **GET /consumption**: Retrieves energy consumption data for a specific meter within a date range.

  Parameters:
    - `meter_id`: ID of the meter.
    - `start_date`: Start date in `YYYY-MM-DD` format.
    - `end_date`: End date in `YYYY-MM-DD` format.
    - `kind_period`: Period type (`daily`, `weekly`, `monthly`).

  Example request:

  ```bash
  GET /consumption?meter_id=1&start_date=2023-06-01&end_date=2023-06-30&kind_period=monthly
  ```

### API Documentation

The API documentation is available in Swagger format. To access it:

1. Start the server.
2. Open your browser and visit:

   ```
   http://localhost:8080/swagger/index.html
   ```

---

## Project Structure

The project is organized as follows:

```
.
├── cmd/                          // Application entry point
│   └── app/
│       └── main.go             // Main function
├── config/                       // Application configuration
│   ├── config.go               // Configuration settings
│   └── env.go                  // Environment variable loader
├── internal/                     // Internal application logic
│   ├── core/                   // Domain layer
│   │   ├── domain/             // Domain entities
│   │   │   └── energy_consumption.go  // Energy consumption entity (struct)
│   │   └── ports/              // Interface definitions (ports)
│   │       ├── repository.go   // Repository interface
│   │       └── address_service.go  // Address service interface
│   ├── application/            // Use cases and business logic
│   │   └── services/
│   │       └── energy_service.go  // Energy service implementation (injects repository & address service)
│   ├── infrastructure/         // External adapters and implementations
│   │   ├── database/
│   │   │   ├── postgres/
│   │   │   │   ├── repository_impl.go  // SQL repository implementation
│   │   │   │   └── migrations/   // Database migration scripts
│   │   └── http/
│   │       └── address_client.go  // HTTP client for the address microservice
│   └── interfaces/             // Controllers, API handlers, and CLI commands
│       └── http/
│           └── handlers/
│               └── energy_handler.go  // API endpoints
├── pkg/                        // Reusable libraries and utilities
│   └── utils/
│     
├── tests/                      // Test suites
│   ├── unit/                 // Unit tests
│   │   ├── core/             // Domain tests
│   │   └── application/      // Business logic tests
│   └── integration/          // Integration tests
│       ├── database/         // Database integration tests
│       └── http/             // HTTP integration tests
├── docs/                     // Project documentation
│   ├── swagger/              // Swagger API specification
│   │   ├── swagger.yaml      // Swagger YAML file
│   │   └── swagger.json      // Generated Swagger JSON file
│   ├── architecture/         // Architecture documentation
│       ├── README.md         // Architectural overview
│       └── diagram           // Architecture diagram
├── Dockerfile                // Docker configuration file
├── docker-compose.yml        // Docker Compose configuration (e.g., for local PostgreSQL)
├── go.mod                    // Go module dependencies
├── .env                      // environment variables file
└── README.md                 // Main project documentation

```

---

## Contact

If you have any questions or suggestions, feel free to reach out:

- **Name**: Karen Lopez
- **Email**: stefanny0214@gmail.com
- **GitHub**: [karen-lopez](https://github.com/karen-lopez)

---

Thank you for using **Consumption Tracker**! 🚀
