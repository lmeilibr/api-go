# Rocketseat Challenge - API RESTful

## Project Structure

- controller/user.go: Handles user-related API requests.
- models/user.go: Defines the User model and its methods.
- routes/routes.go: Registers routes for user endpoints
- db/db.go: Contains database connection logic.
- utils/http_utils.go: Contains utility functions for HTTP responses.
- main.go: Start the Chi app.

## API

| Method | URL                | Description                     |
|--------|--------------------|---------------------------------|
| POST   | /api/users         | Create a new user using the informations sent on the request body |
| GET    | /api/users         | Return all users from the database |
| GET    | /api/users/:id     | Return a user by ID from the database |
| DELETE | /api/users/:id     | Delete a user by ID from the database |
| PUT    | /api/users/:id     | Update a user by ID from the database. Returns the modified user |

## Some comments about the project
- as the DB is in memory, it will be reset every time the application is restarted.
- it's hard to force a DB error, so the functions that return an error will always return nil, but the error handling is in place.
- added more validation to the user model, so it will not accept fields outside the len specified.

## How to run the project
``` go run main.go ```

It will start a server on `http://localhost:8080`, and you can test the API using Postman or any other API client.