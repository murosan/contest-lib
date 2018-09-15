# contest-lib


## Testing

```sh
mkdir -p coverage
go test -coverprofile=coverage/coverage.out ./lib/
go tool cover -html=coverage/coverage.out -o coverage/coverage.html
open ./coverage/coverage.html
```
