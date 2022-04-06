# todong
todong (todo next-gen) is a small HTTP back-end server for managing to-do lists. It is my personal project, with the goal to practice writing production-level code. The main goal of the project is to be a very flexible to-do list back-end.

## Features

- **Configurable web frameworks** - choose between [gorilla/mux](https://github.com/gorilla/mux), [Gin](https://github.com/gin-gonic/gin) or [Fiber](https://github.com/gofiber/fiber), easily with one line in the YAML config file! (e.g. `server: fiber`).

- **Configurable data store**, choose between Postgres (via [Gorm](https://gorm.io)) or [Redis](https://redis.io), easily with one line in the YAML config file! (e.g. `store: redis`). For more info, see [package store](/store/)

- JWT authentication - users can only view and manage their own to-do lists. JWT is issued during logins, while the authentication middleware is in `lib/middleware/authenticate.go`. Both Gin and Fiber handlers utilize their own *but compatible* authorization middleware.

- YAML configuration. For more info, see [package config](/config/)

## Important interfaces
### `httpserver.Server`
`Server` interface abstracts the HTTP servers with `SetUpRoutes()` and `Serve(addr string) error`. `ginserver.ginServer` and `fiberserver.fiberServer` implement `httpserver.Server`.

In `main()`, it calls `httpserver.New()` which takes in a `handler.Adapter`.

### `handler.Adapter`
`Adapter` helps when a `httpserver.Server` is registering handler functions. The type that implment this interface should be able to return correct handler types for each `httpserver.Server`.

### `store.DataGateway`
`DataGateway` abstracts high-level storage, like `GetUserByUuid()` or `DeleteTodo()`.  There are 2 implementations of this interface, by `gormDataGateway` and `redisDataGateway`.
