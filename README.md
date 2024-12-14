# Auth Provider API

## Project Overview
This project is a RESTful API for a program that includes user authentication and authorization. It provides secure user login, registration, email verification, and CRUD operations for user data. The application is built with Go, using the Fiber web framework, and follows clean architecture principles with separate layers for handlers, services, repositories, and middleware.

## Features

### Authentication Endpoints
- **POST `/api/auth/login`**: Authenticates a user and starts a session.
- **POST `/api/auth/logout`**: Logs out a user and invalidates the session.
- **GET `/api/auth/verify-email`**: Verifies a user's email address using a token sent via email.
- **POST `/api/auth/register`**: Registers a new user into the system.

### User Management Endpoints
- **GET `/api/users`**: Lists all users.
- **GET `/api/users/me`**: Retrieves the currently logged-in user's details.
- **GET `/api/users/:id`**: Fetches details of a specific user by ID.
- **PATCH `/api/users`**: Updates the currently logged-in user's details.
- **DELETE `/api/users/:id`**: Deletes a user by ID.

### Middleware
- **JWT Middleware**: Ensures secure access by validating JWT tokens and user sessions.

## Setup Instructions

### Prerequisites
- Go 1.20+
- PostgreSQL or compatible database
- Fiber web framework dependencies

### Installation
1. Clone the repository:
   ```bash
   git clone git@github.com:Xurliman/auth-service.git
   cd auth-service
   ```
2. Install dependencies:
   ```bash
   make install
   ```
3. Set up environment variables in a `.config.yaml` file and rewrite with corresponding credentials:
   ```dotenv
   cp config.yml.example config.yaml
   ```
4. Run database migrations:
   ```bash
   make migrate-up
   ```
5. Optionally run with seeding the database:
   ```bash
   make seed
   ```
6. Start the application:
   ```bash
   make run
   ```

### Usage
The API listens on the port specified in the `.config.yaml` file (default: `8080`). Use tools like Postman or curl to interact with the endpoints.

Example request to log in, but first you need to register and verify the email:
```bash
curl -X POST http://localhost:8080/api/auth/login \
-H "Content-Type: application/json" \
-d '{"email": "user@example.com", "password": "password123"}'
```

## Code Structure

### Main Application Flow
The `main.go` file initializes the application:
1. **Configuration Setup**: Loads environment variables.
2. **Logger Initialization**: Sets up structured logging.
3. **Database Setup**: Connects to the database.
4. **Route Setup**: Configures API routes.
5. **Server Startup**: Starts the Fiber HTTP server.

### Routing
The `routes.Setup` function defines the API endpoints and attaches middleware for authentication and error handling.

### Middleware
- **JWT Middleware**: Validates session cookies and JWT tokens for secured endpoints.
- **Custom Claims**: Extends JWT claims to include user-specific data.

### Seeders
Run `make seed` to populate the database with initial data for testing or development purposes.

## Contributing
Contributions are welcome! Please follow the [contributor guidelines](CONTRIBUTING.md) to submit a pull request or report issues.

---

Feel free to reach out for support or feedback. Enjoy building with the Auth Provider API!

