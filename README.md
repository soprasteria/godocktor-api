# Go API for Docktor model

## Usage

```go
import "github.com/soprasteria/godocktor-api"

dock, err := docktor.Open("localhost"))
if err != nil {
  panic(err)
}
defer dock.Close()

fmt.Printf("Redis exist ? %v\n", dock.Services().IsExist("Redis"))

```
