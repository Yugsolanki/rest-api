# REST API Event Management System

A simple RESTful API for event management, built with Go, Gin, and SQLite. This project allows users to sign up, log in, create events, register for events, and manage event registrations. JWT authentication is used for securing endpoints.

## Features

- User registration and authentication (JWT)
- Create, update, delete, and list events
- Register and cancel registration for events
- SQLite database for persistence
- Passwords hashed with bcrypt
- Clean modular code structure

## Project Structure

```
.
├── app.go                # Main entry point
├── db/                   # Database initialization and migrations
├── models/               # Data models (User, Event)
├── routes/               # API route handlers
├── middlewares/          # Authentication middleware
├── utils/                # Utility functions (hashing, JWT)
├── api-test/             # HTTP request samples for testing
├── go.mod, go.sum        # Go modules
└── api.db                # SQLite database file (created at runtime)
```

## Getting Started

### Prerequisites

- Go 1.18+
- [Git](https://git-scm.com/)
- [SQLite3](https://www.sqlite.org/index.html) (optional, for inspecting the database)

### Installation

1. **Clone the repository**
   ```sh
   git clone https://github.com/yourusername/rest-api-event-management.git
   cd rest-api-event-management
   ```

2. **Install dependencies**
   ```sh
   go mod tidy
   ```

3. **Run the application**
   ```sh
   go run app.go
   ```
   The server will start on `http://localhost:8080`.

## API Endpoints

### Authentication

- `POST /signup` — Register a new user
- `POST /login` — Log in and receive a JWT token

### Events

- `GET /events` — List all events
- `GET /events/:id` — Get a single event
- `POST /events` — Create a new event (authenticated)
- `PUT /events/:id` — Update an event (authenticated, owner only)
- `DELETE /events/:id` — Delete an event (authenticated, owner only)

### Event Registration

- `POST /events/:id/register` — Register for an event (authenticated)
- `DELETE /events/:id/register` — Cancel registration (authenticated)

## Testing the API

Sample HTTP requests are provided in the api-test directory. You can use [VS Code REST Client](https://marketplace.visualstudio.com/items?itemName=humao.rest-client) or [Postman](https://www.postman.com/) to test the endpoints.

## Example: Create Event

```http
POST http://localhost:8080/events
Content-Type: application/json
Authorization: <JWT_TOKEN>

{
  "name": "Sample Event",
  "description": "This is a sample event.",
  "dateTime": "2023-10-01T10:00:00Z",
  "location": "Sample Location"
}
```

## Security

- Passwords are hashed using bcrypt before storing in the database (`utils/hash.go`).
- JWT tokens are used for authentication (`utils/jwt.go`).
- Protected routes require the `Authorization` header with a valid JWT.

---

**Made with Go, Gin, and ❤️**