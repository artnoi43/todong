# todong
todong (todo next-gen) is a small HTTP back-end server for managing to-do lists. It is my personal project, with the goal to practice writing production-level code. The main goal of the project is to be a very flexible to-do list back-end.

## Features

- **Configurable web frameworks** - choose between standard library [`net/http`](https://pkg.go.dev/net/http),  [`gorilla/mux`](https://github.com/gorilla/mux), [Gin](https://github.com/gin-gonic/gin) or [Fiber](https://github.com/gofiber/fiber), easily with one line in the YAML config file! (e.g. `server: fiber`). See [package `/lib/enums`](/lib/enums) for available server types.

- **Configurable data store**, choose between Postgres (via [Gorm](https://gorm.io)) or [Redis](https://redis.io), easily with one line in the YAML config file! (e.g. `store: redis`). For more info, see [package `store`](/data/store/)

- JWT authentication - users can only view and manage their own to-do lists. JWT is issued during logins, while the authentication middleware is in [package `middleware`](/lib/middleware). Both Gin and Fiber handlers utilize their own *but compatible* authorization middleware. On the other hand, `net/http` and `gorilla/mux` share the same handlers ([package `httphandler`](/domain/usecase/handler/httphandler))

- YAML configuration. For more info, see [package config](/config/)

## Important interfaces
### `httpserver.Server`
`Server` interface abstracts the HTTP servers with `SetUpRoutes()` and `Serve(addr string) error`. `ginserver.ginServer` and `fiberserver.fiberServer` implement `httpserver.Server`.

In `main()`, it calls `httpserver.New()` which takes in a `handler.Adapter`.

### `handler.Adapter`
`Adapter` helps when a `httpserver.Server` is registering handler functions. The type that implment this interface should be able to return correct handler types for each `httpserver.Server`.

### `store.DataGateway`
`DataGateway` abstracts high-level storage, like `GetUserByUuid()` or `DeleteTodo()`.  There are 2 implementations of this interface, by `gormDataGateway` and `redisDataGateway`.
