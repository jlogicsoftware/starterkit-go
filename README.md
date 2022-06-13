# Starter Kit in Go

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
