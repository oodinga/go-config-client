[![Go Reference](https://pkg.go.dev/badge/github.com/oodinga/goconfig@v0.4.3.svg)](https://pkg.go.dev/github.com/oodinga/goconfig@v0.4.3)

# GoConfig

**GoConfig** is a Go package that enables developers to seamlessly load configurations from Spring Boot-style config servers. It simplifies the process of managing environment-specific configurations for Go applications.

---

## Installation

Install the package using `go get`:

```bash
go get github.com/oodinga/goconfig
```

---
## Getting Started

### Required Environment Variables

GoConfig requires the following environment variables to be set for proper functionality:

- `app.name` – Name of your application.
- `app.config.profiles.active` – Comma separated list of active profiles (e.g., `dev`, `prod`, etc.).
- `app.config.server.url` – URL of the configuration server.
- `app.config.optional` – Whether loading configurations is optional (`true` or `false`).

These variables can be set in two ways:
1. [Directly as environment variables](#setting-environment-variables).
2. [Using a `.env` file](#using-a-env-file).

---

### Setting Environment Variables

Set the environment variables according to your operating system. 

#### Example (Linux/MacOS):

```bash
export app.name="example-service"
export app.config.profiles.active="dev"
export app.config.server.url="https://localhost:8080"
export app.config.optional="true"
```

### Using a `.env` File

Alternatively, you can store environment variables in a `.env` file in the root of your Go application. GoConfig extends [godotenv](https://pkg.go.dev/github.com/joho/godotenv) for loading `.env` files.

#### Example `.env` File:

```go
app.name="example-service"
app.config.profiles.active="dev"
app.config.server.url="https://localhost:8080"
app.config.optional="true"
```

> **Note:** Customize these values based on your application's configuration on the config server.


---

### Loading Configuration in Your Go Application

#### Standard Import

```go
package main

import (
    "log"
    "os"

    config "github.com/oodinga/goconfig"
)

func main() {
    config.Load()
    log.Println("Service port:", os.Getenv("server.port"))
}
```

#### Auto-Loading Configuration

For a simplified approach, import the `autoload` package. This eliminates the need for explicitly calling `config.Load()`:

```go
package main

import (
    "log"
    "os"

    _ "github.com/oodinga/goconfig/autoload"
)

func main() {
    log.Println("Service port:", os.Getenv("server.port"))
}
```
GoConfig will automatically load the configuration from the server and set them as environment variables for your application.

---

## Example Configuration File

When using a Spring Boot configuration server, configurations can be stored in a database or a Git repository. A sample `application.yaml` file might look like this:

#### `application.yaml`:

```yaml
server:
    port: 8080
db:
    username: user
    password: ******
    url: localhost:3306
```
This configuration file can be hosted on your Spring Boot config server. It includes values such as the server port and database connection details.


### Accessing Configuration in Go

Once the config server is set up, you can access the loaded configuration values in your Go application like this:

```go
package main

import (
    "log"
    "os"

    _ "github.com/oodinga/goconfig/autoload"
)

func main() {
    log.Println("Service port:", os.Getenv("server.port"))
    log.Println("Database username:", os.Getenv("db.username"))
    log.Println("Database URL:", os.Getenv("db.url"))
}
```

This code will retrieve and print configuration values that have been loaded from the Spring Boot config server and set as environment variables.

The output will display the values retrieved from the configuration server, such as:

```bash
Service port: 8080
Database username: user
Database URL: localhost:3306
```

This allows your Go application to dynamically load configuration values without the need to manually set them within your code.

---
## Additional Resources

- Refer to the [Spring Cloud Config documentation](https://docs.spring.io/spring-cloud-config/docs/current/reference/html/) to set up a Spring Boot config server.
- Visit the [GoDoc page](https://pkg.go.dev/github.com/oodinga/goconfig@v0.4.3) for API details.
- Learn more about loading and managing environment variables in Go with the [godotenv package](https://pkg.go.dev/github.com/joho/godotenv).
- Explore the [Go Modules documentation](https://golang.org/doc/go1.11#modules) to better understand how Go handles dependencies.

