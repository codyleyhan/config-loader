# Config Loader

Simple golang package to load a json config file.

## Example usage

For the following file structure and environment variable ENV=dev


```
config
    config.dev.json
    config.test.json
    config.prod.json

app.go
```

config.dev.json
```json
{
    "secret": "super secret",
    "port": 8080,
    "hasBool": true,
    "pie": 3.14,
    "nested": {
        "hello": "world"
    }
}
```

app.go
```go
package main

import "github.com/codyleyhan/config-loader"

func main() {
    conf := config.LoadConfig("./config", "config")

    conf.Get("secret") // "super secret"
    conf.Get("nested.hello") // "world"
    conf.Get("nothing") // ""

    conf.GetInt("port") // 8080
    conf.GetInt("nothing") // 0

    conf.GetUint("port") // 8080
    conf.GetUint("nothing") // 0

    conf.GetBool("hasBool") // true
    conf.GetBool("nothing") // false

    conf.GetFloat("pie") // 3.14
    conf.GetFloat("nothing") // undefined behavior
}

```