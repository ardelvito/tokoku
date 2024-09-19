# E-Commerce Application

This is a monolithic e-commerce application built with Golang for the backend and React for the frontend. It includes features for users to register, login, manage products, and manage shopping carts. The backend handles all business logic, including user authentication, product management, and cart operations, and connects to a PostgreSQL database.

## Features

- User authentication (register, login)
- Product CRUD (Create, Read, Update, Delete)
- Store management (register store, view sales)

## Tech Stack

- **Backend**: Golang
- **Frontend**: React
- **Database**: PostgreSQL
- **API Format**: JSON (RESTful)

## Prerequisites

Before setting up the project, make sure you have the following installed:

- [Go](https://golang.org/doc/install) (version >= 1.18)
- [Node.js](https://nodejs.org/) (for React frontend)
- [PostgreSQL](https://www.postgresql.org/download/)
- [Docker](https://docs.docker.com/get-docker/) (optional for containerization)

## API

- [Postman](https://documenter.getpostman.com/view/25070708/2sAXqs7NYf)

## Setup and Installation

### 1. Clone the repository

```bash
git clone https://github.com/ardelvito/tokoku.git
cd your-tokoku
```

### 2. Backend Setup

#### 2.1 Environment Variables

Create a `.env` file in the root directory of the backend and add the following environment variables:

```
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=root
DB_NAME=tokokudb
JWT_SECRET=ASDQWE123456
```

#### 2.2 Database Setup

Make sure PostgreSQL is running and create a new database:

```bash
psql -U postgres
CREATE DATABASE tokokudb;
```

#### 2.3 Install Go Dependencies

Navigate to the backend folder and install Go dependencies:

```bash
go mod tidy
```

#### 2.4 Run the Backend Server

```bash
go run main.go
```

The server will run on `http://localhost:8080`.

### 3. Frontend Setup

#### 3.1 Install Node.js Dependencies

Navigate to the `frontend` folder and install React dependencies:

```bash
cd frontend
npm install
```

#### 3.2 Run the React Development Server

```bash
npm start
```

The frontend will run on `http://localhost:3000`.

## Environment Variables

The following environment variables are required to configure the application:

| Variable      | Description                                                |
| ------------- | ---------------------------------------------------------- |
| `DB_HOST`     | Hostname for the PostgreSQL server (e.g., `localhost`)     |
| `DB_PORT`     | Port number for PostgreSQL (default: `5432`)               |
| `DB_USER`     | Username for the PostgreSQL database (default: `postgres`) |
| `DB_PASSWORD` | Password for the PostgreSQL database                       |
| `DB_NAME`     | Name of the PostgreSQL database (e.g., `tokokudb`)         |
| `JWT_SECRET`  | Secret key for signing JWT tokens                          |
