# Go API for Docktor model

## Usage

```go
import "gitlab.cdk.corp.sopra/cdk/godocktor-api"

dock, err := docktor.Open("localhost"))
if err != nil {
  panic(err)
}
defer dock.Close()

fmt.Printf("Intools exist ? %v\n", dock.Services.IsExist("Intools"))

```
