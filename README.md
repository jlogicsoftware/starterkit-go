# Starter Kit in Go

## API

The application runs as an HTTP server at port 3000. It provides the following RESTful endpoints:

- POST /auth/signin - accepts username/passwords and returns jwt token and refresh token
- POST /auth/signup - accepts username/passwords and returns jwt token and refresh token
- POST /auth/signout - accepts jwt token and deletes token
- GET /auth/:token -  refreshes sessions and returns jwt token
- GET /auth/me - returns info about currently logged in user

- GET /swagger/ (with trailing slash) - launches swagger UI in browser

- GET /v1/users - returns list of users
- GET /v1/users/:id - returns single user
- POST /v1/users - creates a new user
- PATCH /v1/users/:id - changes user's data
- DELETE /v1/users/:id - deletes a user

## Development

Install Docker images:

```sh
make docker.build
```

- Application is running at [http://localhost:3000](http://localhost:3000)
- Mailhog mail server accessible at [http://localhost:1025/](http://localhost:1025/)
- Postgres Adminer is available at [http://localhost:9000/](http://localhost:9000/)

Adminer credentials:

- system: PostgreSQL
- server: postgres
- username: postgres
- password: password
- database: starterkit

Run the application in a watch mode. Ctl-C to stop.

```sh
make watch
```

For code validation, install dev dependencies first and then run only "check" command:

```sh
make dev
make check
```
