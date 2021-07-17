# sugar

## Setup and run

1. Start a postgres server.
2. Create `/etc/sugar/config/config.toml`. You can edit and use the [sample config](build/config.toml).
3. Run `go build -o sugar`
4. Run `./sugar`

The default credentials is `root:root`

## List of APIs

#### POST /auth/login

#### POST /auth/setpassword

#### POST /door/action

#### POST /door/list

### Admin APIS

#### POST /admin/door/add

#### POST /admin/door/remove

#### POST /admin/user/add

#### POST /admin/user/remove

#### POST /admin/permission/add

#### POST /admin/permission/remove

> Check out these [examples](misc/rest/api.rest)

## File structure

### store

database connections and other application states

### routes

API URL routes and middlewares

### migrations

database table definitions

### pkg

packages like auth,doors,users with their api handlers and db functions

### misc

sample configs, service files, scripts etc
