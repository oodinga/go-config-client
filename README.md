[![Go Reference](https://pkg.go.dev/badge/github.com/oodinga/goconfig@v0.4.3.svg)](https://pkg.go.dev/github.com/oodinga/goconfig@v0.4.3)

# GoConfig
GoConfig is a go package that allows go developers to use spring-boot like config client to load configs exposed via config servers. 

## Installation
``` shell
go get github.com/oodinga/goconfig
```

## Usage
The following env variables are required.
```shell
app.name
app.config.profiles.active
app.config.server.url
app.config.optional
```

These can be set in two ways.

1. [Set environment variables where you are running your application](#setting-env-variables)
2. [Add .env file in the root of your go service](#using-env-file)

### Setting env variables
Follow instructions for setting ENVIROMENT_VARIABLES for your OS.

Example
For linux/MacOs
```shell
export app.name="example-service"
export app.config.profiles.active="dev"
export app.config.server.url="https://localhost:8080"
export app.config.optional="true"
```

### Using .env file
Config variables can also be loaded from .env file. This is an extension of [godotenv](https://pkg.go.dev/github.com/joho/godotenv@v1.5.1)

In your .env file set the following values.

```go
app.name="example-service"
app.config.profiles.active="dev"
app.config.server.url="https://localhost:8080"
app.config.optional="true"
```
> NOTE: Set values according to your application setting on your config server.

Once these variable are set, all you need to do is to import **goconfig** as shown.

```go
package main

import (
    "log"
    "os"

    config "github.com/oodinga/goconfig"
)

func main(){
    config.Load()
    log.Print("Service port: ", os.GetEnv("server.port"))
}
```

This can be simplified by importing the autoload package as shown below.

```go
import (
    "log"
    "os"

    _ "github.com/oodinga/goconfig/autoload"
)

func main(){
    log.Print("Service port: ", os.GetEnv("server.port"))
}
```

GoConfig will autoload the configs from the set config server and set them as environment variables to be used by your application.

### Example config file
When using sprin-boot config server, You can use a db or git as the source of the configuration. A sample yaml file can be set as below.

application.yaml
```yaml
server:
    port: 8080
db:
    username: name
    password: ******
    url: localhost:3600
```

To use these configs 
Using this in your go app.

```go
package main

import (
    "log"
    "os"

    _ "github.com/oodinga/goconfig/autoload"
)

func main(){
    log.Print("Service port: ", os.GetEnv("server.port"))
    log.Print("Database name: ", os.GetEnv("db.name"))
    log.Print("Database url: ", os.GetEnv("db.url"))
}
```

This should log your values to the terminal

Refer to [spring documentation](https://docs.spring.io/spring-cloud-config/docs/current/reference/html/) on how to set a spring config server.

