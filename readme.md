### how to set build info
Add package
```shell
go get github.com/lEx0/go-build-info
```

Simple using in main.go
```go
package main

import (
	"fmt"
	"github.com/lEx0/go-build-info"
)

func main() {
	fmt.Println(build.GetBuildInfo())
}
```

Build main.go
```shell
go build -ldflags "\
  -X github.com/lEx0/go-build-info.version=1.0 \
  -X github.com/lEx0/go-build-info.revision=ddb9c5cf69c5f0cac77d51eb13c8d4e03e605af1 \
  -X github.com/lEx0/go-build-info.branch=master \
  -X github.com/lEx0/go-build-info.buildNumber=2 \
  -X github.com/lEx0/go-build-info.url=http://example.com/build/2 \
  " -v -o main main.go
```

See result
```shell
./main
```


