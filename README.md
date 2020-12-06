## About

Adding default option to os.GetEnv in a smart way

## Usage

```go
package main

import (
  smartenv "github.com/otaviofcs/smartenv"
  "fmt"
)

func main() {
  v := smartenv.Env("RANDOM_ENV", "default one")
  fmt.Println(v)
}
```

Running:

```bash
RANDOM_ENV=opa go run .
```

Running again, now outputting the default env:

```bash
go run .
```
