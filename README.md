# Go API for Docktor model [![Build Status](https://travis-ci.org/soprasteria/godocktor-api.svg?branch=master)](https://travis-ci.org/soprasteria/godocktor-api)

See Docktor : https://github.com/docktor/docktor

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

## License

GNU GENERAL PUBLIC LICENSE 3

See License Files